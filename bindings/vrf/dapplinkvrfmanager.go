// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vrf

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// BN254G2Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// IBLSApkRegistryVrfNoSignerAndSignature is an auto generated low-level Go binding around an user-defined struct.
type IBLSApkRegistryVrfNoSignerAndSignature struct {
	NonSignerPubKeys   []BN254G1Point
	ApkG2              BN254G2Point
	Sigma              BN254G1Point
	TotalDappLinkStake *big.Int
	TotalBtcStake      *big.Int
}

// DappLinkVRFManagerMetaData contains all meta data concerning the DappLinkVRFManager contract.
var DappLinkVRFManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"blsRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"dappLinkAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fulfillRandomWords\",\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_randomWords\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSApkRegistry.VrfNoSignerAndSignature\",\"components\":[{\"name\":\"nonSignerPubKeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"totalDappLinkStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalBtcStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getRequestStatus\",\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"fulfilled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"randomWords\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_dappLinkAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_blsRegistry\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lastRequestId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestIds\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requestMapping\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"fulfilled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requestRandomWords\",\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_numWords\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDappLink\",\"inputs\":[{\"name\":\"_dappLinkAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"FillRandomWords\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"randomWords\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RequestSent\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_numWords\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"current\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608060405234801562000010575f80fd5b50620000216200002760201b60201c565b62000191565b5f620000386200012b60201b60201c565b9050805f0160089054906101000a900460ff161562000083576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8016815f015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff1614620001285767ffffffffffffffff815f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d267ffffffffffffffff6040516200011f919062000176565b60405180910390a15b50565b5f7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00905090565b5f67ffffffffffffffff82169050919050565b620001708162000152565b82525050565b5f6020820190506200018b5f83018462000165565b92915050565b61195d806200019f5f395ff3fe608060405234801561000f575f80fd5b50600436106100cd575f3560e01c806393e4b8801161008a578063d8a4676f11610064578063d8a4676f146101e7578063f0c28a4114610218578063f2fde38b14610236578063fc2a88c314610252576100cd565b806393e4b88014610191578063996869d0146101af578063c0c53b8b146101cb576100cd565b80631b565710146100d15780631b739ef1146100ed578063715018a61461010957806382e215ab146101135780638796ba8c146101435780638da5cb5b14610173575b5f80fd5b6100eb60048036038101906100e6919061112a565b610270565b005b610107600480360381019061010291906111d9565b610444565b005b610111610564565b005b61012d60048036038101906101289190611217565b610577565b60405161013a919061125c565b60405180910390f35b61015d60048036038101906101589190611217565b61059d565b60405161016a9190611284565b60405180910390f35b61017b6105bc565b60405161018891906112dc565b60405180910390f35b6101996105f1565b6040516101a69190611350565b60405180910390f35b6101c960048036038101906101c49190611393565b610616565b005b6101e560048036038101906101e091906113be565b610661565b005b61020160048036038101906101fc9190611217565b610864565b60405161020f9291906114c5565b60405180910390f35b6102206108f4565b60405161022d91906112dc565b60405180910390f35b610250600480360381019061024b9190611393565b610919565b005b61025a61099d565b6040516102679190611284565b60405180910390f35b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146102ff576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102f69061154d565b60405180910390fd5b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16631070de848484846040518463ffffffff1660e01b815260040161035d93929190611770565b606060405180830381865afa158015610378573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061039c9190611821565b505060405180604001604052806001151581526020018581525060045f8781526020019081526020015f205f820151815f015f6101000a81548160ff0219169083151502179055506020820151816001019080519060200190610400929190610c46565b509050507ff3cb4deb0441dd096356debf166f879d78cadc19e4b94053c8bea6d3940de93a858560405161043592919061185f565b60405180910390a15050505050565b61044c6109a3565b60405180604001604052805f151581526020015f67ffffffffffffffff81111561047957610478610d04565b5b6040519080825280602002602001820160405280156104a75781602001602082028036833780820191505090505b5081525060045f8481526020019081526020015f205f820151815f015f6101000a81548160ff02191690831515021790555060208201518160010190805190602001906104f5929190610c46565b509050505f82908060018154018082558091505060019003905f5260205f20015f9091909190915055816001819055507fe697eb68c0228bd7d4e553246a2a86e8402d0895e45092ef8ae87b4cfd29f0168282306040516105589392919061188d565b60405180910390a15050565b61056c6109a3565b6105755f610a2a565b565b6004602052805f5260405f205f91509050805f015f9054906101000a900460ff16905081565b5f81815481106105ab575f80fd5b905f5260205f20015f915090505481565b5f806105c6610afb565b9050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1691505090565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b61061e6109a3565b8060025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f61066a610b22565b90505f815f0160089054906101000a900460ff161590505f825f015f9054906101000a900467ffffffffffffffff1690505f808267ffffffffffffffff161480156106b25750825b90505f60018367ffffffffffffffff161480156106e557505f3073ffffffffffffffffffffffffffffffffffffffff163b145b9050811580156106f3575080155b1561072a576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001855f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508315610777576001855f0160086101000a81548160ff0219169083151502179055505b61078088610b49565b8560035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508660025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550831561085a575f855f0160086101000a81548160ff0219169083151502179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d26001604051610851919061190e565b60405180910390a15b5050505050505050565b5f606060045f8481526020019081526020015f205f015f9054906101000a900460ff1660045f8581526020019081526020015f20600101808054806020026020016040519081016040528092919081815260200182805480156108e457602002820191905f5260205f20905b8154815260200190600101908083116108d0575b5050505050905091509150915091565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6109216109a3565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610991575f6040517f1e4fbdf700000000000000000000000000000000000000000000000000000000815260040161098891906112dc565b60405180910390fd5b61099a81610a2a565b50565b60015481565b6109ab610b5d565b73ffffffffffffffffffffffffffffffffffffffff166109c96105bc565b73ffffffffffffffffffffffffffffffffffffffff1614610a28576109ec610b5d565b6040517f118cdaa7000000000000000000000000000000000000000000000000000000008152600401610a1f91906112dc565b60405180910390fd5b565b5f610a33610afb565b90505f815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905082825f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3505050565b5f7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300905090565b5f7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00905090565b610b51610b64565b610b5a81610ba4565b50565b5f33905090565b610b6c610c28565b610ba2576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b610bac610b64565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610c1c575f6040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401610c1391906112dc565b60405180910390fd5b610c2581610a2a565b50565b5f610c31610b22565b5f0160089054906101000a900460ff16905090565b828054828255905f5260205f20908101928215610c80579160200282015b82811115610c7f578251825591602001919060010190610c64565b5b509050610c8d9190610c91565b5090565b5b80821115610ca8575f815f905550600101610c92565b5090565b5f604051905090565b5f80fd5b5f80fd5b5f819050919050565b610ccf81610cbd565b8114610cd9575f80fd5b50565b5f81359050610cea81610cc6565b92915050565b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610d3a82610cf4565b810181811067ffffffffffffffff82111715610d5957610d58610d04565b5b80604052505050565b5f610d6b610cac565b9050610d778282610d31565b919050565b5f67ffffffffffffffff821115610d9657610d95610d04565b5b602082029050602081019050919050565b5f80fd5b5f610dbd610db884610d7c565b610d62565b90508083825260208201905060208402830185811115610de057610ddf610da7565b5b835b81811015610e095780610df58882610cdc565b845260208401935050602081019050610de2565b5050509392505050565b5f82601f830112610e2757610e26610cf0565b5b8135610e37848260208601610dab565b91505092915050565b5f819050919050565b610e5281610e40565b8114610e5c575f80fd5b50565b5f81359050610e6d81610e49565b92915050565b5f80fd5b5f80fd5b5f67ffffffffffffffff821115610e9557610e94610d04565b5b602082029050602081019050919050565b5f60408284031215610ebb57610eba610e73565b5b610ec56040610d62565b90505f610ed484828501610cdc565b5f830152506020610ee784828501610cdc565b60208301525092915050565b5f610f05610f0084610e7b565b610d62565b90508083825260208201905060408402830185811115610f2857610f27610da7565b5b835b81811015610f515780610f3d8882610ea6565b845260208401935050604081019050610f2a565b5050509392505050565b5f82601f830112610f6f57610f6e610cf0565b5b8135610f7f848260208601610ef3565b91505092915050565b5f67ffffffffffffffff821115610fa257610fa1610d04565b5b602082029050919050565b5f610fbf610fba84610f88565b610d62565b90508060208402830185811115610fd957610fd8610da7565b5b835b818110156110025780610fee8882610cdc565b845260208401935050602081019050610fdb565b5050509392505050565b5f82601f8301126110205761101f610cf0565b5b600261102d848285610fad565b91505092915050565b5f6080828403121561104b5761104a610e73565b5b6110556040610d62565b90505f6110648482850161100c565b5f8301525060406110778482850161100c565b60208301525092915050565b5f610120828403121561109957611098610e73565b5b6110a360a0610d62565b90505f82013567ffffffffffffffff8111156110c2576110c1610e77565b5b6110ce84828501610f5b565b5f8301525060206110e184828501611036565b60208301525060a06110f584828501610ea6565b60408301525060e061110984828501610cdc565b60608301525061010061111e84828501610cdc565b60808301525092915050565b5f805f805f60a0868803121561114357611142610cb5565b5b5f61115088828901610cdc565b955050602086013567ffffffffffffffff81111561117157611170610cb9565b5b61117d88828901610e13565b945050604061118e88828901610e5f565b935050606061119f88828901610cdc565b925050608086013567ffffffffffffffff8111156111c0576111bf610cb9565b5b6111cc88828901611083565b9150509295509295909350565b5f80604083850312156111ef576111ee610cb5565b5b5f6111fc85828601610cdc565b925050602061120d85828601610cdc565b9150509250929050565b5f6020828403121561122c5761122b610cb5565b5b5f61123984828501610cdc565b91505092915050565b5f8115159050919050565b61125681611242565b82525050565b5f60208201905061126f5f83018461124d565b92915050565b61127e81610cbd565b82525050565b5f6020820190506112975f830184611275565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6112c68261129d565b9050919050565b6112d6816112bc565b82525050565b5f6020820190506112ef5f8301846112cd565b92915050565b5f819050919050565b5f61131861131361130e8461129d565b6112f5565b61129d565b9050919050565b5f611329826112fe565b9050919050565b5f61133a8261131f565b9050919050565b61134a81611330565b82525050565b5f6020820190506113635f830184611341565b92915050565b611372816112bc565b811461137c575f80fd5b50565b5f8135905061138d81611369565b92915050565b5f602082840312156113a8576113a7610cb5565b5b5f6113b58482850161137f565b91505092915050565b5f805f606084860312156113d5576113d4610cb5565b5b5f6113e28682870161137f565b93505060206113f38682870161137f565b92505060406114048682870161137f565b9150509250925092565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b61144081610cbd565b82525050565b5f6114518383611437565b60208301905092915050565b5f602082019050919050565b5f6114738261140e565b61147d8185611418565b935061148883611428565b805f5b838110156114b857815161149f8882611446565b97506114aa8361145d565b92505060018101905061148b565b5085935050505092915050565b5f6040820190506114d85f83018561124d565b81810360208301526114ea8184611469565b90509392505050565b5f82825260208201905092915050565b7f446170704c696e6b5652462e6f6e6c79446170704c696e6b00000000000000005f82015250565b5f6115376018836114f3565b915061154282611503565b602082019050919050565b5f6020820190508181035f8301526115648161152b565b9050919050565b61157481610e40565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b604082015f8201516115b75f850182611437565b5060208201516115ca6020850182611437565b50505050565b5f6115db83836115a3565b60408301905092915050565b5f602082019050919050565b5f6115fd8261157a565b6116078185611584565b935061161283611594565b805f5b8381101561164257815161162988826115d0565b9750611634836115e7565b925050600181019050611615565b5085935050505092915050565b5f60029050919050565b5f81905092915050565b5f819050919050565b5f602082019050919050565b6116818161164f565b61168b8184611659565b925061169682611663565b805f5b838110156116c65781516116ad8782611446565b96506116b88361166c565b925050600181019050611699565b505050505050565b608082015f8201516116e25f850182611678565b5060208201516116f56040850182611678565b50505050565b5f61012083015f8301518482035f86015261171682826115f3565b915050602083015161172b60208601826116ce565b50604083015161173e60a08601826115a3565b50606083015161175160e0860182611437565b506080830151611765610100860182611437565b508091505092915050565b5f6060820190506117835f83018661156b565b6117906020830185611275565b81810360408301526117a281846116fb565b9050949350505050565b5f815190506117ba81610cc6565b92915050565b5f604082840312156117d5576117d4610e73565b5b6117df6040610d62565b90505f6117ee848285016117ac565b5f830152506020611801848285016117ac565b60208301525092915050565b5f8151905061181b81610e49565b92915050565b5f806060838503121561183757611836610cb5565b5b5f611844858286016117c0565b92505060406118558582860161180d565b9150509250929050565b5f6040820190506118725f830185611275565b81810360208301526118848184611469565b90509392505050565b5f6060820190506118a05f830186611275565b6118ad6020830185611275565b6118ba60408301846112cd565b949350505050565b5f819050919050565b5f67ffffffffffffffff82169050919050565b5f6118f86118f36118ee846118c2565b6112f5565b6118cb565b9050919050565b611908816118de565b82525050565b5f6020820190506119215f8301846118ff565b9291505056fea26469706673582212203ffc211d22a9cb4b2158092a347dd6bc2138f29e1459e393dabe71d087f9237464736f6c63430008160033",
}

// DappLinkVRFManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use DappLinkVRFManagerMetaData.ABI instead.
var DappLinkVRFManagerABI = DappLinkVRFManagerMetaData.ABI

// DappLinkVRFManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DappLinkVRFManagerMetaData.Bin instead.
var DappLinkVRFManagerBin = DappLinkVRFManagerMetaData.Bin

// DeployDappLinkVRFManager deploys a new Ethereum contract, binding an instance of DappLinkVRFManager to it.
func DeployDappLinkVRFManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DappLinkVRFManager, error) {
	parsed, err := DappLinkVRFManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DappLinkVRFManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DappLinkVRFManager{DappLinkVRFManagerCaller: DappLinkVRFManagerCaller{contract: contract}, DappLinkVRFManagerTransactor: DappLinkVRFManagerTransactor{contract: contract}, DappLinkVRFManagerFilterer: DappLinkVRFManagerFilterer{contract: contract}}, nil
}

// DappLinkVRFManager is an auto generated Go binding around an Ethereum contract.
type DappLinkVRFManager struct {
	DappLinkVRFManagerCaller     // Read-only binding to the contract
	DappLinkVRFManagerTransactor // Write-only binding to the contract
	DappLinkVRFManagerFilterer   // Log filterer for contract events
}

// DappLinkVRFManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type DappLinkVRFManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DappLinkVRFManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DappLinkVRFManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DappLinkVRFManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DappLinkVRFManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DappLinkVRFManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DappLinkVRFManagerSession struct {
	Contract     *DappLinkVRFManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DappLinkVRFManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DappLinkVRFManagerCallerSession struct {
	Contract *DappLinkVRFManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// DappLinkVRFManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DappLinkVRFManagerTransactorSession struct {
	Contract     *DappLinkVRFManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// DappLinkVRFManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type DappLinkVRFManagerRaw struct {
	Contract *DappLinkVRFManager // Generic contract binding to access the raw methods on
}

// DappLinkVRFManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DappLinkVRFManagerCallerRaw struct {
	Contract *DappLinkVRFManagerCaller // Generic read-only contract binding to access the raw methods on
}

// DappLinkVRFManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DappLinkVRFManagerTransactorRaw struct {
	Contract *DappLinkVRFManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDappLinkVRFManager creates a new instance of DappLinkVRFManager, bound to a specific deployed contract.
func NewDappLinkVRFManager(address common.Address, backend bind.ContractBackend) (*DappLinkVRFManager, error) {
	contract, err := bindDappLinkVRFManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DappLinkVRFManager{DappLinkVRFManagerCaller: DappLinkVRFManagerCaller{contract: contract}, DappLinkVRFManagerTransactor: DappLinkVRFManagerTransactor{contract: contract}, DappLinkVRFManagerFilterer: DappLinkVRFManagerFilterer{contract: contract}}, nil
}

