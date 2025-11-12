package vrf_node

import (
	"context"
	"sync/atomic"

	"github.com/WJX2001/vrf-node-new/config"
	"github.com/WJX2001/vrf-node-new/database"
	"github.com/WJX2001/vrf-node-new/event"
	"github.com/ethereum/go-ethereum/log"
)

type VrfNode struct {
	eventsParser *event.EventsParser
	stopped      atomic.Bool
}

func NewVrfNode(ctx context.Context, cfg *config.Config, shutdown context.CancelCauseFunc) (*VrfNode, error) {

	db, err := database.NewDB(ctx, cfg.MasterDB)

	if err != nil {
		log.Error("new database fail", "err", err)
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
		eventsParser: eventsParser,
	}, nil
}

func (vn *VrfNode) Start(ctx context.Context) error {
	err := vn.eventsParser.Start()
	if err != nil {
		return err
	}
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
