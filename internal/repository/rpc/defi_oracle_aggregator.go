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

// DefiOracleReferenceAggregatorABI is the input ABI used to generate the binding from.
const DefiOracleReferenceAggregatorABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"AggregatorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TokenInformationAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TokenInformationChanged\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_aggregator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_logo\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_priceDecimals\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"_isActive\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_canDeposit\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_canBorrow\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_canTrade\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_volatility\",\"type\":\"uint256\"}],\"name\":\"addToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"aggregators\",\"outputs\":[{\"internalType\":\"contractAggregatorInterface\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"findTokenIndex\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_back\",\"type\":\"uint256\"}],\"name\":\"getPreviousPrice\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_back\",\"type\":\"uint256\"}],\"name\":\"getPreviousTimeStamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTimeStamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nativeToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"setAggregator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"logo\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"priceDecimals\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"canDeposit\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"canBorrow\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"canTrade\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"volatility\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokensCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_logo\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_priceDecimals\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"_isActive\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_canDeposit\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_canBorrow\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_canTrade\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_volatility\",\"type\":\"uint256\"}],\"name\":\"updateToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DefiOracleReferenceAggregator is an auto generated Go binding around an Ethereum contract.
type DefiOracleReferenceAggregator struct {
	DefiOracleReferenceAggregatorCaller     // Read-only binding to the contract
	DefiOracleReferenceAggregatorTransactor // Write-only binding to the contract
	DefiOracleReferenceAggregatorFilterer   // Log filterer for contract events
}

// DefiOracleReferenceAggregatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type DefiOracleReferenceAggregatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefiOracleReferenceAggregatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DefiOracleReferenceAggregatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefiOracleReferenceAggregatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DefiOracleReferenceAggregatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefiOracleReferenceAggregatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DefiOracleReferenceAggregatorSession struct {
	Contract     *DefiOracleReferenceAggregator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                  // Call options to use throughout this session
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// DefiOracleReferenceAggregatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DefiOracleReferenceAggregatorCallerSession struct {
	Contract *DefiOracleReferenceAggregatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                        // Call options to use throughout this session
}

// DefiOracleReferenceAggregatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DefiOracleReferenceAggregatorTransactorSession struct {
	Contract     *DefiOracleReferenceAggregatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                        // Transaction auth options to use throughout this session
}

// DefiOracleReferenceAggregatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type DefiOracleReferenceAggregatorRaw struct {
	Contract *DefiOracleReferenceAggregator // Generic contract binding to access the raw methods on
}

// DefiOracleReferenceAggregatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DefiOracleReferenceAggregatorCallerRaw struct {
	Contract *DefiOracleReferenceAggregatorCaller // Generic read-only contract binding to access the raw methods on
}

// DefiOracleReferenceAggregatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DefiOracleReferenceAggregatorTransactorRaw struct {
	Contract *DefiOracleReferenceAggregatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDefiOracleReferenceAggregator creates a new instance of DefiOracleReferenceAggregator, bound to a specific deployed contract.
func NewDefiOracleReferenceAggregator(address common.Address, backend bind.ContractBackend) (*DefiOracleReferenceAggregator, error) {
	contract, err := bindDefiOracleReferenceAggregator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DefiOracleReferenceAggregator{DefiOracleReferenceAggregatorCaller: DefiOracleReferenceAggregatorCaller{contract: contract}, DefiOracleReferenceAggregatorTransactor: DefiOracleReferenceAggregatorTransactor{contract: contract}, DefiOracleReferenceAggregatorFilterer: DefiOracleReferenceAggregatorFilterer{contract: contract}}, nil
}

// NewDefiOracleReferenceAggregatorCaller creates a new read-only instance of DefiOracleReferenceAggregator, bound to a specific deployed contract.
func NewDefiOracleReferenceAggregatorCaller(address common.Address, caller bind.ContractCaller) (*DefiOracleReferenceAggregatorCaller, error) {
	contract, err := bindDefiOracleReferenceAggregator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DefiOracleReferenceAggregatorCaller{contract: contract}, nil
}

