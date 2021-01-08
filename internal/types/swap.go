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

	// Type represents type of event: swap, mint, burn, sync
	Type int `json:"type" bson:"type"`

	// TimeStamp represents time of the swap.
	TimeStamp *hexutil.Uint64 `json:"timestamp" bson:"timestamp"`

	// Pair represents address of the swapped pair.
	Pair common.Address `json:"pair" bson:"pair"`

	// Sender represents address of the sender.
	Sender common.Address `json:"sender" bson:"sender"`

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

	// Reserve0 is a total reserve in time of this event for Token A
	Reserve0 *big.Int `json:"reserve0" bson:"reserve0"`

	// Reserve1 is a total reserve in time of this event for Token B
	Reserve1 *big.Int `json:"reserve1" bson:"reserve1"`
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

// DefiSwapVolume represents a calculated volume for swap pairs in history
type DefiSwapVolume struct {

	// PairAddress is an address of the listed pair
	PairAddress *common.Address

	// Volume is a swap voolume for specified period
	Volume *big.Int

	// IsInFUSD true if volumes can be denominated to fUSD
	IsInFUSD bool

	// DateString represents time tag for this volume
	DateString string
}

// DefiTimePrice represents a calculated price for swap pairs in history
type DefiTimePrice struct {

	// PairAddress is an address of the listed pair
	PairAddress common.Address `json:"address" bson:"pair"`

	// Time represents ISO time tag for this price
	Time string `json:"time" bson:"_id"`

	// opening price for this time period
	Open float64 `json:"open" bson:"open"`

	// closing price for this time period
	Close float64 `json:"close" bson:"close"`

	// lowest price for this time period
	Low float64 `json:"low" bson:"low"`

	// highest price for this time period
	High float64 `json:"high" bson:"high"`

	// average price for this time period
	Average float64 `json:"average" bson:"avg"`
}

// Swap types
const (
	SwapNormal = iota
	SwapMint
	SwapBurn
	SwapSync
)

// DefiTimeReserve represents a reserve for uniswap pair in history
type DefiTimeReserve struct {

	// Time represents UTC ISO time tag for this reserve values
	Time string

	// ReserveClose is close reserve for this time period
	// for both tokens
	ReserveClose []hexutil.Big
}
