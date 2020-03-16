package types

import "github.com/ethereum/go-ethereum/common/hexutil"

// Rewards represents a rewards waiting to be paid structure.
type PendingRewards struct {
	Amount    hexutil.Big
	FromEpoch hexutil.Uint64
	ToEpoch   hexutil.Uint64
}
