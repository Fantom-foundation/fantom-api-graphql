package repository

import (
	"fantom-api-graphql/internal/config"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Erc20Token returns an ERC20 token rfor the given address, if available.
func (p *proxy) Erc20Token(addr *common.Address) (*types.Erc20Token, error) {
	// try the cache first
	token := p.cache.PullErc20Token(addr)
	if token != nil {
		p.log.Debugf("found cached token %s at %s", token.Symbol, token.Address.String())
		return token, nil
	}

	// load the slow way; build the structure and pull needed details
	var err error
	token, err = p.loadErc20TokenDetails(&types.Erc20Token{Address: *addr})
	if err != nil {
		p.log.Errorf("can not load ERC20 token at %s; %s", addr.String(), err.Error())
		return nil, err
	}

	// store to cache and return the result
	p.log.Debugf("found %s at %s", token.Symbol, token.Address.String())
	err = p.cache.PushErc20Token(token)
	if err != nil {
		p.log.Errorf("can not keep ERC20 token %s in cache; %s", addr.String(), err.Error())
	}

	return token, nil
}

// loadErc20TokenDetails loads details of the given ERC20 token using ERC20
// contract calls.
func (p *proxy) loadErc20TokenDetails(token *types.Erc20Token) (*types.Erc20Token, error) {
	var err error

	// get the name
	token.Name, err = p.Erc20Name(&token.Address)
	if err != nil {
		p.log.Errorf("ERC20 token not recognized at %s; %s", token.Address.String(), err.Error())
		return nil, err
	}

	// get symbol
	token.Symbol, err = p.Erc20Symbol(&token.Address)
	if err != nil {
		p.log.Errorf("ERC20 token not recognized at %s; %s", token.Address.String(), err.Error())
		return nil, err
	}

	// get decimals
	token.Decimals, err = p.Erc20Decimals(&token.Address)
	if err != nil {
		p.log.Errorf("ERC20 token not recognized at %s; %s", token.Address.String(), err.Error())
		return nil, err
	}

	return token, nil
}

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

// Erc20LogoURL provides URL address of a logo of the ERC20 token.
func (p *proxy) Erc20LogoURL(addr *common.Address) string {
	// do we know the token?
	logo, ok := p.cfg.TokenLogo[*addr]
	if !ok {
		logo = p.cfg.TokenLogo[common.HexToAddress(config.EmptyAddress)]
	}
	return logo
}
