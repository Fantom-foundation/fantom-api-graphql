// Package types implements different core types of the API.
package types

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Swap represents a basic information provided by the API about finished swap from Uniswap contract.
type Swap struct {

	// BlockNumber represents number of the block where this transaction was in. nil when its pending.
	BlockNumber *hexutil.Uint64 `json:"blockNumber" bson:"blk"`

	// Pair represents address of the swapped pair.
	TimeStamp *hexutil.Uint64 `json:"timestamp" bson:"timestamp"`

	// Pair represents address of the swapped pair.
	Pair common.Address `json:"pair" bson:"pair"`

	// Sender represents address of the sender.
	Sender common.Address `json:"sender" bson:"sender"`

	// To represents address of the receiver.
	To common.Address `json:"to" bson:"to"`

	// Hash represents 32 bytes hash of the transaction.
	Hash Hash `json:"tx" bson:"tx"`

	// Amount0In represents integer of the swap amount.
	Amount0In *big.Int `json:"amount0In" bson:"am0in"`

	// Amount0Out represents integer of the swap amount.
	Amount0Out *big.Int `json:"amount0Out" bson:"am0out"`

	// Amount1In represents integer of the swap amount.
	Amount1In *big.Int `json:"amount1In" bson:"am1in"`

	// Amount1Out represents integer of the swap amount.
	Amount1Out *big.Int `json:"amount1Out" bson:"am1out"`
}

// UnmarshalSwap parses the JSON-encoded block data.
func UnmarshalSwap(data []byte) (*Swap, error) {
	var swap Swap
	err := json.Unmarshal(data, &swap)
	return &swap, err
}

// Marshal returns the JSON encoding of swap.
func (swap *Swap) Marshal() ([]byte, error) {
	return json.Marshal(swap)
}
