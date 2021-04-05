// Package types implements different core types of the API.
package types

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	retypes "github.com/ethereum/go-ethereum/core/types"
	"go.mongodb.org/mongo-driver/bson"
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

	// Logs represents a list of log records created along with the transaction
	Logs []retypes.Log `json:"logs"`
}

// bsonLog represents the transaction log record data structure for BSON formatting.
type BsonLog struct {
	Address string   `bson:"addr"`
	Topics  []string `bson:"top"`
	Data    []byte   `bson:"data"`
	Index   uint     `bson:"ix"`
	Removed bool     `bson:"rm"`
}

// bsonTransaction represents the transaction data structure for BSON formatting.
type BsonTransaction struct {
	Hash      string    `bson:"_id"`
	Ordinal   uint64    `bson:"orx"`
	BlockID   *uint64   `bson:"blk"`
	BlockHash *string   `bson:"blk_h"`
	BlkIndex  *uint64   `bson:"bix"`
	From      string    `bson:"from"`
	To        *string   `bson:"to"`
	Gas       uint64    `bson:"gas"`
	UsedGas   *uint64   `bson:"gas_use"`
	CumGas    *uint64   `bson:"gas_all"`
	GasPrice  string    `bson:"gas_pri"`
	Nonce     uint64    `bson:"nonce"`
	Contract  *string   `bson:"contr"`
	Status    uint64    `bson:"stat"`
	Logs      []BsonLog `bson:"logs"`
}

// mustTransactionIndex always calculate the index of the current transaction
// The ordinal index of a transaction should be unique across a consistent block chain.
// The calculation gives us about 700 years of index space with 50k blocks per second
// rate + 10 years to fix than. Max number of transactions in a block here is 14bits = 16383.
func TransactionIndex(block *Block, trx *Transaction) uint64 {
	if trx.TrxIndex != nil {
		return (uint64(block.Number) << 14) | (uint64(*trx.TrxIndex) & 0x3fff)
	}
	return (uint64(block.Number) << 14) | (binary.BigEndian.Uint64(trx.Hash[:8]) & 0x3FFF)
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

// MarshalBSON creates a BSON representation of the Transaction record.
func (trx *Transaction) MarshalBSON() ([]byte, error) {
	// prep the structure for saving
	pom := BsonTransaction{
		Hash:     trx.Hash.String(),
		Ordinal:  trx.OrdinalIndex,
		From:     trx.From.String(),
		Gas:      uint64(trx.Gas),
		GasPrice: trx.GasPrice.String(),
		Nonce:    uint64(trx.Nonce),
	}

	// transaction has been mined, we have all the extra info, too
	if trx.BlockHash != nil {
		// block hash
		bh := trx.BlockHash.String()
		pom.BlockHash = &bh

		// block number
		bn := uint64(*trx.BlockNumber)
		pom.BlockID = &bn

		// block index
		bx := uint64(*trx.Index)
		pom.BlkIndex = &bx

		// used gas
		gu := uint64(*trx.GasUsed)
		pom.UsedGas = &gu

		// cumulative gas
		gc := uint64(*trx.CumulativeGasUsed)
		pom.CumGas = &gc

		// status
		pom.Status = uint64(*trx.Status)
	}

	// recipient
	if trx.To != nil {
		to := trx.To.String()
		pom.To = &to
	}

	// contract address
	if trx.ContractAddress != nil {
		cn := trx.ContractAddress.String()
		pom.Contract = &cn
	}

	// logs
	pom.Logs = make([]BsonLog, len(trx.Logs))
	for i, lg := range trx.Logs {
		// make the base
		pom.Logs[i] = BsonLog{
			Address: lg.Address.String(),
			Topics:  make([]string, len(lg.Topics)),
			Data:    lg.Data,
			Index:   lg.Index,
			Removed: lg.Removed,
		}

		// add topics
		for ti, th := range lg.Topics {
			pom.Logs[i].Topics[ti] = th.String()
		}
	}

	return bson.Marshal(pom)
}

// UnmarshalBSON updates the value from BSON source.
func (trx *Transaction) UnmarshalBSON(data []byte) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("can not decode stored transaction")
		}
	}()

	// try to decode the BSON data
	var row BsonTransaction
	if err = bson.Unmarshal(data, &row); err != nil {
		return err
	}

	// transfer values
	trx.Hash = HexToHash(row.Hash)
	trx.OrdinalIndex = row.Ordinal
	trx.From = common.HexToAddress(row.From)
	trx.Gas = hexutil.Uint64(row.Gas)
	trx.GasPrice = (hexutil.Big)(*hexutil.MustDecodeBig(row.GasPrice))
	trx.Nonce = hexutil.Uint64(row.Nonce)
	trx.Status = (*hexutil.Uint64)(&row.Status)

	// pointers
	if row.BlockHash != nil {
		// block hash
		bh := HexToHash(*row.BlockHash)
		trx.BlockHash = &bh

		// block number
		bn := hexutil.Uint64(*row.BlockID)
		trx.BlockNumber = &bn

		// index of the transaction in the block
		bx := hexutil.Uint64(*row.BlkIndex)
		trx.Index = &bx
		bi := hexutil.Uint(uint(*row.BlkIndex))
		trx.TrxIndex = &bi

		// used gas
		gu := hexutil.Uint64(*row.UsedGas)
		trx.GasUsed = &gu

		// cumulative gas
		gc := hexutil.Uint64(*row.CumGas)
		trx.CumulativeGasUsed = &gc
	}

	// recipient
	if row.To != nil {
		to := common.HexToAddress(*row.To)
		trx.To = &to
	}

	// contract
	if row.Contract != nil {
		cn := common.HexToAddress(*row.Contract)
		trx.ContractAddress = &cn
	}

	// logs
	trx.Logs = make([]retypes.Log, len(row.Logs))
	for i, lg := range row.Logs {
		// base
		trx.Logs[i] = retypes.Log{
			Address:     common.HexToAddress(lg.Address),
			Topics:      make([]common.Hash, len(lg.Topics)),
			Data:        lg.Data,
			BlockNumber: *row.BlockID,
			TxHash:      common.HexToHash(row.Hash),
			TxIndex:     uint(*row.BlkIndex),
			BlockHash:   common.HexToHash(*row.BlockHash),
			Index:       lg.Index,
			Removed:     lg.Removed,
		}

		// copy topics
		for ti, th := range lg.Topics {
			trx.Logs[i].Topics[ti] = common.HexToHash(th)
		}
	}
	return nil
}
