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

// SfcV1ContractMetaData contains all meta data concerning the SfcV1Contract contract.
var SfcV1ContractMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"minDelegation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bondedRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakersNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"slashedStakeTotalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"short\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"long\",\"type\":\"uint256\"}],\"name\":\"updateGasPowerAllocationRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawDelegation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minStakeDecrease\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"updateBaseRewardPerSec\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"prepareToWithdrawDelegation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"delegationLockPeriodEpochs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"epochSnapshots\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBaseRewardWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalTxRewardWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseRewardPerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeTotalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegationsTotalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxDelegatedRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wrID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"prepareToWithdrawStakePartial\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractCommission\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"}],\"name\":\"_upgradeStakerStorage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxEpochs\",\"type\":\"uint256\"}],\"name\":\"claimValidatorRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"delegationsTotalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"updateStakerMetadata\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"unbondingUnlockPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"increaseDelegation\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakeTotalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakeLockPeriodTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"_syncDelegator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"delegationsNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdrawalRequests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegation\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"unbondingStartDate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commission\",\"type\":\"uint256\"}],\"name\":\"_calcValidatorEpochReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakeLockPeriodEpochs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minDelegationIncrease\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getStakerID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"_calcRawValidatorEpochReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bondedTargetRewardUnlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegationAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commission\",\"type\":\"uint256\"}],\"name\":\"_calcDelegationEpochReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewardsStash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxEpochs\",\"type\":\"uint256\"}],\"name\":\"claimDelegationRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bondedTargetPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentSealedEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakersLastID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardsAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unstashRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dagAdrress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sfcAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"createStakeWithAddresses\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fromEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxEpochs\",\"type\":\"uint256\"}],\"name\":\"calcValidatorRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakerMetadata\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"slashedDelegationsTotalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"validatorCommission\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxStakerMetadataSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isDelegation\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"}],\"name\":\"_rewardsBurnableOnDeactivation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sfcAddress\",\"type\":\"address\"}],\"name\":\"_sfcAddressToStakerID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_fromEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxEpochs\",\"type\":\"uint256\"}],\"name\":\"calcDelegationRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"e\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"}],\"name\":\"epochValidator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatedMe\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseRewardWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"txRewardWeight\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"delegations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"createdEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paidUntilEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toStakerID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"}],\"name\":\"createDelegation\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newSfcAddress\",\"type\":\"address\"}],\"name\":\"updateStakerSfcAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"prepareToWithdrawStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minStakeIncrease\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minDelegationDecrease\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"createStake\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bondedTargetStart\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"increaseStake\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wrID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"prepareToWithdrawDelegationPartial\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"}],\"name\":\"_syncStaker\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"delegationLockPeriodTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wrID\",\"type\":\"uint256\"}],\"name\":\"partialWithdrawByRequest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paidUntilEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatedMe\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dagAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sfcAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dagSfcAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CreatedStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldSfcAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newSfcAddress\",\"type\":\"address\"}],\"name\":\"UpdatedStakerSfcAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"}],\"name\":\"UpdatedStakerMetadata\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"diff\",\"type\":\"uint256\"}],\"name\":\"IncreasedStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toStakerID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CreatedDelegation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"diff\",\"type\":\"uint256\"}],\"name\":\"IncreasedDelegation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromEpoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"untilEpoch\",\"type\":\"uint256\"}],\"name\":\"ClaimedDelegationReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromEpoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"untilEpoch\",\"type\":\"uint256\"}],\"name\":\"ClaimedValidatorReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"auth\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"name\":\"UnstashedRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isDelegation\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BurntRewardStash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"}],\"name\":\"PreparedToWithdrawStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"}],\"name\":\"DeactivatedStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"auth\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wrID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"delegation\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CreatedWithdrawRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"penalty\",\"type\":\"uint256\"}],\"name\":\"WithdrawnStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"}],\"name\":\"PreparedToWithdrawDelegation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"}],\"name\":\"DeactivatedDelegation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"penalty\",\"type\":\"uint256\"}],\"name\":\"WithdrawnDelegation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"auth\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wrID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"delegation\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"penalty\",\"type\":\"uint256\"}],\"name\":\"PartialWithdrawnByRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldStakerID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newStakerID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UpdatedDelegation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delegatedMe\",\"type\":\"uint256\"}],\"name\":\"UpdatedStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"UpdatedBaseRewardPerSec\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"short\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"long\",\"type\":\"uint256\"}],\"name\":\"UpdatedGasPowerAllocationRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]",
}

// SfcV1ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use SfcV1ContractMetaData.ABI instead.
var SfcV1ContractABI = SfcV1ContractMetaData.ABI

// SfcV1Contract is an auto generated Go binding around an Ethereum contract.
type SfcV1Contract struct {
	SfcV1ContractCaller     // Read-only binding to the contract
	SfcV1ContractTransactor // Write-only binding to the contract
	SfcV1ContractFilterer   // Log filterer for contract events
}

// SfcV1ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type SfcV1ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SfcV1ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SfcV1ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SfcV1ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SfcV1ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SfcV1ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SfcV1ContractSession struct {
	Contract     *SfcV1Contract    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SfcV1ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SfcV1ContractCallerSession struct {
	Contract *SfcV1ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SfcV1ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SfcV1ContractTransactorSession struct {
	Contract     *SfcV1ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SfcV1ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type SfcV1ContractRaw struct {
	Contract *SfcV1Contract // Generic contract binding to access the raw methods on
}

// SfcV1ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SfcV1ContractCallerRaw struct {
	Contract *SfcV1ContractCaller // Generic read-only contract binding to access the raw methods on
}

// SfcV1ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SfcV1ContractTransactorRaw struct {
	Contract *SfcV1ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSfcV1Contract creates a new instance of SfcV1Contract, bound to a specific deployed contract.
func NewSfcV1Contract(address common.Address, backend bind.ContractBackend) (*SfcV1Contract, error) {
	contract, err := bindSfcV1Contract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SfcV1Contract{SfcV1ContractCaller: SfcV1ContractCaller{contract: contract}, SfcV1ContractTransactor: SfcV1ContractTransactor{contract: contract}, SfcV1ContractFilterer: SfcV1ContractFilterer{contract: contract}}, nil
}

// NewSfcV1ContractCaller creates a new read-only instance of SfcV1Contract, bound to a specific deployed contract.
func NewSfcV1ContractCaller(address common.Address, caller bind.ContractCaller) (*SfcV1ContractCaller, error) {
	contract, err := bindSfcV1Contract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SfcV1ContractCaller{contract: contract}, nil
}

// NewSfcV1ContractTransactor creates a new write-only instance of SfcV1Contract, bound to a specific deployed contract.
func NewSfcV1ContractTransactor(address common.Address, transactor bind.ContractTransactor) (*SfcV1ContractTransactor, error) {
	contract, err := bindSfcV1Contract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SfcV1ContractTransactor{contract: contract}, nil
}

// NewSfcV1ContractFilterer creates a new log filterer instance of SfcV1Contract, bound to a specific deployed contract.
func NewSfcV1ContractFilterer(address common.Address, filterer bind.ContractFilterer) (*SfcV1ContractFilterer, error) {
	contract, err := bindSfcV1Contract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SfcV1ContractFilterer{contract: contract}, nil
}

// bindSfcV1Contract binds a generic wrapper to an already deployed contract.
func bindSfcV1Contract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SfcV1ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SfcV1Contract *SfcV1ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SfcV1Contract.Contract.SfcV1ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SfcV1Contract *SfcV1ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SfcV1Contract.Contract.SfcV1ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SfcV1Contract *SfcV1ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SfcV1Contract.Contract.SfcV1ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SfcV1Contract *SfcV1ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SfcV1Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SfcV1Contract *SfcV1ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SfcV1Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SfcV1Contract *SfcV1ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SfcV1Contract.Contract.contract.Transact(opts, method, params...)
}

