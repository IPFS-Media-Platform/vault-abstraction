// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package VaultManager

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

// VaultManagerABI is the input ABI used to generate the binding from.
const VaultManagerABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"string\"}],\"name\":\"vaultList\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"vaultAddress\",\"type\":\"address\"}],\"name\":\"VaultAdded\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"addNewVault\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVaults\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// VaultManager is an auto generated Go binding around an Ethereum contract.
type VaultManager struct {
	VaultManagerCaller     // Read-only binding to the contract
	VaultManagerTransactor // Write-only binding to the contract
	VaultManagerFilterer   // Log filterer for contract events
}

// VaultManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaultManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultManagerSession struct {
	Contract     *VaultManager     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultManagerCallerSession struct {
	Contract *VaultManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// VaultManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultManagerTransactorSession struct {
	Contract     *VaultManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// VaultManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultManagerRaw struct {
	Contract *VaultManager // Generic contract binding to access the raw methods on
}

// VaultManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultManagerCallerRaw struct {
	Contract *VaultManagerCaller // Generic read-only contract binding to access the raw methods on
}

// VaultManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultManagerTransactorRaw struct {
	Contract *VaultManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVaultManager creates a new instance of VaultManager, bound to a specific deployed contract.
func NewVaultManager(address common.Address, backend bind.ContractBackend) (*VaultManager, error) {
	contract, err := bindVaultManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VaultManager{VaultManagerCaller: VaultManagerCaller{contract: contract}, VaultManagerTransactor: VaultManagerTransactor{contract: contract}, VaultManagerFilterer: VaultManagerFilterer{contract: contract}}, nil
}

// NewVaultManagerCaller creates a new read-only instance of VaultManager, bound to a specific deployed contract.
func NewVaultManagerCaller(address common.Address, caller bind.ContractCaller) (*VaultManagerCaller, error) {
	contract, err := bindVaultManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaultManagerCaller{contract: contract}, nil
}

// NewVaultManagerTransactor creates a new write-only instance of VaultManager, bound to a specific deployed contract.
func NewVaultManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultManagerTransactor, error) {
	contract, err := bindVaultManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaultManagerTransactor{contract: contract}, nil
}

// NewVaultManagerFilterer creates a new log filterer instance of VaultManager, bound to a specific deployed contract.
func NewVaultManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*VaultManagerFilterer, error) {
	contract, err := bindVaultManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaultManagerFilterer{contract: contract}, nil
}

// bindVaultManager binds a generic wrapper to an already deployed contract.
func bindVaultManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultManager *VaultManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VaultManager.Contract.VaultManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultManager *VaultManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultManager.Contract.VaultManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultManager *VaultManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultManager.Contract.VaultManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultManager *VaultManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VaultManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultManager *VaultManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultManager *VaultManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultManager.Contract.contract.Transact(opts, method, params...)
}

// GetVaults is a free data retrieval call binding the contract method 0x44d00f82.
//
// Solidity: function getVaults() constant returns(address[])
func (_VaultManager *VaultManagerCaller) GetVaults(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _VaultManager.contract.Call(opts, out, "getVaults")
	return *ret0, err
}

// GetVaults is a free data retrieval call binding the contract method 0x44d00f82.
//
// Solidity: function getVaults() constant returns(address[])
func (_VaultManager *VaultManagerSession) GetVaults() ([]common.Address, error) {
	return _VaultManager.Contract.GetVaults(&_VaultManager.CallOpts)
}

// GetVaults is a free data retrieval call binding the contract method 0x44d00f82.
//
// Solidity: function getVaults() constant returns(address[])
func (_VaultManager *VaultManagerCallerSession) GetVaults() ([]common.Address, error) {
	return _VaultManager.Contract.GetVaults(&_VaultManager.CallOpts)
}

// VaultList is a free data retrieval call binding the contract method 0x7e1c8e2d.
//
// Solidity: function vaultList(string ) constant returns(address)
func (_VaultManager *VaultManagerCaller) VaultList(opts *bind.CallOpts, arg0 string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _VaultManager.contract.Call(opts, out, "vaultList", arg0)
	return *ret0, err
}

