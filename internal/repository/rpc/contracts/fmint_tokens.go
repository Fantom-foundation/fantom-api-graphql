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

// DefiFMintTokenRegistryMetaData contains all meta data concerning the DefiFMintTokenRegistry contract.
var DefiFMintTokenRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"TokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"TokenUpdated\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"logo\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"priceDecimals\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"canDeposit\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"canMint\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokensList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokensCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"priceDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"isActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"canDeposit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"canMint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_logo\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_priceDecimals\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"_isActive\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_canDeposit\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_canMint\",\"type\":\"bool\"}],\"name\":\"addToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_logo\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_priceDecimals\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"_isActive\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_canDeposit\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_canMint\",\"type\":\"bool\"}],\"name\":\"updateToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DefiFMintTokenRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use DefiFMintTokenRegistryMetaData.ABI instead.
var DefiFMintTokenRegistryABI = DefiFMintTokenRegistryMetaData.ABI

// DefiFMintTokenRegistry is an auto generated Go binding around an Ethereum contract.
type DefiFMintTokenRegistry struct {
	DefiFMintTokenRegistryCaller     // Read-only binding to the contract
	DefiFMintTokenRegistryTransactor // Write-only binding to the contract
	DefiFMintTokenRegistryFilterer   // Log filterer for contract events
}

// DefiFMintTokenRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type DefiFMintTokenRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefiFMintTokenRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DefiFMintTokenRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefiFMintTokenRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DefiFMintTokenRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefiFMintTokenRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DefiFMintTokenRegistrySession struct {
	Contract     *DefiFMintTokenRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// DefiFMintTokenRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DefiFMintTokenRegistryCallerSession struct {
	Contract *DefiFMintTokenRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// DefiFMintTokenRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DefiFMintTokenRegistryTransactorSession struct {
	Contract     *DefiFMintTokenRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// DefiFMintTokenRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type DefiFMintTokenRegistryRaw struct {
	Contract *DefiFMintTokenRegistry // Generic contract binding to access the raw methods on
}

// DefiFMintTokenRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DefiFMintTokenRegistryCallerRaw struct {
	Contract *DefiFMintTokenRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// DefiFMintTokenRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DefiFMintTokenRegistryTransactorRaw struct {
	Contract *DefiFMintTokenRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDefiFMintTokenRegistry creates a new instance of DefiFMintTokenRegistry, bound to a specific deployed contract.
func NewDefiFMintTokenRegistry(address common.Address, backend bind.ContractBackend) (*DefiFMintTokenRegistry, error) {
	contract, err := bindDefiFMintTokenRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DefiFMintTokenRegistry{DefiFMintTokenRegistryCaller: DefiFMintTokenRegistryCaller{contract: contract}, DefiFMintTokenRegistryTransactor: DefiFMintTokenRegistryTransactor{contract: contract}, DefiFMintTokenRegistryFilterer: DefiFMintTokenRegistryFilterer{contract: contract}}, nil
}

// NewDefiFMintTokenRegistryCaller creates a new read-only instance of DefiFMintTokenRegistry, bound to a specific deployed contract.
func NewDefiFMintTokenRegistryCaller(address common.Address, caller bind.ContractCaller) (*DefiFMintTokenRegistryCaller, error) {
	contract, err := bindDefiFMintTokenRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DefiFMintTokenRegistryCaller{contract: contract}, nil
}

// NewDefiFMintTokenRegistryTransactor creates a new write-only instance of DefiFMintTokenRegistry, bound to a specific deployed contract.
func NewDefiFMintTokenRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*DefiFMintTokenRegistryTransactor, error) {
	contract, err := bindDefiFMintTokenRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DefiFMintTokenRegistryTransactor{contract: contract}, nil
}

// NewDefiFMintTokenRegistryFilterer creates a new log filterer instance of DefiFMintTokenRegistry, bound to a specific deployed contract.
func NewDefiFMintTokenRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*DefiFMintTokenRegistryFilterer, error) {
	contract, err := bindDefiFMintTokenRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DefiFMintTokenRegistryFilterer{contract: contract}, nil
}

// bindDefiFMintTokenRegistry binds a generic wrapper to an already deployed contract.
func bindDefiFMintTokenRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DefiFMintTokenRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DefiFMintTokenRegistry.Contract.DefiFMintTokenRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefiFMintTokenRegistry.Contract.DefiFMintTokenRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DefiFMintTokenRegistry.Contract.DefiFMintTokenRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DefiFMintTokenRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefiFMintTokenRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DefiFMintTokenRegistry.Contract.contract.Transact(opts, method, params...)
}

