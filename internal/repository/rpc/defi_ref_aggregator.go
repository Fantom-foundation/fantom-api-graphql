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

// DefiReferenceAggregatorABI is the input ABI used to generate the binding from.
const DefiReferenceAggregatorABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"AggregatorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TokenInformationAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TokenInformationChanged\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_name\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_symbol\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_decimals\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isActive\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_canDeposit\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_canBorrow\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_canTrade\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_volatility\",\"type\":\"uint256\"}],\"name\":\"addToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"aggregators\",\"outputs\":[{\"internalType\":\"contractAggregatorInterface\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"findTokenIndex\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_back\",\"type\":\"uint256\"}],\"name\":\"getPreviousPrice\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_back\",\"type\":\"uint256\"}],\"name\":\"getPreviousTimeStamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTimeStamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"setAggregator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"symbol\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"canDeposit\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"canBorrow\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"canTrade\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"volatility\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokensCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_decimals\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isActive\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_canDeposit\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_canBorrow\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_canTrade\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_volatility\",\"type\":\"uint256\"}],\"name\":\"updateToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DefiReferenceAggregator is an auto generated Go binding around an Ethereum contract.
type DefiReferenceAggregator struct {
	DefiReferenceAggregatorCaller     // Read-only binding to the contract
	DefiReferenceAggregatorTransactor // Write-only binding to the contract
	DefiReferenceAggregatorFilterer   // Log filterer for contract events
}

// DefiReferenceAggregatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type DefiReferenceAggregatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefiReferenceAggregatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DefiReferenceAggregatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefiReferenceAggregatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DefiReferenceAggregatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefiReferenceAggregatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DefiReferenceAggregatorSession struct {
	Contract     *DefiReferenceAggregator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// DefiReferenceAggregatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DefiReferenceAggregatorCallerSession struct {
	Contract *DefiReferenceAggregatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// DefiReferenceAggregatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DefiReferenceAggregatorTransactorSession struct {
	Contract     *DefiReferenceAggregatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// DefiReferenceAggregatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type DefiReferenceAggregatorRaw struct {
	Contract *DefiReferenceAggregator // Generic contract binding to access the raw methods on
}

// DefiReferenceAggregatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DefiReferenceAggregatorCallerRaw struct {
	Contract *DefiReferenceAggregatorCaller // Generic read-only contract binding to access the raw methods on
}

// DefiReferenceAggregatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DefiReferenceAggregatorTransactorRaw struct {
	Contract *DefiReferenceAggregatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDefiReferenceAggregator creates a new instance of DefiReferenceAggregator, bound to a specific deployed contract.
func NewDefiReferenceAggregator(address common.Address, backend bind.ContractBackend) (*DefiReferenceAggregator, error) {
	contract, err := bindDefiReferenceAggregator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DefiReferenceAggregator{DefiReferenceAggregatorCaller: DefiReferenceAggregatorCaller{contract: contract}, DefiReferenceAggregatorTransactor: DefiReferenceAggregatorTransactor{contract: contract}, DefiReferenceAggregatorFilterer: DefiReferenceAggregatorFilterer{contract: contract}}, nil
}

// NewDefiReferenceAggregatorCaller creates a new read-only instance of DefiReferenceAggregator, bound to a specific deployed contract.
func NewDefiReferenceAggregatorCaller(address common.Address, caller bind.ContractCaller) (*DefiReferenceAggregatorCaller, error) {
	contract, err := bindDefiReferenceAggregator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DefiReferenceAggregatorCaller{contract: contract}, nil
}

// NewDefiReferenceAggregatorTransactor creates a new write-only instance of DefiReferenceAggregator, bound to a specific deployed contract.
func NewDefiReferenceAggregatorTransactor(address common.Address, transactor bind.ContractTransactor) (*DefiReferenceAggregatorTransactor, error) {
	contract, err := bindDefiReferenceAggregator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DefiReferenceAggregatorTransactor{contract: contract}, nil
}

