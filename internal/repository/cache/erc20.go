// Package cache implements bridge to fast in-memory object cache.
package cache

import (
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"strings"
)

const erc20CacheIdPrefix = "erc20_"

// erc20TokenId generates cache id for storing ERC20 token details.
func erc20TokenId(addr *common.Address) string {
	var sb strings.Builder

	// add the prefix and actual address
	sb.WriteString(erc20CacheIdPrefix)
	sb.WriteString(addr.String())

	return sb.String()
}

// PullErc20Token extracts ERC20 token information from the in-memory cache if available.
func (b *MemBridge) PullErc20Token(addr *common.Address) *types.Erc20Token {
	// try to get the account data from the cache
	data, err := b.cache.Get(erc20TokenId(addr))
	if err != nil {
		// cache returns ErrEntryNotFound if the key does not exist
		return nil
	}

	// do we have the data?
	token, err := types.UnmarshalErc20Token(data)
	if err != nil {
		b.log.Criticalf("can not decode ERC20 data from in-memory cache; %s", err.Error())
		return nil
	}

	return token
}

// PushErc20Token stores provided ERC20 token in the in-memory cache.
func (b *MemBridge) PushErc20Token(token *types.Erc20Token) error {
	// we need valid account
	if nil == token {
		return fmt.Errorf("invalid or nil token can not be pushed to the in-memory cache")
	}

	// encode account
	data, err := token.Marshal()
	if err != nil {
		b.log.Criticalf("can not marshal ERC20 token to JSON; %s", err.Error())
		return err
	}

	// set the data to cache
	return b.cache.Set(erc20TokenId(&token.Address), data)
}
