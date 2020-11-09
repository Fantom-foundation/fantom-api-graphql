// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// GovernanceProposalList represents resolvable list of governance
// proposal edges structure.
type GovernanceProposalList struct {
	repo       repository.Repository
	list       *types.GovernanceProposalList
	TotalCount hexutil.Big
}

// GovernanceProposalListEdge represents a single edge of a proposal list structure.
type GovernanceProposalListEdge struct {
	Proposal *GovernanceProposal
	Cursor   Cursor
}

// NewGovernanceProposalList builds new resolvable list of governance proposals.
func NewGovernanceProposalList(gps *types.GovernanceProposalList, repo repository.Repository) *GovernanceProposalList {
	return &GovernanceProposalList{
		repo:       repo,
		list:       gps,
		TotalCount: hexutil.Big(*big.NewInt(int64(gps.Total))),
	}
}

// Proposals resolves list of proposals across all the known governance
// contracts in a browsable structure.
func (rs *rootResolver) GovProposals(args struct {
	Cursor     *Cursor
	Count      int32
	ActiveOnly bool
}) (*GovernanceProposalList, error) {
	// limit query size; the count can be either positive or negative
	// this controls the loading direction
	args.Count = listLimitCount(args.Count, listMaxEdgesPerRequest)

	// prep list of governance contracts we are interested in
	gcl := make([]*common.Address, len(rs.cfg.Governance.Contracts))
	for i, gc := range rs.cfg.Governance.Contracts {
		gcl[i] = &gc.Address
	}

	// get the list of all proposals
	list, err := rs.repo.GovernanceProposals(gcl, (*string)(args.Cursor), args.Count, args.ActiveOnly)
	if err != nil {
		return nil, err
	}

	return NewGovernanceProposalList(list, rs.repo), nil
}

// PageInfo resolves the current page information for the governance proposal list.
func (gpl *GovernanceProposalList) PageInfo() (*ListPageInfo, error) {
	// do we have any items?
	if gpl.list == nil || gpl.list.Collection == nil || len(gpl.list.Collection) == 0 {
		return NewListPageInfo(nil, nil, false, false)
	}

	// get the first and last elements
	first := Cursor(gpl.list.First.String())
	last := Cursor(gpl.list.Last.String())
	return NewListPageInfo(&first, &last, !gpl.list.IsEnd, !gpl.list.IsStart)
}

// Edges resolves list of edges for the linked governance proposal list.
func (gpl *GovernanceProposalList) Edges() []*GovernanceProposalListEdge {
	// do we have any items? return empty list if not
	if gpl.list == nil || gpl.list.Collection == nil || len(gpl.list.Collection) == 0 {
		return make([]*GovernanceProposalListEdge, 0)
	}

	// make the list
	edges := make([]*GovernanceProposalListEdge, len(gpl.list.Collection))
	for i, b := range gpl.list.Collection {
		// make the element
		edge := GovernanceProposalListEdge{
			Proposal: &GovernanceProposal{
				repo:               gpl.repo,
				GovernanceProposal: *b,
			},
			Cursor: Cursor(b.Contract.String()),
		}

		// add it to the list
		edges[i] = &edge
	}

	return edges
}
