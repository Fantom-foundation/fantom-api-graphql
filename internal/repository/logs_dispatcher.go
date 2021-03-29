package repository

import (
	"fantom-api-graphql/internal/logger"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	retypes "github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"sync"
)

// logQueueLength represents the amount og transaction logs
// allowed to be queued at a time before queue writer is slowed down
const logQueueLength = 10000

// eventTrxLog represents a log record to be processed.
type eventTrxLog struct {
	wg *sync.WaitGroup
	retypes.Log
}

// logsDispatcher implements dispatcher of new log events in the blockchain.
type logsDispatcher struct {
	service
	buffer chan *eventTrxLog
}

// trxLogKnownTopics represents a map of known transaction log topics to their handlers.
var trxLogKnownTopics = map[common.Hash]func(*retypes.Log, *logsDispatcher){
	/* SFC1::CreatedDelegation(#address, #toStakerID, amount) */
	common.HexToHash("0xfd8c857fb9acd6f4ad59b8621a2a77825168b7b4b76de9586d08e00d4ed462be"): handleSfcCreatedDelegation,
}

// newLogsDispatcher creates a new transaction logs dispatcher instance.
func newLogsDispatcher(buffer chan *eventTrxLog, repo Repository, log logger.Logger, wg *sync.WaitGroup) *logsDispatcher {
	// create new dispatcher
	return &logsDispatcher{
		service: newService("logs dispatcher", repo, log, wg),
		buffer:  buffer,
	}
}

// run starts the transaction logs dispatcher job
func (ld *logsDispatcher) run() {
	ld.wg.Add(1)
	go ld.dispatch()
}

// dispatch implements the dispatcher reader and router routine.
func (ld *logsDispatcher) dispatch() {
	// log action
	ld.log.Notice("logs dispatcher is running")

	// don't forget to sign off after we are done
	defer func() {
		// log finish
		ld.log.Notice("logs dispatcher is closed")
		ld.wg.Done()
	}()

	// wait for logs and process them
	for {
		// try to read next transaction
		select {
		case log := <-ld.buffer:
			// try to find the topic handler
			handler, ok := trxLogKnownTopics[log.Topics[0]]
			if ok {
				ld.log.Debugf("known topic %s found, processing", log.Topics[0].String())
				handler(&log.Log, ld)
			}

			// mark the processing as finished
			log.wg.Done()

		case <-ld.sigStop:
			return
		}
	}
}

// handleSfcCreatedDelegation handles a new delegation event from SFC v1 and SFC v2 contract.
// event CreatedDelegation(address indexed delegator, uint256 indexed toStakerID, uint256 amount)
func handleSfcCreatedDelegation(log *retypes.Log, ld *logsDispatcher) {
	// get the block
	bn := hexutil.Uint64(log.BlockNumber)
	block, err := ld.repo.BlockByNumber(&bn)
	if err != nil {
		ld.log.Errorf("can not decode CreateDelegation log record; %s", err.Error())
		return
	}

	// decode staker ID (3rd topic) and amount (non-indexed => data)
	stakerID := new(big.Int).SetBytes(log.Topics[2].Bytes())
	amount := new(big.Int).SetBytes(log.Data)

	// make the delegation record
	dl := types.Delegation{
		Transaction:     log.TxHash,
		Address:         common.BytesToAddress(log.Topics[1].Bytes()),
		ToStakerId:      (*hexutil.Big)(stakerID),
		AmountDelegated: (*hexutil.Big)(amount),
		CreatedTime:     block.TimeStamp,
	}

	// store the delegation
	if err := ld.repo.StoreDelegation(&dl); err != nil {
		ld.log.Errorf("failed to store delegation; %s", err.Error())
	}
}
