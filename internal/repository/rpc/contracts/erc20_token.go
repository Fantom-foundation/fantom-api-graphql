// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// ERCTwentyABI is the input ABI used to generate the binding from.
const ERCTwentyABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MinterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MinterRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addPauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isMinter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPauser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renouncePauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERCTwenty is an auto generated Go binding around an Ethereum contract.
type ERCTwenty struct {
	ERCTwentyCaller     // Read-only binding to the contract
	ERCTwentyTransactor // Write-only binding to the contract
	ERCTwentyFilterer   // Log filterer for contract events
}

// ERCTwentyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERCTwentyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERCTwentyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERCTwentyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERCTwentyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERCTwentyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERCTwentySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERCTwentySession struct {
	Contract     *ERCTwenty        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERCTwentyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERCTwentyCallerSession struct {
	Contract *ERCTwentyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ERCTwentyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERCTwentyTransactorSession struct {
	Contract     *ERCTwentyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ERCTwentyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERCTwentyRaw struct {
	Contract *ERCTwenty // Generic contract binding to access the raw methods on
}

// ERCTwentyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERCTwentyCallerRaw struct {
	Contract *ERCTwentyCaller // Generic read-only contract binding to access the raw methods on
}

// ERCTwentyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERCTwentyTransactorRaw struct {
	Contract *ERCTwentyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERCTwenty creates a new instance of ERCTwenty, bound to a specific deployed contract.
func NewERCTwenty(address common.Address, backend bind.ContractBackend) (*ERCTwenty, error) {
	contract, err := bindERCTwenty(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERCTwenty{ERCTwentyCaller: ERCTwentyCaller{contract: contract}, ERCTwentyTransactor: ERCTwentyTransactor{contract: contract}, ERCTwentyFilterer: ERCTwentyFilterer{contract: contract}}, nil
}

// NewERCTwentyCaller creates a new read-only instance of ERCTwenty, bound to a specific deployed contract.
func NewERCTwentyCaller(address common.Address, caller bind.ContractCaller) (*ERCTwentyCaller, error) {
	contract, err := bindERCTwenty(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERCTwentyCaller{contract: contract}, nil
}

// NewERCTwentyTransactor creates a new write-only instance of ERCTwenty, bound to a specific deployed contract.
func NewERCTwentyTransactor(address common.Address, transactor bind.ContractTransactor) (*ERCTwentyTransactor, error) {
	contract, err := bindERCTwenty(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERCTwentyTransactor{contract: contract}, nil
}

// NewERCTwentyFilterer creates a new log filterer instance of ERCTwenty, bound to a specific deployed contract.
func NewERCTwentyFilterer(address common.Address, filterer bind.ContractFilterer) (*ERCTwentyFilterer, error) {
	contract, err := bindERCTwenty(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERCTwentyFilterer{contract: contract}, nil
}

// bindERCTwenty binds a generic wrapper to an already deployed contract.
func bindERCTwenty(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERCTwentyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERCTwenty *ERCTwentyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERCTwenty.Contract.ERCTwentyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERCTwenty *ERCTwentyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERCTwenty.Contract.ERCTwentyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERCTwenty *ERCTwentyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERCTwenty.Contract.ERCTwentyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERCTwenty *ERCTwentyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERCTwenty.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERCTwenty *ERCTwentyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERCTwenty.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERCTwenty *ERCTwentyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERCTwenty.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERCTwenty *ERCTwentyCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERCTwenty.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERCTwenty *ERCTwentySession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERCTwenty.Contract.Allowance(&_ERCTwenty.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERCTwenty *ERCTwentyCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERCTwenty.Contract.Allowance(&_ERCTwenty.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERCTwenty *ERCTwentyCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERCTwenty.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERCTwenty *ERCTwentySession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERCTwenty.Contract.BalanceOf(&_ERCTwenty.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERCTwenty *ERCTwentyCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERCTwenty.Contract.BalanceOf(&_ERCTwenty.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERCTwenty *ERCTwentyCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERCTwenty.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERCTwenty *ERCTwentySession) Decimals() (uint8, error) {
	return _ERCTwenty.Contract.Decimals(&_ERCTwenty.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERCTwenty *ERCTwentyCallerSession) Decimals() (uint8, error) {
	return _ERCTwenty.Contract.Decimals(&_ERCTwenty.CallOpts)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_ERCTwenty *ERCTwentyCaller) IsMinter(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _ERCTwenty.contract.Call(opts, &out, "isMinter", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_ERCTwenty *ERCTwentySession) IsMinter(account common.Address) (bool, error) {
	return _ERCTwenty.Contract.IsMinter(&_ERCTwenty.CallOpts, account)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_ERCTwenty *ERCTwentyCallerSession) IsMinter(account common.Address) (bool, error) {
	return _ERCTwenty.Contract.IsMinter(&_ERCTwenty.CallOpts, account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_ERCTwenty *ERCTwentyCaller) IsPauser(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _ERCTwenty.contract.Call(opts, &out, "isPauser", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_ERCTwenty *ERCTwentySession) IsPauser(account common.Address) (bool, error) {
	return _ERCTwenty.Contract.IsPauser(&_ERCTwenty.CallOpts, account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_ERCTwenty *ERCTwentyCallerSession) IsPauser(account common.Address) (bool, error) {
	return _ERCTwenty.Contract.IsPauser(&_ERCTwenty.CallOpts, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERCTwenty *ERCTwentyCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERCTwenty.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERCTwenty *ERCTwentySession) Name() (string, error) {
	return _ERCTwenty.Contract.Name(&_ERCTwenty.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERCTwenty *ERCTwentyCallerSession) Name() (string, error) {
	return _ERCTwenty.Contract.Name(&_ERCTwenty.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERCTwenty *ERCTwentyCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ERCTwenty.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERCTwenty *ERCTwentySession) Paused() (bool, error) {
	return _ERCTwenty.Contract.Paused(&_ERCTwenty.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERCTwenty *ERCTwentyCallerSession) Paused() (bool, error) {
	return _ERCTwenty.Contract.Paused(&_ERCTwenty.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERCTwenty *ERCTwentyCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERCTwenty.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERCTwenty *ERCTwentySession) Symbol() (string, error) {
	return _ERCTwenty.Contract.Symbol(&_ERCTwenty.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERCTwenty *ERCTwentyCallerSession) Symbol() (string, error) {
	return _ERCTwenty.Contract.Symbol(&_ERCTwenty.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERCTwenty *ERCTwentyCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERCTwenty.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERCTwenty *ERCTwentySession) TotalSupply() (*big.Int, error) {
	return _ERCTwenty.Contract.TotalSupply(&_ERCTwenty.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERCTwenty *ERCTwentyCallerSession) TotalSupply() (*big.Int, error) {
	return _ERCTwenty.Contract.TotalSupply(&_ERCTwenty.CallOpts)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_ERCTwenty *ERCTwentyTransactor) AddMinter(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ERCTwenty.contract.Transact(opts, "addMinter", account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_ERCTwenty *ERCTwentySession) AddMinter(account common.Address) (*types.Transaction, error) {
	return _ERCTwenty.Contract.AddMinter(&_ERCTwenty.TransactOpts, account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_ERCTwenty *ERCTwentyTransactorSession) AddMinter(account common.Address) (*types.Transaction, error) {
	return _ERCTwenty.Contract.AddMinter(&_ERCTwenty.TransactOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_ERCTwenty *ERCTwentyTransactor) AddPauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ERCTwenty.contract.Transact(opts, "addPauser", account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_ERCTwenty *ERCTwentySession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _ERCTwenty.Contract.AddPauser(&_ERCTwenty.TransactOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_ERCTwenty *ERCTwentyTransactorSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _ERCTwenty.Contract.AddPauser(&_ERCTwenty.TransactOpts, account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERCTwenty *ERCTwentyTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERCTwenty.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERCTwenty *ERCTwentySession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERCTwenty.Contract.Approve(&_ERCTwenty.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERCTwenty *ERCTwentyTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERCTwenty.Contract.Approve(&_ERCTwenty.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ERCTwenty *ERCTwentyTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ERCTwenty.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ERCTwenty *ERCTwentySession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ERCTwenty.Contract.Burn(&_ERCTwenty.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ERCTwenty *ERCTwentyTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ERCTwenty.Contract.Burn(&_ERCTwenty.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ERCTwenty *ERCTwentyTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERCTwenty.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ERCTwenty *ERCTwentySession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERCTwenty.Contract.BurnFrom(&_ERCTwenty.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ERCTwenty *ERCTwentyTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERCTwenty.Contract.BurnFrom(&_ERCTwenty.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERCTwenty *ERCTwentyTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERCTwenty.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERCTwenty *ERCTwentySession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERCTwenty.Contract.DecreaseAllowance(&_ERCTwenty.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERCTwenty *ERCTwentyTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERCTwenty.Contract.DecreaseAllowance(&_ERCTwenty.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERCTwenty *ERCTwentyTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERCTwenty.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERCTwenty *ERCTwentySession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERCTwenty.Contract.IncreaseAllowance(&_ERCTwenty.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERCTwenty *ERCTwentyTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERCTwenty.Contract.IncreaseAllowance(&_ERCTwenty.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool)
func (_ERCTwenty *ERCTwentyTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERCTwenty.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool)
func (_ERCTwenty *ERCTwentySession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERCTwenty.Contract.Mint(&_ERCTwenty.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool)
func (_ERCTwenty *ERCTwentyTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERCTwenty.Contract.Mint(&_ERCTwenty.TransactOpts, account, amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ERCTwenty *ERCTwentyTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERCTwenty.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ERCTwenty *ERCTwentySession) Pause() (*types.Transaction, error) {
	return _ERCTwenty.Contract.Pause(&_ERCTwenty.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ERCTwenty *ERCTwentyTransactorSession) Pause() (*types.Transaction, error) {
	return _ERCTwenty.Contract.Pause(&_ERCTwenty.TransactOpts)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_ERCTwenty *ERCTwentyTransactor) RenounceMinter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERCTwenty.contract.Transact(opts, "renounceMinter")
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_ERCTwenty *ERCTwentySession) RenounceMinter() (*types.Transaction, error) {
	return _ERCTwenty.Contract.RenounceMinter(&_ERCTwenty.TransactOpts)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_ERCTwenty *ERCTwentyTransactorSession) RenounceMinter() (*types.Transaction, error) {
	return _ERCTwenty.Contract.RenounceMinter(&_ERCTwenty.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_ERCTwenty *ERCTwentyTransactor) RenouncePauser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERCTwenty.contract.Transact(opts, "renouncePauser")
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_ERCTwenty *ERCTwentySession) RenouncePauser() (*types.Transaction, error) {
	return _ERCTwenty.Contract.RenouncePauser(&_ERCTwenty.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_ERCTwenty *ERCTwentyTransactorSession) RenouncePauser() (*types.Transaction, error) {
	return _ERCTwenty.Contract.RenouncePauser(&_ERCTwenty.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERCTwenty *ERCTwentyTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERCTwenty.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERCTwenty *ERCTwentySession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERCTwenty.Contract.Transfer(&_ERCTwenty.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERCTwenty *ERCTwentyTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERCTwenty.Contract.Transfer(&_ERCTwenty.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERCTwenty *ERCTwentyTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERCTwenty.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERCTwenty *ERCTwentySession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERCTwenty.Contract.TransferFrom(&_ERCTwenty.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERCTwenty *ERCTwentyTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERCTwenty.Contract.TransferFrom(&_ERCTwenty.TransactOpts, from, to, value)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ERCTwenty *ERCTwentyTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERCTwenty.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ERCTwenty *ERCTwentySession) Unpause() (*types.Transaction, error) {
	return _ERCTwenty.Contract.Unpause(&_ERCTwenty.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ERCTwenty *ERCTwentyTransactorSession) Unpause() (*types.Transaction, error) {
	return _ERCTwenty.Contract.Unpause(&_ERCTwenty.TransactOpts)
}

// ERCTwentyApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERCTwenty contract.
type ERCTwentyApprovalIterator struct {
	Event *ERCTwentyApproval // Event containing the contract specifics and raw log

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
func (it *ERCTwentyApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERCTwentyApproval)
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
		it.Event = new(ERCTwentyApproval)
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
func (it *ERCTwentyApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERCTwentyApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERCTwentyApproval represents a Approval event raised by the ERCTwenty contract.
type ERCTwentyApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERCTwenty *ERCTwentyFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERCTwentyApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERCTwenty.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERCTwentyApprovalIterator{contract: _ERCTwenty.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERCTwenty *ERCTwentyFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERCTwentyApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERCTwenty.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERCTwentyApproval)
				if err := _ERCTwenty.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ERCTwenty *ERCTwentyFilterer) ParseApproval(log types.Log) (*ERCTwentyApproval, error) {
	event := new(ERCTwentyApproval)
	if err := _ERCTwenty.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERCTwentyMinterAddedIterator is returned from FilterMinterAdded and is used to iterate over the raw logs and unpacked data for MinterAdded events raised by the ERCTwenty contract.
type ERCTwentyMinterAddedIterator struct {
	Event *ERCTwentyMinterAdded // Event containing the contract specifics and raw log

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
func (it *ERCTwentyMinterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERCTwentyMinterAdded)
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
		it.Event = new(ERCTwentyMinterAdded)
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
func (it *ERCTwentyMinterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERCTwentyMinterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERCTwentyMinterAdded represents a MinterAdded event raised by the ERCTwenty contract.
type ERCTwentyMinterAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterAdded is a free log retrieval operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_ERCTwenty *ERCTwentyFilterer) FilterMinterAdded(opts *bind.FilterOpts, account []common.Address) (*ERCTwentyMinterAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ERCTwenty.contract.FilterLogs(opts, "MinterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &ERCTwentyMinterAddedIterator{contract: _ERCTwenty.contract, event: "MinterAdded", logs: logs, sub: sub}, nil
}

// WatchMinterAdded is a free log subscription operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_ERCTwenty *ERCTwentyFilterer) WatchMinterAdded(opts *bind.WatchOpts, sink chan<- *ERCTwentyMinterAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ERCTwenty.contract.WatchLogs(opts, "MinterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERCTwentyMinterAdded)
				if err := _ERCTwenty.contract.UnpackLog(event, "MinterAdded", log); err != nil {
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

// ParseMinterAdded is a log parse operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_ERCTwenty *ERCTwentyFilterer) ParseMinterAdded(log types.Log) (*ERCTwentyMinterAdded, error) {
	event := new(ERCTwentyMinterAdded)
	if err := _ERCTwenty.contract.UnpackLog(event, "MinterAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERCTwentyMinterRemovedIterator is returned from FilterMinterRemoved and is used to iterate over the raw logs and unpacked data for MinterRemoved events raised by the ERCTwenty contract.
type ERCTwentyMinterRemovedIterator struct {
	Event *ERCTwentyMinterRemoved // Event containing the contract specifics and raw log

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
func (it *ERCTwentyMinterRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERCTwentyMinterRemoved)
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
		it.Event = new(ERCTwentyMinterRemoved)
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
func (it *ERCTwentyMinterRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERCTwentyMinterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERCTwentyMinterRemoved represents a MinterRemoved event raised by the ERCTwenty contract.
type ERCTwentyMinterRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterRemoved is a free log retrieval operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_ERCTwenty *ERCTwentyFilterer) FilterMinterRemoved(opts *bind.FilterOpts, account []common.Address) (*ERCTwentyMinterRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ERCTwenty.contract.FilterLogs(opts, "MinterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &ERCTwentyMinterRemovedIterator{contract: _ERCTwenty.contract, event: "MinterRemoved", logs: logs, sub: sub}, nil
}

// WatchMinterRemoved is a free log subscription operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_ERCTwenty *ERCTwentyFilterer) WatchMinterRemoved(opts *bind.WatchOpts, sink chan<- *ERCTwentyMinterRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ERCTwenty.contract.WatchLogs(opts, "MinterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERCTwentyMinterRemoved)
				if err := _ERCTwenty.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
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

// ParseMinterRemoved is a log parse operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_ERCTwenty *ERCTwentyFilterer) ParseMinterRemoved(log types.Log) (*ERCTwentyMinterRemoved, error) {
	event := new(ERCTwentyMinterRemoved)
	if err := _ERCTwenty.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERCTwentyPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ERCTwenty contract.
type ERCTwentyPausedIterator struct {
	Event *ERCTwentyPaused // Event containing the contract specifics and raw log

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
func (it *ERCTwentyPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERCTwentyPaused)
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
		it.Event = new(ERCTwentyPaused)
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
func (it *ERCTwentyPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERCTwentyPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERCTwentyPaused represents a Paused event raised by the ERCTwenty contract.
type ERCTwentyPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ERCTwenty *ERCTwentyFilterer) FilterPaused(opts *bind.FilterOpts) (*ERCTwentyPausedIterator, error) {

	logs, sub, err := _ERCTwenty.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ERCTwentyPausedIterator{contract: _ERCTwenty.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ERCTwenty *ERCTwentyFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ERCTwentyPaused) (event.Subscription, error) {

	logs, sub, err := _ERCTwenty.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERCTwentyPaused)
				if err := _ERCTwenty.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ERCTwenty *ERCTwentyFilterer) ParsePaused(log types.Log) (*ERCTwentyPaused, error) {
	event := new(ERCTwentyPaused)
	if err := _ERCTwenty.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERCTwentyPauserAddedIterator is returned from FilterPauserAdded and is used to iterate over the raw logs and unpacked data for PauserAdded events raised by the ERCTwenty contract.
type ERCTwentyPauserAddedIterator struct {
	Event *ERCTwentyPauserAdded // Event containing the contract specifics and raw log

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
func (it *ERCTwentyPauserAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERCTwentyPauserAdded)
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
		it.Event = new(ERCTwentyPauserAdded)
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
func (it *ERCTwentyPauserAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERCTwentyPauserAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERCTwentyPauserAdded represents a PauserAdded event raised by the ERCTwenty contract.
type ERCTwentyPauserAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserAdded is a free log retrieval operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address indexed account)
func (_ERCTwenty *ERCTwentyFilterer) FilterPauserAdded(opts *bind.FilterOpts, account []common.Address) (*ERCTwentyPauserAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ERCTwenty.contract.FilterLogs(opts, "PauserAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &ERCTwentyPauserAddedIterator{contract: _ERCTwenty.contract, event: "PauserAdded", logs: logs, sub: sub}, nil
}

// WatchPauserAdded is a free log subscription operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address indexed account)
func (_ERCTwenty *ERCTwentyFilterer) WatchPauserAdded(opts *bind.WatchOpts, sink chan<- *ERCTwentyPauserAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ERCTwenty.contract.WatchLogs(opts, "PauserAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERCTwentyPauserAdded)
				if err := _ERCTwenty.contract.UnpackLog(event, "PauserAdded", log); err != nil {
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

// ParsePauserAdded is a log parse operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address indexed account)
func (_ERCTwenty *ERCTwentyFilterer) ParsePauserAdded(log types.Log) (*ERCTwentyPauserAdded, error) {
	event := new(ERCTwentyPauserAdded)
	if err := _ERCTwenty.contract.UnpackLog(event, "PauserAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERCTwentyPauserRemovedIterator is returned from FilterPauserRemoved and is used to iterate over the raw logs and unpacked data for PauserRemoved events raised by the ERCTwenty contract.
type ERCTwentyPauserRemovedIterator struct {
	Event *ERCTwentyPauserRemoved // Event containing the contract specifics and raw log

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
func (it *ERCTwentyPauserRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERCTwentyPauserRemoved)
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
		it.Event = new(ERCTwentyPauserRemoved)
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
func (it *ERCTwentyPauserRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERCTwentyPauserRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERCTwentyPauserRemoved represents a PauserRemoved event raised by the ERCTwenty contract.
type ERCTwentyPauserRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserRemoved is a free log retrieval operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address indexed account)
func (_ERCTwenty *ERCTwentyFilterer) FilterPauserRemoved(opts *bind.FilterOpts, account []common.Address) (*ERCTwentyPauserRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ERCTwenty.contract.FilterLogs(opts, "PauserRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &ERCTwentyPauserRemovedIterator{contract: _ERCTwenty.contract, event: "PauserRemoved", logs: logs, sub: sub}, nil
}

// WatchPauserRemoved is a free log subscription operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address indexed account)
func (_ERCTwenty *ERCTwentyFilterer) WatchPauserRemoved(opts *bind.WatchOpts, sink chan<- *ERCTwentyPauserRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ERCTwenty.contract.WatchLogs(opts, "PauserRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERCTwentyPauserRemoved)
				if err := _ERCTwenty.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
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

// ParsePauserRemoved is a log parse operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address indexed account)
func (_ERCTwenty *ERCTwentyFilterer) ParsePauserRemoved(log types.Log) (*ERCTwentyPauserRemoved, error) {
	event := new(ERCTwentyPauserRemoved)
	if err := _ERCTwenty.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERCTwentyTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERCTwenty contract.
type ERCTwentyTransferIterator struct {
	Event *ERCTwentyTransfer // Event containing the contract specifics and raw log

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
func (it *ERCTwentyTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERCTwentyTransfer)
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
		it.Event = new(ERCTwentyTransfer)
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
func (it *ERCTwentyTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERCTwentyTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERCTwentyTransfer represents a Transfer event raised by the ERCTwenty contract.
type ERCTwentyTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERCTwenty *ERCTwentyFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERCTwentyTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERCTwenty.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERCTwentyTransferIterator{contract: _ERCTwenty.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERCTwenty *ERCTwentyFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERCTwentyTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERCTwenty.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERCTwentyTransfer)
				if err := _ERCTwenty.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ERCTwenty *ERCTwentyFilterer) ParseTransfer(log types.Log) (*ERCTwentyTransfer, error) {
	event := new(ERCTwentyTransfer)
	if err := _ERCTwenty.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERCTwentyUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ERCTwenty contract.
type ERCTwentyUnpausedIterator struct {
	Event *ERCTwentyUnpaused // Event containing the contract specifics and raw log

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
func (it *ERCTwentyUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERCTwentyUnpaused)
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
		it.Event = new(ERCTwentyUnpaused)
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
func (it *ERCTwentyUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERCTwentyUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERCTwentyUnpaused represents a Unpaused event raised by the ERCTwenty contract.
type ERCTwentyUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ERCTwenty *ERCTwentyFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ERCTwentyUnpausedIterator, error) {

	logs, sub, err := _ERCTwenty.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ERCTwentyUnpausedIterator{contract: _ERCTwenty.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ERCTwenty *ERCTwentyFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ERCTwentyUnpaused) (event.Subscription, error) {

	logs, sub, err := _ERCTwenty.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERCTwentyUnpaused)
				if err := _ERCTwenty.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ERCTwenty *ERCTwentyFilterer) ParseUnpaused(log types.Log) (*ERCTwentyUnpaused, error) {
	event := new(ERCTwentyUnpaused)
	if err := _ERCTwenty.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	return event, nil
}
