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
	FMintTrxTypeDeposit = iota
	FMintTrxTypeWithdraw
	FMintTrxTypeMint
	FMintTrxTypeRepay
)

// FMintAmountDecimalsCorrection represents the correction applied to base fMint trx amount to get apr value.
var FMintAmountDecimalsCorrection = new(big.Int).SetUint64(1000000000)

// FMintTransaction represents a core transaction on fMint contract.
type FMintTransaction struct {
	User        common.Address
	Token       common.Address
	Type        int32
	Amount      hexutil.Big
	Fee         hexutil.Big
	Transaction common.Hash
	TimeStamp   hexutil.Uint64
}

// PK generates a unique primary key for the given fMint transaction.
func (ftx *FMintTransaction) PK() string {
	// make the base PK from the trx hash and log index
	bytes := make([]byte, 32)
	copy(bytes, ftx.User.Bytes()[:10])
	copy(bytes[10:], ftx.Token.Bytes()[:10])
	copy(bytes[20:], ftx.Transaction.Bytes()[:10])

	// xor in the timestamp and the root trx index
	ts := make([]byte, 8)
	binary.BigEndian.PutUint64(ts, ((uint64(ftx.TimeStamp)&0x7FFFFFFFFF)<<8)|(uint64(ftx.Type)&0xFF))
	for i := 0; i < 8; i++ {
		bytes[24+i] = bytes[24+i] ^ ts[i]
	}
	return hexutil.Encode(bytes)
}

// MarshalBSON creates a BSON representation of an fMint transaction.
func (ftx *FMintTransaction) MarshalBSON() ([]byte, error) {
	// calculate the value to 9 digits (and 18 billions remain available)
	val := new(big.Int).Div(ftx.Amount.ToInt(), FMintAmountDecimalsCorrection).Int64()
	fee := new(big.Int).Div(ftx.Fee.ToInt(), FMintAmountDecimalsCorrection).Int64()

	// prep the structure for saving
	pom := struct {
		ID        string    `bson:"_id"`
		Type      int32     `bson:"typ"`
		User      string    `bson:"usr"`
		Token     string    `bson:"tok"`
		Amount    string    `bson:"amo"`
		Fee       string    `bson:"fee"`
		Trx       string    `bson:"trx"`
		TimeStamp time.Time `bson:"stamp"`
		Value     int64     `bson:"val"`
		FeeValue  int64     `bson:"fee_val"`
	}{
		ID:        ftx.PK(),
		User:      ftx.User.String(),
		Token:     ftx.Token.String(),
		Trx:       ftx.Transaction.String(),
		Amount:    ftx.Amount.String(),
		Fee:       ftx.Fee.String(),
		TimeStamp: time.Unix(int64(ftx.TimeStamp), 0),
		Value:     val,
		FeeValue:  fee,
	}
	return bson.Marshal(pom)
}

// UnmarshalBSON updates the value from BSON source.
func (ftx *FMintTransaction) UnmarshalBSON(data []byte) (err error) {
	// capture unmarshal issue
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("can not decode and unmarshal")
		}
	}()

	// try to decode the BSON data
	var row struct {
		Type      int32     `bson:"typ"`
		User      string    `bson:"usr"`
		Token     string    `bson:"tok"`
		Amount    string    `bson:"amo"`
		Fee       string    `bson:"fee"`
		Trx       string    `bson:"trx"`
		TimeStamp time.Time `bson:"stamp"`
	}
	if err = bson.Unmarshal(data, &row); err != nil {
		return err
	}

	// transfer values
	ftx.Type = row.Type
	ftx.User = common.HexToAddress(row.User)
	ftx.Token = common.HexToAddress(row.Token)
	ftx.Transaction = common.HexToHash(row.Trx)
	ftx.Amount = (hexutil.Big)(*hexutil.MustDecodeBig(row.Amount))
	ftx.Fee = (hexutil.Big)(*hexutil.MustDecodeBig(row.Fee))
	ftx.TimeStamp = (hexutil.Uint64)(uint64(row.TimeStamp.Unix()))
	return nil
}
