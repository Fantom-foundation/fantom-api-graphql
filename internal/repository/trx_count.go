/*
Package repository implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Opera full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"sync/atomic"
)

// EstimateTransactionsCount returns an approximate amount of transactions on the network.
func (p *proxy) EstimateTransactionsCount() (hexutil.Uint64, error) {
	// ask for the value
	val, err, _ := p.apiRequestGroup.Do("tx_count", func() (interface{}, error) {
		return p.txCount, nil
	})
	if err != nil {
		return 0, err
	}
	return hexutil.Uint64(val.(uint64)), nil
}

// MustEstimateTransactionsCount resolves estimated transaction pool size.
func (p *proxy) MustEstimateTransactionsCount() hexutil.Uint64 {
	val, err := p.EstimateTransactionsCount()
	if err != nil {
		return 0
	}
	return val
}

// IncTrxCountEstimate bumps the value of transaction counter estimator.
func (p *proxy) IncTrxCountEstimate(diff uint64) {
	atomic.AddUint64(&p.txCount, diff)
}

// UpdateTrxCountEstimate updates the value of transaction counter estimator.
func (p *proxy) UpdateTrxCountEstimate(val uint64) {
	atomic.StoreUint64(&p.txCount, val)
	p.log.Infof("trx count estimator updated to %d", val)
}

// TransactionsCount returns total number of transactions in the block chain.
func (p *proxy) TransactionsCount() (uint64, error) {
	return p.db.TransactionsCount()
}
