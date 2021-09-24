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

// DefiFMintMinterMetaData contains all meta data concerning the DefiFMintMinter contract.
var DefiFMintMinterMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ratio4dec\",\"type\":\"uint256\"}],\"name\":\"CollateralLowestDebtRatioChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee4dec\",\"type\":\"uint256\"}],\"name\":\"MintFeeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Repaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ratio4dec\",\"type\":\"uint256\"}],\"name\":\"RewardEligibilityRatioChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_DEBT_EXCEEDED\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_DEPOSIT_PROHIBITED\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_LOW_ALLOWANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_LOW_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_LOW_BALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_LOW_COLLATERAL_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_MINTING_PROHIBITED\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_NOT_AUTHORIZED\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_NO_COLLATERAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_NO_ERROR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_NO_REWARD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_NO_VALUE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_REWARDS_DEPLETED\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_REWARDS_EARLY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_REWARDS_NONE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_REWARD_CLAIM_REJECTED\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_ZERO_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"addressProvider\",\"outputs\":[{\"internalType\":\"contractIFantomMintAddressProvider\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ratio4dec\",\"type\":\"uint256\"}],\"name\":\"cfgSetLowestCollateralRatio\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee4dec\",\"type\":\"uint256\"}],\"name\":\"cfgSetMintFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ratio4dec\",\"type\":\"uint256\"}],\"name\":\"cfgSetRewardEligibilityRatio\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"collateralCanDecrease\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"collateralLowestDebtRatio4dec\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"collateralRatioDecimalsCorrection\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"debtCanIncrease\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fMintFee4dec\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fMintFeeDigitsCorrection\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"feePool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_ratio\",\"type\":\"uint256\"}],\"name\":\"maxToMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_ratio\",\"type\":\"uint256\"}],\"name\":\"maxToWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_ratio\",\"type\":\"uint256\"}],\"name\":\"minToDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_ratio\",\"type\":\"uint256\"}],\"name\":\"mintMax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mustDeposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mustMint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_ratio\",\"type\":\"uint256\"}],\"name\":\"mustMintMax\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mustRepay\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"mustRepayMax\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mustWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_ratio\",\"type\":\"uint256\"}],\"name\":\"mustWithdrawMax\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"repay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"repayMax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"rewardCanClaim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardEligibilityRatio4dec\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"rewardIsEligible\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_ratio\",\"type\":\"uint256\"}],\"name\":\"withdrawMax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressProvider\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCollateralLowestDebtRatio4dec\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRewardEligibilityRatio4dec\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFMintFee4dec\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCollateralPool\",\"outputs\":[{\"internalType\":\"contractIFantomDeFiTokenStorage\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDebtPool\",\"outputs\":[{\"internalType\":\"contractIFantomDeFiTokenStorage\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"canDeposit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"canMint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"checkCollateralCanDecrease\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"checkDebtCanIncrease\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_add\",\"type\":\"uint256\"}],\"name\":\"debtValueOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_sub\",\"type\":\"uint256\"}],\"name\":\"collateralValueOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_ratio\",\"type\":\"uint256\"}],\"name\":\"getMaxToWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_ratio\",\"type\":\"uint256\"}],\"name\":\"getMaxToMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"rewardUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getExtendedPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_digits\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DefiFMintMinterABI is the input ABI used to generate the binding from.
// Deprecated: Use DefiFMintMinterMetaData.ABI instead.
var DefiFMintMinterABI = DefiFMintMinterMetaData.ABI

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
func (_DefiFMintMinter *DefiFMintMinterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
func (_DefiFMintMinter *DefiFMintMinterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
	var out []interface{}
	err := _DefiFMintMinter.contract.Call(opts, &out, "ERR_DEBT_EXCEEDED")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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

// ERRDEPOSITPROHIBITED is a free data retrieval call binding the contract method 0x2bfcc373.
//
// Solidity: function ERR_DEPOSIT_PROHIBITED() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) ERRDEPOSITPROHIBITED(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DefiFMintMinter.contract.Call(opts, &out, "ERR_DEPOSIT_PROHIBITED")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ERRDEPOSITPROHIBITED is a free data retrieval call binding the contract method 0x2bfcc373.
//
// Solidity: function ERR_DEPOSIT_PROHIBITED() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) ERRDEPOSITPROHIBITED() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRDEPOSITPROHIBITED(&_DefiFMintMinter.CallOpts)
}

// ERRDEPOSITPROHIBITED is a free data retrieval call binding the contract method 0x2bfcc373.
//
// Solidity: function ERR_DEPOSIT_PROHIBITED() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) ERRDEPOSITPROHIBITED() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRDEPOSITPROHIBITED(&_DefiFMintMinter.CallOpts)
}

// ERRLOWALLOWANCE is a free data retrieval call binding the contract method 0x911fc3f1.
//
// Solidity: function ERR_LOW_ALLOWANCE() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) ERRLOWALLOWANCE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DefiFMintMinter.contract.Call(opts, &out, "ERR_LOW_ALLOWANCE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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

// ERRLOWAMOUNT is a free data retrieval call binding the contract method 0xc7ea4889.
//
// Solidity: function ERR_LOW_AMOUNT() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) ERRLOWAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DefiFMintMinter.contract.Call(opts, &out, "ERR_LOW_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ERRLOWAMOUNT is a free data retrieval call binding the contract method 0xc7ea4889.
//
// Solidity: function ERR_LOW_AMOUNT() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) ERRLOWAMOUNT() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRLOWAMOUNT(&_DefiFMintMinter.CallOpts)
}

// ERRLOWAMOUNT is a free data retrieval call binding the contract method 0xc7ea4889.
//
// Solidity: function ERR_LOW_AMOUNT() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) ERRLOWAMOUNT() (*big.Int, error) {
	return _DefiFMintMinter.Contract.ERRLOWAMOUNT(&_DefiFMintMinter.CallOpts)
}