// NewDappLinkVRFManagerCaller creates a new read-only instance of DappLinkVRFManager, bound to a specific deployed contract.
func NewDappLinkVRFManagerCaller(address common.Address, caller bind.ContractCaller) (*DappLinkVRFManagerCaller, error) {
	contract, err := bindDappLinkVRFManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DappLinkVRFManagerCaller{contract: contract}, nil
}

// NewDappLinkVRFManagerTransactor creates a new write-only instance of DappLinkVRFManager, bound to a specific deployed contract.
func NewDappLinkVRFManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*DappLinkVRFManagerTransactor, error) {
	contract, err := bindDappLinkVRFManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DappLinkVRFManagerTransactor{contract: contract}, nil
}

// NewDappLinkVRFManagerFilterer creates a new log filterer instance of DappLinkVRFManager, bound to a specific deployed contract.
func NewDappLinkVRFManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*DappLinkVRFManagerFilterer, error) {
	contract, err := bindDappLinkVRFManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DappLinkVRFManagerFilterer{contract: contract}, nil
}

// bindDappLinkVRFManager binds a generic wrapper to an already deployed contract.
func bindDappLinkVRFManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DappLinkVRFManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DappLinkVRFManager *DappLinkVRFManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DappLinkVRFManager.Contract.DappLinkVRFManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DappLinkVRFManager *DappLinkVRFManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DappLinkVRFManager.Contract.DappLinkVRFManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DappLinkVRFManager *DappLinkVRFManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DappLinkVRFManager.Contract.DappLinkVRFManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DappLinkVRFManager *DappLinkVRFManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DappLinkVRFManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DappLinkVRFManager *DappLinkVRFManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DappLinkVRFManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DappLinkVRFManager *DappLinkVRFManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DappLinkVRFManager.Contract.contract.Transact(opts, method, params...)
}

