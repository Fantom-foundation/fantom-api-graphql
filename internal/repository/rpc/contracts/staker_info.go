// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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
)

// StakerInfoContractMetaData contains all meta data concerning the StakerInfoContract contract.
var StakerInfoContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakerContractAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"}],\"name\":\"InfoUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakerInfos\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakerContractAddress\",\"type\":\"address\"}],\"name\":\"updateStakerContractAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_configUrl\",\"type\":\"string\"}],\"name\":\"updateInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakerID\",\"type\":\"uint256\"}],\"name\":\"getInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// StakerInfoContractABI is the input ABI used to generate the binding from.
// Deprecated: Use StakerInfoContractMetaData.ABI instead.
var StakerInfoContractABI = StakerInfoContractMetaData.ABI

// StakerInfoContract is an auto generated Go binding around an Ethereum contract.
type StakerInfoContract struct {
	StakerInfoContractCaller     // Read-only binding to the contract
	StakerInfoContractTransactor // Write-only binding to the contract
	StakerInfoContractFilterer   // Log filterer for contract events
}

// StakerInfoContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakerInfoContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakerInfoContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakerInfoContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakerInfoContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakerInfoContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakerInfoContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakerInfoContractSession struct {
	Contract     *StakerInfoContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// StakerInfoContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakerInfoContractCallerSession struct {
	Contract *StakerInfoContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// StakerInfoContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakerInfoContractTransactorSession struct {
	Contract     *StakerInfoContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// StakerInfoContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakerInfoContractRaw struct {
	Contract *StakerInfoContract // Generic contract binding to access the raw methods on
}

// StakerInfoContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakerInfoContractCallerRaw struct {
	Contract *StakerInfoContractCaller // Generic read-only contract binding to access the raw methods on
}

// StakerInfoContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakerInfoContractTransactorRaw struct {
	Contract *StakerInfoContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakerInfoContract creates a new instance of StakerInfoContract, bound to a specific deployed contract.
func NewStakerInfoContract(address common.Address, backend bind.ContractBackend) (*StakerInfoContract, error) {
	contract, err := bindStakerInfoContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakerInfoContract{StakerInfoContractCaller: StakerInfoContractCaller{contract: contract}, StakerInfoContractTransactor: StakerInfoContractTransactor{contract: contract}, StakerInfoContractFilterer: StakerInfoContractFilterer{contract: contract}}, nil
}

// NewStakerInfoContractCaller creates a new read-only instance of StakerInfoContract, bound to a specific deployed contract.
func NewStakerInfoContractCaller(address common.Address, caller bind.ContractCaller) (*StakerInfoContractCaller, error) {
	contract, err := bindStakerInfoContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakerInfoContractCaller{contract: contract}, nil
}

// NewStakerInfoContractTransactor creates a new write-only instance of StakerInfoContract, bound to a specific deployed contract.
func NewStakerInfoContractTransactor(address common.Address, transactor bind.ContractTransactor) (*StakerInfoContractTransactor, error) {
	contract, err := bindStakerInfoContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakerInfoContractTransactor{contract: contract}, nil
}

// NewStakerInfoContractFilterer creates a new log filterer instance of StakerInfoContract, bound to a specific deployed contract.
func NewStakerInfoContractFilterer(address common.Address, filterer bind.ContractFilterer) (*StakerInfoContractFilterer, error) {
	contract, err := bindStakerInfoContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakerInfoContractFilterer{contract: contract}, nil
}

// bindStakerInfoContract binds a generic wrapper to an already deployed contract.
func bindStakerInfoContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakerInfoContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakerInfoContract *StakerInfoContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakerInfoContract.Contract.StakerInfoContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakerInfoContract *StakerInfoContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakerInfoContract.Contract.StakerInfoContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakerInfoContract *StakerInfoContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakerInfoContract.Contract.StakerInfoContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakerInfoContract *StakerInfoContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakerInfoContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakerInfoContract *StakerInfoContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakerInfoContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakerInfoContract *StakerInfoContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakerInfoContract.Contract.contract.Transact(opts, method, params...)
}

