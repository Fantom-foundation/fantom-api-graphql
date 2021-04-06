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
	"fantom-api-graphql/internal/repository/rpc/contracts"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/event"
	ftm "github.com/ethereum/go-ethereum/rpc"
	"math/big"
	"strings"
	"sync"
)

// uniswapPairEventQueueSize represents the size of the uniswap pair events queue.
const uniswapPairEventQueueSize = 50

// UniswapPairMonitor represents a monitor capturing events on a given Uniswap pair
type UniswapPairMonitor struct {
	service

	pair     common.Address
	swapChan chan *types.Swap

	// channels capturing events from the subscription on the pair
	swapEventCh chan *contracts.UniswapPairSwap
	mintEventCh chan *contracts.UniswapPairMint
	burnEventCh chan *contracts.UniswapPairBurn
	syncEventCh chan *contracts.UniswapPairSync
	subs        []event.Subscription

	con *ftm.Client
}

// NewUniswapPairMonitor creates a new uniswap monitor instance.
func NewUniswapPairMonitor(
	con *ftm.Client,
	buffer chan *types.Swap,
	repo Repository,
	log logger.Logger,
	wg *sync.WaitGroup,
	pair *common.Address,
) *UniswapPairMonitor {
	var sb strings.Builder
	sb.WriteString("uniswap pair ")
	sb.WriteString(pair.String())
	sb.WriteString(" monitor")

	// create new monitor instance
	pam := UniswapPairMonitor{
		service:  newService(sb.String(), repo, log, wg),
		swapChan: buffer,
		con:      con,
		pair:     *pair,
	}

	// return pair monitor instance
	return &pam
}

// run starts the uniswap pair monitoring.
func (pam *UniswapPairMonitor) run() {
	// inform
	pam.log.Debugf("initiating uniswap pair %s monitor", pam.pair.String())

	// get contract instance for obtaining log events
	contract, err := pam.repo.UniswapPairContract(&pam.pair)
	if err != nil {
		pam.log.Errorf("uniswap contract for pair %s not found; %s", pam.pair.String(), err.Error())
		return
	}

	// creating swap, mint, burn, sync event channels
	pam.swapEventCh = make(chan *contracts.UniswapPairSwap, uniswapPairEventQueueSize)
	pam.mintEventCh = make(chan *contracts.UniswapPairMint, uniswapPairEventQueueSize)
	pam.burnEventCh = make(chan *contracts.UniswapPairBurn, uniswapPairEventQueueSize)
	pam.syncEventCh = make(chan *contracts.UniswapPairSync, uniswapPairEventQueueSize)

	// open subscriptions
	if err = pam.subscribe(contract); err != nil {
		pam.log.Errorf("failed to monitor uniswap pair %s; %s", pam.pair.String(), err.Error())
		return
	}

	// start the monitor
	pam.wg.Add(1)
	go pam.monitor()
}

// subscribe opens subscriptions to all the monitored events on the pair.
func (pam *UniswapPairMonitor) subscribe(pair *contracts.UniswapPair) (err error) {
	// inform
	pam.log.Debugf("opening subscriptions for uniswap pair %s", pam.pair.String())

	// collect subs so we can unsubscribe when done
	pam.subs = make([]event.Subscription, 4)

	// make subscription to Uniswap swap event
	pam.subs[0], err = pair.UniswapPairFilterer.WatchSwap(&bind.WatchOpts{}, pam.swapEventCh, nil, nil)
	if err != nil {
		pam.log.Errorf("can not subscribe for swap action on pair %s; %s", pam.pair.String(), err.Error())
		return err
	}

	// make subscription to Uniswap mint event
	pam.subs[1], err = pair.UniswapPairFilterer.WatchMint(&bind.WatchOpts{}, pam.mintEventCh, nil)
	if err != nil {
		pam.log.Errorf("can not subscribe for uniswap mint event on pair %s; %s", pam.pair.String(), err.Error())
		return err
	}

	// make subscription to Uniswap burn event
	pam.subs[2], err = pair.UniswapPairFilterer.WatchBurn(&bind.WatchOpts{}, pam.burnEventCh, nil, nil)
	if err != nil {
		pam.log.Errorf("can not subscribe for uniswap burn event on pair %s; %s", pam.pair.String(), err.Error())
		return err
	}

	// make subscription to Uniswap sync event
	pam.subs[3], err = pair.UniswapPairFilterer.WatchSync(&bind.WatchOpts{}, pam.syncEventCh)
	if err != nil {
		pam.log.Errorf("can not subscribe for uniswap sync event on pair %s; %s", pam.pair.String(), err.Error())
		return err
	}

	return nil
}

