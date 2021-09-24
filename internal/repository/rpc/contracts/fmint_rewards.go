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

// FMintRewardsDistributionMetaData contains all meta data concerning the FMintRewardsDistribution contract.
var FMintRewardsDistributionMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardPerYear\",\"type\":\"uint256\"}],\"name\":\"RateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardPaid\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_DEBT_EXCEEDED\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_DEPOSIT_PROHIBITED\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_LOW_ALLOWANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_LOW_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_LOW_BALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_LOW_COLLATERAL_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_MINTING_PROHIBITED\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_NOT_AUTHORIZED\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_NO_COLLATERAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_NO_ERROR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_NO_REWARD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_NO_VALUE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_REWARDS_DEPLETED\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_REWARDS_EARLY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_REWARDS_NONE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_REWARD_CLAIM_REJECTED\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_ZERO_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"addressProvider\",\"outputs\":[{\"internalType\":\"contractIFantomMintAddressProvider\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastRewardPush\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minRewardPushInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"mustRewardClaim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardApplicableUntil\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"rewardClaim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"rewardEarned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardEpochEnds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardEpochLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardLastPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardPerSecond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardPerTokenDecimalsCorrection\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardPerTokenPaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardStash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"rewardUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"rewardUpdateGlobal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardUpdated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressProvider\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"mustRewardPush\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"rewardPush\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_perSecond\",\"type\":\"uint256\"}],\"name\":\"rewardUpdateRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"rewardCleanup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"principalBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"principalBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardTokenAddress\",\"outputs\":[{\"internalType\":\"contractERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"rewardCanClaim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"rewardIsEligible\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// FMintRewardsDistributionABI is the input ABI used to generate the binding from.
// Deprecated: Use FMintRewardsDistributionMetaData.ABI instead.
var FMintRewardsDistributionABI = FMintRewardsDistributionMetaData.ABI

// FMintRewardsDistribution is an auto generated Go binding around an Ethereum contract.
type FMintRewardsDistribution struct {
	FMintRewardsDistributionCaller     // Read-only binding to the contract
	FMintRewardsDistributionTransactor // Write-only binding to the contract
	FMintRewardsDistributionFilterer   // Log filterer for contract events
}

// FMintRewardsDistributionCaller is an auto generated read-only Go binding around an Ethereum contract.
type FMintRewardsDistributionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FMintRewardsDistributionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FMintRewardsDistributionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FMintRewardsDistributionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FMintRewardsDistributionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FMintRewardsDistributionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FMintRewardsDistributionSession struct {
	Contract     *FMintRewardsDistribution // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// FMintRewardsDistributionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FMintRewardsDistributionCallerSession struct {
	Contract *FMintRewardsDistributionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// FMintRewardsDistributionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FMintRewardsDistributionTransactorSession struct {
	Contract     *FMintRewardsDistributionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// FMintRewardsDistributionRaw is an auto generated low-level Go binding around an Ethereum contract.
type FMintRewardsDistributionRaw struct {
	Contract *FMintRewardsDistribution // Generic contract binding to access the raw methods on
}

// FMintRewardsDistributionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FMintRewardsDistributionCallerRaw struct {
	Contract *FMintRewardsDistributionCaller // Generic read-only contract binding to access the raw methods on
}

// FMintRewardsDistributionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FMintRewardsDistributionTransactorRaw struct {
	Contract *FMintRewardsDistributionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFMintRewardsDistribution creates a new instance of FMintRewardsDistribution, bound to a specific deployed contract.
func NewFMintRewardsDistribution(address common.Address, backend bind.ContractBackend) (*FMintRewardsDistribution, error) {
	contract, err := bindFMintRewardsDistribution(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FMintRewardsDistribution{FMintRewardsDistributionCaller: FMintRewardsDistributionCaller{contract: contract}, FMintRewardsDistributionTransactor: FMintRewardsDistributionTransactor{contract: contract}, FMintRewardsDistributionFilterer: FMintRewardsDistributionFilterer{contract: contract}}, nil
}

// NewFMintRewardsDistributionCaller creates a new read-only instance of FMintRewardsDistribution, bound to a specific deployed contract.
func NewFMintRewardsDistributionCaller(address common.Address, caller bind.ContractCaller) (*FMintRewardsDistributionCaller, error) {
	contract, err := bindFMintRewardsDistribution(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FMintRewardsDistributionCaller{contract: contract}, nil
}

// NewFMintRewardsDistributionTransactor creates a new write-only instance of FMintRewardsDistribution, bound to a specific deployed contract.
func NewFMintRewardsDistributionTransactor(address common.Address, transactor bind.ContractTransactor) (*FMintRewardsDistributionTransactor, error) {
	contract, err := bindFMintRewardsDistribution(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FMintRewardsDistributionTransactor{contract: contract}, nil
}

// NewFMintRewardsDistributionFilterer creates a new log filterer instance of FMintRewardsDistribution, bound to a specific deployed contract.
func NewFMintRewardsDistributionFilterer(address common.Address, filterer bind.ContractFilterer) (*FMintRewardsDistributionFilterer, error) {
	contract, err := bindFMintRewardsDistribution(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FMintRewardsDistributionFilterer{contract: contract}, nil
}

// bindFMintRewardsDistribution binds a generic wrapper to an already deployed contract.
func bindFMintRewardsDistribution(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FMintRewardsDistributionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FMintRewardsDistribution *FMintRewardsDistributionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FMintRewardsDistribution.Contract.FMintRewardsDistributionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FMintRewardsDistribution *FMintRewardsDistributionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FMintRewardsDistribution.Contract.FMintRewardsDistributionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FMintRewardsDistribution *FMintRewardsDistributionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FMintRewardsDistribution.Contract.FMintRewardsDistributionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FMintRewardsDistribution *FMintRewardsDistributionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FMintRewardsDistribution.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FMintRewardsDistribution *FMintRewardsDistributionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FMintRewardsDistribution.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FMintRewardsDistribution *FMintRewardsDistributionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FMintRewardsDistribution.Contract.contract.Transact(opts, method, params...)
}

// ERRDEBTEXCEEDED is a free data retrieval call binding the contract method 0x372ce3df.
//
// Solidity: function ERR_DEBT_EXCEEDED() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCaller) ERRDEBTEXCEEDED(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FMintRewardsDistribution.contract.Call(opts, &out, "ERR_DEBT_EXCEEDED")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ERRDEBTEXCEEDED is a free data retrieval call binding the contract method 0x372ce3df.
//
// Solidity: function ERR_DEBT_EXCEEDED() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) ERRDEBTEXCEEDED() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.ERRDEBTEXCEEDED(&_FMintRewardsDistribution.CallOpts)
}

// ERRDEBTEXCEEDED is a free data retrieval call binding the contract method 0x372ce3df.
//
// Solidity: function ERR_DEBT_EXCEEDED() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCallerSession) ERRDEBTEXCEEDED() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.ERRDEBTEXCEEDED(&_FMintRewardsDistribution.CallOpts)
}

// ERRDEPOSITPROHIBITED is a free data retrieval call binding the contract method 0x2bfcc373.
//
// Solidity: function ERR_DEPOSIT_PROHIBITED() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCaller) ERRDEPOSITPROHIBITED(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FMintRewardsDistribution.contract.Call(opts, &out, "ERR_DEPOSIT_PROHIBITED")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ERRDEPOSITPROHIBITED is a free data retrieval call binding the contract method 0x2bfcc373.
//
// Solidity: function ERR_DEPOSIT_PROHIBITED() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) ERRDEPOSITPROHIBITED() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.ERRDEPOSITPROHIBITED(&_FMintRewardsDistribution.CallOpts)
}

// ERRDEPOSITPROHIBITED is a free data retrieval call binding the contract method 0x2bfcc373.
//
// Solidity: function ERR_DEPOSIT_PROHIBITED() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCallerSession) ERRDEPOSITPROHIBITED() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.ERRDEPOSITPROHIBITED(&_FMintRewardsDistribution.CallOpts)
}

