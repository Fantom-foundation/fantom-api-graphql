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
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

//go:generate abigen --abi ./contracts/erc20.abi --pkg rpc --type ERCTwenty --out ./erc20_token.go

// Erc20Balance loads the current available balance of and ERC20 token identified by the token
// contract address for an identified owner address.
func (ftm *FtmBridge) Erc20Balance(owner *common.Address, token *common.Address) (hexutil.Big, error) {
	// connect the contract
	contract, err := NewERCTwenty(*token, ftm.eth)
	if err != nil {
		ftm.log.Errorf("can not contact ERC20 contract; %s", err.Error())
		return hexutil.Big{}, err
	}

	// get the balance
	val, err := contract.BalanceOf(nil, *owner)
	if err != nil {
		ftm.log.Errorf("can not ERC20 %s balance for %s; %s", token.String(), owner.String(), err.Error())
		return hexutil.Big{}, err
	}

	// make sur we always have a value; at least zero
	// this should always be the case since the contract should
	// return zero even for unknown owners, but let's be sure here
	if val == nil {
		val = new(big.Int)
	}

	// return the account balance
	return hexutil.Big(*val), nil
}

// Erc20Allowance loads the current amount of ERC20 tokens unlocked for DeFi
// contract by the token owner.
func (ftm *FtmBridge) Erc20Allowance(owner *common.Address, token *common.Address) (hexutil.Big, error) {
	// connect the contract
	contract, err := NewERCTwenty(*token, ftm.eth)
	if err != nil {
		ftm.log.Errorf("can not contact ERC20 contract; %s", err.Error())
		return hexutil.Big{}, err
	}

	// get the amount of tokens allowed for DeFi
	val, err := contract.Allowance(nil, *owner, ftm.defiLiquidityPoolAddress)
	if err != nil {
		ftm.log.Errorf("can not get defi ERC20 %s allowance for %s; %s", token.String(), owner.String(), err.Error())
		return hexutil.Big{}, err
	}

	// make sur we always have a value; at least zero
	// this should always be the case since the contract should
	// return zero even for unknown owners, but let's be sure here
	if val == nil {
		val = new(big.Int)
	}

	// return the account balance
	return hexutil.Big(*val), nil
}
