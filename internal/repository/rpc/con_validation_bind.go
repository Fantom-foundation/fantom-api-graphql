// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rpc

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ContractValidatorABI is the input ABI used to generate the binding from.
const ContractValidatorABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ContractValidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ContractValidationDropped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ValidatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ValidatorDropped\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"list\",\"type\":\"address[]\"}],\"name\":\"addValidators\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"contracts\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"version\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"license\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"contact\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"compiler\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"clVersion\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isOptimized\",\"type\":\"bool\"},{\"internalType\":\"int32\",\"name\":\"optRounds\",\"type\":\"int32\"},{\"internalType\":\"bytes32\",\"name\":\"srcHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"validated\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"dropValidation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"list\",\"type\":\"address[]\"}],\"name\":\"dropValidators\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"version\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"license\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"contact\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"compiler\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"clVersion\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isOptimized\",\"type\":\"bool\"},{\"internalType\":\"int32\",\"name\":\"optRounds\",\"type\":\"int32\"},{\"internalType\":\"bytes\",\"name\":\"source\",\"type\":\"bytes\"}],\"name\":\"validated\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ContractValidator is an auto generated Go binding around an Ethereum contract.
type ContractValidator struct {
	ContractValidatorCaller     // Read-only binding to the contract
	ContractValidatorTransactor // Write-only binding to the contract
	ContractValidatorFilterer   // Log filterer for contract events
}

// ContractValidatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractValidatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractValidatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractValidatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractValidatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractValidatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractValidatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractValidatorSession struct {
	Contract     *ContractValidator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContractValidatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractValidatorCallerSession struct {
	Contract *ContractValidatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ContractValidatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractValidatorTransactorSession struct {
	Contract     *ContractValidatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ContractValidatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractValidatorRaw struct {
	Contract *ContractValidator // Generic contract binding to access the raw methods on
}

// ContractValidatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractValidatorCallerRaw struct {
	Contract *ContractValidatorCaller // Generic read-only contract binding to access the raw methods on
}

// ContractValidatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractValidatorTransactorRaw struct {
	Contract *ContractValidatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractValidator creates a new instance of ContractValidator, bound to a specific deployed contract.
func NewContractValidator(address common.Address, backend bind.ContractBackend) (*ContractValidator, error) {
	contract, err := bindContractValidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractValidator{ContractValidatorCaller: ContractValidatorCaller{contract: contract}, ContractValidatorTransactor: ContractValidatorTransactor{contract: contract}, ContractValidatorFilterer: ContractValidatorFilterer{contract: contract}}, nil
}

// NewContractValidatorCaller creates a new read-only instance of ContractValidator, bound to a specific deployed contract.
func NewContractValidatorCaller(address common.Address, caller bind.ContractCaller) (*ContractValidatorCaller, error) {
	contract, err := bindContractValidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractValidatorCaller{contract: contract}, nil
}

// NewContractValidatorTransactor creates a new write-only instance of ContractValidator, bound to a specific deployed contract.
func NewContractValidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractValidatorTransactor, error) {
	contract, err := bindContractValidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractValidatorTransactor{contract: contract}, nil
}

// NewContractValidatorFilterer creates a new log filterer instance of ContractValidator, bound to a specific deployed contract.
func NewContractValidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractValidatorFilterer, error) {
	contract, err := bindContractValidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractValidatorFilterer{contract: contract}, nil
}

