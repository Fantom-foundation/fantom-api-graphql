// Package svc implements blockchain data processing services.
package svc

import (
	"fantom-api-graphql/internal/types"
	"fmt"
	"math"
	"math/big"
	"time"
)

const (
	// gasPriceSuggestionTickerInterval represents the interval in which we collect
	// suggested gas price amount to get the period data.
	gasPriceSuggestionTickerInterval = 2 * time.Second

	// gasPriceSuggestionPeriodInterval represents the interval in which we close
	// the current suggested gas price period and start a new one.
	gasPriceSuggestionPeriodInterval = 10 * time.Minute

	// gasPriceSuggestionLogInterval represents the interval
	// in which we log the current state of the monitor.
	gasPriceSuggestionLogInterval = 150 * time.Second
)

// gpsMonitor represents a monitor checking the suggested gas price
// periodically and tracking the values in persistent repository.
type gpsMonitor struct {
	service

	// start represents the starting time of the current period
	start *time.Time

	// ticks represent the set of collected gas price ticks in the active period
	ticks []int64

	// count represents the number of ticks collected in the active period
	count int

	// max represents the max suggested value of this period
	max int64

	// min represents the min suggested value of this period
	min int64

	// monTicker controls gas price pulls
	monTicker *time.Ticker

	// fliTicker controls gas price periods flip
	flipTicker *time.Ticker

	// logTicker controls the current state logging
	logTicker *time.Ticker
}

// name returns a human-readable name of the service used by the manager.
func (gps *gpsMonitor) name() string {
	return "gas price monitor"
}

// init prepares the gas monitor to perform its function.
func (gps *gpsMonitor) init() {
	gps.sigStop = make(chan struct{})

	// calculate number of expected ticks and make the tick container for them
	expTicks := int(gasPriceSuggestionPeriodInterval/gasPriceSuggestionTickerInterval) + 1
	gps.ticks = make([]int64, expTicks)
}

// run starts the gas price suggestions tracking
func (gps *gpsMonitor) run() {
	// make sure we are orchestrated
	if gps.mgr == nil {
		panic(fmt.Errorf("no svc manager set on %s", gps.name()))
	}

	// start go routine for processing
	gps.mgr.started(gps)
	go gps.execute()
}

// close terminates the gas price suggestion monitor.
func (gps *gpsMonitor) close() {
	if gps.monTicker != nil {
		gps.monTicker.Stop()
		gps.flipTicker.Stop()
		gps.logTicker.Stop()
	}
	if gps.sigStop != nil {
		close(gps.sigStop)
	}
}

// execute performs periodical check on gas price suggestions
// collecting data by periods and storing them into the off-chain database
// for analysis.
func (gps *gpsMonitor) execute() {
	defer func() {
		gps.mgr.finished(gps)
	}()

	// start control tickers
	gps.monTicker = time.NewTicker(gasPriceSuggestionTickerInterval)
	gps.flipTicker = time.NewTicker(gasPriceSuggestionPeriodInterval)
	gps.logTicker = time.NewTicker(gasPriceSuggestionLogInterval)

	// flip the period to start the first one
	gps.flip()

	// loop here
	for {
		select {
		case <-gps.sigStop:
			return
		case <-gps.monTicker.C:
			gps.read()
		case <-gps.flipTicker.C:
			gps.flip()
		case <-gps.logTicker.C:
			log.Infof("gas price suggestion since %s range [%d, %d] with %d ticks", gps.start.Format(time.RFC850), gps.min, gps.max, gps.count)
		}
	}
}

// read makes the reading on the current gas price suggestion from connected the Opera node.
func (gps *gpsMonitor) read() {
	// get the current suggested gas price value
	gs, err := repo.GasPrice()
	if err != nil {
		log.Errorf("can not read the suggested gas price; %s", err.Error())
		return
	}

	// adjust to desired decimals
	val := new(big.Int).Div(gs.ToInt(), types.TransactionGasCorrection).Int64()

	// add to collected values pool
	gps.ticks[gps.count] = val
	gps.count++

	// check and adjust max
	if val > gps.max {
		gps.max = val
	}

	// check and adjust min
	if val < gps.min {
		gps.min = val
	}
}

// flip closes the current reading period and starts a new one.
// The same function is also used to initialize the first period.
func (gps *gpsMonitor) flip() {
	// where we are now?
	now := time.Now()

	// do we have any collected ticks? store the period if we do
	if gps.count > 0 {
		gps.store(now)
	}

	// reset for the next period
	gps.count = 0
	gps.max = 0
	gps.min = math.MaxInt64
	gps.start = &now
}

// store this period worth of data into the persistent storage.
func (gps *gpsMonitor) store(now time.Time) {
	// prep and store the period data
	err := repo.StoreGasPricePeriod(&types.GasPricePeriod{
		Type:  types.GasPricePeriodTypeSuggestion,
		Open:  gps.ticks[0],
		Close: gps.ticks[gps.count-1],
		Min:   gps.min,
		Max:   gps.max,
		Avg:   gps.avg(),
		From:  *gps.start,
		To:    now,
		Tick:  int64(gasPriceSuggestionTickerInterval),
	})
	if err != nil {
		log.Errorf("could not store gas price period data; %s", err.Error())
	}
}

// avg calculates the average gas price suggestion value of the current period from the collected ticks.
// Please note we assume the ticks' collection time frames are distributed evenly in the period,
// we should be using weighted avg instead if it wouldn't be the case.
func (gps *gpsMonitor) avg() int64 {
	// nothing to calculate here
	if gps.count <= 0 {
		return 0
	}

	// sum all the ticks
	var sum int64
	for i := 0; i < gps.count; i++ {
		sum += gps.ticks[i]
	}

	// avg the ticks
	return int64(float64(sum) / float64(gps.count))
}
