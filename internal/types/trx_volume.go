package types

import (
	"time"
)

// DailyTrxVolume represents a volume of daily transaction aggregation.
type DailyTrxVolume struct {
	Day            string    `bson:"_id"`
	Stamp          time.Time `bson:"stamp"`
	Counter        int64     `bson:"value"`
	AmountAdjusted int64     `bson:"volume"`
	Gas            int64     `bson:"gas"`
}
