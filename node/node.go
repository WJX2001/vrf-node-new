package node

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	"github.com/consensys/gnark-crypto/ecc/bn254"
	tdtypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"

	"github.com/WJX2001/vrf-node-new/bindings/bls"
	"github.com/WJX2001/vrf-node-new/client"
	"github.com/WJX2001/vrf-node-new/config"
	"github.com/WJX2001/vrf-node-new/database"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/the-web3-contracts/vrf-node/sign"

	wsClient "github.com/WJX2001/vrf-node-new/ws/client"
)

type Node struct {
	wg   sync.WaitGroup
	done chan struct{}

	db *database.DB

	privateKey *ecdsa.PrivateKey
	from       common.Address
	ctx        context.Context
	cancel     context.CancelFunc
	stopChan   chan struct{}
	stopped    atomic.Bool

	wsClient *wsClient.WSClients
	KeyPairs *sign.KeyPair

	signTimeOut      time.Duration
	waitScanInterval time.Duration
	signRequestChan  chan tdtypes.RPCRequest
}

func NewNode(ctx context.Context, db *database.DB, privKey *ecdsa.PrivateKey, keyPairs *sign.KeyPair, shouldRegister bool, cfg *config.Config, shutdown context.CancelCauseFunc) (*Node, error) {
	from := crypto.PubkeyToAddress(privKey.PublicKey)
	pubKey := crypto.CompressPubkey(&privKey.PublicKey)
	pubkeyHex := hex.EncodeToString(pubKey)

	log.Info("public key", "publicKey", pubkeyHex)

	if shouldRegister {
		// 执行注册的逻辑
		log.Info("register to operator")
		txKey, txOpt, err := registerOperator(ctx, cfg, privKey, keyPairs)
		if err != nil {
			log.Error("register operator fail", "err", err)
			return nil, err
		}
		log.Info("register success", "txkey", txKey.Hash().String(), "txOpt", txOpt.Hash().String())
	}

	wsCli, err := wsClient.NewWSClient(cfg.Node.WsAddr, "/ws", privKey, pubkeyHex)
	if err != nil {
		log.Error("new ws client fail", "err", err)
		return nil, err
	}

	return &Node{
		wg:   sync.WaitGroup{},
		done: make(chan struct{}),
		db:   db,

		privateKey: privKey,
		from:       from,
		ctx:        ctx,

		wsClient: wsCli,
		KeyPairs: keyPairs,

		signTimeOut:      cfg.Node.SignTimeOut,
		waitScanInterval: cfg.Node.WaitScanInterval,
		signRequestChan:  make(chan tdtypes.RPCRequest, 100),
	}, nil
}

func (n *Node) Start(ctx context.Context) error {
	n.wg.Add(1)
	go n.sign()
	return nil
}

func (n *Node) Stop(ctx context.Context) error {
	n.cancel()
	close(n.done)
	n.wg.Wait()
	n.stopped.Store(true)
	return nil
}

func (n *Node) Stopped() bool {
	return n.stopped.Load()
}

func (n *Node) sign() {
	defer n.wg.Done()
	log.Info("start to sign message")
	go func() {
		defer func() {
			log.Info("exit sign process")
		}()
		for {
			select {
			case <-n.stopChan:
				return
			case req := <-n.signRequestChan:
				log.Error("req", "req", req)
				return
			}
		}
	}()
}

func (n *Node) SignMessage(messageHash string) (*sign.Signature, error) {
	var bSign *sign.Signature
	log.Info("before sign", "messageHash", messageHash)
	bSign = n.KeyPairs.SignMessage(crypto.Keccak256Hash(common.Hex2Bytes(messageHash)))
	log.Info("after sign", "bSign", bSign)
	return bSign, nil
}

