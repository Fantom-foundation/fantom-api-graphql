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
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/event"
	ftm "github.com/ethereum/go-ethereum/rpc"
)

// UniswapMonitor represents a processor capturing new swap events on the blockchain
type UniswapMonitor struct {
	service
	subs         []event.Subscription
	swapEventCh  chan *contracts.UniswapPairSwap
	mintEventCh  chan *contracts.UniswapPairMint
	burnEventCh  chan *contracts.UniswapPairBurn
	syncEventCh  chan *contracts.UniswapPairSync
	newPairEvtCh chan *contracts.UniswapFactoryPairCreated
	swapChan     chan *evtSwap
	con          *ftm.Client
	sub          *ftm.ClientSubscription
}

// NewUniswapMonitor creates a new uniswap monitor instance.
func NewUniswapMonitor(con *ftm.Client, buffer chan *evtSwap, repo Repository, log logger.Logger, wg *sync.WaitGroup) *UniswapMonitor {
	// create new monitor instance
	mo := UniswapMonitor{
		service:  newService("uniswap monitor", repo, log, wg),
		swapChan: buffer,
		con:      con,
	}

	// return monitor instance
	return &mo
}

// run starts monitoring for uniswap swap events
func (um *UniswapMonitor) run() {
	// log
	um.log.Notice("initializing blockchain monitor")

	//get all pairs in blockchain
	pairs, err := um.repo.UniswapPairs()
	if err != nil {
		um.log.Errorf("Uniswap pairs can not be resolved; %s", err.Error())
		return
	}

	// start monitoring for all pairs
	for _, pair := range pairs {
		um.startPairMonitor(&pair)
	}

	// start monitoring for new pairs created
	go um.monitorUniswapFactoryPairCreation()

	// start uniswap monitoring routine to be able to close everything properly on the end
	um.wg.Add(1)
	go um.monitor()

	um.log.Notice("uniswap swap event monitor started")
}

// monitor here just checks for ending signal for propper close of all the other channels and routines
func (um *UniswapMonitor) monitor() {
	// don't forget to sign off after we are done
	defer func() {

		//unsubscribe all swap events
		for _, subscription := range um.subs {
			subscription.Unsubscribe()
		}

		//close all swap event channels
		close(um.swapEventCh)
		close(um.mintEventCh)
		close(um.burnEventCh)
		close(um.syncEventCh)

		//close channel for new pair creation
		close(um.newPairEvtCh)

		// signal to wait group we are done
		um.wg.Done()

		//log
		um.log.Notice("uniswap swap event monitor stopped")
	}()

	// loop here
	for {
		select {
		case <-um.sigStop:
			return
		}
	}
}

// monitorUniswapFactoryPairCreation is goroutine for checking if new pair was created on the blockchain
func (um *UniswapMonitor) monitorUniswapFactoryPairCreation() {
	defer func() {
		um.log.Noticef("Stopping Uniswap factory monitor for new pairs.")
	}()

	// gets the Uniswap factory instance according to config address of the contract
	uniswapFactory, err := um.repo.UniswapFactoryContract()
	if err != nil {
		um.log.Errorf("Uniswap Factory Contract instance was not created; %s", err.Error())
		return
	}

	// create channel for capturing event for new pair creation
	um.newPairEvtCh = make(chan *contracts.UniswapFactoryPairCreated)
	newPairEvtSubscription, err := uniswapFactory.WatchPairCreated(&bind.WatchOpts{}, um.newPairEvtCh, nil, nil)
	if err != nil {
		um.log.Errorf("Cannot subscribe to Uniswap Factory Contract for tracking new pair creation; %s", err.Error())
		return
	}

	// add subscription to array to be able to close it in the end
	um.subs = append(um.subs, newPairEvtSubscription)

	um.log.Noticef("Starting Uniswap factory monitor for new pairs.")

	// listen on subscribed channel for newly created pairs and create monitor for them afterwards
	for val := range um.newPairEvtCh {
		um.startPairMonitor(&val.Pair)
	}
}

