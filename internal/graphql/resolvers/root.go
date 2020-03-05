// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import "fantom-api-graphql/internal/logger"

// RootResolver defines the root resolver to be used to handle API entry points.
type RootResolver struct {
	log  logger.Logger
	repo repository.Repository
}

// New creates a new root resolver instance and initializes it's internal structure.
func New(l logger.Logger, r repository.Repository) *RootResolver {
	return &RootResolver{log: l, repo: r}
}
