// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"time"
)

// defaultTrxVolumeRange defines the range we pull the trx flow for by default.
const defaultTrxVolumeRange = -(90 * 24) * time.Hour

// DailyTrxVolume defines the single day aggregation value.
type DailyTrxVolume struct {
	types.DailyTrxVolume
}

// TrxVolume resolves list of daily aggregations of the network transaction flow.
func (rs *rootResolver) TrxVolume(args struct {
	From *string
	To   *string
}) ([]*DailyTrxVolume, error) {
	// get the date range
	from, to, err := trxVolumeRange(args)
	if err != nil {
		return nil, err
	}

	// load data
	dv, err := repository.R().TrxFlowVolume(from, to)
	if err != nil {
		return nil, err
	}

	// load the list
	list := make([]*DailyTrxVolume, len(dv))
	for i, v := range dv {
		list[i] = &DailyTrxVolume{*v}
	}
	return list, nil
}

// TrxGasSpeed resolves the gas consumption speed speed
// of the network in transactions processed per second.
func (rs *rootResolver) TrxGasSpeed(args struct {
	Range int32
	To    *string
}) (val float64, err error) {
	// make sure to obey the minimal range
	if args.Range < 60 {
		args.Range = 60
	}

	// collect target time
	to := time.Now().UTC()
	if args.To != nil {
		to, err = time.Parse(time.RFC3339, *args.To)
		if err != nil {
			return 0.0, err
		}
	}

	// get the value
	from := to.Add(time.Duration(-args.Range) * time.Second)

	// log what we do
	log.Noticef("calculating gas speed from %s to %s", from.String(), to.String())
	return repository.R().TrxGasSpeed(&from, &to)
}

// TrxSpeed resolves the recent speed of the network in transactions processed per second.
func (rs *rootResolver) TrxSpeed(args struct {
	Range int32
}) (float64, error) {
	// make sure to obey the minimal range
	if args.Range < 60 {
		args.Range = 60
	}
	return repository.R().TrxFlowSpeed(args.Range)
}

// trxVolumeRange generates the time range for trx volume resolver.
func trxVolumeRange(args struct {
	From *string
	To   *string
}) (*time.Time, *time.Time, error) {
	// get the default target time to the last midnight
	var err error
	h, m, s := time.Now().Clock()
	to := time.Now().Add(time.Duration(-(h*3600 + m*60 + s)) * time.Second)

	// parse the target/end time
	if args.To != nil {
		to, err = time.Parse("2006-01-02", *args.To)
		if err != nil {
			return nil, nil, err
		}
	}

	// parse from
	from := to.Add(defaultTrxVolumeRange)
	if args.From != nil {
		from, err = time.Parse("2006-01-02", *args.From)
		if err != nil {
			return nil, nil, err
		}
	}

	// make sure the from is before to
	if from.After(to) {
		return nil, nil, fmt.Errorf("invalid date range received")
	}
	return &from, &to, nil
}

// Amount resolves the amount of native tokens transferred.
func (dtv *DailyTrxVolume) Amount() hexutil.Big {
	val := new(big.Int).Mul(new(big.Int).SetInt64(dtv.DailyTrxVolume.AmountAdjusted), types.TransactionDecimalsCorrection)
	return hexutil.Big(*val)
}

// Volume resolves the number of transactions in Int format.
func (dtv *DailyTrxVolume) Volume() int32 {
	return int32(dtv.DailyTrxVolume.Counter)
}

// Gas resolves the amount of gas consumed by transactions on the network.
func (dtv *DailyTrxVolume) Gas() hexutil.Big {
	val := new(big.Int).SetInt64(dtv.DailyTrxVolume.Gas)
	return hexutil.Big(*val)
}
