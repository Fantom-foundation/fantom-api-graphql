// Package types implements different core types of the API.
package types

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/rand"
)

// Transaction represents a basic information provided by the API about transaction inside Opera blockchain.
type Transaction struct {
	// OrdinalIndex represents the ordinal index of the transaction inside the block chain.
	// It's build from the block number and the index of the transaction inside the block
	// when the transaction is stored in off-chain database.
	OrdinalIndex uint64 `json:"orx" bson:"orx"`

	// BlockHash represents hash of the block where this transaction was in. nil when its pending.
	BlockHash *Hash `json:"blockHash" bson:"-"`

	// BlockNumber represents number of the block where this transaction was in. nil when its pending.
	BlockNumber *hexutil.Uint64 `json:"blockNumber" bson:"-"`

	// From represents address of the sender.
	From common.Address `json:"from" bson:"-"`

	// Gas represents gas provided by the sender.
	Gas hexutil.Uint64 `json:"gas" bson:"-"`

	// Gas represents gas provided by the sender.
	GasUsed *hexutil.Uint64 `json:"gasUsed" bson:"-"`

	// CumulativeGasUsed represents the total amount of gas used when this transaction was executed in the block.
	CumulativeGasUsed *hexutil.Uint64 `json:"-" bson:"-"`

	// GasPrice represents gas price provided by the sender in Wei.
	GasPrice hexutil.Big `json:"gasPrice" bson:"-"`

	// Hash represents 32 bytes hash of the transaction.
	Hash Hash `json:"hash" bson:"-"`

	// Nonce represents the number of transactions made by the sender prior to this one.
	Nonce hexutil.Uint64 `json:"nonce" bson:"-"`

	// To represents the address of the receiver. nil when its a contract creation transaction.
	To *common.Address `json:"to" bson:"-"`

	// ContractAddress represents the address of contract created, if a contract creation transaction, otherwise nil.
	ContractAddress *common.Address `json:"contract" bson:"-"`

	// TrxIndex represents integer of the transaction's index position in the block. nil when its pending.
	TrxIndex *hexutil.Uint `json:"transactionIndex" bson:"-"`

	// Value represents value transferred in Wei.
	Value hexutil.Big `json:"value" bson:"-"`

	// Input represents the data send along with the transaction.
	InputData hexutil.Bytes `json:"input" bson:"-"`

	// Index represents integer of the transaction's index position in the block.
	Index *hexutil.Uint64 `json:"index" bson:"-"`

	// Status represents transaction status; value is either 1 (success) or 0 (failure)
	Status *hexutil.Uint64 `json:"status" bson:"-"`

	// TargetContractType represents the type of the contract the transaction
	// calls, if this is a contract call.
	TargetContractType *string `json:"tContract" bson:"tc"`

	// TargetFunctionCall represents the function the transaction addresses
	// byt the call, if this is a contract call.
	TargetFunctionCall *string `json:"tFunction" bson:"call"`

	// IsErc20Call marks an ERC20 call
	IsErc20Call bool `json:"isErc" bson:"iserc"`

	// IsProcessed mars the transaction as processed through the internal scanner
	IsProcessed bool `json:"isProc" bson:"isok"`
}

// mustTransactionIndex always calculate the index of the current transaction
// The ordinal index of a transaction should be unique across a consistent block chain.
// The calculation gives us about 700 years of index space with 50k blocks per second
// rate + 10 years to fix than. Max number of transactions in a block here is 14bits = 16383.
func TransactionIndex(block *Block, trx *Transaction) uint64 {
	// what is the transaction index
	var txIndex uint64
	if trx.TrxIndex != nil {
		txIndex = uint64(*trx.TrxIndex)
	} else {
		txIndex = uint64((rand.Uint32()<<10)|0xff) & 0x3fff
	}
	return (uint64(block.Number) << 14) | txIndex
}

// UnmarshalTransaction parses the JSON-encoded block data.
func UnmarshalTransaction(data []byte) (*Transaction, error) {
	var trx Transaction
	err := json.Unmarshal(data, &trx)
	return &trx, err
}

// Marshal returns the JSON encoding of transaction.
func (trx *Transaction) Marshal() ([]byte, error) {
	return json.Marshal(trx)
}