// startPairMonitor creates goroutine for given pair address to monitor swaps of this pair
func (um *UniswapMonitor) startPairMonitor(pair *common.Address) {

	// get contract instance for obtaining log events
	contract, err := um.repo.UniswapPairContract(pair)
	if err != nil {
		um.log.Errorf("Uniswap contract for pair %s not found; %s", pair.String(), err.Error())
		return
	}
	// creating swap, mint, burn, sync event channels
	swapEventCh := make(chan *contracts.UniswapPairSwap)
	mintEventCh := make(chan *contracts.UniswapPairMint)
	burnEventCh := make(chan *contracts.UniswapPairBurn)
	syncEventCh := make(chan *contracts.UniswapPairSync)

	// make subscription to Uniswap swap event
	evtSwap, err := contract.UniswapPairFilterer.WatchSwap(&bind.WatchOpts{}, swapEventCh, nil, nil)
	if err != nil {
		um.log.Errorf("Cannot subscribe for swap action on pair %s; %s", pair.String(), err.Error())
		return
	}
	// make subscription to Uniswap mint event
	evtMint, err := contract.UniswapPairFilterer.WatchMint(&bind.WatchOpts{}, mintEventCh, nil)
	if err != nil {
		um.log.Errorf("Cannot subscribe for uniswap mint event on pair %s; %s", pair.String(), err.Error())
		return
	}
	// make subscription to Uniswap burn event
	evtBurn, err := contract.UniswapPairFilterer.WatchBurn(&bind.WatchOpts{}, burnEventCh, nil, nil)
	if err != nil {
		um.log.Errorf("Cannot subscribe for uniswap burn event on pair %s; %s", pair.String(), err.Error())
		return
	}
	// make subscription to Uniswap sync event
	evtSync, err := contract.UniswapPairFilterer.WatchSync(&bind.WatchOpts{}, syncEventCh)
	if err != nil {
		um.log.Errorf("Cannot subscribe for uniswap sync event on pair %s; %s", pair.String(), err.Error())
		return
	}

	//add subscriptions to array to be able to unsubscribe them in the end
	um.subs = append(um.subs, evtMint)
	um.subs = append(um.subs, evtBurn)
	um.subs = append(um.subs, evtSync)
	um.subs = append(um.subs, evtSwap)

	// start routine for handling swaps for given pair
	go um.monitorPair(*pair, um.swapChan, swapEventCh)
	go um.monitorPairMint(*pair, um.swapChan, mintEventCh)
	go um.monitorPairBurn(*pair, um.swapChan, burnEventCh)
	go um.monitorPairSync(*pair, um.swapChan, syncEventCh)

}

// monitorPair uses channel to swap dispatcher to send data from swap event there, for storing them into persistent db
func (um *UniswapMonitor) monitorPair(pair common.Address, swapDispatchChan chan *evtSwap, swapEventCh chan *contracts.UniswapPairSwap) {

	defer func() {
		um.log.Noticef("Stopping Uniswap swap monitor for pair: %s", pair.String())
	}()

	um.log.Noticef("Starting Uniswap swap monitor for pair: %s", pair.String())

	for val := range swapEventCh {
		evtData, err := um.getSwapData(val, pair)
		if err != nil {
			um.log.Errorf("Cannot extract data from swap event. %s", err.Error())
			return
		}
		swapDispatchChan <- evtData
		um.log.Debugf("Sending Uniswap swap for dispatch Tx: %s", evtData.swp.Hash.String())
	}
}

