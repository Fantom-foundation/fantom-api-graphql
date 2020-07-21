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

// DefiLiquidityPoolABI is the input ABI used to generate the binding from.
const DefiLiquidityPoolABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exchangeRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Buy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"PriceOracleChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Repay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exchangeRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Sell\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_collateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_collateralList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_collateralTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_collateralValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_debt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_debtList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_debtTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_debtValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_minted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_mintedTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"colLiquidationRatio4dec\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"colLowestRatio4dec\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"colWarningRatio4dec\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fUsdToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feePool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"loanEntryFee4dec\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nativeToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"poolVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"priceOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ratioDecimalsCorrection\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"setPriceOracleReferenceAggregate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tradeFee4dec\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"collateralListCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"debtListCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"priceDigitsCorrection\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"collateralValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cValue\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"debtValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dValue\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"buy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"sell\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"trade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"borrow\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"repay\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"liquidate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DefiLiquidityPool is an auto generated Go binding around an Ethereum contract.
type DefiLiquidityPool struct {
	DefiLiquidityPoolCaller     // Read-only binding to the contract
	DefiLiquidityPoolTransactor // Write-only binding to the contract
	DefiLiquidityPoolFilterer   // Log filterer for contract events
}

// DefiLiquidityPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type DefiLiquidityPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefiLiquidityPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DefiLiquidityPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefiLiquidityPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DefiLiquidityPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefiLiquidityPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DefiLiquidityPoolSession struct {
	Contract     *DefiLiquidityPool // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DefiLiquidityPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DefiLiquidityPoolCallerSession struct {
	Contract *DefiLiquidityPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// DefiLiquidityPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DefiLiquidityPoolTransactorSession struct {
	Contract     *DefiLiquidityPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// DefiLiquidityPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type DefiLiquidityPoolRaw struct {
	Contract *DefiLiquidityPool // Generic contract binding to access the raw methods on
}

// DefiLiquidityPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DefiLiquidityPoolCallerRaw struct {
	Contract *DefiLiquidityPoolCaller // Generic read-only contract binding to access the raw methods on
}

// DefiLiquidityPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DefiLiquidityPoolTransactorRaw struct {
	Contract *DefiLiquidityPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDefiLiquidityPool creates a new instance of DefiLiquidityPool, bound to a specific deployed contract.
func NewDefiLiquidityPool(address common.Address, backend bind.ContractBackend) (*DefiLiquidityPool, error) {
	contract, err := bindDefiLiquidityPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DefiLiquidityPool{DefiLiquidityPoolCaller: DefiLiquidityPoolCaller{contract: contract}, DefiLiquidityPoolTransactor: DefiLiquidityPoolTransactor{contract: contract}, DefiLiquidityPoolFilterer: DefiLiquidityPoolFilterer{contract: contract}}, nil
}

// NewDefiLiquidityPoolCaller creates a new read-only instance of DefiLiquidityPool, bound to a specific deployed contract.
func NewDefiLiquidityPoolCaller(address common.Address, caller bind.ContractCaller) (*DefiLiquidityPoolCaller, error) {
	contract, err := bindDefiLiquidityPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DefiLiquidityPoolCaller{contract: contract}, nil
}

// NewDefiLiquidityPoolTransactor creates a new write-only instance of DefiLiquidityPool, bound to a specific deployed contract.
func NewDefiLiquidityPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*DefiLiquidityPoolTransactor, error) {
	contract, err := bindDefiLiquidityPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DefiLiquidityPoolTransactor{contract: contract}, nil
}

// NewDefiLiquidityPoolFilterer creates a new log filterer instance of DefiLiquidityPool, bound to a specific deployed contract.
func NewDefiLiquidityPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*DefiLiquidityPoolFilterer, error) {
	contract, err := bindDefiLiquidityPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DefiLiquidityPoolFilterer{contract: contract}, nil
}

// bindDefiLiquidityPool binds a generic wrapper to an already deployed contract.
func bindDefiLiquidityPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DefiLiquidityPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DefiLiquidityPool *DefiLiquidityPoolRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DefiLiquidityPool.Contract.DefiLiquidityPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DefiLiquidityPool *DefiLiquidityPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefiLiquidityPool.Contract.DefiLiquidityPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DefiLiquidityPool *DefiLiquidityPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DefiLiquidityPool.Contract.DefiLiquidityPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DefiLiquidityPool *DefiLiquidityPoolCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DefiLiquidityPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DefiLiquidityPool *DefiLiquidityPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefiLiquidityPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DefiLiquidityPool *DefiLiquidityPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DefiLiquidityPool.Contract.contract.Transact(opts, method, params...)
}

// Collateral is a free data retrieval call binding the contract method 0x1f91e4c8.
//
// Solidity: function _collateral(address , address ) view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolCaller) Collateral(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiLiquidityPool.contract.Call(opts, out, "_collateral", arg0, arg1)
	return *ret0, err
}

// Collateral is a free data retrieval call binding the contract method 0x1f91e4c8.
//
// Solidity: function _collateral(address , address ) view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolSession) Collateral(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DefiLiquidityPool.Contract.Collateral(&_DefiLiquidityPool.CallOpts, arg0, arg1)
}

// Collateral is a free data retrieval call binding the contract method 0x1f91e4c8.
//
// Solidity: function _collateral(address , address ) view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolCallerSession) Collateral(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DefiLiquidityPool.Contract.Collateral(&_DefiLiquidityPool.CallOpts, arg0, arg1)
}

// CollateralList is a free data retrieval call binding the contract method 0xf9946002.
//
// Solidity: function _collateralList(address , uint256 ) view returns(address)
func (_DefiLiquidityPool *DefiLiquidityPoolCaller) CollateralList(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DefiLiquidityPool.contract.Call(opts, out, "_collateralList", arg0, arg1)
	return *ret0, err
}

// CollateralList is a free data retrieval call binding the contract method 0xf9946002.
//
// Solidity: function _collateralList(address , uint256 ) view returns(address)
func (_DefiLiquidityPool *DefiLiquidityPoolSession) CollateralList(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _DefiLiquidityPool.Contract.CollateralList(&_DefiLiquidityPool.CallOpts, arg0, arg1)
}

// CollateralList is a free data retrieval call binding the contract method 0xf9946002.
//
// Solidity: function _collateralList(address , uint256 ) view returns(address)
func (_DefiLiquidityPool *DefiLiquidityPoolCallerSession) CollateralList(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _DefiLiquidityPool.Contract.CollateralList(&_DefiLiquidityPool.CallOpts, arg0, arg1)
}

// CollateralTokens is a free data retrieval call binding the contract method 0xd2d0ec48.
//
// Solidity: function _collateralTokens(address , address ) view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolCaller) CollateralTokens(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiLiquidityPool.contract.Call(opts, out, "_collateralTokens", arg0, arg1)
	return *ret0, err
}

// CollateralTokens is a free data retrieval call binding the contract method 0xd2d0ec48.
//
// Solidity: function _collateralTokens(address , address ) view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolSession) CollateralTokens(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DefiLiquidityPool.Contract.CollateralTokens(&_DefiLiquidityPool.CallOpts, arg0, arg1)
}

// CollateralTokens is a free data retrieval call binding the contract method 0xd2d0ec48.
//
// Solidity: function _collateralTokens(address , address ) view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolCallerSession) CollateralTokens(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DefiLiquidityPool.Contract.CollateralTokens(&_DefiLiquidityPool.CallOpts, arg0, arg1)
}

// StoredCollateralValue is a free data retrieval call binding the contract method 0x25506d4e.
//
// Solidity: function _collateralValue(address ) view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolCaller) StoredCollateralValue(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiLiquidityPool.contract.Call(opts, out, "_collateralValue", arg0)
	return *ret0, err
}