/*
1. 注册 BLS 公钥：使用 BLS 签名、G1/G2 公钥注册到合约
2. 注册操作员：将节点地址注册为操作员
*/
func registerOperator(ctx context.Context, cfg *config.Config, privKey *ecdsa.PrivateKey, keyPairs *sign.KeyPair) (*types.Transaction, *types.Transaction, error) {
	// TODO: 第一阶段 初始化以太坊客户端和合约
	// 连接以太坊客户端，创建带超时的以太坊客户端连接
	ethCli, err := client.DialEthClientWithTimeout(ctx, cfg.Chain.ChainRpcUrl, false)
	if err != nil {
		log.Error("new eth client fail", "err", err)
		return nil, nil, err
	}
	// 创建 BLS 注册表合约实例，使用配置中的 BLS 注册表地址创建合约绑定
	blsRegistry, err := bls.NewBLSApkRegistry(common.HexToAddress(cfg.Chain.BlsRegistryAddress), ethCli)
	if err != nil {
		log.Error("new bls registry fail", "err", err)
		return nil, nil, err
	}
	// 获取合约 ABI，获取 ABI，用于后续构建原始交易
	blsParsed, err := bls.BLSApkRegistryMetaData.GetAbi()
	if err != nil {
		log.Error("fetch bls registry abi fail", "err", err)
		return nil, nil, err
	}

	// 创建原始合约绑定，创建用于发送原始交易的绑定合约
	// 创建用于发送原始交易的绑定合约
	rawBlsContract := bind.NewBoundContract(common.HexToAddress(cfg.Chain.BlsRegistryAddress), *blsParsed, ethCli, ethCli, ethCli)

	// TODO: 第二阶段 准备交易参数
	// 创建交易选项
	topts, err := client.NewTransactOpts(ctx, uint64(cfg.Chain.ChainId), privKey)
	if err != nil {
		log.Error("new transactopt fail", "err", err)
	}

	// 获取节点地址
	// 从私钥公钥计算以太坊地址
	nodeAddr := crypto.PubkeyToAddress(privKey.PublicKey)
	// 获取最新区块号
	latestBlock, _ := ethCli.BlockNumber(ctx)

	// 用于只读调用，指定区块号和调用者地址
	cOpts := &bind.CallOpts{
		BlockNumber: big.NewInt(int64(latestBlock)),
		From:        nodeAddr,
	}

	// 获取公钥注册消息哈希
	msg, err := blsRegistry.GetPubkeyRegMessageHash(cOpts, nodeAddr)
	if err != nil {
		log.Error("get public key register message hash fail", "err", err)
		return nil, nil, err
	}

	// TODO: 第三部分：生成 BLS 签名和参数
	// 计算 BLS 签名
	// 使用 BLS 私钥对消息哈希进行签名（BN254曲线上的标量乘法）
	sigMsg := new(bn254.G1Affine).ScalarMultiplication(sign.NewG1Point(msg.X, msg.Y).G1Affine, keyPairs.PrivKey.BigInt(new(big.Int)))

	// 构建 BLS 注册参数
	// 组装注册参数：BLS签名（G1）、公钥G1、公钥G2
	blsParam := bls.IBLSApkRegistryPubkeyRegistrationParams{
		PubkeyRegistrationSignature: bls.BN254G1Point{
			X: sigMsg.X.BigInt(new(big.Int)),
			Y: sigMsg.Y.BigInt(new(big.Int)),
		},
		PubkeyG1: bls.BN254G1Point{
			X: keyPairs.GetPubKeyG1().X.BigInt(new(big.Int)),
			Y: keyPairs.GetPubKeyG1().Y.BigInt(new(big.Int)),
		},
		PubkeyG2: bls.BN254G2Point{
			X: [2]*big.Int{keyPairs.GetPubKeyG2().X.A1.BigInt(new(big.Int)), keyPairs.GetPubKeyG2().X.A0.BigInt(new(big.Int))},
			Y: [2]*big.Int{keyPairs.GetPubKeyG2().Y.A1.BigInt(new(big.Int)), keyPairs.GetPubKeyG2().Y.A0.BigInt(new(big.Int))},
		},
	}

	// TODO: 第四部分：注册 BLS 公钥
	// 构建注册 BLS 公钥的交易
	regBlsKey, err := blsRegistry.RegisterBLSPublicKey(topts, nodeAddr, blsParam, msg)
	if err != nil {
		log.Error("register bls public key fail", "err", err)
		return nil, nil, err
	}

	// 创建原始交易，使用原始合约绑定创建已签名的交易对象
	blsRegTx, err := rawBlsContract.RawTransact(topts, regBlsKey.Data())
	if err != nil {
		log.Error("raw bls tx fail", "err", err)
		return nil, nil, err
	}
	// 将交易发送到链上
	err = ethCli.SendTransaction(ctx, blsRegTx)
	if err != nil {
		log.Error("send raw transaction fail", "err", err)
		return nil, nil, err
	}

	// 等待交易确认 等待并获取交易回执，确认上链
	_, err = client.GetTransactionReceipt(ctx, ethCli, blsRegTx.Hash())
	if err != nil {
		log.Error("get transaction receipt fail", "err", err)
		return nil, nil, err
	}

	// TODO: 第五步：注册操作员

	// 构建注册操作员的交易
	// 构建 registerOperator 交易 需要先完成 BLS 公钥注册
	regOperator, err := blsRegistry.RegisterOperator(topts, nodeAddr)
	if err != nil {
		log.Error("register operator fail", "err", err)
		return nil, nil, err
	}

	// 创建原始交易
	// 创建已签名的操作员注册交易
	blsRegOptTx, err := rawBlsContract.RawTransact(topts, regOperator.Data())
	if err != nil {
		log.Error("raw bls tx fail", "err", err)
		return nil, nil, err
	}

	// 发送交易，发送操作员注册交易
	err = ethCli.SendTransaction(ctx, blsRegOptTx)
	if err != nil {
		log.Error("send raw transaction fail", "err", err)
		return nil, nil, err
	}

	// 等待交易确认
	_, err = client.GetTransactionReceipt(ctx, ethCli, blsRegOptTx.Hash())
	if err != nil {
		log.Error("get transaction receipt fail", "err", err)
	}
	return blsRegTx, blsRegOptTx, nil
}