// getSwapData loops thru filtered swaps and adds them into the chanel for processing
func (um *UniswapMonitor) getSwapData(swapEvent *contracts.UniswapPairSwap, pair common.Address) (*evtSwap, error) {

	zero := new(big.Int).SetUint64(0)
	blkNumber := hexutil.Uint64(swapEvent.Raw.BlockNumber)

	blk, err := um.repo.BlockByNumber(&blkNumber)
	if err != nil {
		um.log.Errorf("Block was not found for swap on pair %s; %s", pair.String(), err.Error())
		return nil, err
	}

	// log action
	um.log.Debugf("Loading swap from block nr# %d, tx: %s", swapEvent.Raw.BlockNumber, swapEvent.Raw.TxHash.String())

	// prep sending struct and advance transaction index
	swap := &types.Swap{
		Type:        types.SwapNormal,
		BlockNumber: &blk.Number,
		TimeStamp:   &blk.TimeStamp,
		Pair:        pair,
		Sender:      swapEvent.To,
		Hash:        types.BytesToHash(swapEvent.Raw.TxHash.Bytes()),
		Amount0In:   swapEvent.Amount0In,
		Amount0Out:  swapEvent.Amount0Out,
		Amount1In:   swapEvent.Amount1In,
		Amount1Out:  swapEvent.Amount1Out,
		Reserve0:    zero,
		Reserve1:    zero,
	}
	return &evtSwap{swp: swap}, nil
}

// monitorPairMint uses channel to swap dispatcher to send data from uniswap mint event there, for storing them into persistent db
func (um *UniswapMonitor) monitorPairMint(pair common.Address, swapDispatchChan chan *evtSwap, mintEventCh chan *contracts.UniswapPairMint) {

	defer func() {
		um.log.Noticef("Stopping Uniswap mint monitor for pair: %s", pair.String())
	}()

	um.log.Noticef("Starting Uniswap mint monitor for pair: %s", pair.String())

	for val := range mintEventCh {
		evtData, err := um.getMintData(val, pair)
		if err != nil {
			um.log.Errorf("Cannot extract data from uniswap mint event. %s", err.Error())
			return
		}
		swapDispatchChan <- evtData
		um.log.Debugf("Sending Uniswap mint for dispatch Tx: %s", evtData.swp.Hash.String())
	}
}

// getMintData loops thru filtered mint event data and adds them into the chanel for processing
func (um *UniswapMonitor) getMintData(mintEvent *contracts.UniswapPairMint, pair common.Address) (*evtSwap, error) {

	zero := new(big.Int).SetUint64(0)
	blkNumber := hexutil.Uint64(mintEvent.Raw.BlockNumber)

	blk, err := um.repo.BlockByNumber(&blkNumber)
	if err != nil {
		um.log.Errorf("Block was not found for uniswap mint event on pair %s; %s", pair.String(), err.Error())
		return nil, err
	}

	// get actual transaction
	txhash := types.BytesToHash(mintEvent.Raw.TxHash.Bytes())
	tx, _ := um.repo.Transaction(&txhash)

	// log action
	um.log.Debugf("Loading mint event data from block nr# %d, tx: %s", mintEvent.Raw.BlockNumber, mintEvent.Raw.TxHash.String())

	// prep sending struct
	swap := &types.Swap{
		Type:        types.SwapMint,
		BlockNumber: &blk.Number,
		TimeStamp:   &blk.TimeStamp,
		Pair:        pair,
		Sender:      tx.From,
		Hash:        types.BytesToHash(mintEvent.Raw.TxHash.Bytes()),
		Amount0In:   mintEvent.Amount0,
		Amount0Out:  zero,
		Amount1In:   mintEvent.Amount1,
		Amount1Out:  zero,
		Reserve0:    zero,
		Reserve1:    zero,
	}
	return &evtSwap{swp: swap}, nil
}

// monitorPairBurn uses channel to swap dispatcher to send data from uniswap burn event there, for storing them into persistent db
func (um *UniswapMonitor) monitorPairBurn(pair common.Address, swapDispatchChan chan *evtSwap, transferEventCh chan *contracts.UniswapPairBurn) {

	defer func() {
		um.log.Noticef("Stopping Uniswap burn monitor for pair: %s", pair.String())
	}()

	um.log.Noticef("Starting Uniswap burn monitor for pair: %s", pair.String())

	for val := range transferEventCh {
		evtData, err := um.getBurnData(val, pair)
		if err != nil {
			um.log.Errorf("Cannot extract data from uniswap burn event. %s", err.Error())
			return
		}
		swapDispatchChan <- evtData
		um.log.Debugf("Sending Uniswap burn for dispatch Tx: %s", evtData.swp.Hash.String())
	}
}