// NewDefiOracleReferenceAggregatorTransactor creates a new write-only instance of DefiOracleReferenceAggregator, bound to a specific deployed contract.
func NewDefiOracleReferenceAggregatorTransactor(address common.Address, transactor bind.ContractTransactor) (*DefiOracleReferenceAggregatorTransactor, error) {
	contract, err := bindDefiOracleReferenceAggregator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DefiOracleReferenceAggregatorTransactor{contract: contract}, nil
}

// NewDefiOracleReferenceAggregatorFilterer creates a new log filterer instance of DefiOracleReferenceAggregator, bound to a specific deployed contract.
func NewDefiOracleReferenceAggregatorFilterer(address common.Address, filterer bind.ContractFilterer) (*DefiOracleReferenceAggregatorFilterer, error) {
	contract, err := bindDefiOracleReferenceAggregator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DefiOracleReferenceAggregatorFilterer{contract: contract}, nil
}

// bindDefiOracleReferenceAggregator binds a generic wrapper to an already deployed contract.
func bindDefiOracleReferenceAggregator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DefiOracleReferenceAggregatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DefiOracleReferenceAggregator.Contract.DefiOracleReferenceAggregatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefiOracleReferenceAggregator.Contract.DefiOracleReferenceAggregatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DefiOracleReferenceAggregator.Contract.DefiOracleReferenceAggregatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DefiOracleReferenceAggregator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefiOracleReferenceAggregator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DefiOracleReferenceAggregator.Contract.contract.Transact(opts, method, params...)
}

// Aggregators is a free data retrieval call binding the contract method 0x112cdab9.
//
// Solidity: function aggregators(address ) view returns(address)
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorCaller) Aggregators(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DefiOracleReferenceAggregator.contract.Call(opts, out, "aggregators", arg0)
	return *ret0, err
}

// Aggregators is a free data retrieval call binding the contract method 0x112cdab9.
//
// Solidity: function aggregators(address ) view returns(address)
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorSession) Aggregators(arg0 common.Address) (common.Address, error) {
	return _DefiOracleReferenceAggregator.Contract.Aggregators(&_DefiOracleReferenceAggregator.CallOpts, arg0)
}

// Aggregators is a free data retrieval call binding the contract method 0x112cdab9.
//
// Solidity: function aggregators(address ) view returns(address)
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorCallerSession) Aggregators(arg0 common.Address) (common.Address, error) {
	return _DefiOracleReferenceAggregator.Contract.Aggregators(&_DefiOracleReferenceAggregator.CallOpts, arg0)
}

// FindTokenIndex is a free data retrieval call binding the contract method 0x0a11fd07.
//
// Solidity: function findTokenIndex(address _token) view returns(int256)
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorCaller) FindTokenIndex(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiOracleReferenceAggregator.contract.Call(opts, out, "findTokenIndex", _token)
	return *ret0, err
}

// FindTokenIndex is a free data retrieval call binding the contract method 0x0a11fd07.
//
// Solidity: function findTokenIndex(address _token) view returns(int256)
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorSession) FindTokenIndex(_token common.Address) (*big.Int, error) {
	return _DefiOracleReferenceAggregator.Contract.FindTokenIndex(&_DefiOracleReferenceAggregator.CallOpts, _token)
}

// FindTokenIndex is a free data retrieval call binding the contract method 0x0a11fd07.
//
// Solidity: function findTokenIndex(address _token) view returns(int256)
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorCallerSession) FindTokenIndex(_token common.Address) (*big.Int, error) {
	return _DefiOracleReferenceAggregator.Contract.FindTokenIndex(&_DefiOracleReferenceAggregator.CallOpts, _token)
}

// GetPreviousPrice is a free data retrieval call binding the contract method 0xe889e5d6.
//
// Solidity: function getPreviousPrice(address token, uint256 _back) view returns(int256)
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorCaller) GetPreviousPrice(opts *bind.CallOpts, token common.Address, _back *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiOracleReferenceAggregator.contract.Call(opts, out, "getPreviousPrice", token, _back)
	return *ret0, err
}

