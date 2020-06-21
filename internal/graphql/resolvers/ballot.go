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
	repo repository.Repository
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
	return bt.repo.BallotIsFinalized(&bt.Address)
}

// Winner resolves the winning proposal index, if the ballot
// has been finalized already. Returns nil if the ballot is pending.
func (bt *Ballot) Winner() (*hexutil.Uint64, error) {
	return bt.repo.BallotWinner(&bt.Address)
}
