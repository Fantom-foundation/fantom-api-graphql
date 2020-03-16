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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SfcContractABI is the input ABI used to generate the binding from.
const SfcContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"minDelegation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bondedRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakersNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"slashedStakeTotalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"short\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"long\",\"type\":\"uint256\"}],\"name\":\"updateGasPowerAllocationRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawDelegation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"updateBaseRewardPerSec\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"prepareToWithdrawDelegation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"delegationLockPeriodEpochs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"epochSnapshots\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBaseRewardWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalTxRewardWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseRewardPerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeTotalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegationsTotalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxDelegatedRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractCommission\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"delegationsTotalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"updateStakerMetadata\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"calcTotalReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updateCapReachedDate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"unbondingUnlockPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakeTotalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"capReachedDate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakeLockPeriodTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"delegationsNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"unbondingStartDate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakeLockPeriodEpochs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getStakerID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bondedTargetRewardUnlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"unbondingPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentSealedEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakersLastID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardsAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fromEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxEpochs\",\"type\":\"uint256\"}],\"name\":\"calcValidatorRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakerMetadata\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"slashedDelegationsTotalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"validatorCommission\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxStakerMetadataSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_fromEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxEpochs\",\"type\":\"uint256\"}],\"name\":\"calcDelegationRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"e\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"}],\"name\":\"epochValidator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatedMe\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseRewardWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"txRewardWeight\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"delegations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"createdEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paidUntilEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toStakerID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"}],\"name\":\"createDelegation\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"prepareToWithdrawStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minStakeIncrease\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"createStake\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"increaseStake\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatedAmount\",\"type\":\"uint256\"}],\"name\":\"calcDelegationReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"delegationLockPeriodTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fromEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxEpochs\",\"type\":\"uint256\"}],\"name\":\"claimValidatorRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"calcValidatorReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fromEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxEpochs\",\"type\":\"uint256\"}],\"name\":\"claimDelegationRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paidUntilEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatedMe\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CreatedStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"}],\"name\":\"UpdatedStakerMetadata\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"diff\",\"type\":\"uint256\"}],\"name\":\"IncreasedStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toStakerID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CreatedDelegation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromEpoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"untilEpoch\",\"type\":\"uint256\"}],\"name\":\"ClaimedDelegationReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromEpoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"untilEpoch\",\"type\":\"uint256\"}],\"name\":\"ClaimedValidatorReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"}],\"name\":\"PreparedToWithdrawStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"penalty\",\"type\":\"uint256\"}],\"name\":\"WithdrawnStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"}],\"name\":\"PreparedToWithdrawDelegation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"penalty\",\"type\":\"uint256\"}],\"name\":\"WithdrawnDelegation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_capReachedDate\",\"type\":\"uint256\"}],\"name\":\"ChangedCapReachedDate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"UpdatedBaseRewardPerSec\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"short\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"long\",\"type\":\"uint256\"}],\"name\":\"UpdatedGasPowerAllocationRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// SfcContract is an auto generated Go binding around an Ethereum contract.
type SfcContract struct {
	SfcContractCaller     // Read-only binding to the contract
	SfcContractTransactor // Write-only binding to the contract
	SfcContractFilterer   // Log filterer for contract events
}

// SfcContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type SfcContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SfcContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SfcContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SfcContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SfcContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SfcContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SfcContractSession struct {
	Contract     *SfcContract      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SfcContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SfcContractCallerSession struct {
	Contract *SfcContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SfcContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SfcContractTransactorSession struct {
	Contract     *SfcContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SfcContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type SfcContractRaw struct {
	Contract *SfcContract // Generic contract binding to access the raw methods on
}

// SfcContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SfcContractCallerRaw struct {
	Contract *SfcContractCaller // Generic read-only contract binding to access the raw methods on
}

// SfcContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SfcContractTransactorRaw struct {
	Contract *SfcContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSfcContract creates a new instance of SfcContract, bound to a specific deployed contract.
func NewSfcContract(address common.Address, backend bind.ContractBackend) (*SfcContract, error) {
	contract, err := bindSfcContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SfcContract{SfcContractCaller: SfcContractCaller{contract: contract}, SfcContractTransactor: SfcContractTransactor{contract: contract}, SfcContractFilterer: SfcContractFilterer{contract: contract}}, nil
}

// NewSfcContractCaller creates a new read-only instance of SfcContract, bound to a specific deployed contract.
func NewSfcContractCaller(address common.Address, caller bind.ContractCaller) (*SfcContractCaller, error) {
	contract, err := bindSfcContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SfcContractCaller{contract: contract}, nil
}

// NewSfcContractTransactor creates a new write-only instance of SfcContract, bound to a specific deployed contract.
func NewSfcContractTransactor(address common.Address, transactor bind.ContractTransactor) (*SfcContractTransactor, error) {
	contract, err := bindSfcContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SfcContractTransactor{contract: contract}, nil
}

// NewSfcContractFilterer creates a new log filterer instance of SfcContract, bound to a specific deployed contract.
func NewSfcContractFilterer(address common.Address, filterer bind.ContractFilterer) (*SfcContractFilterer, error) {
	contract, err := bindSfcContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SfcContractFilterer{contract: contract}, nil
}

// bindSfcContract binds a generic wrapper to an already deployed contract.
func bindSfcContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SfcContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SfcContract *SfcContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SfcContract.Contract.SfcContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SfcContract *SfcContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SfcContract.Contract.SfcContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SfcContract *SfcContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SfcContract.Contract.SfcContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SfcContract *SfcContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SfcContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SfcContract *SfcContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SfcContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SfcContract *SfcContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SfcContract.Contract.contract.Transact(opts, method, params...)
}

// BondedRatio is a free data retrieval call binding the contract method 0x041d97a3.
//
// Solidity: function bondedRatio() constant returns(uint256)
func (_SfcContract *SfcContractCaller) BondedRatio(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SfcContract.contract.Call(opts, out, "bondedRatio")
	return *ret0, err
}

// BondedRatio is a free data retrieval call binding the contract method 0x041d97a3.
//
// Solidity: function bondedRatio() constant returns(uint256)
func (_SfcContract *SfcContractSession) BondedRatio() (*big.Int, error) {
	return _SfcContract.Contract.BondedRatio(&_SfcContract.CallOpts)
}

// BondedRatio is a free data retrieval call binding the contract method 0x041d97a3.
//
// Solidity: function bondedRatio() constant returns(uint256)
func (_SfcContract *SfcContractCallerSession) BondedRatio() (*big.Int, error) {
	return _SfcContract.Contract.BondedRatio(&_SfcContract.CallOpts)
}

// BondedTargetRewardUnlock is a free data retrieval call binding the contract method 0x6a1cf400.
//
// Solidity: function bondedTargetRewardUnlock() constant returns(uint256)
func (_SfcContract *SfcContractCaller) BondedTargetRewardUnlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SfcContract.contract.Call(opts, out, "bondedTargetRewardUnlock")
	return *ret0, err
}

// BondedTargetRewardUnlock is a free data retrieval call binding the contract method 0x6a1cf400.
//
// Solidity: function bondedTargetRewardUnlock() constant returns(uint256)
func (_SfcContract *SfcContractSession) BondedTargetRewardUnlock() (*big.Int, error) {
	return _SfcContract.Contract.BondedTargetRewardUnlock(&_SfcContract.CallOpts)
}

// BondedTargetRewardUnlock is a free data retrieval call binding the contract method 0x6a1cf400.
//
// Solidity: function bondedTargetRewardUnlock() constant returns(uint256)
func (_SfcContract *SfcContractCallerSession) BondedTargetRewardUnlock() (*big.Int, error) {
	return _SfcContract.Contract.BondedTargetRewardUnlock(&_SfcContract.CallOpts)
}

// CalcDelegationReward is a free data retrieval call binding the contract method 0xe8957988.
//
// Solidity: function calcDelegationReward(uint256 stakerID, uint256 epoch, uint256 delegatedAmount) constant returns(uint256)
func (_SfcContract *SfcContractCaller) CalcDelegationReward(opts *bind.CallOpts, stakerID *big.Int, epoch *big.Int, delegatedAmount *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SfcContract.contract.Call(opts, out, "calcDelegationReward", stakerID, epoch, delegatedAmount)
	return *ret0, err
}

// CalcDelegationReward is a free data retrieval call binding the contract method 0xe8957988.
//
// Solidity: function calcDelegationReward(uint256 stakerID, uint256 epoch, uint256 delegatedAmount) constant returns(uint256)
func (_SfcContract *SfcContractSession) CalcDelegationReward(stakerID *big.Int, epoch *big.Int, delegatedAmount *big.Int) (*big.Int, error) {
	return _SfcContract.Contract.CalcDelegationReward(&_SfcContract.CallOpts, stakerID, epoch, delegatedAmount)
}

// CalcDelegationReward is a free data retrieval call binding the contract method 0xe8957988.
//
// Solidity: function calcDelegationReward(uint256 stakerID, uint256 epoch, uint256 delegatedAmount) constant returns(uint256)
func (_SfcContract *SfcContractCallerSession) CalcDelegationReward(stakerID *big.Int, epoch *big.Int, delegatedAmount *big.Int) (*big.Int, error) {
	return _SfcContract.Contract.CalcDelegationReward(&_SfcContract.CallOpts, stakerID, epoch, delegatedAmount)
}

// CalcDelegationRewards is a free data retrieval call binding the contract method 0xb6732f3d.
//
// Solidity: function calcDelegationRewards(address delegator, uint256 _fromEpoch, uint256 maxEpochs) constant returns(uint256, uint256, uint256)
func (_SfcContract *SfcContractCaller) CalcDelegationRewards(opts *bind.CallOpts, delegator common.Address, _fromEpoch *big.Int, maxEpochs *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _SfcContract.contract.Call(opts, out, "calcDelegationRewards", delegator, _fromEpoch, maxEpochs)
	return *ret0, *ret1, *ret2, err
}

// CalcDelegationRewards is a free data retrieval call binding the contract method 0xb6732f3d.
//
// Solidity: function calcDelegationRewards(address delegator, uint256 _fromEpoch, uint256 maxEpochs) constant returns(uint256, uint256, uint256)
func (_SfcContract *SfcContractSession) CalcDelegationRewards(delegator common.Address, _fromEpoch *big.Int, maxEpochs *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _SfcContract.Contract.CalcDelegationRewards(&_SfcContract.CallOpts, delegator, _fromEpoch, maxEpochs)
}

// CalcDelegationRewards is a free data retrieval call binding the contract method 0xb6732f3d.
//
// Solidity: function calcDelegationRewards(address delegator, uint256 _fromEpoch, uint256 maxEpochs) constant returns(uint256, uint256, uint256)
func (_SfcContract *SfcContractCallerSession) CalcDelegationRewards(delegator common.Address, _fromEpoch *big.Int, maxEpochs *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _SfcContract.Contract.CalcDelegationRewards(&_SfcContract.CallOpts, delegator, _fromEpoch, maxEpochs)
}

// CalcTotalReward is a free data retrieval call binding the contract method 0x3594a5d2.
//
// Solidity: function calcTotalReward(uint256 stakerID, uint256 epoch) constant returns(uint256)
func (_SfcContract *SfcContractCaller) CalcTotalReward(opts *bind.CallOpts, stakerID *big.Int, epoch *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SfcContract.contract.Call(opts, out, "calcTotalReward", stakerID, epoch)
	return *ret0, err
}

// CalcTotalReward is a free data retrieval call binding the contract method 0x3594a5d2.
//
// Solidity: function calcTotalReward(uint256 stakerID, uint256 epoch) constant returns(uint256)
func (_SfcContract *SfcContractSession) CalcTotalReward(stakerID *big.Int, epoch *big.Int) (*big.Int, error) {
	return _SfcContract.Contract.CalcTotalReward(&_SfcContract.CallOpts, stakerID, epoch)
}

// CalcTotalReward is a free data retrieval call binding the contract method 0x3594a5d2.
//
// Solidity: function calcTotalReward(uint256 stakerID, uint256 epoch) constant returns(uint256)
func (_SfcContract *SfcContractCallerSession) CalcTotalReward(stakerID *big.Int, epoch *big.Int) (*big.Int, error) {
	return _SfcContract.Contract.CalcTotalReward(&_SfcContract.CallOpts, stakerID, epoch)
}

// CalcValidatorReward is a free data retrieval call binding the contract method 0xf70f53e0.
//
// Solidity: function calcValidatorReward(uint256 stakerID, uint256 epoch) constant returns(uint256)
func (_SfcContract *SfcContractCaller) CalcValidatorReward(opts *bind.CallOpts, stakerID *big.Int, epoch *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SfcContract.contract.Call(opts, out, "calcValidatorReward", stakerID, epoch)
	return *ret0, err
}

// CalcValidatorReward is a free data retrieval call binding the contract method 0xf70f53e0.
//
// Solidity: function calcValidatorReward(uint256 stakerID, uint256 epoch) constant returns(uint256)
func (_SfcContract *SfcContractSession) CalcValidatorReward(stakerID *big.Int, epoch *big.Int) (*big.Int, error) {
	return _SfcContract.Contract.CalcValidatorReward(&_SfcContract.CallOpts, stakerID, epoch)
}

// CalcValidatorReward is a free data retrieval call binding the contract method 0xf70f53e0.
//
// Solidity: function calcValidatorReward(uint256 stakerID, uint256 epoch) constant returns(uint256)
func (_SfcContract *SfcContractCallerSession) CalcValidatorReward(stakerID *big.Int, epoch *big.Int) (*big.Int, error) {
	return _SfcContract.Contract.CalcValidatorReward(&_SfcContract.CallOpts, stakerID, epoch)
}

