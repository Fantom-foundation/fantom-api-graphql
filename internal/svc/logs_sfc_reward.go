// Package svc implements blockchain data processing services.
package svc

import (
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// handleSfcRewardClaim handles a rewards claim event.
func handleSfcRewardClaim(lr *types.LogRecord, addr common.Address, valID *hexutil.Big, amo *big.Int, isRestake bool) {
	// debug the event
	log.Debugf("%s claimed %d in stake to #%d", addr.String(), amo.Uint64(), valID.ToInt().Uint64())

	// add the rewards claim into the repository
	if err := repo.StoreRewardClaim(&types.RewardClaim{
		Delegator:     addr,
		ToValidatorId: *valID,
		Claimed:       lr.Block.TimeStamp,
		ClaimTrx:      lr.TxHash,
		Amount:        (hexutil.Big)(*amo),
		IsDelegated:   isRestake,
	}); err != nil {
		log.Criticalf("can not store rewards claim; %s", err.Error())
		return
	}

	// check active amount on the delegation
	if err := repo.UpdateDelegationBalance(&addr, valID, func(amo *big.Int) error {
		return makeAdHocDelegation(lr, &addr, valID, amo)
	}); err != nil {
		log.Errorf("failed to update delegation; %s", err.Error())
	}
}

// handleSfcCommonRewardClaim handles the common reward claim on SFC contract.
func handleSfcCommonRewardClaim(lr *types.LogRecord, isRestake bool) {
	// sanity check for data (3x uint256 = 3x32 bytes = 96 bytes)
	if len(lr.Data) != 96 {
		log.Criticalf("%s lr invalid data length; expected 96 bytes, given %d bytes", lr.TxHash.String(), len(lr.Data))
		return
	}

	// extract the basic info about the request
	addr := common.BytesToAddress(lr.Topics[1].Bytes())
	valID := (*hexutil.Big)(new(big.Int).SetBytes(lr.Topics[2].Bytes()))

	// collect values for each reward section
	amoA := new(big.Int).Add(
		new(big.Int).SetBytes(lr.Data[:32]),
		new(big.Int).SetBytes(lr.Data[32:64]),
	)
	amo := new(big.Int).Add(amoA, new(big.Int).SetBytes(lr.Data[64:]))

	// do the handling
	handleSfcRewardClaim(lr, addr, valID, amo, isRestake)
}

// handleSfcRestakeRewards handles a rewards re-stake event.
// event RestakedRewards(address indexed delegator, uint256 indexed toValidatorID, uint256 lockupExtraReward, uint256 lockupBaseReward, uint256 unlockedReward)
func handleSfcRestakeRewards(lr *types.LogRecord) {
	handleSfcCommonRewardClaim(lr, true)

	// adjust locked stake amount (data should be 3x 32 bytes = 96 bytes)
	if len(lr.Data) != 96 {
		log.Criticalf("%s lr invalid data length; expected 96 bytes, given %d bytes", lr.TxHash.String(), len(lr.Data))
		return
	}

	extra := new(big.Int).SetBytes(lr.Data[:32])
	base := new(big.Int).SetBytes(lr.Data[32:64])

	err := repo.AdjustLockedDelegation(
		common.BytesToAddress(lr.Topics[1].Bytes()),
		new(big.Int).SetBytes(lr.Topics[2].Bytes()).Int64(),
		types.LockedDelegationValue(new(big.Int).Add(base, extra)),
	)
	if err != nil {
		log.Errorf("could not adjust locked delegation; %s", err.Error())
	}
}

// handleSfcClaimedRewards handles a rewards re-stake event.
// event ClaimedRewards(address indexed delegator, uint256 indexed toValidatorID, uint256 lockupExtraReward, uint256 lockupBaseReward, uint256 unlockedReward)
func handleSfcClaimedRewards(lr *types.LogRecord) {
	handleSfcCommonRewardClaim(lr, false)
}

// handleSfc1ClaimedDelegationReward handles a delegation reward claim from SFC1.
// event ClaimedDelegationReward(address indexed from, uint256 indexed stakerID, uint256 reward, uint256 fromEpoch, uint256 untilEpoch)
func handleSfc1ClaimedDelegationReward(lr *types.LogRecord) {
	// sanity check for data (3x uint256 = 3x32 bytes = 96 bytes)
	if len(lr.Data) != 96 {
		log.Criticalf("%s lr invalid data length; expected 96 bytes, given %d bytes", lr.TxHash.String(), len(lr.Data))
		return
	}

	// extract the basic info about the request
	addr := common.BytesToAddress(lr.Topics[1].Bytes())
	valID := (*hexutil.Big)(new(big.Int).SetBytes(lr.Topics[2].Bytes()))
	amo := new(big.Int).SetBytes(lr.Data[:32])

	// do the handling
	handleSfcRewardClaim(lr, addr, valID, amo, false)
}

// handleSfc1UnstashedReward handles rewards un-stash request event.
// event UnstashedRewards(address indexed auth, address indexed receiver, uint256 rewards)
func handleSfc1UnstashedReward(lr *types.LogRecord) {
	// sanity check for data (1x uint256 = 1x32 bytes)
	if len(lr.Data) != 32 {
		log.Criticalf("%s lr invalid data length; expected 32 bytes, given %d bytes", lr.TxHash.String(), len(lr.Data))
		return
	}

	// extract the basic info about the request
	addr := common.BytesToAddress(lr.Topics[1].Bytes())
	amo := new(big.Int).SetBytes(lr.Data[:32])

	// debug the event
	log.Debugf("%s un-stashed %d from SFC", addr.String(), amo.Uint64())

	// do the handling; staker ID for the stashed amount is not known here
	handleSfcRewardClaim(lr, addr, (*hexutil.Big)(new(big.Int)), amo, false)
}

// handleSfc1ClaimedValidatorReward handles validator reward claim.
// event ClaimedValidatorReward(uint256 indexed stakerID, uint256 reward, uint256 fromEpoch, uint256 untilEpoch)
func handleSfc1ClaimedValidatorReward(lr *types.LogRecord) {
	// sanity check for data (3x uint256 = 3x32 bytes = 96 bytes)
	if len(lr.Data) != 96 {
		log.Criticalf("%s lr invalid data length; expected 96 bytes, given %d bytes", lr.TxHash.String(), len(lr.Data))
		return
	}

	// get the validator and amount
	valID := (*hexutil.Big)(new(big.Int).SetBytes(lr.Topics[1].Bytes()))
	amo := new(big.Int).SetBytes(lr.Data[:32])

	// get the validator address since this a self stake
	addr, err := repo.ValidatorAddress(valID)
	if err != nil {
		log.Errorf("validator #%d not found; %s", valID.ToInt().Uint64(), err.Error())
		return
	}

	// do the handling
	handleSfcRewardClaim(lr, *addr, valID, amo, false)
}
