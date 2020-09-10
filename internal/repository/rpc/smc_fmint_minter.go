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

// DefiFMintMinterABI is the input ABI used to generate the binding from.
const DefiFMintMinterABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressProvider\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Repaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_DEBT_EXCEEDED\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_LOW_ALLOWANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_LOW_BALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_LOW_COLLATERAL_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_MINTING_PROHIBITED\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_NOT_AUTHORIZED\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_NO_COLLATERAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_NO_ERROR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_NO_REWARD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_NO_VALUE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_REWARDS_DEPLETED\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_REWARDS_EARLY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_REWARDS_NONE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_REWARD_CLAIM_REJECTED\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_ZERO_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"addressProvider\",\"outputs\":[{\"internalType\":\"contractIFantomMintAddressProvider\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"collateralBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"collateralCanDecrease\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"collateralLowestDebtRatio4dec\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"collateralRatioDecimalsCorrection\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"collateralTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"collateralTokensCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"collateralTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"collateralTotalBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"collateralValueOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"debtBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"debtCanIncrease\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"debtTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"debtTokensCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"debtTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"debtTotalBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"debtValueOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fMintFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fMintFeeDigitsCorrection\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fMintPriceDigitsCorrection\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feePool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"principalBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"principalBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardApplicableUntil\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"rewardCanClaim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"rewardClaim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"rewardEarned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardEligibilityRatio4dec\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardEpochEnds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardEpochLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardLastPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"rewardNotifyAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardPerTokenDecimalsCorrection\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardPerTokenPaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardStash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardUpdated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"tokenValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardPoolAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardDistributionAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenRegistryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPriceDigitsCorrection\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"repay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DefiFMintMinter is an auto generated Go binding around an Ethereum contract.
type DefiFMintMinter struct {
	DefiFMintMinterCaller     // Read-only binding to the contract
	DefiFMintMinterTransactor // Write-only binding to the contract
	DefiFMintMinterFilterer   // Log filterer for contract events
}

// DefiFMintMinterCaller is an auto generated read-only Go binding around an Ethereum contract.
type DefiFMintMinterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefiFMintMinterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DefiFMintMinterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefiFMintMinterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DefiFMintMinterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefiFMintMinterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DefiFMintMinterSession struct {
	Contract     *DefiFMintMinter  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DefiFMintMinterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DefiFMintMinterCallerSession struct {
	Contract *DefiFMintMinterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// DefiFMintMinterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DefiFMintMinterTransactorSession struct {
	Contract     *DefiFMintMinterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// DefiFMintMinterRaw is an auto generated low-level Go binding around an Ethereum contract.
type DefiFMintMinterRaw struct {
	Contract *DefiFMintMinter // Generic contract binding to access the raw methods on
}

// DefiFMintMinterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DefiFMintMinterCallerRaw struct {
	Contract *DefiFMintMinterCaller // Generic read-only contract binding to access the raw methods on
}

// DefiFMintMinterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DefiFMintMinterTransactorRaw struct {
	Contract *DefiFMintMinterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDefiFMintMinter creates a new instance of DefiFMintMinter, bound to a specific deployed contract.
func NewDefiFMintMinter(address common.Address, backend bind.ContractBackend) (*DefiFMintMinter, error) {
	contract, err := bindDefiFMintMinter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DefiFMintMinter{DefiFMintMinterCaller: DefiFMintMinterCaller{contract: contract}, DefiFMintMinterTransactor: DefiFMintMinterTransactor{contract: contract}, DefiFMintMinterFilterer: DefiFMintMinterFilterer{contract: contract}}, nil
}

// NewDefiFMintMinterCaller creates a new read-only instance of DefiFMintMinter, bound to a specific deployed contract.
func NewDefiFMintMinterCaller(address common.Address, caller bind.ContractCaller) (*DefiFMintMinterCaller, error) {
	contract, err := bindDefiFMintMinter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DefiFMintMinterCaller{contract: contract}, nil
}

// NewDefiFMintMinterTransactor creates a new write-only instance of DefiFMintMinter, bound to a specific deployed contract.
func NewDefiFMintMinterTransactor(address common.Address, transactor bind.ContractTransactor) (*DefiFMintMinterTransactor, error) {
	contract, err := bindDefiFMintMinter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DefiFMintMinterTransactor{contract: contract}, nil
}

// NewDefiFMintMinterFilterer creates a new log filterer instance of DefiFMintMinter, bound to a specific deployed contract.
func NewDefiFMintMinterFilterer(address common.Address, filterer bind.ContractFilterer) (*DefiFMintMinterFilterer, error) {
	contract, err := bindDefiFMintMinter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DefiFMintMinterFilterer{contract: contract}, nil
}

// bindDefiFMintMinter binds a generic wrapper to an already deployed contract.
func bindDefiFMintMinter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DefiFMintMinterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DefiFMintMinter *DefiFMintMinterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DefiFMintMinter.Contract.DefiFMintMinterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DefiFMintMinter *DefiFMintMinterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.DefiFMintMinterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DefiFMintMinter *DefiFMintMinterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.DefiFMintMinterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DefiFMintMinter *DefiFMintMinterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DefiFMintMinter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DefiFMintMinter *DefiFMintMinterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DefiFMintMinter *DefiFMintMinterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.contract.Transact(opts, method, params...)
}

// ERRDEBTEXCEEDED is a free data retrieval call binding the contract method 0x372ce3df.
//
// Solidity: function ERR_DEBT_EXCEEDED() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) ERRDEBTEXCEEDED(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "ERR_DEBT_EXCEEDED")
	return *ret0, err
}

// ERRDEBTEXCEEDED is a free data retrieval call binding the contract method 0x372ce3df.
//
// Solidity: function ERR_DEBT_EXCEEDED() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) ERRDEBTEXCEEDED() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRDEBTEXCEEDED(&_DefiFMintMinter.CallOpts)
}

// ERRDEBTEXCEEDED is a free data retrieval call binding the contract method 0x372ce3df.
//
// Solidity: function ERR_DEBT_EXCEEDED() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) ERRDEBTEXCEEDED() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRDEBTEXCEEDED(&_DefiFMintMinter.CallOpts)
}

// ERRLOWALLOWANCE is a free data retrieval call binding the contract method 0x911fc3f1.
//
// Solidity: function ERR_LOW_ALLOWANCE() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) ERRLOWALLOWANCE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "ERR_LOW_ALLOWANCE")
	return *ret0, err
}

// ERRLOWALLOWANCE is a free data retrieval call binding the contract method 0x911fc3f1.
//
// Solidity: function ERR_LOW_ALLOWANCE() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) ERRLOWALLOWANCE() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRLOWALLOWANCE(&_DefiFMintMinter.CallOpts)
}

// ERRLOWALLOWANCE is a free data retrieval call binding the contract method 0x911fc3f1.
//
// Solidity: function ERR_LOW_ALLOWANCE() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) ERRLOWALLOWANCE() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRLOWALLOWANCE(&_DefiFMintMinter.CallOpts)
}

// ERRLOWBALANCE is a free data retrieval call binding the contract method 0x1ac919b0.
//
// Solidity: function ERR_LOW_BALANCE() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) ERRLOWBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "ERR_LOW_BALANCE")
	return *ret0, err
}

// ERRLOWBALANCE is a free data retrieval call binding the contract method 0x1ac919b0.
//
// Solidity: function ERR_LOW_BALANCE() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) ERRLOWBALANCE() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRLOWBALANCE(&_DefiFMintMinter.CallOpts)
}

// ERRLOWBALANCE is a free data retrieval call binding the contract method 0x1ac919b0.
//
// Solidity: function ERR_LOW_BALANCE() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) ERRLOWBALANCE() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRLOWBALANCE(&_DefiFMintMinter.CallOpts)
}

// ERRLOWCOLLATERALRATIO is a free data retrieval call binding the contract method 0x04b62f29.
//
// Solidity: function ERR_LOW_COLLATERAL_RATIO() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) ERRLOWCOLLATERALRATIO(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "ERR_LOW_COLLATERAL_RATIO")
	return *ret0, err
}

// ERRLOWCOLLATERALRATIO is a free data retrieval call binding the contract method 0x04b62f29.
//
// Solidity: function ERR_LOW_COLLATERAL_RATIO() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) ERRLOWCOLLATERALRATIO() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRLOWCOLLATERALRATIO(&_DefiFMintMinter.CallOpts)
}

// ERRLOWCOLLATERALRATIO is a free data retrieval call binding the contract method 0x04b62f29.
//
// Solidity: function ERR_LOW_COLLATERAL_RATIO() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) ERRLOWCOLLATERALRATIO() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRLOWCOLLATERALRATIO(&_DefiFMintMinter.CallOpts)
}

// ERRMINTINGPROHIBITED is a free data retrieval call binding the contract method 0x8c7b9980.
//
// Solidity: function ERR_MINTING_PROHIBITED() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) ERRMINTINGPROHIBITED(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "ERR_MINTING_PROHIBITED")
	return *ret0, err
}

// ERRMINTINGPROHIBITED is a free data retrieval call binding the contract method 0x8c7b9980.
//
// Solidity: function ERR_MINTING_PROHIBITED() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) ERRMINTINGPROHIBITED() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRMINTINGPROHIBITED(&_DefiFMintMinter.CallOpts)
}

