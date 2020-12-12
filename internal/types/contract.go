// Package types implements different core types of the API.
package types

import (
	"encoding/json"
	"fantom-api-graphql/internal/repository/rpc/contracts"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Contract represents an Opera smart contract at the blockchain.
type Contract struct {
	// Type represents a general type of the contract.
	Type string `json:"type" bson:"type"`

	// OrdinalIndex is the ordinal contract index in the database.
	OrdinalIndex uint64 `json:"index" bson:"orx"`

	// Address represents the address of the contract
	Address    Address `json:"address"`
	RawAddress string  `bson:"adr"`

	// TransactionHash represents the hash of the contract deployment transaction.
	TransactionHash    Hash   `json:"tx"`
	RawTransactionHash string `bson:"trx"`

	// TimeStamp represents the unix timestamp of the contract deployment.
	TimeStamp hexutil.Uint64 `json:"timestamp" bson:"ts"`

	// Name of the smart contract, if available.
	Name string `json:"name" bson:"name,omitempty"`

	// Smart contract version identifier, if available.
	Version string `json:"ver,omitempty" bson:"ver,omitempty"`

	// SupportContact represents a contact to the smart contract support, if available.
	SupportContact string `json:"contact,omitempty" bson:"url,omitempty"`

	// License represents an optional contact open source license
	// being used.
	License string `json:"license,omitempty" bson:"lic,omitempty"`

	// Smart contract compiler identifier, if available.
	Compiler string `json:"cv,omitempty" bson:"comp,omitempty"`

	// IsOptimized signals that the contract byte code was optimized
	// during compilation.
	IsOptimized bool `json:"optimized" bson:"is_opt"`

	// OptimizeRuns represents number of optimization runs used
	// during the contract compilation.
	OptimizeRuns int32 `json:"optimizeRuns" bson:"opt_run,omitempty"`

	// SourceCode is the smart contract source code, if available.
	SourceCode string `json:"sol,omitempty" bson:"src,omitempty"`

	// SourceCodeHash represents a hash code of the stored contract
	// source code. Is nil if the source code is not available.
	SourceCodeHash    *Hash   `json:"soh,omitempty"`
	RawSourceCodeHash *string `bson:"src_hash,omitempty"`

	// ABI definition of the smart contract, if available.
	Abi string `json:"abi,omitempty" bson:"abi,omitempty"`

	// Validated represents the unix timestamp
	//of the contract source validation against deployed byte code.
	Validated *hexutil.Uint64 `json:"ok,omitempty" bson:"is_ok,omitempty"`
}

// UnmarshalContract parses the JSON-encoded smart contract data.
func UnmarshalContract(data []byte) (*Contract, error) {
	var sc Contract
	err := json.Unmarshal(data, &sc)
	return &sc, err
}

// Marshal returns the JSON encoding of Contract.
func (sc *Contract) Marshal() ([]byte, error) {
	return json.Marshal(sc)
}

// NewGenericContract creates new generic contract record
func NewGenericContract(addr *common.Address, block *Block, trx *Transaction) *Contract {
	// make the contract
	return &Contract{
		Type:            AccountTypeContract,
		OrdinalIndex:    TransactionIndex(block, trx),
		Address:         Address(*addr),
		TransactionHash: trx.Hash,
		TimeStamp:       block.TimeStamp,
		Name:            "",
		Version:         "",
		SupportContact:  "",
		License:         "",
		Compiler:        "",
		IsOptimized:     false,
		OptimizeRuns:    0,
		SourceCode:      "",
		SourceCodeHash:  nil,
		Abi:             "",
		Validated:       nil,
	}
}

// NewErc20Contract creates new basic ERC20 contract
func NewErc20Contract(addr *common.Address, name string, block *Block, trx *Transaction) *Contract {
	// make the contract
	con := NewGenericContract(addr, block, trx)

	// set additional details
	con.Type = AccountTypeERC20Token
	con.Abi = contracts.ERCTwentyABI
	con.Name = name
	return con
}

// NewSfcContract creates new Special Purpose Contract reference
func NewSfcContract(addr *common.Address, ver uint64, block *Block, trx *Transaction) *Contract {
	// make the contract
	con := NewGenericContract(addr, block, trx)

	// set additional details
	con.Type = AccountTypeSFC
	con.Name = "SFC Contract"
	con.Version = fmt.Sprintf("%s.%s.%s",
		string([]byte{byte((ver >> 16) & 255)}),
		string([]byte{byte((ver >> 8) & 255)}),
		string([]byte{byte(ver & 255)}))
	con.SupportContact = "https://fantom.foundation"
	con.License = "MIT"
	con.Compiler = "Solidity"
	con.SourceCode = "https://github.com/Fantom-foundation/fantom-sfc"
	con.Abi = contracts.SfcContractABI
	con.Validated = &block.TimeStamp
	return con
}

// NewStiContract creates new Staker Information Contract reference
func NewStiContract(addr *common.Address, block *Block, trx *Transaction) *Contract {
	// make the contract
	con := NewGenericContract(addr, block, trx)

	// set additional details
	con.Name = "Staker Info Contract"
	con.Version = "1.4.0"
	con.SupportContact = "https://github.com/block42-blockchain-company/fantom-staker-info"
	con.License = "MIT"
	con.Compiler = "Solidity"
	con.SourceCode = "https://github.com/block42-blockchain-company/fantom-staker-info"
	con.Abi = contracts.StakerInfoContractABI
	return con
}
