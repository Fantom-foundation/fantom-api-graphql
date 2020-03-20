// Package cache implements bridge to fast in-memory object cache.
package cache

import (
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"strings"
)

// stiCacheKeyPrefix is the prefix used for cache key to store staker information.
const stiCacheKeyPrefix = "staker_info_"

// PullStakerInfo extracts staker information from the in-memory cache if available.
func (b *MemBridge) PullStakerInfo(id hexutil.Uint64) *types.StakerInfo {
	// try to get the account data from the cache
	data, err := b.cache.Get(getStakerInfoKey(id))
	if err != nil {
		// cache returns ErrEntryNotFound if the key does not exist
		return nil
	}

	// do we have the data?
	sti, err := types.UnmarshalStakerInfo(data)
	if err != nil {
		b.log.Criticalf("can not decode staker information data from in-memory cache; %s", err.Error())
		return nil
	}

	return sti
}

// PushStakerInfo stores provided staker information in the in-memory cache.
func (b *MemBridge) PushStakerInfo(id hexutil.Uint64, sti types.StakerInfo) error {
	// encode account
	data, err := sti.Marshal()
	if err != nil {
		b.log.Criticalf("can not marshal staker information to JSON; %s", err.Error())
		return err
	}

	// set the data to cache by block number
	return b.cache.Set(getStakerInfoKey(id), data)
}

// getPriceKeyBySymbol build a cache key for the given price symbol.
func getStakerInfoKey(id hexutil.Uint64) string {
	// use the builder to make the string
	var sb strings.Builder

	sb.WriteString(stiCacheKeyPrefix)
	sb.WriteString(id.String())

	return sb.String()
}
