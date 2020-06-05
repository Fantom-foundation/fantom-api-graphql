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
	Address            common.Address
	Receiver           common.Address
	StakerID           hexutil.Uint64
	WithdrawRequestID  hexutil.Big
	IsDelegation       bool
	Amount             hexutil.Big
	RequestBlockNumber hexutil.Uint64

	// struct members for finalized withdraw
	WithdrawBlockNumber *hexutil.Uint64
	WithdrawPenalty     *hexutil.Big
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
