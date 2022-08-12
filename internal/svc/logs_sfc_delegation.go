// Package svc implements blockchain data processing services.
package svc

import (
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"time"
)

// handleDelegationLog handles a new delegation event from logs.
func handleNewDelegation(lr *types.LogRecord, stakerID *big.Int, addr common.Address, amo *big.Int) {
	// get the validator address
	val, err := repo.ValidatorAddress((*hexutil.Big)(stakerID))
	if err != nil {
		log.Errorf("unknown validator #%d; %s", stakerID.Uint64(), err.Error())
		return
	}

	// pull the current value of the stake
	staked, err := repo.DelegationAmountStaked(&addr, (*hexutil.Big)(stakerID))
	if err != nil {
		log.Errorf("delegation balance not available for %s to %d; %s", addr.String(), stakerID.Uint64(), err.Error())
		return
	}

	// make the delegation record
	dl := types.Delegation{
		Transaction:     lr.Trx.Hash,
		Address:         addr,
		ToStakerId:      (*hexutil.Big)(stakerID),
		ToStakerAddress: *val,
		AmountDelegated: (*hexutil.Big)(staked),
		AmountStaked:    (*hexutil.Big)(amo),
		CreatedTime:     lr.Block.TimeStamp,
	}

	// store the delegation
	if err := repo.StoreDelegation(&dl); err != nil {
		log.Errorf("failed to store delegation; %s", err.Error())
	}
}

// handleSfcCreatedDelegation handles a new delegation event from SFC v1 and SFC v2 contract
// and also the new delegation event from SFC3 contract with the same structure.
// (SFCv1, SFCv2) event CreatedDelegation(address indexed delegator, uint256 indexed toStakerID, uint256 amount)
// (SFCv3) event Delegated(address indexed delegator, uint256 indexed toValidatorID, uint256 amount)
func handleSfcCreatedDelegation(lr *types.LogRecord) {
	handleNewDelegation(
		lr,
		new(big.Int).SetBytes(lr.Topics[2].Bytes()),
		common.BytesToAddress(lr.Topics[1].Bytes()),
		new(big.Int).SetBytes(lr.Data),
	)
}

// handleSfc1IncreasedDelegation handles delegation amount increase event in SFC v1 and SFC v2.
// SFC1::IncreasedDelegation(address indexed delegator, uint256 indexed stakerID, uint256 newAmount, uint256 diff);
// SFC1::PreparedToWithdrawDelegation(address indexed delegator, uint256 indexed stakerID)
func handleSfc1IncreasedDelegation(lr *types.LogRecord) {
	// get the validator ID
	addr := common.BytesToAddress(lr.Topics[1].Bytes())
	valID := new(big.Int).SetBytes(lr.Topics[2].Bytes())

	// update the balance
	if err := repo.UpdateDelegationBalance(&addr, (*hexutil.Big)(valID), func(amo *big.Int) error {
		return makeAdHocDelegation(lr, &addr, (*hexutil.Big)(valID), amo)
	}); err != nil {
		log.Errorf("failed to update delegation; %s", err.Error())
	}
}

// handleSfcUndelegated handles new withdrawal request from SFCv3 contract.
// We ignore withdrawals from previous SFC versions since after the upgrade all the pending
// withdrawals will be settled.
// event Undelegated(address indexed delegator, uint256 indexed toValidatorID, uint256 indexed wrID, uint256 amount)
func handleSfcUndelegated(lr *types.LogRecord) {
	// sanity check for data (1x uint256 = 32 bytes)
	if len(lr.Data) != 32 {
		log.Criticalf("%s lr invalid data length; expected 32 bytes, %d bytes given, %d topics given", lr.TxHash.String(), len(lr.Data), len(lr.Topics))
		return
	}

	// create withdraw request
	handleNewWithdrawRequest(
		types.WithdrawTypeUndelegated,
		common.BytesToAddress(lr.Topics[1].Bytes()),
		new(big.Int).SetBytes(lr.Topics[2].Bytes()),
		new(big.Int).SetBytes(lr.Topics[3].Bytes()),
		new(big.Int).SetBytes(lr.Data[:]),
		lr,
	)
}