// CanDeposit is a free data retrieval call binding the contract method 0x4bf0d331.
//
// Solidity: function canDeposit(address _token) view returns(bool)
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistryCaller) CanDeposit(opts *bind.CallOpts, _token common.Address) (bool, error) {
	var out []interface{}
	err := _DefiFMintTokenRegistry.contract.Call(opts, &out, "canDeposit", _token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanDeposit is a free data retrieval call binding the contract method 0x4bf0d331.
//
// Solidity: function canDeposit(address _token) view returns(bool)
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistrySession) CanDeposit(_token common.Address) (bool, error) {
	return _DefiFMintTokenRegistry.Contract.CanDeposit(&_DefiFMintTokenRegistry.CallOpts, _token)
}

// CanDeposit is a free data retrieval call binding the contract method 0x4bf0d331.
//
// Solidity: function canDeposit(address _token) view returns(bool)
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistryCallerSession) CanDeposit(_token common.Address) (bool, error) {
	return _DefiFMintTokenRegistry.Contract.CanDeposit(&_DefiFMintTokenRegistry.CallOpts, _token)
}

// CanMint is a free data retrieval call binding the contract method 0xc2ba4744.
//
// Solidity: function canMint(address _token) view returns(bool)
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistryCaller) CanMint(opts *bind.CallOpts, _token common.Address) (bool, error) {
	var out []interface{}
	err := _DefiFMintTokenRegistry.contract.Call(opts, &out, "canMint", _token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanMint is a free data retrieval call binding the contract method 0xc2ba4744.
//
// Solidity: function canMint(address _token) view returns(bool)
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistrySession) CanMint(_token common.Address) (bool, error) {
	return _DefiFMintTokenRegistry.Contract.CanMint(&_DefiFMintTokenRegistry.CallOpts, _token)
}

// CanMint is a free data retrieval call binding the contract method 0xc2ba4744.
//
// Solidity: function canMint(address _token) view returns(bool)
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistryCallerSession) CanMint(_token common.Address) (bool, error) {
	return _DefiFMintTokenRegistry.Contract.CanMint(&_DefiFMintTokenRegistry.CallOpts, _token)
}

