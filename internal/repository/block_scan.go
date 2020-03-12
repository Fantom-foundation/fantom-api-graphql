/*
Repository package implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Opera/Lachesis full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import (
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"sync"
)

// scanBufferCapacity is the number of transaction kept in buffer before scanner gets throttled.
const scanBufferCapacity = 250

// scanTransaction represents a single scanned transaction record to be processed.
type scanTransaction struct {
	block *types.Block
	trx   *types.Transaction
}

// ScanChain performs blockchain scan and stores found data in off-chain storage.
// It returns a channel which can be used to singal the scanner to terminate prematurely.
func (p *proxy) ScanBlockChain(wg *sync.WaitGroup) chan bool {
	// get the newest known transaction
	lnb, err := p.db.LastKnownBlock()
	if err != nil {
		p.log.Critical("can not scan the block chain")
		return nil
	}

	// log what we do
	p.log.Infof("starting blockchain scan from block %d", lnb)

	// make the termination signaling channel and add self to the wait group
	// no one may be listening, if scanner already finished its job, so use buffered one
	sigStop := make(chan bool, 1)

	// make the scan results channel
	buffer := make(chan *scanTransaction, scanBufferCapacity)

	// start the digger, use scan buffer to send digger blocks to the receiver
	wg.Add(1)
	go digTransactions(lnb, sigStop, buffer, wg, p)

	// do the processing
	wg.Add(1)
	go dispatchTransactions(buffer, wg, p)

	// return the stop signal hook
	return sigStop
}

// digTransactions will extract transactions from blockchain and send them to the output channel for processing.
func digTransactions(start uint64, sigStop chan bool, buffer chan *scanTransaction, wg *sync.WaitGroup, repo *proxy) {
	// don't forget to sign off after we are done
	defer func() {
		// log finish
		repo.log.Info("blockchain scanner finished")

		// we write to the channel so we close it before leaving
		close(buffer)
		wg.Done()
	}()

	// what is the current block
	var (
		block     *types.Block
		trx       *types.Transaction
		toSend    *scanTransaction
		err       error
		index     int
		blkNumber = hexutil.Uint64(start)
	)

	// do the scan
	for {
		// do we need a new block?
		if block == nil || block.Txs == nil || len(block.Txs) <= index {
			// log action
			repo.log.Debugf("scanner reached block %d", uint64(blkNumber))

			// try to get the next block
			block, err = repo.BlockByNumber(&blkNumber)
			if err != nil {
				return
			}

			// reset the current block tx index and advance to the next block
			index = 0
			blkNumber += 1
		}

		// do we have something to send?
		if toSend != nil || len(block.Txs) > index {
			// do we need to pull next transaction?
			if toSend == nil {
				// log action
				repo.log.Debugf("loading transaction %d of block %d", index, uint64(block.Number))

				// get transaction
				trx, err = repo.Transaction(block.Txs[index])
				if err != nil {
					return
				}

				// prep sending struct and advance transaction index
				toSend = &scanTransaction{block: block, trx: trx}
				index++
			}

			// try to send
			select {
			case <-sigStop:
				// stop signal received?
				return
			case buffer <- toSend:
				// we did send it and now we need next one
				toSend = nil
			default:
			}
		}
	}
}

// dispatchTransactions will receive scanned transactions and send them into the off-chain database.
func dispatchTransactions(buffer chan *scanTransaction, wg *sync.WaitGroup, repo *proxy) {
	// don't forget to sign off after we are done
	defer func() {
		// log finish
		repo.log.Info("blockchain scanner tx dispatch finished")
		wg.Done()
	}()

	// wait for transactions and process them
	for {
		// try to read next transaction
		toDispatch, ok := <-buffer
		if ok == false {
			// we are done, no more data will come
			return
		}

		// validate
		if toDispatch.block == nil || toDispatch.trx == nil {
			repo.log.Critical("invalid tx received from the scanner")
		}

		// dispatch the received
		err := repo.db.AddTransaction(toDispatch.block, toDispatch.trx)
		if err != nil {
			repo.log.Error("could not dispatch transaction")
			repo.log.Error(err)
		}
	}
}
