// Package types implements different core types of the API.
package types

import (
	"encoding/binary"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson"
)

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
	CreatedEpoch      hexutil.Uint64
	Amount            *hexutil.Big

	// struct members for finalized withdraw
	WithdrawTrx  *common.Hash
	WithdrawTime *hexutil.Uint64
}

// Uid returns a unique identifier for the given withdraw request.
func (wr *WithdrawRequest) Uid() uint64 {
	return uint64(wr.CreatedTime)<<16 | (wr.StakerID.ToInt().Uint64()&0xFF)<<8 | (binary.BigEndian.Uint64(wr.RequestTrx[:8]) & 0xFF)
}

// MarshalBSON creates a BSON representation of the withdraw request record.
func (wr *WithdrawRequest) MarshalBSON() ([]byte, error) {
	pom := struct {
		Uid     uint64  `bson:"_id"`
		ReqID   string  `bson:"req_id"`
		ReqTrx  string  `bson:"req_trx"`
		Addr    string  `bson:"addr"`
		To      string  `bson:"to"`
		CrTime  uint64  `bson:"cr_time"`
		CrEpoch uint64  `bson:"cr_epoch"`
		Amount  string  `bson:"amount"`
		FinTrx  *string `bson:"fin_trx"`
		FinTime *uint64 `bson:"fin_time"`
	}{
		Uid:     wr.Uid(),
		ReqID:   wr.WithdrawRequestID.String(),
		ReqTrx:  wr.RequestTrx.String(),
		Addr:    wr.Address.String(),
		To:      wr.StakerID.String(),
		CrTime:  uint64(wr.CreatedTime),
		CrEpoch: uint64(wr.CreatedEpoch),
		Amount:  wr.Amount.String(),
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
		CrEpoch uint64  `bson:"cr_epoch"`
		Amount  string  `bson:"amount"`
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
	wr.CreatedEpoch = hexutil.Uint64(row.CrEpoch)
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

// WithdrawRequestsByAge represents a list of withdraw requests sortable
// by their age of creation. New requests are on top.
type WithdrawRequestsByAge []*WithdrawRequest

// Len returns size of the withdraw requests list.
func (wr WithdrawRequestsByAge) Len() int {
	return len(wr)
}

// Less compares two withdraw requests and returns true if the first is lower than the last.
// We use it to sort withdraw requests by time created having newer on top.
func (wr WithdrawRequestsByAge) Less(i, j int) bool {
	return uint64(wr[i].CreatedTime) > uint64(wr[j].CreatedTime)
}

// Swap changes position of two withdraw requests in the list.
func (wr WithdrawRequestsByAge) Swap(i, j int) {
	wr[i], wr[j] = wr[j], wr[i]
}