// StoredCollateralValue is a free data retrieval call binding the contract method 0x25506d4e.
//
// Solidity: function _collateralValue(address ) view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolSession) StoredCollateralValue(arg0 common.Address) (*big.Int, error) {
	return _DefiLiquidityPool.Contract.StoredCollateralValue(&_DefiLiquidityPool.CallOpts, arg0)
}

// StoredCollateralValue is a free data retrieval call binding the contract method 0x25506d4e.
//
// Solidity: function _collateralValue(address ) view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolCallerSession) StoredCollateralValue(arg0 common.Address) (*big.Int, error) {
	return _DefiLiquidityPool.Contract.StoredCollateralValue(&_DefiLiquidityPool.CallOpts, arg0)
}

// Debt is a free data retrieval call binding the contract method 0x2f6d0bf9.
//
// Solidity: function _debt(address , address ) view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolCaller) Debt(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiLiquidityPool.contract.Call(opts, out, "_debt", arg0, arg1)
	return *ret0, err
}

// Debt is a free data retrieval call binding the contract method 0x2f6d0bf9.
//
// Solidity: function _debt(address , address ) view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolSession) Debt(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DefiLiquidityPool.Contract.Debt(&_DefiLiquidityPool.CallOpts, arg0, arg1)
}

// Debt is a free data retrieval call binding the contract method 0x2f6d0bf9.
//
// Solidity: function _debt(address , address ) view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolCallerSession) Debt(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DefiLiquidityPool.Contract.Debt(&_DefiLiquidityPool.CallOpts, arg0, arg1)
}

// DebtList is a free data retrieval call binding the contract method 0x0b08dcfd.
//
// Solidity: function _debtList(address , uint256 ) view returns(address)
func (_DefiLiquidityPool *DefiLiquidityPoolCaller) DebtList(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DefiLiquidityPool.contract.Call(opts, out, "_debtList", arg0, arg1)
	return *ret0, err
}

// DebtList is a free data retrieval call binding the contract method 0x0b08dcfd.
//
// Solidity: function _debtList(address , uint256 ) view returns(address)
func (_DefiLiquidityPool *DefiLiquidityPoolSession) DebtList(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _DefiLiquidityPool.Contract.DebtList(&_DefiLiquidityPool.CallOpts, arg0, arg1)
}

// DebtList is a free data retrieval call binding the contract method 0x0b08dcfd.
//
// Solidity: function _debtList(address , uint256 ) view returns(address)
func (_DefiLiquidityPool *DefiLiquidityPoolCallerSession) DebtList(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _DefiLiquidityPool.Contract.DebtList(&_DefiLiquidityPool.CallOpts, arg0, arg1)
}

// DebtTokens is a free data retrieval call binding the contract method 0x9a8e7150.
//
// Solidity: function _debtTokens(address , address ) view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolCaller) DebtTokens(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiLiquidityPool.contract.Call(opts, out, "_debtTokens", arg0, arg1)
	return *ret0, err
}

// DebtTokens is a free data retrieval call binding the contract method 0x9a8e7150.
//
// Solidity: function _debtTokens(address , address ) view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolSession) DebtTokens(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DefiLiquidityPool.Contract.DebtTokens(&_DefiLiquidityPool.CallOpts, arg0, arg1)
}

// DebtTokens is a free data retrieval call binding the contract method 0x9a8e7150.
//
// Solidity: function _debtTokens(address , address ) view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolCallerSession) DebtTokens(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DefiLiquidityPool.Contract.DebtTokens(&_DefiLiquidityPool.CallOpts, arg0, arg1)
}

// StoredDebtValue is a free data retrieval call binding the contract method 0xa13cb2f3.
//
// Solidity: function _debtValue(address ) view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolCaller) StoredDebtValue(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiLiquidityPool.contract.Call(opts, out, "_debtValue", arg0)
	return *ret0, err
}

// StoredDebtValue is a free data retrieval call binding the contract method 0xa13cb2f3.
//
// Solidity: function _debtValue(address ) view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolSession) StoredDebtValue(arg0 common.Address) (*big.Int, error) {
	return _DefiLiquidityPool.Contract.StoredDebtValue(&_DefiLiquidityPool.CallOpts, arg0)
}

// StoredDebtValue is a free data retrieval call binding the contract method 0xa13cb2f3.
//
// Solidity: function _debtValue(address ) view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolCallerSession) StoredDebtValue(arg0 common.Address) (*big.Int, error) {
	return _DefiLiquidityPool.Contract.StoredDebtValue(&_DefiLiquidityPool.CallOpts, arg0)
}

// Minted is a free data retrieval call binding the contract method 0x6040237b.
//
// Solidity: function _minted(address , address ) view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolCaller) Minted(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiLiquidityPool.contract.Call(opts, out, "_minted", arg0, arg1)
	return *ret0, err
}

// Minted is a free data retrieval call binding the contract method 0x6040237b.
//
// Solidity: function _minted(address , address ) view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolSession) Minted(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DefiLiquidityPool.Contract.Minted(&_DefiLiquidityPool.CallOpts, arg0, arg1)
}

// Minted is a free data retrieval call binding the contract method 0x6040237b.
//
// Solidity: function _minted(address , address ) view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolCallerSession) Minted(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DefiLiquidityPool.Contract.Minted(&_DefiLiquidityPool.CallOpts, arg0, arg1)
}

// MintedTokens is a free data retrieval call binding the contract method 0xf0ac2911.
//
// Solidity: function _mintedTokens(address , address ) view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolCaller) MintedTokens(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiLiquidityPool.contract.Call(opts, out, "_mintedTokens", arg0, arg1)
	return *ret0, err
}

// MintedTokens is a free data retrieval call binding the contract method 0xf0ac2911.
//
// Solidity: function _mintedTokens(address , address ) view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolSession) MintedTokens(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DefiLiquidityPool.Contract.MintedTokens(&_DefiLiquidityPool.CallOpts, arg0, arg1)
}

// MintedTokens is a free data retrieval call binding the contract method 0xf0ac2911.
//
// Solidity: function _mintedTokens(address , address ) view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolCallerSession) MintedTokens(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DefiLiquidityPool.Contract.MintedTokens(&_DefiLiquidityPool.CallOpts, arg0, arg1)
}

// ColLiquidationRatio4dec is a free data retrieval call binding the contract method 0x81e9d826.
//
// Solidity: function colLiquidationRatio4dec() view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolCaller) ColLiquidationRatio4dec(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiLiquidityPool.contract.Call(opts, out, "colLiquidationRatio4dec")
	return *ret0, err
}

// ColLiquidationRatio4dec is a free data retrieval call binding the contract method 0x81e9d826.
//
// Solidity: function colLiquidationRatio4dec() view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolSession) ColLiquidationRatio4dec() (*big.Int, error) {
	return _DefiLiquidityPool.Contract.ColLiquidationRatio4dec(&_DefiLiquidityPool.CallOpts)
}

// ColLiquidationRatio4dec is a free data retrieval call binding the contract method 0x81e9d826.
//
// Solidity: function colLiquidationRatio4dec() view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolCallerSession) ColLiquidationRatio4dec() (*big.Int, error) {
	return _DefiLiquidityPool.Contract.ColLiquidationRatio4dec(&_DefiLiquidityPool.CallOpts)
}

// ColLowestRatio4dec is a free data retrieval call binding the contract method 0xaf4a20e7.
//
// Solidity: function colLowestRatio4dec() view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolCaller) ColLowestRatio4dec(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiLiquidityPool.contract.Call(opts, out, "colLowestRatio4dec")
	return *ret0, err
}

