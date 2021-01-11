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
	"math/big"
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
	return &uniswapScanner{
		service: newService("uniswap scanner", repo, log, wg),
		buffer:  buffer,
		isDone:  isDone,
	}
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
		sws.log.Notice("blockchain swap scanner done")

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

	// variable for last known swap block number
	var lastBlkNr uint64

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
		blkNr := sws.processSwaps(contract, &pair, &opts)
		if lastBlkNr < blkNr {
			lastBlkNr = blkNr
		}
	}

	// store/update last processed swap block number into db
	sws.repo.UniswapUpdateLastKnownSwapBlock(lastBlkNr)
}

// getUniswapBlock returns block with given number or the actual block if the number os same
func (sws *uniswapScanner) getUniswapBlock(blockNr uint64, actualBlock *types.Block) *types.Block {

	if actualBlock != nil && blockNr == uint64(actualBlock.Number) {
		return actualBlock
	}

	blkNumber := hexutil.Uint64(blockNr)
	blk, err := sws.repo.BlockByNumber(&blkNumber)
	if err != nil {
		sws.log.Errorf("Uniswap block number %i was not found: %s", blockNr, err.Error())
		return nil
	}
	return blk
}

// processSwaps loops thru filtered swaps and adds them into the chanel for processing
func (sws *uniswapScanner) processSwaps(contract *contracts.UniswapPair, pair *common.Address, filter *bind.FilterOpts) uint64 {

	var (
		swap      *types.Swap
		lastBlkNr uint64
	)
	// get the iterator for uniswap swap event according to provided filter
	itSwap, err := contract.FilterSwap(filter, nil, nil)
	if err != nil {
		sws.log.Errorf("Cannot get swap iterator: %s", err)
		return 0
	}
	zero := new(big.Int).SetInt64(0)

	for itSwap.Next() {
		blk := sws.getUniswapBlock(itSwap.Event.Raw.BlockNumber, nil)
		sws.log.Debugf("Loading uniswap swap event from block nr# %d, tx: %s, time: %s", itSwap.Event.Raw.BlockNumber, itSwap.Event.Raw.TxHash.String(), blk.TimeStamp.String())
		swap = sws.newSwap(*blk, pair, itSwap.Event.Raw.TxHash)
		swap.Type = types.SwapNormal
		swap.Sender = itSwap.Event.To
		swap.Hash = types.BytesToHash(itSwap.Event.Raw.TxHash.Bytes())
		swap.Amount0In = itSwap.Event.Amount0In
		swap.Amount0Out = itSwap.Event.Amount0Out
		swap.Amount1In = itSwap.Event.Amount1In
		swap.Amount1Out = itSwap.Event.Amount1Out
		swap.Reserve0 = zero
		swap.Reserve1 = zero

		if lastBlkNr < uint64(blk.Number) {
			lastBlkNr = uint64(blk.Number)
		}
		if sws.isStopSignal() {
			return 0
		}
		sws.buffer <- &evtSwap{swp: swap}
	}

	// get the iterator for uniswap mint event according to provided filter
	itMint, err := contract.FilterMint(filter, nil)
	if err != nil {
		sws.log.Errorf("Cannot get uniswap mint iterator: %s", err)
		return 0
	}

	for itMint.Next() {
		blk := sws.getUniswapBlock(itMint.Event.Raw.BlockNumber, nil)
		txhash := types.BytesToHash(itMint.Event.Raw.TxHash.Bytes())
		tx, _ := sws.repo.Transaction(&txhash)
		sws.log.Debugf("Loading uniswap mint event from block nr# %d, tx: %s, time: %s", itMint.Event.Raw.BlockNumber, itMint.Event.Raw.TxHash.String(), blk.TimeStamp.String())
		swap = sws.newSwap(*blk, pair, itMint.Event.Raw.TxHash)
		swap.Type = types.SwapMint
		swap.Sender = tx.From
		swap.Hash = types.BytesToHash(itMint.Event.Raw.TxHash.Bytes())
		swap.Amount0In = itMint.Event.Amount0
		swap.Amount0Out = zero
		swap.Amount1In = itMint.Event.Amount1
		swap.Amount1Out = zero
		swap.Reserve0 = zero
		swap.Reserve1 = zero

		if lastBlkNr < uint64(blk.Number) {
			lastBlkNr = uint64(blk.Number)
		}
		if sws.isStopSignal() {
			return 0
		}
		sws.buffer <- &evtSwap{swp: swap}
	}

	// get the iterator for uniswap burn event according to provided filter
	itBurn, err := contract.FilterBurn(filter, nil, nil)
	if err != nil {
		sws.log.Errorf("Cannot get uniswap burn iterator: %s", err)
		return 0
	}

	for itBurn.Next() {
		blk := sws.getUniswapBlock(itBurn.Event.Raw.BlockNumber, nil)
		txhash := types.BytesToHash(itBurn.Event.Raw.TxHash.Bytes())
		tx, _ := sws.repo.Transaction(&txhash)
		sws.log.Debugf("Loading uniswap burn event from block nr# %d, tx: %s, time: %s", itBurn.Event.Raw.BlockNumber, itBurn.Event.Raw.TxHash.String(), blk.TimeStamp.String())
		swap = sws.newSwap(*blk, pair, itBurn.Event.Raw.TxHash)
		swap.Type = types.SwapBurn
		swap.Sender = tx.From
		swap.Hash = types.BytesToHash(itBurn.Event.Raw.TxHash.Bytes())
		swap.Amount0In = zero
		swap.Amount0Out = itBurn.Event.Amount0
		swap.Amount1In = zero
		swap.Amount1Out = itBurn.Event.Amount1
		swap.Reserve0 = zero
		swap.Reserve1 = zero

		if lastBlkNr < uint64(blk.Number) {
			lastBlkNr = uint64(blk.Number)
		}
		if sws.isStopSignal() {
			return 0
		}
		sws.buffer <- &evtSwap{swp: swap}
	}

	// get the iterator for uniswap sync event according to provided filter
	itSync, err := contract.FilterSync(filter)
	if err != nil {
		sws.log.Errorf("Cannot get uniswap sync iterator: %s", err)
		return 0
	}

	for itSync.Next() {
		blk := sws.getUniswapBlock(itSync.Event.Raw.BlockNumber, nil)
		sws.log.Debugf("Loading uniswap sync event from block nr# %d, tx: %s, time: %s", itSync.Event.Raw.BlockNumber, itSync.Event.Raw.TxHash.String(), blk.TimeStamp.String())
		swap = sws.newSwap(*blk, pair, itSync.Event.Raw.TxHash)
		swap.Type = types.SwapSync
		swap.Hash = types.BytesToHash(itSync.Event.Raw.TxHash.Bytes())
		swap.Amount0In = zero
		swap.Amount0Out = zero
		swap.Amount1In = zero
		swap.Amount1Out = zero
		swap.Reserve0 = itSync.Event.Reserve0
		swap.Reserve1 = itSync.Event.Reserve1

		if lastBlkNr < uint64(blk.Number) {
			lastBlkNr = uint64(blk.Number)
		}
		if sws.isStopSignal() {
			return 0
		}
		sws.buffer <- &evtSwap{swp: swap}
	}
	return lastBlkNr
}

func (sws *uniswapScanner) isStopSignal() bool {
	select {
	case <-sws.sigStop:
		// stop signal received?
		return true
	default:
	}
	return false
}

// newSwap creates a new swap instance with predefined data
func (sws *uniswapScanner) newSwap(block types.Block, pair *common.Address, txHash common.Hash) *types.Swap {
	return &types.Swap{
		OrdIndex:    sws.getOrdinalIndex(&block, &txHash),
		BlockNumber: &block.Number,
		TimeStamp:   &block.TimeStamp,
		Pair:        *pair,
	}
}

// getOrdinalIndex resolves ordinal index for block and transaction input
func (sws *uniswapScanner) getOrdinalIndex(blk *types.Block, txHash *common.Hash) uint64 {
	// get transaction
	txhash := types.BytesToHash(txHash.Bytes())
	tx, err := sws.repo.Transaction(&txhash)
	if err != nil {
		sws.log.Errorf("Transaction was not found for block %s and txHash %s; %s", blk.Number.String(), txhash.String(), err.Error())
		return 0
	}
	return types.TransactionIndex(blk, tx)
}
