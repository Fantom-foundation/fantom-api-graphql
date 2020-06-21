// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
)

// Vote represents resolvable vote in a ballot.
type Vote struct {
	types.Vote
	repo repository.Repository
}

// NewVote creates a new resolvable ballot structure
func NewVote(vo *types.Vote, repo repository.Repository) *Vote {
	return &Vote{Vote: *vo, repo: repo}
}

// Account resolves the account related to the Vote.
func (vo Vote) Account() (*Account, error) {
	// get the account detail by address
	acc, err := vo.repo.Account(&vo.Voter)
	if err != nil {
		return nil, err
	}

	// return the account detail
	return NewAccount(acc, vo.repo), nil
}

// Votes resolves list of votes for the given voter address
// and list of ballots.
func (rs *rootResolver) Votes(args *struct {
	Voter   common.Address
	Ballots []common.Address
}) ([]Vote, error) {
	// get the list of votes
	votes, err := rs.repo.Votes(args.Voter, args.Ballots)
	if err != nil {
		return nil, err
	}

	// prep the resolvable list of votes
	list := make([]Vote, 0)

	// loop the list to make output
	for _, v := range votes {
		vo := NewVote(&v, rs.repo)
		list = append(list, *vo)
	}

	return list, nil
}
