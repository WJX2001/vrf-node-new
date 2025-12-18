package node

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/pkg/errors"

	"github.com/WJX2001/vrf-node-new/bindings/vrf"
	"github.com/WJX2001/vrf-node-new/txmgr"
)

var (
	errMaxPriorityFeePerGasNotFound = errors.New(
		"Method eth_maxPriorityFeePerGas not found",
	)
	FallbackGasTipCap = big.NewInt(1500000000)
)

type CallerConfig struct {
	ChainClient               *ethclient.Client
	ChainId                   *big.Int
	DappLinkVrfManagerAddress common.Address
	CallerAddress             common.Address
	PrivateKey                *ecdsa.PrivateKey
	NumConfirmations          uint64
	SafeAbortNonceToLowCount  uint64
}

type Caller struct {
	Ctx                     context.Context
	Cfg                     *CallerConfig
	DappLinkVrfContracts    *vrf.DappLinkVRFManager
	RawDappLinkVrfContracts *bind.BoundContract
	DappLinkVrfContractsAbi *abi.ABI
	TxMgr                   txmgr.TxManager
}

func NewCaller(Ctx context.Context, Cfg *CallerConfig) (*Caller, error) {
	dappLinkVrfContracts, err := vrf.NewDappLinkVRFManager(Cfg.DappLinkVrfManagerAddress, Cfg.ChainClient)
	if err != nil {
		log.Error("New DappLink Vrf Manager Fail", "err", err)
		return nil, err
	}

	parsed, err := abi.JSON(strings.NewReader(vrf.DappLinkVRFManagerMetaData.ABI))
	if err != nil {
		log.Error("abi parsed fail", "err", err)
		return nil, err
	}

	dappLinkVrfContractsAbi, err := vrf.DappLinkVRFManagerMetaData.GetAbi()
	if err != nil {
		log.Error("get abi fail", "err", err)
		return nil, err
	}

	rawDappLinkVrfContracts := bind.NewBoundContract(Cfg.DappLinkVrfManagerAddress, parsed, Cfg.ChainClient, Cfg.ChainClient, Cfg.ChainClient)

	txManagerConfig := txmgr.Config{
		ResubmissionTimeout:       time.Second * 5,
		ReceiptQueryInterval:      time.Second,
		NumConfirmations:          Cfg.NumConfirmations,
		SafeAbortNonceTooLowCount: Cfg.SafeAbortNonceToLowCount,
	}

	txManager := txmgr.NewSimpleTxManager(txManagerConfig, Cfg.ChainClient)

	return &Caller{
		Ctx:                     Ctx,
		Cfg:                     Cfg,
		DappLinkVrfContracts:    dappLinkVrfContracts,
		RawDappLinkVrfContracts: rawDappLinkVrfContracts,
		DappLinkVrfContractsAbi: dappLinkVrfContractsAbi,
		TxMgr:                   txManager,
	}, nil
}

// 判断错误是否为 不支持 eth_maxPriorityFeePerGas 方法的错误
// 检查错误信息中是否包含 Method eth_maxPriorityFeePerGas not found
func (caller *Caller) isMaxPriorityFeePerGasNotFoundError(err error) bool {
	return strings.Contains(err.Error(), errMaxPriorityFeePerGasNotFound.Error())
}

func (caller *Caller) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return caller.Cfg.ChainClient.SendTransaction(ctx, tx)
}

func (caller *Caller) UpdateGasPrice(ctx context.Context, tx *types.Transaction) (*types.Transaction, error) {
	var opts *bind.TransactOpts
	var err error

	/**
	* 1.从私钥创建交易选项
	* 2.自动设置 From 地址 （从私钥推导）
	* 3.配置签名器，使用指定的 ChainId 进行 EIP-155 签名
	* 4.返回的 opts 可用于调用合约方法
	 */
	opts, err = bind.NewKeyedTransactorWithChainID(caller.Cfg.PrivateKey, caller.Cfg.ChainId)
	if err != nil {
		log.Error("new keyed transactor with chain id fail", "err", err)
		return nil, err
	}

	opts.Context = ctx
	opts.Nonce = new(big.Int).SetUint64(tx.Nonce())
	// 这个参数控制是否自动发送交易，设置为true时，只构建和签名交易，不发送到链上，返回已签名的交易对象
	opts.NoSend = true

	/**
	重新构建交易
	 1. 使用原交易的 Data() 调用数据
	 2. 从链上获取最新的 gas 价格
	 3. 使用新的 gas 价格构建交易
	 4. 使用私钥签名交易
	 5. 返回已签名但未发送的交易
	*/
	finalTx, err := caller.RawDappLinkVrfContracts.RawTransact(opts, tx.Data())
	switch {
	case err == nil:
		return finalTx, nil
	case caller.isMaxPriorityFeePerGasNotFoundError(err):
		log.Info("Don't support priority fee")
		opts.GasTipCap = FallbackGasTipCap
		return caller.RawDappLinkVrfContracts.RawTransact(opts, tx.Data())
	default:
		return nil, err
	}
}

func (caller *Caller) fulfillRandomWords(ctx context.Context, requestId *big.Int, randomList []*big.Int) (*types.Transaction, error) {
	nonce, err := caller.Cfg.ChainClient.NonceAt(ctx, caller.Cfg.CallerAddress, nil)
	if err != nil {
		log.Error("get eth nonce fail", "err", err)
		return nil, err
	}

	opts, err := bind.NewKeyedTransactorWithChainID(caller.Cfg.PrivateKey, caller.Cfg.ChainId)
	if err != nil {
		log.Error("new keyed transactor with chain id fail", "err", err)
		return nil, err
	}
	opts.Context = ctx
	opts.Nonce = new(big.Int).SetUint64(nonce)
	opts.NoSend = true

	var msgHash [32]byte
	var blsParam vrf.IBLSApkRegistryVrfNoSignerAndSignature

	tx, err := caller.DappLinkVrfContracts.FulfillRandomWords(opts, requestId, randomList, msgHash, big.NewInt(100), blsParam)
	if err != nil {
		log.Error("fulfill random words fail", "err", err)
		return nil, err
	}
	// 错误处理 兼容旧链
	/**
	* 使用场景：
	1. 交易待发送但 gas 价格已过时
	2. 需要重新估算 gas 价格
	3. 兼容不支持 EIP-1599 的旧链 （使用回退值）
	*/
	switch {
	case err == nil:
		return tx, nil

	case caller.isMaxPriorityFeePerGasNotFoundError(err):
		log.Info("Don't support priority fee")
		opts.GasTipCap = FallbackGasTipCap
		return caller.DappLinkVrfContracts.FulfillRandomWords(opts, requestId, randomList, msgHash, big.NewInt(100), blsParam)
	default:
		return nil, err
	}
}

func (caller *Caller) FulfillRandomWords(requestId *big.Int, randomList []*big.Int) (*types.Receipt, error) {
	tx, err := caller.fulfillRandomWords(caller.Ctx, requestId, randomList)
	if err != nil {
		log.Error("build request random words tx fail", "err", err)
		return nil, err
	}
	updateGasPrice := func(ctx context.Context) (*types.Transaction, error) {
		return caller.UpdateGasPrice(ctx, tx)
	}
	receipt, err := caller.TxMgr.Send(caller.Ctx, updateGasPrice, caller.SendTransaction)
	if err != nil {
		log.Error("send tx fail", "err", err)
		return nil, err
	}
	return receipt, nil
}