// IsActive is a free data retrieval call binding the contract method 0x9f8a13d7.
//
// Solidity: function isActive(address _token) view returns(bool)
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistryCaller) IsActive(opts *bind.CallOpts, _token common.Address) (bool, error) {
	var out []interface{}
	err := _DefiFMintTokenRegistry.contract.Call(opts, &out, "isActive", _token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActive is a free data retrieval call binding the contract method 0x9f8a13d7.
//
// Solidity: function isActive(address _token) view returns(bool)
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistrySession) IsActive(_token common.Address) (bool, error) {
	return _DefiFMintTokenRegistry.Contract.IsActive(&_DefiFMintTokenRegistry.CallOpts, _token)
}

// IsActive is a free data retrieval call binding the contract method 0x9f8a13d7.
//
// Solidity: function isActive(address _token) view returns(bool)
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistryCallerSession) IsActive(_token common.Address) (bool, error) {
	return _DefiFMintTokenRegistry.Contract.IsActive(&_DefiFMintTokenRegistry.CallOpts, _token)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistryCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DefiFMintTokenRegistry.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistrySession) IsOwner() (bool, error) {
	return _DefiFMintTokenRegistry.Contract.IsOwner(&_DefiFMintTokenRegistry.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistryCallerSession) IsOwner() (bool, error) {
	return _DefiFMintTokenRegistry.Contract.IsOwner(&_DefiFMintTokenRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DefiFMintTokenRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistrySession) Owner() (common.Address, error) {
	return _DefiFMintTokenRegistry.Contract.Owner(&_DefiFMintTokenRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistryCallerSession) Owner() (common.Address, error) {
	return _DefiFMintTokenRegistry.Contract.Owner(&_DefiFMintTokenRegistry.CallOpts)
}

// PriceDecimals is a free data retrieval call binding the contract method 0xcefe0f21.
//
// Solidity: function priceDecimals(address _token) view returns(uint8)
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistryCaller) PriceDecimals(opts *bind.CallOpts, _token common.Address) (uint8, error) {
	var out []interface{}
	err := _DefiFMintTokenRegistry.contract.Call(opts, &out, "priceDecimals", _token)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PriceDecimals is a free data retrieval call binding the contract method 0xcefe0f21.
//
// Solidity: function priceDecimals(address _token) view returns(uint8)
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistrySession) PriceDecimals(_token common.Address) (uint8, error) {
	return _DefiFMintTokenRegistry.Contract.PriceDecimals(&_DefiFMintTokenRegistry.CallOpts, _token)
}

// PriceDecimals is a free data retrieval call binding the contract method 0xcefe0f21.
//
// Solidity: function priceDecimals(address _token) view returns(uint8)
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistryCallerSession) PriceDecimals(_token common.Address) (uint8, error) {
	return _DefiFMintTokenRegistry.Contract.PriceDecimals(&_DefiFMintTokenRegistry.CallOpts, _token)
}

// Tokens is a free data retrieval call binding the contract method 0xe4860339.
//
// Solidity: function tokens(address ) view returns(uint256 id, string name, string symbol, uint8 decimals, string logo, address oracle, uint8 priceDecimals, bool isActive, bool canDeposit, bool canMint)
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistryCaller) Tokens(opts *bind.CallOpts, arg0 common.Address) (struct {
	Id            *big.Int
	Name          string
	Symbol        string
	Decimals      uint8
	Logo          string
	Oracle        common.Address
	PriceDecimals uint8
	IsActive      bool
	CanDeposit    bool
	CanMint       bool
}, error) {
	var out []interface{}
	err := _DefiFMintTokenRegistry.contract.Call(opts, &out, "tokens", arg0)

	outstruct := new(struct {
		Id            *big.Int
		Name          string
		Symbol        string
		Decimals      uint8
		Logo          string
		Oracle        common.Address
		PriceDecimals uint8
		IsActive      bool
		CanDeposit    bool
		CanMint       bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Symbol = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Decimals = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.Logo = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Oracle = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.PriceDecimals = *abi.ConvertType(out[6], new(uint8)).(*uint8)
	outstruct.IsActive = *abi.ConvertType(out[7], new(bool)).(*bool)
	outstruct.CanDeposit = *abi.ConvertType(out[8], new(bool)).(*bool)
	outstruct.CanMint = *abi.ConvertType(out[9], new(bool)).(*bool)

	return *outstruct, err

}

// Tokens is a free data retrieval call binding the contract method 0xe4860339.
//
// Solidity: function tokens(address ) view returns(uint256 id, string name, string symbol, uint8 decimals, string logo, address oracle, uint8 priceDecimals, bool isActive, bool canDeposit, bool canMint)
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistrySession) Tokens(arg0 common.Address) (struct {
	Id            *big.Int
	Name          string
	Symbol        string
	Decimals      uint8
	Logo          string
	Oracle        common.Address
	PriceDecimals uint8
	IsActive      bool
	CanDeposit    bool
	CanMint       bool
}, error) {
	return _DefiFMintTokenRegistry.Contract.Tokens(&_DefiFMintTokenRegistry.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0xe4860339.
//
// Solidity: function tokens(address ) view returns(uint256 id, string name, string symbol, uint8 decimals, string logo, address oracle, uint8 priceDecimals, bool isActive, bool canDeposit, bool canMint)
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistryCallerSession) Tokens(arg0 common.Address) (struct {
	Id            *big.Int
	Name          string
	Symbol        string
	Decimals      uint8
	Logo          string
	Oracle        common.Address
	PriceDecimals uint8
	IsActive      bool
	CanDeposit    bool
	CanMint       bool
}, error) {
	return _DefiFMintTokenRegistry.Contract.Tokens(&_DefiFMintTokenRegistry.CallOpts, arg0)
}

// TokensCount is a free data retrieval call binding the contract method 0xa64ed8ba.
//
// Solidity: function tokensCount() view returns(uint256)
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistryCaller) TokensCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DefiFMintTokenRegistry.contract.Call(opts, &out, "tokensCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokensCount is a free data retrieval call binding the contract method 0xa64ed8ba.
//
// Solidity: function tokensCount() view returns(uint256)
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistrySession) TokensCount() (*big.Int, error) {
	return _DefiFMintTokenRegistry.Contract.TokensCount(&_DefiFMintTokenRegistry.CallOpts)
}

// TokensCount is a free data retrieval call binding the contract method 0xa64ed8ba.
//
// Solidity: function tokensCount() view returns(uint256)
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistryCallerSession) TokensCount() (*big.Int, error) {
	return _DefiFMintTokenRegistry.Contract.TokensCount(&_DefiFMintTokenRegistry.CallOpts)
}

