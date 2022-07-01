// Package svc implements blockchain data processing services.
package svc

import (
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	retypes "github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/atomic"
	"sync"
	"time"
)

// trxAddressQueueCapacity is the number of addresses kept in the dispatch buffer.
const trxAddressQueueCapacity = 1000

// trxLogQueueCapacity is the number of transaction logs kept in the dispatch buffer.
const trxLogQueueCapacity = 5000

// trxDispatchBlockUpdateTicker represents the period of block registry updater.
const trxDispatchBlockUpdateTicker = 15 * time.Second

// eventAcc represents a structure of a mentioned account.
type eventAcc struct {
	watchDog *sync.WaitGroup
	addr     *common.Address
	act      string
	blk      *types.Block
	trx      *types.Transaction
	deploy   *common.Hash
}

// trxDispatcher implements dispatcher of new transactions in the blockchain.
type trxDispatcher struct {
	service
	onTransaction  chan *types.Transaction
	bot            *time.Ticker
	blkObserver    *atomic.Uint64
	inTransaction  chan *eventTrx
	outTransaction chan *eventTrx
	outAccount     chan *eventAcc
	outLog         chan *types.LogRecord
}

// name returns the name of the service used by orchestrator.
func (trd *trxDispatcher) name() string {
	return "transaction dispatcher"
}

// init prepares the transaction dispatcher to perform its function.
func (trd *trxDispatcher) init() {
	trd.sigStop = make(chan struct{})
	trd.blkObserver = atomic.NewUint64(1)
	trd.outAccount = make(chan *eventAcc, trxAddressQueueCapacity)
	trd.outLog = make(chan *types.LogRecord, trxLogQueueCapacity)
	trd.outTransaction = make(chan *eventTrx, trxLogQueueCapacity)
}

// run starts the transaction dispatcher job
func (trd *trxDispatcher) run() {
	// make sure we are orchestrated
	if trd.mgr == nil {
		panic(fmt.Errorf("no svc manager set on %s", trd.name()))
	}

	// start the block observer ticker
	trd.bot = time.NewTicker(trxDispatchBlockUpdateTicker)

	// signal orchestrator we started and go
	trd.mgr.started(trd)
	go trd.execute()
}

// close terminates the block dispatcher.
func (trd *trxDispatcher) close() {
	if trd.bot != nil {
		trd.bot.Stop()
	}
	if trd.sigStop != nil {
		close(trd.sigStop)
	}
}

// execute implements the dispatcher reader and router routine.
func (trd *trxDispatcher) execute() {
	// don't forget to sign off after we are done
	defer func() {
		close(trd.outAccount)
		close(trd.outLog)
		close(trd.outTransaction)

		trd.mgr.finished(trd)
	}()

	// wait for transactions and process them
	for {
		// try to read next transaction
		select {
		case <-trd.sigStop:
			return
		case <-trd.bot.C:
			trd.updateLastSeenBlock()
		case evt, ok := <-trd.inTransaction:
			// is the channel even available for reading
			if !ok {
				log.Notice("trx channel closed, terminating %s", trd.name())
				return
			}

			if evt.blk == nil || evt.trx == nil {
				log.Criticalf("dispatcher dry loop")
				continue
			}
			trd.process(evt)
		}
	}
}

// updateLastSeenBlock updates the information about last known block
// in the persistent database.
func (trd *trxDispatcher) updateLastSeenBlock() {
	// get the current value
	lsb := trd.blkObserver.Load()
	log.Noticef("last seen block is #%d", lsb)

	// make the change in the database so the progress persists
	err := repo.UpdateLastKnownBlock((*hexutil.Uint64)(&lsb))
	if err != nil {
		log.Errorf("could not update last seen block; %s", err.Error())
	}
}

// process the given transaction event into the required targets.
func (trd *trxDispatcher) process(evt *eventTrx) {
	// send the transaction out for burns processing
	select {
	case trd.outTransaction <- evt:
	case <-trd.sigStop:
		return
	}

	// process transaction accounts; exit if terminated
	var wg sync.WaitGroup
	if !trd.pushAccounts(evt, &wg) {
		return
	}

	// process transaction logs; exit if terminated
	for _, lg := range evt.trx.Logs {
		if !trd.pushLog(lg, evt.blk, evt.trx, &wg) {
			return
		}
	}

	// store the transaction into the database once the processing is done
	// we spawn a lot of go-routines here, so we should test the optimal queue length above
	go trd.waitAndStore(evt, &wg)

	// broadcast new transaction; if it can not be broadcast quickly, skip
	select {
	case trd.onTransaction <- evt.trx:
	case <-time.After(200 * time.Millisecond):
	case <-trd.sigStop:
	}
}

// waitAndStore waits for the transaction processing to finish and stores the transaction into db.
func (trd *trxDispatcher) waitAndStore(evt *eventTrx, wg *sync.WaitGroup) {
	// wait until all the sub-processors finish their job
	wg.Wait()
	if err := repo.StoreTransaction(evt.blk, evt.trx); err != nil {
		log.Errorf("can not store trx %s from block #%d", evt.trx.Hash.String(), evt.blk.Number)
	}

	repo.IncTrxCountEstimate(1)
	repo.CacheTransaction(evt.trx)
	trd.blkObserver.Store(uint64(evt.blk.Number))
}

// pushAccounts pushes given transaction accounts on both sides observing terminate signal on process.
func (trd *trxDispatcher) pushAccounts(evt *eventTrx, wg *sync.WaitGroup) bool {
	// the sender is always present
	if !trd.pushAccount(types.AccountTypeWallet, &evt.trx.From, evt.blk, evt.trx, wg) {
		return false
	}

	// do we have a recipient?
	if evt.trx.To != nil && !trd.pushAccount(types.AccountTypeWallet, evt.trx.To, evt.blk, evt.trx, wg) {
		return false
	}

	// if there is no contract created, we are done here
	if evt.trx.ContractAddress == nil {
		return true
	}

	// queue the new contract to be processed as well
	log.Debugf("contract %s found at trx %s", evt.trx.ContractAddress.String(), evt.trx.Hash.String())
	return trd.pushAccount(types.AccountTypeContract, evt.trx.ContractAddress, evt.blk, evt.trx, wg)
}

// pushAccount pushes given account event to output queue observing terminate signal.
func (trd *trxDispatcher) pushAccount(at string, adr *common.Address, blk *types.Block, trx *types.Transaction, wg *sync.WaitGroup) bool {
	wg.Add(1)
	select {
	case trd.outAccount <- &eventAcc{
		watchDog: wg,
		addr:     adr,
		act:      at,
		blk:      blk,
		trx:      trx,
		deploy:   nil,
	}:
	case <-trd.sigStop:
		return false
	}
	return true
}

// pushLog pushes specified log record into a processing queue observing terminate signal.
func (trd *trxDispatcher) pushLog(lg retypes.Log, blk *types.Block, trx *types.Transaction, wg *sync.WaitGroup) bool {
	wg.Add(1)
	select {
	case trd.outLog <- &types.LogRecord{
		WatchDog: wg,
		Block:    blk,
		Trx:      trx,
		Log:      lg,
	}:
	case <-trd.sigStop:
		return false
	}
	return true
}
