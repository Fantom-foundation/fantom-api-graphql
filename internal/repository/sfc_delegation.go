/*
Package repository implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Opera/Lachesis full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import (
	"fantom-api-graphql/internal/repository/db"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	retypes "github.com/ethereum/go-ethereum/core/types"
	"go.mongodb.org/mongo-driver/bson"
	"math/big"
)

// IsDelegating returns if the given address is an SFC delegator.
func (p *proxy) IsDelegating(addr *common.Address) (bool, error) {
	// count only active delegations (with non-zero value)
	count, err := p.db.DelegationsCountFiltered(&bson.D{
		{types.FiDelegationAddress, addr.String()},
		{types.FiDelegationValue, bson.D{{"$gt", 0}}},
	})
	if err != nil {
		p.log.Errorf("can not check delegation by address; %s", addr.String())
		return false, err
	}
	return 0 < count, nil
}

// StoreDelegation stores the delegation in persistent database.
func (p *proxy) StoreDelegation(dl *types.Delegation) error {
	return p.db.AddDelegation(dl)
}

// UpdateDelegationBalance updates active balance of the given delegation.
func (p *proxy) UpdateDelegationBalance(addr *common.Address, valID *hexutil.Big, unknownDelegation func(*big.Int) error) error {
	// pull the current value
	val, err := p.DelegationAmountStaked(addr, valID)
	if err != nil {
		p.log.Errorf("delegation balance not available for %s to %d; %s", addr.String(), valID.ToInt().Uint64(), err.Error())
		return err
	}

	// do the update
	err = p.updateDelegationBalance(addr, valID, val)
	if err == nil {
		return nil
	}

	// unknown delegation detected?
	if err == db.ErrUnknownDelegation {
		p.log.Errorf("delegation %s to %d not known yet", addr.String(), valID.ToInt().Uint64())
		return unknownDelegation(val)
	}

	// some other error
	p.log.Errorf("delegation %s to %d update failed; %s", addr.String(), valID.ToInt().Uint64(), err.Error())
	return err
}

// updateDelegationBalance performs delegation balance update if needed.
func (p *proxy) updateDelegationBalance(addr *common.Address, valID *hexutil.Big, amo *big.Int) error {
	// get the delegation detail
	dlg, err := p.Delegation(addr, valID)
	if err != nil {
		return err
	}

	// do we need to update? if the amount did not change, skip the update
	if dlg.AmountStaked.ToInt().Cmp(amo) == 0 {
		return nil
	}

	// update the delegation in DB and memory
	dlg.AmountDelegated = (*hexutil.Big)(amo)
	err = p.db.UpdateDelegationBalance(addr, valID, dlg.AmountDelegated)
	if nil == err {
		p.cache.PushDelegation(dlg)
	}
	return err
}

// Delegation returns a detail of delegation for the given address.
func (p *proxy) Delegation(adr *common.Address, valID *hexutil.Big) (*types.Delegation, error) {
	// log what we do
	p.log.Debugf("accessing delegation of %s to #%d", adr.String(), valID.ToInt().Uint64())

	// try cache first
	dlg := p.cache.PullDelegation(*adr, valID)
	if dlg != nil {
		return dlg, nil
	}

	// pull from DB instead; do we actually have it?
	dlg, err := p.db.Delegation(adr, valID)
	if err != nil {
		return nil, err
	}

	// store to cache for future reference
	p.cache.PushDelegation(dlg)
	return dlg, nil
}

// DelegationAmountStaked returns the current amount of staked tokens for the given delegation.
func (p *proxy) DelegationAmountStaked(addr *common.Address, valID *hexutil.Big) (*big.Int, error) {
	val, err := p.rpc.AmountStaked(addr, (*big.Int)(valID))
	if err != nil {
		p.log.Errorf("can not get amount delegated by %s to %d; %s", addr.String(), valID.ToInt().Uint64(), err.Error())
		return nil, err
	}
	// log and return
	p.log.Debugf("%s delegated %d to %d", addr.String(), val.Uint64(), valID.ToInt().Uint64())
	return val, nil
}

// DelegationsByAddress returns a list of all delegations of a given delegator address.
func (p *proxy) DelegationsByAddress(addr *common.Address, cursor *common.Hash, count int32) (*types.DelegationList, error) {
	p.log.Debugf("loading delegations of %s", addr.String())
	return p.db.Delegations(cursor, count, &bson.D{{types.FiDelegationAddress, addr.String()}})
}

// DelegationsByAddressAll returns a list of all delegations of the given address un-paged.
func (p *proxy) DelegationsByAddressAll(addr *common.Address) ([]*types.Delegation, error) {
	p.log.Debugf("loading all delegations of %s", addr.String())
	return p.db.DelegationsAll(&bson.D{{types.FiDelegationAddress, addr.String()}})
}

// DelegationsOfValidator extract a list of delegations for a given validator.
func (p *proxy) DelegationsOfValidator(valID *hexutil.Big, cursor *common.Hash, count int32) (*types.DelegationList, error) {
	p.log.Debugf("loading delegations of #%d", valID.ToInt().Uint64())
	return p.db.Delegations(cursor, count, &bson.D{{types.FiDelegationToValidator, valID.String()}})
}

// DelegationLock returns delegation lock information using SFC contract binding.
func (p *proxy) DelegationLock(addr *common.Address, valID *hexutil.Big) (*types.DelegationLock, error) {
	p.log.Debugf("loading lock information for %s to #%d", addr.String(), valID.ToInt().Uint64())
	return p.rpc.DelegationLock(addr, valID)
}

// DelegationAmountUnlocked returns delegation lock information using SFC contract binding.
func (p *proxy) DelegationAmountUnlocked(addr *common.Address, valID *big.Int) (hexutil.Big, error) {
	p.log.Debugf("loading unlocked amount for %s to #%d", addr.String(), valID.Uint64())

	// get the amount
	val, err := p.rpc.AmountStakeUnlocked(addr, valID)
	if err != nil {
		return hexutil.Big{}, err
	}
	return hexutil.Big(*val), nil
}

// DelegationUnlockPenalty returns the amount of penalty applied on given stake unlock.
func (p *proxy) DelegationUnlockPenalty(addr *common.Address, valID *big.Int, amount *big.Int) (hexutil.Big, error) {
	p.log.Debugf("checking unlock of %d penalty for %s to #%d", amount.Uint64(), addr.String(), valID.Uint64())

	// get the amount
	val, err := p.rpc.StakeUnlockPenalty(addr, valID, amount)
	if err != nil {
		return hexutil.Big{}, err
	}
	return hexutil.Big(*val), nil
}

// PendingRewards returns a detail of pending rewards for the given delegation address and validator ID.
func (p *proxy) PendingRewards(addr *common.Address, valID *hexutil.Big) (*types.PendingRewards, error) {
	p.log.Debugf("loading pending rewards of %s to #%d", addr.String(), valID.ToInt().Uint64())
	return p.rpc.PendingRewards(addr, valID.ToInt())
}

// DelegationOutstandingSFTM returns the amount of sFTM tokens for the delegation
// identified by the delegator address and the stakerId.
func (p *proxy) DelegationOutstandingSFTM(addr *common.Address, toStaker *hexutil.Big) (*hexutil.Big, error) {
	val, err := p.rpc.DelegationOutstandingSFTM(addr, toStaker.ToInt())
	if err != nil {
		return nil, err
	}
	return (*hexutil.Big)(val), nil
}

// DelegationTokenizerUnlocked returns the status of SFC Tokenizer lock
// for a delegation identified by the address and staker id.
func (p *proxy) DelegationTokenizerUnlocked(addr *common.Address, toStaker *hexutil.Big) (bool, error) {
	return p.rpc.DelegationTokenizerUnlocked(addr, toStaker.ToInt())
}

// DelegationFluidStakingActive signals if the delegation is upgraded to Fluid Staking model.
func (p *proxy) DelegationFluidStakingActive(_ *common.Address, _ *hexutil.Big) (bool, error) {
	return true, nil
}

// handleDelegationLog handles a new delegation event from logs.
func handleNewDelegation(blk hexutil.Uint64, trx *common.Hash, stakerID *big.Int, addr common.Address, amo *big.Int, ld *logsDispatcher) {
	// get the block
	block, err := ld.repo.BlockByNumber(&blk)
	if err != nil {
		ld.log.Errorf("can not decode delegation log record; %s", err.Error())
		return
	}

	// get the validator address
	val, err := R().ValidatorAddress((*hexutil.Big)(stakerID))
	if err != nil {
		ld.log.Errorf("unknown validator #%d; %s", stakerID.Uint64(), err.Error())
		return
	}

	// pull the current value of the stake
	stAmo, err := ld.repo.DelegationAmountStaked(&addr, (*hexutil.Big)(stakerID))
	if err != nil {
		ld.log.Errorf("delegation balance not available for %s to %d; %s", addr.String(), stakerID.Uint64(), err.Error())
		return
	}

	// make the delegation record
	dl := types.Delegation{
		Transaction:     *trx,
		Address:         addr,
		ToStakerId:      (*hexutil.Big)(stakerID),
		ToStakerAddress: *val,
		AmountDelegated: (*hexutil.Big)(stAmo),
		AmountStaked:    (*hexutil.Big)(amo),
		CreatedTime:     block.TimeStamp,
	}

	// store the delegation
	if err := ld.repo.StoreDelegation(&dl); err != nil {
		ld.log.Errorf("failed to store delegation; %s", err.Error())
	}
}

// handleSfcCreatedDelegation handles a new delegation event from SFC v1 and SFC v2 contract
// and also the new delegation event from SFC3 contract with the same structure.
// (SFCv1, SFCv2) event CreatedDelegation(address indexed delegator, uint256 indexed toStakerID, uint256 amount)
// (SFCv3) event Delegated(address indexed delegator, uint256 indexed toValidatorID, uint256 amount)
func handleSfcCreatedDelegation(log *retypes.Log, ld *logsDispatcher) {
	handleNewDelegation(
		hexutil.Uint64(log.BlockNumber),
		&log.TxHash,
		new(big.Int).SetBytes(log.Topics[2].Bytes()),
		common.BytesToAddress(log.Topics[1].Bytes()),
		new(big.Int).SetBytes(log.Data),
		ld,
	)
}

// handleSfc1IncreasedDelegation handles delegation amount increase event in SFC v1 and SFC v2.
// SFC1::IncreasedDelegation(address indexed delegator, uint256 indexed stakerID, uint256 newAmount, uint256 diff);
// SFC1::PreparedToWithdrawDelegation(address indexed delegator, uint256 indexed stakerID)
func handleSfc1IncreasedDelegation(log *retypes.Log, ld *logsDispatcher) {
	// get the validator ID
	addr := common.BytesToAddress(log.Topics[1].Bytes())
	valID := new(big.Int).SetBytes(log.Topics[2].Bytes())

	// update the balance
	if err := ld.repo.UpdateDelegationBalance(&addr, (*hexutil.Big)(valID), func(amo *big.Int) error {
		return ld.makeAdHocDelegation(log, &addr, (*hexutil.Big)(valID), amo)
	}); err != nil {
		ld.log.Errorf("failed to update delegation; %s", err.Error())
	}
}

// handleSfcUndelegated handles new withdrawal request from SFCv3 contract.
// We ignore withdrawals from previous SFC versions since after the upgrade all the pending
// withdrawals will be settled.
// event Undelegated(address indexed delegator, uint256 indexed toValidatorID, uint256 indexed wrID, uint256 amount)
func handleSfcUndelegated(log *retypes.Log, ld *logsDispatcher) {
	// sanity check for data (1x uint256 = 32 bytes)
	if len(log.Data) != 32 {
		ld.log.Criticalf("%s log invalid data length; expected 32 bytes, %d bytes given, %d topics given", log.TxHash.String(), len(log.Data), len(log.Topics))
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
		ld,
	)
}

// handleNewWithdrawRequest will create a new withdraw request for the given stake.
func handleNewWithdrawRequest(wrt string, adr common.Address, valID *big.Int, reqID *big.Int, amo *big.Int, log *retypes.Log, ld *logsDispatcher) {
	// get the block
	blk := hexutil.Uint64(log.BlockNumber)
	block, err := ld.repo.BlockByNumber(&blk)
	if err != nil {
		ld.log.Errorf("can not decode un-delegation log record; %s", err.Error())
		return
	}

	// make the request
	wr := types.WithdrawRequest{
		Type:              wrt,
		RequestTrx:        log.TxHash,
		WithdrawRequestID: (*hexutil.Big)(reqID),
		Address:           adr,
		StakerID:          (*hexutil.Big)(valID),
		CreatedTime:       block.TimeStamp,
		Amount:            (*hexutil.Big)(amo),
	}

	// store the request
	if err := ld.repo.StoreWithdrawRequest(&wr); err != nil {
		ld.log.Errorf("failed to store new withdraw request; %s", err.Error())
	}

	// check active amount on the delegation
	if err := ld.repo.UpdateDelegationBalance(&wr.Address, wr.StakerID, func(amo *big.Int) error {
		return ld.makeAdHocDelegation(log, &wr.Address, wr.StakerID, amo)
	}); err != nil {
		ld.log.Errorf("failed to update delegation; %s", err.Error())
	}
}

// handleFinishedWithdrawRequest handles withdrawal request finalisation event.
func handleFinishedWithdrawRequest(adr common.Address, valID *big.Int, reqID *big.Int, penalty *big.Int, log *retypes.Log, ld *logsDispatcher) {
	// try to get the request from database
	req, err := ld.repo.WithdrawRequest(&adr, (*hexutil.Big)(valID), (*hexutil.Big)(reqID))
	if err != nil {
		ld.log.Errorf("can not load withdraw requests to finalise; %s", err.Error())
		return
	}

	// get the block
	blk := hexutil.Uint64(log.BlockNumber)
	block, err := ld.repo.BlockByNumber(&blk)
	if err != nil {
		ld.log.Errorf("can not decode un-delegation log record; %s", err.Error())
		return
	}

	// update the request to have the finalization details
	req.WithdrawTime = &block.TimeStamp
	req.WithdrawTrx = &log.TxHash
	req.Penalty = (*hexutil.Big)(penalty)

	// store the updated request
	if err := ld.repo.UpdateWithdrawRequest(req); err != nil {
		ld.log.Errorf("failed to store finalized withdraw request; %s", err.Error())
	}

	// check active amount on the delegation
	if err := ld.repo.UpdateDelegationBalance(&req.Address, req.StakerID, func(amo *big.Int) error {
		return ld.makeAdHocDelegation(log, &adr, (*hexutil.Big)(valID), amo)
	}); err != nil {
		ld.log.Errorf("failed to update delegation; %s", err.Error())
	}
}

// handleSfc1DeactivatedDelegation handles SFC1 delegation deactivation request.
// SFC1::DeactivatedDelegation(address indexed delegator, uint256 indexed stakerID)
// SFC1::PreparedToWithdrawDelegation(address indexed delegator, uint256 indexed stakerID)
func handleSfc1DeactivatedDelegation(log *retypes.Log, ld *logsDispatcher) {
	// sanity check for data
	if len(log.Data) != 0 {
		ld.log.Criticalf("%s log invalid data length; expected 0 bytes, %d bytes given, %d topics given", log.TxHash.String(), len(log.Data), len(log.Topics))
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
		ld,
	)
}

// handleSfc1CreatedWithdrawRequest handles withdraw request creation event in SFC1/2
// CreatedWithdrawRequest(address indexed auth, address indexed receiver, uint256 indexed stakerID, uint256 wrID, bool delegation, uint256 amount)
func handleSfc1CreatedWithdrawRequest(log *retypes.Log, ld *logsDispatcher) {
	// sanity check for data (2x uint256 + 1x bool = 96 bytes)
	if len(log.Data) != 96 {
		ld.log.Criticalf("%s log invalid data length; expected 96 bytes, %d bytes given", log.TxHash.String(), len(log.Data))
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
		ld,
	)
}

// handleSfc1PartialWithdrawByRequest handles SFC1 withdraw finalization event.
// PartialWithdrawnByRequest(address indexed auth, address indexed receiver, uint256 indexed stakerID, uint256 wrID, bool delegation, uint256 penalty)
func handleSfc1PartialWithdrawByRequest(log *retypes.Log, ld *logsDispatcher) {
	// sanity check for data (2x uint256 + 1x bool = 96 bytes)
	if len(log.Data) != 96 {
		ld.log.Criticalf("%s log invalid data length; expected 96 bytes, %d bytes given, %d topics given", log.TxHash.String(), len(log.Data), len(log.Topics))
		return
	}

	// finish the request
	handleFinishedWithdrawRequest(
		/* address */ common.BytesToAddress(log.Topics[1].Bytes()),
		/* valID */ new(big.Int).SetBytes(log.Topics[3].Bytes()),
		/* reqID*/ new(big.Int).SetBytes(log.Data[:32]),
		/* penalty */ new(big.Int).SetBytes(log.Data[64:]),
		log,
		ld,
	)
}

