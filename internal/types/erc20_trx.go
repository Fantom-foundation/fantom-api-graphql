// Package types implements different core types of the API.
package types

import (
	"encoding/binary"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson"
	"math/big"
	"time"
)

const (
	FiTokenTransactionPk        = "_id"
	FiTokenTransactionCallHash  = "trx"
	FiTokenTransactionOrdinal   = "orx"
	FiTokenTransactionTokenType = "tty"
	FiTokenTransactionToken     = "tok"
	FiTokenTransactionTokenId   = "tid"
	FiTokenTransactionType      = "type"
	FiTokenTransactionSender    = "from"
	FiTokenTransactionRecipient = "to"

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
)

// TokenTransaction represents an operation with ERC20 token.
type TokenTransaction struct {
	ID           string         `json:"_id"`
	Transaction  common.Hash    `json:"trx"`  // hash of the transaction
	TrxIndex     hexutil.Uint64 `json:"tix"`  // index of the transaction in the block
	TokenAddress common.Address `json:"erc"`  // contract address
	TokenType    string         `json:"tty"`  // ERC20/ERC721/ERC1155...
	Type         int32          `json:"type"` // Transfer/Mint/Approval...
	Sender       common.Address `json:"from"`
	Recipient    common.Address `json:"to"`
	Amount       hexutil.Big    `json:"amo"`
	TokenId      hexutil.Big    `json:"tid"` // for multi-token contracts (ERC-721/ERC-1155)
	TimeStamp    hexutil.Uint64 `json:"ts"`  // when the block(!) was collated
	BlockNumber  uint64         // number of the block
	LogIndex     uint           // index of the log in the block - only for OrdinalIndex / Pk generating
	Seq          uint16         // index of transfer in one log event - only for Pk generating
}

// BsonErc20Transaction represents the BSON i/o struct for an ERC20 operation.
// Used for saving transactions into MongoDB storage.
type BsonErc20Transaction struct {
	ID        string    `bson:"_id"`
	Trx       string    `bson:"trx"`
	Tix       uint64    `bson:"tix"`
	Orx       uint64    `bson:"orx"`
	Token     string    `bson:"tok"`
	TokenType string    `bson:"tty"`
	Type      int32     `bson:"type"`
	From      string    `bson:"from"`
	To        string    `bson:"to"`
	Amo       string    `bson:"amo"`
	TokenId   string    `bson:"tid"`
	TimeStamp uint64    `bson:"ts"`
	Value     int64     `bson:"val"`
	Stamp     time.Time `bson:"stamp"`
}

// Pk generates unique identifier of the ERC20 transaction from the transaction data.
func (etx *TokenTransaction) Pk() string {
	bytes := make([]byte, 14)
	binary.BigEndian.PutUint64(bytes[0:8], etx.BlockNumber)       // unique number of the block
	binary.BigEndian.PutUint32(bytes[8:12], uint32(etx.LogIndex)) // index of log event in the block
	binary.BigEndian.PutUint16(bytes[12:14], etx.Seq)             // one log event can be batch of multiple transfers
	return hexutil.Encode(bytes)
}

// OrdinalIndex returns an ordinal index (field for deterministic sorting) for the given ERC20 transaction.
// We construct the UID from the time the transaction was processed (40 bits = 1099511627775s = 34000 years),
// salted by the transaction hash, the event log index (index of the log in the block)
// and sequence number of transaction in batch transfer event.
func (etx *TokenTransaction) OrdinalIndex() uint64 {
	ordinal := make([]byte, 8)
	binary.BigEndian.PutUint64(ordinal, (uint64(etx.TimeStamp)&0x7FFFFFFFFF)<<24)

	trxHash := etx.Transaction.Bytes()

	logIndex := make([]byte, 4)
	binary.BigEndian.PutUint32(logIndex, uint32(etx.LogIndex))

	seq := make([]byte, 2)
	binary.BigEndian.PutUint16(seq, etx.Seq)

	// use transaction hash as base of salt
	// XOR with logIndex to distinguish individual contract emitted events
	// XOR with seq to distinguish multiple transfers in one batch transfer event
	ordinal[5] = trxHash[0] ^ logIndex[1] ^ seq[1]
	ordinal[6] = trxHash[1] ^ logIndex[2] ^ seq[0]
	ordinal[7] = trxHash[2] ^ logIndex[3]
	return binary.BigEndian.Uint64(ordinal)
}

// MarshalBSON creates a BSON representation of the ERC20 transaction record.
func (etx *TokenTransaction) MarshalBSON() ([]byte, error) {
	// calculate transfer value for ERC20 tokens
	var val *big.Int
	if etx.TokenType == AccountTypeERC20Token {
		val = new(big.Int).Div(etx.Amount.ToInt(), TransactionDecimalsCorrection)
	} else {
		val = etx.Amount.ToInt()
	}

	// make the record and encode it
	return bson.Marshal(BsonErc20Transaction{
		ID:        etx.Pk(),
		Trx:       etx.Transaction.String(),
		Tix:       uint64(etx.TrxIndex),
		Orx:       etx.OrdinalIndex(),
		Token:     etx.TokenAddress.String(),
		Type:      etx.Type,
		TokenType: etx.TokenType,
		From:      etx.Sender.String(),
		To:        etx.Recipient.String(),
		Amo:       etx.Amount.String(),
		TokenId:   etx.TokenId.String(),
		TimeStamp: uint64(etx.TimeStamp),
		Value:     val.Int64(),
		Stamp:     time.Unix(int64(etx.TimeStamp), 0),
	})
}

// UnmarshalBSON updates the value from BSON source.
func (etx *TokenTransaction) UnmarshalBSON(data []byte) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("can not decode ERC20 transfer; %s", err.Error())
		}
	}()

	// try to decode the BSON data
	var row BsonErc20Transaction
	if err = bson.Unmarshal(data, &row); err != nil {
		return err
	}

	// copy the data
	etx.ID = row.ID
	etx.Transaction = common.HexToHash(row.Trx)
	etx.TrxIndex = hexutil.Uint64(row.Tix)
	etx.TimeStamp = hexutil.Uint64(row.TimeStamp)
	etx.TokenAddress = common.HexToAddress(row.Token)
	etx.TokenType = row.TokenType
	etx.Type = row.Type
	etx.Sender = common.HexToAddress(row.From)
	etx.Recipient = common.HexToAddress(row.To)
	etx.Amount = (hexutil.Big)(*hexutil.MustDecodeBig(row.Amo))
	etx.TokenId = (hexutil.Big)(*hexutil.MustDecodeBig(row.TokenId))
	return nil
}
