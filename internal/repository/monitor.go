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
)

// monBlocksBufferCapacity is the number of new blocks kept in the block processing channel.
const monBlocksBufferCapacity = 5000

// BlockMonitor represents a subscription processor capturing new blockchain blocks.
type blockMonitor struct {
	service

	txChan   chan *evtTransaction
	blkChan  chan types.Block
	procChan chan types.Block
	reScan   chan bool
	con      *ftm.Client
	sub      *ftm.ClientSubscription

	// event broadcast channels
	onBlock       chan *types.Block
	onTransaction chan *types.Transaction
}

// NewBlockMonitor creates a new block monitor instance.
func NewBlockMonitor(con *ftm.Client, buffer chan *evtTransaction, rescan chan bool, repo Repository, log logger.Logger, wg *sync.WaitGroup) *blockMonitor {
	// create new scanner instance
	mo := blockMonitor{
		service: newService("block monitor", repo, log, wg),
		txChan:  buffer,
		reScan:  rescan,
		con:     con,
	}

	// start the scanner job
	return &mo
}

// run starts monitoring for new transaction
func (bm *blockMonitor) run() {
	// log
	bm.log.Notice("initializing blockchain monitor")

	// init the monitor
	err := bm.init()
	if err != nil {
		bm.log.Critical("can not monitor blockchain")
		bm.log.Critical(err)
		return
	}

	// start go routine for processing
	bm.wg.Add(1)
	go bm.process()

	// start go routine for subscription reader
	bm.wg.Add(1)
	go bm.monitor()

	// log what we do
	bm.log.Notice("new blocks monitoring and processing started")
}

// init prepares to monitor the Opera/Lachesis blockchain for new blocks.
func (bm *blockMonitor) init() error {
	// open block channel
	bm.blkChan = make(chan types.Block, monBlocksBufferCapacity)
	bm.procChan = make(chan types.Block, monBlocksBufferCapacity)

	// subscribe
	return bm.subscribe()
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
	// don't forget to sign off after we are done
	defer func() {
		// unsubscribe
		bm.sub.Unsubscribe()

		// close block processing channels
		close(bm.blkChan)
		close(bm.procChan)

		// log finish
		bm.log.Notice("block monitor done")

		// signal to wait group we are done
		bm.wg.Done()
	}()

	// received block
	var block types.Block

	// loop here
	for {
		select {
		case <-bm.sigStop:
			return
		case err := <-bm.sub.Err():
			if err != nil {
				// log issue
				bm.log.Error("monitor subscription error; %s", err.Error())

				// signal orchestrator to schedule re-scan and restart subscription
				bm.reScan <- true
			} else {
				bm.log.Notice("monitor subscription has been closed")
			}
			return
		case block = <-bm.blkChan:
			// log the action
			bm.log.Debugf("new block #%d arrived", uint64(block.Number))

			// extract full block information
			block, err := bm.repo.BlockByNumber(&block.Number)
			if err != nil {
				bm.log.Errorf("can not process block; %s", err.Error())
			} else {
				// push for processing
				bm.procChan <- *block

				// notify event
				bm.onBlock <- block
			}
		}
	}
}

// process pulls blocks from processing queue and act on it as needed.
func (bm *blockMonitor) process() {
	// don't forget to sign off after we are done
	defer func() {
		// log finish
		bm.log.Notice("block processor done")

		// signal to wait group we are done
		bm.wg.Done()
	}()

	for {
		// wait to receive next block
		block, ok := <-bm.procChan

		// if the channel is closed, no more data will arrive here
		if !ok {
			return
		}

		// any transactions in the block?
		if block.Txs != nil && len(block.Txs) > 0 {
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

				// prep sending struct and push it to the queue
				event := evtTransaction{block: &block, trx: trx}
				bm.txChan <- &event

				// notify new transaction
				bm.onTransaction <- trx
			}

			// log action
			bm.log.Debugf("block #%d processed", uint64(block.Number))
		}
	}
}
