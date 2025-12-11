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

// DappLinkVRFManagerMetaData contains all meta data concerning the DappLinkVRFManager contract.
var DappLinkVRFManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"blsRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"dappLinkAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fulfillRandomWords\",\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_randomWords\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getRequestStatus\",\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"fulfilled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"randomWords\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_dappLinkAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_blsRegistry\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lastRequestId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestIds\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requestMapping\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"fulfilled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requestRandomWords\",\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_numWords\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDappLink\",\"inputs\":[{\"name\":\"_dappLinkAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"FillRandomWords\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"randomWords\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RequestSent\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_numWords\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"current\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608060405234801561000f575f5ffd5b5061001e61002360201b60201c565b61019e565b5f61003261012160201b60201c565b9050805f0160089054906101000a900460ff161561007c576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8016815f015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff161461011e5767ffffffffffffffff815f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d267ffffffffffffffff6040516101159190610185565b60405180910390a15b50565b5f5f61013161013a60201b60201c565b90508091505090565b5f7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005f1b905090565b5f67ffffffffffffffff82169050919050565b61017f81610163565b82525050565b5f6020820190506101985f830184610176565b92915050565b61129d806101ab5f395ff3fe608060405234801561000f575f5ffd5b50600436106100cd575f3560e01c806393e4b8801161008a578063d8a4676f11610064578063d8a4676f146101e7578063f0c28a4114610218578063f2fde38b14610236578063fc2a88c314610252576100cd565b806393e4b88014610191578063996869d0146101af578063c0c53b8b146101cb576100cd565b80631b739ef1146100d157806338ba4614146100ed578063715018a61461010957806382e215ab146101135780638796ba8c146101435780638da5cb5b14610173575b5f5ffd5b6100eb60048036038101906100e69190610c63565b610270565b005b61010760048036038101906101029190610df1565b610390565b005b6101116104c2565b005b61012d60048036038101906101289190610e4b565b6104d5565b60405161013a9190610e90565b60405180910390f35b61015d60048036038101906101589190610e4b565b6104fb565b60405161016a9190610eb8565b60405180910390f35b61017b61051a565b6040516101889190610f10565b60405180910390f35b61019961054f565b6040516101a69190610f84565b60405180910390f35b6101c960048036038101906101c49190610fc7565b610574565b005b6101e560048036038101906101e09190610ff2565b6105bf565b005b61020160048036038101906101fc9190610e4b565b6107c2565b60405161020f9291906110f9565b60405180910390f35b610220610852565b60405161022d9190610f10565b60405180910390f35b610250600480360381019061024b9190610fc7565b610877565b005b61025a6108fb565b6040516102679190610eb8565b60405180910390f35b610278610901565b60405180604001604052805f151581526020015f67ffffffffffffffff8111156102a5576102a4610cb5565b5b6040519080825280602002602001820160405280156102d35781602001602082028036833780820191505090505b5081525060045f8481526020019081526020015f205f820151815f015f6101000a81548160ff0219169083151502179055506020820151816001019080519060200190610321929190610bb9565b509050505f82908060018154018082558091505060019003905f5260205f20015f9091909190915055816001819055507fe697eb68c0228bd7d4e553246a2a86e8402d0895e45092ef8ae87b4cfd29f01682823060405161038493929190611127565b60405180910390a15050565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461041f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610416906111b6565b60405180910390fd5b60405180604001604052806001151581526020018281525060045f8481526020019081526020015f205f820151815f015f6101000a81548160ff0219169083151502179055506020820151816001019080519060200190610481929190610bb9565b509050507ff3cb4deb0441dd096356debf166f879d78cadc19e4b94053c8bea6d3940de93a82826040516104b69291906111d4565b60405180910390a15050565b6104ca610901565b6104d35f610988565b565b6004602052805f5260405f205f91509050805f015f9054906101000a900460ff16905081565b5f8181548110610509575f80fd5b905f5260205f20015f915090505481565b5f5f610524610a59565b9050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1691505090565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b61057c610901565b8060025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f6105c8610a80565b90505f815f0160089054906101000a900460ff161590505f825f015f9054906101000a900467ffffffffffffffff1690505f5f8267ffffffffffffffff161480156106105750825b90505f60018367ffffffffffffffff1614801561064357505f3073ffffffffffffffffffffffffffffffffffffffff163b145b905081158015610651575080155b15610688576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001855f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555083156106d5576001855f0160086101000a81548160ff0219169083151502179055505b6106de88610a93565b8560035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508660025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083156107b8575f855f0160086101000a81548160ff0219169083151502179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d260016040516107af919061124e565b60405180910390a15b5050505050505050565b5f606060045f8481526020019081526020015f205f015f9054906101000a900460ff1660045f8581526020019081526020015f206001018080548060200260200160405190810160405280929190818152602001828054801561084257602002820191905f5260205f20905b81548152602001906001019080831161082e575b5050505050905091509150915091565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b61087f610901565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036108ef575f6040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016108e69190610f10565b60405180910390fd5b6108f881610988565b50565b60015481565b610909610aa7565b73ffffffffffffffffffffffffffffffffffffffff1661092761051a565b73ffffffffffffffffffffffffffffffffffffffff16146109865761094a610aa7565b6040517f118cdaa700000000000000000000000000000000000000000000000000000000815260040161097d9190610f10565b60405180910390fd5b565b5f610991610a59565b90505f815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905082825f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3505050565b5f7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300905090565b5f5f610a8a610aae565b90508091505090565b610a9b610ad7565b610aa481610b17565b50565b5f33905090565b5f7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005f1b905090565b610adf610b9b565b610b15576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b610b1f610ad7565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610b8f575f6040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401610b869190610f10565b60405180910390fd5b610b9881610988565b50565b5f610ba4610a80565b5f0160089054906101000a900460ff16905090565b828054828255905f5260205f20908101928215610bf3579160200282015b82811115610bf2578251825591602001919060010190610bd7565b5b509050610c009190610c04565b5090565b5b80821115610c1b575f815f905550600101610c05565b5090565b5f604051905090565b5f5ffd5b5f5ffd5b5f819050919050565b610c4281610c30565b8114610c4c575f5ffd5b50565b5f81359050610c5d81610c39565b92915050565b5f5f60408385031215610c7957610c78610c28565b5b5f610c8685828601610c4f565b9250506020610c9785828601610c4f565b9150509250929050565b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610ceb82610ca5565b810181811067ffffffffffffffff82111715610d0a57610d09610cb5565b5b80604052505050565b5f610d1c610c1f565b9050610d288282610ce2565b919050565b5f67ffffffffffffffff821115610d4757610d46610cb5565b5b602082029050602081019050919050565b5f5ffd5b5f610d6e610d6984610d2d565b610d13565b90508083825260208201905060208402830185811115610d9157610d90610d58565b5b835b81811015610dba5780610da68882610c4f565b845260208401935050602081019050610d93565b5050509392505050565b5f82601f830112610dd857610dd7610ca1565b5b8135610de8848260208601610d5c565b91505092915050565b5f5f60408385031215610e0757610e06610c28565b5b5f610e1485828601610c4f565b925050602083013567ffffffffffffffff811115610e3557610e34610c2c565b5b610e4185828601610dc4565b9150509250929050565b5f60208284031215610e6057610e5f610c28565b5b5f610e6d84828501610c4f565b91505092915050565b5f8115159050919050565b610e8a81610e76565b82525050565b5f602082019050610ea35f830184610e81565b92915050565b610eb281610c30565b82525050565b5f602082019050610ecb5f830184610ea9565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610efa82610ed1565b9050919050565b610f0a81610ef0565b82525050565b5f602082019050610f235f830184610f01565b92915050565b5f819050919050565b5f610f4c610f47610f4284610ed1565b610f29565b610ed1565b9050919050565b5f610f5d82610f32565b9050919050565b5f610f6e82610f53565b9050919050565b610f7e81610f64565b82525050565b5f602082019050610f975f830184610f75565b92915050565b610fa681610ef0565b8114610fb0575f5ffd5b50565b5f81359050610fc181610f9d565b92915050565b5f60208284031215610fdc57610fdb610c28565b5b5f610fe984828501610fb3565b91505092915050565b5f5f5f6060848603121561100957611008610c28565b5b5f61101686828701610fb3565b935050602061102786828701610fb3565b925050604061103886828701610fb3565b9150509250925092565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b61107481610c30565b82525050565b5f611085838361106b565b60208301905092915050565b5f602082019050919050565b5f6110a782611042565b6110b1818561104c565b93506110bc8361105c565b805f5b838110156110ec5781516110d3888261107a565b97506110de83611091565b9250506001810190506110bf565b5085935050505092915050565b5f60408201905061110c5f830185610e81565b818103602083015261111e818461109d565b90509392505050565b5f60608201905061113a5f830186610ea9565b6111476020830185610ea9565b6111546040830184610f01565b949350505050565b5f82825260208201905092915050565b7f446170704c696e6b5652462e6f6e6c79446170704c696e6b00000000000000005f82015250565b5f6111a060188361115c565b91506111ab8261116c565b602082019050919050565b5f6020820190508181035f8301526111cd81611194565b9050919050565b5f6040820190506111e75f830185610ea9565b81810360208301526111f9818461109d565b90509392505050565b5f819050919050565b5f67ffffffffffffffff82169050919050565b5f61123861123361122e84611202565b610f29565b61120b565b9050919050565b6112488161121e565b82525050565b5f6020820190506112615f83018461123f565b9291505056fea26469706673582212204b07eada964248f760fa4e0838e02f0a39094ddaa70ba1b5ff51d90b6f65182664736f6c634300081e0033",
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

