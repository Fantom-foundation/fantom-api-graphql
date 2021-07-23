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
	retypes "github.com/ethereum/go-ethereum/core/types"
	"go.mongodb.org/mongo-driver/bson"
	"math/big"
	"time"
)

// StoreRewardClaim stores reward claim record in the persistent repository.
func (p *proxy) StoreRewardClaim(rc *types.RewardClaim) error {
	return p.db.AddRewardClaim(rc)
}

// RewardClaims provides a list of reward claims for the given delegation and/or filter.
func (p *proxy) RewardClaims(adr *common.Address, valID *big.Int, cursor *string, count int32) (*types.RewardClaimsList, error) {
	// prep the filter
	fi := bson.D{}

	// add delegator address to the filter
	if adr != nil {
		fi = append(fi, bson.E{
			Key:   types.FiRewardClaimAddress,
			Value: adr.String(),
		})
	}

	// add validator ID to the filter
	if valID != nil {
		fi = append(fi, bson.E{
			Key:   types.FiRewardClaimToValidator,
			Value: (*hexutil.Big)(valID).String(),
		})
	}
	return p.db.RewardClaims(cursor, count, &fi)
}

// RewardsClaimed returns sum of all claimed rewards for the given delegator address and validator ID.
func (p *proxy) RewardsClaimed(adr *common.Address, valId *big.Int, since *int64, until *int64) (*big.Int, error) {
	// prep the filter
	fi := bson.D{}

	// filter by delegator address
	if adr != nil {
		fi = append(fi, bson.E{
			Key:   types.FiRewardClaimAddress,
			Value: adr.String(),
		})
	}

	// filter by validator ID
	if valId != nil {
		fi = append(fi, bson.E{
			Key:   types.FiRewardClaimToValidator,
			Value: (*hexutil.Big)(valId).String(),
		})
	}

	// starting time stamp provided
	if since != nil {
		fi = append(fi, bson.E{
			Key:   types.FiRewardClaimedTimeStamp,
			Value: bson.D{{Key: "$gte", Value: time.Unix(*since, 0)}},
		})
	}

	// ending time stamp provided
	if until != nil {
		fi = append(fi, bson.E{
			Key:   types.FiRewardClaimedTimeStamp,
			Value: bson.D{{Key: "$lte", Value: time.Unix(*until, 0)}},
		})
	}
	return p.db.RewardsSumValue(&fi)
}

