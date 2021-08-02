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

// handleDelegationLog handles a new delegation event from logs.
func handleNewDelegation(log *types.LogRecord, stakerID *big.Int, addr common.Address, amo *big.Int) {
	// get the validator address
	val, err := R().ValidatorAddress((*hexutil.Big)(stakerID))
	if err != nil {
		R().Log().Errorf("unknown validator #%d; %s", stakerID.Uint64(), err.Error())
		return
	}

	// pull the current value of the stake
	staked, err := R().DelegationAmountStaked(&addr, (*hexutil.Big)(stakerID))
	if err != nil {
		R().Log().Errorf("delegation balance not available for %s to %d; %s", addr.String(), stakerID.Uint64(), err.Error())
		return
	}

	// make the delegation record
	dl := types.Delegation{
		Transaction:     log.Trx.Hash,
		Address:         addr,
		ToStakerId:      (*hexutil.Big)(stakerID),
		ToStakerAddress: *val,
		AmountDelegated: (*hexutil.Big)(staked),
		AmountStaked:    (*hexutil.Big)(amo),
		CreatedTime:     log.Block.TimeStamp,
	}

	// store the delegation
	if err := R().StoreDelegation(&dl); err != nil {
		R().Log().Errorf("failed to store delegation; %s", err.Error())
	}
}

// handleSfcCreatedDelegation handles a new delegation event from SFC v1 and SFC v2 contract
// and also the new delegation event from SFC3 contract with the same structure.
// (SFCv1, SFCv2) event CreatedDelegation(address indexed delegator, uint256 indexed toStakerID, uint256 amount)
// (SFCv3) event Delegated(address indexed delegator, uint256 indexed toValidatorID, uint256 amount)
func handleSfcCreatedDelegation(log *types.LogRecord) {
	handleNewDelegation(
		log,
		new(big.Int).SetBytes(log.Topics[2].Bytes()),
		common.BytesToAddress(log.Topics[1].Bytes()),
		new(big.Int).SetBytes(log.Data),
	)
}

// handleSfc1IncreasedDelegation handles delegation amount increase event in SFC v1 and SFC v2.
// SFC1::IncreasedDelegation(address indexed delegator, uint256 indexed stakerID, uint256 newAmount, uint256 diff);
// SFC1::PreparedToWithdrawDelegation(address indexed delegator, uint256 indexed stakerID)
func handleSfc1IncreasedDelegation(log *types.LogRecord) {
	// get the validator ID
	addr := common.BytesToAddress(log.Topics[1].Bytes())
	valID := new(big.Int).SetBytes(log.Topics[2].Bytes())

	// update the balance
	if err := R().UpdateDelegationBalance(&addr, (*hexutil.Big)(valID), func(amo *big.Int) error {
		return makeAdHocDelegation(log, &addr, (*hexutil.Big)(valID), amo)
	}); err != nil {
		R().Log().Errorf("failed to update delegation; %s", err.Error())
	}
}

// handleSfcUndelegated handles new withdrawal request from SFCv3 contract.
// We ignore withdrawals from previous SFC versions since after the upgrade all the pending
// withdrawals will be settled.
// event Undelegated(address indexed delegator, uint256 indexed toValidatorID, uint256 indexed wrID, uint256 amount)
func handleSfcUndelegated(log *types.LogRecord) {
	// sanity check for data (1x uint256 = 32 bytes)
	if len(log.Data) != 32 {
		R().Log().Criticalf("%s log invalid data length; expected 32 bytes, %d bytes given, %d topics given", log.TxHash.String(), len(log.Data), len(log.Topics))
		return
	}

	// create withdraw request
	handleNewWithdrawRequest(
		types.WithdrawTypeUndelegated,
		common.BytesToAddress(log.Topics[1].Bytes()),
		new(big.Int).SetBytes(log.Topics[2].Bytes()),
		new(big.Int).SetBytes(log.Topics[3].Bytes()),
		new(big.Int).SetBytes(log.Data[:]),
		log,
	)
}

