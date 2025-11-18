package event

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/WJX2001/vrf-node-new/common/bigint"
	"github.com/WJX2001/vrf-node-new/common/tasks"
	"github.com/WJX2001/vrf-node-new/database"
	"github.com/WJX2001/vrf-node-new/database/common"
	"github.com/WJX2001/vrf-node-new/database/worker"
	"github.com/WJX2001/vrf-node-new/event/contracts"
	"github.com/WJX2001/vrf-node-new/synchronizer/retry"
	"github.com/ethereum/go-ethereum/log"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventsParserConfig struct {
	DappLinkVrfAddress        string
	DappLinkVrfFactoryAddress string
	EventLoopInterval         time.Duration
	StartHeight               *big.Int
	BlockSize                 uint64
}

type EventsParser struct {
	db *database.DB

	dappLinkVrf        *contracts.DappLinkVrfManager
	dappLinkVrfFactory *contracts.DappLinkVrfFactory

	epConf            *EventsParserConfig
	latestBlockHeader *common.BlockHeader

	resourceCtx    context.Context
	resourceCancel context.CancelFunc
	tasks          tasks.Group
}

func NewEventsParser(db *database.DB, epConf *EventsParserConfig, shutdown context.CancelCauseFunc) (*EventsParser, error) {
	dappLinkVrf, err := contracts.NewDappLinkVrfManager()
	if err != nil {
		log.Error("new dapplink vrf fail", "err", err)
		return nil, err
	}
	dappLinkVrfFactory, err := contracts.NewDappLinkVrfFactory()
	if err != nil {
		log.Error("new dapplink vrf factory fail", "err", err)
		return nil, err
	}

	ltBlockHeader, err := db.EventBlocks.LatestEventBlockHeader()
	if err != nil {
		log.Error("fetch latest block header fail", "err", err)
		return nil, err
	}

	resCtx, resCancel := context.WithCancel(context.Background())

	return &EventsParser{
		db:                 db,
		dappLinkVrf:        dappLinkVrf,
		dappLinkVrfFactory: dappLinkVrfFactory,
		epConf:             epConf,
		latestBlockHeader:  ltBlockHeader,
		resourceCtx:        resCtx,
		resourceCancel:     resCancel,
		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in event parser: %w", err))
		}},
	}, nil
}

func (ep *EventsParser) Start() error {
	tickerSyncer := time.NewTicker(ep.epConf.EventLoopInterval)
	ep.tasks.Go(func() error {
		for range tickerSyncer.C {
			log.Info("start parse event logs")
			err := ep.ProcessEvent()
			if err != nil {
				log.Info("process event error", "err", err)
				return err
			}
		}
		return nil
	})
	return nil
}

func (ep *EventsParser) ProcessEvent() error {
	lastBlockNumber := ep.epConf.StartHeight
	if ep.latestBlockHeader != nil {
		lastBlockNumber = ep.latestBlockHeader.Number
	}
	log.Info("process event latest block number", "lastBlockNumber", lastBlockNumber)

	latestHeaderScope := func(db *gorm.DB) *gorm.DB {
		newQuery := db.Session(&gorm.Session{NewDB: true})
		headers := newQuery.Model(common.BlockHeader{}).Where("number > ?", lastBlockNumber)
		return db.Where("number = (?)", newQuery.Table("(?) as block_numbers", headers.Order("number ASC").Limit(int(ep.epConf.BlockSize))).Select("MAX(number)"))
	}

	if latestHeaderScope == nil {
		return nil
	}

	latestBlockHeader, err := ep.db.Blocks.BlockHeaderWithScope(latestHeaderScope)
	if err != nil {
		log.Error("get latest block header with scope fail", "err", err)
		return err
	} else if latestBlockHeader == nil {
		log.Debug("no new block for process event")
		return nil
	}
	fromHeight, toHeight := new(big.Int).Add(lastBlockNumber, bigint.One), latestBlockHeader.Number
	eventBlocks := make([]worker.EventBlocks, 0, toHeight.Uint64()-fromHeight.Uint64())
	for index := fromHeight.Uint64(); index < toHeight.Uint64(); index++ {
		blockHeader, err := ep.db.Blocks.BlockHeaderByNumber(big.NewInt(int64(index)))
		if err != nil {
			return err
		}
		evBlock := worker.EventBlocks{
			GUID:       uuid.New(),
			Hash:       blockHeader.Hash,
			ParentHash: blockHeader.ParentHash,
			Number:     blockHeader.Number,
			Timestamp:  blockHeader.Timestamp,
		}
		eventBlocks = append(eventBlocks, evBlock)
	}

	requestSentList, fillRandomWordList, err := ep.dappLinkVrf.ProcessDappLinkVrfManagerEvent(ep.db, ep.epConf.DappLinkVrfAddress, fromHeight, toHeight)
	if err != nil {
		log.Error("process dapplink vrf event fail", "err", err)
		return err
	}
	proxyCreatedList, err := ep.dappLinkVrfFactory.ProcessDappLinkVrfFactoryEvent(ep.db, ep.epConf.DappLinkVrfFactoryAddress, fromHeight, toHeight)
	if err != nil {
		return err
	}

	retryStrategy := &retry.ExponentialStrategy{Min: 1000, Max: 20_000, MaxJitter: 250}
	if _, err := retry.Do[interface{}](ep.resourceCtx, 10, retryStrategy, func() (interface{}, error) {
		if err := ep.db.Transaction(func(tx *database.DB) error {
			if len(requestSentList) > 0 {
				err := ep.db.RequestSend.StoreRequestSend(requestSentList)
				if err != nil {
					log.Error("store request send fail", "err", err)
					return err
				}
			}

			if len(fillRandomWordList) > 0 {
				err := ep.db.FillRandomWords.StoreFillRandomWords(fillRandomWordList)
				if err != nil {
					log.Error("store fill random words fail", "err", err)
					return err
				}
			}

			if len(proxyCreatedList) > 0 {
				err := ep.db.PoxyCreated.StorePoxyCreated(proxyCreatedList)
				if err != nil {
					log.Error("store proxy created fail", "err", err)
					return err
				}
			}

			if len(eventBlocks) > 0 {
				err := ep.db.EventBlocks.StoreEventBlocks(eventBlocks)
				if err != nil {
					log.Error("store event blocks fail", "err", err)
					return err
				}
			}
			return nil
		}); err != nil {
			log.Debug("unable to persist batch", err)
			return nil, fmt.Errorf("unable to persist batch: %w", err)
		}
		return nil, nil
	}); err != nil {
		return err
	}
	ep.latestBlockHeader = latestBlockHeader
	return nil
}

func (ep *EventsParser) Close() error {
	ep.resourceCancel()
	return ep.tasks.Wait()
}
