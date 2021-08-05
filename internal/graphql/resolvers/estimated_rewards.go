package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// constants used in reward calculations
const (
	erwSecondsInDay   uint64 = 86400
	erwSecondsInWeek  uint64 = 86400 * 7
	erwSecondsInYear  uint64 = 31556926
	erwSecondsInMonth        = erwSecondsInYear / 12
)

// EstimatedRewards represents resolvable estimated rewards structure.
type EstimatedRewards struct {
	Staked      hexutil.Uint64
	TotalStaked hexutil.Big
	LastEpoch   Epoch
}

// weiToFtmDecimals represents decimal conversion between WEI and FTM units.
var weiToFtmDecimals = new(big.Int).SetUint64(1000000000000000000)

// NewEstimatedRewards builds new resolvable estimated rewards structure.
func NewEstimatedRewards(ep *types.Epoch, amount *hexutil.Uint64, total *hexutil.Big) EstimatedRewards {
	return EstimatedRewards{
		Staked:      *amount,
		TotalStaked: *total,
		LastEpoch:   Epoch{*ep},
	}
}

// estimateRewardsByAddress instantiates the estimated rewards for specified address if possible.
func (rs *rootResolver) estimateRewardsByAddress(addr *common.Address, ep *types.Epoch, total *hexutil.Big) (EstimatedRewards, error) {
	// try to get the address involved
	acc, err := repository.R().Account(addr)
	if err != nil {
		log.Error("invalid address or address not found")
		return EstimatedRewards{}, fmt.Errorf("address not found")
	}

	// inform to debug
	log.Debugf("calculating rewards estimation for address [%s]", acc.Address.String())

	// get the address balance
	balance, err := repository.R().AccountBalance(&acc.Address)
	if err != nil {
		log.Errorf("can not get balance for address [%s]", acc.Address.String())
		return EstimatedRewards{}, fmt.Errorf("address balance not found")
	}

	// get the value of the balance as Uint64 value
	val := hexutil.Uint64(new(big.Int).Div(balance.ToInt(), weiToFtmDecimals).Uint64())
	return NewEstimatedRewards(ep, &val, total), nil
}

// EstimateRewards resolves reward estimation for the given address or amount staked.
func (rs *rootResolver) EstimateRewards(args *struct {
	Address *common.Address
	Amount  *hexutil.Uint64
}) (EstimatedRewards, error) {
	// at least one of the parameters must be present
	if args == nil || (args.Address == nil && args.Amount == nil) {
		log.Error("can not calculate estimated rewards without parameters")
		return EstimatedRewards{}, fmt.Errorf("missing both address and amount")
	}

	// get the latest sealed epoch
	// the data could be delayed behind the real-time sealed epoch due to caching,
	// but we don't need that precise reflection here
	ep, err := repository.R().CurrentSealedEpoch()
	if err != nil {
		log.Errorf("can not get the current sealed epoch information; %s", err.Error())
		return EstimatedRewards{}, fmt.Errorf("current sealed epoch not found")
	}

	// get the current total staked amount
	total, err := repository.R().TotalStaked()
	if err != nil {
		log.Errorf("can not get the current total staked amount; %s", err.Error())
		return EstimatedRewards{}, fmt.Errorf("current total staked amount not found")
	}

	// if address is specified, pull the estimation from it
	if args.Address != nil {
		return rs.estimateRewardsByAddress(args.Address, ep, total)
	}
	return NewEstimatedRewards(ep, args.Amount, total), nil
}

// canCalculateRewards checks if the reward can actually be calculated
func (erw EstimatedRewards) canCalculateRewards() bool {
	zero := new(big.Int)
	return erw.Staked > 0 &&
		erw.LastEpoch.BaseRewardPerSecond.ToInt().Cmp(zero) > 0 &&
		erw.TotalStaked.ToInt().Cmp(zero) > 0
}

// getRewards calculates the rewards value for the given time period.
func (erw EstimatedRewards) getRewards(period uint64) hexutil.Big {
	// validate that we can actually calculate the value
	if !erw.canCalculateRewards() {
		fmt.Printf("can not calculate!")
		return hexutil.Big{}
	}

	// prep values and calculate results
	// (perSecond * period * stakedAmount) / totalStakedAmount
	base := new(big.Int).Mul(erw.LastEpoch.BaseRewardPerSecond.ToInt(), new(big.Int).SetUint64(period))
	staked := new(big.Int).Mul(base, new(big.Int).SetUint64(uint64(erw.Staked)))
	val := new(big.Int).Div(staked, erw.TotalStaked.ToInt())

	// return the value
	return hexutil.Big(*val)
}

// DailyReward calculates daily rewards for the given rewards estimation.
func (erw EstimatedRewards) DailyReward() hexutil.Big {
	return erw.getRewards(erwSecondsInDay)
}

// WeeklyReward calculates daily rewards for the given rewards estimation.
func (erw EstimatedRewards) WeeklyReward() hexutil.Big {
	return erw.getRewards(erwSecondsInWeek)
}

// MonthlyReward calculates daily rewards for the given rewards estimation.
func (erw EstimatedRewards) MonthlyReward() hexutil.Big {
	return erw.getRewards(erwSecondsInMonth)
}

// YearlyReward calculates daily rewards for the given rewards estimation.
func (erw EstimatedRewards) YearlyReward() hexutil.Big {
	return erw.getRewards(erwSecondsInYear)
}

// CurrentRewardRateYearly calculates average reward rate
// for any staked amount in average per year
func (erw EstimatedRewards) CurrentRewardRateYearly() int32 {
	// make sure we can calculate the yearly rate safely
	zero := new(big.Int)
	if erw.LastEpoch.BaseRewardPerSecond.ToInt().Cmp(zero) <= 0 || erw.TotalStaked.ToInt().Cmp(zero) <= 0 {
		return 0
	}

	// prep and calculate the rate
	base := new(big.Int).Mul(erw.LastEpoch.BaseRewardPerSecond.ToInt(), new(big.Int).SetUint64(erwSecondsInYear))
	toPct := new(big.Int).Mul(base, new(big.Int).SetUint64(100))
	val := new(big.Int).Div(toPct, erw.TotalStaked.ToInt())

	// get the value
	return int32(val.Int64())
}
