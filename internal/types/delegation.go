// Package types implements different core types of the API.
package types

import (
	"encoding/binary"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"time"
)

// DelegationDecimalsCorrection is used to adjust decimal precision of a delegation active value.
var DelegationDecimalsCorrection = new(big.Int).SetUint64(1000000000)

// Delegation represents a delegator in Opera blockchain.
type Delegation struct {
	ID              string         `json:"id" bson:"_id"`
	Address         common.Address `json:"addr" bson:"adr"`
	ToStakerId      *hexutil.Big   `json:"to_id" bson:"to"`
	Transaction     common.Hash    `json:"trx" bson:"trx"`
	ToStakerAddress common.Address `json:"to_addr" bson:"toad"`
	Created         time.Time      `json:"created" bson:"stamp"`
	Index           int64          `bson:"orx"`
	Value           int64          `bson:"value"`

	// AmountStaked represents the current staked amount
	AmountStaked *hexutil.Big `json:"amount" bson:"amo"`

	// AmountDelegated is the original amount delegated
	AmountDelegated *hexutil.Big `json:"initial" bson:"act"`
}

// OrdinalIndex returns an ordinal index for the given delegation.
// We construct the UID from the time the delegation was created (40 bits = 1099511627775s = 34000 years),
// a part of the creation transaction hash and part of the target validator index (12 bits = 4096).
func (dl *Delegation) OrdinalIndex() int64 {
	return int64((uint64(dl.Created.Unix())&0x7FFFFFFFFF)<<24 | (dl.ToStakerId.ToInt().Uint64()&0xFFF)<<12 | (binary.BigEndian.Uint64(dl.Transaction[:8]) & 0xFFF))
}

// WithOrdinalIndex updates an ordinal index (field for deterministic sorting) for the given token transaction.
// We construct the index from the time the transaction was processed (40 bits = 1099511627775s = 34000 years),
// salted by the transaction hash, the event log index (index of the log in the block)
// and sequence number of transaction in batch transfer event.
func (dl *Delegation) WithOrdinalIndex() *Delegation {
	dl.Index = dl.OrdinalIndex()
	return dl
}

// WithDelegatedValue updates current delegated value on the given delegation.
func (dl *Delegation) WithDelegatedValue() *Delegation {
	dl.Value = new(big.Int).Div(dl.AmountDelegated.ToInt(), DelegationDecimalsCorrection).Int64()
	return dl
}
