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

// SfcTokenizerMetaData contains all meta data concerning the SfcTokenizer contract.
var SfcTokenizerMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"_updateSFTMTokenAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"}],\"name\":\"allowedToWithdrawStake\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"}],\"name\":\"checkAllowedToWithdrawStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"redeemSFTM\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"outstandingSFTM\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"sFTMTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toStakerID\",\"type\":\"uint256\"}],\"name\":\"mintSFTM\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]",
}

// SfcTokenizerABI is the input ABI used to generate the binding from.
// Deprecated: Use SfcTokenizerMetaData.ABI instead.
var SfcTokenizerABI = SfcTokenizerMetaData.ABI

// SfcTokenizer is an auto generated Go binding around an Ethereum contract.
type SfcTokenizer struct {
	SfcTokenizerCaller     // Read-only binding to the contract
	SfcTokenizerTransactor // Write-only binding to the contract
	SfcTokenizerFilterer   // Log filterer for contract events
}

// SfcTokenizerCaller is an auto generated read-only Go binding around an Ethereum contract.
type SfcTokenizerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SfcTokenizerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SfcTokenizerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SfcTokenizerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SfcTokenizerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SfcTokenizerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SfcTokenizerSession struct {
	Contract     *SfcTokenizer     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SfcTokenizerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SfcTokenizerCallerSession struct {
	Contract *SfcTokenizerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// SfcTokenizerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SfcTokenizerTransactorSession struct {
	Contract     *SfcTokenizerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// SfcTokenizerRaw is an auto generated low-level Go binding around an Ethereum contract.
type SfcTokenizerRaw struct {
	Contract *SfcTokenizer // Generic contract binding to access the raw methods on
}

// SfcTokenizerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SfcTokenizerCallerRaw struct {
	Contract *SfcTokenizerCaller // Generic read-only contract binding to access the raw methods on
}

// SfcTokenizerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SfcTokenizerTransactorRaw struct {
	Contract *SfcTokenizerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSfcTokenizer creates a new instance of SfcTokenizer, bound to a specific deployed contract.
func NewSfcTokenizer(address common.Address, backend bind.ContractBackend) (*SfcTokenizer, error) {
	contract, err := bindSfcTokenizer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SfcTokenizer{SfcTokenizerCaller: SfcTokenizerCaller{contract: contract}, SfcTokenizerTransactor: SfcTokenizerTransactor{contract: contract}, SfcTokenizerFilterer: SfcTokenizerFilterer{contract: contract}}, nil
}

// NewSfcTokenizerCaller creates a new read-only instance of SfcTokenizer, bound to a specific deployed contract.
func NewSfcTokenizerCaller(address common.Address, caller bind.ContractCaller) (*SfcTokenizerCaller, error) {
	contract, err := bindSfcTokenizer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SfcTokenizerCaller{contract: contract}, nil
}

// NewSfcTokenizerTransactor creates a new write-only instance of SfcTokenizer, bound to a specific deployed contract.
func NewSfcTokenizerTransactor(address common.Address, transactor bind.ContractTransactor) (*SfcTokenizerTransactor, error) {
	contract, err := bindSfcTokenizer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SfcTokenizerTransactor{contract: contract}, nil
}

// NewSfcTokenizerFilterer creates a new log filterer instance of SfcTokenizer, bound to a specific deployed contract.
func NewSfcTokenizerFilterer(address common.Address, filterer bind.ContractFilterer) (*SfcTokenizerFilterer, error) {
	contract, err := bindSfcTokenizer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SfcTokenizerFilterer{contract: contract}, nil
}

// bindSfcTokenizer binds a generic wrapper to an already deployed contract.
func bindSfcTokenizer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SfcTokenizerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SfcTokenizer *SfcTokenizerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SfcTokenizer.Contract.SfcTokenizerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SfcTokenizer *SfcTokenizerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SfcTokenizer.Contract.SfcTokenizerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SfcTokenizer *SfcTokenizerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SfcTokenizer.Contract.SfcTokenizerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SfcTokenizer *SfcTokenizerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SfcTokenizer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SfcTokenizer *SfcTokenizerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SfcTokenizer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SfcTokenizer *SfcTokenizerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SfcTokenizer.Contract.contract.Transact(opts, method, params...)
}

