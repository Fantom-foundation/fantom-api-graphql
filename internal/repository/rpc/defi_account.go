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
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// emptyAddress represents an empty address returned from a contract call.
const emptyAddress = `0x0000000000000000000000000000000000000000`

// DefiAccount loads details of a DeFi account identified by the owner address.
func (ftm *FtmBridge) DefiAccount(owner *common.Address) (*types.DefiAccount, error) {
	// connect the contract
	contract, err := NewDefiLiquidityPool(ftm.defiLiquidityPoolAddress, ftm.eth)
	if err != nil {
		ftm.log.Errorf("can not contact liquidity pool contract; %s", err.Error())
		return nil, err
	}

	// make the container
	da := types.DefiAccount{Address: *owner}

	// load list of collateral tokens
	da.CollateralList, err = ftm.defiAccountTokensList(*owner, contract.CollateralList)
	if err != nil {
		ftm.log.Errorf("collateral tokens list loader failed; %s", err.Error())
		return nil, err
	}

	// load list of debt tokens
	da.DebtList, err = ftm.defiAccountTokensList(*owner, contract.DebtList)
	if err != nil {
		ftm.log.Errorf("debt tokens list loader failed; %s", err.Error())
		return nil, err
	}

	// get the current values of the account tokens on both collateral and debt
	da.CollateralValue, da.DebtValue, err = ftm.defiAccountTokensValue(*owner, contract)
	if err != nil {
		ftm.log.Errorf("can not pull account tokens value; %s", err.Error())
		return nil, err
	}

	// return the account detail
	return &da, nil
}

// defiAccountTokensList loads DeFi tokens list from provided contract array access function.
// It's used to pull collateral tokens, and debt tokens associated with the account.
func (ftm *FtmBridge) defiAccountTokensList(
	owner common.Address,
	loader func(*bind.CallOpts, common.Address, *big.Int) (common.Address, error),
) ([]common.Address, error) {
	// make container
	list := make([]common.Address, 0)

	// setup the initial state and helpers
	tokenFound := true
	index := new(big.Int)
	add := big.NewInt(1)

	// use loader to pull the data
	for tokenFound {
		// get next address if possible
		addr, err := loader(nil, owner, index)
		if err != nil {
			// next address not available
			ftm.log.Debugf("next defi address not available; %s", err.Error())
			break
		}

		// do we have a valid address?
		tokenFound = addr.String() != emptyAddress
		if tokenFound {
			// add the address to the list
			list = append(list, addr)

			// advance the index to get the next address
			index = new(big.Int).Add(index, add)
		}
	}

	return list, nil
}

// DefiAccount loads details of a DeFi account identified by the owner address.
func (ftm *FtmBridge) defiAccountTokensValue(
	owner common.Address,
	contract *DefiLiquidityPool,
) (hexutil.Big, hexutil.Big, error) {
	// get joined collateral value
	cValue, err := contract.CollateralValue(nil, owner)
	if err != nil {
		ftm.log.Errorf("joined collateral value loader failed")
		return hexutil.Big{}, hexutil.Big{}, err
	}

	// get joined debt value
	dValue, err := contract.DebtValue(nil, owner)
	if err != nil {
		ftm.log.Errorf("joined debt value loader failed")
		return hexutil.Big{}, hexutil.Big{}, err
	}

	// return the value we got
	return hexutil.Big(*cValue), hexutil.Big(*dValue), nil
}
