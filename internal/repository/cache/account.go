// Cache package implements bridge to fast in-memory object cache.
package cache

import (
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
)

// PullAccount extracts account information from the in-memory cache if available.
func (b *Bridge) PullAccount(addr *common.Address) *types.Account {
	// try to get the account data from the cache
	data, err := b.cache.Get(addr.Hex())
	if err != nil {
		// cache returns ErrEntryNotFound if the key does not exist
		return nil
	}

	// do we have the data?
	acc, err := types.UnmarshalAccount(data)
	if err != nil {
		b.log.Criticalf("can not decode account data from in-memory cache; %s", err.Error())
		return nil
	}

	return &acc
}

// PushAccount stores provided account in the in-memory cache.
func (b *Bridge) PushAccount(acc *types.Account) error {
	// we need valid account
	if nil == acc {
		return fmt.Errorf("invalid or nil account can not be pushed to the in-memory cache")
	}

	// encode account
	data, err := acc.Marshal()
	if err != nil {
		b.log.Criticalf("can not marshal account to JSON; %s", err.Error())
		return err
	}

	// set the data to cache
	return b.cache.Set(acc.Address.Hex(), data)
}