// handleNewWithdrawRequest will create a new withdrawal request for the given stake.
func handleNewWithdrawRequest(wrt string, adr common.Address, valID *big.Int, reqID *big.Int, amo *big.Int, log *types.LogRecord) {
	// make the request
	wr := types.WithdrawRequest{
		Type:              wrt,
		RequestTrx:        log.TxHash,
		WithdrawRequestID: (*hexutil.Big)(reqID),
		Address:           adr,
		StakerID:          (*hexutil.Big)(valID),
		CreatedTime:       log.Block.TimeStamp,
		Amount:            (*hexutil.Big)(amo),
	}

	// log what we do
	R().Log().Debugf("new withdrawal of type %s by %s to #%d, req %s at %s",
		wrt,
		adr.String(),
		valID.Uint64(),
		((*hexutil.Big)(reqID)).String(),
		log.TxHash.String(),
	)

	// store the request
	if err := R().StoreWithdrawRequest(&wr); err != nil {
		R().Log().Errorf("failed to store new withdraw request; %s", err.Error())
	}

	// check active amount on the delegation
	if err := R().UpdateDelegationBalance(&wr.Address, wr.StakerID, func(amo *big.Int) error {
		return makeAdHocDelegation(log, &wr.Address, wr.StakerID, amo)
	}); err != nil {
		R().Log().Errorf("failed to update delegation; %s", err.Error())
	}
}

// handleFinishedWithdrawRequest handles withdrawal request finalisation event.
func handleFinishedWithdrawRequest(adr common.Address, valID *big.Int, reqID *big.Int, penalty *big.Int, log *types.LogRecord) {
	// make sure the delegation balance will be updated
	defer func() {
		// check active amount on the delegation
		if err := R().UpdateDelegationBalance(&adr, (*hexutil.Big)(valID), func(amo *big.Int) error {
			return makeAdHocDelegation(log, &adr, (*hexutil.Big)(valID), amo)
		}); err != nil {
			R().Log().Errorf("failed to update delegation; %s", err.Error())
		}
	}()

	// log what we do
	R().Log().Debugf("closing withdrawal by %s to #%d, req %s at %s",
		adr.String(),
		valID.Uint64(),
		((*hexutil.Big)(reqID)).String(),
		log.TxHash.String(),
	)

	// try to get the request from database
	req, err := R().WithdrawRequest(&adr, (*hexutil.Big)(valID), (*hexutil.Big)(reqID))
	if err != nil {
		R().Log().Errorf("can not load withdraw requests to finalise; %s", err.Error())
		return
	}

	// update the request to have the finalization details
	req.WithdrawTime = &log.Block.TimeStamp
	req.WithdrawTrx = &log.TxHash
	req.Penalty = (*hexutil.Big)(penalty)

	// store the updated request
	if err := R().UpdateWithdrawRequest(req); err != nil {
		R().Log().Errorf("failed to store finalized withdraw request; %s", err.Error())
	}
}

// handleSfc1DeactivatedDelegation handles SFC1 delegation deactivation request.
// SFC1::DeactivatedDelegation(address indexed delegator, uint256 indexed stakerID)
// SFC1::PreparedToWithdrawDelegation(address indexed delegator, uint256 indexed stakerID)
func handleSfc1DeactivatedDelegation(log *types.LogRecord) {
	// sanity check for data
	if len(log.Data) != 0 {
		R().Log().Criticalf("%s log invalid data length; expected 0 bytes, %d bytes given, %d topics given", log.TxHash.String(), len(log.Data), len(log.Topics))
		return
	}

	// create withdraw request
	zero := new(big.Int)
	handleNewWithdrawRequest(
		types.WithdrawTypeDeactivatedDlg,
		/* address */ common.BytesToAddress(log.Topics[1].Bytes()),
		/* valID */ new(big.Int).SetBytes(log.Topics[2].Bytes()),
		/* reqID*/ zero,
		/* amount */ zero,
		log,
	)
}

// handleSfc1CreatedWithdrawRequest handles withdraw request creation event in SFC1/2
// CreatedWithdrawRequest(address indexed auth, address indexed receiver, uint256 indexed stakerID, uint256 wrID, bool delegation, uint256 amount)
func handleSfc1CreatedWithdrawRequest(log *types.LogRecord) {
	// sanity check for data (2x uint256 + 1x bool = 96 bytes)
	if len(log.Data) != 96 {
		R().Log().Criticalf("%s log invalid data length; expected 96 bytes, %d bytes given", log.TxHash.String(), len(log.Data))
		return
	}

	// create withdraw request
	handleNewWithdrawRequest(
		types.WithdrawTypeWithdrawRequest,
		/* address */ common.BytesToAddress(log.Topics[1].Bytes()),
		/* valID */ new(big.Int).SetBytes(log.Topics[3].Bytes()),
		/* reqID*/ new(big.Int).SetBytes(log.Data[:32]),
		/* amount */ new(big.Int).SetBytes(log.Data[64:]),
		log,
	)
}

