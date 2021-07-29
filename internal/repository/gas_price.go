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
	"fantom-api-graphql/internal/types"
	"math"
	"math/big"
	"sync"
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

// gasPriceSuggestionMonitor represents a monitor checking the suggested gas price
// periodically and making a record of
type gasPriceSuggestionMonitor struct {
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
}

// newGasPriceSuggestionMonitor creates a new gas price suggestion monitor instance.
func newGasPriceSuggestionMonitor(repo Repository, log logger.Logger, wg *sync.WaitGroup) *gasPriceSuggestionMonitor {
	// calculate expected amount of ticks based on the size of single period and the ticker interval
	expTicks := int(gasPriceSuggestionPeriodInterval / gasPriceSuggestionTickerInterval)

	// create new blockScanner instance
	return &gasPriceSuggestionMonitor{
		service: newService("gas price suggestion tracker", repo, log, wg),
		ticks:   make([]int64, expTicks),
	}
}

// run starts the gas price suggestions tracking
func (gps *gasPriceSuggestionMonitor) run() {
	// start go routine for processing
	gps.wg.Add(1)
	go gps.monitor()
}

// monitor executes periodical check on gas price suggestions
// collecting data by periods and storing them into the off-chain database
// for analysis.
func (gps *gasPriceSuggestionMonitor) monitor() {
	// create the ticker used to signal another read and an interval flip
	var rt = time.NewTicker(gasPriceSuggestionTickerInterval)
	var pt = time.NewTicker(gasPriceSuggestionPeriodInterval)
	var lt = time.NewTicker(gasPriceSuggestionLogInterval)

	// don't forget to sign off after we are done
	defer func() {
		// stop ticking on both tickers
		rt.Stop()
		pt.Stop()

		// signal and log
		gps.wg.Done()
		gps.log.Notice("gas price suggestion tracking is closed")
	}()

	// flip the period to start the first one
	gps.flip()
	gps.log.Notice("gas price suggestion tracking is active")

	// loop here
	for {
		select {
		case <-gps.sigStop:
			return
		case <-rt.C:
			// do the reading
			gps.read()
		case <-pt.C:
			// flip the period
			gps.flip()
		case <-lt.C:
			// log the status
			gps.log.Infof("gas price suggestion since %s range [%d, %d] with %d ticks", gps.start.Format(time.RFC850), gps.min, gps.max, gps.count)
		}
	}
}

// read makes the reading on the current gas price suggestion
// from the Opera node.
func (gps *gasPriceSuggestionMonitor) read() {
	// get the current suggested gas price value
	gs, err := gps.repo.GasPrice()
	if err != nil {
		gps.log.Errorf("can not read the suggested gas price; %s", err.Error())
		return
	}

	// adjust to desired decimals
	val := new(big.Int).Div(gs.ToInt(), types.TransactionGasCorrection).Int64()

	// add to ticks pool
	gps.ticks[gps.count] = val
	gps.count++

	// check max
	if val > gps.max {
		gps.max = val
	}

	// check min
	if val < gps.min {
		gps.min = val
	}
}

// flip closes the current reading period and starts a new one.
// The same function is also used to initialize the first period.
func (gps *gasPriceSuggestionMonitor) flip() {
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

// store stores this period into the persistent storage.
func (gps *gasPriceSuggestionMonitor) store(now time.Time) {
	// prep and store the period data
	err := gps.repo.StoreGasPricePeriod(&types.GasPricePeriod{
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
		gps.log.Errorf("could not store gas price period data; %s", err.Error())
	}
}

// avg calculates the average gas price suggestion value of the current period from the collected ticks.
// Please note we assume the ticks collection time frames are distributed evenly in the period,
// we should be using weighted avg instead if it wouldn't be the case.
func (gps *gasPriceSuggestionMonitor) avg() int64 {
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

// StoreGasPricePeriod stores the given gas price period data in the persistent storage
func (p *proxy) StoreGasPricePeriod(gp *types.GasPricePeriod) error {
	return p.db.AddGasPricePeriod(gp)
}