// CalcValidatorRewards is a free data retrieval call binding the contract method 0x96060e71.
//
// Solidity: function calcValidatorRewards(uint256 stakerID, uint256 _fromEpoch, uint256 maxEpochs) constant returns(uint256, uint256, uint256)
func (_SfcContract *SfcContractCaller) CalcValidatorRewards(opts *bind.CallOpts, stakerID *big.Int, _fromEpoch *big.Int, maxEpochs *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _SfcContract.contract.Call(opts, out, "calcValidatorRewards", stakerID, _fromEpoch, maxEpochs)
	return *ret0, *ret1, *ret2, err
}

// CalcValidatorRewards is a free data retrieval call binding the contract method 0x96060e71.
//
// Solidity: function calcValidatorRewards(uint256 stakerID, uint256 _fromEpoch, uint256 maxEpochs) constant returns(uint256, uint256, uint256)
func (_SfcContract *SfcContractSession) CalcValidatorRewards(stakerID *big.Int, _fromEpoch *big.Int, maxEpochs *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _SfcContract.Contract.CalcValidatorRewards(&_SfcContract.CallOpts, stakerID, _fromEpoch, maxEpochs)
}

// CalcValidatorRewards is a free data retrieval call binding the contract method 0x96060e71.
//
// Solidity: function calcValidatorRewards(uint256 stakerID, uint256 _fromEpoch, uint256 maxEpochs) constant returns(uint256, uint256, uint256)
func (_SfcContract *SfcContractCallerSession) CalcValidatorRewards(stakerID *big.Int, _fromEpoch *big.Int, maxEpochs *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _SfcContract.Contract.CalcValidatorRewards(&_SfcContract.CallOpts, stakerID, _fromEpoch, maxEpochs)
}

// CapReachedDate is a free data retrieval call binding the contract method 0x3f4ef95f.
//
// Solidity: function capReachedDate() constant returns(uint256)
func (_SfcContract *SfcContractCaller) CapReachedDate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SfcContract.contract.Call(opts, out, "capReachedDate")
	return *ret0, err
}

// CapReachedDate is a free data retrieval call binding the contract method 0x3f4ef95f.
//
// Solidity: function capReachedDate() constant returns(uint256)
func (_SfcContract *SfcContractSession) CapReachedDate() (*big.Int, error) {
	return _SfcContract.Contract.CapReachedDate(&_SfcContract.CallOpts)
}

// CapReachedDate is a free data retrieval call binding the contract method 0x3f4ef95f.
//
// Solidity: function capReachedDate() constant returns(uint256)
func (_SfcContract *SfcContractCallerSession) CapReachedDate() (*big.Int, error) {
	return _SfcContract.Contract.CapReachedDate(&_SfcContract.CallOpts)
}

// ContractCommission is a free data retrieval call binding the contract method 0x2709275e.
//
// Solidity: function contractCommission() constant returns(uint256)
func (_SfcContract *SfcContractCaller) ContractCommission(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SfcContract.contract.Call(opts, out, "contractCommission")
	return *ret0, err
}

// ContractCommission is a free data retrieval call binding the contract method 0x2709275e.
//
// Solidity: function contractCommission() constant returns(uint256)
func (_SfcContract *SfcContractSession) ContractCommission() (*big.Int, error) {
	return _SfcContract.Contract.ContractCommission(&_SfcContract.CallOpts)
}

