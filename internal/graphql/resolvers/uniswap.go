// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// UniswapPair represents a pair of tokens in Uniswap protocol.
type UniswapPair struct {
	PairAddress common.Address
}

// UniswapPairVolume represents swap volume data
type UniswapPairVolume struct {
	*UniswapPair
	PairAddress common.Address
	InFUSD      bool
	TokenPrice  hexutil.Big
}

// DefiTimeVolume represents swap volume for given pair and time interval
type DefiTimeVolume struct {
	PairAddress common.Address
	Time        string
	Value       hexutil.Big
}

// DefiTimeReserve represents the time reserve on a given Uniswap pair.
type DefiTimeReserve struct {
	types.DefiTimeReserve
	*UniswapPair
}

// NewUniswapPair creates a new instance of resolvable UniswapPair token.
func NewUniswapPair(adr *common.Address) *UniswapPair {
	// make the instance of the token
	return &UniswapPair{
		PairAddress: *adr,
	}
}

// defiUniswapPairs load list of Uniswap pairs once in concurrent threads.
func (rs *rootResolver) defiUniswapPairs() []*UniswapPair {
	// make sure to do this only once
	list, err, _ := rs.cg.Do("uniswap-pairs", func() (interface{}, error) {
		// get the list of pair addresses
		pairs, err := repository.R().UniswapKnownPairs()
		if err != nil || pairs == nil {
			return make([]*UniswapPair, 0), nil
		}

		// build the output list
		list := make([]*UniswapPair, len(pairs))
		for i, adr := range pairs {
			list[i] = NewUniswapPair(&adr)
		}
		return list, nil
	})
	if err != nil {
		return make([]*UniswapPair, 0)
	}
	return list.([]*UniswapPair)
}

// DefiUniswapPairs resolves list of
func (rs *rootResolver) DefiUniswapPairs() []*UniswapPair {
	return rs.defiUniswapPairs()
}

// DefiUniswapAmountsOut resolves a list of output amounts for the given
// input amount and a list of tokens to be used to make the swap operation.
func (rs *rootResolver) DefiUniswapAmountsOut(args *struct {
	AmountIn hexutil.Big
	Tokens   []common.Address
}) ([]hexutil.Big, error) {
	return repository.R().UniswapAmountsOut(args.AmountIn, args.Tokens)
}

// DefiUniswapAmountsIn resolves a list of input amounts for the given
// output amount and a list of tokens to be used to make the swap operation.
func (rs *rootResolver) DefiUniswapAmountsIn(args *struct {
	AmountOut hexutil.Big
	Tokens    []common.Address
}) ([]hexutil.Big, error) {
	return repository.R().UniswapAmountsIn(args.AmountOut, args.Tokens)
}

// DefiUniswapQuoteLiquidity resolves a list of optimal amounts of tokens
// to be added to both sides of a pair on addLiquidity call.
func (rs *rootResolver) DefiUniswapQuoteLiquidity(args *struct {
	Tokens    []common.Address
	AmountsIn []hexutil.Big
}) ([]hexutil.Big, error) {
	// make sure the number of tokens make sense
	if args.Tokens == nil || len(args.Tokens) != 2 {
		return nil, fmt.Errorf("invalid tokens pair given")
	}

	// make sure the number of input prices make sense
	if args.AmountsIn == nil || len(args.AmountsIn) != 2 {
		return nil, fmt.Errorf("invalid input amounts pair given")
	}

	// get the pair address for the given set of tokens
	pair, err := repository.R().UniswapPair(&args.Tokens[0], &args.Tokens[1])
	if err != nil {
		return nil, err
	}

	// get normalized tokens order
	tokens, err := repository.R().UniswapTokens(pair)
	if err != nil {
		return nil, err
	}

	// make sure to call the amounts correctly
	if tokens[0] == args.Tokens[0] {
		return rs.uniswapOptimalLiquidity(pair, &args.AmountsIn[0], &args.AmountsIn[1])
	}

	// tokens came in in reversed order
	if tokens[0] == args.Tokens[1] {
		val, err := rs.uniswapOptimalLiquidity(pair, &args.AmountsIn[1], &args.AmountsIn[0])
		if err != nil {
			return nil, err
		}

		// reverse the order of value in response
		return []hexutil.Big{val[1], val[0]}, nil
	}

	// sanity check, tokens don't match the original pair?
	return nil, fmt.Errorf("the pair tokens don't match with input tokens")
}

