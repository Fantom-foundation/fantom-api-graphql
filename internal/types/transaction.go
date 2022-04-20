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
	"math/big"
	"time"
)

// trxLargeInputWall represents the largest transaction input block we store in the off-chain database.
// Larger inputs (like contract deployments) need to be loaded from the blockchain directly if needed.
const trxLargeInputWall = 32 * 8

// TransactionDecimalsCorrection is used to manipulate precision of a transaction amount value,
// so it can be stored in database as INT64 without loosing too much data
var TransactionDecimalsCorrection = new(big.Int).SetUint64(1000000000)

// TransactionGasCorrection is used to restore the precision on the transaction gas value calculations.
var TransactionGasCorrection = new(big.Int).SetUint64(10000000)

// Transaction represents basic information provided by the API about transaction inside Opera blockchain.
type Transaction struct {
	// BlockHash represents hash of the block where this transaction was in. nil when its pending.
	BlockHash *common.Hash `json:"blockHash"`

	// BlockNumber represents number of the block where this transaction was in. nil when its pending.
	BlockNumber *hexutil.Uint64 `json:"blockNumber"`

	// TimeStamp represents the time stamp of the transaction.
	TimeStamp time.Time

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
	Hash common.Hash `json:"hash"`

	// Nonce represents the number of transactions made by the sender prior to this one.
	Nonce hexutil.Uint64 `json:"nonce"`

	// To represents the address of the receiver. nil when its a contract creation transaction.
	To *common.Address `json:"to,omitempty"`

	// ContractAddress represents the address of contract created, if a contract creation transaction, otherwise nil.
	ContractAddress *common.Address `json:"contract,omitempty"`

	// TrxIndex represents integer of the transaction's index position in the block. nil when its pending.
	TrxIndex *hexutil.Uint `json:"transactionIndex,omitempty"`

	// Value represents value transferred in Wei.
	Value hexutil.Big `json:"value"`

	// Input represents the data send along with the transaction.
	InputData hexutil.Bytes `json:"input"`

	// LargeInput indicates the input data block is large and needs to be loaded separately.
	LargeInput bool `json:"-"`

	// Index represents integer of the transaction's index position in the block.
	Index *hexutil.Uint64 `json:"index,omitempty"`

	// Status represents transaction status; value is either 1 (success) or 0 (failure)
	Status *hexutil.Uint64 `json:"status,omitempty"`

	// Logs represents a list of log records created along with the transaction
	Logs []retypes.Log `json:"logs"`
}

// BsonLog represents the transaction log record data structure for BSON formatting.
type BsonLog struct {
	Address string   `bson:"addr"`
	Topics  []string `bson:"top"`
	Data    []byte   `bson:"data"`
	Index   uint     `bson:"ix"`
	Removed bool     `bson:"rm"`
}

// BsonTransaction represents the transaction data structure for BSON formatting.
type BsonTransaction struct {
	Hash       string    `bson:"_id"`
	Ordinal    uint64    `bson:"orx"`
	BlockID    *uint64   `bson:"blk"`
	BlockHash  *string   `bson:"blk_h"`
	BlkIndex   *uint64   `bson:"bix"`
	From       string    `bson:"from"`
	To         *string   `bson:"to"`
	Value      string    `bson:"value"`
	Amount     int64     `bson:"amo"`
	LargeInput bool      `bson:"large"`
	Input      []byte    `bson:"input"`
	Gas        int64     `bson:"gas_lim"`
	UsedGas    *uint64   `bson:"gas_use"`
	CumGas     *uint64   `bson:"gas_cum"`
	GasPrice   string    `bson:"gas_pri"`
	GasGWei    int64     `bson:"gwx100"`
	Nonce      int64     `bson:"nonce"`
	Contract   *string   `bson:"contr"`
	Status     uint64    `bson:"stat"`
	Stamp      time.Time `bson:"stamp"`
	Logs       []BsonLog `bson:"logs"`
}

// Uid calculates an ordinal index of the transaction referenced.
// The ordinal index of a transaction should be unique across a consistent block chain.
// The calculation gives us about 700 years of index space with 50k blocks per second
// rate + 10 years to fix than. Max number of transactions in a block here is 14bits = 16383.
func (trx *Transaction) Uid() uint64 {
	// is this a processed transaction?
	if trx.Index != nil {
		return (uint64(*trx.BlockNumber) << 14) | (uint64(*trx.Index)&0x3fff)&0x7FFFFFFFFFFFFFFF
	}
	// pending transaction
	return binary.BigEndian.Uint64(trx.Hash[:8]) & 0x7FFFFFFFFFFFFFFF
}

// Marshal returns the JSON encoding of transaction.
func (trx *Transaction) Marshal() ([]byte, error) {
	return json.Marshal(trx)
}

// MarshalBSON creates a BSON representation of the Transaction record.
func (trx *Transaction) MarshalBSON() ([]byte, error) {
	// calculate the value to 9 digits (and 18 billions remain available)
	val := new(big.Int).Div(trx.Value.ToInt(), TransactionDecimalsCorrection)
	gWei := new(big.Int).Div(trx.GasPrice.ToInt(), TransactionGasCorrection)

	// prep the structure for saving
	pom := BsonTransaction{
		Hash:       trx.Hash.String(),
		Ordinal:    trx.Uid(),
		From:       trx.From.String(),
		Gas:        int64(trx.Gas),
		GasPrice:   trx.GasPrice.String(),
		GasGWei:    gWei.Int64(),
		Nonce:      int64(trx.Nonce),
		Value:      trx.Value.String(),
		Amount:     val.Int64(),
		LargeInput: len(trx.InputData) > trxLargeInputWall,
		Stamp:      trx.TimeStamp,
	}

	// store the input data along with the trx
	if !pom.LargeInput {
		pom.Input = trx.InputData
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

		// cumulative gas
		gc := uint64(*trx.CumulativeGasUsed) & 0x7FFFFFFFFFFFFFFF
		pom.CumGas = &gc

		// used gas
		// overcome an error in node API, it reports invalid gas usage on some trx
		gu := uint64(*trx.GasUsed) & 0x7FFFFFFFFFFFFFFF
		if *trx.GasUsed > *trx.CumulativeGasUsed {
			gu = gc
		}
		pom.UsedGas = &gu

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
	trx.Hash = common.HexToHash(row.Hash)
	trx.From = common.HexToAddress(row.From)
	trx.Gas = hexutil.Uint64(row.Gas)
	trx.GasPrice = (hexutil.Big)(*hexutil.MustDecodeBig(row.GasPrice))
	trx.Nonce = hexutil.Uint64(row.Nonce)
	trx.Status = (*hexutil.Uint64)(&row.Status)
	trx.InputData = row.Input
	trx.LargeInput = row.LargeInput
	trx.TimeStamp = row.Stamp

	// try to decode the value
	tv, err := hexutil.DecodeBig(row.Value)
	if err == nil && tv != nil {
		trx.Value = (hexutil.Big)(*tv)
	}

	// pointers
	if row.BlockHash != nil {
		// block hash
		bh := common.HexToHash(*row.BlockHash)
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