// NewDefiReferenceAggregatorFilterer creates a new log filterer instance of DefiReferenceAggregator, bound to a specific deployed contract.
func NewDefiReferenceAggregatorFilterer(address common.Address, filterer bind.ContractFilterer) (*DefiReferenceAggregatorFilterer, error) {
	contract, err := bindDefiReferenceAggregator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DefiReferenceAggregatorFilterer{contract: contract}, nil
}

// bindDefiReferenceAggregator binds a generic wrapper to an already deployed contract.
func bindDefiReferenceAggregator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DefiReferenceAggregatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DefiReferenceAggregator *DefiReferenceAggregatorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DefiReferenceAggregator.Contract.DefiReferenceAggregatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DefiReferenceAggregator *DefiReferenceAggregatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefiReferenceAggregator.Contract.DefiReferenceAggregatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DefiReferenceAggregator *DefiReferenceAggregatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DefiReferenceAggregator.Contract.DefiReferenceAggregatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DefiReferenceAggregator *DefiReferenceAggregatorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DefiReferenceAggregator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DefiReferenceAggregator *DefiReferenceAggregatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefiReferenceAggregator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DefiReferenceAggregator *DefiReferenceAggregatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DefiReferenceAggregator.Contract.contract.Transact(opts, method, params...)
}

// Aggregators is a free data retrieval call binding the contract method 0x112cdab9.
//
// Solidity: function aggregators(address ) view returns(address)
func (_DefiReferenceAggregator *DefiReferenceAggregatorCaller) Aggregators(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DefiReferenceAggregator.contract.Call(opts, out, "aggregators", arg0)
	return *ret0, err
}

// Aggregators is a free data retrieval call binding the contract method 0x112cdab9.
//
// Solidity: function aggregators(address ) view returns(address)
func (_DefiReferenceAggregator *DefiReferenceAggregatorSession) Aggregators(arg0 common.Address) (common.Address, error) {
	return _DefiReferenceAggregator.Contract.Aggregators(&_DefiReferenceAggregator.CallOpts, arg0)
}

// Aggregators is a free data retrieval call binding the contract method 0x112cdab9.
//
// Solidity: function aggregators(address ) view returns(address)
func (_DefiReferenceAggregator *DefiReferenceAggregatorCallerSession) Aggregators(arg0 common.Address) (common.Address, error) {
	return _DefiReferenceAggregator.Contract.Aggregators(&_DefiReferenceAggregator.CallOpts, arg0)
}

// FindTokenIndex is a free data retrieval call binding the contract method 0x0a11fd07.
//
// Solidity: function findTokenIndex(address _addr) view returns(int256)
func (_DefiReferenceAggregator *DefiReferenceAggregatorCaller) FindTokenIndex(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiReferenceAggregator.contract.Call(opts, out, "findTokenIndex", _addr)
	return *ret0, err
}

// FindTokenIndex is a free data retrieval call binding the contract method 0x0a11fd07.
//
// Solidity: function findTokenIndex(address _addr) view returns(int256)
func (_DefiReferenceAggregator *DefiReferenceAggregatorSession) FindTokenIndex(_addr common.Address) (*big.Int, error) {
	return _DefiReferenceAggregator.Contract.FindTokenIndex(&_DefiReferenceAggregator.CallOpts, _addr)
}

// FindTokenIndex is a free data retrieval call binding the contract method 0x0a11fd07.
//
// Solidity: function findTokenIndex(address _addr) view returns(int256)
func (_DefiReferenceAggregator *DefiReferenceAggregatorCallerSession) FindTokenIndex(_addr common.Address) (*big.Int, error) {
	return _DefiReferenceAggregator.Contract.FindTokenIndex(&_DefiReferenceAggregator.CallOpts, _addr)
}

// GetPreviousPrice is a free data retrieval call binding the contract method 0xe889e5d6.
//
// Solidity: function getPreviousPrice(address token, uint256 _back) view returns(int256)
func (_DefiReferenceAggregator *DefiReferenceAggregatorCaller) GetPreviousPrice(opts *bind.CallOpts, token common.Address, _back *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiReferenceAggregator.contract.Call(opts, out, "getPreviousPrice", token, _back)
	return *ret0, err
}