// ContractCommission is a free data retrieval call binding the contract method 0x2709275e.
//
// Solidity: function contractCommission() constant returns(uint256)
func (_SfcContract *SfcContractCallerSession) ContractCommission() (*big.Int, error) {
	return _SfcContract.Contract.ContractCommission(&_SfcContract.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() constant returns(uint256)
func (_SfcContract *SfcContractCaller) CurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SfcContract.contract.Call(opts, out, "currentEpoch")
	return *ret0, err
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() constant returns(uint256)
func (_SfcContract *SfcContractSession) CurrentEpoch() (*big.Int, error) {
	return _SfcContract.Contract.CurrentEpoch(&_SfcContract.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() constant returns(uint256)
func (_SfcContract *SfcContractCallerSession) CurrentEpoch() (*big.Int, error) {
	return _SfcContract.Contract.CurrentEpoch(&_SfcContract.CallOpts)
}

// CurrentSealedEpoch is a free data retrieval call binding the contract method 0x7cacb1d6.
//
// Solidity: function currentSealedEpoch() constant returns(uint256)
func (_SfcContract *SfcContractCaller) CurrentSealedEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SfcContract.contract.Call(opts, out, "currentSealedEpoch")
	return *ret0, err
}

// CurrentSealedEpoch is a free data retrieval call binding the contract method 0x7cacb1d6.
//
// Solidity: function currentSealedEpoch() constant returns(uint256)
func (_SfcContract *SfcContractSession) CurrentSealedEpoch() (*big.Int, error) {
	return _SfcContract.Contract.CurrentSealedEpoch(&_SfcContract.CallOpts)
}

// CurrentSealedEpoch is a free data retrieval call binding the contract method 0x7cacb1d6.
//
// Solidity: function currentSealedEpoch() constant returns(uint256)
func (_SfcContract *SfcContractCallerSession) CurrentSealedEpoch() (*big.Int, error) {
	return _SfcContract.Contract.CurrentSealedEpoch(&_SfcContract.CallOpts)
}

// DelegationLockPeriodEpochs is a free data retrieval call binding the contract method 0x1d58179c.
//
// Solidity: function delegationLockPeriodEpochs() constant returns(uint256)
func (_SfcContract *SfcContractCaller) DelegationLockPeriodEpochs(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SfcContract.contract.Call(opts, out, "delegationLockPeriodEpochs")
	return *ret0, err
}

// DelegationLockPeriodEpochs is a free data retrieval call binding the contract method 0x1d58179c.
//
// Solidity: function delegationLockPeriodEpochs() constant returns(uint256)
func (_SfcContract *SfcContractSession) DelegationLockPeriodEpochs() (*big.Int, error) {
	return _SfcContract.Contract.DelegationLockPeriodEpochs(&_SfcContract.CallOpts)
}

// DelegationLockPeriodEpochs is a free data retrieval call binding the contract method 0x1d58179c.
//
// Solidity: function delegationLockPeriodEpochs() constant returns(uint256)
func (_SfcContract *SfcContractCallerSession) DelegationLockPeriodEpochs() (*big.Int, error) {
	return _SfcContract.Contract.DelegationLockPeriodEpochs(&_SfcContract.CallOpts)
}

// DelegationLockPeriodTime is a free data retrieval call binding the contract method 0xec6a7f1c.
//
// Solidity: function delegationLockPeriodTime() constant returns(uint256)
func (_SfcContract *SfcContractCaller) DelegationLockPeriodTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SfcContract.contract.Call(opts, out, "delegationLockPeriodTime")
	return *ret0, err
}

// DelegationLockPeriodTime is a free data retrieval call binding the contract method 0xec6a7f1c.
//
// Solidity: function delegationLockPeriodTime() constant returns(uint256)
func (_SfcContract *SfcContractSession) DelegationLockPeriodTime() (*big.Int, error) {
	return _SfcContract.Contract.DelegationLockPeriodTime(&_SfcContract.CallOpts)
}

// DelegationLockPeriodTime is a free data retrieval call binding the contract method 0xec6a7f1c.
//
// Solidity: function delegationLockPeriodTime() constant returns(uint256)
func (_SfcContract *SfcContractCallerSession) DelegationLockPeriodTime() (*big.Int, error) {
	return _SfcContract.Contract.DelegationLockPeriodTime(&_SfcContract.CallOpts)
}

// Delegations is a free data retrieval call binding the contract method 0xbffe3486.
//
// Solidity: function delegations(address ) constant returns(uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime, uint256 amount, uint256 paidUntilEpoch, uint256 toStakerID)
func (_SfcContract *SfcContractCaller) Delegations(opts *bind.CallOpts, arg0 common.Address) (struct {
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	DeactivatedEpoch *big.Int
	DeactivatedTime  *big.Int
	Amount           *big.Int
	PaidUntilEpoch   *big.Int
	ToStakerID       *big.Int
}, error) {
	ret := new(struct {
		CreatedEpoch     *big.Int
		CreatedTime      *big.Int
		DeactivatedEpoch *big.Int
		DeactivatedTime  *big.Int
		Amount           *big.Int
		PaidUntilEpoch   *big.Int
		ToStakerID       *big.Int
	})
	out := ret
	err := _SfcContract.contract.Call(opts, out, "delegations", arg0)
	return *ret, err
}

// Delegations is a free data retrieval call binding the contract method 0xbffe3486.
//
// Solidity: function delegations(address ) constant returns(uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime, uint256 amount, uint256 paidUntilEpoch, uint256 toStakerID)
func (_SfcContract *SfcContractSession) Delegations(arg0 common.Address) (struct {
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	DeactivatedEpoch *big.Int
	DeactivatedTime  *big.Int
	Amount           *big.Int
	PaidUntilEpoch   *big.Int
	ToStakerID       *big.Int
}, error) {
	return _SfcContract.Contract.Delegations(&_SfcContract.CallOpts, arg0)
}

// Delegations is a free data retrieval call binding the contract method 0xbffe3486.
//
// Solidity: function delegations(address ) constant returns(uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime, uint256 amount, uint256 paidUntilEpoch, uint256 toStakerID)
func (_SfcContract *SfcContractCallerSession) Delegations(arg0 common.Address) (struct {
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	DeactivatedEpoch *big.Int
	DeactivatedTime  *big.Int
	Amount           *big.Int
	PaidUntilEpoch   *big.Int
	ToStakerID       *big.Int
}, error) {
	return _SfcContract.Contract.Delegations(&_SfcContract.CallOpts, arg0)
}

// DelegationsNum is a free data retrieval call binding the contract method 0x4bd202dc.
//
// Solidity: function delegationsNum() constant returns(uint256)
func (_SfcContract *SfcContractCaller) DelegationsNum(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SfcContract.contract.Call(opts, out, "delegationsNum")
	return *ret0, err
}

// DelegationsNum is a free data retrieval call binding the contract method 0x4bd202dc.
//
// Solidity: function delegationsNum() constant returns(uint256)
func (_SfcContract *SfcContractSession) DelegationsNum() (*big.Int, error) {
	return _SfcContract.Contract.DelegationsNum(&_SfcContract.CallOpts)
}

// DelegationsNum is a free data retrieval call binding the contract method 0x4bd202dc.
//
// Solidity: function delegationsNum() constant returns(uint256)
func (_SfcContract *SfcContractCallerSession) DelegationsNum() (*big.Int, error) {
	return _SfcContract.Contract.DelegationsNum(&_SfcContract.CallOpts)
}

// DelegationsTotalAmount is a free data retrieval call binding the contract method 0x30fa9929.
//
// Solidity: function delegationsTotalAmount() constant returns(uint256)
func (_SfcContract *SfcContractCaller) DelegationsTotalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SfcContract.contract.Call(opts, out, "delegationsTotalAmount")
	return *ret0, err
}

// DelegationsTotalAmount is a free data retrieval call binding the contract method 0x30fa9929.
//
// Solidity: function delegationsTotalAmount() constant returns(uint256)
func (_SfcContract *SfcContractSession) DelegationsTotalAmount() (*big.Int, error) {
	return _SfcContract.Contract.DelegationsTotalAmount(&_SfcContract.CallOpts)
}

// DelegationsTotalAmount is a free data retrieval call binding the contract method 0x30fa9929.
//
// Solidity: function delegationsTotalAmount() constant returns(uint256)
func (_SfcContract *SfcContractCallerSession) DelegationsTotalAmount() (*big.Int, error) {
	return _SfcContract.Contract.DelegationsTotalAmount(&_SfcContract.CallOpts)
}

// EpochSnapshots is a free data retrieval call binding the contract method 0x1e8a6956.
//
// Solidity: function epochSnapshots(uint256 ) constant returns(uint256 endTime, uint256 duration, uint256 epochFee, uint256 totalBaseRewardWeight, uint256 totalTxRewardWeight, uint256 baseRewardPerSecond, uint256 stakeTotalAmount, uint256 delegationsTotalAmount, uint256 totalSupply)
func (_SfcContract *SfcContractCaller) EpochSnapshots(opts *bind.CallOpts, arg0 *big.Int) (struct {
	EndTime                *big.Int
	Duration               *big.Int
	EpochFee               *big.Int
	TotalBaseRewardWeight  *big.Int
	TotalTxRewardWeight    *big.Int
	BaseRewardPerSecond    *big.Int
	StakeTotalAmount       *big.Int
	DelegationsTotalAmount *big.Int
	TotalSupply            *big.Int
}, error) {
	ret := new(struct {
		EndTime                *big.Int
		Duration               *big.Int
		EpochFee               *big.Int
		TotalBaseRewardWeight  *big.Int
		TotalTxRewardWeight    *big.Int
		BaseRewardPerSecond    *big.Int
		StakeTotalAmount       *big.Int
		DelegationsTotalAmount *big.Int
		TotalSupply            *big.Int
	})
	out := ret
	err := _SfcContract.contract.Call(opts, out, "epochSnapshots", arg0)
	return *ret, err
}

// EpochSnapshots is a free data retrieval call binding the contract method 0x1e8a6956.
//
// Solidity: function epochSnapshots(uint256 ) constant returns(uint256 endTime, uint256 duration, uint256 epochFee, uint256 totalBaseRewardWeight, uint256 totalTxRewardWeight, uint256 baseRewardPerSecond, uint256 stakeTotalAmount, uint256 delegationsTotalAmount, uint256 totalSupply)
func (_SfcContract *SfcContractSession) EpochSnapshots(arg0 *big.Int) (struct {
	EndTime                *big.Int
	Duration               *big.Int
	EpochFee               *big.Int
	TotalBaseRewardWeight  *big.Int
	TotalTxRewardWeight    *big.Int
	BaseRewardPerSecond    *big.Int
	StakeTotalAmount       *big.Int
	DelegationsTotalAmount *big.Int
	TotalSupply            *big.Int
}, error) {
	return _SfcContract.Contract.EpochSnapshots(&_SfcContract.CallOpts, arg0)
}

// EpochSnapshots is a free data retrieval call binding the contract method 0x1e8a6956.
//
// Solidity: function epochSnapshots(uint256 ) constant returns(uint256 endTime, uint256 duration, uint256 epochFee, uint256 totalBaseRewardWeight, uint256 totalTxRewardWeight, uint256 baseRewardPerSecond, uint256 stakeTotalAmount, uint256 delegationsTotalAmount, uint256 totalSupply)
func (_SfcContract *SfcContractCallerSession) EpochSnapshots(arg0 *big.Int) (struct {
	EndTime                *big.Int
	Duration               *big.Int
	EpochFee               *big.Int
	TotalBaseRewardWeight  *big.Int
	TotalTxRewardWeight    *big.Int
	BaseRewardPerSecond    *big.Int
	StakeTotalAmount       *big.Int
	DelegationsTotalAmount *big.Int
	TotalSupply            *big.Int
}, error) {
	return _SfcContract.Contract.EpochSnapshots(&_SfcContract.CallOpts, arg0)
}

// EpochValidator is a free data retrieval call binding the contract method 0xb9029d50.
//
// Solidity: function epochValidator(uint256 e, uint256 v) constant returns(uint256 stakeAmount, uint256 delegatedMe, uint256 baseRewardWeight, uint256 txRewardWeight)
func (_SfcContract *SfcContractCaller) EpochValidator(opts *bind.CallOpts, e *big.Int, v *big.Int) (struct {
	StakeAmount      *big.Int
	DelegatedMe      *big.Int
	BaseRewardWeight *big.Int
	TxRewardWeight   *big.Int
}, error) {
	ret := new(struct {
		StakeAmount      *big.Int
		DelegatedMe      *big.Int
		BaseRewardWeight *big.Int
		TxRewardWeight   *big.Int
	})
	out := ret
	err := _SfcContract.contract.Call(opts, out, "epochValidator", e, v)
	return *ret, err
}

// EpochValidator is a free data retrieval call binding the contract method 0xb9029d50.
//
// Solidity: function epochValidator(uint256 e, uint256 v) constant returns(uint256 stakeAmount, uint256 delegatedMe, uint256 baseRewardWeight, uint256 txRewardWeight)
func (_SfcContract *SfcContractSession) EpochValidator(e *big.Int, v *big.Int) (struct {
	StakeAmount      *big.Int
	DelegatedMe      *big.Int
	BaseRewardWeight *big.Int
	TxRewardWeight   *big.Int
}, error) {
	return _SfcContract.Contract.EpochValidator(&_SfcContract.CallOpts, e, v)
}

// EpochValidator is a free data retrieval call binding the contract method 0xb9029d50.
//
// Solidity: function epochValidator(uint256 e, uint256 v) constant returns(uint256 stakeAmount, uint256 delegatedMe, uint256 baseRewardWeight, uint256 txRewardWeight)
func (_SfcContract *SfcContractCallerSession) EpochValidator(e *big.Int, v *big.Int) (struct {
	StakeAmount      *big.Int
	DelegatedMe      *big.Int
	BaseRewardWeight *big.Int
	TxRewardWeight   *big.Int
}, error) {
	return _SfcContract.Contract.EpochValidator(&_SfcContract.CallOpts, e, v)
}

// GetStakerID is a free data retrieval call binding the contract method 0x63321e27.
//
// Solidity: function getStakerID(address addr) constant returns(uint256)
func (_SfcContract *SfcContractCaller) GetStakerID(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SfcContract.contract.Call(opts, out, "getStakerID", addr)
	return *ret0, err
}

// GetStakerID is a free data retrieval call binding the contract method 0x63321e27.
//
// Solidity: function getStakerID(address addr) constant returns(uint256)
func (_SfcContract *SfcContractSession) GetStakerID(addr common.Address) (*big.Int, error) {
	return _SfcContract.Contract.GetStakerID(&_SfcContract.CallOpts, addr)
}

// GetStakerID is a free data retrieval call binding the contract method 0x63321e27.
//
// Solidity: function getStakerID(address addr) constant returns(uint256)
func (_SfcContract *SfcContractCallerSession) GetStakerID(addr common.Address) (*big.Int, error) {
	return _SfcContract.Contract.GetStakerID(&_SfcContract.CallOpts, addr)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_SfcContract *SfcContractCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SfcContract.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_SfcContract *SfcContractSession) IsOwner() (bool, error) {
	return _SfcContract.Contract.IsOwner(&_SfcContract.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_SfcContract *SfcContractCallerSession) IsOwner() (bool, error) {
	return _SfcContract.Contract.IsOwner(&_SfcContract.CallOpts)
}

// MaxDelegatedRatio is a free data retrieval call binding the contract method 0x2265f284.
//
// Solidity: function maxDelegatedRatio() constant returns(uint256)
func (_SfcContract *SfcContractCaller) MaxDelegatedRatio(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SfcContract.contract.Call(opts, out, "maxDelegatedRatio")
	return *ret0, err
}

// MaxDelegatedRatio is a free data retrieval call binding the contract method 0x2265f284.
//
// Solidity: function maxDelegatedRatio() constant returns(uint256)
func (_SfcContract *SfcContractSession) MaxDelegatedRatio() (*big.Int, error) {
	return _SfcContract.Contract.MaxDelegatedRatio(&_SfcContract.CallOpts)
}

// MaxDelegatedRatio is a free data retrieval call binding the contract method 0x2265f284.
//
// Solidity: function maxDelegatedRatio() constant returns(uint256)
func (_SfcContract *SfcContractCallerSession) MaxDelegatedRatio() (*big.Int, error) {
	return _SfcContract.Contract.MaxDelegatedRatio(&_SfcContract.CallOpts)
}

// MaxStakerMetadataSize is a free data retrieval call binding the contract method 0xab2273c0.
//
// Solidity: function maxStakerMetadataSize() constant returns(uint256)
func (_SfcContract *SfcContractCaller) MaxStakerMetadataSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SfcContract.contract.Call(opts, out, "maxStakerMetadataSize")
	return *ret0, err
}

// MaxStakerMetadataSize is a free data retrieval call binding the contract method 0xab2273c0.
//
// Solidity: function maxStakerMetadataSize() constant returns(uint256)
func (_SfcContract *SfcContractSession) MaxStakerMetadataSize() (*big.Int, error) {
	return _SfcContract.Contract.MaxStakerMetadataSize(&_SfcContract.CallOpts)
}

// MaxStakerMetadataSize is a free data retrieval call binding the contract method 0xab2273c0.
//
// Solidity: function maxStakerMetadataSize() constant returns(uint256)
func (_SfcContract *SfcContractCallerSession) MaxStakerMetadataSize() (*big.Int, error) {
	return _SfcContract.Contract.MaxStakerMetadataSize(&_SfcContract.CallOpts)
}

// MinDelegation is a free data retrieval call binding the contract method 0x02985992.
//
// Solidity: function minDelegation() constant returns(uint256)
func (_SfcContract *SfcContractCaller) MinDelegation(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SfcContract.contract.Call(opts, out, "minDelegation")
	return *ret0, err
}

// MinDelegation is a free data retrieval call binding the contract method 0x02985992.
//
// Solidity: function minDelegation() constant returns(uint256)
func (_SfcContract *SfcContractSession) MinDelegation() (*big.Int, error) {
	return _SfcContract.Contract.MinDelegation(&_SfcContract.CallOpts)
}

// MinDelegation is a free data retrieval call binding the contract method 0x02985992.
//
// Solidity: function minDelegation() constant returns(uint256)
func (_SfcContract *SfcContractCallerSession) MinDelegation() (*big.Int, error) {
	return _SfcContract.Contract.MinDelegation(&_SfcContract.CallOpts)
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() constant returns(uint256)
func (_SfcContract *SfcContractCaller) MinStake(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SfcContract.contract.Call(opts, out, "minStake")
	return *ret0, err
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() constant returns(uint256)
func (_SfcContract *SfcContractSession) MinStake() (*big.Int, error) {
	return _SfcContract.Contract.MinStake(&_SfcContract.CallOpts)
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() constant returns(uint256)
func (_SfcContract *SfcContractCallerSession) MinStake() (*big.Int, error) {
	return _SfcContract.Contract.MinStake(&_SfcContract.CallOpts)
}

// MinStakeIncrease is a free data retrieval call binding the contract method 0xc4b5dd7e.
//
// Solidity: function minStakeIncrease() constant returns(uint256)
func (_SfcContract *SfcContractCaller) MinStakeIncrease(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SfcContract.contract.Call(opts, out, "minStakeIncrease")
	return *ret0, err
}

// MinStakeIncrease is a free data retrieval call binding the contract method 0xc4b5dd7e.
//
// Solidity: function minStakeIncrease() constant returns(uint256)
func (_SfcContract *SfcContractSession) MinStakeIncrease() (*big.Int, error) {
	return _SfcContract.Contract.MinStakeIncrease(&_SfcContract.CallOpts)
}

// MinStakeIncrease is a free data retrieval call binding the contract method 0xc4b5dd7e.
//
// Solidity: function minStakeIncrease() constant returns(uint256)
func (_SfcContract *SfcContractCallerSession) MinStakeIncrease() (*big.Int, error) {
	return _SfcContract.Contract.MinStakeIncrease(&_SfcContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SfcContract *SfcContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SfcContract.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SfcContract *SfcContractSession) Owner() (common.Address, error) {
	return _SfcContract.Contract.Owner(&_SfcContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SfcContract *SfcContractCallerSession) Owner() (common.Address, error) {
	return _SfcContract.Contract.Owner(&_SfcContract.CallOpts)
}

// RewardsAllowed is a free data retrieval call binding the contract method 0x8447c4df.
//
// Solidity: function rewardsAllowed() constant returns(bool)
func (_SfcContract *SfcContractCaller) RewardsAllowed(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SfcContract.contract.Call(opts, out, "rewardsAllowed")
	return *ret0, err
}

// RewardsAllowed is a free data retrieval call binding the contract method 0x8447c4df.
//
// Solidity: function rewardsAllowed() constant returns(bool)
func (_SfcContract *SfcContractSession) RewardsAllowed() (bool, error) {
	return _SfcContract.Contract.RewardsAllowed(&_SfcContract.CallOpts)
}

// RewardsAllowed is a free data retrieval call binding the contract method 0x8447c4df.
//
// Solidity: function rewardsAllowed() constant returns(bool)
func (_SfcContract *SfcContractCallerSession) RewardsAllowed() (bool, error) {
	return _SfcContract.Contract.RewardsAllowed(&_SfcContract.CallOpts)
}

// SlashedDelegationsTotalAmount is a free data retrieval call binding the contract method 0xa70da4d2.
//
// Solidity: function slashedDelegationsTotalAmount() constant returns(uint256)
func (_SfcContract *SfcContractCaller) SlashedDelegationsTotalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SfcContract.contract.Call(opts, out, "slashedDelegationsTotalAmount")
	return *ret0, err
}

// SlashedDelegationsTotalAmount is a free data retrieval call binding the contract method 0xa70da4d2.
//
// Solidity: function slashedDelegationsTotalAmount() constant returns(uint256)
func (_SfcContract *SfcContractSession) SlashedDelegationsTotalAmount() (*big.Int, error) {
	return _SfcContract.Contract.SlashedDelegationsTotalAmount(&_SfcContract.CallOpts)
}

// SlashedDelegationsTotalAmount is a free data retrieval call binding the contract method 0xa70da4d2.
//
// Solidity: function slashedDelegationsTotalAmount() constant returns(uint256)
func (_SfcContract *SfcContractCallerSession) SlashedDelegationsTotalAmount() (*big.Int, error) {
	return _SfcContract.Contract.SlashedDelegationsTotalAmount(&_SfcContract.CallOpts)
}

// SlashedStakeTotalAmount is a free data retrieval call binding the contract method 0x0a29180c.
//
// Solidity: function slashedStakeTotalAmount() constant returns(uint256)
func (_SfcContract *SfcContractCaller) SlashedStakeTotalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SfcContract.contract.Call(opts, out, "slashedStakeTotalAmount")
	return *ret0, err
}

// SlashedStakeTotalAmount is a free data retrieval call binding the contract method 0x0a29180c.
//
// Solidity: function slashedStakeTotalAmount() constant returns(uint256)
func (_SfcContract *SfcContractSession) SlashedStakeTotalAmount() (*big.Int, error) {
	return _SfcContract.Contract.SlashedStakeTotalAmount(&_SfcContract.CallOpts)
}

// SlashedStakeTotalAmount is a free data retrieval call binding the contract method 0x0a29180c.
//
// Solidity: function slashedStakeTotalAmount() constant returns(uint256)
func (_SfcContract *SfcContractCallerSession) SlashedStakeTotalAmount() (*big.Int, error) {
	return _SfcContract.Contract.SlashedStakeTotalAmount(&_SfcContract.CallOpts)
}

// StakeLockPeriodEpochs is a free data retrieval call binding the contract method 0x54d77ed2.
//
// Solidity: function stakeLockPeriodEpochs() constant returns(uint256)
func (_SfcContract *SfcContractCaller) StakeLockPeriodEpochs(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SfcContract.contract.Call(opts, out, "stakeLockPeriodEpochs")
	return *ret0, err
}

// StakeLockPeriodEpochs is a free data retrieval call binding the contract method 0x54d77ed2.
//
// Solidity: function stakeLockPeriodEpochs() constant returns(uint256)
func (_SfcContract *SfcContractSession) StakeLockPeriodEpochs() (*big.Int, error) {
	return _SfcContract.Contract.StakeLockPeriodEpochs(&_SfcContract.CallOpts)
}

// StakeLockPeriodEpochs is a free data retrieval call binding the contract method 0x54d77ed2.
//
// Solidity: function stakeLockPeriodEpochs() constant returns(uint256)
func (_SfcContract *SfcContractCallerSession) StakeLockPeriodEpochs() (*big.Int, error) {
	return _SfcContract.Contract.StakeLockPeriodEpochs(&_SfcContract.CallOpts)
}

// StakeLockPeriodTime is a free data retrieval call binding the contract method 0x3fee10a8.
//
// Solidity: function stakeLockPeriodTime() constant returns(uint256)
func (_SfcContract *SfcContractCaller) StakeLockPeriodTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SfcContract.contract.Call(opts, out, "stakeLockPeriodTime")
	return *ret0, err
}

// StakeLockPeriodTime is a free data retrieval call binding the contract method 0x3fee10a8.
//
// Solidity: function stakeLockPeriodTime() constant returns(uint256)
func (_SfcContract *SfcContractSession) StakeLockPeriodTime() (*big.Int, error) {
	return _SfcContract.Contract.StakeLockPeriodTime(&_SfcContract.CallOpts)
}

// StakeLockPeriodTime is a free data retrieval call binding the contract method 0x3fee10a8.
//
// Solidity: function stakeLockPeriodTime() constant returns(uint256)
func (_SfcContract *SfcContractCallerSession) StakeLockPeriodTime() (*big.Int, error) {
	return _SfcContract.Contract.StakeLockPeriodTime(&_SfcContract.CallOpts)
}

// StakeTotalAmount is a free data retrieval call binding the contract method 0x3d0317fe.
//
// Solidity: function stakeTotalAmount() constant returns(uint256)
func (_SfcContract *SfcContractCaller) StakeTotalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SfcContract.contract.Call(opts, out, "stakeTotalAmount")
	return *ret0, err
}

// StakeTotalAmount is a free data retrieval call binding the contract method 0x3d0317fe.
//
// Solidity: function stakeTotalAmount() constant returns(uint256)
func (_SfcContract *SfcContractSession) StakeTotalAmount() (*big.Int, error) {
	return _SfcContract.Contract.StakeTotalAmount(&_SfcContract.CallOpts)
}

// StakeTotalAmount is a free data retrieval call binding the contract method 0x3d0317fe.
//
// Solidity: function stakeTotalAmount() constant returns(uint256)
func (_SfcContract *SfcContractCallerSession) StakeTotalAmount() (*big.Int, error) {
	return _SfcContract.Contract.StakeTotalAmount(&_SfcContract.CallOpts)
}

// StakerMetadata is a free data retrieval call binding the contract method 0x98ec2de5.
//
// Solidity: function stakerMetadata(uint256 ) constant returns(bytes)
func (_SfcContract *SfcContractCaller) StakerMetadata(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _SfcContract.contract.Call(opts, out, "stakerMetadata", arg0)
	return *ret0, err
}

// StakerMetadata is a free data retrieval call binding the contract method 0x98ec2de5.
//
// Solidity: function stakerMetadata(uint256 ) constant returns(bytes)
func (_SfcContract *SfcContractSession) StakerMetadata(arg0 *big.Int) ([]byte, error) {
	return _SfcContract.Contract.StakerMetadata(&_SfcContract.CallOpts, arg0)
}

// StakerMetadata is a free data retrieval call binding the contract method 0x98ec2de5.
//
// Solidity: function stakerMetadata(uint256 ) constant returns(bytes)
func (_SfcContract *SfcContractCallerSession) StakerMetadata(arg0 *big.Int) ([]byte, error) {
	return _SfcContract.Contract.StakerMetadata(&_SfcContract.CallOpts, arg0)
}

// Stakers is a free data retrieval call binding the contract method 0xfd5e6dd1.
//
// Solidity: function stakers(uint256 ) constant returns(uint256 status, uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime, uint256 stakeAmount, uint256 paidUntilEpoch, uint256 delegatedMe, address stakerAddress)
func (_SfcContract *SfcContractCaller) Stakers(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Status           *big.Int
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	DeactivatedEpoch *big.Int
	DeactivatedTime  *big.Int
	StakeAmount      *big.Int
	PaidUntilEpoch   *big.Int
	DelegatedMe      *big.Int
	StakerAddress    common.Address
}, error) {
	ret := new(struct {
		Status           *big.Int
		CreatedEpoch     *big.Int
		CreatedTime      *big.Int
		DeactivatedEpoch *big.Int
		DeactivatedTime  *big.Int
		StakeAmount      *big.Int
		PaidUntilEpoch   *big.Int
		DelegatedMe      *big.Int
		StakerAddress    common.Address
	})
	out := ret
	err := _SfcContract.contract.Call(opts, out, "stakers", arg0)
	return *ret, err
}

// Stakers is a free data retrieval call binding the contract method 0xfd5e6dd1.
//
// Solidity: function stakers(uint256 ) constant returns(uint256 status, uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime, uint256 stakeAmount, uint256 paidUntilEpoch, uint256 delegatedMe, address stakerAddress)
func (_SfcContract *SfcContractSession) Stakers(arg0 *big.Int) (struct {
	Status           *big.Int
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	DeactivatedEpoch *big.Int
	DeactivatedTime  *big.Int
	StakeAmount      *big.Int
	PaidUntilEpoch   *big.Int
	DelegatedMe      *big.Int
	StakerAddress    common.Address
}, error) {
	return _SfcContract.Contract.Stakers(&_SfcContract.CallOpts, arg0)
}

// Stakers is a free data retrieval call binding the contract method 0xfd5e6dd1.
//
// Solidity: function stakers(uint256 ) constant returns(uint256 status, uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime, uint256 stakeAmount, uint256 paidUntilEpoch, uint256 delegatedMe, address stakerAddress)
func (_SfcContract *SfcContractCallerSession) Stakers(arg0 *big.Int) (struct {
	Status           *big.Int
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	DeactivatedEpoch *big.Int
	DeactivatedTime  *big.Int
	StakeAmount      *big.Int
	PaidUntilEpoch   *big.Int
	DelegatedMe      *big.Int
	StakerAddress    common.Address
}, error) {
	return _SfcContract.Contract.Stakers(&_SfcContract.CallOpts, arg0)
}

// StakersLastID is a free data retrieval call binding the contract method 0x81d9dc7a.
//
// Solidity: function stakersLastID() constant returns(uint256)
func (_SfcContract *SfcContractCaller) StakersLastID(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SfcContract.contract.Call(opts, out, "stakersLastID")
	return *ret0, err
}

// StakersLastID is a free data retrieval call binding the contract method 0x81d9dc7a.
//
// Solidity: function stakersLastID() constant returns(uint256)
func (_SfcContract *SfcContractSession) StakersLastID() (*big.Int, error) {
	return _SfcContract.Contract.StakersLastID(&_SfcContract.CallOpts)
}

// StakersLastID is a free data retrieval call binding the contract method 0x81d9dc7a.
//
// Solidity: function stakersLastID() constant returns(uint256)
func (_SfcContract *SfcContractCallerSession) StakersLastID() (*big.Int, error) {
	return _SfcContract.Contract.StakersLastID(&_SfcContract.CallOpts)
}

// StakersNum is a free data retrieval call binding the contract method 0x08728f6e.
//
// Solidity: function stakersNum() constant returns(uint256)
func (_SfcContract *SfcContractCaller) StakersNum(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SfcContract.contract.Call(opts, out, "stakersNum")
	return *ret0, err
}

// StakersNum is a free data retrieval call binding the contract method 0x08728f6e.
//
// Solidity: function stakersNum() constant returns(uint256)
func (_SfcContract *SfcContractSession) StakersNum() (*big.Int, error) {
	return _SfcContract.Contract.StakersNum(&_SfcContract.CallOpts)
}

// StakersNum is a free data retrieval call binding the contract method 0x08728f6e.
//
// Solidity: function stakersNum() constant returns(uint256)
func (_SfcContract *SfcContractCallerSession) StakersNum() (*big.Int, error) {
	return _SfcContract.Contract.StakersNum(&_SfcContract.CallOpts)
}

// UnbondingPeriod is a free data retrieval call binding the contract method 0x6cf6d675.
//
// Solidity: function unbondingPeriod() constant returns(uint256)
func (_SfcContract *SfcContractCaller) UnbondingPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SfcContract.contract.Call(opts, out, "unbondingPeriod")
	return *ret0, err
}