// VaultList is a free data retrieval call binding the contract method 0x7e1c8e2d.
//
// Solidity: function vaultList(string ) constant returns(address)
func (_VaultManager *VaultManagerSession) VaultList(arg0 string) (common.Address, error) {
	return _VaultManager.Contract.VaultList(&_VaultManager.CallOpts, arg0)
}

// VaultList is a free data retrieval call binding the contract method 0x7e1c8e2d.
//
// Solidity: function vaultList(string ) constant returns(address)
func (_VaultManager *VaultManagerCallerSession) VaultList(arg0 string) (common.Address, error) {
	return _VaultManager.Contract.VaultList(&_VaultManager.CallOpts, arg0)
}

// AddNewVault is a paid mutator transaction binding the contract method 0x3f6378f4.
//
// Solidity: function addNewVault(string _name) returns(address)
func (_VaultManager *VaultManagerTransactor) AddNewVault(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _VaultManager.contract.Transact(opts, "addNewVault", _name)
}

// AddNewVault is a paid mutator transaction binding the contract method 0x3f6378f4.
//
// Solidity: function addNewVault(string _name) returns(address)
func (_VaultManager *VaultManagerSession) AddNewVault(_name string) (*types.Transaction, error) {
	return _VaultManager.Contract.AddNewVault(&_VaultManager.TransactOpts, _name)
}

// AddNewVault is a paid mutator transaction binding the contract method 0x3f6378f4.
//
// Solidity: function addNewVault(string _name) returns(address)
func (_VaultManager *VaultManagerTransactorSession) AddNewVault(_name string) (*types.Transaction, error) {
	return _VaultManager.Contract.AddNewVault(&_VaultManager.TransactOpts, _name)
}

// VaultManagerVaultAddedIterator is returned from FilterVaultAdded and is used to iterate over the raw logs and unpacked data for VaultAdded events raised by the VaultManager contract.
type VaultManagerVaultAddedIterator struct {
	Event *VaultManagerVaultAdded // Event containing the contract specifics and raw log

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
func (it *VaultManagerVaultAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultManagerVaultAdded)
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
		it.Event = new(VaultManagerVaultAdded)
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
func (it *VaultManagerVaultAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultManagerVaultAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultManagerVaultAdded represents a VaultAdded event raised by the VaultManager contract.
type VaultManagerVaultAdded struct {
	Name         string
	VaultAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterVaultAdded is a free log retrieval operation binding the contract event 0x4104dba30cb4c9a6b91a5229bc9214a5710fa836b100dffdd213346c0de42dcb.
//
// Solidity: event VaultAdded(string name, address vaultAddress)
func (_VaultManager *VaultManagerFilterer) FilterVaultAdded(opts *bind.FilterOpts) (*VaultManagerVaultAddedIterator, error) {

	logs, sub, err := _VaultManager.contract.FilterLogs(opts, "VaultAdded")
	if err != nil {
		return nil, err
	}
	return &VaultManagerVaultAddedIterator{contract: _VaultManager.contract, event: "VaultAdded", logs: logs, sub: sub}, nil
}

// WatchVaultAdded is a free log subscription operation binding the contract event 0x4104dba30cb4c9a6b91a5229bc9214a5710fa836b100dffdd213346c0de42dcb.
//
// Solidity: event VaultAdded(string name, address vaultAddress)
func (_VaultManager *VaultManagerFilterer) WatchVaultAdded(opts *bind.WatchOpts, sink chan<- *VaultManagerVaultAdded) (event.Subscription, error) {

	logs, sub, err := _VaultManager.contract.WatchLogs(opts, "VaultAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultManagerVaultAdded)
				if err := _VaultManager.contract.UnpackLog(event, "VaultAdded", log); err != nil {
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

// ParseVaultAdded is a log parse operation binding the contract event 0x4104dba30cb4c9a6b91a5229bc9214a5710fa836b100dffdd213346c0de42dcb.
//
// Solidity: event VaultAdded(string name, address vaultAddress)
func (_VaultManager *VaultManagerFilterer) ParseVaultAdded(log types.Log) (*VaultManagerVaultAdded, error) {
	event := new(VaultManagerVaultAdded)
	if err := _VaultManager.contract.UnpackLog(event, "VaultAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}