// ColLowestRatio4dec is a free data retrieval call binding the contract method 0xaf4a20e7.
//
// Solidity: function colLowestRatio4dec() view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolSession) ColLowestRatio4dec() (*big.Int, error) {
	return _DefiLiquidityPool.Contract.ColLowestRatio4dec(&_DefiLiquidityPool.CallOpts)
}

// ColLowestRatio4dec is a free data retrieval call binding the contract method 0xaf4a20e7.
//
// Solidity: function colLowestRatio4dec() view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolCallerSession) ColLowestRatio4dec() (*big.Int, error) {
	return _DefiLiquidityPool.Contract.ColLowestRatio4dec(&_DefiLiquidityPool.CallOpts)
}

// ColWarningRatio4dec is a free data retrieval call binding the contract method 0x19132c68.
//
// Solidity: function colWarningRatio4dec() view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolCaller) ColWarningRatio4dec(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiLiquidityPool.contract.Call(opts, out, "colWarningRatio4dec")
	return *ret0, err
}

// ColWarningRatio4dec is a free data retrieval call binding the contract method 0x19132c68.
//
// Solidity: function colWarningRatio4dec() view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolSession) ColWarningRatio4dec() (*big.Int, error) {
	return _DefiLiquidityPool.Contract.ColWarningRatio4dec(&_DefiLiquidityPool.CallOpts)
}

// ColWarningRatio4dec is a free data retrieval call binding the contract method 0x19132c68.
//
// Solidity: function colWarningRatio4dec() view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolCallerSession) ColWarningRatio4dec() (*big.Int, error) {
	return _DefiLiquidityPool.Contract.ColWarningRatio4dec(&_DefiLiquidityPool.CallOpts)
}

// CollateralListCount is a free data retrieval call binding the contract method 0x49222d72.
//
// Solidity: function collateralListCount(address _addr) view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolCaller) CollateralListCount(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiLiquidityPool.contract.Call(opts, out, "collateralListCount", _addr)
	return *ret0, err
}

// CollateralListCount is a free data retrieval call binding the contract method 0x49222d72.
//
// Solidity: function collateralListCount(address _addr) view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolSession) CollateralListCount(_addr common.Address) (*big.Int, error) {
	return _DefiLiquidityPool.Contract.CollateralListCount(&_DefiLiquidityPool.CallOpts, _addr)
}

// CollateralListCount is a free data retrieval call binding the contract method 0x49222d72.
//
// Solidity: function collateralListCount(address _addr) view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolCallerSession) CollateralListCount(_addr common.Address) (*big.Int, error) {
	return _DefiLiquidityPool.Contract.CollateralListCount(&_DefiLiquidityPool.CallOpts, _addr)
}

// CollateralValue is a free data retrieval call binding the contract method 0x7544aa44.
//
// Solidity: function collateralValue(address _user) view returns(uint256 cValue)
func (_DefiLiquidityPool *DefiLiquidityPoolCaller) CollateralValue(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiLiquidityPool.contract.Call(opts, out, "collateralValue", _user)
	return *ret0, err
}

// CollateralValue is a free data retrieval call binding the contract method 0x7544aa44.
//
// Solidity: function collateralValue(address _user) view returns(uint256 cValue)
func (_DefiLiquidityPool *DefiLiquidityPoolSession) CollateralValue(_user common.Address) (*big.Int, error) {
	return _DefiLiquidityPool.Contract.CollateralValue(&_DefiLiquidityPool.CallOpts, _user)
}

// CollateralValue is a free data retrieval call binding the contract method 0x7544aa44.
//
// Solidity: function collateralValue(address _user) view returns(uint256 cValue)
func (_DefiLiquidityPool *DefiLiquidityPoolCallerSession) CollateralValue(_user common.Address) (*big.Int, error) {
	return _DefiLiquidityPool.Contract.CollateralValue(&_DefiLiquidityPool.CallOpts, _user)
}

// DebtListCount is a free data retrieval call binding the contract method 0xd8459223.
//
// Solidity: function debtListCount(address _addr) view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolCaller) DebtListCount(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiLiquidityPool.contract.Call(opts, out, "debtListCount", _addr)
	return *ret0, err
}

// DebtListCount is a free data retrieval call binding the contract method 0xd8459223.
//
// Solidity: function debtListCount(address _addr) view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolSession) DebtListCount(_addr common.Address) (*big.Int, error) {
	return _DefiLiquidityPool.Contract.DebtListCount(&_DefiLiquidityPool.CallOpts, _addr)
}

// DebtListCount is a free data retrieval call binding the contract method 0xd8459223.
//
// Solidity: function debtListCount(address _addr) view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolCallerSession) DebtListCount(_addr common.Address) (*big.Int, error) {
	return _DefiLiquidityPool.Contract.DebtListCount(&_DefiLiquidityPool.CallOpts, _addr)
}

// DebtValue is a free data retrieval call binding the contract method 0xb9637e49.
//
// Solidity: function debtValue(address _user) view returns(uint256 dValue)
func (_DefiLiquidityPool *DefiLiquidityPoolCaller) DebtValue(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiLiquidityPool.contract.Call(opts, out, "debtValue", _user)
	return *ret0, err
}

// DebtValue is a free data retrieval call binding the contract method 0xb9637e49.
//
// Solidity: function debtValue(address _user) view returns(uint256 dValue)
func (_DefiLiquidityPool *DefiLiquidityPoolSession) DebtValue(_user common.Address) (*big.Int, error) {
	return _DefiLiquidityPool.Contract.DebtValue(&_DefiLiquidityPool.CallOpts, _user)
}

// DebtValue is a free data retrieval call binding the contract method 0xb9637e49.
//
// Solidity: function debtValue(address _user) view returns(uint256 dValue)
func (_DefiLiquidityPool *DefiLiquidityPoolCallerSession) DebtValue(_user common.Address) (*big.Int, error) {
	return _DefiLiquidityPool.Contract.DebtValue(&_DefiLiquidityPool.CallOpts, _user)
}

// FUsdToken is a free data retrieval call binding the contract method 0x0e47f738.
//
// Solidity: function fUsdToken() view returns(address)
func (_DefiLiquidityPool *DefiLiquidityPoolCaller) FUsdToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DefiLiquidityPool.contract.Call(opts, out, "fUsdToken")
	return *ret0, err
}

// FUsdToken is a free data retrieval call binding the contract method 0x0e47f738.
//
// Solidity: function fUsdToken() view returns(address)
func (_DefiLiquidityPool *DefiLiquidityPoolSession) FUsdToken() (common.Address, error) {
	return _DefiLiquidityPool.Contract.FUsdToken(&_DefiLiquidityPool.CallOpts)
}

// FUsdToken is a free data retrieval call binding the contract method 0x0e47f738.
//
// Solidity: function fUsdToken() view returns(address)
func (_DefiLiquidityPool *DefiLiquidityPoolCallerSession) FUsdToken() (common.Address, error) {
	return _DefiLiquidityPool.Contract.FUsdToken(&_DefiLiquidityPool.CallOpts)
}

// FeePool is a free data retrieval call binding the contract method 0xae2e933b.
//
// Solidity: function feePool() view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolCaller) FeePool(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiLiquidityPool.contract.Call(opts, out, "feePool")
	return *ret0, err
}

// FeePool is a free data retrieval call binding the contract method 0xae2e933b.
//
// Solidity: function feePool() view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolSession) FeePool() (*big.Int, error) {
	return _DefiLiquidityPool.Contract.FeePool(&_DefiLiquidityPool.CallOpts)
}

