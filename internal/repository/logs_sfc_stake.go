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

// handleSfcCreatedStake handles a new stake event from SFC v1 and SFC v2 contract.
// event CreatedStake(uint256 indexed stakerID, address indexed dagSfcAddress, uint256 amount)
func handleSfcCreatedStake(log *types.LogRecord) {
	handleNewDelegation(
		log,
		new(big.Int).SetBytes(log.Topics[1].Bytes()),
		common.BytesToAddress(log.Topics[2].Bytes()),
		new(big.Int).SetBytes(log.Data),
	)
}

// handleSfc1IncreasedStake handles a stake increase event from SFC v1 and SFC v2 contract.
// event IncreasedStake(uint256 indexed stakerID, uint256 newAmount, uint256 diff)
func handleSfc1IncreasedStake(log *types.LogRecord) {
	// get the validator address
	valID := (*hexutil.Big)(new(big.Int).SetBytes(log.Topics[1].Bytes()))
	addr, err := R().ValidatorAddress(valID)
	if err != nil {
		R().Log().Errorf("unknown validator #%d; %s", valID.ToInt().Uint64(), err.Error())
		return
	}

	// update the balance
	if err := R().UpdateDelegationBalance(addr, valID, func(amo *big.Int) error {
		return makeAdHocDelegation(log, addr, valID, amo)
	}); err != nil {
		R().Log().Errorf("failed to update delegation; %s", err.Error())
	}
}

// handleSfc1WithdrawnStake handles a withdrawal request finalization event from SFC1.
// event WithdrawnStake(uint256 indexed stakerID, uint256 penalty)
func handleSfc1WithdrawnStake(log *types.LogRecord) {
	// sanity check for data (1 uint256 = 32 bytes)
	if len(log.Data) != 32 {
		R().Log().Criticalf("%s log invalid data length; expected 32 bytes, %d bytes given, %d topics given", log.TxHash.String(), len(log.Data), len(log.Topics))
		return
	}

	// get the validator address (we handle the stake as self-delegation in SFC3 context)
	valID := new(big.Int).SetBytes(log.Topics[1].Bytes())
	addr, err := R().ValidatorAddress((*hexutil.Big)(valID))
	if err != nil {
		R().Log().Errorf("unknown validator #%d; %s", valID.Uint64(), err.Error())
		return
	}

	// handle request closure
	handleFinishedWithdrawRequest(*addr, valID, new(big.Int), new(big.Int).SetBytes(log.Data[:]), log)
}

// handleSfc1DeactivatedStake handles a validator stake deactivation event
// so the stake could be withdrawn later, this basically creates a withdraw request on the stake
// without request ID.
// DeactivatedStake(uint256 indexed stakerID)
func handleSfc1DeactivatedStake(log *types.LogRecord) {
	// sanity check for data
	if len(log.Data) != 0 {
		R().Log().Criticalf("%s log invalid data length; expected 0 bytes, %d bytes given, %d topics given", log.TxHash.String(), len(log.Data), len(log.Topics))
		return
	}

	// get the validator address (we handle the stake as self-delegation in SFC3 context)
	valID := new(big.Int).SetBytes(log.Topics[1].Bytes())
	addr, err := R().ValidatorAddress((*hexutil.Big)(valID))
	if err != nil {
		R().Log().Errorf("unknown validator #%d; %s", valID.Uint64(), err.Error())
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
	)
}

// handleSfc1UpdatedStake handles stake update event; the event is emitted
// when a stake and/or delegation amount changes on a validator  account.
// UpdatedStake(uint256 indexed stakerID, uint256 amount, uint256 delegatedMe)
func handleSfc1UpdatedStake(log *types.LogRecord) {
	// get the validator address (we handle the stake as self-delegation in SFC3 context)
	valID := (*hexutil.Big)(new(big.Int).SetBytes(log.Topics[1].Bytes()))
	addr, err := R().ValidatorAddress(valID)
	if err != nil {
		R().Log().Errorf("unknown validator #%d; %s", valID.ToInt().Uint64(), err.Error())
		return
	}

	// update the balance
	if err := R().UpdateDelegationBalance(addr, valID, func(amo *big.Int) error {
		R().Log().Criticalf("expected validator %d stake not found at %s", valID.ToInt().Uint64(), log.TxHash.String())
		return makeAdHocDelegation(log, addr, valID, amo)
	}); err != nil {
		R().Log().Errorf("failed to update delegation; %s", err.Error())
	}
}
