/*
Package rpc implements bridge to Opera full node API interface.

We recommend using local IPC for fast and the most efficient inter-process communication between the API server
and an Opera/Opera node. Any remote RPC connection will work, but the performance may be significantly degraded
by extra networking overhead of remote RPC calls.

You should also consider security implications of opening Opera RPC interface for a remote access.
If you considering it as your deployment strategy, you should establish encrypted channel between the API server
and Opera RPC interface with connection limited to specified endpoints.

We strongly discourage opening Opera RPC interface for unrestricted Internet access.
*/
package rpc

import (
	"fantom-api-graphql/internal/repository/rpc/contracts"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

//go:generate tools/abigen.sh --abi ./contracts/abi/uniswap-factory.abi --pkg contracts --type UniswapFactory --out ./contracts/uniswap_factory.go
//go:generate tools/abigen.sh --abi ./contracts/abi/uniswap-pair.abi --pkg contracts --type UniswapPair --out ./contracts/uniswap_pair.go
//go:generate tools/abigen.sh --abi ./contracts/abi/uniswap-router.abi --pkg contracts --type UniswapRouter --out ./contracts/uniswap_router.go

// NativeTokenAddress returns an address of native token.
func (ftm *FtmBridge) NativeTokenAddress() (*common.Address, error) {
	// get the router contract if possible
	contract, err := contracts.NewUniswapRouter(ftm.uniswapConfig.Router, ftm.eth)
	if err != nil {
		ftm.log.Errorf("Uniswap router contract not found; %s", err.Error())
		return nil, err
	}

	// get the native token address
	adr, err := contract.WETH(nil)
	if err != nil {
		ftm.log.Errorf("Native token address not available; %s", err.Error())
		return nil, err
	}

	return &adr, nil
}

// UniswapPair returns an address of an Uniswap pair for the given tokens.
func (ftm *FtmBridge) UniswapPair(tokenA *common.Address, tokenB *common.Address) (*common.Address, error) {
	// get the router contract if possible
	contract, err := contracts.NewUniswapFactory(ftm.uniswapConfig.Core, ftm.eth)
	if err != nil {
		ftm.log.Errorf("Uniswap factory contract not found; %s", err.Error())
		return nil, err
	}

	// try to get the pair
	pair, err := contract.GetPair(nil, *tokenA, *tokenB)
	if err != nil {
		ftm.log.Errorf("Uniswap pair not found for tokens %s and %s; %s", tokenA.String(), tokenB.String(), err.Error())
		return nil, err
	}

	return &pair, nil
}

// UniswapPairs returns list of all token pairs managed by Uniswap core.
func (ftm *FtmBridge) UniswapPairs(whiteListedOnly bool) ([]common.Address, error) {
	// get the router contract if possible
	contract, err := contracts.NewUniswapFactory(ftm.uniswapConfig.Core, ftm.eth)
	if err != nil {
		ftm.log.Errorf("Uniswap factory contract not found; %s", err.Error())
		return nil, err
	}

	// get the number of pairs
	length, err := contract.AllPairsLength(nil)
	if err != nil || length == nil {
		ftm.log.Error("Uniswap pairs array length not available")
		return nil, err
	}

	// prep pairs container
	list := make([]common.Address, 0)

	// loop to pull all the pairs
	index := new(big.Int)
	for i := uint64(0); i < length.Uint64(); i++ {
		// get the pair address
		adr, err := contract.AllPairs(nil, index.SetUint64(i))
		if err != nil {
			ftm.log.Errorf("error loading Uniswap pair; %s", err.Error())
			continue
		}

		// check the pair (is it on the white list?)
		// add the address to the list, if it's ok
		if !whiteListedOnly || ftm.isUniswapPairWhitelisted(&adr) {
			list = append(list, adr)
		}
	}

	return list, nil
}

// isUniswapPairWhitelisted checks if the given Uniswap pair is whitelisted
// on our configuration.
func (ftm *FtmBridge) isUniswapPairWhitelisted(pair *common.Address) bool {
	// decode the pair address
	pairAddr := pair.String()

	// check the pair address against all the white listed pairs in config
	for _, addr := range ftm.uniswapConfig.PairsWhiteList {
		if strings.EqualFold(addr.String(), pairAddr) {
			return true
		}
	}

	// we did not find it
	return false
}

// UniswapQuoteInput calculates optimal input on sibling token based on input amount and
// self reserves of the analyzed token.
func (ftm *FtmBridge) UniswapQuoteInput(
	amountA hexutil.Big,
	reserveA hexutil.Big,
	reserveB hexutil.Big,
) (hexutil.Big, error) {
	// get the router contract if possible
	contract, err := contracts.NewUniswapRouter(ftm.uniswapConfig.Router, ftm.eth)
	if err != nil {
		ftm.log.Errorf("Uniswap router contract not found; %s", err.Error())
		return hexutil.Big{}, err
	}

	// get the quote amount
	amount, err := contract.Quote(nil, amountA.ToInt(), reserveA.ToInt(), reserveB.ToInt())
	if err != nil {
		ftm.log.Errorf("can not calculate Uniswap quote; %s", err.Error())
		return hexutil.Big{}, err
	}

	return hexutil.Big(*amount), nil
}

// UniswapAmountsOut resolves a list of output amounts for the given
// input amount and a list of tokens to be used to make the swap operation.
func (ftm *FtmBridge) UniswapAmountsOut(amountIn hexutil.Big, tokens []common.Address) ([]hexutil.Big, error) {
	// get the router contract if possible
	contract, err := contracts.NewUniswapRouter(ftm.uniswapConfig.Router, ftm.eth)
	if err != nil {
		ftm.log.Errorf("Uniswap router contract not found; %s", err.Error())
		return nil, err
	}

	return ftm.uniswapLoadAmounts(contract.GetAmountsOut, (*big.Int)(&amountIn), tokens)
}

// UniswapAmountsIn resolves a list of input amounts for the given
// output amount and a list of tokens to be used to make the swap operation.
func (ftm *FtmBridge) UniswapAmountsIn(amountOut hexutil.Big, tokens []common.Address) ([]hexutil.Big, error) {
	// get the router contract if possible
	contract, err := contracts.NewUniswapRouter(ftm.uniswapConfig.Router, ftm.eth)
	if err != nil {
		ftm.log.Errorf("Uniswap router contract not found; %s", err.Error())
		return nil, err
	}

	return ftm.uniswapLoadAmounts(contract.GetAmountsIn, (*big.Int)(&amountOut), tokens)
}

// uniswapLoadAmounts loads a list of swap amounts using the given loader function
// of the Uniswap router contract.
func (ftm *FtmBridge) uniswapLoadAmounts(
	loader func(*bind.CallOpts, *big.Int, []common.Address) ([]*big.Int, error),
	amount *big.Int,
	tokens []common.Address,
) ([]hexutil.Big, error) {
	// get the amounts list
	out, err := loader(nil, amount, tokens)
	if err != nil {
		ftm.log.Errorf("can not load swap amounts; %s", err.Error())
		return nil, err
	}

	// make the container
	list := make([]hexutil.Big, len(out))
	for i, val := range out {
		list[i] = hexutil.Big(*val)
	}

	return list, nil
}

// UniswapTokens returns list of addresses of tokens involved in a Uniswap pair.
func (ftm *FtmBridge) UniswapTokens(pair *common.Address) ([]common.Address, error) {
	// get the pair contract if possible
	contract, err := contracts.NewUniswapPair(*pair, ftm.eth)
	if err != nil {
		ftm.log.Errorf("Uniswap pair %s not found; %s", pair.String(), err.Error())
		return nil, err
	}

	// make the address container
	tokens := make([]common.Address, 2)

	// get the first token
	tokens[0], err = contract.Token0(nil)
	if err != nil {
		ftm.log.Errorf("Uniswap pair %s token A not found; %s", pair.String(), err.Error())
		return nil, err
	}

	// get the second token
	tokens[1], err = contract.Token1(nil)
	if err != nil {
		ftm.log.Errorf("Uniswap pair %s token B not found; %s", pair.String(), err.Error())
		return nil, err
	}

	return tokens, nil
}

// UniswapReserves returns list of token reserve amounts in a Uniswap pair.
func (ftm *FtmBridge) UniswapReserves(pair *common.Address) ([]hexutil.Big, error) {
	// get the reserves record from the contract
	rs, err := ftm.uniswapReservesRecord(pair)
	if err != nil {
		ftm.log.Errorf("Uniswap pair %s reserves not available; %s", pair.String(), err.Error())
		return nil, err
	}

	// make the reserves container
	reserves := make([]hexutil.Big, 2)
	reserves[0] = hexutil.Big(*rs.Reserve0)
	reserves[1] = hexutil.Big(*rs.Reserve1)

	return reserves, nil
}

// UniswapReservesTimeStamp returns the timestamp of the reserves of a Uniswap pair.
func (ftm *FtmBridge) UniswapReservesTimeStamp(pair *common.Address) (hexutil.Uint64, error) {
	// get the reserves record from the contract
	rs, err := ftm.uniswapReservesRecord(pair)
	if err != nil {
		ftm.log.Errorf("Uniswap pair %s reserves not available; %s", pair.String(), err.Error())
		return 0, err
	}

	return hexutil.Uint64(rs.BlockTimestampLast), nil
}

// UniswapCumulativePrices returns list of token cumulative prices of a Uniswap pair.
func (ftm *FtmBridge) UniswapCumulativePrices(pair *common.Address) ([]hexutil.Big, error) {
	// get the pair contract if possible
	contract, err := contracts.NewUniswapPair(*pair, ftm.eth)
	if err != nil {
		ftm.log.Errorf("Uniswap pair %s not found; %s", pair.String(), err.Error())
		return nil, err
	}

	// make the address container
	prices := make([]hexutil.Big, 2)
	var price *big.Int

	// get the first token price
	price, err = contract.Price0CumulativeLast(nil)
	if err != nil {
		ftm.log.Errorf("Uniswap pair %s cumulative price A not available; %s", pair.String(), err.Error())
		return nil, err
	}

	// add the price to list
	prices[0] = hexutil.Big(*price)

	// get the second token price
	price, err = contract.Price1CumulativeLast(nil)
	if err != nil {
		ftm.log.Errorf("Uniswap pair %s cumulative price B not available; %s", pair.String(), err.Error())
		return nil, err
	}

	// add the price to list
	prices[1] = hexutil.Big(*price)
	return prices, nil
}

// uniswapReservesRecord loads the reserves record for the given Uniswap pair.
func (ftm *FtmBridge) uniswapReservesRecord(pair *common.Address) (*struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	// get the pair contract if possible
	contract, err := contracts.NewUniswapPair(*pair, ftm.eth)
	if err != nil {
		ftm.log.Errorf("Uniswap pair %s not found; %s", pair.String(), err.Error())
		return nil, err
	}

	// try to get the reserves
	rs, err := contract.GetReserves(nil)
	return &rs, err
}

// UniswapLastKValue returns the last value of the pool control coefficient.
func (ftm *FtmBridge) UniswapLastKValue(pair *common.Address) (hexutil.Big, error) {
	// get the pair contract if possible
	contract, err := contracts.NewUniswapPair(*pair, ftm.eth)
	if err != nil {
		ftm.log.Errorf("Uniswap pair %s not found; %s", pair.String(), err.Error())
		return hexutil.Big{}, err
	}

	// try to get the reserves
	k, err := contract.KLast(nil)
	if err != nil {
		ftm.log.Errorf("Uniswap pair %s coefficient K not available; %s", pair.String(), err.Error())
		return hexutil.Big{}, err
	}

	return hexutil.Big(*k), nil
}

// UniswapPairContract returns instance of this contract according to given pair address
func (ftm *FtmBridge) UniswapPairContract(pairAddres *common.Address) (*contracts.UniswapPair, error) {
	contract, err := contracts.NewUniswapPair(*pairAddres, ftm.eth)
	if err != nil {
		ftm.log.Errorf("Uniswap pair contract %s not found; %s", pairAddres.String(), err.Error())
		return nil, err
	}
	return contract, nil
}

// UniswapFactoryContract returns an instance of an Uniswap factory
func (ftm *FtmBridge) UniswapFactoryContract() (*contracts.UniswapFactory, error) {
	// get the router contract if possible
	contract, err := contracts.NewUniswapFactory(ftm.uniswapConfig.Core, ftm.eth)
	if err != nil {
		ftm.log.Errorf("Uniswap factory contract not found; %s", err.Error())
		return nil, err
	}
	return contract, nil
}
