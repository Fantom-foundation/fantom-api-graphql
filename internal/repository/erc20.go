package repository

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Erc20Name provides information about the name of the ERC20 token.
func (p *proxy) Erc20Name(token *common.Address) (string, error) {
	return p.rpc.Erc20Name(token)
}

// Erc20Symbol provides information about the symbol of the ERC20 token.
func (p *proxy) Erc20Symbol(token *common.Address) (string, error) {
	return p.rpc.Erc20Symbol(token)
}

// Erc20Name provides information about the decimals of the ERC20 token.
func (p *proxy) Erc20Decimals(token *common.Address) (int32, error) {
	return p.rpc.Erc20Decimals(token)
}

// Erc20BalanceOf load the current available balance of and ERC20 token identified by the token
// contract address for an identified owner address.
func (p *proxy) Erc20BalanceOf(token *common.Address, owner *common.Address) (hexutil.Big, error) {
	return p.rpc.Erc20BalanceOf(token, owner)
}

// Erc20Allowance loads the current amount of ERC20 tokens unlocked for DeFi
// contract by the token owner.
func (p *proxy) Erc20Allowance(token *common.Address, owner *common.Address, spender *common.Address) (hexutil.Big, error) {
	return p.rpc.Erc20Allowance(token, owner, spender)
}

// Erc20TotalSupply provides information about all available tokens
func (p *proxy) Erc20TotalSupply(token *common.Address) (hexutil.Big, error) {
	return p.rpc.Erc20TotalSupply(token)
}

// Erc20TokensList returns a list of known ERC20 tokens ordered by their activity.
func (p *proxy) Erc20TokensList(count int32) ([]common.Address, error) {
	return p.db.Erc20TokensList(count)
}
