package main

import (
	"context"

	vrf_node "github.com/WJX2001/vrf-node-new"
	"github.com/WJX2001/vrf-node-new/common/cliapp"
	"github.com/WJX2001/vrf-node-new/config"
	flag2 "github.com/WJX2001/vrf-node-new/flags"
	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli/v2"
)

func runDappLinkVrfNode(ctx *cli.Context, shutdown context.CancelCauseFunc) (cliapp.Lifecycle, error) {
	cfg, err := config.LoadConfig(ctx)
	if err != nil {
		log.Error("failed to load config", "err", err)
		return nil, err
	}
	return vrf_node.NewVrfNode(ctx.Context, &cfg, shutdown)
}

func NewCli(GitCommit string, GitDate string) *cli.App {
	flags := flag2.Flags
	return &cli.App{
		Version:              "v0.0.1",
		Description:          "An indexer of all optimism events with a serving api layer",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:        "index",
				Flags:       flags,
				Description: "Runs the indexing service",
				Action:      cliapp.LifecycleCmd(runDappLinkVrfNode),
			},
		},
	}
}
