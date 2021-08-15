package repository

import (
	"fantom-api-graphql/internal/repository/rpc/contracts"
	"fantom-api-graphql/internal/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// NativeTokenAddress returns address of the native token wrapper, if available.
func (p *proxy) NativeTokenAddress() (*common.Address, error) {
	return p.rpc.NativeTokenAddress()
}

// UniswapPairs returns list of all token pairs managed by Uniswap core.
// We use cache to store the list temporarily, the list is refreshed from RCP when the cache record expires.
func (p *proxy) UniswapPairs() ([]common.Address, error) {
	// try the cache first
	l := p.cache.PullAllPairsList()
	if l != nil {
		return l, nil
	}

	// use RPC to get the fresh list and cache it for future use
	l, err := p.rpc.UniswapPairs(false)
	if err != nil {
		log.Errorf("uniswap pairs not available; %s", err.Error())
		return nil, err
	}
	p.cache.PushAllPairsList(l)
	return l, nil
}

// UniswapKnownPairs returns list of all known and whitelisted token pairs managed by Uniswap core.
func (p *proxy) UniswapKnownPairs() ([]common.Address, error) {
	return p.rpc.UniswapPairs(true)
}

// UniswapPair returns an address of an Uniswap pair for the given tokens.
func (p *proxy) UniswapPair(tokenA *common.Address, tokenB *common.Address) (*common.Address, error) {
	return p.rpc.UniswapPair(tokenA, tokenB)
}

// UniswapAmountsOut resolves a list of output amounts for the given
// input amount and a list of tokens to be used to make the swap operation.
func (p *proxy) UniswapAmountsOut(amountIn hexutil.Big, tokens []common.Address) ([]hexutil.Big, error) {
	return p.rpc.UniswapAmountsOut(amountIn, tokens)
}

// UniswapAmountsIn resolves a list of input amounts for the given
// output amount and a list of tokens to be used to make the swap operation.
func (p *proxy) UniswapAmountsIn(amountOut hexutil.Big, tokens []common.Address) ([]hexutil.Big, error) {
	return p.rpc.UniswapAmountsIn(amountOut, tokens)
}

// UniswapQuoteInput calculates optimal input on sibling token based on input amount and
// self reserves of the analyzed token.
func (p *proxy) UniswapQuoteInput(
	amountIn hexutil.Big,
	reserveMy hexutil.Big,
	reserveSibling hexutil.Big,
) (hexutil.Big, error) {
	return p.rpc.UniswapQuoteInput(amountIn, reserveMy, reserveSibling)
}

// UniswapTokens returns list of addresses of tokens involved in a Uniswap pair.
func (p *proxy) UniswapTokens(pair *common.Address) ([]common.Address, error) {
	var err error
	// try cache first
	tl := p.cache.PullUniswapPairTokens(pair)
	if nil == tl {
		// load the hard way
		tl, err = p.rpc.UniswapTokens(pair)
		if err == nil {
			// push to cache for future use
			p.cache.PushUniswapPairTokens(pair, tl)
		}
	}
	return tl, err
}

// UniswapReserves returns list of token reserve amounts in a Uniswap pair.
func (p *proxy) UniswapReserves(pair *common.Address) ([]hexutil.Big, error) {
	return p.rpc.UniswapReserves(pair)
}

// UniswapReservesTimeStamp returns the timestamp of the reserves of a Uniswap pair.
func (p *proxy) UniswapReservesTimeStamp(pair *common.Address) (hexutil.Uint64, error) {
	return p.rpc.UniswapReservesTimeStamp(pair)
}

// UniswapCumulativePrices returns list of token cumulative prices of a Uniswap pair.
func (p *proxy) UniswapCumulativePrices(pair *common.Address) ([]hexutil.Big, error) {
	return p.rpc.UniswapCumulativePrices(pair)
}

// UniswapLastKValue returns the last value of the pool control coefficient.
func (p *proxy) UniswapLastKValue(pair *common.Address) (hexutil.Big, error) {
	return p.rpc.UniswapLastKValue(pair)
}

// UniswapPairContract returns instance of this contract according to given pair address
func (p *proxy) UniswapPairContract(pairAddres *common.Address) (*contracts.UniswapPair, error) {
	return p.rpc.UniswapPairContract(pairAddres)
}

// UniswapAdd notifies a new incoming swap from blockchain to the repository.
func (p *proxy) UniswapAdd(swap *types.Swap) error {
	return p.db.UniswapAdd(swap)
}

// LastKnownSwapBlock returns number of the last block known to the repository with the swap event.
func (p *proxy) LastKnownSwapBlock() (uint64, error) {
	return p.db.LastKnownSwapBlock()
}

// UniswapUpdateLastKnownSwapBlock stores a last correctly saved swap block number into persistent storage.
func (p *proxy) UniswapUpdateLastKnownSwapBlock(blkNumber uint64) error {
	return p.db.UniswapUpdateLastKnownSwapBlock(blkNumber)
}

// UniswapFactoryContract returns an instance of an Uniswap factory
func (p *proxy) UniswapFactoryContract() (*contracts.UniswapFactory, error) {
	return p.rpc.UniswapFactoryContract()
}

// UniswapVolume returns swap volume for specified uniswap pair
// If toTime = 0, then it resolves volumes till now
func (p *proxy) UniswapVolume(pairAddress *common.Address, fromTime int64, toTime int64) (types.DefiSwapVolume, error) {
	return p.db.UniswapVolume(pairAddress, fromTime, toTime)
}

// UniswapTimeVolumes returns daily swap volume for specified uniswap pair and period of time
// If toTime = 0, then it resolves volumes till now
func (p *proxy) UniswapTimeVolumes(pairAddress *common.Address, resolution string, fromTime int64, toTime int64) ([]types.DefiSwapVolume, error) {
	return p.db.UniswapTimeVolumes(pairAddress, resolution, fromTime, toTime)
}

// UniswapTimePrices resolves price of swap trades for specified pair grouped by date interval.
// If toTime is 0, then it calculates prices till now
func (p *proxy) UniswapTimePrices(pairAddress *common.Address, resolution string, fromTime int64, toTime int64, direction int32) ([]types.DefiTimePrice, error) {
	return p.db.UniswapTimePrices(pairAddress, resolution, fromTime, toTime, direction)
}

// UniswapTimeReserves resolves reserves of uniswap trades for specified pair grouped by date interval.
// If toTime is 0, then it calculates prices till now
func (p *proxy) UniswapTimeReserves(pairAddress *common.Address, resolution string, fromTime int64, toTime int64) ([]types.DefiTimeReserve, error) {
	return p.db.UniswapTimeReserves(pairAddress, resolution, fromTime, toTime)
}

// UniswapActions provides list of uniswap actions stored in the persistent storage.
func (p *proxy) UniswapActions(pairAddress *common.Address, cursor *string, count int32, actionType int32) (*types.UniswapActionList, error) {
	return p.db.UniswapActions(pairAddress, cursor, count, actionType)
}