// GetPreviousPrice is a free data retrieval call binding the contract method 0xe889e5d6.
//
// Solidity: function getPreviousPrice(address token, uint256 _back) view returns(int256)
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorSession) GetPreviousPrice(token common.Address, _back *big.Int) (*big.Int, error) {
	return _DefiOracleReferenceAggregator.Contract.GetPreviousPrice(&_DefiOracleReferenceAggregator.CallOpts, token, _back)
}

// GetPreviousPrice is a free data retrieval call binding the contract method 0xe889e5d6.
//
// Solidity: function getPreviousPrice(address token, uint256 _back) view returns(int256)
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorCallerSession) GetPreviousPrice(token common.Address, _back *big.Int) (*big.Int, error) {
	return _DefiOracleReferenceAggregator.Contract.GetPreviousPrice(&_DefiOracleReferenceAggregator.CallOpts, token, _back)
}

// GetPreviousTimeStamp is a free data retrieval call binding the contract method 0xe0a3581e.
//
// Solidity: function getPreviousTimeStamp(address token, uint256 _back) view returns(uint256)
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorCaller) GetPreviousTimeStamp(opts *bind.CallOpts, token common.Address, _back *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiOracleReferenceAggregator.contract.Call(opts, out, "getPreviousTimeStamp", token, _back)
	return *ret0, err
}

// GetPreviousTimeStamp is a free data retrieval call binding the contract method 0xe0a3581e.
//
// Solidity: function getPreviousTimeStamp(address token, uint256 _back) view returns(uint256)
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorSession) GetPreviousTimeStamp(token common.Address, _back *big.Int) (*big.Int, error) {
	return _DefiOracleReferenceAggregator.Contract.GetPreviousTimeStamp(&_DefiOracleReferenceAggregator.CallOpts, token, _back)
}

// GetPreviousTimeStamp is a free data retrieval call binding the contract method 0xe0a3581e.
//
// Solidity: function getPreviousTimeStamp(address token, uint256 _back) view returns(uint256)
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorCallerSession) GetPreviousTimeStamp(token common.Address, _back *big.Int) (*big.Int, error) {
	return _DefiOracleReferenceAggregator.Contract.GetPreviousTimeStamp(&_DefiOracleReferenceAggregator.CallOpts, token, _back)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address token) view returns(int256)
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorCaller) GetPrice(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiOracleReferenceAggregator.contract.Call(opts, out, "getPrice", token)
	return *ret0, err
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address token) view returns(int256)
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorSession) GetPrice(token common.Address) (*big.Int, error) {
	return _DefiOracleReferenceAggregator.Contract.GetPrice(&_DefiOracleReferenceAggregator.CallOpts, token)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address token) view returns(int256)
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorCallerSession) GetPrice(token common.Address) (*big.Int, error) {
	return _DefiOracleReferenceAggregator.Contract.GetPrice(&_DefiOracleReferenceAggregator.CallOpts, token)
}

// GetTimeStamp is a free data retrieval call binding the contract method 0xeb470ebf.
//
// Solidity: function getTimeStamp(address token) view returns(uint256)
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorCaller) GetTimeStamp(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiOracleReferenceAggregator.contract.Call(opts, out, "getTimeStamp", token)
	return *ret0, err
}

// GetTimeStamp is a free data retrieval call binding the contract method 0xeb470ebf.
//
// Solidity: function getTimeStamp(address token) view returns(uint256)
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorSession) GetTimeStamp(token common.Address) (*big.Int, error) {
	return _DefiOracleReferenceAggregator.Contract.GetTimeStamp(&_DefiOracleReferenceAggregator.CallOpts, token)
}

// GetTimeStamp is a free data retrieval call binding the contract method 0xeb470ebf.
//
// Solidity: function getTimeStamp(address token) view returns(uint256)
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorCallerSession) GetTimeStamp(token common.Address) (*big.Int, error) {
	return _DefiOracleReferenceAggregator.Contract.GetTimeStamp(&_DefiOracleReferenceAggregator.CallOpts, token)
}

