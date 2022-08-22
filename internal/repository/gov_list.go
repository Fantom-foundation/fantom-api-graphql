/*
Package repository implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Opera full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import (
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"sort"
	"time"
)

// GovernanceProposalsByStart represents a list of governance proposals
// sortable by their voting start.
type GovernanceProposalsByStart []*types.GovernanceProposal

// Len returns size of the proposals list.
func (s GovernanceProposalsByStart) Len() int {
	return len(s)
}

// Less compares two proposals and returns true if the first is lower than the last.
// We use it to sort proposals by the starting date.
func (s GovernanceProposalsByStart) Less(i, j int) bool {
	return uint64(s[i].VotingStarts) > uint64(s[j].VotingStarts)
}

// Swap changes position of two proposals in the list.
func (s GovernanceProposalsByStart) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// GovernanceProposalsBy loads list of proposals of the given Governance contract.
func (p *proxy) GovernanceProposalsBy(gov common.Address) ([]*types.GovernanceProposal, error) {
	return p.rpc.GovernanceProposalsBy(gov)
}

// GovernanceProposals loads list of proposals from given set of Governance contracts.
func (p *proxy) GovernanceProposals(govs []common.Address, cursor *string, count int32, activeOnly bool) (*types.GovernanceProposalList, error) {
	// prep the container
	var result []*types.GovernanceProposal

	// loop all the governance contracts queried
	for _, addr := range govs {
		// get the list
		gpl, err := p.rpc.GovernanceProposalsBy(addr)
		if err != nil {
			p.log.Errorf("list of proposals not available for %s; %s", addr.String(), err.Error())
			return nil, err
		}

		// merge to output
		if result == nil {
			// make this list the output
			result = gpl
		} else {
			// append proposals to the result set
			result = append(result, gpl...)
		}
	}

	// do we have any proposals? if not, make an empty response
	if result == nil {
		result = make([]*types.GovernanceProposal, 0)
	}

	// sort the response
	sort.Sort(GovernanceProposalsByStart(result))

	// log the action
	p.log.Debugf("filtering governance proposals")
	return filterGovernanceProposals(result, cursor, count, activeOnly), nil
}

// filterActiveGovernanceProposals returns only active proposals from the input list.
func filterActiveGovernanceProposals(list []*types.GovernanceProposal) []*types.GovernanceProposal {
	// make a new list
	result := make([]*types.GovernanceProposal, 0)
	now := time.Now().UTC().Unix()

	// loop the input list
	for _, gp := range list {
		// we decide just by dates range here
		if int64(gp.VotingStarts) < now && int64(gp.VotingMustEnd) > now {
			result = append(result, gp)
		}
	}

	return result
}

// findGovernanceProposal scans the list of proposals and return index
// of the one specified by the given cursor
func findGovernanceProposal(list []*types.GovernanceProposal, cursor string) int64 {
	// loop the list
	for i, gp := range list {
		if gp.Contract.String() == cursor {
			return int64(i)
		}
	}
	return -1
}

// governanceProposalsListStart determines where to start the list from given criteria.
func governanceProposalsListStart(list []*types.GovernanceProposal, cursor *string, count int32) (int64, bool) {
	var startAt int64
	var down = true

	// try to find the proposal identified by the cursor if any
	if cursor != nil {
		startAt = findGovernanceProposal(list, *cursor)
	}

	// no starting point? assume to start from top
	if startAt < 0 || cursor == nil {
		startAt = 0
	}

	// what if the direction of counting is reversed?
	if count < 0 {
		if startAt == 0 || cursor == nil {
			startAt = int64(len(list) - 1)
		}
		down = false
	}

	return startAt, down
}

// filterGovernanceProposals filters the list of governance proposals to match the filters given.
func filterGovernanceProposals(list []*types.GovernanceProposal, cursor *string, count int32, activeOnly bool) *types.GovernanceProposalList {
	// get the active only if requested
	if activeOnly {
		list = filterActiveGovernanceProposals(list)
	}

	// an empty list?
	if 0 == len(list) {
		return &types.GovernanceProposalList{
			Collection: list,
			Total:      0,
			First:      common.Address{},
			Last:       common.Address{},
			IsStart:    true,
			IsEnd:      true,
		}
	}

	// calculate the start and count
	startAt, down := governanceProposalsListStart(list, cursor, count)
	endAt := startAt + int64(count)
	if down {
		if endAt > int64(len(list)) {
			endAt = int64(len(list))
		}
	} else {
		if endAt < 0 {
			endAt = 0
		}
		startAt, endAt = endAt, startAt
	}

	// return the proposals list details
	return &types.GovernanceProposalList{
		Collection: list[startAt:endAt],
		Total:      uint64(len(list)),
		First:      list[0].Contract,
		Last:       list[len(list)-1].Contract,
		IsStart:    startAt == 0,
		IsEnd:      endAt == int64(len(list)),
	}
}