// UnbondingPeriod is a free data retrieval call binding the contract method 0x6cf6d675.
//
// Solidity: function unbondingPeriod() constant returns(uint256)
func (_SfcContract *SfcContractSession) UnbondingPeriod() (*big.Int, error) {
	return _SfcContract.Contract.UnbondingPeriod(&_SfcContract.CallOpts)
}

// UnbondingPeriod is a free data retrieval call binding the contract method 0x6cf6d675.
//
// Solidity: function unbondingPeriod() constant returns(uint256)
func (_SfcContract *SfcContractCallerSession) UnbondingPeriod() (*big.Int, error) {
	return _SfcContract.Contract.UnbondingPeriod(&_SfcContract.CallOpts)
}

// UnbondingStartDate is a free data retrieval call binding the contract method 0x53a72586.
//
// Solidity: function unbondingStartDate() constant returns(uint256)
func (_SfcContract *SfcContractCaller) UnbondingStartDate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SfcContract.contract.Call(opts, out, "unbondingStartDate")
	return *ret0, err
}

// UnbondingStartDate is a free data retrieval call binding the contract method 0x53a72586.
//
// Solidity: function unbondingStartDate() constant returns(uint256)
func (_SfcContract *SfcContractSession) UnbondingStartDate() (*big.Int, error) {
	return _SfcContract.Contract.UnbondingStartDate(&_SfcContract.CallOpts)
}

// UnbondingStartDate is a free data retrieval call binding the contract method 0x53a72586.
//
// Solidity: function unbondingStartDate() constant returns(uint256)
func (_SfcContract *SfcContractCallerSession) UnbondingStartDate() (*big.Int, error) {
	return _SfcContract.Contract.UnbondingStartDate(&_SfcContract.CallOpts)
}

// UnbondingUnlockPeriod is a free data retrieval call binding the contract method 0x3a0af4d4.
//
// Solidity: function unbondingUnlockPeriod() constant returns(uint256)
func (_SfcContract *SfcContractCaller) UnbondingUnlockPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SfcContract.contract.Call(opts, out, "unbondingUnlockPeriod")
	return *ret0, err
}

// UnbondingUnlockPeriod is a free data retrieval call binding the contract method 0x3a0af4d4.
//
// Solidity: function unbondingUnlockPeriod() constant returns(uint256)
func (_SfcContract *SfcContractSession) UnbondingUnlockPeriod() (*big.Int, error) {
	return _SfcContract.Contract.UnbondingUnlockPeriod(&_SfcContract.CallOpts)
}

// UnbondingUnlockPeriod is a free data retrieval call binding the contract method 0x3a0af4d4.
//
// Solidity: function unbondingUnlockPeriod() constant returns(uint256)
func (_SfcContract *SfcContractCallerSession) UnbondingUnlockPeriod() (*big.Int, error) {
	return _SfcContract.Contract.UnbondingUnlockPeriod(&_SfcContract.CallOpts)
}

// ValidatorCommission is a free data retrieval call binding the contract method 0xa7786515.
//
// Solidity: function validatorCommission() constant returns(uint256)
func (_SfcContract *SfcContractCaller) ValidatorCommission(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SfcContract.contract.Call(opts, out, "validatorCommission")
	return *ret0, err
}

// ValidatorCommission is a free data retrieval call binding the contract method 0xa7786515.
//
// Solidity: function validatorCommission() constant returns(uint256)
func (_SfcContract *SfcContractSession) ValidatorCommission() (*big.Int, error) {
	return _SfcContract.Contract.ValidatorCommission(&_SfcContract.CallOpts)
}

// ValidatorCommission is a free data retrieval call binding the contract method 0xa7786515.
//
// Solidity: function validatorCommission() constant returns(uint256)
func (_SfcContract *SfcContractCallerSession) ValidatorCommission() (*big.Int, error) {
	return _SfcContract.Contract.ValidatorCommission(&_SfcContract.CallOpts)
}

// ClaimDelegationRewards is a paid mutator transaction binding the contract method 0xf99837e6.
//
// Solidity: function claimDelegationRewards(uint256 _fromEpoch, uint256 maxEpochs) returns()
func (_SfcContract *SfcContractTransactor) ClaimDelegationRewards(opts *bind.TransactOpts, _fromEpoch *big.Int, maxEpochs *big.Int) (*types.Transaction, error) {
	return _SfcContract.contract.Transact(opts, "claimDelegationRewards", _fromEpoch, maxEpochs)
}

