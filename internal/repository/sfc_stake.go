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
	handleNewDelegation(
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
	// get the validator address
	valID := (*hexutil.Big)(new(big.Int).SetBytes(log.Topics[1].Bytes()))
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
	// sanity check for data (1 uint256 = 32 bytes)
	if len(log.Data) != 32 {
		ld.log.Criticalf("%s log invalid data length; expected 32 bytes, %d bytes given, %d topics given", log.TxHash.String(), len(log.Data), len(log.Topics))
		return
	}

	// get the validator address (we handle the stake as self-delegation in SFC3 context)
	valID := new(big.Int).SetBytes(log.Topics[1].Bytes())
	addr, err := R().ValidatorAddress((*hexutil.Big)(valID))
	if err != nil {
		ld.log.Errorf("unknown validator #%d; %s", valID.Uint64(), err.Error())
		return
	}

	// handle request closure
	handleFinishedWithdrawRequest(*addr, valID, new(big.Int), new(big.Int).SetBytes(log.Data[:]), log, ld)
}

// handleSfc1DeactivatedStake handles a validator stake deactivation event
// so the stake could be withdrawn later, this basically creates a withdraw request on the stake
// without request ID.
// DeactivatedStake(uint256 indexed stakerID)
func handleSfc1DeactivatedStake(log *retypes.Log, ld *logsDispatcher) {
	// sanity check for data
	if len(log.Data) != 0 {
		ld.log.Criticalf("%s log invalid data length; expected 0 bytes, %d bytes given, %d topics given", log.TxHash.String(), len(log.Data), len(log.Topics))
		return
	}

	// get the validator address (we handle the stake as self-delegation in SFC3 context)
	valID := new(big.Int).SetBytes(log.Topics[1].Bytes())
	addr, err := R().ValidatorAddress((*hexutil.Big)(valID))
	if err != nil {
		ld.log.Errorf("unknown validator #%d; %s", valID.Uint64(), err.Error())
		return
	}

	// create withdraw request
	zero := new(big.Int)
	handleNewWithdrawRequest(
		types.WithdrawTypeDeactivatedVal,
		/* address */ *addr,
		/* valID */ valID,
		/* reqID*/ zero,
		/* amount */ zero,
		log,
		ld,
	)
}

// handleSfc1UpdatedStake handles stake update event.
// UpdatedStake(uint256 indexed stakerID, uint256 amount, uint256 delegatedMe)
func handleSfc1UpdatedStake(log *retypes.Log, ld *logsDispatcher) {
	// get the validator address (we handle the stake as self-delegation in SFC3 context)
	valID := (*hexutil.Big)(new(big.Int).SetBytes(log.Topics[1].Bytes()))
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