// GetPreviousPrice is a free data retrieval call binding the contract method 0xe889e5d6.
//
// Solidity: function getPreviousPrice(address token, uint256 _back) view returns(int256)
func (_DefiReferenceAggregator *DefiReferenceAggregatorSession) GetPreviousPrice(token common.Address, _back *big.Int) (*big.Int, error) {
	return _DefiReferenceAggregator.Contract.GetPreviousPrice(&_DefiReferenceAggregator.CallOpts, token, _back)
}

// GetPreviousPrice is a free data retrieval call binding the contract method 0xe889e5d6.
//
// Solidity: function getPreviousPrice(address token, uint256 _back) view returns(int256)
func (_DefiReferenceAggregator *DefiReferenceAggregatorCallerSession) GetPreviousPrice(token common.Address, _back *big.Int) (*big.Int, error) {
	return _DefiReferenceAggregator.Contract.GetPreviousPrice(&_DefiReferenceAggregator.CallOpts, token, _back)
}

// GetPreviousTimeStamp is a free data retrieval call binding the contract method 0xe0a3581e.
//
// Solidity: function getPreviousTimeStamp(address token, uint256 _back) view returns(uint256)
func (_DefiReferenceAggregator *DefiReferenceAggregatorCaller) GetPreviousTimeStamp(opts *bind.CallOpts, token common.Address, _back *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiReferenceAggregator.contract.Call(opts, out, "getPreviousTimeStamp", token, _back)
	return *ret0, err
}

// GetPreviousTimeStamp is a free data retrieval call binding the contract method 0xe0a3581e.
//
// Solidity: function getPreviousTimeStamp(address token, uint256 _back) view returns(uint256)
func (_DefiReferenceAggregator *DefiReferenceAggregatorSession) GetPreviousTimeStamp(token common.Address, _back *big.Int) (*big.Int, error) {
	return _DefiReferenceAggregator.Contract.GetPreviousTimeStamp(&_DefiReferenceAggregator.CallOpts, token, _back)
}

// GetPreviousTimeStamp is a free data retrieval call binding the contract method 0xe0a3581e.
//
// Solidity: function getPreviousTimeStamp(address token, uint256 _back) view returns(uint256)
func (_DefiReferenceAggregator *DefiReferenceAggregatorCallerSession) GetPreviousTimeStamp(token common.Address, _back *big.Int) (*big.Int, error) {
	return _DefiReferenceAggregator.Contract.GetPreviousTimeStamp(&_DefiReferenceAggregator.CallOpts, token, _back)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address token) view returns(int256)
func (_DefiReferenceAggregator *DefiReferenceAggregatorCaller) GetPrice(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiReferenceAggregator.contract.Call(opts, out, "getPrice", token)
	return *ret0, err
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address token) view returns(int256)
func (_DefiReferenceAggregator *DefiReferenceAggregatorSession) GetPrice(token common.Address) (*big.Int, error) {
	return _DefiReferenceAggregator.Contract.GetPrice(&_DefiReferenceAggregator.CallOpts, token)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address token) view returns(int256)
func (_DefiReferenceAggregator *DefiReferenceAggregatorCallerSession) GetPrice(token common.Address) (*big.Int, error) {
	return _DefiReferenceAggregator.Contract.GetPrice(&_DefiReferenceAggregator.CallOpts, token)
}

// GetTimeStamp is a free data retrieval call binding the contract method 0xeb470ebf.
//
// Solidity: function getTimeStamp(address token) view returns(uint256)
func (_DefiReferenceAggregator *DefiReferenceAggregatorCaller) GetTimeStamp(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiReferenceAggregator.contract.Call(opts, out, "getTimeStamp", token)
	return *ret0, err
}

