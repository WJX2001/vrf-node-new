package node

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/WJX2001/vrf-node-new/common/bigint"
)

var (
	ErrHeaderTraversalAheadOfProvider            = errors.New("the HeaderTraversal's internal state is ahead of the provider")
	ErrHeaderTraversalAndProviderMismatchedState = errors.New("the HeaderTraversal and provider have diverged in state")
)

type HeaderTraversal struct {
	ethClient EthClient
	chainId   uint

	latestHeader        *types.Header
	lastTraversedHeader *types.Header

	blockConfirmationDepth *big.Int
}

func NewHeaderTraversal(ethClient EthClient, fromHeader *types.Header, confDepth *big.Int, chainId uint) *HeaderTraversal {
	return &HeaderTraversal{
		ethClient:              ethClient,
		lastTraversedHeader:    fromHeader,
		blockConfirmationDepth: confDepth,
		chainId:                chainId,
	}
}

func (f *HeaderTraversal) LatestHeader() *types.Header {
	return f.latestHeader
}

func (f *HeaderTraversal) LastTraversedHeader() *types.Header {
	return f.lastTraversedHeader
}

func (f *HeaderTraversal) NextHeaders(maxSize uint64) ([]types.Header, error) {
	latestHeader, err := f.ethClient.BlockHeaderByNumber(nil)
	if err != nil {
		return nil, fmt.Errorf("unable to query latest block: %w", err)
	} else if latestHeader == nil {
		return nil, fmt.Errorf("latest header unreported")
	} else {
		f.latestHeader = latestHeader
	}

	// 计算安全同步高度（考虑确认深度）
	endHeight := new(big.Int).Sub(latestHeader.Number, f.blockConfirmationDepth)
	if endHeight.Sign() < 0 {
		// No blocks with the provided confirmation depth available
		return nil, nil
	}

	// 检查是否已同步到目标高度
	if f.lastTraversedHeader != nil {
		cmp := f.lastTraversedHeader.Number.Cmp(endHeight)
		if cmp == 0 {
			return nil, nil
		} else if cmp > 0 {
			return nil, ErrHeaderTraversalAheadOfProvider
		}
	}

	nextHeight := bigint.Zero

	// 确定起始高度
	if f.lastTraversedHeader != nil {
		// 如果已有同步记录，从 lastTraversedHeader + 1 开始
		// 否则从0开始（首次同步）
		nextHeight = new(big.Int).Add(f.lastTraversedHeader.Number, bigint.One)
	}

	// 限制批量大小并获取区块头
	endHeight = bigint.Clamp(nextHeight, endHeight, maxSize)
	headers, err := f.ethClient.BlockHeadersByRange(nextHeight, endHeight, f.chainId)
	if err != nil {
		return nil, fmt.Errorf("error querying blocks by range: %w", err)
	}

	// 验证区块链连续性
	// 第一个区块的 parentHash 应该等于 lastTraversedHeader.Hash()
	numHeaders := len(headers)
	if numHeaders == 0 {
		return nil, nil
	} else if f.lastTraversedHeader != nil && headers[0].ParentHash != f.lastTraversedHeader.Hash() {
		fmt.Println(f.lastTraversedHeader.Number)
		fmt.Println(headers[0].Number)
		fmt.Println(len(headers))
		return nil, ErrHeaderTraversalAndProviderMismatchedState
	}
	f.lastTraversedHeader = &headers[numHeaders-1]
	return headers, nil
}
