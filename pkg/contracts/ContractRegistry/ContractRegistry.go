// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ContractRegistry

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

// ContractRegistryABI is the input ABI used to generate the binding from.
const ContractRegistryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"ContractAdded\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"addContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"getContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ContractRegistry is an auto generated Go binding around an Ethereum contract.
type ContractRegistry struct {
	ContractRegistryCaller     // Read-only binding to the contract
	ContractRegistryTransactor // Write-only binding to the contract
	ContractRegistryFilterer   // Log filterer for contract events
}

// ContractRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractRegistrySession struct {
	Contract     *ContractRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractRegistryCallerSession struct {
	Contract *ContractRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ContractRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractRegistryTransactorSession struct {
	Contract     *ContractRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ContractRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRegistryRaw struct {
	Contract *ContractRegistry // Generic contract binding to access the raw methods on
}

// ContractRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractRegistryCallerRaw struct {
	Contract *ContractRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ContractRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractRegistryTransactorRaw struct {
	Contract *ContractRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractRegistry creates a new instance of ContractRegistry, bound to a specific deployed contract.
func NewContractRegistry(address common.Address, backend bind.ContractBackend) (*ContractRegistry, error) {
	contract, err := bindContractRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractRegistry{ContractRegistryCaller: ContractRegistryCaller{contract: contract}, ContractRegistryTransactor: ContractRegistryTransactor{contract: contract}, ContractRegistryFilterer: ContractRegistryFilterer{contract: contract}}, nil
}

// NewContractRegistryCaller creates a new read-only instance of ContractRegistry, bound to a specific deployed contract.
func NewContractRegistryCaller(address common.Address, caller bind.ContractCaller) (*ContractRegistryCaller, error) {
	contract, err := bindContractRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryCaller{contract: contract}, nil
}

// NewContractRegistryTransactor creates a new write-only instance of ContractRegistry, bound to a specific deployed contract.
func NewContractRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractRegistryTransactor, error) {
	contract, err := bindContractRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryTransactor{contract: contract}, nil
}

// NewContractRegistryFilterer creates a new log filterer instance of ContractRegistry, bound to a specific deployed contract.
func NewContractRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractRegistryFilterer, error) {
	contract, err := bindContractRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryFilterer{contract: contract}, nil
}

// bindContractRegistry binds a generic wrapper to an already deployed contract.
func bindContractRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractRegistry *ContractRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ContractRegistry.Contract.ContractRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractRegistry *ContractRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractRegistry.Contract.ContractRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractRegistry *ContractRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractRegistry.Contract.ContractRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractRegistry *ContractRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ContractRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractRegistry *ContractRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractRegistry *ContractRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetContract is a free data retrieval call binding the contract method 0x35817773.
//
// Solidity: function getContract(string _name) constant returns(address)
func (_ContractRegistry *ContractRegistryCaller) GetContract(opts *bind.CallOpts, _name string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractRegistry.contract.Call(opts, out, "getContract", _name)
	return *ret0, err
}

// GetContract is a free data retrieval call binding the contract method 0x35817773.
//
// Solidity: function getContract(string _name) constant returns(address)
func (_ContractRegistry *ContractRegistrySession) GetContract(_name string) (common.Address, error) {
	return _ContractRegistry.Contract.GetContract(&_ContractRegistry.CallOpts, _name)
}

// GetContract is a free data retrieval call binding the contract method 0x35817773.
//
// Solidity: function getContract(string _name) constant returns(address)
func (_ContractRegistry *ContractRegistryCallerSession) GetContract(_name string) (common.Address, error) {
	return _ContractRegistry.Contract.GetContract(&_ContractRegistry.CallOpts, _name)
}

// AddContract is a paid mutator transaction binding the contract method 0xbf5b6016.
//
// Solidity: function addContract(string _name, address _address) returns()
func (_ContractRegistry *ContractRegistryTransactor) AddContract(opts *bind.TransactOpts, _name string, _address common.Address) (*types.Transaction, error) {
	return _ContractRegistry.contract.Transact(opts, "addContract", _name, _address)
}

// AddContract is a paid mutator transaction binding the contract method 0xbf5b6016.
//
// Solidity: function addContract(string _name, address _address) returns()
func (_ContractRegistry *ContractRegistrySession) AddContract(_name string, _address common.Address) (*types.Transaction, error) {
	return _ContractRegistry.Contract.AddContract(&_ContractRegistry.TransactOpts, _name, _address)
}

// AddContract is a paid mutator transaction binding the contract method 0xbf5b6016.
//
// Solidity: function addContract(string _name, address _address) returns()
func (_ContractRegistry *ContractRegistryTransactorSession) AddContract(_name string, _address common.Address) (*types.Transaction, error) {
	return _ContractRegistry.Contract.AddContract(&_ContractRegistry.TransactOpts, _name, _address)
}

// ContractRegistryContractAddedIterator is returned from FilterContractAdded and is used to iterate over the raw logs and unpacked data for ContractAdded events raised by the ContractRegistry contract.
type ContractRegistryContractAddedIterator struct {
	Event *ContractRegistryContractAdded // Event containing the contract specifics and raw log

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
func (it *ContractRegistryContractAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRegistryContractAdded)
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
		it.Event = new(ContractRegistryContractAdded)
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
func (it *ContractRegistryContractAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRegistryContractAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRegistryContractAdded represents a ContractAdded event raised by the ContractRegistry contract.
type ContractRegistryContractAdded struct {
	Name            string
	ContractAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterContractAdded is a free log retrieval operation binding the contract event 0x8b4ef7d4e5bc8f098e6f637ac0acf4aee47b3f027efea6307264b06b4bc9d298.
//
// Solidity: event ContractAdded(string name, address indexed contractAddress)
func (_ContractRegistry *ContractRegistryFilterer) FilterContractAdded(opts *bind.FilterOpts, contractAddress []common.Address) (*ContractRegistryContractAddedIterator, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _ContractRegistry.contract.FilterLogs(opts, "ContractAdded", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryContractAddedIterator{contract: _ContractRegistry.contract, event: "ContractAdded", logs: logs, sub: sub}, nil
}

// WatchContractAdded is a free log subscription operation binding the contract event 0x8b4ef7d4e5bc8f098e6f637ac0acf4aee47b3f027efea6307264b06b4bc9d298.
//
// Solidity: event ContractAdded(string name, address indexed contractAddress)
func (_ContractRegistry *ContractRegistryFilterer) WatchContractAdded(opts *bind.WatchOpts, sink chan<- *ContractRegistryContractAdded, contractAddress []common.Address) (event.Subscription, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _ContractRegistry.contract.WatchLogs(opts, "ContractAdded", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRegistryContractAdded)
				if err := _ContractRegistry.contract.UnpackLog(event, "ContractAdded", log); err != nil {
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

// ParseContractAdded is a log parse operation binding the contract event 0x8b4ef7d4e5bc8f098e6f637ac0acf4aee47b3f027efea6307264b06b4bc9d298.
//
// Solidity: event ContractAdded(string name, address indexed contractAddress)
func (_ContractRegistry *ContractRegistryFilterer) ParseContractAdded(log types.Log) (*ContractRegistryContractAdded, error) {
	event := new(ContractRegistryContractAdded)
	if err := _ContractRegistry.contract.UnpackLog(event, "ContractAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}
