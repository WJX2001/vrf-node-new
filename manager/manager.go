package manager

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/WJX2001/vrf-node-new/bindings/bls"
	"github.com/WJX2001/vrf-node-new/bindings/vrf"
	"github.com/WJX2001/vrf-node-new/client"
	"github.com/WJX2001/vrf-node-new/config"
	"github.com/WJX2001/vrf-node-new/database"
	"github.com/WJX2001/vrf-node-new/manager/router"
	"github.com/WJX2001/vrf-node-new/manager/types"
	"github.com/WJX2001/vrf-node-new/sign"
	"github.com/WJX2001/vrf-node-new/ws/server"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/gin-gonic/gin"
)

var (
	errNotEnoughSignNode = errors.New("not enough available nodes to sign")
	errNotEnoughVoteNode = errors.New("not enough available nodes to vote")
)

type Manager struct {
	wg              sync.WaitGroup
	done            chan struct{}
	log             log.Logger
	db              *database.DB
	wsServer        server.IWebsocketManager // WebSocket 管理器接口
	NodeMembers     []string                 // 节点成员列表
	httpAddr        string                   // HTTP 地址
	httpServer      *http.Server             // HTTP 服务器
	sStakeUrl       string                   // 质押URL
	mu              sync.Mutex               // 互斥锁
	ctx             context.Context
	stopped         atomic.Bool             // 原子布尔值，表示是否已经停止
	ethChainID      uint64                  // 链ID
	privateKey      *ecdsa.PrivateKey       // 私钥
	from            common.Address          // 地址
	ethClient       *ethclient.Client       // 以太坊客户端
	vrfContract     *vrf.DappLinkVRFManager // VRF 合约实例
	vrfContractAddr common.Address          // VRF 合约地址
	rawVrfContract  *bind.BoundContract     // 原始绑定合约
	barContract     *bls.BLSApkRegistry     // BLS 公钥注册表合约
	barContractAddr common.Address          // BLS 合约地址
	rawBarContract  *bind.BoundContract     // 原始 BLS 绑定合约
	isFirstBatch    bool                    // 是否首次批次
	signTimeout     time.Duration           // 签名超时（10秒）
	fPTimeout       time.Duration           // 工作循环超时（10秒）
}

func NewManager(ctx context.Context, db *database.DB, wsServer server.IWebsocketManager, cfg *config.Config, priv *ecdsa.PrivateKey, shutdown context.CancelCauseFunc) (*Manager, error) {
	// 连接以太坊客户端
	ethCli, err := client.DialEthClientWithTimeout(ctx, cfg.Chain.ChainRpcUrl, false)
	if err != nil {
		return nil, err
	}
	// 初始化 VRF 合约
	vrfContract, err := vrf.NewDappLinkVRFManager(common.HexToAddress(cfg.Chain.DappLinkVrfContractAddress), ethCli)
	if err != nil {
		return nil, err
	}

	// 将 ABI 字符串 -> 解析成 ABI 对象 -> 用来编码/解码合约调用
	// ABI字符串为接口说明书 转化为可执行对象，将 ABI 字符串解析成 ABI对象，用于后续构造交易
	fParsed, err := abi.JSON(strings.NewReader(
		vrf.DappLinkVRFFactoryMetaData.ABI,
	))
	if err != nil {
		return nil, err
	}
	// 解析 VRF ABI 创建原始绑定合约
	rawVrfContract := bind.NewBoundContract(
		common.HexToAddress(cfg.Chain.DappLinkVrfContractAddress), fParsed, ethCli, ethCli,
		ethCli,
	)

	barContract, err := bls.NewBLSApkRegistry(common.HexToAddress(cfg.Chain.BlsRegistryAddress), ethCli)
	if err != nil {
		return nil, err
	}

	// 解析 BLS ABI
	bParsed, err := bls.BLSApkRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	rawBarContract := bind.NewBoundContract(
		common.HexToAddress(cfg.Chain.BlsRegistryAddress), *bParsed, ethCli, ethCli,
		ethCli,
	)

	nodeMemberS := strings.Split(cfg.Manager.NodeMembers, ",")
	for _, nodeMember := range nodeMemberS {
		fmt.Println("nodeMember===", nodeMember)
	}

	return &Manager{
		done:            make(chan struct{}),
		db:              db,
		wsServer:        wsServer,
		httpAddr:        cfg.Manager.HttpAddr,
		NodeMembers:     nodeMemberS,
		ctx:             ctx,
		privateKey:      priv,
		from:            crypto.PubkeyToAddress(priv.PublicKey),
		signTimeout:     time.Second * 10,
		fPTimeout:       time.Second * 10,
		ethChainID:      uint64(cfg.Chain.ChainId),
		ethClient:       ethCli,
		vrfContract:     vrfContract,
		vrfContractAddr: common.HexToAddress(cfg.Chain.DappLinkVrfContractAddress),
		rawVrfContract:  rawVrfContract,
		barContract:     barContract,
		barContractAddr: common.HexToAddress(cfg.Chain.BlsRegistryAddress),
		rawBarContract:  rawBarContract,
	}, nil
}

func (m *Manager) Start(ctx context.Context) error {
	waitNodeTicker := time.NewTicker(5 * time.Second)
	var done bool
	for !done {
		select {
		case <-waitNodeTicker.C:
			availableNodes := m.availableNodes(m.NodeMembers)
			if len(availableNodes) < len(m.NodeMembers) {
				m.log.Warn("wait node to connect", "availableNodesNum", len(availableNodes), "connectedNodeNum", len(m.NodeMembers))
				continue
			} else {
				done = true
				break
			}
		}
	}

	registry := router.NewRegistry(m, m.db)
	r := gin.Default()
	registry.Register(r)

	var s *http.Server
	s = &http.Server{
		Addr:    m.httpAddr,
		Handler: r,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			m.log.Error("api server starts failed", "err", err)
		}
	}()
	m.httpServer = s

	m.wg.Add(1)
	go m.work()
	log.Info("manager is starting......")
	return nil
}