// NativeToken is a free data retrieval call binding the contract method 0xe1758bd8.
//
// Solidity: function nativeToken() view returns(address)
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorCaller) NativeToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DefiOracleReferenceAggregator.contract.Call(opts, out, "nativeToken")
	return *ret0, err
}

// NativeToken is a free data retrieval call binding the contract method 0xe1758bd8.
//
// Solidity: function nativeToken() view returns(address)
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorSession) NativeToken() (common.Address, error) {
	return _DefiOracleReferenceAggregator.Contract.NativeToken(&_DefiOracleReferenceAggregator.CallOpts)
}

// NativeToken is a free data retrieval call binding the contract method 0xe1758bd8.
//
// Solidity: function nativeToken() view returns(address)
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorCallerSession) NativeToken() (common.Address, error) {
	return _DefiOracleReferenceAggregator.Contract.NativeToken(&_DefiOracleReferenceAggregator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DefiOracleReferenceAggregator.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorSession) Owner() (common.Address, error) {
	return _DefiOracleReferenceAggregator.Contract.Owner(&_DefiOracleReferenceAggregator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorCallerSession) Owner() (common.Address, error) {
	return _DefiOracleReferenceAggregator.Contract.Owner(&_DefiOracleReferenceAggregator.CallOpts)
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address token, string name, string symbol, string logo, uint8 decimals, uint8 priceDecimals, bool isActive, bool canDeposit, bool canBorrow, bool canTrade, uint256 volatility)
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorCaller) Tokens(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Token         common.Address
	Name          string
	Symbol        string
	Logo          string
	Decimals      uint8
	PriceDecimals uint8
	IsActive      bool
	CanDeposit    bool
	CanBorrow     bool
	CanTrade      bool
	Volatility    *big.Int
}, error) {
	ret := new(struct {
		Token         common.Address
		Name          string
		Symbol        string
		Logo          string
		Decimals      uint8
		PriceDecimals uint8
		IsActive      bool
		CanDeposit    bool
		CanBorrow     bool
		CanTrade      bool
		Volatility    *big.Int
	})
	out := ret
	err := _DefiOracleReferenceAggregator.contract.Call(opts, out, "tokens", arg0)
	return *ret, err
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address token, string name, string symbol, string logo, uint8 decimals, uint8 priceDecimals, bool isActive, bool canDeposit, bool canBorrow, bool canTrade, uint256 volatility)
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorSession) Tokens(arg0 *big.Int) (struct {
	Token         common.Address
	Name          string
	Symbol        string
	Logo          string
	Decimals      uint8
	PriceDecimals uint8
	IsActive      bool
	CanDeposit    bool
	CanBorrow     bool
	CanTrade      bool
	Volatility    *big.Int
}, error) {
	return _DefiOracleReferenceAggregator.Contract.Tokens(&_DefiOracleReferenceAggregator.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address token, string name, string symbol, string logo, uint8 decimals, uint8 priceDecimals, bool isActive, bool canDeposit, bool canBorrow, bool canTrade, uint256 volatility)
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorCallerSession) Tokens(arg0 *big.Int) (struct {
	Token         common.Address
	Name          string
	Symbol        string
	Logo          string
	Decimals      uint8
	PriceDecimals uint8
	IsActive      bool
	CanDeposit    bool
	CanBorrow     bool
	CanTrade      bool
	Volatility    *big.Int
}, error) {
	return _DefiOracleReferenceAggregator.Contract.Tokens(&_DefiOracleReferenceAggregator.CallOpts, arg0)
}

// TokensCount is a free data retrieval call binding the contract method 0xa64ed8ba.
//
// Solidity: function tokensCount() view returns(uint256)
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorCaller) TokensCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiOracleReferenceAggregator.contract.Call(opts, out, "tokensCount")
	return *ret0, err
}

// TokensCount is a free data retrieval call binding the contract method 0xa64ed8ba.
//
// Solidity: function tokensCount() view returns(uint256)
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorSession) TokensCount() (*big.Int, error) {
	return _DefiOracleReferenceAggregator.Contract.TokensCount(&_DefiOracleReferenceAggregator.CallOpts)
}

