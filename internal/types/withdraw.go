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

// WithdrawDecimalsCorrection is used to manipulate precision of a withdrawal value
// so it can be stored in database as UINT64 without loosing too much data
var WithdrawDecimalsCorrection = new(big.Int).SetUint64(1000000000)

// WithdrawRequest represents a withdraw request in Opera staking
// SFC contract. When partial withdraw is requested either on staking or delegation,
// this record is created in the SFC contract to track the withdrawal process.
type WithdrawRequest struct {
	// struct members for initiated withdraw
	RequestTrx        common.Hash
	WithdrawRequestID *hexutil.Big
	Address           common.Address
	StakerID          *hexutil.Big
	CreatedTime       hexutil.Uint64
	Amount            *hexutil.Big

	// struct members for finalized withdraw
	WithdrawTrx  *common.Hash
	WithdrawTime *hexutil.Uint64
}

// Uid returns a unique identifier for the given withdraw request.
func (wr *WithdrawRequest) Uid() uint64 {
	return (uint64(wr.CreatedTime)&0xFFFFFFFFFF)<<24 | (wr.StakerID.ToInt().Uint64()&0xFFF)<<12 | (binary.BigEndian.Uint64(wr.RequestTrx[:8]) & 0xFFF)
}

// MarshalBSON returns a BSON document for the withdraw request.
func (wr *WithdrawRequest) MarshalBSON() ([]byte, error) {
	// calculate the value to 9 digits (and 18 billions remain available)
	val := new(big.Int).Div(wr.Amount.ToInt(), WithdrawDecimalsCorrection)

	// prep the structure for saving
	pom := struct {
		Uid     uint64  `bson:"_id"`
		ReqID   string  `bson:"req_id"`
		ReqTrx  string  `bson:"req_trx"`
		Addr    string  `bson:"addr"`
		To      string  `bson:"to"`
		CrTime  uint64  `bson:"cr_time"`
		Amount  string  `bson:"amount"`
		Value   uint64  `bson:"value"`
		FinTrx  *string `bson:"fin_trx"`
		FinTime *uint64 `bson:"fin_time"`
	}{
		Uid:    wr.Uid(),
		ReqID:  wr.WithdrawRequestID.String(),
		ReqTrx: wr.RequestTrx.String(),
		Addr:   wr.Address.String(),
		To:     wr.StakerID.String(),
		CrTime: uint64(wr.CreatedTime),
		Amount: wr.Amount.String(),
		Value:  val.Uint64(),
	}
	if wr.WithdrawTrx != nil {
		val := wr.WithdrawTrx.String()
		pom.FinTrx = &val
	}
	if wr.WithdrawTime != nil {
		pom.FinTime = (*uint64)(wr.WithdrawTime)
	}
	return bson.Marshal(pom)
}

// UnmarshalBSON updates the value from BSON source.
func (wr *WithdrawRequest) UnmarshalBSON(data []byte) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("can not decode BIG number in withdrawal unmarshal")
		}
	}()

	// try to decode the BSON data
	var row struct {
		Uid     uint64  `bson:"_id"`
		ReqID   string  `bson:"req_id"`
		ReqTrx  string  `bson:"req_trx"`
		Addr    string  `bson:"addr"`
		To      string  `bson:"to"`
		CrTime  uint64  `bson:"cr_time"`
		Amount  string  `bson:"amount"`
		Value   uint64  `bson:"value"`
		FinTrx  *string `bson:"fin_trx"`
		FinTime *uint64 `bson:"fin_time"`
	}
	if err = bson.Unmarshal(data, &row); err != nil {
		return err
	}

	// transfer values
	wr.WithdrawRequestID = (*hexutil.Big)(hexutil.MustDecodeBig(row.ReqID))
	wr.RequestTrx = common.HexToHash(row.ReqTrx)
	wr.Address = common.HexToAddress(row.Addr)
	wr.StakerID = (*hexutil.Big)(hexutil.MustDecodeBig(row.To))
	wr.CreatedTime = hexutil.Uint64(row.CrTime)
	wr.Amount = (*hexutil.Big)(hexutil.MustDecodeBig(row.Amount))
	wr.WithdrawTrx = nil
	wr.WithdrawTime = nil
	if row.FinTrx != nil {
		val := common.HexToHash(*row.FinTrx)
		wr.WithdrawTrx = &val
	}
	if row.FinTime != nil {
		wr.WithdrawTime = (*hexutil.Uint64)(row.FinTime)
	}
	return nil
}