// ERRLOWALLOWANCE is a free data retrieval call binding the contract method 0x911fc3f1.
//
// Solidity: function ERR_LOW_ALLOWANCE() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCaller) ERRLOWALLOWANCE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FMintRewardsDistribution.contract.Call(opts, &out, "ERR_LOW_ALLOWANCE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ERRLOWALLOWANCE is a free data retrieval call binding the contract method 0x911fc3f1.
//
// Solidity: function ERR_LOW_ALLOWANCE() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) ERRLOWALLOWANCE() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.ERRLOWALLOWANCE(&_FMintRewardsDistribution.CallOpts)
}

// ERRLOWALLOWANCE is a free data retrieval call binding the contract method 0x911fc3f1.
//
// Solidity: function ERR_LOW_ALLOWANCE() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCallerSession) ERRLOWALLOWANCE() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.ERRLOWALLOWANCE(&_FMintRewardsDistribution.CallOpts)
}

// ERRLOWAMOUNT is a free data retrieval call binding the contract method 0xc7ea4889.
//
// Solidity: function ERR_LOW_AMOUNT() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCaller) ERRLOWAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FMintRewardsDistribution.contract.Call(opts, &out, "ERR_LOW_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ERRLOWAMOUNT is a free data retrieval call binding the contract method 0xc7ea4889.
//
// Solidity: function ERR_LOW_AMOUNT() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) ERRLOWAMOUNT() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.ERRLOWAMOUNT(&_FMintRewardsDistribution.CallOpts)
}

// ERRLOWAMOUNT is a free data retrieval call binding the contract method 0xc7ea4889.
//
// Solidity: function ERR_LOW_AMOUNT() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCallerSession) ERRLOWAMOUNT() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.ERRLOWAMOUNT(&_FMintRewardsDistribution.CallOpts)
}

// ERRLOWBALANCE is a free data retrieval call binding the contract method 0x1ac919b0.
//
// Solidity: function ERR_LOW_BALANCE() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCaller) ERRLOWBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FMintRewardsDistribution.contract.Call(opts, &out, "ERR_LOW_BALANCE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ERRLOWBALANCE is a free data retrieval call binding the contract method 0x1ac919b0.
//
// Solidity: function ERR_LOW_BALANCE() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) ERRLOWBALANCE() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.ERRLOWBALANCE(&_FMintRewardsDistribution.CallOpts)
}

// ERRLOWBALANCE is a free data retrieval call binding the contract method 0x1ac919b0.
//
// Solidity: function ERR_LOW_BALANCE() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCallerSession) ERRLOWBALANCE() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.ERRLOWBALANCE(&_FMintRewardsDistribution.CallOpts)
}

// ERRLOWCOLLATERALRATIO is a free data retrieval call binding the contract method 0x04b62f29.
//
// Solidity: function ERR_LOW_COLLATERAL_RATIO() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCaller) ERRLOWCOLLATERALRATIO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FMintRewardsDistribution.contract.Call(opts, &out, "ERR_LOW_COLLATERAL_RATIO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ERRLOWCOLLATERALRATIO is a free data retrieval call binding the contract method 0x04b62f29.
//
// Solidity: function ERR_LOW_COLLATERAL_RATIO() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) ERRLOWCOLLATERALRATIO() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.ERRLOWCOLLATERALRATIO(&_FMintRewardsDistribution.CallOpts)
}

// ERRLOWCOLLATERALRATIO is a free data retrieval call binding the contract method 0x04b62f29.
//
// Solidity: function ERR_LOW_COLLATERAL_RATIO() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCallerSession) ERRLOWCOLLATERALRATIO() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.ERRLOWCOLLATERALRATIO(&_FMintRewardsDistribution.CallOpts)
}

// ERRMINTINGPROHIBITED is a free data retrieval call binding the contract method 0x8c7b9980.
//
// Solidity: function ERR_MINTING_PROHIBITED() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCaller) ERRMINTINGPROHIBITED(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FMintRewardsDistribution.contract.Call(opts, &out, "ERR_MINTING_PROHIBITED")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ERRMINTINGPROHIBITED is a free data retrieval call binding the contract method 0x8c7b9980.
//
// Solidity: function ERR_MINTING_PROHIBITED() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) ERRMINTINGPROHIBITED() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.ERRMINTINGPROHIBITED(&_FMintRewardsDistribution.CallOpts)
}

// ERRMINTINGPROHIBITED is a free data retrieval call binding the contract method 0x8c7b9980.
//
// Solidity: function ERR_MINTING_PROHIBITED() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCallerSession) ERRMINTINGPROHIBITED() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.ERRMINTINGPROHIBITED(&_FMintRewardsDistribution.CallOpts)
}

// ERRNOTAUTHORIZED is a free data retrieval call binding the contract method 0xbc99d6ae.
//
// Solidity: function ERR_NOT_AUTHORIZED() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCaller) ERRNOTAUTHORIZED(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FMintRewardsDistribution.contract.Call(opts, &out, "ERR_NOT_AUTHORIZED")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ERRNOTAUTHORIZED is a free data retrieval call binding the contract method 0xbc99d6ae.
//
// Solidity: function ERR_NOT_AUTHORIZED() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) ERRNOTAUTHORIZED() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.ERRNOTAUTHORIZED(&_FMintRewardsDistribution.CallOpts)
}

// ERRNOTAUTHORIZED is a free data retrieval call binding the contract method 0xbc99d6ae.
//
// Solidity: function ERR_NOT_AUTHORIZED() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCallerSession) ERRNOTAUTHORIZED() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.ERRNOTAUTHORIZED(&_FMintRewardsDistribution.CallOpts)
}

// ERRNOCOLLATERAL is a free data retrieval call binding the contract method 0xb76361c2.
//
// Solidity: function ERR_NO_COLLATERAL() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCaller) ERRNOCOLLATERAL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FMintRewardsDistribution.contract.Call(opts, &out, "ERR_NO_COLLATERAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ERRNOCOLLATERAL is a free data retrieval call binding the contract method 0xb76361c2.
//
// Solidity: function ERR_NO_COLLATERAL() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) ERRNOCOLLATERAL() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.ERRNOCOLLATERAL(&_FMintRewardsDistribution.CallOpts)
}

// ERRNOCOLLATERAL is a free data retrieval call binding the contract method 0xb76361c2.
//
// Solidity: function ERR_NO_COLLATERAL() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCallerSession) ERRNOCOLLATERAL() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.ERRNOCOLLATERAL(&_FMintRewardsDistribution.CallOpts)
}

