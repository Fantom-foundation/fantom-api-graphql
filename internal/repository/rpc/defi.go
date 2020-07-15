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

//go:generate abigen --abi ./contracts/defi_rf_aggregator.abi --pkg rpc --type DefiReferenceAggregator --out ./defi_ref_aggregator.go

import (
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"strings"
)

// DefiTokens resolves list of DeFi tokens available for the DeFi functions.
func (ftm *FtmBridge) DefiTokens() ([]types.DefiToken, error) {
	// connect the contract
	contract, err := NewDefiReferenceAggregator(ftm.defiRfAggregatorAddress, ftm.eth)
	if err != nil {
		ftm.log.Errorf("can not open reference aggregator contract connection; %s", err.Error())
		return nil, err
	}

	return ftm.defiTokensList(contract)
}

// decodeToken decodes the contract internal token representation
// into the API structure.
func decodeToken(tk struct {
	Addr       common.Address
	Name       [32]byte
	Symbol     [32]byte
	Decimals   *big.Int
	IsActive   bool
	CanDeposit bool
	CanBorrow  bool
	CanTrade   bool
	Volatility *big.Int
}) (types.DefiToken, error) {
	// make sure the number of decimals is available
	if tk.Decimals == nil {
		return types.DefiToken{}, fmt.Errorf("decimals not available")
	}

	// make sure the volatility is available
	if tk.Volatility == nil {
		return types.DefiToken{}, fmt.Errorf("volatility index not available")
	}

	// decode and return
	return types.DefiToken{
		Address:         tk.Addr,
		Name:            strings.TrimRight(string(tk.Name[:]), "\u0000 "),
		Symbol:          strings.TrimRight(string(tk.Symbol[:]), "\u0000 "),
		Decimals:        int32(tk.Decimals.Uint64()),
		IsActive:        tk.IsActive,
		CanDeposit:      tk.CanDeposit,
		CanBorrow:       tk.CanBorrow,
		CanTrade:        tk.CanTrade,
		VolatilityIndex: hexutil.Big(*tk.Volatility),
	}, nil
}

// defiTokensList load list of DeFi tokens from the smart contract.
func (ftm *FtmBridge) defiTokensList(contract *DefiReferenceAggregator) ([]types.DefiToken, error) {
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