// ERRLOWBALANCE is a free data retrieval call binding the contract method 0x1ac919b0.
//
// Solidity: function ERR_LOW_BALANCE() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) ERRLOWBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DefiFMintMinter.contract.Call(opts, &out, "ERR_LOW_BALANCE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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
	var out []interface{}
	err := _DefiFMintMinter.contract.Call(opts, &out, "ERR_LOW_COLLATERAL_RATIO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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
	var out []interface{}
	err := _DefiFMintMinter.contract.Call(opts, &out, "ERR_MINTING_PROHIBITED")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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
	var out []interface{}
	err := _DefiFMintMinter.contract.Call(opts, &out, "ERR_NOT_AUTHORIZED")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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
	var out []interface{}
	err := _DefiFMintMinter.contract.Call(opts, &out, "ERR_NO_COLLATERAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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
	var out []interface{}
	err := _DefiFMintMinter.contract.Call(opts, &out, "ERR_NO_ERROR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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
	var out []interface{}
	err := _DefiFMintMinter.contract.Call(opts, &out, "ERR_NO_REWARD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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
	var out []interface{}
	err := _DefiFMintMinter.contract.Call(opts, &out, "ERR_NO_VALUE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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
	var out []interface{}
	err := _DefiFMintMinter.contract.Call(opts, &out, "ERR_REWARDS_DEPLETED")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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
	var out []interface{}
	err := _DefiFMintMinter.contract.Call(opts, &out, "ERR_REWARDS_EARLY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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
	var out []interface{}
	err := _DefiFMintMinter.contract.Call(opts, &out, "ERR_REWARDS_NONE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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
	var out []interface{}
	err := _DefiFMintMinter.contract.Call(opts, &out, "ERR_REWARD_CLAIM_REJECTED")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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
	var out []interface{}
	err := _DefiFMintMinter.contract.Call(opts, &out, "ERR_ZERO_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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
	var out []interface{}
	err := _DefiFMintMinter.contract.Call(opts, &out, "addressProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

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

// CanDeposit is a free data retrieval call binding the contract method 0x4bf0d331.
//
// Solidity: function canDeposit(address _token) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterCaller) CanDeposit(opts *bind.CallOpts, _token common.Address) (bool, error) {
	var out []interface{}
	err := _DefiFMintMinter.contract.Call(opts, &out, "canDeposit", _token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanDeposit is a free data retrieval call binding the contract method 0x4bf0d331.
//
// Solidity: function canDeposit(address _token) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterSession) CanDeposit(_token common.Address) (bool, error) {
	return _DefiFMintMinter.Contract.CanDeposit(&_DefiFMintMinter.CallOpts, _token)
}

// CanDeposit is a free data retrieval call binding the contract method 0x4bf0d331.
//
// Solidity: function canDeposit(address _token) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) CanDeposit(_token common.Address) (bool, error) {
	return _DefiFMintMinter.Contract.CanDeposit(&_DefiFMintMinter.CallOpts, _token)
}

// CanMint is a free data retrieval call binding the contract method 0xc2ba4744.
//
// Solidity: function canMint(address _token) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterCaller) CanMint(opts *bind.CallOpts, _token common.Address) (bool, error) {
	var out []interface{}
	err := _DefiFMintMinter.contract.Call(opts, &out, "canMint", _token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanMint is a free data retrieval call binding the contract method 0xc2ba4744.
//
// Solidity: function canMint(address _token) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterSession) CanMint(_token common.Address) (bool, error) {
	return _DefiFMintMinter.Contract.CanMint(&_DefiFMintMinter.CallOpts, _token)
}

// CanMint is a free data retrieval call binding the contract method 0xc2ba4744.
//
// Solidity: function canMint(address _token) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) CanMint(_token common.Address) (bool, error) {
	return _DefiFMintMinter.Contract.CanMint(&_DefiFMintMinter.CallOpts, _token)
}

// CheckCollateralCanDecrease is a free data retrieval call binding the contract method 0xa03a2689.
//
// Solidity: function checkCollateralCanDecrease(address _account, address _token, uint256 _amount) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterCaller) CheckCollateralCanDecrease(opts *bind.CallOpts, _account common.Address, _token common.Address, _amount *big.Int) (bool, error) {
	var out []interface{}
	err := _DefiFMintMinter.contract.Call(opts, &out, "checkCollateralCanDecrease", _account, _token, _amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckCollateralCanDecrease is a free data retrieval call binding the contract method 0xa03a2689.
//
// Solidity: function checkCollateralCanDecrease(address _account, address _token, uint256 _amount) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterSession) CheckCollateralCanDecrease(_account common.Address, _token common.Address, _amount *big.Int) (bool, error) {
	return _DefiFMintMinter.Contract.CheckCollateralCanDecrease(&_DefiFMintMinter.CallOpts, _account, _token, _amount)
}

// CheckCollateralCanDecrease is a free data retrieval call binding the contract method 0xa03a2689.
//
// Solidity: function checkCollateralCanDecrease(address _account, address _token, uint256 _amount) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) CheckCollateralCanDecrease(_account common.Address, _token common.Address, _amount *big.Int) (bool, error) {
	return _DefiFMintMinter.Contract.CheckCollateralCanDecrease(&_DefiFMintMinter.CallOpts, _account, _token, _amount)
}

// CheckDebtCanIncrease is a free data retrieval call binding the contract method 0x4764efb0.
//
// Solidity: function checkDebtCanIncrease(address _account, address _token, uint256 _amount) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterCaller) CheckDebtCanIncrease(opts *bind.CallOpts, _account common.Address, _token common.Address, _amount *big.Int) (bool, error) {
	var out []interface{}
	err := _DefiFMintMinter.contract.Call(opts, &out, "checkDebtCanIncrease", _account, _token, _amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckDebtCanIncrease is a free data retrieval call binding the contract method 0x4764efb0.
//
// Solidity: function checkDebtCanIncrease(address _account, address _token, uint256 _amount) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterSession) CheckDebtCanIncrease(_account common.Address, _token common.Address, _amount *big.Int) (bool, error) {
	return _DefiFMintMinter.Contract.CheckDebtCanIncrease(&_DefiFMintMinter.CallOpts, _account, _token, _amount)
}

// CheckDebtCanIncrease is a free data retrieval call binding the contract method 0x4764efb0.
//
// Solidity: function checkDebtCanIncrease(address _account, address _token, uint256 _amount) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) CheckDebtCanIncrease(_account common.Address, _token common.Address, _amount *big.Int) (bool, error) {
	return _DefiFMintMinter.Contract.CheckDebtCanIncrease(&_DefiFMintMinter.CallOpts, _account, _token, _amount)
}

// CollateralCanDecrease is a free data retrieval call binding the contract method 0xf4305a99.
//
// Solidity: function collateralCanDecrease(address _account, address _token, uint256 _amount) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterCaller) CollateralCanDecrease(opts *bind.CallOpts, _account common.Address, _token common.Address, _amount *big.Int) (bool, error) {
	var out []interface{}
	err := _DefiFMintMinter.contract.Call(opts, &out, "collateralCanDecrease", _account, _token, _amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

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
	var out []interface{}
	err := _DefiFMintMinter.contract.Call(opts, &out, "collateralLowestDebtRatio4dec")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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
	var out []interface{}
	err := _DefiFMintMinter.contract.Call(opts, &out, "collateralRatioDecimalsCorrection")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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

// CollateralValueOf is a free data retrieval call binding the contract method 0x5a13fd77.
//
// Solidity: function collateralValueOf(address _account, address _token, uint256 _sub) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) CollateralValueOf(opts *bind.CallOpts, _account common.Address, _token common.Address, _sub *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DefiFMintMinter.contract.Call(opts, &out, "collateralValueOf", _account, _token, _sub)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollateralValueOf is a free data retrieval call binding the contract method 0x5a13fd77.
//
// Solidity: function collateralValueOf(address _account, address _token, uint256 _sub) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) CollateralValueOf(_account common.Address, _token common.Address, _sub *big.Int) (*big.Int, error) {
	return _DefiFMintMinter.Contract.CollateralValueOf(&_DefiFMintMinter.CallOpts, _account, _token, _sub)
}

// CollateralValueOf is a free data retrieval call binding the contract method 0x5a13fd77.
//
// Solidity: function collateralValueOf(address _account, address _token, uint256 _sub) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) CollateralValueOf(_account common.Address, _token common.Address, _sub *big.Int) (*big.Int, error) {
	return _DefiFMintMinter.Contract.CollateralValueOf(&_DefiFMintMinter.CallOpts, _account, _token, _sub)
}

// DebtCanIncrease is a free data retrieval call binding the contract method 0x905ca247.
//
// Solidity: function debtCanIncrease(address _account, address _token, uint256 _amount) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterCaller) DebtCanIncrease(opts *bind.CallOpts, _account common.Address, _token common.Address, _amount *big.Int) (bool, error) {
	var out []interface{}
	err := _DefiFMintMinter.contract.Call(opts, &out, "debtCanIncrease", _account, _token, _amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

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

// DebtValueOf is a free data retrieval call binding the contract method 0xb36607e7.
//
// Solidity: function debtValueOf(address _account, address _token, uint256 _add) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) DebtValueOf(opts *bind.CallOpts, _account common.Address, _token common.Address, _add *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DefiFMintMinter.contract.Call(opts, &out, "debtValueOf", _account, _token, _add)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtValueOf is a free data retrieval call binding the contract method 0xb36607e7.
//
// Solidity: function debtValueOf(address _account, address _token, uint256 _add) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) DebtValueOf(_account common.Address, _token common.Address, _add *big.Int) (*big.Int, error) {
	return _DefiFMintMinter.Contract.DebtValueOf(&_DefiFMintMinter.CallOpts, _account, _token, _add)
}

// DebtValueOf is a free data retrieval call binding the contract method 0xb36607e7.
//
// Solidity: function debtValueOf(address _account, address _token, uint256 _add) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) DebtValueOf(_account common.Address, _token common.Address, _add *big.Int) (*big.Int, error) {
	return _DefiFMintMinter.Contract.DebtValueOf(&_DefiFMintMinter.CallOpts, _account, _token, _add)
}

// FMintFee4dec is a free data retrieval call binding the contract method 0x572511bc.
//
// Solidity: function fMintFee4dec() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) FMintFee4dec(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DefiFMintMinter.contract.Call(opts, &out, "fMintFee4dec")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FMintFee4dec is a free data retrieval call binding the contract method 0x572511bc.
//
// Solidity: function fMintFee4dec() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) FMintFee4dec() (*big.Int, error) {
	return _DefiFMintMinter.Contract.FMintFee4dec(&_DefiFMintMinter.CallOpts)
}

// FMintFee4dec is a free data retrieval call binding the contract method 0x572511bc.
//
// Solidity: function fMintFee4dec() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) FMintFee4dec() (*big.Int, error) {
	return _DefiFMintMinter.Contract.FMintFee4dec(&_DefiFMintMinter.CallOpts)
}

// FMintFeeDigitsCorrection is a free data retrieval call binding the contract method 0xcbf02fd5.
//
// Solidity: function fMintFeeDigitsCorrection() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) FMintFeeDigitsCorrection(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DefiFMintMinter.contract.Call(opts, &out, "fMintFeeDigitsCorrection")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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

// FeePool is a free data retrieval call binding the contract method 0xfd26fef1.
//
// Solidity: function feePool(address ) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) FeePool(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DefiFMintMinter.contract.Call(opts, &out, "feePool", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeePool is a free data retrieval call binding the contract method 0xfd26fef1.
//
// Solidity: function feePool(address ) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) FeePool(arg0 common.Address) (*big.Int, error) {
	return _DefiFMintMinter.Contract.FeePool(&_DefiFMintMinter.CallOpts, arg0)
}

