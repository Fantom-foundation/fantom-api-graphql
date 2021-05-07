/*
Package repository implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Opera/Lachesis full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"sync/atomic"
	"time"
)

// trxEstimator helps with counting transaction in repository.
type trxEstimator struct {
	count      uint64
	nextUpdate time.Time
}

// EstimateTransactionsCount returns an approximate amount of transactions on the network.
func (p *proxy) EstimateTransactionsCount() (hexutil.Uint64, error) {
	// ask for the value
	val, err, _ := p.apiRequestGroup.Do("tx_count", func() (interface{}, error) {
		// do we need to pull a recent value from DB?
		if p.txe.nextUpdate.IsZero() || p.txe.nextUpdate.Before(time.Now()) {
			// pull the value from DB
			val, err := p.db.TransactionsCount()
			if err != nil {
				return uint64(0), err
			}

			// store the value
			atomic.StoreUint64(&p.txe.count, val)
			p.txe.nextUpdate = time.Now().Add(15 * time.Minute)
		}
		return p.txe.count, nil
	})
	if err != nil {
		return 0, err
	}
	return hexutil.Uint64(val.(uint64)), nil
}

// TransactionsCountInc updates the value of transaction counter estimator.
func (p *proxy) TransactionsCountInc(diff uint64) {
	atomic.AddUint64(&p.txe.count, diff)
}

// TransactionsCount returns total number of transactions in the block chain.
func (p *proxy) TransactionsCount() (uint64, error) {
	return p.db.TransactionsCount()
}
