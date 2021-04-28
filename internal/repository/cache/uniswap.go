// Package cache implements bridge to fast in-memory object cache.
package cache

import (
	"github.com/ethereum/go-ethereum/common"
	"strings"
)

// uniswapPairTokensPrefix represents a prefix used for uniswap pair tokens caching key.
const uniswapPairTokensPrefix = "unp"

// uniswapPairTokensKey generates cache key for uniswap tokens pair entry.
func uniswapPairTokensKey(pair *common.Address) string {
	var sb strings.Builder
	sb.WriteString(uniswapPairTokensPrefix)
	sb.WriteString(pair.String())
	return sb.String()
}

// PushGovernanceTotalWeight stores governance total weight information
// in the in-memory cache.
func (b *MemBridge) PushUniswapPairTokens(pair *common.Address, tl []common.Address) {
	// nothing to store or bad data?
	if pair == nil || tl == nil || 2 != len(tl) {
		return
	}

	// set the data to cache
	// just stick those addresses together into a single slice of 40 bytes, 20 for each address
	// we don't need to worry about rewriting tl[0] bytes slice by the append call
	// since it's strictly 20 bytes long and so a new underlying array will be allocated by append()
	if err := b.cache.Set(uniswapPairTokensKey(pair), append(tl[0].Bytes(), tl[1].Bytes()...)); err != nil {
		b.log.Errorf("can not store uniswap pair %s tokens; %s", pair.String(), err.Error())
	}
}

// PullUniswapPairTokens tries to load a uniswap pair tokens from the cache.
func (b *MemBridge) PullUniswapPairTokens(pair *common.Address) []common.Address {
	// nothing to load if the pair is not given
	if pair == nil {
		return nil
	}

	// try to get the data from cache
	data, err := b.cache.Get(uniswapPairTokensKey(pair))
	if err != nil {
		// cache returns ErrEntryNotFound if the key does not exist
		return nil
	}

	// restore addresses from received raw data
	// bytes <0..19> belong to the first address; bytes 20+ are for the second one
	tl := make([]common.Address, 2)
	tl[0].SetBytes(data[:20])
	tl[1].SetBytes(data[20:])
	return tl
}