// GetTimeStamp is a free data retrieval call binding the contract method 0xeb470ebf.
//
// Solidity: function getTimeStamp(address token) view returns(uint256)
func (_DefiReferenceAggregator *DefiReferenceAggregatorSession) GetTimeStamp(token common.Address) (*big.Int, error) {
	return _DefiReferenceAggregator.Contract.GetTimeStamp(&_DefiReferenceAggregator.CallOpts, token)
}

// GetTimeStamp is a free data retrieval call binding the contract method 0xeb470ebf.
//
// Solidity: function getTimeStamp(address token) view returns(uint256)
func (_DefiReferenceAggregator *DefiReferenceAggregatorCallerSession) GetTimeStamp(token common.Address) (*big.Int, error) {
	return _DefiReferenceAggregator.Contract.GetTimeStamp(&_DefiReferenceAggregator.CallOpts, token)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DefiReferenceAggregator *DefiReferenceAggregatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DefiReferenceAggregator.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DefiReferenceAggregator *DefiReferenceAggregatorSession) Owner() (common.Address, error) {
	return _DefiReferenceAggregator.Contract.Owner(&_DefiReferenceAggregator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DefiReferenceAggregator *DefiReferenceAggregatorCallerSession) Owner() (common.Address, error) {
	return _DefiReferenceAggregator.Contract.Owner(&_DefiReferenceAggregator.CallOpts)
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address addr, bytes32 name, bytes32 symbol, uint256 decimals, bool isActive, bool canDeposit, bool canBorrow, bool canTrade, uint256 volatility)
func (_DefiReferenceAggregator *DefiReferenceAggregatorCaller) Tokens(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Addr       common.Address
	Name       [32]byte
	Symbol     [32]byte
	Decimals   *big.Int
	IsActive   bool
	CanDeposit bool
	CanBorrow  bool
	CanTrade   bool
	Volatility *big.Int
}, error) {
	ret := new(struct {
		Addr       common.Address
		Name       [32]byte
		Symbol     [32]byte
		Decimals   *big.Int
		IsActive   bool
		CanDeposit bool
		CanBorrow  bool
		CanTrade   bool
		Volatility *big.Int
	})
	out := ret
	err := _DefiReferenceAggregator.contract.Call(opts, out, "tokens", arg0)
	return *ret, err
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address addr, bytes32 name, bytes32 symbol, uint256 decimals, bool isActive, bool canDeposit, bool canBorrow, bool canTrade, uint256 volatility)
func (_DefiReferenceAggregator *DefiReferenceAggregatorSession) Tokens(arg0 *big.Int) (struct {
	Addr       common.Address
	Name       [32]byte
	Symbol     [32]byte
	Decimals   *big.Int
	IsActive   bool
	CanDeposit bool
	CanBorrow  bool
	CanTrade   bool
	Volatility *big.Int
}, error) {
	return _DefiReferenceAggregator.Contract.Tokens(&_DefiReferenceAggregator.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address addr, bytes32 name, bytes32 symbol, uint256 decimals, bool isActive, bool canDeposit, bool canBorrow, bool canTrade, uint256 volatility)
func (_DefiReferenceAggregator *DefiReferenceAggregatorCallerSession) Tokens(arg0 *big.Int) (struct {
	Addr       common.Address
	Name       [32]byte
	Symbol     [32]byte
	Decimals   *big.Int
	IsActive   bool
	CanDeposit bool
	CanBorrow  bool
	CanTrade   bool
	Volatility *big.Int
}, error) {
	return _DefiReferenceAggregator.Contract.Tokens(&_DefiReferenceAggregator.CallOpts, arg0)
}

// TokensCount is a free data retrieval call binding the contract method 0xa64ed8ba.
//
// Solidity: function tokensCount() view returns(uint256)
func (_DefiReferenceAggregator *DefiReferenceAggregatorCaller) TokensCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiReferenceAggregator.contract.Call(opts, out, "tokensCount")
	return *ret0, err
}

// TokensCount is a free data retrieval call binding the contract method 0xa64ed8ba.
//
// Solidity: function tokensCount() view returns(uint256)
func (_DefiReferenceAggregator *DefiReferenceAggregatorSession) TokensCount() (*big.Int, error) {
	return _DefiReferenceAggregator.Contract.TokensCount(&_DefiReferenceAggregator.CallOpts)
}

