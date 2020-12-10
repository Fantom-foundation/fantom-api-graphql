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
	subsCh       []chan *contracts.UniswapPairSwap
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
		for _, subCh := range um.subsCh {
			close(subCh)
		}

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
	// creating swap event channel
	swapEventCh := make(chan *contracts.UniswapPairSwap)

	// make subscription to Uniswap swaps
	evt, err := contract.UniswapPairFilterer.WatchSwap(&bind.WatchOpts{}, swapEventCh, nil, nil)
	if err != nil {
		um.log.Errorf("Cannot subscribe for swap action on pair %s; %s", pair.String(), err.Error())
		return
	}
	// add channel and subsription to arrays to be able to close them
	um.subs = append(um.subs, evt)
	um.subsCh = append(um.subsCh, swapEventCh)

	// start routine for handling swaps for given pair
	go um.monitorPair(*pair, um.swapChan, swapEventCh, &evt)

}

// monitorPair uses channel to swap dispatcher to send data from swap event there, for storing them into persistent db
func (um *UniswapMonitor) monitorPair(pair common.Address, swapDispatchChan chan *evtSwap, swapEventCh chan *contracts.UniswapPairSwap, evt *event.Subscription) {

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
		BlockNumber: &blk.Number,
		TimeStamp:   &blk.TimeStamp,
		Pair:        pair,
		Sender:      swapEvent.Sender,
		To:          swapEvent.To,
		Hash:        types.BytesToHash(swapEvent.Raw.TxHash.Bytes()),
		Amount0In:   swapEvent.Amount0In,
		Amount0Out:  swapEvent.Amount0Out,
		Amount1In:   swapEvent.Amount1In,
		Amount1Out:  swapEvent.Amount1Out,
	}
	return &evtSwap{swp: swap}, nil
}