// ERRMINTINGPROHIBITED is a free data retrieval call binding the contract method 0x8c7b9980.
//
// Solidity: function ERR_MINTING_PROHIBITED() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) ERRMINTINGPROHIBITED() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRMINTINGPROHIBITED(&_DefiFMintMinter.CallOpts)
}

// ERRNOTAUTHORIZED is a free data retrieval call binding the contract method 0xbc99d6ae.
//
// Solidity: function ERR_NOT_AUTHORIZED() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) ERRNOTAUTHORIZED(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "ERR_NOT_AUTHORIZED")
	return *ret0, err
}

// ERRNOTAUTHORIZED is a free data retrieval call binding the contract method 0xbc99d6ae.
//
// Solidity: function ERR_NOT_AUTHORIZED() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) ERRNOTAUTHORIZED() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRNOTAUTHORIZED(&_DefiFMintMinter.CallOpts)
}

// ERRNOTAUTHORIZED is a free data retrieval call binding the contract method 0xbc99d6ae.
//
// Solidity: function ERR_NOT_AUTHORIZED() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) ERRNOTAUTHORIZED() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRNOTAUTHORIZED(&_DefiFMintMinter.CallOpts)
}

// ERRNOCOLLATERAL is a free data retrieval call binding the contract method 0xb76361c2.
//
// Solidity: function ERR_NO_COLLATERAL() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) ERRNOCOLLATERAL(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "ERR_NO_COLLATERAL")
	return *ret0, err
}

// ERRNOCOLLATERAL is a free data retrieval call binding the contract method 0xb76361c2.
//
// Solidity: function ERR_NO_COLLATERAL() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) ERRNOCOLLATERAL() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRNOCOLLATERAL(&_DefiFMintMinter.CallOpts)
}

// ERRNOCOLLATERAL is a free data retrieval call binding the contract method 0xb76361c2.
//
// Solidity: function ERR_NO_COLLATERAL() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) ERRNOCOLLATERAL() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRNOCOLLATERAL(&_DefiFMintMinter.CallOpts)
}

// ERRNOERROR is a free data retrieval call binding the contract method 0x35052d6e.
//
// Solidity: function ERR_NO_ERROR() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) ERRNOERROR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "ERR_NO_ERROR")
	return *ret0, err
}

// ERRNOERROR is a free data retrieval call binding the contract method 0x35052d6e.
//
// Solidity: function ERR_NO_ERROR() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) ERRNOERROR() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRNOERROR(&_DefiFMintMinter.CallOpts)
}

// ERRNOERROR is a free data retrieval call binding the contract method 0x35052d6e.
//
// Solidity: function ERR_NO_ERROR() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) ERRNOERROR() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRNOERROR(&_DefiFMintMinter.CallOpts)
}

// ERRNOREWARD is a free data retrieval call binding the contract method 0x73a93af6.
//
// Solidity: function ERR_NO_REWARD() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) ERRNOREWARD(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "ERR_NO_REWARD")
	return *ret0, err
}

// ERRNOREWARD is a free data retrieval call binding the contract method 0x73a93af6.
//
// Solidity: function ERR_NO_REWARD() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) ERRNOREWARD() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRNOREWARD(&_DefiFMintMinter.CallOpts)
}

// ERRNOREWARD is a free data retrieval call binding the contract method 0x73a93af6.
//
// Solidity: function ERR_NO_REWARD() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) ERRNOREWARD() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRNOREWARD(&_DefiFMintMinter.CallOpts)
}

// ERRNOVALUE is a free data retrieval call binding the contract method 0x69d1cb27.
//
// Solidity: function ERR_NO_VALUE() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) ERRNOVALUE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "ERR_NO_VALUE")
	return *ret0, err
}

// ERRNOVALUE is a free data retrieval call binding the contract method 0x69d1cb27.
//
// Solidity: function ERR_NO_VALUE() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) ERRNOVALUE() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRNOVALUE(&_DefiFMintMinter.CallOpts)
}

// ERRNOVALUE is a free data retrieval call binding the contract method 0x69d1cb27.
//
// Solidity: function ERR_NO_VALUE() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) ERRNOVALUE() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRNOVALUE(&_DefiFMintMinter.CallOpts)
}

// ERRREWARDSDEPLETED is a free data retrieval call binding the contract method 0x0a19dd33.
//
// Solidity: function ERR_REWARDS_DEPLETED() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) ERRREWARDSDEPLETED(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "ERR_REWARDS_DEPLETED")
	return *ret0, err
}

// ERRREWARDSDEPLETED is a free data retrieval call binding the contract method 0x0a19dd33.
//
// Solidity: function ERR_REWARDS_DEPLETED() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) ERRREWARDSDEPLETED() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRREWARDSDEPLETED(&_DefiFMintMinter.CallOpts)
}

// ERRREWARDSDEPLETED is a free data retrieval call binding the contract method 0x0a19dd33.
//
// Solidity: function ERR_REWARDS_DEPLETED() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) ERRREWARDSDEPLETED() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRREWARDSDEPLETED(&_DefiFMintMinter.CallOpts)
}

// ERRREWARDSEARLY is a free data retrieval call binding the contract method 0x67fc176b.
//
// Solidity: function ERR_REWARDS_EARLY() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) ERRREWARDSEARLY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "ERR_REWARDS_EARLY")
	return *ret0, err
}

// ERRREWARDSEARLY is a free data retrieval call binding the contract method 0x67fc176b.
//
// Solidity: function ERR_REWARDS_EARLY() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) ERRREWARDSEARLY() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRREWARDSEARLY(&_DefiFMintMinter.CallOpts)
}

// ERRREWARDSEARLY is a free data retrieval call binding the contract method 0x67fc176b.
//
// Solidity: function ERR_REWARDS_EARLY() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) ERRREWARDSEARLY() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRREWARDSEARLY(&_DefiFMintMinter.CallOpts)
}

// ERRREWARDSNONE is a free data retrieval call binding the contract method 0xc7222c72.
//
// Solidity: function ERR_REWARDS_NONE() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) ERRREWARDSNONE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "ERR_REWARDS_NONE")
	return *ret0, err
}

// ERRREWARDSNONE is a free data retrieval call binding the contract method 0xc7222c72.
//
// Solidity: function ERR_REWARDS_NONE() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) ERRREWARDSNONE() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRREWARDSNONE(&_DefiFMintMinter.CallOpts)
}

// ERRREWARDSNONE is a free data retrieval call binding the contract method 0xc7222c72.
//
// Solidity: function ERR_REWARDS_NONE() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) ERRREWARDSNONE() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRREWARDSNONE(&_DefiFMintMinter.CallOpts)
}

// ERRREWARDCLAIMREJECTED is a free data retrieval call binding the contract method 0x4846e345.
//
// Solidity: function ERR_REWARD_CLAIM_REJECTED() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) ERRREWARDCLAIMREJECTED(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "ERR_REWARD_CLAIM_REJECTED")
	return *ret0, err
}

// ERRREWARDCLAIMREJECTED is a free data retrieval call binding the contract method 0x4846e345.
//
// Solidity: function ERR_REWARD_CLAIM_REJECTED() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) ERRREWARDCLAIMREJECTED() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRREWARDCLAIMREJECTED(&_DefiFMintMinter.CallOpts)
}

// ERRREWARDCLAIMREJECTED is a free data retrieval call binding the contract method 0x4846e345.
//
// Solidity: function ERR_REWARD_CLAIM_REJECTED() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) ERRREWARDCLAIMREJECTED() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRREWARDCLAIMREJECTED(&_DefiFMintMinter.CallOpts)
}

// ERRZEROAMOUNT is a free data retrieval call binding the contract method 0x0aff90bb.
//
// Solidity: function ERR_ZERO_AMOUNT() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) ERRZEROAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "ERR_ZERO_AMOUNT")
	return *ret0, err
}

// ERRZEROAMOUNT is a free data retrieval call binding the contract method 0x0aff90bb.
//
// Solidity: function ERR_ZERO_AMOUNT() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) ERRZEROAMOUNT() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRZEROAMOUNT(&_DefiFMintMinter.CallOpts)
}

// ERRZEROAMOUNT is a free data retrieval call binding the contract method 0x0aff90bb.
//
// Solidity: function ERR_ZERO_AMOUNT() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) ERRZEROAMOUNT() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRZEROAMOUNT(&_DefiFMintMinter.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_DefiFMintMinter *DefiFMintMinterCaller) AddressProvider(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "addressProvider")
	return *ret0, err
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_DefiFMintMinter *DefiFMintMinterSession) AddressProvider() (common.Address, error) {
	return _DefiFMintMinter.Contract.AddressProvider(&_DefiFMintMinter.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) AddressProvider() (common.Address, error) {
	return _DefiFMintMinter.Contract.AddressProvider(&_DefiFMintMinter.CallOpts)
}

// CollateralBalance is a free data retrieval call binding the contract method 0xe7602b9d.
//
// Solidity: function collateralBalance(address , address ) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) CollateralBalance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "collateralBalance", arg0, arg1)
	return *ret0, err
}