// ERRNOERROR is a free data retrieval call binding the contract method 0x35052d6e.
//
// Solidity: function ERR_NO_ERROR() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCaller) ERRNOERROR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FMintRewardsDistribution.contract.Call(opts, &out, "ERR_NO_ERROR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ERRNOERROR is a free data retrieval call binding the contract method 0x35052d6e.
//
// Solidity: function ERR_NO_ERROR() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) ERRNOERROR() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.ERRNOERROR(&_FMintRewardsDistribution.CallOpts)
}

// ERRNOERROR is a free data retrieval call binding the contract method 0x35052d6e.
//
// Solidity: function ERR_NO_ERROR() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCallerSession) ERRNOERROR() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.ERRNOERROR(&_FMintRewardsDistribution.CallOpts)
}

// ERRNOREWARD is a free data retrieval call binding the contract method 0x73a93af6.
//
// Solidity: function ERR_NO_REWARD() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCaller) ERRNOREWARD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FMintRewardsDistribution.contract.Call(opts, &out, "ERR_NO_REWARD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ERRNOREWARD is a free data retrieval call binding the contract method 0x73a93af6.
//
// Solidity: function ERR_NO_REWARD() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) ERRNOREWARD() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.ERRNOREWARD(&_FMintRewardsDistribution.CallOpts)
}

// ERRNOREWARD is a free data retrieval call binding the contract method 0x73a93af6.
//
// Solidity: function ERR_NO_REWARD() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCallerSession) ERRNOREWARD() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.ERRNOREWARD(&_FMintRewardsDistribution.CallOpts)
}

// ERRNOVALUE is a free data retrieval call binding the contract method 0x69d1cb27.
//
// Solidity: function ERR_NO_VALUE() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCaller) ERRNOVALUE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FMintRewardsDistribution.contract.Call(opts, &out, "ERR_NO_VALUE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ERRNOVALUE is a free data retrieval call binding the contract method 0x69d1cb27.
//
// Solidity: function ERR_NO_VALUE() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) ERRNOVALUE() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.ERRNOVALUE(&_FMintRewardsDistribution.CallOpts)
}

// ERRNOVALUE is a free data retrieval call binding the contract method 0x69d1cb27.
//
// Solidity: function ERR_NO_VALUE() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCallerSession) ERRNOVALUE() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.ERRNOVALUE(&_FMintRewardsDistribution.CallOpts)
}

// ERRREWARDSDEPLETED is a free data retrieval call binding the contract method 0x0a19dd33.
//
// Solidity: function ERR_REWARDS_DEPLETED() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCaller) ERRREWARDSDEPLETED(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FMintRewardsDistribution.contract.Call(opts, &out, "ERR_REWARDS_DEPLETED")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ERRREWARDSDEPLETED is a free data retrieval call binding the contract method 0x0a19dd33.
//
// Solidity: function ERR_REWARDS_DEPLETED() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) ERRREWARDSDEPLETED() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.ERRREWARDSDEPLETED(&_FMintRewardsDistribution.CallOpts)
}

// ERRREWARDSDEPLETED is a free data retrieval call binding the contract method 0x0a19dd33.
//
// Solidity: function ERR_REWARDS_DEPLETED() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCallerSession) ERRREWARDSDEPLETED() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.ERRREWARDSDEPLETED(&_FMintRewardsDistribution.CallOpts)
}

// ERRREWARDSEARLY is a free data retrieval call binding the contract method 0x67fc176b.
//
// Solidity: function ERR_REWARDS_EARLY() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCaller) ERRREWARDSEARLY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FMintRewardsDistribution.contract.Call(opts, &out, "ERR_REWARDS_EARLY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ERRREWARDSEARLY is a free data retrieval call binding the contract method 0x67fc176b.
//
// Solidity: function ERR_REWARDS_EARLY() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) ERRREWARDSEARLY() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.ERRREWARDSEARLY(&_FMintRewardsDistribution.CallOpts)
}

// ERRREWARDSEARLY is a free data retrieval call binding the contract method 0x67fc176b.
//
// Solidity: function ERR_REWARDS_EARLY() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCallerSession) ERRREWARDSEARLY() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.ERRREWARDSEARLY(&_FMintRewardsDistribution.CallOpts)
}

// ERRREWARDSNONE is a free data retrieval call binding the contract method 0xc7222c72.
//
// Solidity: function ERR_REWARDS_NONE() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCaller) ERRREWARDSNONE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FMintRewardsDistribution.contract.Call(opts, &out, "ERR_REWARDS_NONE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ERRREWARDSNONE is a free data retrieval call binding the contract method 0xc7222c72.
//
// Solidity: function ERR_REWARDS_NONE() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) ERRREWARDSNONE() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.ERRREWARDSNONE(&_FMintRewardsDistribution.CallOpts)
}

// ERRREWARDSNONE is a free data retrieval call binding the contract method 0xc7222c72.
//
// Solidity: function ERR_REWARDS_NONE() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCallerSession) ERRREWARDSNONE() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.ERRREWARDSNONE(&_FMintRewardsDistribution.CallOpts)
}

// ERRREWARDCLAIMREJECTED is a free data retrieval call binding the contract method 0x4846e345.
//
// Solidity: function ERR_REWARD_CLAIM_REJECTED() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCaller) ERRREWARDCLAIMREJECTED(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FMintRewardsDistribution.contract.Call(opts, &out, "ERR_REWARD_CLAIM_REJECTED")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ERRREWARDCLAIMREJECTED is a free data retrieval call binding the contract method 0x4846e345.
//
// Solidity: function ERR_REWARD_CLAIM_REJECTED() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) ERRREWARDCLAIMREJECTED() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.ERRREWARDCLAIMREJECTED(&_FMintRewardsDistribution.CallOpts)
}

// ERRREWARDCLAIMREJECTED is a free data retrieval call binding the contract method 0x4846e345.
//
// Solidity: function ERR_REWARD_CLAIM_REJECTED() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCallerSession) ERRREWARDCLAIMREJECTED() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.ERRREWARDCLAIMREJECTED(&_FMintRewardsDistribution.CallOpts)
}

// ERRZEROAMOUNT is a free data retrieval call binding the contract method 0x0aff90bb.
//
// Solidity: function ERR_ZERO_AMOUNT() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCaller) ERRZEROAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FMintRewardsDistribution.contract.Call(opts, &out, "ERR_ZERO_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ERRZEROAMOUNT is a free data retrieval call binding the contract method 0x0aff90bb.
//
// Solidity: function ERR_ZERO_AMOUNT() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) ERRZEROAMOUNT() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.ERRZEROAMOUNT(&_FMintRewardsDistribution.CallOpts)
}