// bindContractValidator binds a generic wrapper to an already deployed contract.
func bindContractValidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractValidatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractValidator *ContractValidatorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ContractValidator.Contract.ContractValidatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractValidator *ContractValidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractValidator.Contract.ContractValidatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractValidator *ContractValidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractValidator.Contract.ContractValidatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractValidator *ContractValidatorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ContractValidator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractValidator *ContractValidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractValidator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractValidator *ContractValidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractValidator.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_ContractValidator *ContractValidatorCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractValidator.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_ContractValidator *ContractValidatorSession) Admin() (common.Address, error) {
	return _ContractValidator.Contract.Admin(&_ContractValidator.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_ContractValidator *ContractValidatorCallerSession) Admin() (common.Address, error) {
	return _ContractValidator.Contract.Admin(&_ContractValidator.CallOpts)
}

// Contracts is a free data retrieval call binding the contract method 0x69dc9ff3.
//
// Solidity: function contracts(address ) view returns(string name, bytes32 version, bytes32 license, string contact, bytes32 compiler, bytes32 clVersion, bool isOptimized, int32 optRounds, bytes32 srcHash, uint256 validated)
func (_ContractValidator *ContractValidatorCaller) Contracts(opts *bind.CallOpts, arg0 common.Address) (struct {
	Name        string
	Version     [32]byte
	License     [32]byte
	Contact     string
	Compiler    [32]byte
	ClVersion   [32]byte
	IsOptimized bool
	OptRounds   int32
	SrcHash     [32]byte
	Validated   *big.Int
}, error) {
	ret := new(struct {
		Name        string
		Version     [32]byte
		License     [32]byte
		Contact     string
		Compiler    [32]byte
		ClVersion   [32]byte
		IsOptimized bool
		OptRounds   int32
		SrcHash     [32]byte
		Validated   *big.Int
	})
	out := ret
	err := _ContractValidator.contract.Call(opts, out, "contracts", arg0)
	return *ret, err
}

// Contracts is a free data retrieval call binding the contract method 0x69dc9ff3.
//
// Solidity: function contracts(address ) view returns(string name, bytes32 version, bytes32 license, string contact, bytes32 compiler, bytes32 clVersion, bool isOptimized, int32 optRounds, bytes32 srcHash, uint256 validated)
func (_ContractValidator *ContractValidatorSession) Contracts(arg0 common.Address) (struct {
	Name        string
	Version     [32]byte
	License     [32]byte
	Contact     string
	Compiler    [32]byte
	ClVersion   [32]byte
	IsOptimized bool
	OptRounds   int32
	SrcHash     [32]byte
	Validated   *big.Int
}, error) {
	return _ContractValidator.Contract.Contracts(&_ContractValidator.CallOpts, arg0)
}

// Contracts is a free data retrieval call binding the contract method 0x69dc9ff3.
//
// Solidity: function contracts(address ) view returns(string name, bytes32 version, bytes32 license, string contact, bytes32 compiler, bytes32 clVersion, bool isOptimized, int32 optRounds, bytes32 srcHash, uint256 validated)
func (_ContractValidator *ContractValidatorCallerSession) Contracts(arg0 common.Address) (struct {
	Name        string
	Version     [32]byte
	License     [32]byte
	Contact     string
	Compiler    [32]byte
	ClVersion   [32]byte
	IsOptimized bool
	OptRounds   int32
	SrcHash     [32]byte
	Validated   *big.Int
}, error) {
	return _ContractValidator.Contract.Contracts(&_ContractValidator.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(bool)
func (_ContractValidator *ContractValidatorCaller) Validators(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ContractValidator.contract.Call(opts, out, "validators", arg0)
	return *ret0, err
}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(bool)
func (_ContractValidator *ContractValidatorSession) Validators(arg0 common.Address) (bool, error) {
	return _ContractValidator.Contract.Validators(&_ContractValidator.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(bool)
func (_ContractValidator *ContractValidatorCallerSession) Validators(arg0 common.Address) (bool, error) {
	return _ContractValidator.Contract.Validators(&_ContractValidator.CallOpts, arg0)
}

// AddValidators is a paid mutator transaction binding the contract method 0x70223952.
//
// Solidity: function addValidators(address[] list) returns()
func (_ContractValidator *ContractValidatorTransactor) AddValidators(opts *bind.TransactOpts, list []common.Address) (*types.Transaction, error) {
	return _ContractValidator.contract.Transact(opts, "addValidators", list)
}

// AddValidators is a paid mutator transaction binding the contract method 0x70223952.
//
// Solidity: function addValidators(address[] list) returns()
func (_ContractValidator *ContractValidatorSession) AddValidators(list []common.Address) (*types.Transaction, error) {
	return _ContractValidator.Contract.AddValidators(&_ContractValidator.TransactOpts, list)
}

// AddValidators is a paid mutator transaction binding the contract method 0x70223952.
//
// Solidity: function addValidators(address[] list) returns()
func (_ContractValidator *ContractValidatorTransactorSession) AddValidators(list []common.Address) (*types.Transaction, error) {
	return _ContractValidator.Contract.AddValidators(&_ContractValidator.TransactOpts, list)
}

// DropValidation is a paid mutator transaction binding the contract method 0xfb2997cc.
//
// Solidity: function dropValidation(address addr) returns()
func (_ContractValidator *ContractValidatorTransactor) DropValidation(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _ContractValidator.contract.Transact(opts, "dropValidation", addr)
}

// DropValidation is a paid mutator transaction binding the contract method 0xfb2997cc.
//
// Solidity: function dropValidation(address addr) returns()
func (_ContractValidator *ContractValidatorSession) DropValidation(addr common.Address) (*types.Transaction, error) {
	return _ContractValidator.Contract.DropValidation(&_ContractValidator.TransactOpts, addr)
}

// DropValidation is a paid mutator transaction binding the contract method 0xfb2997cc.
//
// Solidity: function dropValidation(address addr) returns()
func (_ContractValidator *ContractValidatorTransactorSession) DropValidation(addr common.Address) (*types.Transaction, error) {
	return _ContractValidator.Contract.DropValidation(&_ContractValidator.TransactOpts, addr)
}

// DropValidators is a paid mutator transaction binding the contract method 0x6d0da93b.
//
// Solidity: function dropValidators(address[] list) returns()
func (_ContractValidator *ContractValidatorTransactor) DropValidators(opts *bind.TransactOpts, list []common.Address) (*types.Transaction, error) {
	return _ContractValidator.contract.Transact(opts, "dropValidators", list)
}

// DropValidators is a paid mutator transaction binding the contract method 0x6d0da93b.
//
// Solidity: function dropValidators(address[] list) returns()
func (_ContractValidator *ContractValidatorSession) DropValidators(list []common.Address) (*types.Transaction, error) {
	return _ContractValidator.Contract.DropValidators(&_ContractValidator.TransactOpts, list)
}

// DropValidators is a paid mutator transaction binding the contract method 0x6d0da93b.
//
// Solidity: function dropValidators(address[] list) returns()
func (_ContractValidator *ContractValidatorTransactorSession) DropValidators(list []common.Address) (*types.Transaction, error) {
	return _ContractValidator.Contract.DropValidators(&_ContractValidator.TransactOpts, list)
}

// Validated is a paid mutator transaction binding the contract method 0x16409ce5.
//
// Solidity: function validated(address addr, string name, bytes32 version, bytes32 license, string contact, bytes32 compiler, bytes32 clVersion, bool isOptimized, int32 optRounds, bytes source) returns()
func (_ContractValidator *ContractValidatorTransactor) Validated(opts *bind.TransactOpts, addr common.Address, name string, version [32]byte, license [32]byte, contact string, compiler [32]byte, clVersion [32]byte, isOptimized bool, optRounds int32, source []byte) (*types.Transaction, error) {
	return _ContractValidator.contract.Transact(opts, "validated", addr, name, version, license, contact, compiler, clVersion, isOptimized, optRounds, source)
}

// Validated is a paid mutator transaction binding the contract method 0x16409ce5.
//
// Solidity: function validated(address addr, string name, bytes32 version, bytes32 license, string contact, bytes32 compiler, bytes32 clVersion, bool isOptimized, int32 optRounds, bytes source) returns()
func (_ContractValidator *ContractValidatorSession) Validated(addr common.Address, name string, version [32]byte, license [32]byte, contact string, compiler [32]byte, clVersion [32]byte, isOptimized bool, optRounds int32, source []byte) (*types.Transaction, error) {
	return _ContractValidator.Contract.Validated(&_ContractValidator.TransactOpts, addr, name, version, license, contact, compiler, clVersion, isOptimized, optRounds, source)
}

// Validated is a paid mutator transaction binding the contract method 0x16409ce5.
//
// Solidity: function validated(address addr, string name, bytes32 version, bytes32 license, string contact, bytes32 compiler, bytes32 clVersion, bool isOptimized, int32 optRounds, bytes source) returns()
func (_ContractValidator *ContractValidatorTransactorSession) Validated(addr common.Address, name string, version [32]byte, license [32]byte, contact string, compiler [32]byte, clVersion [32]byte, isOptimized bool, optRounds int32, source []byte) (*types.Transaction, error) {
	return _ContractValidator.Contract.Validated(&_ContractValidator.TransactOpts, addr, name, version, license, contact, compiler, clVersion, isOptimized, optRounds, source)
}

// ContractValidatorContractValidatedIterator is returned from FilterContractValidated and is used to iterate over the raw logs and unpacked data for ContractValidated events raised by the ContractValidator contract.
type ContractValidatorContractValidatedIterator struct {
	Event *ContractValidatorContractValidated // Event containing the contract specifics and raw log

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
func (it *ContractValidatorContractValidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractValidatorContractValidated)
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
		it.Event = new(ContractValidatorContractValidated)
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
func (it *ContractValidatorContractValidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractValidatorContractValidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractValidatorContractValidated represents a ContractValidated event raised by the ContractValidator contract.
type ContractValidatorContractValidated struct {
	Addr      common.Address
	SrcHash   [32]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterContractValidated is a free log retrieval operation binding the contract event 0x3b15bde36b64b2a249391c52a2357c68839cecbe92e4daee30577bab8bb5808a.
//
// Solidity: event ContractValidated(address indexed addr, bytes32 srcHash, uint256 timestamp)
func (_ContractValidator *ContractValidatorFilterer) FilterContractValidated(opts *bind.FilterOpts, addr []common.Address) (*ContractValidatorContractValidatedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _ContractValidator.contract.FilterLogs(opts, "ContractValidated", addrRule)
	if err != nil {
		return nil, err
	}
	return &ContractValidatorContractValidatedIterator{contract: _ContractValidator.contract, event: "ContractValidated", logs: logs, sub: sub}, nil
}

// WatchContractValidated is a free log subscription operation binding the contract event 0x3b15bde36b64b2a249391c52a2357c68839cecbe92e4daee30577bab8bb5808a.
//
// Solidity: event ContractValidated(address indexed addr, bytes32 srcHash, uint256 timestamp)
func (_ContractValidator *ContractValidatorFilterer) WatchContractValidated(opts *bind.WatchOpts, sink chan<- *ContractValidatorContractValidated, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _ContractValidator.contract.WatchLogs(opts, "ContractValidated", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractValidatorContractValidated)
				if err := _ContractValidator.contract.UnpackLog(event, "ContractValidated", log); err != nil {
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

// ParseContractValidated is a log parse operation binding the contract event 0x3b15bde36b64b2a249391c52a2357c68839cecbe92e4daee30577bab8bb5808a.
//
// Solidity: event ContractValidated(address indexed addr, bytes32 srcHash, uint256 timestamp)
func (_ContractValidator *ContractValidatorFilterer) ParseContractValidated(log types.Log) (*ContractValidatorContractValidated, error) {
	event := new(ContractValidatorContractValidated)
	if err := _ContractValidator.contract.UnpackLog(event, "ContractValidated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ContractValidatorContractValidationDroppedIterator is returned from FilterContractValidationDropped and is used to iterate over the raw logs and unpacked data for ContractValidationDropped events raised by the ContractValidator contract.
type ContractValidatorContractValidationDroppedIterator struct {
	Event *ContractValidatorContractValidationDropped // Event containing the contract specifics and raw log

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
func (it *ContractValidatorContractValidationDroppedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractValidatorContractValidationDropped)
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
		it.Event = new(ContractValidatorContractValidationDropped)
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
func (it *ContractValidatorContractValidationDroppedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractValidatorContractValidationDroppedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractValidatorContractValidationDropped represents a ContractValidationDropped event raised by the ContractValidator contract.
type ContractValidatorContractValidationDropped struct {
	Addr      common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterContractValidationDropped is a free log retrieval operation binding the contract event 0x80195ee4e23ad0ccf8ea361b589784cdc784902c93f4484375faa6a13c219379.
//
// Solidity: event ContractValidationDropped(address indexed addr, uint256 timestamp)
func (_ContractValidator *ContractValidatorFilterer) FilterContractValidationDropped(opts *bind.FilterOpts, addr []common.Address) (*ContractValidatorContractValidationDroppedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _ContractValidator.contract.FilterLogs(opts, "ContractValidationDropped", addrRule)
	if err != nil {
		return nil, err
	}
	return &ContractValidatorContractValidationDroppedIterator{contract: _ContractValidator.contract, event: "ContractValidationDropped", logs: logs, sub: sub}, nil
}

// WatchContractValidationDropped is a free log subscription operation binding the contract event 0x80195ee4e23ad0ccf8ea361b589784cdc784902c93f4484375faa6a13c219379.
//
// Solidity: event ContractValidationDropped(address indexed addr, uint256 timestamp)
func (_ContractValidator *ContractValidatorFilterer) WatchContractValidationDropped(opts *bind.WatchOpts, sink chan<- *ContractValidatorContractValidationDropped, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _ContractValidator.contract.WatchLogs(opts, "ContractValidationDropped", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractValidatorContractValidationDropped)
				if err := _ContractValidator.contract.UnpackLog(event, "ContractValidationDropped", log); err != nil {
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

// ParseContractValidationDropped is a log parse operation binding the contract event 0x80195ee4e23ad0ccf8ea361b589784cdc784902c93f4484375faa6a13c219379.
//
// Solidity: event ContractValidationDropped(address indexed addr, uint256 timestamp)
func (_ContractValidator *ContractValidatorFilterer) ParseContractValidationDropped(log types.Log) (*ContractValidatorContractValidationDropped, error) {
	event := new(ContractValidatorContractValidationDropped)
	if err := _ContractValidator.contract.UnpackLog(event, "ContractValidationDropped", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ContractValidatorValidatorAddedIterator is returned from FilterValidatorAdded and is used to iterate over the raw logs and unpacked data for ValidatorAdded events raised by the ContractValidator contract.
type ContractValidatorValidatorAddedIterator struct {
	Event *ContractValidatorValidatorAdded // Event containing the contract specifics and raw log

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
func (it *ContractValidatorValidatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractValidatorValidatorAdded)
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
		it.Event = new(ContractValidatorValidatorAdded)
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
func (it *ContractValidatorValidatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractValidatorValidatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractValidatorValidatorAdded represents a ValidatorAdded event raised by the ContractValidator contract.
type ContractValidatorValidatorAdded struct {
	Addr      common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorAdded is a free log retrieval operation binding the contract event 0x9000b209805850a65058f21361a9978cb30f1413ed555553ab52a59b440b5d99.
//
// Solidity: event ValidatorAdded(address indexed addr, uint256 timestamp)
func (_ContractValidator *ContractValidatorFilterer) FilterValidatorAdded(opts *bind.FilterOpts, addr []common.Address) (*ContractValidatorValidatorAddedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _ContractValidator.contract.FilterLogs(opts, "ValidatorAdded", addrRule)
	if err != nil {
		return nil, err
	}
	return &ContractValidatorValidatorAddedIterator{contract: _ContractValidator.contract, event: "ValidatorAdded", logs: logs, sub: sub}, nil
}

// WatchValidatorAdded is a free log subscription operation binding the contract event 0x9000b209805850a65058f21361a9978cb30f1413ed555553ab52a59b440b5d99.
//
// Solidity: event ValidatorAdded(address indexed addr, uint256 timestamp)
func (_ContractValidator *ContractValidatorFilterer) WatchValidatorAdded(opts *bind.WatchOpts, sink chan<- *ContractValidatorValidatorAdded, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _ContractValidator.contract.WatchLogs(opts, "ValidatorAdded", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractValidatorValidatorAdded)
				if err := _ContractValidator.contract.UnpackLog(event, "ValidatorAdded", log); err != nil {
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

// ParseValidatorAdded is a log parse operation binding the contract event 0x9000b209805850a65058f21361a9978cb30f1413ed555553ab52a59b440b5d99.
//
// Solidity: event ValidatorAdded(address indexed addr, uint256 timestamp)
func (_ContractValidator *ContractValidatorFilterer) ParseValidatorAdded(log types.Log) (*ContractValidatorValidatorAdded, error) {
	event := new(ContractValidatorValidatorAdded)
	if err := _ContractValidator.contract.UnpackLog(event, "ValidatorAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ContractValidatorValidatorDroppedIterator is returned from FilterValidatorDropped and is used to iterate over the raw logs and unpacked data for ValidatorDropped events raised by the ContractValidator contract.
type ContractValidatorValidatorDroppedIterator struct {
	Event *ContractValidatorValidatorDropped // Event containing the contract specifics and raw log

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
func (it *ContractValidatorValidatorDroppedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractValidatorValidatorDropped)
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
		it.Event = new(ContractValidatorValidatorDropped)
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
func (it *ContractValidatorValidatorDroppedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractValidatorValidatorDroppedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractValidatorValidatorDropped represents a ValidatorDropped event raised by the ContractValidator contract.
type ContractValidatorValidatorDropped struct {
	Addr      common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorDropped is a free log retrieval operation binding the contract event 0x7e6683db3b399c619e98bd9ea52059b78e79a1acaaf6a12438f9fbb2b9679a61.
//
// Solidity: event ValidatorDropped(address indexed addr, uint256 timestamp)
func (_ContractValidator *ContractValidatorFilterer) FilterValidatorDropped(opts *bind.FilterOpts, addr []common.Address) (*ContractValidatorValidatorDroppedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _ContractValidator.contract.FilterLogs(opts, "ValidatorDropped", addrRule)
	if err != nil {
		return nil, err
	}
	return &ContractValidatorValidatorDroppedIterator{contract: _ContractValidator.contract, event: "ValidatorDropped", logs: logs, sub: sub}, nil
}

// WatchValidatorDropped is a free log subscription operation binding the contract event 0x7e6683db3b399c619e98bd9ea52059b78e79a1acaaf6a12438f9fbb2b9679a61.
//
// Solidity: event ValidatorDropped(address indexed addr, uint256 timestamp)
func (_ContractValidator *ContractValidatorFilterer) WatchValidatorDropped(opts *bind.WatchOpts, sink chan<- *ContractValidatorValidatorDropped, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _ContractValidator.contract.WatchLogs(opts, "ValidatorDropped", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractValidatorValidatorDropped)
				if err := _ContractValidator.contract.UnpackLog(event, "ValidatorDropped", log); err != nil {
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

// ParseValidatorDropped is a log parse operation binding the contract event 0x7e6683db3b399c619e98bd9ea52059b78e79a1acaaf6a12438f9fbb2b9679a61.
//
// Solidity: event ValidatorDropped(address indexed addr, uint256 timestamp)
func (_ContractValidator *ContractValidatorFilterer) ParseValidatorDropped(log types.Log) (*ContractValidatorValidatorDropped, error) {
	event := new(ContractValidatorValidatorDropped)
	if err := _ContractValidator.contract.UnpackLog(event, "ValidatorDropped", log); err != nil {
		return nil, err
	}
	return event, nil
}