// uniswapQuoteLiquidity calculates the optimal liquidity advance on addLiquidity call.
func (rs *rootResolver) uniswapOptimalLiquidity(
	pair *common.Address,
	amountAIn *hexutil.Big,
	amountBIn *hexutil.Big,
) ([]hexutil.Big, error) {
	// get amount of reserves
	reserves, err := repository.R().UniswapReserves(pair)
	if err != nil {
		return nil, err
	}

	// no liquidity on the pair at all? simply confirm desired values
	if 0 == reserves[0].ToInt().Cmp(zeroInt) && 0 == reserves[1].ToInt().Cmp(zeroInt) {
		return []hexutil.Big{*amountAIn, *amountBIn}, nil
	}

	// get side B optimal
	optimalB, err := repository.R().UniswapQuoteInput(*amountAIn, reserves[0], reserves[1])
	if err != nil {
		return nil, err
	}

	// optimal amount on B side is lower or the same as the input amount on B side?
	if 0 > optimalB.ToInt().Cmp(amountBIn.ToInt()) || 0 == optimalB.ToInt().Cmp(amountBIn.ToInt()) {
		return []hexutil.Big{*amountAIn, optimalB}, nil
	}

	// optimal B si higher than the input offered; calculate optimal A from the reversed reserves
	optimalA, err := repository.R().UniswapQuoteInput(*amountBIn, reserves[1], reserves[0])
	if err != nil {
		return nil, err
	}

	// optimal A must be lower or same as the desired input
	if 0 < optimalA.ToInt().Cmp(amountAIn.ToInt()) {
		return nil, fmt.Errorf("neither optimal value matches inputs")
	}

	return []hexutil.Big{optimalA, *amountBIn}, nil
}

// Tokens resolves a list of tokens of the given Uniswap pair.
func (up *UniswapPair) Tokens() ([]*ERC20Token, error) {
	// load addresses
	tokens, err := repository.R().UniswapTokens(&up.PairAddress)
	if err != nil {
		return nil, err
	}

	// make a sanity check, the pair should contain exactly two tokens
	// since it's called pair for a good reason
	if 2 != len(tokens) {
		return nil, fmt.Errorf("invalid pair tokens list")
	}

	// make the list container
	list := make([]*ERC20Token, len(tokens))
	for i, adr := range tokens {
		erc := NewErc20Token(&adr)
		list[i] = erc
	}
	return list, nil
}

// DefiUniswapVolumes returns all swap pairs and their information for swap volumes
func (rs *rootResolver) DefiUniswapVolumes() []*UniswapPairVolume {
	// get all the pairs
	pairs := rs.defiUniswapPairs()

	// create empty list as a result object
	list := make([]*UniswapPairVolume, len(pairs))
	for i, pair := range pairs {
		// get thr pair tokens
		tl, err := repository.R().UniswapTokens(&pair.PairAddress)
		if err != nil {
			return list
		}

		// get token price for denomination
		isDenominated := true
		tokenAPrice, err := repository.R().DefiTokenPrice(&tl[0])
		if err != nil {
			tokenAPrice = hexutil.Big{}
			isDenominated = false
		}

		// fill result list
		list[i] = &UniswapPairVolume{
			UniswapPair: pair,
			PairAddress: pair.PairAddress,
			TokenPrice:  tokenAPrice,
			InFUSD:      isDenominated,
		}
	}
	return list
}

func (upv *UniswapPairVolume) getVolumeTillNow(fromTime int64) (hexutil.Big, error) {
	toTime := time.Now().UTC().Unix()
	swapVolume, err := repository.R().UniswapVolume(&upv.PairAddress, fromTime, toTime)
	if err != nil {
		return hexutil.Big{}, err
	}
	return hexutil.Big(*swapVolume.Volume), nil
}

// DailyVolume returns swap volume for last 24 hours
func (upv *UniswapPairVolume) DailyVolume() (hexutil.Big, error) {
	fromTime := time.Now().UTC().AddDate(0, 0, -1).Unix()
	return upv.getVolumeTillNow(fromTime)
}

// WeeklyVolume returns swap volume for last 7 days
func (upv *UniswapPairVolume) WeeklyVolume() (hexutil.Big, error) {
	fromTime := time.Now().UTC().AddDate(0, 0, -7).Unix()
	return upv.getVolumeTillNow(fromTime)
}

// MonthlyVolume returns swap volume for last month
func (upv *UniswapPairVolume) MonthlyVolume() (hexutil.Big, error) {
	fromTime := time.Now().UTC().AddDate(0, -1, 0).Unix()
	return upv.getVolumeTillNow(fromTime)
}

// YearlyVolume returns swap volume for last year
func (upv *UniswapPairVolume) YearlyVolume() (hexutil.Big, error) {
	fromTime := time.Now().UTC().AddDate(-1, 0, 0).Unix()
	return upv.getVolumeTillNow(fromTime)
}

// IsInFUSD indicates if TokenA from the pair has a price value to be able
// to calculate value in fUSD
func (upv *UniswapPairVolume) IsInFUSD() (bool, error) {
	return upv.InFUSD, nil
}

