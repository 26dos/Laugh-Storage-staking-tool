// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// DcPayChannelStakeInfo is an auto generated low-level Go binding around an user-defined struct.
type DcPayChannelStakeInfo struct {
	StakeId   *big.Int
	Amount    *big.Int
	StartTime *big.Int
	LockTime  *big.Int
	Active    bool
}

// DcPayChannelMetaData contains all meta data concerning the DcPayChannel contract.
var DcPayChannelMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakeId\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakeId\",\"type\":\"uint256\"}],\"name\":\"Unstaked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FIL_PER_10T\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LOCK_TIME_20_DAYS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LOCK_TIME_40_DAYS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_DC\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getAllStakeInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stakeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"internalType\":\"structDcPayChannel.StakeInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getStakeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeType\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeId\",\"type\":\"uint256\"}],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561001057600080fd5b506108d6806100206000396000f3fe6080604052600436106100c65760003560e01c8063a69df4b51161007f578063d1d7e10a11610059578063d1d7e10a14610272578063e30c39781461028e578063f158adf2146102ac578063f2fde38b146102c357600080fd5b8063a69df4b5146101d4578063bb224d6e146101e9578063c3453153146101ff57600080fd5b806316934fc4146100d257806361cbfa501461013e578063715018a61461016357806379ba50971461017a5780637b0472f01461018f5780638da5cb5b146101a257600080fd5b366100cd57005b600080fd5b3480156100de57600080fd5b506101176100ed3660046107df565b60976020526000908152604090208054600182015460028301546003909301549192909160ff1684565b60408051948552602085019390935291830152151560608201526080015b60405180910390f35b34801561014a57600080fd5b506101556234bc0081565b604051908152602001610135565b34801561016f57600080fd5b506101786102e3565b005b34801561018657600080fd5b506101786102f7565b61017861019d36600461080f565b610376565b3480156101ae57600080fd5b506033546001600160a01b03165b6040516001600160a01b039091168152602001610135565b3480156101e057600080fd5b5061017861052c565b3480156101f557600080fd5b5061015561080081565b34801561020b57600080fd5b5061011761021a3660046107df565b6001600160a01b0316600090815260976020908152604091829020825160808101845281548082526001830154938201849052600283015494820185905260039092015460ff16151560609091018190529093919291565b34801561027e57600080fd5b50610155670de0b6b3a764000081565b34801561029a57600080fd5b506065546001600160a01b03166101bc565b3480156102b857600080fd5b50610155621a5e0081565b3480156102cf57600080fd5b506101786102de3660046107df565b6106ae565b6102eb61071f565b6102f56000610779565b565b60655433906001600160a01b0316811461036a5760405162461bcd60e51b815260206004820152602960248201527f4f776e61626c6532537465703a2063616c6c6572206973206e6f7420746865206044820152683732bb9037bbb732b960b91b60648201526084015b60405180910390fd5b61037381610779565b50565b8015806103835750806001145b6103c45760405162461bcd60e51b8152602060048201526012602482015271496e76616c6964207374616b65207479706560701b6044820152606401610361565b6108008211156104075760405162461bcd60e51b815260206004820152600e60248201526d45786365656473206d617820444360901b6044820152606401610361565b60008115610418576234bc0061041d565b621a5e005b90506000670de0b6b3a7640000610435600a86610847565b61043f9190610869565b9050803410156104915760405162461bcd60e51b815260206004820152601c60248201527f496e73756666696369656e742046494c20666f72207374616b696e67000000006044820152606401610361565b6040805160808101825234808252426020808401918252838501878152600160608601818152336000818152609786528990209751885594519187019190915590516002860155516003909401805460ff191694151594909417909355835191825291810185905290917f1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee90910160405180910390a250505050565b336000908152609760205260409020600381015460ff166105815760405162461bcd60e51b815260206004820152600f60248201526e4e6f20616374697665207374616b6560881b6044820152606401610361565b806002015481600101546105959190610888565b4210156105dc5760405162461bcd60e51b815260206004820152601560248201527414dd185ad9481a5cc81cdd1a5b1b081b1bd8dad959605a1b6044820152606401610361565b805460038201805460ff19169055604051600090339083908381818185875af1925050503d806000811461062c576040519150601f19603f3d011682016040523d82523d6000602084013e610631565b606091505b50509050806106745760405162461bcd60e51b815260206004820152600f60248201526e151c985b9cd9995c8819985a5b1959608a1b6044820152606401610361565b60405182815233907f0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f759060200160405180910390a2505050565b6106b661071f565b606580546001600160a01b0383166001600160a01b031990911681179091556106e76033546001600160a01b031690565b6001600160a01b03167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e2270060405160405180910390a350565b6033546001600160a01b031633146102f55760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610361565b606580546001600160a01b031916905561037381603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000602082840312156107f157600080fd5b81356001600160a01b038116811461080857600080fd5b9392505050565b6000806040838503121561082257600080fd5b50508035926020909101359150565b634e487b7160e01b600052601160045260246000fd5b60008261086457634e487b7160e01b600052601260045260246000fd5b500490565b600081600019048311821515161561088357610883610831565b500290565b6000821982111561089b5761089b610831565b50019056fea2646970667358221220344aea29010ffc0505176c4964266462d824338dbbc0a2d78c7e50abe2af9c0a64736f6c63430008090033",
}

