// Package types implements different core types of the API.
package types

import (
	"encoding/binary"
	"fantom-api-graphql/internal/repository/db/registry"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
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
	ID           string         `bson:"_id"`
	UserAddress  common.Address `bson:"usr"`
	TokenAddress common.Address `bson:"tok"`
	Type         int32          `bson:"typ"`
	Amount       hexutil.Big    `bson:"amo"`
	Fee          hexutil.Big    `bson:"fee"`
	FeeValue     int64          `bson:"fee_val"`
	TrxHash      common.Hash    `bson:"trx"`
	TrxIndex     int64          `bson:"tix"`
	TimeStamp    hexutil.Uint64 `bson:"stamp"`
	Value        int64          `bson:"val"`
	Ordinal      int64          `bson:"orx"`
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
	ftx.Value = new(big.Int).Div(ftx.Amount.ToInt(), FMintAmountDecimalsCorrection).Int64()
	ftx.FeeValue = new(big.Int).Div(ftx.Fee.ToInt(), FMintAmountDecimalsCorrection).Int64()
	ftx.ID = ftx.Pk()
	ftx.Ordinal = ftx.OrdinalIndex()

	return registry.Marshal(ftx)
}