// handleNewWithdrawRequest will create a new withdrawal request for the given stake.
func handleNewWithdrawRequest(wrt string, adr common.Address, valID *big.Int, reqID *big.Int, amo *big.Int, lr *types.LogRecord) {
	// make the request
	wr := types.WithdrawRequest{
		Type:              wrt,
		RequestTrx:        lr.TxHash,
		WithdrawRequestID: (*hexutil.Big)(reqID),
		Address:           adr,
		StakerID:          (*hexutil.Big)(valID),
		CreatedTime:       lr.Block.TimeStamp,
		Amount:            (*hexutil.Big)(amo),
	}

	// lr what we do
	log.Debugf("new withdrawal of type %s by %s to #%d, req %s at %s",
		wrt,
		adr.String(),
		valID.Uint64(),
		((*hexutil.Big)(reqID)).String(),
		lr.TxHash.String(),
	)

	// store the request
	if err := repo.StoreWithdrawRequest(&wr); err != nil {
		log.Errorf("failed to store new withdraw request; %s", err.Error())
	}

	// check active amount on the delegation
	if err := repo.UpdateDelegationBalance(&wr.Address, wr.StakerID, func(amo *big.Int) error {
		return makeAdHocDelegation(lr, &wr.Address, wr.StakerID, amo)
	}); err != nil {
		log.Errorf("failed to update delegation; %s", err.Error())
	}
}

// handleFinishedWithdrawRequest handles withdrawal request finalisation event.
func handleFinishedWithdrawRequest(adr common.Address, valID *big.Int, reqID *big.Int, penalty *big.Int, lr *types.LogRecord) {
	// make sure the delegation balance will be updated
	defer func() {
		// check active amount on the delegation
		if err := repo.UpdateDelegationBalance(&adr, (*hexutil.Big)(valID), func(amo *big.Int) error {
			return makeAdHocDelegation(lr, &adr, (*hexutil.Big)(valID), amo)
		}); err != nil {
			log.Errorf("failed to update delegation; %s", err.Error())
		}
	}()

	// lr what we do
	log.Debugf("closing withdrawal by %s to #%d, req %s at %s",
		adr.String(),
		valID.Uint64(),
		((*hexutil.Big)(reqID)).String(),
		lr.TxHash.String(),
	)

	// try to get the request from database
	req, err := repo.WithdrawRequest(&adr, (*hexutil.Big)(valID), (*hexutil.Big)(reqID))
	if err != nil {
		log.Errorf("can not load withdraw requests to finalise; %s", err.Error())
		return
	}

	// update the request to have the finalization details
	req.WithdrawTime = &lr.Block.TimeStamp
	req.WithdrawTrx = &lr.TxHash
	req.Penalty = (*hexutil.Big)(penalty)

	// store the updated request
	if err := repo.UpdateWithdrawRequest(req); err != nil {
		log.Errorf("failed to store finalized withdraw request; %s", err.Error())
	}
}

// handleSfc1DeactivatedDelegation handles SFC1 delegation deactivation request.
// SFC1::DeactivatedDelegation(address indexed delegator, uint256 indexed stakerID)
// SFC1::PreparedToWithdrawDelegation(address indexed delegator, uint256 indexed stakerID)
func handleSfc1DeactivatedDelegation(lr *types.LogRecord) {
	// sanity check for data
	if len(lr.Data) != 0 {
		log.Criticalf("%s lr invalid data length; expected 0 bytes, %d bytes given, %d topics given", lr.TxHash.String(), len(lr.Data), len(lr.Topics))
		return
	}

	// create withdraw request
	zero := new(big.Int)
	handleNewWithdrawRequest(
		types.WithdrawTypeDeactivatedDlg,
		/* address */ common.BytesToAddress(lr.Topics[1].Bytes()),
		/* valID */ new(big.Int).SetBytes(lr.Topics[2].Bytes()),
		/* reqID*/ zero,
		/* amount */ zero,
		lr,
	)
}

