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

// DefiFMintAddressProviderMetaData contains all meta data concerning the DefiFMintAddressProvider contract.
var DefiFMintAddressProviderMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"CollateralPoolChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"DebtPoolChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"MinterChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"PriceOracleChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"RewardDistributionChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"RewardTokenChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"TokenRegistryChanged\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPriceOracleProxy\",\"outputs\":[{\"internalType\":\"contractIPriceOracleProxy\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setPriceOracleProxy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTokenRegistry\",\"outputs\":[{\"internalType\":\"contractIFantomMintTokenRegistry\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setTokenRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRewardDistribution\",\"outputs\":[{\"internalType\":\"contractIFantomMintRewardManager\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setRewardDistribution\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRewardToken\",\"outputs\":[{\"internalType\":\"contractERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setRewardToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFantomMint\",\"outputs\":[{\"internalType\":\"contractIFantomMintBalanceGuard\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setFantomMint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCollateralPool\",\"outputs\":[{\"internalType\":\"contractIFantomDeFiTokenStorage\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setCollateralPool\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDebtPool\",\"outputs\":[{\"internalType\":\"contractIFantomDeFiTokenStorage\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setDebtPool\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DefiFMintAddressProviderABI is the input ABI used to generate the binding from.
// Deprecated: Use DefiFMintAddressProviderMetaData.ABI instead.
var DefiFMintAddressProviderABI = DefiFMintAddressProviderMetaData.ABI

// DefiFMintAddressProvider is an auto generated Go binding around an Ethereum contract.
type DefiFMintAddressProvider struct {
	DefiFMintAddressProviderCaller     // Read-only binding to the contract
	DefiFMintAddressProviderTransactor // Write-only binding to the contract
	DefiFMintAddressProviderFilterer   // Log filterer for contract events
}

// DefiFMintAddressProviderCaller is an auto generated read-only Go binding around an Ethereum contract.
type DefiFMintAddressProviderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefiFMintAddressProviderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DefiFMintAddressProviderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefiFMintAddressProviderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DefiFMintAddressProviderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefiFMintAddressProviderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DefiFMintAddressProviderSession struct {
	Contract     *DefiFMintAddressProvider // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// DefiFMintAddressProviderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DefiFMintAddressProviderCallerSession struct {
	Contract *DefiFMintAddressProviderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// DefiFMintAddressProviderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DefiFMintAddressProviderTransactorSession struct {
	Contract     *DefiFMintAddressProviderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// DefiFMintAddressProviderRaw is an auto generated low-level Go binding around an Ethereum contract.
type DefiFMintAddressProviderRaw struct {
	Contract *DefiFMintAddressProvider // Generic contract binding to access the raw methods on
}

// DefiFMintAddressProviderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DefiFMintAddressProviderCallerRaw struct {
	Contract *DefiFMintAddressProviderCaller // Generic read-only contract binding to access the raw methods on
}

// DefiFMintAddressProviderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DefiFMintAddressProviderTransactorRaw struct {
	Contract *DefiFMintAddressProviderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDefiFMintAddressProvider creates a new instance of DefiFMintAddressProvider, bound to a specific deployed contract.
func NewDefiFMintAddressProvider(address common.Address, backend bind.ContractBackend) (*DefiFMintAddressProvider, error) {
	contract, err := bindDefiFMintAddressProvider(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DefiFMintAddressProvider{DefiFMintAddressProviderCaller: DefiFMintAddressProviderCaller{contract: contract}, DefiFMintAddressProviderTransactor: DefiFMintAddressProviderTransactor{contract: contract}, DefiFMintAddressProviderFilterer: DefiFMintAddressProviderFilterer{contract: contract}}, nil
}

// NewDefiFMintAddressProviderCaller creates a new read-only instance of DefiFMintAddressProvider, bound to a specific deployed contract.
func NewDefiFMintAddressProviderCaller(address common.Address, caller bind.ContractCaller) (*DefiFMintAddressProviderCaller, error) {
	contract, err := bindDefiFMintAddressProvider(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DefiFMintAddressProviderCaller{contract: contract}, nil
}

// NewDefiFMintAddressProviderTransactor creates a new write-only instance of DefiFMintAddressProvider, bound to a specific deployed contract.
func NewDefiFMintAddressProviderTransactor(address common.Address, transactor bind.ContractTransactor) (*DefiFMintAddressProviderTransactor, error) {
	contract, err := bindDefiFMintAddressProvider(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DefiFMintAddressProviderTransactor{contract: contract}, nil
}

// NewDefiFMintAddressProviderFilterer creates a new log filterer instance of DefiFMintAddressProvider, bound to a specific deployed contract.
func NewDefiFMintAddressProviderFilterer(address common.Address, filterer bind.ContractFilterer) (*DefiFMintAddressProviderFilterer, error) {
	contract, err := bindDefiFMintAddressProvider(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DefiFMintAddressProviderFilterer{contract: contract}, nil
}

// bindDefiFMintAddressProvider binds a generic wrapper to an already deployed contract.
func bindDefiFMintAddressProvider(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DefiFMintAddressProviderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DefiFMintAddressProvider *DefiFMintAddressProviderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DefiFMintAddressProvider.Contract.DefiFMintAddressProviderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DefiFMintAddressProvider *DefiFMintAddressProviderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefiFMintAddressProvider.Contract.DefiFMintAddressProviderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DefiFMintAddressProvider *DefiFMintAddressProviderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DefiFMintAddressProvider.Contract.DefiFMintAddressProviderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DefiFMintAddressProvider *DefiFMintAddressProviderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DefiFMintAddressProvider.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DefiFMintAddressProvider *DefiFMintAddressProviderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefiFMintAddressProvider.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DefiFMintAddressProvider *DefiFMintAddressProviderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DefiFMintAddressProvider.Contract.contract.Transact(opts, method, params...)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 _id) view returns(address)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderCaller) GetAddress(opts *bind.CallOpts, _id [32]byte) (common.Address, error) {
	var out []interface{}
	err := _DefiFMintAddressProvider.contract.Call(opts, &out, "getAddress", _id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 _id) view returns(address)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderSession) GetAddress(_id [32]byte) (common.Address, error) {
	return _DefiFMintAddressProvider.Contract.GetAddress(&_DefiFMintAddressProvider.CallOpts, _id)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 _id) view returns(address)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderCallerSession) GetAddress(_id [32]byte) (common.Address, error) {
	return _DefiFMintAddressProvider.Contract.GetAddress(&_DefiFMintAddressProvider.CallOpts, _id)
}

// GetCollateralPool is a free data retrieval call binding the contract method 0x73c9641d.
//
// Solidity: function getCollateralPool() view returns(address)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderCaller) GetCollateralPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DefiFMintAddressProvider.contract.Call(opts, &out, "getCollateralPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCollateralPool is a free data retrieval call binding the contract method 0x73c9641d.
//
// Solidity: function getCollateralPool() view returns(address)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderSession) GetCollateralPool() (common.Address, error) {
	return _DefiFMintAddressProvider.Contract.GetCollateralPool(&_DefiFMintAddressProvider.CallOpts)
}

// GetCollateralPool is a free data retrieval call binding the contract method 0x73c9641d.
//
// Solidity: function getCollateralPool() view returns(address)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderCallerSession) GetCollateralPool() (common.Address, error) {
	return _DefiFMintAddressProvider.Contract.GetCollateralPool(&_DefiFMintAddressProvider.CallOpts)
}