// FeePool is a free data retrieval call binding the contract method 0xfd26fef1.
//
// Solidity: function feePool(address ) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) FeePool(arg0 common.Address) (*big.Int, error) {
	return _DefiFMintMinter.Contract.FeePool(&_DefiFMintMinter.CallOpts, arg0)
}

// GetCollateralLowestDebtRatio4dec is a free data retrieval call binding the contract method 0xd65cb5aa.
//
// Solidity: function getCollateralLowestDebtRatio4dec() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) GetCollateralLowestDebtRatio4dec(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DefiFMintMinter.contract.Call(opts, &out, "getCollateralLowestDebtRatio4dec")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCollateralLowestDebtRatio4dec is a free data retrieval call binding the contract method 0xd65cb5aa.
//
// Solidity: function getCollateralLowestDebtRatio4dec() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) GetCollateralLowestDebtRatio4dec() (*big.Int, error) {
	return _DefiFMintMinter.Contract.GetCollateralLowestDebtRatio4dec(&_DefiFMintMinter.CallOpts)
}

// GetCollateralLowestDebtRatio4dec is a free data retrieval call binding the contract method 0xd65cb5aa.
//
// Solidity: function getCollateralLowestDebtRatio4dec() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) GetCollateralLowestDebtRatio4dec() (*big.Int, error) {
	return _DefiFMintMinter.Contract.GetCollateralLowestDebtRatio4dec(&_DefiFMintMinter.CallOpts)
}

// GetCollateralPool is a free data retrieval call binding the contract method 0x73c9641d.
//
// Solidity: function getCollateralPool() view returns(address)
func (_DefiFMintMinter *DefiFMintMinterCaller) GetCollateralPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DefiFMintMinter.contract.Call(opts, &out, "getCollateralPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCollateralPool is a free data retrieval call binding the contract method 0x73c9641d.
//
// Solidity: function getCollateralPool() view returns(address)
func (_DefiFMintMinter *DefiFMintMinterSession) GetCollateralPool() (common.Address, error) {
	return _DefiFMintMinter.Contract.GetCollateralPool(&_DefiFMintMinter.CallOpts)
}

// GetCollateralPool is a free data retrieval call binding the contract method 0x73c9641d.
//
// Solidity: function getCollateralPool() view returns(address)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) GetCollateralPool() (common.Address, error) {
	return _DefiFMintMinter.Contract.GetCollateralPool(&_DefiFMintMinter.CallOpts)
}

// GetDebtPool is a free data retrieval call binding the contract method 0x03ec357f.
//
// Solidity: function getDebtPool() view returns(address)
func (_DefiFMintMinter *DefiFMintMinterCaller) GetDebtPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DefiFMintMinter.contract.Call(opts, &out, "getDebtPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDebtPool is a free data retrieval call binding the contract method 0x03ec357f.
//
// Solidity: function getDebtPool() view returns(address)
func (_DefiFMintMinter *DefiFMintMinterSession) GetDebtPool() (common.Address, error) {
	return _DefiFMintMinter.Contract.GetDebtPool(&_DefiFMintMinter.CallOpts)
}

// GetDebtPool is a free data retrieval call binding the contract method 0x03ec357f.
//
// Solidity: function getDebtPool() view returns(address)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) GetDebtPool() (common.Address, error) {
	return _DefiFMintMinter.Contract.GetDebtPool(&_DefiFMintMinter.CallOpts)
}

// GetExtendedPrice is a free data retrieval call binding the contract method 0x35870d2e.
//
// Solidity: function getExtendedPrice(address _token) view returns(uint256 _price, uint256 _digits)
func (_DefiFMintMinter *DefiFMintMinterCaller) GetExtendedPrice(opts *bind.CallOpts, _token common.Address) (struct {
	Price  *big.Int
	Digits *big.Int
}, error) {
	var out []interface{}
	err := _DefiFMintMinter.contract.Call(opts, &out, "getExtendedPrice", _token)

	outstruct := new(struct {
		Price  *big.Int
		Digits *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Digits = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetExtendedPrice is a free data retrieval call binding the contract method 0x35870d2e.
//
// Solidity: function getExtendedPrice(address _token) view returns(uint256 _price, uint256 _digits)
func (_DefiFMintMinter *DefiFMintMinterSession) GetExtendedPrice(_token common.Address) (struct {
	Price  *big.Int
	Digits *big.Int
}, error) {
	return _DefiFMintMinter.Contract.GetExtendedPrice(&_DefiFMintMinter.CallOpts, _token)
}

// GetExtendedPrice is a free data retrieval call binding the contract method 0x35870d2e.
//
// Solidity: function getExtendedPrice(address _token) view returns(uint256 _price, uint256 _digits)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) GetExtendedPrice(_token common.Address) (struct {
	Price  *big.Int
	Digits *big.Int
}, error) {
	return _DefiFMintMinter.Contract.GetExtendedPrice(&_DefiFMintMinter.CallOpts, _token)
}

// GetFMintFee4dec is a free data retrieval call binding the contract method 0x7c4b7a86.
//
// Solidity: function getFMintFee4dec() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) GetFMintFee4dec(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DefiFMintMinter.contract.Call(opts, &out, "getFMintFee4dec")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFMintFee4dec is a free data retrieval call binding the contract method 0x7c4b7a86.
//
// Solidity: function getFMintFee4dec() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) GetFMintFee4dec() (*big.Int, error) {
	return _DefiFMintMinter.Contract.GetFMintFee4dec(&_DefiFMintMinter.CallOpts)
}

