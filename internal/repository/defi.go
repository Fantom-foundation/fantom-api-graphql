package repository

import (
	"fantom-api-graphql/internal/repository/rpc"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// DefiAccount loads details of a DeFi account identified by the owner address.
func (p *proxy) DefiAccount(owner common.Address) (*types.DefiAccount, error) {
	return p.rpc.DefiAccount(&owner)
}

// DefiToken loads details of a single DeFi token by it's address.
func (p *proxy) DefiToken(token *common.Address) (*types.DefiToken, error) {
	return p.rpc.DefiToken(token)
}

// DefiTokens resolves list of DeFi tokens available for the DeFi functions.
func (p *proxy) DefiTokens() ([]types.DefiToken, error) {
	return p.rpc.DefiTokens()
}

// DefiConfiguration resolves the current DeFi contract settings.
func (p *proxy) DefiConfiguration() (*types.DefiSettings, error) {
	return p.rpc.DefiConfiguration()
}

// DefiTokenBalance loads balance of a single DeFi token by it's address.
func (p *proxy) DefiTokenBalance(owner *common.Address, token *common.Address, tt string) (hexutil.Big, error) {
	return p.rpc.DefiTokenBalance(owner, token, tt)
}

// DefiTokenValue loads value of a single DeFi token by it's address in fUSD.
func (p *proxy) DefiTokenValue(owner *common.Address, token *common.Address, tt string) (hexutil.Big, error) {
	return p.rpc.DefiTokenValue(owner, token, tt)
}

// DefiTokenPrice loads the current price of the given token
// from on-chain price oracle.
func (p *proxy) DefiTokenPrice(token *common.Address) (hexutil.Big, error) {
	return p.rpc.DefiTokenPrice(token)
}

// Erc20Balance load the current available balance of and ERC20 token identified by the token
// contract address for an identified owner address.
func (p *proxy) Erc20Balance(owner *common.Address, token *common.Address) (hexutil.Big, error) {
	// make sure to treat native FTM tokens differently
	if token.String() == rpc.NativeTokenAddress {
		val, err := p.rpc.AccountBalance(owner)
		if err != nil {
			return hexutil.Big{}, err
		}
		return *val, nil
	}

	// handle ERC20 token balance
	return p.rpc.Erc20Balance(owner, token)
}

// Erc20Allowance loads the current amount of ERC20 tokens unlocked for DeFi
// contract by the token owner.
func (p *proxy) Erc20Allowance(owner *common.Address, token *common.Address) (hexutil.Big, error) {
	// if this is native FTM token, all available tokens
	// are allowed to be handled by DeFi calls (if supported by DeFi)
	if token.String() == rpc.NativeTokenAddress {
		return p.Erc20Balance(owner, token)
	}

	// handle ERC20 token allowance
	return p.rpc.Erc20Allowance(owner, token)
}