// handleSfcRewardClaim handles a rewards claim event.
func handleSfcRewardClaim(log *retypes.Log, ld *logsDispatcher, isRestake bool) {
	// sanity check for data (3x uint256 = 3x32 bytes = 96 bytes)
	if len(log.Data) != 96 {
		ld.log.Criticalf("%s log invalid data length; expected 96 bytes, given %d bytes", log.TxHash.String(), len(log.Data))
		return
	}

	// get the block
	blk := hexutil.Uint64(log.BlockNumber)
	block, err := ld.repo.BlockByNumber(&blk)
	if err != nil {
		ld.log.Errorf("can not decode delegation log record; %s", err.Error())
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

	// debug the event
	ld.log.Debugf("%s claimed %d in stake to #%d", addr.String(), amo.Uint64(), valID.ToInt().Uint64())

	// add the rewards claim into the repository
	if err := ld.repo.StoreRewardClaim(&types.RewardClaim{
		Delegator:     addr,
		ToValidatorId: *valID,
		Claimed:       block.TimeStamp,
		ClaimTrx:      log.TxHash,
		Amount:        (hexutil.Big)(*amo),
		IsDelegated:   isRestake,
	}); err != nil {
		ld.log.Criticalf("can not store rewards claim; %s", err.Error())
		return
	}

	// check active amount on the delegation
	if err := ld.repo.UpdateDelegationBalance(&addr, valID, func(amo *big.Int) error {
		return ld.makeAdHocDelegation(log, &addr, valID, amo)
	}); err != nil {
		ld.log.Errorf("failed to update delegation; %s", err.Error())
	}
}

// handleSfcRestakeRewards handles a rewards re-stake event.
// event RestakedRewards(address indexed delegator, uint256 indexed toValidatorID, uint256 lockupExtraReward, uint256 lockupBaseReward, uint256 unlockedReward)
func handleSfcRestakeRewards(log *retypes.Log, ld *logsDispatcher) {
	handleSfcRewardClaim(log, ld, true)
}

// handleSfcClaimedRewards handles a rewards re-stake event.
// event ClaimedRewards(address indexed delegator, uint256 indexed toValidatorID, uint256 lockupExtraReward, uint256 lockupBaseReward, uint256 unlockedReward)
func handleSfcClaimedRewards(log *retypes.Log, ld *logsDispatcher) {
	handleSfcRewardClaim(log, ld, false)
}

// handleSfc1ClaimedDelegationReward handles a delegation reward claim from SFC1.
// event ClaimedDelegationReward(address indexed from, uint256 indexed stakerID, uint256 reward, uint256 fromEpoch, uint256 untilEpoch)
func handleSfc1ClaimedDelegationReward(log *retypes.Log, ld *logsDispatcher) {
	// sanity check for data (3x uint256 = 3x32 bytes = 96 bytes)
	if len(log.Data) != 96 {
		ld.log.Criticalf("%s log invalid data length; expected 96 bytes, given %d bytes", log.TxHash.String(), len(log.Data))
		return
	}

	// get the block
	blk := hexutil.Uint64(log.BlockNumber)
	block, err := ld.repo.BlockByNumber(&blk)
	if err != nil {
		ld.log.Errorf("can not decode delegation log record; %s", err.Error())
		return
	}

	// extract the basic info about the request
	addr := common.BytesToAddress(log.Topics[1].Bytes())
	valID := (*hexutil.Big)(new(big.Int).SetBytes(log.Topics[2].Bytes()))
	amo := new(big.Int).SetBytes(log.Data[:32])

	// debug the event
	ld.log.Debugf("%s claimed %d in stake to #%d", addr.String(), amo.Uint64(), valID.ToInt().Uint64())

	// add the rewards claim into the repository
	if err := ld.repo.StoreRewardClaim(&types.RewardClaim{
		Delegator:     addr,
		ToValidatorId: *valID,
		Claimed:       block.TimeStamp,
		ClaimTrx:      log.TxHash,
		Amount:        (hexutil.Big)(*amo),
		IsDelegated:   false,
	}); err != nil {
		ld.log.Criticalf("can not store rewards claim; %s", err.Error())
		return
	}

	// check active amount on the delegation
	if err := ld.repo.UpdateDelegationBalance(&addr, valID, func(amo *big.Int) error {
		return ld.makeAdHocDelegation(log, &addr, valID, amo)
	}); err != nil {
		ld.log.Errorf("failed to update delegation; %s", err.Error())
	}
}

// handleSfc1UnstashedReward handles rewards un-stash request event.
// event UnstashedRewards(address indexed auth, address indexed receiver, uint256 rewards)
func handleSfc1UnstashedReward(log *retypes.Log, ld *logsDispatcher) {
	// sanity check for data (1x uint256 = 1x32 bytes)
	if len(log.Data) != 32 {
		ld.log.Criticalf("%s log invalid data length; expected 32 bytes, given %d bytes", log.TxHash.String(), len(log.Data))
		return
	}

	// get the block
	blk := hexutil.Uint64(log.BlockNumber)
	block, err := ld.repo.BlockByNumber(&blk)
	if err != nil {
		ld.log.Errorf("can not decode delegation log record; %s", err.Error())
		return
	}

	// extract the basic info about the request
	addr := common.BytesToAddress(log.Topics[1].Bytes())
	amo := new(big.Int).SetBytes(log.Data[:32])

	// debug the event
	ld.log.Debugf("%s un-stashed %d from SFC", addr.String(), amo.Uint64())

	// add the rewards claim into the repository
	if err := ld.repo.StoreRewardClaim(&types.RewardClaim{
		Delegator:     addr,
		ToValidatorId: (hexutil.Big)(*new(big.Int)),
		Claimed:       block.TimeStamp,
		ClaimTrx:      log.TxHash,
		Amount:        (hexutil.Big)(*amo),
		IsDelegated:   false,
	}); err != nil {
		ld.log.Criticalf("can not store rewards un-stash; %s", err.Error())
		return
	}
}

// handleSfc1ClaimedValidatorReward handles validator reward claim.
// event ClaimedValidatorReward(uint256 indexed stakerID, uint256 reward, uint256 fromEpoch, uint256 untilEpoch)
func handleSfc1ClaimedValidatorReward(log *retypes.Log, ld *logsDispatcher) {
	// sanity check for data (3x uint256 = 3x32 bytes = 96 bytes)
	if len(log.Data) != 96 {
		ld.log.Criticalf("%s log invalid data length; expected 96 bytes, given %d bytes", log.TxHash.String(), len(log.Data))
		return
	}

	// get the block
	blk := hexutil.Uint64(log.BlockNumber)
	block, err := ld.repo.BlockByNumber(&blk)
	if err != nil {
		ld.log.Errorf("can not decode delegation log record; %s", err.Error())
		return
	}

	// get the validator and amount
	valID := (*hexutil.Big)(new(big.Int).SetBytes(log.Topics[1].Bytes()))
	amo := new(big.Int).SetBytes(log.Data[:32])

	// get the validator address since this a self stake
	val, err := ld.repo.ValidatorAddress(valID)
	if err != nil {
		ld.log.Errorf("validator #%d not found; %s", valID.ToInt().Uint64(), err.Error())
		return
	}

	// debug the event
	ld.log.Debugf("%s claimed %d in stake to #%d", val.String(), amo.Uint64(), valID.ToInt().Uint64())

	// add the rewards claim into the repository
	if err := ld.repo.StoreRewardClaim(&types.RewardClaim{
		Delegator:     *val,
		ToValidatorId: *valID,
		Claimed:       block.TimeStamp,
		ClaimTrx:      log.TxHash,
		Amount:        (hexutil.Big)(*amo),
		IsDelegated:   false,
	}); err != nil {
		ld.log.Criticalf("can not store rewards claim; %s", err.Error())
		return
	}

	// check active amount on the delegation
	if err := ld.repo.UpdateDelegationBalance(val, valID, func(amo *big.Int) error {
		return ld.makeAdHocDelegation(log, val, valID, amo)
	}); err != nil {
		ld.log.Errorf("failed to update delegation; %s", err.Error())
	}
}
