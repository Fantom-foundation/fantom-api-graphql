/*
Package repository implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Opera/Lachesis full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import (
	"errors"
	"fantom-api-graphql/internal/types"
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

// Transaction returns a transaction at Opera blockchain by a hash, nil if not found.
// If the transaction is not found, ErrTransactionNotFound error is returned.
func (p *proxy) Transaction(hash *common.Hash) (*types.Transaction, error) {
	// log
	p.log.Debugf("requested transaction %s", hash.String())

	// try to use the in-memory cache
	if trx := p.cache.PullTransaction(hash); trx != nil {
		// log and return
		p.log.Debugf("transaction %s loaded from cache", hash.String())
		return trx, nil
	}

	// return the value
	return p.loadTransaction(hash)
}

// loadTransaction loads the transaction hard way using RPC and DB.
func (p *proxy) loadTransaction(hash *common.Hash) (*types.Transaction, error) {
	// we need to go to RPC
	trx, err := p.rpc.Transaction(hash)
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

// SendTransaction sends raw signed and RLP encoded transaction to the block chain.
func (p *proxy) SendTransaction(tx hexutil.Bytes) (*types.Transaction, error) {
	// log
	p.log.Debugf("requested transaction submit for %s", tx.String())

	// try to send it and get the tx hash
	hash, err := p.rpc.SendTransaction(tx)
	if err != nil {
		p.log.Errorf("can not send transaction to block chain; %s", err.Error())
		return nil, err
	}

	// we do have the hash so we can use it to get the transaction details
	// we always need to go to RPC and we will not try to store the transaction in cache yet
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

	// log and return
	p.log.Debugf("transaction %s created and found", hash.String())

	// return the value
	return trx, nil
}

// Transactions pulls list of transaction hashes starting on the specified cursor.
// If the initial transaction cursor is not provided, we start on top, or bottom based on count value.
//
// No-number boundaries are handled as follows:
// 	- For positive count we start from the most recent transaction and scan to older transactions.
// 	- For negative count we start from the first transaction and scan to newer transactions.
func (p *proxy) Transactions(cursor *string, count int32) (*types.TransactionList, error) {
	// go to the database for the list of hashes of transaction searched
	return p.db.Transactions(cursor, count, nil)
}

// TransactionsCount returns total number of transactions in the block chain.
func (p *proxy) TransactionsCount() (hexutil.Uint64, error) {
	// get the number of transactions registered
	tc, err := p.db.TransactionsCount()
	if err != nil {
		return hexutil.Uint64(0), err
	}

	return hexutil.Uint64(tc), nil
}
