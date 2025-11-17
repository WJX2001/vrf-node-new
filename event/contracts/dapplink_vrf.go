package contracts

import (
	"math/big"
	"time"

	"github.com/WJX2001/vrf-node-new/bindings/vrf"
	"github.com/WJX2001/vrf-node-new/database"
	"github.com/WJX2001/vrf-node-new/database/event"
	"github.com/WJX2001/vrf-node-new/database/worker"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/google/uuid"
)

type DappLinkVrfManager struct {
	DappLinkVrfAbi    *abi.ABI
	DappLinkVrfFilter *vrf.DappLinkVRFManagerFilterer
}

func NewDappLinkVrfManager() (*DappLinkVrfManager, error) {
	dappLinkVrfAbi, err := vrf.DappLinkVRFManagerMetaData.GetAbi()
	if err != nil {
		log.Error("get dapplink vrf factory meta data abi fail", "err", err)
		return nil, err
	}

	dappLinkVrfFilter, err := vrf.NewDappLinkVRFManagerFilterer(common.Address{}, nil)

	if err != nil {
		log.Error("new dapplink vrf manager filter fail", "err", err)
		return nil, err
	}

	return &DappLinkVrfManager{
		DappLinkVrfAbi:    dappLinkVrfAbi,
		DappLinkVrfFilter: dappLinkVrfFilter,
	}, nil
}

func (dvm *DappLinkVrfManager) ProcessDappLinkVrfManagerEvent(
	db *database.DB,
	dappLinkVrfAddress string,
	startBlock, endBlock *big.Int) ([]worker.RequestSend, []worker.FillRandomWords, error) {

	var requestSendList []worker.RequestSend
	var fillRandomWordsList []worker.FillRandomWords

	// 构建过滤器：按合约地址过滤
	contractFilter := event.ContractEvent{ContractAddress: common.HexToAddress(dappLinkVrfAddress)}
	// 从 contract_events 表查询指定区块范围内该合约的所有事件
	contractEventList, err := db.ContractEvent.ContractEventsWithFilter(contractFilter, startBlock, endBlock)

	if err != nil {
		log.Error("Query contracts event list fail", "err", err)
		return requestSendList, fillRandomWordsList, err
	}

	// 遍历每个事件
	// 记录日志用于调试
	for _, contractEvent := range contractEventList {
		log.Info("==========================================")
		log.Info("DappLink Vrf Manager Contracts EventList", "contractEventListLength", len(contractEventList))
		log.Info("DappLink Vrf Manager Contracts EventList", "contractEvent.EventSignature.String()", contractEvent.EventSignature.String(), "dvm.DappLinkVrfAbi.Events[RequestSent].ID.String()", dvm.DappLinkVrfAbi.Events["RequestSent"].ID.String())
		log.Info("==========================================")

		if contractEvent.EventSignature.String() == dvm.DappLinkVrfAbi.Events["RequestSent"].ID.String() {
			// 将 RLP 日志解析为结构化数据
			requestSent, errParse := dvm.DappLinkVrfFilter.ParseRequestSent(*contractEvent.RLPLog)
			if errParse != nil {
				log.Error("Parse request send event fail", "errParse", errParse)
				return requestSendList, fillRandomWordsList, errParse
			}
			log.Info("Parse Request Sent event success", "RequestId", requestSent.RequestId, "NumWords", requestSent.NumWords)
			rs := worker.RequestSend{
				GUID:       uuid.New(),            // 生成唯一 ID
				RequestId:  requestSent.RequestId, // 请求 ID
				VrfAddress: requestSent.Current,   // VRF 合约地址
				NumWords:   requestSent.NumWords,  //
				Status:     0,
				Timestamp:  uint64(time.Now().Unix()),
			}
			requestSendList = append(requestSendList, rs)
		}

		if contractEvent.EventSignature.String() == dvm.DappLinkVrfAbi.Events["FillRandomWords"].ID.String() {
			fillRandomWords, errParse := dvm.DappLinkVrfFilter.ParseFillRandomWords(*contractEvent.RLPLog)
			if errParse != nil {
				log.Error("Parse fill random words fail", "errParse", errParse)
				return requestSendList, fillRandomWordsList, errParse
			}
			log.Info("Parse fillRandomWords event success", "RequestId", fillRandomWords.RequestId, "RandomWords", fillRandomWords.RandomWords)

			var randomWords string
			for _, rword := range fillRandomWords.RandomWords {
				randomWords = rword.String()
			}

			frw := worker.FillRandomWords{
				GUID:        uuid.New(),
				RequestId:   fillRandomWords.RequestId,
				RandomWords: randomWords,
				Timestamp:   uint64(time.Now().Unix()),
			}
			fillRandomWordsList = append(fillRandomWordsList, frw)
		}
	}
	return requestSendList, fillRandomWordsList, nil
}
