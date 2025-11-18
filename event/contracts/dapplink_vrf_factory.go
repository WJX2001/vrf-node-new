package contracts

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/google/uuid"

	"github.com/WJX2001/vrf-node-new/bindings/vrf"
	"github.com/WJX2001/vrf-node-new/database"
	"github.com/WJX2001/vrf-node-new/database/event"
	"github.com/WJX2001/vrf-node-new/database/worker"
)

type DappLinkVrfFactory struct {
	DlVrfFactoryAbi    *abi.ABI
	DlVrfFactoryFilter *vrf.DappLinkVRFFactoryFilterer
}

/**

Factory 是一个“工厂”，用来批量创建 VRF 代理合约；你的节点监听创建事件，并把新地址记录下来。
1. 用户调用 Factory.createProxy() 的时候 创建一个新的 VRF 代理合约
2. 触发 ProxyCreated 事件 （包含新代理地址）
3. 你的 vrf-node 监听到这个事件
4. 将代理地址存储到 proxy_created 表中
*/

func NewDappLinkVrfFactory() (*DappLinkVrfFactory, error) {
	dappLinkVrfFactoryAbi, err := vrf.DappLinkVRFFactoryMetaData.GetAbi()
	if err != nil {
		log.Error("get dapplink vrf factory abi fail", "err", err)
		return nil, err
	}
	dappLinkVrfFactoryFilterer, err := vrf.NewDappLinkVRFFactoryFilterer(common.Address{}, nil)
	if err != nil {
		log.Error("new dapplink vrf factory filter fail", "err", err)
		return nil, err
	}
	return &DappLinkVrfFactory{
		DlVrfFactoryAbi:    dappLinkVrfFactoryAbi,
		DlVrfFactoryFilter: dappLinkVrfFactoryFilterer,
	}, nil
}

func (dvff *DappLinkVrfFactory) ProcessDappLinkVrfFactoryEvent(db *database.DB, dappLinkVrfFactoryAddres string, startBlock, endBlock *big.Int) ([]worker.PoxyCreated, error) {
	var proxyCreatedList []worker.PoxyCreated
	contactFilter := event.ContractEvent{ContractAddress: common.HexToAddress(dappLinkVrfFactoryAddres)}
	contractEventList, err := db.ContractEvent.ContractEventsWithFilter(contactFilter, startBlock, endBlock)
	if err != nil {
		log.Error("query contacts event fail", "err", err)
		return proxyCreatedList, err
	}
	for _, contractEvent := range contractEventList {
		if contractEvent.EventSignature.String() == dvff.DlVrfFactoryAbi.Events["ProxyCreated"].ID.String() {
			proxyCreated, err := dvff.DlVrfFactoryFilter.ParseProxyCreated(*contractEvent.RLPLog)
			if err != nil {
				log.Error("proxy created fail", "err", err)
				return proxyCreatedList, err
			}
			log.Info("proxy created event", "MintProxyAddress", proxyCreated.MintProxyAddress)
			pc := worker.PoxyCreated{
				GUID:         uuid.New(),
				ProxyAddress: proxyCreated.MintProxyAddress,
				Timestamp:    uint64(time.Now().Unix()),
			}
			proxyCreatedList = append(proxyCreatedList, pc)
		}
	}
	return proxyCreatedList, nil
}
