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
	types.GovernanceProposalList
}

// GovernanceProposalListEdge represents a single edge of a proposal list structure.
type GovernanceProposalListEdge struct {
	Proposal *GovernanceProposal
	Cursor   Cursor
}

// NewGovernanceProposalList builds new resolvable list of governance proposals.
func NewGovernanceProposalList(gps *types.GovernanceProposalList) *GovernanceProposalList {
	return &GovernanceProposalList{*gps}
}

// GovProposals resolves list of proposals across all the known governance
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
	gcl := make([]common.Address, len(cfg.Governance.Contracts))
	for i, gc := range cfg.Governance.Contracts {
		gcl[i] = gc.Address
	}

	// get the list of all proposals
	list, err := repository.R().GovernanceProposals(gcl, (*string)(args.Cursor), args.Count, args.ActiveOnly)
	if err != nil {
		return nil, err
	}

	return NewGovernanceProposalList(list), nil
}

// TotalCount resolves the total number of governance proposals in the list.
func (gpl *GovernanceProposalList) TotalCount() hexutil.Big {
	val := new(big.Int).SetUint64(gpl.Total)
	return (hexutil.Big)(*val)
}

// PageInfo resolves the current page information for the governance proposal list.
func (gpl *GovernanceProposalList) PageInfo() (*ListPageInfo, error) {
	// do we have any items?
	if gpl.Collection == nil || len(gpl.Collection) == 0 {
		return NewListPageInfo(nil, nil, false, false)
	}

	// get the first and last elements
	first := Cursor(gpl.First.String())
	last := Cursor(gpl.Last.String())
	return NewListPageInfo(&first, &last, !gpl.IsEnd, !gpl.IsStart)
}

// Edges resolves list of edges for the linked governance proposal list.
func (gpl *GovernanceProposalList) Edges() []*GovernanceProposalListEdge {
	// do we have any items? return empty list if not
	if gpl.Collection == nil || len(gpl.Collection) == 0 {
		return make([]*GovernanceProposalListEdge, 0)
	}

	// make the list
	edges := make([]*GovernanceProposalListEdge, len(gpl.Collection))
	for i, gp := range gpl.Collection {
		edges[i] = &GovernanceProposalListEdge{
			Proposal: NewGovernanceProposal(gp),
			Cursor:   Cursor(gp.Contract.String()),
		}
	}

	return edges
}
