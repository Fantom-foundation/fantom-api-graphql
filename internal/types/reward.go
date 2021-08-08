// Package types implements different core types of the API.
package types

import (
	"encoding/binary"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson"
	"math/big"
	"time"
)

const (
	FiRewardClaimPk          = "_id"
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
	Delegator     common.Address
	ToValidatorId hexutil.Big
	Claimed       hexutil.Uint64
	ClaimTrx      common.Hash
	Amount        hexutil.Big
	IsDelegated   bool
}

// BsonRewardClaim represents BSON rew structure of the reward claim.
type BsonRewardClaim struct {
	ID        string    `bson:"_id"`
	Ordinal   uint64    `bson:"orx"`
	Addr      string    `bson:"addr"`
	To        string    `bson:"to"`
	ClaimTime uint64    `bson:"when"`
	TimeStamp time.Time `bson:"stamp"`
	Amount    string    `bson:"amount"`
	Value     uint64    `bson:"value"`
	IsDlg     bool      `bson:"red"`
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
	val := new(big.Int).Div(rwc.Amount.ToInt(), RewardDecimalsCorrection)

	// prep the structure for saving
	pom := BsonRewardClaim{
		ID:        rwc.Pk(),
		Ordinal:   rwc.OrdinalIndex(),
		Addr:      rwc.Delegator.String(),
		To:        rwc.ToValidatorId.String(),
		ClaimTime: uint64(rwc.Claimed),
		TimeStamp: time.Unix(int64(rwc.Claimed), 0),
		Amount:    rwc.Amount.String(),
		Value:     val.Uint64(),
		IsDlg:     rwc.IsDelegated,
	}
	return bson.Marshal(pom)
}

// UnmarshalBSON updates the value from BSON source.
func (rwc *RewardClaim) UnmarshalBSON(data []byte) (err error) {
	// capture unmarshal issue
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("can not decode and unmarshal")
		}
	}()

	// try to decode the BSON data
	var row BsonRewardClaim
	if err = bson.Unmarshal(data, &row); err != nil {
		return err
	}

	// transfer values
	rwc.Delegator = common.HexToAddress(row.Addr)
	rwc.ToValidatorId = (hexutil.Big)(*hexutil.MustDecodeBig(row.To))
	rwc.Claimed = hexutil.Uint64(row.ClaimTime)
	rwc.ClaimTrx = common.HexToHash(row.ID)
	rwc.Amount = (hexutil.Big)(*hexutil.MustDecodeBig(row.Amount))
	rwc.IsDelegated = row.IsDlg
	return nil
}
