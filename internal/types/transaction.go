package types

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// Transaction represents a basic information provided by the API about transaction inside Opera blockchain.
type Transaction struct {
	Hash         Hash
	Sender       common.Address
	Recipient    *common.Address
	Amount       *big.Int
	Payload      []byte
	AccountNonce uint64
	GasLimit     uint64
	GasUsed      uint64
	GasPrice     big.Int
	Fee          big.Int
	TxIndex      uint64
	BlockHash    Hash
}