// TokensCount is a free data retrieval call binding the contract method 0xa64ed8ba.
//
// Solidity: function tokensCount() view returns(uint256)
func (_DefiReferenceAggregator *DefiReferenceAggregatorCallerSession) TokensCount() (*big.Int, error) {
	return _DefiReferenceAggregator.Contract.TokensCount(&_DefiReferenceAggregator.CallOpts)
}

// AddToken is a paid mutator transaction binding the contract method 0x592bda0f.
//
// Solidity: function addToken(address _addr, bytes32 _name, bytes32 _symbol, uint256 _decimals, bool _isActive, bool _canDeposit, bool _canBorrow, bool _canTrade, uint256 _volatility) returns()
func (_DefiReferenceAggregator *DefiReferenceAggregatorTransactor) AddToken(opts *bind.TransactOpts, _addr common.Address, _name [32]byte, _symbol [32]byte, _decimals *big.Int, _isActive bool, _canDeposit bool, _canBorrow bool, _canTrade bool, _volatility *big.Int) (*types.Transaction, error) {
	return _DefiReferenceAggregator.contract.Transact(opts, "addToken", _addr, _name, _symbol, _decimals, _isActive, _canDeposit, _canBorrow, _canTrade, _volatility)
}

// AddToken is a paid mutator transaction binding the contract method 0x592bda0f.
//
// Solidity: function addToken(address _addr, bytes32 _name, bytes32 _symbol, uint256 _decimals, bool _isActive, bool _canDeposit, bool _canBorrow, bool _canTrade, uint256 _volatility) returns()
func (_DefiReferenceAggregator *DefiReferenceAggregatorSession) AddToken(_addr common.Address, _name [32]byte, _symbol [32]byte, _decimals *big.Int, _isActive bool, _canDeposit bool, _canBorrow bool, _canTrade bool, _volatility *big.Int) (*types.Transaction, error) {
	return _DefiReferenceAggregator.Contract.AddToken(&_DefiReferenceAggregator.TransactOpts, _addr, _name, _symbol, _decimals, _isActive, _canDeposit, _canBorrow, _canTrade, _volatility)
}

// AddToken is a paid mutator transaction binding the contract method 0x592bda0f.
//
// Solidity: function addToken(address _addr, bytes32 _name, bytes32 _symbol, uint256 _decimals, bool _isActive, bool _canDeposit, bool _canBorrow, bool _canTrade, uint256 _volatility) returns()
func (_DefiReferenceAggregator *DefiReferenceAggregatorTransactorSession) AddToken(_addr common.Address, _name [32]byte, _symbol [32]byte, _decimals *big.Int, _isActive bool, _canDeposit bool, _canBorrow bool, _canTrade bool, _volatility *big.Int) (*types.Transaction, error) {
	return _DefiReferenceAggregator.Contract.AddToken(&_DefiReferenceAggregator.TransactOpts, _addr, _name, _symbol, _decimals, _isActive, _canDeposit, _canBorrow, _canTrade, _volatility)
}

// SetAggregator is a paid mutator transaction binding the contract method 0x7394a7f6.
//
// Solidity: function setAggregator(address token, address aggregator) returns()
func (_DefiReferenceAggregator *DefiReferenceAggregatorTransactor) SetAggregator(opts *bind.TransactOpts, token common.Address, aggregator common.Address) (*types.Transaction, error) {
	return _DefiReferenceAggregator.contract.Transact(opts, "setAggregator", token, aggregator)
}

// SetAggregator is a paid mutator transaction binding the contract method 0x7394a7f6.
//
// Solidity: function setAggregator(address token, address aggregator) returns()
func (_DefiReferenceAggregator *DefiReferenceAggregatorSession) SetAggregator(token common.Address, aggregator common.Address) (*types.Transaction, error) {
	return _DefiReferenceAggregator.Contract.SetAggregator(&_DefiReferenceAggregator.TransactOpts, token, aggregator)
}

