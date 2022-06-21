// Package types implements different core types of the API.
package types

import (
	"encoding/binary"
	"fantom-api-graphql/internal/repository/db/registry"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"time"
)

const (
	FiRewardClaimOrdinal     = "orx"
	FiRewardClaimAddress     = "addr"
	FiRewardClaimToValidator = "to"
	FiRewardClaimedValue     = "value"
	FiRewardClaimedTimeStamp = "stamp"
)

// RewardDecimalsCorrection is used to manipulate precision of a rewards value,
// so it can be stored in database as UINT64 without loosing too much data
var RewardDecimalsCorrection = new(big.Int).SetUint64(1000000000)

// RewardClaim represents a reward claim record in Opera staking
// SFC contract. A reward can be claimed directly towards the account balance,
// or it can be re-staked into the SFC contract as an increased delegation.
type RewardClaim struct {
	ClaimTrx      common.Hash    `bson:"_id"`
	Delegator     common.Address `bson:"addr"`
	ToValidatorId hexutil.Big    `bson:"to"`
	Claimed       hexutil.Uint64 `bson:"when"`
	Amount        hexutil.Big    `bson:"amount"`
	IsDelegated   bool           `bson:"red"`
	TimeStamp     time.Time      `bson:"stamp"`
	Value         uint64         `bson:"value"`
	Ordinal       uint64         `bson:"orx"`
}

// Pk returns a unique primary key of the claim.
func (rwc *RewardClaim) Pk() string {
	return rwc.ClaimTrx.String()
}

// OrdinalIndex returns an ordinal index for the given reward claim request.
func (rwc *RewardClaim) OrdinalIndex() uint64 {
	return (uint64(rwc.Claimed)&0x7FFFFFFFFF)<<24 | (binary.BigEndian.Uint64(rwc.ClaimTrx[:8]) & 0xFFFFFF)
}

// MarshalBSON creates a BSON representation of the reward claim request record.
func (rwc *RewardClaim) MarshalBSON() ([]byte, error) {
	// calculate the value to 9 digits (and 18 billions remain available)
	rwc.Value = new(big.Int).Div(rwc.Amount.ToInt(), RewardDecimalsCorrection).Uint64()
	rwc.Ordinal = rwc.OrdinalIndex()
	rwc.TimeStamp = time.Unix(int64(rwc.Claimed), 0)

	return registry.Marshal(rwc)
}
