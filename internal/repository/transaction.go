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
	p.log.Infof("requested transaction %s", hash.String())

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
	p.log.Infof("transaction %s loaded from rpc", hash.String())

	// push the transaction to the cache to speed things up next time
	err = p.cache.PushTransaction(trx)
	if err != nil {
		p.log.Errorf("can not store transaction in cache; %s", err.Error())
	}

	// return the value
	return trx, nil
}
