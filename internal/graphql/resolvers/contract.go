// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
)

// Contract represents resolvable blockchain smart contract structure.
type Contract struct {
	repo repository.Repository
	types.Contract
}

// NewContract builds new resolvable smart contract structure.
func NewContract(con *types.Contract, repo repository.Repository) *Contract {
	return &Contract{
		repo:     repo,
		Contract: *con,
	}
}

// DeployedBy resolves the deployment transaction of the contract.
func (con *Contract) DeployedBy() (*Transaction, error) {
	tr, err := con.repo.Transaction(&con.TransactionHash)
	if err != nil {
		return nil, err
	}

	// make the resolvable transaction object
	trx := NewTransaction(tr, con.repo)
	return trx, nil
}
