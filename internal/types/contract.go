// Package types implements different core types of the API.
package types

import (
	"encoding/binary"
	"encoding/json"
	"fantom-api-graphql/internal/repository/rpc/contracts"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson"
)

// Contract represents an Opera smart contract at the blockchain.
type Contract struct {
	// Type represents a general type of the contract.
	Type string `json:"type"`

	// Address represents the address of the contract
	Address common.Address `json:"address"`

	// TransactionHash represents the hash of the contract deployment transaction.
	TransactionHash common.Hash `json:"tx"`

	// TimeStamp represents the unix timestamp of the contract deployment.
	TimeStamp hexutil.Uint64 `json:"timestamp"`

	// Name of the smart contract, if available.
	Name string `json:"name"`

	// Smart contract version identifier, if available.
	Version string `json:"ver,omitempty"`

	// SupportContact represents a contact to the smart contract support, if available.
	SupportContact string `json:"contact,omitempty"`

	// License represents an optional contact open source license
	// being used.
	License string `json:"license,omitempty"`

	// Smart contract compiler identifier, if available.
	Compiler string `json:"cv,omitempty"`

	// IsOptimized signals that the contract byte code was optimized
	// during compilation.
	IsOptimized bool `json:"optimized"`

	// OptimizeRuns represents number of optimization runs used
	// during the contract compilation.
	OptimizeRuns int32 `json:"optimizeRuns"`

	// SourceCode is the smart contract source code, if available.
	SourceCode string `json:"sol,omitempty"`

	// SourceCodeHash represents a hash code of the stored contract
	// source code. Is nil if the source code is not available.
	SourceCodeHash *common.Hash `json:"soh,omitempty"`

	// ABI definition of the smart contract, if available.
	Abi string `json:"abi,omitempty" bson:"abi,omitempty"`

	// Validated represents the unix timestamp
	//of the contract source validation against deployed byte code.
	Validated *hexutil.Uint64 `json:"ok,omitempty" bson:"is_ok,omitempty"`
}

// BsonContract represents the contract data structure for BSON formatting.
type BsonContract struct {
	Address   string  `bson:"_id"`
	Type      string  `bson:"type"`
	Name      string  `bson:"name"`
	Ordinal   uint64  `bson:"orx"`
	Trx       string  `bson:"trx"`
	Created   uint64  `bson:"ts"`
	Version   string  `bson:"ver"`
	Support   string  `bson:"sup"`
	License   string  `bson:"lic"`
	Compiler  string  `bson:"sol"`
	IsOpt     bool    `bson:"is_opt"`
	OptRuns   int32   `bson:"opt"`
	Src       string  `bson:"src"`
	Abi       string  `bson:"abi"`
	SrcHash   *string `bson:"src_h"`
	Validated *uint64 `bson:"val"`
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

// Uid generates unique identifier of the contract record.
func (sc *Contract) Uid() uint64 {
	return (uint64(sc.TimeStamp)&0xFFFFFFFFFF)<<24 | (binary.BigEndian.Uint64(sc.TransactionHash[:8]) & 0xFFFFFF)
}

// NewGenericContract creates new generic contract record
func NewGenericContract(addr *common.Address, block *Block, trx *Transaction) *Contract {
	// make the contract
	return &Contract{
		Type:            AccountTypeContract,
		Address:         *addr,
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

// NewErcTokenContract creates new basic ERC20/ERC721/ERC1155 contract
func NewErcTokenContract(addr *common.Address, name string, block *Block, trx *Transaction, tType string, abi string) *Contract {
	// make the contract
	con := NewGenericContract(addr, block, trx)

	// set additional details
	con.Type = tType
	con.Abi = abi
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
	con.SourceCode = "https://github.com/Fantom-foundation/opera-sfc"
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

// MarshalBSON creates a BSON representation of the Contract record.
func (sc *Contract) MarshalBSON() ([]byte, error) {
	// prep the row
	row := BsonContract{
		Address:  sc.Address.String(),
		Type:     sc.Type,
		Name:     sc.Name,
		Ordinal:  sc.Uid(),
		Trx:      sc.TransactionHash.String(),
		Created:  uint64(sc.TimeStamp),
		Version:  sc.Version,
		Support:  sc.SupportContact,
		License:  sc.License,
		Compiler: sc.Compiler,
		IsOpt:    sc.IsOptimized,
		OptRuns:  sc.OptimizeRuns,
		Src:      sc.SourceCode,
		Abi:      sc.Abi,
	}
	// is validated?
	if sc.Validated != nil {
		row.Validated = (*uint64)(sc.Validated)
	}
	// do we have source code hash?
	if sc.SourceCodeHash != nil {
		val := sc.SourceCodeHash.String()
		row.SrcHash = &val
	}
	return bson.Marshal(row)
}

// UnmarshalBSON updates the value from BSON source.
func (sc *Contract) UnmarshalBSON(data []byte) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("can not decode stored contract")
		}
	}()

	// try to decode the BSON data
	var row BsonContract
	if err = bson.Unmarshal(data, &row); err != nil {
		return err
	}

	// transfer data
	sc.Address = common.HexToAddress(row.Address)
	sc.Type = row.Type
	sc.Name = row.Name
	sc.TransactionHash = common.HexToHash(row.Trx)
	sc.TimeStamp = hexutil.Uint64(row.Created)
	sc.Version = row.Version
	sc.SupportContact = row.Support
	sc.License = row.License
	sc.Compiler = row.Compiler
	sc.IsOptimized = row.IsOpt
	sc.OptimizeRuns = row.OptRuns
	sc.SourceCode = row.Src
	sc.Abi = row.Abi
	if row.Validated != nil {
		sc.Validated = (*hexutil.Uint64)(row.Validated)
	}
	if row.SrcHash != nil {
		val := common.HexToHash(*row.SrcHash)
		sc.SourceCodeHash = &val
	}
	return nil
}