// GetFMintFee4dec is a free data retrieval call binding the contract method 0x7c4b7a86.
//
// Solidity: function getFMintFee4dec() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) GetFMintFee4dec() (*big.Int, error) {
	return _DefiFMintMinter.Contract.GetFMintFee4dec(&_DefiFMintMinter.CallOpts)
}

// GetMaxToMint is a free data retrieval call binding the contract method 0x54a36bcf.
//
// Solidity: function getMaxToMint(address _account, address _token, uint256 _ratio) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) GetMaxToMint(opts *bind.CallOpts, _account common.Address, _token common.Address, _ratio *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DefiFMintMinter.contract.Call(opts, &out, "getMaxToMint", _account, _token, _ratio)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxToMint is a free data retrieval call binding the contract method 0x54a36bcf.
//
// Solidity: function getMaxToMint(address _account, address _token, uint256 _ratio) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) GetMaxToMint(_account common.Address, _token common.Address, _ratio *big.Int) (*big.Int, error) {
	return _DefiFMintMinter.Contract.GetMaxToMint(&_DefiFMintMinter.CallOpts, _account, _token, _ratio)
}

// GetMaxToMint is a free data retrieval call binding the contract method 0x54a36bcf.
//
// Solidity: function getMaxToMint(address _account, address _token, uint256 _ratio) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) GetMaxToMint(_account common.Address, _token common.Address, _ratio *big.Int) (*big.Int, error) {
	return _DefiFMintMinter.Contract.GetMaxToMint(&_DefiFMintMinter.CallOpts, _account, _token, _ratio)
}

// GetMaxToWithdraw is a free data retrieval call binding the contract method 0x850e102f.
//
// Solidity: function getMaxToWithdraw(address _account, address _token, uint256 _ratio) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) GetMaxToWithdraw(opts *bind.CallOpts, _account common.Address, _token common.Address, _ratio *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DefiFMintMinter.contract.Call(opts, &out, "getMaxToWithdraw", _account, _token, _ratio)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxToWithdraw is a free data retrieval call binding the contract method 0x850e102f.
//
// Solidity: function getMaxToWithdraw(address _account, address _token, uint256 _ratio) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) GetMaxToWithdraw(_account common.Address, _token common.Address, _ratio *big.Int) (*big.Int, error) {
	return _DefiFMintMinter.Contract.GetMaxToWithdraw(&_DefiFMintMinter.CallOpts, _account, _token, _ratio)
}

// GetMaxToWithdraw is a free data retrieval call binding the contract method 0x850e102f.
//
// Solidity: function getMaxToWithdraw(address _account, address _token, uint256 _ratio) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) GetMaxToWithdraw(_account common.Address, _token common.Address, _ratio *big.Int) (*big.Int, error) {
	return _DefiFMintMinter.Contract.GetMaxToWithdraw(&_DefiFMintMinter.CallOpts, _account, _token, _ratio)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address _token) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) GetPrice(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DefiFMintMinter.contract.Call(opts, &out, "getPrice", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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

// GetRewardEligibilityRatio4dec is a free data retrieval call binding the contract method 0x59eb3570.
//
// Solidity: function getRewardEligibilityRatio4dec() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) GetRewardEligibilityRatio4dec(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DefiFMintMinter.contract.Call(opts, &out, "getRewardEligibilityRatio4dec")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardEligibilityRatio4dec is a free data retrieval call binding the contract method 0x59eb3570.
//
// Solidity: function getRewardEligibilityRatio4dec() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) GetRewardEligibilityRatio4dec() (*big.Int, error) {
	return _DefiFMintMinter.Contract.GetRewardEligibilityRatio4dec(&_DefiFMintMinter.CallOpts)
}

// GetRewardEligibilityRatio4dec is a free data retrieval call binding the contract method 0x59eb3570.
//
// Solidity: function getRewardEligibilityRatio4dec() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) GetRewardEligibilityRatio4dec() (*big.Int, error) {
	return _DefiFMintMinter.Contract.GetRewardEligibilityRatio4dec(&_DefiFMintMinter.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DefiFMintMinter.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

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

// MaxToMint is a free data retrieval call binding the contract method 0xd4ca4d74.
//
// Solidity: function maxToMint(address _account, address _token, uint256 _ratio) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) MaxToMint(opts *bind.CallOpts, _account common.Address, _token common.Address, _ratio *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DefiFMintMinter.contract.Call(opts, &out, "maxToMint", _account, _token, _ratio)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxToMint is a free data retrieval call binding the contract method 0xd4ca4d74.
//
// Solidity: function maxToMint(address _account, address _token, uint256 _ratio) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) MaxToMint(_account common.Address, _token common.Address, _ratio *big.Int) (*big.Int, error) {
	return _DefiFMintMinter.Contract.MaxToMint(&_DefiFMintMinter.CallOpts, _account, _token, _ratio)
}

// MaxToMint is a free data retrieval call binding the contract method 0xd4ca4d74.
//
// Solidity: function maxToMint(address _account, address _token, uint256 _ratio) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) MaxToMint(_account common.Address, _token common.Address, _ratio *big.Int) (*big.Int, error) {
	return _DefiFMintMinter.Contract.MaxToMint(&_DefiFMintMinter.CallOpts, _account, _token, _ratio)
}

// MaxToWithdraw is a free data retrieval call binding the contract method 0xdc2b18cc.
//
// Solidity: function maxToWithdraw(address _account, address _token, uint256 _ratio) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) MaxToWithdraw(opts *bind.CallOpts, _account common.Address, _token common.Address, _ratio *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DefiFMintMinter.contract.Call(opts, &out, "maxToWithdraw", _account, _token, _ratio)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxToWithdraw is a free data retrieval call binding the contract method 0xdc2b18cc.
//
// Solidity: function maxToWithdraw(address _account, address _token, uint256 _ratio) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) MaxToWithdraw(_account common.Address, _token common.Address, _ratio *big.Int) (*big.Int, error) {
	return _DefiFMintMinter.Contract.MaxToWithdraw(&_DefiFMintMinter.CallOpts, _account, _token, _ratio)
}

// MaxToWithdraw is a free data retrieval call binding the contract method 0xdc2b18cc.
//
// Solidity: function maxToWithdraw(address _account, address _token, uint256 _ratio) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) MaxToWithdraw(_account common.Address, _token common.Address, _ratio *big.Int) (*big.Int, error) {
	return _DefiFMintMinter.Contract.MaxToWithdraw(&_DefiFMintMinter.CallOpts, _account, _token, _ratio)
}

// MinToDeposit is a free data retrieval call binding the contract method 0x90c36959.
//
// Solidity: function minToDeposit(address _account, address _token, uint256 _ratio) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) MinToDeposit(opts *bind.CallOpts, _account common.Address, _token common.Address, _ratio *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DefiFMintMinter.contract.Call(opts, &out, "minToDeposit", _account, _token, _ratio)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinToDeposit is a free data retrieval call binding the contract method 0x90c36959.
//
// Solidity: function minToDeposit(address _account, address _token, uint256 _ratio) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) MinToDeposit(_account common.Address, _token common.Address, _ratio *big.Int) (*big.Int, error) {
	return _DefiFMintMinter.Contract.MinToDeposit(&_DefiFMintMinter.CallOpts, _account, _token, _ratio)
}