// BlsRegistry is a free data retrieval call binding the contract method 0x93e4b880.
//
// Solidity: function blsRegistry() view returns(address)
func (_DappLinkVRFManager *DappLinkVRFManagerCaller) BlsRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DappLinkVRFManager.contract.Call(opts, &out, "blsRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsRegistry is a free data retrieval call binding the contract method 0x93e4b880.
//
// Solidity: function blsRegistry() view returns(address)
func (_DappLinkVRFManager *DappLinkVRFManagerSession) BlsRegistry() (common.Address, error) {
	return _DappLinkVRFManager.Contract.BlsRegistry(&_DappLinkVRFManager.CallOpts)
}

// BlsRegistry is a free data retrieval call binding the contract method 0x93e4b880.
//
// Solidity: function blsRegistry() view returns(address)
func (_DappLinkVRFManager *DappLinkVRFManagerCallerSession) BlsRegistry() (common.Address, error) {
	return _DappLinkVRFManager.Contract.BlsRegistry(&_DappLinkVRFManager.CallOpts)
}

// DappLinkAddress is a free data retrieval call binding the contract method 0xf0c28a41.
//
// Solidity: function dappLinkAddress() view returns(address)
func (_DappLinkVRFManager *DappLinkVRFManagerCaller) DappLinkAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DappLinkVRFManager.contract.Call(opts, &out, "dappLinkAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DappLinkAddress is a free data retrieval call binding the contract method 0xf0c28a41.
//
// Solidity: function dappLinkAddress() view returns(address)
func (_DappLinkVRFManager *DappLinkVRFManagerSession) DappLinkAddress() (common.Address, error) {
	return _DappLinkVRFManager.Contract.DappLinkAddress(&_DappLinkVRFManager.CallOpts)
}

// DappLinkAddress is a free data retrieval call binding the contract method 0xf0c28a41.
//
// Solidity: function dappLinkAddress() view returns(address)
func (_DappLinkVRFManager *DappLinkVRFManagerCallerSession) DappLinkAddress() (common.Address, error) {
	return _DappLinkVRFManager.Contract.DappLinkAddress(&_DappLinkVRFManager.CallOpts)
}

// GetRequestStatus is a free data retrieval call binding the contract method 0xd8a4676f.
//
// Solidity: function getRequestStatus(uint256 _requestId) view returns(bool fulfilled, uint256[] randomWords)
func (_DappLinkVRFManager *DappLinkVRFManagerCaller) GetRequestStatus(opts *bind.CallOpts, _requestId *big.Int) (struct {
	Fulfilled   bool
	RandomWords []*big.Int
}, error) {
	var out []interface{}
	err := _DappLinkVRFManager.contract.Call(opts, &out, "getRequestStatus", _requestId)

	outstruct := new(struct {
		Fulfilled   bool
		RandomWords []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fulfilled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.RandomWords = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetRequestStatus is a free data retrieval call binding the contract method 0xd8a4676f.
//
// Solidity: function getRequestStatus(uint256 _requestId) view returns(bool fulfilled, uint256[] randomWords)
func (_DappLinkVRFManager *DappLinkVRFManagerSession) GetRequestStatus(_requestId *big.Int) (struct {
	Fulfilled   bool
	RandomWords []*big.Int
}, error) {
	return _DappLinkVRFManager.Contract.GetRequestStatus(&_DappLinkVRFManager.CallOpts, _requestId)
}

// GetRequestStatus is a free data retrieval call binding the contract method 0xd8a4676f.
//
// Solidity: function getRequestStatus(uint256 _requestId) view returns(bool fulfilled, uint256[] randomWords)
func (_DappLinkVRFManager *DappLinkVRFManagerCallerSession) GetRequestStatus(_requestId *big.Int) (struct {
	Fulfilled   bool
	RandomWords []*big.Int
}, error) {
	return _DappLinkVRFManager.Contract.GetRequestStatus(&_DappLinkVRFManager.CallOpts, _requestId)
}

// LastRequestId is a free data retrieval call binding the contract method 0xfc2a88c3.
//
// Solidity: function lastRequestId() view returns(uint256)
func (_DappLinkVRFManager *DappLinkVRFManagerCaller) LastRequestId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DappLinkVRFManager.contract.Call(opts, &out, "lastRequestId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRequestId is a free data retrieval call binding the contract method 0xfc2a88c3.
//
// Solidity: function lastRequestId() view returns(uint256)
func (_DappLinkVRFManager *DappLinkVRFManagerSession) LastRequestId() (*big.Int, error) {
	return _DappLinkVRFManager.Contract.LastRequestId(&_DappLinkVRFManager.CallOpts)
}

// LastRequestId is a free data retrieval call binding the contract method 0xfc2a88c3.
//
// Solidity: function lastRequestId() view returns(uint256)
func (_DappLinkVRFManager *DappLinkVRFManagerCallerSession) LastRequestId() (*big.Int, error) {
	return _DappLinkVRFManager.Contract.LastRequestId(&_DappLinkVRFManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DappLinkVRFManager *DappLinkVRFManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DappLinkVRFManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DappLinkVRFManager *DappLinkVRFManagerSession) Owner() (common.Address, error) {
	return _DappLinkVRFManager.Contract.Owner(&_DappLinkVRFManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DappLinkVRFManager *DappLinkVRFManagerCallerSession) Owner() (common.Address, error) {
	return _DappLinkVRFManager.Contract.Owner(&_DappLinkVRFManager.CallOpts)
}

// RequestIds is a free data retrieval call binding the contract method 0x8796ba8c.
//
// Solidity: function requestIds(uint256 ) view returns(uint256)
func (_DappLinkVRFManager *DappLinkVRFManagerCaller) RequestIds(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DappLinkVRFManager.contract.Call(opts, &out, "requestIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestIds is a free data retrieval call binding the contract method 0x8796ba8c.
//
// Solidity: function requestIds(uint256 ) view returns(uint256)
func (_DappLinkVRFManager *DappLinkVRFManagerSession) RequestIds(arg0 *big.Int) (*big.Int, error) {
	return _DappLinkVRFManager.Contract.RequestIds(&_DappLinkVRFManager.CallOpts, arg0)
}

// RequestIds is a free data retrieval call binding the contract method 0x8796ba8c.
//
// Solidity: function requestIds(uint256 ) view returns(uint256)
func (_DappLinkVRFManager *DappLinkVRFManagerCallerSession) RequestIds(arg0 *big.Int) (*big.Int, error) {
	return _DappLinkVRFManager.Contract.RequestIds(&_DappLinkVRFManager.CallOpts, arg0)
}

// RequestMapping is a free data retrieval call binding the contract method 0x82e215ab.
//
// Solidity: function requestMapping(uint256 ) view returns(bool fulfilled)
func (_DappLinkVRFManager *DappLinkVRFManagerCaller) RequestMapping(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _DappLinkVRFManager.contract.Call(opts, &out, "requestMapping", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RequestMapping is a free data retrieval call binding the contract method 0x82e215ab.
//
// Solidity: function requestMapping(uint256 ) view returns(bool fulfilled)
func (_DappLinkVRFManager *DappLinkVRFManagerSession) RequestMapping(arg0 *big.Int) (bool, error) {
	return _DappLinkVRFManager.Contract.RequestMapping(&_DappLinkVRFManager.CallOpts, arg0)
}

// RequestMapping is a free data retrieval call binding the contract method 0x82e215ab.
//
// Solidity: function requestMapping(uint256 ) view returns(bool fulfilled)
func (_DappLinkVRFManager *DappLinkVRFManagerCallerSession) RequestMapping(arg0 *big.Int) (bool, error) {
	return _DappLinkVRFManager.Contract.RequestMapping(&_DappLinkVRFManager.CallOpts, arg0)
}

// FulfillRandomWords is a paid mutator transaction binding the contract method 0x1b565710.
//
// Solidity: function fulfillRandomWords(uint256 _requestId, uint256[] _randomWords, bytes32 msgHash, uint256 referenceBlockNumber, ((uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint256,uint256) params) returns()
func (_DappLinkVRFManager *DappLinkVRFManagerTransactor) FulfillRandomWords(opts *bind.TransactOpts, _requestId *big.Int, _randomWords []*big.Int, msgHash [32]byte, referenceBlockNumber *big.Int, params IBLSApkRegistryVrfNoSignerAndSignature) (*types.Transaction, error) {
	return _DappLinkVRFManager.contract.Transact(opts, "fulfillRandomWords", _requestId, _randomWords, msgHash, referenceBlockNumber, params)
}

// FulfillRandomWords is a paid mutator transaction binding the contract method 0x1b565710.
//
// Solidity: function fulfillRandomWords(uint256 _requestId, uint256[] _randomWords, bytes32 msgHash, uint256 referenceBlockNumber, ((uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint256,uint256) params) returns()
func (_DappLinkVRFManager *DappLinkVRFManagerSession) FulfillRandomWords(_requestId *big.Int, _randomWords []*big.Int, msgHash [32]byte, referenceBlockNumber *big.Int, params IBLSApkRegistryVrfNoSignerAndSignature) (*types.Transaction, error) {
	return _DappLinkVRFManager.Contract.FulfillRandomWords(&_DappLinkVRFManager.TransactOpts, _requestId, _randomWords, msgHash, referenceBlockNumber, params)
}

// FulfillRandomWords is a paid mutator transaction binding the contract method 0x1b565710.
//
// Solidity: function fulfillRandomWords(uint256 _requestId, uint256[] _randomWords, bytes32 msgHash, uint256 referenceBlockNumber, ((uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint256,uint256) params) returns()
func (_DappLinkVRFManager *DappLinkVRFManagerTransactorSession) FulfillRandomWords(_requestId *big.Int, _randomWords []*big.Int, msgHash [32]byte, referenceBlockNumber *big.Int, params IBLSApkRegistryVrfNoSignerAndSignature) (*types.Transaction, error) {
	return _DappLinkVRFManager.Contract.FulfillRandomWords(&_DappLinkVRFManager.TransactOpts, _requestId, _randomWords, msgHash, referenceBlockNumber, params)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address initialOwner, address _dappLinkAddress, address _blsRegistry) returns()
func (_DappLinkVRFManager *DappLinkVRFManagerTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, _dappLinkAddress common.Address, _blsRegistry common.Address) (*types.Transaction, error) {
	return _DappLinkVRFManager.contract.Transact(opts, "initialize", initialOwner, _dappLinkAddress, _blsRegistry)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address initialOwner, address _dappLinkAddress, address _blsRegistry) returns()
func (_DappLinkVRFManager *DappLinkVRFManagerSession) Initialize(initialOwner common.Address, _dappLinkAddress common.Address, _blsRegistry common.Address) (*types.Transaction, error) {
	return _DappLinkVRFManager.Contract.Initialize(&_DappLinkVRFManager.TransactOpts, initialOwner, _dappLinkAddress, _blsRegistry)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address initialOwner, address _dappLinkAddress, address _blsRegistry) returns()
func (_DappLinkVRFManager *DappLinkVRFManagerTransactorSession) Initialize(initialOwner common.Address, _dappLinkAddress common.Address, _blsRegistry common.Address) (*types.Transaction, error) {
	return _DappLinkVRFManager.Contract.Initialize(&_DappLinkVRFManager.TransactOpts, initialOwner, _dappLinkAddress, _blsRegistry)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DappLinkVRFManager *DappLinkVRFManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DappLinkVRFManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DappLinkVRFManager *DappLinkVRFManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _DappLinkVRFManager.Contract.RenounceOwnership(&_DappLinkVRFManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DappLinkVRFManager *DappLinkVRFManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DappLinkVRFManager.Contract.RenounceOwnership(&_DappLinkVRFManager.TransactOpts)
}

// RequestRandomWords is a paid mutator transaction binding the contract method 0x1b739ef1.
//
// Solidity: function requestRandomWords(uint256 _requestId, uint256 _numWords) returns()
func (_DappLinkVRFManager *DappLinkVRFManagerTransactor) RequestRandomWords(opts *bind.TransactOpts, _requestId *big.Int, _numWords *big.Int) (*types.Transaction, error) {
	return _DappLinkVRFManager.contract.Transact(opts, "requestRandomWords", _requestId, _numWords)
}

// RequestRandomWords is a paid mutator transaction binding the contract method 0x1b739ef1.
//
// Solidity: function requestRandomWords(uint256 _requestId, uint256 _numWords) returns()
func (_DappLinkVRFManager *DappLinkVRFManagerSession) RequestRandomWords(_requestId *big.Int, _numWords *big.Int) (*types.Transaction, error) {
	return _DappLinkVRFManager.Contract.RequestRandomWords(&_DappLinkVRFManager.TransactOpts, _requestId, _numWords)
}

// RequestRandomWords is a paid mutator transaction binding the contract method 0x1b739ef1.
//
// Solidity: function requestRandomWords(uint256 _requestId, uint256 _numWords) returns()
func (_DappLinkVRFManager *DappLinkVRFManagerTransactorSession) RequestRandomWords(_requestId *big.Int, _numWords *big.Int) (*types.Transaction, error) {
	return _DappLinkVRFManager.Contract.RequestRandomWords(&_DappLinkVRFManager.TransactOpts, _requestId, _numWords)
}

// SetDappLink is a paid mutator transaction binding the contract method 0x996869d0.
//
// Solidity: function setDappLink(address _dappLinkAddress) returns()
func (_DappLinkVRFManager *DappLinkVRFManagerTransactor) SetDappLink(opts *bind.TransactOpts, _dappLinkAddress common.Address) (*types.Transaction, error) {
	return _DappLinkVRFManager.contract.Transact(opts, "setDappLink", _dappLinkAddress)
}

// SetDappLink is a paid mutator transaction binding the contract method 0x996869d0.
//
// Solidity: function setDappLink(address _dappLinkAddress) returns()
func (_DappLinkVRFManager *DappLinkVRFManagerSession) SetDappLink(_dappLinkAddress common.Address) (*types.Transaction, error) {
	return _DappLinkVRFManager.Contract.SetDappLink(&_DappLinkVRFManager.TransactOpts, _dappLinkAddress)
}

// SetDappLink is a paid mutator transaction binding the contract method 0x996869d0.
//
// Solidity: function setDappLink(address _dappLinkAddress) returns()
func (_DappLinkVRFManager *DappLinkVRFManagerTransactorSession) SetDappLink(_dappLinkAddress common.Address) (*types.Transaction, error) {
	return _DappLinkVRFManager.Contract.SetDappLink(&_DappLinkVRFManager.TransactOpts, _dappLinkAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DappLinkVRFManager *DappLinkVRFManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DappLinkVRFManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DappLinkVRFManager *DappLinkVRFManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DappLinkVRFManager.Contract.TransferOwnership(&_DappLinkVRFManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DappLinkVRFManager *DappLinkVRFManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DappLinkVRFManager.Contract.TransferOwnership(&_DappLinkVRFManager.TransactOpts, newOwner)
}

// DappLinkVRFManagerFillRandomWordsIterator is returned from FilterFillRandomWords and is used to iterate over the raw logs and unpacked data for FillRandomWords events raised by the DappLinkVRFManager contract.
type DappLinkVRFManagerFillRandomWordsIterator struct {
	Event *DappLinkVRFManagerFillRandomWords // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DappLinkVRFManagerFillRandomWordsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DappLinkVRFManagerFillRandomWords)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DappLinkVRFManagerFillRandomWords)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DappLinkVRFManagerFillRandomWordsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DappLinkVRFManagerFillRandomWordsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DappLinkVRFManagerFillRandomWords represents a FillRandomWords event raised by the DappLinkVRFManager contract.
type DappLinkVRFManagerFillRandomWords struct {
	RequestId   *big.Int
	RandomWords []*big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFillRandomWords is a free log retrieval operation binding the contract event 0xf3cb4deb0441dd096356debf166f879d78cadc19e4b94053c8bea6d3940de93a.
//
// Solidity: event FillRandomWords(uint256 requestId, uint256[] randomWords)
func (_DappLinkVRFManager *DappLinkVRFManagerFilterer) FilterFillRandomWords(opts *bind.FilterOpts) (*DappLinkVRFManagerFillRandomWordsIterator, error) {

	logs, sub, err := _DappLinkVRFManager.contract.FilterLogs(opts, "FillRandomWords")
	if err != nil {
		return nil, err
	}
	return &DappLinkVRFManagerFillRandomWordsIterator{contract: _DappLinkVRFManager.contract, event: "FillRandomWords", logs: logs, sub: sub}, nil
}

// WatchFillRandomWords is a free log subscription operation binding the contract event 0xf3cb4deb0441dd096356debf166f879d78cadc19e4b94053c8bea6d3940de93a.
//
// Solidity: event FillRandomWords(uint256 requestId, uint256[] randomWords)
func (_DappLinkVRFManager *DappLinkVRFManagerFilterer) WatchFillRandomWords(opts *bind.WatchOpts, sink chan<- *DappLinkVRFManagerFillRandomWords) (event.Subscription, error) {

	logs, sub, err := _DappLinkVRFManager.contract.WatchLogs(opts, "FillRandomWords")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DappLinkVRFManagerFillRandomWords)
				if err := _DappLinkVRFManager.contract.UnpackLog(event, "FillRandomWords", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFillRandomWords is a log parse operation binding the contract event 0xf3cb4deb0441dd096356debf166f879d78cadc19e4b94053c8bea6d3940de93a.
//
// Solidity: event FillRandomWords(uint256 requestId, uint256[] randomWords)
func (_DappLinkVRFManager *DappLinkVRFManagerFilterer) ParseFillRandomWords(log types.Log) (*DappLinkVRFManagerFillRandomWords, error) {
	event := new(DappLinkVRFManagerFillRandomWords)
	if err := _DappLinkVRFManager.contract.UnpackLog(event, "FillRandomWords", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DappLinkVRFManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the DappLinkVRFManager contract.
type DappLinkVRFManagerInitializedIterator struct {
	Event *DappLinkVRFManagerInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DappLinkVRFManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DappLinkVRFManagerInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DappLinkVRFManagerInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DappLinkVRFManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DappLinkVRFManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DappLinkVRFManagerInitialized represents a Initialized event raised by the DappLinkVRFManager contract.
type DappLinkVRFManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_DappLinkVRFManager *DappLinkVRFManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*DappLinkVRFManagerInitializedIterator, error) {

	logs, sub, err := _DappLinkVRFManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DappLinkVRFManagerInitializedIterator{contract: _DappLinkVRFManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_DappLinkVRFManager *DappLinkVRFManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DappLinkVRFManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _DappLinkVRFManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DappLinkVRFManagerInitialized)
				if err := _DappLinkVRFManager.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_DappLinkVRFManager *DappLinkVRFManagerFilterer) ParseInitialized(log types.Log) (*DappLinkVRFManagerInitialized, error) {
	event := new(DappLinkVRFManagerInitialized)
	if err := _DappLinkVRFManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DappLinkVRFManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DappLinkVRFManager contract.
type DappLinkVRFManagerOwnershipTransferredIterator struct {
	Event *DappLinkVRFManagerOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DappLinkVRFManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DappLinkVRFManagerOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DappLinkVRFManagerOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DappLinkVRFManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DappLinkVRFManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DappLinkVRFManagerOwnershipTransferred represents a OwnershipTransferred event raised by the DappLinkVRFManager contract.
type DappLinkVRFManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DappLinkVRFManager *DappLinkVRFManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DappLinkVRFManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DappLinkVRFManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DappLinkVRFManagerOwnershipTransferredIterator{contract: _DappLinkVRFManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DappLinkVRFManager *DappLinkVRFManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DappLinkVRFManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DappLinkVRFManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DappLinkVRFManagerOwnershipTransferred)
				if err := _DappLinkVRFManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DappLinkVRFManager *DappLinkVRFManagerFilterer) ParseOwnershipTransferred(log types.Log) (*DappLinkVRFManagerOwnershipTransferred, error) {
	event := new(DappLinkVRFManagerOwnershipTransferred)
	if err := _DappLinkVRFManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DappLinkVRFManagerRequestSentIterator is returned from FilterRequestSent and is used to iterate over the raw logs and unpacked data for RequestSent events raised by the DappLinkVRFManager contract.
type DappLinkVRFManagerRequestSentIterator struct {
	Event *DappLinkVRFManagerRequestSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DappLinkVRFManagerRequestSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DappLinkVRFManagerRequestSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DappLinkVRFManagerRequestSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DappLinkVRFManagerRequestSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DappLinkVRFManagerRequestSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DappLinkVRFManagerRequestSent represents a RequestSent event raised by the DappLinkVRFManager contract.
type DappLinkVRFManagerRequestSent struct {
	RequestId *big.Int
	NumWords  *big.Int
	Current   common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRequestSent is a free log retrieval operation binding the contract event 0xe697eb68c0228bd7d4e553246a2a86e8402d0895e45092ef8ae87b4cfd29f016.
//
// Solidity: event RequestSent(uint256 requestId, uint256 _numWords, address current)
func (_DappLinkVRFManager *DappLinkVRFManagerFilterer) FilterRequestSent(opts *bind.FilterOpts) (*DappLinkVRFManagerRequestSentIterator, error) {

	logs, sub, err := _DappLinkVRFManager.contract.FilterLogs(opts, "RequestSent")
	if err != nil {
		return nil, err
	}
	return &DappLinkVRFManagerRequestSentIterator{contract: _DappLinkVRFManager.contract, event: "RequestSent", logs: logs, sub: sub}, nil
}

// WatchRequestSent is a free log subscription operation binding the contract event 0xe697eb68c0228bd7d4e553246a2a86e8402d0895e45092ef8ae87b4cfd29f016.
//
// Solidity: event RequestSent(uint256 requestId, uint256 _numWords, address current)
func (_DappLinkVRFManager *DappLinkVRFManagerFilterer) WatchRequestSent(opts *bind.WatchOpts, sink chan<- *DappLinkVRFManagerRequestSent) (event.Subscription, error) {

	logs, sub, err := _DappLinkVRFManager.contract.WatchLogs(opts, "RequestSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DappLinkVRFManagerRequestSent)
				if err := _DappLinkVRFManager.contract.UnpackLog(event, "RequestSent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRequestSent is a log parse operation binding the contract event 0xe697eb68c0228bd7d4e553246a2a86e8402d0895e45092ef8ae87b4cfd29f016.
//
// Solidity: event RequestSent(uint256 requestId, uint256 _numWords, address current)
func (_DappLinkVRFManager *DappLinkVRFManagerFilterer) ParseRequestSent(log types.Log) (*DappLinkVRFManagerRequestSent, error) {
	event := new(DappLinkVRFManagerRequestSent)
	if err := _DappLinkVRFManager.contract.UnpackLog(event, "RequestSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
