package repository

import (
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
