// Package types implements different core types of the API.
package types

import (
	"encoding/binary"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// RewardDecimalsCorrection is used to manipulate precision of a rewards value,
// so it can be stored in database as UINT64 without loosing too much data
var RewardDecimalsCorrection = new(big.Int).SetUint64(1000000000)

// RewardClaim represents a reward claim record in Opera staking
// SFC contract. A reward can be claimed directly towards the account balance,
// or it can be re-staked into the SFC contract as an increased delegation.
type RewardClaim struct {
	ClaimTrx         common.Hash    `bson:"_id"`
	Delegator        common.Address `bson:"addr"`
	ToValidatorId    hexutil.Big    `bson:"to"`
	ClaimedTimeStamp uint64         `bson:"when"`
	Amount           hexutil.Big    `bson:"amount"`
	IsDelegated      bool           `bson:"red"`
	Value            uint64         `bson:"value"`
	Ordinal          uint64         `bson:"orx"`
}

// ComputedOrdinalIndex returns an ordinal index for the given reward claim request.
func (rwc *RewardClaim) ComputedOrdinalIndex() uint64 {
	return (uint64(rwc.ClaimedTimeStamp)&0x7FFFFFFFFF)<<24 | (binary.BigEndian.Uint64(rwc.ClaimTrx[:8]) & 0xFFFFFF)
}

// ComputedValue returns computed value.
func (rwc *RewardClaim) ComputedValue() uint64 {
	return new(big.Int).Div(rwc.Amount.ToInt(), RewardDecimalsCorrection).Uint64()
}
