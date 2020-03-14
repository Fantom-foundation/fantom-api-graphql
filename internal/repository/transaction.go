/*
Repository package implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Opera/Lachesis full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import (
	"errors"
	"fantom-api-graphql/internal/types"
	eth "github.com/ethereum/go-ethereum/rpc"
)

// ErrTransactionNotFound represents an error returned if a transaction can not be found.
var ErrTransactionNotFound = errors.New("requested transaction can not be found in Opera blockchain")

// Transaction returns a transaction at Opera blockchain by a hash, nil if not found.
// If the transaction is not found, ErrTransactionNotFound error is returned.
func (p *proxy) Transaction(hash *types.Hash) (*types.Transaction, error) {
	// log
	p.log.Debugf("requested transaction %s", hash.String())

	// try to use the in-memory cache
	if trx := p.cache.PullTransaction(hash); trx != nil {
		// log and return
		p.log.Infof("transaction %s loaded from cache", hash.String())
		return trx, nil
	}

	// we need to go to RPC
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
	p.log.Debugf("transaction %s loaded from rpc", hash.String())

	// push the transaction to the cache to speed things up next time
	err = p.cache.PushTransaction(trx)
	if err != nil {
		p.log.Errorf("can not store transaction in cache; %s", err.Error())
	}

	// return the value
	return trx, nil
}

// Transactions pulls list of transaction hashes starting on the specified cursor.
// If the initial transaction cursor is not provided, we start on top, or bottom based on count value.
//
// No-number boundaries are handled as follows:
// 	- For positive count we start from the most recent transaction and scan to older transactions.
// 	- For negative count we start from the first transaction and scan to newer transactions.
func (p *proxy) Transactions(cursor *string, count int32) (*types.TransactionHashList, error) {
	// go to the database for the list of hashes of transaction searched
	return p.db.Transactions(cursor, count)
}
