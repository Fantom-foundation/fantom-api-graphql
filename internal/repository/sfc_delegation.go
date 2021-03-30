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
)

// IsDelegating returns if the given address is an SFC delegator.
func (p *proxy) IsDelegating(addr *common.Address) (bool, error) {
	count, err := p.db.DelegationsCountFiltered(&bson.D{{"addr", addr.String()}})
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

// Delegation returns a detail of delegation for the given address.
func (p *proxy) Delegation(addr *common.Address, valID *hexutil.Big) (*types.Delegation, error) {
	p.log.Debugf("loading delegation of %s to #%d", addr.String(), valID.ToInt().Uint64())
	return p.db.Delegation(addr, valID)
}

// DelegationsByAddress returns a list of all delegations of a given delegator address.
func (p *proxy) DelegationsByAddress(addr *common.Address, cursor *string, count int32) (*types.DelegationList, error) {
	p.log.Debugf("loading delegations of %s", addr.String())
	return p.db.Delegations(cursor, count, &bson.D{{"addr", addr.String()}})
}

// DelegationsOfValidator extract a list of delegations for a given validator.
func (p *proxy) DelegationsOfValidator(valID *hexutil.Big, cursor *string, count int32) (*types.DelegationList, error) {
	p.log.Debugf("loading delegations of #%d", valID.ToInt().Uint64())
	return p.db.Delegations(cursor, count, &bson.D{{"to", valID.String()}})
}

// DelegationLock returns delegation lock information using SFC contract binding.
func (p *proxy) DelegationLock(dlg *types.Delegation) (*types.DelegationLock, error) {
	p.log.Debugf("loading lock information for %s to #%d", dlg.Address.String(), dlg.ToStakerId.ToInt().Uint64())
	return p.rpc.DelegationLock(dlg)
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
func (p *proxy) DelegationFluidStakingActive(dl *types.Delegation) (bool, error) {
	return p.rpc.DelegationFluidStakingActive(dl)
}

// handleDelegationLog handles a new delegation event from logs.
func handleDelegationLog(blk hexutil.Uint64, trx *common.Hash, stakerID *big.Int, addr common.Address, amo *big.Int, ld *logsDispatcher) {
	// get the block
	block, err := ld.repo.BlockByNumber(&blk)
	if err != nil {
		ld.log.Errorf("can not decode delegation log record; %s", err.Error())
		return
	}

	// make the delegation record
	dl := types.Delegation{
		Transaction:     *trx,
		Address:         addr,
		ToStakerId:      (*hexutil.Big)(stakerID),
		AmountDelegated: (*hexutil.Big)(amo),
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
	handleDelegationLog(
		hexutil.Uint64(log.BlockNumber),
		&log.TxHash,
		new(big.Int).SetBytes(log.Topics[2].Bytes()),
		common.BytesToAddress(log.Topics[1].Bytes()),
		new(big.Int).SetBytes(log.Data),
		ld,
	)
}

// handleSfcCreatedStake handles a new stake event from SFC v1 and SFC v2 contract.
// event CreatedStake(uint256 indexed stakerID, address indexed dagSfcAddress, uint256 amount);
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

// handleSfcUndelegated handles new withdrawal request from SFCv3 contract.
// We ignore withdrawals from previous SFC versions since after the upgrade all the pending
// withdrawals will be settled.
// event Undelegated(address indexed delegator, uint256 indexed toValidatorID, uint256 indexed wrID, uint256 amount)
func handleSfcUndelegated(log *retypes.Log, ld *logsDispatcher) {
	// get the block
	blk := hexutil.Uint64(log.BlockNumber)
	block, err := ld.repo.BlockByNumber(&blk)
	if err != nil {
		ld.log.Errorf("can not decode un-delegation log record; %s", err.Error())
		return
	}

	// make the request
	wr := types.WithdrawRequest{
		RequestTrx:        log.TxHash,
		WithdrawRequestID: (*hexutil.Big)(new(big.Int).SetBytes(log.Topics[3].Bytes())),
		Address:           common.BytesToAddress(log.Topics[1].Bytes()),
		StakerID:          (*hexutil.Big)(new(big.Int).SetBytes(log.Topics[2].Bytes())),
		CreatedTime:       block.TimeStamp,
		Amount:            (*hexutil.Big)(new(big.Int).SetBytes(log.Data)),
	}

	// store the request
	if err := ld.repo.StoreWithdrawRequest(&wr); err != nil {
		ld.log.Errorf("failed to store new withdraw request; %s", err.Error())
	}
}

// handleSfcWithdrawn handles a withdrawal request finalization event.
// event Withdrawn(address indexed delegator, uint256 indexed toValidatorID, uint256 indexed wrID, uint256 amount)
func handleSfcWithdrawn(log *retypes.Log, ld *logsDispatcher) {
	// extract the basic info about the request
	addr := common.BytesToAddress(log.Topics[1].Bytes())
	valID := (*hexutil.Big)(new(big.Int).SetBytes(log.Topics[2].Bytes()))
	reqID := (*hexutil.Big)(new(big.Int).SetBytes(log.Topics[3].Bytes()))

	// try to get the request from database
	req, err := ld.repo.WithdrawRequest(&addr, valID, reqID)
	if err != nil {
		ld.log.Errorf("can not load withdraw requests to process finalization; %s", err.Error())
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

	// store the updated request
	if err := ld.repo.StoreWithdrawRequest(req); err != nil {
		ld.log.Errorf("failed to store finalized withdraw request; %s", err.Error())
	}
}
