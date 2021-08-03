// Package svc implements blockchain data processing services.
package svc

import (
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// handleSfcCreatedStake handles a new stake event from SFC v1 and SFC v2 contract.
// event CreatedStake(uint256 indexed stakerID, address indexed dagSfcAddress, uint256 amount)
func handleSfcCreatedStake(lr *types.LogRecord) {
	handleNewDelegation(
		lr,
		new(big.Int).SetBytes(lr.Topics[1].Bytes()),
		common.BytesToAddress(lr.Topics[2].Bytes()),
		new(big.Int).SetBytes(lr.Data),
	)
}

// handleSfc1IncreasedStake handles a stake increase event from SFC v1 and SFC v2 contract.
// event IncreasedStake(uint256 indexed stakerID, uint256 newAmount, uint256 diff)
func handleSfc1IncreasedStake(lr *types.LogRecord) {
	// get the validator address
	valID := (*hexutil.Big)(new(big.Int).SetBytes(lr.Topics[1].Bytes()))
	addr, err := repo.ValidatorAddress(valID)
	if err != nil {
		log.Errorf("unknown validator #%d; %s", valID.ToInt().Uint64(), err.Error())
		return
	}

	// update the balance
	if err := repo.UpdateDelegationBalance(addr, valID, func(amo *big.Int) error {
		return makeAdHocDelegation(lr, addr, valID, amo)
	}); err != nil {
		log.Errorf("failed to update delegation; %s", err.Error())
	}
}

// handleSfc1WithdrawnStake handles a withdrawal request finalization event from SFC1.
// event WithdrawnStake(uint256 indexed stakerID, uint256 penalty)
func handleSfc1WithdrawnStake(lr *types.LogRecord) {
	// sanity check for data (1 uint256 = 32 bytes)
	if len(lr.Data) != 32 {
		log.Criticalf("%s lr invalid data length; expected 32 bytes, %d bytes given, %d topics given", lr.TxHash.String(), len(lr.Data), len(lr.Topics))
		return
	}

	// get the validator address (we handle the stake as self-delegation in SFC3 context)
	valID := new(big.Int).SetBytes(lr.Topics[1].Bytes())
	addr, err := repo.ValidatorAddress((*hexutil.Big)(valID))
	if err != nil {
		log.Errorf("unknown validator #%d; %s", valID.Uint64(), err.Error())
		return
	}

	// handle request closure
	handleFinishedWithdrawRequest(*addr, valID, new(big.Int), new(big.Int).SetBytes(lr.Data[:]), lr)
}

// handleSfc1DeactivatedStake handles a validator stake deactivation event
// so the stake could be withdrawn later, this basically creates a withdrawal request on the stake
// without request ID.
// DeactivatedStake(uint256 indexed stakerID)
func handleSfc1DeactivatedStake(lr *types.LogRecord) {
	// sanity check for data
	if len(lr.Data) != 0 {
		log.Criticalf("%s lr invalid data length; expected 0 bytes, %d bytes given, %d topics given", lr.TxHash.String(), len(lr.Data), len(lr.Topics))
		return
	}

	// get the validator address (we handle the stake as self-delegation in SFC3 context)
	valID := new(big.Int).SetBytes(lr.Topics[1].Bytes())
	addr, err := repo.ValidatorAddress((*hexutil.Big)(valID))
	if err != nil {
		log.Errorf("unknown validator #%d; %s", valID.Uint64(), err.Error())
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
		lr,
	)
}

// handleSfc1UpdatedStake handles stake update event; the event is emitted
// when a stake and/or delegation amount changes on a validator  account.
// UpdatedStake(uint256 indexed stakerID, uint256 amount, uint256 delegatedMe)
func handleSfc1UpdatedStake(lr *types.LogRecord) {
	// get the validator address (we handle the stake as self-delegation in SFC3 context)
	valID := (*hexutil.Big)(new(big.Int).SetBytes(lr.Topics[1].Bytes()))
	addr, err := repo.ValidatorAddress(valID)
	if err != nil {
		log.Errorf("unknown validator #%d; %s", valID.ToInt().Uint64(), err.Error())
		return
	}

	// update the balance
	if err := repo.UpdateDelegationBalance(addr, valID, func(amo *big.Int) error {
		log.Criticalf("expected validator %d stake not found at %s", valID.ToInt().Uint64(), lr.TxHash.String())
		return makeAdHocDelegation(lr, addr, valID, amo)
	}); err != nil {
		log.Errorf("failed to update delegation; %s", err.Error())
	}
}
