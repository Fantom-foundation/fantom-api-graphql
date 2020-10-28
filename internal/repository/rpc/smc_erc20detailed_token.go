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

// ERCTwentyDetailedABI is the input ABI used to generate the binding from.
const ERCTwentyDetailedABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ERCTwentyDetailed is an auto generated Go binding around an Ethereum contract.
type ERCTwentyDetailed struct {
	ERCTwentyDetailedCaller     // Read-only binding to the contract
	ERCTwentyDetailedTransactor // Write-only binding to the contract
	ERCTwentyDetailedFilterer   // Log filterer for contract events
}

// ERCTwentyDetailedCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERCTwentyDetailedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERCTwentyDetailedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERCTwentyDetailedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERCTwentyDetailedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERCTwentyDetailedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERCTwentyDetailedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERCTwentyDetailedSession struct {
	Contract     *ERCTwentyDetailed // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ERCTwentyDetailedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERCTwentyDetailedCallerSession struct {
	Contract *ERCTwentyDetailedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ERCTwentyDetailedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERCTwentyDetailedTransactorSession struct {
	Contract     *ERCTwentyDetailedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ERCTwentyDetailedRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERCTwentyDetailedRaw struct {
	Contract *ERCTwentyDetailed // Generic contract binding to access the raw methods on
}

// ERCTwentyDetailedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERCTwentyDetailedCallerRaw struct {
	Contract *ERCTwentyDetailedCaller // Generic read-only contract binding to access the raw methods on
}

// ERCTwentyDetailedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERCTwentyDetailedTransactorRaw struct {
	Contract *ERCTwentyDetailedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERCTwentyDetailed creates a new instance of ERCTwentyDetailed, bound to a specific deployed contract.