// handleSfc1CreatedWithdrawRequest handles withdraw request creation event in SFC1/2
// CreatedWithdrawRequest(address indexed auth, address indexed receiver, uint256 indexed stakerID, uint256 wrID, bool delegation, uint256 amount)
func handleSfc1CreatedWithdrawRequest(lr *types.LogRecord) {
	// sanity check for data (2x uint256 + 1x bool = 96 bytes)
	if len(lr.Data) != 96 {
		log.Criticalf("%s lr invalid data length; expected 96 bytes, %d bytes given", lr.TxHash.String(), len(lr.Data))
		return
	}

	// create withdraw request
	handleNewWithdrawRequest(
		types.WithdrawTypeWithdrawRequest,
		/* address */ common.BytesToAddress(lr.Topics[1].Bytes()),
		/* valID */ new(big.Int).SetBytes(lr.Topics[3].Bytes()),
		/* reqID*/ new(big.Int).SetBytes(lr.Data[:32]),
		/* amount */ new(big.Int).SetBytes(lr.Data[64:]),
		lr,
	)
}

// handleSfc1PartialWithdrawByRequest handles SFC1 withdraw finalization event.
// PartialWithdrawnByRequest(address indexed auth, address indexed receiver, uint256 indexed stakerID, uint256 wrID, bool delegation, uint256 penalty)
func handleSfc1PartialWithdrawByRequest(lr *types.LogRecord) {
	// sanity check for data (2x uint256 + 1x bool = 96 bytes)
	if len(lr.Data) != 96 {
		log.Criticalf("%s lr invalid data length; expected 96 bytes, %d bytes given, %d topics given", lr.TxHash.String(), len(lr.Data), len(lr.Topics))
		return
	}

	// finish the request
	handleFinishedWithdrawRequest(
		/* address */ common.BytesToAddress(lr.Topics[1].Bytes()),
		/* valID */ new(big.Int).SetBytes(lr.Topics[3].Bytes()),
		/* reqID*/ new(big.Int).SetBytes(lr.Data[:32]),
		/* penalty */ new(big.Int).SetBytes(lr.Data[64:]),
		lr,
	)
}

// handleSfc1UpdatedDelegation handles delegation update event.
// UpdatedDelegation(address indexed delegator, uint256 indexed oldStakerID, uint256 indexed newStakerID, uint256 amount)
func handleSfc1UpdatedDelegation(lr *types.LogRecord) {
	// sanity check for data (4x topic + 1x uint256 (value) = 32 bytes)
	if len(lr.Topics) != 4 || len(lr.Data) != 32 {
		log.Criticalf("%s not UpdatedDelegation; expected 32 bytes, %d bytes given; expected 4 topics, %d topics given", lr.TxHash.String(), len(lr.Data), len(lr.Topics))
		return
	}

	// check active amount on the delegation
	addr := common.BytesToAddress(lr.Topics[1].Bytes())
	valID := (*hexutil.Big)(new(big.Int).SetBytes(lr.Topics[2].Bytes()))
	if err := repo.UpdateDelegationBalance(&addr, valID, func(amo *big.Int) error {
		return makeAdHocDelegation(lr, &addr, valID, amo)
	}); err != nil {
		log.Errorf("failed to update delegation; %s", err.Error())
	}

	// this should have created a new delegation
	handleNewDelegation(
		lr,
		new(big.Int).SetBytes(lr.Topics[3].Bytes()),
		addr,
		new(big.Int).SetBytes(lr.Data[:]),
	)
}

// handleSfcWithdrawn handles a withdrawal request finalization event.
// event Withdrawn(address indexed delegator, uint256 indexed toValidatorID, uint256 indexed wrID, uint256 amount)
func handleSfcWithdrawn(lr *types.LogRecord) {
	// sanity check for data (4x topic + 1x + 1 x uint256 = 32 bytes)
	if len(lr.Topics) != 4 || len(lr.Data) != 32 {
		log.Criticalf("%s is not event Withdrawn; expected 32 bytes, %d bytes given; expected 4 topics, %d given", lr.TxHash.String(), len(lr.Data), len(lr.Topics))
		return
	}

	// finish the request
	handleFinishedWithdrawRequest(
		common.BytesToAddress(lr.Topics[1].Bytes()),
		new(big.Int).SetBytes(lr.Topics[2].Bytes()),
		new(big.Int).SetBytes(lr.Topics[3].Bytes()),
		new(big.Int),
		lr,
	)
}