// MinToDeposit is a free data retrieval call binding the contract method 0x90c36959.
//
// Solidity: function minToDeposit(address _account, address _token, uint256 _ratio) view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) MinToDeposit(_account common.Address, _token common.Address, _ratio *big.Int) (*big.Int, error) {
	return _DefiFMintMinter.Contract.MinToDeposit(&_DefiFMintMinter.CallOpts, _account, _token, _ratio)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DefiFMintMinter *DefiFMintMinterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DefiFMintMinter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

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

// RewardCanClaim is a free data retrieval call binding the contract method 0xda0a0432.
//
// Solidity: function rewardCanClaim(address _account) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterCaller) RewardCanClaim(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var out []interface{}
	err := _DefiFMintMinter.contract.Call(opts, &out, "rewardCanClaim", _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

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

// RewardEligibilityRatio4dec is a free data retrieval call binding the contract method 0x50fca4bd.
//
// Solidity: function rewardEligibilityRatio4dec() view returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterCaller) RewardEligibilityRatio4dec(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DefiFMintMinter.contract.Call(opts, &out, "rewardEligibilityRatio4dec")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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

// RewardIsEligible is a free data retrieval call binding the contract method 0x6aee9c13.
//
// Solidity: function rewardIsEligible(address _account) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterCaller) RewardIsEligible(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var out []interface{}
	err := _DefiFMintMinter.contract.Call(opts, &out, "rewardIsEligible", _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RewardIsEligible is a free data retrieval call binding the contract method 0x6aee9c13.
//
// Solidity: function rewardIsEligible(address _account) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterSession) RewardIsEligible(_account common.Address) (bool, error) {
	return _DefiFMintMinter.Contract.RewardIsEligible(&_DefiFMintMinter.CallOpts, _account)
}

// RewardIsEligible is a free data retrieval call binding the contract method 0x6aee9c13.
//
// Solidity: function rewardIsEligible(address _account) view returns(bool)
func (_DefiFMintMinter *DefiFMintMinterCallerSession) RewardIsEligible(_account common.Address) (bool, error) {
	return _DefiFMintMinter.Contract.RewardIsEligible(&_DefiFMintMinter.CallOpts, _account)
}

// CfgSetLowestCollateralRatio is a paid mutator transaction binding the contract method 0xf1bd51ea.
//
// Solidity: function cfgSetLowestCollateralRatio(uint256 _ratio4dec) returns()
func (_DefiFMintMinter *DefiFMintMinterTransactor) CfgSetLowestCollateralRatio(opts *bind.TransactOpts, _ratio4dec *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.contract.Transact(opts, "cfgSetLowestCollateralRatio", _ratio4dec)
}

// CfgSetLowestCollateralRatio is a paid mutator transaction binding the contract method 0xf1bd51ea.
//
// Solidity: function cfgSetLowestCollateralRatio(uint256 _ratio4dec) returns()
func (_DefiFMintMinter *DefiFMintMinterSession) CfgSetLowestCollateralRatio(_ratio4dec *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.CfgSetLowestCollateralRatio(&_DefiFMintMinter.TransactOpts, _ratio4dec)
}

// CfgSetLowestCollateralRatio is a paid mutator transaction binding the contract method 0xf1bd51ea.
//
// Solidity: function cfgSetLowestCollateralRatio(uint256 _ratio4dec) returns()
func (_DefiFMintMinter *DefiFMintMinterTransactorSession) CfgSetLowestCollateralRatio(_ratio4dec *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.CfgSetLowestCollateralRatio(&_DefiFMintMinter.TransactOpts, _ratio4dec)
}

// CfgSetMintFee is a paid mutator transaction binding the contract method 0x51ce9561.
//
// Solidity: function cfgSetMintFee(uint256 _fee4dec) returns()
func (_DefiFMintMinter *DefiFMintMinterTransactor) CfgSetMintFee(opts *bind.TransactOpts, _fee4dec *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.contract.Transact(opts, "cfgSetMintFee", _fee4dec)
}

// CfgSetMintFee is a paid mutator transaction binding the contract method 0x51ce9561.
//
// Solidity: function cfgSetMintFee(uint256 _fee4dec) returns()
func (_DefiFMintMinter *DefiFMintMinterSession) CfgSetMintFee(_fee4dec *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.CfgSetMintFee(&_DefiFMintMinter.TransactOpts, _fee4dec)
}

// CfgSetMintFee is a paid mutator transaction binding the contract method 0x51ce9561.
//
// Solidity: function cfgSetMintFee(uint256 _fee4dec) returns()
func (_DefiFMintMinter *DefiFMintMinterTransactorSession) CfgSetMintFee(_fee4dec *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.CfgSetMintFee(&_DefiFMintMinter.TransactOpts, _fee4dec)
}

// CfgSetRewardEligibilityRatio is a paid mutator transaction binding the contract method 0xd3f70634.
//
// Solidity: function cfgSetRewardEligibilityRatio(uint256 _ratio4dec) returns()
func (_DefiFMintMinter *DefiFMintMinterTransactor) CfgSetRewardEligibilityRatio(opts *bind.TransactOpts, _ratio4dec *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.contract.Transact(opts, "cfgSetRewardEligibilityRatio", _ratio4dec)
}

// CfgSetRewardEligibilityRatio is a paid mutator transaction binding the contract method 0xd3f70634.
//
// Solidity: function cfgSetRewardEligibilityRatio(uint256 _ratio4dec) returns()
func (_DefiFMintMinter *DefiFMintMinterSession) CfgSetRewardEligibilityRatio(_ratio4dec *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.CfgSetRewardEligibilityRatio(&_DefiFMintMinter.TransactOpts, _ratio4dec)
}

// CfgSetRewardEligibilityRatio is a paid mutator transaction binding the contract method 0xd3f70634.
//
// Solidity: function cfgSetRewardEligibilityRatio(uint256 _ratio4dec) returns()
func (_DefiFMintMinter *DefiFMintMinterTransactorSession) CfgSetRewardEligibilityRatio(_ratio4dec *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.CfgSetRewardEligibilityRatio(&_DefiFMintMinter.TransactOpts, _ratio4dec)
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

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address owner, address _addressProvider) returns()
func (_DefiFMintMinter *DefiFMintMinterTransactor) Initialize(opts *bind.TransactOpts, owner common.Address, _addressProvider common.Address) (*types.Transaction, error) {
	return _DefiFMintMinter.contract.Transact(opts, "initialize", owner, _addressProvider)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address owner, address _addressProvider) returns()
func (_DefiFMintMinter *DefiFMintMinterSession) Initialize(owner common.Address, _addressProvider common.Address) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.Initialize(&_DefiFMintMinter.TransactOpts, owner, _addressProvider)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address owner, address _addressProvider) returns()
func (_DefiFMintMinter *DefiFMintMinterTransactorSession) Initialize(owner common.Address, _addressProvider common.Address) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.Initialize(&_DefiFMintMinter.TransactOpts, owner, _addressProvider)
}

// Initialize0 is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_DefiFMintMinter *DefiFMintMinterTransactor) Initialize0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefiFMintMinter.contract.Transact(opts, "initialize0")
}

// Initialize0 is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_DefiFMintMinter *DefiFMintMinterSession) Initialize0() (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.Initialize0(&_DefiFMintMinter.TransactOpts)
}

// Initialize0 is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_DefiFMintMinter *DefiFMintMinterTransactorSession) Initialize0() (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.Initialize0(&_DefiFMintMinter.TransactOpts)
}

