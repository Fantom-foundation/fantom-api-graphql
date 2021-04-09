// Package types implements different core types of the API.
package types

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	FiErc20TransactionPk        = "_id"
	FiErc20TransactionOrdinal   = "orx"
	FiErc20TransactionSender    = "from"
	FiErc20TransactionRecipient = "to"
	FiErc20TransactionAmount    = "amo"
)

// Erc20Transaction represents an operation with ERC20 token.
type Erc20Transaction struct {
	Transaction common.Hash    `json:"trx"`
	TrxIndex    hexutil.Uint64 `json:"tix"`
	Token       common.Address `json:"erc"`
	Type        int32          `json:"type"`
	Sender      common.Address `json:"from"`
	Recipient   common.Address `json:"to"`
	Amount      hexutil.Big    `json:"amo"`
	TimeStamp   hexutil.Uint64 `json:"ts"`
}

// BsonErc20Transaction represents the BSON i/o struct for an ERC20 operation.
type BsonErc20Transaction struct {
	ID      string `bson:"_id"`
	Tix     uint64 `bson:"tix"`
	Orx     uint64 `bson:"orx"`
	Token   string `bson:"tok"`
	Type    int32  `bson:"type"`
	From    string `bson:"from"`
	To      string `bson:"to"`
	Amo     string `bson:"amo"`
	Created uint64 `bson:"ts"`
}

// Pk generates unique identifier of the ERC20 transaction.
func (etx *Erc20Transaction) Pk() string {
	return ""
}

// OrdinalIndex returns an ordinal index for the given ERC20 transaction.
// We construct the UID from the time the transaction was processed (40 bits = 1099511627775s = 34000 years),
func (etx *Erc20Transaction) OrdinalIndex() uint64 {
	return (uint64(etx.TimeStamp)&0x7FFFFFFFFF)<<24 | uint64(etx.TrxIndex)&0xFFFFFF
}

// MarshalBSON creates a BSON representation of the ERC20 transaction record.
func (etx *Erc20Transaction) MarshalBSON() ([]byte, error) {
	return bson.Marshal(BsonErc20Transaction{
		ID:      etx.Pk(),
		Tix:     uint64(etx.TrxIndex),
		Orx:     etx.OrdinalIndex(),
		Token:   etx.Token.String(),
		Type:    etx.Type,
		From:    etx.Sender.String(),
		To:      etx.Recipient.String(),
		Amo:     etx.Amount.String(),
		Created: uint64(etx.TimeStamp),
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
	return nil
}
