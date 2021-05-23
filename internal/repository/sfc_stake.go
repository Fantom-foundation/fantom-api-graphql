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

// makeAdHocDelegation creates a new delegation with unknown initial conditions in case of an attempt
// to update a non-existent delegation.
func (ld *logsDispatcher) makeAdHocDelegation(log *retypes.Log, addr *common.Address, valID *hexutil.Big, amo *big.Int) error {
	// get the block
	block, err := ld.repo.BlockByNumber((*hexutil.Uint64)(&log.BlockNumber))
	if err != nil {
		return err
	}

	// get staker address
	val, err := R().ValidatorAddress(valID)
	if err != nil {
		return err
	}

	// do the insert
	ld.log.Noticef("creating ad-hoc delegation of %s to %d", addr.String(), valID.ToInt().Uint64())
	return ld.repo.StoreDelegation(&types.Delegation{
		Transaction:     log.TxHash,
		Address:         *addr,
		ToStakerId:      valID,
		ToStakerAddress: *val,
		AmountDelegated: (*hexutil.Big)(amo),
		AmountStaked:    (*hexutil.Big)(amo),
		CreatedTime:     block.TimeStamp,
	})
}

// handleSfcCreatedStake handles a new stake event from SFC v1 and SFC v2 contract.
// event CreatedStake(uint256 indexed stakerID, address indexed dagSfcAddress, uint256 amount)
func handleSfcCreatedStake(log *retypes.Log, ld *logsDispatcher) {
	handleDelegationLog(
		hexutil.Uint64(log.BlockNumber),
		&log.TxHash,
		new(big.Int).SetBytes(log.Topics[1].Bytes()),
		common.BytesToAddress(log.Topics[2].Bytes()),
		new(big.Int).SetBytes(log.Data),
		ld,
	)
}

// handleSfc1IncreasedStake handles a stake increase event from SFC v1 and SFC v2 contract.
// event IncreasedStake(uint256 indexed stakerID, uint256 newAmount, uint256 diff)
func handleSfc1IncreasedStake(log *retypes.Log, ld *logsDispatcher) {
	// get the validator ID
	valID := (*hexutil.Big)(new(big.Int).SetBytes(log.Topics[1].Bytes()))

	// get the validator address
	addr, err := R().ValidatorAddress(valID)
	if err != nil {
		ld.log.Errorf("unknown validator #%d; %s", valID.ToInt().Uint64(), err.Error())
		return
	}

	// update the balance
	if err := ld.repo.UpdateDelegationBalance(addr, valID, func(amo *big.Int) error {
		return ld.makeAdHocDelegation(log, addr, valID, amo)
	}); err != nil {
		ld.log.Errorf("failed to update delegation; %s", err.Error())
	}
}

// handleSfc1WithdrawnStake handles a withdrawal request finalization event from SFC1.
// event WithdrawnStake(uint256 indexed stakerID, uint256 penalty)
func handleSfc1WithdrawnStake(log *retypes.Log, ld *logsDispatcher) {
	// extract the basic info about the request
	valID := (*hexutil.Big)(new(big.Int).SetBytes(log.Topics[1].Bytes()))

	// get the validator address (we handle the stake as self-delegation in SFC3 context)
	addr, err := R().ValidatorAddress(valID)
	if err != nil {
		ld.log.Errorf("unknown validator #%d; %s", valID.ToInt().Uint64(), err.Error())
		return
	}

	// check active amount on the delegation
	if err := ld.repo.UpdateDelegationBalance(addr, valID, func(amo *big.Int) error {
		return ld.makeAdHocDelegation(log, addr, valID, amo)
	}); err != nil {
		ld.log.Errorf("failed to update delegation; %s", err.Error())
	}
}

// handleSfc1DeactivatedStake handles a validator stake deactivation event.
// DeactivatedStake(uint256 indexed stakerID)
func handleSfc1DeactivatedStake(log *retypes.Log, ld *logsDispatcher) {
	// re-use similar function
	handleSfc1WithdrawnStake(log, ld)
}

// handleSfc1UpdatedStake handles stake update event.
// UpdatedStake(uint256 indexed stakerID, uint256 amount, uint256 delegatedMe)
func handleSfc1UpdatedStake(log *retypes.Log, ld *logsDispatcher) {
	// re-use similar function
	handleSfc1WithdrawnStake(log, ld)
}
