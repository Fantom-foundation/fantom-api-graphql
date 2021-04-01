// Package cache implements bridge to fast in-memory object cache.
package cache

import (
	"math/big"
)

// sfcMaxDelegatedRatioKey represents the key used to store SFC delegation ratio.
const sfcMaxDelegatedRatioKey = "sfc_dlg_ratio"

// PullSfcMaxDelegatedRatio extract the ratio from cache, if possible.
func (b *MemBridge) PullSfcMaxDelegatedRatio() *big.Int {
	// try to get the account data from the cache
	data, err := b.cache.Get(sfcMaxDelegatedRatioKey)
	if err != nil {
		return nil
	}
	return new(big.Int).SetBytes(data)
}

// PushSfcMaxDelegatedRatio stores the ratio in cache, if possible.
func (b *MemBridge) PushSfcMaxDelegatedRatio(val *big.Int) {
	if val == nil {
		return
	}
	if err := b.cache.Set(sfcMaxDelegatedRatioKey, val.Bytes()); err != nil {
		b.log.Errorf("can not store SFC delegation ratio value")
	}
}