// GetDebtPool is a free data retrieval call binding the contract method 0x03ec357f.
//
// Solidity: function getDebtPool() view returns(address)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderCaller) GetDebtPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DefiFMintAddressProvider.contract.Call(opts, &out, "getDebtPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDebtPool is a free data retrieval call binding the contract method 0x03ec357f.
//
// Solidity: function getDebtPool() view returns(address)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderSession) GetDebtPool() (common.Address, error) {
	return _DefiFMintAddressProvider.Contract.GetDebtPool(&_DefiFMintAddressProvider.CallOpts)
}

// GetDebtPool is a free data retrieval call binding the contract method 0x03ec357f.
//
// Solidity: function getDebtPool() view returns(address)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderCallerSession) GetDebtPool() (common.Address, error) {
	return _DefiFMintAddressProvider.Contract.GetDebtPool(&_DefiFMintAddressProvider.CallOpts)
}

// GetFantomMint is a free data retrieval call binding the contract method 0x44969711.
//
// Solidity: function getFantomMint() view returns(address)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderCaller) GetFantomMint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DefiFMintAddressProvider.contract.Call(opts, &out, "getFantomMint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFantomMint is a free data retrieval call binding the contract method 0x44969711.
//
// Solidity: function getFantomMint() view returns(address)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderSession) GetFantomMint() (common.Address, error) {
	return _DefiFMintAddressProvider.Contract.GetFantomMint(&_DefiFMintAddressProvider.CallOpts)
}

// GetFantomMint is a free data retrieval call binding the contract method 0x44969711.
//
// Solidity: function getFantomMint() view returns(address)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderCallerSession) GetFantomMint() (common.Address, error) {
	return _DefiFMintAddressProvider.Contract.GetFantomMint(&_DefiFMintAddressProvider.CallOpts)
}

// GetPriceOracleProxy is a free data retrieval call binding the contract method 0x045bb7f8.
//
// Solidity: function getPriceOracleProxy() view returns(address)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderCaller) GetPriceOracleProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DefiFMintAddressProvider.contract.Call(opts, &out, "getPriceOracleProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPriceOracleProxy is a free data retrieval call binding the contract method 0x045bb7f8.
//
// Solidity: function getPriceOracleProxy() view returns(address)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderSession) GetPriceOracleProxy() (common.Address, error) {
	return _DefiFMintAddressProvider.Contract.GetPriceOracleProxy(&_DefiFMintAddressProvider.CallOpts)
}

// GetPriceOracleProxy is a free data retrieval call binding the contract method 0x045bb7f8.
//
// Solidity: function getPriceOracleProxy() view returns(address)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderCallerSession) GetPriceOracleProxy() (common.Address, error) {
	return _DefiFMintAddressProvider.Contract.GetPriceOracleProxy(&_DefiFMintAddressProvider.CallOpts)
}

// GetRewardDistribution is a free data retrieval call binding the contract method 0x84d9319e.
//
// Solidity: function getRewardDistribution() view returns(address)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderCaller) GetRewardDistribution(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DefiFMintAddressProvider.contract.Call(opts, &out, "getRewardDistribution")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRewardDistribution is a free data retrieval call binding the contract method 0x84d9319e.
//
// Solidity: function getRewardDistribution() view returns(address)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderSession) GetRewardDistribution() (common.Address, error) {
	return _DefiFMintAddressProvider.Contract.GetRewardDistribution(&_DefiFMintAddressProvider.CallOpts)
}

// GetRewardDistribution is a free data retrieval call binding the contract method 0x84d9319e.
//
// Solidity: function getRewardDistribution() view returns(address)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderCallerSession) GetRewardDistribution() (common.Address, error) {
	return _DefiFMintAddressProvider.Contract.GetRewardDistribution(&_DefiFMintAddressProvider.CallOpts)
}

// GetRewardToken is a free data retrieval call binding the contract method 0x69940d79.
//
// Solidity: function getRewardToken() view returns(address)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderCaller) GetRewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DefiFMintAddressProvider.contract.Call(opts, &out, "getRewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRewardToken is a free data retrieval call binding the contract method 0x69940d79.
//
// Solidity: function getRewardToken() view returns(address)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderSession) GetRewardToken() (common.Address, error) {
	return _DefiFMintAddressProvider.Contract.GetRewardToken(&_DefiFMintAddressProvider.CallOpts)
}

