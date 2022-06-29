// Package svc implements blockchain data processing services.
package svc

import (
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"time"
)

// trxBufferCapacity is the number of new packed transactions kept in the trx channel.
const trxBufferCapacity = 50000

// eventTrx represents a packed transaction event
// sent between block dispatcher and transaction dispatcher
type eventTrx struct {
	blk *types.Block
	trx *types.Transaction
}

// blockDispatcher implements a service responsible for processing new blocks on the blockchain.
type blockDispatcher struct {
	service
	onBlock        chan *types.Block
	inBlock        chan *types.Block
	outTransaction chan *eventTrx
	outDispatched  chan uint64
}

// name returns the name of the service used by orchestrator.
func (bld *blockDispatcher) name() string {
	return "block dispatcher"
}

// init prepares the block dispatcher to perform its function.
func (bld *blockDispatcher) init() {
	bld.sigStop = make(chan struct{})
	bld.outTransaction = make(chan *eventTrx, trxBufferCapacity)
	bld.outDispatched = make(chan uint64, blsBlockBufferCapacity)
}

// run starts the block dispatcher
func (bld *blockDispatcher) run() {
	// make sure we are orchestrated
	if bld.mgr == nil {
		panic(fmt.Errorf("no svc manager set on %s", bld.name()))
	}

	// signal orchestrator we started and go
	bld.mgr.started(bld)
	go bld.execute()
}

// execute collects blocks from an input channel
// and processes them.
func (bld *blockDispatcher) execute() {
	// make sure to clean up
	defer func() {
		// close our channels
		close(bld.outTransaction)
		close(bld.outDispatched)

		// signal we are done
		bld.mgr.finished(bld)
	}()

	// loop here
	for {
		select {
		case <-bld.sigStop:
			return

		case blk, ok := <-bld.inBlock:
			// do we have a working channel?
			if !ok {
				log.Notice("block channel closed, terminating %s", bld.name())
				return
			}

			// process the new block
			log.Debugf("block #%d arrived", uint64(blk.Number))
			if !bld.process(blk) {
				continue
			}

			// broadcast the block event
			select {
			case bld.onBlock <- blk:
			case <-time.After(200 * time.Millisecond):
			}

			// add the block to the ring
			repo.CacheBlock(blk)
		}
	}
}

// process the given block by loading its content and sending block transactions
// into the trx dispatcher. Observe terminate signal.
func (bld *blockDispatcher) process(blk *types.Block) bool {
	// dispatched block number is used by the block scanner
	// to keep track of the work done vs. work pending
	select {
	case bld.outDispatched <- uint64(blk.Number):
	case <-bld.sigStop:
		return false
	}

	if blk.Txs == nil || len(blk.Txs) == 0 {
		log.Debugf("empty block #%d processed", blk.Number)
		return true
	}

	log.Debugf("%d transaction found in block #%d", len(blk.Txs), blk.Number)
	if !bld.processTxs(blk) {
		return false
	}

	log.Debugf("block #%d processed", blk.Number)
	return true
}

// processTxs loops all the transactions in the block and pushes them
// into the transaction dispatcher queue observing the term signal.
func (bld *blockDispatcher) processTxs(blk *types.Block) bool {
	for i, th := range blk.Txs {
		log.Debugf("loading trx #%d from block #%d", i, blk.Number)
		trx := bld.load(blk, th)
		if trx != nil {
			// queue and broadcast the transaction
			select {
			case bld.outTransaction <- &eventTrx{
				blk: blk,
				trx: trx,
			}:
			case <-bld.sigStop:
				return false
			}
		}
	}
	return true
}

// load a transaction detail from repository, if possible.
func (bld *blockDispatcher) load(blk *types.Block, th *common.Hash) *types.Transaction {
	// get transaction
	trx, err := repo.Transaction(th)
	if err != nil {
		log.Errorf("transaction %s detail not available; %s", th.String(), err.Error())
		return nil
	}

	// update time stamp using the block data
	trx.TimeStamp = time.Unix(int64(blk.TimeStamp), 0)
	return trx
}
