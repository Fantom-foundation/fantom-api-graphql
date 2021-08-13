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
	FiFMintTransactionId        = "_id"
	FiFMintTransactionToken     = "tok"
	FiFMintTransactionUser      = "usr"
	FiFMintTransactionTimestamp = "stamp"
	FiFMintTransactionOrdinal   = "orx"
)

// define types of fMint operations used on the protocol
const (
	FMintTrxTypeDeposit = iota
	FMintTrxTypeWithdraw
	FMintTrxTypeMint
	FMintTrxTypeRepay
)

// FMintAmountDecimalsCorrection represents the correction applied to base fMint trx amount to get apr value.
var FMintAmountDecimalsCorrection = new(big.Int).SetUint64(1000000000000)

// FMintUserTokens represents a user with its fMint tokens aggregated
// for the specific purpose.
type FMintUserTokens struct {
	Purpose int32
	User    common.Address
	Tokens  []common.Address
}

// FMintTransaction represents a core transaction on fMint contract.
type FMintTransaction struct {
	UserAddress  common.Address
	TokenAddress common.Address
	Type         int32
	Amount       hexutil.Big
	Fee          hexutil.Big
	TrxHash      common.Hash
	TrxIndex     int64
	TimeStamp    hexutil.Uint64
}

// Pk generates a unique primary key for the given fMint transaction.
func (ftx *FMintTransaction) Pk() string {
	// make the base PK from the trx hash and log index
	bytes := make([]byte, 32)
	copy(bytes, ftx.UserAddress.Bytes()[:10])
	copy(bytes[10:], ftx.TokenAddress.Bytes()[:10])
	copy(bytes[20:], ftx.TrxHash.Bytes()[:10])

	// xor in the timestamp and the root trx index
	ts := make([]byte, 8)
	binary.BigEndian.PutUint64(ts, ((uint64(ftx.TimeStamp)&0x7FFFFFFFFF)<<8)|(uint64(ftx.Type)&0xFF))
	for i := 0; i < 8; i++ {
		bytes[24+i] = bytes[24+i] ^ ts[i]
	}
	return hexutil.Encode(bytes)
}

// OrdinalIndex returns an ordinal index for the given fMint transaction.
func (ftx *FMintTransaction) OrdinalIndex() int64 {
	return ((int64(ftx.TimeStamp)<<14)&0x7FFFFFFFFFFFFFFF | (int64(ftx.TrxIndex) & 0x3fff)) ^ (int64(ftx.TrxHash[0]) << 8)
}

// MarshalBSON creates a BSON representation of an fMint transaction.
func (ftx *FMintTransaction) MarshalBSON() ([]byte, error) {
	// calculate the value to 9 digits (and 18 billions remain available)
	val := new(big.Int).Div(ftx.Amount.ToInt(), FMintAmountDecimalsCorrection).Int64()
	fee := new(big.Int).Div(ftx.Fee.ToInt(), FMintAmountDecimalsCorrection).Int64()

	// prep the structure for saving
	pom := struct {
		ID        string    `bson:"_id"`
		Ordinal   int64     `bson:"orx"`
		Type      int32     `bson:"typ"`
		User      string    `bson:"usr"`
		Token     string    `bson:"tok"`
		Amount    string    `bson:"amo"`
		Fee       string    `bson:"fee"`
		Trx       string    `bson:"trx"`
		TrxIndex  int64     `bson:"tix"`
		TimeStamp time.Time `bson:"stamp"`
		Value     int64     `bson:"val"`
		FeeValue  int64     `bson:"fee_val"`
	}{
		ID:        ftx.Pk(),
		Ordinal:   ftx.OrdinalIndex(),
		Type:      ftx.Type,
		User:      ftx.UserAddress.String(),
		Token:     ftx.TokenAddress.String(),
		Trx:       ftx.TrxHash.String(),
		TrxIndex:  ftx.TrxIndex,
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
		TrxHash   string    `bson:"trx"`
		TrxIndex  int64     `bson:"tix"`
		TimeStamp time.Time `bson:"stamp"`
	}
	if err = bson.Unmarshal(data, &row); err != nil {
		return err
	}

	// transfer values
	ftx.Type = row.Type
	ftx.UserAddress = common.HexToAddress(row.User)
	ftx.TokenAddress = common.HexToAddress(row.Token)
	ftx.TrxHash = common.HexToHash(row.TrxHash)
	ftx.TrxIndex = row.TrxIndex
	ftx.Amount = (hexutil.Big)(*hexutil.MustDecodeBig(row.Amount))
	ftx.Fee = (hexutil.Big)(*hexutil.MustDecodeBig(row.Fee))
	ftx.TimeStamp = (hexutil.Uint64)(uint64(row.TimeStamp.Unix()))
	return nil
}
