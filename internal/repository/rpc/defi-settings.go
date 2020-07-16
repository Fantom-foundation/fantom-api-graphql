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

import (
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

//go:generate abigen --abi ./contracts/defi_liquidity_pool.abi --pkg rpc --type DefiLiquidityPool --out ./defi_liquidity_pool.go --alias _collateralValue=StoredCollateralValue,_debtValue=StoredDebtValue

// tConfigItemsLoaders defines a map between DeFi config elements and their respective loaders.
type tConfigItemsLoaders map[*hexutil.Big]func(*bind.CallOpts) (*big.Int, error)

// DefiConfiguration resolves the current DeFi contract settings.
func (ftm *FtmBridge) DefiConfiguration() (*types.DefiSettings, error) {
	// connect the contract
	contract, err := NewDefiLiquidityPool(ftm.defiLiquidityPoolAddress, ftm.eth)
	if err != nil {
		ftm.log.Errorf("can not open liquidity pool contract connection; %s", err.Error())
		return nil, err
	}

	// create the container
	ds := types.DefiSettings{}
	loaders := tConfigItemsLoaders{
		&ds.TradeFee4:               contract.TradeFee4dec,
		&ds.LoanFee4:                contract.LoanEntryFee4dec,
		&ds.MinCollateralRatio4:     contract.ColLowestRatio4dec,
		&ds.WarningCollateralRatio4: contract.ColWarningRatio4dec,
		&ds.LiqCollateralRatio4:     contract.ColLiquidationRatio4dec,
	}

	// load all the configured values
	if err := ftm.pullSetOfDefiConfigValues(loaders); err != nil {
		ftm.log.Errorf("can not pull defi config values; %s", err.Error())
		return nil, err
	}

	// load the decimals correction
	if ds.Decimals, err = ftm.pullDefiDecimalCorrection(contract); err != nil {
		ftm.log.Errorf("can not pull defi decimals correction; %s", err.Error())
		return nil, err
	}

	// load the decimals correction
	if ds.PriceOracleAggregate, err = ftm.pullDefiPriceAggregatorAddress(contract); err != nil {
		ftm.log.Errorf("can not pull defi price oracle aggregate address; %s", err.Error())
		return nil, err
	}

	// number of decimals
	return &ds, nil
}

// pullSetOfDefiConfigValues pulls set of DeFi configuration values for the given
// config loaders map.
func (ftm *FtmBridge) pullDefiDecimalCorrection(con *DefiLiquidityPool) (int32, error) {
	// load the decimals correction
	val, err := ftm.pullDefiConfigValue(con.RatioDecimalsCorrection)
	if err != nil {
		ftm.log.Errorf("can not pull decimals correction; %s", err.Error())
		return 0, err
	}

	// calculate number of decimals
	var dec int32
	var value = val.ToInt().Uint64()
	for value > 1 {
		value /= 10
		dec++
	}

	// convert and return
	return dec, nil
}

// pullDefiPriceAggregatorAddress pulls address of the price oracle aggregator contract used by DeFi.
func (ftm *FtmBridge) pullDefiPriceAggregatorAddress(con *DefiLiquidityPool) (common.Address, error) {
	// load the decimals correction
	addr, err := con.PriceOracle(nil)
	if err != nil {
		ftm.log.Errorf("can not pull price oracle aggregator address; %s", err.Error())
		return common.Address{}, err
	}

	// convert and return
	return addr, nil
}

// pullSetOfDefiConfigValues pulls set of DeFi configuration values for the given
// config loaders map.
func (ftm *FtmBridge) pullSetOfDefiConfigValues(loaders tConfigItemsLoaders) error {
	// collect loaders error
	var err error

	// loop the map and load the values
	for ref, fn := range loaders {
		*ref, err = ftm.pullDefiConfigValue(fn)
		if err != nil {
			return err
		}
	}

	return nil
}

// tradeFee4 pulls DeFi trading fee from the Liquidity Pool contract.
func (ftm *FtmBridge) pullDefiConfigValue(cf func(*bind.CallOpts) (*big.Int, error)) (hexutil.Big, error) {
	// pull the trading fee value
	val, err := cf(nil)
	if err != nil {
		return hexutil.Big{}, err
	}

	// do we have the value? we should always have
	if val == nil {
		return hexutil.Big{}, fmt.Errorf("defi config value not available")
	}

	return hexutil.Big(*val), nil
}
