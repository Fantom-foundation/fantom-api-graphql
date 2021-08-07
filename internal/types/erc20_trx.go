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
	FiErc20TransactionPk        = "_id"
	FiErc20TransactionOrdinal   = "orx"
	FiErc20TransactionToken     = "tok"
	FiErc20TransactionSender    = "from"
	FiErc20TransactionRecipient = "to"
	FiErc20TransactionType      = "type"
	FiErc20TransactionStamp     = "stamp"

	// ERC20TrxTypeTransfer represents transaction for transfers.
	ERC20TrxTypeTransfer     = 1
	ERC20TrxTypeNameTransfer = "TRANSFER"
	ERC20TrxTypeNameMint     = "MINT"
	ERC20TrxTypeNameBurn     = "BURN"

	// ERC20TrxTypeApproval represents transaction for granting approvals.
	ERC20TrxTypeApproval     = 2
	ERC20TrxTypeNameApproval = "APPROVAL"
)

// Erc20Transaction represents an operation with ERC20 token.
type Erc20Transaction struct {
	ID           string         `json:"_id"`
	Transaction  common.Hash    `json:"trx"`
	TrxIndex     hexutil.Uint64 `json:"tix"`
	TokenAddress common.Address `json:"erc"`
	TokenType    string         `json:"tty"`
	Type         int32          `json:"type"`
	Sender       common.Address `json:"from"`
	Recipient    common.Address `json:"to"`
	Amount       hexutil.Big    `json:"amo"`
	TimeStamp    hexutil.Uint64 `json:"ts"`
}

// BsonErc20Transaction represents the BSON i/o struct for an ERC20 operation.
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
	Created   uint64    `bson:"ts"`
	Value     int64     `bson:"val"`
	Stamp     time.Time `bson:"stamp"`
}

// Erc20TrxTypeByName returns numeric type of the ERC20 transaction by its name.
// Returns nil if the name is not recognized.
func Erc20TrxTypeByName(name string) *int32 {
	// decode the transaction type filter
	var txType *int32
	switch name {
	case ERC20TrxTypeNameMint:
		i := int32(ERC20TrxTypeTransfer)
		txType = &i
	case ERC20TrxTypeNameBurn:
		i := int32(ERC20TrxTypeTransfer)
		txType = &i
	case ERC20TrxTypeNameTransfer:
		i := int32(ERC20TrxTypeTransfer)
		txType = &i
	case ERC20TrxTypeNameApproval:
		i := int32(ERC20TrxTypeApproval)
		txType = &i
	}
	return txType
}

// Pk generates unique identifier of the ERC20 transaction.
// We use 10 bytes of the sender address, 10 bytes of the recipient address,
// and 10 bytes of the token address. The last 8 bytes are ten XOR-ed
// with the transaction time stamp.
func (etx *Erc20Transaction) Pk() string {
	// make the base PK from the trx hash and log index
	bytes := make([]byte, 32)
	copy(bytes, etx.Sender.Bytes()[:10])
	copy(bytes[10:], etx.Recipient.Bytes()[:10])
	copy(bytes[20:], etx.TokenAddress.Bytes()[:10])

	// xor in the timestamp and the root trx index
	ts := make([]byte, 8)
	binary.BigEndian.PutUint64(ts, ((uint64(etx.TimeStamp)&0x7FFFFFFFFF)<<24)|(uint64(etx.TrxIndex)&0xFFFFFF))
	for i := 0; i < 8; i++ {
		bytes[24+i] = bytes[24+i] ^ ts[i]
	}
	return hexutil.Encode(bytes)
}

// OrdinalIndex returns an ordinal index for the given ERC20 transaction.
// We construct the UID from the time the transaction was processed (40 bits = 1099511627775s = 34000 years),
// and the small fraction of the token address to distinguish between different transfers on the same block.
func (etx *Erc20Transaction) OrdinalIndex() uint64 {
	ts := make([]byte, 8)
	binary.BigEndian.PutUint64(ts, (uint64(etx.TimeStamp)&0x7FFFFFFFFF)<<24)
	copy(ts[5:], etx.TokenAddress.Bytes()[:3])
	return binary.BigEndian.Uint64(ts)
}

// MarshalBSON creates a BSON representation of the ERC20 transaction record.
func (etx *Erc20Transaction) MarshalBSON() ([]byte, error) {
	// calculate transfer value for ERC20 tokens
	val := new(big.Int)
	if etx.TokenType == AccountTypeERC20Token {
		val = val.Div(etx.Amount.ToInt(), TransactionDecimalsCorrection)
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
		Created:   uint64(etx.TimeStamp),
		Value:     val.Int64(),
		Stamp:     time.Unix(int64(etx.TimeStamp), 0),
	})
}

// UnmarshalBSON updates the value from BSON source.
func (etx *Erc20Transaction) UnmarshalBSON(data []byte) (err error) {
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
	etx.TimeStamp = hexutil.Uint64(row.Created)
	etx.TokenAddress = common.HexToAddress(row.Token)
	etx.TokenType = row.TokenType
	etx.Type = row.Type
	etx.Sender = common.HexToAddress(row.From)
	etx.Recipient = common.HexToAddress(row.To)
	etx.Amount = (hexutil.Big)(*hexutil.MustDecodeBig(row.Amo))
	return nil
}
