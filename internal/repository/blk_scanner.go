/*
Package repository implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Opera/Lachesis full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import (
	"fantom-api-graphql/internal/logger"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"sync"
	"time"
)

// blockScanner implements blockchain blockScanner used to extract blockchain data to off-chain storage.
type blockScanner struct {
	service
	buffer chan *eventTransaction
	isDone chan bool
}

// newBlockScanner creates new blockchain blockScanner service.
func newBlockScanner(buffer chan *eventTransaction, isDone chan bool, repo Repository, log logger.Logger, wg *sync.WaitGroup) *blockScanner {
	// create new blockScanner instance
	return &blockScanner{
		service: newService("block scanner", repo, log, wg),
		buffer:  buffer,
		isDone:  isDone,
	}
}

// scan initializes the blockScanner and starts scanning
func (bls *blockScanner) run() {
	// get the newest known transaction
	lnb, err := bls.repo.LastKnownBlock()
	if err != nil {
		bls.log.Critical("can not scan blockchain; %sys", err.Error())
		return
	}

	// log what we do
	bls.log.Noticef("block chain scan starts from block #%d", lnb)

	// start blockScanner
	bls.wg.Add(1)
	go bls.scan(lnb)
}

// scan performs the actual blockScanner operation on the missing blocks starting
// from the identified last known block id/number.
func (bls *blockScanner) scan(lnb uint64) {
	// what is the current block
	var (
		current = hexutil.Uint64(lnb)
		block   *types.Block
		err     error
		index   int
		trx     *types.Transaction
		toSend  *eventTransaction
		stopLog = make(chan bool, 1)
	)

	// don't forget to sign off after we are done
	defer func() {
		// signal we are done with the sync
		bls.isDone <- true
		stopLog <- true

		// log finish
		bls.log.Notice("block scanner done")

		// signal to wait group we are done
		bls.wg.Done()
	}()

	// inform about block scanner progress sparsely to prevent log flood
	go func() {
		// track the progress
		start := time.Now()

		bls.log.Infof("block scanner on block #%d", uint64(current))
		for {
			select {
			case <-stopLog:
				bls.log.Infof("block scanner finished on block #%d", uint64(current))
				return
			case <-time.After(5 * time.Second):
				eta := time.Now().Add(time.Duration(int64(time.Now().Sub(start).Nanoseconds() * (int64(lnb) / int64(current)))))
				bls.log.Infof("block scanner reached block #%d, ETA %s", uint64(current), eta.Format(time.Stamp))
			}
		}
	}()

	// do the scan
	for {
		// do we need a new block?
		if block == nil || block.Txs == nil || len(block.Txs) <= index {
			// try to get the next block
			block, err = bls.repo.BlockByNumber(&current)
			if err != nil {
				return
			}

			// reset the current block tx index and advance to the next block
			index = 0
			current += 1
		}

		// do we have something to send?
		if toSend != nil || len(block.Txs) > index {
			// do we need to pull next transaction?
			if toSend == nil {
				// log action
				bls.log.Debugf("loading transaction #%d of block #%d", index, uint64(block.Number))

				// get transaction
				trx, err = bls.repo.Transaction(block.Txs[index])
				if err != nil {
					return
				}

				// prep sending struct and advance transaction index
				toSend = &eventTransaction{block: block, trx: trx}
				index++
			}

			// try to send
			select {
			case <-bls.sigStop:
				// stop signal received?
				return
			case bls.buffer <- toSend:
				// we did send it and now we need next one
				toSend = nil
			default:
			}
		}
	}
}