// DcPayChannelABI is the input ABI used to generate the binding from.
// Deprecated: Use DcPayChannelMetaData.ABI instead.
var DcPayChannelABI = DcPayChannelMetaData.ABI

// DcPayChannelBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DcPayChannelMetaData.Bin instead.
var DcPayChannelBin = DcPayChannelMetaData.Bin

// DeployDcPayChannel deploys a new Ethereum contract, binding an instance of DcPayChannel to it.
func DeployDcPayChannel(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DcPayChannel, error) {
	parsed, err := DcPayChannelMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DcPayChannelBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DcPayChannel{DcPayChannelCaller: DcPayChannelCaller{contract: contract}, DcPayChannelTransactor: DcPayChannelTransactor{contract: contract}, DcPayChannelFilterer: DcPayChannelFilterer{contract: contract}}, nil
}

// DcPayChannel is an auto generated Go binding around an Ethereum contract.
type DcPayChannel struct {
	DcPayChannelCaller     // Read-only binding to the contract
	DcPayChannelTransactor // Write-only binding to the contract
	DcPayChannelFilterer   // Log filterer for contract events
}

// DcPayChannelCaller is an auto generated read-only Go binding around an Ethereum contract.
type DcPayChannelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DcPayChannelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DcPayChannelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DcPayChannelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DcPayChannelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DcPayChannelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DcPayChannelSession struct {
	Contract     *DcPayChannel     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DcPayChannelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DcPayChannelCallerSession struct {
	Contract *DcPayChannelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// DcPayChannelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DcPayChannelTransactorSession struct {
	Contract     *DcPayChannelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// DcPayChannelRaw is an auto generated low-level Go binding around an Ethereum contract.
type DcPayChannelRaw struct {
	Contract *DcPayChannel // Generic contract binding to access the raw methods on
}

// DcPayChannelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DcPayChannelCallerRaw struct {
	Contract *DcPayChannelCaller // Generic read-only contract binding to access the raw methods on
}

// DcPayChannelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DcPayChannelTransactorRaw struct {
	Contract *DcPayChannelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDcPayChannel creates a new instance of DcPayChannel, bound to a specific deployed contract.
func NewDcPayChannel(address common.Address, backend bind.ContractBackend) (*DcPayChannel, error) {
	contract, err := bindDcPayChannel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DcPayChannel{DcPayChannelCaller: DcPayChannelCaller{contract: contract}, DcPayChannelTransactor: DcPayChannelTransactor{contract: contract}, DcPayChannelFilterer: DcPayChannelFilterer{contract: contract}}, nil
}

// NewDcPayChannelCaller creates a new read-only instance of DcPayChannel, bound to a specific deployed contract.
func NewDcPayChannelCaller(address common.Address, caller bind.ContractCaller) (*DcPayChannelCaller, error) {
	contract, err := bindDcPayChannel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DcPayChannelCaller{contract: contract}, nil
}

// NewDcPayChannelTransactor creates a new write-only instance of DcPayChannel, bound to a specific deployed contract.
func NewDcPayChannelTransactor(address common.Address, transactor bind.ContractTransactor) (*DcPayChannelTransactor, error) {
	contract, err := bindDcPayChannel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DcPayChannelTransactor{contract: contract}, nil
}

// NewDcPayChannelFilterer creates a new log filterer instance of DcPayChannel, bound to a specific deployed contract.
func NewDcPayChannelFilterer(address common.Address, filterer bind.ContractFilterer) (*DcPayChannelFilterer, error) {
	contract, err := bindDcPayChannel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DcPayChannelFilterer{contract: contract}, nil
}

// bindDcPayChannel binds a generic wrapper to an already deployed contract.
func bindDcPayChannel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DcPayChannelMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DcPayChannel *DcPayChannelRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DcPayChannel.Contract.DcPayChannelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DcPayChannel *DcPayChannelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DcPayChannel.Contract.DcPayChannelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DcPayChannel *DcPayChannelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DcPayChannel.Contract.DcPayChannelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DcPayChannel *DcPayChannelCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DcPayChannel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DcPayChannel *DcPayChannelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DcPayChannel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DcPayChannel *DcPayChannelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DcPayChannel.Contract.contract.Transact(opts, method, params...)
}

