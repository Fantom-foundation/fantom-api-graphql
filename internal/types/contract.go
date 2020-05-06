// Package types implements different core types of the API.
package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Contract represents an Opera smart contract at the blockchain.
type Contract struct {
	Address common.Address `json:"address"`

	// TransactionHash represents the hash of the contract deployment transaction.
	TransactionHash Hash `json:"tx"`

	// TimeStamp represents the unix timestamp of the contract deployment.
	TimeStamp hexutil.Uint64 `json:"timestamp"`

	// SourceCode is the smart contract source code, if available.
	SourceCode string `json:"sol"`

	// Validated represents the unix timestamp
	//of the contract source validation against deployed byte code.
	Validated *hexutil.Uint64 `json:"ok"`
}
