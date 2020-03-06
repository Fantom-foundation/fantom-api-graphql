package types

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// Account represents an Opera account at the blockchain.
type Account struct {
	Addr    common.Address
	Balance big.Int
	Nonce   uint64
	Txs     []*AccountTransaction
}

// TransactionDirection represents a direction of a Transaction at blockchain from an Account perspective.
type TransactionDirection int

const (
	INCOMING TransactionDirection = iota
	OUTGOING
)

// AccountTransaction represents a Transaction associated with an Account at the blockchain.
type AccountTransaction struct {
	Account   *Account
	Direction TransactionDirection
	Hash      Hash
}