// GetInfo is a free data retrieval call binding the contract method 0x1a3cd59a.
//
// Solidity: function getInfo(uint256 _stakerID) view returns(string)
func (_StakerInfoContract *StakerInfoContractCaller) GetInfo(opts *bind.CallOpts, _stakerID *big.Int) (string, error) {
	var out []interface{}
	err := _StakerInfoContract.contract.Call(opts, &out, "getInfo", _stakerID)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetInfo is a free data retrieval call binding the contract method 0x1a3cd59a.
//
// Solidity: function getInfo(uint256 _stakerID) view returns(string)
func (_StakerInfoContract *StakerInfoContractSession) GetInfo(_stakerID *big.Int) (string, error) {
	return _StakerInfoContract.Contract.GetInfo(&_StakerInfoContract.CallOpts, _stakerID)
}

// GetInfo is a free data retrieval call binding the contract method 0x1a3cd59a.
//
// Solidity: function getInfo(uint256 _stakerID) view returns(string)
func (_StakerInfoContract *StakerInfoContractCallerSession) GetInfo(_stakerID *big.Int) (string, error) {
	return _StakerInfoContract.Contract.GetInfo(&_StakerInfoContract.CallOpts, _stakerID)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_StakerInfoContract *StakerInfoContractCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakerInfoContract.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_StakerInfoContract *StakerInfoContractSession) IsOwner() (bool, error) {
	return _StakerInfoContract.Contract.IsOwner(&_StakerInfoContract.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_StakerInfoContract *StakerInfoContractCallerSession) IsOwner() (bool, error) {
	return _StakerInfoContract.Contract.IsOwner(&_StakerInfoContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakerInfoContract *StakerInfoContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakerInfoContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakerInfoContract *StakerInfoContractSession) Owner() (common.Address, error) {
	return _StakerInfoContract.Contract.Owner(&_StakerInfoContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakerInfoContract *StakerInfoContractCallerSession) Owner() (common.Address, error) {
	return _StakerInfoContract.Contract.Owner(&_StakerInfoContract.CallOpts)
}

// StakerInfos is a free data retrieval call binding the contract method 0x33470433.
//
// Solidity: function stakerInfos(uint256 ) view returns(string)
func (_StakerInfoContract *StakerInfoContractCaller) StakerInfos(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _StakerInfoContract.contract.Call(opts, &out, "stakerInfos", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// StakerInfos is a free data retrieval call binding the contract method 0x33470433.
//
// Solidity: function stakerInfos(uint256 ) view returns(string)
func (_StakerInfoContract *StakerInfoContractSession) StakerInfos(arg0 *big.Int) (string, error) {
	return _StakerInfoContract.Contract.StakerInfos(&_StakerInfoContract.CallOpts, arg0)
}

// StakerInfos is a free data retrieval call binding the contract method 0x33470433.
//
// Solidity: function stakerInfos(uint256 ) view returns(string)
func (_StakerInfoContract *StakerInfoContractCallerSession) StakerInfos(arg0 *big.Int) (string, error) {
	return _StakerInfoContract.Contract.StakerInfos(&_StakerInfoContract.CallOpts, arg0)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakerInfoContract *StakerInfoContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakerInfoContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakerInfoContract *StakerInfoContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakerInfoContract.Contract.RenounceOwnership(&_StakerInfoContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakerInfoContract *StakerInfoContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakerInfoContract.Contract.RenounceOwnership(&_StakerInfoContract.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakerInfoContract *StakerInfoContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StakerInfoContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakerInfoContract *StakerInfoContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakerInfoContract.Contract.TransferOwnership(&_StakerInfoContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakerInfoContract *StakerInfoContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakerInfoContract.Contract.TransferOwnership(&_StakerInfoContract.TransactOpts, newOwner)
}

// UpdateInfo is a paid mutator transaction binding the contract method 0xab29511b.
//
// Solidity: function updateInfo(string _configUrl) returns()
func (_StakerInfoContract *StakerInfoContractTransactor) UpdateInfo(opts *bind.TransactOpts, _configUrl string) (*types.Transaction, error) {
	return _StakerInfoContract.contract.Transact(opts, "updateInfo", _configUrl)
}

// UpdateInfo is a paid mutator transaction binding the contract method 0xab29511b.
//
// Solidity: function updateInfo(string _configUrl) returns()
func (_StakerInfoContract *StakerInfoContractSession) UpdateInfo(_configUrl string) (*types.Transaction, error) {
	return _StakerInfoContract.Contract.UpdateInfo(&_StakerInfoContract.TransactOpts, _configUrl)
}

// UpdateInfo is a paid mutator transaction binding the contract method 0xab29511b.
//
// Solidity: function updateInfo(string _configUrl) returns()
func (_StakerInfoContract *StakerInfoContractTransactorSession) UpdateInfo(_configUrl string) (*types.Transaction, error) {
	return _StakerInfoContract.Contract.UpdateInfo(&_StakerInfoContract.TransactOpts, _configUrl)
}

// UpdateStakerContractAddress is a paid mutator transaction binding the contract method 0xfb9cba30.
//
// Solidity: function updateStakerContractAddress(address _stakerContractAddress) returns()
func (_StakerInfoContract *StakerInfoContractTransactor) UpdateStakerContractAddress(opts *bind.TransactOpts, _stakerContractAddress common.Address) (*types.Transaction, error) {
	return _StakerInfoContract.contract.Transact(opts, "updateStakerContractAddress", _stakerContractAddress)
}

// UpdateStakerContractAddress is a paid mutator transaction binding the contract method 0xfb9cba30.
//
// Solidity: function updateStakerContractAddress(address _stakerContractAddress) returns()
func (_StakerInfoContract *StakerInfoContractSession) UpdateStakerContractAddress(_stakerContractAddress common.Address) (*types.Transaction, error) {
	return _StakerInfoContract.Contract.UpdateStakerContractAddress(&_StakerInfoContract.TransactOpts, _stakerContractAddress)
}

// UpdateStakerContractAddress is a paid mutator transaction binding the contract method 0xfb9cba30.
//
// Solidity: function updateStakerContractAddress(address _stakerContractAddress) returns()
func (_StakerInfoContract *StakerInfoContractTransactorSession) UpdateStakerContractAddress(_stakerContractAddress common.Address) (*types.Transaction, error) {
	return _StakerInfoContract.Contract.UpdateStakerContractAddress(&_StakerInfoContract.TransactOpts, _stakerContractAddress)
}

// StakerInfoContractInfoUpdatedIterator is returned from FilterInfoUpdated and is used to iterate over the raw logs and unpacked data for InfoUpdated events raised by the StakerInfoContract contract.
type StakerInfoContractInfoUpdatedIterator struct {
	Event *StakerInfoContractInfoUpdated // Event containing the contract specifics and raw log

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
func (it *StakerInfoContractInfoUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakerInfoContractInfoUpdated)
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
		it.Event = new(StakerInfoContractInfoUpdated)
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
func (it *StakerInfoContractInfoUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakerInfoContractInfoUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakerInfoContractInfoUpdated represents a InfoUpdated event raised by the StakerInfoContract contract.
type StakerInfoContractInfoUpdated struct {
	StakerID *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInfoUpdated is a free log retrieval operation binding the contract event 0x3a668b70276c6b5af986be90ab9921c67bbef483987bb44cd5145c4984e59f24.
//
// Solidity: event InfoUpdated(uint256 stakerID)
func (_StakerInfoContract *StakerInfoContractFilterer) FilterInfoUpdated(opts *bind.FilterOpts) (*StakerInfoContractInfoUpdatedIterator, error) {

	logs, sub, err := _StakerInfoContract.contract.FilterLogs(opts, "InfoUpdated")
	if err != nil {
		return nil, err
	}
	return &StakerInfoContractInfoUpdatedIterator{contract: _StakerInfoContract.contract, event: "InfoUpdated", logs: logs, sub: sub}, nil
}

// WatchInfoUpdated is a free log subscription operation binding the contract event 0x3a668b70276c6b5af986be90ab9921c67bbef483987bb44cd5145c4984e59f24.
//
// Solidity: event InfoUpdated(uint256 stakerID)
func (_StakerInfoContract *StakerInfoContractFilterer) WatchInfoUpdated(opts *bind.WatchOpts, sink chan<- *StakerInfoContractInfoUpdated) (event.Subscription, error) {

	logs, sub, err := _StakerInfoContract.contract.WatchLogs(opts, "InfoUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakerInfoContractInfoUpdated)
				if err := _StakerInfoContract.contract.UnpackLog(event, "InfoUpdated", log); err != nil {
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

// ParseInfoUpdated is a log parse operation binding the contract event 0x3a668b70276c6b5af986be90ab9921c67bbef483987bb44cd5145c4984e59f24.
//
// Solidity: event InfoUpdated(uint256 stakerID)
func (_StakerInfoContract *StakerInfoContractFilterer) ParseInfoUpdated(log types.Log) (*StakerInfoContractInfoUpdated, error) {
	event := new(StakerInfoContractInfoUpdated)
	if err := _StakerInfoContract.contract.UnpackLog(event, "InfoUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakerInfoContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StakerInfoContract contract.
type StakerInfoContractOwnershipTransferredIterator struct {
	Event *StakerInfoContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakerInfoContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakerInfoContractOwnershipTransferred)
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
		it.Event = new(StakerInfoContractOwnershipTransferred)
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
func (it *StakerInfoContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakerInfoContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakerInfoContractOwnershipTransferred represents a OwnershipTransferred event raised by the StakerInfoContract contract.
type StakerInfoContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakerInfoContract *StakerInfoContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakerInfoContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakerInfoContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakerInfoContractOwnershipTransferredIterator{contract: _StakerInfoContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakerInfoContract *StakerInfoContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakerInfoContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakerInfoContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakerInfoContractOwnershipTransferred)
				if err := _StakerInfoContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_StakerInfoContract *StakerInfoContractFilterer) ParseOwnershipTransferred(log types.Log) (*StakerInfoContractOwnershipTransferred, error) {
	event := new(StakerInfoContractOwnershipTransferred)
	if err := _StakerInfoContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
