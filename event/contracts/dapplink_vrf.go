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

func (dvm *DappLinkVrfManager) ProcessDappLinkVrfManagerEvent(db *database.DB, dappLinkVrfAddress string, startBlock, endBlock *big.Int) ([]worker.RequestSend, []worker.FillRandomWords, error) {
	var requestSendList []worker.RequestSend
	var fillRandomWordsList []worker.FillRandomWords

	contractFiler := event.ContractEvent{ContractAddress: common.HexToAddress(dappLinkVrfAddress)}
	contractEventList, err := db.ContractEvent.ContractEventsWithFilter(contractFiler, startBlock, endBlock)
	if err != nil {
		log.Error("Query contracts event list fail", "err", err)
		return requestSendList, fillRandomWordsList, err
	}
	for _, contractEvent := range contractEventList {
		log.Info("==========================================")
		log.Info("DappLink Vrf Manager Contracts EventList", "contractEventListLength", len(contractEventList))
		log.Info("DappLink Vrf Manager Contracts EventList", "contractEvent.EventSignature.String()", contractEvent.EventSignature.String(), "dvm.DappLinkVrfAbi.Events[RequestSent].ID.String()", dvm.DappLinkVrfAbi.Events["RequestSent"].ID.String())
		log.Info("==========================================")
		if contractEvent.EventSignature.String() == dvm.DappLinkVrfAbi.Events["RequestSent"].ID.String() {
			requestSent, errParse := dvm.DappLinkVrfFilter.ParseRequestSent(*contractEvent.RLPLog)
			if errParse != nil {
				log.Error("Parse request send event fail", "errParse", errParse)
				return requestSendList, fillRandomWordsList, errParse
			}
			log.Info("Parse Request Sent event success", "RequestId", requestSent.RequestId, "NumWords", requestSent.NumWords)
			rs := worker.RequestSend{
				GUID:       uuid.New(),
				RequestId:  requestSent.RequestId,
				VrfAddress: requestSent.Current,
				NumWords:   requestSent.NumWords,
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
