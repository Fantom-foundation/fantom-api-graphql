// Package types implements different core types of the API.
package types

import (
	"encoding/binary"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson"
	"math/big"
)

// RewardDecimalsCorrection is used to manipulate precision of a rewards value
// so it can be stored in database as UINT64 without loosing too much data
var RewardDecimalsCorrection = new(big.Int).SetUint64(1000000000)

// RewardClaim represents a reward claim record in Opera staking
// SFC contract. A reward can be claimed directly towards the account balance
// or it can be re-staked into the SFC contract as an increased delegation.
type RewardClaim struct {
	Delegator     common.Address
	ToValidatorId hexutil.Big
	Claimed       hexutil.Uint64
	ClaimTrx      common.Hash
	Amount        hexutil.Big
	IsDelegated   bool
}

// Uid returns a unique identifier for the given reward claim request.
func (rwc *RewardClaim) Uid() uint64 {
	return (uint64(rwc.Claimed)&0xFFFFFFFFFF)<<24 | (rwc.ToValidatorId.ToInt().Uint64()&0xFFF)<<12 | (binary.BigEndian.Uint64(rwc.ClaimTrx[:8]) & 0xFFF)
}

// MarshalBSON creates a BSON representation of the reward claim request record.
func (rwc *RewardClaim) MarshalBSON() ([]byte, error) {
	// calculate the value to 9 digits (and 18 billions remain available)
	val := new(big.Int).Div(rwc.Amount.ToInt(), RewardDecimalsCorrection)

	// prep the structure for saving
	pom := struct {
		Uid       uint64 `bson:"_id"`
		Addr      string `bson:"addr"`
		To        string `bson:"to"`
		ClaimTime uint64 `bson:"when"`
		Trx       string `bson:"trx"`
		Amount    string `bson:"amount"`
		Value     uint64 `bson:"value"`
		IsDlg     bool   `bson:"red"`
	}{
		Uid:       rwc.Uid(),
		Addr:      rwc.Delegator.String(),
		To:        rwc.ToValidatorId.String(),
		ClaimTime: uint64(rwc.Claimed),
		Trx:       rwc.ClaimTrx.String(),
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
	var row struct {
		Uid       uint64 `bson:"_id"`
		Addr      string `bson:"addr"`
		To        string `bson:"to"`
		ClaimTime uint64 `bson:"when"`
		Trx       string `bson:"trx"`
		Amount    string `bson:"amount"`
		Value     uint64 `bson:"value"`
		IsDlg     bool   `bson:"red"`
	}
	if err = bson.Unmarshal(data, &row); err != nil {
		return err
	}

	// transfer values
	rwc.Delegator = common.HexToAddress(row.Addr)
	rwc.ToValidatorId = (hexutil.Big)(*hexutil.MustDecodeBig(row.To))
	rwc.Claimed = hexutil.Uint64(row.ClaimTime)
	rwc.ClaimTrx = common.HexToHash(row.Trx)
	rwc.Amount = (hexutil.Big)(*hexutil.MustDecodeBig(row.Amount))
	rwc.IsDelegated = row.IsDlg

	return nil
}
