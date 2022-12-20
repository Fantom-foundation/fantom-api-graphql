// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/graph-gophers/graphql-go"
	"time"
)

// defaultFeeFlowListStartingRange represents the default range we provide the fee flow for.
const defaultFeeFlowListStartingRange = -30 * 24 * time.Hour

// FeeFlowDaily represents a tick structure of the daily fee flow.
type FeeFlowDaily types.FtmDailyBurn

// DailyFeeFlow resolves the flow of transaction fees over time.
func (rs *rootResolver) DailyFeeFlow(args struct {
	From *graphql.Time
	To   *graphql.Time
}) ([]*FeeFlowDaily, error) {
	now := time.Now().UTC()

	// create a standard default range ending now
	if args.To == nil || args.To.After(now) {
		args.To = &graphql.Time{Time: now.Add(-24 * time.Hour)}
	}

	// adjust starting time
	if args.From == nil || args.From.After(args.To.Time) {
		from := args.To.Add(defaultFeeFlowListStartingRange)
		args.From = &graphql.Time{Time: from}
	}

	// load the data
	list, err := repository.R().FeeFlow(args.From.Time, args.To.Time)
	if err != nil {
		log.Criticalf("failed to load fee flow list; %s", err.Error())
		return nil, err
	}

	// retype to the expressed GraphQL list
	out := make([]*FeeFlowDaily, len(list))
	for i, f := range list {
		out[i] = (*FeeFlowDaily)(f)
	}

	return out, nil
}

// Date returns the aggregated tick date resolvable as GraphQL Time type.
func (fd *FeeFlowDaily) Date() graphql.Time {
	return graphql.Time{Time: fd.TickDate}
}

// Fee returns Long encoded fee amount.
func (fd *FeeFlowDaily) Fee() hexutil.Uint64 {
	return (hexutil.Uint64)(uint64(fd.FeeAmount))
}

// FeeFTM returns Long encoded fee amount in FTM units.
func (fd *FeeFlowDaily) FeeFTM() float64 {
	return float64(fd.FeeAmount) / types.BurnFTMDecimalsCorrection
}

// Burned returns Long encoded burned fee slice amount.
func (fd *FeeFlowDaily) Burned() hexutil.Uint64 {
	return (hexutil.Uint64)(uint64(fd.BurnedAmount))
}

// BurnedFTM returns Long encoded burned fee slice amount in FTM units.
func (fd *FeeFlowDaily) BurnedFTM() float64 {
	return float64(fd.BurnedAmount) / types.BurnFTMDecimalsCorrection
}

// Treasury returns Long encoded treasury fee slice amount.
func (fd *FeeFlowDaily) Treasury() hexutil.Uint64 {
	return (hexutil.Uint64)(uint64(fd.TreasuryAmount))
}

// TreasuryFTM returns Long encoded treasury fee slice amount in FTM units.
func (fd *FeeFlowDaily) TreasuryFTM() float64 {
	return float64(fd.TreasuryAmount) / types.BurnFTMDecimalsCorrection
}

// Rewards returns Long encoded rewards fee slice amount.
func (fd *FeeFlowDaily) Rewards() hexutil.Uint64 {
	return (hexutil.Uint64)(uint64(fd.RewardsAmount))
}

// RewardsFTM returns Long encoded rewards fee slice amount in FTM units.
func (fd *FeeFlowDaily) RewardsFTM() float64 {
	return float64(fd.RewardsAmount) / types.BurnFTMDecimalsCorrection
}
