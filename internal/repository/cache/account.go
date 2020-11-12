// Package cache implements bridge to fast in-memory object cache.
package cache

import (
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"strings"
)

const accountExistenceCacheIdPrefix = "acc_"

// accountId generates cache id for storing account details.
func accountId(addr *common.Address) string {
	var sb strings.Builder

	// add the prefix and actual address
	sb.WriteString(accountExistenceCacheIdPrefix)
	sb.WriteString(addr.String())

	return sb.String()
}

// PullAccount extracts account information from the in-memory cache if available.
func (b *MemBridge) PullAccount(addr *common.Address) *types.Account {
	// try to get the account data from the cache
	data, err := b.cache.Get(accountId(addr))
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

	return acc
}

// PushAccount stores provided account in the in-memory cache.
func (b *MemBridge) PushAccount(acc *types.Account) error {
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
	return b.cache.Set(accountId(&acc.Address), data)
}

// CheckAccountKnown verifies if the cache is aware of the account existence
// in the database.
func (b *MemBridge) CheckAccountKnown(addr *common.Address) *bool {
	// try to get the account data from the cache
	data, err := b.cache.Get(addr.Hex())
	if err != nil {
		return nil
	}

	// check if the account is known
	val := int(data[0])&1 > 0
	return &val
}

// PushAccountKnown caches the known account state.
func (b *MemBridge) PushAccountKnown(addr *common.Address) {
	// cache the account existence
	err := b.cache.Set(addr.Hex(), []byte{1})
	if err != nil {
		b.log.Errorf("can not cache account %s existence; %s", addr.String(), err.Error())
	}
}