// monitor does the actual monitoring and processing job for monitored uniswap events.
func (pam *UniswapPairMonitor) monitor() {
	// don't forget to sign off after we are done
	defer func() {
		// unsubscribe from events
		pam.log.Debugf("closing subscriptions for uniswap pair %s", pam.pair.String())
		for _, sub := range pam.subs {
			sub.Unsubscribe()
		}

		// signal to wait group we are done
		pam.log.Notice("uniswap pair %s monitor is closed", pam.pair.String())
		pam.wg.Done()
	}()

	// inform about the active state
	pam.log.Debugf("uniswap pair %s is being monitored", pam.pair.String())

	// loop until terminated
	for {
		select {
		case <-pam.sigStop:
			// we have been terminated
			return
		case swap := <-pam.swapEventCh:
			// new swap have been detected
			swapData, err := pam.getSwapData(swap)
			if err != nil {
				pam.log.Criticalf("can not extract data from swap event. %s", err.Error())
				continue
			}

			// send the event for processing
			pam.swapChan <- swapData
			pam.log.Debugf("sending Uniswap swap for dispatch Tx: %s", swapData.Hash.String())

		case mint := <-pam.mintEventCh:
			// new swap have been detected
			swapData, err := pam.getMintData(mint)
			if err != nil {
				pam.log.Criticalf("can not extract data from mint event. %s", err.Error())
				continue
			}

			// send the event for processing
			pam.swapChan <- swapData
			pam.log.Debugf("sending Uniswap mint for dispatch Tx: %s", swapData.Hash.String())

		case burn := <-pam.burnEventCh:
			// new swap have been detected
			swapData, err := pam.getBurnData(burn)
			if err != nil {
				pam.log.Criticalf("can not extract data from burn event. %s", err.Error())
				continue
			}

			// send the event for processing
			pam.swapChan <- swapData
			pam.log.Debugf("sending Uniswap burn for dispatch Tx: %s", swapData.Hash.String())

		case snc := <-pam.syncEventCh:
			// new swap have been detected
			swapData, err := pam.getSyncData(snc)
			if err != nil {
				pam.log.Criticalf("can not extract data from sync event. %s", err.Error())
				continue
			}

			// send the event for processing
			pam.swapChan <- swapData
			pam.log.Debugf("sending Uniswap sync for dispatch Tx: %s", swapData.Hash.String())
		}
	}
}

// getOrdinalIndex resolves ordinal index for block and transaction input
func (pam *UniswapPairMonitor) getOrdinalIndex(blk *types.Block, txHash *common.Hash) uint64 {
	// get transaction
	th := common.BytesToHash(txHash.Bytes())
	tx, err := pam.repo.Transaction(&th)
	if err != nil {
		pam.log.Errorf("transaction not found for block %s and hash %s; %s", blk.Number.String(), txHash.String(), err.Error())
		return 0
	}
	return types.TransactionIndex(blk, tx)
}

// getSwapData loops thru filtered swaps and adds them into the channel for processing
func (pam *UniswapPairMonitor) getSwapData(swapEvent *contracts.UniswapPairSwap) (*types.Swap, error) {
	zero := new(big.Int).SetUint64(0)
	blkNumber := hexutil.Uint64(swapEvent.Raw.BlockNumber)

	// get the block information
	blk, err := pam.repo.BlockByNumber(&blkNumber)
	if err != nil {
		pam.log.Errorf("block was not found for swap on pair %s; %s", pam.pair.String(), err.Error())
		return nil, err
	}

	// log action
	pam.log.Debugf("loading swap from block nr# %d, tx: %s", swapEvent.Raw.BlockNumber, swapEvent.Raw.TxHash.String())

	// prep sending struct and advance transaction index
	swap := &types.Swap{
		OrdIndex:    pam.getOrdinalIndex(blk, &swapEvent.Raw.TxHash),
		Type:        types.SwapNormal,
		BlockNumber: &blk.Number,
		TimeStamp:   &blk.TimeStamp,
		Pair:        pam.pair,
		Sender:      swapEvent.To,
		Hash:        common.BytesToHash(swapEvent.Raw.TxHash.Bytes()),
		Amount0In:   swapEvent.Amount0In,
		Amount0Out:  swapEvent.Amount0Out,
		Amount1In:   swapEvent.Amount1In,
		Amount1Out:  swapEvent.Amount1Out,
		Reserve0:    zero,
		Reserve1:    zero,
	}
	return swap, nil
}

