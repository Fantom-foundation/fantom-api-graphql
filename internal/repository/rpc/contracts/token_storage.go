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

// DeFiTokenStorageMetaData contains all meta data concerning the DeFiTokenStorage contract.
var DeFiTokenStorageMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"addressProvider\",\"outputs\":[{\"internalType\":\"contractIFantomMintAddressProvider\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"valueDustAdjustment\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressProvider\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_dustAdt\",\"type\":\"bool\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"tokenValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"total\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"totalOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"totalOfInc\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"totalOfDec\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"add\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"sub\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokensCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DeFiTokenStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use DeFiTokenStorageMetaData.ABI instead.
var DeFiTokenStorageABI = DeFiTokenStorageMetaData.ABI

// DeFiTokenStorage is an auto generated Go binding around an Ethereum contract.
type DeFiTokenStorage struct {
	DeFiTokenStorageCaller     // Read-only binding to the contract
	DeFiTokenStorageTransactor // Write-only binding to the contract
	DeFiTokenStorageFilterer   // Log filterer for contract events
}

// DeFiTokenStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type DeFiTokenStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeFiTokenStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DeFiTokenStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeFiTokenStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DeFiTokenStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeFiTokenStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DeFiTokenStorageSession struct {
	Contract     *DeFiTokenStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DeFiTokenStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DeFiTokenStorageCallerSession struct {
	Contract *DeFiTokenStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// DeFiTokenStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DeFiTokenStorageTransactorSession struct {
	Contract     *DeFiTokenStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// DeFiTokenStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type DeFiTokenStorageRaw struct {
	Contract *DeFiTokenStorage // Generic contract binding to access the raw methods on
}

// DeFiTokenStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DeFiTokenStorageCallerRaw struct {
	Contract *DeFiTokenStorageCaller // Generic read-only contract binding to access the raw methods on
}

// DeFiTokenStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DeFiTokenStorageTransactorRaw struct {
	Contract *DeFiTokenStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDeFiTokenStorage creates a new instance of DeFiTokenStorage, bound to a specific deployed contract.
func NewDeFiTokenStorage(address common.Address, backend bind.ContractBackend) (*DeFiTokenStorage, error) {
	contract, err := bindDeFiTokenStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DeFiTokenStorage{DeFiTokenStorageCaller: DeFiTokenStorageCaller{contract: contract}, DeFiTokenStorageTransactor: DeFiTokenStorageTransactor{contract: contract}, DeFiTokenStorageFilterer: DeFiTokenStorageFilterer{contract: contract}}, nil
}

// NewDeFiTokenStorageCaller creates a new read-only instance of DeFiTokenStorage, bound to a specific deployed contract.
func NewDeFiTokenStorageCaller(address common.Address, caller bind.ContractCaller) (*DeFiTokenStorageCaller, error) {
	contract, err := bindDeFiTokenStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DeFiTokenStorageCaller{contract: contract}, nil
}

// NewDeFiTokenStorageTransactor creates a new write-only instance of DeFiTokenStorage, bound to a specific deployed contract.
func NewDeFiTokenStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*DeFiTokenStorageTransactor, error) {
	contract, err := bindDeFiTokenStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DeFiTokenStorageTransactor{contract: contract}, nil
}

// NewDeFiTokenStorageFilterer creates a new log filterer instance of DeFiTokenStorage, bound to a specific deployed contract.
func NewDeFiTokenStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*DeFiTokenStorageFilterer, error) {
	contract, err := bindDeFiTokenStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DeFiTokenStorageFilterer{contract: contract}, nil
}

