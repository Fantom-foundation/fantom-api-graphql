/*
Package repository implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Opera/Lachesis full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import (
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// handleSfcRewardClaim handles a rewards claim event.
func handleSfcRewardClaim(log *types.LogRecord, addr common.Address, valID *hexutil.Big, amo *big.Int, isRestake bool) {
	// debug the event
	R().Log().Debugf("%s claimed %d in stake to #%d", addr.String(), amo.Uint64(), valID.ToInt().Uint64())

	// add the rewards claim into the repository
	if err := R().StoreRewardClaim(&types.RewardClaim{
		Delegator:     addr,
		ToValidatorId: *valID,
		Claimed:       log.Block.TimeStamp,
		ClaimTrx:      log.TxHash,
		Amount:        (hexutil.Big)(*amo),
		IsDelegated:   isRestake,
	}); err != nil {
		R().Log().Criticalf("can not store rewards claim; %s", err.Error())
		return
	}

	// check active amount on the delegation
	if err := R().UpdateDelegationBalance(&addr, valID, func(amo *big.Int) error {
		return makeAdHocDelegation(log, &addr, valID, amo)
	}); err != nil {
		R().Log().Errorf("failed to update delegation; %s", err.Error())
	}
}

// handleSfcCommonRewardClaim handles the common reward claim on SFC contract.
func handleSfcCommonRewardClaim(log *types.LogRecord, isRestake bool) {
	// sanity check for data (3x uint256 = 3x32 bytes = 96 bytes)
	if len(log.Data) != 96 {
		R().Log().Criticalf("%s log invalid data length; expected 96 bytes, given %d bytes", log.TxHash.String(), len(log.Data))
		return
	}

	// extract the basic info about the request
	addr := common.BytesToAddress(log.Topics[1].Bytes())
	valID := (*hexutil.Big)(new(big.Int).SetBytes(log.Topics[2].Bytes()))

	// collect values for each reward section
	amoA := new(big.Int).Add(
		new(big.Int).SetBytes(log.Data[:32]),
		new(big.Int).SetBytes(log.Data[32:64]),
	)
	amo := new(big.Int).Add(amoA, new(big.Int).SetBytes(log.Data[64:]))

	// do the handling
	handleSfcRewardClaim(log, addr, valID, amo, isRestake)
}

// handleSfcRestakeRewards handles a rewards re-stake event.
// event RestakedRewards(address indexed delegator, uint256 indexed toValidatorID, uint256 lockupExtraReward, uint256 lockupBaseReward, uint256 unlockedReward)
func handleSfcRestakeRewards(log *types.LogRecord) {
	handleSfcCommonRewardClaim(log, true)
}

// handleSfcClaimedRewards handles a rewards re-stake event.
// event ClaimedRewards(address indexed delegator, uint256 indexed toValidatorID, uint256 lockupExtraReward, uint256 lockupBaseReward, uint256 unlockedReward)
func handleSfcClaimedRewards(log *types.LogRecord) {
	handleSfcCommonRewardClaim(log, false)
}

// handleSfc1ClaimedDelegationReward handles a delegation reward claim from SFC1.
// event ClaimedDelegationReward(address indexed from, uint256 indexed stakerID, uint256 reward, uint256 fromEpoch, uint256 untilEpoch)
func handleSfc1ClaimedDelegationReward(log *types.LogRecord) {
	// sanity check for data (3x uint256 = 3x32 bytes = 96 bytes)
	if len(log.Data) != 96 {
		R().Log().Criticalf("%s log invalid data length; expected 96 bytes, given %d bytes", log.TxHash.String(), len(log.Data))
		return
	}

	// extract the basic info about the request
	addr := common.BytesToAddress(log.Topics[1].Bytes())
	valID := (*hexutil.Big)(new(big.Int).SetBytes(log.Topics[2].Bytes()))
	amo := new(big.Int).SetBytes(log.Data[:32])

	// do the handling
	handleSfcRewardClaim(log, addr, valID, amo, false)
}

// handleSfc1UnstashedReward handles rewards un-stash request event.
// event UnstashedRewards(address indexed auth, address indexed receiver, uint256 rewards)
func handleSfc1UnstashedReward(log *types.LogRecord) {
	// sanity check for data (1x uint256 = 1x32 bytes)
	if len(log.Data) != 32 {
		R().Log().Criticalf("%s log invalid data length; expected 32 bytes, given %d bytes", log.TxHash.String(), len(log.Data))
		return
	}

	// extract the basic info about the request
	addr := common.BytesToAddress(log.Topics[1].Bytes())
	amo := new(big.Int).SetBytes(log.Data[:32])

	// debug the event
	R().Log().Debugf("%s un-stashed %d from SFC", addr.String(), amo.Uint64())

	// do the handling; staker ID for the stashed amount is not known here
	handleSfcRewardClaim(log, addr, (*hexutil.Big)(new(big.Int)), amo, false)
}

// handleSfc1ClaimedValidatorReward handles validator reward claim.
// event ClaimedValidatorReward(uint256 indexed stakerID, uint256 reward, uint256 fromEpoch, uint256 untilEpoch)
func handleSfc1ClaimedValidatorReward(log *types.LogRecord) {
	// sanity check for data (3x uint256 = 3x32 bytes = 96 bytes)
	if len(log.Data) != 96 {
		R().Log().Criticalf("%s log invalid data length; expected 96 bytes, given %d bytes", log.TxHash.String(), len(log.Data))
		return
	}

	// get the validator and amount
	valID := (*hexutil.Big)(new(big.Int).SetBytes(log.Topics[1].Bytes()))
	amo := new(big.Int).SetBytes(log.Data[:32])

	// get the validator address since this a self stake
	addr, err := R().ValidatorAddress(valID)
	if err != nil {
		R().Log().Errorf("validator #%d not found; %s", valID.ToInt().Uint64(), err.Error())
		return
	}

	// do the handling
	handleSfcRewardClaim(log, *addr, valID, amo, false)
}
