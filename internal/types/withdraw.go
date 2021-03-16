// Package types implements different core types of the API.
package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// WithdrawRequest represents a withdraw request in Opera staking
// SFC contract. When partial withdraw is requested either on staking or delegation,
// this record is created in the SFC contract to track the withdrawal process.
type WithdrawRequest struct {
	// struct members for initiated withdraw
	Address            common.Address `bson:"addr"`
	StakerID           hexutil.Uint64 `bson:"staker"`
	WithdrawRequestID  hexutil.Big    `bson:"reqID"`
	IsDelegation       bool           `bson:"isDlg"`
	Amount             hexutil.Big    `bson:"amo"`
	RequestBlockNumber hexutil.Uint64 `bson:"reqBlk"`
	Receiver           common.Address `bson:"to"`

	// struct members for finalized withdraw
	WithdrawBlockNumber *hexutil.Uint64 `bson:"finBlk"`
	WithdrawAmount      *hexutil.Big    `bson:"finAmo"`
	WithdrawPenalty     *hexutil.Big    `bson:"finPen"`
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
	return uint64(wr[i].RequestBlockNumber) > uint64(wr[j].RequestBlockNumber)
}

// Swap changes position of two withdraw requests in the list.
func (wr WithdrawRequestsByAge) Swap(i, j int) {
	wr[i], wr[j] = wr[j], wr[i]
}
