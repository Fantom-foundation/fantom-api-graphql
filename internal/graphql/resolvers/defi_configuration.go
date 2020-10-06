// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/config"
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
)

// DefiConfiguration represents a resolvable DeFi Configuration instance.
type DefiConfiguration struct {
	repo repository.Repository
	cfg  *config.Config
	types.DefiSettings
}

// NewDefiConfiguration creates a new instance of resolvable DeFi token.
func NewDefiConfiguration(cf *types.DefiSettings, cfg *config.Config, repo repository.Repository) *DefiConfiguration {
	return &DefiConfiguration{
		repo:         repo,
		cfg:          cfg,
		DefiSettings: *cf,
	}
}

// DefiConfiguration resolves the current DeFi contract settings.
func (rs *rootResolver) DefiConfiguration() (*DefiConfiguration, error) {
	// pass the call to repository
	st, err := rs.repo.DefiConfiguration()
	if err != nil {
		return nil, err
	}

	return NewDefiConfiguration(st, rs.cfg, rs.repo), nil
}

// UniswapCoreFactory returns the address of the Uniswap factory contract
// from the app configuration.
func (dfc *DefiConfiguration) UniswapCoreFactory() common.Address {
	return common.HexToAddress(dfc.cfg.DefiUniswapCore)
}

// UniswapRouter returns the address of the Uniswap router contract
// from the app configuration.
func (dfc *DefiConfiguration) UniswapRouter() common.Address {
	return common.HexToAddress(dfc.cfg.DefiUniswapRouter)
}
