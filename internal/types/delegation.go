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
	// FiDelegationPk defines primary key column of the delegation table.
	FiDelegationPk = "_id"

	// FiDelegationOrdinal defines ordinal index column of the delegation table.
	FiDelegationOrdinal = "orx"

	// FiDelegationAddress defines delegation address column of the delegation table.
	FiDelegationAddress = "adr"

	// FiDelegationToValidator defines id of the validator column of the delegation table.
	FiDelegationToValidator = "to"

	// FiDelegationTransaction defines transaction has column of the delegation table.
	FiDelegationTransaction = "trx"

	// FiDelegationToValidatorAddress defines validator address column of the delegation table.
	FiDelegationToValidatorAddress = "toad"

	// FiDelegationAmountActive defines amount delegated column of the delegation table.
	FiDelegationAmountActive = "act"

	// FiDelegationValue defines value of the delegation column of the delegation table.
	FiDelegationValue = "val"

	// FiDelegationStamp defines time stamp column of the delegation table.
	FiDelegationStamp = "stamp"
)

// Delegation represents a delegator in Opera blockchain.
type Delegation struct {
	ID              string         `json:"id"`
	Transaction     common.Hash    `json:"trx"`
	Address         common.Address `json:"address"`
	ToStakerId      *hexutil.Big   `json:"toStakerID"`
	ToStakerAddress common.Address `json:"toStakerAddr"`
	CreatedTime     hexutil.Uint64 `json:"createdTime"`
	Index           uint64         `json:"ordinalIndex"`

	// AmountStaked represents the original staked amount
	AmountStaked *hexutil.Big `json:"amountStaked"`

	// AmountDelegated is the current amount delegated
	AmountDelegated *hexutil.Big `json:"amountDelegated"`
}

// DelegationDecimalsCorrection is used to adjust decimal precision of a delegation active value.
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
	return bson.Marshal(struct {
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
	}{
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
	var row struct {
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
	if err = bson.Unmarshal(data, &row); err != nil {
		return err
	}

	// transfer values
	dl.ID = row.ID
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