func NewERCTwentyDetailed(address common.Address, backend bind.ContractBackend) (*ERCTwentyDetailed, error) {
	contract, err := bindERCTwentyDetailed(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERCTwentyDetailed{ERCTwentyDetailedCaller: ERCTwentyDetailedCaller{contract: contract}, ERCTwentyDetailedTransactor: ERCTwentyDetailedTransactor{contract: contract}, ERCTwentyDetailedFilterer: ERCTwentyDetailedFilterer{contract: contract}}, nil
}

// NewERCTwentyDetailedCaller creates a new read-only instance of ERCTwentyDetailed, bound to a specific deployed contract.
func NewERCTwentyDetailedCaller(address common.Address, caller bind.ContractCaller) (*ERCTwentyDetailedCaller, error) {
	contract, err := bindERCTwentyDetailed(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERCTwentyDetailedCaller{contract: contract}, nil
}

// NewERCTwentyDetailedTransactor creates a new write-only instance of ERCTwentyDetailed, bound to a specific deployed contract.
func NewERCTwentyDetailedTransactor(address common.Address, transactor bind.ContractTransactor) (*ERCTwentyDetailedTransactor, error) {
	contract, err := bindERCTwentyDetailed(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERCTwentyDetailedTransactor{contract: contract}, nil
}

// NewERCTwentyDetailedFilterer creates a new log filterer instance of ERCTwentyDetailed, bound to a specific deployed contract.
func NewERCTwentyDetailedFilterer(address common.Address, filterer bind.ContractFilterer) (*ERCTwentyDetailedFilterer, error) {
	contract, err := bindERCTwentyDetailed(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERCTwentyDetailedFilterer{contract: contract}, nil
}

// bindERCTwentyDetailed binds a generic wrapper to an already deployed contract.
func bindERCTwentyDetailed(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERCTwentyDetailedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERCTwentyDetailed *ERCTwentyDetailedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERCTwentyDetailed.Contract.ERCTwentyDetailedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERCTwentyDetailed *ERCTwentyDetailedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERCTwentyDetailed.Contract.ERCTwentyDetailedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERCTwentyDetailed *ERCTwentyDetailedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERCTwentyDetailed.Contract.ERCTwentyDetailedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERCTwentyDetailed *ERCTwentyDetailedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERCTwentyDetailed.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERCTwentyDetailed *ERCTwentyDetailedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERCTwentyDetailed.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERCTwentyDetailed *ERCTwentyDetailedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERCTwentyDetailed.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERCTwentyDetailed *ERCTwentyDetailedCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERCTwentyDetailed.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERCTwentyDetailed *ERCTwentyDetailedSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERCTwentyDetailed.Contract.Allowance(&_ERCTwentyDetailed.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERCTwentyDetailed *ERCTwentyDetailedCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERCTwentyDetailed.Contract.Allowance(&_ERCTwentyDetailed.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERCTwentyDetailed *ERCTwentyDetailedCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERCTwentyDetailed.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERCTwentyDetailed *ERCTwentyDetailedSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERCTwentyDetailed.Contract.BalanceOf(&_ERCTwentyDetailed.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERCTwentyDetailed *ERCTwentyDetailedCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERCTwentyDetailed.Contract.BalanceOf(&_ERCTwentyDetailed.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERCTwentyDetailed *ERCTwentyDetailedCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERCTwentyDetailed.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERCTwentyDetailed *ERCTwentyDetailedSession) Decimals() (uint8, error) {
	return _ERCTwentyDetailed.Contract.Decimals(&_ERCTwentyDetailed.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERCTwentyDetailed *ERCTwentyDetailedCallerSession) Decimals() (uint8, error) {
	return _ERCTwentyDetailed.Contract.Decimals(&_ERCTwentyDetailed.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERCTwentyDetailed *ERCTwentyDetailedCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERCTwentyDetailed.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERCTwentyDetailed *ERCTwentyDetailedSession) Name() (string, error) {
	return _ERCTwentyDetailed.Contract.Name(&_ERCTwentyDetailed.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERCTwentyDetailed *ERCTwentyDetailedCallerSession) Name() (string, error) {
	return _ERCTwentyDetailed.Contract.Name(&_ERCTwentyDetailed.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERCTwentyDetailed *ERCTwentyDetailedCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERCTwentyDetailed.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERCTwentyDetailed *ERCTwentyDetailedSession) Symbol() (string, error) {
	return _ERCTwentyDetailed.Contract.Symbol(&_ERCTwentyDetailed.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERCTwentyDetailed *ERCTwentyDetailedCallerSession) Symbol() (string, error) {
	return _ERCTwentyDetailed.Contract.Symbol(&_ERCTwentyDetailed.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERCTwentyDetailed *ERCTwentyDetailedCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERCTwentyDetailed.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERCTwentyDetailed *ERCTwentyDetailedSession) TotalSupply() (*big.Int, error) {
	return _ERCTwentyDetailed.Contract.TotalSupply(&_ERCTwentyDetailed.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERCTwentyDetailed *ERCTwentyDetailedCallerSession) TotalSupply() (*big.Int, error) {
	return _ERCTwentyDetailed.Contract.TotalSupply(&_ERCTwentyDetailed.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERCTwentyDetailed *ERCTwentyDetailedTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERCTwentyDetailed.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERCTwentyDetailed *ERCTwentyDetailedSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERCTwentyDetailed.Contract.Approve(&_ERCTwentyDetailed.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERCTwentyDetailed *ERCTwentyDetailedTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERCTwentyDetailed.Contract.Approve(&_ERCTwentyDetailed.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERCTwentyDetailed *ERCTwentyDetailedTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERCTwentyDetailed.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERCTwentyDetailed *ERCTwentyDetailedSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERCTwentyDetailed.Contract.Transfer(&_ERCTwentyDetailed.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERCTwentyDetailed *ERCTwentyDetailedTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERCTwentyDetailed.Contract.Transfer(&_ERCTwentyDetailed.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERCTwentyDetailed *ERCTwentyDetailedTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERCTwentyDetailed.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERCTwentyDetailed *ERCTwentyDetailedSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERCTwentyDetailed.Contract.TransferFrom(&_ERCTwentyDetailed.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERCTwentyDetailed *ERCTwentyDetailedTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERCTwentyDetailed.Contract.TransferFrom(&_ERCTwentyDetailed.TransactOpts, sender, recipient, amount)
}

// ERCTwentyDetailedApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERCTwentyDetailed contract.
type ERCTwentyDetailedApprovalIterator struct {
	Event *ERCTwentyDetailedApproval // Event containing the contract specifics and raw log

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
func (it *ERCTwentyDetailedApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERCTwentyDetailedApproval)
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
		it.Event = new(ERCTwentyDetailedApproval)
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
func (it *ERCTwentyDetailedApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERCTwentyDetailedApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERCTwentyDetailedApproval represents a Approval event raised by the ERCTwentyDetailed contract.
type ERCTwentyDetailedApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERCTwentyDetailed *ERCTwentyDetailedFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERCTwentyDetailedApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERCTwentyDetailed.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERCTwentyDetailedApprovalIterator{contract: _ERCTwentyDetailed.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERCTwentyDetailed *ERCTwentyDetailedFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERCTwentyDetailedApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERCTwentyDetailed.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERCTwentyDetailedApproval)
				if err := _ERCTwentyDetailed.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERCTwentyDetailed *ERCTwentyDetailedFilterer) ParseApproval(log types.Log) (*ERCTwentyDetailedApproval, error) {
	event := new(ERCTwentyDetailedApproval)
	if err := _ERCTwentyDetailed.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERCTwentyDetailedTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERCTwentyDetailed contract.
type ERCTwentyDetailedTransferIterator struct {
	Event *ERCTwentyDetailedTransfer // Event containing the contract specifics and raw log

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
func (it *ERCTwentyDetailedTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERCTwentyDetailedTransfer)
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
		it.Event = new(ERCTwentyDetailedTransfer)
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
func (it *ERCTwentyDetailedTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERCTwentyDetailedTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERCTwentyDetailedTransfer represents a Transfer event raised by the ERCTwentyDetailed contract.
type ERCTwentyDetailedTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERCTwentyDetailed *ERCTwentyDetailedFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERCTwentyDetailedTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERCTwentyDetailed.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERCTwentyDetailedTransferIterator{contract: _ERCTwentyDetailed.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERCTwentyDetailed *ERCTwentyDetailedFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERCTwentyDetailedTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERCTwentyDetailed.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERCTwentyDetailedTransfer)
				if err := _ERCTwentyDetailed.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERCTwentyDetailed *ERCTwentyDetailedFilterer) ParseTransfer(log types.Log) (*ERCTwentyDetailedTransfer, error) {
	event := new(ERCTwentyDetailedTransfer)
	if err := _ERCTwentyDetailed.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
