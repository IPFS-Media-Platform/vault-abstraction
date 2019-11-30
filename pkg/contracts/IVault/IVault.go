// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IVault

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

// IVaultABI is the input ABI used to generate the binding from.
const IVaultABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_ipfsHash\",\"type\":\"string\"}],\"name\":\"add\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IVault is an auto generated Go binding around an Ethereum contract.
type IVault struct {
	IVaultCaller     // Read-only binding to the contract
	IVaultTransactor // Write-only binding to the contract
	IVaultFilterer   // Log filterer for contract events
}

// IVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type IVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IVaultSession struct {
	Contract     *IVault           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IVaultCallerSession struct {
	Contract *IVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IVaultTransactorSession struct {
	Contract     *IVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type IVaultRaw struct {
	Contract *IVault // Generic contract binding to access the raw methods on
}

// IVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IVaultCallerRaw struct {
	Contract *IVaultCaller // Generic read-only contract binding to access the raw methods on
}

// IVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IVaultTransactorRaw struct {
	Contract *IVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIVault creates a new instance of IVault, bound to a specific deployed contract.
func NewIVault(address common.Address, backend bind.ContractBackend) (*IVault, error) {
	contract, err := bindIVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IVault{IVaultCaller: IVaultCaller{contract: contract}, IVaultTransactor: IVaultTransactor{contract: contract}, IVaultFilterer: IVaultFilterer{contract: contract}}, nil
}

// NewIVaultCaller creates a new read-only instance of IVault, bound to a specific deployed contract.
func NewIVaultCaller(address common.Address, caller bind.ContractCaller) (*IVaultCaller, error) {
	contract, err := bindIVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IVaultCaller{contract: contract}, nil
}

// NewIVaultTransactor creates a new write-only instance of IVault, bound to a specific deployed contract.
func NewIVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*IVaultTransactor, error) {
	contract, err := bindIVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IVaultTransactor{contract: contract}, nil
}

// NewIVaultFilterer creates a new log filterer instance of IVault, bound to a specific deployed contract.
func NewIVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*IVaultFilterer, error) {
	contract, err := bindIVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IVaultFilterer{contract: contract}, nil
}

// bindIVault binds a generic wrapper to an already deployed contract.
func bindIVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IVaultABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVault *IVaultRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IVault.Contract.IVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVault *IVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVault.Contract.IVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVault *IVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVault.Contract.IVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVault *IVaultCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVault *IVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVault *IVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVault.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() constant returns(string)
func (_IVault *IVaultCaller) Get(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _IVault.contract.Call(opts, out, "get")
	return *ret0, err
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() constant returns(string)
func (_IVault *IVaultSession) Get() (string, error) {
	return _IVault.Contract.Get(&_IVault.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() constant returns(string)
func (_IVault *IVaultCallerSession) Get() (string, error) {
	return _IVault.Contract.Get(&_IVault.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0xb0c8f9dc.
//
// Solidity: function add(string _ipfsHash) returns()
func (_IVault *IVaultTransactor) Add(opts *bind.TransactOpts, _ipfsHash string) (*types.Transaction, error) {
	return _IVault.contract.Transact(opts, "add", _ipfsHash)
}

// Add is a paid mutator transaction binding the contract method 0xb0c8f9dc.
//
// Solidity: function add(string _ipfsHash) returns()
func (_IVault *IVaultSession) Add(_ipfsHash string) (*types.Transaction, error) {
	return _IVault.Contract.Add(&_IVault.TransactOpts, _ipfsHash)
}

// Add is a paid mutator transaction binding the contract method 0xb0c8f9dc.
//
// Solidity: function add(string _ipfsHash) returns()
func (_IVault *IVaultTransactorSession) Add(_ipfsHash string) (*types.Transaction, error) {
	return _IVault.Contract.Add(&_IVault.TransactOpts, _ipfsHash)
}
