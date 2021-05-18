// Package cache implements bridge to fast in-memory object cache.
package cache

import (
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
)

// PullTransaction extracts transaction information from the in-memory cache if available.
func (b *MemBridge) PullTransaction(hash *common.Hash) *types.Transaction {
	// try to get the account data from the cache
	data, err := b.cache.Get(hash.String())
	if err != nil {
		// cache returns ErrEntryNotFound if the key does not exist
		return nil
	}

	// do we have the data?
	trx := new(types.Transaction)
	if err := trx.UnmarshalBSON(data); err != nil {
		b.log.Criticalf("can not decode transaction data from in-memory cache; %s", err.Error())
		return nil
	}
	return trx
}

// PushTransaction stores provided transaction in the in-memory cache.
func (b *MemBridge) PushTransaction(trx *types.Transaction) {
	// we need valid account
	if nil == trx {
		b.log.Errorf("undefined transaction can not be pushed to the in-memory cache")
		return
	}

	// encode account
	data, err := trx.MarshalBSON()
	if err != nil {
		b.log.Criticalf("can not marshal transaction %s; %s", trx.Hash.String(), err.Error())
		return
	}

	// set the data to cache by block number
	if err := b.cache.Set(trx.Hash.String(), data); err != nil {
		b.log.Criticalf("can not cache transaction %s; %s", trx.Hash.String(), err.Error())
	}
}