// AllowedToWithdrawStake is a free data retrieval call binding the contract method 0x21d585c3.
//
// Solidity: function allowedToWithdrawStake(address sender, uint256 stakerID) view returns(bool)
func (_SfcTokenizer *SfcTokenizerCaller) AllowedToWithdrawStake(opts *bind.CallOpts, sender common.Address, stakerID *big.Int) (bool, error) {
	var out []interface{}
	err := _SfcTokenizer.contract.Call(opts, &out, "allowedToWithdrawStake", sender, stakerID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedToWithdrawStake is a free data retrieval call binding the contract method 0x21d585c3.
//
// Solidity: function allowedToWithdrawStake(address sender, uint256 stakerID) view returns(bool)
func (_SfcTokenizer *SfcTokenizerSession) AllowedToWithdrawStake(sender common.Address, stakerID *big.Int) (bool, error) {
	return _SfcTokenizer.Contract.AllowedToWithdrawStake(&_SfcTokenizer.CallOpts, sender, stakerID)
}

// AllowedToWithdrawStake is a free data retrieval call binding the contract method 0x21d585c3.
//
// Solidity: function allowedToWithdrawStake(address sender, uint256 stakerID) view returns(bool)
func (_SfcTokenizer *SfcTokenizerCallerSession) AllowedToWithdrawStake(sender common.Address, stakerID *big.Int) (bool, error) {
	return _SfcTokenizer.Contract.AllowedToWithdrawStake(&_SfcTokenizer.CallOpts, sender, stakerID)
}

// CheckAllowedToWithdrawStake is a free data retrieval call binding the contract method 0x41326a96.
//
// Solidity: function checkAllowedToWithdrawStake(address sender, uint256 stakerID) view returns()
func (_SfcTokenizer *SfcTokenizerCaller) CheckAllowedToWithdrawStake(opts *bind.CallOpts, sender common.Address, stakerID *big.Int) error {
	var out []interface{}
	err := _SfcTokenizer.contract.Call(opts, &out, "checkAllowedToWithdrawStake", sender, stakerID)

	if err != nil {
		return err
	}

	return err

}

// CheckAllowedToWithdrawStake is a free data retrieval call binding the contract method 0x41326a96.
//
// Solidity: function checkAllowedToWithdrawStake(address sender, uint256 stakerID) view returns()
func (_SfcTokenizer *SfcTokenizerSession) CheckAllowedToWithdrawStake(sender common.Address, stakerID *big.Int) error {
	return _SfcTokenizer.Contract.CheckAllowedToWithdrawStake(&_SfcTokenizer.CallOpts, sender, stakerID)
}

// CheckAllowedToWithdrawStake is a free data retrieval call binding the contract method 0x41326a96.
//
// Solidity: function checkAllowedToWithdrawStake(address sender, uint256 stakerID) view returns()
func (_SfcTokenizer *SfcTokenizerCallerSession) CheckAllowedToWithdrawStake(sender common.Address, stakerID *big.Int) error {
	return _SfcTokenizer.Contract.CheckAllowedToWithdrawStake(&_SfcTokenizer.CallOpts, sender, stakerID)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_SfcTokenizer *SfcTokenizerCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SfcTokenizer.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_SfcTokenizer *SfcTokenizerSession) IsOwner() (bool, error) {
	return _SfcTokenizer.Contract.IsOwner(&_SfcTokenizer.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_SfcTokenizer *SfcTokenizerCallerSession) IsOwner() (bool, error) {
	return _SfcTokenizer.Contract.IsOwner(&_SfcTokenizer.CallOpts)
}

// OutstandingSFTM is a free data retrieval call binding the contract method 0x5b04f4c8.
//
// Solidity: function outstandingSFTM(address , uint256 ) view returns(uint256)
func (_SfcTokenizer *SfcTokenizerCaller) OutstandingSFTM(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SfcTokenizer.contract.Call(opts, &out, "outstandingSFTM", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OutstandingSFTM is a free data retrieval call binding the contract method 0x5b04f4c8.
//
// Solidity: function outstandingSFTM(address , uint256 ) view returns(uint256)
func (_SfcTokenizer *SfcTokenizerSession) OutstandingSFTM(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _SfcTokenizer.Contract.OutstandingSFTM(&_SfcTokenizer.CallOpts, arg0, arg1)
}

// OutstandingSFTM is a free data retrieval call binding the contract method 0x5b04f4c8.
//
// Solidity: function outstandingSFTM(address , uint256 ) view returns(uint256)
func (_SfcTokenizer *SfcTokenizerCallerSession) OutstandingSFTM(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _SfcTokenizer.Contract.OutstandingSFTM(&_SfcTokenizer.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SfcTokenizer *SfcTokenizerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SfcTokenizer.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SfcTokenizer *SfcTokenizerSession) Owner() (common.Address, error) {
	return _SfcTokenizer.Contract.Owner(&_SfcTokenizer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SfcTokenizer *SfcTokenizerCallerSession) Owner() (common.Address, error) {
	return _SfcTokenizer.Contract.Owner(&_SfcTokenizer.CallOpts)
}

// SFTMTokenAddress is a free data retrieval call binding the contract method 0x693c5e24.
//
// Solidity: function sFTMTokenAddress() view returns(address)
func (_SfcTokenizer *SfcTokenizerCaller) SFTMTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SfcTokenizer.contract.Call(opts, &out, "sFTMTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SFTMTokenAddress is a free data retrieval call binding the contract method 0x693c5e24.
//
// Solidity: function sFTMTokenAddress() view returns(address)
func (_SfcTokenizer *SfcTokenizerSession) SFTMTokenAddress() (common.Address, error) {
	return _SfcTokenizer.Contract.SFTMTokenAddress(&_SfcTokenizer.CallOpts)
}

// SFTMTokenAddress is a free data retrieval call binding the contract method 0x693c5e24.
//
// Solidity: function sFTMTokenAddress() view returns(address)
func (_SfcTokenizer *SfcTokenizerCallerSession) SFTMTokenAddress() (common.Address, error) {
	return _SfcTokenizer.Contract.SFTMTokenAddress(&_SfcTokenizer.CallOpts)
}

// UpdateSFTMTokenAddress is a paid mutator transaction binding the contract method 0x06f03e22.
//
// Solidity: function _updateSFTMTokenAddress(address addr) returns()
func (_SfcTokenizer *SfcTokenizerTransactor) UpdateSFTMTokenAddress(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _SfcTokenizer.contract.Transact(opts, "_updateSFTMTokenAddress", addr)
}

// UpdateSFTMTokenAddress is a paid mutator transaction binding the contract method 0x06f03e22.
//
// Solidity: function _updateSFTMTokenAddress(address addr) returns()
func (_SfcTokenizer *SfcTokenizerSession) UpdateSFTMTokenAddress(addr common.Address) (*types.Transaction, error) {
	return _SfcTokenizer.Contract.UpdateSFTMTokenAddress(&_SfcTokenizer.TransactOpts, addr)
}

// UpdateSFTMTokenAddress is a paid mutator transaction binding the contract method 0x06f03e22.
//
// Solidity: function _updateSFTMTokenAddress(address addr) returns()
func (_SfcTokenizer *SfcTokenizerTransactorSession) UpdateSFTMTokenAddress(addr common.Address) (*types.Transaction, error) {
	return _SfcTokenizer.Contract.UpdateSFTMTokenAddress(&_SfcTokenizer.TransactOpts, addr)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_SfcTokenizer *SfcTokenizerTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SfcTokenizer.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_SfcTokenizer *SfcTokenizerSession) Initialize() (*types.Transaction, error) {
	return _SfcTokenizer.Contract.Initialize(&_SfcTokenizer.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_SfcTokenizer *SfcTokenizerTransactorSession) Initialize() (*types.Transaction, error) {
	return _SfcTokenizer.Contract.Initialize(&_SfcTokenizer.TransactOpts)
}

// MintSFTM is a paid mutator transaction binding the contract method 0xca2d6fd8.
//
// Solidity: function mintSFTM(uint256 toStakerID) returns()
func (_SfcTokenizer *SfcTokenizerTransactor) MintSFTM(opts *bind.TransactOpts, toStakerID *big.Int) (*types.Transaction, error) {
	return _SfcTokenizer.contract.Transact(opts, "mintSFTM", toStakerID)
}

// MintSFTM is a paid mutator transaction binding the contract method 0xca2d6fd8.
//
// Solidity: function mintSFTM(uint256 toStakerID) returns()
func (_SfcTokenizer *SfcTokenizerSession) MintSFTM(toStakerID *big.Int) (*types.Transaction, error) {
	return _SfcTokenizer.Contract.MintSFTM(&_SfcTokenizer.TransactOpts, toStakerID)
}

// MintSFTM is a paid mutator transaction binding the contract method 0xca2d6fd8.
//
// Solidity: function mintSFTM(uint256 toStakerID) returns()
func (_SfcTokenizer *SfcTokenizerTransactorSession) MintSFTM(toStakerID *big.Int) (*types.Transaction, error) {
	return _SfcTokenizer.Contract.MintSFTM(&_SfcTokenizer.TransactOpts, toStakerID)
}

// RedeemSFTM is a paid mutator transaction binding the contract method 0x4a37ec9d.
//
// Solidity: function redeemSFTM(uint256 stakerID, uint256 amount) returns()
func (_SfcTokenizer *SfcTokenizerTransactor) RedeemSFTM(opts *bind.TransactOpts, stakerID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _SfcTokenizer.contract.Transact(opts, "redeemSFTM", stakerID, amount)
}

// RedeemSFTM is a paid mutator transaction binding the contract method 0x4a37ec9d.
//
// Solidity: function redeemSFTM(uint256 stakerID, uint256 amount) returns()
func (_SfcTokenizer *SfcTokenizerSession) RedeemSFTM(stakerID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _SfcTokenizer.Contract.RedeemSFTM(&_SfcTokenizer.TransactOpts, stakerID, amount)
}

// RedeemSFTM is a paid mutator transaction binding the contract method 0x4a37ec9d.
//
// Solidity: function redeemSFTM(uint256 stakerID, uint256 amount) returns()
func (_SfcTokenizer *SfcTokenizerTransactorSession) RedeemSFTM(stakerID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _SfcTokenizer.Contract.RedeemSFTM(&_SfcTokenizer.TransactOpts, stakerID, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SfcTokenizer *SfcTokenizerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SfcTokenizer.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SfcTokenizer *SfcTokenizerSession) RenounceOwnership() (*types.Transaction, error) {
	return _SfcTokenizer.Contract.RenounceOwnership(&_SfcTokenizer.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SfcTokenizer *SfcTokenizerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SfcTokenizer.Contract.RenounceOwnership(&_SfcTokenizer.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SfcTokenizer *SfcTokenizerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SfcTokenizer.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SfcTokenizer *SfcTokenizerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SfcTokenizer.Contract.TransferOwnership(&_SfcTokenizer.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SfcTokenizer *SfcTokenizerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SfcTokenizer.Contract.TransferOwnership(&_SfcTokenizer.TransactOpts, newOwner)
}

// SfcTokenizerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SfcTokenizer contract.
type SfcTokenizerOwnershipTransferredIterator struct {
	Event *SfcTokenizerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SfcTokenizerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcTokenizerOwnershipTransferred)
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
		it.Event = new(SfcTokenizerOwnershipTransferred)
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
func (it *SfcTokenizerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcTokenizerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcTokenizerOwnershipTransferred represents a OwnershipTransferred event raised by the SfcTokenizer contract.
type SfcTokenizerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SfcTokenizer *SfcTokenizerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SfcTokenizerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SfcTokenizer.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SfcTokenizerOwnershipTransferredIterator{contract: _SfcTokenizer.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SfcTokenizer *SfcTokenizerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SfcTokenizerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SfcTokenizer.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcTokenizerOwnershipTransferred)
				if err := _SfcTokenizer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SfcTokenizer *SfcTokenizerFilterer) ParseOwnershipTransferred(log types.Log) (*SfcTokenizerOwnershipTransferred, error) {
	event := new(SfcTokenizerOwnershipTransferred)
	if err := _SfcTokenizer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
