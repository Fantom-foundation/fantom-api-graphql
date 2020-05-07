// Package types implements different core types of the API.
package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Contract represents an Opera smart contract at the blockchain.
type Contract struct {
	// OrdinalIndex is the ordinal contract index in the database.
	OrdinalIndex uint64

	// Address represents the address of the contract
	Address common.Address `json:"address"`

	// TransactionHash represents the hash of the contract deployment transaction.
	TransactionHash Hash `json:"tx"`

	// TimeStamp represents the unix timestamp of the contract deployment.
	TimeStamp hexutil.Uint64 `json:"timestamp"`

	// Name of the smart contract, if available.
	Name string `json:"name"`

	// Smart contract version identifier, if available.
	Version string `json:"ver"`

	// Smart contract compiler identifier, if available.
	Compiler string `json:"cv"`

	// SourceCode is the smart contract source code, if available.
	SourceCode string `json:"sol"`

	// ABI definition of the smart contract, if available.
	Abi string `json:"abi"`

	// Validated represents the unix timestamp
	//of the contract source validation against deployed byte code.
	Validated *hexutil.Uint64 `json:"ok"`
}