// work 是一个 定时任务循环，定期执行签名收集和链上提交
func (m *Manager) work() {
	fpTicker := time.NewTicker(m.fPTimeout)
	defer m.wg.Done()

	for {
		select {
		case <-fpTicker.C:
			var signature *sign.G1Point
			var g2Point *sign.G2Point
			var NonSignerPubkeys []vrf.BN254G1Point

			request := types.SignMsgRequest{}

			// 调用 SignMsgBatch 收集多节点签名
			res, err := m.SignMsgBatch(request)
			if errors.Is(err, errNotEnoughSignNode) || errors.Is(err, errNotEnoughVoteNode) {
				m.log.Error("not enough available nodes to sign or not enough available nodes to vote")
				continue
			} else if err != nil {
				m.log.Error("failed to sign msg", "err", err)
				continue
			}
			m.log.Info("success to sign msg", "signature", res.Signature)

			signature = res.Signature
			g2Point = res.G2Point

			// 转换未签名者公钥格式
			for _, v := range res.NonSignerPubkeys {
				NonSignerPubkeys = append(NonSignerPubkeys, vrf.BN254G1Point{
					X: v.X.BigInt(new(big.Int)),
					Y: v.Y.BigInt(new(big.Int)),
				})
			}

			// 创建交易选项 包含 私钥、链ID Gas 等
			opts, err := client.NewTransactOpts(m.ctx, m.ethChainID, m.privateKey)
			if err != nil {
				m.log.Error("failed to new transact opts", "err", err)
				continue
			}
			// 构建合约调用参数
			vrfNonSignerAndSignature := vrf.IBLSApkRegistryVrfNoSignerAndSignature{
				NonSignerPubKeys: NonSignerPubkeys, // 未签名者公钥列表
				// 聚合公钥 （G2点）
				ApkG2: vrf.BN254G2Point{
					X: [2]*big.Int{g2Point.X.A1.BigInt(new(big.Int)), g2Point.X.A0.BigInt(new(big.Int))},
					Y: [2]*big.Int{g2Point.Y.A1.BigInt(new(big.Int)), g2Point.Y.A0.BigInt(new(big.Int))},
				},
				// 聚合签名（G1点）
				Sigma: vrf.BN254G1Point{
					X: signature.X.BigInt(new(big.Int)),
					Y: signature.Y.BigInt(new(big.Int)),
				},
				// 质押总量
				TotalBtcStake:      big.NewInt(100),
				TotalDappLinkStake: big.NewInt(100),
			}
			// 构建交易
			var aaa []*big.Int
			var bb [32]byte

			tx, err := m.vrfContract.FulfillRandomWords(opts, big.NewInt(1), aaa, bb, big.NewInt(1), vrfNonSignerAndSignature)
			if err != nil {
				m.log.Error("failed to craft VerifyFinalitySignature transaction", "err", err)
				continue
			}
			// 创建原始交易
			rTx, err := m.rawVrfContract.RawTransact(opts, tx.Data())
			if err != nil {
				m.log.Error("failed to raw VerifyFinalitySignature transaction", "err", err)
				continue
			}

			// 发送到链上
			err = m.ethClient.SendTransaction(m.ctx, tx)
			if err != nil {
				m.log.Error("failed to send VerifyFinalitySignature transaction", "err", err)
				continue
			}

			// 等待交易确认，获取交易回执 确认上链
			receipt, err := client.GetTransactionReceipt(m.ctx, m.ethClient, rTx.Hash())
			if err != nil {
				m.log.Error("failed to get verify finality transaction receipt", "err", err)
				continue
			}
			m.log.Info("success to send verify finality signature transaction", "tx_hash", receipt.TxHash.String())
		case <-m.done:
			return
		}
	}
}

// 协调多个节点进行 BLS 聚合签名，返回聚合签名结果
func (m *Manager) SignMsgBatch(request types.SignMsgRequest) (*types.SignResult, error) {
	m.log.Info("received sign request")

	activeMember, err := m.db.Member.GetActiveMember()
	if err != nil {
		m.log.Error("failed to get active member from db", "err", err)
		return nil, err
	}
	availableNodes := m.availableNodes(activeMember)
	if len(availableNodes) == 0 {
		m.log.Warn("not enough sign node", "availableNodes", availableNodes)
		return nil, errNotEnoughSignNode
	}

	ctx := types.NewContext().WithAvailableNodes(availableNodes).WithRequestId(randomRequestId())

	var resp types.SignResult
	var signErr error
	resp, signErr = m.sign(ctx, request, types.SignMsgBatch)
	if signErr != nil {
		return nil, signErr
	}
	if resp.Signature == nil {
		return nil, errNotEnoughVoteNode
	}
	return &resp, nil

}

func (m *Manager) availableNodes(nodeMembers []string) []string {
	aliveNodes := m.wsServer.AliveNodes()
	log.Info("check available nodes", "expected", fmt.Sprintf("%v", nodeMembers), "alive nodes", fmt.Sprintf("%v", aliveNodes))
	availableNodes := make([]string, 0)
	for _, n := range aliveNodes {
		if ExistsIgnoreCase(nodeMembers, n) {
			availableNodes = append(availableNodes, n)
		}
	}
	return availableNodes
}

func randomRequestId() string {
	code := fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
	return time.Now().Format("20060102150405") + code
}

func ExistsIgnoreCase(slice []string, target string) bool {
	for _, item := range slice {
		if strings.EqualFold(item, target) {
			return true
		}
	}
	return false
}