// CollateralBalance is a free data retrieval call binding the contract method 0xe7602b9d.
//
// Solidity: function collateralBalance(address , address ) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) CollateralBalance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DefiFMintMinter.Contract.CollateralBalance(&_DefiFMintMinter.CallOpts, arg0, arg1)
}

// CollateralBalance is a free data retrieval call binding the contract method 0xe7602b9d.
//
// Solidity: function collateralBalance(address , address ) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) CollateralBalance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DefiFMintMinter.Contract.CollateralBalance(&_DefiFMintMinter.CallOpts, arg0, arg1)
}

// CollateralCanDecrease is a free data retrieval call binding the contract method 0xf4305a99.
//
// Solidity: function collateralCanDecrease(address _account, address _token, uint256 _amount) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterCaller) CollateralCanDecrease(opts *bind.CallOpts, _account common.Address, _token common.Address, _amount *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "collateralCanDecrease", _account, _token, _amount)
	return *ret0, err
}

// CollateralCanDecrease is a free data retrieval call binding the contract method 0xf4305a99.
//
// Solidity: function collateralCanDecrease(address _account, address _token, uint256 _amount) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterSession) CollateralCanDecrease(_account common.Address, _token common.Address, _amount *big.Int) (bool, error) {
	return _DefiFMintMinter.Contract.CollateralCanDecrease(&_DefiFMintMinter.CallOpts, _account, _token, _amount)
}

// CollateralCanDecrease is a free data retrieval call binding the contract method 0xf4305a99.
//
// Solidity: function collateralCanDecrease(address _account, address _token, uint256 _amount) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) CollateralCanDecrease(_account common.Address, _token common.Address, _amount *big.Int) (bool, error) {
	return _DefiFMintMinter.Contract.CollateralCanDecrease(&_DefiFMintMinter.CallOpts, _account, _token, _amount)
}

// CollateralLowestDebtRatio4dec is a free data retrieval call binding the contract method 0x3b8b09b7.
//
// Solidity: function collateralLowestDebtRatio4dec() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) CollateralLowestDebtRatio4dec(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "collateralLowestDebtRatio4dec")
	return *ret0, err
}

// CollateralLowestDebtRatio4dec is a free data retrieval call binding the contract method 0x3b8b09b7.
//
// Solidity: function collateralLowestDebtRatio4dec() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) CollateralLowestDebtRatio4dec() (*big.Int, error) {
	return _DefiFMintMinter.Contract.CollateralLowestDebtRatio4dec(&_DefiFMintMinter.CallOpts)
}

// CollateralLowestDebtRatio4dec is a free data retrieval call binding the contract method 0x3b8b09b7.
//
// Solidity: function collateralLowestDebtRatio4dec() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) CollateralLowestDebtRatio4dec() (*big.Int, error) {
	return _DefiFMintMinter.Contract.CollateralLowestDebtRatio4dec(&_DefiFMintMinter.CallOpts)
}

// CollateralRatioDecimalsCorrection is a free data retrieval call binding the contract method 0xe69993ac.
//
// Solidity: function collateralRatioDecimalsCorrection() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) CollateralRatioDecimalsCorrection(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "collateralRatioDecimalsCorrection")
	return *ret0, err
}

// CollateralRatioDecimalsCorrection is a free data retrieval call binding the contract method 0xe69993ac.
//
// Solidity: function collateralRatioDecimalsCorrection() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) CollateralRatioDecimalsCorrection() (*big.Int, error) {
	return _DefiFMintMinter.Contract.CollateralRatioDecimalsCorrection(&_DefiFMintMinter.CallOpts)
}

// CollateralRatioDecimalsCorrection is a free data retrieval call binding the contract method 0xe69993ac.
//
// Solidity: function collateralRatioDecimalsCorrection() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) CollateralRatioDecimalsCorrection() (*big.Int, error) {
	return _DefiFMintMinter.Contract.CollateralRatioDecimalsCorrection(&_DefiFMintMinter.CallOpts)
}

// CollateralTokens is a free data retrieval call binding the contract method 0x172c48c7.
//
// Solidity: function collateralTokens(uint256 ) view returns(address)
func (_DefiFMintMinter *DefiFMintMinterCaller) CollateralTokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "collateralTokens", arg0)
	return *ret0, err
}

// CollateralTokens is a free data retrieval call binding the contract method 0x172c48c7.
//
// Solidity: function collateralTokens(uint256 ) view returns(address)
func (_DefiFMintMinter *DefiFMintMinterSession) CollateralTokens(arg0 *big.Int) (common.Address, error) {
	return _DefiFMintMinter.Contract.CollateralTokens(&_DefiFMintMinter.CallOpts, arg0)
}

// CollateralTokens is a free data retrieval call binding the contract method 0x172c48c7.
//
// Solidity: function collateralTokens(uint256 ) view returns(address)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) CollateralTokens(arg0 *big.Int) (common.Address, error) {
	return _DefiFMintMinter.Contract.CollateralTokens(&_DefiFMintMinter.CallOpts, arg0)
}

// CollateralTokensCount is a free data retrieval call binding the contract method 0x458936f5.
//
// Solidity: function collateralTokensCount() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) CollateralTokensCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "collateralTokensCount")
	return *ret0, err
}

// CollateralTokensCount is a free data retrieval call binding the contract method 0x458936f5.
//
// Solidity: function collateralTokensCount() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) CollateralTokensCount() (*big.Int, error) {
	return _DefiFMintMinter.Contract.CollateralTokensCount(&_DefiFMintMinter.CallOpts)
}

// CollateralTokensCount is a free data retrieval call binding the contract method 0x458936f5.
//
// Solidity: function collateralTokensCount() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) CollateralTokensCount() (*big.Int, error) {
	return _DefiFMintMinter.Contract.CollateralTokensCount(&_DefiFMintMinter.CallOpts)
}

// CollateralTotal is a free data retrieval call binding the contract method 0x96427a2b.
//
// Solidity: function collateralTotal() view returns(uint256 value)
func (_DefiFMintMinter *DefiFMintMinterCaller) CollateralTotal(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "collateralTotal")
	return *ret0, err
}

// CollateralTotal is a free data retrieval call binding the contract method 0x96427a2b.
//
// Solidity: function collateralTotal() view returns(uint256 value)
func (_DefiFMintMinter *DefiFMintMinterSession) CollateralTotal() (*big.Int, error) {
	return _DefiFMintMinter.Contract.CollateralTotal(&_DefiFMintMinter.CallOpts)
}

// CollateralTotal is a free data retrieval call binding the contract method 0x96427a2b.
//
// Solidity: function collateralTotal() view returns(uint256 value)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) CollateralTotal() (*big.Int, error) {
	return _DefiFMintMinter.Contract.CollateralTotal(&_DefiFMintMinter.CallOpts)
}

// CollateralTotalBalance is a free data retrieval call binding the contract method 0x1d842131.
//
// Solidity: function collateralTotalBalance(address ) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) CollateralTotalBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "collateralTotalBalance", arg0)
	return *ret0, err
}

// CollateralTotalBalance is a free data retrieval call binding the contract method 0x1d842131.
//
// Solidity: function collateralTotalBalance(address ) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) CollateralTotalBalance(arg0 common.Address) (*big.Int, error) {
	return _DefiFMintMinter.Contract.CollateralTotalBalance(&_DefiFMintMinter.CallOpts, arg0)
}

// CollateralTotalBalance is a free data retrieval call binding the contract method 0x1d842131.
//
// Solidity: function collateralTotalBalance(address ) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) CollateralTotalBalance(arg0 common.Address) (*big.Int, error) {
	return _DefiFMintMinter.Contract.CollateralTotalBalance(&_DefiFMintMinter.CallOpts, arg0)
}

// CollateralValueOf is a free data retrieval call binding the contract method 0x3a65a350.
//
// Solidity: function collateralValueOf(address _account) view returns(uint256 value)
func (_DefiFMintMinter *DefiFMintMinterCaller) CollateralValueOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "collateralValueOf", _account)
	return *ret0, err
}

// CollateralValueOf is a free data retrieval call binding the contract method 0x3a65a350.
//
// Solidity: function collateralValueOf(address _account) view returns(uint256 value)
func (_DefiFMintMinter *DefiFMintMinterSession) CollateralValueOf(_account common.Address) (*big.Int, error) {
	return _DefiFMintMinter.Contract.CollateralValueOf(&_DefiFMintMinter.CallOpts, _account)
}

// CollateralValueOf is a free data retrieval call binding the contract method 0x3a65a350.
//
// Solidity: function collateralValueOf(address _account) view returns(uint256 value)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) CollateralValueOf(_account common.Address) (*big.Int, error) {
	return _DefiFMintMinter.Contract.CollateralValueOf(&_DefiFMintMinter.CallOpts, _account)
}

// DebtBalance is a free data retrieval call binding the contract method 0xad8f240e.
//
// Solidity: function debtBalance(address , address ) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) DebtBalance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "debtBalance", arg0, arg1)
	return *ret0, err
}

// DebtBalance is a free data retrieval call binding the contract method 0xad8f240e.
//
// Solidity: function debtBalance(address , address ) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) DebtBalance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DefiFMintMinter.Contract.DebtBalance(&_DefiFMintMinter.CallOpts, arg0, arg1)
}

