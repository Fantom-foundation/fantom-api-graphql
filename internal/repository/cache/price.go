// Package cache implements bridge to fast in-memory object cache.
package cache

import (
	"fantom-api-graphql/internal/types"
	"fmt"
	"strings"
)

// priceCacheKeyPrefix is the prefix used for cache key to store price information.
const priceCacheKeyPrefix = "price_FTM_2_"

// PullPrice extracts price information from the in-memory cache if available.
func (b *MemBridge) PullPrice(symbol string) *types.Price {
	// try to get the account data from the cache
	data, err := b.cache.Get(getPriceKeyBySymbol(symbol))
	if err != nil {
		// cache returns ErrEntryNotFound if the key does not exist
		return nil
	}

	// do we have the data?
	pri, err := types.UnmarshalPrice(data)
	if err != nil {
		b.log.Criticalf("can not decode price data from in-memory cache; %s", err.Error())
		return nil
	}

	return &pri
}

// PushPrice stores provided price in the in-memory cache.
func (b *MemBridge) PushPrice(sym string, pri *types.Price) error {
	// we need valid price
	if nil == pri {
		return fmt.Errorf("undefined price can not be pushed to the in-memory cache")
	}

	// encode account
	data, err := pri.Marshal()
	if err != nil {
		b.log.Criticalf("can not marshal price to JSON; %s", err.Error())
		return err
	}

	// set the data to cache by block number
	return b.cache.Set(getPriceKeyBySymbol(sym), data)
}

// getPriceKeyBySymbol build a cache key for the given price symbol.
func getPriceKeyBySymbol(sym string) string {
	// use the builder to make the string
	var sb strings.Builder

	sb.WriteString(priceCacheKeyPrefix)
	sb.WriteString(sym)

	return sb.String()
}
