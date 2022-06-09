// Package types implements different core types of the API.
package types

import (
	"encoding/binary"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"math/big"
	"time"
)

const (
	// TokenTrxTypeTransfer represents token transfer transaction.
	TokenTrxTypeTransfer = 1

	// TokenTrxTypeApproval represents token transfer approval.
	TokenTrxTypeApproval = 2

	// TokenTrxTypeMint represents token minting (transfer from 0x0).
	TokenTrxTypeMint = 3

	// TokenTrxTypeBurn represents token burning (transfer into 0x0).
	TokenTrxTypeBurn = 4

	// TokenTransactionTargetDecimals represents the number of decimals we want
	// for normalized transaction value, calculated from the amount of tokens transferred.
	TokenTransactionTargetDecimals = 4
)

// TokenTransaction represents an operation with ERC20 token.
type TokenTransaction struct {
	ID           string         `bson:"_id"`
	TokenAddress common.Address `bson:"addr"`    // contract address
	TrxType      int32          `bson:"tx_type"` // Transfer/Mint/Approval...
	Sender       common.Address `bson:"from"`
	Recipient    common.Address `bson:"to"`
	Amount       hexutil.Big    `bson:"amo"`
	Value        int64          `bson:"value"`
	Decimals     int8           `bson:"decimals"`

	Transaction  common.Hash `bson:"trx"`       // hash of a chain transaction bearing the token transaction
	TimeStamp    time.Time   `bson:"ts"`        // when the block(!) was collated
	BlockNumber  int64       `bson:"block"`     // number of the block
	TrxIndex     int64       `bson:"trx_index"` // index of the transaction in the block
	LogIndex     uint32      `bson:"log_index"` // index of the log in the block
	OrdinalIndex int64       `bson:"ordinal"`   // pre-generated ordinal index of the transaction
}

// Pk generates unique identifier of the token transaction from the transaction data.
func (etx *TokenTransaction) Pk() string {
	bytes := make([]byte, 12)
	binary.BigEndian.PutUint64(bytes[0:8], uint64(etx.BlockNumber)) // unique number of the block
	binary.BigEndian.PutUint32(bytes[8:12], etx.LogIndex)           // index of log event in the block
	return hexutil.Encode(bytes)
}

// Index generates ordinal index used for transactions list segmentation.
func (etx *TokenTransaction) Index() int64 {
	ordinal := make([]byte, 8)
	binary.BigEndian.PutUint64(ordinal, uint64((etx.TimeStamp.Unix()&0x7FFFFFFFFF)<<24))

	logIndex := make([]byte, 4)
	binary.BigEndian.PutUint32(logIndex, etx.LogIndex)

	// use transaction hash as base of salt
	// XOR with logIndex to distinguish individual contract emitted events
	trxHash := etx.Transaction.Bytes()
	ordinal[5] = trxHash[0] ^ logIndex[1]
	ordinal[6] = trxHash[1] ^ logIndex[2]
	ordinal[7] = trxHash[2] ^ logIndex[3]

	return int64(binary.BigEndian.Uint64(ordinal))
}

// WithOrdinalIndex updates an ordinal index (field for deterministic sorting) for the given token transaction.
// We construct the index from the time the transaction was processed (40 bits = 1099511627775s = 34000 years),
// salted by the transaction hash, the event log index (index of the log in the block)
// and sequence number of transaction in batch transfer event.
func (etx *TokenTransaction) WithOrdinalIndex() *TokenTransaction {
	etx.OrdinalIndex = etx.Index()
	return etx
}

// WithValue updates the normalized value of the
func (etx *TokenTransaction) WithValue(decimals int8) *TokenTransaction {
	etx.Decimals = decimals

	// if actual number of decimals on the token is lower than the target
	// all we need to do is to get the int64 value from the amount directly
	if decimals <= TokenTransactionTargetDecimals {
		etx.Value = etx.Amount.ToInt().Int64()
		return etx
	}

	// we need to reduce decimals to get the desired precision; so we divide the amount by 10^(decimals diff)
	etx.Value = new(big.Int).Div(
		etx.Amount.ToInt(),
		math.Exp(big.NewInt(10), big.NewInt(int64(decimals-TokenTransactionTargetDecimals))),
	).Int64()
	return etx
}
