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

// ErcWrappedFtmMetaData contains all meta data concerning the ErcWrappedFtm contract.
var ErcWrappedFtmMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_INVALID_ZERO_VALUE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_NO_ERROR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addPauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPauser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renouncePauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ErcWrappedFtmABI is the input ABI used to generate the binding from.
// Deprecated: Use ErcWrappedFtmMetaData.ABI instead.
var ErcWrappedFtmABI = ErcWrappedFtmMetaData.ABI

// ErcWrappedFtm is an auto generated Go binding around an Ethereum contract.
type ErcWrappedFtm struct {
	ErcWrappedFtmCaller     // Read-only binding to the contract
	ErcWrappedFtmTransactor // Write-only binding to the contract
	ErcWrappedFtmFilterer   // Log filterer for contract events
}

// ErcWrappedFtmCaller is an auto generated read-only Go binding around an Ethereum contract.
type ErcWrappedFtmCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ErcWrappedFtmTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ErcWrappedFtmTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ErcWrappedFtmFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ErcWrappedFtmFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ErcWrappedFtmSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ErcWrappedFtmSession struct {
	Contract     *ErcWrappedFtm    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ErcWrappedFtmCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ErcWrappedFtmCallerSession struct {
	Contract *ErcWrappedFtmCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ErcWrappedFtmTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ErcWrappedFtmTransactorSession struct {
	Contract     *ErcWrappedFtmTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ErcWrappedFtmRaw is an auto generated low-level Go binding around an Ethereum contract.
type ErcWrappedFtmRaw struct {
	Contract *ErcWrappedFtm // Generic contract binding to access the raw methods on
}

// ErcWrappedFtmCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ErcWrappedFtmCallerRaw struct {
	Contract *ErcWrappedFtmCaller // Generic read-only contract binding to access the raw methods on
}

// ErcWrappedFtmTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ErcWrappedFtmTransactorRaw struct {
	Contract *ErcWrappedFtmTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErcWrappedFtm creates a new instance of ErcWrappedFtm, bound to a specific deployed contract.
func NewErcWrappedFtm(address common.Address, backend bind.ContractBackend) (*ErcWrappedFtm, error) {
	contract, err := bindErcWrappedFtm(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ErcWrappedFtm{ErcWrappedFtmCaller: ErcWrappedFtmCaller{contract: contract}, ErcWrappedFtmTransactor: ErcWrappedFtmTransactor{contract: contract}, ErcWrappedFtmFilterer: ErcWrappedFtmFilterer{contract: contract}}, nil
}

// NewErcWrappedFtmCaller creates a new read-only instance of ErcWrappedFtm, bound to a specific deployed contract.
func NewErcWrappedFtmCaller(address common.Address, caller bind.ContractCaller) (*ErcWrappedFtmCaller, error) {
	contract, err := bindErcWrappedFtm(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ErcWrappedFtmCaller{contract: contract}, nil
}

// NewErcWrappedFtmTransactor creates a new write-only instance of ErcWrappedFtm, bound to a specific deployed contract.
func NewErcWrappedFtmTransactor(address common.Address, transactor bind.ContractTransactor) (*ErcWrappedFtmTransactor, error) {
	contract, err := bindErcWrappedFtm(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ErcWrappedFtmTransactor{contract: contract}, nil
}

// NewErcWrappedFtmFilterer creates a new log filterer instance of ErcWrappedFtm, bound to a specific deployed contract.
func NewErcWrappedFtmFilterer(address common.Address, filterer bind.ContractFilterer) (*ErcWrappedFtmFilterer, error) {
	contract, err := bindErcWrappedFtm(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ErcWrappedFtmFilterer{contract: contract}, nil
}

// bindErcWrappedFtm binds a generic wrapper to an already deployed contract.
func bindErcWrappedFtm(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ErcWrappedFtmABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ErcWrappedFtm *ErcWrappedFtmRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ErcWrappedFtm.Contract.ErcWrappedFtmCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ErcWrappedFtm *ErcWrappedFtmRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ErcWrappedFtm.Contract.ErcWrappedFtmTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ErcWrappedFtm *ErcWrappedFtmRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ErcWrappedFtm.Contract.ErcWrappedFtmTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ErcWrappedFtm *ErcWrappedFtmCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ErcWrappedFtm.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ErcWrappedFtm *ErcWrappedFtmTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ErcWrappedFtm.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ErcWrappedFtm *ErcWrappedFtmTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ErcWrappedFtm.Contract.contract.Transact(opts, method, params...)
}

// ERRINVALIDZEROVALUE is a free data retrieval call binding the contract method 0x6d7497b3.
//
// Solidity: function ERR_INVALID_ZERO_VALUE() view returns(uint256)
func (_ErcWrappedFtm *ErcWrappedFtmCaller) ERRINVALIDZEROVALUE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ErcWrappedFtm.contract.Call(opts, &out, "ERR_INVALID_ZERO_VALUE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ERRINVALIDZEROVALUE is a free data retrieval call binding the contract method 0x6d7497b3.
//
// Solidity: function ERR_INVALID_ZERO_VALUE() view returns(uint256)
func (_ErcWrappedFtm *ErcWrappedFtmSession) ERRINVALIDZEROVALUE() (*big.Int, error) {
	return _ErcWrappedFtm.Contract.ERRINVALIDZEROVALUE(&_ErcWrappedFtm.CallOpts)
}

// ERRINVALIDZEROVALUE is a free data retrieval call binding the contract method 0x6d7497b3.
//
// Solidity: function ERR_INVALID_ZERO_VALUE() view returns(uint256)
func (_ErcWrappedFtm *ErcWrappedFtmCallerSession) ERRINVALIDZEROVALUE() (*big.Int, error) {
	return _ErcWrappedFtm.Contract.ERRINVALIDZEROVALUE(&_ErcWrappedFtm.CallOpts)
}

// ERRNOERROR is a free data retrieval call binding the contract method 0x35052d6e.
//
// Solidity: function ERR_NO_ERROR() view returns(uint256)
func (_ErcWrappedFtm *ErcWrappedFtmCaller) ERRNOERROR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ErcWrappedFtm.contract.Call(opts, &out, "ERR_NO_ERROR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ERRNOERROR is a free data retrieval call binding the contract method 0x35052d6e.
//
// Solidity: function ERR_NO_ERROR() view returns(uint256)
func (_ErcWrappedFtm *ErcWrappedFtmSession) ERRNOERROR() (*big.Int, error) {
	return _ErcWrappedFtm.Contract.ERRNOERROR(&_ErcWrappedFtm.CallOpts)
}

// ERRNOERROR is a free data retrieval call binding the contract method 0x35052d6e.
//
// Solidity: function ERR_NO_ERROR() view returns(uint256)
func (_ErcWrappedFtm *ErcWrappedFtmCallerSession) ERRNOERROR() (*big.Int, error) {
	return _ErcWrappedFtm.Contract.ERRNOERROR(&_ErcWrappedFtm.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ErcWrappedFtm *ErcWrappedFtmCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ErcWrappedFtm.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ErcWrappedFtm *ErcWrappedFtmSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ErcWrappedFtm.Contract.Allowance(&_ErcWrappedFtm.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ErcWrappedFtm *ErcWrappedFtmCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ErcWrappedFtm.Contract.Allowance(&_ErcWrappedFtm.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ErcWrappedFtm *ErcWrappedFtmCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ErcWrappedFtm.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ErcWrappedFtm *ErcWrappedFtmSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ErcWrappedFtm.Contract.BalanceOf(&_ErcWrappedFtm.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ErcWrappedFtm *ErcWrappedFtmCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ErcWrappedFtm.Contract.BalanceOf(&_ErcWrappedFtm.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ErcWrappedFtm *ErcWrappedFtmCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ErcWrappedFtm.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ErcWrappedFtm *ErcWrappedFtmSession) Decimals() (uint8, error) {
	return _ErcWrappedFtm.Contract.Decimals(&_ErcWrappedFtm.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ErcWrappedFtm *ErcWrappedFtmCallerSession) Decimals() (uint8, error) {
	return _ErcWrappedFtm.Contract.Decimals(&_ErcWrappedFtm.CallOpts)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_ErcWrappedFtm *ErcWrappedFtmCaller) IsPauser(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _ErcWrappedFtm.contract.Call(opts, &out, "isPauser", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_ErcWrappedFtm *ErcWrappedFtmSession) IsPauser(account common.Address) (bool, error) {
	return _ErcWrappedFtm.Contract.IsPauser(&_ErcWrappedFtm.CallOpts, account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_ErcWrappedFtm *ErcWrappedFtmCallerSession) IsPauser(account common.Address) (bool, error) {
	return _ErcWrappedFtm.Contract.IsPauser(&_ErcWrappedFtm.CallOpts, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ErcWrappedFtm *ErcWrappedFtmCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ErcWrappedFtm.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ErcWrappedFtm *ErcWrappedFtmSession) Name() (string, error) {
	return _ErcWrappedFtm.Contract.Name(&_ErcWrappedFtm.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ErcWrappedFtm *ErcWrappedFtmCallerSession) Name() (string, error) {
	return _ErcWrappedFtm.Contract.Name(&_ErcWrappedFtm.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ErcWrappedFtm *ErcWrappedFtmCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ErcWrappedFtm.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ErcWrappedFtm *ErcWrappedFtmSession) Paused() (bool, error) {
	return _ErcWrappedFtm.Contract.Paused(&_ErcWrappedFtm.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ErcWrappedFtm *ErcWrappedFtmCallerSession) Paused() (bool, error) {
	return _ErcWrappedFtm.Contract.Paused(&_ErcWrappedFtm.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ErcWrappedFtm *ErcWrappedFtmCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ErcWrappedFtm.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ErcWrappedFtm *ErcWrappedFtmSession) Symbol() (string, error) {
	return _ErcWrappedFtm.Contract.Symbol(&_ErcWrappedFtm.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ErcWrappedFtm *ErcWrappedFtmCallerSession) Symbol() (string, error) {
	return _ErcWrappedFtm.Contract.Symbol(&_ErcWrappedFtm.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ErcWrappedFtm *ErcWrappedFtmCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ErcWrappedFtm.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ErcWrappedFtm *ErcWrappedFtmSession) TotalSupply() (*big.Int, error) {
	return _ErcWrappedFtm.Contract.TotalSupply(&_ErcWrappedFtm.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ErcWrappedFtm *ErcWrappedFtmCallerSession) TotalSupply() (*big.Int, error) {
	return _ErcWrappedFtm.Contract.TotalSupply(&_ErcWrappedFtm.CallOpts)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_ErcWrappedFtm *ErcWrappedFtmTransactor) AddPauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ErcWrappedFtm.contract.Transact(opts, "addPauser", account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_ErcWrappedFtm *ErcWrappedFtmSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _ErcWrappedFtm.Contract.AddPauser(&_ErcWrappedFtm.TransactOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_ErcWrappedFtm *ErcWrappedFtmTransactorSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _ErcWrappedFtm.Contract.AddPauser(&_ErcWrappedFtm.TransactOpts, account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ErcWrappedFtm *ErcWrappedFtmTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ErcWrappedFtm.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ErcWrappedFtm *ErcWrappedFtmSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ErcWrappedFtm.Contract.Approve(&_ErcWrappedFtm.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ErcWrappedFtm *ErcWrappedFtmTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ErcWrappedFtm.Contract.Approve(&_ErcWrappedFtm.TransactOpts, spender, value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ErcWrappedFtm *ErcWrappedFtmTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ErcWrappedFtm.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ErcWrappedFtm *ErcWrappedFtmSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ErcWrappedFtm.Contract.DecreaseAllowance(&_ErcWrappedFtm.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ErcWrappedFtm *ErcWrappedFtmTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ErcWrappedFtm.Contract.DecreaseAllowance(&_ErcWrappedFtm.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns(uint256)
func (_ErcWrappedFtm *ErcWrappedFtmTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ErcWrappedFtm.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns(uint256)
func (_ErcWrappedFtm *ErcWrappedFtmSession) Deposit() (*types.Transaction, error) {
	return _ErcWrappedFtm.Contract.Deposit(&_ErcWrappedFtm.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns(uint256)
func (_ErcWrappedFtm *ErcWrappedFtmTransactorSession) Deposit() (*types.Transaction, error) {
	return _ErcWrappedFtm.Contract.Deposit(&_ErcWrappedFtm.TransactOpts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ErcWrappedFtm *ErcWrappedFtmTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ErcWrappedFtm.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ErcWrappedFtm *ErcWrappedFtmSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ErcWrappedFtm.Contract.IncreaseAllowance(&_ErcWrappedFtm.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ErcWrappedFtm *ErcWrappedFtmTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ErcWrappedFtm.Contract.IncreaseAllowance(&_ErcWrappedFtm.TransactOpts, spender, addedValue)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ErcWrappedFtm *ErcWrappedFtmTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ErcWrappedFtm.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ErcWrappedFtm *ErcWrappedFtmSession) Pause() (*types.Transaction, error) {
	return _ErcWrappedFtm.Contract.Pause(&_ErcWrappedFtm.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ErcWrappedFtm *ErcWrappedFtmTransactorSession) Pause() (*types.Transaction, error) {
	return _ErcWrappedFtm.Contract.Pause(&_ErcWrappedFtm.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_ErcWrappedFtm *ErcWrappedFtmTransactor) RenouncePauser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ErcWrappedFtm.contract.Transact(opts, "renouncePauser")
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_ErcWrappedFtm *ErcWrappedFtmSession) RenouncePauser() (*types.Transaction, error) {
	return _ErcWrappedFtm.Contract.RenouncePauser(&_ErcWrappedFtm.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_ErcWrappedFtm *ErcWrappedFtmTransactorSession) RenouncePauser() (*types.Transaction, error) {
	return _ErcWrappedFtm.Contract.RenouncePauser(&_ErcWrappedFtm.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ErcWrappedFtm *ErcWrappedFtmTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ErcWrappedFtm.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ErcWrappedFtm *ErcWrappedFtmSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ErcWrappedFtm.Contract.Transfer(&_ErcWrappedFtm.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ErcWrappedFtm *ErcWrappedFtmTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ErcWrappedFtm.Contract.Transfer(&_ErcWrappedFtm.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ErcWrappedFtm *ErcWrappedFtmTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ErcWrappedFtm.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ErcWrappedFtm *ErcWrappedFtmSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ErcWrappedFtm.Contract.TransferFrom(&_ErcWrappedFtm.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ErcWrappedFtm *ErcWrappedFtmTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ErcWrappedFtm.Contract.TransferFrom(&_ErcWrappedFtm.TransactOpts, from, to, value)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ErcWrappedFtm *ErcWrappedFtmTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ErcWrappedFtm.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ErcWrappedFtm *ErcWrappedFtmSession) Unpause() (*types.Transaction, error) {
	return _ErcWrappedFtm.Contract.Unpause(&_ErcWrappedFtm.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ErcWrappedFtm *ErcWrappedFtmTransactorSession) Unpause() (*types.Transaction, error) {
	return _ErcWrappedFtm.Contract.Unpause(&_ErcWrappedFtm.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns(uint256)
func (_ErcWrappedFtm *ErcWrappedFtmTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ErcWrappedFtm.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns(uint256)
func (_ErcWrappedFtm *ErcWrappedFtmSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _ErcWrappedFtm.Contract.Withdraw(&_ErcWrappedFtm.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns(uint256)
func (_ErcWrappedFtm *ErcWrappedFtmTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _ErcWrappedFtm.Contract.Withdraw(&_ErcWrappedFtm.TransactOpts, amount)
}

// ErcWrappedFtmApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ErcWrappedFtm contract.
type ErcWrappedFtmApprovalIterator struct {
	Event *ErcWrappedFtmApproval // Event containing the contract specifics and raw log

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
func (it *ErcWrappedFtmApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ErcWrappedFtmApproval)
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
		it.Event = new(ErcWrappedFtmApproval)
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
func (it *ErcWrappedFtmApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ErcWrappedFtmApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ErcWrappedFtmApproval represents a Approval event raised by the ErcWrappedFtm contract.
type ErcWrappedFtmApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ErcWrappedFtm *ErcWrappedFtmFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ErcWrappedFtmApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ErcWrappedFtm.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ErcWrappedFtmApprovalIterator{contract: _ErcWrappedFtm.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ErcWrappedFtm *ErcWrappedFtmFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ErcWrappedFtmApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ErcWrappedFtm.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ErcWrappedFtmApproval)
				if err := _ErcWrappedFtm.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ErcWrappedFtm *ErcWrappedFtmFilterer) ParseApproval(log types.Log) (*ErcWrappedFtmApproval, error) {
	event := new(ErcWrappedFtmApproval)
	if err := _ErcWrappedFtm.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ErcWrappedFtmPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ErcWrappedFtm contract.
type ErcWrappedFtmPausedIterator struct {
	Event *ErcWrappedFtmPaused // Event containing the contract specifics and raw log

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
func (it *ErcWrappedFtmPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ErcWrappedFtmPaused)
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
		it.Event = new(ErcWrappedFtmPaused)
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
func (it *ErcWrappedFtmPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ErcWrappedFtmPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ErcWrappedFtmPaused represents a Paused event raised by the ErcWrappedFtm contract.
type ErcWrappedFtmPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ErcWrappedFtm *ErcWrappedFtmFilterer) FilterPaused(opts *bind.FilterOpts) (*ErcWrappedFtmPausedIterator, error) {

	logs, sub, err := _ErcWrappedFtm.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ErcWrappedFtmPausedIterator{contract: _ErcWrappedFtm.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ErcWrappedFtm *ErcWrappedFtmFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ErcWrappedFtmPaused) (event.Subscription, error) {

	logs, sub, err := _ErcWrappedFtm.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ErcWrappedFtmPaused)
				if err := _ErcWrappedFtm.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_ErcWrappedFtm *ErcWrappedFtmFilterer) ParsePaused(log types.Log) (*ErcWrappedFtmPaused, error) {
	event := new(ErcWrappedFtmPaused)
	if err := _ErcWrappedFtm.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ErcWrappedFtmPauserAddedIterator is returned from FilterPauserAdded and is used to iterate over the raw logs and unpacked data for PauserAdded events raised by the ErcWrappedFtm contract.
type ErcWrappedFtmPauserAddedIterator struct {
	Event *ErcWrappedFtmPauserAdded // Event containing the contract specifics and raw log

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
func (it *ErcWrappedFtmPauserAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ErcWrappedFtmPauserAdded)
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
		it.Event = new(ErcWrappedFtmPauserAdded)
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
func (it *ErcWrappedFtmPauserAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ErcWrappedFtmPauserAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ErcWrappedFtmPauserAdded represents a PauserAdded event raised by the ErcWrappedFtm contract.
type ErcWrappedFtmPauserAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserAdded is a free log retrieval operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address indexed account)
func (_ErcWrappedFtm *ErcWrappedFtmFilterer) FilterPauserAdded(opts *bind.FilterOpts, account []common.Address) (*ErcWrappedFtmPauserAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ErcWrappedFtm.contract.FilterLogs(opts, "PauserAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &ErcWrappedFtmPauserAddedIterator{contract: _ErcWrappedFtm.contract, event: "PauserAdded", logs: logs, sub: sub}, nil
}

// WatchPauserAdded is a free log subscription operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address indexed account)
func (_ErcWrappedFtm *ErcWrappedFtmFilterer) WatchPauserAdded(opts *bind.WatchOpts, sink chan<- *ErcWrappedFtmPauserAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ErcWrappedFtm.contract.WatchLogs(opts, "PauserAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ErcWrappedFtmPauserAdded)
				if err := _ErcWrappedFtm.contract.UnpackLog(event, "PauserAdded", log); err != nil {
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
func (_ErcWrappedFtm *ErcWrappedFtmFilterer) ParsePauserAdded(log types.Log) (*ErcWrappedFtmPauserAdded, error) {
	event := new(ErcWrappedFtmPauserAdded)
	if err := _ErcWrappedFtm.contract.UnpackLog(event, "PauserAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ErcWrappedFtmPauserRemovedIterator is returned from FilterPauserRemoved and is used to iterate over the raw logs and unpacked data for PauserRemoved events raised by the ErcWrappedFtm contract.
type ErcWrappedFtmPauserRemovedIterator struct {
	Event *ErcWrappedFtmPauserRemoved // Event containing the contract specifics and raw log

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
func (it *ErcWrappedFtmPauserRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ErcWrappedFtmPauserRemoved)
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
		it.Event = new(ErcWrappedFtmPauserRemoved)
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
func (it *ErcWrappedFtmPauserRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ErcWrappedFtmPauserRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ErcWrappedFtmPauserRemoved represents a PauserRemoved event raised by the ErcWrappedFtm contract.
type ErcWrappedFtmPauserRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserRemoved is a free log retrieval operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address indexed account)
func (_ErcWrappedFtm *ErcWrappedFtmFilterer) FilterPauserRemoved(opts *bind.FilterOpts, account []common.Address) (*ErcWrappedFtmPauserRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ErcWrappedFtm.contract.FilterLogs(opts, "PauserRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &ErcWrappedFtmPauserRemovedIterator{contract: _ErcWrappedFtm.contract, event: "PauserRemoved", logs: logs, sub: sub}, nil
}

// WatchPauserRemoved is a free log subscription operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address indexed account)
func (_ErcWrappedFtm *ErcWrappedFtmFilterer) WatchPauserRemoved(opts *bind.WatchOpts, sink chan<- *ErcWrappedFtmPauserRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ErcWrappedFtm.contract.WatchLogs(opts, "PauserRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ErcWrappedFtmPauserRemoved)
				if err := _ErcWrappedFtm.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
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
func (_ErcWrappedFtm *ErcWrappedFtmFilterer) ParsePauserRemoved(log types.Log) (*ErcWrappedFtmPauserRemoved, error) {
	event := new(ErcWrappedFtmPauserRemoved)
	if err := _ErcWrappedFtm.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ErcWrappedFtmTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ErcWrappedFtm contract.
type ErcWrappedFtmTransferIterator struct {
	Event *ErcWrappedFtmTransfer // Event containing the contract specifics and raw log

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
func (it *ErcWrappedFtmTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ErcWrappedFtmTransfer)
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
		it.Event = new(ErcWrappedFtmTransfer)
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
func (it *ErcWrappedFtmTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ErcWrappedFtmTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ErcWrappedFtmTransfer represents a Transfer event raised by the ErcWrappedFtm contract.
type ErcWrappedFtmTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ErcWrappedFtm *ErcWrappedFtmFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ErcWrappedFtmTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ErcWrappedFtm.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ErcWrappedFtmTransferIterator{contract: _ErcWrappedFtm.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ErcWrappedFtm *ErcWrappedFtmFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ErcWrappedFtmTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ErcWrappedFtm.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ErcWrappedFtmTransfer)
				if err := _ErcWrappedFtm.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ErcWrappedFtm *ErcWrappedFtmFilterer) ParseTransfer(log types.Log) (*ErcWrappedFtmTransfer, error) {
	event := new(ErcWrappedFtmTransfer)
	if err := _ErcWrappedFtm.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ErcWrappedFtmUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ErcWrappedFtm contract.
type ErcWrappedFtmUnpausedIterator struct {
	Event *ErcWrappedFtmUnpaused // Event containing the contract specifics and raw log

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
func (it *ErcWrappedFtmUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ErcWrappedFtmUnpaused)
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
		it.Event = new(ErcWrappedFtmUnpaused)
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
func (it *ErcWrappedFtmUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ErcWrappedFtmUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ErcWrappedFtmUnpaused represents a Unpaused event raised by the ErcWrappedFtm contract.
type ErcWrappedFtmUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ErcWrappedFtm *ErcWrappedFtmFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ErcWrappedFtmUnpausedIterator, error) {

	logs, sub, err := _ErcWrappedFtm.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ErcWrappedFtmUnpausedIterator{contract: _ErcWrappedFtm.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ErcWrappedFtm *ErcWrappedFtmFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ErcWrappedFtmUnpaused) (event.Subscription, error) {

	logs, sub, err := _ErcWrappedFtm.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ErcWrappedFtmUnpaused)
				if err := _ErcWrappedFtm.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_ErcWrappedFtm *ErcWrappedFtmFilterer) ParseUnpaused(log types.Log) (*ErcWrappedFtmUnpaused, error) {
	event := new(ErcWrappedFtmUnpaused)
	if err := _ErcWrappedFtm.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
