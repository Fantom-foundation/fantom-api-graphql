// Package svc implements blockchain data processing services.
package svc

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"time"
)

// stiScannerTickDuration represents the delay between STI info pull attempts.
const stiScannerTickDuration = 5 * time.Second

// stiScanner implements staker information scanner service.
type stiScanner struct {
	service
	scanTick *time.Ticker
	current  uint64
	top      uint64
}

// name returns the name of the service used by orchestrator.
func (sti *stiScanner) name() string {
	return "staker info scanner"
}

// run initializes and starts the staker information monitor.
func (sti *stiScanner) run() {
	// make sure we are orchestrated
	if sti.mgr == nil {
		panic(fmt.Errorf("no svc manager set on %s", sti.name()))
	}

	// signal orchestrator we started and go
	sti.mgr.started(sti)
	go sti.execute()
}

// close terminates the staker information scanner.
func (sti *stiScanner) close() {
	if sti.scanTick != nil {
		sti.scanTick.Stop()
	}

	if sti.sigStop != nil {
		sti.sigStop <- true
	}
}

// execute runs the staker information monitoring task.
func (sti *stiScanner) execute() {
	// make sure to clean up on exit
	defer func() {
		close(sti.sigStop)
		sti.mgr.finished(sti)
	}()

	// start the ticker
	sti.scanTick = time.NewTicker(stiScannerTickDuration)

	// loop before terminated
	for {
		// wait for stop or timeout
		select {
		case <-sti.sigStop:
			return
		case <-sti.scanTick.C:
			sti.next()
		}
	}
}

// syncTop re-checks the top validator ID.
// The top is used as a range of the scan loop.
func (sti *stiScanner) syncTop() error {
	// the top is updated on the scan loop reset only
	if sti.current == 0 {
		// the last staker id will be used as the new top
		ls, err := repo.LastValidatorId()
		if err != nil {
			log.Errorf("last staker id not available; %s", err.Error())
			return err
		}

		// remember the ceiling for this loop
		log.Noticef("%d validators found", ls)
		sti.top = ls

		// start the loop, the first val has #1
		sti.current++
	}
	return nil
}

// next tries to download and store staker information for the next val on the list.
func (sti *stiScanner) next() {
	// make sure we have the right ceiling, do not continue if the sync failed
	if err := sti.syncTop(); err != nil {
		return
	}

	// request the staker information
	log.Debugf("loading staker #%d information", sti.current)
	stakerID := new(big.Int).SetUint64(sti.current)
	info, err := repo.PullStakerInfo((*hexutil.Big)(stakerID))
	if err == nil && info != nil {
		err = repo.StoreStakerInfo((*hexutil.Big)(stakerID), info)
	}

	// anything failed?
	if err != nil || info == nil {
		log.Debugf("staker #%d information not available", sti.current)
	}

	// advance to the next staker closing the loop at the top
	sti.current++
	if sti.top < sti.current {
		sti.current = 0
	}
}
