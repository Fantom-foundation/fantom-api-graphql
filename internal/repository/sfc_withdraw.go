/*
Package repository implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Opera full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import (
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson"
	"math/big"
)

// StoreWithdrawRequest stores the given withdraw request in persistent storage.
func (p *proxy) StoreWithdrawRequest(wr *types.WithdrawRequest) error {
	return p.db.AddWithdrawal(wr)
}

// UpdateWithdrawRequest stores the given updated withdraw request in persistent storage.
func (p *proxy) UpdateWithdrawRequest(wr *types.WithdrawRequest) error {
	return p.db.UpdateWithdrawal(wr)
}

// WithdrawRequest extracts details of a withdraw request specified by the delegator, validator and request ID.
func (p *proxy) WithdrawRequest(addr *common.Address, valID *hexutil.Big, reqID *hexutil.Big) (*types.WithdrawRequest, error) {
	return p.db.Withdrawal(addr, valID, reqID)
}

// WithdrawRequests extracts a list of partial withdraw requests for the given address.
func (p *proxy) WithdrawRequests(addr *common.Address, stakerID *hexutil.Big, cursor *string, count int32, activeOnly bool) (*types.WithdrawRequestList, error) {
	if addr == nil {
		return nil, fmt.Errorf("address not given")
	}

	// get all the requests for the given delegator address
	if stakerID == nil {
		// log the action and pull the list for all vals
		p.log.Debugf("loading withdraw requests of %s to any validator", addr.String())
		return p.db.Withdrawals(cursor, count, &bson.D{{Key: types.FiWithdrawalAddress, Value: addr.String()}})
	}

	// log the action and pull the list for specific address and val
	p.log.Debugf("loading withdraw requests of %s to #%d", addr.String(), stakerID.ToInt().Uint64())
	filter := bson.D{
		{Key: types.FiWithdrawalAddress, Value: addr.String()},
		{Key: types.FiWithdrawalToValidator, Value: stakerID.String()},
	}

	// limit the output to active withdrawals only
	if activeOnly {
		filter = append(filter, bson.E{Key: types.FiWithdrawalFinTrx, Value: bson.D{{Key: "$type", Value: 10}}})
	}
	return p.db.Withdrawals(cursor, count, &filter)
}

// WithdrawRequestsPendingTotal is the total value of all pending withdrawal requests
// for the given delegator and target staker ID.
func (p *proxy) WithdrawRequestsPendingTotal(addr *common.Address, stakerID *hexutil.Big) (*big.Int, error) {
	if addr == nil {
		return nil, fmt.Errorf("address not given")
	}

	// all withdrawals for the address regardless of the target staker
	if stakerID == nil {
		return p.db.WithdrawalsSumValue(&bson.D{
			{Key: types.FiWithdrawalAddress, Value: addr.String()},
			{Key: types.FiWithdrawalFinTrx, Value: bson.D{{Key: "$type", Value: 10}}},
		})
	}

	// specific delegation withdrawal
	return p.db.WithdrawalsSumValue(&bson.D{
		{Key: types.FiWithdrawalAddress, Value: addr.String()},
		{Key: types.FiWithdrawalToValidator, Value: stakerID.String()},
		{Key: types.FiWithdrawalFinTrx, Value: bson.D{{Key: "$type", Value: 10}}},
	})
}
