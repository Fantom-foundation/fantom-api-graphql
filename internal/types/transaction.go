package types

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Transaction represents a basic information provided by the API about transaction inside Opera blockchain.
type Transaction struct {
	// BlockHash represents hash of the block where this transaction was in. nil when its pending.
	BlockHash *Hash `json:"blockHash"`

	// BlockNumber represents number of the block where this transaction was in. nil when its pending.
	BlockNumber *hexutil.Uint64 `json:"blockNumber"`

	// From represents address of the sender.
	From common.Address `json:"from"`

	// Gas represents gas provided by the sender.
	Gas hexutil.Uint64 `json:"gas"`

	// Gas represents gas provided by the sender.
	GasUsed *hexutil.Uint64 `json:"gasUsed"`

	// CumulativeGasUsed represents the total amount of gas used when this transaction was executed in the block.
	CumulativeGasUsed *hexutil.Uint64 `json:"-"`

	// GasPrice represents gas price provided by the sender in Wei.
	GasPrice hexutil.Big `json:"gasPrice"`

	// Hash represents 32 bytes hash of the transaction.
	Hash Hash `json:"hash"`

	// Nonce represents the number of transactions made by the sender prior to this one.
	Nonce hexutil.Uint64 `json:"nonce"`

	// To represents the address of the receiver. nil when its a contract creation transaction.
	To *common.Address `json:"to"`

	// ContractAddress represents the address of contract created, if a contract creation transaction, otherwise nil.
	ContractAddress *common.Address `json:"-"`

	// TrxIndex represents integer of the transaction's index position in the block. nil when its pending.
	TrxIndex *hexutil.Uint `json:"transactionIndex"`

	// Value represents value transferred in Wei.
	Value hexutil.Big `json:"value"`

	// Input represents the data send along with the transaction.
	InputData hexutil.Bytes `json:"input"`

	// Index represents integer of the transaction's index position in the block.
	Index *hexutil.Uint64 `json:"index"`

	// Status represents transaction status; value is either 1 (success) or 0 (failure)
	Status *hexutil.Uint64 `json:"status"`
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