// DebtBalance is a free data retrieval call binding the contract method 0xad8f240e.
//
// Solidity: function debtBalance(address , address ) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) DebtBalance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DefiFMintMinter.Contract.DebtBalance(&_DefiFMintMinter.CallOpts, arg0, arg1)
}

// DebtCanIncrease is a free data retrieval call binding the contract method 0x905ca247.
//
// Solidity: function debtCanIncrease(address _account, address _token, uint256 _amount) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterCaller) DebtCanIncrease(opts *bind.CallOpts, _account common.Address, _token common.Address, _amount *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "debtCanIncrease", _account, _token, _amount)
	return *ret0, err
}

// DebtCanIncrease is a free data retrieval call binding the contract method 0x905ca247.
//
// Solidity: function debtCanIncrease(address _account, address _token, uint256 _amount) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterSession) DebtCanIncrease(_account common.Address, _token common.Address, _amount *big.Int) (bool, error) {
	return _DefiFMintMinter.Contract.DebtCanIncrease(&_DefiFMintMinter.CallOpts, _account, _token, _amount)
}

// DebtCanIncrease is a free data retrieval call binding the contract method 0x905ca247.
//
// Solidity: function debtCanIncrease(address _account, address _token, uint256 _amount) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) DebtCanIncrease(_account common.Address, _token common.Address, _amount *big.Int) (bool, error) {
	return _DefiFMintMinter.Contract.DebtCanIncrease(&_DefiFMintMinter.CallOpts, _account, _token, _amount)
}

// DebtTokens is a free data retrieval call binding the contract method 0xf48f3b40.
//
// Solidity: function debtTokens(uint256 ) view returns(address)
func (_DefiFMintMinter *DefiFMintMinterCaller) DebtTokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "debtTokens", arg0)
	return *ret0, err
}

// DebtTokens is a free data retrieval call binding the contract method 0xf48f3b40.
//
// Solidity: function debtTokens(uint256 ) view returns(address)
func (_DefiFMintMinter *DefiFMintMinterSession) DebtTokens(arg0 *big.Int) (common.Address, error) {
	return _DefiFMintMinter.Contract.DebtTokens(&_DefiFMintMinter.CallOpts, arg0)
}

// DebtTokens is a free data retrieval call binding the contract method 0xf48f3b40.
//
// Solidity: function debtTokens(uint256 ) view returns(address)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) DebtTokens(arg0 *big.Int) (common.Address, error) {
	return _DefiFMintMinter.Contract.DebtTokens(&_DefiFMintMinter.CallOpts, arg0)
}

// DebtTokensCount is a free data retrieval call binding the contract method 0x9ee4736d.
//
// Solidity: function debtTokensCount() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) DebtTokensCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "debtTokensCount")
	return *ret0, err
}

// DebtTokensCount is a free data retrieval call binding the contract method 0x9ee4736d.
//
// Solidity: function debtTokensCount() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) DebtTokensCount() (*big.Int, error) {
	return _DefiFMintMinter.Contract.DebtTokensCount(&_DefiFMintMinter.CallOpts)
}

// DebtTokensCount is a free data retrieval call binding the contract method 0x9ee4736d.
//
// Solidity: function debtTokensCount() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) DebtTokensCount() (*big.Int, error) {
	return _DefiFMintMinter.Contract.DebtTokensCount(&_DefiFMintMinter.CallOpts)
}

// DebtTotal is a free data retrieval call binding the contract method 0xe7e7e387.
//
// Solidity: function debtTotal() view returns(uint256 value)
func (_DefiFMintMinter *DefiFMintMinterCaller) DebtTotal(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "debtTotal")
	return *ret0, err
}

// DebtTotal is a free data retrieval call binding the contract method 0xe7e7e387.
//
// Solidity: function debtTotal() view returns(uint256 value)
func (_DefiFMintMinter *DefiFMintMinterSession) DebtTotal() (*big.Int, error) {
	return _DefiFMintMinter.Contract.DebtTotal(&_DefiFMintMinter.CallOpts)
}

// DebtTotal is a free data retrieval call binding the contract method 0xe7e7e387.
//
// Solidity: function debtTotal() view returns(uint256 value)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) DebtTotal() (*big.Int, error) {
	return _DefiFMintMinter.Contract.DebtTotal(&_DefiFMintMinter.CallOpts)
}

// DebtTotalBalance is a free data retrieval call binding the contract method 0xf2423ea1.
//
// Solidity: function debtTotalBalance(address ) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) DebtTotalBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "debtTotalBalance", arg0)
	return *ret0, err
}

// DebtTotalBalance is a free data retrieval call binding the contract method 0xf2423ea1.
//
// Solidity: function debtTotalBalance(address ) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) DebtTotalBalance(arg0 common.Address) (*big.Int, error) {
	return _DefiFMintMinter.Contract.DebtTotalBalance(&_DefiFMintMinter.CallOpts, arg0)
}

// DebtTotalBalance is a free data retrieval call binding the contract method 0xf2423ea1.
//
// Solidity: function debtTotalBalance(address ) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) DebtTotalBalance(arg0 common.Address) (*big.Int, error) {
	return _DefiFMintMinter.Contract.DebtTotalBalance(&_DefiFMintMinter.CallOpts, arg0)
}

// DebtValueOf is a free data retrieval call binding the contract method 0x2f573910.
//
// Solidity: function debtValueOf(address _account) view returns(uint256 value)
func (_DefiFMintMinter *DefiFMintMinterCaller) DebtValueOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "debtValueOf", _account)
	return *ret0, err
}

// DebtValueOf is a free data retrieval call binding the contract method 0x2f573910.
//
// Solidity: function debtValueOf(address _account) view returns(uint256 value)
func (_DefiFMintMinter *DefiFMintMinterSession) DebtValueOf(_account common.Address) (*big.Int, error) {
	return _DefiFMintMinter.Contract.DebtValueOf(&_DefiFMintMinter.CallOpts, _account)
}

// DebtValueOf is a free data retrieval call binding the contract method 0x2f573910.
//
// Solidity: function debtValueOf(address _account) view returns(uint256 value)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) DebtValueOf(_account common.Address) (*big.Int, error) {
	return _DefiFMintMinter.Contract.DebtValueOf(&_DefiFMintMinter.CallOpts, _account)
}

// FMintFee is a free data retrieval call binding the contract method 0x9ccf1201.
//
// Solidity: function fMintFee() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) FMintFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "fMintFee")
	return *ret0, err
}

// FMintFee is a free data retrieval call binding the contract method 0x9ccf1201.
//
// Solidity: function fMintFee() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) FMintFee() (*big.Int, error) {
	return _DefiFMintMinter.Contract.FMintFee(&_DefiFMintMinter.CallOpts)
}

// FMintFee is a free data retrieval call binding the contract method 0x9ccf1201.
//
// Solidity: function fMintFee() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) FMintFee() (*big.Int, error) {
	return _DefiFMintMinter.Contract.FMintFee(&_DefiFMintMinter.CallOpts)
}

// FMintFeeDigitsCorrection is a free data retrieval call binding the contract method 0xcbf02fd5.
//
// Solidity: function fMintFeeDigitsCorrection() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) FMintFeeDigitsCorrection(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "fMintFeeDigitsCorrection")
	return *ret0, err
}

// FMintFeeDigitsCorrection is a free data retrieval call binding the contract method 0xcbf02fd5.
//
// Solidity: function fMintFeeDigitsCorrection() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) FMintFeeDigitsCorrection() (*big.Int, error) {
	return _DefiFMintMinter.Contract.FMintFeeDigitsCorrection(&_DefiFMintMinter.CallOpts)
}

// FMintFeeDigitsCorrection is a free data retrieval call binding the contract method 0xcbf02fd5.
//
// Solidity: function fMintFeeDigitsCorrection() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) FMintFeeDigitsCorrection() (*big.Int, error) {
	return _DefiFMintMinter.Contract.FMintFeeDigitsCorrection(&_DefiFMintMinter.CallOpts)
}

// FMintPriceDigitsCorrection is a free data retrieval call binding the contract method 0x65b61259.
//
// Solidity: function fMintPriceDigitsCorrection() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) FMintPriceDigitsCorrection(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "fMintPriceDigitsCorrection")
	return *ret0, err
}

// FMintPriceDigitsCorrection is a free data retrieval call binding the contract method 0x65b61259.
//
// Solidity: function fMintPriceDigitsCorrection() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) FMintPriceDigitsCorrection() (*big.Int, error) {
	return _DefiFMintMinter.Contract.FMintPriceDigitsCorrection(&_DefiFMintMinter.CallOpts)
}

// FMintPriceDigitsCorrection is a free data retrieval call binding the contract method 0x65b61259.
//
// Solidity: function fMintPriceDigitsCorrection() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) FMintPriceDigitsCorrection() (*big.Int, error) {
	return _DefiFMintMinter.Contract.FMintPriceDigitsCorrection(&_DefiFMintMinter.CallOpts)
}

// FeePool is a free data retrieval call binding the contract method 0xae2e933b.
//
// Solidity: function feePool() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) FeePool(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "feePool")
	return *ret0, err
}