// handleSfc1WithdrawnDelegation handles a withdrawal request finalization event from SFC1.
// event WithdrawnDelegation(address indexed delegator, uint256 indexed stakerID, uint256 penalty)
func handleSfc1WithdrawnDelegation(lr *types.LogRecord) {
	// sanity check for data (3x topic + 1x + 1 x uint256 = 32 bytes)
	if len(lr.Topics) != 3 || len(lr.Data) != 32 {
		log.Criticalf("%s is not event Withdrawn; expected 32 bytes, %d bytes given; expected 3 topics, %d given", lr.TxHash.String(), len(lr.Data), len(lr.Topics))
		return
	}

	// extract the basic info about the request
	addr := common.BytesToAddress(lr.Topics[1].Bytes())
	valID := new(big.Int).SetBytes(lr.Topics[2].Bytes())

	// close the previous request
	zero := new(big.Int)
	handleFinishedWithdrawRequest(addr, valID, zero, zero, lr)
}

// makeAdHocDelegation creates a new delegation in case an expected existing delegation
// could not be found on a new lr event processing.
func makeAdHocDelegation(lr *types.LogRecord, addr *common.Address, stakerID *hexutil.Big, amo *big.Int) error {
	// get staker address
	val, err := repo.ValidatorAddress(stakerID)
	if err != nil {
		return err
	}

	// do the insert
	log.Noticef("creating ad-hoc delegation of %s to #%d", addr.String(), stakerID.ToInt().Uint64())
	return repo.StoreDelegation(&types.Delegation{
		Transaction:     lr.TxHash,
		Address:         *addr,
		ToStakerId:      stakerID,
		ToStakerAddress: *val,
		AmountDelegated: (*hexutil.Big)(amo),
		AmountStaked:    (*hexutil.Big)(amo),
		CreatedTime:     lr.Block.TimeStamp,
	})
}

// handleLockedUpStake handles a new and/or updated delegation processed on the SFC.
// event LockedUpStake(address indexed delegator, uint256 indexed validatorID, uint256 duration, uint256 amount)
func handleLockedUpStake(lr *types.LogRecord) {
	// sanity check for data (3x topic + 2x uint256 = 64 bytes)
	if len(lr.Topics) != 3 || len(lr.Data) != 64 {
		log.Criticalf("%s is not event LockedUpStake; expected 64 bytes, %d bytes given; expected 3 topics, %d given", lr.TxHash.String(), len(lr.Data), len(lr.Topics))
		return
	}

	// extract the basic info about the lock
	lock := types.LockedDelegation{
		Delegator:   common.BytesToAddress(lr.Topics[1].Bytes()),
		ValidatorId: new(big.Int).SetBytes(lr.Topics[2].Bytes()).Int64(),
		Locked:      time.Unix(int64(lr.Block.TimeStamp), 0),
		Duration:    new(big.Int).SetBytes(lr.Data[:32]).Int64(),
	}
	lock.LockedUntil = time.Unix(int64(lr.Block.TimeStamp), 0).Add(time.Duration(lock.Duration) * time.Second)
	lock.SetAmount((*hexutil.Big)(new(big.Int).SetBytes(lr.Data[32:])))

	// log what happened
	log.Noticef("delegation of %s to #%d locked until %s", lock.Delegator.String(), lock.ValidatorId, lock.LockedUntil.String())
	err := repo.StoreLockedDelegation(&lock)
	if err != nil {
		log.Errorf("failed to store locked delegation; %s", err.Error())
	}
}

// handleUnlockedStake handles a new and/or updated delegation processed on the SFC.
// event UnlockedStake(address indexed delegator, uint256 indexed validatorID, uint256 amount, uint256 penalty)
func handleUnlockedStake(lr *types.LogRecord) {
	// sanity check for data (3x topic + 2x uint256 = 64 bytes)
	if len(lr.Topics) != 3 || len(lr.Data) != 64 {
		log.Criticalf("%s is not event UnlockedStake; expected 64 bytes, %d bytes given; expected 3 topics, %d given", lr.TxHash.String(), len(lr.Data), len(lr.Topics))
		return
	}

	err := repo.AdjustLockedDelegation(
		common.BytesToAddress(lr.Topics[1].Bytes()),
		new(big.Int).SetBytes(lr.Topics[2].Bytes()).Int64(),
		-types.LockedDelegationValue(new(big.Int).SetBytes(lr.Data[32:])),
	)
	if err != nil {
		log.Errorf("failed to adjust locked delegation value; %s", err.Error())
	}
}
