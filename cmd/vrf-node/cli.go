package main

import (
	"context"

	vrf_node "github.com/WJX2001/vrf-node-new"
	"github.com/WJX2001/vrf-node-new/common/cliapp"
	"github.com/WJX2001/vrf-node-new/common/opio"
	"github.com/WJX2001/vrf-node-new/config"
	"github.com/WJX2001/vrf-node-new/database"
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

func runMigrations(ctx *cli.Context) error {
	log.Info("Running migrations...")
	cfg, err := config.LoadConfig(ctx)
	if err != nil {
		log.Error("failed to load config", "err", err)
		return err
	}
	ctx.Context = opio.CancelOnInterrupt(ctx.Context)
	db, err := database.NewDB(ctx.Context, cfg.MasterDB)
	if err != nil {
		log.Error("failed to connect to database", "err", err)
		return err
	}
	defer func(db *database.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)
	return db.ExecuteSQLMigration(cfg.Migrations)
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
			{
				Name:        "migrate",
				Flags:       flags,
				Description: "Runs the database migrations",
				Action:      runMigrations,
			},
			{
				Name:        "version",
				Description: "print version",
				Action: func(ctx *cli.Context) error {
					cli.ShowVersion(ctx)
					return nil
				},
			},
		},
	}
}
