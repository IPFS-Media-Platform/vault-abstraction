// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IContractRegistry

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IContractRegistryABI is the input ABI used to generate the binding from.
const IContractRegistryABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"addContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"getContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IContractRegistry is an auto generated Go binding around an Ethereum contract.
type IContractRegistry struct {
	IContractRegistryCaller     // Read-only binding to the contract
	IContractRegistryTransactor // Write-only binding to the contract
	IContractRegistryFilterer   // Log filterer for contract events
}

// IContractRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IContractRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IContractRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IContractRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IContractRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IContractRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IContractRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IContractRegistrySession struct {
	Contract     *IContractRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IContractRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IContractRegistryCallerSession struct {
	Contract *IContractRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IContractRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IContractRegistryTransactorSession struct {
	Contract     *IContractRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IContractRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IContractRegistryRaw struct {
	Contract *IContractRegistry // Generic contract binding to access the raw methods on
}

// IContractRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IContractRegistryCallerRaw struct {
	Contract *IContractRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// IContractRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IContractRegistryTransactorRaw struct {
	Contract *IContractRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIContractRegistry creates a new instance of IContractRegistry, bound to a specific deployed contract.
func NewIContractRegistry(address common.Address, backend bind.ContractBackend) (*IContractRegistry, error) {
	contract, err := bindIContractRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IContractRegistry{IContractRegistryCaller: IContractRegistryCaller{contract: contract}, IContractRegistryTransactor: IContractRegistryTransactor{contract: contract}, IContractRegistryFilterer: IContractRegistryFilterer{contract: contract}}, nil
}

// NewIContractRegistryCaller creates a new read-only instance of IContractRegistry, bound to a specific deployed contract.
func NewIContractRegistryCaller(address common.Address, caller bind.ContractCaller) (*IContractRegistryCaller, error) {
	contract, err := bindIContractRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IContractRegistryCaller{contract: contract}, nil
}

// NewIContractRegistryTransactor creates a new write-only instance of IContractRegistry, bound to a specific deployed contract.
func NewIContractRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*IContractRegistryTransactor, error) {
	contract, err := bindIContractRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IContractRegistryTransactor{contract: contract}, nil
}

// NewIContractRegistryFilterer creates a new log filterer instance of IContractRegistry, bound to a specific deployed contract.
func NewIContractRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*IContractRegistryFilterer, error) {
	contract, err := bindIContractRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IContractRegistryFilterer{contract: contract}, nil
}

// bindIContractRegistry binds a generic wrapper to an already deployed contract.
func bindIContractRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IContractRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IContractRegistry *IContractRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IContractRegistry.Contract.IContractRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IContractRegistry *IContractRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IContractRegistry.Contract.IContractRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IContractRegistry *IContractRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IContractRegistry.Contract.IContractRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IContractRegistry *IContractRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IContractRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IContractRegistry *IContractRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IContractRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IContractRegistry *IContractRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IContractRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetContract is a free data retrieval call binding the contract method 0x35817773.
//
// Solidity: function getContract(string _name) constant returns(address)
func (_IContractRegistry *IContractRegistryCaller) GetContract(opts *bind.CallOpts, _name string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IContractRegistry.contract.Call(opts, out, "getContract", _name)
	return *ret0, err
}

// GetContract is a free data retrieval call binding the contract method 0x35817773.
//
// Solidity: function getContract(string _name) constant returns(address)
func (_IContractRegistry *IContractRegistrySession) GetContract(_name string) (common.Address, error) {
	return _IContractRegistry.Contract.GetContract(&_IContractRegistry.CallOpts, _name)
}

// GetContract is a free data retrieval call binding the contract method 0x35817773.
//
// Solidity: function getContract(string _name) constant returns(address)
func (_IContractRegistry *IContractRegistryCallerSession) GetContract(_name string) (common.Address, error) {
	return _IContractRegistry.Contract.GetContract(&_IContractRegistry.CallOpts, _name)
}

// AddContract is a paid mutator transaction binding the contract method 0xbf5b6016.
//
// Solidity: function addContract(string _name, address _address) returns()
func (_IContractRegistry *IContractRegistryTransactor) AddContract(opts *bind.TransactOpts, _name string, _address common.Address) (*types.Transaction, error) {
	return _IContractRegistry.contract.Transact(opts, "addContract", _name, _address)
}

// AddContract is a paid mutator transaction binding the contract method 0xbf5b6016.
//
// Solidity: function addContract(string _name, address _address) returns()
func (_IContractRegistry *IContractRegistrySession) AddContract(_name string, _address common.Address) (*types.Transaction, error) {
	return _IContractRegistry.Contract.AddContract(&_IContractRegistry.TransactOpts, _name, _address)
}

// AddContract is a paid mutator transaction binding the contract method 0xbf5b6016.
//
// Solidity: function addContract(string _name, address _address) returns()
func (_IContractRegistry *IContractRegistryTransactorSession) AddContract(_name string, _address common.Address) (*types.Transaction, error) {
	return _IContractRegistry.Contract.AddContract(&_IContractRegistry.TransactOpts, _name, _address)
}