// ERRZEROAMOUNT is a free data retrieval call binding the contract method 0x0aff90bb.
//
// Solidity: function ERR_ZERO_AMOUNT() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCallerSession) ERRZEROAMOUNT() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.ERRZEROAMOUNT(&_FMintRewardsDistribution.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_FMintRewardsDistribution *FMintRewardsDistributionCaller) AddressProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FMintRewardsDistribution.contract.Call(opts, &out, "addressProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) AddressProvider() (common.Address, error) {
	return _FMintRewardsDistribution.Contract.AddressProvider(&_FMintRewardsDistribution.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_FMintRewardsDistribution *FMintRewardsDistributionCallerSession) AddressProvider() (common.Address, error) {
	return _FMintRewardsDistribution.Contract.AddressProvider(&_FMintRewardsDistribution.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_FMintRewardsDistribution *FMintRewardsDistributionCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FMintRewardsDistribution.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) IsOwner() (bool, error) {
	return _FMintRewardsDistribution.Contract.IsOwner(&_FMintRewardsDistribution.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_FMintRewardsDistribution *FMintRewardsDistributionCallerSession) IsOwner() (bool, error) {
	return _FMintRewardsDistribution.Contract.IsOwner(&_FMintRewardsDistribution.CallOpts)
}

// LastRewardPush is a free data retrieval call binding the contract method 0xa664150a.
//
// Solidity: function lastRewardPush() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCaller) LastRewardPush(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FMintRewardsDistribution.contract.Call(opts, &out, "lastRewardPush")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRewardPush is a free data retrieval call binding the contract method 0xa664150a.
//
// Solidity: function lastRewardPush() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) LastRewardPush() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.LastRewardPush(&_FMintRewardsDistribution.CallOpts)
}

// LastRewardPush is a free data retrieval call binding the contract method 0xa664150a.
//
// Solidity: function lastRewardPush() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCallerSession) LastRewardPush() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.LastRewardPush(&_FMintRewardsDistribution.CallOpts)
}

// MinRewardPushInterval is a free data retrieval call binding the contract method 0xc0464d45.
//
// Solidity: function minRewardPushInterval() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCaller) MinRewardPushInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FMintRewardsDistribution.contract.Call(opts, &out, "minRewardPushInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinRewardPushInterval is a free data retrieval call binding the contract method 0xc0464d45.
//
// Solidity: function minRewardPushInterval() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) MinRewardPushInterval() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.MinRewardPushInterval(&_FMintRewardsDistribution.CallOpts)
}

// MinRewardPushInterval is a free data retrieval call binding the contract method 0xc0464d45.
//
// Solidity: function minRewardPushInterval() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCallerSession) MinRewardPushInterval() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.MinRewardPushInterval(&_FMintRewardsDistribution.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FMintRewardsDistribution *FMintRewardsDistributionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FMintRewardsDistribution.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) Owner() (common.Address, error) {
	return _FMintRewardsDistribution.Contract.Owner(&_FMintRewardsDistribution.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FMintRewardsDistribution *FMintRewardsDistributionCallerSession) Owner() (common.Address, error) {
	return _FMintRewardsDistribution.Contract.Owner(&_FMintRewardsDistribution.CallOpts)
}

// PrincipalBalance is a free data retrieval call binding the contract method 0xa83e53ac.
//
// Solidity: function principalBalance() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCaller) PrincipalBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FMintRewardsDistribution.contract.Call(opts, &out, "principalBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PrincipalBalance is a free data retrieval call binding the contract method 0xa83e53ac.
//
// Solidity: function principalBalance() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) PrincipalBalance() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.PrincipalBalance(&_FMintRewardsDistribution.CallOpts)
}

// PrincipalBalance is a free data retrieval call binding the contract method 0xa83e53ac.
//
// Solidity: function principalBalance() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCallerSession) PrincipalBalance() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.PrincipalBalance(&_FMintRewardsDistribution.CallOpts)
}

// PrincipalBalanceOf is a free data retrieval call binding the contract method 0xc634dfaa.
//
// Solidity: function principalBalanceOf(address _account) view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCaller) PrincipalBalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FMintRewardsDistribution.contract.Call(opts, &out, "principalBalanceOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PrincipalBalanceOf is a free data retrieval call binding the contract method 0xc634dfaa.
//
// Solidity: function principalBalanceOf(address _account) view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) PrincipalBalanceOf(_account common.Address) (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.PrincipalBalanceOf(&_FMintRewardsDistribution.CallOpts, _account)
}

// PrincipalBalanceOf is a free data retrieval call binding the contract method 0xc634dfaa.
//
// Solidity: function principalBalanceOf(address _account) view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCallerSession) PrincipalBalanceOf(_account common.Address) (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.PrincipalBalanceOf(&_FMintRewardsDistribution.CallOpts, _account)
}

// RewardApplicableUntil is a free data retrieval call binding the contract method 0xdb16e0b5.
//
// Solidity: function rewardApplicableUntil() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCaller) RewardApplicableUntil(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FMintRewardsDistribution.contract.Call(opts, &out, "rewardApplicableUntil")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardApplicableUntil is a free data retrieval call binding the contract method 0xdb16e0b5.
//
// Solidity: function rewardApplicableUntil() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) RewardApplicableUntil() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.RewardApplicableUntil(&_FMintRewardsDistribution.CallOpts)
}

// RewardApplicableUntil is a free data retrieval call binding the contract method 0xdb16e0b5.
//
// Solidity: function rewardApplicableUntil() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCallerSession) RewardApplicableUntil() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.RewardApplicableUntil(&_FMintRewardsDistribution.CallOpts)
}

// RewardCanClaim is a free data retrieval call binding the contract method 0xda0a0432.
//
// Solidity: function rewardCanClaim(address _account) view returns(bool)
func (_FMintRewardsDistribution *FMintRewardsDistributionCaller) RewardCanClaim(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var out []interface{}
	err := _FMintRewardsDistribution.contract.Call(opts, &out, "rewardCanClaim", _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RewardCanClaim is a free data retrieval call binding the contract method 0xda0a0432.
//
// Solidity: function rewardCanClaim(address _account) view returns(bool)
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) RewardCanClaim(_account common.Address) (bool, error) {
	return _FMintRewardsDistribution.Contract.RewardCanClaim(&_FMintRewardsDistribution.CallOpts, _account)
}

// RewardCanClaim is a free data retrieval call binding the contract method 0xda0a0432.
//
// Solidity: function rewardCanClaim(address _account) view returns(bool)
func (_FMintRewardsDistribution *FMintRewardsDistributionCallerSession) RewardCanClaim(_account common.Address) (bool, error) {
	return _FMintRewardsDistribution.Contract.RewardCanClaim(&_FMintRewardsDistribution.CallOpts, _account)
}

// RewardEarned is a free data retrieval call binding the contract method 0x16ba6bf3.
//
// Solidity: function rewardEarned(address _account) view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCaller) RewardEarned(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FMintRewardsDistribution.contract.Call(opts, &out, "rewardEarned", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardEarned is a free data retrieval call binding the contract method 0x16ba6bf3.
//
// Solidity: function rewardEarned(address _account) view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) RewardEarned(_account common.Address) (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.RewardEarned(&_FMintRewardsDistribution.CallOpts, _account)
}

// RewardEarned is a free data retrieval call binding the contract method 0x16ba6bf3.
//
// Solidity: function rewardEarned(address _account) view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCallerSession) RewardEarned(_account common.Address) (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.RewardEarned(&_FMintRewardsDistribution.CallOpts, _account)
}

// RewardEpochEnds is a free data retrieval call binding the contract method 0x3ce9b316.
//
// Solidity: function rewardEpochEnds() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCaller) RewardEpochEnds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FMintRewardsDistribution.contract.Call(opts, &out, "rewardEpochEnds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardEpochEnds is a free data retrieval call binding the contract method 0x3ce9b316.
//
// Solidity: function rewardEpochEnds() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) RewardEpochEnds() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.RewardEpochEnds(&_FMintRewardsDistribution.CallOpts)
}