// TokensCount is a free data retrieval call binding the contract method 0xa64ed8ba.
//
// Solidity: function tokensCount() view returns(uint256)
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorCallerSession) TokensCount() (*big.Int, error) {
	return _DefiOracleReferenceAggregator.Contract.TokensCount(&_DefiOracleReferenceAggregator.CallOpts)
}

// AddToken is a paid mutator transaction binding the contract method 0x6e36fabc.
//
// Solidity: function addToken(address _token, address _aggregator, string _logo, uint8 _priceDecimals, bool _isActive, bool _canDeposit, bool _canBorrow, bool _canTrade, uint256 _volatility) returns()
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorTransactor) AddToken(opts *bind.TransactOpts, _token common.Address, _aggregator common.Address, _logo string, _priceDecimals uint8, _isActive bool, _canDeposit bool, _canBorrow bool, _canTrade bool, _volatility *big.Int) (*types.Transaction, error) {
	return _DefiOracleReferenceAggregator.contract.Transact(opts, "addToken", _token, _aggregator, _logo, _priceDecimals, _isActive, _canDeposit, _canBorrow, _canTrade, _volatility)
}

// AddToken is a paid mutator transaction binding the contract method 0x6e36fabc.
//
// Solidity: function addToken(address _token, address _aggregator, string _logo, uint8 _priceDecimals, bool _isActive, bool _canDeposit, bool _canBorrow, bool _canTrade, uint256 _volatility) returns()
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorSession) AddToken(_token common.Address, _aggregator common.Address, _logo string, _priceDecimals uint8, _isActive bool, _canDeposit bool, _canBorrow bool, _canTrade bool, _volatility *big.Int) (*types.Transaction, error) {
	return _DefiOracleReferenceAggregator.Contract.AddToken(&_DefiOracleReferenceAggregator.TransactOpts, _token, _aggregator, _logo, _priceDecimals, _isActive, _canDeposit, _canBorrow, _canTrade, _volatility)
}

// AddToken is a paid mutator transaction binding the contract method 0x6e36fabc.
//
// Solidity: function addToken(address _token, address _aggregator, string _logo, uint8 _priceDecimals, bool _isActive, bool _canDeposit, bool _canBorrow, bool _canTrade, uint256 _volatility) returns()
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorTransactorSession) AddToken(_token common.Address, _aggregator common.Address, _logo string, _priceDecimals uint8, _isActive bool, _canDeposit bool, _canBorrow bool, _canTrade bool, _volatility *big.Int) (*types.Transaction, error) {
	return _DefiOracleReferenceAggregator.Contract.AddToken(&_DefiOracleReferenceAggregator.TransactOpts, _token, _aggregator, _logo, _priceDecimals, _isActive, _canDeposit, _canBorrow, _canTrade, _volatility)
}

// SetAggregator is a paid mutator transaction binding the contract method 0x7394a7f6.
//
// Solidity: function setAggregator(address token, address aggregator) returns()
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorTransactor) SetAggregator(opts *bind.TransactOpts, token common.Address, aggregator common.Address) (*types.Transaction, error) {
	return _DefiOracleReferenceAggregator.contract.Transact(opts, "setAggregator", token, aggregator)
}

// SetAggregator is a paid mutator transaction binding the contract method 0x7394a7f6.
//
// Solidity: function setAggregator(address token, address aggregator) returns()
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorSession) SetAggregator(token common.Address, aggregator common.Address) (*types.Transaction, error) {
	return _DefiOracleReferenceAggregator.Contract.SetAggregator(&_DefiOracleReferenceAggregator.TransactOpts, token, aggregator)
}

// SetAggregator is a paid mutator transaction binding the contract method 0x7394a7f6.
//
// Solidity: function setAggregator(address token, address aggregator) returns()
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorTransactorSession) SetAggregator(token common.Address, aggregator common.Address) (*types.Transaction, error) {
	return _DefiOracleReferenceAggregator.Contract.SetAggregator(&_DefiOracleReferenceAggregator.TransactOpts, token, aggregator)
}