// FeePool is a free data retrieval call binding the contract method 0xae2e933b.
//
// Solidity: function feePool() view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolCallerSession) FeePool() (*big.Int, error) {
	return _DefiLiquidityPool.Contract.FeePool(&_DefiLiquidityPool.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_DefiLiquidityPool *DefiLiquidityPoolCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DefiLiquidityPool.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_DefiLiquidityPool *DefiLiquidityPoolSession) IsOwner() (bool, error) {
	return _DefiLiquidityPool.Contract.IsOwner(&_DefiLiquidityPool.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_DefiLiquidityPool *DefiLiquidityPoolCallerSession) IsOwner() (bool, error) {
	return _DefiLiquidityPool.Contract.IsOwner(&_DefiLiquidityPool.CallOpts)
}

// LoanEntryFee4dec is a free data retrieval call binding the contract method 0xace25b11.
//
// Solidity: function loanEntryFee4dec() view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolCaller) LoanEntryFee4dec(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiLiquidityPool.contract.Call(opts, out, "loanEntryFee4dec")
	return *ret0, err
}

// LoanEntryFee4dec is a free data retrieval call binding the contract method 0xace25b11.
//
// Solidity: function loanEntryFee4dec() view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolSession) LoanEntryFee4dec() (*big.Int, error) {
	return _DefiLiquidityPool.Contract.LoanEntryFee4dec(&_DefiLiquidityPool.CallOpts)
}

// LoanEntryFee4dec is a free data retrieval call binding the contract method 0xace25b11.
//
// Solidity: function loanEntryFee4dec() view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolCallerSession) LoanEntryFee4dec() (*big.Int, error) {
	return _DefiLiquidityPool.Contract.LoanEntryFee4dec(&_DefiLiquidityPool.CallOpts)
}

// NativeToken is a free data retrieval call binding the contract method 0xe1758bd8.
//
// Solidity: function nativeToken() view returns(address)
func (_DefiLiquidityPool *DefiLiquidityPoolCaller) NativeToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DefiLiquidityPool.contract.Call(opts, out, "nativeToken")
	return *ret0, err
}

// NativeToken is a free data retrieval call binding the contract method 0xe1758bd8.
//
// Solidity: function nativeToken() view returns(address)
func (_DefiLiquidityPool *DefiLiquidityPoolSession) NativeToken() (common.Address, error) {
	return _DefiLiquidityPool.Contract.NativeToken(&_DefiLiquidityPool.CallOpts)
}

// NativeToken is a free data retrieval call binding the contract method 0xe1758bd8.
//
// Solidity: function nativeToken() view returns(address)
func (_DefiLiquidityPool *DefiLiquidityPoolCallerSession) NativeToken() (common.Address, error) {
	return _DefiLiquidityPool.Contract.NativeToken(&_DefiLiquidityPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DefiLiquidityPool *DefiLiquidityPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DefiLiquidityPool.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DefiLiquidityPool *DefiLiquidityPoolSession) Owner() (common.Address, error) {
	return _DefiLiquidityPool.Contract.Owner(&_DefiLiquidityPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DefiLiquidityPool *DefiLiquidityPoolCallerSession) Owner() (common.Address, error) {
	return _DefiLiquidityPool.Contract.Owner(&_DefiLiquidityPool.CallOpts)
}

// PoolVersion is a free data retrieval call binding the contract method 0xda815731.
//
// Solidity: function poolVersion() view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolCaller) PoolVersion(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiLiquidityPool.contract.Call(opts, out, "poolVersion")
	return *ret0, err
}

// PoolVersion is a free data retrieval call binding the contract method 0xda815731.
//
// Solidity: function poolVersion() view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolSession) PoolVersion() (*big.Int, error) {
	return _DefiLiquidityPool.Contract.PoolVersion(&_DefiLiquidityPool.CallOpts)
}

// PoolVersion is a free data retrieval call binding the contract method 0xda815731.
//
// Solidity: function poolVersion() view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolCallerSession) PoolVersion() (*big.Int, error) {
	return _DefiLiquidityPool.Contract.PoolVersion(&_DefiLiquidityPool.CallOpts)
}

// PriceDigitsCorrection is a free data retrieval call binding the contract method 0x11a0abf1.
//
// Solidity: function priceDigitsCorrection() pure returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolCaller) PriceDigitsCorrection(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiLiquidityPool.contract.Call(opts, out, "priceDigitsCorrection")
	return *ret0, err
}

// PriceDigitsCorrection is a free data retrieval call binding the contract method 0x11a0abf1.
//
// Solidity: function priceDigitsCorrection() pure returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolSession) PriceDigitsCorrection() (*big.Int, error) {
	return _DefiLiquidityPool.Contract.PriceDigitsCorrection(&_DefiLiquidityPool.CallOpts)
}

// PriceDigitsCorrection is a free data retrieval call binding the contract method 0x11a0abf1.
//
// Solidity: function priceDigitsCorrection() pure returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolCallerSession) PriceDigitsCorrection() (*big.Int, error) {
	return _DefiLiquidityPool.Contract.PriceDigitsCorrection(&_DefiLiquidityPool.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x2630c12f.
//
// Solidity: function priceOracle() view returns(address)
func (_DefiLiquidityPool *DefiLiquidityPoolCaller) PriceOracle(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DefiLiquidityPool.contract.Call(opts, out, "priceOracle")
	return *ret0, err
}

// PriceOracle is a free data retrieval call binding the contract method 0x2630c12f.
//
// Solidity: function priceOracle() view returns(address)
func (_DefiLiquidityPool *DefiLiquidityPoolSession) PriceOracle() (common.Address, error) {
	return _DefiLiquidityPool.Contract.PriceOracle(&_DefiLiquidityPool.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x2630c12f.
//
// Solidity: function priceOracle() view returns(address)
func (_DefiLiquidityPool *DefiLiquidityPoolCallerSession) PriceOracle() (common.Address, error) {
	return _DefiLiquidityPool.Contract.PriceOracle(&_DefiLiquidityPool.CallOpts)
}

// RatioDecimalsCorrection is a free data retrieval call binding the contract method 0xae04d884.
//
// Solidity: function ratioDecimalsCorrection() view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolCaller) RatioDecimalsCorrection(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiLiquidityPool.contract.Call(opts, out, "ratioDecimalsCorrection")
	return *ret0, err
}

// RatioDecimalsCorrection is a free data retrieval call binding the contract method 0xae04d884.
//
// Solidity: function ratioDecimalsCorrection() view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolSession) RatioDecimalsCorrection() (*big.Int, error) {
	return _DefiLiquidityPool.Contract.RatioDecimalsCorrection(&_DefiLiquidityPool.CallOpts)
}

// RatioDecimalsCorrection is a free data retrieval call binding the contract method 0xae04d884.
//
// Solidity: function ratioDecimalsCorrection() view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolCallerSession) RatioDecimalsCorrection() (*big.Int, error) {
	return _DefiLiquidityPool.Contract.RatioDecimalsCorrection(&_DefiLiquidityPool.CallOpts)
}

// TradeFee4dec is a free data retrieval call binding the contract method 0xde61bcda.
//
// Solidity: function tradeFee4dec() view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolCaller) TradeFee4dec(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiLiquidityPool.contract.Call(opts, out, "tradeFee4dec")
	return *ret0, err
}

// TradeFee4dec is a free data retrieval call binding the contract method 0xde61bcda.
//
// Solidity: function tradeFee4dec() view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolSession) TradeFee4dec() (*big.Int, error) {
	return _DefiLiquidityPool.Contract.TradeFee4dec(&_DefiLiquidityPool.CallOpts)
}

// TradeFee4dec is a free data retrieval call binding the contract method 0xde61bcda.
//
// Solidity: function tradeFee4dec() view returns(uint256)
func (_DefiLiquidityPool *DefiLiquidityPoolCallerSession) TradeFee4dec() (*big.Int, error) {
	return _DefiLiquidityPool.Contract.TradeFee4dec(&_DefiLiquidityPool.CallOpts)
}

