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
	FiWithdrawalType        = "type"
	FiWithdrawalOrdinal     = "orx"
	FiWithdrawalRequestID   = "req_id"
	FiWithdrawalAddress     = "adr"
	FiWithdrawalToValidator = "to"
	FiWithdrawalCreated     = "crt"
	FiWithdrawalStamp       = "stamp"
	FiWithdrawalValue       = "val"
	FiWithdrawalSlash       = "slash"
	FiWithdrawalRequestTrx  = "req_trx"
	FiWithdrawalFinTrx      = "fin_trx"
	FiWithdrawalFinTime     = "fin_time"

	WithdrawTypeUndelegated     = "SFC3:Undelegated"
	WithdrawTypeWithdrawRequest = "SFC1:WithdrawRequest"
	WithdrawTypeDeactivatedDlg  = "SFC1:DeactivatedDelegation"
	WithdrawTypeDeactivatedVal  = "SFC1:DeactivatedStake"
)

// WithdrawRequest represents a withdraw request in Opera staking
// SFC contract. When partial withdraw is requested either on staking or delegation,
// this record is created in the SFC contract to track the withdrawal process.
type WithdrawRequest struct {
	// struct members for initiated withdraw
	ID                string          `bson:"_id"`
	WithdrawRequestID *hexutil.Big    `bson:"req_id"`
	Address           common.Address  `bson:"adr"`
	StakerID          *hexutil.Big    `bson:"to"`
	CreatedTime       hexutil.Uint64  `bson:"crt"`
	TimeStamp         time.Time       `bson:"stamp"`
	Amount            *hexutil.Big    `bson:"amo"`
	Type              string          `bson:"type"`
	Ordinal           uint64          `bson:"orx"`
	Value             uint64          `bson:"val"`
	RequestTrx        common.Hash     `bson:"req_trx"`
	WithdrawTrx       *common.Hash    `bson:"fin_trx"`
	WithdrawTime      *hexutil.Uint64 `bson:"fin_time"`
	Penalty           *hexutil.Big    `bson:"slash"`
}

// WithdrawDecimalsCorrection is used to manipulate precision of a withdrawal value
// so it can be stored in database as UINT64 without loosing too much data
var WithdrawDecimalsCorrection = new(big.Int).SetUint64(1000000000)

// OrdinalIndex returns an ordinal index of the withdraw request.
func (wr *WithdrawRequest) OrdinalIndex() uint64 {
	return (uint64(wr.CreatedTime)&0xFFFFFFFFFF)<<24 | (wr.StakerID.ToInt().Uint64()&0xFFF)<<12 | (binary.BigEndian.Uint64(wr.RequestTrx[:8]) & 0xFFF)
}

// MarshalBSON returns a BSON document for the withdrawal request.
func (wr *WithdrawRequest) MarshalBSON() ([]byte, error) {
	// calculate the value to 9 digits (and 18 billions remain available)
	wr.Value = new(big.Int).Div(wr.Amount.ToInt(), WithdrawDecimalsCorrection).Uint64()
	wr.Ordinal = wr.OrdinalIndex()
	wr.ID = wr.RequestTrx.String()
	wr.TimeStamp = time.Unix(int64(wr.CreatedTime), 0)

	return registry.Marshal(wr)
}