// getBurnData loops thru filtered burn event data and adds them into the chanel for processing
func (um *UniswapMonitor) getBurnData(burnEvent *contracts.UniswapPairBurn, pair common.Address) (*evtSwap, error) {

	blkNumber := hexutil.Uint64(burnEvent.Raw.BlockNumber)

	blk, err := um.repo.BlockByNumber(&blkNumber)
	if err != nil {
		um.log.Errorf("Block was not found for uniswap burn event on pair %s; %s", pair.String(), err.Error())
		return nil, err
	}

	// log action
	um.log.Debugf("Loading burn event data from block nr# %d, tx: %s", burnEvent.Raw.BlockNumber, burnEvent.Raw.TxHash.String())

	// get actual transaction
	txhash := types.BytesToHash(burnEvent.Raw.TxHash.Bytes())
	tx, _ := um.repo.Transaction(&txhash)
	zero := new(big.Int).SetUint64(0)

	// prep sending struct
	swap := &types.Swap{
		Type:        types.SwapBurn,
		BlockNumber: &blk.Number,
		TimeStamp:   &blk.TimeStamp,
		Pair:        pair,
		Sender:      tx.From,
		Hash:        types.BytesToHash(burnEvent.Raw.TxHash.Bytes()),
		Amount0In:   zero,
		Amount0Out:  burnEvent.Amount0,
		Amount1In:   zero,
		Amount1Out:  burnEvent.Amount1,
		Reserve0:    zero,
		Reserve1:    zero,
	}
	return &evtSwap{swp: swap}, nil
}

// monitorPairSync uses channel to swap dispatcher to send data from uniswap sync event there, for storing them into persistent db
func (um *UniswapMonitor) monitorPairSync(pair common.Address, swapDispatchChan chan *evtSwap, syncEventCh chan *contracts.UniswapPairSync) {

	defer func() {
		um.log.Noticef("Stopping Uniswap sync monitor for pair: %s", pair.String())
	}()

	um.log.Noticef("Starting Uniswap sync monitor for pair: %s", pair.String())

	for val := range syncEventCh {
		evtData, err := um.getSyncData(val, pair)
		if err != nil {
			um.log.Errorf("Cannot extract data from uniswap sync event. %s", err.Error())
			return
		}
		swapDispatchChan <- evtData
		um.log.Debugf("Sending Uniswap sync for dispatch Tx: %s", evtData.swp.Hash.String())
	}
}

// getSyncData loops thru filtered uniswap sync event data and adds them into the chanel for processing
func (um *UniswapMonitor) getSyncData(syncEvent *contracts.UniswapPairSync, pair common.Address) (*evtSwap, error) {

	blkNumber := hexutil.Uint64(syncEvent.Raw.BlockNumber)

	blk, err := um.repo.BlockByNumber(&blkNumber)
	if err != nil {
		um.log.Errorf("Block was not found for uniswap sync event on pair %s; %s", pair.String(), err.Error())
		return nil, err
	}

	// log action
	um.log.Debugf("Loading uniswap sync event data from block nr# %d, tx: %s, reserve0: %s, reserve1: %s",
		syncEvent.Raw.BlockNumber, syncEvent.Raw.TxHash.String(),
		syncEvent.Reserve0.String(), syncEvent.Reserve1.String())

	zero := new(big.Int).SetUint64(0)

	// prep sending struct
	swap := &types.Swap{
		Type:        types.SwapSync,
		BlockNumber: &blk.Number,
		TimeStamp:   &blk.TimeStamp,
		Pair:        pair,
		Hash:        types.BytesToHash(syncEvent.Raw.TxHash.Bytes()),
		Amount0In:   zero,
		Amount0Out:  zero,
		Amount1In:   zero,
		Amount1Out:  zero,
		Reserve0:    syncEvent.Reserve0,
		Reserve1:    syncEvent.Reserve1,
	}
	return &evtSwap{swp: swap}, nil
}
