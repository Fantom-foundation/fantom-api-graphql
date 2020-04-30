package cache

import (
	"fantom-api-graphql/internal/types"
	"fmt"
)

// lastEpochCacheKey represents the in-memory cache key for the latest sealed Epoch data.
const lastEpochCacheKey = "last-sealed-epoch"

// PullLastEpoch extracts information about the latest Epoch from the in-memory cache if available.
func (b *MemBridge) PullLastEpoch() *types.Epoch {
	// try to get the Epoch data from the cache
	data, err := b.cache.Get(lastEpochCacheKey)
	if err != nil {
		// cache returns ErrEntryNotFound if the key does not exist
		return nil
	}

	// do we have the data?
	ep, err := types.UnmarshalEpoch(data)
	if err != nil {
		b.log.Criticalf("can not decode epoch data from in-memory cache; %s", err.Error())
		return nil
	}

	return ep
}

// PushLastEpoch stores provided latest sealed Epoch in the in-memory cache.
func (b *MemBridge) PushLastEpoch(ep *types.Epoch) error {
	// we need valid account
	if nil == ep {
		return fmt.Errorf("undefined epoch can not be pushed to the in-memory cache")
	}

	// encode account
	data, err := ep.Marshal()
	if err != nil {
		b.log.Criticalf("can not marshal epoch to JSON; %s", err.Error())
		return err
	}

	// set the data to cache by block number
	return b.cache.Set(lastEpochCacheKey, data)
}
