// Package types implements different core types of the API.
package types

import (
	retypes "github.com/ethereum/go-ethereum/core/types"
	"sync"
)

// LogRecord represents a log record to be processed.
type LogRecord struct {
	WatchDog *sync.WaitGroup
	Block    *Block
	Trx      *Transaction
	retypes.Log
}