// TokensList is a free data retrieval call binding the contract method 0x4d12e34e.
//
// Solidity: function tokensList(uint256 ) view returns(address)
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistryCaller) TokensList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DefiFMintTokenRegistry.contract.Call(opts, &out, "tokensList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokensList is a free data retrieval call binding the contract method 0x4d12e34e.
//
// Solidity: function tokensList(uint256 ) view returns(address)
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistrySession) TokensList(arg0 *big.Int) (common.Address, error) {
	return _DefiFMintTokenRegistry.Contract.TokensList(&_DefiFMintTokenRegistry.CallOpts, arg0)
}

// TokensList is a free data retrieval call binding the contract method 0x4d12e34e.
//
// Solidity: function tokensList(uint256 ) view returns(address)
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistryCallerSession) TokensList(arg0 *big.Int) (common.Address, error) {
	return _DefiFMintTokenRegistry.Contract.TokensList(&_DefiFMintTokenRegistry.CallOpts, arg0)
}

// AddToken is a paid mutator transaction binding the contract method 0x6a77a214.
//
// Solidity: function addToken(address _token, string _logo, address _oracle, uint8 _priceDecimals, bool _isActive, bool _canDeposit, bool _canMint) returns()
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistryTransactor) AddToken(opts *bind.TransactOpts, _token common.Address, _logo string, _oracle common.Address, _priceDecimals uint8, _isActive bool, _canDeposit bool, _canMint bool) (*types.Transaction, error) {
	return _DefiFMintTokenRegistry.contract.Transact(opts, "addToken", _token, _logo, _oracle, _priceDecimals, _isActive, _canDeposit, _canMint)
}

// AddToken is a paid mutator transaction binding the contract method 0x6a77a214.
//
// Solidity: function addToken(address _token, string _logo, address _oracle, uint8 _priceDecimals, bool _isActive, bool _canDeposit, bool _canMint) returns()
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistrySession) AddToken(_token common.Address, _logo string, _oracle common.Address, _priceDecimals uint8, _isActive bool, _canDeposit bool, _canMint bool) (*types.Transaction, error) {
	return _DefiFMintTokenRegistry.Contract.AddToken(&_DefiFMintTokenRegistry.TransactOpts, _token, _logo, _oracle, _priceDecimals, _isActive, _canDeposit, _canMint)
}