// getMintData loops thru filtered mint event data and adds them into the channel for processing
func (pam *UniswapPairMonitor) getMintData(mintEvent *contracts.UniswapPairMint) (*types.Swap, error) {
	zero := new(big.Int).SetUint64(0)
	blkNumber := hexutil.Uint64(mintEvent.Raw.BlockNumber)

	blk, err := pam.repo.BlockByNumber(&blkNumber)
	if err != nil {
		pam.log.Errorf("Block was not found for uniswap mint event on pair %s; %s", pam.pair.String(), err.Error())
		return nil, err
	}

	// get actual transaction
	th := common.BytesToHash(mintEvent.Raw.TxHash.Bytes())
	tx, err := pam.repo.Transaction(&th)
	if err != nil {
		pam.log.Errorf("Transaction was not found for uniswap mint event on pair %s; %s", pam.pair.String(), err.Error())
		return nil, err
	}

	// log action
	pam.log.Debugf("loading mint event data from block nr# %d, tx: %s", mintEvent.Raw.BlockNumber, mintEvent.Raw.TxHash.String())

	// prep sending struct
	swap := &types.Swap{
		OrdIndex:    pam.getOrdinalIndex(blk, &mintEvent.Raw.TxHash),
		Type:        types.SwapMint,
		BlockNumber: &blk.Number,
		TimeStamp:   &blk.TimeStamp,
		Pair:        pam.pair,
		Sender:      tx.From,
		Hash:        common.BytesToHash(mintEvent.Raw.TxHash.Bytes()),
		Amount0In:   mintEvent.Amount0,
		Amount0Out:  zero,
		Amount1In:   mintEvent.Amount1,
		Amount1Out:  zero,
		Reserve0:    zero,
		Reserve1:    zero,
	}
	return swap, nil
}

// getBurnData loops thru filtered burn event data and adds them into the channel for processing
func (pam *UniswapPairMonitor) getBurnData(burnEvent *contracts.UniswapPairBurn) (*types.Swap, error) {
	zero := new(big.Int).SetUint64(0)
	blkNumber := hexutil.Uint64(burnEvent.Raw.BlockNumber)

	blk, err := pam.repo.BlockByNumber(&blkNumber)
	if err != nil {
		pam.log.Errorf("block was not found for uniswap burn event on pair %s; %s", pam.pair.String(), err.Error())
		return nil, err
	}

	// log action
	pam.log.Debugf("loading burn event data from block nr# %d, tx: %s", burnEvent.Raw.BlockNumber, burnEvent.Raw.TxHash.String())

	// get actual transaction
	th := common.BytesToHash(burnEvent.Raw.TxHash.Bytes())
	tx, err := pam.repo.Transaction(&th)
	if err != nil {
		pam.log.Errorf("Transaction was not found for uniswap burn event on pair %s; %s", pam.pair.String(), err.Error())
		return nil, err
	}

	// prep sending struct
	swap := &types.Swap{
		OrdIndex:    pam.getOrdinalIndex(blk, &burnEvent.Raw.TxHash),
		Type:        types.SwapBurn,
		BlockNumber: &blk.Number,
		TimeStamp:   &blk.TimeStamp,
		Pair:        pam.pair,
		Sender:      tx.From,
		Hash:        common.BytesToHash(burnEvent.Raw.TxHash.Bytes()),
		Amount0In:   zero,
		Amount0Out:  burnEvent.Amount0,
		Amount1In:   zero,
		Amount1Out:  burnEvent.Amount1,
		Reserve0:    zero,
		Reserve1:    zero,
	}
	return swap, nil
}

// getSyncData loops thru filtered uniswap sync event data and adds them into the chanel for processing
func (pam *UniswapPairMonitor) getSyncData(syncEvent *contracts.UniswapPairSync) (*types.Swap, error) {
	zero := new(big.Int).SetUint64(0)
	blkNumber := hexutil.Uint64(syncEvent.Raw.BlockNumber)

	blk, err := pam.repo.BlockByNumber(&blkNumber)
	if err != nil {
		pam.log.Errorf("Block was not found for uniswap sync event on pair %s; %s", pam.pair.String(), err.Error())
		return nil, err
	}

	// log action
	pam.log.Debugf("Loading uniswap sync event data from block nr# %d, tx: %s, reserve0: %s, reserve1: %s",
		syncEvent.Raw.BlockNumber, syncEvent.Raw.TxHash.String(),
		syncEvent.Reserve0.String(), syncEvent.Reserve1.String())

	// prep sending struct
	swap := &types.Swap{
		OrdIndex:    pam.getOrdinalIndex(blk, &syncEvent.Raw.TxHash),
		Type:        types.SwapSync,
		BlockNumber: &blk.Number,
		TimeStamp:   &blk.TimeStamp,
		Pair:        pam.pair,
		Hash:        common.BytesToHash(syncEvent.Raw.TxHash.Bytes()),
		Amount0In:   zero,
		Amount0Out:  zero,
		Amount1In:   zero,
		Amount1Out:  zero,
		Reserve0:    syncEvent.Reserve0,
		Reserve1:    syncEvent.Reserve1,
	}
	return swap, nil
}