// ClaimDelegationRewards is a paid mutator transaction binding the contract method 0xf99837e6.
//
// Solidity: function claimDelegationRewards(uint256 _fromEpoch, uint256 maxEpochs) returns()
func (_SfcContract *SfcContractSession) ClaimDelegationRewards(_fromEpoch *big.Int, maxEpochs *big.Int) (*types.Transaction, error) {
	return _SfcContract.Contract.ClaimDelegationRewards(&_SfcContract.TransactOpts, _fromEpoch, maxEpochs)
}

// ClaimDelegationRewards is a paid mutator transaction binding the contract method 0xf99837e6.
//
// Solidity: function claimDelegationRewards(uint256 _fromEpoch, uint256 maxEpochs) returns()
func (_SfcContract *SfcContractTransactorSession) ClaimDelegationRewards(_fromEpoch *big.Int, maxEpochs *big.Int) (*types.Transaction, error) {
	return _SfcContract.Contract.ClaimDelegationRewards(&_SfcContract.TransactOpts, _fromEpoch, maxEpochs)
}

// ClaimValidatorRewards is a paid mutator transaction binding the contract method 0xf0f947c8.
//
// Solidity: function claimValidatorRewards(uint256 _fromEpoch, uint256 maxEpochs) returns()
func (_SfcContract *SfcContractTransactor) ClaimValidatorRewards(opts *bind.TransactOpts, _fromEpoch *big.Int, maxEpochs *big.Int) (*types.Transaction, error) {
	return _SfcContract.contract.Transact(opts, "claimValidatorRewards", _fromEpoch, maxEpochs)
}

// ClaimValidatorRewards is a paid mutator transaction binding the contract method 0xf0f947c8.
//
// Solidity: function claimValidatorRewards(uint256 _fromEpoch, uint256 maxEpochs) returns()
func (_SfcContract *SfcContractSession) ClaimValidatorRewards(_fromEpoch *big.Int, maxEpochs *big.Int) (*types.Transaction, error) {
	return _SfcContract.Contract.ClaimValidatorRewards(&_SfcContract.TransactOpts, _fromEpoch, maxEpochs)
}

// ClaimValidatorRewards is a paid mutator transaction binding the contract method 0xf0f947c8.
//
// Solidity: function claimValidatorRewards(uint256 _fromEpoch, uint256 maxEpochs) returns()
func (_SfcContract *SfcContractTransactorSession) ClaimValidatorRewards(_fromEpoch *big.Int, maxEpochs *big.Int) (*types.Transaction, error) {
	return _SfcContract.Contract.ClaimValidatorRewards(&_SfcContract.TransactOpts, _fromEpoch, maxEpochs)
}

// CreateDelegation is a paid mutator transaction binding the contract method 0xc312eb07.
//
// Solidity: function createDelegation(uint256 to) returns()
func (_SfcContract *SfcContractTransactor) CreateDelegation(opts *bind.TransactOpts, to *big.Int) (*types.Transaction, error) {
	return _SfcContract.contract.Transact(opts, "createDelegation", to)
}

// CreateDelegation is a paid mutator transaction binding the contract method 0xc312eb07.
//
// Solidity: function createDelegation(uint256 to) returns()
func (_SfcContract *SfcContractSession) CreateDelegation(to *big.Int) (*types.Transaction, error) {
	return _SfcContract.Contract.CreateDelegation(&_SfcContract.TransactOpts, to)
}

// CreateDelegation is a paid mutator transaction binding the contract method 0xc312eb07.
//
// Solidity: function createDelegation(uint256 to) returns()
func (_SfcContract *SfcContractTransactorSession) CreateDelegation(to *big.Int) (*types.Transaction, error) {
	return _SfcContract.Contract.CreateDelegation(&_SfcContract.TransactOpts, to)
}

// CreateStake is a paid mutator transaction binding the contract method 0xcc8c2120.
//
// Solidity: function createStake(bytes metadata) returns()
func (_SfcContract *SfcContractTransactor) CreateStake(opts *bind.TransactOpts, metadata []byte) (*types.Transaction, error) {
	return _SfcContract.contract.Transact(opts, "createStake", metadata)
}

// CreateStake is a paid mutator transaction binding the contract method 0xcc8c2120.
//
// Solidity: function createStake(bytes metadata) returns()
func (_SfcContract *SfcContractSession) CreateStake(metadata []byte) (*types.Transaction, error) {
	return _SfcContract.Contract.CreateStake(&_SfcContract.TransactOpts, metadata)
}

// CreateStake is a paid mutator transaction binding the contract method 0xcc8c2120.
//
// Solidity: function createStake(bytes metadata) returns()
func (_SfcContract *SfcContractTransactorSession) CreateStake(metadata []byte) (*types.Transaction, error) {
	return _SfcContract.Contract.CreateStake(&_SfcContract.TransactOpts, metadata)
}

// IncreaseStake is a paid mutator transaction binding the contract method 0xd9e257ef.
//
// Solidity: function increaseStake() returns()
func (_SfcContract *SfcContractTransactor) IncreaseStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SfcContract.contract.Transact(opts, "increaseStake")
}

// IncreaseStake is a paid mutator transaction binding the contract method 0xd9e257ef.
//
// Solidity: function increaseStake() returns()
func (_SfcContract *SfcContractSession) IncreaseStake() (*types.Transaction, error) {
	return _SfcContract.Contract.IncreaseStake(&_SfcContract.TransactOpts)
}

// IncreaseStake is a paid mutator transaction binding the contract method 0xd9e257ef.
//
// Solidity: function increaseStake() returns()
func (_SfcContract *SfcContractTransactorSession) IncreaseStake() (*types.Transaction, error) {
	return _SfcContract.Contract.IncreaseStake(&_SfcContract.TransactOpts)
}

// PrepareToWithdrawDelegation is a paid mutator transaction binding the contract method 0x1c333318.
//
// Solidity: function prepareToWithdrawDelegation() returns()
func (_SfcContract *SfcContractTransactor) PrepareToWithdrawDelegation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SfcContract.contract.Transact(opts, "prepareToWithdrawDelegation")
}

// PrepareToWithdrawDelegation is a paid mutator transaction binding the contract method 0x1c333318.
//
// Solidity: function prepareToWithdrawDelegation() returns()
func (_SfcContract *SfcContractSession) PrepareToWithdrawDelegation() (*types.Transaction, error) {
	return _SfcContract.Contract.PrepareToWithdrawDelegation(&_SfcContract.TransactOpts)
}

// PrepareToWithdrawDelegation is a paid mutator transaction binding the contract method 0x1c333318.
//
// Solidity: function prepareToWithdrawDelegation() returns()
func (_SfcContract *SfcContractTransactorSession) PrepareToWithdrawDelegation() (*types.Transaction, error) {
	return _SfcContract.Contract.PrepareToWithdrawDelegation(&_SfcContract.TransactOpts)
}

// PrepareToWithdrawStake is a paid mutator transaction binding the contract method 0xc41b6405.
//
// Solidity: function prepareToWithdrawStake() returns()
func (_SfcContract *SfcContractTransactor) PrepareToWithdrawStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SfcContract.contract.Transact(opts, "prepareToWithdrawStake")
}

// PrepareToWithdrawStake is a paid mutator transaction binding the contract method 0xc41b6405.
//
// Solidity: function prepareToWithdrawStake() returns()
func (_SfcContract *SfcContractSession) PrepareToWithdrawStake() (*types.Transaction, error) {
	return _SfcContract.Contract.PrepareToWithdrawStake(&_SfcContract.TransactOpts)
}

