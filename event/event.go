package event

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/WJX2001/vrf-node-new/common/tasks"
	"github.com/WJX2001/vrf-node-new/database"
	"github.com/ethereum/go-ethereum/log"
)

type EventsParserConfig struct {
	DappLinkVrfAddress        string
	DappLinkVrfFactoryAddress string
	EventLoopInterval         time.Duration
	StartHeight               *big.Int
	BlockSize                 uint64 // 每次扫的块
}

type EventsParser struct {
	db             *database.DB
	epConf         *EventsParserConfig
	resourceCtx    context.Context
	resourceCancel context.CancelFunc
	tasks          tasks.Group
}

func NewEventsParser(db *database.DB, epConf *EventsParserConfig, shutdown context.CancelCauseFunc) (*EventsParser, error) {

	resCtx, resCancel := context.WithCancel(context.Background())

	return &EventsParser{
		db:             db,
		epConf:         epConf,
		resourceCtx:    resCtx,
		resourceCancel: resCancel,
		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in event parser: %w", err))
		}},
	}, nil
}

func (ep *EventsParser) Start() error {
	tickerSyncer := time.NewTicker(ep.epConf.EventLoopInterval)
	ep.tasks.Go(func() error {
		for range tickerSyncer.C {
			log.Info("xxxxxx")
		}
		return nil
	})
	return nil
}

func (ep *EventsParser) Close() error {
	ep.resourceCancel()
	return ep.tasks.Wait()
}
