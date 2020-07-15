// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import "fantom-api-graphql/internal/types"

// DefiTokens resolves list of DeFi tokens available for the DeFi functions.
func (rs *rootResolver) DefiTokens() ([]types.DefiToken, error) {
	// pass the call to repository
	return rs.repo.DefiTokens()
}
