package repository

import "fantom-api-graphql/internal/types"

// DefiTokens resolves list of DeFi tokens available for the DeFi functions.
func (p *proxy) DefiTokens() ([]types.DefiToken, error) {
	return p.rpc.DefiTokens()
}
