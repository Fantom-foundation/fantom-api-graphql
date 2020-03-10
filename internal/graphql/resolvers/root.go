// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/logger"
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// ApiResolver represents the API interface expected to handle API access points
type ApiResolver interface {
	// Account resolves blockchain account by address.
	Account(struct{ Address common.Address }) (*Account, error)

	// Block resolves blockchain block by number or by hash. If neither is provided, the most recent block is given.
	Block(*struct {
		Number *hexutil.Uint64
		Hash   *types.Hash
	}) (*Block, error)

	// Transaction resolves blockchain transaction by hash.
	Transaction(*struct{ Hash types.Hash }) (*Transaction, error)
}

// rootResolver represents the ApiResolver implementation.
type rootResolver struct {
	log  logger.Logger
	repo repository.Repository
}

// New creates a new root resolver instance and initializes it's internal structure.
func New(log logger.Logger, repo repository.Repository) ApiResolver {
	return &rootResolver{log: log, repo: repo}
}
