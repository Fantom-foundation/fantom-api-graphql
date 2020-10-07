package repository

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// NativeTokenAddress returns address of the native token wrapper, if available.
func (p *proxy) NativeTokenAddress() (*common.Address, error) {
	return p.rpc.NativeTokenAddress()
}

// UniswapPairs returns list of all token pairs managed by Uniswap core.
func (p *proxy) UniswapPairs() ([]common.Address, error) {
	return p.rpc.UniswapPairs()
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

// UniswapTokens returns list of addresses of tokens involved in a Uniswap pair.
func (p *proxy) UniswapTokens(pair *common.Address) ([]common.Address, error) {
	return p.rpc.UniswapTokens(pair)
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
