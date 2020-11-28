// Package cache implements bridge to fast in-memory object cache.
package cache

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"strings"
)

const (
	governanceTotalWeightPrefix = "gtw_"
)

// getGovernanceTotalWeightKey build a cache key
// for the given governance total weight value.
func getGovernanceTotalWeightKey(addr *common.Address) string {
	var sb strings.Builder
	// add the prefix and than the address
	sb.WriteString(governanceTotalWeightPrefix)
	sb.WriteString(addr.String())

	return sb.String()
}

// PullGovernanceTotalWeight extracts governance total weight information
// from the in-memory cache if available.
func (b *MemBridge) PullGovernanceTotalWeight(gov *common.Address) *hexutil.Big {
	// nothing to do
	if gov == nil {
		return nil
	}

	// try to get the account data from the cache
	data, err := b.cache.Get(getGovernanceTotalWeightKey(gov))
	if err != nil {
		// cache returns ErrEntryNotFound if the key does not exist
		return nil
	}

	// do we have the data?
	pri := new(hexutil.Big)
	if err := pri.UnmarshalText(data); err != nil {
		b.log.Criticalf("can not decode governance total weight; %s", err.Error())
		return nil
	}

	return pri
}

// PushGovernanceTotalWeight stores governance total weight information
// in the in-memory cache.
func (b *MemBridge) PushGovernanceTotalWeight(gov *common.Address, val *hexutil.Big) error {
	// nothing to store
	if gov == nil || val == nil {
		return nil
	}

	// encode the value
	data, err := val.MarshalText()
	if err != nil {
		b.log.Criticalf("can not encode governance total weight; %s", err.Error())
		return err
	}

	// set the data to cache by address
	return b.cache.Set(getGovernanceTotalWeightKey(gov), data)
}
