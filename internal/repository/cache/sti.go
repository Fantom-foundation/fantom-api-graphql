// Package cache implements bridge to fast in-memory object cache.
package cache

import (
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"strings"
)

// stiCacheKeyPrefix is the prefix used for cache key to store staker information.
const stiCacheKeyPrefix = "staker_info_"

// stiTotalStakedKey is the cache key used to store total staked amount
const stiTotalStakedKey = "staked_total"

// PullStakerInfo extracts staker information from the in-memory cache if available.
func (b *MemBridge) PullStakerInfo(id *hexutil.Big) *types.StakerInfo {
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
func (b *MemBridge) PushStakerInfo(id *hexutil.Big, sti *types.StakerInfo) error {
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
func getStakerInfoKey(id *hexutil.Big) string {
	// use the builder to make the string
	var sb strings.Builder

	sb.WriteString(stiCacheKeyPrefix)
	sb.WriteString(id.String())

	return sb.String()
}

// PullTotalStaked extracts total staked amount from the in-memory cache if available.
func (b *MemBridge) PullTotalStaked() *hexutil.Big {
	// try to get the account data from the cache
	data, err := b.cache.Get(stiTotalStakedKey)
	if err != nil {
		return nil
	}

	// do we have the data?
	val := new(big.Int).SetBytes(data)
	return (*hexutil.Big)(val)
}

// PushTotalStaked stores provided total staked amount information in the in-memory cache.
func (b *MemBridge) PushTotalStaked(amount *hexutil.Big) error {
	// we must have the value
	if amount == nil {
		b.log.Criticalf("can not store invalid amount")
		return fmt.Errorf("amount not provided")
	}

	// encode account
	return b.cache.Set(stiTotalStakedKey, amount.ToInt().Bytes())
}
