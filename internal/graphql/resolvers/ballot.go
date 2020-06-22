// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"time"
)

// Ballot represents resolvable ballot record.
type Ballot struct {
	types.Ballot
	isFinalized *bool
	repo        repository.Repository
}

// NewBallot creates a new resolvable ballot structure
func NewBallot(bt *types.Ballot, repo repository.Repository) *Ballot {
	return &Ballot{Ballot: *bt, repo: repo}
}

// Ballot resolves details of a voting ballot by it's address.
func (rs *rootResolver) Ballot(args *struct{ Address common.Address }) (*Ballot, error) {
	// get the ballot for the address
	bl, err := rs.repo.BallotByAddress(&args.Address)
	if err != nil {
		return nil, err
	}

	// return the resolved ballot record
	return NewBallot(bl, rs.repo), nil
}

// IsOpen resolves open status of a ballot.
func (bt *Ballot) IsOpen() bool {
	// get the current time
	now := time.Now()
	return int64(bt.Start) < now.Unix() && int64(bt.End) > now.Unix()
}

// IsFinalized resolves finalized status of a ballot.
// If a ballot has been finalized, the winner is known.
func (bt *Ballot) IsFinalized() (bool, error) {
	// do we have the value pre-cached?
	if bt.isFinalized == nil {
		fin, err := bt.repo.BallotIsFinalized(&bt.Address)
		if err != nil {
			return false, err
		}

		// rewrite the value
		bt.isFinalized = &fin
	}

	return *bt.isFinalized, nil
}

// Winner resolves the winning proposal index, if the ballot
// has been finalized already. Returns nil if the ballot is pending.
func (bt *Ballot) Winner() (*hexutil.Uint64, error) {
	// make sure the ballot is already finalized
	fin, err := bt.IsFinalized()
	if err != nil {
		return nil, err
	}

	// if the ballot is not resolved, do not dare
	if !fin {
		return nil, nil
	}

	// check the ballot winner
	return bt.repo.BallotWinner(&bt.Address)
}

// Contract resolves managing smart contract of the ballot.
func (bt *Ballot) Contract() (*Contract, error) {
	// get a contract to be validated if any
	sc, err := bt.repo.Contract(&bt.Address)
	if err != nil {
		return nil, err
	}

	return NewContract(sc, bt.repo), nil
}
