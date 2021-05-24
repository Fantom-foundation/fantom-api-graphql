/*
Package repository implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Opera/Lachesis full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import (
	"context"
	"fantom-api-graphql/internal/logger"
	"fantom-api-graphql/internal/types"
	ftm "github.com/ethereum/go-ethereum/rpc"
	"sync"
	"time"
)

// monBlocksBufferCapacity is the number of new blocks kept in the block processing channel.
const monBlocksBufferCapacity = 50000

// blockMonitor represents a subscription processor capturing new blockchain blocks.
type blockMonitor struct {
	service

	txChan      chan *eventTransaction
	blkChan     chan types.Block
	procChan    chan types.Block
	sigProcStop chan bool
	reScan      chan bool
	con         *ftm.Client
	sub         *ftm.ClientSubscription

	// event broadcast channels
	onBlock       chan *types.Block
	onTransaction chan *types.Transaction
}

// NewBlockMonitor creates a new block monitor instance.
func NewBlockMonitor(con *ftm.Client, buffer chan *eventTransaction, rescan chan bool, repo Repository, log logger.Logger, wg *sync.WaitGroup) *blockMonitor {
	// create new blockScanner instance
	return &blockMonitor{
		service:     newService("block monitor", repo, log, wg),
		txChan:      buffer,
		reScan:      rescan,
		con:         con,
		sigProcStop: make(chan bool, 1),
		blkChan:     make(chan types.Block, monBlocksBufferCapacity),
		procChan:    make(chan types.Block, monBlocksBufferCapacity),
	}
}

// run starts monitoring for new transaction
func (bm *blockMonitor) run() {
	// start go routine for processing
	bm.wg.Add(1)
	go bm.process()

	// start go routine for subscription reader
	bm.wg.Add(1)
	go bm.monitor()
}

// subscribe opens a subscription on the connected Opera/Lachesis full node.
func (bm *blockMonitor) subscribe() error {
	// open subscription
	sub, err := bm.con.Subscribe(context.Background(), "eth", bm.blkChan, "newHeads")
	if err != nil {
		bm.log.Error("can not subscribe to blockchain")
		bm.log.Error(err)
		return err
	}

	// keep the subscription
	bm.sub = sub
	return nil
}

// monitor consumes new blocks from the block channel and route them to target functions.
func (bm *blockMonitor) monitor() {
	// inform about the monitor
	bm.log.Notice("block monitor is running")

	// don't forget to sign off after we are done
	defer func() {
		// make sure to spread the word
		bm.sigProcStop <- true

		// unsubscribe
		bm.log.Notice("block monitor unsubscribed")
		if bm.sub != nil {
			bm.sub.Unsubscribe()
			bm.sub = nil
		}

		// log finish
		bm.log.Notice("block monitor is closed")

		// signal to wait group we are done
		bm.wg.Done()
	}()

	// open subscription
	if err := bm.subscribe(); err != nil {
		bm.log.Criticalf("block monitor subscription failed; %s", err.Error())
		return
	}

	// loop here
	for {
		select {
		case <-bm.sigStop:
			return

		case err, ok := <-bm.sub.Err():
			// do we have a working channel?
			if !ok {
				bm.log.Notice("block monitor subscription has been closed")
				return
			}

			// log issue
			bm.log.Error("block monitor subscription error; %s", err.Error())

			// signal orchestrator to schedule re-scan and restart subscription
			bm.reScan <- true
			return

		case blk := <-bm.blkChan:
			// log the action
			bm.log.Debugf("new block #%d arrived", uint64(blk.Number))

			// extract full block information
			block, err := bm.repo.BlockByNumber(&blk.Number)
			if err != nil {
				bm.log.Errorf("can not process block; %s", err.Error())
				continue
			}

			// push for processing
			bm.procChan <- *block

			// notify event
			bm.onBlock <- block

			// add to the ring cache
			bm.repo.CacheBlock(block)
		}
	}
}

// process pulls blocks from processing queue and act on it as needed.
func (bm *blockMonitor) process() {
	// inform about the monitor
	bm.log.Notice("block processor is running")

	// don't forget to sign off after we are done
	defer func() {
		// log finish
		bm.log.Notice("block processor is closed")

		// signal to wait group we are done
		bm.wg.Done()
	}()

	for {
		select {
		case <-bm.sigProcStop:
			return

		case block, ok := <-bm.procChan:
			// if the channel is closed, no more data will arrive here
			if !ok {
				return
			}

			// any transactions in the block? process them
			if block.Txs != nil && len(block.Txs) > 0 {
				bm.handle(&block)
			}
		}
	}
}

// handle pulls transactions from the incoming block to be processed in the trx dispatch queue.
func (bm *blockMonitor) handle(block *types.Block) {
	// log action
	bm.log.Debugf("block #%d has %d transactions to process", uint64(block.Number), len(block.Txs))

	// loop transaction hashes
	for i, hash := range block.Txs {
		// log action
		bm.log.Debugf("processing transaction #%d of block #%d", i, uint64(block.Number))

		// get transaction
		trx, err := bm.repo.Transaction(hash)
		if err != nil {
			bm.log.Errorf("error getting transaction detail; %s", err.Error())
			continue
		}

		// update time stamp using the block data
		trx.TimeStamp = time.Unix(int64(block.TimeStamp), 0)

		// prep sending struct and push it to the queue
		event := eventTransaction{block: block, trx: trx}
		bm.txChan <- &event

		// notify new transaction
		bm.onTransaction <- trx
	}

	// log action
	bm.log.Debugf("block #%d processed", uint64(block.Number))
}
