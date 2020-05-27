// Package types implements different core types of the API.
package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// WithdrawRequest represents a partial withdraw request in Opera staking
// SFC contract. When partial withdraw is requested either on staking or delegation,
// this record is created in the SFC contract to track the withdrawal process.
type WithdrawRequest struct {
	Address           common.Address
	Receiver          common.Address
	StakerID          hexutil.Uint64
	WithdrawRequestID hexutil.Big
	IsDelegation      bool
	Amount            hexutil.Big
}