// UpdateToken is a paid mutator transaction binding the contract method 0x3ca29277.
//
// Solidity: function updateToken(address _token, string _logo, uint8 _priceDecimals, bool _isActive, bool _canDeposit, bool _canBorrow, bool _canTrade, uint256 _volatility) returns()
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorTransactor) UpdateToken(opts *bind.TransactOpts, _token common.Address, _logo string, _priceDecimals uint8, _isActive bool, _canDeposit bool, _canBorrow bool, _canTrade bool, _volatility *big.Int) (*types.Transaction, error) {
	return _DefiOracleReferenceAggregator.contract.Transact(opts, "updateToken", _token, _logo, _priceDecimals, _isActive, _canDeposit, _canBorrow, _canTrade, _volatility)
}

// UpdateToken is a paid mutator transaction binding the contract method 0x3ca29277.
//
// Solidity: function updateToken(address _token, string _logo, uint8 _priceDecimals, bool _isActive, bool _canDeposit, bool _canBorrow, bool _canTrade, uint256 _volatility) returns()
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorSession) UpdateToken(_token common.Address, _logo string, _priceDecimals uint8, _isActive bool, _canDeposit bool, _canBorrow bool, _canTrade bool, _volatility *big.Int) (*types.Transaction, error) {
	return _DefiOracleReferenceAggregator.Contract.UpdateToken(&_DefiOracleReferenceAggregator.TransactOpts, _token, _logo, _priceDecimals, _isActive, _canDeposit, _canBorrow, _canTrade, _volatility)
}

// UpdateToken is a paid mutator transaction binding the contract method 0x3ca29277.
//
// Solidity: function updateToken(address _token, string _logo, uint8 _priceDecimals, bool _isActive, bool _canDeposit, bool _canBorrow, bool _canTrade, uint256 _volatility) returns()
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorTransactorSession) UpdateToken(_token common.Address, _logo string, _priceDecimals uint8, _isActive bool, _canDeposit bool, _canBorrow bool, _canTrade bool, _volatility *big.Int) (*types.Transaction, error) {
	return _DefiOracleReferenceAggregator.Contract.UpdateToken(&_DefiOracleReferenceAggregator.TransactOpts, _token, _logo, _priceDecimals, _isActive, _canDeposit, _canBorrow, _canTrade, _volatility)
}

