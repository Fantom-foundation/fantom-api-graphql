// Package types implements different core types of the API.
package types

import "github.com/ethereum/go-ethereum/common/hexutil"

// PendingRewards represents a rewards waiting to be claimed structure.
type PendingRewards struct {
	Staker hexutil.Uint64
	Amount hexutil.Big
}
