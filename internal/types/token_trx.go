// Package types implements different core types of the API.
package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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

	// TokenTrxTypeApprovalForAll represents universal token transfer approval.
	TokenTrxTypeApprovalForAll = 5

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
	Decimals     uint8          `bson:"decimals"`

	Transaction  common.Hash `bson:"trx"`       // hash of a chain transaction bearing the token transaction
	TimeStamp    time.Time   `bson:"ts"`        // when the block(!) was collated
	BlockNumber  uint64      `bson:"block"`     // number of the block
	TrxIndex     uint64      `bson:"trx_index"` // index of the transaction in the block
	LogIndex     uint64      `bson:"log_index"` // index of the log in the block
	OrdinalIndex int64       `bson:"ordinal"`   // pre-generated ordinal index of the transaction
}
