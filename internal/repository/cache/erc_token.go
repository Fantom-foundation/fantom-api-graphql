// Package cache implements bridge to fast in-memory object cache.
package cache

import (
	"encoding/json"
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"strings"
)

const (
	Erc20CacheIdPrefix   = "erc20_"
	Erc721CacheIdPrefix  = "erc721_"
)

// ErcTokenId generates cache id for storing ERC token contract.
func ErcTokenId(addr *common.Address, tp string) string {
	var sb strings.Builder

	// add the prefix and actual address
	sb.WriteString(tp)
	sb.WriteString(addr.String())

	return sb.String()
}

// PullErc20Token extracts ERC20 token information from the in-memory cache if available.
func (b *MemBridge) PullErc20Token(addr *common.Address) *types.Erc20Token {
	// try to get the account data from the cache
	data, err := b.cache.Get(ErcTokenId(addr, Erc20CacheIdPrefix))
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
	return b.cache.Set(ErcTokenId(&token.Address, Erc20CacheIdPrefix), data)
}

// PullErc721Contract pulls ERC-721 token contract details from cache, if available.
func (b *MemBridge) PullErc721Contract(addr *common.Address) *types.Erc721Contract {
	// try to get the account data from the cache
	data, err := b.cache.Get(ErcTokenId(addr, Erc721CacheIdPrefix))
	if err != nil {
		return nil
	}

	// do we have the data?
	var token types.Erc721Contract
	err = json.Unmarshal(data, &token)
	if err != nil {
		b.log.Criticalf("can not decode ERC721 data from in-memory cache; %s", err.Error())
		return nil
	}
	return &token
}

// PushErc721Contract stores provided ERC-721 contract details in memory cache.
func (b *MemBridge) PushErc721Contract(tok *types.Erc721Contract) error {
	if nil == tok {
		return fmt.Errorf("invalid or nil token can not be pushed to the in-memory cache")
	}

	// encode account
	data, err := json.Marshal(tok)
	if err != nil {
		b.log.Criticalf("can not marshal ERC721 token to JSON; %s", err.Error())
		return err
	}
	return b.cache.Set(ErcTokenId(&tok.Address, Erc721CacheIdPrefix), data)
}