// DefiOracleReferenceAggregatorAggregatorChangedIterator is returned from FilterAggregatorChanged and is used to iterate over the raw logs and unpacked data for AggregatorChanged events raised by the DefiOracleReferenceAggregator contract.
type DefiOracleReferenceAggregatorAggregatorChangedIterator struct {
	Event *DefiOracleReferenceAggregatorAggregatorChanged // Event containing the contract specifics and raw log

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
func (it *DefiOracleReferenceAggregatorAggregatorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefiOracleReferenceAggregatorAggregatorChanged)
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
		it.Event = new(DefiOracleReferenceAggregatorAggregatorChanged)
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
func (it *DefiOracleReferenceAggregatorAggregatorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefiOracleReferenceAggregatorAggregatorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefiOracleReferenceAggregatorAggregatorChanged represents a AggregatorChanged event raised by the DefiOracleReferenceAggregator contract.
type DefiOracleReferenceAggregatorAggregatorChanged struct {
	Token      common.Address
	Aggregator common.Address
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAggregatorChanged is a free log retrieval operation binding the contract event 0x51012bc010926f3f948f99b90ed487a209375a900797c69e70dbc17f09264357.
//
// Solidity: event AggregatorChanged(address indexed token, address aggregator, uint256 timestamp)
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorFilterer) FilterAggregatorChanged(opts *bind.FilterOpts, token []common.Address) (*DefiOracleReferenceAggregatorAggregatorChangedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _DefiOracleReferenceAggregator.contract.FilterLogs(opts, "AggregatorChanged", tokenRule)
	if err != nil {
		return nil, err
	}
	return &DefiOracleReferenceAggregatorAggregatorChangedIterator{contract: _DefiOracleReferenceAggregator.contract, event: "AggregatorChanged", logs: logs, sub: sub}, nil
}

// WatchAggregatorChanged is a free log subscription operation binding the contract event 0x51012bc010926f3f948f99b90ed487a209375a900797c69e70dbc17f09264357.
//
// Solidity: event AggregatorChanged(address indexed token, address aggregator, uint256 timestamp)
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorFilterer) WatchAggregatorChanged(opts *bind.WatchOpts, sink chan<- *DefiOracleReferenceAggregatorAggregatorChanged, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _DefiOracleReferenceAggregator.contract.WatchLogs(opts, "AggregatorChanged", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefiOracleReferenceAggregatorAggregatorChanged)
				if err := _DefiOracleReferenceAggregator.contract.UnpackLog(event, "AggregatorChanged", log); err != nil {
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

// ParseAggregatorChanged is a log parse operation binding the contract event 0x51012bc010926f3f948f99b90ed487a209375a900797c69e70dbc17f09264357.
//
// Solidity: event AggregatorChanged(address indexed token, address aggregator, uint256 timestamp)
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorFilterer) ParseAggregatorChanged(log types.Log) (*DefiOracleReferenceAggregatorAggregatorChanged, error) {
	event := new(DefiOracleReferenceAggregatorAggregatorChanged)
	if err := _DefiOracleReferenceAggregator.contract.UnpackLog(event, "AggregatorChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DefiOracleReferenceAggregatorTokenInformationAddedIterator is returned from FilterTokenInformationAdded and is used to iterate over the raw logs and unpacked data for TokenInformationAdded events raised by the DefiOracleReferenceAggregator contract.
type DefiOracleReferenceAggregatorTokenInformationAddedIterator struct {
	Event *DefiOracleReferenceAggregatorTokenInformationAdded // Event containing the contract specifics and raw log

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
func (it *DefiOracleReferenceAggregatorTokenInformationAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefiOracleReferenceAggregatorTokenInformationAdded)
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
		it.Event = new(DefiOracleReferenceAggregatorTokenInformationAdded)
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
func (it *DefiOracleReferenceAggregatorTokenInformationAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefiOracleReferenceAggregatorTokenInformationAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefiOracleReferenceAggregatorTokenInformationAdded represents a TokenInformationAdded event raised by the DefiOracleReferenceAggregator contract.
type DefiOracleReferenceAggregatorTokenInformationAdded struct {
	Token     common.Address
	Name      string
	Index     *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokenInformationAdded is a free log retrieval operation binding the contract event 0x9dfeb273d84d99e25fc9417bf306e35a15ad839fe52835c0f31c7c1a4b64bb80.
//
// Solidity: event TokenInformationAdded(address indexed token, string name, uint256 index, uint256 timestamp)
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorFilterer) FilterTokenInformationAdded(opts *bind.FilterOpts, token []common.Address) (*DefiOracleReferenceAggregatorTokenInformationAddedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _DefiOracleReferenceAggregator.contract.FilterLogs(opts, "TokenInformationAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return &DefiOracleReferenceAggregatorTokenInformationAddedIterator{contract: _DefiOracleReferenceAggregator.contract, event: "TokenInformationAdded", logs: logs, sub: sub}, nil
}

// WatchTokenInformationAdded is a free log subscription operation binding the contract event 0x9dfeb273d84d99e25fc9417bf306e35a15ad839fe52835c0f31c7c1a4b64bb80.
//
// Solidity: event TokenInformationAdded(address indexed token, string name, uint256 index, uint256 timestamp)
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorFilterer) WatchTokenInformationAdded(opts *bind.WatchOpts, sink chan<- *DefiOracleReferenceAggregatorTokenInformationAdded, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _DefiOracleReferenceAggregator.contract.WatchLogs(opts, "TokenInformationAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefiOracleReferenceAggregatorTokenInformationAdded)
				if err := _DefiOracleReferenceAggregator.contract.UnpackLog(event, "TokenInformationAdded", log); err != nil {
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

// ParseTokenInformationAdded is a log parse operation binding the contract event 0x9dfeb273d84d99e25fc9417bf306e35a15ad839fe52835c0f31c7c1a4b64bb80.
//
// Solidity: event TokenInformationAdded(address indexed token, string name, uint256 index, uint256 timestamp)
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorFilterer) ParseTokenInformationAdded(log types.Log) (*DefiOracleReferenceAggregatorTokenInformationAdded, error) {
	event := new(DefiOracleReferenceAggregatorTokenInformationAdded)
	if err := _DefiOracleReferenceAggregator.contract.UnpackLog(event, "TokenInformationAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DefiOracleReferenceAggregatorTokenInformationChangedIterator is returned from FilterTokenInformationChanged and is used to iterate over the raw logs and unpacked data for TokenInformationChanged events raised by the DefiOracleReferenceAggregator contract.
type DefiOracleReferenceAggregatorTokenInformationChangedIterator struct {
	Event *DefiOracleReferenceAggregatorTokenInformationChanged // Event containing the contract specifics and raw log

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
func (it *DefiOracleReferenceAggregatorTokenInformationChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefiOracleReferenceAggregatorTokenInformationChanged)
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
		it.Event = new(DefiOracleReferenceAggregatorTokenInformationChanged)
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
func (it *DefiOracleReferenceAggregatorTokenInformationChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefiOracleReferenceAggregatorTokenInformationChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefiOracleReferenceAggregatorTokenInformationChanged represents a TokenInformationChanged event raised by the DefiOracleReferenceAggregator contract.
type DefiOracleReferenceAggregatorTokenInformationChanged struct {
	Token     common.Address
	Index     *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokenInformationChanged is a free log retrieval operation binding the contract event 0x736225fc71c38a467b7b94748f75ea4268f87bcb9678d6eda334a82d88bbec80.
//
// Solidity: event TokenInformationChanged(address indexed token, uint256 index, uint256 timestamp)
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorFilterer) FilterTokenInformationChanged(opts *bind.FilterOpts, token []common.Address) (*DefiOracleReferenceAggregatorTokenInformationChangedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _DefiOracleReferenceAggregator.contract.FilterLogs(opts, "TokenInformationChanged", tokenRule)
	if err != nil {
		return nil, err
	}
	return &DefiOracleReferenceAggregatorTokenInformationChangedIterator{contract: _DefiOracleReferenceAggregator.contract, event: "TokenInformationChanged", logs: logs, sub: sub}, nil
}

// WatchTokenInformationChanged is a free log subscription operation binding the contract event 0x736225fc71c38a467b7b94748f75ea4268f87bcb9678d6eda334a82d88bbec80.
//
// Solidity: event TokenInformationChanged(address indexed token, uint256 index, uint256 timestamp)
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorFilterer) WatchTokenInformationChanged(opts *bind.WatchOpts, sink chan<- *DefiOracleReferenceAggregatorTokenInformationChanged, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _DefiOracleReferenceAggregator.contract.WatchLogs(opts, "TokenInformationChanged", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefiOracleReferenceAggregatorTokenInformationChanged)
				if err := _DefiOracleReferenceAggregator.contract.UnpackLog(event, "TokenInformationChanged", log); err != nil {
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

// ParseTokenInformationChanged is a log parse operation binding the contract event 0x736225fc71c38a467b7b94748f75ea4268f87bcb9678d6eda334a82d88bbec80.
//
// Solidity: event TokenInformationChanged(address indexed token, uint256 index, uint256 timestamp)
func (_DefiOracleReferenceAggregator *DefiOracleReferenceAggregatorFilterer) ParseTokenInformationChanged(log types.Log) (*DefiOracleReferenceAggregatorTokenInformationChanged, error) {
	event := new(DefiOracleReferenceAggregatorTokenInformationChanged)
	if err := _DefiOracleReferenceAggregator.contract.UnpackLog(event, "TokenInformationChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}