// FulfillRandomWords is a paid mutator transaction binding the contract method 0x38ba4614.
//
// Solidity: function fulfillRandomWords(uint256 _requestId, uint256[] _randomWords) returns()
func (_DappLinkVRFManager *DappLinkVRFManagerTransactor) FulfillRandomWords(opts *bind.TransactOpts, _requestId *big.Int, _randomWords []*big.Int) (*types.Transaction, error) {
	return _DappLinkVRFManager.contract.Transact(opts, "fulfillRandomWords", _requestId, _randomWords)
}

// FulfillRandomWords is a paid mutator transaction binding the contract method 0x38ba4614.
//
// Solidity: function fulfillRandomWords(uint256 _requestId, uint256[] _randomWords) returns()
func (_DappLinkVRFManager *DappLinkVRFManagerSession) FulfillRandomWords(_requestId *big.Int, _randomWords []*big.Int) (*types.Transaction, error) {
	return _DappLinkVRFManager.Contract.FulfillRandomWords(&_DappLinkVRFManager.TransactOpts, _requestId, _randomWords)
}

// FulfillRandomWords is a paid mutator transaction binding the contract method 0x38ba4614.
//
// Solidity: function fulfillRandomWords(uint256 _requestId, uint256[] _randomWords) returns()
func (_DappLinkVRFManager *DappLinkVRFManagerTransactorSession) FulfillRandomWords(_requestId *big.Int, _randomWords []*big.Int) (*types.Transaction, error) {
	return _DappLinkVRFManager.Contract.FulfillRandomWords(&_DappLinkVRFManager.TransactOpts, _requestId, _randomWords)
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