// Borrow is a paid mutator transaction binding the contract method 0x4b8a3529.
//
// Solidity: function borrow(address _token, uint256 _amount) returns()
func (_DefiLiquidityPool *DefiLiquidityPoolTransactor) Borrow(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiLiquidityPool.contract.Transact(opts, "borrow", _token, _amount)
}

// Borrow is a paid mutator transaction binding the contract method 0x4b8a3529.
//
// Solidity: function borrow(address _token, uint256 _amount) returns()
func (_DefiLiquidityPool *DefiLiquidityPoolSession) Borrow(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiLiquidityPool.Contract.Borrow(&_DefiLiquidityPool.TransactOpts, _token, _amount)
}

// Borrow is a paid mutator transaction binding the contract method 0x4b8a3529.
//
// Solidity: function borrow(address _token, uint256 _amount) returns()
func (_DefiLiquidityPool *DefiLiquidityPoolTransactorSession) Borrow(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiLiquidityPool.Contract.Borrow(&_DefiLiquidityPool.TransactOpts, _token, _amount)
}

// Buy is a paid mutator transaction binding the contract method 0xcce7ec13.
//
// Solidity: function buy(address _token, uint256 _amount) returns()
func (_DefiLiquidityPool *DefiLiquidityPoolTransactor) Buy(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiLiquidityPool.contract.Transact(opts, "buy", _token, _amount)
}

// Buy is a paid mutator transaction binding the contract method 0xcce7ec13.
//
// Solidity: function buy(address _token, uint256 _amount) returns()
func (_DefiLiquidityPool *DefiLiquidityPoolSession) Buy(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiLiquidityPool.Contract.Buy(&_DefiLiquidityPool.TransactOpts, _token, _amount)
}

