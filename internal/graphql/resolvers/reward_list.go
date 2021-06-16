// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// RewardClaimList represents resolvable list of blockchain reward claim edges structure.
type RewardClaimList struct {
	types.RewardClaimsList
}

// RewardClaimListEdge represents a single edge of a reward claim list structure.
type RewardClaimListEdge struct {
	Claim *RewardClaim
}

// NewRewardClaimList builds new resolvable list of reward claims.
func NewRewardClaimList(dl *types.RewardClaimsList) *RewardClaimList {
	return &RewardClaimList{*dl}
}

// TotalCount resolves the total number of delegations in the list.
func (rl *RewardClaimList) TotalCount() hexutil.Uint64 {
	return hexutil.Uint64(rl.Total)
}

// PageInfo resolves the current page information for the reward claims list.
func (rl *RewardClaimList) PageInfo() (*ListPageInfo, error) {
	// do we have any items?
	if rl.Collection == nil || len(rl.Collection) == 0 {
		return NewListPageInfo(nil, nil, false, false)
	}

	// get the first and last elements
	first := Cursor(rl.Collection[0].Pk())
	last := Cursor(rl.Collection[len(rl.Collection)-1].Pk())
	return NewListPageInfo(&first, &last, !rl.IsEnd, !rl.IsStart)
}

// Edges resolves list of reward claim list edges for the linked block list.
func (rl *RewardClaimList) Edges() []*RewardClaimListEdge {
	// do we have any items? return empty list if not
	if rl.Collection == nil || len(rl.Collection) == 0 {
		return make([]*RewardClaimListEdge, 0)
	}

	// make the list
	edges := make([]*RewardClaimListEdge, len(rl.Collection))
	for i, d := range rl.Collection {
		edges[i] = &RewardClaimListEdge{Claim: NewRewardClaim(d)}
	}
	return edges
}

// Cursor generates the list edge cursor.
func (rce *RewardClaimListEdge) Cursor() Cursor {
	return Cursor(rce.Claim.Pk())
}
