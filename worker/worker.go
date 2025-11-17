package worker

import (
	"context"
	"fmt"
	"time"

	"github.com/WJX2001/vrf-node-new/common/tasks"
	"github.com/WJX2001/vrf-node-new/database"
	"github.com/ethereum/go-ethereum/log"
)

type WorkerConfig struct {
	LoopInternal time.Duration
}

type Worker struct {
	workerConf *WorkerConfig
	db         *database.DB

	resourceCtx    context.Context
	resourceCancel context.CancelFunc
	tasks          tasks.Group
}

func NewWorker(db *database.DB, workconf *WorkerConfig, shutdown context.CancelCauseFunc) (*Worker, error) {

	resCtx, resCancel := context.WithCancel(context.Background())

	return &Worker{
		db:             db,
		workerConf:     workconf,
		resourceCtx:    resCtx,
		resourceCancel: resCancel,
		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in worker processor: %w", err))
		}},
	}, nil
}

func (wk *Worker) Start() error {
	tickerEventWorker := time.NewTicker(wk.workerConf.LoopInternal)
	wk.tasks.Go(func() error {
		for range tickerEventWorker.C {
			log.Info("starting worker processor...")
		}
		return nil
	})
	return nil
}

func (wk *Worker) Close() error {
	wk.resourceCancel()
	return wk.tasks.Wait()
}