// Initialize1 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner) returns()
func (_DefiFMintMinter *DefiFMintMinterTransactor) Initialize1(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _DefiFMintMinter.contract.Transact(opts, "initialize1", owner)
}

// Initialize1 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner) returns()
func (_DefiFMintMinter *DefiFMintMinterSession) Initialize1(owner common.Address) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.Initialize1(&_DefiFMintMinter.TransactOpts, owner)
}

// Initialize1 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner) returns()
func (_DefiFMintMinter *DefiFMintMinterTransactorSession) Initialize1(owner common.Address) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.Initialize1(&_DefiFMintMinter.TransactOpts, owner)
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

// MintMax is a paid mutator transaction binding the contract method 0x50fd96a8.
//
// Solidity: function mintMax(address _token, uint256 _ratio) returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterTransactor) MintMax(opts *bind.TransactOpts, _token common.Address, _ratio *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.contract.Transact(opts, "mintMax", _token, _ratio)
}

// MintMax is a paid mutator transaction binding the contract method 0x50fd96a8.
//
// Solidity: function mintMax(address _token, uint256 _ratio) returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) MintMax(_token common.Address, _ratio *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.MintMax(&_DefiFMintMinter.TransactOpts, _token, _ratio)
}

// MintMax is a paid mutator transaction binding the contract method 0x50fd96a8.
//
// Solidity: function mintMax(address _token, uint256 _ratio) returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterTransactorSession) MintMax(_token common.Address, _ratio *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.MintMax(&_DefiFMintMinter.TransactOpts, _token, _ratio)
}

// MustDeposit is a paid mutator transaction binding the contract method 0xa02bda7a.
//
// Solidity: function mustDeposit(address _token, uint256 _amount) returns()
func (_DefiFMintMinter *DefiFMintMinterTransactor) MustDeposit(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.contract.Transact(opts, "mustDeposit", _token, _amount)
}

// MustDeposit is a paid mutator transaction binding the contract method 0xa02bda7a.
//
// Solidity: function mustDeposit(address _token, uint256 _amount) returns()
func (_DefiFMintMinter *DefiFMintMinterSession) MustDeposit(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.MustDeposit(&_DefiFMintMinter.TransactOpts, _token, _amount)
}

// MustDeposit is a paid mutator transaction binding the contract method 0xa02bda7a.
//
// Solidity: function mustDeposit(address _token, uint256 _amount) returns()
func (_DefiFMintMinter *DefiFMintMinterTransactorSession) MustDeposit(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.MustDeposit(&_DefiFMintMinter.TransactOpts, _token, _amount)
}

// MustMint is a paid mutator transaction binding the contract method 0x893ebfd5.
//
// Solidity: function mustMint(address _token, uint256 _amount) returns()
func (_DefiFMintMinter *DefiFMintMinterTransactor) MustMint(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.contract.Transact(opts, "mustMint", _token, _amount)
}

// MustMint is a paid mutator transaction binding the contract method 0x893ebfd5.
//
// Solidity: function mustMint(address _token, uint256 _amount) returns()
func (_DefiFMintMinter *DefiFMintMinterSession) MustMint(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.MustMint(&_DefiFMintMinter.TransactOpts, _token, _amount)
}

// MustMint is a paid mutator transaction binding the contract method 0x893ebfd5.
//
// Solidity: function mustMint(address _token, uint256 _amount) returns()
func (_DefiFMintMinter *DefiFMintMinterTransactorSession) MustMint(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.MustMint(&_DefiFMintMinter.TransactOpts, _token, _amount)
}

// MustMintMax is a paid mutator transaction binding the contract method 0x572f9225.
//
// Solidity: function mustMintMax(address _token, uint256 _ratio) returns()
func (_DefiFMintMinter *DefiFMintMinterTransactor) MustMintMax(opts *bind.TransactOpts, _token common.Address, _ratio *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.contract.Transact(opts, "mustMintMax", _token, _ratio)
}

// MustMintMax is a paid mutator transaction binding the contract method 0x572f9225.
//
// Solidity: function mustMintMax(address _token, uint256 _ratio) returns()
func (_DefiFMintMinter *DefiFMintMinterSession) MustMintMax(_token common.Address, _ratio *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.MustMintMax(&_DefiFMintMinter.TransactOpts, _token, _ratio)
}

// MustMintMax is a paid mutator transaction binding the contract method 0x572f9225.
//
// Solidity: function mustMintMax(address _token, uint256 _ratio) returns()
func (_DefiFMintMinter *DefiFMintMinterTransactorSession) MustMintMax(_token common.Address, _ratio *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.MustMintMax(&_DefiFMintMinter.TransactOpts, _token, _ratio)
}

// MustRepay is a paid mutator transaction binding the contract method 0x557c138b.
//
// Solidity: function mustRepay(address _token, uint256 _amount) returns()
func (_DefiFMintMinter *DefiFMintMinterTransactor) MustRepay(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.contract.Transact(opts, "mustRepay", _token, _amount)
}

// MustRepay is a paid mutator transaction binding the contract method 0x557c138b.
//
// Solidity: function mustRepay(address _token, uint256 _amount) returns()
func (_DefiFMintMinter *DefiFMintMinterSession) MustRepay(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.MustRepay(&_DefiFMintMinter.TransactOpts, _token, _amount)
}

// MustRepay is a paid mutator transaction binding the contract method 0x557c138b.
//
// Solidity: function mustRepay(address _token, uint256 _amount) returns()
func (_DefiFMintMinter *DefiFMintMinterTransactorSession) MustRepay(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.MustRepay(&_DefiFMintMinter.TransactOpts, _token, _amount)
}

// MustRepayMax is a paid mutator transaction binding the contract method 0xcc7b9327.
//
// Solidity: function mustRepayMax(address _token) returns()
func (_DefiFMintMinter *DefiFMintMinterTransactor) MustRepayMax(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _DefiFMintMinter.contract.Transact(opts, "mustRepayMax", _token)
}

// MustRepayMax is a paid mutator transaction binding the contract method 0xcc7b9327.
//
// Solidity: function mustRepayMax(address _token) returns()
func (_DefiFMintMinter *DefiFMintMinterSession) MustRepayMax(_token common.Address) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.MustRepayMax(&_DefiFMintMinter.TransactOpts, _token)
}

// MustRepayMax is a paid mutator transaction binding the contract method 0xcc7b9327.
//
// Solidity: function mustRepayMax(address _token) returns()
func (_DefiFMintMinter *DefiFMintMinterTransactorSession) MustRepayMax(_token common.Address) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.MustRepayMax(&_DefiFMintMinter.TransactOpts, _token)
}

// MustWithdraw is a paid mutator transaction binding the contract method 0x0feea739.
//
// Solidity: function mustWithdraw(address _token, uint256 _amount) returns()
func (_DefiFMintMinter *DefiFMintMinterTransactor) MustWithdraw(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.contract.Transact(opts, "mustWithdraw", _token, _amount)
}

// MustWithdraw is a paid mutator transaction binding the contract method 0x0feea739.
//
// Solidity: function mustWithdraw(address _token, uint256 _amount) returns()
func (_DefiFMintMinter *DefiFMintMinterSession) MustWithdraw(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.MustWithdraw(&_DefiFMintMinter.TransactOpts, _token, _amount)
}

// MustWithdraw is a paid mutator transaction binding the contract method 0x0feea739.
//
// Solidity: function mustWithdraw(address _token, uint256 _amount) returns()
func (_DefiFMintMinter *DefiFMintMinterTransactorSession) MustWithdraw(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.MustWithdraw(&_DefiFMintMinter.TransactOpts, _token, _amount)
}

