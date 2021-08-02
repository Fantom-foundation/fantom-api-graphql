/*
Package repository implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Opera/Lachesis full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import (
	"fantom-api-graphql/internal/repository"
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
	or      *Orchestrator
	sigStop chan bool
	onBlock chan *types.Block

	inBlock        chan *types.Block
	outTransaction chan *eventTrx
}

// name returns the name of the service.
func (bld *blockDispatcher) name() string {
	return "block dispatcher"
}

// init prepares the block dispatcher to perform its function.
func (bld *blockDispatcher) init() {
	// make channels we need
	bld.sigStop = make(chan bool, 1)
	bld.outTransaction = make(chan *eventTrx, trxBufferCapacity)
}

// run starts the block dispatcher
func (bld *blockDispatcher) run() {
	// make sure we are orchestrated
	if bld.or == nil {
		panic(fmt.Errorf("no orchestrator set"))
	}

	// signal orchestrator we started and go
	bld.or.started(bld)
	go bld.dispatch()
}

// close terminates the block dispatcher.
func (bld *blockDispatcher) close() {
	bld.sigStop <- true
}

// dispatch collects blocks from an input channel
// and processes them.
func (bld *blockDispatcher) dispatch() {
	// make sure to clean up
	defer func() {
		// close our channels
		close(bld.sigStop)
		close(bld.outTransaction)

		// signal we are done
		bld.or.finished(bld)
	}()

	// loop here
	for {
		select {
		case <-bld.sigStop:
			return

		case blk, ok := <-bld.inBlock:
			// do we have a working channel?
			if !ok {
				bld.or.log.Notice("block channel closed, terminating %s", bld.name())
				return
			}

			// process the new block
			bld.or.log.Debugf("block #%d arrived", uint64(blk.Number))
			bld.process(blk)

			// broadcast the block event
			bld.onBlock <- blk

			// add the block to the ring
			repository.R().CacheBlock(blk)
		}
	}
}

// process the given block by loading its content and sending block transactions
// into the trx dispatcher.
func (bld *blockDispatcher) process(blk *types.Block) {
	// any transactions within?
	if blk.Txs == nil || len(blk.Txs) == 0 {
		return
	}

	// log the situation
	bld.or.log.Debugf("%d transaction found in block #%d", len(blk.Txs), blk.Number)

	// loop transactions
	for i, th := range blk.Txs {
		// log
		bld.or.log.Debugf("loading trx #%d from block #%d", i, blk.Number)

		// load the transaction
		trx := bld.load(blk, th)
		if trx != nil {
			// queue and broadcast the transaction
			bld.outTransaction <- &eventTrx{
				blk: blk,
				trx: trx,
			}
		}
	}

	// log the situation
	bld.or.log.Debugf("block #%d processed", blk.Number)
}

// load a transaction detail from repository, if possible.
func (bld *blockDispatcher) load(blk *types.Block, th *common.Hash) *types.Transaction {
	// get transaction
	trx, err := repository.R().Transaction(th)
	if err != nil {
		bld.or.log.Errorf("transaction %s detail not available; %s", th.String(), err.Error())
		return nil
	}

	// update time stamp using the block data
	trx.TimeStamp = time.Unix(int64(blk.TimeStamp), 0)
	return trx
}