// FeePool is a free data retrieval call binding the contract method 0xae2e933b.
//
// Solidity: function feePool() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) FeePool() (*big.Int, error) {
	return _DefiFMintMinter.Contract.FeePool(&_DefiFMintMinter.CallOpts)
}

// FeePool is a free data retrieval call binding the contract method 0xae2e933b.
//
// Solidity: function feePool() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) FeePool() (*big.Int, error) {
	return _DefiFMintMinter.Contract.FeePool(&_DefiFMintMinter.CallOpts)
}

// FeeTokenAddress is a free data retrieval call binding the contract method 0xb8df0dea.
//
// Solidity: function feeTokenAddress() view returns(address)
func (_DefiFMintMinter *DefiFMintMinterCaller) FeeTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "feeTokenAddress")
	return *ret0, err
}

// FeeTokenAddress is a free data retrieval call binding the contract method 0xb8df0dea.
//
// Solidity: function feeTokenAddress() view returns(address)
func (_DefiFMintMinter *DefiFMintMinterSession) FeeTokenAddress() (common.Address, error) {
	return _DefiFMintMinter.Contract.FeeTokenAddress(&_DefiFMintMinter.CallOpts)
}

// FeeTokenAddress is a free data retrieval call binding the contract method 0xb8df0dea.
//
// Solidity: function feeTokenAddress() view returns(address)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) FeeTokenAddress() (common.Address, error) {
	return _DefiFMintMinter.Contract.FeeTokenAddress(&_DefiFMintMinter.CallOpts)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address _token) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) GetPrice(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "getPrice", _token)
	return *ret0, err
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address _token) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) GetPrice(_token common.Address) (*big.Int, error) {
	return _DefiFMintMinter.Contract.GetPrice(&_DefiFMintMinter.CallOpts, _token)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address _token) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) GetPrice(_token common.Address) (*big.Int, error) {
	return _DefiFMintMinter.Contract.GetPrice(&_DefiFMintMinter.CallOpts, _token)
}

// GetPriceDigitsCorrection is a free data retrieval call binding the contract method 0xfadeada2.
//
// Solidity: function getPriceDigitsCorrection() pure returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) GetPriceDigitsCorrection(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "getPriceDigitsCorrection")
	return *ret0, err
}

// GetPriceDigitsCorrection is a free data retrieval call binding the contract method 0xfadeada2.
//
// Solidity: function getPriceDigitsCorrection() pure returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) GetPriceDigitsCorrection() (*big.Int, error) {
	return _DefiFMintMinter.Contract.GetPriceDigitsCorrection(&_DefiFMintMinter.CallOpts)
}

// GetPriceDigitsCorrection is a free data retrieval call binding the contract method 0xfadeada2.
//
// Solidity: function getPriceDigitsCorrection() pure returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) GetPriceDigitsCorrection() (*big.Int, error) {
	return _DefiFMintMinter.Contract.GetPriceDigitsCorrection(&_DefiFMintMinter.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterSession) IsOwner() (bool, error) {
	return _DefiFMintMinter.Contract.IsOwner(&_DefiFMintMinter.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) IsOwner() (bool, error) {
	return _DefiFMintMinter.Contract.IsOwner(&_DefiFMintMinter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DefiFMintMinter *DefiFMintMinterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DefiFMintMinter *DefiFMintMinterSession) Owner() (common.Address, error) {
	return _DefiFMintMinter.Contract.Owner(&_DefiFMintMinter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) Owner() (common.Address, error) {
	return _DefiFMintMinter.Contract.Owner(&_DefiFMintMinter.CallOpts)
}

// PrincipalBalance is a free data retrieval call binding the contract method 0xa83e53ac.
//
// Solidity: function principalBalance() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) PrincipalBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "principalBalance")
	return *ret0, err
}

// PrincipalBalance is a free data retrieval call binding the contract method 0xa83e53ac.
//
// Solidity: function principalBalance() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) PrincipalBalance() (*big.Int, error) {
	return _DefiFMintMinter.Contract.PrincipalBalance(&_DefiFMintMinter.CallOpts)
}

// PrincipalBalance is a free data retrieval call binding the contract method 0xa83e53ac.
//
// Solidity: function principalBalance() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) PrincipalBalance() (*big.Int, error) {
	return _DefiFMintMinter.Contract.PrincipalBalance(&_DefiFMintMinter.CallOpts)
}

// PrincipalBalanceOf is a free data retrieval call binding the contract method 0xc634dfaa.
//
// Solidity: function principalBalanceOf(address _account) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) PrincipalBalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "principalBalanceOf", _account)
	return *ret0, err
}

// PrincipalBalanceOf is a free data retrieval call binding the contract method 0xc634dfaa.
//
// Solidity: function principalBalanceOf(address _account) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) PrincipalBalanceOf(_account common.Address) (*big.Int, error) {
	return _DefiFMintMinter.Contract.PrincipalBalanceOf(&_DefiFMintMinter.CallOpts, _account)
}

// PrincipalBalanceOf is a free data retrieval call binding the contract method 0xc634dfaa.
//
// Solidity: function principalBalanceOf(address _account) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) PrincipalBalanceOf(_account common.Address) (*big.Int, error) {
	return _DefiFMintMinter.Contract.PrincipalBalanceOf(&_DefiFMintMinter.CallOpts, _account)
}

// RewardApplicableUntil is a free data retrieval call binding the contract method 0xdb16e0b5.
//
// Solidity: function rewardApplicableUntil() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) RewardApplicableUntil(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "rewardApplicableUntil")
	return *ret0, err
}

// RewardApplicableUntil is a free data retrieval call binding the contract method 0xdb16e0b5.
//
// Solidity: function rewardApplicableUntil() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) RewardApplicableUntil() (*big.Int, error) {
	return _DefiFMintMinter.Contract.RewardApplicableUntil(&_DefiFMintMinter.CallOpts)
}

// RewardApplicableUntil is a free data retrieval call binding the contract method 0xdb16e0b5.
//
// Solidity: function rewardApplicableUntil() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) RewardApplicableUntil() (*big.Int, error) {
	return _DefiFMintMinter.Contract.RewardApplicableUntil(&_DefiFMintMinter.CallOpts)
}

// RewardCanClaim is a free data retrieval call binding the contract method 0xda0a0432.
//
// Solidity: function rewardCanClaim(address _account) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterCaller) RewardCanClaim(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "rewardCanClaim", _account)
	return *ret0, err
}

// RewardCanClaim is a free data retrieval call binding the contract method 0xda0a0432.
//
// Solidity: function rewardCanClaim(address _account) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterSession) RewardCanClaim(_account common.Address) (bool, error) {
	return _DefiFMintMinter.Contract.RewardCanClaim(&_DefiFMintMinter.CallOpts, _account)
}

// RewardCanClaim is a free data retrieval call binding the contract method 0xda0a0432.
//
// Solidity: function rewardCanClaim(address _account) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) RewardCanClaim(_account common.Address) (bool, error) {
	return _DefiFMintMinter.Contract.RewardCanClaim(&_DefiFMintMinter.CallOpts, _account)
}

// RewardDistributionAddress is a free data retrieval call binding the contract method 0x72e13848.
//
// Solidity: function rewardDistributionAddress() view returns(address)
func (_DefiFMintMinter *DefiFMintMinterCaller) RewardDistributionAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "rewardDistributionAddress")
	return *ret0, err
}

// RewardDistributionAddress is a free data retrieval call binding the contract method 0x72e13848.
//
// Solidity: function rewardDistributionAddress() view returns(address)
func (_DefiFMintMinter *DefiFMintMinterSession) RewardDistributionAddress() (common.Address, error) {
	return _DefiFMintMinter.Contract.RewardDistributionAddress(&_DefiFMintMinter.CallOpts)
}

// RewardDistributionAddress is a free data retrieval call binding the contract method 0x72e13848.
//
// Solidity: function rewardDistributionAddress() view returns(address)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) RewardDistributionAddress() (common.Address, error) {
	return _DefiFMintMinter.Contract.RewardDistributionAddress(&_DefiFMintMinter.CallOpts)
}

// RewardEarned is a free data retrieval call binding the contract method 0x16ba6bf3.
//
// Solidity: function rewardEarned(address _account) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) RewardEarned(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "rewardEarned", _account)
	return *ret0, err
}

// RewardEarned is a free data retrieval call binding the contract method 0x16ba6bf3.
//
// Solidity: function rewardEarned(address _account) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) RewardEarned(_account common.Address) (*big.Int, error) {
	return _DefiFMintMinter.Contract.RewardEarned(&_DefiFMintMinter.CallOpts, _account)
}

// RewardEarned is a free data retrieval call binding the contract method 0x16ba6bf3.
//
// Solidity: function rewardEarned(address _account) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) RewardEarned(_account common.Address) (*big.Int, error) {
	return _DefiFMintMinter.Contract.RewardEarned(&_DefiFMintMinter.CallOpts, _account)
}

// RewardEligibilityRatio4dec is a free data retrieval call binding the contract method 0x50fca4bd.
//
// Solidity: function rewardEligibilityRatio4dec() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) RewardEligibilityRatio4dec(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "rewardEligibilityRatio4dec")
	return *ret0, err
}