// RewardEpochEnds is a free data retrieval call binding the contract method 0x3ce9b316.
//
// Solidity: function rewardEpochEnds() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCallerSession) RewardEpochEnds() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.RewardEpochEnds(&_FMintRewardsDistribution.CallOpts)
}

// RewardEpochLength is a free data retrieval call binding the contract method 0x20a0a0e9.
//
// Solidity: function rewardEpochLength() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCaller) RewardEpochLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FMintRewardsDistribution.contract.Call(opts, &out, "rewardEpochLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardEpochLength is a free data retrieval call binding the contract method 0x20a0a0e9.
//
// Solidity: function rewardEpochLength() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) RewardEpochLength() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.RewardEpochLength(&_FMintRewardsDistribution.CallOpts)
}

// RewardEpochLength is a free data retrieval call binding the contract method 0x20a0a0e9.
//
// Solidity: function rewardEpochLength() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCallerSession) RewardEpochLength() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.RewardEpochLength(&_FMintRewardsDistribution.CallOpts)
}

// RewardIsEligible is a free data retrieval call binding the contract method 0x6aee9c13.
//
// Solidity: function rewardIsEligible(address _account) view returns(bool)
func (_FMintRewardsDistribution *FMintRewardsDistributionCaller) RewardIsEligible(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var out []interface{}
	err := _FMintRewardsDistribution.contract.Call(opts, &out, "rewardIsEligible", _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RewardIsEligible is a free data retrieval call binding the contract method 0x6aee9c13.
//
// Solidity: function rewardIsEligible(address _account) view returns(bool)
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) RewardIsEligible(_account common.Address) (bool, error) {
	return _FMintRewardsDistribution.Contract.RewardIsEligible(&_FMintRewardsDistribution.CallOpts, _account)
}

// RewardIsEligible is a free data retrieval call binding the contract method 0x6aee9c13.
//
// Solidity: function rewardIsEligible(address _account) view returns(bool)
func (_FMintRewardsDistribution *FMintRewardsDistributionCallerSession) RewardIsEligible(_account common.Address) (bool, error) {
	return _FMintRewardsDistribution.Contract.RewardIsEligible(&_FMintRewardsDistribution.CallOpts, _account)
}

// RewardLastPerToken is a free data retrieval call binding the contract method 0x544bb473.
//
// Solidity: function rewardLastPerToken() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCaller) RewardLastPerToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FMintRewardsDistribution.contract.Call(opts, &out, "rewardLastPerToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardLastPerToken is a free data retrieval call binding the contract method 0x544bb473.
//
// Solidity: function rewardLastPerToken() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) RewardLastPerToken() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.RewardLastPerToken(&_FMintRewardsDistribution.CallOpts)
}

// RewardLastPerToken is a free data retrieval call binding the contract method 0x544bb473.
//
// Solidity: function rewardLastPerToken() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCallerSession) RewardLastPerToken() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.RewardLastPerToken(&_FMintRewardsDistribution.CallOpts)
}

// RewardPerSecond is a free data retrieval call binding the contract method 0x8f10369a.
//
// Solidity: function rewardPerSecond() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCaller) RewardPerSecond(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FMintRewardsDistribution.contract.Call(opts, &out, "rewardPerSecond")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerSecond is a free data retrieval call binding the contract method 0x8f10369a.
//
// Solidity: function rewardPerSecond() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) RewardPerSecond() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.RewardPerSecond(&_FMintRewardsDistribution.CallOpts)
}

// RewardPerSecond is a free data retrieval call binding the contract method 0x8f10369a.
//
// Solidity: function rewardPerSecond() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCallerSession) RewardPerSecond() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.RewardPerSecond(&_FMintRewardsDistribution.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCaller) RewardPerToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FMintRewardsDistribution.contract.Call(opts, &out, "rewardPerToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) RewardPerToken() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.RewardPerToken(&_FMintRewardsDistribution.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCallerSession) RewardPerToken() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.RewardPerToken(&_FMintRewardsDistribution.CallOpts)
}

// RewardPerTokenDecimalsCorrection is a free data retrieval call binding the contract method 0x64631d9b.
//
// Solidity: function rewardPerTokenDecimalsCorrection() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCaller) RewardPerTokenDecimalsCorrection(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FMintRewardsDistribution.contract.Call(opts, &out, "rewardPerTokenDecimalsCorrection")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerTokenDecimalsCorrection is a free data retrieval call binding the contract method 0x64631d9b.
//
// Solidity: function rewardPerTokenDecimalsCorrection() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) RewardPerTokenDecimalsCorrection() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.RewardPerTokenDecimalsCorrection(&_FMintRewardsDistribution.CallOpts)
}

// RewardPerTokenDecimalsCorrection is a free data retrieval call binding the contract method 0x64631d9b.
//
// Solidity: function rewardPerTokenDecimalsCorrection() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCallerSession) RewardPerTokenDecimalsCorrection() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.RewardPerTokenDecimalsCorrection(&_FMintRewardsDistribution.CallOpts)
}

// RewardPerTokenPaid is a free data retrieval call binding the contract method 0x653a8da1.
//
// Solidity: function rewardPerTokenPaid(address ) view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCaller) RewardPerTokenPaid(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FMintRewardsDistribution.contract.Call(opts, &out, "rewardPerTokenPaid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerTokenPaid is a free data retrieval call binding the contract method 0x653a8da1.
//
// Solidity: function rewardPerTokenPaid(address ) view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) RewardPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.RewardPerTokenPaid(&_FMintRewardsDistribution.CallOpts, arg0)
}

// RewardPerTokenPaid is a free data retrieval call binding the contract method 0x653a8da1.
//
// Solidity: function rewardPerTokenPaid(address ) view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCallerSession) RewardPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.RewardPerTokenPaid(&_FMintRewardsDistribution.CallOpts, arg0)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCaller) RewardRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FMintRewardsDistribution.contract.Call(opts, &out, "rewardRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) RewardRate() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.RewardRate(&_FMintRewardsDistribution.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCallerSession) RewardRate() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.RewardRate(&_FMintRewardsDistribution.CallOpts)
}

// RewardStash is a free data retrieval call binding the contract method 0xf2392c8d.
//
// Solidity: function rewardStash(address ) view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCaller) RewardStash(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FMintRewardsDistribution.contract.Call(opts, &out, "rewardStash", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardStash is a free data retrieval call binding the contract method 0xf2392c8d.
//
// Solidity: function rewardStash(address ) view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) RewardStash(arg0 common.Address) (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.RewardStash(&_FMintRewardsDistribution.CallOpts, arg0)
}

// RewardStash is a free data retrieval call binding the contract method 0xf2392c8d.
//
// Solidity: function rewardStash(address ) view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCallerSession) RewardStash(arg0 common.Address) (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.RewardStash(&_FMintRewardsDistribution.CallOpts, arg0)
}

// RewardTokenAddress is a free data retrieval call binding the contract method 0x125f9e33.
//
// Solidity: function rewardTokenAddress() view returns(address)
func (_FMintRewardsDistribution *FMintRewardsDistributionCaller) RewardTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FMintRewardsDistribution.contract.Call(opts, &out, "rewardTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardTokenAddress is a free data retrieval call binding the contract method 0x125f9e33.
//
// Solidity: function rewardTokenAddress() view returns(address)
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) RewardTokenAddress() (common.Address, error) {
	return _FMintRewardsDistribution.Contract.RewardTokenAddress(&_FMintRewardsDistribution.CallOpts)
}

