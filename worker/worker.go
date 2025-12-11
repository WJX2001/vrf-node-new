package worker

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	"github.com/WJX2001/vrf-node-new/common/tasks"
	"github.com/WJX2001/vrf-node-new/database"
	"github.com/WJX2001/vrf-node-new/node"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
)

type WorkerConfig struct {
	ChainClient               *ethclient.Client
	ChainId                   *big.Int
	DappLinkVrfManagerAddress common.Address
	CallerAddress             common.Address
	PrivateKey                *ecdsa.PrivateKey
	NumConfirmations          uint64
	SafeAbortNonceToLowCount  uint64
	LoopInternal              time.Duration
}

type Worker struct {
	workerConf *WorkerConfig
	db         *database.DB
	caller     *node.Caller

	resourceCtx    context.Context
	resourceCancel context.CancelFunc
	tasks          tasks.Group
}

func NewWorker(db *database.DB, workconf *WorkerConfig, shutdown context.CancelCauseFunc) (*Worker, error) {

	callerConf := &node.CallerConfig{
		ChainClient:               workconf.ChainClient,
		ChainId:                   workconf.ChainId,
		DappLinkVrfManagerAddress: workconf.DappLinkVrfManagerAddress,
		CallerAddress:             workconf.CallerAddress,
		PrivateKey:                workconf.PrivateKey,
		NumConfirmations:          workconf.NumConfirmations,
		SafeAbortNonceToLowCount:  workconf.SafeAbortNonceToLowCount,
	}

	callerF, err := node.NewCaller(context.Background(), callerConf)
	if err != nil {
		log.Error("new caller fail", "err", err)
		return nil, err
	}

	resCtx, resCancel := context.WithCancel(context.Background())

	return &Worker{
		db:             db,
		workerConf:     workconf,
		caller:         callerF,
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

			var randomList []*big.Int
			randomList = append(randomList, big.NewInt(1000))
			randomList = append(randomList, big.NewInt(1001))
			randomList = append(randomList, big.NewInt(1002))

			err := wk.ProcessCallerVrf(randomList)
			if err != nil {
				log.Error("Process caller vrf manager fail", "err", err)
				return err
			}
		}
		return nil
	})
	return nil
}

func (wk *Worker) Close() error {
	wk.resourceCancel()
	return wk.tasks.Wait()
}

func (wk *Worker) ProcessCallerVrf(randomList []*big.Int) error {
	// 获取 RequestSend 的合约事件
	requestUnSentList, err := wk.db.RequestSend.QueryUnHandleRequestSendList()
	if err != nil {
		log.Error("query unhandle request send list fail", "err", err)
		return err
	}
	log.Info("handle requestUnSentList", "length", len(requestUnSentList))
	for _, requestUnSent := range requestUnSentList {
		txReceipt, err := wk.caller.FulfillRandomWords(requestUnSent.RequestId, randomList) // 组装随机数发送到链上
		if err != nil {
			log.Error("FulfillRandomWords fail", "err", err)
			return err
		}

		if txReceipt.Status == 1 {
			err := wk.db.RequestSend.MarkRequestSendFinish(requestUnSent) // 更新 RequestSend 合约事件的状态
			if err != nil {
				log.Error("mark request sent event list fail", "err", err)
				return err
			}
		}
	}
	return nil
}
