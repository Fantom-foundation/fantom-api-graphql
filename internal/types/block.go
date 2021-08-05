// Package types implements different core types of the API.
package types

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Block represents basic information provided by the API about block inside Opera blockchain.
type Block struct {
	// Number represents the block number. nil when its pending block.
	Number hexutil.Uint64 `json:"number"`

	// Hash represents hash of the block. nil when its pending block.
	Hash common.Hash `json:"hash"`

	// ParentHash represents hash of the parent block.
	ParentHash common.Hash `json:"parentHash"`

	// Miner represents the address of the beneficiary to whom the mining rewards were given.
	Miner common.Address `json:"miner"`

	// StateRoot represents the hash of the trie state root.
	StateRoot  common.Hash `json:"stateRoot"`

	// Difficulty represents integer of the difficulty for this block.
	Difficulty hexutil.Uint64 `json:"difficulty"`

	// Size represents the size of this block in bytes.
	Size hexutil.Uint64 `json:"size"`

	// GasLimit represents the maximum gas allowed in this block.
	GasLimit hexutil.Uint64 `json:"gasLimit"`

	// GasUsed represents the actual total used gas by all transactions in this block.
	GasUsed hexutil.Uint64 `json:"gasUsed"`

	// TimeStamp represents the unix timestamp for when the block was collated.
	TimeStamp hexutil.Uint64 `json:"timestamp"`

	// Txs represents array of 32 bytes hashes of transactions included in the block.
	Txs []*common.Hash `json:"transactions"`
}

// UnmarshalBlock parses the JSON-encoded block data.
func UnmarshalBlock(data []byte) (*Block, error) {
	var blk Block
	err := json.Unmarshal(data, &blk)
	return &blk, err
}

// Marshal returns the JSON encoding of block.
func (b *Block) Marshal() ([]byte, error) {
	return json.Marshal(b)
}
