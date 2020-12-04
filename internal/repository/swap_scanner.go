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
	"fantom-api-graphql/internal/repository/rpc/contracts"
	"fantom-api-graphql/internal/types"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// uniswapScanner implements blockchain scanner used to extract blockchain data to off-chain storage.
type uniswapScanner struct {
	service
	buffer chan *evtSwap
	isDone chan bool
}

// newUniswapScanner creates new blockchain scanner service.
func newUniswapScanner(buffer chan *evtSwap, isDone chan bool, repo Repository, log logger.Logger, wg *sync.WaitGroup) *uniswapScanner {
	// create new scanner instance
	sc := uniswapScanner{
		service: newService("uniswap scanner", repo, log, wg),
		buffer:  buffer,
		isDone:  isDone,
	}

	// start the scanner job
	sc.run()
	return &sc
}

// scan initializes the scanner and starts scanning
func (sws *uniswapScanner) run() {

	// get the newest known swap block number
	lnb, err := sws.repo.LastKnownSwapBlock()
	if err != nil {
		sws.log.Critical("can not scan blockchain; %s", err.Error())
		return
	}

	// log what we do
	sws.log.Noticef("blockchain swap scan starts from block #%d", lnb)

	// start scanner
	sws.wg.Add(1)
	go sws.scan(lnb)
}

// scan performs the actual scanner operation on the missing blocks starting
// from the identified last known block id/number.
func (sws *uniswapScanner) scan(lnb uint64) {
	// don't forget to sign off after we are done
	defer func() {
		// signal we are done with the sync
		sws.isDone <- true

		// log finish
		sws.log.Notice("blockchain scanner done")

		// signal to wait group we are done
		sws.wg.Done()
	}()

	//get all pairs in blockchain
	pairs, err := sws.repo.UniswapPairs()
	if err != nil {
		sws.log.Errorf("Uniswap pairs can not be resolved; %s", err.Error())
		return
	}

	// skip last block if there is any
	if lnb > 0 {
		lnb++
	}

	// create range for the scan from last known swap block to the latest one
	opts := bind.FilterOpts{
		Start:   lnb,
		End:     nil,
		Context: context.Background(),
	}

	// get filter for all pairs
	for _, pair := range pairs {

		sws.log.Debugf("starting swap scan for pair %s", pair.String())

		// get contract instance for obtaining log events
		contract, err := sws.repo.UniswapPairContract(&pair)
		if err != nil {
			sws.log.Errorf("Uniswap pair %s not found; %s", pair.String(), err.Error())
			return
		}

		// process swaps
		processSwaps(sws, contract, &pair, &opts)

		// here will proccess also Burn, Mint events
	}
}

// processSwaps loops thru filtered swaps and adds them into the chanel for processing
func processSwaps(sws *uniswapScanner, contract *contracts.UniswapPair, pair *common.Address, filter *bind.FilterOpts) {

	var (
		toSend *evtSwap
		swap   *types.Swap
	)

	// get the iterator if all swaps according to provided filter
	it, err := contract.FilterSwap(filter, nil, nil)
	if err != nil {
		sws.log.Errorf("Cannot get swap iterator: %s", err)
		return
	}

	// loop thru all swaps

	for {

		if toSend == nil {
			if it.Next() {
				blkNumber := hexutil.Uint64(it.Event.Raw.BlockNumber)

				blk, err := sws.repo.BlockByNumber(&blkNumber)
				if err != nil {
					sws.log.Errorf("Block was not found for pair %s; %s", pair.String(), err.Error())
					return
				}

				// log action
				sws.log.Debugf("Loading swap from block nr# %d, tx: %s", it.Event.Raw.BlockNumber, it.Event.Raw.TxHash.String())

				// prep sending struct and advance transaction index
				swap = &types.Swap{
					BlockNumber: &blk.Number,
					TimeStamp:   &blk.TimeStamp,
					Pair:        *pair,
					Sender:      it.Event.Sender,
					To:          it.Event.To,
					Hash:        types.BytesToHash(it.Event.Raw.TxHash.Bytes()),
					Amount0In:   it.Event.Amount0In,
					Amount0Out:  it.Event.Amount0Out,
					Amount1In:   it.Event.Amount1In,
					Amount1Out:  it.Event.Amount1Out,
				}
				toSend = &evtSwap{swp: swap}
			} else {
				return
			}
		}

		// try to send
		select {
		case <-sws.sigStop:
			// stop signal received?
			return
		case sws.buffer <- toSend:
			// we did send it and now we need next one
			toSend = nil
		default:
		}
	}

}