// CalcDelegationEpochReward is a free data retrieval call binding the contract method 0x6af49928.
//
// Solidity: function _calcDelegationEpochReward(uint256 stakerID, uint256 epoch, uint256 delegationAmount, uint256 commission) view returns(uint256)
func (_SfcV1Contract *SfcV1ContractCaller) CalcDelegationEpochReward(opts *bind.CallOpts, stakerID *big.Int, epoch *big.Int, delegationAmount *big.Int, commission *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SfcV1Contract.contract.Call(opts, &out, "_calcDelegationEpochReward", stakerID, epoch, delegationAmount, commission)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcDelegationEpochReward is a free data retrieval call binding the contract method 0x6af49928.
//
// Solidity: function _calcDelegationEpochReward(uint256 stakerID, uint256 epoch, uint256 delegationAmount, uint256 commission) view returns(uint256)
func (_SfcV1Contract *SfcV1ContractSession) CalcDelegationEpochReward(stakerID *big.Int, epoch *big.Int, delegationAmount *big.Int, commission *big.Int) (*big.Int, error) {
	return _SfcV1Contract.Contract.CalcDelegationEpochReward(&_SfcV1Contract.CallOpts, stakerID, epoch, delegationAmount, commission)
}

// CalcDelegationEpochReward is a free data retrieval call binding the contract method 0x6af49928.
//
// Solidity: function _calcDelegationEpochReward(uint256 stakerID, uint256 epoch, uint256 delegationAmount, uint256 commission) view returns(uint256)
func (_SfcV1Contract *SfcV1ContractCallerSession) CalcDelegationEpochReward(stakerID *big.Int, epoch *big.Int, delegationAmount *big.Int, commission *big.Int) (*big.Int, error) {
	return _SfcV1Contract.Contract.CalcDelegationEpochReward(&_SfcV1Contract.CallOpts, stakerID, epoch, delegationAmount, commission)
}

// CalcRawValidatorEpochReward is a free data retrieval call binding the contract method 0x6459cf86.
//
// Solidity: function _calcRawValidatorEpochReward(uint256 stakerID, uint256 epoch) view returns(uint256)
func (_SfcV1Contract *SfcV1ContractCaller) CalcRawValidatorEpochReward(opts *bind.CallOpts, stakerID *big.Int, epoch *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SfcV1Contract.contract.Call(opts, &out, "_calcRawValidatorEpochReward", stakerID, epoch)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcRawValidatorEpochReward is a free data retrieval call binding the contract method 0x6459cf86.
//
// Solidity: function _calcRawValidatorEpochReward(uint256 stakerID, uint256 epoch) view returns(uint256)
func (_SfcV1Contract *SfcV1ContractSession) CalcRawValidatorEpochReward(stakerID *big.Int, epoch *big.Int) (*big.Int, error) {
	return _SfcV1Contract.Contract.CalcRawValidatorEpochReward(&_SfcV1Contract.CallOpts, stakerID, epoch)
}

// CalcRawValidatorEpochReward is a free data retrieval call binding the contract method 0x6459cf86.
//
// Solidity: function _calcRawValidatorEpochReward(uint256 stakerID, uint256 epoch) view returns(uint256)
func (_SfcV1Contract *SfcV1ContractCallerSession) CalcRawValidatorEpochReward(stakerID *big.Int, epoch *big.Int) (*big.Int, error) {
	return _SfcV1Contract.Contract.CalcRawValidatorEpochReward(&_SfcV1Contract.CallOpts, stakerID, epoch)
}

// CalcValidatorEpochReward is a free data retrieval call binding the contract method 0x53bbb754.
//
// Solidity: function _calcValidatorEpochReward(uint256 stakerID, uint256 epoch, uint256 commission) view returns(uint256)
func (_SfcV1Contract *SfcV1ContractCaller) CalcValidatorEpochReward(opts *bind.CallOpts, stakerID *big.Int, epoch *big.Int, commission *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SfcV1Contract.contract.Call(opts, &out, "_calcValidatorEpochReward", stakerID, epoch, commission)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcValidatorEpochReward is a free data retrieval call binding the contract method 0x53bbb754.
//
// Solidity: function _calcValidatorEpochReward(uint256 stakerID, uint256 epoch, uint256 commission) view returns(uint256)
func (_SfcV1Contract *SfcV1ContractSession) CalcValidatorEpochReward(stakerID *big.Int, epoch *big.Int, commission *big.Int) (*big.Int, error) {
	return _SfcV1Contract.Contract.CalcValidatorEpochReward(&_SfcV1Contract.CallOpts, stakerID, epoch, commission)
}

// CalcValidatorEpochReward is a free data retrieval call binding the contract method 0x53bbb754.
//
// Solidity: function _calcValidatorEpochReward(uint256 stakerID, uint256 epoch, uint256 commission) view returns(uint256)
func (_SfcV1Contract *SfcV1ContractCallerSession) CalcValidatorEpochReward(stakerID *big.Int, epoch *big.Int, commission *big.Int) (*big.Int, error) {
	return _SfcV1Contract.Contract.CalcValidatorEpochReward(&_SfcV1Contract.CallOpts, stakerID, epoch, commission)
}

// RewardsBurnableOnDeactivation is a free data retrieval call binding the contract method 0xb1a3ebfa.
//
// Solidity: function _rewardsBurnableOnDeactivation(bool isDelegation, uint256 stakerID) view returns(bool)
func (_SfcV1Contract *SfcV1ContractCaller) RewardsBurnableOnDeactivation(opts *bind.CallOpts, isDelegation bool, stakerID *big.Int) (bool, error) {
	var out []interface{}
	err := _SfcV1Contract.contract.Call(opts, &out, "_rewardsBurnableOnDeactivation", isDelegation, stakerID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RewardsBurnableOnDeactivation is a free data retrieval call binding the contract method 0xb1a3ebfa.
//
// Solidity: function _rewardsBurnableOnDeactivation(bool isDelegation, uint256 stakerID) view returns(bool)
func (_SfcV1Contract *SfcV1ContractSession) RewardsBurnableOnDeactivation(isDelegation bool, stakerID *big.Int) (bool, error) {
	return _SfcV1Contract.Contract.RewardsBurnableOnDeactivation(&_SfcV1Contract.CallOpts, isDelegation, stakerID)
}

// RewardsBurnableOnDeactivation is a free data retrieval call binding the contract method 0xb1a3ebfa.
//
// Solidity: function _rewardsBurnableOnDeactivation(bool isDelegation, uint256 stakerID) view returns(bool)
func (_SfcV1Contract *SfcV1ContractCallerSession) RewardsBurnableOnDeactivation(isDelegation bool, stakerID *big.Int) (bool, error) {
	return _SfcV1Contract.Contract.RewardsBurnableOnDeactivation(&_SfcV1Contract.CallOpts, isDelegation, stakerID)
}

// SfcAddressToStakerID is a free data retrieval call binding the contract method 0xb42cb58d.
//
// Solidity: function _sfcAddressToStakerID(address sfcAddress) view returns(uint256)
func (_SfcV1Contract *SfcV1ContractCaller) SfcAddressToStakerID(opts *bind.CallOpts, sfcAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SfcV1Contract.contract.Call(opts, &out, "_sfcAddressToStakerID", sfcAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SfcAddressToStakerID is a free data retrieval call binding the contract method 0xb42cb58d.
//
// Solidity: function _sfcAddressToStakerID(address sfcAddress) view returns(uint256)
func (_SfcV1Contract *SfcV1ContractSession) SfcAddressToStakerID(sfcAddress common.Address) (*big.Int, error) {
	return _SfcV1Contract.Contract.SfcAddressToStakerID(&_SfcV1Contract.CallOpts, sfcAddress)
}

// SfcAddressToStakerID is a free data retrieval call binding the contract method 0xb42cb58d.
//
// Solidity: function _sfcAddressToStakerID(address sfcAddress) view returns(uint256)
func (_SfcV1Contract *SfcV1ContractCallerSession) SfcAddressToStakerID(sfcAddress common.Address) (*big.Int, error) {
	return _SfcV1Contract.Contract.SfcAddressToStakerID(&_SfcV1Contract.CallOpts, sfcAddress)
}

// BondedRatio is a free data retrieval call binding the contract method 0x041d97a3.
//
// Solidity: function bondedRatio() view returns(uint256)
func (_SfcV1Contract *SfcV1ContractCaller) BondedRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SfcV1Contract.contract.Call(opts, &out, "bondedRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BondedRatio is a free data retrieval call binding the contract method 0x041d97a3.
//
// Solidity: function bondedRatio() view returns(uint256)
func (_SfcV1Contract *SfcV1ContractSession) BondedRatio() (*big.Int, error) {
	return _SfcV1Contract.Contract.BondedRatio(&_SfcV1Contract.CallOpts)
}

// BondedRatio is a free data retrieval call binding the contract method 0x041d97a3.
//
// Solidity: function bondedRatio() view returns(uint256)
func (_SfcV1Contract *SfcV1ContractCallerSession) BondedRatio() (*big.Int, error) {
	return _SfcV1Contract.Contract.BondedRatio(&_SfcV1Contract.CallOpts)
}

// BondedTargetPeriod is a free data retrieval call binding the contract method 0x7b8c6b02.
//
// Solidity: function bondedTargetPeriod() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractCaller) BondedTargetPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SfcV1Contract.contract.Call(opts, &out, "bondedTargetPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BondedTargetPeriod is a free data retrieval call binding the contract method 0x7b8c6b02.
//
// Solidity: function bondedTargetPeriod() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractSession) BondedTargetPeriod() (*big.Int, error) {
	return _SfcV1Contract.Contract.BondedTargetPeriod(&_SfcV1Contract.CallOpts)
}

// BondedTargetPeriod is a free data retrieval call binding the contract method 0x7b8c6b02.
//
// Solidity: function bondedTargetPeriod() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractCallerSession) BondedTargetPeriod() (*big.Int, error) {
	return _SfcV1Contract.Contract.BondedTargetPeriod(&_SfcV1Contract.CallOpts)
}

// BondedTargetRewardUnlock is a free data retrieval call binding the contract method 0x6a1cf400.
//
// Solidity: function bondedTargetRewardUnlock() view returns(uint256)
func (_SfcV1Contract *SfcV1ContractCaller) BondedTargetRewardUnlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SfcV1Contract.contract.Call(opts, &out, "bondedTargetRewardUnlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BondedTargetRewardUnlock is a free data retrieval call binding the contract method 0x6a1cf400.
//
// Solidity: function bondedTargetRewardUnlock() view returns(uint256)
func (_SfcV1Contract *SfcV1ContractSession) BondedTargetRewardUnlock() (*big.Int, error) {
	return _SfcV1Contract.Contract.BondedTargetRewardUnlock(&_SfcV1Contract.CallOpts)
}

// BondedTargetRewardUnlock is a free data retrieval call binding the contract method 0x6a1cf400.
//
// Solidity: function bondedTargetRewardUnlock() view returns(uint256)
func (_SfcV1Contract *SfcV1ContractCallerSession) BondedTargetRewardUnlock() (*big.Int, error) {
	return _SfcV1Contract.Contract.BondedTargetRewardUnlock(&_SfcV1Contract.CallOpts)
}

// BondedTargetStart is a free data retrieval call binding the contract method 0xce5aa000.
//
// Solidity: function bondedTargetStart() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractCaller) BondedTargetStart(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SfcV1Contract.contract.Call(opts, &out, "bondedTargetStart")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BondedTargetStart is a free data retrieval call binding the contract method 0xce5aa000.
//
// Solidity: function bondedTargetStart() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractSession) BondedTargetStart() (*big.Int, error) {
	return _SfcV1Contract.Contract.BondedTargetStart(&_SfcV1Contract.CallOpts)
}

// BondedTargetStart is a free data retrieval call binding the contract method 0xce5aa000.
//
// Solidity: function bondedTargetStart() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractCallerSession) BondedTargetStart() (*big.Int, error) {
	return _SfcV1Contract.Contract.BondedTargetStart(&_SfcV1Contract.CallOpts)
}

// CalcDelegationRewards is a free data retrieval call binding the contract method 0xb6732f3d.
//
// Solidity: function calcDelegationRewards(address delegator, uint256 _fromEpoch, uint256 maxEpochs) view returns(uint256, uint256, uint256)
func (_SfcV1Contract *SfcV1ContractCaller) CalcDelegationRewards(opts *bind.CallOpts, delegator common.Address, _fromEpoch *big.Int, maxEpochs *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _SfcV1Contract.contract.Call(opts, &out, "calcDelegationRewards", delegator, _fromEpoch, maxEpochs)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// CalcDelegationRewards is a free data retrieval call binding the contract method 0xb6732f3d.
//
// Solidity: function calcDelegationRewards(address delegator, uint256 _fromEpoch, uint256 maxEpochs) view returns(uint256, uint256, uint256)
func (_SfcV1Contract *SfcV1ContractSession) CalcDelegationRewards(delegator common.Address, _fromEpoch *big.Int, maxEpochs *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _SfcV1Contract.Contract.CalcDelegationRewards(&_SfcV1Contract.CallOpts, delegator, _fromEpoch, maxEpochs)
}

// CalcDelegationRewards is a free data retrieval call binding the contract method 0xb6732f3d.
//
// Solidity: function calcDelegationRewards(address delegator, uint256 _fromEpoch, uint256 maxEpochs) view returns(uint256, uint256, uint256)
func (_SfcV1Contract *SfcV1ContractCallerSession) CalcDelegationRewards(delegator common.Address, _fromEpoch *big.Int, maxEpochs *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _SfcV1Contract.Contract.CalcDelegationRewards(&_SfcV1Contract.CallOpts, delegator, _fromEpoch, maxEpochs)
}

// CalcValidatorRewards is a free data retrieval call binding the contract method 0x96060e71.
//
// Solidity: function calcValidatorRewards(uint256 stakerID, uint256 _fromEpoch, uint256 maxEpochs) view returns(uint256, uint256, uint256)
func (_SfcV1Contract *SfcV1ContractCaller) CalcValidatorRewards(opts *bind.CallOpts, stakerID *big.Int, _fromEpoch *big.Int, maxEpochs *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _SfcV1Contract.contract.Call(opts, &out, "calcValidatorRewards", stakerID, _fromEpoch, maxEpochs)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// CalcValidatorRewards is a free data retrieval call binding the contract method 0x96060e71.
//
// Solidity: function calcValidatorRewards(uint256 stakerID, uint256 _fromEpoch, uint256 maxEpochs) view returns(uint256, uint256, uint256)
func (_SfcV1Contract *SfcV1ContractSession) CalcValidatorRewards(stakerID *big.Int, _fromEpoch *big.Int, maxEpochs *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _SfcV1Contract.Contract.CalcValidatorRewards(&_SfcV1Contract.CallOpts, stakerID, _fromEpoch, maxEpochs)
}

// CalcValidatorRewards is a free data retrieval call binding the contract method 0x96060e71.
//
// Solidity: function calcValidatorRewards(uint256 stakerID, uint256 _fromEpoch, uint256 maxEpochs) view returns(uint256, uint256, uint256)
func (_SfcV1Contract *SfcV1ContractCallerSession) CalcValidatorRewards(stakerID *big.Int, _fromEpoch *big.Int, maxEpochs *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _SfcV1Contract.Contract.CalcValidatorRewards(&_SfcV1Contract.CallOpts, stakerID, _fromEpoch, maxEpochs)
}

// ContractCommission is a free data retrieval call binding the contract method 0x2709275e.
//
// Solidity: function contractCommission() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractCaller) ContractCommission(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SfcV1Contract.contract.Call(opts, &out, "contractCommission")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ContractCommission is a free data retrieval call binding the contract method 0x2709275e.
//
// Solidity: function contractCommission() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractSession) ContractCommission() (*big.Int, error) {
	return _SfcV1Contract.Contract.ContractCommission(&_SfcV1Contract.CallOpts)
}

// ContractCommission is a free data retrieval call binding the contract method 0x2709275e.
//
// Solidity: function contractCommission() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractCallerSession) ContractCommission() (*big.Int, error) {
	return _SfcV1Contract.Contract.ContractCommission(&_SfcV1Contract.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_SfcV1Contract *SfcV1ContractCaller) CurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SfcV1Contract.contract.Call(opts, &out, "currentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_SfcV1Contract *SfcV1ContractSession) CurrentEpoch() (*big.Int, error) {
	return _SfcV1Contract.Contract.CurrentEpoch(&_SfcV1Contract.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_SfcV1Contract *SfcV1ContractCallerSession) CurrentEpoch() (*big.Int, error) {
	return _SfcV1Contract.Contract.CurrentEpoch(&_SfcV1Contract.CallOpts)
}

// CurrentSealedEpoch is a free data retrieval call binding the contract method 0x7cacb1d6.
//
// Solidity: function currentSealedEpoch() view returns(uint256)
func (_SfcV1Contract *SfcV1ContractCaller) CurrentSealedEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SfcV1Contract.contract.Call(opts, &out, "currentSealedEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentSealedEpoch is a free data retrieval call binding the contract method 0x7cacb1d6.
//
// Solidity: function currentSealedEpoch() view returns(uint256)
func (_SfcV1Contract *SfcV1ContractSession) CurrentSealedEpoch() (*big.Int, error) {
	return _SfcV1Contract.Contract.CurrentSealedEpoch(&_SfcV1Contract.CallOpts)
}

// CurrentSealedEpoch is a free data retrieval call binding the contract method 0x7cacb1d6.
//
// Solidity: function currentSealedEpoch() view returns(uint256)
func (_SfcV1Contract *SfcV1ContractCallerSession) CurrentSealedEpoch() (*big.Int, error) {
	return _SfcV1Contract.Contract.CurrentSealedEpoch(&_SfcV1Contract.CallOpts)
}

// DelegationLockPeriodEpochs is a free data retrieval call binding the contract method 0x1d58179c.
//
// Solidity: function delegationLockPeriodEpochs() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractCaller) DelegationLockPeriodEpochs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SfcV1Contract.contract.Call(opts, &out, "delegationLockPeriodEpochs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegationLockPeriodEpochs is a free data retrieval call binding the contract method 0x1d58179c.
//
// Solidity: function delegationLockPeriodEpochs() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractSession) DelegationLockPeriodEpochs() (*big.Int, error) {
	return _SfcV1Contract.Contract.DelegationLockPeriodEpochs(&_SfcV1Contract.CallOpts)
}

// DelegationLockPeriodEpochs is a free data retrieval call binding the contract method 0x1d58179c.
//
// Solidity: function delegationLockPeriodEpochs() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractCallerSession) DelegationLockPeriodEpochs() (*big.Int, error) {
	return _SfcV1Contract.Contract.DelegationLockPeriodEpochs(&_SfcV1Contract.CallOpts)
}

// DelegationLockPeriodTime is a free data retrieval call binding the contract method 0xec6a7f1c.
//
// Solidity: function delegationLockPeriodTime() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractCaller) DelegationLockPeriodTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SfcV1Contract.contract.Call(opts, &out, "delegationLockPeriodTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegationLockPeriodTime is a free data retrieval call binding the contract method 0xec6a7f1c.
//
// Solidity: function delegationLockPeriodTime() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractSession) DelegationLockPeriodTime() (*big.Int, error) {
	return _SfcV1Contract.Contract.DelegationLockPeriodTime(&_SfcV1Contract.CallOpts)
}

// DelegationLockPeriodTime is a free data retrieval call binding the contract method 0xec6a7f1c.
//
// Solidity: function delegationLockPeriodTime() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractCallerSession) DelegationLockPeriodTime() (*big.Int, error) {
	return _SfcV1Contract.Contract.DelegationLockPeriodTime(&_SfcV1Contract.CallOpts)
}

// Delegations is a free data retrieval call binding the contract method 0xbffe3486.
//
// Solidity: function delegations(address ) view returns(uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime, uint256 amount, uint256 paidUntilEpoch, uint256 toStakerID)
func (_SfcV1Contract *SfcV1ContractCaller) Delegations(opts *bind.CallOpts, arg0 common.Address) (struct {
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	DeactivatedEpoch *big.Int
	DeactivatedTime  *big.Int
	Amount           *big.Int
	PaidUntilEpoch   *big.Int
	ToStakerID       *big.Int
}, error) {
	var out []interface{}
	err := _SfcV1Contract.contract.Call(opts, &out, "delegations", arg0)

	outstruct := new(struct {
		CreatedEpoch     *big.Int
		CreatedTime      *big.Int
		DeactivatedEpoch *big.Int
		DeactivatedTime  *big.Int
		Amount           *big.Int
		PaidUntilEpoch   *big.Int
		ToStakerID       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CreatedEpoch = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CreatedTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.DeactivatedEpoch = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.DeactivatedTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.PaidUntilEpoch = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.ToStakerID = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Delegations is a free data retrieval call binding the contract method 0xbffe3486.
//
// Solidity: function delegations(address ) view returns(uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime, uint256 amount, uint256 paidUntilEpoch, uint256 toStakerID)
func (_SfcV1Contract *SfcV1ContractSession) Delegations(arg0 common.Address) (struct {
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	DeactivatedEpoch *big.Int
	DeactivatedTime  *big.Int
	Amount           *big.Int
	PaidUntilEpoch   *big.Int
	ToStakerID       *big.Int
}, error) {
	return _SfcV1Contract.Contract.Delegations(&_SfcV1Contract.CallOpts, arg0)
}

// Delegations is a free data retrieval call binding the contract method 0xbffe3486.
//
// Solidity: function delegations(address ) view returns(uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime, uint256 amount, uint256 paidUntilEpoch, uint256 toStakerID)
func (_SfcV1Contract *SfcV1ContractCallerSession) Delegations(arg0 common.Address) (struct {
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	DeactivatedEpoch *big.Int
	DeactivatedTime  *big.Int
	Amount           *big.Int
	PaidUntilEpoch   *big.Int
	ToStakerID       *big.Int
}, error) {
	return _SfcV1Contract.Contract.Delegations(&_SfcV1Contract.CallOpts, arg0)
}

// DelegationsNum is a free data retrieval call binding the contract method 0x4bd202dc.
//
// Solidity: function delegationsNum() view returns(uint256)
func (_SfcV1Contract *SfcV1ContractCaller) DelegationsNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SfcV1Contract.contract.Call(opts, &out, "delegationsNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegationsNum is a free data retrieval call binding the contract method 0x4bd202dc.
//
// Solidity: function delegationsNum() view returns(uint256)
func (_SfcV1Contract *SfcV1ContractSession) DelegationsNum() (*big.Int, error) {
	return _SfcV1Contract.Contract.DelegationsNum(&_SfcV1Contract.CallOpts)
}

// DelegationsNum is a free data retrieval call binding the contract method 0x4bd202dc.
//
// Solidity: function delegationsNum() view returns(uint256)
func (_SfcV1Contract *SfcV1ContractCallerSession) DelegationsNum() (*big.Int, error) {
	return _SfcV1Contract.Contract.DelegationsNum(&_SfcV1Contract.CallOpts)
}

// DelegationsTotalAmount is a free data retrieval call binding the contract method 0x30fa9929.
//
// Solidity: function delegationsTotalAmount() view returns(uint256)
func (_SfcV1Contract *SfcV1ContractCaller) DelegationsTotalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SfcV1Contract.contract.Call(opts, &out, "delegationsTotalAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegationsTotalAmount is a free data retrieval call binding the contract method 0x30fa9929.
//
// Solidity: function delegationsTotalAmount() view returns(uint256)
func (_SfcV1Contract *SfcV1ContractSession) DelegationsTotalAmount() (*big.Int, error) {
	return _SfcV1Contract.Contract.DelegationsTotalAmount(&_SfcV1Contract.CallOpts)
}

// DelegationsTotalAmount is a free data retrieval call binding the contract method 0x30fa9929.
//
// Solidity: function delegationsTotalAmount() view returns(uint256)
func (_SfcV1Contract *SfcV1ContractCallerSession) DelegationsTotalAmount() (*big.Int, error) {
	return _SfcV1Contract.Contract.DelegationsTotalAmount(&_SfcV1Contract.CallOpts)
}

// EpochSnapshots is a free data retrieval call binding the contract method 0x1e8a6956.
//
// Solidity: function epochSnapshots(uint256 ) view returns(uint256 endTime, uint256 duration, uint256 epochFee, uint256 totalBaseRewardWeight, uint256 totalTxRewardWeight, uint256 baseRewardPerSecond, uint256 stakeTotalAmount, uint256 delegationsTotalAmount, uint256 totalSupply)
func (_SfcV1Contract *SfcV1ContractCaller) EpochSnapshots(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	var out []interface{}
	err := _SfcV1Contract.contract.Call(opts, &out, "epochSnapshots", arg0)

	outstruct := new(struct {
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
	if err != nil {
		return *outstruct, err
	}

	outstruct.EndTime = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Duration = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EpochFee = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TotalBaseRewardWeight = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TotalTxRewardWeight = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.BaseRewardPerSecond = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.StakeTotalAmount = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.DelegationsTotalAmount = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.TotalSupply = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// EpochSnapshots is a free data retrieval call binding the contract method 0x1e8a6956.
//
// Solidity: function epochSnapshots(uint256 ) view returns(uint256 endTime, uint256 duration, uint256 epochFee, uint256 totalBaseRewardWeight, uint256 totalTxRewardWeight, uint256 baseRewardPerSecond, uint256 stakeTotalAmount, uint256 delegationsTotalAmount, uint256 totalSupply)
func (_SfcV1Contract *SfcV1ContractSession) EpochSnapshots(arg0 *big.Int) (struct {
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
	return _SfcV1Contract.Contract.EpochSnapshots(&_SfcV1Contract.CallOpts, arg0)
}

// EpochSnapshots is a free data retrieval call binding the contract method 0x1e8a6956.
//
// Solidity: function epochSnapshots(uint256 ) view returns(uint256 endTime, uint256 duration, uint256 epochFee, uint256 totalBaseRewardWeight, uint256 totalTxRewardWeight, uint256 baseRewardPerSecond, uint256 stakeTotalAmount, uint256 delegationsTotalAmount, uint256 totalSupply)
func (_SfcV1Contract *SfcV1ContractCallerSession) EpochSnapshots(arg0 *big.Int) (struct {
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
	return _SfcV1Contract.Contract.EpochSnapshots(&_SfcV1Contract.CallOpts, arg0)
}

// EpochValidator is a free data retrieval call binding the contract method 0xb9029d50.
//
// Solidity: function epochValidator(uint256 e, uint256 v) view returns(uint256 stakeAmount, uint256 delegatedMe, uint256 baseRewardWeight, uint256 txRewardWeight)
func (_SfcV1Contract *SfcV1ContractCaller) EpochValidator(opts *bind.CallOpts, e *big.Int, v *big.Int) (struct {
	StakeAmount      *big.Int
	DelegatedMe      *big.Int
	BaseRewardWeight *big.Int
	TxRewardWeight   *big.Int
}, error) {
	var out []interface{}
	err := _SfcV1Contract.contract.Call(opts, &out, "epochValidator", e, v)

	outstruct := new(struct {
		StakeAmount      *big.Int
		DelegatedMe      *big.Int
		BaseRewardWeight *big.Int
		TxRewardWeight   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StakeAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.DelegatedMe = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BaseRewardWeight = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TxRewardWeight = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// EpochValidator is a free data retrieval call binding the contract method 0xb9029d50.
//
// Solidity: function epochValidator(uint256 e, uint256 v) view returns(uint256 stakeAmount, uint256 delegatedMe, uint256 baseRewardWeight, uint256 txRewardWeight)
func (_SfcV1Contract *SfcV1ContractSession) EpochValidator(e *big.Int, v *big.Int) (struct {
	StakeAmount      *big.Int
	DelegatedMe      *big.Int
	BaseRewardWeight *big.Int
	TxRewardWeight   *big.Int
}, error) {
	return _SfcV1Contract.Contract.EpochValidator(&_SfcV1Contract.CallOpts, e, v)
}

// EpochValidator is a free data retrieval call binding the contract method 0xb9029d50.
//
// Solidity: function epochValidator(uint256 e, uint256 v) view returns(uint256 stakeAmount, uint256 delegatedMe, uint256 baseRewardWeight, uint256 txRewardWeight)
func (_SfcV1Contract *SfcV1ContractCallerSession) EpochValidator(e *big.Int, v *big.Int) (struct {
	StakeAmount      *big.Int
	DelegatedMe      *big.Int
	BaseRewardWeight *big.Int
	TxRewardWeight   *big.Int
}, error) {
	return _SfcV1Contract.Contract.EpochValidator(&_SfcV1Contract.CallOpts, e, v)
}

// GetStakerID is a free data retrieval call binding the contract method 0x63321e27.
//
// Solidity: function getStakerID(address addr) view returns(uint256)
func (_SfcV1Contract *SfcV1ContractCaller) GetStakerID(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SfcV1Contract.contract.Call(opts, &out, "getStakerID", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakerID is a free data retrieval call binding the contract method 0x63321e27.
//
// Solidity: function getStakerID(address addr) view returns(uint256)
func (_SfcV1Contract *SfcV1ContractSession) GetStakerID(addr common.Address) (*big.Int, error) {
	return _SfcV1Contract.Contract.GetStakerID(&_SfcV1Contract.CallOpts, addr)
}

// GetStakerID is a free data retrieval call binding the contract method 0x63321e27.
//
// Solidity: function getStakerID(address addr) view returns(uint256)
func (_SfcV1Contract *SfcV1ContractCallerSession) GetStakerID(addr common.Address) (*big.Int, error) {
	return _SfcV1Contract.Contract.GetStakerID(&_SfcV1Contract.CallOpts, addr)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_SfcV1Contract *SfcV1ContractCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SfcV1Contract.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_SfcV1Contract *SfcV1ContractSession) IsOwner() (bool, error) {
	return _SfcV1Contract.Contract.IsOwner(&_SfcV1Contract.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_SfcV1Contract *SfcV1ContractCallerSession) IsOwner() (bool, error) {
	return _SfcV1Contract.Contract.IsOwner(&_SfcV1Contract.CallOpts)
}

// MaxDelegatedRatio is a free data retrieval call binding the contract method 0x2265f284.
//
// Solidity: function maxDelegatedRatio() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractCaller) MaxDelegatedRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SfcV1Contract.contract.Call(opts, &out, "maxDelegatedRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDelegatedRatio is a free data retrieval call binding the contract method 0x2265f284.
//
// Solidity: function maxDelegatedRatio() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractSession) MaxDelegatedRatio() (*big.Int, error) {
	return _SfcV1Contract.Contract.MaxDelegatedRatio(&_SfcV1Contract.CallOpts)
}

// MaxDelegatedRatio is a free data retrieval call binding the contract method 0x2265f284.
//
// Solidity: function maxDelegatedRatio() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractCallerSession) MaxDelegatedRatio() (*big.Int, error) {
	return _SfcV1Contract.Contract.MaxDelegatedRatio(&_SfcV1Contract.CallOpts)
}

// MaxStakerMetadataSize is a free data retrieval call binding the contract method 0xab2273c0.
//
// Solidity: function maxStakerMetadataSize() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractCaller) MaxStakerMetadataSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SfcV1Contract.contract.Call(opts, &out, "maxStakerMetadataSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxStakerMetadataSize is a free data retrieval call binding the contract method 0xab2273c0.
//
// Solidity: function maxStakerMetadataSize() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractSession) MaxStakerMetadataSize() (*big.Int, error) {
	return _SfcV1Contract.Contract.MaxStakerMetadataSize(&_SfcV1Contract.CallOpts)
}

// MaxStakerMetadataSize is a free data retrieval call binding the contract method 0xab2273c0.
//
// Solidity: function maxStakerMetadataSize() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractCallerSession) MaxStakerMetadataSize() (*big.Int, error) {
	return _SfcV1Contract.Contract.MaxStakerMetadataSize(&_SfcV1Contract.CallOpts)
}

// MinDelegation is a free data retrieval call binding the contract method 0x02985992.
//
// Solidity: function minDelegation() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractCaller) MinDelegation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SfcV1Contract.contract.Call(opts, &out, "minDelegation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDelegation is a free data retrieval call binding the contract method 0x02985992.
//
// Solidity: function minDelegation() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractSession) MinDelegation() (*big.Int, error) {
	return _SfcV1Contract.Contract.MinDelegation(&_SfcV1Contract.CallOpts)
}

// MinDelegation is a free data retrieval call binding the contract method 0x02985992.
//
// Solidity: function minDelegation() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractCallerSession) MinDelegation() (*big.Int, error) {
	return _SfcV1Contract.Contract.MinDelegation(&_SfcV1Contract.CallOpts)
}

// MinDelegationDecrease is a free data retrieval call binding the contract method 0xcb1c4e67.
//
// Solidity: function minDelegationDecrease() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractCaller) MinDelegationDecrease(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SfcV1Contract.contract.Call(opts, &out, "minDelegationDecrease")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDelegationDecrease is a free data retrieval call binding the contract method 0xcb1c4e67.
//
// Solidity: function minDelegationDecrease() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractSession) MinDelegationDecrease() (*big.Int, error) {
	return _SfcV1Contract.Contract.MinDelegationDecrease(&_SfcV1Contract.CallOpts)
}

// MinDelegationDecrease is a free data retrieval call binding the contract method 0xcb1c4e67.
//
// Solidity: function minDelegationDecrease() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractCallerSession) MinDelegationDecrease() (*big.Int, error) {
	return _SfcV1Contract.Contract.MinDelegationDecrease(&_SfcV1Contract.CallOpts)
}

// MinDelegationIncrease is a free data retrieval call binding the contract method 0x60c7e37f.
//
// Solidity: function minDelegationIncrease() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractCaller) MinDelegationIncrease(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SfcV1Contract.contract.Call(opts, &out, "minDelegationIncrease")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDelegationIncrease is a free data retrieval call binding the contract method 0x60c7e37f.
//
// Solidity: function minDelegationIncrease() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractSession) MinDelegationIncrease() (*big.Int, error) {
	return _SfcV1Contract.Contract.MinDelegationIncrease(&_SfcV1Contract.CallOpts)
}

// MinDelegationIncrease is a free data retrieval call binding the contract method 0x60c7e37f.
//
// Solidity: function minDelegationIncrease() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractCallerSession) MinDelegationIncrease() (*big.Int, error) {
	return _SfcV1Contract.Contract.MinDelegationIncrease(&_SfcV1Contract.CallOpts)
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractCaller) MinStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SfcV1Contract.contract.Call(opts, &out, "minStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractSession) MinStake() (*big.Int, error) {
	return _SfcV1Contract.Contract.MinStake(&_SfcV1Contract.CallOpts)
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractCallerSession) MinStake() (*big.Int, error) {
	return _SfcV1Contract.Contract.MinStake(&_SfcV1Contract.CallOpts)
}

// MinStakeDecrease is a free data retrieval call binding the contract method 0x19ddb54f.
//
// Solidity: function minStakeDecrease() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractCaller) MinStakeDecrease(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SfcV1Contract.contract.Call(opts, &out, "minStakeDecrease")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinStakeDecrease is a free data retrieval call binding the contract method 0x19ddb54f.
//
// Solidity: function minStakeDecrease() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractSession) MinStakeDecrease() (*big.Int, error) {
	return _SfcV1Contract.Contract.MinStakeDecrease(&_SfcV1Contract.CallOpts)
}

// MinStakeDecrease is a free data retrieval call binding the contract method 0x19ddb54f.
//
// Solidity: function minStakeDecrease() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractCallerSession) MinStakeDecrease() (*big.Int, error) {
	return _SfcV1Contract.Contract.MinStakeDecrease(&_SfcV1Contract.CallOpts)
}

// MinStakeIncrease is a free data retrieval call binding the contract method 0xc4b5dd7e.
//
// Solidity: function minStakeIncrease() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractCaller) MinStakeIncrease(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SfcV1Contract.contract.Call(opts, &out, "minStakeIncrease")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinStakeIncrease is a free data retrieval call binding the contract method 0xc4b5dd7e.
//
// Solidity: function minStakeIncrease() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractSession) MinStakeIncrease() (*big.Int, error) {
	return _SfcV1Contract.Contract.MinStakeIncrease(&_SfcV1Contract.CallOpts)
}

// MinStakeIncrease is a free data retrieval call binding the contract method 0xc4b5dd7e.
//
// Solidity: function minStakeIncrease() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractCallerSession) MinStakeIncrease() (*big.Int, error) {
	return _SfcV1Contract.Contract.MinStakeIncrease(&_SfcV1Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SfcV1Contract *SfcV1ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SfcV1Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SfcV1Contract *SfcV1ContractSession) Owner() (common.Address, error) {
	return _SfcV1Contract.Contract.Owner(&_SfcV1Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SfcV1Contract *SfcV1ContractCallerSession) Owner() (common.Address, error) {
	return _SfcV1Contract.Contract.Owner(&_SfcV1Contract.CallOpts)
}

// RewardsAllowed is a free data retrieval call binding the contract method 0x8447c4df.
//
// Solidity: function rewardsAllowed() view returns(bool)
func (_SfcV1Contract *SfcV1ContractCaller) RewardsAllowed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SfcV1Contract.contract.Call(opts, &out, "rewardsAllowed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RewardsAllowed is a free data retrieval call binding the contract method 0x8447c4df.
//
// Solidity: function rewardsAllowed() view returns(bool)
func (_SfcV1Contract *SfcV1ContractSession) RewardsAllowed() (bool, error) {
	return _SfcV1Contract.Contract.RewardsAllowed(&_SfcV1Contract.CallOpts)
}

// RewardsAllowed is a free data retrieval call binding the contract method 0x8447c4df.
//
// Solidity: function rewardsAllowed() view returns(bool)
func (_SfcV1Contract *SfcV1ContractCallerSession) RewardsAllowed() (bool, error) {
	return _SfcV1Contract.Contract.RewardsAllowed(&_SfcV1Contract.CallOpts)
}

// RewardsStash is a free data retrieval call binding the contract method 0x6f498663.
//
// Solidity: function rewardsStash(address , uint256 ) view returns(uint256 amount)
func (_SfcV1Contract *SfcV1ContractCaller) RewardsStash(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SfcV1Contract.contract.Call(opts, &out, "rewardsStash", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardsStash is a free data retrieval call binding the contract method 0x6f498663.
//
// Solidity: function rewardsStash(address , uint256 ) view returns(uint256 amount)
func (_SfcV1Contract *SfcV1ContractSession) RewardsStash(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _SfcV1Contract.Contract.RewardsStash(&_SfcV1Contract.CallOpts, arg0, arg1)
}

// RewardsStash is a free data retrieval call binding the contract method 0x6f498663.
//
// Solidity: function rewardsStash(address , uint256 ) view returns(uint256 amount)
func (_SfcV1Contract *SfcV1ContractCallerSession) RewardsStash(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _SfcV1Contract.Contract.RewardsStash(&_SfcV1Contract.CallOpts, arg0, arg1)
}

// SlashedDelegationsTotalAmount is a free data retrieval call binding the contract method 0xa70da4d2.
//
// Solidity: function slashedDelegationsTotalAmount() view returns(uint256)
func (_SfcV1Contract *SfcV1ContractCaller) SlashedDelegationsTotalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SfcV1Contract.contract.Call(opts, &out, "slashedDelegationsTotalAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlashedDelegationsTotalAmount is a free data retrieval call binding the contract method 0xa70da4d2.
//
// Solidity: function slashedDelegationsTotalAmount() view returns(uint256)
func (_SfcV1Contract *SfcV1ContractSession) SlashedDelegationsTotalAmount() (*big.Int, error) {
	return _SfcV1Contract.Contract.SlashedDelegationsTotalAmount(&_SfcV1Contract.CallOpts)
}

// SlashedDelegationsTotalAmount is a free data retrieval call binding the contract method 0xa70da4d2.
//
// Solidity: function slashedDelegationsTotalAmount() view returns(uint256)
func (_SfcV1Contract *SfcV1ContractCallerSession) SlashedDelegationsTotalAmount() (*big.Int, error) {
	return _SfcV1Contract.Contract.SlashedDelegationsTotalAmount(&_SfcV1Contract.CallOpts)
}

// SlashedStakeTotalAmount is a free data retrieval call binding the contract method 0x0a29180c.
//
// Solidity: function slashedStakeTotalAmount() view returns(uint256)
func (_SfcV1Contract *SfcV1ContractCaller) SlashedStakeTotalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SfcV1Contract.contract.Call(opts, &out, "slashedStakeTotalAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlashedStakeTotalAmount is a free data retrieval call binding the contract method 0x0a29180c.
//
// Solidity: function slashedStakeTotalAmount() view returns(uint256)
func (_SfcV1Contract *SfcV1ContractSession) SlashedStakeTotalAmount() (*big.Int, error) {
	return _SfcV1Contract.Contract.SlashedStakeTotalAmount(&_SfcV1Contract.CallOpts)
}

// SlashedStakeTotalAmount is a free data retrieval call binding the contract method 0x0a29180c.
//
// Solidity: function slashedStakeTotalAmount() view returns(uint256)
func (_SfcV1Contract *SfcV1ContractCallerSession) SlashedStakeTotalAmount() (*big.Int, error) {
	return _SfcV1Contract.Contract.SlashedStakeTotalAmount(&_SfcV1Contract.CallOpts)
}

// StakeLockPeriodEpochs is a free data retrieval call binding the contract method 0x54d77ed2.
//
// Solidity: function stakeLockPeriodEpochs() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractCaller) StakeLockPeriodEpochs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SfcV1Contract.contract.Call(opts, &out, "stakeLockPeriodEpochs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeLockPeriodEpochs is a free data retrieval call binding the contract method 0x54d77ed2.
//
// Solidity: function stakeLockPeriodEpochs() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractSession) StakeLockPeriodEpochs() (*big.Int, error) {
	return _SfcV1Contract.Contract.StakeLockPeriodEpochs(&_SfcV1Contract.CallOpts)
}

// StakeLockPeriodEpochs is a free data retrieval call binding the contract method 0x54d77ed2.
//
// Solidity: function stakeLockPeriodEpochs() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractCallerSession) StakeLockPeriodEpochs() (*big.Int, error) {
	return _SfcV1Contract.Contract.StakeLockPeriodEpochs(&_SfcV1Contract.CallOpts)
}

// StakeLockPeriodTime is a free data retrieval call binding the contract method 0x3fee10a8.
//
// Solidity: function stakeLockPeriodTime() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractCaller) StakeLockPeriodTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SfcV1Contract.contract.Call(opts, &out, "stakeLockPeriodTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeLockPeriodTime is a free data retrieval call binding the contract method 0x3fee10a8.
//
// Solidity: function stakeLockPeriodTime() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractSession) StakeLockPeriodTime() (*big.Int, error) {
	return _SfcV1Contract.Contract.StakeLockPeriodTime(&_SfcV1Contract.CallOpts)
}

// StakeLockPeriodTime is a free data retrieval call binding the contract method 0x3fee10a8.
//
// Solidity: function stakeLockPeriodTime() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractCallerSession) StakeLockPeriodTime() (*big.Int, error) {
	return _SfcV1Contract.Contract.StakeLockPeriodTime(&_SfcV1Contract.CallOpts)
}

// StakeTotalAmount is a free data retrieval call binding the contract method 0x3d0317fe.
//
// Solidity: function stakeTotalAmount() view returns(uint256)
func (_SfcV1Contract *SfcV1ContractCaller) StakeTotalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SfcV1Contract.contract.Call(opts, &out, "stakeTotalAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeTotalAmount is a free data retrieval call binding the contract method 0x3d0317fe.
//
// Solidity: function stakeTotalAmount() view returns(uint256)
func (_SfcV1Contract *SfcV1ContractSession) StakeTotalAmount() (*big.Int, error) {
	return _SfcV1Contract.Contract.StakeTotalAmount(&_SfcV1Contract.CallOpts)
}

// StakeTotalAmount is a free data retrieval call binding the contract method 0x3d0317fe.
//
// Solidity: function stakeTotalAmount() view returns(uint256)
func (_SfcV1Contract *SfcV1ContractCallerSession) StakeTotalAmount() (*big.Int, error) {
	return _SfcV1Contract.Contract.StakeTotalAmount(&_SfcV1Contract.CallOpts)
}

// StakerMetadata is a free data retrieval call binding the contract method 0x98ec2de5.
//
// Solidity: function stakerMetadata(uint256 ) view returns(bytes)
func (_SfcV1Contract *SfcV1ContractCaller) StakerMetadata(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _SfcV1Contract.contract.Call(opts, &out, "stakerMetadata", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// StakerMetadata is a free data retrieval call binding the contract method 0x98ec2de5.
//
// Solidity: function stakerMetadata(uint256 ) view returns(bytes)
func (_SfcV1Contract *SfcV1ContractSession) StakerMetadata(arg0 *big.Int) ([]byte, error) {
	return _SfcV1Contract.Contract.StakerMetadata(&_SfcV1Contract.CallOpts, arg0)
}

// StakerMetadata is a free data retrieval call binding the contract method 0x98ec2de5.
//
// Solidity: function stakerMetadata(uint256 ) view returns(bytes)
func (_SfcV1Contract *SfcV1ContractCallerSession) StakerMetadata(arg0 *big.Int) ([]byte, error) {
	return _SfcV1Contract.Contract.StakerMetadata(&_SfcV1Contract.CallOpts, arg0)
}

// Stakers is a free data retrieval call binding the contract method 0xfd5e6dd1.
//
// Solidity: function stakers(uint256 ) view returns(uint256 status, uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime, uint256 stakeAmount, uint256 paidUntilEpoch, uint256 delegatedMe, address dagAddress, address sfcAddress)
func (_SfcV1Contract *SfcV1ContractCaller) Stakers(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Status           *big.Int
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	DeactivatedEpoch *big.Int
	DeactivatedTime  *big.Int
	StakeAmount      *big.Int
	PaidUntilEpoch   *big.Int
	DelegatedMe      *big.Int
	DagAddress       common.Address
	SfcAddress       common.Address
}, error) {
	var out []interface{}
	err := _SfcV1Contract.contract.Call(opts, &out, "stakers", arg0)

	outstruct := new(struct {
		Status           *big.Int
		CreatedEpoch     *big.Int
		CreatedTime      *big.Int
		DeactivatedEpoch *big.Int
		DeactivatedTime  *big.Int
		StakeAmount      *big.Int
		PaidUntilEpoch   *big.Int
		DelegatedMe      *big.Int
		DagAddress       common.Address
		SfcAddress       common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CreatedEpoch = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.CreatedTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.DeactivatedEpoch = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.DeactivatedTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.StakeAmount = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.PaidUntilEpoch = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.DelegatedMe = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.DagAddress = *abi.ConvertType(out[8], new(common.Address)).(*common.Address)
	outstruct.SfcAddress = *abi.ConvertType(out[9], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Stakers is a free data retrieval call binding the contract method 0xfd5e6dd1.
//
// Solidity: function stakers(uint256 ) view returns(uint256 status, uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime, uint256 stakeAmount, uint256 paidUntilEpoch, uint256 delegatedMe, address dagAddress, address sfcAddress)
func (_SfcV1Contract *SfcV1ContractSession) Stakers(arg0 *big.Int) (struct {
	Status           *big.Int
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	DeactivatedEpoch *big.Int
	DeactivatedTime  *big.Int
	StakeAmount      *big.Int
	PaidUntilEpoch   *big.Int
	DelegatedMe      *big.Int
	DagAddress       common.Address
	SfcAddress       common.Address
}, error) {
	return _SfcV1Contract.Contract.Stakers(&_SfcV1Contract.CallOpts, arg0)
}

// Stakers is a free data retrieval call binding the contract method 0xfd5e6dd1.
//
// Solidity: function stakers(uint256 ) view returns(uint256 status, uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime, uint256 stakeAmount, uint256 paidUntilEpoch, uint256 delegatedMe, address dagAddress, address sfcAddress)
func (_SfcV1Contract *SfcV1ContractCallerSession) Stakers(arg0 *big.Int) (struct {
	Status           *big.Int
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	DeactivatedEpoch *big.Int
	DeactivatedTime  *big.Int
	StakeAmount      *big.Int
	PaidUntilEpoch   *big.Int
	DelegatedMe      *big.Int
	DagAddress       common.Address
	SfcAddress       common.Address
}, error) {
	return _SfcV1Contract.Contract.Stakers(&_SfcV1Contract.CallOpts, arg0)
}

// StakersLastID is a free data retrieval call binding the contract method 0x81d9dc7a.
//
// Solidity: function stakersLastID() view returns(uint256)
func (_SfcV1Contract *SfcV1ContractCaller) StakersLastID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SfcV1Contract.contract.Call(opts, &out, "stakersLastID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakersLastID is a free data retrieval call binding the contract method 0x81d9dc7a.
//
// Solidity: function stakersLastID() view returns(uint256)
func (_SfcV1Contract *SfcV1ContractSession) StakersLastID() (*big.Int, error) {
	return _SfcV1Contract.Contract.StakersLastID(&_SfcV1Contract.CallOpts)
}

// StakersLastID is a free data retrieval call binding the contract method 0x81d9dc7a.
//
// Solidity: function stakersLastID() view returns(uint256)
func (_SfcV1Contract *SfcV1ContractCallerSession) StakersLastID() (*big.Int, error) {
	return _SfcV1Contract.Contract.StakersLastID(&_SfcV1Contract.CallOpts)
}

// StakersNum is a free data retrieval call binding the contract method 0x08728f6e.
//
// Solidity: function stakersNum() view returns(uint256)
func (_SfcV1Contract *SfcV1ContractCaller) StakersNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SfcV1Contract.contract.Call(opts, &out, "stakersNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakersNum is a free data retrieval call binding the contract method 0x08728f6e.
//
// Solidity: function stakersNum() view returns(uint256)
func (_SfcV1Contract *SfcV1ContractSession) StakersNum() (*big.Int, error) {
	return _SfcV1Contract.Contract.StakersNum(&_SfcV1Contract.CallOpts)
}

// StakersNum is a free data retrieval call binding the contract method 0x08728f6e.
//
// Solidity: function stakersNum() view returns(uint256)
func (_SfcV1Contract *SfcV1ContractCallerSession) StakersNum() (*big.Int, error) {
	return _SfcV1Contract.Contract.StakersNum(&_SfcV1Contract.CallOpts)
}

// UnbondingStartDate is a free data retrieval call binding the contract method 0x53a72586.
//
// Solidity: function unbondingStartDate() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractCaller) UnbondingStartDate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SfcV1Contract.contract.Call(opts, &out, "unbondingStartDate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnbondingStartDate is a free data retrieval call binding the contract method 0x53a72586.
//
// Solidity: function unbondingStartDate() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractSession) UnbondingStartDate() (*big.Int, error) {
	return _SfcV1Contract.Contract.UnbondingStartDate(&_SfcV1Contract.CallOpts)
}

// UnbondingStartDate is a free data retrieval call binding the contract method 0x53a72586.
//
// Solidity: function unbondingStartDate() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractCallerSession) UnbondingStartDate() (*big.Int, error) {
	return _SfcV1Contract.Contract.UnbondingStartDate(&_SfcV1Contract.CallOpts)
}

// UnbondingUnlockPeriod is a free data retrieval call binding the contract method 0x3a0af4d4.
//
// Solidity: function unbondingUnlockPeriod() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractCaller) UnbondingUnlockPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SfcV1Contract.contract.Call(opts, &out, "unbondingUnlockPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnbondingUnlockPeriod is a free data retrieval call binding the contract method 0x3a0af4d4.
//
// Solidity: function unbondingUnlockPeriod() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractSession) UnbondingUnlockPeriod() (*big.Int, error) {
	return _SfcV1Contract.Contract.UnbondingUnlockPeriod(&_SfcV1Contract.CallOpts)
}

// UnbondingUnlockPeriod is a free data retrieval call binding the contract method 0x3a0af4d4.
//
// Solidity: function unbondingUnlockPeriod() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractCallerSession) UnbondingUnlockPeriod() (*big.Int, error) {
	return _SfcV1Contract.Contract.UnbondingUnlockPeriod(&_SfcV1Contract.CallOpts)
}

// ValidatorCommission is a free data retrieval call binding the contract method 0xa7786515.
//
// Solidity: function validatorCommission() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractCaller) ValidatorCommission(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SfcV1Contract.contract.Call(opts, &out, "validatorCommission")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorCommission is a free data retrieval call binding the contract method 0xa7786515.
//
// Solidity: function validatorCommission() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractSession) ValidatorCommission() (*big.Int, error) {
	return _SfcV1Contract.Contract.ValidatorCommission(&_SfcV1Contract.CallOpts)
}

// ValidatorCommission is a free data retrieval call binding the contract method 0xa7786515.
//
// Solidity: function validatorCommission() pure returns(uint256)
func (_SfcV1Contract *SfcV1ContractCallerSession) ValidatorCommission() (*big.Int, error) {
	return _SfcV1Contract.Contract.ValidatorCommission(&_SfcV1Contract.CallOpts)
}

// WithdrawalRequests is a free data retrieval call binding the contract method 0x4e5a2328.
//
// Solidity: function withdrawalRequests(address , uint256 ) view returns(uint256 stakerID, uint256 epoch, uint256 time, uint256 amount, bool delegation)
func (_SfcV1Contract *SfcV1ContractCaller) WithdrawalRequests(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	StakerID   *big.Int
	Epoch      *big.Int
	Time       *big.Int
	Amount     *big.Int
	Delegation bool
}, error) {
	var out []interface{}
	err := _SfcV1Contract.contract.Call(opts, &out, "withdrawalRequests", arg0, arg1)

	outstruct := new(struct {
		StakerID   *big.Int
		Epoch      *big.Int
		Time       *big.Int
		Amount     *big.Int
		Delegation bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StakerID = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Epoch = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Time = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Delegation = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// WithdrawalRequests is a free data retrieval call binding the contract method 0x4e5a2328.
//
// Solidity: function withdrawalRequests(address , uint256 ) view returns(uint256 stakerID, uint256 epoch, uint256 time, uint256 amount, bool delegation)
func (_SfcV1Contract *SfcV1ContractSession) WithdrawalRequests(arg0 common.Address, arg1 *big.Int) (struct {
	StakerID   *big.Int
	Epoch      *big.Int
	Time       *big.Int
	Amount     *big.Int
	Delegation bool
}, error) {
	return _SfcV1Contract.Contract.WithdrawalRequests(&_SfcV1Contract.CallOpts, arg0, arg1)
}

// WithdrawalRequests is a free data retrieval call binding the contract method 0x4e5a2328.
//
// Solidity: function withdrawalRequests(address , uint256 ) view returns(uint256 stakerID, uint256 epoch, uint256 time, uint256 amount, bool delegation)
func (_SfcV1Contract *SfcV1ContractCallerSession) WithdrawalRequests(arg0 common.Address, arg1 *big.Int) (struct {
	StakerID   *big.Int
	Epoch      *big.Int
	Time       *big.Int
	Amount     *big.Int
	Delegation bool
}, error) {
	return _SfcV1Contract.Contract.WithdrawalRequests(&_SfcV1Contract.CallOpts, arg0, arg1)
}

// SyncDelegator is a paid mutator transaction binding the contract method 0x424219a4.
//
// Solidity: function _syncDelegator(address delegator) returns()
func (_SfcV1Contract *SfcV1ContractTransactor) SyncDelegator(opts *bind.TransactOpts, delegator common.Address) (*types.Transaction, error) {
	return _SfcV1Contract.contract.Transact(opts, "_syncDelegator", delegator)
}

// SyncDelegator is a paid mutator transaction binding the contract method 0x424219a4.
//
// Solidity: function _syncDelegator(address delegator) returns()
func (_SfcV1Contract *SfcV1ContractSession) SyncDelegator(delegator common.Address) (*types.Transaction, error) {
	return _SfcV1Contract.Contract.SyncDelegator(&_SfcV1Contract.TransactOpts, delegator)
}

// SyncDelegator is a paid mutator transaction binding the contract method 0x424219a4.
//
// Solidity: function _syncDelegator(address delegator) returns()
func (_SfcV1Contract *SfcV1ContractTransactorSession) SyncDelegator(delegator common.Address) (*types.Transaction, error) {
	return _SfcV1Contract.Contract.SyncDelegator(&_SfcV1Contract.TransactOpts, delegator)
}

// SyncStaker is a paid mutator transaction binding the contract method 0xeac3baf2.
//
// Solidity: function _syncStaker(uint256 stakerID) returns()
func (_SfcV1Contract *SfcV1ContractTransactor) SyncStaker(opts *bind.TransactOpts, stakerID *big.Int) (*types.Transaction, error) {
	return _SfcV1Contract.contract.Transact(opts, "_syncStaker", stakerID)
}

// SyncStaker is a paid mutator transaction binding the contract method 0xeac3baf2.
//
// Solidity: function _syncStaker(uint256 stakerID) returns()
func (_SfcV1Contract *SfcV1ContractSession) SyncStaker(stakerID *big.Int) (*types.Transaction, error) {
	return _SfcV1Contract.Contract.SyncStaker(&_SfcV1Contract.TransactOpts, stakerID)
}

// SyncStaker is a paid mutator transaction binding the contract method 0xeac3baf2.
//
// Solidity: function _syncStaker(uint256 stakerID) returns()
func (_SfcV1Contract *SfcV1ContractTransactorSession) SyncStaker(stakerID *big.Int) (*types.Transaction, error) {
	return _SfcV1Contract.Contract.SyncStaker(&_SfcV1Contract.TransactOpts, stakerID)
}

// UpgradeStakerStorage is a paid mutator transaction binding the contract method 0x28dca8ff.
//
// Solidity: function _upgradeStakerStorage(uint256 stakerID) returns()
func (_SfcV1Contract *SfcV1ContractTransactor) UpgradeStakerStorage(opts *bind.TransactOpts, stakerID *big.Int) (*types.Transaction, error) {
	return _SfcV1Contract.contract.Transact(opts, "_upgradeStakerStorage", stakerID)
}

// UpgradeStakerStorage is a paid mutator transaction binding the contract method 0x28dca8ff.
//
// Solidity: function _upgradeStakerStorage(uint256 stakerID) returns()
func (_SfcV1Contract *SfcV1ContractSession) UpgradeStakerStorage(stakerID *big.Int) (*types.Transaction, error) {
	return _SfcV1Contract.Contract.UpgradeStakerStorage(&_SfcV1Contract.TransactOpts, stakerID)
}

// UpgradeStakerStorage is a paid mutator transaction binding the contract method 0x28dca8ff.
//
// Solidity: function _upgradeStakerStorage(uint256 stakerID) returns()
func (_SfcV1Contract *SfcV1ContractTransactorSession) UpgradeStakerStorage(stakerID *big.Int) (*types.Transaction, error) {
	return _SfcV1Contract.Contract.UpgradeStakerStorage(&_SfcV1Contract.TransactOpts, stakerID)
}

// ClaimDelegationRewards is a paid mutator transaction binding the contract method 0x793c45ce.
//
// Solidity: function claimDelegationRewards(uint256 maxEpochs) returns()
func (_SfcV1Contract *SfcV1ContractTransactor) ClaimDelegationRewards(opts *bind.TransactOpts, maxEpochs *big.Int) (*types.Transaction, error) {
	return _SfcV1Contract.contract.Transact(opts, "claimDelegationRewards", maxEpochs)
}

// ClaimDelegationRewards is a paid mutator transaction binding the contract method 0x793c45ce.
//
// Solidity: function claimDelegationRewards(uint256 maxEpochs) returns()
func (_SfcV1Contract *SfcV1ContractSession) ClaimDelegationRewards(maxEpochs *big.Int) (*types.Transaction, error) {
	return _SfcV1Contract.Contract.ClaimDelegationRewards(&_SfcV1Contract.TransactOpts, maxEpochs)
}

// ClaimDelegationRewards is a paid mutator transaction binding the contract method 0x793c45ce.
//
// Solidity: function claimDelegationRewards(uint256 maxEpochs) returns()
func (_SfcV1Contract *SfcV1ContractTransactorSession) ClaimDelegationRewards(maxEpochs *big.Int) (*types.Transaction, error) {
	return _SfcV1Contract.Contract.ClaimDelegationRewards(&_SfcV1Contract.TransactOpts, maxEpochs)
}

// ClaimValidatorRewards is a paid mutator transaction binding the contract method 0x295cccba.
//
// Solidity: function claimValidatorRewards(uint256 maxEpochs) returns()
func (_SfcV1Contract *SfcV1ContractTransactor) ClaimValidatorRewards(opts *bind.TransactOpts, maxEpochs *big.Int) (*types.Transaction, error) {
	return _SfcV1Contract.contract.Transact(opts, "claimValidatorRewards", maxEpochs)
}

// ClaimValidatorRewards is a paid mutator transaction binding the contract method 0x295cccba.
//
// Solidity: function claimValidatorRewards(uint256 maxEpochs) returns()
func (_SfcV1Contract *SfcV1ContractSession) ClaimValidatorRewards(maxEpochs *big.Int) (*types.Transaction, error) {
	return _SfcV1Contract.Contract.ClaimValidatorRewards(&_SfcV1Contract.TransactOpts, maxEpochs)
}

// ClaimValidatorRewards is a paid mutator transaction binding the contract method 0x295cccba.
//
// Solidity: function claimValidatorRewards(uint256 maxEpochs) returns()
func (_SfcV1Contract *SfcV1ContractTransactorSession) ClaimValidatorRewards(maxEpochs *big.Int) (*types.Transaction, error) {
	return _SfcV1Contract.Contract.ClaimValidatorRewards(&_SfcV1Contract.TransactOpts, maxEpochs)
}

// CreateDelegation is a paid mutator transaction binding the contract method 0xc312eb07.
//
// Solidity: function createDelegation(uint256 to) payable returns()
func (_SfcV1Contract *SfcV1ContractTransactor) CreateDelegation(opts *bind.TransactOpts, to *big.Int) (*types.Transaction, error) {
	return _SfcV1Contract.contract.Transact(opts, "createDelegation", to)
}

// CreateDelegation is a paid mutator transaction binding the contract method 0xc312eb07.
//
// Solidity: function createDelegation(uint256 to) payable returns()
func (_SfcV1Contract *SfcV1ContractSession) CreateDelegation(to *big.Int) (*types.Transaction, error) {
	return _SfcV1Contract.Contract.CreateDelegation(&_SfcV1Contract.TransactOpts, to)
}

// CreateDelegation is a paid mutator transaction binding the contract method 0xc312eb07.
//
// Solidity: function createDelegation(uint256 to) payable returns()
func (_SfcV1Contract *SfcV1ContractTransactorSession) CreateDelegation(to *big.Int) (*types.Transaction, error) {
	return _SfcV1Contract.Contract.CreateDelegation(&_SfcV1Contract.TransactOpts, to)
}

// CreateStake is a paid mutator transaction binding the contract method 0xcc8c2120.
//
// Solidity: function createStake(bytes metadata) payable returns()
func (_SfcV1Contract *SfcV1ContractTransactor) CreateStake(opts *bind.TransactOpts, metadata []byte) (*types.Transaction, error) {
	return _SfcV1Contract.contract.Transact(opts, "createStake", metadata)
}

// CreateStake is a paid mutator transaction binding the contract method 0xcc8c2120.
//
// Solidity: function createStake(bytes metadata) payable returns()
func (_SfcV1Contract *SfcV1ContractSession) CreateStake(metadata []byte) (*types.Transaction, error) {
	return _SfcV1Contract.Contract.CreateStake(&_SfcV1Contract.TransactOpts, metadata)
}

// CreateStake is a paid mutator transaction binding the contract method 0xcc8c2120.
//
// Solidity: function createStake(bytes metadata) payable returns()
func (_SfcV1Contract *SfcV1ContractTransactorSession) CreateStake(metadata []byte) (*types.Transaction, error) {
	return _SfcV1Contract.Contract.CreateStake(&_SfcV1Contract.TransactOpts, metadata)
}

// CreateStakeWithAddresses is a paid mutator transaction binding the contract method 0x90475ae4.
//
// Solidity: function createStakeWithAddresses(address dagAdrress, address sfcAddress, bytes metadata) payable returns()
func (_SfcV1Contract *SfcV1ContractTransactor) CreateStakeWithAddresses(opts *bind.TransactOpts, dagAdrress common.Address, sfcAddress common.Address, metadata []byte) (*types.Transaction, error) {
	return _SfcV1Contract.contract.Transact(opts, "createStakeWithAddresses", dagAdrress, sfcAddress, metadata)
}

// CreateStakeWithAddresses is a paid mutator transaction binding the contract method 0x90475ae4.
//
// Solidity: function createStakeWithAddresses(address dagAdrress, address sfcAddress, bytes metadata) payable returns()
func (_SfcV1Contract *SfcV1ContractSession) CreateStakeWithAddresses(dagAdrress common.Address, sfcAddress common.Address, metadata []byte) (*types.Transaction, error) {
	return _SfcV1Contract.Contract.CreateStakeWithAddresses(&_SfcV1Contract.TransactOpts, dagAdrress, sfcAddress, metadata)
}

// CreateStakeWithAddresses is a paid mutator transaction binding the contract method 0x90475ae4.
//
// Solidity: function createStakeWithAddresses(address dagAdrress, address sfcAddress, bytes metadata) payable returns()
func (_SfcV1Contract *SfcV1ContractTransactorSession) CreateStakeWithAddresses(dagAdrress common.Address, sfcAddress common.Address, metadata []byte) (*types.Transaction, error) {
	return _SfcV1Contract.Contract.CreateStakeWithAddresses(&_SfcV1Contract.TransactOpts, dagAdrress, sfcAddress, metadata)
}

// IncreaseDelegation is a paid mutator transaction binding the contract method 0x3a274ff6.
//
// Solidity: function increaseDelegation() payable returns()
func (_SfcV1Contract *SfcV1ContractTransactor) IncreaseDelegation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SfcV1Contract.contract.Transact(opts, "increaseDelegation")
}

// IncreaseDelegation is a paid mutator transaction binding the contract method 0x3a274ff6.
//
// Solidity: function increaseDelegation() payable returns()
func (_SfcV1Contract *SfcV1ContractSession) IncreaseDelegation() (*types.Transaction, error) {
	return _SfcV1Contract.Contract.IncreaseDelegation(&_SfcV1Contract.TransactOpts)
}

// IncreaseDelegation is a paid mutator transaction binding the contract method 0x3a274ff6.
//
// Solidity: function increaseDelegation() payable returns()
func (_SfcV1Contract *SfcV1ContractTransactorSession) IncreaseDelegation() (*types.Transaction, error) {
	return _SfcV1Contract.Contract.IncreaseDelegation(&_SfcV1Contract.TransactOpts)
}

// IncreaseStake is a paid mutator transaction binding the contract method 0xd9e257ef.
//
// Solidity: function increaseStake() payable returns()
func (_SfcV1Contract *SfcV1ContractTransactor) IncreaseStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SfcV1Contract.contract.Transact(opts, "increaseStake")
}

// IncreaseStake is a paid mutator transaction binding the contract method 0xd9e257ef.
//
// Solidity: function increaseStake() payable returns()
func (_SfcV1Contract *SfcV1ContractSession) IncreaseStake() (*types.Transaction, error) {
	return _SfcV1Contract.Contract.IncreaseStake(&_SfcV1Contract.TransactOpts)
}

// IncreaseStake is a paid mutator transaction binding the contract method 0xd9e257ef.
//
// Solidity: function increaseStake() payable returns()
func (_SfcV1Contract *SfcV1ContractTransactorSession) IncreaseStake() (*types.Transaction, error) {
	return _SfcV1Contract.Contract.IncreaseStake(&_SfcV1Contract.TransactOpts)
}

// PartialWithdrawByRequest is a paid mutator transaction binding the contract method 0xf8b18d8a.
//
// Solidity: function partialWithdrawByRequest(uint256 wrID) returns()
func (_SfcV1Contract *SfcV1ContractTransactor) PartialWithdrawByRequest(opts *bind.TransactOpts, wrID *big.Int) (*types.Transaction, error) {
	return _SfcV1Contract.contract.Transact(opts, "partialWithdrawByRequest", wrID)
}

// PartialWithdrawByRequest is a paid mutator transaction binding the contract method 0xf8b18d8a.
//
// Solidity: function partialWithdrawByRequest(uint256 wrID) returns()
func (_SfcV1Contract *SfcV1ContractSession) PartialWithdrawByRequest(wrID *big.Int) (*types.Transaction, error) {
	return _SfcV1Contract.Contract.PartialWithdrawByRequest(&_SfcV1Contract.TransactOpts, wrID)
}

// PartialWithdrawByRequest is a paid mutator transaction binding the contract method 0xf8b18d8a.
//
// Solidity: function partialWithdrawByRequest(uint256 wrID) returns()
func (_SfcV1Contract *SfcV1ContractTransactorSession) PartialWithdrawByRequest(wrID *big.Int) (*types.Transaction, error) {
	return _SfcV1Contract.Contract.PartialWithdrawByRequest(&_SfcV1Contract.TransactOpts, wrID)
}

// PrepareToWithdrawDelegation is a paid mutator transaction binding the contract method 0x1c333318.
//
// Solidity: function prepareToWithdrawDelegation() returns()
func (_SfcV1Contract *SfcV1ContractTransactor) PrepareToWithdrawDelegation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SfcV1Contract.contract.Transact(opts, "prepareToWithdrawDelegation")
}

// PrepareToWithdrawDelegation is a paid mutator transaction binding the contract method 0x1c333318.
//
// Solidity: function prepareToWithdrawDelegation() returns()
func (_SfcV1Contract *SfcV1ContractSession) PrepareToWithdrawDelegation() (*types.Transaction, error) {
	return _SfcV1Contract.Contract.PrepareToWithdrawDelegation(&_SfcV1Contract.TransactOpts)
}

// PrepareToWithdrawDelegation is a paid mutator transaction binding the contract method 0x1c333318.
//
// Solidity: function prepareToWithdrawDelegation() returns()
func (_SfcV1Contract *SfcV1ContractTransactorSession) PrepareToWithdrawDelegation() (*types.Transaction, error) {
	return _SfcV1Contract.Contract.PrepareToWithdrawDelegation(&_SfcV1Contract.TransactOpts)
}

// PrepareToWithdrawDelegationPartial is a paid mutator transaction binding the contract method 0xe7ff9e78.
//
// Solidity: function prepareToWithdrawDelegationPartial(uint256 wrID, uint256 amount) returns()
func (_SfcV1Contract *SfcV1ContractTransactor) PrepareToWithdrawDelegationPartial(opts *bind.TransactOpts, wrID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _SfcV1Contract.contract.Transact(opts, "prepareToWithdrawDelegationPartial", wrID, amount)
}

// PrepareToWithdrawDelegationPartial is a paid mutator transaction binding the contract method 0xe7ff9e78.
//
// Solidity: function prepareToWithdrawDelegationPartial(uint256 wrID, uint256 amount) returns()
func (_SfcV1Contract *SfcV1ContractSession) PrepareToWithdrawDelegationPartial(wrID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _SfcV1Contract.Contract.PrepareToWithdrawDelegationPartial(&_SfcV1Contract.TransactOpts, wrID, amount)
}

// PrepareToWithdrawDelegationPartial is a paid mutator transaction binding the contract method 0xe7ff9e78.
//
// Solidity: function prepareToWithdrawDelegationPartial(uint256 wrID, uint256 amount) returns()
func (_SfcV1Contract *SfcV1ContractTransactorSession) PrepareToWithdrawDelegationPartial(wrID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _SfcV1Contract.Contract.PrepareToWithdrawDelegationPartial(&_SfcV1Contract.TransactOpts, wrID, amount)
}

// PrepareToWithdrawStake is a paid mutator transaction binding the contract method 0xc41b6405.
//
// Solidity: function prepareToWithdrawStake() returns()
func (_SfcV1Contract *SfcV1ContractTransactor) PrepareToWithdrawStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SfcV1Contract.contract.Transact(opts, "prepareToWithdrawStake")
}

// PrepareToWithdrawStake is a paid mutator transaction binding the contract method 0xc41b6405.
//
// Solidity: function prepareToWithdrawStake() returns()
func (_SfcV1Contract *SfcV1ContractSession) PrepareToWithdrawStake() (*types.Transaction, error) {
	return _SfcV1Contract.Contract.PrepareToWithdrawStake(&_SfcV1Contract.TransactOpts)
}

// PrepareToWithdrawStake is a paid mutator transaction binding the contract method 0xc41b6405.
//
// Solidity: function prepareToWithdrawStake() returns()
func (_SfcV1Contract *SfcV1ContractTransactorSession) PrepareToWithdrawStake() (*types.Transaction, error) {
	return _SfcV1Contract.Contract.PrepareToWithdrawStake(&_SfcV1Contract.TransactOpts)
}

// PrepareToWithdrawStakePartial is a paid mutator transaction binding the contract method 0x26682c71.
//
// Solidity: function prepareToWithdrawStakePartial(uint256 wrID, uint256 amount) returns()
func (_SfcV1Contract *SfcV1ContractTransactor) PrepareToWithdrawStakePartial(opts *bind.TransactOpts, wrID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _SfcV1Contract.contract.Transact(opts, "prepareToWithdrawStakePartial", wrID, amount)
}

// PrepareToWithdrawStakePartial is a paid mutator transaction binding the contract method 0x26682c71.
//
// Solidity: function prepareToWithdrawStakePartial(uint256 wrID, uint256 amount) returns()
func (_SfcV1Contract *SfcV1ContractSession) PrepareToWithdrawStakePartial(wrID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _SfcV1Contract.Contract.PrepareToWithdrawStakePartial(&_SfcV1Contract.TransactOpts, wrID, amount)
}

// PrepareToWithdrawStakePartial is a paid mutator transaction binding the contract method 0x26682c71.
//
// Solidity: function prepareToWithdrawStakePartial(uint256 wrID, uint256 amount) returns()
func (_SfcV1Contract *SfcV1ContractTransactorSession) PrepareToWithdrawStakePartial(wrID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _SfcV1Contract.Contract.PrepareToWithdrawStakePartial(&_SfcV1Contract.TransactOpts, wrID, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SfcV1Contract *SfcV1ContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SfcV1Contract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SfcV1Contract *SfcV1ContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _SfcV1Contract.Contract.RenounceOwnership(&_SfcV1Contract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SfcV1Contract *SfcV1ContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SfcV1Contract.Contract.RenounceOwnership(&_SfcV1Contract.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SfcV1Contract *SfcV1ContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SfcV1Contract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SfcV1Contract *SfcV1ContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SfcV1Contract.Contract.TransferOwnership(&_SfcV1Contract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SfcV1Contract *SfcV1ContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SfcV1Contract.Contract.TransferOwnership(&_SfcV1Contract.TransactOpts, newOwner)
}

// UnstashRewards is a paid mutator transaction binding the contract method 0x876f7e2a.
//
// Solidity: function unstashRewards() returns()
func (_SfcV1Contract *SfcV1ContractTransactor) UnstashRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SfcV1Contract.contract.Transact(opts, "unstashRewards")
}

// UnstashRewards is a paid mutator transaction binding the contract method 0x876f7e2a.
//
// Solidity: function unstashRewards() returns()
func (_SfcV1Contract *SfcV1ContractSession) UnstashRewards() (*types.Transaction, error) {
	return _SfcV1Contract.Contract.UnstashRewards(&_SfcV1Contract.TransactOpts)
}

// UnstashRewards is a paid mutator transaction binding the contract method 0x876f7e2a.
//
// Solidity: function unstashRewards() returns()
func (_SfcV1Contract *SfcV1ContractTransactorSession) UnstashRewards() (*types.Transaction, error) {
	return _SfcV1Contract.Contract.UnstashRewards(&_SfcV1Contract.TransactOpts)
}

// UpdateBaseRewardPerSec is a paid mutator transaction binding the contract method 0x1b593d8a.
//
// Solidity: function updateBaseRewardPerSec(uint256 value) returns()
func (_SfcV1Contract *SfcV1ContractTransactor) UpdateBaseRewardPerSec(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _SfcV1Contract.contract.Transact(opts, "updateBaseRewardPerSec", value)
}

// UpdateBaseRewardPerSec is a paid mutator transaction binding the contract method 0x1b593d8a.
//
// Solidity: function updateBaseRewardPerSec(uint256 value) returns()
func (_SfcV1Contract *SfcV1ContractSession) UpdateBaseRewardPerSec(value *big.Int) (*types.Transaction, error) {
	return _SfcV1Contract.Contract.UpdateBaseRewardPerSec(&_SfcV1Contract.TransactOpts, value)
}

// UpdateBaseRewardPerSec is a paid mutator transaction binding the contract method 0x1b593d8a.
//
// Solidity: function updateBaseRewardPerSec(uint256 value) returns()
func (_SfcV1Contract *SfcV1ContractTransactorSession) UpdateBaseRewardPerSec(value *big.Int) (*types.Transaction, error) {
	return _SfcV1Contract.Contract.UpdateBaseRewardPerSec(&_SfcV1Contract.TransactOpts, value)
}

// UpdateGasPowerAllocationRate is a paid mutator transaction binding the contract method 0x119e351a.
//
// Solidity: function updateGasPowerAllocationRate(uint256 short, uint256 long) returns()
func (_SfcV1Contract *SfcV1ContractTransactor) UpdateGasPowerAllocationRate(opts *bind.TransactOpts, short *big.Int, long *big.Int) (*types.Transaction, error) {
	return _SfcV1Contract.contract.Transact(opts, "updateGasPowerAllocationRate", short, long)
}

// UpdateGasPowerAllocationRate is a paid mutator transaction binding the contract method 0x119e351a.
//
// Solidity: function updateGasPowerAllocationRate(uint256 short, uint256 long) returns()
func (_SfcV1Contract *SfcV1ContractSession) UpdateGasPowerAllocationRate(short *big.Int, long *big.Int) (*types.Transaction, error) {
	return _SfcV1Contract.Contract.UpdateGasPowerAllocationRate(&_SfcV1Contract.TransactOpts, short, long)
}

// UpdateGasPowerAllocationRate is a paid mutator transaction binding the contract method 0x119e351a.
//
// Solidity: function updateGasPowerAllocationRate(uint256 short, uint256 long) returns()
func (_SfcV1Contract *SfcV1ContractTransactorSession) UpdateGasPowerAllocationRate(short *big.Int, long *big.Int) (*types.Transaction, error) {
	return _SfcV1Contract.Contract.UpdateGasPowerAllocationRate(&_SfcV1Contract.TransactOpts, short, long)
}

// UpdateStakerMetadata is a paid mutator transaction binding the contract method 0x33a14912.
//
// Solidity: function updateStakerMetadata(bytes metadata) returns()
func (_SfcV1Contract *SfcV1ContractTransactor) UpdateStakerMetadata(opts *bind.TransactOpts, metadata []byte) (*types.Transaction, error) {
	return _SfcV1Contract.contract.Transact(opts, "updateStakerMetadata", metadata)
}

// UpdateStakerMetadata is a paid mutator transaction binding the contract method 0x33a14912.
//
// Solidity: function updateStakerMetadata(bytes metadata) returns()
func (_SfcV1Contract *SfcV1ContractSession) UpdateStakerMetadata(metadata []byte) (*types.Transaction, error) {
	return _SfcV1Contract.Contract.UpdateStakerMetadata(&_SfcV1Contract.TransactOpts, metadata)
}

// UpdateStakerMetadata is a paid mutator transaction binding the contract method 0x33a14912.
//
// Solidity: function updateStakerMetadata(bytes metadata) returns()
func (_SfcV1Contract *SfcV1ContractTransactorSession) UpdateStakerMetadata(metadata []byte) (*types.Transaction, error) {
	return _SfcV1Contract.Contract.UpdateStakerMetadata(&_SfcV1Contract.TransactOpts, metadata)
}

// UpdateStakerSfcAddress is a paid mutator transaction binding the contract method 0xc3d74f1a.
//
// Solidity: function updateStakerSfcAddress(address newSfcAddress) returns()
func (_SfcV1Contract *SfcV1ContractTransactor) UpdateStakerSfcAddress(opts *bind.TransactOpts, newSfcAddress common.Address) (*types.Transaction, error) {
	return _SfcV1Contract.contract.Transact(opts, "updateStakerSfcAddress", newSfcAddress)
}

// UpdateStakerSfcAddress is a paid mutator transaction binding the contract method 0xc3d74f1a.
//
// Solidity: function updateStakerSfcAddress(address newSfcAddress) returns()
func (_SfcV1Contract *SfcV1ContractSession) UpdateStakerSfcAddress(newSfcAddress common.Address) (*types.Transaction, error) {
	return _SfcV1Contract.Contract.UpdateStakerSfcAddress(&_SfcV1Contract.TransactOpts, newSfcAddress)
}

// UpdateStakerSfcAddress is a paid mutator transaction binding the contract method 0xc3d74f1a.
//
// Solidity: function updateStakerSfcAddress(address newSfcAddress) returns()
func (_SfcV1Contract *SfcV1ContractTransactorSession) UpdateStakerSfcAddress(newSfcAddress common.Address) (*types.Transaction, error) {
	return _SfcV1Contract.Contract.UpdateStakerSfcAddress(&_SfcV1Contract.TransactOpts, newSfcAddress)
}

// WithdrawDelegation is a paid mutator transaction binding the contract method 0x16bfdd81.
//
// Solidity: function withdrawDelegation() returns()
func (_SfcV1Contract *SfcV1ContractTransactor) WithdrawDelegation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SfcV1Contract.contract.Transact(opts, "withdrawDelegation")
}

// WithdrawDelegation is a paid mutator transaction binding the contract method 0x16bfdd81.
//
// Solidity: function withdrawDelegation() returns()
func (_SfcV1Contract *SfcV1ContractSession) WithdrawDelegation() (*types.Transaction, error) {
	return _SfcV1Contract.Contract.WithdrawDelegation(&_SfcV1Contract.TransactOpts)
}

// WithdrawDelegation is a paid mutator transaction binding the contract method 0x16bfdd81.
//
// Solidity: function withdrawDelegation() returns()
func (_SfcV1Contract *SfcV1ContractTransactorSession) WithdrawDelegation() (*types.Transaction, error) {
	return _SfcV1Contract.Contract.WithdrawDelegation(&_SfcV1Contract.TransactOpts)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_SfcV1Contract *SfcV1ContractTransactor) WithdrawStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SfcV1Contract.contract.Transact(opts, "withdrawStake")
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_SfcV1Contract *SfcV1ContractSession) WithdrawStake() (*types.Transaction, error) {
	return _SfcV1Contract.Contract.WithdrawStake(&_SfcV1Contract.TransactOpts)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_SfcV1Contract *SfcV1ContractTransactorSession) WithdrawStake() (*types.Transaction, error) {
	return _SfcV1Contract.Contract.WithdrawStake(&_SfcV1Contract.TransactOpts)
}

// SfcV1ContractBurntRewardStashIterator is returned from FilterBurntRewardStash and is used to iterate over the raw logs and unpacked data for BurntRewardStash events raised by the SfcV1Contract contract.
type SfcV1ContractBurntRewardStashIterator struct {
	Event *SfcV1ContractBurntRewardStash // Event containing the contract specifics and raw log

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
func (it *SfcV1ContractBurntRewardStashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcV1ContractBurntRewardStash)
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
		it.Event = new(SfcV1ContractBurntRewardStash)
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
func (it *SfcV1ContractBurntRewardStashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcV1ContractBurntRewardStashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcV1ContractBurntRewardStash represents a BurntRewardStash event raised by the SfcV1Contract contract.
type SfcV1ContractBurntRewardStash struct {
	Addr         common.Address
	StakerID     *big.Int
	IsDelegation bool
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBurntRewardStash is a free log retrieval operation binding the contract event 0x0ea92567e76d40ddc52d2c1d74a521a59329a38b50411451de6ad2e565466d0f.
//
// Solidity: event BurntRewardStash(address indexed addr, uint256 indexed stakerID, bool isDelegation, uint256 amount)
func (_SfcV1Contract *SfcV1ContractFilterer) FilterBurntRewardStash(opts *bind.FilterOpts, addr []common.Address, stakerID []*big.Int) (*SfcV1ContractBurntRewardStashIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _SfcV1Contract.contract.FilterLogs(opts, "BurntRewardStash", addrRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &SfcV1ContractBurntRewardStashIterator{contract: _SfcV1Contract.contract, event: "BurntRewardStash", logs: logs, sub: sub}, nil
}

// WatchBurntRewardStash is a free log subscription operation binding the contract event 0x0ea92567e76d40ddc52d2c1d74a521a59329a38b50411451de6ad2e565466d0f.
//
// Solidity: event BurntRewardStash(address indexed addr, uint256 indexed stakerID, bool isDelegation, uint256 amount)
func (_SfcV1Contract *SfcV1ContractFilterer) WatchBurntRewardStash(opts *bind.WatchOpts, sink chan<- *SfcV1ContractBurntRewardStash, addr []common.Address, stakerID []*big.Int) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _SfcV1Contract.contract.WatchLogs(opts, "BurntRewardStash", addrRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcV1ContractBurntRewardStash)
				if err := _SfcV1Contract.contract.UnpackLog(event, "BurntRewardStash", log); err != nil {
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

// ParseBurntRewardStash is a log parse operation binding the contract event 0x0ea92567e76d40ddc52d2c1d74a521a59329a38b50411451de6ad2e565466d0f.
//
// Solidity: event BurntRewardStash(address indexed addr, uint256 indexed stakerID, bool isDelegation, uint256 amount)
func (_SfcV1Contract *SfcV1ContractFilterer) ParseBurntRewardStash(log types.Log) (*SfcV1ContractBurntRewardStash, error) {
	event := new(SfcV1ContractBurntRewardStash)
	if err := _SfcV1Contract.contract.UnpackLog(event, "BurntRewardStash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SfcV1ContractClaimedDelegationRewardIterator is returned from FilterClaimedDelegationReward and is used to iterate over the raw logs and unpacked data for ClaimedDelegationReward events raised by the SfcV1Contract contract.
type SfcV1ContractClaimedDelegationRewardIterator struct {
	Event *SfcV1ContractClaimedDelegationReward // Event containing the contract specifics and raw log

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
func (it *SfcV1ContractClaimedDelegationRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcV1ContractClaimedDelegationReward)
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
		it.Event = new(SfcV1ContractClaimedDelegationReward)
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
func (it *SfcV1ContractClaimedDelegationRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcV1ContractClaimedDelegationRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcV1ContractClaimedDelegationReward represents a ClaimedDelegationReward event raised by the SfcV1Contract contract.
type SfcV1ContractClaimedDelegationReward struct {
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
func (_SfcV1Contract *SfcV1ContractFilterer) FilterClaimedDelegationReward(opts *bind.FilterOpts, from []common.Address, stakerID []*big.Int) (*SfcV1ContractClaimedDelegationRewardIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _SfcV1Contract.contract.FilterLogs(opts, "ClaimedDelegationReward", fromRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &SfcV1ContractClaimedDelegationRewardIterator{contract: _SfcV1Contract.contract, event: "ClaimedDelegationReward", logs: logs, sub: sub}, nil
}

// WatchClaimedDelegationReward is a free log subscription operation binding the contract event 0x2676e1697cf4731b93ddb4ef54e0e5a98c06cccbbbb2202848a3c6286595e6ce.
//
// Solidity: event ClaimedDelegationReward(address indexed from, uint256 indexed stakerID, uint256 reward, uint256 fromEpoch, uint256 untilEpoch)
func (_SfcV1Contract *SfcV1ContractFilterer) WatchClaimedDelegationReward(opts *bind.WatchOpts, sink chan<- *SfcV1ContractClaimedDelegationReward, from []common.Address, stakerID []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _SfcV1Contract.contract.WatchLogs(opts, "ClaimedDelegationReward", fromRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcV1ContractClaimedDelegationReward)
				if err := _SfcV1Contract.contract.UnpackLog(event, "ClaimedDelegationReward", log); err != nil {
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
func (_SfcV1Contract *SfcV1ContractFilterer) ParseClaimedDelegationReward(log types.Log) (*SfcV1ContractClaimedDelegationReward, error) {
	event := new(SfcV1ContractClaimedDelegationReward)
	if err := _SfcV1Contract.contract.UnpackLog(event, "ClaimedDelegationReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SfcV1ContractClaimedValidatorRewardIterator is returned from FilterClaimedValidatorReward and is used to iterate over the raw logs and unpacked data for ClaimedValidatorReward events raised by the SfcV1Contract contract.
type SfcV1ContractClaimedValidatorRewardIterator struct {
	Event *SfcV1ContractClaimedValidatorReward // Event containing the contract specifics and raw log

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
func (it *SfcV1ContractClaimedValidatorRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcV1ContractClaimedValidatorReward)
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
		it.Event = new(SfcV1ContractClaimedValidatorReward)
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
func (it *SfcV1ContractClaimedValidatorRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcV1ContractClaimedValidatorRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcV1ContractClaimedValidatorReward represents a ClaimedValidatorReward event raised by the SfcV1Contract contract.
type SfcV1ContractClaimedValidatorReward struct {
	StakerID   *big.Int
	Reward     *big.Int
	FromEpoch  *big.Int
	UntilEpoch *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterClaimedValidatorReward is a free log retrieval operation binding the contract event 0x2ea54c2b22a07549d19fb5eb8e4e48ebe1c653117215e94d5468c5612750d35c.
//
// Solidity: event ClaimedValidatorReward(uint256 indexed stakerID, uint256 reward, uint256 fromEpoch, uint256 untilEpoch)
func (_SfcV1Contract *SfcV1ContractFilterer) FilterClaimedValidatorReward(opts *bind.FilterOpts, stakerID []*big.Int) (*SfcV1ContractClaimedValidatorRewardIterator, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _SfcV1Contract.contract.FilterLogs(opts, "ClaimedValidatorReward", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &SfcV1ContractClaimedValidatorRewardIterator{contract: _SfcV1Contract.contract, event: "ClaimedValidatorReward", logs: logs, sub: sub}, nil
}

// WatchClaimedValidatorReward is a free log subscription operation binding the contract event 0x2ea54c2b22a07549d19fb5eb8e4e48ebe1c653117215e94d5468c5612750d35c.
//
// Solidity: event ClaimedValidatorReward(uint256 indexed stakerID, uint256 reward, uint256 fromEpoch, uint256 untilEpoch)
func (_SfcV1Contract *SfcV1ContractFilterer) WatchClaimedValidatorReward(opts *bind.WatchOpts, sink chan<- *SfcV1ContractClaimedValidatorReward, stakerID []*big.Int) (event.Subscription, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _SfcV1Contract.contract.WatchLogs(opts, "ClaimedValidatorReward", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcV1ContractClaimedValidatorReward)
				if err := _SfcV1Contract.contract.UnpackLog(event, "ClaimedValidatorReward", log); err != nil {
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
func (_SfcV1Contract *SfcV1ContractFilterer) ParseClaimedValidatorReward(log types.Log) (*SfcV1ContractClaimedValidatorReward, error) {
	event := new(SfcV1ContractClaimedValidatorReward)
	if err := _SfcV1Contract.contract.UnpackLog(event, "ClaimedValidatorReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SfcV1ContractCreatedDelegationIterator is returned from FilterCreatedDelegation and is used to iterate over the raw logs and unpacked data for CreatedDelegation events raised by the SfcV1Contract contract.
type SfcV1ContractCreatedDelegationIterator struct {
	Event *SfcV1ContractCreatedDelegation // Event containing the contract specifics and raw log

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
func (it *SfcV1ContractCreatedDelegationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcV1ContractCreatedDelegation)
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
		it.Event = new(SfcV1ContractCreatedDelegation)
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
func (it *SfcV1ContractCreatedDelegationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcV1ContractCreatedDelegationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcV1ContractCreatedDelegation represents a CreatedDelegation event raised by the SfcV1Contract contract.
type SfcV1ContractCreatedDelegation struct {
	Delegator  common.Address
	ToStakerID *big.Int
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCreatedDelegation is a free log retrieval operation binding the contract event 0xfd8c857fb9acd6f4ad59b8621a2a77825168b7b4b76de9586d08e00d4ed462be.
//
// Solidity: event CreatedDelegation(address indexed delegator, uint256 indexed toStakerID, uint256 amount)
func (_SfcV1Contract *SfcV1ContractFilterer) FilterCreatedDelegation(opts *bind.FilterOpts, delegator []common.Address, toStakerID []*big.Int) (*SfcV1ContractCreatedDelegationIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var toStakerIDRule []interface{}
	for _, toStakerIDItem := range toStakerID {
		toStakerIDRule = append(toStakerIDRule, toStakerIDItem)
	}

	logs, sub, err := _SfcV1Contract.contract.FilterLogs(opts, "CreatedDelegation", delegatorRule, toStakerIDRule)
	if err != nil {
		return nil, err
	}
	return &SfcV1ContractCreatedDelegationIterator{contract: _SfcV1Contract.contract, event: "CreatedDelegation", logs: logs, sub: sub}, nil
}

// WatchCreatedDelegation is a free log subscription operation binding the contract event 0xfd8c857fb9acd6f4ad59b8621a2a77825168b7b4b76de9586d08e00d4ed462be.
//
// Solidity: event CreatedDelegation(address indexed delegator, uint256 indexed toStakerID, uint256 amount)
func (_SfcV1Contract *SfcV1ContractFilterer) WatchCreatedDelegation(opts *bind.WatchOpts, sink chan<- *SfcV1ContractCreatedDelegation, delegator []common.Address, toStakerID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var toStakerIDRule []interface{}
	for _, toStakerIDItem := range toStakerID {
		toStakerIDRule = append(toStakerIDRule, toStakerIDItem)
	}

	logs, sub, err := _SfcV1Contract.contract.WatchLogs(opts, "CreatedDelegation", delegatorRule, toStakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcV1ContractCreatedDelegation)
				if err := _SfcV1Contract.contract.UnpackLog(event, "CreatedDelegation", log); err != nil {
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
// Solidity: event CreatedDelegation(address indexed delegator, uint256 indexed toStakerID, uint256 amount)
func (_SfcV1Contract *SfcV1ContractFilterer) ParseCreatedDelegation(log types.Log) (*SfcV1ContractCreatedDelegation, error) {
	event := new(SfcV1ContractCreatedDelegation)
	if err := _SfcV1Contract.contract.UnpackLog(event, "CreatedDelegation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SfcV1ContractCreatedStakeIterator is returned from FilterCreatedStake and is used to iterate over the raw logs and unpacked data for CreatedStake events raised by the SfcV1Contract contract.
type SfcV1ContractCreatedStakeIterator struct {
	Event *SfcV1ContractCreatedStake // Event containing the contract specifics and raw log

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
func (it *SfcV1ContractCreatedStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcV1ContractCreatedStake)
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
		it.Event = new(SfcV1ContractCreatedStake)
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
func (it *SfcV1ContractCreatedStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcV1ContractCreatedStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcV1ContractCreatedStake represents a CreatedStake event raised by the SfcV1Contract contract.
type SfcV1ContractCreatedStake struct {
	StakerID      *big.Int
	DagSfcAddress common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCreatedStake is a free log retrieval operation binding the contract event 0x0697dfe5062b9db8108e4b31254f47a912ae6bbb78837667b2e923a6f5160d39.
//
// Solidity: event CreatedStake(uint256 indexed stakerID, address indexed dagSfcAddress, uint256 amount)
func (_SfcV1Contract *SfcV1ContractFilterer) FilterCreatedStake(opts *bind.FilterOpts, stakerID []*big.Int, dagSfcAddress []common.Address) (*SfcV1ContractCreatedStakeIterator, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}
	var dagSfcAddressRule []interface{}
	for _, dagSfcAddressItem := range dagSfcAddress {
		dagSfcAddressRule = append(dagSfcAddressRule, dagSfcAddressItem)
	}

	logs, sub, err := _SfcV1Contract.contract.FilterLogs(opts, "CreatedStake", stakerIDRule, dagSfcAddressRule)
	if err != nil {
		return nil, err
	}
	return &SfcV1ContractCreatedStakeIterator{contract: _SfcV1Contract.contract, event: "CreatedStake", logs: logs, sub: sub}, nil
}

// WatchCreatedStake is a free log subscription operation binding the contract event 0x0697dfe5062b9db8108e4b31254f47a912ae6bbb78837667b2e923a6f5160d39.
//
// Solidity: event CreatedStake(uint256 indexed stakerID, address indexed dagSfcAddress, uint256 amount)
func (_SfcV1Contract *SfcV1ContractFilterer) WatchCreatedStake(opts *bind.WatchOpts, sink chan<- *SfcV1ContractCreatedStake, stakerID []*big.Int, dagSfcAddress []common.Address) (event.Subscription, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}
	var dagSfcAddressRule []interface{}
	for _, dagSfcAddressItem := range dagSfcAddress {
		dagSfcAddressRule = append(dagSfcAddressRule, dagSfcAddressItem)
	}

	logs, sub, err := _SfcV1Contract.contract.WatchLogs(opts, "CreatedStake", stakerIDRule, dagSfcAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcV1ContractCreatedStake)
				if err := _SfcV1Contract.contract.UnpackLog(event, "CreatedStake", log); err != nil {
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
// Solidity: event CreatedStake(uint256 indexed stakerID, address indexed dagSfcAddress, uint256 amount)
func (_SfcV1Contract *SfcV1ContractFilterer) ParseCreatedStake(log types.Log) (*SfcV1ContractCreatedStake, error) {
	event := new(SfcV1ContractCreatedStake)
	if err := _SfcV1Contract.contract.UnpackLog(event, "CreatedStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SfcV1ContractCreatedWithdrawRequestIterator is returned from FilterCreatedWithdrawRequest and is used to iterate over the raw logs and unpacked data for CreatedWithdrawRequest events raised by the SfcV1Contract contract.
type SfcV1ContractCreatedWithdrawRequestIterator struct {
	Event *SfcV1ContractCreatedWithdrawRequest // Event containing the contract specifics and raw log

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
func (it *SfcV1ContractCreatedWithdrawRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcV1ContractCreatedWithdrawRequest)
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
		it.Event = new(SfcV1ContractCreatedWithdrawRequest)
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
func (it *SfcV1ContractCreatedWithdrawRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcV1ContractCreatedWithdrawRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcV1ContractCreatedWithdrawRequest represents a CreatedWithdrawRequest event raised by the SfcV1Contract contract.
type SfcV1ContractCreatedWithdrawRequest struct {
	Auth       common.Address
	Receiver   common.Address
	StakerID   *big.Int
	WrID       *big.Int
	Delegation bool
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCreatedWithdrawRequest is a free log retrieval operation binding the contract event 0xde2d2a87af2fa2de55bde86f04143144eb632fa6be266dc224341a371fb8916d.
//
// Solidity: event CreatedWithdrawRequest(address indexed auth, address indexed receiver, uint256 indexed stakerID, uint256 wrID, bool delegation, uint256 amount)
func (_SfcV1Contract *SfcV1ContractFilterer) FilterCreatedWithdrawRequest(opts *bind.FilterOpts, auth []common.Address, receiver []common.Address, stakerID []*big.Int) (*SfcV1ContractCreatedWithdrawRequestIterator, error) {

	var authRule []interface{}
	for _, authItem := range auth {
		authRule = append(authRule, authItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _SfcV1Contract.contract.FilterLogs(opts, "CreatedWithdrawRequest", authRule, receiverRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &SfcV1ContractCreatedWithdrawRequestIterator{contract: _SfcV1Contract.contract, event: "CreatedWithdrawRequest", logs: logs, sub: sub}, nil
}

// WatchCreatedWithdrawRequest is a free log subscription operation binding the contract event 0xde2d2a87af2fa2de55bde86f04143144eb632fa6be266dc224341a371fb8916d.
//
// Solidity: event CreatedWithdrawRequest(address indexed auth, address indexed receiver, uint256 indexed stakerID, uint256 wrID, bool delegation, uint256 amount)
func (_SfcV1Contract *SfcV1ContractFilterer) WatchCreatedWithdrawRequest(opts *bind.WatchOpts, sink chan<- *SfcV1ContractCreatedWithdrawRequest, auth []common.Address, receiver []common.Address, stakerID []*big.Int) (event.Subscription, error) {

	var authRule []interface{}
	for _, authItem := range auth {
		authRule = append(authRule, authItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _SfcV1Contract.contract.WatchLogs(opts, "CreatedWithdrawRequest", authRule, receiverRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcV1ContractCreatedWithdrawRequest)
				if err := _SfcV1Contract.contract.UnpackLog(event, "CreatedWithdrawRequest", log); err != nil {
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

// ParseCreatedWithdrawRequest is a log parse operation binding the contract event 0xde2d2a87af2fa2de55bde86f04143144eb632fa6be266dc224341a371fb8916d.
//
// Solidity: event CreatedWithdrawRequest(address indexed auth, address indexed receiver, uint256 indexed stakerID, uint256 wrID, bool delegation, uint256 amount)
func (_SfcV1Contract *SfcV1ContractFilterer) ParseCreatedWithdrawRequest(log types.Log) (*SfcV1ContractCreatedWithdrawRequest, error) {
	event := new(SfcV1ContractCreatedWithdrawRequest)
	if err := _SfcV1Contract.contract.UnpackLog(event, "CreatedWithdrawRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SfcV1ContractDeactivatedDelegationIterator is returned from FilterDeactivatedDelegation and is used to iterate over the raw logs and unpacked data for DeactivatedDelegation events raised by the SfcV1Contract contract.
type SfcV1ContractDeactivatedDelegationIterator struct {
	Event *SfcV1ContractDeactivatedDelegation // Event containing the contract specifics and raw log

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
func (it *SfcV1ContractDeactivatedDelegationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcV1ContractDeactivatedDelegation)
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
		it.Event = new(SfcV1ContractDeactivatedDelegation)
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
func (it *SfcV1ContractDeactivatedDelegationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcV1ContractDeactivatedDelegationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcV1ContractDeactivatedDelegation represents a DeactivatedDelegation event raised by the SfcV1Contract contract.
type SfcV1ContractDeactivatedDelegation struct {
	Delegator common.Address
	StakerID  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeactivatedDelegation is a free log retrieval operation binding the contract event 0x912c4125a208704a342cbdc4726795d26556b0170b7fc95bc706d5cb1f506469.
//
// Solidity: event DeactivatedDelegation(address indexed delegator, uint256 indexed stakerID)
func (_SfcV1Contract *SfcV1ContractFilterer) FilterDeactivatedDelegation(opts *bind.FilterOpts, delegator []common.Address, stakerID []*big.Int) (*SfcV1ContractDeactivatedDelegationIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _SfcV1Contract.contract.FilterLogs(opts, "DeactivatedDelegation", delegatorRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &SfcV1ContractDeactivatedDelegationIterator{contract: _SfcV1Contract.contract, event: "DeactivatedDelegation", logs: logs, sub: sub}, nil
}

// WatchDeactivatedDelegation is a free log subscription operation binding the contract event 0x912c4125a208704a342cbdc4726795d26556b0170b7fc95bc706d5cb1f506469.
//
// Solidity: event DeactivatedDelegation(address indexed delegator, uint256 indexed stakerID)
func (_SfcV1Contract *SfcV1ContractFilterer) WatchDeactivatedDelegation(opts *bind.WatchOpts, sink chan<- *SfcV1ContractDeactivatedDelegation, delegator []common.Address, stakerID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _SfcV1Contract.contract.WatchLogs(opts, "DeactivatedDelegation", delegatorRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcV1ContractDeactivatedDelegation)
				if err := _SfcV1Contract.contract.UnpackLog(event, "DeactivatedDelegation", log); err != nil {
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

// ParseDeactivatedDelegation is a log parse operation binding the contract event 0x912c4125a208704a342cbdc4726795d26556b0170b7fc95bc706d5cb1f506469.
//
// Solidity: event DeactivatedDelegation(address indexed delegator, uint256 indexed stakerID)
func (_SfcV1Contract *SfcV1ContractFilterer) ParseDeactivatedDelegation(log types.Log) (*SfcV1ContractDeactivatedDelegation, error) {
	event := new(SfcV1ContractDeactivatedDelegation)
	if err := _SfcV1Contract.contract.UnpackLog(event, "DeactivatedDelegation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SfcV1ContractDeactivatedStakeIterator is returned from FilterDeactivatedStake and is used to iterate over the raw logs and unpacked data for DeactivatedStake events raised by the SfcV1Contract contract.
type SfcV1ContractDeactivatedStakeIterator struct {
	Event *SfcV1ContractDeactivatedStake // Event containing the contract specifics and raw log

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
func (it *SfcV1ContractDeactivatedStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcV1ContractDeactivatedStake)
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
		it.Event = new(SfcV1ContractDeactivatedStake)
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
func (it *SfcV1ContractDeactivatedStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcV1ContractDeactivatedStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcV1ContractDeactivatedStake represents a DeactivatedStake event raised by the SfcV1Contract contract.
type SfcV1ContractDeactivatedStake struct {
	StakerID *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeactivatedStake is a free log retrieval operation binding the contract event 0xf7c308d0d978cce3aec157d1b34e355db4636b4e71ce91b4f5ec9e7a4f5cdc60.
//
// Solidity: event DeactivatedStake(uint256 indexed stakerID)
func (_SfcV1Contract *SfcV1ContractFilterer) FilterDeactivatedStake(opts *bind.FilterOpts, stakerID []*big.Int) (*SfcV1ContractDeactivatedStakeIterator, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _SfcV1Contract.contract.FilterLogs(opts, "DeactivatedStake", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &SfcV1ContractDeactivatedStakeIterator{contract: _SfcV1Contract.contract, event: "DeactivatedStake", logs: logs, sub: sub}, nil
}

// WatchDeactivatedStake is a free log subscription operation binding the contract event 0xf7c308d0d978cce3aec157d1b34e355db4636b4e71ce91b4f5ec9e7a4f5cdc60.
//
// Solidity: event DeactivatedStake(uint256 indexed stakerID)
func (_SfcV1Contract *SfcV1ContractFilterer) WatchDeactivatedStake(opts *bind.WatchOpts, sink chan<- *SfcV1ContractDeactivatedStake, stakerID []*big.Int) (event.Subscription, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _SfcV1Contract.contract.WatchLogs(opts, "DeactivatedStake", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcV1ContractDeactivatedStake)
				if err := _SfcV1Contract.contract.UnpackLog(event, "DeactivatedStake", log); err != nil {
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

// ParseDeactivatedStake is a log parse operation binding the contract event 0xf7c308d0d978cce3aec157d1b34e355db4636b4e71ce91b4f5ec9e7a4f5cdc60.
//
// Solidity: event DeactivatedStake(uint256 indexed stakerID)
func (_SfcV1Contract *SfcV1ContractFilterer) ParseDeactivatedStake(log types.Log) (*SfcV1ContractDeactivatedStake, error) {
	event := new(SfcV1ContractDeactivatedStake)
	if err := _SfcV1Contract.contract.UnpackLog(event, "DeactivatedStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SfcV1ContractIncreasedDelegationIterator is returned from FilterIncreasedDelegation and is used to iterate over the raw logs and unpacked data for IncreasedDelegation events raised by the SfcV1Contract contract.
type SfcV1ContractIncreasedDelegationIterator struct {
	Event *SfcV1ContractIncreasedDelegation // Event containing the contract specifics and raw log

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
func (it *SfcV1ContractIncreasedDelegationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcV1ContractIncreasedDelegation)
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
		it.Event = new(SfcV1ContractIncreasedDelegation)
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
func (it *SfcV1ContractIncreasedDelegationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcV1ContractIncreasedDelegationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcV1ContractIncreasedDelegation represents a IncreasedDelegation event raised by the SfcV1Contract contract.
type SfcV1ContractIncreasedDelegation struct {
	Delegator common.Address
	StakerID  *big.Int
	NewAmount *big.Int
	Diff      *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterIncreasedDelegation is a free log retrieval operation binding the contract event 0x4ca781bfe171e588a2661d5a7f2f5f59df879c53489063552fbad2145b707fc1.
//
// Solidity: event IncreasedDelegation(address indexed delegator, uint256 indexed stakerID, uint256 newAmount, uint256 diff)
func (_SfcV1Contract *SfcV1ContractFilterer) FilterIncreasedDelegation(opts *bind.FilterOpts, delegator []common.Address, stakerID []*big.Int) (*SfcV1ContractIncreasedDelegationIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _SfcV1Contract.contract.FilterLogs(opts, "IncreasedDelegation", delegatorRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &SfcV1ContractIncreasedDelegationIterator{contract: _SfcV1Contract.contract, event: "IncreasedDelegation", logs: logs, sub: sub}, nil
}

// WatchIncreasedDelegation is a free log subscription operation binding the contract event 0x4ca781bfe171e588a2661d5a7f2f5f59df879c53489063552fbad2145b707fc1.
//
// Solidity: event IncreasedDelegation(address indexed delegator, uint256 indexed stakerID, uint256 newAmount, uint256 diff)
func (_SfcV1Contract *SfcV1ContractFilterer) WatchIncreasedDelegation(opts *bind.WatchOpts, sink chan<- *SfcV1ContractIncreasedDelegation, delegator []common.Address, stakerID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _SfcV1Contract.contract.WatchLogs(opts, "IncreasedDelegation", delegatorRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcV1ContractIncreasedDelegation)
				if err := _SfcV1Contract.contract.UnpackLog(event, "IncreasedDelegation", log); err != nil {
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

// ParseIncreasedDelegation is a log parse operation binding the contract event 0x4ca781bfe171e588a2661d5a7f2f5f59df879c53489063552fbad2145b707fc1.
//
// Solidity: event IncreasedDelegation(address indexed delegator, uint256 indexed stakerID, uint256 newAmount, uint256 diff)
func (_SfcV1Contract *SfcV1ContractFilterer) ParseIncreasedDelegation(log types.Log) (*SfcV1ContractIncreasedDelegation, error) {
	event := new(SfcV1ContractIncreasedDelegation)
	if err := _SfcV1Contract.contract.UnpackLog(event, "IncreasedDelegation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SfcV1ContractIncreasedStakeIterator is returned from FilterIncreasedStake and is used to iterate over the raw logs and unpacked data for IncreasedStake events raised by the SfcV1Contract contract.
type SfcV1ContractIncreasedStakeIterator struct {
	Event *SfcV1ContractIncreasedStake // Event containing the contract specifics and raw log

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
func (it *SfcV1ContractIncreasedStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcV1ContractIncreasedStake)
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
		it.Event = new(SfcV1ContractIncreasedStake)
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
func (it *SfcV1ContractIncreasedStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcV1ContractIncreasedStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcV1ContractIncreasedStake represents a IncreasedStake event raised by the SfcV1Contract contract.
type SfcV1ContractIncreasedStake struct {
	StakerID  *big.Int
	NewAmount *big.Int
	Diff      *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterIncreasedStake is a free log retrieval operation binding the contract event 0xa1d93e9a2a16bf4c2d0cdc6f47fe0fa054c741c96b3dac1297c79eaca31714e9.
//
// Solidity: event IncreasedStake(uint256 indexed stakerID, uint256 newAmount, uint256 diff)
func (_SfcV1Contract *SfcV1ContractFilterer) FilterIncreasedStake(opts *bind.FilterOpts, stakerID []*big.Int) (*SfcV1ContractIncreasedStakeIterator, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _SfcV1Contract.contract.FilterLogs(opts, "IncreasedStake", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &SfcV1ContractIncreasedStakeIterator{contract: _SfcV1Contract.contract, event: "IncreasedStake", logs: logs, sub: sub}, nil
}

// WatchIncreasedStake is a free log subscription operation binding the contract event 0xa1d93e9a2a16bf4c2d0cdc6f47fe0fa054c741c96b3dac1297c79eaca31714e9.
//
// Solidity: event IncreasedStake(uint256 indexed stakerID, uint256 newAmount, uint256 diff)
func (_SfcV1Contract *SfcV1ContractFilterer) WatchIncreasedStake(opts *bind.WatchOpts, sink chan<- *SfcV1ContractIncreasedStake, stakerID []*big.Int) (event.Subscription, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _SfcV1Contract.contract.WatchLogs(opts, "IncreasedStake", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcV1ContractIncreasedStake)
				if err := _SfcV1Contract.contract.UnpackLog(event, "IncreasedStake", log); err != nil {
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
func (_SfcV1Contract *SfcV1ContractFilterer) ParseIncreasedStake(log types.Log) (*SfcV1ContractIncreasedStake, error) {
	event := new(SfcV1ContractIncreasedStake)
	if err := _SfcV1Contract.contract.UnpackLog(event, "IncreasedStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SfcV1ContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SfcV1Contract contract.
type SfcV1ContractOwnershipTransferredIterator struct {
	Event *SfcV1ContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SfcV1ContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcV1ContractOwnershipTransferred)
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
		it.Event = new(SfcV1ContractOwnershipTransferred)
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
func (it *SfcV1ContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcV1ContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcV1ContractOwnershipTransferred represents a OwnershipTransferred event raised by the SfcV1Contract contract.
type SfcV1ContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SfcV1Contract *SfcV1ContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SfcV1ContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SfcV1Contract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SfcV1ContractOwnershipTransferredIterator{contract: _SfcV1Contract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SfcV1Contract *SfcV1ContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SfcV1ContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SfcV1Contract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcV1ContractOwnershipTransferred)
				if err := _SfcV1Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SfcV1Contract *SfcV1ContractFilterer) ParseOwnershipTransferred(log types.Log) (*SfcV1ContractOwnershipTransferred, error) {
	event := new(SfcV1ContractOwnershipTransferred)
	if err := _SfcV1Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SfcV1ContractPartialWithdrawnByRequestIterator is returned from FilterPartialWithdrawnByRequest and is used to iterate over the raw logs and unpacked data for PartialWithdrawnByRequest events raised by the SfcV1Contract contract.
type SfcV1ContractPartialWithdrawnByRequestIterator struct {
	Event *SfcV1ContractPartialWithdrawnByRequest // Event containing the contract specifics and raw log

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
func (it *SfcV1ContractPartialWithdrawnByRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcV1ContractPartialWithdrawnByRequest)
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
		it.Event = new(SfcV1ContractPartialWithdrawnByRequest)
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
func (it *SfcV1ContractPartialWithdrawnByRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcV1ContractPartialWithdrawnByRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcV1ContractPartialWithdrawnByRequest represents a PartialWithdrawnByRequest event raised by the SfcV1Contract contract.
type SfcV1ContractPartialWithdrawnByRequest struct {
	Auth       common.Address
	Receiver   common.Address
	StakerID   *big.Int
	WrID       *big.Int
	Delegation bool
	Penalty    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPartialWithdrawnByRequest is a free log retrieval operation binding the contract event 0xd5304dabc5bd47105b6921889d1b528c4b2223250248a916afd129b1c0512ddd.
//
// Solidity: event PartialWithdrawnByRequest(address indexed auth, address indexed receiver, uint256 indexed stakerID, uint256 wrID, bool delegation, uint256 penalty)
func (_SfcV1Contract *SfcV1ContractFilterer) FilterPartialWithdrawnByRequest(opts *bind.FilterOpts, auth []common.Address, receiver []common.Address, stakerID []*big.Int) (*SfcV1ContractPartialWithdrawnByRequestIterator, error) {

	var authRule []interface{}
	for _, authItem := range auth {
		authRule = append(authRule, authItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _SfcV1Contract.contract.FilterLogs(opts, "PartialWithdrawnByRequest", authRule, receiverRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &SfcV1ContractPartialWithdrawnByRequestIterator{contract: _SfcV1Contract.contract, event: "PartialWithdrawnByRequest", logs: logs, sub: sub}, nil
}

// WatchPartialWithdrawnByRequest is a free log subscription operation binding the contract event 0xd5304dabc5bd47105b6921889d1b528c4b2223250248a916afd129b1c0512ddd.
//
// Solidity: event PartialWithdrawnByRequest(address indexed auth, address indexed receiver, uint256 indexed stakerID, uint256 wrID, bool delegation, uint256 penalty)
func (_SfcV1Contract *SfcV1ContractFilterer) WatchPartialWithdrawnByRequest(opts *bind.WatchOpts, sink chan<- *SfcV1ContractPartialWithdrawnByRequest, auth []common.Address, receiver []common.Address, stakerID []*big.Int) (event.Subscription, error) {

	var authRule []interface{}
	for _, authItem := range auth {
		authRule = append(authRule, authItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _SfcV1Contract.contract.WatchLogs(opts, "PartialWithdrawnByRequest", authRule, receiverRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcV1ContractPartialWithdrawnByRequest)
				if err := _SfcV1Contract.contract.UnpackLog(event, "PartialWithdrawnByRequest", log); err != nil {
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

// ParsePartialWithdrawnByRequest is a log parse operation binding the contract event 0xd5304dabc5bd47105b6921889d1b528c4b2223250248a916afd129b1c0512ddd.
//
// Solidity: event PartialWithdrawnByRequest(address indexed auth, address indexed receiver, uint256 indexed stakerID, uint256 wrID, bool delegation, uint256 penalty)
func (_SfcV1Contract *SfcV1ContractFilterer) ParsePartialWithdrawnByRequest(log types.Log) (*SfcV1ContractPartialWithdrawnByRequest, error) {
	event := new(SfcV1ContractPartialWithdrawnByRequest)
	if err := _SfcV1Contract.contract.UnpackLog(event, "PartialWithdrawnByRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SfcV1ContractPreparedToWithdrawDelegationIterator is returned from FilterPreparedToWithdrawDelegation and is used to iterate over the raw logs and unpacked data for PreparedToWithdrawDelegation events raised by the SfcV1Contract contract.
type SfcV1ContractPreparedToWithdrawDelegationIterator struct {
	Event *SfcV1ContractPreparedToWithdrawDelegation // Event containing the contract specifics and raw log

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
func (it *SfcV1ContractPreparedToWithdrawDelegationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcV1ContractPreparedToWithdrawDelegation)
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
		it.Event = new(SfcV1ContractPreparedToWithdrawDelegation)
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
func (it *SfcV1ContractPreparedToWithdrawDelegationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcV1ContractPreparedToWithdrawDelegationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcV1ContractPreparedToWithdrawDelegation represents a PreparedToWithdrawDelegation event raised by the SfcV1Contract contract.
type SfcV1ContractPreparedToWithdrawDelegation struct {
	Delegator common.Address
	StakerID  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPreparedToWithdrawDelegation is a free log retrieval operation binding the contract event 0x5b1eea49e405ef6d509836aac841959c30bb0673b1fd70859bfc6ae5e4ee3df2.
//
// Solidity: event PreparedToWithdrawDelegation(address indexed delegator, uint256 indexed stakerID)
func (_SfcV1Contract *SfcV1ContractFilterer) FilterPreparedToWithdrawDelegation(opts *bind.FilterOpts, delegator []common.Address, stakerID []*big.Int) (*SfcV1ContractPreparedToWithdrawDelegationIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _SfcV1Contract.contract.FilterLogs(opts, "PreparedToWithdrawDelegation", delegatorRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &SfcV1ContractPreparedToWithdrawDelegationIterator{contract: _SfcV1Contract.contract, event: "PreparedToWithdrawDelegation", logs: logs, sub: sub}, nil
}

// WatchPreparedToWithdrawDelegation is a free log subscription operation binding the contract event 0x5b1eea49e405ef6d509836aac841959c30bb0673b1fd70859bfc6ae5e4ee3df2.
//
// Solidity: event PreparedToWithdrawDelegation(address indexed delegator, uint256 indexed stakerID)
func (_SfcV1Contract *SfcV1ContractFilterer) WatchPreparedToWithdrawDelegation(opts *bind.WatchOpts, sink chan<- *SfcV1ContractPreparedToWithdrawDelegation, delegator []common.Address, stakerID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _SfcV1Contract.contract.WatchLogs(opts, "PreparedToWithdrawDelegation", delegatorRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcV1ContractPreparedToWithdrawDelegation)
				if err := _SfcV1Contract.contract.UnpackLog(event, "PreparedToWithdrawDelegation", log); err != nil {
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
// Solidity: event PreparedToWithdrawDelegation(address indexed delegator, uint256 indexed stakerID)
func (_SfcV1Contract *SfcV1ContractFilterer) ParsePreparedToWithdrawDelegation(log types.Log) (*SfcV1ContractPreparedToWithdrawDelegation, error) {
	event := new(SfcV1ContractPreparedToWithdrawDelegation)
	if err := _SfcV1Contract.contract.UnpackLog(event, "PreparedToWithdrawDelegation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SfcV1ContractPreparedToWithdrawStakeIterator is returned from FilterPreparedToWithdrawStake and is used to iterate over the raw logs and unpacked data for PreparedToWithdrawStake events raised by the SfcV1Contract contract.
type SfcV1ContractPreparedToWithdrawStakeIterator struct {
	Event *SfcV1ContractPreparedToWithdrawStake // Event containing the contract specifics and raw log

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
func (it *SfcV1ContractPreparedToWithdrawStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcV1ContractPreparedToWithdrawStake)
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
		it.Event = new(SfcV1ContractPreparedToWithdrawStake)
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
func (it *SfcV1ContractPreparedToWithdrawStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcV1ContractPreparedToWithdrawStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcV1ContractPreparedToWithdrawStake represents a PreparedToWithdrawStake event raised by the SfcV1Contract contract.
type SfcV1ContractPreparedToWithdrawStake struct {
	StakerID *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPreparedToWithdrawStake is a free log retrieval operation binding the contract event 0x84244546a9da4942f506db48ff90ebc240c73bb399e3e47d58843c6bb60e7185.
//
// Solidity: event PreparedToWithdrawStake(uint256 indexed stakerID)
func (_SfcV1Contract *SfcV1ContractFilterer) FilterPreparedToWithdrawStake(opts *bind.FilterOpts, stakerID []*big.Int) (*SfcV1ContractPreparedToWithdrawStakeIterator, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _SfcV1Contract.contract.FilterLogs(opts, "PreparedToWithdrawStake", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &SfcV1ContractPreparedToWithdrawStakeIterator{contract: _SfcV1Contract.contract, event: "PreparedToWithdrawStake", logs: logs, sub: sub}, nil
}

// WatchPreparedToWithdrawStake is a free log subscription operation binding the contract event 0x84244546a9da4942f506db48ff90ebc240c73bb399e3e47d58843c6bb60e7185.
//
// Solidity: event PreparedToWithdrawStake(uint256 indexed stakerID)
func (_SfcV1Contract *SfcV1ContractFilterer) WatchPreparedToWithdrawStake(opts *bind.WatchOpts, sink chan<- *SfcV1ContractPreparedToWithdrawStake, stakerID []*big.Int) (event.Subscription, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _SfcV1Contract.contract.WatchLogs(opts, "PreparedToWithdrawStake", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcV1ContractPreparedToWithdrawStake)
				if err := _SfcV1Contract.contract.UnpackLog(event, "PreparedToWithdrawStake", log); err != nil {
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
func (_SfcV1Contract *SfcV1ContractFilterer) ParsePreparedToWithdrawStake(log types.Log) (*SfcV1ContractPreparedToWithdrawStake, error) {
	event := new(SfcV1ContractPreparedToWithdrawStake)
	if err := _SfcV1Contract.contract.UnpackLog(event, "PreparedToWithdrawStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SfcV1ContractUnstashedRewardsIterator is returned from FilterUnstashedRewards and is used to iterate over the raw logs and unpacked data for UnstashedRewards events raised by the SfcV1Contract contract.
type SfcV1ContractUnstashedRewardsIterator struct {
	Event *SfcV1ContractUnstashedRewards // Event containing the contract specifics and raw log

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
func (it *SfcV1ContractUnstashedRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcV1ContractUnstashedRewards)
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
		it.Event = new(SfcV1ContractUnstashedRewards)
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
func (it *SfcV1ContractUnstashedRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcV1ContractUnstashedRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcV1ContractUnstashedRewards represents a UnstashedRewards event raised by the SfcV1Contract contract.
type SfcV1ContractUnstashedRewards struct {
	Auth     common.Address
	Receiver common.Address
	Rewards  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUnstashedRewards is a free log retrieval operation binding the contract event 0x80b36a0e929d7e7925087e54acfeecf4c6043e451b9d71ac5e908b66f9e5d126.
//
// Solidity: event UnstashedRewards(address indexed auth, address indexed receiver, uint256 rewards)
func (_SfcV1Contract *SfcV1ContractFilterer) FilterUnstashedRewards(opts *bind.FilterOpts, auth []common.Address, receiver []common.Address) (*SfcV1ContractUnstashedRewardsIterator, error) {

	var authRule []interface{}
	for _, authItem := range auth {
		authRule = append(authRule, authItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _SfcV1Contract.contract.FilterLogs(opts, "UnstashedRewards", authRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &SfcV1ContractUnstashedRewardsIterator{contract: _SfcV1Contract.contract, event: "UnstashedRewards", logs: logs, sub: sub}, nil
}

// WatchUnstashedRewards is a free log subscription operation binding the contract event 0x80b36a0e929d7e7925087e54acfeecf4c6043e451b9d71ac5e908b66f9e5d126.
//
// Solidity: event UnstashedRewards(address indexed auth, address indexed receiver, uint256 rewards)
func (_SfcV1Contract *SfcV1ContractFilterer) WatchUnstashedRewards(opts *bind.WatchOpts, sink chan<- *SfcV1ContractUnstashedRewards, auth []common.Address, receiver []common.Address) (event.Subscription, error) {

	var authRule []interface{}
	for _, authItem := range auth {
		authRule = append(authRule, authItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _SfcV1Contract.contract.WatchLogs(opts, "UnstashedRewards", authRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcV1ContractUnstashedRewards)
				if err := _SfcV1Contract.contract.UnpackLog(event, "UnstashedRewards", log); err != nil {
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

// ParseUnstashedRewards is a log parse operation binding the contract event 0x80b36a0e929d7e7925087e54acfeecf4c6043e451b9d71ac5e908b66f9e5d126.
//
// Solidity: event UnstashedRewards(address indexed auth, address indexed receiver, uint256 rewards)
func (_SfcV1Contract *SfcV1ContractFilterer) ParseUnstashedRewards(log types.Log) (*SfcV1ContractUnstashedRewards, error) {
	event := new(SfcV1ContractUnstashedRewards)
	if err := _SfcV1Contract.contract.UnpackLog(event, "UnstashedRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SfcV1ContractUpdatedBaseRewardPerSecIterator is returned from FilterUpdatedBaseRewardPerSec and is used to iterate over the raw logs and unpacked data for UpdatedBaseRewardPerSec events raised by the SfcV1Contract contract.
type SfcV1ContractUpdatedBaseRewardPerSecIterator struct {
	Event *SfcV1ContractUpdatedBaseRewardPerSec // Event containing the contract specifics and raw log

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
func (it *SfcV1ContractUpdatedBaseRewardPerSecIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcV1ContractUpdatedBaseRewardPerSec)
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
		it.Event = new(SfcV1ContractUpdatedBaseRewardPerSec)
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
func (it *SfcV1ContractUpdatedBaseRewardPerSecIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcV1ContractUpdatedBaseRewardPerSecIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcV1ContractUpdatedBaseRewardPerSec represents a UpdatedBaseRewardPerSec event raised by the SfcV1Contract contract.
type SfcV1ContractUpdatedBaseRewardPerSec struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdatedBaseRewardPerSec is a free log retrieval operation binding the contract event 0x8cd9dae1bbea2bc8a5e80ffce2c224727a25925130a03ae100619a8861ae2396.
//
// Solidity: event UpdatedBaseRewardPerSec(uint256 value)
func (_SfcV1Contract *SfcV1ContractFilterer) FilterUpdatedBaseRewardPerSec(opts *bind.FilterOpts) (*SfcV1ContractUpdatedBaseRewardPerSecIterator, error) {

	logs, sub, err := _SfcV1Contract.contract.FilterLogs(opts, "UpdatedBaseRewardPerSec")
	if err != nil {
		return nil, err
	}
	return &SfcV1ContractUpdatedBaseRewardPerSecIterator{contract: _SfcV1Contract.contract, event: "UpdatedBaseRewardPerSec", logs: logs, sub: sub}, nil
}

// WatchUpdatedBaseRewardPerSec is a free log subscription operation binding the contract event 0x8cd9dae1bbea2bc8a5e80ffce2c224727a25925130a03ae100619a8861ae2396.
//
// Solidity: event UpdatedBaseRewardPerSec(uint256 value)
func (_SfcV1Contract *SfcV1ContractFilterer) WatchUpdatedBaseRewardPerSec(opts *bind.WatchOpts, sink chan<- *SfcV1ContractUpdatedBaseRewardPerSec) (event.Subscription, error) {

	logs, sub, err := _SfcV1Contract.contract.WatchLogs(opts, "UpdatedBaseRewardPerSec")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcV1ContractUpdatedBaseRewardPerSec)
				if err := _SfcV1Contract.contract.UnpackLog(event, "UpdatedBaseRewardPerSec", log); err != nil {
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
func (_SfcV1Contract *SfcV1ContractFilterer) ParseUpdatedBaseRewardPerSec(log types.Log) (*SfcV1ContractUpdatedBaseRewardPerSec, error) {
	event := new(SfcV1ContractUpdatedBaseRewardPerSec)
	if err := _SfcV1Contract.contract.UnpackLog(event, "UpdatedBaseRewardPerSec", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SfcV1ContractUpdatedDelegationIterator is returned from FilterUpdatedDelegation and is used to iterate over the raw logs and unpacked data for UpdatedDelegation events raised by the SfcV1Contract contract.
type SfcV1ContractUpdatedDelegationIterator struct {
	Event *SfcV1ContractUpdatedDelegation // Event containing the contract specifics and raw log

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
func (it *SfcV1ContractUpdatedDelegationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcV1ContractUpdatedDelegation)
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
		it.Event = new(SfcV1ContractUpdatedDelegation)
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
func (it *SfcV1ContractUpdatedDelegationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcV1ContractUpdatedDelegationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcV1ContractUpdatedDelegation represents a UpdatedDelegation event raised by the SfcV1Contract contract.
type SfcV1ContractUpdatedDelegation struct {
	Delegator   common.Address
	OldStakerID *big.Int
	NewStakerID *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdatedDelegation is a free log retrieval operation binding the contract event 0x19b46b9014e4dc8ca74f505b8921797c6a8a489860217d15b3c7d741637dfcff.
//
// Solidity: event UpdatedDelegation(address indexed delegator, uint256 indexed oldStakerID, uint256 indexed newStakerID, uint256 amount)
func (_SfcV1Contract *SfcV1ContractFilterer) FilterUpdatedDelegation(opts *bind.FilterOpts, delegator []common.Address, oldStakerID []*big.Int, newStakerID []*big.Int) (*SfcV1ContractUpdatedDelegationIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var oldStakerIDRule []interface{}
	for _, oldStakerIDItem := range oldStakerID {
		oldStakerIDRule = append(oldStakerIDRule, oldStakerIDItem)
	}
	var newStakerIDRule []interface{}
	for _, newStakerIDItem := range newStakerID {
		newStakerIDRule = append(newStakerIDRule, newStakerIDItem)
	}

	logs, sub, err := _SfcV1Contract.contract.FilterLogs(opts, "UpdatedDelegation", delegatorRule, oldStakerIDRule, newStakerIDRule)
	if err != nil {
		return nil, err
	}
	return &SfcV1ContractUpdatedDelegationIterator{contract: _SfcV1Contract.contract, event: "UpdatedDelegation", logs: logs, sub: sub}, nil
}

// WatchUpdatedDelegation is a free log subscription operation binding the contract event 0x19b46b9014e4dc8ca74f505b8921797c6a8a489860217d15b3c7d741637dfcff.
//
// Solidity: event UpdatedDelegation(address indexed delegator, uint256 indexed oldStakerID, uint256 indexed newStakerID, uint256 amount)
func (_SfcV1Contract *SfcV1ContractFilterer) WatchUpdatedDelegation(opts *bind.WatchOpts, sink chan<- *SfcV1ContractUpdatedDelegation, delegator []common.Address, oldStakerID []*big.Int, newStakerID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var oldStakerIDRule []interface{}
	for _, oldStakerIDItem := range oldStakerID {
		oldStakerIDRule = append(oldStakerIDRule, oldStakerIDItem)
	}
	var newStakerIDRule []interface{}
	for _, newStakerIDItem := range newStakerID {
		newStakerIDRule = append(newStakerIDRule, newStakerIDItem)
	}

	logs, sub, err := _SfcV1Contract.contract.WatchLogs(opts, "UpdatedDelegation", delegatorRule, oldStakerIDRule, newStakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcV1ContractUpdatedDelegation)
				if err := _SfcV1Contract.contract.UnpackLog(event, "UpdatedDelegation", log); err != nil {
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

// ParseUpdatedDelegation is a log parse operation binding the contract event 0x19b46b9014e4dc8ca74f505b8921797c6a8a489860217d15b3c7d741637dfcff.
//
// Solidity: event UpdatedDelegation(address indexed delegator, uint256 indexed oldStakerID, uint256 indexed newStakerID, uint256 amount)
func (_SfcV1Contract *SfcV1ContractFilterer) ParseUpdatedDelegation(log types.Log) (*SfcV1ContractUpdatedDelegation, error) {
	event := new(SfcV1ContractUpdatedDelegation)
	if err := _SfcV1Contract.contract.UnpackLog(event, "UpdatedDelegation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SfcV1ContractUpdatedGasPowerAllocationRateIterator is returned from FilterUpdatedGasPowerAllocationRate and is used to iterate over the raw logs and unpacked data for UpdatedGasPowerAllocationRate events raised by the SfcV1Contract contract.
type SfcV1ContractUpdatedGasPowerAllocationRateIterator struct {
	Event *SfcV1ContractUpdatedGasPowerAllocationRate // Event containing the contract specifics and raw log

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
func (it *SfcV1ContractUpdatedGasPowerAllocationRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcV1ContractUpdatedGasPowerAllocationRate)
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
		it.Event = new(SfcV1ContractUpdatedGasPowerAllocationRate)
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
func (it *SfcV1ContractUpdatedGasPowerAllocationRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcV1ContractUpdatedGasPowerAllocationRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcV1ContractUpdatedGasPowerAllocationRate represents a UpdatedGasPowerAllocationRate event raised by the SfcV1Contract contract.
type SfcV1ContractUpdatedGasPowerAllocationRate struct {
	Short *big.Int
	Long  *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdatedGasPowerAllocationRate is a free log retrieval operation binding the contract event 0x95ae5488127de4bc98492f4487556e7af9f37eb4b6d5e94f6d849e03ff76cc7c.
//
// Solidity: event UpdatedGasPowerAllocationRate(uint256 short, uint256 long)
func (_SfcV1Contract *SfcV1ContractFilterer) FilterUpdatedGasPowerAllocationRate(opts *bind.FilterOpts) (*SfcV1ContractUpdatedGasPowerAllocationRateIterator, error) {

	logs, sub, err := _SfcV1Contract.contract.FilterLogs(opts, "UpdatedGasPowerAllocationRate")
	if err != nil {
		return nil, err
	}
	return &SfcV1ContractUpdatedGasPowerAllocationRateIterator{contract: _SfcV1Contract.contract, event: "UpdatedGasPowerAllocationRate", logs: logs, sub: sub}, nil
}

// WatchUpdatedGasPowerAllocationRate is a free log subscription operation binding the contract event 0x95ae5488127de4bc98492f4487556e7af9f37eb4b6d5e94f6d849e03ff76cc7c.
//
// Solidity: event UpdatedGasPowerAllocationRate(uint256 short, uint256 long)
func (_SfcV1Contract *SfcV1ContractFilterer) WatchUpdatedGasPowerAllocationRate(opts *bind.WatchOpts, sink chan<- *SfcV1ContractUpdatedGasPowerAllocationRate) (event.Subscription, error) {

	logs, sub, err := _SfcV1Contract.contract.WatchLogs(opts, "UpdatedGasPowerAllocationRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcV1ContractUpdatedGasPowerAllocationRate)
				if err := _SfcV1Contract.contract.UnpackLog(event, "UpdatedGasPowerAllocationRate", log); err != nil {
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
func (_SfcV1Contract *SfcV1ContractFilterer) ParseUpdatedGasPowerAllocationRate(log types.Log) (*SfcV1ContractUpdatedGasPowerAllocationRate, error) {
	event := new(SfcV1ContractUpdatedGasPowerAllocationRate)
	if err := _SfcV1Contract.contract.UnpackLog(event, "UpdatedGasPowerAllocationRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SfcV1ContractUpdatedStakeIterator is returned from FilterUpdatedStake and is used to iterate over the raw logs and unpacked data for UpdatedStake events raised by the SfcV1Contract contract.
type SfcV1ContractUpdatedStakeIterator struct {
	Event *SfcV1ContractUpdatedStake // Event containing the contract specifics and raw log

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
func (it *SfcV1ContractUpdatedStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcV1ContractUpdatedStake)
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
		it.Event = new(SfcV1ContractUpdatedStake)
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
func (it *SfcV1ContractUpdatedStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcV1ContractUpdatedStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcV1ContractUpdatedStake represents a UpdatedStake event raised by the SfcV1Contract contract.
type SfcV1ContractUpdatedStake struct {
	StakerID    *big.Int
	Amount      *big.Int
	DelegatedMe *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdatedStake is a free log retrieval operation binding the contract event 0x509404fa75ce234a1273cf9f7918bcf54e0ef19f2772e4f71b6526606a723b7c.
//
// Solidity: event UpdatedStake(uint256 indexed stakerID, uint256 amount, uint256 delegatedMe)
func (_SfcV1Contract *SfcV1ContractFilterer) FilterUpdatedStake(opts *bind.FilterOpts, stakerID []*big.Int) (*SfcV1ContractUpdatedStakeIterator, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _SfcV1Contract.contract.FilterLogs(opts, "UpdatedStake", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &SfcV1ContractUpdatedStakeIterator{contract: _SfcV1Contract.contract, event: "UpdatedStake", logs: logs, sub: sub}, nil
}

// WatchUpdatedStake is a free log subscription operation binding the contract event 0x509404fa75ce234a1273cf9f7918bcf54e0ef19f2772e4f71b6526606a723b7c.
//
// Solidity: event UpdatedStake(uint256 indexed stakerID, uint256 amount, uint256 delegatedMe)
func (_SfcV1Contract *SfcV1ContractFilterer) WatchUpdatedStake(opts *bind.WatchOpts, sink chan<- *SfcV1ContractUpdatedStake, stakerID []*big.Int) (event.Subscription, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _SfcV1Contract.contract.WatchLogs(opts, "UpdatedStake", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcV1ContractUpdatedStake)
				if err := _SfcV1Contract.contract.UnpackLog(event, "UpdatedStake", log); err != nil {
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

// ParseUpdatedStake is a log parse operation binding the contract event 0x509404fa75ce234a1273cf9f7918bcf54e0ef19f2772e4f71b6526606a723b7c.
//
// Solidity: event UpdatedStake(uint256 indexed stakerID, uint256 amount, uint256 delegatedMe)
func (_SfcV1Contract *SfcV1ContractFilterer) ParseUpdatedStake(log types.Log) (*SfcV1ContractUpdatedStake, error) {
	event := new(SfcV1ContractUpdatedStake)
	if err := _SfcV1Contract.contract.UnpackLog(event, "UpdatedStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SfcV1ContractUpdatedStakerMetadataIterator is returned from FilterUpdatedStakerMetadata and is used to iterate over the raw logs and unpacked data for UpdatedStakerMetadata events raised by the SfcV1Contract contract.
type SfcV1ContractUpdatedStakerMetadataIterator struct {
	Event *SfcV1ContractUpdatedStakerMetadata // Event containing the contract specifics and raw log

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
func (it *SfcV1ContractUpdatedStakerMetadataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcV1ContractUpdatedStakerMetadata)
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
		it.Event = new(SfcV1ContractUpdatedStakerMetadata)
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
func (it *SfcV1ContractUpdatedStakerMetadataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcV1ContractUpdatedStakerMetadataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcV1ContractUpdatedStakerMetadata represents a UpdatedStakerMetadata event raised by the SfcV1Contract contract.
type SfcV1ContractUpdatedStakerMetadata struct {
	StakerID *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdatedStakerMetadata is a free log retrieval operation binding the contract event 0xb7a99a0df6a9e15c2689e6a55811ef76cdb514c67d4a0e37fcb125ada0e3cd83.
//
// Solidity: event UpdatedStakerMetadata(uint256 indexed stakerID)
func (_SfcV1Contract *SfcV1ContractFilterer) FilterUpdatedStakerMetadata(opts *bind.FilterOpts, stakerID []*big.Int) (*SfcV1ContractUpdatedStakerMetadataIterator, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _SfcV1Contract.contract.FilterLogs(opts, "UpdatedStakerMetadata", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &SfcV1ContractUpdatedStakerMetadataIterator{contract: _SfcV1Contract.contract, event: "UpdatedStakerMetadata", logs: logs, sub: sub}, nil
}

// WatchUpdatedStakerMetadata is a free log subscription operation binding the contract event 0xb7a99a0df6a9e15c2689e6a55811ef76cdb514c67d4a0e37fcb125ada0e3cd83.
//
// Solidity: event UpdatedStakerMetadata(uint256 indexed stakerID)
func (_SfcV1Contract *SfcV1ContractFilterer) WatchUpdatedStakerMetadata(opts *bind.WatchOpts, sink chan<- *SfcV1ContractUpdatedStakerMetadata, stakerID []*big.Int) (event.Subscription, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _SfcV1Contract.contract.WatchLogs(opts, "UpdatedStakerMetadata", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcV1ContractUpdatedStakerMetadata)
				if err := _SfcV1Contract.contract.UnpackLog(event, "UpdatedStakerMetadata", log); err != nil {
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
func (_SfcV1Contract *SfcV1ContractFilterer) ParseUpdatedStakerMetadata(log types.Log) (*SfcV1ContractUpdatedStakerMetadata, error) {
	event := new(SfcV1ContractUpdatedStakerMetadata)
	if err := _SfcV1Contract.contract.UnpackLog(event, "UpdatedStakerMetadata", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SfcV1ContractUpdatedStakerSfcAddressIterator is returned from FilterUpdatedStakerSfcAddress and is used to iterate over the raw logs and unpacked data for UpdatedStakerSfcAddress events raised by the SfcV1Contract contract.
type SfcV1ContractUpdatedStakerSfcAddressIterator struct {
	Event *SfcV1ContractUpdatedStakerSfcAddress // Event containing the contract specifics and raw log

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
func (it *SfcV1ContractUpdatedStakerSfcAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcV1ContractUpdatedStakerSfcAddress)
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
		it.Event = new(SfcV1ContractUpdatedStakerSfcAddress)
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
func (it *SfcV1ContractUpdatedStakerSfcAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcV1ContractUpdatedStakerSfcAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcV1ContractUpdatedStakerSfcAddress represents a UpdatedStakerSfcAddress event raised by the SfcV1Contract contract.
type SfcV1ContractUpdatedStakerSfcAddress struct {
	StakerID      *big.Int
	OldSfcAddress common.Address
	NewSfcAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedStakerSfcAddress is a free log retrieval operation binding the contract event 0x7cc102ee500cbca85691c9642080562e8f012b04d27f5b7f389453672b206946.
//
// Solidity: event UpdatedStakerSfcAddress(uint256 indexed stakerID, address indexed oldSfcAddress, address indexed newSfcAddress)
func (_SfcV1Contract *SfcV1ContractFilterer) FilterUpdatedStakerSfcAddress(opts *bind.FilterOpts, stakerID []*big.Int, oldSfcAddress []common.Address, newSfcAddress []common.Address) (*SfcV1ContractUpdatedStakerSfcAddressIterator, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}
	var oldSfcAddressRule []interface{}
	for _, oldSfcAddressItem := range oldSfcAddress {
		oldSfcAddressRule = append(oldSfcAddressRule, oldSfcAddressItem)
	}
	var newSfcAddressRule []interface{}
	for _, newSfcAddressItem := range newSfcAddress {
		newSfcAddressRule = append(newSfcAddressRule, newSfcAddressItem)
	}

	logs, sub, err := _SfcV1Contract.contract.FilterLogs(opts, "UpdatedStakerSfcAddress", stakerIDRule, oldSfcAddressRule, newSfcAddressRule)
	if err != nil {
		return nil, err
	}
	return &SfcV1ContractUpdatedStakerSfcAddressIterator{contract: _SfcV1Contract.contract, event: "UpdatedStakerSfcAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedStakerSfcAddress is a free log subscription operation binding the contract event 0x7cc102ee500cbca85691c9642080562e8f012b04d27f5b7f389453672b206946.
//
// Solidity: event UpdatedStakerSfcAddress(uint256 indexed stakerID, address indexed oldSfcAddress, address indexed newSfcAddress)
func (_SfcV1Contract *SfcV1ContractFilterer) WatchUpdatedStakerSfcAddress(opts *bind.WatchOpts, sink chan<- *SfcV1ContractUpdatedStakerSfcAddress, stakerID []*big.Int, oldSfcAddress []common.Address, newSfcAddress []common.Address) (event.Subscription, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}
	var oldSfcAddressRule []interface{}
	for _, oldSfcAddressItem := range oldSfcAddress {
		oldSfcAddressRule = append(oldSfcAddressRule, oldSfcAddressItem)
	}
	var newSfcAddressRule []interface{}
	for _, newSfcAddressItem := range newSfcAddress {
		newSfcAddressRule = append(newSfcAddressRule, newSfcAddressItem)
	}

	logs, sub, err := _SfcV1Contract.contract.WatchLogs(opts, "UpdatedStakerSfcAddress", stakerIDRule, oldSfcAddressRule, newSfcAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcV1ContractUpdatedStakerSfcAddress)
				if err := _SfcV1Contract.contract.UnpackLog(event, "UpdatedStakerSfcAddress", log); err != nil {
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

// ParseUpdatedStakerSfcAddress is a log parse operation binding the contract event 0x7cc102ee500cbca85691c9642080562e8f012b04d27f5b7f389453672b206946.
//
// Solidity: event UpdatedStakerSfcAddress(uint256 indexed stakerID, address indexed oldSfcAddress, address indexed newSfcAddress)
func (_SfcV1Contract *SfcV1ContractFilterer) ParseUpdatedStakerSfcAddress(log types.Log) (*SfcV1ContractUpdatedStakerSfcAddress, error) {
	event := new(SfcV1ContractUpdatedStakerSfcAddress)
	if err := _SfcV1Contract.contract.UnpackLog(event, "UpdatedStakerSfcAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SfcV1ContractWithdrawnDelegationIterator is returned from FilterWithdrawnDelegation and is used to iterate over the raw logs and unpacked data for WithdrawnDelegation events raised by the SfcV1Contract contract.
type SfcV1ContractWithdrawnDelegationIterator struct {
	Event *SfcV1ContractWithdrawnDelegation // Event containing the contract specifics and raw log

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
func (it *SfcV1ContractWithdrawnDelegationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcV1ContractWithdrawnDelegation)
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
		it.Event = new(SfcV1ContractWithdrawnDelegation)
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
func (it *SfcV1ContractWithdrawnDelegationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcV1ContractWithdrawnDelegationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcV1ContractWithdrawnDelegation represents a WithdrawnDelegation event raised by the SfcV1Contract contract.
type SfcV1ContractWithdrawnDelegation struct {
	Delegator common.Address
	StakerID  *big.Int
	Penalty   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnDelegation is a free log retrieval operation binding the contract event 0x87e86b3710b72c10173ca52c6a9f9cf2df27e77ed177741a8b4feb12bb7a606f.
//
// Solidity: event WithdrawnDelegation(address indexed delegator, uint256 indexed stakerID, uint256 penalty)
func (_SfcV1Contract *SfcV1ContractFilterer) FilterWithdrawnDelegation(opts *bind.FilterOpts, delegator []common.Address, stakerID []*big.Int) (*SfcV1ContractWithdrawnDelegationIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _SfcV1Contract.contract.FilterLogs(opts, "WithdrawnDelegation", delegatorRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &SfcV1ContractWithdrawnDelegationIterator{contract: _SfcV1Contract.contract, event: "WithdrawnDelegation", logs: logs, sub: sub}, nil
}

// WatchWithdrawnDelegation is a free log subscription operation binding the contract event 0x87e86b3710b72c10173ca52c6a9f9cf2df27e77ed177741a8b4feb12bb7a606f.
//
// Solidity: event WithdrawnDelegation(address indexed delegator, uint256 indexed stakerID, uint256 penalty)
func (_SfcV1Contract *SfcV1ContractFilterer) WatchWithdrawnDelegation(opts *bind.WatchOpts, sink chan<- *SfcV1ContractWithdrawnDelegation, delegator []common.Address, stakerID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _SfcV1Contract.contract.WatchLogs(opts, "WithdrawnDelegation", delegatorRule, stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcV1ContractWithdrawnDelegation)
				if err := _SfcV1Contract.contract.UnpackLog(event, "WithdrawnDelegation", log); err != nil {
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
// Solidity: event WithdrawnDelegation(address indexed delegator, uint256 indexed stakerID, uint256 penalty)
func (_SfcV1Contract *SfcV1ContractFilterer) ParseWithdrawnDelegation(log types.Log) (*SfcV1ContractWithdrawnDelegation, error) {
	event := new(SfcV1ContractWithdrawnDelegation)
	if err := _SfcV1Contract.contract.UnpackLog(event, "WithdrawnDelegation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SfcV1ContractWithdrawnStakeIterator is returned from FilterWithdrawnStake and is used to iterate over the raw logs and unpacked data for WithdrawnStake events raised by the SfcV1Contract contract.
type SfcV1ContractWithdrawnStakeIterator struct {
	Event *SfcV1ContractWithdrawnStake // Event containing the contract specifics and raw log

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
func (it *SfcV1ContractWithdrawnStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcV1ContractWithdrawnStake)
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
		it.Event = new(SfcV1ContractWithdrawnStake)
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
func (it *SfcV1ContractWithdrawnStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcV1ContractWithdrawnStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcV1ContractWithdrawnStake represents a WithdrawnStake event raised by the SfcV1Contract contract.
type SfcV1ContractWithdrawnStake struct {
	StakerID *big.Int
	Penalty  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnStake is a free log retrieval operation binding the contract event 0x8c6548258f8f12a9d4b593fa89a223417ed901d4ee9712ba09beb4d56f5262b6.
//
// Solidity: event WithdrawnStake(uint256 indexed stakerID, uint256 penalty)
func (_SfcV1Contract *SfcV1ContractFilterer) FilterWithdrawnStake(opts *bind.FilterOpts, stakerID []*big.Int) (*SfcV1ContractWithdrawnStakeIterator, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _SfcV1Contract.contract.FilterLogs(opts, "WithdrawnStake", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return &SfcV1ContractWithdrawnStakeIterator{contract: _SfcV1Contract.contract, event: "WithdrawnStake", logs: logs, sub: sub}, nil
}

// WatchWithdrawnStake is a free log subscription operation binding the contract event 0x8c6548258f8f12a9d4b593fa89a223417ed901d4ee9712ba09beb4d56f5262b6.
//
// Solidity: event WithdrawnStake(uint256 indexed stakerID, uint256 penalty)
func (_SfcV1Contract *SfcV1ContractFilterer) WatchWithdrawnStake(opts *bind.WatchOpts, sink chan<- *SfcV1ContractWithdrawnStake, stakerID []*big.Int) (event.Subscription, error) {

	var stakerIDRule []interface{}
	for _, stakerIDItem := range stakerID {
		stakerIDRule = append(stakerIDRule, stakerIDItem)
	}

	logs, sub, err := _SfcV1Contract.contract.WatchLogs(opts, "WithdrawnStake", stakerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcV1ContractWithdrawnStake)
				if err := _SfcV1Contract.contract.UnpackLog(event, "WithdrawnStake", log); err != nil {
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
func (_SfcV1Contract *SfcV1ContractFilterer) ParseWithdrawnStake(log types.Log) (*SfcV1ContractWithdrawnStake, error) {
	event := new(SfcV1ContractWithdrawnStake)
	if err := _SfcV1Contract.contract.UnpackLog(event, "WithdrawnStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