// bindDeFiTokenStorage binds a generic wrapper to an already deployed contract.
func bindDeFiTokenStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DeFiTokenStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DeFiTokenStorage *DeFiTokenStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DeFiTokenStorage.Contract.DeFiTokenStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DeFiTokenStorage *DeFiTokenStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeFiTokenStorage.Contract.DeFiTokenStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DeFiTokenStorage *DeFiTokenStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DeFiTokenStorage.Contract.DeFiTokenStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DeFiTokenStorage *DeFiTokenStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DeFiTokenStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DeFiTokenStorage *DeFiTokenStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeFiTokenStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DeFiTokenStorage *DeFiTokenStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DeFiTokenStorage.Contract.contract.Transact(opts, method, params...)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_DeFiTokenStorage *DeFiTokenStorageCaller) AddressProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DeFiTokenStorage.contract.Call(opts, &out, "addressProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_DeFiTokenStorage *DeFiTokenStorageSession) AddressProvider() (common.Address, error) {
	return _DeFiTokenStorage.Contract.AddressProvider(&_DeFiTokenStorage.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_DeFiTokenStorage *DeFiTokenStorageCallerSession) AddressProvider() (common.Address, error) {
	return _DeFiTokenStorage.Contract.AddressProvider(&_DeFiTokenStorage.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb203bb99.
//
// Solidity: function balance(address , address ) view returns(uint256)
func (_DeFiTokenStorage *DeFiTokenStorageCaller) Balance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DeFiTokenStorage.contract.Call(opts, &out, "balance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0xb203bb99.
//
// Solidity: function balance(address , address ) view returns(uint256)
func (_DeFiTokenStorage *DeFiTokenStorageSession) Balance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DeFiTokenStorage.Contract.Balance(&_DeFiTokenStorage.CallOpts, arg0, arg1)
}

// Balance is a free data retrieval call binding the contract method 0xb203bb99.
//
// Solidity: function balance(address , address ) view returns(uint256)
func (_DeFiTokenStorage *DeFiTokenStorageCallerSession) Balance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DeFiTokenStorage.Contract.Balance(&_DeFiTokenStorage.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(address _account, address _token) view returns(uint256)
func (_DeFiTokenStorage *DeFiTokenStorageCaller) BalanceOf(opts *bind.CallOpts, _account common.Address, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DeFiTokenStorage.contract.Call(opts, &out, "balanceOf", _account, _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(address _account, address _token) view returns(uint256)
func (_DeFiTokenStorage *DeFiTokenStorageSession) BalanceOf(_account common.Address, _token common.Address) (*big.Int, error) {
	return _DeFiTokenStorage.Contract.BalanceOf(&_DeFiTokenStorage.CallOpts, _account, _token)
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(address _account, address _token) view returns(uint256)
func (_DeFiTokenStorage *DeFiTokenStorageCallerSession) BalanceOf(_account common.Address, _token common.Address) (*big.Int, error) {
	return _DeFiTokenStorage.Contract.BalanceOf(&_DeFiTokenStorage.CallOpts, _account, _token)
}

// TokenValue is a free data retrieval call binding the contract method 0xf1821783.
//
// Solidity: function tokenValue(address _token, uint256 _amount) view returns(uint256 value)
func (_DeFiTokenStorage *DeFiTokenStorageCaller) TokenValue(opts *bind.CallOpts, _token common.Address, _amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DeFiTokenStorage.contract.Call(opts, &out, "tokenValue", _token, _amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenValue is a free data retrieval call binding the contract method 0xf1821783.
//
// Solidity: function tokenValue(address _token, uint256 _amount) view returns(uint256 value)
func (_DeFiTokenStorage *DeFiTokenStorageSession) TokenValue(_token common.Address, _amount *big.Int) (*big.Int, error) {
	return _DeFiTokenStorage.Contract.TokenValue(&_DeFiTokenStorage.CallOpts, _token, _amount)
}

// TokenValue is a free data retrieval call binding the contract method 0xf1821783.
//
// Solidity: function tokenValue(address _token, uint256 _amount) view returns(uint256 value)
func (_DeFiTokenStorage *DeFiTokenStorageCallerSession) TokenValue(_token common.Address, _amount *big.Int) (*big.Int, error) {
	return _DeFiTokenStorage.Contract.TokenValue(&_DeFiTokenStorage.CallOpts, _token, _amount)
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address)
func (_DeFiTokenStorage *DeFiTokenStorageCaller) Tokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DeFiTokenStorage.contract.Call(opts, &out, "tokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address)
func (_DeFiTokenStorage *DeFiTokenStorageSession) Tokens(arg0 *big.Int) (common.Address, error) {
	return _DeFiTokenStorage.Contract.Tokens(&_DeFiTokenStorage.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address)
func (_DeFiTokenStorage *DeFiTokenStorageCallerSession) Tokens(arg0 *big.Int) (common.Address, error) {
	return _DeFiTokenStorage.Contract.Tokens(&_DeFiTokenStorage.CallOpts, arg0)
}

// TokensCount is a free data retrieval call binding the contract method 0xa64ed8ba.
//
// Solidity: function tokensCount() view returns(uint256)
func (_DeFiTokenStorage *DeFiTokenStorageCaller) TokensCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DeFiTokenStorage.contract.Call(opts, &out, "tokensCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokensCount is a free data retrieval call binding the contract method 0xa64ed8ba.
//
// Solidity: function tokensCount() view returns(uint256)
func (_DeFiTokenStorage *DeFiTokenStorageSession) TokensCount() (*big.Int, error) {
	return _DeFiTokenStorage.Contract.TokensCount(&_DeFiTokenStorage.CallOpts)
}

// TokensCount is a free data retrieval call binding the contract method 0xa64ed8ba.
//
// Solidity: function tokensCount() view returns(uint256)
func (_DeFiTokenStorage *DeFiTokenStorageCallerSession) TokensCount() (*big.Int, error) {
	return _DeFiTokenStorage.Contract.TokensCount(&_DeFiTokenStorage.CallOpts)
}

// Total is a free data retrieval call binding the contract method 0x2ddbd13a.
//
// Solidity: function total() view returns(uint256 value)
func (_DeFiTokenStorage *DeFiTokenStorageCaller) Total(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DeFiTokenStorage.contract.Call(opts, &out, "total")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Total is a free data retrieval call binding the contract method 0x2ddbd13a.
//
// Solidity: function total() view returns(uint256 value)
func (_DeFiTokenStorage *DeFiTokenStorageSession) Total() (*big.Int, error) {
	return _DeFiTokenStorage.Contract.Total(&_DeFiTokenStorage.CallOpts)
}

// Total is a free data retrieval call binding the contract method 0x2ddbd13a.
//
// Solidity: function total() view returns(uint256 value)
func (_DeFiTokenStorage *DeFiTokenStorageCallerSession) Total() (*big.Int, error) {
	return _DeFiTokenStorage.Contract.Total(&_DeFiTokenStorage.CallOpts)
}

// TotalBalance is a free data retrieval call binding the contract method 0x6eacd398.
//
// Solidity: function totalBalance(address ) view returns(uint256)
func (_DeFiTokenStorage *DeFiTokenStorageCaller) TotalBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DeFiTokenStorage.contract.Call(opts, &out, "totalBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBalance is a free data retrieval call binding the contract method 0x6eacd398.
//
// Solidity: function totalBalance(address ) view returns(uint256)
func (_DeFiTokenStorage *DeFiTokenStorageSession) TotalBalance(arg0 common.Address) (*big.Int, error) {
	return _DeFiTokenStorage.Contract.TotalBalance(&_DeFiTokenStorage.CallOpts, arg0)
}

// TotalBalance is a free data retrieval call binding the contract method 0x6eacd398.
//
// Solidity: function totalBalance(address ) view returns(uint256)
func (_DeFiTokenStorage *DeFiTokenStorageCallerSession) TotalBalance(arg0 common.Address) (*big.Int, error) {
	return _DeFiTokenStorage.Contract.TotalBalance(&_DeFiTokenStorage.CallOpts, arg0)
}

// TotalOf is a free data retrieval call binding the contract method 0x912c2673.
//
// Solidity: function totalOf(address _account) view returns(uint256)
func (_DeFiTokenStorage *DeFiTokenStorageCaller) TotalOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DeFiTokenStorage.contract.Call(opts, &out, "totalOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalOf is a free data retrieval call binding the contract method 0x912c2673.
//
// Solidity: function totalOf(address _account) view returns(uint256)
func (_DeFiTokenStorage *DeFiTokenStorageSession) TotalOf(_account common.Address) (*big.Int, error) {
	return _DeFiTokenStorage.Contract.TotalOf(&_DeFiTokenStorage.CallOpts, _account)
}

// TotalOf is a free data retrieval call binding the contract method 0x912c2673.
//
// Solidity: function totalOf(address _account) view returns(uint256)
func (_DeFiTokenStorage *DeFiTokenStorageCallerSession) TotalOf(_account common.Address) (*big.Int, error) {
	return _DeFiTokenStorage.Contract.TotalOf(&_DeFiTokenStorage.CallOpts, _account)
}

// TotalOfDec is a free data retrieval call binding the contract method 0x65be454d.
//
// Solidity: function totalOfDec(address _account, address _token, uint256 _amount) view returns(uint256 value)
func (_DeFiTokenStorage *DeFiTokenStorageCaller) TotalOfDec(opts *bind.CallOpts, _account common.Address, _token common.Address, _amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DeFiTokenStorage.contract.Call(opts, &out, "totalOfDec", _account, _token, _amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalOfDec is a free data retrieval call binding the contract method 0x65be454d.
//
// Solidity: function totalOfDec(address _account, address _token, uint256 _amount) view returns(uint256 value)
func (_DeFiTokenStorage *DeFiTokenStorageSession) TotalOfDec(_account common.Address, _token common.Address, _amount *big.Int) (*big.Int, error) {
	return _DeFiTokenStorage.Contract.TotalOfDec(&_DeFiTokenStorage.CallOpts, _account, _token, _amount)
}

// TotalOfDec is a free data retrieval call binding the contract method 0x65be454d.
//
// Solidity: function totalOfDec(address _account, address _token, uint256 _amount) view returns(uint256 value)
func (_DeFiTokenStorage *DeFiTokenStorageCallerSession) TotalOfDec(_account common.Address, _token common.Address, _amount *big.Int) (*big.Int, error) {
	return _DeFiTokenStorage.Contract.TotalOfDec(&_DeFiTokenStorage.CallOpts, _account, _token, _amount)
}

// TotalOfInc is a free data retrieval call binding the contract method 0x660eab83.
//
// Solidity: function totalOfInc(address _account, address _token, uint256 _amount) view returns(uint256 value)
func (_DeFiTokenStorage *DeFiTokenStorageCaller) TotalOfInc(opts *bind.CallOpts, _account common.Address, _token common.Address, _amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DeFiTokenStorage.contract.Call(opts, &out, "totalOfInc", _account, _token, _amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalOfInc is a free data retrieval call binding the contract method 0x660eab83.
//
// Solidity: function totalOfInc(address _account, address _token, uint256 _amount) view returns(uint256 value)
func (_DeFiTokenStorage *DeFiTokenStorageSession) TotalOfInc(_account common.Address, _token common.Address, _amount *big.Int) (*big.Int, error) {
	return _DeFiTokenStorage.Contract.TotalOfInc(&_DeFiTokenStorage.CallOpts, _account, _token, _amount)
}

// TotalOfInc is a free data retrieval call binding the contract method 0x660eab83.
//
// Solidity: function totalOfInc(address _account, address _token, uint256 _amount) view returns(uint256 value)
func (_DeFiTokenStorage *DeFiTokenStorageCallerSession) TotalOfInc(_account common.Address, _token common.Address, _amount *big.Int) (*big.Int, error) {
	return _DeFiTokenStorage.Contract.TotalOfInc(&_DeFiTokenStorage.CallOpts, _account, _token, _amount)
}

// ValueDustAdjustment is a free data retrieval call binding the contract method 0x496a1140.
//
// Solidity: function valueDustAdjustment() view returns(bool)
func (_DeFiTokenStorage *DeFiTokenStorageCaller) ValueDustAdjustment(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DeFiTokenStorage.contract.Call(opts, &out, "valueDustAdjustment")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValueDustAdjustment is a free data retrieval call binding the contract method 0x496a1140.
//
// Solidity: function valueDustAdjustment() view returns(bool)
func (_DeFiTokenStorage *DeFiTokenStorageSession) ValueDustAdjustment() (bool, error) {
	return _DeFiTokenStorage.Contract.ValueDustAdjustment(&_DeFiTokenStorage.CallOpts)
}

// ValueDustAdjustment is a free data retrieval call binding the contract method 0x496a1140.
//
// Solidity: function valueDustAdjustment() view returns(bool)
func (_DeFiTokenStorage *DeFiTokenStorageCallerSession) ValueDustAdjustment() (bool, error) {
	return _DeFiTokenStorage.Contract.ValueDustAdjustment(&_DeFiTokenStorage.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0x551f8e2a.
//
// Solidity: function add(address _account, address _token, uint256 _amount) returns()
func (_DeFiTokenStorage *DeFiTokenStorageTransactor) Add(opts *bind.TransactOpts, _account common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DeFiTokenStorage.contract.Transact(opts, "add", _account, _token, _amount)
}

// Add is a paid mutator transaction binding the contract method 0x551f8e2a.
//
// Solidity: function add(address _account, address _token, uint256 _amount) returns()
func (_DeFiTokenStorage *DeFiTokenStorageSession) Add(_account common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DeFiTokenStorage.Contract.Add(&_DeFiTokenStorage.TransactOpts, _account, _token, _amount)
}

// Add is a paid mutator transaction binding the contract method 0x551f8e2a.
//
// Solidity: function add(address _account, address _token, uint256 _amount) returns()
func (_DeFiTokenStorage *DeFiTokenStorageTransactorSession) Add(_account common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DeFiTokenStorage.Contract.Add(&_DeFiTokenStorage.TransactOpts, _account, _token, _amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x400ada75.
//
// Solidity: function initialize(address _addressProvider, bool _dustAdt) returns()
func (_DeFiTokenStorage *DeFiTokenStorageTransactor) Initialize(opts *bind.TransactOpts, _addressProvider common.Address, _dustAdt bool) (*types.Transaction, error) {
	return _DeFiTokenStorage.contract.Transact(opts, "initialize", _addressProvider, _dustAdt)
}

// Initialize is a paid mutator transaction binding the contract method 0x400ada75.
//
// Solidity: function initialize(address _addressProvider, bool _dustAdt) returns()
func (_DeFiTokenStorage *DeFiTokenStorageSession) Initialize(_addressProvider common.Address, _dustAdt bool) (*types.Transaction, error) {
	return _DeFiTokenStorage.Contract.Initialize(&_DeFiTokenStorage.TransactOpts, _addressProvider, _dustAdt)
}

// Initialize is a paid mutator transaction binding the contract method 0x400ada75.
//
// Solidity: function initialize(address _addressProvider, bool _dustAdt) returns()
func (_DeFiTokenStorage *DeFiTokenStorageTransactorSession) Initialize(_addressProvider common.Address, _dustAdt bool) (*types.Transaction, error) {
	return _DeFiTokenStorage.Contract.Initialize(&_DeFiTokenStorage.TransactOpts, _addressProvider, _dustAdt)
}

// Sub is a paid mutator transaction binding the contract method 0x55ceeb31.
//
// Solidity: function sub(address _account, address _token, uint256 _amount) returns()
func (_DeFiTokenStorage *DeFiTokenStorageTransactor) Sub(opts *bind.TransactOpts, _account common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DeFiTokenStorage.contract.Transact(opts, "sub", _account, _token, _amount)
}

// Sub is a paid mutator transaction binding the contract method 0x55ceeb31.
//
// Solidity: function sub(address _account, address _token, uint256 _amount) returns()
func (_DeFiTokenStorage *DeFiTokenStorageSession) Sub(_account common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DeFiTokenStorage.Contract.Sub(&_DeFiTokenStorage.TransactOpts, _account, _token, _amount)
}

// Sub is a paid mutator transaction binding the contract method 0x55ceeb31.
//
// Solidity: function sub(address _account, address _token, uint256 _amount) returns()
func (_DeFiTokenStorage *DeFiTokenStorageTransactorSession) Sub(_account common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DeFiTokenStorage.Contract.Sub(&_DeFiTokenStorage.TransactOpts, _account, _token, _amount)
}