// GetRewardToken is a free data retrieval call binding the contract method 0x69940d79.
//
// Solidity: function getRewardToken() view returns(address)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderCallerSession) GetRewardToken() (common.Address, error) {
	return _DefiFMintAddressProvider.Contract.GetRewardToken(&_DefiFMintAddressProvider.CallOpts)
}

// GetTokenRegistry is a free data retrieval call binding the contract method 0x057838bd.
//
// Solidity: function getTokenRegistry() view returns(address)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderCaller) GetTokenRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DefiFMintAddressProvider.contract.Call(opts, &out, "getTokenRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenRegistry is a free data retrieval call binding the contract method 0x057838bd.
//
// Solidity: function getTokenRegistry() view returns(address)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderSession) GetTokenRegistry() (common.Address, error) {
	return _DefiFMintAddressProvider.Contract.GetTokenRegistry(&_DefiFMintAddressProvider.CallOpts)
}

// GetTokenRegistry is a free data retrieval call binding the contract method 0x057838bd.
//
// Solidity: function getTokenRegistry() view returns(address)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderCallerSession) GetTokenRegistry() (common.Address, error) {
	return _DefiFMintAddressProvider.Contract.GetTokenRegistry(&_DefiFMintAddressProvider.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DefiFMintAddressProvider.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderSession) IsOwner() (bool, error) {
	return _DefiFMintAddressProvider.Contract.IsOwner(&_DefiFMintAddressProvider.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderCallerSession) IsOwner() (bool, error) {
	return _DefiFMintAddressProvider.Contract.IsOwner(&_DefiFMintAddressProvider.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DefiFMintAddressProvider.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderSession) Owner() (common.Address, error) {
	return _DefiFMintAddressProvider.Contract.Owner(&_DefiFMintAddressProvider.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderCallerSession) Owner() (common.Address, error) {
	return _DefiFMintAddressProvider.Contract.Owner(&_DefiFMintAddressProvider.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner) returns()
func (_DefiFMintAddressProvider *DefiFMintAddressProviderTransactor) Initialize(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _DefiFMintAddressProvider.contract.Transact(opts, "initialize", owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner) returns()
func (_DefiFMintAddressProvider *DefiFMintAddressProviderSession) Initialize(owner common.Address) (*types.Transaction, error) {
	return _DefiFMintAddressProvider.Contract.Initialize(&_DefiFMintAddressProvider.TransactOpts, owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner) returns()
func (_DefiFMintAddressProvider *DefiFMintAddressProviderTransactorSession) Initialize(owner common.Address) (*types.Transaction, error) {
	return _DefiFMintAddressProvider.Contract.Initialize(&_DefiFMintAddressProvider.TransactOpts, owner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DefiFMintAddressProvider *DefiFMintAddressProviderTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefiFMintAddressProvider.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DefiFMintAddressProvider *DefiFMintAddressProviderSession) RenounceOwnership() (*types.Transaction, error) {
	return _DefiFMintAddressProvider.Contract.RenounceOwnership(&_DefiFMintAddressProvider.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DefiFMintAddressProvider *DefiFMintAddressProviderTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DefiFMintAddressProvider.Contract.RenounceOwnership(&_DefiFMintAddressProvider.TransactOpts)
}

// SetCollateralPool is a paid mutator transaction binding the contract method 0x1ba28878.
//
// Solidity: function setCollateralPool(address _addr) returns()
func (_DefiFMintAddressProvider *DefiFMintAddressProviderTransactor) SetCollateralPool(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _DefiFMintAddressProvider.contract.Transact(opts, "setCollateralPool", _addr)
}

// SetCollateralPool is a paid mutator transaction binding the contract method 0x1ba28878.
//
// Solidity: function setCollateralPool(address _addr) returns()
func (_DefiFMintAddressProvider *DefiFMintAddressProviderSession) SetCollateralPool(_addr common.Address) (*types.Transaction, error) {
	return _DefiFMintAddressProvider.Contract.SetCollateralPool(&_DefiFMintAddressProvider.TransactOpts, _addr)
}

// SetCollateralPool is a paid mutator transaction binding the contract method 0x1ba28878.
//
// Solidity: function setCollateralPool(address _addr) returns()
func (_DefiFMintAddressProvider *DefiFMintAddressProviderTransactorSession) SetCollateralPool(_addr common.Address) (*types.Transaction, error) {
	return _DefiFMintAddressProvider.Contract.SetCollateralPool(&_DefiFMintAddressProvider.TransactOpts, _addr)
}

// SetDebtPool is a paid mutator transaction binding the contract method 0x42ae8684.
//
// Solidity: function setDebtPool(address _addr) returns()
func (_DefiFMintAddressProvider *DefiFMintAddressProviderTransactor) SetDebtPool(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _DefiFMintAddressProvider.contract.Transact(opts, "setDebtPool", _addr)
}

// SetDebtPool is a paid mutator transaction binding the contract method 0x42ae8684.
//
// Solidity: function setDebtPool(address _addr) returns()
func (_DefiFMintAddressProvider *DefiFMintAddressProviderSession) SetDebtPool(_addr common.Address) (*types.Transaction, error) {
	return _DefiFMintAddressProvider.Contract.SetDebtPool(&_DefiFMintAddressProvider.TransactOpts, _addr)
}

// SetDebtPool is a paid mutator transaction binding the contract method 0x42ae8684.
//
// Solidity: function setDebtPool(address _addr) returns()
func (_DefiFMintAddressProvider *DefiFMintAddressProviderTransactorSession) SetDebtPool(_addr common.Address) (*types.Transaction, error) {
	return _DefiFMintAddressProvider.Contract.SetDebtPool(&_DefiFMintAddressProvider.TransactOpts, _addr)
}

// SetFantomMint is a paid mutator transaction binding the contract method 0xfcc3c074.
//
// Solidity: function setFantomMint(address _addr) returns()
func (_DefiFMintAddressProvider *DefiFMintAddressProviderTransactor) SetFantomMint(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _DefiFMintAddressProvider.contract.Transact(opts, "setFantomMint", _addr)
}

// SetFantomMint is a paid mutator transaction binding the contract method 0xfcc3c074.
//
// Solidity: function setFantomMint(address _addr) returns()
func (_DefiFMintAddressProvider *DefiFMintAddressProviderSession) SetFantomMint(_addr common.Address) (*types.Transaction, error) {
	return _DefiFMintAddressProvider.Contract.SetFantomMint(&_DefiFMintAddressProvider.TransactOpts, _addr)
}

// SetFantomMint is a paid mutator transaction binding the contract method 0xfcc3c074.
//
// Solidity: function setFantomMint(address _addr) returns()
func (_DefiFMintAddressProvider *DefiFMintAddressProviderTransactorSession) SetFantomMint(_addr common.Address) (*types.Transaction, error) {
	return _DefiFMintAddressProvider.Contract.SetFantomMint(&_DefiFMintAddressProvider.TransactOpts, _addr)
}

// SetPriceOracleProxy is a paid mutator transaction binding the contract method 0xcc653b9a.
//
// Solidity: function setPriceOracleProxy(address _addr) returns()
func (_DefiFMintAddressProvider *DefiFMintAddressProviderTransactor) SetPriceOracleProxy(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _DefiFMintAddressProvider.contract.Transact(opts, "setPriceOracleProxy", _addr)
}

// SetPriceOracleProxy is a paid mutator transaction binding the contract method 0xcc653b9a.
//
// Solidity: function setPriceOracleProxy(address _addr) returns()
func (_DefiFMintAddressProvider *DefiFMintAddressProviderSession) SetPriceOracleProxy(_addr common.Address) (*types.Transaction, error) {
	return _DefiFMintAddressProvider.Contract.SetPriceOracleProxy(&_DefiFMintAddressProvider.TransactOpts, _addr)
}

// SetPriceOracleProxy is a paid mutator transaction binding the contract method 0xcc653b9a.
//
// Solidity: function setPriceOracleProxy(address _addr) returns()
func (_DefiFMintAddressProvider *DefiFMintAddressProviderTransactorSession) SetPriceOracleProxy(_addr common.Address) (*types.Transaction, error) {
	return _DefiFMintAddressProvider.Contract.SetPriceOracleProxy(&_DefiFMintAddressProvider.TransactOpts, _addr)
}

// SetRewardDistribution is a paid mutator transaction binding the contract method 0x0d68b761.
//
// Solidity: function setRewardDistribution(address _addr) returns()
func (_DefiFMintAddressProvider *DefiFMintAddressProviderTransactor) SetRewardDistribution(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _DefiFMintAddressProvider.contract.Transact(opts, "setRewardDistribution", _addr)
}

// SetRewardDistribution is a paid mutator transaction binding the contract method 0x0d68b761.
//
// Solidity: function setRewardDistribution(address _addr) returns()
func (_DefiFMintAddressProvider *DefiFMintAddressProviderSession) SetRewardDistribution(_addr common.Address) (*types.Transaction, error) {
	return _DefiFMintAddressProvider.Contract.SetRewardDistribution(&_DefiFMintAddressProvider.TransactOpts, _addr)
}

// SetRewardDistribution is a paid mutator transaction binding the contract method 0x0d68b761.
//
// Solidity: function setRewardDistribution(address _addr) returns()
func (_DefiFMintAddressProvider *DefiFMintAddressProviderTransactorSession) SetRewardDistribution(_addr common.Address) (*types.Transaction, error) {
	return _DefiFMintAddressProvider.Contract.SetRewardDistribution(&_DefiFMintAddressProvider.TransactOpts, _addr)
}

// SetRewardToken is a paid mutator transaction binding the contract method 0x8aee8127.
//
// Solidity: function setRewardToken(address _addr) returns()
func (_DefiFMintAddressProvider *DefiFMintAddressProviderTransactor) SetRewardToken(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _DefiFMintAddressProvider.contract.Transact(opts, "setRewardToken", _addr)
}

// SetRewardToken is a paid mutator transaction binding the contract method 0x8aee8127.
//
// Solidity: function setRewardToken(address _addr) returns()
func (_DefiFMintAddressProvider *DefiFMintAddressProviderSession) SetRewardToken(_addr common.Address) (*types.Transaction, error) {
	return _DefiFMintAddressProvider.Contract.SetRewardToken(&_DefiFMintAddressProvider.TransactOpts, _addr)
}

// SetRewardToken is a paid mutator transaction binding the contract method 0x8aee8127.
//
// Solidity: function setRewardToken(address _addr) returns()
func (_DefiFMintAddressProvider *DefiFMintAddressProviderTransactorSession) SetRewardToken(_addr common.Address) (*types.Transaction, error) {
	return _DefiFMintAddressProvider.Contract.SetRewardToken(&_DefiFMintAddressProvider.TransactOpts, _addr)
}

// SetTokenRegistry is a paid mutator transaction binding the contract method 0x35a5af92.
//
// Solidity: function setTokenRegistry(address _addr) returns()
func (_DefiFMintAddressProvider *DefiFMintAddressProviderTransactor) SetTokenRegistry(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _DefiFMintAddressProvider.contract.Transact(opts, "setTokenRegistry", _addr)
}

// SetTokenRegistry is a paid mutator transaction binding the contract method 0x35a5af92.
//
// Solidity: function setTokenRegistry(address _addr) returns()
func (_DefiFMintAddressProvider *DefiFMintAddressProviderSession) SetTokenRegistry(_addr common.Address) (*types.Transaction, error) {
	return _DefiFMintAddressProvider.Contract.SetTokenRegistry(&_DefiFMintAddressProvider.TransactOpts, _addr)
}

// SetTokenRegistry is a paid mutator transaction binding the contract method 0x35a5af92.
//
// Solidity: function setTokenRegistry(address _addr) returns()
func (_DefiFMintAddressProvider *DefiFMintAddressProviderTransactorSession) SetTokenRegistry(_addr common.Address) (*types.Transaction, error) {
	return _DefiFMintAddressProvider.Contract.SetTokenRegistry(&_DefiFMintAddressProvider.TransactOpts, _addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DefiFMintAddressProvider *DefiFMintAddressProviderTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DefiFMintAddressProvider.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DefiFMintAddressProvider *DefiFMintAddressProviderSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DefiFMintAddressProvider.Contract.TransferOwnership(&_DefiFMintAddressProvider.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DefiFMintAddressProvider *DefiFMintAddressProviderTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DefiFMintAddressProvider.Contract.TransferOwnership(&_DefiFMintAddressProvider.TransactOpts, newOwner)
}

// DefiFMintAddressProviderCollateralPoolChangedIterator is returned from FilterCollateralPoolChanged and is used to iterate over the raw logs and unpacked data for CollateralPoolChanged events raised by the DefiFMintAddressProvider contract.
type DefiFMintAddressProviderCollateralPoolChangedIterator struct {
	Event *DefiFMintAddressProviderCollateralPoolChanged // Event containing the contract specifics and raw log

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
func (it *DefiFMintAddressProviderCollateralPoolChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefiFMintAddressProviderCollateralPoolChanged)
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
		it.Event = new(DefiFMintAddressProviderCollateralPoolChanged)
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
func (it *DefiFMintAddressProviderCollateralPoolChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefiFMintAddressProviderCollateralPoolChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefiFMintAddressProviderCollateralPoolChanged represents a CollateralPoolChanged event raised by the DefiFMintAddressProvider contract.
type DefiFMintAddressProviderCollateralPoolChanged struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCollateralPoolChanged is a free log retrieval operation binding the contract event 0x9ee268d502b2200c71bbd3cba8222b4f501a7b505c684bd40423fd446bb29fad.
//
// Solidity: event CollateralPoolChanged(address newAddress)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderFilterer) FilterCollateralPoolChanged(opts *bind.FilterOpts) (*DefiFMintAddressProviderCollateralPoolChangedIterator, error) {

	logs, sub, err := _DefiFMintAddressProvider.contract.FilterLogs(opts, "CollateralPoolChanged")
	if err != nil {
		return nil, err
	}
	return &DefiFMintAddressProviderCollateralPoolChangedIterator{contract: _DefiFMintAddressProvider.contract, event: "CollateralPoolChanged", logs: logs, sub: sub}, nil
}

// WatchCollateralPoolChanged is a free log subscription operation binding the contract event 0x9ee268d502b2200c71bbd3cba8222b4f501a7b505c684bd40423fd446bb29fad.
//
// Solidity: event CollateralPoolChanged(address newAddress)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderFilterer) WatchCollateralPoolChanged(opts *bind.WatchOpts, sink chan<- *DefiFMintAddressProviderCollateralPoolChanged) (event.Subscription, error) {

	logs, sub, err := _DefiFMintAddressProvider.contract.WatchLogs(opts, "CollateralPoolChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefiFMintAddressProviderCollateralPoolChanged)
				if err := _DefiFMintAddressProvider.contract.UnpackLog(event, "CollateralPoolChanged", log); err != nil {
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

// ParseCollateralPoolChanged is a log parse operation binding the contract event 0x9ee268d502b2200c71bbd3cba8222b4f501a7b505c684bd40423fd446bb29fad.
//
// Solidity: event CollateralPoolChanged(address newAddress)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderFilterer) ParseCollateralPoolChanged(log types.Log) (*DefiFMintAddressProviderCollateralPoolChanged, error) {
	event := new(DefiFMintAddressProviderCollateralPoolChanged)
	if err := _DefiFMintAddressProvider.contract.UnpackLog(event, "CollateralPoolChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DefiFMintAddressProviderDebtPoolChangedIterator is returned from FilterDebtPoolChanged and is used to iterate over the raw logs and unpacked data for DebtPoolChanged events raised by the DefiFMintAddressProvider contract.
type DefiFMintAddressProviderDebtPoolChangedIterator struct {
	Event *DefiFMintAddressProviderDebtPoolChanged // Event containing the contract specifics and raw log

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
func (it *DefiFMintAddressProviderDebtPoolChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefiFMintAddressProviderDebtPoolChanged)
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
		it.Event = new(DefiFMintAddressProviderDebtPoolChanged)
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
func (it *DefiFMintAddressProviderDebtPoolChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefiFMintAddressProviderDebtPoolChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefiFMintAddressProviderDebtPoolChanged represents a DebtPoolChanged event raised by the DefiFMintAddressProvider contract.
type DefiFMintAddressProviderDebtPoolChanged struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDebtPoolChanged is a free log retrieval operation binding the contract event 0xf10b554a663200a2ae53269b5aeb591082984e03e47f76ec558f283c01b116d4.
//
// Solidity: event DebtPoolChanged(address newAddress)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderFilterer) FilterDebtPoolChanged(opts *bind.FilterOpts) (*DefiFMintAddressProviderDebtPoolChangedIterator, error) {

	logs, sub, err := _DefiFMintAddressProvider.contract.FilterLogs(opts, "DebtPoolChanged")
	if err != nil {
		return nil, err
	}
	return &DefiFMintAddressProviderDebtPoolChangedIterator{contract: _DefiFMintAddressProvider.contract, event: "DebtPoolChanged", logs: logs, sub: sub}, nil
}

// WatchDebtPoolChanged is a free log subscription operation binding the contract event 0xf10b554a663200a2ae53269b5aeb591082984e03e47f76ec558f283c01b116d4.
//
// Solidity: event DebtPoolChanged(address newAddress)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderFilterer) WatchDebtPoolChanged(opts *bind.WatchOpts, sink chan<- *DefiFMintAddressProviderDebtPoolChanged) (event.Subscription, error) {

	logs, sub, err := _DefiFMintAddressProvider.contract.WatchLogs(opts, "DebtPoolChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefiFMintAddressProviderDebtPoolChanged)
				if err := _DefiFMintAddressProvider.contract.UnpackLog(event, "DebtPoolChanged", log); err != nil {
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

// ParseDebtPoolChanged is a log parse operation binding the contract event 0xf10b554a663200a2ae53269b5aeb591082984e03e47f76ec558f283c01b116d4.
//
// Solidity: event DebtPoolChanged(address newAddress)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderFilterer) ParseDebtPoolChanged(log types.Log) (*DefiFMintAddressProviderDebtPoolChanged, error) {
	event := new(DefiFMintAddressProviderDebtPoolChanged)
	if err := _DefiFMintAddressProvider.contract.UnpackLog(event, "DebtPoolChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DefiFMintAddressProviderMinterChangedIterator is returned from FilterMinterChanged and is used to iterate over the raw logs and unpacked data for MinterChanged events raised by the DefiFMintAddressProvider contract.
type DefiFMintAddressProviderMinterChangedIterator struct {
	Event *DefiFMintAddressProviderMinterChanged // Event containing the contract specifics and raw log

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
func (it *DefiFMintAddressProviderMinterChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefiFMintAddressProviderMinterChanged)
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
		it.Event = new(DefiFMintAddressProviderMinterChanged)
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
func (it *DefiFMintAddressProviderMinterChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefiFMintAddressProviderMinterChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefiFMintAddressProviderMinterChanged represents a MinterChanged event raised by the DefiFMintAddressProvider contract.
type DefiFMintAddressProviderMinterChanged struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMinterChanged is a free log retrieval operation binding the contract event 0xb6b8f1859c5c352e5ffad07d0f77e384ac725512c015bd3a3ffc885831c8a425.
//
// Solidity: event MinterChanged(address newAddress)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderFilterer) FilterMinterChanged(opts *bind.FilterOpts) (*DefiFMintAddressProviderMinterChangedIterator, error) {

	logs, sub, err := _DefiFMintAddressProvider.contract.FilterLogs(opts, "MinterChanged")
	if err != nil {
		return nil, err
	}
	return &DefiFMintAddressProviderMinterChangedIterator{contract: _DefiFMintAddressProvider.contract, event: "MinterChanged", logs: logs, sub: sub}, nil
}

// WatchMinterChanged is a free log subscription operation binding the contract event 0xb6b8f1859c5c352e5ffad07d0f77e384ac725512c015bd3a3ffc885831c8a425.
//
// Solidity: event MinterChanged(address newAddress)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderFilterer) WatchMinterChanged(opts *bind.WatchOpts, sink chan<- *DefiFMintAddressProviderMinterChanged) (event.Subscription, error) {

	logs, sub, err := _DefiFMintAddressProvider.contract.WatchLogs(opts, "MinterChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefiFMintAddressProviderMinterChanged)
				if err := _DefiFMintAddressProvider.contract.UnpackLog(event, "MinterChanged", log); err != nil {
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

// ParseMinterChanged is a log parse operation binding the contract event 0xb6b8f1859c5c352e5ffad07d0f77e384ac725512c015bd3a3ffc885831c8a425.
//
// Solidity: event MinterChanged(address newAddress)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderFilterer) ParseMinterChanged(log types.Log) (*DefiFMintAddressProviderMinterChanged, error) {
	event := new(DefiFMintAddressProviderMinterChanged)
	if err := _DefiFMintAddressProvider.contract.UnpackLog(event, "MinterChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DefiFMintAddressProviderOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DefiFMintAddressProvider contract.
type DefiFMintAddressProviderOwnershipTransferredIterator struct {
	Event *DefiFMintAddressProviderOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DefiFMintAddressProviderOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefiFMintAddressProviderOwnershipTransferred)
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
		it.Event = new(DefiFMintAddressProviderOwnershipTransferred)
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
func (it *DefiFMintAddressProviderOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefiFMintAddressProviderOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefiFMintAddressProviderOwnershipTransferred represents a OwnershipTransferred event raised by the DefiFMintAddressProvider contract.
type DefiFMintAddressProviderOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DefiFMintAddressProviderOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DefiFMintAddressProvider.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DefiFMintAddressProviderOwnershipTransferredIterator{contract: _DefiFMintAddressProvider.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DefiFMintAddressProviderOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DefiFMintAddressProvider.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefiFMintAddressProviderOwnershipTransferred)
				if err := _DefiFMintAddressProvider.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_DefiFMintAddressProvider *DefiFMintAddressProviderFilterer) ParseOwnershipTransferred(log types.Log) (*DefiFMintAddressProviderOwnershipTransferred, error) {
	event := new(DefiFMintAddressProviderOwnershipTransferred)
	if err := _DefiFMintAddressProvider.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DefiFMintAddressProviderPriceOracleChangedIterator is returned from FilterPriceOracleChanged and is used to iterate over the raw logs and unpacked data for PriceOracleChanged events raised by the DefiFMintAddressProvider contract.
type DefiFMintAddressProviderPriceOracleChangedIterator struct {
	Event *DefiFMintAddressProviderPriceOracleChanged // Event containing the contract specifics and raw log

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
func (it *DefiFMintAddressProviderPriceOracleChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefiFMintAddressProviderPriceOracleChanged)
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
		it.Event = new(DefiFMintAddressProviderPriceOracleChanged)
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
func (it *DefiFMintAddressProviderPriceOracleChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefiFMintAddressProviderPriceOracleChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefiFMintAddressProviderPriceOracleChanged represents a PriceOracleChanged event raised by the DefiFMintAddressProvider contract.
type DefiFMintAddressProviderPriceOracleChanged struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPriceOracleChanged is a free log retrieval operation binding the contract event 0xb36d86785c7d32b1ad714bb705e00e93eccc37b8cf47549043e61e10908ad251.
//
// Solidity: event PriceOracleChanged(address newAddress)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderFilterer) FilterPriceOracleChanged(opts *bind.FilterOpts) (*DefiFMintAddressProviderPriceOracleChangedIterator, error) {

	logs, sub, err := _DefiFMintAddressProvider.contract.FilterLogs(opts, "PriceOracleChanged")
	if err != nil {
		return nil, err
	}
	return &DefiFMintAddressProviderPriceOracleChangedIterator{contract: _DefiFMintAddressProvider.contract, event: "PriceOracleChanged", logs: logs, sub: sub}, nil
}

// WatchPriceOracleChanged is a free log subscription operation binding the contract event 0xb36d86785c7d32b1ad714bb705e00e93eccc37b8cf47549043e61e10908ad251.
//
// Solidity: event PriceOracleChanged(address newAddress)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderFilterer) WatchPriceOracleChanged(opts *bind.WatchOpts, sink chan<- *DefiFMintAddressProviderPriceOracleChanged) (event.Subscription, error) {

	logs, sub, err := _DefiFMintAddressProvider.contract.WatchLogs(opts, "PriceOracleChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefiFMintAddressProviderPriceOracleChanged)
				if err := _DefiFMintAddressProvider.contract.UnpackLog(event, "PriceOracleChanged", log); err != nil {
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

// ParsePriceOracleChanged is a log parse operation binding the contract event 0xb36d86785c7d32b1ad714bb705e00e93eccc37b8cf47549043e61e10908ad251.
//
// Solidity: event PriceOracleChanged(address newAddress)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderFilterer) ParsePriceOracleChanged(log types.Log) (*DefiFMintAddressProviderPriceOracleChanged, error) {
	event := new(DefiFMintAddressProviderPriceOracleChanged)
	if err := _DefiFMintAddressProvider.contract.UnpackLog(event, "PriceOracleChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DefiFMintAddressProviderRewardDistributionChangedIterator is returned from FilterRewardDistributionChanged and is used to iterate over the raw logs and unpacked data for RewardDistributionChanged events raised by the DefiFMintAddressProvider contract.
type DefiFMintAddressProviderRewardDistributionChangedIterator struct {
	Event *DefiFMintAddressProviderRewardDistributionChanged // Event containing the contract specifics and raw log

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
func (it *DefiFMintAddressProviderRewardDistributionChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefiFMintAddressProviderRewardDistributionChanged)
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
		it.Event = new(DefiFMintAddressProviderRewardDistributionChanged)
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
func (it *DefiFMintAddressProviderRewardDistributionChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefiFMintAddressProviderRewardDistributionChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefiFMintAddressProviderRewardDistributionChanged represents a RewardDistributionChanged event raised by the DefiFMintAddressProvider contract.
type DefiFMintAddressProviderRewardDistributionChanged struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRewardDistributionChanged is a free log retrieval operation binding the contract event 0xfe09426f22c44354b62f360c333309adadd6392ae248adc902f3006c7c4b9205.
//
// Solidity: event RewardDistributionChanged(address newAddress)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderFilterer) FilterRewardDistributionChanged(opts *bind.FilterOpts) (*DefiFMintAddressProviderRewardDistributionChangedIterator, error) {

	logs, sub, err := _DefiFMintAddressProvider.contract.FilterLogs(opts, "RewardDistributionChanged")
	if err != nil {
		return nil, err
	}
	return &DefiFMintAddressProviderRewardDistributionChangedIterator{contract: _DefiFMintAddressProvider.contract, event: "RewardDistributionChanged", logs: logs, sub: sub}, nil
}

// WatchRewardDistributionChanged is a free log subscription operation binding the contract event 0xfe09426f22c44354b62f360c333309adadd6392ae248adc902f3006c7c4b9205.
//
// Solidity: event RewardDistributionChanged(address newAddress)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderFilterer) WatchRewardDistributionChanged(opts *bind.WatchOpts, sink chan<- *DefiFMintAddressProviderRewardDistributionChanged) (event.Subscription, error) {

	logs, sub, err := _DefiFMintAddressProvider.contract.WatchLogs(opts, "RewardDistributionChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefiFMintAddressProviderRewardDistributionChanged)
				if err := _DefiFMintAddressProvider.contract.UnpackLog(event, "RewardDistributionChanged", log); err != nil {
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

// ParseRewardDistributionChanged is a log parse operation binding the contract event 0xfe09426f22c44354b62f360c333309adadd6392ae248adc902f3006c7c4b9205.
//
// Solidity: event RewardDistributionChanged(address newAddress)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderFilterer) ParseRewardDistributionChanged(log types.Log) (*DefiFMintAddressProviderRewardDistributionChanged, error) {
	event := new(DefiFMintAddressProviderRewardDistributionChanged)
	if err := _DefiFMintAddressProvider.contract.UnpackLog(event, "RewardDistributionChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DefiFMintAddressProviderRewardTokenChangedIterator is returned from FilterRewardTokenChanged and is used to iterate over the raw logs and unpacked data for RewardTokenChanged events raised by the DefiFMintAddressProvider contract.
type DefiFMintAddressProviderRewardTokenChangedIterator struct {
	Event *DefiFMintAddressProviderRewardTokenChanged // Event containing the contract specifics and raw log

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
func (it *DefiFMintAddressProviderRewardTokenChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefiFMintAddressProviderRewardTokenChanged)
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
		it.Event = new(DefiFMintAddressProviderRewardTokenChanged)
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
func (it *DefiFMintAddressProviderRewardTokenChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefiFMintAddressProviderRewardTokenChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefiFMintAddressProviderRewardTokenChanged represents a RewardTokenChanged event raised by the DefiFMintAddressProvider contract.
type DefiFMintAddressProviderRewardTokenChanged struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRewardTokenChanged is a free log retrieval operation binding the contract event 0xb74d956cf6ec7842d08ebf0ab19ec03a88c1efd4a50ea4349d30f9c4ce512e98.
//
// Solidity: event RewardTokenChanged(address newAddress)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderFilterer) FilterRewardTokenChanged(opts *bind.FilterOpts) (*DefiFMintAddressProviderRewardTokenChangedIterator, error) {

	logs, sub, err := _DefiFMintAddressProvider.contract.FilterLogs(opts, "RewardTokenChanged")
	if err != nil {
		return nil, err
	}
	return &DefiFMintAddressProviderRewardTokenChangedIterator{contract: _DefiFMintAddressProvider.contract, event: "RewardTokenChanged", logs: logs, sub: sub}, nil
}

// WatchRewardTokenChanged is a free log subscription operation binding the contract event 0xb74d956cf6ec7842d08ebf0ab19ec03a88c1efd4a50ea4349d30f9c4ce512e98.
//
// Solidity: event RewardTokenChanged(address newAddress)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderFilterer) WatchRewardTokenChanged(opts *bind.WatchOpts, sink chan<- *DefiFMintAddressProviderRewardTokenChanged) (event.Subscription, error) {

	logs, sub, err := _DefiFMintAddressProvider.contract.WatchLogs(opts, "RewardTokenChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefiFMintAddressProviderRewardTokenChanged)
				if err := _DefiFMintAddressProvider.contract.UnpackLog(event, "RewardTokenChanged", log); err != nil {
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

// ParseRewardTokenChanged is a log parse operation binding the contract event 0xb74d956cf6ec7842d08ebf0ab19ec03a88c1efd4a50ea4349d30f9c4ce512e98.
//
// Solidity: event RewardTokenChanged(address newAddress)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderFilterer) ParseRewardTokenChanged(log types.Log) (*DefiFMintAddressProviderRewardTokenChanged, error) {
	event := new(DefiFMintAddressProviderRewardTokenChanged)
	if err := _DefiFMintAddressProvider.contract.UnpackLog(event, "RewardTokenChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DefiFMintAddressProviderTokenRegistryChangedIterator is returned from FilterTokenRegistryChanged and is used to iterate over the raw logs and unpacked data for TokenRegistryChanged events raised by the DefiFMintAddressProvider contract.
type DefiFMintAddressProviderTokenRegistryChangedIterator struct {
	Event *DefiFMintAddressProviderTokenRegistryChanged // Event containing the contract specifics and raw log

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
func (it *DefiFMintAddressProviderTokenRegistryChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefiFMintAddressProviderTokenRegistryChanged)
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
		it.Event = new(DefiFMintAddressProviderTokenRegistryChanged)
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
func (it *DefiFMintAddressProviderTokenRegistryChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefiFMintAddressProviderTokenRegistryChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefiFMintAddressProviderTokenRegistryChanged represents a TokenRegistryChanged event raised by the DefiFMintAddressProvider contract.
type DefiFMintAddressProviderTokenRegistryChanged struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTokenRegistryChanged is a free log retrieval operation binding the contract event 0xb6f925ec7d36d613e5d1aa87c0de3ee16a0167e6bdfa2ea254e5fee9870a941e.
//
// Solidity: event TokenRegistryChanged(address newAddress)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderFilterer) FilterTokenRegistryChanged(opts *bind.FilterOpts) (*DefiFMintAddressProviderTokenRegistryChangedIterator, error) {

	logs, sub, err := _DefiFMintAddressProvider.contract.FilterLogs(opts, "TokenRegistryChanged")
	if err != nil {
		return nil, err
	}
	return &DefiFMintAddressProviderTokenRegistryChangedIterator{contract: _DefiFMintAddressProvider.contract, event: "TokenRegistryChanged", logs: logs, sub: sub}, nil
}

// WatchTokenRegistryChanged is a free log subscription operation binding the contract event 0xb6f925ec7d36d613e5d1aa87c0de3ee16a0167e6bdfa2ea254e5fee9870a941e.
//
// Solidity: event TokenRegistryChanged(address newAddress)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderFilterer) WatchTokenRegistryChanged(opts *bind.WatchOpts, sink chan<- *DefiFMintAddressProviderTokenRegistryChanged) (event.Subscription, error) {

	logs, sub, err := _DefiFMintAddressProvider.contract.WatchLogs(opts, "TokenRegistryChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefiFMintAddressProviderTokenRegistryChanged)
				if err := _DefiFMintAddressProvider.contract.UnpackLog(event, "TokenRegistryChanged", log); err != nil {
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

// ParseTokenRegistryChanged is a log parse operation binding the contract event 0xb6f925ec7d36d613e5d1aa87c0de3ee16a0167e6bdfa2ea254e5fee9870a941e.
//
// Solidity: event TokenRegistryChanged(address newAddress)
func (_DefiFMintAddressProvider *DefiFMintAddressProviderFilterer) ParseTokenRegistryChanged(log types.Log) (*DefiFMintAddressProviderTokenRegistryChanged, error) {
	event := new(DefiFMintAddressProviderTokenRegistryChanged)
	if err := _DefiFMintAddressProvider.contract.UnpackLog(event, "TokenRegistryChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