// RewardTokenAddress is a free data retrieval call binding the contract method 0x125f9e33.
//
// Solidity: function rewardTokenAddress() view returns(address)
func (_FMintRewardsDistribution *FMintRewardsDistributionCallerSession) RewardTokenAddress() (common.Address, error) {
	return _FMintRewardsDistribution.Contract.RewardTokenAddress(&_FMintRewardsDistribution.CallOpts)
}

// RewardUpdated is a free data retrieval call binding the contract method 0x6e718e04.
//
// Solidity: function rewardUpdated() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCaller) RewardUpdated(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FMintRewardsDistribution.contract.Call(opts, &out, "rewardUpdated")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardUpdated is a free data retrieval call binding the contract method 0x6e718e04.
//
// Solidity: function rewardUpdated() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) RewardUpdated() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.RewardUpdated(&_FMintRewardsDistribution.CallOpts)
}

// RewardUpdated is a free data retrieval call binding the contract method 0x6e718e04.
//
// Solidity: function rewardUpdated() view returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionCallerSession) RewardUpdated() (*big.Int, error) {
	return _FMintRewardsDistribution.Contract.RewardUpdated(&_FMintRewardsDistribution.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address owner, address _addressProvider) returns()
func (_FMintRewardsDistribution *FMintRewardsDistributionTransactor) Initialize(opts *bind.TransactOpts, owner common.Address, _addressProvider common.Address) (*types.Transaction, error) {
	return _FMintRewardsDistribution.contract.Transact(opts, "initialize", owner, _addressProvider)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address owner, address _addressProvider) returns()
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) Initialize(owner common.Address, _addressProvider common.Address) (*types.Transaction, error) {
	return _FMintRewardsDistribution.Contract.Initialize(&_FMintRewardsDistribution.TransactOpts, owner, _addressProvider)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address owner, address _addressProvider) returns()
func (_FMintRewardsDistribution *FMintRewardsDistributionTransactorSession) Initialize(owner common.Address, _addressProvider common.Address) (*types.Transaction, error) {
	return _FMintRewardsDistribution.Contract.Initialize(&_FMintRewardsDistribution.TransactOpts, owner, _addressProvider)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address sender) returns()
func (_FMintRewardsDistribution *FMintRewardsDistributionTransactor) Initialize0(opts *bind.TransactOpts, sender common.Address) (*types.Transaction, error) {
	return _FMintRewardsDistribution.contract.Transact(opts, "initialize0", sender)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address sender) returns()
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) Initialize0(sender common.Address) (*types.Transaction, error) {
	return _FMintRewardsDistribution.Contract.Initialize0(&_FMintRewardsDistribution.TransactOpts, sender)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address sender) returns()
func (_FMintRewardsDistribution *FMintRewardsDistributionTransactorSession) Initialize0(sender common.Address) (*types.Transaction, error) {
	return _FMintRewardsDistribution.Contract.Initialize0(&_FMintRewardsDistribution.TransactOpts, sender)
}

// MustRewardClaim is a paid mutator transaction binding the contract method 0x101df8b5.
//
// Solidity: function mustRewardClaim() returns()
func (_FMintRewardsDistribution *FMintRewardsDistributionTransactor) MustRewardClaim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FMintRewardsDistribution.contract.Transact(opts, "mustRewardClaim")
}

// MustRewardClaim is a paid mutator transaction binding the contract method 0x101df8b5.
//
// Solidity: function mustRewardClaim() returns()
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) MustRewardClaim() (*types.Transaction, error) {
	return _FMintRewardsDistribution.Contract.MustRewardClaim(&_FMintRewardsDistribution.TransactOpts)
}

// MustRewardClaim is a paid mutator transaction binding the contract method 0x101df8b5.
//
// Solidity: function mustRewardClaim() returns()
func (_FMintRewardsDistribution *FMintRewardsDistributionTransactorSession) MustRewardClaim() (*types.Transaction, error) {
	return _FMintRewardsDistribution.Contract.MustRewardClaim(&_FMintRewardsDistribution.TransactOpts)
}

// MustRewardPush is a paid mutator transaction binding the contract method 0x9b7ea007.
//
// Solidity: function mustRewardPush() returns()
func (_FMintRewardsDistribution *FMintRewardsDistributionTransactor) MustRewardPush(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FMintRewardsDistribution.contract.Transact(opts, "mustRewardPush")
}

// MustRewardPush is a paid mutator transaction binding the contract method 0x9b7ea007.
//
// Solidity: function mustRewardPush() returns()
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) MustRewardPush() (*types.Transaction, error) {
	return _FMintRewardsDistribution.Contract.MustRewardPush(&_FMintRewardsDistribution.TransactOpts)
}

// MustRewardPush is a paid mutator transaction binding the contract method 0x9b7ea007.
//
// Solidity: function mustRewardPush() returns()
func (_FMintRewardsDistribution *FMintRewardsDistributionTransactorSession) MustRewardPush() (*types.Transaction, error) {
	return _FMintRewardsDistribution.Contract.MustRewardPush(&_FMintRewardsDistribution.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FMintRewardsDistribution *FMintRewardsDistributionTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FMintRewardsDistribution.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) RenounceOwnership() (*types.Transaction, error) {
	return _FMintRewardsDistribution.Contract.RenounceOwnership(&_FMintRewardsDistribution.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FMintRewardsDistribution *FMintRewardsDistributionTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FMintRewardsDistribution.Contract.RenounceOwnership(&_FMintRewardsDistribution.TransactOpts)
}

// RewardClaim is a paid mutator transaction binding the contract method 0x6409f921.
//
// Solidity: function rewardClaim() returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionTransactor) RewardClaim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FMintRewardsDistribution.contract.Transact(opts, "rewardClaim")
}

// RewardClaim is a paid mutator transaction binding the contract method 0x6409f921.
//
// Solidity: function rewardClaim() returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) RewardClaim() (*types.Transaction, error) {
	return _FMintRewardsDistribution.Contract.RewardClaim(&_FMintRewardsDistribution.TransactOpts)
}

// RewardClaim is a paid mutator transaction binding the contract method 0x6409f921.
//
// Solidity: function rewardClaim() returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionTransactorSession) RewardClaim() (*types.Transaction, error) {
	return _FMintRewardsDistribution.Contract.RewardClaim(&_FMintRewardsDistribution.TransactOpts)
}

// RewardCleanup is a paid mutator transaction binding the contract method 0xc759df4e.
//
// Solidity: function rewardCleanup(address _recipient) returns()
func (_FMintRewardsDistribution *FMintRewardsDistributionTransactor) RewardCleanup(opts *bind.TransactOpts, _recipient common.Address) (*types.Transaction, error) {
	return _FMintRewardsDistribution.contract.Transact(opts, "rewardCleanup", _recipient)
}

// RewardCleanup is a paid mutator transaction binding the contract method 0xc759df4e.
//
// Solidity: function rewardCleanup(address _recipient) returns()
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) RewardCleanup(_recipient common.Address) (*types.Transaction, error) {
	return _FMintRewardsDistribution.Contract.RewardCleanup(&_FMintRewardsDistribution.TransactOpts, _recipient)
}