// SetAggregator is a paid mutator transaction binding the contract method 0x7394a7f6.
//
// Solidity: function setAggregator(address token, address aggregator) returns()
func (_DefiReferenceAggregator *DefiReferenceAggregatorTransactorSession) SetAggregator(token common.Address, aggregator common.Address) (*types.Transaction, error) {
	return _DefiReferenceAggregator.Contract.SetAggregator(&_DefiReferenceAggregator.TransactOpts, token, aggregator)
}

// UpdateToken is a paid mutator transaction binding the contract method 0x5b2e0ff4.
//
// Solidity: function updateToken(address _addr, uint256 _decimals, bool _isActive, bool _canDeposit, bool _canBorrow, bool _canTrade, uint256 _volatility) returns()
func (_DefiReferenceAggregator *DefiReferenceAggregatorTransactor) UpdateToken(opts *bind.TransactOpts, _addr common.Address, _decimals *big.Int, _isActive bool, _canDeposit bool, _canBorrow bool, _canTrade bool, _volatility *big.Int) (*types.Transaction, error) {
	return _DefiReferenceAggregator.contract.Transact(opts, "updateToken", _addr, _decimals, _isActive, _canDeposit, _canBorrow, _canTrade, _volatility)
}

// UpdateToken is a paid mutator transaction binding the contract method 0x5b2e0ff4.
//
// Solidity: function updateToken(address _addr, uint256 _decimals, bool _isActive, bool _canDeposit, bool _canBorrow, bool _canTrade, uint256 _volatility) returns()
func (_DefiReferenceAggregator *DefiReferenceAggregatorSession) UpdateToken(_addr common.Address, _decimals *big.Int, _isActive bool, _canDeposit bool, _canBorrow bool, _canTrade bool, _volatility *big.Int) (*types.Transaction, error) {
	return _DefiReferenceAggregator.Contract.UpdateToken(&_DefiReferenceAggregator.TransactOpts, _addr, _decimals, _isActive, _canDeposit, _canBorrow, _canTrade, _volatility)
}

// UpdateToken is a paid mutator transaction binding the contract method 0x5b2e0ff4.
//
// Solidity: function updateToken(address _addr, uint256 _decimals, bool _isActive, bool _canDeposit, bool _canBorrow, bool _canTrade, uint256 _volatility) returns()
func (_DefiReferenceAggregator *DefiReferenceAggregatorTransactorSession) UpdateToken(_addr common.Address, _decimals *big.Int, _isActive bool, _canDeposit bool, _canBorrow bool, _canTrade bool, _volatility *big.Int) (*types.Transaction, error) {
	return _DefiReferenceAggregator.Contract.UpdateToken(&_DefiReferenceAggregator.TransactOpts, _addr, _decimals, _isActive, _canDeposit, _canBorrow, _canTrade, _volatility)
}

