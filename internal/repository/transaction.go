/*
Package repository implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Opera full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import (
	"bytes"
	"errors"
	"fantom-api-graphql/internal/repository/cache"
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	eth "github.com/ethereum/go-ethereum/rpc"
)

// ErrTransactionNotFound represents an error returned if a transaction can not be found.
var ErrTransactionNotFound = errors.New("requested transaction can not be found in Opera blockchain")

// StoreTransaction notifies a new incoming transaction from blockchain to the repository.
func (p *proxy) StoreTransaction(block *types.Block, trx *types.Transaction) error {
	return p.db.AddTransaction(block, trx)
}

// CacheTransaction puts a transaction to the internal ring cache.
func (p *proxy) CacheTransaction(trx *types.Transaction) {
	p.cache.AddTransaction(trx)
}

// Transaction returns a transaction at Opera blockchain by a hash, nil if not found.
// If the transaction is not found, ErrTransactionNotFound error is returned.
func (p *proxy) Transaction(hash *common.Hash) (*types.Transaction, error) {
	p.log.Debugf("requested transaction %s", hash.String())

	// try to use the in-memory cache
	if trx := p.cache.PullTransaction(hash); trx != nil {
		p.log.Debugf("transaction %s loaded from cache", hash.String())
		return trx, nil
	}

	// return the value
	trx, err := p.LoadTransaction(hash)
	if err != nil {
		return nil, err
	}

	// push the transaction to the cache to speed things up next time
	// we don't cache pending transactions since it would cause issues
	// when re-loading data of such transactions on the client side
	if trx.BlockHash == nil {
		p.log.Debugf("pending transaction %s found", trx.Hash)
		return trx, nil
	}

	// store to cache
	p.cache.PushTransaction(trx)
	return trx, nil
}

// LoadTransaction returns a transaction at Opera blockchain
// by a hash loaded directly from the node.
func (p *proxy) LoadTransaction(hash *common.Hash) (*types.Transaction, error) {
	return p.rpc.Transaction(hash)
}

// SendTransaction sends raw signed and RLP encoded transaction to the block chain.
func (p *proxy) SendTransaction(tx hexutil.Bytes) (*types.Transaction, error) {
	p.log.Debugf("announcing trx %s", tx.String())

	// try to send it and get the tx hash
	hash, err := p.rpc.SendTransaction(tx)
	if err != nil {
		p.log.Errorf("can not send transaction to block chain; %s", err.Error())
		return nil, err
	}

	// check the hash makes sense by comparing it to empty hash
	if bytes.Compare(hash.Bytes(), common.Hash{}.Bytes()) == 0 {
		p.log.Criticalf("transaction not send; %s", tx.String())
		return nil, fmt.Errorf("transaction could not be send")
	}

	// we do have the hash, so we can use it to get the transaction details
	// we always need to go to RPC, and we will not try to store the transaction in cache yet
	trx, err := p.rpc.Transaction(hash)
	if err != nil {
		// transaction simply not found?
		if err == eth.ErrNoResult {
			p.log.Warning("transaction not found in the blockchain")
			return nil, ErrTransactionNotFound
		}

		// something went wrong
		return nil, err
	}

	// do we have the transaction we expected?
	if bytes.Compare(hash.Bytes(), trx.Hash.Bytes()) != 0 {
		p.log.Criticalf("transaction %s not confirmed, got %s", hash.String(), trx.Hash.String())
		return nil, fmt.Errorf("transaction %s could not be confirmed", hash.String())
	}

	// log transaction hash
	p.log.Noticef("trx %s from %s submitted", hash.String(), trx.From.String())
	return trx, nil
}

// Transactions pulls list of transaction hashes starting on the specified cursor.
// If the initial transaction cursor is not provided, we start on top, or bottom based on count value.
//
// No-number boundaries are handled as follows:
// 	- For positive count we start from the most recent transaction and scan to older transactions.
// 	- For negative count we start from the first transaction and scan to newer transactions.
func (p *proxy) Transactions(cursor *string, count int32) (*types.TransactionList, error) {
	// we may be able to pull the list faster than from the db
	if cursor == nil && count > 0 && count < cache.TransactionRingCacheSize {
		// pull the quick list
		tl := p.cache.ListTransactions(int(count))

		// does it make sense? if so, make the list from it
		if len(tl) > 0 {
			return &types.TransactionList{
				Collection: tl,
				Total:      uint64(p.MustEstimateTransactionsCount()),
				First:      tl[0].Uid(),
				Last:       tl[len(tl)-1].Uid(),
				IsStart:    true,
				IsEnd:      false,
				Filter:     nil,
			}, nil
		}
	}

	// use slow trx list pulling
	return p.db.Transactions(cursor, count, nil)
}

// StoreGasPricePeriod stores the given gas price period data in the persistent storage
func (p *proxy) StoreGasPricePeriod(gp *types.GasPricePeriod) error {
	return p.db.AddGasPricePeriod(gp)
}