// handleSfc1UpdatedDelegation handles delegation update event.
// UpdatedDelegation(address indexed delegator, uint256 indexed oldStakerID, uint256 indexed newStakerID, uint256 amount)
func handleSfc1UpdatedDelegation(log *retypes.Log, ld *logsDispatcher) {
	// sanity check for data (1x uint256 (value) = 32 bytes)
	if len(log.Data) != 32 {
		ld.log.Criticalf("%s log invalid data length; expected 32 bytes, %d bytes given", log.TxHash.String(), len(log.Data))
		return
	}

	// check active amount on the delegation
	addr := common.BytesToAddress(log.Topics[1].Bytes())
	valID := (*hexutil.Big)(new(big.Int).SetBytes(log.Topics[2].Bytes()))
	if err := ld.repo.UpdateDelegationBalance(&addr, valID, func(amo *big.Int) error {
		return ld.makeAdHocDelegation(log, &addr, valID, amo)
	}); err != nil {
		ld.log.Errorf("failed to update delegation; %s", err.Error())
	}

	// this should have created a new delegation
	handleNewDelegation(
		hexutil.Uint64(log.BlockNumber),
		&log.TxHash,
		new(big.Int).SetBytes(log.Topics[3].Bytes()),
		addr,
		new(big.Int).SetBytes(log.Data[:]),
		ld,
	)
}

// handleSfcWithdrawn handles a withdrawal request finalization event.
// event Withdrawn(address indexed delegator, uint256 indexed toValidatorID, uint256 indexed wrID, uint256 amount)
func handleSfcWithdrawn(log *retypes.Log, ld *logsDispatcher) {
	// finish the request
	handleFinishedWithdrawRequest(
		common.BytesToAddress(log.Topics[1].Bytes()),
		new(big.Int).SetBytes(log.Topics[2].Bytes()),
		new(big.Int).SetBytes(log.Topics[3].Bytes()),
		new(big.Int),
		log,
		ld,
	)
}

// handleSfc1WithdrawnDelegation handles a withdrawal request finalization event from SFC1.
// event WithdrawnDelegation(address indexed delegator, uint256 indexed stakerID, uint256 penalty)
func handleSfc1WithdrawnDelegation(log *retypes.Log, ld *logsDispatcher) {
	// extract the basic info about the request
	addr := common.BytesToAddress(log.Topics[1].Bytes())
	valID := new(big.Int).SetBytes(log.Topics[2].Bytes())

	// close the previous request
	zero := new(big.Int)
	handleFinishedWithdrawRequest(addr, valID, zero, zero, log, ld)
}