// DefiReferenceAggregatorAggregatorChangedIterator is returned from FilterAggregatorChanged and is used to iterate over the raw logs and unpacked data for AggregatorChanged events raised by the DefiReferenceAggregator contract.
type DefiReferenceAggregatorAggregatorChangedIterator struct {
	Event *DefiReferenceAggregatorAggregatorChanged // Event containing the contract specifics and raw log

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
func (it *DefiReferenceAggregatorAggregatorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefiReferenceAggregatorAggregatorChanged)
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
		it.Event = new(DefiReferenceAggregatorAggregatorChanged)
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
func (it *DefiReferenceAggregatorAggregatorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefiReferenceAggregatorAggregatorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefiReferenceAggregatorAggregatorChanged represents a AggregatorChanged event raised by the DefiReferenceAggregator contract.
type DefiReferenceAggregatorAggregatorChanged struct {
	Token      common.Address
	Aggregator common.Address
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAggregatorChanged is a free log retrieval operation binding the contract event 0x51012bc010926f3f948f99b90ed487a209375a900797c69e70dbc17f09264357.
//
// Solidity: event AggregatorChanged(address indexed token, address aggregator, uint256 timestamp)
func (_DefiReferenceAggregator *DefiReferenceAggregatorFilterer) FilterAggregatorChanged(opts *bind.FilterOpts, token []common.Address) (*DefiReferenceAggregatorAggregatorChangedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _DefiReferenceAggregator.contract.FilterLogs(opts, "AggregatorChanged", tokenRule)
	if err != nil {
		return nil, err
	}
	return &DefiReferenceAggregatorAggregatorChangedIterator{contract: _DefiReferenceAggregator.contract, event: "AggregatorChanged", logs: logs, sub: sub}, nil
}

// WatchAggregatorChanged is a free log subscription operation binding the contract event 0x51012bc010926f3f948f99b90ed487a209375a900797c69e70dbc17f09264357.
//
// Solidity: event AggregatorChanged(address indexed token, address aggregator, uint256 timestamp)
func (_DefiReferenceAggregator *DefiReferenceAggregatorFilterer) WatchAggregatorChanged(opts *bind.WatchOpts, sink chan<- *DefiReferenceAggregatorAggregatorChanged, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _DefiReferenceAggregator.contract.WatchLogs(opts, "AggregatorChanged", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefiReferenceAggregatorAggregatorChanged)
				if err := _DefiReferenceAggregator.contract.UnpackLog(event, "AggregatorChanged", log); err != nil {
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
func (_DefiReferenceAggregator *DefiReferenceAggregatorFilterer) ParseAggregatorChanged(log types.Log) (*DefiReferenceAggregatorAggregatorChanged, error) {
	event := new(DefiReferenceAggregatorAggregatorChanged)
	if err := _DefiReferenceAggregator.contract.UnpackLog(event, "AggregatorChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DefiReferenceAggregatorTokenInformationAddedIterator is returned from FilterTokenInformationAdded and is used to iterate over the raw logs and unpacked data for TokenInformationAdded events raised by the DefiReferenceAggregator contract.
type DefiReferenceAggregatorTokenInformationAddedIterator struct {
	Event *DefiReferenceAggregatorTokenInformationAdded // Event containing the contract specifics and raw log

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
func (it *DefiReferenceAggregatorTokenInformationAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefiReferenceAggregatorTokenInformationAdded)
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
		it.Event = new(DefiReferenceAggregatorTokenInformationAdded)
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
func (it *DefiReferenceAggregatorTokenInformationAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefiReferenceAggregatorTokenInformationAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefiReferenceAggregatorTokenInformationAdded represents a TokenInformationAdded event raised by the DefiReferenceAggregator contract.
type DefiReferenceAggregatorTokenInformationAdded struct {
	Token     common.Address
	Index     *big.Int
	Name      [32]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokenInformationAdded is a free log retrieval operation binding the contract event 0xe09375923f54cf7415d35fcba0e0c9b19b26966ca902f1d51da971a40fb6d5e7.
//
// Solidity: event TokenInformationAdded(address indexed token, uint256 index, bytes32 name, uint256 timestamp)
func (_DefiReferenceAggregator *DefiReferenceAggregatorFilterer) FilterTokenInformationAdded(opts *bind.FilterOpts, token []common.Address) (*DefiReferenceAggregatorTokenInformationAddedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _DefiReferenceAggregator.contract.FilterLogs(opts, "TokenInformationAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return &DefiReferenceAggregatorTokenInformationAddedIterator{contract: _DefiReferenceAggregator.contract, event: "TokenInformationAdded", logs: logs, sub: sub}, nil
}

// WatchTokenInformationAdded is a free log subscription operation binding the contract event 0xe09375923f54cf7415d35fcba0e0c9b19b26966ca902f1d51da971a40fb6d5e7.
//
// Solidity: event TokenInformationAdded(address indexed token, uint256 index, bytes32 name, uint256 timestamp)
func (_DefiReferenceAggregator *DefiReferenceAggregatorFilterer) WatchTokenInformationAdded(opts *bind.WatchOpts, sink chan<- *DefiReferenceAggregatorTokenInformationAdded, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _DefiReferenceAggregator.contract.WatchLogs(opts, "TokenInformationAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefiReferenceAggregatorTokenInformationAdded)
				if err := _DefiReferenceAggregator.contract.UnpackLog(event, "TokenInformationAdded", log); err != nil {
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

// ParseTokenInformationAdded is a log parse operation binding the contract event 0xe09375923f54cf7415d35fcba0e0c9b19b26966ca902f1d51da971a40fb6d5e7.
//
// Solidity: event TokenInformationAdded(address indexed token, uint256 index, bytes32 name, uint256 timestamp)
func (_DefiReferenceAggregator *DefiReferenceAggregatorFilterer) ParseTokenInformationAdded(log types.Log) (*DefiReferenceAggregatorTokenInformationAdded, error) {
	event := new(DefiReferenceAggregatorTokenInformationAdded)
	if err := _DefiReferenceAggregator.contract.UnpackLog(event, "TokenInformationAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DefiReferenceAggregatorTokenInformationChangedIterator is returned from FilterTokenInformationChanged and is used to iterate over the raw logs and unpacked data for TokenInformationChanged events raised by the DefiReferenceAggregator contract.
type DefiReferenceAggregatorTokenInformationChangedIterator struct {
	Event *DefiReferenceAggregatorTokenInformationChanged // Event containing the contract specifics and raw log

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
func (it *DefiReferenceAggregatorTokenInformationChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefiReferenceAggregatorTokenInformationChanged)
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
		it.Event = new(DefiReferenceAggregatorTokenInformationChanged)
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
func (it *DefiReferenceAggregatorTokenInformationChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefiReferenceAggregatorTokenInformationChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefiReferenceAggregatorTokenInformationChanged represents a TokenInformationChanged event raised by the DefiReferenceAggregator contract.
type DefiReferenceAggregatorTokenInformationChanged struct {
	Token     common.Address
	Index     *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokenInformationChanged is a free log retrieval operation binding the contract event 0x736225fc71c38a467b7b94748f75ea4268f87bcb9678d6eda334a82d88bbec80.
//
// Solidity: event TokenInformationChanged(address indexed token, uint256 index, uint256 timestamp)
func (_DefiReferenceAggregator *DefiReferenceAggregatorFilterer) FilterTokenInformationChanged(opts *bind.FilterOpts, token []common.Address) (*DefiReferenceAggregatorTokenInformationChangedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _DefiReferenceAggregator.contract.FilterLogs(opts, "TokenInformationChanged", tokenRule)
	if err != nil {
		return nil, err
	}
	return &DefiReferenceAggregatorTokenInformationChangedIterator{contract: _DefiReferenceAggregator.contract, event: "TokenInformationChanged", logs: logs, sub: sub}, nil
}

// WatchTokenInformationChanged is a free log subscription operation binding the contract event 0x736225fc71c38a467b7b94748f75ea4268f87bcb9678d6eda334a82d88bbec80.
//
// Solidity: event TokenInformationChanged(address indexed token, uint256 index, uint256 timestamp)
func (_DefiReferenceAggregator *DefiReferenceAggregatorFilterer) WatchTokenInformationChanged(opts *bind.WatchOpts, sink chan<- *DefiReferenceAggregatorTokenInformationChanged, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _DefiReferenceAggregator.contract.WatchLogs(opts, "TokenInformationChanged", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefiReferenceAggregatorTokenInformationChanged)
				if err := _DefiReferenceAggregator.contract.UnpackLog(event, "TokenInformationChanged", log); err != nil {
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
func (_DefiReferenceAggregator *DefiReferenceAggregatorFilterer) ParseTokenInformationChanged(log types.Log) (*DefiReferenceAggregatorTokenInformationChanged, error) {
	event := new(DefiReferenceAggregatorTokenInformationChanged)
	if err := _DefiReferenceAggregator.contract.UnpackLog(event, "TokenInformationChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}