// PrepareToWithdrawStake is a paid mutator transaction binding the contract method 0xc41b6405.
//
// Solidity: function prepareToWithdrawStake() returns()
func (_SfcContract *SfcContractTransactorSession) PrepareToWithdrawStake() (*types.Transaction, error) {
	return _SfcContract.Contract.PrepareToWithdrawStake(&_SfcContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SfcContract *SfcContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SfcContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SfcContract *SfcContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _SfcContract.Contract.RenounceOwnership(&_SfcContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SfcContract *SfcContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SfcContract.Contract.RenounceOwnership(&_SfcContract.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SfcContract *SfcContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SfcContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SfcContract *SfcContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SfcContract.Contract.TransferOwnership(&_SfcContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SfcContract *SfcContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SfcContract.Contract.TransferOwnership(&_SfcContract.TransactOpts, newOwner)
}

// UpdateBaseRewardPerSec is a paid mutator transaction binding the contract method 0x1b593d8a.
//
// Solidity: function updateBaseRewardPerSec(uint256 value) returns()
func (_SfcContract *SfcContractTransactor) UpdateBaseRewardPerSec(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _SfcContract.contract.Transact(opts, "updateBaseRewardPerSec", value)
}

// UpdateBaseRewardPerSec is a paid mutator transaction binding the contract method 0x1b593d8a.
//
// Solidity: function updateBaseRewardPerSec(uint256 value) returns()
func (_SfcContract *SfcContractSession) UpdateBaseRewardPerSec(value *big.Int) (*types.Transaction, error) {
	return _SfcContract.Contract.UpdateBaseRewardPerSec(&_SfcContract.TransactOpts, value)
}

// UpdateBaseRewardPerSec is a paid mutator transaction binding the contract method 0x1b593d8a.
//
// Solidity: function updateBaseRewardPerSec(uint256 value) returns()
func (_SfcContract *SfcContractTransactorSession) UpdateBaseRewardPerSec(value *big.Int) (*types.Transaction, error) {
	return _SfcContract.Contract.UpdateBaseRewardPerSec(&_SfcContract.TransactOpts, value)
}

// UpdateCapReachedDate is a paid mutator transaction binding the contract method 0x39190939.
//
// Solidity: function updateCapReachedDate() returns()
func (_SfcContract *SfcContractTransactor) UpdateCapReachedDate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SfcContract.contract.Transact(opts, "updateCapReachedDate")
}

// UpdateCapReachedDate is a paid mutator transaction binding the contract method 0x39190939.
//
// Solidity: function updateCapReachedDate() returns()
func (_SfcContract *SfcContractSession) UpdateCapReachedDate() (*types.Transaction, error) {
	return _SfcContract.Contract.UpdateCapReachedDate(&_SfcContract.TransactOpts)
}

// UpdateCapReachedDate is a paid mutator transaction binding the contract method 0x39190939.
//
// Solidity: function updateCapReachedDate() returns()
func (_SfcContract *SfcContractTransactorSession) UpdateCapReachedDate() (*types.Transaction, error) {
	return _SfcContract.Contract.UpdateCapReachedDate(&_SfcContract.TransactOpts)
}

// UpdateGasPowerAllocationRate is a paid mutator transaction binding the contract method 0x119e351a.
//
// Solidity: function updateGasPowerAllocationRate(uint256 short, uint256 long) returns()
func (_SfcContract *SfcContractTransactor) UpdateGasPowerAllocationRate(opts *bind.TransactOpts, short *big.Int, long *big.Int) (*types.Transaction, error) {
	return _SfcContract.contract.Transact(opts, "updateGasPowerAllocationRate", short, long)
}

// UpdateGasPowerAllocationRate is a paid mutator transaction binding the contract method 0x119e351a.
//
// Solidity: function updateGasPowerAllocationRate(uint256 short, uint256 long) returns()
func (_SfcContract *SfcContractSession) UpdateGasPowerAllocationRate(short *big.Int, long *big.Int) (*types.Transaction, error) {
	return _SfcContract.Contract.UpdateGasPowerAllocationRate(&_SfcContract.TransactOpts, short, long)
}

// UpdateGasPowerAllocationRate is a paid mutator transaction binding the contract method 0x119e351a.
//
// Solidity: function updateGasPowerAllocationRate(uint256 short, uint256 long) returns()
func (_SfcContract *SfcContractTransactorSession) UpdateGasPowerAllocationRate(short *big.Int, long *big.Int) (*types.Transaction, error) {
	return _SfcContract.Contract.UpdateGasPowerAllocationRate(&_SfcContract.TransactOpts, short, long)
}

// UpdateStakerMetadata is a paid mutator transaction binding the contract method 0x33a14912.
//
// Solidity: function updateStakerMetadata(bytes metadata) returns()
func (_SfcContract *SfcContractTransactor) UpdateStakerMetadata(opts *bind.TransactOpts, metadata []byte) (*types.Transaction, error) {
	return _SfcContract.contract.Transact(opts, "updateStakerMetadata", metadata)
}

// UpdateStakerMetadata is a paid mutator transaction binding the contract method 0x33a14912.
//
// Solidity: function updateStakerMetadata(bytes metadata) returns()
func (_SfcContract *SfcContractSession) UpdateStakerMetadata(metadata []byte) (*types.Transaction, error) {
	return _SfcContract.Contract.UpdateStakerMetadata(&_SfcContract.TransactOpts, metadata)
}

// UpdateStakerMetadata is a paid mutator transaction binding the contract method 0x33a14912.
//
// Solidity: function updateStakerMetadata(bytes metadata) returns()
func (_SfcContract *SfcContractTransactorSession) UpdateStakerMetadata(metadata []byte) (*types.Transaction, error) {
	return _SfcContract.Contract.UpdateStakerMetadata(&_SfcContract.TransactOpts, metadata)
}

// WithdrawDelegation is a paid mutator transaction binding the contract method 0x16bfdd81.
//
// Solidity: function withdrawDelegation() returns()
func (_SfcContract *SfcContractTransactor) WithdrawDelegation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SfcContract.contract.Transact(opts, "withdrawDelegation")
}

// WithdrawDelegation is a paid mutator transaction binding the contract method 0x16bfdd81.
//
// Solidity: function withdrawDelegation() returns()
func (_SfcContract *SfcContractSession) WithdrawDelegation() (*types.Transaction, error) {
	return _SfcContract.Contract.WithdrawDelegation(&_SfcContract.TransactOpts)
}

// WithdrawDelegation is a paid mutator transaction binding the contract method 0x16bfdd81.
//
// Solidity: function withdrawDelegation() returns()
func (_SfcContract *SfcContractTransactorSession) WithdrawDelegation() (*types.Transaction, error) {
	return _SfcContract.Contract.WithdrawDelegation(&_SfcContract.TransactOpts)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_SfcContract *SfcContractTransactor) WithdrawStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SfcContract.contract.Transact(opts, "withdrawStake")
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_SfcContract *SfcContractSession) WithdrawStake() (*types.Transaction, error) {
	return _SfcContract.Contract.WithdrawStake(&_SfcContract.TransactOpts)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_SfcContract *SfcContractTransactorSession) WithdrawStake() (*types.Transaction, error) {
	return _SfcContract.Contract.WithdrawStake(&_SfcContract.TransactOpts)
}

// SfcContractChangedCapReachedDateIterator is returned from FilterChangedCapReachedDate and is used to iterate over the raw logs and unpacked data for ChangedCapReachedDate events raised by the SfcContract contract.
type SfcContractChangedCapReachedDateIterator struct {
	Event *SfcContractChangedCapReachedDate // Event containing the contract specifics and raw log

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
func (it *SfcContractChangedCapReachedDateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcContractChangedCapReachedDate)
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
		it.Event = new(SfcContractChangedCapReachedDate)
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
func (it *SfcContractChangedCapReachedDateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcContractChangedCapReachedDateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcContractChangedCapReachedDate represents a ChangedCapReachedDate event raised by the SfcContract contract.
type SfcContractChangedCapReachedDate struct {
	CapReachedDate *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterChangedCapReachedDate is a free log retrieval operation binding the contract event 0xe1470d87e1b53ff05cc02eba0dd5210423e111eaad054cfd6752324671d271ec.
//
// Solidity: event ChangedCapReachedDate(uint256 _capReachedDate)
func (_SfcContract *SfcContractFilterer) FilterChangedCapReachedDate(opts *bind.FilterOpts) (*SfcContractChangedCapReachedDateIterator, error) {

	logs, sub, err := _SfcContract.contract.FilterLogs(opts, "ChangedCapReachedDate")
	if err != nil {
		return nil, err
	}
	return &SfcContractChangedCapReachedDateIterator{contract: _SfcContract.contract, event: "ChangedCapReachedDate", logs: logs, sub: sub}, nil
}

// WatchChangedCapReachedDate is a free log subscription operation binding the contract event 0xe1470d87e1b53ff05cc02eba0dd5210423e111eaad054cfd6752324671d271ec.
//
// Solidity: event ChangedCapReachedDate(uint256 _capReachedDate)
func (_SfcContract *SfcContractFilterer) WatchChangedCapReachedDate(opts *bind.WatchOpts, sink chan<- *SfcContractChangedCapReachedDate) (event.Subscription, error) {

	logs, sub, err := _SfcContract.contract.WatchLogs(opts, "ChangedCapReachedDate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcContractChangedCapReachedDate)
				if err := _SfcContract.contract.UnpackLog(event, "ChangedCapReachedDate", log); err != nil {
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

// ParseChangedCapReachedDate is a log parse operation binding the contract event 0xe1470d87e1b53ff05cc02eba0dd5210423e111eaad054cfd6752324671d271ec.
//
// Solidity: event ChangedCapReachedDate(uint256 _capReachedDate)
func (_SfcContract *SfcContractFilterer) ParseChangedCapReachedDate(log types.Log) (*SfcContractChangedCapReachedDate, error) {
	event := new(SfcContractChangedCapReachedDate)
	if err := _SfcContract.contract.UnpackLog(event, "ChangedCapReachedDate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SfcContractClaimedDelegationRewardIterator is returned from FilterClaimedDelegationReward and is used to iterate over the raw logs and unpacked data for ClaimedDelegationReward events raised by the SfcContract contract.
type SfcContractClaimedDelegationRewardIterator struct {
	Event *SfcContractClaimedDelegationReward // Event containing the contract specifics and raw log

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
func (it *SfcContractClaimedDelegationRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcContractClaimedDelegationReward)
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
		it.Event = new(SfcContractClaimedDelegationReward)
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
func (it *SfcContractClaimedDelegationRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcContractClaimedDelegationRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcContractClaimedDelegationReward represents a ClaimedDelegationReward event raised by the SfcContract contract.
type SfcContractClaimedDelegationReward struct {
	From       common.Address
	StakerID   *big.Int
	Reward     *big.Int
	FromEpoch  *big.Int
	UntilEpoch *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterClaimedDelegationReward is a free log retrieval operation binding the contract event 0x2676e1697cf4731b93ddb4ef54e0e5a98c06cccbbbb2202848a3c6286595e6ce.
//
// Solidity: event ClaimedDelegationReward(address indexed from, uint256 indexed stakerID, uint256 reward, uint256 fromEpoch, uint256 untilEpoch)
func (_SfcContract *SfcContractFilterer) FilterClaimedDelegationReward(opts *bind.FilterOpts, from []common.Address, stakerID []*big.Int) (*SfcContractClaimedDelegationRewardIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _SfcContract.contract.FilterLogs(opts, "ClaimedDelegationReward", fromRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &SfcContractClaimedDelegationRewardIterator{contract: _SfcContract.contract, event: "ClaimedDelegationReward", logs: logs, sub: sub}, nil
}

// WatchClaimedDelegationReward is a free log subscription operation binding the contract event 0x2676e1697cf4731b93ddb4ef54e0e5a98c06cccbbbb2202848a3c6286595e6ce.
//
// Solidity: event ClaimedDelegationReward(address indexed from, uint256 indexed stakerID, uint256 reward, uint256 fromEpoch, uint256 untilEpoch)
func (_SfcContract *SfcContractFilterer) WatchClaimedDelegationReward(opts *bind.WatchOpts, sink chan<- *SfcContractClaimedDelegationReward, from []common.Address, stakerID []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _SfcContract.contract.WatchLogs(opts, "ClaimedDelegationReward", fromRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcContractClaimedDelegationReward)
				if err := _SfcContract.contract.UnpackLog(event, "ClaimedDelegationReward", log); err != nil {
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

// ParseClaimedDelegationReward is a log parse operation binding the contract event 0x2676e1697cf4731b93ddb4ef54e0e5a98c06cccbbbb2202848a3c6286595e6ce.
//
// Solidity: event ClaimedDelegationReward(address indexed from, uint256 indexed stakerID, uint256 reward, uint256 fromEpoch, uint256 untilEpoch)
func (_SfcContract *SfcContractFilterer) ParseClaimedDelegationReward(log types.Log) (*SfcContractClaimedDelegationReward, error) {
	event := new(SfcContractClaimedDelegationReward)
	if err := _SfcContract.contract.UnpackLog(event, "ClaimedDelegationReward", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SfcContractClaimedValidatorRewardIterator is returned from FilterClaimedValidatorReward and is used to iterate over the raw logs and unpacked data for ClaimedValidatorReward events raised by the SfcContract contract.
type SfcContractClaimedValidatorRewardIterator struct {
	Event *SfcContractClaimedValidatorReward // Event containing the contract specifics and raw log

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
func (it *SfcContractClaimedValidatorRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcContractClaimedValidatorReward)
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
		it.Event = new(SfcContractClaimedValidatorReward)
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
func (it *SfcContractClaimedValidatorRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcContractClaimedValidatorRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcContractClaimedValidatorReward represents a ClaimedValidatorReward event raised by the SfcContract contract.
type SfcContractClaimedValidatorReward struct {
	StakerID   *big.Int
	Reward     *big.Int
	FromEpoch  *big.Int
	UntilEpoch *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterClaimedValidatorReward is a free log retrieval operation binding the contract event 0x2ea54c2b22a07549d19fb5eb8e4e48ebe1c653117215e94d5468c5612750d35c.
//
// Solidity: event ClaimedValidatorReward(uint256 indexed stakerID, uint256 reward, uint256 fromEpoch, uint256 untilEpoch)
func (_SfcContract *SfcContractFilterer) FilterClaimedValidatorReward(opts *bind.FilterOpts, stakerID []*big.Int) (*SfcContractClaimedValidatorRewardIterator, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _SfcContract.contract.FilterLogs(opts, "ClaimedValidatorReward", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &SfcContractClaimedValidatorRewardIterator{contract: _SfcContract.contract, event: "ClaimedValidatorReward", logs: logs, sub: sub}, nil
}

// WatchClaimedValidatorReward is a free log subscription operation binding the contract event 0x2ea54c2b22a07549d19fb5eb8e4e48ebe1c653117215e94d5468c5612750d35c.
//
// Solidity: event ClaimedValidatorReward(uint256 indexed stakerID, uint256 reward, uint256 fromEpoch, uint256 untilEpoch)
func (_SfcContract *SfcContractFilterer) WatchClaimedValidatorReward(opts *bind.WatchOpts, sink chan<- *SfcContractClaimedValidatorReward, stakerID []*big.Int) (event.Subscription, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _SfcContract.contract.WatchLogs(opts, "ClaimedValidatorReward", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcContractClaimedValidatorReward)
				if err := _SfcContract.contract.UnpackLog(event, "ClaimedValidatorReward", log); err != nil {
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

// ParseClaimedValidatorReward is a log parse operation binding the contract event 0x2ea54c2b22a07549d19fb5eb8e4e48ebe1c653117215e94d5468c5612750d35c.
//
// Solidity: event ClaimedValidatorReward(uint256 indexed stakerID, uint256 reward, uint256 fromEpoch, uint256 untilEpoch)
func (_SfcContract *SfcContractFilterer) ParseClaimedValidatorReward(log types.Log) (*SfcContractClaimedValidatorReward, error) {
	event := new(SfcContractClaimedValidatorReward)
	if err := _SfcContract.contract.UnpackLog(event, "ClaimedValidatorReward", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SfcContractCreatedDelegationIterator is returned from FilterCreatedDelegation and is used to iterate over the raw logs and unpacked data for CreatedDelegation events raised by the SfcContract contract.
type SfcContractCreatedDelegationIterator struct {
	Event *SfcContractCreatedDelegation // Event containing the contract specifics and raw log

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
func (it *SfcContractCreatedDelegationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcContractCreatedDelegation)
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
		it.Event = new(SfcContractCreatedDelegation)
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
func (it *SfcContractCreatedDelegationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcContractCreatedDelegationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcContractCreatedDelegation represents a CreatedDelegation event raised by the SfcContract contract.
type SfcContractCreatedDelegation struct {
	From       common.Address
	ToStakerID *big.Int
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCreatedDelegation is a free log retrieval operation binding the contract event 0xfd8c857fb9acd6f4ad59b8621a2a77825168b7b4b76de9586d08e00d4ed462be.
//
// Solidity: event CreatedDelegation(address indexed from, uint256 indexed toStakerID, uint256 amount)
func (_SfcContract *SfcContractFilterer) FilterCreatedDelegation(opts *bind.FilterOpts, from []common.Address, toStakerID []*big.Int) (*SfcContractCreatedDelegationIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toStakerIDRule []interface{}
	for _, toStakerIDItem := range toStakerID {
		toStakerIDRule = append(toStakerIDRule, toStakerIDItem)
	}

	logs, sub, err := _SfcContract.contract.FilterLogs(opts, "CreatedDelegation", fromRule, toStakerIDRule)
	if err != nil {
		return nil, err
	}
	return &SfcContractCreatedDelegationIterator{contract: _SfcContract.contract, event: "CreatedDelegation", logs: logs, sub: sub}, nil
}

// WatchCreatedDelegation is a free log subscription operation binding the contract event 0xfd8c857fb9acd6f4ad59b8621a2a77825168b7b4b76de9586d08e00d4ed462be.
//
// Solidity: event CreatedDelegation(address indexed from, uint256 indexed toStakerID, uint256 amount)
func (_SfcContract *SfcContractFilterer) WatchCreatedDelegation(opts *bind.WatchOpts, sink chan<- *SfcContractCreatedDelegation, from []common.Address, toStakerID []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toStakerIDRule []interface{}
	for _, toStakerIDItem := range toStakerID {
		toStakerIDRule = append(toStakerIDRule, toStakerIDItem)
	}

	logs, sub, err := _SfcContract.contract.WatchLogs(opts, "CreatedDelegation", fromRule, toStakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcContractCreatedDelegation)
				if err := _SfcContract.contract.UnpackLog(event, "CreatedDelegation", log); err != nil {
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

// ParseCreatedDelegation is a log parse operation binding the contract event 0xfd8c857fb9acd6f4ad59b8621a2a77825168b7b4b76de9586d08e00d4ed462be.
//
// Solidity: event CreatedDelegation(address indexed from, uint256 indexed toStakerID, uint256 amount)
func (_SfcContract *SfcContractFilterer) ParseCreatedDelegation(log types.Log) (*SfcContractCreatedDelegation, error) {
	event := new(SfcContractCreatedDelegation)
	if err := _SfcContract.contract.UnpackLog(event, "CreatedDelegation", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SfcContractCreatedStakeIterator is returned from FilterCreatedStake and is used to iterate over the raw logs and unpacked data for CreatedStake events raised by the SfcContract contract.
type SfcContractCreatedStakeIterator struct {
	Event *SfcContractCreatedStake // Event containing the contract specifics and raw log

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
func (it *SfcContractCreatedStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcContractCreatedStake)
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
		it.Event = new(SfcContractCreatedStake)
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
func (it *SfcContractCreatedStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcContractCreatedStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcContractCreatedStake represents a CreatedStake event raised by the SfcContract contract.
type SfcContractCreatedStake struct {
	StakerID      *big.Int
	StakerAddress common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCreatedStake is a free log retrieval operation binding the contract event 0x0697dfe5062b9db8108e4b31254f47a912ae6bbb78837667b2e923a6f5160d39.
//
// Solidity: event CreatedStake(uint256 indexed stakerID, address indexed stakerAddress, uint256 amount)
func (_SfcContract *SfcContractFilterer) FilterCreatedStake(opts *bind.FilterOpts, stakerID []*big.Int, stakerAddress []common.Address) (*SfcContractCreatedStakeIterator, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}
	var stakerAddressRule []interface{}
	for _, stakerAddressItem := range stakerAddress {
		stakerAddressRule = append(stakerAddressRule, stakerAddressItem)
	}

	logs, sub, err := _SfcContract.contract.FilterLogs(opts, "CreatedStake", stakerIDRule, stakerAddressRule)
	if err != nil {
		return nil, err
	}
	return &SfcContractCreatedStakeIterator{contract: _SfcContract.contract, event: "CreatedStake", logs: logs, sub: sub}, nil
}

// WatchCreatedStake is a free log subscription operation binding the contract event 0x0697dfe5062b9db8108e4b31254f47a912ae6bbb78837667b2e923a6f5160d39.
//
// Solidity: event CreatedStake(uint256 indexed stakerID, address indexed stakerAddress, uint256 amount)
func (_SfcContract *SfcContractFilterer) WatchCreatedStake(opts *bind.WatchOpts, sink chan<- *SfcContractCreatedStake, stakerID []*big.Int, stakerAddress []common.Address) (event.Subscription, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}
	var stakerAddressRule []interface{}
	for _, stakerAddressItem := range stakerAddress {
		stakerAddressRule = append(stakerAddressRule, stakerAddressItem)
	}

	logs, sub, err := _SfcContract.contract.WatchLogs(opts, "CreatedStake", stakerIDRule, stakerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcContractCreatedStake)
				if err := _SfcContract.contract.UnpackLog(event, "CreatedStake", log); err != nil {
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

// ParseCreatedStake is a log parse operation binding the contract event 0x0697dfe5062b9db8108e4b31254f47a912ae6bbb78837667b2e923a6f5160d39.
//
// Solidity: event CreatedStake(uint256 indexed stakerID, address indexed stakerAddress, uint256 amount)
func (_SfcContract *SfcContractFilterer) ParseCreatedStake(log types.Log) (*SfcContractCreatedStake, error) {
	event := new(SfcContractCreatedStake)
	if err := _SfcContract.contract.UnpackLog(event, "CreatedStake", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SfcContractIncreasedStakeIterator is returned from FilterIncreasedStake and is used to iterate over the raw logs and unpacked data for IncreasedStake events raised by the SfcContract contract.
type SfcContractIncreasedStakeIterator struct {
	Event *SfcContractIncreasedStake // Event containing the contract specifics and raw log

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
func (it *SfcContractIncreasedStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcContractIncreasedStake)
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
		it.Event = new(SfcContractIncreasedStake)
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
func (it *SfcContractIncreasedStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcContractIncreasedStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcContractIncreasedStake represents a IncreasedStake event raised by the SfcContract contract.
type SfcContractIncreasedStake struct {
	StakerID  *big.Int
	NewAmount *big.Int
	Diff      *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterIncreasedStake is a free log retrieval operation binding the contract event 0xa1d93e9a2a16bf4c2d0cdc6f47fe0fa054c741c96b3dac1297c79eaca31714e9.
//
// Solidity: event IncreasedStake(uint256 indexed stakerID, uint256 newAmount, uint256 diff)
func (_SfcContract *SfcContractFilterer) FilterIncreasedStake(opts *bind.FilterOpts, stakerID []*big.Int) (*SfcContractIncreasedStakeIterator, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _SfcContract.contract.FilterLogs(opts, "IncreasedStake", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &SfcContractIncreasedStakeIterator{contract: _SfcContract.contract, event: "IncreasedStake", logs: logs, sub: sub}, nil
}

// WatchIncreasedStake is a free log subscription operation binding the contract event 0xa1d93e9a2a16bf4c2d0cdc6f47fe0fa054c741c96b3dac1297c79eaca31714e9.
//
// Solidity: event IncreasedStake(uint256 indexed stakerID, uint256 newAmount, uint256 diff)
func (_SfcContract *SfcContractFilterer) WatchIncreasedStake(opts *bind.WatchOpts, sink chan<- *SfcContractIncreasedStake, stakerID []*big.Int) (event.Subscription, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _SfcContract.contract.WatchLogs(opts, "IncreasedStake", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcContractIncreasedStake)
				if err := _SfcContract.contract.UnpackLog(event, "IncreasedStake", log); err != nil {
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

// ParseIncreasedStake is a log parse operation binding the contract event 0xa1d93e9a2a16bf4c2d0cdc6f47fe0fa054c741c96b3dac1297c79eaca31714e9.
//
// Solidity: event IncreasedStake(uint256 indexed stakerID, uint256 newAmount, uint256 diff)
func (_SfcContract *SfcContractFilterer) ParseIncreasedStake(log types.Log) (*SfcContractIncreasedStake, error) {
	event := new(SfcContractIncreasedStake)
	if err := _SfcContract.contract.UnpackLog(event, "IncreasedStake", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SfcContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SfcContract contract.
type SfcContractOwnershipTransferredIterator struct {
	Event *SfcContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SfcContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcContractOwnershipTransferred)
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
		it.Event = new(SfcContractOwnershipTransferred)
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
func (it *SfcContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcContractOwnershipTransferred represents a OwnershipTransferred event raised by the SfcContract contract.
type SfcContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SfcContract *SfcContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SfcContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SfcContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SfcContractOwnershipTransferredIterator{contract: _SfcContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SfcContract *SfcContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SfcContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SfcContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcContractOwnershipTransferred)
				if err := _SfcContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SfcContract *SfcContractFilterer) ParseOwnershipTransferred(log types.Log) (*SfcContractOwnershipTransferred, error) {
	event := new(SfcContractOwnershipTransferred)
	if err := _SfcContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SfcContractPreparedToWithdrawDelegationIterator is returned from FilterPreparedToWithdrawDelegation and is used to iterate over the raw logs and unpacked data for PreparedToWithdrawDelegation events raised by the SfcContract contract.
type SfcContractPreparedToWithdrawDelegationIterator struct {
	Event *SfcContractPreparedToWithdrawDelegation // Event containing the contract specifics and raw log

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
func (it *SfcContractPreparedToWithdrawDelegationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcContractPreparedToWithdrawDelegation)
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
		it.Event = new(SfcContractPreparedToWithdrawDelegation)
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
func (it *SfcContractPreparedToWithdrawDelegationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcContractPreparedToWithdrawDelegationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcContractPreparedToWithdrawDelegation represents a PreparedToWithdrawDelegation event raised by the SfcContract contract.
type SfcContractPreparedToWithdrawDelegation struct {
	From     common.Address
	StakerID *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPreparedToWithdrawDelegation is a free log retrieval operation binding the contract event 0x5b1eea49e405ef6d509836aac841959c30bb0673b1fd70859bfc6ae5e4ee3df2.
//
// Solidity: event PreparedToWithdrawDelegation(address indexed from, uint256 indexed stakerID)
func (_SfcContract *SfcContractFilterer) FilterPreparedToWithdrawDelegation(opts *bind.FilterOpts, from []common.Address, stakerID []*big.Int) (*SfcContractPreparedToWithdrawDelegationIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _SfcContract.contract.FilterLogs(opts, "PreparedToWithdrawDelegation", fromRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &SfcContractPreparedToWithdrawDelegationIterator{contract: _SfcContract.contract, event: "PreparedToWithdrawDelegation", logs: logs, sub: sub}, nil
}

// WatchPreparedToWithdrawDelegation is a free log subscription operation binding the contract event 0x5b1eea49e405ef6d509836aac841959c30bb0673b1fd70859bfc6ae5e4ee3df2.
//
// Solidity: event PreparedToWithdrawDelegation(address indexed from, uint256 indexed stakerID)
func (_SfcContract *SfcContractFilterer) WatchPreparedToWithdrawDelegation(opts *bind.WatchOpts, sink chan<- *SfcContractPreparedToWithdrawDelegation, from []common.Address, stakerID []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _SfcContract.contract.WatchLogs(opts, "PreparedToWithdrawDelegation", fromRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcContractPreparedToWithdrawDelegation)
				if err := _SfcContract.contract.UnpackLog(event, "PreparedToWithdrawDelegation", log); err != nil {
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

// ParsePreparedToWithdrawDelegation is a log parse operation binding the contract event 0x5b1eea49e405ef6d509836aac841959c30bb0673b1fd70859bfc6ae5e4ee3df2.
//
// Solidity: event PreparedToWithdrawDelegation(address indexed from, uint256 indexed stakerID)
func (_SfcContract *SfcContractFilterer) ParsePreparedToWithdrawDelegation(log types.Log) (*SfcContractPreparedToWithdrawDelegation, error) {
	event := new(SfcContractPreparedToWithdrawDelegation)
	if err := _SfcContract.contract.UnpackLog(event, "PreparedToWithdrawDelegation", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SfcContractPreparedToWithdrawStakeIterator is returned from FilterPreparedToWithdrawStake and is used to iterate over the raw logs and unpacked data for PreparedToWithdrawStake events raised by the SfcContract contract.
type SfcContractPreparedToWithdrawStakeIterator struct {
	Event *SfcContractPreparedToWithdrawStake // Event containing the contract specifics and raw log

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
func (it *SfcContractPreparedToWithdrawStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcContractPreparedToWithdrawStake)
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
		it.Event = new(SfcContractPreparedToWithdrawStake)
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
func (it *SfcContractPreparedToWithdrawStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcContractPreparedToWithdrawStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcContractPreparedToWithdrawStake represents a PreparedToWithdrawStake event raised by the SfcContract contract.
type SfcContractPreparedToWithdrawStake struct {
	StakerID *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPreparedToWithdrawStake is a free log retrieval operation binding the contract event 0x84244546a9da4942f506db48ff90ebc240c73bb399e3e47d58843c6bb60e7185.
//
// Solidity: event PreparedToWithdrawStake(uint256 indexed stakerID)
func (_SfcContract *SfcContractFilterer) FilterPreparedToWithdrawStake(opts *bind.FilterOpts, stakerID []*big.Int) (*SfcContractPreparedToWithdrawStakeIterator, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _SfcContract.contract.FilterLogs(opts, "PreparedToWithdrawStake", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &SfcContractPreparedToWithdrawStakeIterator{contract: _SfcContract.contract, event: "PreparedToWithdrawStake", logs: logs, sub: sub}, nil
}

// WatchPreparedToWithdrawStake is a free log subscription operation binding the contract event 0x84244546a9da4942f506db48ff90ebc240c73bb399e3e47d58843c6bb60e7185.
//
// Solidity: event PreparedToWithdrawStake(uint256 indexed stakerID)
func (_SfcContract *SfcContractFilterer) WatchPreparedToWithdrawStake(opts *bind.WatchOpts, sink chan<- *SfcContractPreparedToWithdrawStake, stakerID []*big.Int) (event.Subscription, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _SfcContract.contract.WatchLogs(opts, "PreparedToWithdrawStake", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcContractPreparedToWithdrawStake)
				if err := _SfcContract.contract.UnpackLog(event, "PreparedToWithdrawStake", log); err != nil {
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

// ParsePreparedToWithdrawStake is a log parse operation binding the contract event 0x84244546a9da4942f506db48ff90ebc240c73bb399e3e47d58843c6bb60e7185.
//
// Solidity: event PreparedToWithdrawStake(uint256 indexed stakerID)
func (_SfcContract *SfcContractFilterer) ParsePreparedToWithdrawStake(log types.Log) (*SfcContractPreparedToWithdrawStake, error) {
	event := new(SfcContractPreparedToWithdrawStake)
	if err := _SfcContract.contract.UnpackLog(event, "PreparedToWithdrawStake", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SfcContractUpdatedBaseRewardPerSecIterator is returned from FilterUpdatedBaseRewardPerSec and is used to iterate over the raw logs and unpacked data for UpdatedBaseRewardPerSec events raised by the SfcContract contract.
type SfcContractUpdatedBaseRewardPerSecIterator struct {
	Event *SfcContractUpdatedBaseRewardPerSec // Event containing the contract specifics and raw log

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
func (it *SfcContractUpdatedBaseRewardPerSecIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcContractUpdatedBaseRewardPerSec)
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
		it.Event = new(SfcContractUpdatedBaseRewardPerSec)
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
func (it *SfcContractUpdatedBaseRewardPerSecIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcContractUpdatedBaseRewardPerSecIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcContractUpdatedBaseRewardPerSec represents a UpdatedBaseRewardPerSec event raised by the SfcContract contract.
type SfcContractUpdatedBaseRewardPerSec struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdatedBaseRewardPerSec is a free log retrieval operation binding the contract event 0x8cd9dae1bbea2bc8a5e80ffce2c224727a25925130a03ae100619a8861ae2396.
//
// Solidity: event UpdatedBaseRewardPerSec(uint256 value)
func (_SfcContract *SfcContractFilterer) FilterUpdatedBaseRewardPerSec(opts *bind.FilterOpts) (*SfcContractUpdatedBaseRewardPerSecIterator, error) {

	logs, sub, err := _SfcContract.contract.FilterLogs(opts, "UpdatedBaseRewardPerSec")
	if err != nil {
		return nil, err
	}
	return &SfcContractUpdatedBaseRewardPerSecIterator{contract: _SfcContract.contract, event: "UpdatedBaseRewardPerSec", logs: logs, sub: sub}, nil
}

// WatchUpdatedBaseRewardPerSec is a free log subscription operation binding the contract event 0x8cd9dae1bbea2bc8a5e80ffce2c224727a25925130a03ae100619a8861ae2396.
//
// Solidity: event UpdatedBaseRewardPerSec(uint256 value)
func (_SfcContract *SfcContractFilterer) WatchUpdatedBaseRewardPerSec(opts *bind.WatchOpts, sink chan<- *SfcContractUpdatedBaseRewardPerSec) (event.Subscription, error) {

	logs, sub, err := _SfcContract.contract.WatchLogs(opts, "UpdatedBaseRewardPerSec")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcContractUpdatedBaseRewardPerSec)
				if err := _SfcContract.contract.UnpackLog(event, "UpdatedBaseRewardPerSec", log); err != nil {
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

// ParseUpdatedBaseRewardPerSec is a log parse operation binding the contract event 0x8cd9dae1bbea2bc8a5e80ffce2c224727a25925130a03ae100619a8861ae2396.
//
// Solidity: event UpdatedBaseRewardPerSec(uint256 value)
func (_SfcContract *SfcContractFilterer) ParseUpdatedBaseRewardPerSec(log types.Log) (*SfcContractUpdatedBaseRewardPerSec, error) {
	event := new(SfcContractUpdatedBaseRewardPerSec)
	if err := _SfcContract.contract.UnpackLog(event, "UpdatedBaseRewardPerSec", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SfcContractUpdatedGasPowerAllocationRateIterator is returned from FilterUpdatedGasPowerAllocationRate and is used to iterate over the raw logs and unpacked data for UpdatedGasPowerAllocationRate events raised by the SfcContract contract.
type SfcContractUpdatedGasPowerAllocationRateIterator struct {
	Event *SfcContractUpdatedGasPowerAllocationRate // Event containing the contract specifics and raw log

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
func (it *SfcContractUpdatedGasPowerAllocationRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcContractUpdatedGasPowerAllocationRate)
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
		it.Event = new(SfcContractUpdatedGasPowerAllocationRate)
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
func (it *SfcContractUpdatedGasPowerAllocationRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcContractUpdatedGasPowerAllocationRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcContractUpdatedGasPowerAllocationRate represents a UpdatedGasPowerAllocationRate event raised by the SfcContract contract.
type SfcContractUpdatedGasPowerAllocationRate struct {
	Short *big.Int
	Long  *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdatedGasPowerAllocationRate is a free log retrieval operation binding the contract event 0x95ae5488127de4bc98492f4487556e7af9f37eb4b6d5e94f6d849e03ff76cc7c.
//
// Solidity: event UpdatedGasPowerAllocationRate(uint256 short, uint256 long)
func (_SfcContract *SfcContractFilterer) FilterUpdatedGasPowerAllocationRate(opts *bind.FilterOpts) (*SfcContractUpdatedGasPowerAllocationRateIterator, error) {

	logs, sub, err := _SfcContract.contract.FilterLogs(opts, "UpdatedGasPowerAllocationRate")
	if err != nil {
		return nil, err
	}
	return &SfcContractUpdatedGasPowerAllocationRateIterator{contract: _SfcContract.contract, event: "UpdatedGasPowerAllocationRate", logs: logs, sub: sub}, nil
}

// WatchUpdatedGasPowerAllocationRate is a free log subscription operation binding the contract event 0x95ae5488127de4bc98492f4487556e7af9f37eb4b6d5e94f6d849e03ff76cc7c.
//
// Solidity: event UpdatedGasPowerAllocationRate(uint256 short, uint256 long)
func (_SfcContract *SfcContractFilterer) WatchUpdatedGasPowerAllocationRate(opts *bind.WatchOpts, sink chan<- *SfcContractUpdatedGasPowerAllocationRate) (event.Subscription, error) {

	logs, sub, err := _SfcContract.contract.WatchLogs(opts, "UpdatedGasPowerAllocationRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcContractUpdatedGasPowerAllocationRate)
				if err := _SfcContract.contract.UnpackLog(event, "UpdatedGasPowerAllocationRate", log); err != nil {
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

// ParseUpdatedGasPowerAllocationRate is a log parse operation binding the contract event 0x95ae5488127de4bc98492f4487556e7af9f37eb4b6d5e94f6d849e03ff76cc7c.
//
// Solidity: event UpdatedGasPowerAllocationRate(uint256 short, uint256 long)
func (_SfcContract *SfcContractFilterer) ParseUpdatedGasPowerAllocationRate(log types.Log) (*SfcContractUpdatedGasPowerAllocationRate, error) {
	event := new(SfcContractUpdatedGasPowerAllocationRate)
	if err := _SfcContract.contract.UnpackLog(event, "UpdatedGasPowerAllocationRate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SfcContractUpdatedStakerMetadataIterator is returned from FilterUpdatedStakerMetadata and is used to iterate over the raw logs and unpacked data for UpdatedStakerMetadata events raised by the SfcContract contract.
type SfcContractUpdatedStakerMetadataIterator struct {
	Event *SfcContractUpdatedStakerMetadata // Event containing the contract specifics and raw log

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
func (it *SfcContractUpdatedStakerMetadataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcContractUpdatedStakerMetadata)
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
		it.Event = new(SfcContractUpdatedStakerMetadata)
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
func (it *SfcContractUpdatedStakerMetadataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcContractUpdatedStakerMetadataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcContractUpdatedStakerMetadata represents a UpdatedStakerMetadata event raised by the SfcContract contract.
type SfcContractUpdatedStakerMetadata struct {
	StakerID *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdatedStakerMetadata is a free log retrieval operation binding the contract event 0xb7a99a0df6a9e15c2689e6a55811ef76cdb514c67d4a0e37fcb125ada0e3cd83.
//
// Solidity: event UpdatedStakerMetadata(uint256 indexed stakerID)
func (_SfcContract *SfcContractFilterer) FilterUpdatedStakerMetadata(opts *bind.FilterOpts, stakerID []*big.Int) (*SfcContractUpdatedStakerMetadataIterator, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _SfcContract.contract.FilterLogs(opts, "UpdatedStakerMetadata", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &SfcContractUpdatedStakerMetadataIterator{contract: _SfcContract.contract, event: "UpdatedStakerMetadata", logs: logs, sub: sub}, nil
}

// WatchUpdatedStakerMetadata is a free log subscription operation binding the contract event 0xb7a99a0df6a9e15c2689e6a55811ef76cdb514c67d4a0e37fcb125ada0e3cd83.
//
// Solidity: event UpdatedStakerMetadata(uint256 indexed stakerID)
func (_SfcContract *SfcContractFilterer) WatchUpdatedStakerMetadata(opts *bind.WatchOpts, sink chan<- *SfcContractUpdatedStakerMetadata, stakerID []*big.Int) (event.Subscription, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _SfcContract.contract.WatchLogs(opts, "UpdatedStakerMetadata", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcContractUpdatedStakerMetadata)
				if err := _SfcContract.contract.UnpackLog(event, "UpdatedStakerMetadata", log); err != nil {
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

// ParseUpdatedStakerMetadata is a log parse operation binding the contract event 0xb7a99a0df6a9e15c2689e6a55811ef76cdb514c67d4a0e37fcb125ada0e3cd83.
//
// Solidity: event UpdatedStakerMetadata(uint256 indexed stakerID)
func (_SfcContract *SfcContractFilterer) ParseUpdatedStakerMetadata(log types.Log) (*SfcContractUpdatedStakerMetadata, error) {
	event := new(SfcContractUpdatedStakerMetadata)
	if err := _SfcContract.contract.UnpackLog(event, "UpdatedStakerMetadata", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SfcContractWithdrawnDelegationIterator is returned from FilterWithdrawnDelegation and is used to iterate over the raw logs and unpacked data for WithdrawnDelegation events raised by the SfcContract contract.
type SfcContractWithdrawnDelegationIterator struct {
	Event *SfcContractWithdrawnDelegation // Event containing the contract specifics and raw log

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
func (it *SfcContractWithdrawnDelegationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcContractWithdrawnDelegation)
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
		it.Event = new(SfcContractWithdrawnDelegation)
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
func (it *SfcContractWithdrawnDelegationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcContractWithdrawnDelegationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcContractWithdrawnDelegation represents a WithdrawnDelegation event raised by the SfcContract contract.
type SfcContractWithdrawnDelegation struct {
	From     common.Address
	StakerID *big.Int
	Penalty  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnDelegation is a free log retrieval operation binding the contract event 0x87e86b3710b72c10173ca52c6a9f9cf2df27e77ed177741a8b4feb12bb7a606f.
//
// Solidity: event WithdrawnDelegation(address indexed from, uint256 indexed stakerID, uint256 penalty)
func (_SfcContract *SfcContractFilterer) FilterWithdrawnDelegation(opts *bind.FilterOpts, from []common.Address, stakerID []*big.Int) (*SfcContractWithdrawnDelegationIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _SfcContract.contract.FilterLogs(opts, "WithdrawnDelegation", fromRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &SfcContractWithdrawnDelegationIterator{contract: _SfcContract.contract, event: "WithdrawnDelegation", logs: logs, sub: sub}, nil
}

// WatchWithdrawnDelegation is a free log subscription operation binding the contract event 0x87e86b3710b72c10173ca52c6a9f9cf2df27e77ed177741a8b4feb12bb7a606f.
//
// Solidity: event WithdrawnDelegation(address indexed from, uint256 indexed stakerID, uint256 penalty)
func (_SfcContract *SfcContractFilterer) WatchWithdrawnDelegation(opts *bind.WatchOpts, sink chan<- *SfcContractWithdrawnDelegation, from []common.Address, stakerID []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _SfcContract.contract.WatchLogs(opts, "WithdrawnDelegation", fromRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcContractWithdrawnDelegation)
				if err := _SfcContract.contract.UnpackLog(event, "WithdrawnDelegation", log); err != nil {
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

// ParseWithdrawnDelegation is a log parse operation binding the contract event 0x87e86b3710b72c10173ca52c6a9f9cf2df27e77ed177741a8b4feb12bb7a606f.
//
// Solidity: event WithdrawnDelegation(address indexed from, uint256 indexed stakerID, uint256 penalty)
func (_SfcContract *SfcContractFilterer) ParseWithdrawnDelegation(log types.Log) (*SfcContractWithdrawnDelegation, error) {
	event := new(SfcContractWithdrawnDelegation)
	if err := _SfcContract.contract.UnpackLog(event, "WithdrawnDelegation", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SfcContractWithdrawnStakeIterator is returned from FilterWithdrawnStake and is used to iterate over the raw logs and unpacked data for WithdrawnStake events raised by the SfcContract contract.
type SfcContractWithdrawnStakeIterator struct {
	Event *SfcContractWithdrawnStake // Event containing the contract specifics and raw log

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
func (it *SfcContractWithdrawnStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcContractWithdrawnStake)
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
		it.Event = new(SfcContractWithdrawnStake)
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
func (it *SfcContractWithdrawnStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcContractWithdrawnStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcContractWithdrawnStake represents a WithdrawnStake event raised by the SfcContract contract.
type SfcContractWithdrawnStake struct {
	StakerID *big.Int
	Penalty  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnStake is a free log retrieval operation binding the contract event 0x8c6548258f8f12a9d4b593fa89a223417ed901d4ee9712ba09beb4d56f5262b6.
//
// Solidity: event WithdrawnStake(uint256 indexed stakerID, uint256 penalty)
func (_SfcContract *SfcContractFilterer) FilterWithdrawnStake(opts *bind.FilterOpts, stakerID []*big.Int) (*SfcContractWithdrawnStakeIterator, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _SfcContract.contract.FilterLogs(opts, "WithdrawnStake", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &SfcContractWithdrawnStakeIterator{contract: _SfcContract.contract, event: "WithdrawnStake", logs: logs, sub: sub}, nil
}

// WatchWithdrawnStake is a free log subscription operation binding the contract event 0x8c6548258f8f12a9d4b593fa89a223417ed901d4ee9712ba09beb4d56f5262b6.
//
// Solidity: event WithdrawnStake(uint256 indexed stakerID, uint256 penalty)
func (_SfcContract *SfcContractFilterer) WatchWithdrawnStake(opts *bind.WatchOpts, sink chan<- *SfcContractWithdrawnStake, stakerID []*big.Int) (event.Subscription, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _SfcContract.contract.WatchLogs(opts, "WithdrawnStake", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcContractWithdrawnStake)
				if err := _SfcContract.contract.UnpackLog(event, "WithdrawnStake", log); err != nil {
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

// ParseWithdrawnStake is a log parse operation binding the contract event 0x8c6548258f8f12a9d4b593fa89a223417ed901d4ee9712ba09beb4d56f5262b6.
//
// Solidity: event WithdrawnStake(uint256 indexed stakerID, uint256 penalty)
func (_SfcContract *SfcContractFilterer) ParseWithdrawnStake(log types.Log) (*SfcContractWithdrawnStake, error) {
	event := new(SfcContractWithdrawnStake)
	if err := _SfcContract.contract.UnpackLog(event, "WithdrawnStake", log); err != nil {
		return nil, err
	}
	return event, nil
}
