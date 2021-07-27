package cache

import (
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"strings"
)

// lastEpochCacheKey represents the in-memory cache key for the latest sealed Epoch data.
const epochCacheKey = "epoch"

// epochKey generates cache key for the given epoch number.
func epochKey(id *hexutil.Uint64) string {
	var sb strings.Builder
	sb.WriteString(epochCacheKey)
	sb.WriteString(id.String())
	return sb.String()
}

// PullEpoch extracts information about the given Epoch from the in-memory cache if available.
func (b *MemBridge) PullEpoch(id *hexutil.Uint64) *types.Epoch {
	// try to get the Epoch data from the cache
	data, err := b.cache.Get(epochKey(id))
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

// PushEpoch stores provided Epoch in the in-memory cache.
func (b *MemBridge) PushEpoch(ep *types.Epoch) {
	// we need valid account
	if nil == ep {
		return
	}

	// encode account
	data, err := ep.Marshal()
	if err != nil {
		b.log.Criticalf("can not marshal epoch to JSON; %s", err.Error())
		return
	}

	// set the data to cache by block number
	if err := b.cache.Set(epochKey(&ep.Id), data); err != nil {
		b.log.Errorf("can not cache epoch #%d; %s", ep.Id, err.Error())
	}
}
