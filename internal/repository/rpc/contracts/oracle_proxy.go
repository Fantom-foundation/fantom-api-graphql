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

// PriceOracleProxyInterfaceMetaData contains all meta data concerning the PriceOracleProxyInterface contract.
var PriceOracleProxyInterfaceMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PriceOracleProxyInterfaceABI is the input ABI used to generate the binding from.
// Deprecated: Use PriceOracleProxyInterfaceMetaData.ABI instead.
var PriceOracleProxyInterfaceABI = PriceOracleProxyInterfaceMetaData.ABI

// PriceOracleProxyInterface is an auto generated Go binding around an Ethereum contract.
type PriceOracleProxyInterface struct {
	PriceOracleProxyInterfaceCaller     // Read-only binding to the contract
	PriceOracleProxyInterfaceTransactor // Write-only binding to the contract
	PriceOracleProxyInterfaceFilterer   // Log filterer for contract events
}

// PriceOracleProxyInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type PriceOracleProxyInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceOracleProxyInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PriceOracleProxyInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceOracleProxyInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PriceOracleProxyInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceOracleProxyInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PriceOracleProxyInterfaceSession struct {
	Contract     *PriceOracleProxyInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// PriceOracleProxyInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PriceOracleProxyInterfaceCallerSession struct {
	Contract *PriceOracleProxyInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// PriceOracleProxyInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PriceOracleProxyInterfaceTransactorSession struct {
	Contract     *PriceOracleProxyInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// PriceOracleProxyInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type PriceOracleProxyInterfaceRaw struct {
	Contract *PriceOracleProxyInterface // Generic contract binding to access the raw methods on
}

// PriceOracleProxyInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PriceOracleProxyInterfaceCallerRaw struct {
	Contract *PriceOracleProxyInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// PriceOracleProxyInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PriceOracleProxyInterfaceTransactorRaw struct {
	Contract *PriceOracleProxyInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPriceOracleProxyInterface creates a new instance of PriceOracleProxyInterface, bound to a specific deployed contract.
func NewPriceOracleProxyInterface(address common.Address, backend bind.ContractBackend) (*PriceOracleProxyInterface, error) {
	contract, err := bindPriceOracleProxyInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PriceOracleProxyInterface{PriceOracleProxyInterfaceCaller: PriceOracleProxyInterfaceCaller{contract: contract}, PriceOracleProxyInterfaceTransactor: PriceOracleProxyInterfaceTransactor{contract: contract}, PriceOracleProxyInterfaceFilterer: PriceOracleProxyInterfaceFilterer{contract: contract}}, nil
}

// NewPriceOracleProxyInterfaceCaller creates a new read-only instance of PriceOracleProxyInterface, bound to a specific deployed contract.
func NewPriceOracleProxyInterfaceCaller(address common.Address, caller bind.ContractCaller) (*PriceOracleProxyInterfaceCaller, error) {
	contract, err := bindPriceOracleProxyInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PriceOracleProxyInterfaceCaller{contract: contract}, nil
}

// NewPriceOracleProxyInterfaceTransactor creates a new write-only instance of PriceOracleProxyInterface, bound to a specific deployed contract.
func NewPriceOracleProxyInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*PriceOracleProxyInterfaceTransactor, error) {
	contract, err := bindPriceOracleProxyInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PriceOracleProxyInterfaceTransactor{contract: contract}, nil
}

// NewPriceOracleProxyInterfaceFilterer creates a new log filterer instance of PriceOracleProxyInterface, bound to a specific deployed contract.
func NewPriceOracleProxyInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*PriceOracleProxyInterfaceFilterer, error) {
	contract, err := bindPriceOracleProxyInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PriceOracleProxyInterfaceFilterer{contract: contract}, nil
}

// bindPriceOracleProxyInterface binds a generic wrapper to an already deployed contract.
func bindPriceOracleProxyInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PriceOracleProxyInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceOracleProxyInterface *PriceOracleProxyInterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceOracleProxyInterface.Contract.PriceOracleProxyInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceOracleProxyInterface *PriceOracleProxyInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceOracleProxyInterface.Contract.PriceOracleProxyInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceOracleProxyInterface *PriceOracleProxyInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceOracleProxyInterface.Contract.PriceOracleProxyInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceOracleProxyInterface *PriceOracleProxyInterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceOracleProxyInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceOracleProxyInterface *PriceOracleProxyInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceOracleProxyInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceOracleProxyInterface *PriceOracleProxyInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceOracleProxyInterface.Contract.contract.Transact(opts, method, params...)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address _token) view returns(uint256)
func (_PriceOracleProxyInterface *PriceOracleProxyInterfaceCaller) GetPrice(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PriceOracleProxyInterface.contract.Call(opts, &out, "getPrice", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address _token) view returns(uint256)
func (_PriceOracleProxyInterface *PriceOracleProxyInterfaceSession) GetPrice(_token common.Address) (*big.Int, error) {
	return _PriceOracleProxyInterface.Contract.GetPrice(&_PriceOracleProxyInterface.CallOpts, _token)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address _token) view returns(uint256)
func (_PriceOracleProxyInterface *PriceOracleProxyInterfaceCallerSession) GetPrice(_token common.Address) (*big.Int, error) {
	return _PriceOracleProxyInterface.Contract.GetPrice(&_PriceOracleProxyInterface.CallOpts, _token)
}
