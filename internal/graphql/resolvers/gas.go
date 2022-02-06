package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/graph-gophers/graphql-go"
	"math"
	"time"
)

// maxGasPriceListRange represents the max time span of the gas price list.
const maxGasPriceListRange = 90 * 24 * time.Hour

// GasPriceTick represents a tick of the gas price.
type GasPriceTick types.GasPricePeriod

// GasPriceList resolves a list of gas price ticks for the given time period.
func (rs *rootResolver) GasPriceList(args struct {
	From time.Time
	To   *time.Time
}) ([]*GasPriceTick, error) {
	// make sure to set the end time
	if args.To == nil {
		now := time.Now()
		args.To = &now
	}

	// validate start time; replace it with the max allowed if not valid
	diff := args.To.Sub(args.From)
	if diff < 0 || diff > maxGasPriceListRange {
		args.From = args.To.Add(-maxGasPriceListRange)
		diff = maxGasPriceListRange
	}

	// pull the data
	ticks, err := repository.R().GasPriceTicks(&args.From, args.To)
	if err != nil {
		return nil, err
	}

	// any ticks at all?
	if len(ticks) == 0 {
		return make([]*GasPriceTick, 0), nil
	}

	return aggregateGasPriceTicks(ticks, gasPriceListTickSize(diff)), nil
}

// aggregateGasPriceTicks collects aggregated ticks of gas price into an array with the specified tick duration.
func aggregateGasPriceTicks(ticks []types.GasPricePeriod, tickLen time.Duration) []*GasPriceTick {
	list := make([]*GasPriceTick, 0)

	var tickCount int
	var this *GasPriceTick = nil
	for _, next := range ticks {
		// did we reach the aggregated tick length?
		if this != nil && this.From.Add(tickLen).Before(next.From) {
			// recalculate average value
			if tickCount > 1 {
				this.Avg = (int64)(math.Round((float64)(this.Avg) / (float64)(tickCount)))
			}

			// push to the output
			list = append(list, this)
			this = nil
		}

		// make new aggregate if needed
		if this == nil {
			this = (*GasPriceTick)(&next)
			tickCount = 1
			continue
		}

		// add the next tick to the current aggregate
		this.To = next.To
		this.Close = next.Close
		this.Avg += next.Avg
		this.Tick += next.Tick
		if this.Min > next.Min {
			this.Min = next.Min
		}
		if this.Max < next.Max {
			this.Max = next.Max
		}
		tickCount++
	}

	return list
}

// gasPriceListTickSize derives individual gas price tick size based on the full list range.
// * 24 hour or less => 5 minutes
// * 1 day to 2 days => 20 minutes
// * 2 days to 7 days => 30 minutes
// * 7 days to 14 days => 1 hour
// * 14 days to 30 days => 4 hour
// * more than 30 days => 1 day
func gasPriceListTickSize(diff time.Duration) time.Duration {
	if diff <= 24*time.Hour {
		return 5 * time.Minute
	}

	if diff <= 48*time.Hour {
		return 20 * time.Minute
	}

	if diff <= 7*24*time.Hour {
		return 30 * time.Minute
	}

	if diff <= 14*24*time.Hour {
		return time.Hour
	}

	if diff <= 30*24*time.Hour {
		return 4 * time.Hour
	}

	return 24 * time.Hour
}

// FromTime returns the starting time of the tick.
func (gt *GasPriceTick) FromTime() graphql.Time {
	return graphql.Time{Time: gt.From}
}

// ToTime returns the ending time of the tick.
func (gt *GasPriceTick) ToTime() graphql.Time {
	return graphql.Time{Time: gt.To}
}

// OpenPrice returns the opening price of the tick.
func (gt *GasPriceTick) OpenPrice() hexutil.Uint64 {
	return hexutil.Uint64(gt.Open)
}

// ClosePrice returns the closing price of the tick.
func (gt *GasPriceTick) ClosePrice() hexutil.Uint64 {
	return hexutil.Uint64(gt.Close)
}

// MinPrice returns the minimal price reached within the tick.
func (gt *GasPriceTick) MinPrice() hexutil.Uint64 {
	return hexutil.Uint64(gt.Min)
}

// MaxPrice returns the maximal price reached within the tick.
func (gt *GasPriceTick) MaxPrice() hexutil.Uint64 {
	return hexutil.Uint64(gt.Max)
}

// AvgPrice returns the numerically average price reached within the tick.
func (gt *GasPriceTick) AvgPrice() hexutil.Uint64 {
	return hexutil.Uint64(gt.Avg)
}
