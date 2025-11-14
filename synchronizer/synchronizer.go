package synchronizer

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/WJX2001/vrf-node-new/common/tasks"
	"github.com/WJX2001/vrf-node-new/config"
	"github.com/WJX2001/vrf-node-new/database"
	"github.com/WJX2001/vrf-node-new/synchronizer/node"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

type Synchronizer struct {
	ethClient       node.EthClient
	db              *database.DB
	headerTraversal *node.HeaderTraversal
	headers         []types.Header
	latestHeader    *types.Header
	chainCfg        *config.ChainConfig
	resourceCtx     context.Context
	resourceCancel  context.CancelFunc
	tasks           tasks.Group
}

func NewSynchronizer(cfg *config.Config, db *database.DB, client node.EthClient, shutdown context.CancelCauseFunc) (*Synchronizer, error) {
	latestHeader, err := db.Blocks.LatestBlockHeader()
	if err != nil {
		log.Error("query latest block header fail", "err", err)
		return nil, err
	}

	var fromHeader *types.Header
	if latestHeader != nil { // 数据库中已经有同步的区块，从最新区块头继续同步
		fromHeader = latestHeader.RLPHeader.Header()
	} else if cfg.Chain.StartingHeight > 0 { // 数据库为空，但配置了起始高度，从链上获取该高度的区块头作为起始点
		header, err := client.BlockHeaderByNumber(big.NewInt(int64(cfg.Chain.StartingHeight)))
		if err != nil {
			log.Error("get block from chain fail", "err", err)
			return nil, err
		}
		fromHeader = header
	} else { // 数据库为空未配置起始高度
		log.Info("no eth block indexed state")
	}

	headerTraversal := node.NewHeaderTraversal(client, fromHeader, big.NewInt(0), cfg.Chain.ChainId)

	resCtx, resCancel := context.WithCancel(context.Background())

	return &Synchronizer{
		ethClient:       client,
		db:              db,
		headerTraversal: headerTraversal,
		latestHeader:    fromHeader,
		chainCfg:        &cfg.Chain,
		resourceCtx:     resCtx,
		resourceCancel:  resCancel,
		tasks: tasks.Group{
			HandleCrit: func(err error) {
				shutdown(fmt.Errorf("critical error in Synchronizer %w", err))
			},
		},
	}, err
}

func (syncer *Synchronizer) Start() error {
	tickerSyncer := time.NewTicker(syncer.chainCfg.MainLoopInterval)
	syncer.tasks.Go(func() error {
		for range tickerSyncer.C {
			fmt.Println("aaaaaa")
		}
		return nil
	})
	return nil
}

func (syncer *Synchronizer) Close() error {
	syncer.resourceCancel()
	return nil
}
