// Package types implements different core types of the API.
package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// DeactivatedDelegation represents a delegation deactivation request in Opera staking
// SFC contract. When full withdraw is requested on delegation,
// this record is created in the SFC contract to track the withdrawal process.
type DeactivatedDelegation struct {
	Address            common.Address
	StakerID           hexutil.Uint64
	RequestBlockNumber hexutil.Uint64

	// struct members for finalized withdraw
	WithdrawBlockNumber *hexutil.Uint64
	WithdrawPenalty     *hexutil.Big
}

// DeactivatedDelegationByAge represents a list of deactivated delegation
// requests sortable by their age of creation. New requests are on top.
type DeactivatedDelegationByAge []*DeactivatedDelegation

// Len returns size of the deactivated delegation requests list.
func (dd DeactivatedDelegationByAge) Len() int {
	return len(dd)
}

// Less compares two deactivated delegation requests and returns true
// if the first is lower than the last.
// We use it to sort deactivated delegation requests
// by time created having newer on top.
func (dd DeactivatedDelegationByAge) Less(i, j int) bool {
	return uint64(dd[i].RequestBlockNumber) > uint64(dd[j].RequestBlockNumber)
}

// Swap changes position of two withdraw requests in the list.
func (dd DeactivatedDelegationByAge) Swap(i, j int) {
	dd[i], dd[j] = dd[j], dd[i]
}