// Buy is a paid mutator transaction binding the contract method 0xcce7ec13.
//
// Solidity: function buy(address _token, uint256 _amount) returns()
func (_DefiLiquidityPool *DefiLiquidityPoolTransactorSession) Buy(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiLiquidityPool.Contract.Buy(&_DefiLiquidityPool.TransactOpts, _token, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _token, uint256 _amount) payable returns()
func (_DefiLiquidityPool *DefiLiquidityPoolTransactor) Deposit(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiLiquidityPool.contract.Transact(opts, "deposit", _token, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _token, uint256 _amount) payable returns()
func (_DefiLiquidityPool *DefiLiquidityPoolSession) Deposit(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiLiquidityPool.Contract.Deposit(&_DefiLiquidityPool.TransactOpts, _token, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _token, uint256 _amount) payable returns()
func (_DefiLiquidityPool *DefiLiquidityPoolTransactorSession) Deposit(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiLiquidityPool.Contract.Deposit(&_DefiLiquidityPool.TransactOpts, _token, _amount)
}

// Liquidate is a paid mutator transaction binding the contract method 0x2f865568.
//
// Solidity: function liquidate(address _owner) returns()
func (_DefiLiquidityPool *DefiLiquidityPoolTransactor) Liquidate(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _DefiLiquidityPool.contract.Transact(opts, "liquidate", _owner)
}

// Liquidate is a paid mutator transaction binding the contract method 0x2f865568.
//
// Solidity: function liquidate(address _owner) returns()
func (_DefiLiquidityPool *DefiLiquidityPoolSession) Liquidate(_owner common.Address) (*types.Transaction, error) {
	return _DefiLiquidityPool.Contract.Liquidate(&_DefiLiquidityPool.TransactOpts, _owner)
}

// Liquidate is a paid mutator transaction binding the contract method 0x2f865568.
//
// Solidity: function liquidate(address _owner) returns()
func (_DefiLiquidityPool *DefiLiquidityPoolTransactorSession) Liquidate(_owner common.Address) (*types.Transaction, error) {
	return _DefiLiquidityPool.Contract.Liquidate(&_DefiLiquidityPool.TransactOpts, _owner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DefiLiquidityPool *DefiLiquidityPoolTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefiLiquidityPool.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DefiLiquidityPool *DefiLiquidityPoolSession) RenounceOwnership() (*types.Transaction, error) {
	return _DefiLiquidityPool.Contract.RenounceOwnership(&_DefiLiquidityPool.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DefiLiquidityPool *DefiLiquidityPoolTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DefiLiquidityPool.Contract.RenounceOwnership(&_DefiLiquidityPool.TransactOpts)
}

// Repay is a paid mutator transaction binding the contract method 0x22867d78.
//
// Solidity: function repay(address _token, uint256 _amount) returns()
func (_DefiLiquidityPool *DefiLiquidityPoolTransactor) Repay(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiLiquidityPool.contract.Transact(opts, "repay", _token, _amount)
}

// Repay is a paid mutator transaction binding the contract method 0x22867d78.
//
// Solidity: function repay(address _token, uint256 _amount) returns()
func (_DefiLiquidityPool *DefiLiquidityPoolSession) Repay(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiLiquidityPool.Contract.Repay(&_DefiLiquidityPool.TransactOpts, _token, _amount)
}

// Repay is a paid mutator transaction binding the contract method 0x22867d78.
//
// Solidity: function repay(address _token, uint256 _amount) returns()
func (_DefiLiquidityPool *DefiLiquidityPoolTransactorSession) Repay(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiLiquidityPool.Contract.Repay(&_DefiLiquidityPool.TransactOpts, _token, _amount)
}

// Sell is a paid mutator transaction binding the contract method 0x6c197ff5.
//
// Solidity: function sell(address _token, uint256 _amount) returns()
func (_DefiLiquidityPool *DefiLiquidityPoolTransactor) Sell(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiLiquidityPool.contract.Transact(opts, "sell", _token, _amount)
}

// Sell is a paid mutator transaction binding the contract method 0x6c197ff5.
//
// Solidity: function sell(address _token, uint256 _amount) returns()
func (_DefiLiquidityPool *DefiLiquidityPoolSession) Sell(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiLiquidityPool.Contract.Sell(&_DefiLiquidityPool.TransactOpts, _token, _amount)
}

// Sell is a paid mutator transaction binding the contract method 0x6c197ff5.
//
// Solidity: function sell(address _token, uint256 _amount) returns()
func (_DefiLiquidityPool *DefiLiquidityPoolTransactorSession) Sell(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiLiquidityPool.Contract.Sell(&_DefiLiquidityPool.TransactOpts, _token, _amount)
}

// SetPriceOracleReferenceAggregate is a paid mutator transaction binding the contract method 0xa441d055.
//
// Solidity: function setPriceOracleReferenceAggregate(address oracle) returns()
func (_DefiLiquidityPool *DefiLiquidityPoolTransactor) SetPriceOracleReferenceAggregate(opts *bind.TransactOpts, oracle common.Address) (*types.Transaction, error) {
	return _DefiLiquidityPool.contract.Transact(opts, "setPriceOracleReferenceAggregate", oracle)
}

// SetPriceOracleReferenceAggregate is a paid mutator transaction binding the contract method 0xa441d055.
//
// Solidity: function setPriceOracleReferenceAggregate(address oracle) returns()
func (_DefiLiquidityPool *DefiLiquidityPoolSession) SetPriceOracleReferenceAggregate(oracle common.Address) (*types.Transaction, error) {
	return _DefiLiquidityPool.Contract.SetPriceOracleReferenceAggregate(&_DefiLiquidityPool.TransactOpts, oracle)
}

// SetPriceOracleReferenceAggregate is a paid mutator transaction binding the contract method 0xa441d055.
//
// Solidity: function setPriceOracleReferenceAggregate(address oracle) returns()
func (_DefiLiquidityPool *DefiLiquidityPoolTransactorSession) SetPriceOracleReferenceAggregate(oracle common.Address) (*types.Transaction, error) {
	return _DefiLiquidityPool.Contract.SetPriceOracleReferenceAggregate(&_DefiLiquidityPool.TransactOpts, oracle)
}

// Trade is a paid mutator transaction binding the contract method 0x5253baae.
//
// Solidity: function trade(address _fromToken, address _toToken, uint256 _amount) returns()
func (_DefiLiquidityPool *DefiLiquidityPoolTransactor) Trade(opts *bind.TransactOpts, _fromToken common.Address, _toToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiLiquidityPool.contract.Transact(opts, "trade", _fromToken, _toToken, _amount)
}

// Trade is a paid mutator transaction binding the contract method 0x5253baae.
//
// Solidity: function trade(address _fromToken, address _toToken, uint256 _amount) returns()
func (_DefiLiquidityPool *DefiLiquidityPoolSession) Trade(_fromToken common.Address, _toToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiLiquidityPool.Contract.Trade(&_DefiLiquidityPool.TransactOpts, _fromToken, _toToken, _amount)
}

// Trade is a paid mutator transaction binding the contract method 0x5253baae.
//
// Solidity: function trade(address _fromToken, address _toToken, uint256 _amount) returns()
func (_DefiLiquidityPool *DefiLiquidityPoolTransactorSession) Trade(_fromToken common.Address, _toToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiLiquidityPool.Contract.Trade(&_DefiLiquidityPool.TransactOpts, _fromToken, _toToken, _amount)
}

// TransferBalance is a paid mutator transaction binding the contract method 0x63b6b31f.
//
// Solidity: function transferBalance(address to) returns()
func (_DefiLiquidityPool *DefiLiquidityPoolTransactor) TransferBalance(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _DefiLiquidityPool.contract.Transact(opts, "transferBalance", to)
}

// TransferBalance is a paid mutator transaction binding the contract method 0x63b6b31f.
//
// Solidity: function transferBalance(address to) returns()
func (_DefiLiquidityPool *DefiLiquidityPoolSession) TransferBalance(to common.Address) (*types.Transaction, error) {
	return _DefiLiquidityPool.Contract.TransferBalance(&_DefiLiquidityPool.TransactOpts, to)
}

// TransferBalance is a paid mutator transaction binding the contract method 0x63b6b31f.
//
// Solidity: function transferBalance(address to) returns()
func (_DefiLiquidityPool *DefiLiquidityPoolTransactorSession) TransferBalance(to common.Address) (*types.Transaction, error) {
	return _DefiLiquidityPool.Contract.TransferBalance(&_DefiLiquidityPool.TransactOpts, to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DefiLiquidityPool *DefiLiquidityPoolTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DefiLiquidityPool.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DefiLiquidityPool *DefiLiquidityPoolSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DefiLiquidityPool.Contract.TransferOwnership(&_DefiLiquidityPool.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DefiLiquidityPool *DefiLiquidityPoolTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DefiLiquidityPool.Contract.TransferOwnership(&_DefiLiquidityPool.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _token, uint256 _amount) returns()
func (_DefiLiquidityPool *DefiLiquidityPoolTransactor) Withdraw(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiLiquidityPool.contract.Transact(opts, "withdraw", _token, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _token, uint256 _amount) returns()
func (_DefiLiquidityPool *DefiLiquidityPoolSession) Withdraw(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiLiquidityPool.Contract.Withdraw(&_DefiLiquidityPool.TransactOpts, _token, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _token, uint256 _amount) returns()
func (_DefiLiquidityPool *DefiLiquidityPoolTransactorSession) Withdraw(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiLiquidityPool.Contract.Withdraw(&_DefiLiquidityPool.TransactOpts, _token, _amount)
}

// DefiLiquidityPoolBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the DefiLiquidityPool contract.
type DefiLiquidityPoolBorrowIterator struct {
	Event *DefiLiquidityPoolBorrow // Event containing the contract specifics and raw log

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
func (it *DefiLiquidityPoolBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefiLiquidityPoolBorrow)
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
		it.Event = new(DefiLiquidityPoolBorrow)
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
func (it *DefiLiquidityPoolBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefiLiquidityPoolBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefiLiquidityPoolBorrow represents a Borrow event raised by the DefiLiquidityPool contract.
type DefiLiquidityPoolBorrow struct {
	Token     common.Address
	User      common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0xc1561b330e73faa7d5d1ac03c968d8f359b0191ccdb9cc002cf7d8eb6ae038cb.
//
// Solidity: event Borrow(address indexed token, address indexed user, uint256 amount, uint256 timestamp)
func (_DefiLiquidityPool *DefiLiquidityPoolFilterer) FilterBorrow(opts *bind.FilterOpts, token []common.Address, user []common.Address) (*DefiLiquidityPoolBorrowIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DefiLiquidityPool.contract.FilterLogs(opts, "Borrow", tokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return &DefiLiquidityPoolBorrowIterator{contract: _DefiLiquidityPool.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0xc1561b330e73faa7d5d1ac03c968d8f359b0191ccdb9cc002cf7d8eb6ae038cb.
//
// Solidity: event Borrow(address indexed token, address indexed user, uint256 amount, uint256 timestamp)
func (_DefiLiquidityPool *DefiLiquidityPoolFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *DefiLiquidityPoolBorrow, token []common.Address, user []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DefiLiquidityPool.contract.WatchLogs(opts, "Borrow", tokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefiLiquidityPoolBorrow)
				if err := _DefiLiquidityPool.contract.UnpackLog(event, "Borrow", log); err != nil {
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

// ParseBorrow is a log parse operation binding the contract event 0xc1561b330e73faa7d5d1ac03c968d8f359b0191ccdb9cc002cf7d8eb6ae038cb.
//
// Solidity: event Borrow(address indexed token, address indexed user, uint256 amount, uint256 timestamp)
func (_DefiLiquidityPool *DefiLiquidityPoolFilterer) ParseBorrow(log types.Log) (*DefiLiquidityPoolBorrow, error) {
	event := new(DefiLiquidityPoolBorrow)
	if err := _DefiLiquidityPool.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DefiLiquidityPoolBuyIterator is returned from FilterBuy and is used to iterate over the raw logs and unpacked data for Buy events raised by the DefiLiquidityPool contract.
type DefiLiquidityPoolBuyIterator struct {
	Event *DefiLiquidityPoolBuy // Event containing the contract specifics and raw log

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
func (it *DefiLiquidityPoolBuyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefiLiquidityPoolBuy)
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
		it.Event = new(DefiLiquidityPoolBuy)
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
func (it *DefiLiquidityPoolBuyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefiLiquidityPoolBuyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefiLiquidityPoolBuy represents a Buy event raised by the DefiLiquidityPool contract.
type DefiLiquidityPoolBuy struct {
	Token        common.Address
	User         common.Address
	Amount       *big.Int
	ExchangeRate *big.Int
	Timestamp    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBuy is a free log retrieval operation binding the contract event 0x00f93dbdb72854b6b6fb35433086556f2635fc83c37080c667496fecfa650fb4.
//
// Solidity: event Buy(address indexed token, address indexed user, uint256 amount, uint256 exchangeRate, uint256 timestamp)
func (_DefiLiquidityPool *DefiLiquidityPoolFilterer) FilterBuy(opts *bind.FilterOpts, token []common.Address, user []common.Address) (*DefiLiquidityPoolBuyIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DefiLiquidityPool.contract.FilterLogs(opts, "Buy", tokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return &DefiLiquidityPoolBuyIterator{contract: _DefiLiquidityPool.contract, event: "Buy", logs: logs, sub: sub}, nil
}

// WatchBuy is a free log subscription operation binding the contract event 0x00f93dbdb72854b6b6fb35433086556f2635fc83c37080c667496fecfa650fb4.
//
// Solidity: event Buy(address indexed token, address indexed user, uint256 amount, uint256 exchangeRate, uint256 timestamp)
func (_DefiLiquidityPool *DefiLiquidityPoolFilterer) WatchBuy(opts *bind.WatchOpts, sink chan<- *DefiLiquidityPoolBuy, token []common.Address, user []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DefiLiquidityPool.contract.WatchLogs(opts, "Buy", tokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefiLiquidityPoolBuy)
				if err := _DefiLiquidityPool.contract.UnpackLog(event, "Buy", log); err != nil {
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

// ParseBuy is a log parse operation binding the contract event 0x00f93dbdb72854b6b6fb35433086556f2635fc83c37080c667496fecfa650fb4.
//
// Solidity: event Buy(address indexed token, address indexed user, uint256 amount, uint256 exchangeRate, uint256 timestamp)
func (_DefiLiquidityPool *DefiLiquidityPoolFilterer) ParseBuy(log types.Log) (*DefiLiquidityPoolBuy, error) {
	event := new(DefiLiquidityPoolBuy)
	if err := _DefiLiquidityPool.contract.UnpackLog(event, "Buy", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DefiLiquidityPoolDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the DefiLiquidityPool contract.
type DefiLiquidityPoolDepositIterator struct {
	Event *DefiLiquidityPoolDeposit // Event containing the contract specifics and raw log

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
func (it *DefiLiquidityPoolDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefiLiquidityPoolDeposit)
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
		it.Event = new(DefiLiquidityPoolDeposit)
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
func (it *DefiLiquidityPoolDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefiLiquidityPoolDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefiLiquidityPoolDeposit represents a Deposit event raised by the DefiLiquidityPool contract.
type DefiLiquidityPoolDeposit struct {
	Token     common.Address
	User      common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed token, address indexed user, uint256 amount, uint256 timestamp)
func (_DefiLiquidityPool *DefiLiquidityPoolFilterer) FilterDeposit(opts *bind.FilterOpts, token []common.Address, user []common.Address) (*DefiLiquidityPoolDepositIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DefiLiquidityPool.contract.FilterLogs(opts, "Deposit", tokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return &DefiLiquidityPoolDepositIterator{contract: _DefiLiquidityPool.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed token, address indexed user, uint256 amount, uint256 timestamp)
func (_DefiLiquidityPool *DefiLiquidityPoolFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *DefiLiquidityPoolDeposit, token []common.Address, user []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DefiLiquidityPool.contract.WatchLogs(opts, "Deposit", tokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefiLiquidityPoolDeposit)
				if err := _DefiLiquidityPool.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed token, address indexed user, uint256 amount, uint256 timestamp)
func (_DefiLiquidityPool *DefiLiquidityPoolFilterer) ParseDeposit(log types.Log) (*DefiLiquidityPoolDeposit, error) {
	event := new(DefiLiquidityPoolDeposit)
	if err := _DefiLiquidityPool.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DefiLiquidityPoolOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DefiLiquidityPool contract.
type DefiLiquidityPoolOwnershipTransferredIterator struct {
	Event *DefiLiquidityPoolOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DefiLiquidityPoolOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefiLiquidityPoolOwnershipTransferred)
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
		it.Event = new(DefiLiquidityPoolOwnershipTransferred)
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
func (it *DefiLiquidityPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefiLiquidityPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefiLiquidityPoolOwnershipTransferred represents a OwnershipTransferred event raised by the DefiLiquidityPool contract.
type DefiLiquidityPoolOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DefiLiquidityPool *DefiLiquidityPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DefiLiquidityPoolOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DefiLiquidityPool.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DefiLiquidityPoolOwnershipTransferredIterator{contract: _DefiLiquidityPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DefiLiquidityPool *DefiLiquidityPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DefiLiquidityPoolOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DefiLiquidityPool.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefiLiquidityPoolOwnershipTransferred)
				if err := _DefiLiquidityPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_DefiLiquidityPool *DefiLiquidityPoolFilterer) ParseOwnershipTransferred(log types.Log) (*DefiLiquidityPoolOwnershipTransferred, error) {
	event := new(DefiLiquidityPoolOwnershipTransferred)
	if err := _DefiLiquidityPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DefiLiquidityPoolPriceOracleChangedIterator is returned from FilterPriceOracleChanged and is used to iterate over the raw logs and unpacked data for PriceOracleChanged events raised by the DefiLiquidityPool contract.
type DefiLiquidityPoolPriceOracleChangedIterator struct {
	Event *DefiLiquidityPoolPriceOracleChanged // Event containing the contract specifics and raw log

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
func (it *DefiLiquidityPoolPriceOracleChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefiLiquidityPoolPriceOracleChanged)
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
		it.Event = new(DefiLiquidityPoolPriceOracleChanged)
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
func (it *DefiLiquidityPoolPriceOracleChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefiLiquidityPoolPriceOracleChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefiLiquidityPoolPriceOracleChanged represents a PriceOracleChanged event raised by the DefiLiquidityPool contract.
type DefiLiquidityPoolPriceOracleChanged struct {
	Oracle    common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPriceOracleChanged is a free log retrieval operation binding the contract event 0x09213e56b42d8171eecfb5137a54ea062b7386085d7606c0ec108a84ef341b95.
//
// Solidity: event PriceOracleChanged(address oracle, uint256 timestamp)
func (_DefiLiquidityPool *DefiLiquidityPoolFilterer) FilterPriceOracleChanged(opts *bind.FilterOpts) (*DefiLiquidityPoolPriceOracleChangedIterator, error) {

	logs, sub, err := _DefiLiquidityPool.contract.FilterLogs(opts, "PriceOracleChanged")
	if err != nil {
		return nil, err
	}
	return &DefiLiquidityPoolPriceOracleChangedIterator{contract: _DefiLiquidityPool.contract, event: "PriceOracleChanged", logs: logs, sub: sub}, nil
}

// WatchPriceOracleChanged is a free log subscription operation binding the contract event 0x09213e56b42d8171eecfb5137a54ea062b7386085d7606c0ec108a84ef341b95.
//
// Solidity: event PriceOracleChanged(address oracle, uint256 timestamp)
func (_DefiLiquidityPool *DefiLiquidityPoolFilterer) WatchPriceOracleChanged(opts *bind.WatchOpts, sink chan<- *DefiLiquidityPoolPriceOracleChanged) (event.Subscription, error) {

	logs, sub, err := _DefiLiquidityPool.contract.WatchLogs(opts, "PriceOracleChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefiLiquidityPoolPriceOracleChanged)
				if err := _DefiLiquidityPool.contract.UnpackLog(event, "PriceOracleChanged", log); err != nil {
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

// ParsePriceOracleChanged is a log parse operation binding the contract event 0x09213e56b42d8171eecfb5137a54ea062b7386085d7606c0ec108a84ef341b95.
//
// Solidity: event PriceOracleChanged(address oracle, uint256 timestamp)
func (_DefiLiquidityPool *DefiLiquidityPoolFilterer) ParsePriceOracleChanged(log types.Log) (*DefiLiquidityPoolPriceOracleChanged, error) {
	event := new(DefiLiquidityPoolPriceOracleChanged)
	if err := _DefiLiquidityPool.contract.UnpackLog(event, "PriceOracleChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DefiLiquidityPoolRepayIterator is returned from FilterRepay and is used to iterate over the raw logs and unpacked data for Repay events raised by the DefiLiquidityPool contract.
type DefiLiquidityPoolRepayIterator struct {
	Event *DefiLiquidityPoolRepay // Event containing the contract specifics and raw log

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
func (it *DefiLiquidityPoolRepayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefiLiquidityPoolRepay)
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
		it.Event = new(DefiLiquidityPoolRepay)
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
func (it *DefiLiquidityPoolRepayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefiLiquidityPoolRepayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefiLiquidityPoolRepay represents a Repay event raised by the DefiLiquidityPool contract.
type DefiLiquidityPoolRepay struct {
	Token     common.Address
	User      common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRepay is a free log retrieval operation binding the contract event 0xe4a1ae657f49cb1fb1c7d3a94ae6093565c4c8c0e03de488f79c377c3c3a24e0.
//
// Solidity: event Repay(address indexed token, address indexed user, uint256 amount, uint256 timestamp)
func (_DefiLiquidityPool *DefiLiquidityPoolFilterer) FilterRepay(opts *bind.FilterOpts, token []common.Address, user []common.Address) (*DefiLiquidityPoolRepayIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DefiLiquidityPool.contract.FilterLogs(opts, "Repay", tokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return &DefiLiquidityPoolRepayIterator{contract: _DefiLiquidityPool.contract, event: "Repay", logs: logs, sub: sub}, nil
}

// WatchRepay is a free log subscription operation binding the contract event 0xe4a1ae657f49cb1fb1c7d3a94ae6093565c4c8c0e03de488f79c377c3c3a24e0.
//
// Solidity: event Repay(address indexed token, address indexed user, uint256 amount, uint256 timestamp)
func (_DefiLiquidityPool *DefiLiquidityPoolFilterer) WatchRepay(opts *bind.WatchOpts, sink chan<- *DefiLiquidityPoolRepay, token []common.Address, user []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DefiLiquidityPool.contract.WatchLogs(opts, "Repay", tokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefiLiquidityPoolRepay)
				if err := _DefiLiquidityPool.contract.UnpackLog(event, "Repay", log); err != nil {
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

// ParseRepay is a log parse operation binding the contract event 0xe4a1ae657f49cb1fb1c7d3a94ae6093565c4c8c0e03de488f79c377c3c3a24e0.
//
// Solidity: event Repay(address indexed token, address indexed user, uint256 amount, uint256 timestamp)
func (_DefiLiquidityPool *DefiLiquidityPoolFilterer) ParseRepay(log types.Log) (*DefiLiquidityPoolRepay, error) {
	event := new(DefiLiquidityPoolRepay)
	if err := _DefiLiquidityPool.contract.UnpackLog(event, "Repay", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DefiLiquidityPoolSellIterator is returned from FilterSell and is used to iterate over the raw logs and unpacked data for Sell events raised by the DefiLiquidityPool contract.
type DefiLiquidityPoolSellIterator struct {
	Event *DefiLiquidityPoolSell // Event containing the contract specifics and raw log

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
func (it *DefiLiquidityPoolSellIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefiLiquidityPoolSell)
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
		it.Event = new(DefiLiquidityPoolSell)
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
func (it *DefiLiquidityPoolSellIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefiLiquidityPoolSellIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefiLiquidityPoolSell represents a Sell event raised by the DefiLiquidityPool contract.
type DefiLiquidityPoolSell struct {
	Token        common.Address
	User         common.Address
	Amount       *big.Int
	ExchangeRate *big.Int
	Timestamp    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSell is a free log retrieval operation binding the contract event 0x01fbb57444511e3de5b26ac09ad6bec45c3f9a1e59dd4a0f2b13a240d18476ce.
//
// Solidity: event Sell(address indexed token, address indexed user, uint256 amount, uint256 exchangeRate, uint256 timestamp)
func (_DefiLiquidityPool *DefiLiquidityPoolFilterer) FilterSell(opts *bind.FilterOpts, token []common.Address, user []common.Address) (*DefiLiquidityPoolSellIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DefiLiquidityPool.contract.FilterLogs(opts, "Sell", tokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return &DefiLiquidityPoolSellIterator{contract: _DefiLiquidityPool.contract, event: "Sell", logs: logs, sub: sub}, nil
}

// WatchSell is a free log subscription operation binding the contract event 0x01fbb57444511e3de5b26ac09ad6bec45c3f9a1e59dd4a0f2b13a240d18476ce.
//
// Solidity: event Sell(address indexed token, address indexed user, uint256 amount, uint256 exchangeRate, uint256 timestamp)
func (_DefiLiquidityPool *DefiLiquidityPoolFilterer) WatchSell(opts *bind.WatchOpts, sink chan<- *DefiLiquidityPoolSell, token []common.Address, user []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DefiLiquidityPool.contract.WatchLogs(opts, "Sell", tokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefiLiquidityPoolSell)
				if err := _DefiLiquidityPool.contract.UnpackLog(event, "Sell", log); err != nil {
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

// ParseSell is a log parse operation binding the contract event 0x01fbb57444511e3de5b26ac09ad6bec45c3f9a1e59dd4a0f2b13a240d18476ce.
//
// Solidity: event Sell(address indexed token, address indexed user, uint256 amount, uint256 exchangeRate, uint256 timestamp)
func (_DefiLiquidityPool *DefiLiquidityPoolFilterer) ParseSell(log types.Log) (*DefiLiquidityPoolSell, error) {
	event := new(DefiLiquidityPoolSell)
	if err := _DefiLiquidityPool.contract.UnpackLog(event, "Sell", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DefiLiquidityPoolWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the DefiLiquidityPool contract.
type DefiLiquidityPoolWithdrawIterator struct {
	Event *DefiLiquidityPoolWithdraw // Event containing the contract specifics and raw log

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
func (it *DefiLiquidityPoolWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefiLiquidityPoolWithdraw)
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
		it.Event = new(DefiLiquidityPoolWithdraw)
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
func (it *DefiLiquidityPoolWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefiLiquidityPoolWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefiLiquidityPoolWithdraw represents a Withdraw event raised by the DefiLiquidityPool contract.
type DefiLiquidityPoolWithdraw struct {
	Token     common.Address
	User      common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf341246adaac6f497bc2a656f546ab9e182111d630394f0c57c710a59a2cb567.
//
// Solidity: event Withdraw(address indexed token, address indexed user, uint256 amount, uint256 timestamp)
func (_DefiLiquidityPool *DefiLiquidityPoolFilterer) FilterWithdraw(opts *bind.FilterOpts, token []common.Address, user []common.Address) (*DefiLiquidityPoolWithdrawIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DefiLiquidityPool.contract.FilterLogs(opts, "Withdraw", tokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return &DefiLiquidityPoolWithdrawIterator{contract: _DefiLiquidityPool.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf341246adaac6f497bc2a656f546ab9e182111d630394f0c57c710a59a2cb567.
//
// Solidity: event Withdraw(address indexed token, address indexed user, uint256 amount, uint256 timestamp)
func (_DefiLiquidityPool *DefiLiquidityPoolFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *DefiLiquidityPoolWithdraw, token []common.Address, user []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DefiLiquidityPool.contract.WatchLogs(opts, "Withdraw", tokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefiLiquidityPoolWithdraw)
				if err := _DefiLiquidityPool.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xf341246adaac6f497bc2a656f546ab9e182111d630394f0c57c710a59a2cb567.
//
// Solidity: event Withdraw(address indexed token, address indexed user, uint256 amount, uint256 timestamp)
func (_DefiLiquidityPool *DefiLiquidityPoolFilterer) ParseWithdraw(log types.Log) (*DefiLiquidityPoolWithdraw, error) {
	event := new(DefiLiquidityPoolWithdraw)
	if err := _DefiLiquidityPool.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}