// DefiTimeVolumes resolves swap volumes for given pair
// If dates are not given, then it returns last month values
func (rs *rootResolver) DefiTimeVolumes(args *struct {
	Address    common.Address
	Resolution *string
	FromDate   *int32
	ToDate     *int32
}) []*DefiTimeVolume {
	// decode dates
	var fDate int64
	if args.FromDate != nil {
		fDate = (int64)(*args.FromDate)
	} else {
		fDate = time.Now().UTC().AddDate(0, -1, 0).Unix()
	}
	tDate := checkDate(args.ToDate)

	// decode resolution
	resolution := ""
	if args.Resolution != nil {
		resolution = *args.Resolution
	}

	// get volumes from DB repository
	swapVolumes, err := repository.R().UniswapTimeVolumes(&args.Address, resolution, fDate, tDate)
	if err != nil {
		log.Errorf("Can not get swap volumes from DB repository: %s", err.Error())
		return make([]*DefiTimeVolume, 0)
	}

	// iterate thru results and build final list
	list := make([]*DefiTimeVolume, len(swapVolumes))
	for i, volume := range swapVolumes {
		list[i] = &DefiTimeVolume{
			PairAddress: *volume.PairAddress,
			Time:        volume.DateString,
			Value:       hexutil.Big(*volume.Volume),
		}
	}
	return list
}

// DefiTimePrices resolves swap prices for given pair
// If dates are not given, then it returns last month values
func (rs *rootResolver) DefiTimePrices(args *struct {
	Address    common.Address
	Resolution *string
	FromDate   *int32
	ToDate     *int32
	Direction  *int32
}) []types.DefiTimePrice {
	//check date values
	var fDate int64
	if args.FromDate != nil {
		fDate = (int64)(*args.FromDate)
	} else {
		fDate = time.Now().UTC().AddDate(0, -1, 0).Unix()
	}
	tDate := checkDate(args.ToDate)

	//check resolution value
	resolution := ""
	if args.Resolution != nil {
		resolution = *args.Resolution
	}

	// check direction value
	var dir int32
	if args.Direction != nil {
		dir = *args.Direction
	}

	// get prices from DB repository
	swapPrices, err := repository.R().UniswapTimePrices(&args.Address, resolution, fDate, tDate, dir)
	if err != nil {
		log.Errorf("Can not get uniswap prices from DB repository: %s", err.Error())
		return make([]types.DefiTimePrice, 0)
	}
	return swapPrices
}

// Reserves resolves a list of token reserves of the given Uniswap pair.
func (up *UniswapPair) Reserves() ([]hexutil.Big, error) {
	return repository.R().UniswapReserves(&up.PairAddress)
}

// ReservesTimeStamp resolves reserves of the given Uniswap pair.
func (up *UniswapPair) ReservesTimeStamp() (hexutil.Uint64, error) {
	return repository.R().UniswapReservesTimeStamp(&up.PairAddress)
}

// CumulativePrices resolves a list of token cumulative prices
// of the given Uniswap pair.
func (up *UniswapPair) CumulativePrices() ([]hexutil.Big, error) {
	return repository.R().UniswapCumulativePrices(&up.PairAddress)
}

// TotalSupply resolves the total amount of pair tokens, e.g. the share pool
// of the given Uniswap pair.
func (up *UniswapPair) TotalSupply() (hexutil.Big, error) {
	return repository.R().Erc20TotalSupply(&up.PairAddress)
}

// ShareOf resolves the total amount of a share of the given user on the given Uniswap pair.
func (up *UniswapPair) ShareOf(args *struct{ User common.Address }) (hexutil.Big, error) {
	return repository.R().Erc20BalanceOf(&up.PairAddress, &args.User)
}

// LastKValue resolves the last value of the pool control coefficient.
func (up *UniswapPair) LastKValue() (hexutil.Big, error) {
	return repository.R().UniswapLastKValue(&up.PairAddress)
}

func checkDate(td *int32) int64 {
	if td != nil {
		return (int64)(*td)
	}
	return 0
}

// DefiTimeReserves resolves uniswap reserves for given pair
// If dates are not given, then it returns last month values
func (rs *rootResolver) DefiTimeReserves(args *struct {
	Address    common.Address
	Resolution *string
	FromDate   *int32
	ToDate     *int32
}) []DefiTimeReserve {
	//check date values
	var fDate int64
	if args.FromDate != nil {
		fDate = (int64)(*args.FromDate)
	} else {
		fDate = time.Now().UTC().AddDate(0, -1, 0).Unix()
	}
	tDate := checkDate(args.ToDate)

	//check resolution value
	resolution := ""
	if args.Resolution != nil {
		resolution = *args.Resolution
	}

	// get reserves from DB repository
	timeReserves, err := repository.R().UniswapTimeReserves(&args.Address, resolution, fDate, tDate)
	if err != nil {
		log.Errorf("Can not get uniswap reserves from DB repository: %s", err.Error())
		return make([]DefiTimeReserve, 0)
	}

	list := make([]DefiTimeReserve, len(timeReserves))
	for i, timeReserve := range timeReserves {
		list[i] = DefiTimeReserve{
			DefiTimeReserve: timeReserve,
			UniswapPair:     NewUniswapPair(&args.Address),
		}
	}
	return list
}