// FILPER10T is a free data retrieval call binding the contract method 0xd1d7e10a.
//
// Solidity: function FIL_PER_10T() view returns(uint256)
func (_DcPayChannel *DcPayChannelCaller) FILPER10T(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DcPayChannel.contract.Call(opts, &out, "FIL_PER_10T")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FILPER10T is a free data retrieval call binding the contract method 0xd1d7e10a.
//
// Solidity: function FIL_PER_10T() view returns(uint256)
func (_DcPayChannel *DcPayChannelSession) FILPER10T() (*big.Int, error) {
	return _DcPayChannel.Contract.FILPER10T(&_DcPayChannel.CallOpts)
}

// FILPER10T is a free data retrieval call binding the contract method 0xd1d7e10a.
//
// Solidity: function FIL_PER_10T() view returns(uint256)
func (_DcPayChannel *DcPayChannelCallerSession) FILPER10T() (*big.Int, error) {
	return _DcPayChannel.Contract.FILPER10T(&_DcPayChannel.CallOpts)
}

// LOCKTIME20DAYS is a free data retrieval call binding the contract method 0xf158adf2.
//
// Solidity: function LOCK_TIME_20_DAYS() view returns(uint256)
func (_DcPayChannel *DcPayChannelCaller) LOCKTIME20DAYS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DcPayChannel.contract.Call(opts, &out, "LOCK_TIME_20_DAYS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LOCKTIME20DAYS is a free data retrieval call binding the contract method 0xf158adf2.
//
// Solidity: function LOCK_TIME_20_DAYS() view returns(uint256)
func (_DcPayChannel *DcPayChannelSession) LOCKTIME20DAYS() (*big.Int, error) {
	return _DcPayChannel.Contract.LOCKTIME20DAYS(&_DcPayChannel.CallOpts)
}

// LOCKTIME20DAYS is a free data retrieval call binding the contract method 0xf158adf2.
//
// Solidity: function LOCK_TIME_20_DAYS() view returns(uint256)
func (_DcPayChannel *DcPayChannelCallerSession) LOCKTIME20DAYS() (*big.Int, error) {
	return _DcPayChannel.Contract.LOCKTIME20DAYS(&_DcPayChannel.CallOpts)
}

// LOCKTIME40DAYS is a free data retrieval call binding the contract method 0x61cbfa50.
//
// Solidity: function LOCK_TIME_40_DAYS() view returns(uint256)
func (_DcPayChannel *DcPayChannelCaller) LOCKTIME40DAYS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DcPayChannel.contract.Call(opts, &out, "LOCK_TIME_40_DAYS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LOCKTIME40DAYS is a free data retrieval call binding the contract method 0x61cbfa50.
//
// Solidity: function LOCK_TIME_40_DAYS() view returns(uint256)
func (_DcPayChannel *DcPayChannelSession) LOCKTIME40DAYS() (*big.Int, error) {
	return _DcPayChannel.Contract.LOCKTIME40DAYS(&_DcPayChannel.CallOpts)
}

// LOCKTIME40DAYS is a free data retrieval call binding the contract method 0x61cbfa50.
//
// Solidity: function LOCK_TIME_40_DAYS() view returns(uint256)
func (_DcPayChannel *DcPayChannelCallerSession) LOCKTIME40DAYS() (*big.Int, error) {
	return _DcPayChannel.Contract.LOCKTIME40DAYS(&_DcPayChannel.CallOpts)
}

// MAXDC is a free data retrieval call binding the contract method 0xbb224d6e.
//
// Solidity: function MAX_DC() view returns(uint256)
func (_DcPayChannel *DcPayChannelCaller) MAXDC(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DcPayChannel.contract.Call(opts, &out, "MAX_DC")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXDC is a free data retrieval call binding the contract method 0xbb224d6e.
//
// Solidity: function MAX_DC() view returns(uint256)
func (_DcPayChannel *DcPayChannelSession) MAXDC() (*big.Int, error) {
	return _DcPayChannel.Contract.MAXDC(&_DcPayChannel.CallOpts)
}

// MAXDC is a free data retrieval call binding the contract method 0xbb224d6e.
//
// Solidity: function MAX_DC() view returns(uint256)
func (_DcPayChannel *DcPayChannelCallerSession) MAXDC() (*big.Int, error) {
	return _DcPayChannel.Contract.MAXDC(&_DcPayChannel.CallOpts)
}

// GetAllStakeInfo is a free data retrieval call binding the contract method 0x559566b9.
//
// Solidity: function getAllStakeInfo(address user) view returns((uint256,uint256,uint256,uint256,bool)[])
func (_DcPayChannel *DcPayChannelCaller) GetAllStakeInfo(opts *bind.CallOpts, user common.Address) ([]DcPayChannelStakeInfo, error) {
	var out []interface{}
	err := _DcPayChannel.contract.Call(opts, &out, "getAllStakeInfo", user)

	if err != nil {
		return *new([]DcPayChannelStakeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]DcPayChannelStakeInfo)).(*[]DcPayChannelStakeInfo)

	return out0, err

}

// GetAllStakeInfo is a free data retrieval call binding the contract method 0x559566b9.
//
// Solidity: function getAllStakeInfo(address user) view returns((uint256,uint256,uint256,uint256,bool)[])
func (_DcPayChannel *DcPayChannelSession) GetAllStakeInfo(user common.Address) ([]DcPayChannelStakeInfo, error) {
	return _DcPayChannel.Contract.GetAllStakeInfo(&_DcPayChannel.CallOpts, user)
}

// GetAllStakeInfo is a free data retrieval call binding the contract method 0x559566b9.
//
// Solidity: function getAllStakeInfo(address user) view returns((uint256,uint256,uint256,uint256,bool)[])
func (_DcPayChannel *DcPayChannelCallerSession) GetAllStakeInfo(user common.Address) ([]DcPayChannelStakeInfo, error) {
	return _DcPayChannel.Contract.GetAllStakeInfo(&_DcPayChannel.CallOpts, user)
}

// GetStakeCount is a free data retrieval call binding the contract method 0xcf57ee69.
//
// Solidity: function getStakeCount(address user) view returns(uint256)
func (_DcPayChannel *DcPayChannelCaller) GetStakeCount(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DcPayChannel.contract.Call(opts, &out, "getStakeCount", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakeCount is a free data retrieval call binding the contract method 0xcf57ee69.
//
// Solidity: function getStakeCount(address user) view returns(uint256)
func (_DcPayChannel *DcPayChannelSession) GetStakeCount(user common.Address) (*big.Int, error) {
	return _DcPayChannel.Contract.GetStakeCount(&_DcPayChannel.CallOpts, user)
}

// GetStakeCount is a free data retrieval call binding the contract method 0xcf57ee69.
//
// Solidity: function getStakeCount(address user) view returns(uint256)
func (_DcPayChannel *DcPayChannelCallerSession) GetStakeCount(user common.Address) (*big.Int, error) {
	return _DcPayChannel.Contract.GetStakeCount(&_DcPayChannel.CallOpts, user)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DcPayChannel *DcPayChannelCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DcPayChannel.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DcPayChannel *DcPayChannelSession) Owner() (common.Address, error) {
	return _DcPayChannel.Contract.Owner(&_DcPayChannel.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DcPayChannel *DcPayChannelCallerSession) Owner() (common.Address, error) {
	return _DcPayChannel.Contract.Owner(&_DcPayChannel.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_DcPayChannel *DcPayChannelCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DcPayChannel.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_DcPayChannel *DcPayChannelSession) PendingOwner() (common.Address, error) {
	return _DcPayChannel.Contract.PendingOwner(&_DcPayChannel.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_DcPayChannel *DcPayChannelCallerSession) PendingOwner() (common.Address, error) {
	return _DcPayChannel.Contract.PendingOwner(&_DcPayChannel.CallOpts)
}

// Stakes is a free data retrieval call binding the contract method 0x584b62a1.
//
// Solidity: function stakes(address , uint256 ) view returns(uint256 amount, uint256 startTime, uint256 lockTime, bool active)
func (_DcPayChannel *DcPayChannelCaller) Stakes(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Amount    *big.Int
	StartTime *big.Int
	LockTime  *big.Int
	Active    bool
}, error) {
	var out []interface{}
	err := _DcPayChannel.contract.Call(opts, &out, "stakes", arg0, arg1)

	outstruct := new(struct {
		Amount    *big.Int
		StartTime *big.Int
		LockTime  *big.Int
		Active    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.StartTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LockTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Active = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// Stakes is a free data retrieval call binding the contract method 0x584b62a1.
//
// Solidity: function stakes(address , uint256 ) view returns(uint256 amount, uint256 startTime, uint256 lockTime, bool active)
func (_DcPayChannel *DcPayChannelSession) Stakes(arg0 common.Address, arg1 *big.Int) (struct {
	Amount    *big.Int
	StartTime *big.Int
	LockTime  *big.Int
	Active    bool
}, error) {
	return _DcPayChannel.Contract.Stakes(&_DcPayChannel.CallOpts, arg0, arg1)
}

// Stakes is a free data retrieval call binding the contract method 0x584b62a1.
//
// Solidity: function stakes(address , uint256 ) view returns(uint256 amount, uint256 startTime, uint256 lockTime, bool active)
func (_DcPayChannel *DcPayChannelCallerSession) Stakes(arg0 common.Address, arg1 *big.Int) (struct {
	Amount    *big.Int
	StartTime *big.Int
	LockTime  *big.Int
	Active    bool
}, error) {
	return _DcPayChannel.Contract.Stakes(&_DcPayChannel.CallOpts, arg0, arg1)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_DcPayChannel *DcPayChannelTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DcPayChannel.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_DcPayChannel *DcPayChannelSession) AcceptOwnership() (*types.Transaction, error) {
	return _DcPayChannel.Contract.AcceptOwnership(&_DcPayChannel.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_DcPayChannel *DcPayChannelTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _DcPayChannel.Contract.AcceptOwnership(&_DcPayChannel.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DcPayChannel *DcPayChannelTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DcPayChannel.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DcPayChannel *DcPayChannelSession) RenounceOwnership() (*types.Transaction, error) {
	return _DcPayChannel.Contract.RenounceOwnership(&_DcPayChannel.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DcPayChannel *DcPayChannelTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DcPayChannel.Contract.RenounceOwnership(&_DcPayChannel.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0x7b0472f0.
//
// Solidity: function stake(uint256 dcAmount, uint256 stakeType) payable returns()
func (_DcPayChannel *DcPayChannelTransactor) Stake(opts *bind.TransactOpts, dcAmount *big.Int, stakeType *big.Int) (*types.Transaction, error) {
	return _DcPayChannel.contract.Transact(opts, "stake", dcAmount, stakeType)
}

// Stake is a paid mutator transaction binding the contract method 0x7b0472f0.
//
// Solidity: function stake(uint256 dcAmount, uint256 stakeType) payable returns()
func (_DcPayChannel *DcPayChannelSession) Stake(dcAmount *big.Int, stakeType *big.Int) (*types.Transaction, error) {
	return _DcPayChannel.Contract.Stake(&_DcPayChannel.TransactOpts, dcAmount, stakeType)
}

// Stake is a paid mutator transaction binding the contract method 0x7b0472f0.
//
// Solidity: function stake(uint256 dcAmount, uint256 stakeType) payable returns()
func (_DcPayChannel *DcPayChannelTransactorSession) Stake(dcAmount *big.Int, stakeType *big.Int) (*types.Transaction, error) {
	return _DcPayChannel.Contract.Stake(&_DcPayChannel.TransactOpts, dcAmount, stakeType)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DcPayChannel *DcPayChannelTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DcPayChannel.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DcPayChannel *DcPayChannelSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DcPayChannel.Contract.TransferOwnership(&_DcPayChannel.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DcPayChannel *DcPayChannelTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DcPayChannel.Contract.TransferOwnership(&_DcPayChannel.TransactOpts, newOwner)
}

// Unlock is a paid mutator transaction binding the contract method 0x6198e339.
//
// Solidity: function unlock(uint256 stakeId) returns()
func (_DcPayChannel *DcPayChannelTransactor) Unlock(opts *bind.TransactOpts, stakeId *big.Int) (*types.Transaction, error) {
	return _DcPayChannel.contract.Transact(opts, "unlock", stakeId)
}

// Unlock is a paid mutator transaction binding the contract method 0x6198e339.
//
// Solidity: function unlock(uint256 stakeId) returns()
func (_DcPayChannel *DcPayChannelSession) Unlock(stakeId *big.Int) (*types.Transaction, error) {
	return _DcPayChannel.Contract.Unlock(&_DcPayChannel.TransactOpts, stakeId)
}

// Unlock is a paid mutator transaction binding the contract method 0x6198e339.
//
// Solidity: function unlock(uint256 stakeId) returns()
func (_DcPayChannel *DcPayChannelTransactorSession) Unlock(stakeId *big.Int) (*types.Transaction, error) {
	return _DcPayChannel.Contract.Unlock(&_DcPayChannel.TransactOpts, stakeId)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DcPayChannel *DcPayChannelTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DcPayChannel.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DcPayChannel *DcPayChannelSession) Receive() (*types.Transaction, error) {
	return _DcPayChannel.Contract.Receive(&_DcPayChannel.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DcPayChannel *DcPayChannelTransactorSession) Receive() (*types.Transaction, error) {
	return _DcPayChannel.Contract.Receive(&_DcPayChannel.TransactOpts)
}

// DcPayChannelInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the DcPayChannel contract.
type DcPayChannelInitializedIterator struct {
	Event *DcPayChannelInitialized // Event containing the contract specifics and raw log

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
func (it *DcPayChannelInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DcPayChannelInitialized)
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
		it.Event = new(DcPayChannelInitialized)
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
func (it *DcPayChannelInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DcPayChannelInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DcPayChannelInitialized represents a Initialized event raised by the DcPayChannel contract.
type DcPayChannelInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DcPayChannel *DcPayChannelFilterer) FilterInitialized(opts *bind.FilterOpts) (*DcPayChannelInitializedIterator, error) {

	logs, sub, err := _DcPayChannel.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DcPayChannelInitializedIterator{contract: _DcPayChannel.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DcPayChannel *DcPayChannelFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DcPayChannelInitialized) (event.Subscription, error) {

	logs, sub, err := _DcPayChannel.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DcPayChannelInitialized)
				if err := _DcPayChannel.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DcPayChannel *DcPayChannelFilterer) ParseInitialized(log types.Log) (*DcPayChannelInitialized, error) {
	event := new(DcPayChannelInitialized)
	if err := _DcPayChannel.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DcPayChannelOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the DcPayChannel contract.
type DcPayChannelOwnershipTransferStartedIterator struct {
	Event *DcPayChannelOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *DcPayChannelOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DcPayChannelOwnershipTransferStarted)
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
		it.Event = new(DcPayChannelOwnershipTransferStarted)
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
func (it *DcPayChannelOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DcPayChannelOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DcPayChannelOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the DcPayChannel contract.
type DcPayChannelOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_DcPayChannel *DcPayChannelFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DcPayChannelOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DcPayChannel.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DcPayChannelOwnershipTransferStartedIterator{contract: _DcPayChannel.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_DcPayChannel *DcPayChannelFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *DcPayChannelOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DcPayChannel.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DcPayChannelOwnershipTransferStarted)
				if err := _DcPayChannel.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_DcPayChannel *DcPayChannelFilterer) ParseOwnershipTransferStarted(log types.Log) (*DcPayChannelOwnershipTransferStarted, error) {
	event := new(DcPayChannelOwnershipTransferStarted)
	if err := _DcPayChannel.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DcPayChannelOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DcPayChannel contract.
type DcPayChannelOwnershipTransferredIterator struct {
	Event *DcPayChannelOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DcPayChannelOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DcPayChannelOwnershipTransferred)
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
		it.Event = new(DcPayChannelOwnershipTransferred)
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
func (it *DcPayChannelOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DcPayChannelOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DcPayChannelOwnershipTransferred represents a OwnershipTransferred event raised by the DcPayChannel contract.
type DcPayChannelOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DcPayChannel *DcPayChannelFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DcPayChannelOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DcPayChannel.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DcPayChannelOwnershipTransferredIterator{contract: _DcPayChannel.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DcPayChannel *DcPayChannelFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DcPayChannelOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DcPayChannel.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DcPayChannelOwnershipTransferred)
				if err := _DcPayChannel.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_DcPayChannel *DcPayChannelFilterer) ParseOwnershipTransferred(log types.Log) (*DcPayChannelOwnershipTransferred, error) {
	event := new(DcPayChannelOwnershipTransferred)
	if err := _DcPayChannel.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DcPayChannelStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the DcPayChannel contract.
type DcPayChannelStakedIterator struct {
	Event *DcPayChannelStaked // Event containing the contract specifics and raw log

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
func (it *DcPayChannelStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DcPayChannelStaked)
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
		it.Event = new(DcPayChannelStaked)
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
func (it *DcPayChannelStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DcPayChannelStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DcPayChannelStaked represents a Staked event raised by the DcPayChannel contract.
type DcPayChannelStaked struct {
	User     common.Address
	Amount   *big.Int
	LockTime *big.Int
	StakeId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0xb4caaf29adda3eefee3ad552a8e85058589bf834c7466cae4ee58787f70589ed.
//
// Solidity: event Staked(address indexed user, uint256 amount, uint256 lockTime, uint256 stakeId)
func (_DcPayChannel *DcPayChannelFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address) (*DcPayChannelStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DcPayChannel.contract.FilterLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return &DcPayChannelStakedIterator{contract: _DcPayChannel.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0xb4caaf29adda3eefee3ad552a8e85058589bf834c7466cae4ee58787f70589ed.
//
// Solidity: event Staked(address indexed user, uint256 amount, uint256 lockTime, uint256 stakeId)
func (_DcPayChannel *DcPayChannelFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *DcPayChannelStaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DcPayChannel.contract.WatchLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DcPayChannelStaked)
				if err := _DcPayChannel.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0xb4caaf29adda3eefee3ad552a8e85058589bf834c7466cae4ee58787f70589ed.
//
// Solidity: event Staked(address indexed user, uint256 amount, uint256 lockTime, uint256 stakeId)
func (_DcPayChannel *DcPayChannelFilterer) ParseStaked(log types.Log) (*DcPayChannelStaked, error) {
	event := new(DcPayChannelStaked)
	if err := _DcPayChannel.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DcPayChannelUnstakedIterator is returned from FilterUnstaked and is used to iterate over the raw logs and unpacked data for Unstaked events raised by the DcPayChannel contract.
type DcPayChannelUnstakedIterator struct {
	Event *DcPayChannelUnstaked // Event containing the contract specifics and raw log

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
func (it *DcPayChannelUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DcPayChannelUnstaked)
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
		it.Event = new(DcPayChannelUnstaked)
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
func (it *DcPayChannelUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DcPayChannelUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DcPayChannelUnstaked represents a Unstaked event raised by the DcPayChannel contract.
type DcPayChannelUnstaked struct {
	User    common.Address
	Amount  *big.Int
	StakeId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnstaked is a free log retrieval operation binding the contract event 0x7fc4727e062e336010f2c282598ef5f14facb3de68cf8195c2f23e1454b2b74e.
//
// Solidity: event Unstaked(address indexed user, uint256 amount, uint256 stakeId)
func (_DcPayChannel *DcPayChannelFilterer) FilterUnstaked(opts *bind.FilterOpts, user []common.Address) (*DcPayChannelUnstakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DcPayChannel.contract.FilterLogs(opts, "Unstaked", userRule)
	if err != nil {
		return nil, err
	}
	return &DcPayChannelUnstakedIterator{contract: _DcPayChannel.contract, event: "Unstaked", logs: logs, sub: sub}, nil
}

// WatchUnstaked is a free log subscription operation binding the contract event 0x7fc4727e062e336010f2c282598ef5f14facb3de68cf8195c2f23e1454b2b74e.
//
// Solidity: event Unstaked(address indexed user, uint256 amount, uint256 stakeId)
func (_DcPayChannel *DcPayChannelFilterer) WatchUnstaked(opts *bind.WatchOpts, sink chan<- *DcPayChannelUnstaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DcPayChannel.contract.WatchLogs(opts, "Unstaked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DcPayChannelUnstaked)
				if err := _DcPayChannel.contract.UnpackLog(event, "Unstaked", log); err != nil {
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

// ParseUnstaked is a log parse operation binding the contract event 0x7fc4727e062e336010f2c282598ef5f14facb3de68cf8195c2f23e1454b2b74e.
//
// Solidity: event Unstaked(address indexed user, uint256 amount, uint256 stakeId)
func (_DcPayChannel *DcPayChannelFilterer) ParseUnstaked(log types.Log) (*DcPayChannelUnstaked, error) {
	event := new(DcPayChannelUnstaked)
	if err := _DcPayChannel.contract.UnpackLog(event, "Unstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