// MustWithdrawMax is a paid mutator transaction binding the contract method 0x1e3eb86a.
//
// Solidity: function mustWithdrawMax(address _token, uint256 _ratio) returns()
func (_DefiFMintMinter *DefiFMintMinterTransactor) MustWithdrawMax(opts *bind.TransactOpts, _token common.Address, _ratio *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.contract.Transact(opts, "mustWithdrawMax", _token, _ratio)
}

// MustWithdrawMax is a paid mutator transaction binding the contract method 0x1e3eb86a.
//
// Solidity: function mustWithdrawMax(address _token, uint256 _ratio) returns()
func (_DefiFMintMinter *DefiFMintMinterSession) MustWithdrawMax(_token common.Address, _ratio *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.MustWithdrawMax(&_DefiFMintMinter.TransactOpts, _token, _ratio)
}

// MustWithdrawMax is a paid mutator transaction binding the contract method 0x1e3eb86a.
//
// Solidity: function mustWithdrawMax(address _token, uint256 _ratio) returns()
func (_DefiFMintMinter *DefiFMintMinterTransactorSession) MustWithdrawMax(_token common.Address, _ratio *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.MustWithdrawMax(&_DefiFMintMinter.TransactOpts, _token, _ratio)
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

// RepayMax is a paid mutator transaction binding the contract method 0xbcc2ff51.
//
// Solidity: function repayMax(address _token) returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterTransactor) RepayMax(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _DefiFMintMinter.contract.Transact(opts, "repayMax", _token)
}

// RepayMax is a paid mutator transaction binding the contract method 0xbcc2ff51.
//
// Solidity: function repayMax(address _token) returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) RepayMax(_token common.Address) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.RepayMax(&_DefiFMintMinter.TransactOpts, _token)
}

// RepayMax is a paid mutator transaction binding the contract method 0xbcc2ff51.
//
// Solidity: function repayMax(address _token) returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterTransactorSession) RepayMax(_token common.Address) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.RepayMax(&_DefiFMintMinter.TransactOpts, _token)
}

// RewardUpdate is a paid mutator transaction binding the contract method 0x48ebb08d.
//
// Solidity: function rewardUpdate(address _account) returns()
func (_DefiFMintMinter *DefiFMintMinterTransactor) RewardUpdate(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _DefiFMintMinter.contract.Transact(opts, "rewardUpdate", _account)
}

// RewardUpdate is a paid mutator transaction binding the contract method 0x48ebb08d.
//
// Solidity: function rewardUpdate(address _account) returns()
func (_DefiFMintMinter *DefiFMintMinterSession) RewardUpdate(_account common.Address) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.RewardUpdate(&_DefiFMintMinter.TransactOpts, _account)
}

// RewardUpdate is a paid mutator transaction binding the contract method 0x48ebb08d.
//
// Solidity: function rewardUpdate(address _account) returns()
func (_DefiFMintMinter *DefiFMintMinterTransactorSession) RewardUpdate(_account common.Address) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.RewardUpdate(&_DefiFMintMinter.TransactOpts, _account)
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

// WithdrawMax is a paid mutator transaction binding the contract method 0xdb3243f2.
//
// Solidity: function withdrawMax(address _token, uint256 _ratio) returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterTransactor) WithdrawMax(opts *bind.TransactOpts, _token common.Address, _ratio *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.contract.Transact(opts, "withdrawMax", _token, _ratio)
}

// WithdrawMax is a paid mutator transaction binding the contract method 0xdb3243f2.
//
// Solidity: function withdrawMax(address _token, uint256 _ratio) returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterSession) WithdrawMax(_token common.Address, _ratio *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.WithdrawMax(&_DefiFMintMinter.TransactOpts, _token, _ratio)
}

// WithdrawMax is a paid mutator transaction binding the contract method 0xdb3243f2.
//
// Solidity: function withdrawMax(address _token, uint256 _ratio) returns(uint256)
func (_DefiFMintMinter *DefiFMintMinterTransactorSession) WithdrawMax(_token common.Address, _ratio *big.Int) (*types.Transaction, error) {
	return _DefiFMintMinter.Contract.WithdrawMax(&_DefiFMintMinter.TransactOpts, _token, _ratio)
}

