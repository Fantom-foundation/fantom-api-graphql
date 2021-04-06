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
	"math/big"
)

// StoreRewardClaim stores reward claim record in the persistent repository.
func (p *proxy) StoreRewardClaim(rc *types.RewardClaim) error {
	return p.db.AddRewardClaim(rc)
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
	ld.log.Debugf("%s claimed %d in stake to  #%d", addr.String(), amo.Uint64(), valID.ToInt().Uint64())

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
	go ld.repo.UpdateDelegationActiveAmount(&addr, valID)
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
