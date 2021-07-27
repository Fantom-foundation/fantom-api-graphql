// Package cache implements bridge to fast in-memory object cache.
package cache

import (
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/allegro/bigcache"
	"github.com/ethereum/go-ethereum/common"
	"strings"
)

// contractCacheIdPrefix is used to identify smart contracts inside in-memory cache
const contractCacheIdPrefix = "sc_"

// contractId generates cache id for storing smart contract details.
func contractId(addr *common.Address) string {
	var sb strings.Builder

	// add the prefix and actual address
	sb.WriteString(contractCacheIdPrefix)
	sb.WriteString(addr.String())

	return sb.String()
}

// PullContract extracts smart contract information from the in-memory cache if available.
func (b *MemBridge) PullContract(addr *common.Address) *types.Contract {
	// try to get the account data from the cache
	data, err := b.cache.Get(contractId(addr))
	if err != nil {
		// cache returns ErrEntryNotFound if the key does not exist
		return nil
	}

	// do we have the data?
	sc, err := types.UnmarshalContract(data)
	if err != nil {
		b.log.Criticalf("can not decode contract data from in-memory cache; %s", err.Error())
		return nil
	}

	return sc
}

// PushContract stores provided contract in the in-memory cache.
func (b *MemBridge) PushContract(sc *types.Contract) error {
	// we need valid account
	if nil == sc {
		return fmt.Errorf("invalid or nil contract can not be pushed to the in-memory cache")
	}

	// encode contract
	data, err := sc.Marshal()
	if err != nil {
		b.log.Criticalf("can not marshal contract to JSON; %s", err.Error())
		return err
	}

	// set the data to cache
	return b.cache.Set(contractId(&sc.Address), data)
}

// EvictContract makes sure the contract of the given address
// is not kept in the cache.
func (b *MemBridge) EvictContract(addr *common.Address) {
	// make sure the request makes sense
	if nil == addr {
		b.log.Errorf("can not evict empty contract")
		return
	}

	// delete the record, if the is any
	err := b.cache.Delete(contractId(addr))
	if err != nil && err != bigcache.ErrEntryNotFound {
		b.log.Criticalf("cache error %s", err.Error())
	}
}