// RewardCleanup is a paid mutator transaction binding the contract method 0xc759df4e.
//
// Solidity: function rewardCleanup(address _recipient) returns()
func (_FMintRewardsDistribution *FMintRewardsDistributionTransactorSession) RewardCleanup(_recipient common.Address) (*types.Transaction, error) {
	return _FMintRewardsDistribution.Contract.RewardCleanup(&_FMintRewardsDistribution.TransactOpts, _recipient)
}

// RewardPush is a paid mutator transaction binding the contract method 0x185463a4.
//
// Solidity: function rewardPush() returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionTransactor) RewardPush(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FMintRewardsDistribution.contract.Transact(opts, "rewardPush")
}

// RewardPush is a paid mutator transaction binding the contract method 0x185463a4.
//
// Solidity: function rewardPush() returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) RewardPush() (*types.Transaction, error) {
	return _FMintRewardsDistribution.Contract.RewardPush(&_FMintRewardsDistribution.TransactOpts)
}

// RewardPush is a paid mutator transaction binding the contract method 0x185463a4.
//
// Solidity: function rewardPush() returns(uint256)
func (_FMintRewardsDistribution *FMintRewardsDistributionTransactorSession) RewardPush() (*types.Transaction, error) {
	return _FMintRewardsDistribution.Contract.RewardPush(&_FMintRewardsDistribution.TransactOpts)
}

// RewardUpdate is a paid mutator transaction binding the contract method 0x48ebb08d.
//
// Solidity: function rewardUpdate(address _account) returns()
func (_FMintRewardsDistribution *FMintRewardsDistributionTransactor) RewardUpdate(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _FMintRewardsDistribution.contract.Transact(opts, "rewardUpdate", _account)
}

// RewardUpdate is a paid mutator transaction binding the contract method 0x48ebb08d.
//
// Solidity: function rewardUpdate(address _account) returns()
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) RewardUpdate(_account common.Address) (*types.Transaction, error) {
	return _FMintRewardsDistribution.Contract.RewardUpdate(&_FMintRewardsDistribution.TransactOpts, _account)
}

// RewardUpdate is a paid mutator transaction binding the contract method 0x48ebb08d.
//
// Solidity: function rewardUpdate(address _account) returns()
func (_FMintRewardsDistribution *FMintRewardsDistributionTransactorSession) RewardUpdate(_account common.Address) (*types.Transaction, error) {
	return _FMintRewardsDistribution.Contract.RewardUpdate(&_FMintRewardsDistribution.TransactOpts, _account)
}

// RewardUpdateGlobal is a paid mutator transaction binding the contract method 0x64dd213f.
//
// Solidity: function rewardUpdateGlobal() returns()
func (_FMintRewardsDistribution *FMintRewardsDistributionTransactor) RewardUpdateGlobal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FMintRewardsDistribution.contract.Transact(opts, "rewardUpdateGlobal")
}

// RewardUpdateGlobal is a paid mutator transaction binding the contract method 0x64dd213f.
//
// Solidity: function rewardUpdateGlobal() returns()
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) RewardUpdateGlobal() (*types.Transaction, error) {
	return _FMintRewardsDistribution.Contract.RewardUpdateGlobal(&_FMintRewardsDistribution.TransactOpts)
}

// RewardUpdateGlobal is a paid mutator transaction binding the contract method 0x64dd213f.
//
// Solidity: function rewardUpdateGlobal() returns()
func (_FMintRewardsDistribution *FMintRewardsDistributionTransactorSession) RewardUpdateGlobal() (*types.Transaction, error) {
	return _FMintRewardsDistribution.Contract.RewardUpdateGlobal(&_FMintRewardsDistribution.TransactOpts)
}

// RewardUpdateRate is a paid mutator transaction binding the contract method 0xef24d9a1.
//
// Solidity: function rewardUpdateRate(uint256 _perSecond) returns()
func (_FMintRewardsDistribution *FMintRewardsDistributionTransactor) RewardUpdateRate(opts *bind.TransactOpts, _perSecond *big.Int) (*types.Transaction, error) {
	return _FMintRewardsDistribution.contract.Transact(opts, "rewardUpdateRate", _perSecond)
}

// RewardUpdateRate is a paid mutator transaction binding the contract method 0xef24d9a1.
//
// Solidity: function rewardUpdateRate(uint256 _perSecond) returns()
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) RewardUpdateRate(_perSecond *big.Int) (*types.Transaction, error) {
	return _FMintRewardsDistribution.Contract.RewardUpdateRate(&_FMintRewardsDistribution.TransactOpts, _perSecond)
}