// RewardEligibilityRatio4dec is a free data retrieval call binding the contract method 0x50fca4bd.
//
// Solidity: function rewardEligibilityRatio4dec() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) RewardEligibilityRatio4dec() (*big.Int, error) {
	return _DefiFMintMinter.Contract.RewardEligibilityRatio4dec(&_DefiFMintMinter.CallOpts)
}

// RewardEligibilityRatio4dec is a free data retrieval call binding the contract method 0x50fca4bd.
//
// Solidity: function rewardEligibilityRatio4dec() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) RewardEligibilityRatio4dec() (*big.Int, error) {
	return _DefiFMintMinter.Contract.RewardEligibilityRatio4dec(&_DefiFMintMinter.CallOpts)
}

// RewardEpochEnds is a free data retrieval call binding the contract method 0x3ce9b316.
//
// Solidity: function rewardEpochEnds() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) RewardEpochEnds(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "rewardEpochEnds")
	return *ret0, err
}

// RewardEpochEnds is a free data retrieval call binding the contract method 0x3ce9b316.
//
// Solidity: function rewardEpochEnds() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) RewardEpochEnds() (*big.Int, error) {
	return _DefiFMintMinter.Contract.RewardEpochEnds(&_DefiFMintMinter.CallOpts)
}

// RewardEpochEnds is a free data retrieval call binding the contract method 0x3ce9b316.
//
// Solidity: function rewardEpochEnds() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) RewardEpochEnds() (*big.Int, error) {
	return _DefiFMintMinter.Contract.RewardEpochEnds(&_DefiFMintMinter.CallOpts)
}

// RewardEpochLength is a free data retrieval call binding the contract method 0x20a0a0e9.
//
// Solidity: function rewardEpochLength() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) RewardEpochLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "rewardEpochLength")
	return *ret0, err
}

// RewardEpochLength is a free data retrieval call binding the contract method 0x20a0a0e9.
//
// Solidity: function rewardEpochLength() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) RewardEpochLength() (*big.Int, error) {
	return _DefiFMintMinter.Contract.RewardEpochLength(&_DefiFMintMinter.CallOpts)
}

// RewardEpochLength is a free data retrieval call binding the contract method 0x20a0a0e9.
//
// Solidity: function rewardEpochLength() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) RewardEpochLength() (*big.Int, error) {
	return _DefiFMintMinter.Contract.RewardEpochLength(&_DefiFMintMinter.CallOpts)
}

// RewardLastPerToken is a free data retrieval call binding the contract method 0x544bb473.
//
// Solidity: function rewardLastPerToken() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) RewardLastPerToken(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "rewardLastPerToken")
	return *ret0, err
}

// RewardLastPerToken is a free data retrieval call binding the contract method 0x544bb473.
//
// Solidity: function rewardLastPerToken() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) RewardLastPerToken() (*big.Int, error) {
	return _DefiFMintMinter.Contract.RewardLastPerToken(&_DefiFMintMinter.CallOpts)
}

// RewardLastPerToken is a free data retrieval call binding the contract method 0x544bb473.
//
// Solidity: function rewardLastPerToken() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) RewardLastPerToken() (*big.Int, error) {
	return _DefiFMintMinter.Contract.RewardLastPerToken(&_DefiFMintMinter.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) RewardPerToken(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "rewardPerToken")
	return *ret0, err
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) RewardPerToken() (*big.Int, error) {
	return _DefiFMintMinter.Contract.RewardPerToken(&_DefiFMintMinter.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) RewardPerToken() (*big.Int, error) {
	return _DefiFMintMinter.Contract.RewardPerToken(&_DefiFMintMinter.CallOpts)
}

// RewardPerTokenDecimalsCorrection is a free data retrieval call binding the contract method 0x64631d9b.
//
// Solidity: function rewardPerTokenDecimalsCorrection() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) RewardPerTokenDecimalsCorrection(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "rewardPerTokenDecimalsCorrection")
	return *ret0, err
}

// RewardPerTokenDecimalsCorrection is a free data retrieval call binding the contract method 0x64631d9b.
//
// Solidity: function rewardPerTokenDecimalsCorrection() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) RewardPerTokenDecimalsCorrection() (*big.Int, error) {
	return _DefiFMintMinter.Contract.RewardPerTokenDecimalsCorrection(&_DefiFMintMinter.CallOpts)
}

// RewardPerTokenDecimalsCorrection is a free data retrieval call binding the contract method 0x64631d9b.
//
// Solidity: function rewardPerTokenDecimalsCorrection() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) RewardPerTokenDecimalsCorrection() (*big.Int, error) {
	return _DefiFMintMinter.Contract.RewardPerTokenDecimalsCorrection(&_DefiFMintMinter.CallOpts)
}

// RewardPerTokenPaid is a free data retrieval call binding the contract method 0x653a8da1.
//
// Solidity: function rewardPerTokenPaid(address ) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) RewardPerTokenPaid(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "rewardPerTokenPaid", arg0)
	return *ret0, err
}

// RewardPerTokenPaid is a free data retrieval call binding the contract method 0x653a8da1.
//
// Solidity: function rewardPerTokenPaid(address ) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) RewardPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _DefiFMintMinter.Contract.RewardPerTokenPaid(&_DefiFMintMinter.CallOpts, arg0)
}

// RewardPerTokenPaid is a free data retrieval call binding the contract method 0x653a8da1.
//
// Solidity: function rewardPerTokenPaid(address ) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) RewardPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _DefiFMintMinter.Contract.RewardPerTokenPaid(&_DefiFMintMinter.CallOpts, arg0)
}

// RewardPoolAddress is a free data retrieval call binding the contract method 0x845a51ec.
//
// Solidity: function rewardPoolAddress() view returns(address)
func (_DefiFMintMinter *DefiFMintMinterCaller) RewardPoolAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "rewardPoolAddress")
	return *ret0, err
}

// RewardPoolAddress is a free data retrieval call binding the contract method 0x845a51ec.
//
// Solidity: function rewardPoolAddress() view returns(address)
func (_DefiFMintMinter *DefiFMintMinterSession) RewardPoolAddress() (common.Address, error) {
	return _DefiFMintMinter.Contract.RewardPoolAddress(&_DefiFMintMinter.CallOpts)
}

// RewardPoolAddress is a free data retrieval call binding the contract method 0x845a51ec.
//
// Solidity: function rewardPoolAddress() view returns(address)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) RewardPoolAddress() (common.Address, error) {
	return _DefiFMintMinter.Contract.RewardPoolAddress(&_DefiFMintMinter.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) RewardRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "rewardRate")
	return *ret0, err
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) RewardRate() (*big.Int, error) {
	return _DefiFMintMinter.Contract.RewardRate(&_DefiFMintMinter.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) RewardRate() (*big.Int, error) {
	return _DefiFMintMinter.Contract.RewardRate(&_DefiFMintMinter.CallOpts)
}

// RewardStash is a free data retrieval call binding the contract method 0xf2392c8d.
//
// Solidity: function rewardStash(address ) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) RewardStash(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "rewardStash", arg0)
	return *ret0, err
}

// RewardStash is a free data retrieval call binding the contract method 0xf2392c8d.
//
// Solidity: function rewardStash(address ) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) RewardStash(arg0 common.Address) (*big.Int, error) {
	return _DefiFMintMinter.Contract.RewardStash(&_DefiFMintMinter.CallOpts, arg0)
}

// RewardStash is a free data retrieval call binding the contract method 0xf2392c8d.
//
// Solidity: function rewardStash(address ) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) RewardStash(arg0 common.Address) (*big.Int, error) {
	return _DefiFMintMinter.Contract.RewardStash(&_DefiFMintMinter.CallOpts, arg0)
}

// RewardUpdated is a free data retrieval call binding the contract method 0x6e718e04.
//
// Solidity: function rewardUpdated() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) RewardUpdated(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "rewardUpdated")
	return *ret0, err
}

// RewardUpdated is a free data retrieval call binding the contract method 0x6e718e04.
//
// Solidity: function rewardUpdated() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) RewardUpdated() (*big.Int, error) {
	return _DefiFMintMinter.Contract.RewardUpdated(&_DefiFMintMinter.CallOpts)
}

// RewardUpdated is a free data retrieval call binding the contract method 0x6e718e04.
//
// Solidity: function rewardUpdated() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) RewardUpdated() (*big.Int, error) {
	return _DefiFMintMinter.Contract.RewardUpdated(&_DefiFMintMinter.CallOpts)
}

// TokenRegistryAddress is a free data retrieval call binding the contract method 0x5be2aca0.
//
// Solidity: function tokenRegistryAddress() view returns(address)
func (_DefiFMintMinter *DefiFMintMinterCaller) TokenRegistryAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "tokenRegistryAddress")
	return *ret0, err
}

// TokenRegistryAddress is a free data retrieval call binding the contract method 0x5be2aca0.
//
// Solidity: function tokenRegistryAddress() view returns(address)
func (_DefiFMintMinter *DefiFMintMinterSession) TokenRegistryAddress() (common.Address, error) {
	return _DefiFMintMinter.Contract.TokenRegistryAddress(&_DefiFMintMinter.CallOpts)
}

// TokenRegistryAddress is a free data retrieval call binding the contract method 0x5be2aca0.
//
// Solidity: function tokenRegistryAddress() view returns(address)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) TokenRegistryAddress() (common.Address, error) {
	return _DefiFMintMinter.Contract.TokenRegistryAddress(&_DefiFMintMinter.CallOpts)
}

