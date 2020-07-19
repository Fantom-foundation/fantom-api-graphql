/*
Package rpc implements bridge to Lachesis full node API interface.

We recommend using local IPC for fast and the most efficient inter-process communication between the API server
and an Opera/Lachesis node. Any remote RPC connection will work, but the performance may be significantly degraded
by extra networking overhead of remote RPC calls.

You should also consider security implications of opening Lachesis RPC interface for a remote access.
If you considering it as your deployment strategy, you should establish encrypted channel between the API server
and Lachesis RPC interface with connection limited to specified endpoints.

We strongly discourage opening Lachesis RPC interface for unrestricted Internet access.
*/
package rpc

//go:generate abigen --abi ./contracts/defi_oracle_aggregator.abi --pkg rpc --type DefiOracleReferenceAggregator --out ./defi_oracle_aggregator.go

import (
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// DefiToken loads details of a single DeFi token by it's address.
func (ftm *FtmBridge) DefiToken(token *common.Address) (*types.DefiToken, error) {
	// connect the contract
	contract, err := NewDefiOracleReferenceAggregator(ftm.defiRfAggregatorAddress, ftm.eth)
	if err != nil {
		ftm.log.Errorf("can not open reference aggregator contract connection; %s", err.Error())
		return nil, err
	}

	// get the token index
	ix, err := contract.FindTokenIndex(nil, *token)
	if err != nil {
		ftm.log.Errorf("can not get token index by address; %s", err.Error())
		return nil, err
	}

	// if the index valid?
	if 0 > ix.Int64() {
		ftm.log.Debugf("token %s not found", token.String())
		return nil, nil
	}

	// get the token details
	tk, err := contract.Tokens(nil, ix)
	if err != nil {
		ftm.log.Errorf("can not get token %s details for index %d; %s", token.String(), ix.Int64(), err.Error())
		return nil, err
	}

	// decode token details
	dt, err := decodeToken(tk)
	if err != nil {
		ftm.log.Errorf("can not decode token %s details for index %d; %s", token.String(), ix.Int64(), err.Error())
		return nil, err
	}

	return &dt, nil
}

// DefiTokens resolves list of DeFi tokens available for the DeFi functions.
func (ftm *FtmBridge) DefiTokens() ([]types.DefiToken, error) {
	// connect the contract
	contract, err := NewDefiOracleReferenceAggregator(ftm.defiRfAggregatorAddress, ftm.eth)
	if err != nil {
		ftm.log.Errorf("can not open reference aggregator contract connection; %s", err.Error())
		return nil, err
	}

	return ftm.defiTokensList(contract)
}

// decodeToken decodes the contract internal token representation
// into the API structure.
func decodeToken(tk struct {
	Token         common.Address
	Name          string
	Symbol        string
	Logo          string
	Decimals      uint8
	PriceDecimals uint8
	IsActive      bool
	CanDeposit    bool
	CanBorrow     bool
	CanTrade      bool
	Volatility    *big.Int
}) (types.DefiToken, error) {
	// make sure the volatility is available
	if tk.Volatility == nil {
		return types.DefiToken{}, fmt.Errorf("volatility index not available")
	}

	// decode and return
	return types.DefiToken{
		Address:         tk.Token,
		Name:            tk.Name,
		Symbol:          tk.Symbol,
		LogoUrl:         tk.Logo,
		Decimals:        int32(tk.Decimals),
		PriceDecimals:   int32(tk.PriceDecimals),
		IsActive:        tk.IsActive,
		CanDeposit:      tk.CanDeposit,
		CanBorrow:       tk.CanBorrow,
		CanTrade:        tk.CanTrade,
		VolatilityIndex: hexutil.Big(*tk.Volatility),
	}, nil
}

// defiTokensList load list of DeFi tokens from the smart contract.
func (ftm *FtmBridge) defiTokensList(contract *DefiOracleReferenceAggregator) ([]types.DefiToken, error) {
	// get the number of tokens in the reference aggregator
	count, err := contract.TokensCount(nil)
	if err != nil {
		ftm.log.Errorf("can not read aggregator tokens range; %s", err.Error())
		return nil, err
	}

	// make a container for tokens
	list := make([]types.DefiToken, 0)
	index := new(big.Int)

	// load all the tokens in the contract
	for i := uint64(0); i < count.Uint64(); i++ {
		// read the indexed token from contract
		prop, err := contract.Tokens(nil, index.SetUint64(i))
		if err != nil {
			ftm.log.Errorf("can not read token %d details from the contract; %s", i, err.Error())
			return nil, err
		}

		// decode the token
		tk, err := decodeToken(prop)
		if err != nil {
			ftm.log.Errorf("invalid token information received from aggregate for #%d; %s", i, err.Error())
			continue
		}

		// add the token
		list = append(list, tk)
	}

	return list, nil
}

// DefiTokenBalance loads balance of a single DeFi token by it's address.
func (ftm *FtmBridge) DefiTokenBalance(owner *common.Address, token *common.Address, tt string) (hexutil.Big, error) {
	// connect the contract
	contract, err := NewDefiLiquidityPool(ftm.defiLiquidityPoolAddress, ftm.eth)
	if err != nil {
		ftm.log.Errorf("can not open liquidity pool contract connection; %s", err.Error())
		return hexutil.Big{}, err
	}

	// get the collateral token balance
	var val *big.Int

	// pull the right value based to token type
	if tt == "COLLATERAL" {
		val, err = contract.CollateralTokens(nil, *owner, *token)
	} else {
		val, err = contract.DebtTokens(nil, *owner, *token)
	}

	// do we have the value?
	if val == nil {
		ftm.log.Debugf("token %s balance not available for owner %s", token.String(), owner.String())
		return hexutil.Big{}, err
	}

	return hexutil.Big(*val), err
}

// DefiTokenValue loads value of a single DeFi token by it's address in fUSD.
func (ftm *FtmBridge) DefiTokenValue(owner *common.Address, token *common.Address, tt string) (hexutil.Big, error) {
	// get the balance
	balance, err := ftm.DefiTokenBalance(owner, token, tt)
	if err != nil {
		ftm.log.Errorf("can not get token balance; %s", err.Error())
		return hexutil.Big{}, err
	}

	// get the price for the given token from oracle
	val, err := ftm.DefiTokenPrice(token)
	if err != nil {
		ftm.log.Errorf("price not available for token %s; %s", token.String(), err.Error())
		return hexutil.Big{}, err
	}

	// calculate the target value
	value := new(big.Int).Mul(val.ToInt(), balance.ToInt())
	return hexutil.Big(*value), nil
}

// DefiTokenPrice loads the current price of the given token
// from on-chain price oracle.
func (ftm *FtmBridge) DefiTokenPrice(token *common.Address) (hexutil.Big, error) {
	// log actions
	ftm.log.Debugf("connecting price oracle %s", ftm.defiRfAggregatorAddress.String())

	// get the price oracle address
	priceOracle, err := NewDefiOracleReferenceAggregator(ftm.defiRfAggregatorAddress, ftm.eth)
	if err != nil {
		ftm.log.Errorf("can not open price oracle contract connection; %s", err.Error())
		return hexutil.Big{}, err
	}

	// get the price for the given token from oracle
	val, err := priceOracle.GetPrice(nil, *token)
	if err != nil {
		ftm.log.Errorf("price not available for token %s; %s", token.String(), err.Error())
		return hexutil.Big{}, err
	}

	// do we have the value?
	if val == nil {
		ftm.log.Debugf("token %s has not value", token.String())
		return hexutil.Big{}, nil
	}

	return hexutil.Big(*val), nil
}