// RewardUpdateRate is a paid mutator transaction binding the contract method 0xef24d9a1.
//
// Solidity: function rewardUpdateRate(uint256 _perSecond) returns()
func (_FMintRewardsDistribution *FMintRewardsDistributionTransactorSession) RewardUpdateRate(_perSecond *big.Int) (*types.Transaction, error) {
	return _FMintRewardsDistribution.Contract.RewardUpdateRate(&_FMintRewardsDistribution.TransactOpts, _perSecond)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FMintRewardsDistribution *FMintRewardsDistributionTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FMintRewardsDistribution.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FMintRewardsDistribution *FMintRewardsDistributionSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FMintRewardsDistribution.Contract.TransferOwnership(&_FMintRewardsDistribution.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FMintRewardsDistribution *FMintRewardsDistributionTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FMintRewardsDistribution.Contract.TransferOwnership(&_FMintRewardsDistribution.TransactOpts, newOwner)
}

// FMintRewardsDistributionOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FMintRewardsDistribution contract.
type FMintRewardsDistributionOwnershipTransferredIterator struct {
	Event *FMintRewardsDistributionOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FMintRewardsDistributionOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FMintRewardsDistributionOwnershipTransferred)
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
		it.Event = new(FMintRewardsDistributionOwnershipTransferred)
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
func (it *FMintRewardsDistributionOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FMintRewardsDistributionOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FMintRewardsDistributionOwnershipTransferred represents a OwnershipTransferred event raised by the FMintRewardsDistribution contract.
type FMintRewardsDistributionOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FMintRewardsDistribution *FMintRewardsDistributionFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FMintRewardsDistributionOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FMintRewardsDistribution.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FMintRewardsDistributionOwnershipTransferredIterator{contract: _FMintRewardsDistribution.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FMintRewardsDistribution *FMintRewardsDistributionFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FMintRewardsDistributionOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FMintRewardsDistribution.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FMintRewardsDistributionOwnershipTransferred)
				if err := _FMintRewardsDistribution.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FMintRewardsDistribution *FMintRewardsDistributionFilterer) ParseOwnershipTransferred(log types.Log) (*FMintRewardsDistributionOwnershipTransferred, error) {
	event := new(FMintRewardsDistributionOwnershipTransferred)
	if err := _FMintRewardsDistribution.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FMintRewardsDistributionRateUpdatedIterator is returned from FilterRateUpdated and is used to iterate over the raw logs and unpacked data for RateUpdated events raised by the FMintRewardsDistribution contract.
type FMintRewardsDistributionRateUpdatedIterator struct {
	Event *FMintRewardsDistributionRateUpdated // Event containing the contract specifics and raw log

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
func (it *FMintRewardsDistributionRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FMintRewardsDistributionRateUpdated)
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
		it.Event = new(FMintRewardsDistributionRateUpdated)
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
func (it *FMintRewardsDistributionRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FMintRewardsDistributionRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FMintRewardsDistributionRateUpdated represents a RateUpdated event raised by the FMintRewardsDistribution contract.
type FMintRewardsDistributionRateUpdated struct {
	RewardPerYear *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRateUpdated is a free log retrieval operation binding the contract event 0xe65c987b2e4668e09ba867026921588005b2b2063607a1e7e7d91683c8f91b7b.
//
// Solidity: event RateUpdated(uint256 rewardPerYear)
func (_FMintRewardsDistribution *FMintRewardsDistributionFilterer) FilterRateUpdated(opts *bind.FilterOpts) (*FMintRewardsDistributionRateUpdatedIterator, error) {

	logs, sub, err := _FMintRewardsDistribution.contract.FilterLogs(opts, "RateUpdated")
	if err != nil {
		return nil, err
	}
	return &FMintRewardsDistributionRateUpdatedIterator{contract: _FMintRewardsDistribution.contract, event: "RateUpdated", logs: logs, sub: sub}, nil
}

// WatchRateUpdated is a free log subscription operation binding the contract event 0xe65c987b2e4668e09ba867026921588005b2b2063607a1e7e7d91683c8f91b7b.
//
// Solidity: event RateUpdated(uint256 rewardPerYear)
func (_FMintRewardsDistribution *FMintRewardsDistributionFilterer) WatchRateUpdated(opts *bind.WatchOpts, sink chan<- *FMintRewardsDistributionRateUpdated) (event.Subscription, error) {

	logs, sub, err := _FMintRewardsDistribution.contract.WatchLogs(opts, "RateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FMintRewardsDistributionRateUpdated)
				if err := _FMintRewardsDistribution.contract.UnpackLog(event, "RateUpdated", log); err != nil {
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

// ParseRateUpdated is a log parse operation binding the contract event 0xe65c987b2e4668e09ba867026921588005b2b2063607a1e7e7d91683c8f91b7b.
//
// Solidity: event RateUpdated(uint256 rewardPerYear)
func (_FMintRewardsDistribution *FMintRewardsDistributionFilterer) ParseRateUpdated(log types.Log) (*FMintRewardsDistributionRateUpdated, error) {
	event := new(FMintRewardsDistributionRateUpdated)
	if err := _FMintRewardsDistribution.contract.UnpackLog(event, "RateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FMintRewardsDistributionRewardAddedIterator is returned from FilterRewardAdded and is used to iterate over the raw logs and unpacked data for RewardAdded events raised by the FMintRewardsDistribution contract.
type FMintRewardsDistributionRewardAddedIterator struct {
	Event *FMintRewardsDistributionRewardAdded // Event containing the contract specifics and raw log

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
func (it *FMintRewardsDistributionRewardAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FMintRewardsDistributionRewardAdded)
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
		it.Event = new(FMintRewardsDistributionRewardAdded)
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
func (it *FMintRewardsDistributionRewardAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FMintRewardsDistributionRewardAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FMintRewardsDistributionRewardAdded represents a RewardAdded event raised by the FMintRewardsDistribution contract.
type FMintRewardsDistributionRewardAdded struct {
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardAdded is a free log retrieval operation binding the contract event 0xde88a922e0d3b88b24e9623efeb464919c6bf9f66857a65e2bfcf2ce87a9433d.
//
// Solidity: event RewardAdded(uint256 reward)
func (_FMintRewardsDistribution *FMintRewardsDistributionFilterer) FilterRewardAdded(opts *bind.FilterOpts) (*FMintRewardsDistributionRewardAddedIterator, error) {

	logs, sub, err := _FMintRewardsDistribution.contract.FilterLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return &FMintRewardsDistributionRewardAddedIterator{contract: _FMintRewardsDistribution.contract, event: "RewardAdded", logs: logs, sub: sub}, nil
}

// WatchRewardAdded is a free log subscription operation binding the contract event 0xde88a922e0d3b88b24e9623efeb464919c6bf9f66857a65e2bfcf2ce87a9433d.
//
// Solidity: event RewardAdded(uint256 reward)
func (_FMintRewardsDistribution *FMintRewardsDistributionFilterer) WatchRewardAdded(opts *bind.WatchOpts, sink chan<- *FMintRewardsDistributionRewardAdded) (event.Subscription, error) {

	logs, sub, err := _FMintRewardsDistribution.contract.WatchLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FMintRewardsDistributionRewardAdded)
				if err := _FMintRewardsDistribution.contract.UnpackLog(event, "RewardAdded", log); err != nil {
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

// ParseRewardAdded is a log parse operation binding the contract event 0xde88a922e0d3b88b24e9623efeb464919c6bf9f66857a65e2bfcf2ce87a9433d.
//
// Solidity: event RewardAdded(uint256 reward)
func (_FMintRewardsDistribution *FMintRewardsDistributionFilterer) ParseRewardAdded(log types.Log) (*FMintRewardsDistributionRewardAdded, error) {
	event := new(FMintRewardsDistributionRewardAdded)
	if err := _FMintRewardsDistribution.contract.UnpackLog(event, "RewardAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FMintRewardsDistributionRewardPaidIterator is returned from FilterRewardPaid and is used to iterate over the raw logs and unpacked data for RewardPaid events raised by the FMintRewardsDistribution contract.
type FMintRewardsDistributionRewardPaidIterator struct {
	Event *FMintRewardsDistributionRewardPaid // Event containing the contract specifics and raw log

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
func (it *FMintRewardsDistributionRewardPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FMintRewardsDistributionRewardPaid)
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
		it.Event = new(FMintRewardsDistributionRewardPaid)
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
func (it *FMintRewardsDistributionRewardPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FMintRewardsDistributionRewardPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FMintRewardsDistributionRewardPaid represents a RewardPaid event raised by the FMintRewardsDistribution contract.
type FMintRewardsDistributionRewardPaid struct {
	User   common.Address
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardPaid is a free log retrieval operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_FMintRewardsDistribution *FMintRewardsDistributionFilterer) FilterRewardPaid(opts *bind.FilterOpts, user []common.Address) (*FMintRewardsDistributionRewardPaidIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _FMintRewardsDistribution.contract.FilterLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return &FMintRewardsDistributionRewardPaidIterator{contract: _FMintRewardsDistribution.contract, event: "RewardPaid", logs: logs, sub: sub}, nil
}

// WatchRewardPaid is a free log subscription operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_FMintRewardsDistribution *FMintRewardsDistributionFilterer) WatchRewardPaid(opts *bind.WatchOpts, sink chan<- *FMintRewardsDistributionRewardPaid, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _FMintRewardsDistribution.contract.WatchLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FMintRewardsDistributionRewardPaid)
				if err := _FMintRewardsDistribution.contract.UnpackLog(event, "RewardPaid", log); err != nil {
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

// ParseRewardPaid is a log parse operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_FMintRewardsDistribution *FMintRewardsDistributionFilterer) ParseRewardPaid(log types.Log) (*FMintRewardsDistributionRewardPaid, error) {
	event := new(FMintRewardsDistributionRewardPaid)
	if err := _FMintRewardsDistribution.contract.UnpackLog(event, "RewardPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
