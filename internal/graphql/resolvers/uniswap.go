package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// UniswapPair represents a pair of tokens in Uniswap protocol.
type UniswapPair struct {
	repo        repository.Repository
	PairAddress common.Address
}

// NewUniswapPair creates a new instance of resolvable UniswapPair token.
func NewUniswapPair(adr *common.Address, repo repository.Repository) *UniswapPair {
	// make the instance of the token
	return &UniswapPair{
		repo:        repo,
		PairAddress: *adr,
	}
}

// DefiUniswapPairs resolves list of
func (rs *rootResolver) DefiUniswapPairs() []*UniswapPair {
	// prep container for the pairs list
	list := make([]*UniswapPair, 0)

	// get the list of pair addresses
	pairs, err := rs.repo.UniswapPairs()
	if err != nil || pairs == nil {
		return list
	}

	// loop all addresses and build the output list
	for _, adr := range pairs {
		uPair := NewUniswapPair(&adr, rs.repo)
		list = append(list, uPair)
	}

	return list
}

// DefiUniswapAmountsOut resolves a list of output amounts for the given
// input amount and a list of tokens to be used to make the swap operation.
func (rs *rootResolver) DefiUniswapAmountsOut(args *struct {
	AmountIn hexutil.Big
	Tokens   []common.Address
}) ([]hexutil.Big, error) {
	return rs.repo.UniswapAmountsOut(args.AmountIn, args.Tokens)
}

// DefiUniswapAmountsOut resolves a list of input amounts for the given
// output amount and a list of tokens to be used to make the swap operation.
func (rs *rootResolver) DefiUniswapAmountsIn(args *struct {
	AmountOut hexutil.Big
	Tokens    []common.Address
}) ([]hexutil.Big, error) {
	return rs.repo.UniswapAmountsIn(args.AmountOut, args.Tokens)
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
	pair, err := rs.repo.UniswapPair(&args.Tokens[0], &args.Tokens[1])
	if err != nil {
		return nil, err
	}

	// get normalized tokens order
	tokens, err := rs.repo.UniswapTokens(pair)
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
	reserves, err := rs.repo.UniswapReserves(pair)
	if err != nil {
		return nil, err
	}

	// no liquidity on the pair at all? simply confirm desired values
	zeroInt := new(big.Int)
	if 0 == reserves[0].ToInt().Cmp(zeroInt) && 0 == reserves[1].ToInt().Cmp(zeroInt) {
		return []hexutil.Big{*amountAIn, *amountBIn}, nil
	}

	// get side B optimal
	optimalB, err := rs.repo.UniswapQuoteInput(*amountAIn, reserves[0], reserves[1])
	if err != nil {
		return nil, err
	}

	// optimal amount on B side is lower or the same as the input amount on B side?
	if 0 > optimalB.ToInt().Cmp(amountBIn.ToInt()) || 0 == optimalB.ToInt().Cmp(amountBIn.ToInt()) {
		return []hexutil.Big{*amountAIn, optimalB}, nil
	}

	// optimal B si higher than the input offered; calculate optimal A from the reversed reserves
	optimalA, err := rs.repo.UniswapQuoteInput(*amountBIn, reserves[1], reserves[0])
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
	tokens, err := up.repo.UniswapTokens(&up.PairAddress)
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

	// populate a resolvable list of tokens involved with the pair
	for i, adr := range tokens {
		erc := NewErc20Token(&adr, up.repo)
		list[i] = erc
	}

	return list, nil
}

// Reserves resolves a list of token reserves of the given Uniswap pair.
func (up *UniswapPair) Reserves() ([]hexutil.Big, error) {
	return up.repo.UniswapReserves(&up.PairAddress)
}

// ReservesTimeStamp resolves reserves of the given Uniswap pair.
func (up *UniswapPair) ReservesTimeStamp() (hexutil.Uint64, error) {
	return up.repo.UniswapReservesTimeStamp(&up.PairAddress)
}

// CumulativePrices resolves a list of token cumulative prices
// of the given Uniswap pair.
func (up *UniswapPair) CumulativePrices() ([]hexutil.Big, error) {
	return up.repo.UniswapCumulativePrices(&up.PairAddress)
}

// TotalSupply resolves the total amount of pair tokens, e.g. the share pool
// of the given Uniswap pair.
func (up *UniswapPair) TotalSupply() (hexutil.Big, error) {
	return up.repo.Erc20TotalSupply(&up.PairAddress)
}

// TotalSupply resolves the total amount of pair tokens, e.g. the share pool
// of the given Uniswap pair.
func (up *UniswapPair) ShareOf(args *struct{ User common.Address }) (hexutil.Big, error) {
	return up.repo.Erc20BalanceOf(&up.PairAddress, &args.User)
}

// LastKValue resolves the last value of the pool control coefficient.
func (up *UniswapPair) LastKValue() (hexutil.Big, error) {
	return up.repo.UniswapLastKValue(&up.PairAddress)
}