// TokenValue is a free data retrieval call binding the contract method 0xf1821783.
//
// Solidity: function tokenValue(address _token, uint256 _amount) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) TokenValue(opts *bind.CallOpts, _token common.Address, _amount *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DefiFMintMinter.contract.Call(opts, out, "tokenValue", _token, _amount)
	return *ret0, err
}

// TokenValue is a free data retrieval call binding the contract method 0xf1821783.
//
// Solidity: function tokenValue(address _token, uint256 _amount) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) TokenValue(_token common.Address, _amount *big.Int) (*big.Int, error) {
	return _DefiFMintMinter.Contract.TokenValue(&_DefiFMintMinter.CallOpts, _token, _amount)
}

// TokenValue is a free data retrieval call binding the contract method 0xf1821783.
//
// Solidity: function tokenValue(address _token, uint256 _amount) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) TokenValue(_token common.Address, _amount *big.Int) (*big.Int, error) {
	return _DefiFMintMinter.Contract.TokenValue(&_DefiFMintMinter.CallOpts, _token, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _token, uint256 _amount) returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterTransactor) Deposit(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.contract.Transact(opts, "deposit", _token, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _token, uint256 _amount) returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) Deposit(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.Deposit(&_DefiFMintMinter.TransactOpts, _token, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _token, uint256 _amount) returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterTransactorSession) Deposit(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.Deposit(&_DefiFMintMinter.TransactOpts, _token, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _token, uint256 _amount) returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterTransactor) Mint(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.contract.Transact(opts, "mint", _token, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _token, uint256 _amount) returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) Mint(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.Mint(&_DefiFMintMinter.TransactOpts, _token, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _token, uint256 _amount) returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterTransactorSession) Mint(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.Mint(&_DefiFMintMinter.TransactOpts, _token, _amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DefiFMintMinter *DefiFMintMinterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefiFMintMinter.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DefiFMintMinter *DefiFMintMinterSession) RenounceOwnership() (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.RenounceOwnership(&_DefiFMintMinter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DefiFMintMinter *DefiFMintMinterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.RenounceOwnership(&_DefiFMintMinter.TransactOpts)
}

// Repay is a paid mutator transaction binding the contract method 0x22867d78.
//
// Solidity: function repay(address _token, uint256 _amount) returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterTransactor) Repay(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.contract.Transact(opts, "repay", _token, _amount)
}

// Repay is a paid mutator transaction binding the contract method 0x22867d78.
//
// Solidity: function repay(address _token, uint256 _amount) returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) Repay(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.Repay(&_DefiFMintMinter.TransactOpts, _token, _amount)
}

// Repay is a paid mutator transaction binding the contract method 0x22867d78.
//
// Solidity: function repay(address _token, uint256 _amount) returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterTransactorSession) Repay(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.Repay(&_DefiFMintMinter.TransactOpts, _token, _amount)
}

// RewardClaim is a paid mutator transaction binding the contract method 0x6409f921.
//
// Solidity: function rewardClaim() returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterTransactor) RewardClaim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefiFMintMinter.contract.Transact(opts, "rewardClaim")
}

// RewardClaim is a paid mutator transaction binding the contract method 0x6409f921.
//
// Solidity: function rewardClaim() returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) RewardClaim() (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.RewardClaim(&_DefiFMintMinter.TransactOpts)
}

// RewardClaim is a paid mutator transaction binding the contract method 0x6409f921.
//
// Solidity: function rewardClaim() returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterTransactorSession) RewardClaim() (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.RewardClaim(&_DefiFMintMinter.TransactOpts)
}

// RewardNotifyAmount is a paid mutator transaction binding the contract method 0x2eac3b87.
//
// Solidity: function rewardNotifyAmount(uint256 reward) returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterTransactor) RewardNotifyAmount(opts *bind.TransactOpts, reward *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.contract.Transact(opts, "rewardNotifyAmount", reward)
}

// RewardNotifyAmount is a paid mutator transaction binding the contract method 0x2eac3b87.
//
// Solidity: function rewardNotifyAmount(uint256 reward) returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) RewardNotifyAmount(reward *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.RewardNotifyAmount(&_DefiFMintMinter.TransactOpts, reward)
}

// RewardNotifyAmount is a paid mutator transaction binding the contract method 0x2eac3b87.
//
// Solidity: function rewardNotifyAmount(uint256 reward) returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterTransactorSession) RewardNotifyAmount(reward *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.RewardNotifyAmount(&_DefiFMintMinter.TransactOpts, reward)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DefiFMintMinter *DefiFMintMinterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DefiFMintMinter.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DefiFMintMinter *DefiFMintMinterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.TransferOwnership(&_DefiFMintMinter.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DefiFMintMinter *DefiFMintMinterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.TransferOwnership(&_DefiFMintMinter.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _token, uint256 _amount) returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterTransactor) Withdraw(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.contract.Transact(opts, "withdraw", _token, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _token, uint256 _amount) returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) Withdraw(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.Withdraw(&_DefiFMintMinter.TransactOpts, _token, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _token, uint256 _amount) returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterTransactorSession) Withdraw(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.Withdraw(&_DefiFMintMinter.TransactOpts, _token, _amount)
}

// DefiFMintMinterDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the DefiFMintMinter contract.
type DefiFMintMinterDepositedIterator struct {
	Event *DefiFMintMinterDeposited // Event containing the contract specifics and raw log

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
func (it *DefiFMintMinterDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefiFMintMinterDeposited)
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
		it.Event = new(DefiFMintMinterDeposited)
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
func (it *DefiFMintMinterDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefiFMintMinterDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefiFMintMinterDeposited represents a Deposited event raised by the DefiFMintMinter contract.
type DefiFMintMinterDeposited struct {
	Token  common.Address
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(address indexed token, address indexed user, uint256 amount)
func (_DefiFMintMinter *DefiFMintMinterFilterer) FilterDeposited(opts *bind.FilterOpts, token []common.Address, user []common.Address) (*DefiFMintMinterDepositedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DefiFMintMinter.contract.FilterLogs(opts, "Deposited", tokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return &DefiFMintMinterDepositedIterator{contract: _DefiFMintMinter.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(address indexed token, address indexed user, uint256 amount)
func (_DefiFMintMinter *DefiFMintMinterFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *DefiFMintMinterDeposited, token []common.Address, user []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DefiFMintMinter.contract.WatchLogs(opts, "Deposited", tokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefiFMintMinterDeposited)
				if err := _DefiFMintMinter.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(address indexed token, address indexed user, uint256 amount)
func (_DefiFMintMinter *DefiFMintMinterFilterer) ParseDeposited(log types.Log) (*DefiFMintMinterDeposited, error) {
	event := new(DefiFMintMinterDeposited)
	if err := _DefiFMintMinter.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DefiFMintMinterMintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the DefiFMintMinter contract.
type DefiFMintMinterMintedIterator struct {
	Event *DefiFMintMinterMinted // Event containing the contract specifics and raw log

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
func (it *DefiFMintMinterMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefiFMintMinterMinted)
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
		it.Event = new(DefiFMintMinterMinted)
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
func (it *DefiFMintMinterMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefiFMintMinterMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefiFMintMinterMinted represents a Minted event raised by the DefiFMintMinter contract.
type DefiFMintMinterMinted struct {
	Token  common.Address
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0x9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f0.
//
// Solidity: event Minted(address indexed token, address indexed user, uint256 amount)
func (_DefiFMintMinter *DefiFMintMinterFilterer) FilterMinted(opts *bind.FilterOpts, token []common.Address, user []common.Address) (*DefiFMintMinterMintedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DefiFMintMinter.contract.FilterLogs(opts, "Minted", tokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return &DefiFMintMinterMintedIterator{contract: _DefiFMintMinter.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0x9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f0.
//
// Solidity: event Minted(address indexed token, address indexed user, uint256 amount)
func (_DefiFMintMinter *DefiFMintMinterFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *DefiFMintMinterMinted, token []common.Address, user []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DefiFMintMinter.contract.WatchLogs(opts, "Minted", tokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefiFMintMinterMinted)
				if err := _DefiFMintMinter.contract.UnpackLog(event, "Minted", log); err != nil {
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

// ParseMinted is a log parse operation binding the contract event 0x9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f0.
//
// Solidity: event Minted(address indexed token, address indexed user, uint256 amount)
func (_DefiFMintMinter *DefiFMintMinterFilterer) ParseMinted(log types.Log) (*DefiFMintMinterMinted, error) {
	event := new(DefiFMintMinterMinted)
	if err := _DefiFMintMinter.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DefiFMintMinterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DefiFMintMinter contract.
type DefiFMintMinterOwnershipTransferredIterator struct {
	Event *DefiFMintMinterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DefiFMintMinterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefiFMintMinterOwnershipTransferred)
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
		it.Event = new(DefiFMintMinterOwnershipTransferred)
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
func (it *DefiFMintMinterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefiFMintMinterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefiFMintMinterOwnershipTransferred represents a OwnershipTransferred event raised by the DefiFMintMinter contract.
type DefiFMintMinterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DefiFMintMinter *DefiFMintMinterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DefiFMintMinterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DefiFMintMinter.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DefiFMintMinterOwnershipTransferredIterator{contract: _DefiFMintMinter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DefiFMintMinter *DefiFMintMinterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DefiFMintMinterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DefiFMintMinter.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefiFMintMinterOwnershipTransferred)
				if err := _DefiFMintMinter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_DefiFMintMinter *DefiFMintMinterFilterer) ParseOwnershipTransferred(log types.Log) (*DefiFMintMinterOwnershipTransferred, error) {
	event := new(DefiFMintMinterOwnershipTransferred)
	if err := _DefiFMintMinter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DefiFMintMinterRepaidIterator is returned from FilterRepaid and is used to iterate over the raw logs and unpacked data for Repaid events raised by the DefiFMintMinter contract.
type DefiFMintMinterRepaidIterator struct {
	Event *DefiFMintMinterRepaid // Event containing the contract specifics and raw log

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
func (it *DefiFMintMinterRepaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefiFMintMinterRepaid)
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
		it.Event = new(DefiFMintMinterRepaid)
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
func (it *DefiFMintMinterRepaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefiFMintMinterRepaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefiFMintMinterRepaid represents a Repaid event raised by the DefiFMintMinter contract.
type DefiFMintMinterRepaid struct {
	Token  common.Address
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRepaid is a free log retrieval operation binding the contract event 0x0a3fbbea70e93f2daafa3102f5c9a1b8315e6d7a1e43e4bc020bc1162327470a.
//
// Solidity: event Repaid(address indexed token, address indexed user, uint256 amount)
func (_DefiFMintMinter *DefiFMintMinterFilterer) FilterRepaid(opts *bind.FilterOpts, token []common.Address, user []common.Address) (*DefiFMintMinterRepaidIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DefiFMintMinter.contract.FilterLogs(opts, "Repaid", tokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return &DefiFMintMinterRepaidIterator{contract: _DefiFMintMinter.contract, event: "Repaid", logs: logs, sub: sub}, nil
}

// WatchRepaid is a free log subscription operation binding the contract event 0x0a3fbbea70e93f2daafa3102f5c9a1b8315e6d7a1e43e4bc020bc1162327470a.
//
// Solidity: event Repaid(address indexed token, address indexed user, uint256 amount)
func (_DefiFMintMinter *DefiFMintMinterFilterer) WatchRepaid(opts *bind.WatchOpts, sink chan<- *DefiFMintMinterRepaid, token []common.Address, user []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DefiFMintMinter.contract.WatchLogs(opts, "Repaid", tokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefiFMintMinterRepaid)
				if err := _DefiFMintMinter.contract.UnpackLog(event, "Repaid", log); err != nil {
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

// ParseRepaid is a log parse operation binding the contract event 0x0a3fbbea70e93f2daafa3102f5c9a1b8315e6d7a1e43e4bc020bc1162327470a.
//
// Solidity: event Repaid(address indexed token, address indexed user, uint256 amount)
func (_DefiFMintMinter *DefiFMintMinterFilterer) ParseRepaid(log types.Log) (*DefiFMintMinterRepaid, error) {
	event := new(DefiFMintMinterRepaid)
	if err := _DefiFMintMinter.contract.UnpackLog(event, "Repaid", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DefiFMintMinterRewardAddedIterator is returned from FilterRewardAdded and is used to iterate over the raw logs and unpacked data for RewardAdded events raised by the DefiFMintMinter contract.
type DefiFMintMinterRewardAddedIterator struct {
	Event *DefiFMintMinterRewardAdded // Event containing the contract specifics and raw log

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
func (it *DefiFMintMinterRewardAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefiFMintMinterRewardAdded)
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
		it.Event = new(DefiFMintMinterRewardAdded)
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
func (it *DefiFMintMinterRewardAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefiFMintMinterRewardAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefiFMintMinterRewardAdded represents a RewardAdded event raised by the DefiFMintMinter contract.
type DefiFMintMinterRewardAdded struct {
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardAdded is a free log retrieval operation binding the contract event 0xde88a922e0d3b88b24e9623efeb464919c6bf9f66857a65e2bfcf2ce87a9433d.
//
// Solidity: event RewardAdded(uint256 reward)
func (_DefiFMintMinter *DefiFMintMinterFilterer) FilterRewardAdded(opts *bind.FilterOpts) (*DefiFMintMinterRewardAddedIterator, error) {

	logs, sub, err := _DefiFMintMinter.contract.FilterLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return &DefiFMintMinterRewardAddedIterator{contract: _DefiFMintMinter.contract, event: "RewardAdded", logs: logs, sub: sub}, nil
}

// WatchRewardAdded is a free log subscription operation binding the contract event 0xde88a922e0d3b88b24e9623efeb464919c6bf9f66857a65e2bfcf2ce87a9433d.
//
// Solidity: event RewardAdded(uint256 reward)
func (_DefiFMintMinter *DefiFMintMinterFilterer) WatchRewardAdded(opts *bind.WatchOpts, sink chan<- *DefiFMintMinterRewardAdded) (event.Subscription, error) {

	logs, sub, err := _DefiFMintMinter.contract.WatchLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefiFMintMinterRewardAdded)
				if err := _DefiFMintMinter.contract.UnpackLog(event, "RewardAdded", log); err != nil {
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
func (_DefiFMintMinter *DefiFMintMinterFilterer) ParseRewardAdded(log types.Log) (*DefiFMintMinterRewardAdded, error) {
	event := new(DefiFMintMinterRewardAdded)
	if err := _DefiFMintMinter.contract.UnpackLog(event, "RewardAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DefiFMintMinterRewardPaidIterator is returned from FilterRewardPaid and is used to iterate over the raw logs and unpacked data for RewardPaid events raised by the DefiFMintMinter contract.
type DefiFMintMinterRewardPaidIterator struct {
	Event *DefiFMintMinterRewardPaid // Event containing the contract specifics and raw log

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
func (it *DefiFMintMinterRewardPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefiFMintMinterRewardPaid)
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
		it.Event = new(DefiFMintMinterRewardPaid)
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
func (it *DefiFMintMinterRewardPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefiFMintMinterRewardPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefiFMintMinterRewardPaid represents a RewardPaid event raised by the DefiFMintMinter contract.
type DefiFMintMinterRewardPaid struct {
	User   common.Address
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardPaid is a free log retrieval operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_DefiFMintMinter *DefiFMintMinterFilterer) FilterRewardPaid(opts *bind.FilterOpts, user []common.Address) (*DefiFMintMinterRewardPaidIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DefiFMintMinter.contract.FilterLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return &DefiFMintMinterRewardPaidIterator{contract: _DefiFMintMinter.contract, event: "RewardPaid", logs: logs, sub: sub}, nil
}

// WatchRewardPaid is a free log subscription operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_DefiFMintMinter *DefiFMintMinterFilterer) WatchRewardPaid(opts *bind.WatchOpts, sink chan<- *DefiFMintMinterRewardPaid, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DefiFMintMinter.contract.WatchLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefiFMintMinterRewardPaid)
				if err := _DefiFMintMinter.contract.UnpackLog(event, "RewardPaid", log); err != nil {
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
func (_DefiFMintMinter *DefiFMintMinterFilterer) ParseRewardPaid(log types.Log) (*DefiFMintMinterRewardPaid, error) {
	event := new(DefiFMintMinterRewardPaid)
	if err := _DefiFMintMinter.contract.UnpackLog(event, "RewardPaid", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DefiFMintMinterWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the DefiFMintMinter contract.
type DefiFMintMinterWithdrawnIterator struct {
	Event *DefiFMintMinterWithdrawn // Event containing the contract specifics and raw log

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
func (it *DefiFMintMinterWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefiFMintMinterWithdrawn)
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
		it.Event = new(DefiFMintMinterWithdrawn)
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
func (it *DefiFMintMinterWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefiFMintMinterWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefiFMintMinterWithdrawn represents a Withdrawn event raised by the DefiFMintMinter contract.
type DefiFMintMinterWithdrawn struct {
	Token  common.Address
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed token, address indexed user, uint256 amount)
func (_DefiFMintMinter *DefiFMintMinterFilterer) FilterWithdrawn(opts *bind.FilterOpts, token []common.Address, user []common.Address) (*DefiFMintMinterWithdrawnIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DefiFMintMinter.contract.FilterLogs(opts, "Withdrawn", tokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return &DefiFMintMinterWithdrawnIterator{contract: _DefiFMintMinter.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed token, address indexed user, uint256 amount)
func (_DefiFMintMinter *DefiFMintMinterFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *DefiFMintMinterWithdrawn, token []common.Address, user []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _DefiFMintMinter.contract.WatchLogs(opts, "Withdrawn", tokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefiFMintMinterWithdrawn)
				if err := _DefiFMintMinter.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed token, address indexed user, uint256 amount)
func (_DefiFMintMinter *DefiFMintMinterFilterer) ParseWithdrawn(log types.Log) (*DefiFMintMinterWithdrawn, error) {
	event := new(DefiFMintMinterWithdrawn)
	if err := _DefiFMintMinter.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	return event, nil
}
