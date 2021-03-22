/*
Package repository implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Opera/Lachesis full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import (
	"bytes"
	"fantom-api-graphql/internal/logger"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"sync"
	"time"
)

// how much time to wait between calls
var stiInitDelay = 100 * time.Millisecond
var stiIdleDelay = 10 * time.Second

// stiMonitor implements staker information monitoring service.
type stiMonitor struct {
	service
	onInit        bool
	currentStaker uint64
	topStaker     uint64
}

// newBlockScanner creates new blockchain blockScanner service.
func newStiMonitor(repo Repository, log logger.Logger, wg *sync.WaitGroup) *stiMonitor {
	// create new blockScanner instance
	return &stiMonitor{
		service: newService("stm monitor", repo, log, wg),
		onInit:  true,
	}
}

// run initializes and starts the staker information monitor.
func (sti *stiMonitor) run() {
	// log what we do
	sti.log.Notice("staker information monitor started")

	// start blockScanner
	sti.wg.Add(1)
	go sti.monitor()
}

// monitor runs the staker information monitoring task.
func (sti *stiMonitor) monitor() {
	// don't forget to sign off after we are done
	defer func() {
		// log finish
		sti.log.Notice("staker information monitor done")

		// signal to wait group we are done
		sti.wg.Done()
	}()

	// we will be changing the delay based on the round we do
	var delay time.Duration

	// loop before terminated
	for {
		// what delay to use
		delay = stiIdleDelay
		if sti.onInit {
			delay = stiInitDelay
		}

		// wait for stop or timeout
		select {
		case <-sti.sigStop:
			// stop signal received?
			return
		case <-time.After(delay):
			sti.next()
		}
	}
}

// next tries to download and store next staker information.
func (sti *stiMonitor) next() {
	// make sure we have the right ceiling
	if sti.currentStaker == 0 {
		// try to get the last staker id
		ls, err := sti.repo.LastStakerId()
		if err != nil {
			sti.log.Errorf("could not get the last staker id; %s", err.Error())
			return
		}

		// remember the ceiling for this loop
		sti.topStaker = uint64(ls)
		sti.currentStaker++
	}

	// inform
	sti.log.Debugf("updating information about staker #%d", sti.currentStaker)

	// request current staker
	info, err := sti.repo.PullStakerInfo(hexutil.Uint64(sti.currentStaker))
	if err == nil && info != nil {
		// got info? store it in cache
		err = sti.repo.StoreStakerInfo(hexutil.Uint64(sti.currentStaker), *info)
	}

	// anything failed?
	if err != nil || info == nil {
		// log issue, but we still continue to the next staker
		sti.log.Debugf("staker information of #%d not available", sti.currentStaker)
	}

	// advance to the next staker
	sti.currentStaker++

	// close the loop when we reached the last staker; from now on the monitor will loop on idle
	if sti.topStaker < sti.currentStaker {
		sti.currentStaker = 0
		sti.onInit = false
	}
}

// StakerInfo extracts an extended staker information from smart contact.
func (p *proxy) PullStakerInfo(id hexutil.Uint64) (*types.StakerInfo, error) {
	return p.rpc.StakerInfo(id)
}

// StoreStakerInfo stores staker information to in-memory cache for future use.
func (p *proxy) StoreStakerInfo(id hexutil.Uint64, sti types.StakerInfo) error {
	// push to in-memory cache
	err := p.cache.PushStakerInfo(id, sti)
	if err != nil {
		p.log.Error("staker info can net be kept")
		return err
	}

	return nil
}

// RetrieveStakerInfo gets staker information from in-memory if available.
func (p *proxy) RetrieveStakerInfo(id hexutil.Uint64) *types.StakerInfo {
	return p.cache.PullStakerInfo(id)
}

// IsStiContract returns true if the given address points to the STI contract.
func (p *proxy) IsStiContract(addr *common.Address) bool {
	return bytes.Equal(addr.Bytes(), p.cfg.Staking.StiContract.Bytes())
}