// AddToken is a paid mutator transaction binding the contract method 0x6a77a214.
//
// Solidity: function addToken(address _token, string _logo, address _oracle, uint8 _priceDecimals, bool _isActive, bool _canDeposit, bool _canMint) returns()
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistryTransactorSession) AddToken(_token common.Address, _logo string, _oracle common.Address, _priceDecimals uint8, _isActive bool, _canDeposit bool, _canMint bool) (*types.Transaction, error) {
	return _DefiFMintTokenRegistry.Contract.AddToken(&_DefiFMintTokenRegistry.TransactOpts, _token, _logo, _oracle, _priceDecimals, _isActive, _canDeposit, _canMint)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner) returns()
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistryTransactor) Initialize(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _DefiFMintTokenRegistry.contract.Transact(opts, "initialize", owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner) returns()
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistrySession) Initialize(owner common.Address) (*types.Transaction, error) {
	return _DefiFMintTokenRegistry.Contract.Initialize(&_DefiFMintTokenRegistry.TransactOpts, owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner) returns()
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistryTransactorSession) Initialize(owner common.Address) (*types.Transaction, error) {
	return _DefiFMintTokenRegistry.Contract.Initialize(&_DefiFMintTokenRegistry.TransactOpts, owner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefiFMintTokenRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _DefiFMintTokenRegistry.Contract.RenounceOwnership(&_DefiFMintTokenRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DefiFMintTokenRegistry.Contract.RenounceOwnership(&_DefiFMintTokenRegistry.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DefiFMintTokenRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DefiFMintTokenRegistry.Contract.TransferOwnership(&_DefiFMintTokenRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DefiFMintTokenRegistry.Contract.TransferOwnership(&_DefiFMintTokenRegistry.TransactOpts, newOwner)
}

// UpdateToken is a paid mutator transaction binding the contract method 0xd75c1c4c.
//
// Solidity: function updateToken(address _token, string _logo, address _oracle, uint8 _priceDecimals, bool _isActive, bool _canDeposit, bool _canMint) returns()
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistryTransactor) UpdateToken(opts *bind.TransactOpts, _token common.Address, _logo string, _oracle common.Address, _priceDecimals uint8, _isActive bool, _canDeposit bool, _canMint bool) (*types.Transaction, error) {
	return _DefiFMintTokenRegistry.contract.Transact(opts, "updateToken", _token, _logo, _oracle, _priceDecimals, _isActive, _canDeposit, _canMint)
}

// UpdateToken is a paid mutator transaction binding the contract method 0xd75c1c4c.
//
// Solidity: function updateToken(address _token, string _logo, address _oracle, uint8 _priceDecimals, bool _isActive, bool _canDeposit, bool _canMint) returns()
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistrySession) UpdateToken(_token common.Address, _logo string, _oracle common.Address, _priceDecimals uint8, _isActive bool, _canDeposit bool, _canMint bool) (*types.Transaction, error) {
	return _DefiFMintTokenRegistry.Contract.UpdateToken(&_DefiFMintTokenRegistry.TransactOpts, _token, _logo, _oracle, _priceDecimals, _isActive, _canDeposit, _canMint)
}

// UpdateToken is a paid mutator transaction binding the contract method 0xd75c1c4c.
//
// Solidity: function updateToken(address _token, string _logo, address _oracle, uint8 _priceDecimals, bool _isActive, bool _canDeposit, bool _canMint) returns()
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistryTransactorSession) UpdateToken(_token common.Address, _logo string, _oracle common.Address, _priceDecimals uint8, _isActive bool, _canDeposit bool, _canMint bool) (*types.Transaction, error) {
	return _DefiFMintTokenRegistry.Contract.UpdateToken(&_DefiFMintTokenRegistry.TransactOpts, _token, _logo, _oracle, _priceDecimals, _isActive, _canDeposit, _canMint)
}

// DefiFMintTokenRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DefiFMintTokenRegistry contract.
type DefiFMintTokenRegistryOwnershipTransferredIterator struct {
	Event *DefiFMintTokenRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DefiFMintTokenRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefiFMintTokenRegistryOwnershipTransferred)
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
		it.Event = new(DefiFMintTokenRegistryOwnershipTransferred)
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
func (it *DefiFMintTokenRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefiFMintTokenRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefiFMintTokenRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the DefiFMintTokenRegistry contract.
type DefiFMintTokenRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DefiFMintTokenRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DefiFMintTokenRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DefiFMintTokenRegistryOwnershipTransferredIterator{contract: _DefiFMintTokenRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DefiFMintTokenRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DefiFMintTokenRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefiFMintTokenRegistryOwnershipTransferred)
				if err := _DefiFMintTokenRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*DefiFMintTokenRegistryOwnershipTransferred, error) {
	event := new(DefiFMintTokenRegistryOwnershipTransferred)
	if err := _DefiFMintTokenRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DefiFMintTokenRegistryTokenAddedIterator is returned from FilterTokenAdded and is used to iterate over the raw logs and unpacked data for TokenAdded events raised by the DefiFMintTokenRegistry contract.
type DefiFMintTokenRegistryTokenAddedIterator struct {
	Event *DefiFMintTokenRegistryTokenAdded // Event containing the contract specifics and raw log

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
func (it *DefiFMintTokenRegistryTokenAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefiFMintTokenRegistryTokenAdded)
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
		it.Event = new(DefiFMintTokenRegistryTokenAdded)
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
func (it *DefiFMintTokenRegistryTokenAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefiFMintTokenRegistryTokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefiFMintTokenRegistryTokenAdded represents a TokenAdded event raised by the DefiFMintTokenRegistry contract.
type DefiFMintTokenRegistryTokenAdded struct {
	Token common.Address
	Name  string
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenAdded is a free log retrieval operation binding the contract event 0x4af7419360b60cfcf01ac8a5c1487814e666a0af47877d73e82476772ac9150f.
//
// Solidity: event TokenAdded(address indexed token, string name, uint256 index)
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistryFilterer) FilterTokenAdded(opts *bind.FilterOpts, token []common.Address) (*DefiFMintTokenRegistryTokenAddedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _DefiFMintTokenRegistry.contract.FilterLogs(opts, "TokenAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return &DefiFMintTokenRegistryTokenAddedIterator{contract: _DefiFMintTokenRegistry.contract, event: "TokenAdded", logs: logs, sub: sub}, nil
}

// WatchTokenAdded is a free log subscription operation binding the contract event 0x4af7419360b60cfcf01ac8a5c1487814e666a0af47877d73e82476772ac9150f.
//
// Solidity: event TokenAdded(address indexed token, string name, uint256 index)
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistryFilterer) WatchTokenAdded(opts *bind.WatchOpts, sink chan<- *DefiFMintTokenRegistryTokenAdded, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _DefiFMintTokenRegistry.contract.WatchLogs(opts, "TokenAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefiFMintTokenRegistryTokenAdded)
				if err := _DefiFMintTokenRegistry.contract.UnpackLog(event, "TokenAdded", log); err != nil {
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

// ParseTokenAdded is a log parse operation binding the contract event 0x4af7419360b60cfcf01ac8a5c1487814e666a0af47877d73e82476772ac9150f.
//
// Solidity: event TokenAdded(address indexed token, string name, uint256 index)
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistryFilterer) ParseTokenAdded(log types.Log) (*DefiFMintTokenRegistryTokenAdded, error) {
	event := new(DefiFMintTokenRegistryTokenAdded)
	if err := _DefiFMintTokenRegistry.contract.UnpackLog(event, "TokenAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DefiFMintTokenRegistryTokenUpdatedIterator is returned from FilterTokenUpdated and is used to iterate over the raw logs and unpacked data for TokenUpdated events raised by the DefiFMintTokenRegistry contract.
type DefiFMintTokenRegistryTokenUpdatedIterator struct {
	Event *DefiFMintTokenRegistryTokenUpdated // Event containing the contract specifics and raw log

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
func (it *DefiFMintTokenRegistryTokenUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefiFMintTokenRegistryTokenUpdated)
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
		it.Event = new(DefiFMintTokenRegistryTokenUpdated)
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
func (it *DefiFMintTokenRegistryTokenUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefiFMintTokenRegistryTokenUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefiFMintTokenRegistryTokenUpdated represents a TokenUpdated event raised by the DefiFMintTokenRegistry contract.
type DefiFMintTokenRegistryTokenUpdated struct {
	Token common.Address
	Name  string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenUpdated is a free log retrieval operation binding the contract event 0x7dfa4f44638df9ca9c035c37f4954edb0383135db7751b81208a86345775a159.
//
// Solidity: event TokenUpdated(address indexed token, string name)
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistryFilterer) FilterTokenUpdated(opts *bind.FilterOpts, token []common.Address) (*DefiFMintTokenRegistryTokenUpdatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _DefiFMintTokenRegistry.contract.FilterLogs(opts, "TokenUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &DefiFMintTokenRegistryTokenUpdatedIterator{contract: _DefiFMintTokenRegistry.contract, event: "TokenUpdated", logs: logs, sub: sub}, nil
}

// WatchTokenUpdated is a free log subscription operation binding the contract event 0x7dfa4f44638df9ca9c035c37f4954edb0383135db7751b81208a86345775a159.
//
// Solidity: event TokenUpdated(address indexed token, string name)
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistryFilterer) WatchTokenUpdated(opts *bind.WatchOpts, sink chan<- *DefiFMintTokenRegistryTokenUpdated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _DefiFMintTokenRegistry.contract.WatchLogs(opts, "TokenUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefiFMintTokenRegistryTokenUpdated)
				if err := _DefiFMintTokenRegistry.contract.UnpackLog(event, "TokenUpdated", log); err != nil {
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

// ParseTokenUpdated is a log parse operation binding the contract event 0x7dfa4f44638df9ca9c035c37f4954edb0383135db7751b81208a86345775a159.
//
// Solidity: event TokenUpdated(address indexed token, string name)
func (_DefiFMintTokenRegistry *DefiFMintTokenRegistryFilterer) ParseTokenUpdated(log types.Log) (*DefiFMintTokenRegistryTokenUpdated, error) {
	event := new(DefiFMintTokenRegistryTokenUpdated)
	if err := _DefiFMintTokenRegistry.contract.UnpackLog(event, "TokenUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