// DefiFMintMinterCollateralLowestDebtRatioChangedIterator is returned from FilterCollateralLowestDebtRatioChanged and is used to iterate over the raw logs and unpacked data for CollateralLowestDebtRatioChanged events raised by the DefiFMintMinter contract.
type DefiFMintMinterCollateralLowestDebtRatioChangedIterator struct {
	Event *DefiFMintMinterCollateralLowestDebtRatioChanged // Event containing the contract specifics and raw log

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
func (it *DefiFMintMinterCollateralLowestDebtRatioChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefiFMintMinterCollateralLowestDebtRatioChanged)
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
		it.Event = new(DefiFMintMinterCollateralLowestDebtRatioChanged)
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
func (it *DefiFMintMinterCollateralLowestDebtRatioChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefiFMintMinterCollateralLowestDebtRatioChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefiFMintMinterCollateralLowestDebtRatioChanged represents a CollateralLowestDebtRatioChanged event raised by the DefiFMintMinter contract.
type DefiFMintMinterCollateralLowestDebtRatioChanged struct {
	Ratio4dec *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCollateralLowestDebtRatioChanged is a free log retrieval operation binding the contract event 0x03b166133cc99dd16eff1cc93a1a34996f2710564ca9563fe1ddd539293f3e68.
//
// Solidity: event CollateralLowestDebtRatioChanged(uint256 ratio4dec)
func (_DefiFMintMinter *DefiFMintMinterFilterer) FilterCollateralLowestDebtRatioChanged(opts *bind.FilterOpts) (*DefiFMintMinterCollateralLowestDebtRatioChangedIterator, error) {

	logs, sub, err := _DefiFMintMinter.contract.FilterLogs(opts, "CollateralLowestDebtRatioChanged")
	if err != nil {
		return nil, err
	}
	return &DefiFMintMinterCollateralLowestDebtRatioChangedIterator{contract: _DefiFMintMinter.contract, event: "CollateralLowestDebtRatioChanged", logs: logs, sub: sub}, nil
}

// WatchCollateralLowestDebtRatioChanged is a free log subscription operation binding the contract event 0x03b166133cc99dd16eff1cc93a1a34996f2710564ca9563fe1ddd539293f3e68.
//
// Solidity: event CollateralLowestDebtRatioChanged(uint256 ratio4dec)
func (_DefiFMintMinter *DefiFMintMinterFilterer) WatchCollateralLowestDebtRatioChanged(opts *bind.WatchOpts, sink chan<- *DefiFMintMinterCollateralLowestDebtRatioChanged) (event.Subscription, error) {

	logs, sub, err := _DefiFMintMinter.contract.WatchLogs(opts, "CollateralLowestDebtRatioChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefiFMintMinterCollateralLowestDebtRatioChanged)
				if err := _DefiFMintMinter.contract.UnpackLog(event, "CollateralLowestDebtRatioChanged", log); err != nil {
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

// ParseCollateralLowestDebtRatioChanged is a log parse operation binding the contract event 0x03b166133cc99dd16eff1cc93a1a34996f2710564ca9563fe1ddd539293f3e68.
//
// Solidity: event CollateralLowestDebtRatioChanged(uint256 ratio4dec)
func (_DefiFMintMinter *DefiFMintMinterFilterer) ParseCollateralLowestDebtRatioChanged(log types.Log) (*DefiFMintMinterCollateralLowestDebtRatioChanged, error) {
	event := new(DefiFMintMinterCollateralLowestDebtRatioChanged)
	if err := _DefiFMintMinter.contract.UnpackLog(event, "CollateralLowestDebtRatioChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	event.Raw = log
	return event, nil
}

// DefiFMintMinterMintFeeChangedIterator is returned from FilterMintFeeChanged and is used to iterate over the raw logs and unpacked data for MintFeeChanged events raised by the DefiFMintMinter contract.
type DefiFMintMinterMintFeeChangedIterator struct {
	Event *DefiFMintMinterMintFeeChanged // Event containing the contract specifics and raw log

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
func (it *DefiFMintMinterMintFeeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefiFMintMinterMintFeeChanged)
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
		it.Event = new(DefiFMintMinterMintFeeChanged)
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
func (it *DefiFMintMinterMintFeeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefiFMintMinterMintFeeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefiFMintMinterMintFeeChanged represents a MintFeeChanged event raised by the DefiFMintMinter contract.
type DefiFMintMinterMintFeeChanged struct {
	Fee4dec *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMintFeeChanged is a free log retrieval operation binding the contract event 0xe427e272b122e738fd867ac5defcedb2bc9362341166a49d793d8b230f75670c.
//
// Solidity: event MintFeeChanged(uint256 fee4dec)
func (_DefiFMintMinter *DefiFMintMinterFilterer) FilterMintFeeChanged(opts *bind.FilterOpts) (*DefiFMintMinterMintFeeChangedIterator, error) {

	logs, sub, err := _DefiFMintMinter.contract.FilterLogs(opts, "MintFeeChanged")
	if err != nil {
		return nil, err
	}
	return &DefiFMintMinterMintFeeChangedIterator{contract: _DefiFMintMinter.contract, event: "MintFeeChanged", logs: logs, sub: sub}, nil
}

// WatchMintFeeChanged is a free log subscription operation binding the contract event 0xe427e272b122e738fd867ac5defcedb2bc9362341166a49d793d8b230f75670c.
//
// Solidity: event MintFeeChanged(uint256 fee4dec)
func (_DefiFMintMinter *DefiFMintMinterFilterer) WatchMintFeeChanged(opts *bind.WatchOpts, sink chan<- *DefiFMintMinterMintFeeChanged) (event.Subscription, error) {

	logs, sub, err := _DefiFMintMinter.contract.WatchLogs(opts, "MintFeeChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefiFMintMinterMintFeeChanged)
				if err := _DefiFMintMinter.contract.UnpackLog(event, "MintFeeChanged", log); err != nil {
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

// ParseMintFeeChanged is a log parse operation binding the contract event 0xe427e272b122e738fd867ac5defcedb2bc9362341166a49d793d8b230f75670c.
//
// Solidity: event MintFeeChanged(uint256 fee4dec)
func (_DefiFMintMinter *DefiFMintMinterFilterer) ParseMintFeeChanged(log types.Log) (*DefiFMintMinterMintFeeChanged, error) {
	event := new(DefiFMintMinterMintFeeChanged)
	if err := _DefiFMintMinter.contract.UnpackLog(event, "MintFeeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
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
	Fee    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0x03f17d66ad3bf18e9412eb06582908831508cdb9b8da9cddb1431f645a5b8632.
//
// Solidity: event Minted(address indexed token, address indexed user, uint256 amount, uint256 fee)
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

// WatchMinted is a free log subscription operation binding the contract event 0x03f17d66ad3bf18e9412eb06582908831508cdb9b8da9cddb1431f645a5b8632.
//
// Solidity: event Minted(address indexed token, address indexed user, uint256 amount, uint256 fee)
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

// ParseMinted is a log parse operation binding the contract event 0x03f17d66ad3bf18e9412eb06582908831508cdb9b8da9cddb1431f645a5b8632.
//
// Solidity: event Minted(address indexed token, address indexed user, uint256 amount, uint256 fee)
func (_DefiFMintMinter *DefiFMintMinterFilterer) ParseMinted(log types.Log) (*DefiFMintMinterMinted, error) {
	event := new(DefiFMintMinterMinted)
	if err := _DefiFMintMinter.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
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
	event.Raw = log
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
	event.Raw = log
	return event, nil
}

// DefiFMintMinterRewardEligibilityRatioChangedIterator is returned from FilterRewardEligibilityRatioChanged and is used to iterate over the raw logs and unpacked data for RewardEligibilityRatioChanged events raised by the DefiFMintMinter contract.
type DefiFMintMinterRewardEligibilityRatioChangedIterator struct {
	Event *DefiFMintMinterRewardEligibilityRatioChanged // Event containing the contract specifics and raw log

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
func (it *DefiFMintMinterRewardEligibilityRatioChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefiFMintMinterRewardEligibilityRatioChanged)
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
		it.Event = new(DefiFMintMinterRewardEligibilityRatioChanged)
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
func (it *DefiFMintMinterRewardEligibilityRatioChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefiFMintMinterRewardEligibilityRatioChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefiFMintMinterRewardEligibilityRatioChanged represents a RewardEligibilityRatioChanged event raised by the DefiFMintMinter contract.
type DefiFMintMinterRewardEligibilityRatioChanged struct {
	Ratio4dec *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRewardEligibilityRatioChanged is a free log retrieval operation binding the contract event 0x3ec85924f12f4be0739c4a0a45218af5f95364180a8e6650aedaad068db44b79.
//
// Solidity: event RewardEligibilityRatioChanged(uint256 ratio4dec)
func (_DefiFMintMinter *DefiFMintMinterFilterer) FilterRewardEligibilityRatioChanged(opts *bind.FilterOpts) (*DefiFMintMinterRewardEligibilityRatioChangedIterator, error) {

	logs, sub, err := _DefiFMintMinter.contract.FilterLogs(opts, "RewardEligibilityRatioChanged")
	if err != nil {
		return nil, err
	}
	return &DefiFMintMinterRewardEligibilityRatioChangedIterator{contract: _DefiFMintMinter.contract, event: "RewardEligibilityRatioChanged", logs: logs, sub: sub}, nil
}

// WatchRewardEligibilityRatioChanged is a free log subscription operation binding the contract event 0x3ec85924f12f4be0739c4a0a45218af5f95364180a8e6650aedaad068db44b79.
//
// Solidity: event RewardEligibilityRatioChanged(uint256 ratio4dec)
func (_DefiFMintMinter *DefiFMintMinterFilterer) WatchRewardEligibilityRatioChanged(opts *bind.WatchOpts, sink chan<- *DefiFMintMinterRewardEligibilityRatioChanged) (event.Subscription, error) {

	logs, sub, err := _DefiFMintMinter.contract.WatchLogs(opts, "RewardEligibilityRatioChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefiFMintMinterRewardEligibilityRatioChanged)
				if err := _DefiFMintMinter.contract.UnpackLog(event, "RewardEligibilityRatioChanged", log); err != nil {
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

// ParseRewardEligibilityRatioChanged is a log parse operation binding the contract event 0x3ec85924f12f4be0739c4a0a45218af5f95364180a8e6650aedaad068db44b79.
//
// Solidity: event RewardEligibilityRatioChanged(uint256 ratio4dec)
func (_DefiFMintMinter *DefiFMintMinterFilterer) ParseRewardEligibilityRatioChanged(log types.Log) (*DefiFMintMinterRewardEligibilityRatioChanged, error) {
	event := new(DefiFMintMinterRewardEligibilityRatioChanged)
	if err := _DefiFMintMinter.contract.UnpackLog(event, "RewardEligibilityRatioChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
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
	event.Raw = log
	return event, nil
}
