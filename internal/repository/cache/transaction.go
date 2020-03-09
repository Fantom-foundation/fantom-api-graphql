// Cache package implements bridge to fast in-memory object cache.
package cache

import (
	"fantom-api-graphql/internal/types"
	"fmt"
)

// PullTransaction extracts transaction information from the in-memory cache if available.
func (b *MemBridge) PullTransaction(hash *types.Hash) *types.Transaction {
	// try to get the account data from the cache
	data, err := b.cache.Get(hash.String())
	if err != nil {
		// cache returns ErrEntryNotFound if the key does not exist
		return nil
	}

	// do we have the data?
	trx, err := types.UnmarshalTransaction(data)
	if err != nil {
		b.log.Criticalf("can not decode transaction data from in-memory cache; %s", err.Error())
		return nil
	}

	return trx
}

// PushTransaction stores provided transaction in the in-memory cache.
func (b *MemBridge) PushTransaction(trx *types.Transaction) error {
	// we need valid account
	if nil == trx {
		return fmt.Errorf("undefined transaction can not be pushed to the in-memory cache")
	}

	// encode account
	data, err := trx.Marshal()
	if err != nil {
		b.log.Criticalf("can not marshal transaction to JSON; %s", err.Error())
		return err
	}

	// set the data to cache by block number
	return b.cache.Set(trx.Hash.String(), data)
}
