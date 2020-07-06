/*
Package repository implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Opera/Lachesis full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import (
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
)

// ballotsListExtraCount represents the number ballots we will pull extra
// from the database to compensate lack of finalization info in there.
const ballotsListExtraCount = 10

// Ballots returns list of ballots at Opera blockchain.
func (p *proxy) Ballots(cursor *string, count int32) (*types.BallotList, error) {
	return p.db.Ballots(cursor, count)
}

// BallotIsFinalized returns the finalized status of the ballot by it's address.
func (p *proxy) BallotIsFinalized(addr *common.Address) (bool, error) {
	return p.rpc.BallotIsFinalized(addr)
}

// BallotWinner returns the winning proposal index, or nil if not decided.
func (p *proxy) BallotWinner(addr *common.Address) (*uint64, error) {
	return p.rpc.BallotWinner(addr)
}

// BallotByAddress returns a ballot information by the contract address.
func (p *proxy) BallotByAddress(addr *common.Address) (*types.Ballot, error) {
	return p.db.Ballot(addr)
}

// Votes returns a list of votes for the given votes and list of ballots.
func (p *proxy) Votes(voter common.Address, ballots []common.Address) ([]types.Vote, error) {
	return p.rpc.Votes(voter, ballots)
}

// BallotsActive returns a list of at most <count> currently active Ballots.
func (p *proxy) BallotsActive(count uint32) ([]types.Ballot, error) {
	return p.db.BallotsActive(count)
}

// BallotsClosed returns a list of <count> recently closed Ballots.
// If the finalized is set to false, the list contains Ballots waiting
// to be resolved.
func (p *proxy) BallotsClosed(finalized bool, count uint32) ([]types.Ballot, error) {
	// get a list of recently closed ballots
	list, err := p.db.BallotsClosed(count + ballotsListExtraCount)
	if err != nil {
		p.log.Errorf("can not pull the ballots list from database; %s", err.Error())
		return nil, err
	}

	// loop the list and filter out what's not needed
	// work over the underlying array to reduce filtering memory footprint
	// and avoid unnecessary allocations of new Ballot records
	var listIndex int
	for _, bal := range list {
		// check if the ballot is finalized
		isFin, err := p.rpc.BallotIsFinalized(&bal.Address)
		if err != nil {
			p.log.Errorf("can not check ballot [%s] finalization status; %s", bal.Address.String(), err.Error())
			return nil, err
		}

		// decide if this ballot will be kept.
		// if it doesn't match the filter, skip it
		if (isFin && !finalized) || (!isFin && finalized) {
			continue
		}

		// reassign back to the slice
		list[listIndex] = bal
		listIndex++

		// break when the <count> item is reached
		if listIndex >= int(count) {
			break
		}
	}

	// return only the subset
	return list[:listIndex], nil
}
