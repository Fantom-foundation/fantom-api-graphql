package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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
