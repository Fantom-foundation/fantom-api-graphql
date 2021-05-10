/*
Package repository implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Opera/Lachesis full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import (
	"fantom-api-graphql/internal/config"
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
	cmd    *config.RepoCmd
}

// newBlockScanner creates new blockchain blockScanner service.
func newBlockScanner(buffer chan *eventTransaction, isDone chan bool, repo Repository, log logger.Logger, wg *sync.WaitGroup, cmd *config.RepoCmd) *blockScanner {
	// create new blockScanner instance
	return &blockScanner{
		service: newService("block scanner", repo, log, wg),
		buffer:  buffer,
		isDone:  isDone,
		cmd:     cmd,
	}
}

// scan initializes the blockScanner and starts scanning
func (bls *blockScanner) run() {
	// get the scanner range
	start, end, err := bls.scanRange()
	if err != nil {
		bls.log.Errorf("scanner can not proceed; %s", err.Error())
		return
	}

	// log what we do
	bls.log.Noticef("block scan starts at #%d", start)

	// start blockScanner
	bls.wg.Add(1)
	go bls.scan(start, end)
}

// scanRange provides the range the block scanner should run for.
func (bls *blockScanner) scanRange() (uint64, *uint64, error) {
	// get the newest known transaction
	lnb, err := bls.repo.LastKnownBlock()
	if err != nil {
		bls.log.Critical("can not scan blockchain; %s", err.Error())
		return 0, nil, err
	}

	// scan to
	var end *uint64
	if bls.cmd.BlockScanEnd > 0 {
		bls.log.Debugf("block scanner stops at #%d", bls.cmd.BlockScanEnd)
		end = &bls.cmd.BlockScanEnd
	}

	// start from
	if bls.cmd.BlockScanStart > 0 && bls.cmd.BlockScanStart <= lnb {
		bls.log.Debugf("block scanner start defined at #%d", bls.cmd.BlockScanStart)
		return bls.cmd.BlockScanStart, end, nil
	}

	// apply re-scan
	if lnb > bls.cmd.BlockScanReScan {
		bls.log.Debugf("last known block is #%d, re-scanning %d blocks", lnb, bls.cmd.BlockScanReScan)
		lnb = lnb - bls.cmd.BlockScanReScan
	}
	return lnb, end, nil
}

// logProgress will log the progress of the scanner on tics.
func (bls *blockScanner) logProgress(block *types.Block, txIndex *int, stop chan bool) {
	start := time.Now()
	tick := time.NewTicker(5 * time.Second)

	// inform we have ended
	defer func() {
		tick.Stop()
		bls.log.Infof("block scanner finished in %s", time.Now().Sub(start).String())
	}()

	for {
		select {
		case <-stop:
			return
		case <-tick.C:
			// do we have a block to display
			if block == nil {
				continue
			}

			// track the progress
			if bh, err := bls.repo.BlockHeight(); err == nil && bh != nil {
				pass := time.Now().Sub(start)
				bls.log.Infof("block #%d of #%d, trx #%d of %d; runs for %s since %s", uint64(block.Number), bh.ToInt().Uint64(), *txIndex, len(block.Txs), pass.String(), start.String())
			}
		}
	}
}

// scan performs the actual blockScanner operation on the missing blocks starting
// from the identified last known block id/number.
func (bls *blockScanner) scan(from uint64, to *uint64) {
	// what is the current block
	var (
		current = hexutil.Uint64(from)
		block   *types.Block
		txIndex int
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
	go bls.logProgress(block, &txIndex, stopLog)

	// do the scan
	for {
		// capture stop signal
		select {
		case <-bls.sigStop:
			return
		default:
		}

		// do we need a new block? if so, get next block to be scanned
		if block == nil || block.Txs == nil || len(block.Txs) <= txIndex {
			block = bls.pullBlock(&current, to, &txIndex)
			if block == nil {
				return
			}
		}

		// process next block transaction
		ev, err := bls.procNextTransaction(block, &txIndex)
		if err != nil {
			return
		}

		// anything to dispatch?
		if ev != nil {
			select {
			case <-bls.sigStop:
				return
			case bls.buffer <- ev:
			}
		}
	}
}

// pullBlock pulls the next block to be scanned, if available.
func (bls *blockScanner) pullBlock(blk *hexutil.Uint64, to *uint64, txIndex *int) *types.Block {
	// are we over the set range?
	if to != nil && uint64(*blk) > *to {
		return nil
	}

	// try to get the next block
	// if not available we assume the scanner reached the end of the chain
	block, err := bls.repo.BlockByNumber(blk)
	if err != nil {
		return nil
	}

	// add the block to the ring cache
	bls.repo.CacheBlock(block)

	// reset the current block tx index and advance to the next block
	*txIndex = 0
	*blk += 1
	return block
}

// procNextTransaction processes the next block transaction, if any.
func (bls *blockScanner) procNextTransaction(block *types.Block, txIndex *int) (*eventTransaction, error) {
	// do we have something to send?
	if block == nil || len(block.Txs) <= *txIndex {
		return nil, nil
	}

	// log action
	bls.log.Debugf("loading trx #%d of %d from block #%d", *txIndex, len(block.Txs), uint64(block.Number))

	// get transaction
	trx, err := bls.repo.LoadTransaction(block.Txs[*txIndex])
	if err != nil {
		bls.log.Criticalf("transaction not available; %s", err.Error())
		return nil, err
	}

	// update time stamp using the block data
	trx.TimeStamp = time.Unix(int64(block.TimeStamp), 0)

	// advance trx index
	*txIndex++
	return &eventTransaction{
		block: block,
		trx:   trx,
	}, err
}
