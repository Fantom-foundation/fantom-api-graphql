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
	FiDelegationPk                 = "_id"
	FiDelegationOrdinal            = "orx"
	FiDelegationAddress            = "adr"
	FiDelegationToValidator        = "to"
	FiDelegationTransaction        = "trx"
	FiDelegationToValidatorAddress = "toad"
	FiDelegationAmountActive       = "act"
	FiDelegationValue              = "val"
	FiDelegationStamp              = "stamp"
)

// Delegation represents a delegator in Opera blockchain.
type Delegation struct {
	Transaction     common.Hash    `json:"trx"`
	Address         common.Address `json:"address"`
	ToStakerId      *hexutil.Big   `json:"toStakerID"`
	ToStakerAddress common.Address `json:"toStakerAddr"`
	CreatedTime     hexutil.Uint64 `json:"createdTime"`
	Index           uint64         `json:"ordinalIndex"`

	// AmountStaked represents the current staked amount
	AmountStaked *hexutil.Big `json:"amountStaked"`

	// AmountDelegated is the original amount delegated
	AmountDelegated *hexutil.Big `json:"amountDelegated"`
}

// BsonDelegation represents the BSON i/o struct for a delegation.
type BsonDelegation struct {
	ID     string    `bson:"_id"`
	Orx    uint64    `bson:"orx"`
	Trx    string    `bson:"trx"`
	Addr   string    `bson:"adr"`
	To     string    `bson:"to"`
	ToAddr string    `bson:"toad"`
	CrTime uint64    `bson:"crt"`
	Staked string    `bson:"amo"`
	Active string    `bson:"act"`
	Value  uint64    `bson:"val"`
	Stamp  time.Time `bson:"stamp"`
}

// DelegationDecimalsCorrection is used to manipulate precision of a delegation active value
// so it can be stored in database as UINT64 without loosing too much data
var DelegationDecimalsCorrection = new(big.Int).SetUint64(1000000000)

// OrdinalIndex returns an ordinal index for the given delegation.
// We construct the UID from the time the delegation was created (40 bits = 1099511627775s = 34000 years),
// a part of the creation transaction hash and part of the target validator index (12 bits = 4096).
func (dl *Delegation) OrdinalIndex() uint64 {
	return (uint64(dl.CreatedTime)&0x7FFFFFFFFF)<<24 | (dl.ToStakerId.ToInt().Uint64()&0xFFF)<<12 | (binary.BigEndian.Uint64(dl.Transaction[:8]) & 0xFFF)
}

// MarshalBSON creates a BSON representation of the delegation record.
func (dl *Delegation) MarshalBSON() ([]byte, error) {
	// calculate the value to 9 digits (and 18 billions remain available)
	val := new(big.Int).Div(dl.AmountDelegated.ToInt(), DelegationDecimalsCorrection).Uint64()

	// do BSON encoding
	return bson.Marshal(BsonDelegation{
		ID:     dl.Transaction.String(),
		Orx:    dl.OrdinalIndex(),
		Trx:    dl.Transaction.String(),
		Addr:   dl.Address.String(),
		To:     dl.ToStakerId.String(),
		ToAddr: dl.ToStakerAddress.String(),
		CrTime: uint64(dl.CreatedTime),
		Staked: dl.AmountStaked.String(),
		Active: dl.AmountDelegated.String(),
		Value:  val,
		Stamp:  time.Unix(int64(dl.CreatedTime), 0),
	})
}

// UnmarshalBSON updates the value from BSON source.
func (dl *Delegation) UnmarshalBSON(data []byte) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("can not decode BIG number in delegation unmarshal")
		}
	}()

	// try to decode the BSON data
	var row BsonDelegation
	if err = bson.Unmarshal(data, &row); err != nil {
		return err
	}

	// transfer values
	dl.Transaction = common.HexToHash(row.Trx)
	dl.Address = common.HexToAddress(row.Addr)
	dl.ToStakerId = (*hexutil.Big)(hexutil.MustDecodeBig(row.To))
	dl.ToStakerAddress = common.HexToAddress(row.ToAddr)
	dl.CreatedTime = hexutil.Uint64(row.CrTime)
	dl.AmountStaked = (*hexutil.Big)(hexutil.MustDecodeBig(row.Staked))
	dl.AmountDelegated = (*hexutil.Big)(hexutil.MustDecodeBig(row.Active))
	dl.Index = row.Orx
	return nil
}
