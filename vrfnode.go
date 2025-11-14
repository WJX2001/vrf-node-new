package vrf_node

import (
	"context"
	"sync/atomic"

	"github.com/WJX2001/vrf-node-new/config"
	"github.com/WJX2001/vrf-node-new/database"
	"github.com/WJX2001/vrf-node-new/event"
	"github.com/WJX2001/vrf-node-new/synchronizer"
	"github.com/WJX2001/vrf-node-new/synchronizer/node"
	"github.com/ethereum/go-ethereum/log"
)

type VrfNode struct {
	db           *database.DB
	synchronizer *synchronizer.Synchronizer
	eventsParser *event.EventsParser
	stopped      atomic.Bool
}

func NewVrfNode(ctx context.Context, cfg *config.Config, shutdown context.CancelCauseFunc) (*VrfNode, error) {
	ethClient, err := node.DialEthClient(ctx, cfg.Chain.ChainRpcUrl)
	if err != nil {
		log.Error("new eth syncer client fail", "err", err)
		return nil, err
	}

	if err != nil {
		log.Error("new eth syncer client fail", "err", err)
		return nil, err
	}
	db, err := database.NewDB(ctx, cfg.MasterDB)

	if err != nil {
		log.Error("new database fail", "err", err)
		return nil, err
	}

	syncer, err := synchronizer.NewSynchronizer(cfg, db, ethClient, shutdown)
	if err != nil {
		log.Error("new synchronizer fail", "err", err)
		return nil, err
	}
	epConfig := &event.EventsParserConfig{
		DappLinkVrfAddress:        cfg.Chain.DappLinkVrfContractAddress,
		DappLinkVrfFactoryAddress: cfg.Chain.DappLinkVrfFactoryContractAddress,
		EventLoopInterval:         cfg.Chain.EventInterval,
		StartHeight:               nil, // TODO: 根据需要设置
		BlockSize:                 0,   // TODO: 根据需要设置
	}

	eventsParser, err := event.NewEventsParser(db, epConfig, shutdown)
	if err != nil {
		log.Error("new events parser fail", "err", err)
		return nil, err
	}
	return &VrfNode{
		db:           db,
		synchronizer: syncer,
		eventsParser: eventsParser,
	}, nil
}

func (vn *VrfNode) Start(ctx context.Context) error {
	err := vn.synchronizer.Start()
	if err != nil {
		return err
	}
	// err = vn.eventsParser.Start()
	// if err != nil {
	// 	return err
	// }
	return err
}

func (vn *VrfNode) Stop(ctx context.Context) error {
	err := vn.eventsParser.Close()
	if err != nil {
		return err
	}
	return nil
}

func (vn *VrfNode) Stopped() bool {
	return vn.stopped.Load()
}