// handleSfc1PartialWithdrawByRequest handles SFC1 withdraw finalization event.
// PartialWithdrawnByRequest(address indexed auth, address indexed receiver, uint256 indexed stakerID, uint256 wrID, bool delegation, uint256 penalty)
func handleSfc1PartialWithdrawByRequest(log *types.LogRecord) {
	// sanity check for data (2x uint256 + 1x bool = 96 bytes)
	if len(log.Data) != 96 {
		R().Log().Criticalf("%s log invalid data length; expected 96 bytes, %d bytes given, %d topics given", log.TxHash.String(), len(log.Data), len(log.Topics))
		return
	}

	// finish the request
	handleFinishedWithdrawRequest(
		/* address */ common.BytesToAddress(log.Topics[1].Bytes()),
		/* valID */ new(big.Int).SetBytes(log.Topics[3].Bytes()),
		/* reqID*/ new(big.Int).SetBytes(log.Data[:32]),
		/* penalty */ new(big.Int).SetBytes(log.Data[64:]),
		log,
	)
}

// handleSfc1UpdatedDelegation handles delegation update event.
// UpdatedDelegation(address indexed delegator, uint256 indexed oldStakerID, uint256 indexed newStakerID, uint256 amount)
func handleSfc1UpdatedDelegation(log *types.LogRecord) {
	// sanity check for data (1x uint256 (value) = 32 bytes)
	if len(log.Data) != 32 {
		R().Log().Criticalf("%s log invalid data length; expected 32 bytes, %d bytes given", log.TxHash.String(), len(log.Data))
		return
	}

	// check active amount on the delegation
	addr := common.BytesToAddress(log.Topics[1].Bytes())
	valID := (*hexutil.Big)(new(big.Int).SetBytes(log.Topics[2].Bytes()))
	if err := R().UpdateDelegationBalance(&addr, valID, func(amo *big.Int) error {
		return makeAdHocDelegation(log, &addr, valID, amo)
	}); err != nil {
		R().Log().Errorf("failed to update delegation; %s", err.Error())
	}

	// this should have created a new delegation
	handleNewDelegation(
		log,
		new(big.Int).SetBytes(log.Topics[3].Bytes()),
		addr,
		new(big.Int).SetBytes(log.Data[:]),
	)
}

// handleSfcWithdrawn handles a withdrawal request finalization event.
// event Withdrawn(address indexed delegator, uint256 indexed toValidatorID, uint256 indexed wrID, uint256 amount)
func handleSfcWithdrawn(log *types.LogRecord) {
	// finish the request
	handleFinishedWithdrawRequest(
		common.BytesToAddress(log.Topics[1].Bytes()),
		new(big.Int).SetBytes(log.Topics[2].Bytes()),
		new(big.Int).SetBytes(log.Topics[3].Bytes()),
		new(big.Int),
		log,
	)
}

// handleSfc1WithdrawnDelegation handles a withdrawal request finalization event from SFC1.
// event WithdrawnDelegation(address indexed delegator, uint256 indexed stakerID, uint256 penalty)
func handleSfc1WithdrawnDelegation(log *types.LogRecord) {
	// extract the basic info about the request
	addr := common.BytesToAddress(log.Topics[1].Bytes())
	valID := new(big.Int).SetBytes(log.Topics[2].Bytes())

	// close the previous request
	zero := new(big.Int)
	handleFinishedWithdrawRequest(addr, valID, zero, zero, log)
}

// makeAdHocDelegation creates a new delegation in case an expected existing delegation
// could not be found on a new log event processing.
func makeAdHocDelegation(log *types.LogRecord, addr *common.Address, stakerID *hexutil.Big, amo *big.Int) error {
	// get staker address
	val, err := R().ValidatorAddress(stakerID)
	if err != nil {
		return err
	}

	// do the insert
	R().Log().Noticef("creating ad-hoc delegation of %s to #%d", addr.String(), stakerID.ToInt().Uint64())
	return R().StoreDelegation(&types.Delegation{
		Transaction:     log.TxHash,
		Address:         *addr,
		ToStakerId:      stakerID,
		ToStakerAddress: *val,
		AmountDelegated: (*hexutil.Big)(amo),
		AmountStaked:    (*hexutil.Big)(amo),
		CreatedTime:     log.Block.TimeStamp,
	})
}
