// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"strconv"
)

// ContractList represents resolvable list of blockchain smart contract edges structure.
type ContractList struct {
	repo repository.Repository
	list *types.ContractList
}

// ContractListEdge represents a single edge of a smart contract list structure.
type ContractListEdge struct {
	Contract *Contract
	Cursor   Cursor
}

// NewContractList builds new resolvable list of smart contracts.
func NewContractList(cl *types.ContractList, repo repository.Repository) *ContractList {
	return &ContractList{
		repo: repo,
		list: cl,
	}
}

// Contracts resolves list of blockchain smart contracts encapsulated in a listable structure.
func (rs *rootResolver) Contracts(args *struct {
	ValidatedOnly bool
	Cursor        *Cursor
	Count         int32
}) (*ContractList, error) {
	// limit query size; the count can be either positive or negative
	// this controls the loading direction
	args.Count = listLimitCount(args.Count, listMaxEdgesPerRequest)

	// get the contract list from repository
	cl, err := rs.repo.Contracts(args.ValidatedOnly, (*string)(args.Cursor), args.Count)
	if err != nil {
		rs.log.Errorf("can not get contracts list; %s", err.Error())
		return nil, err
	}

	return NewContractList(cl, rs.repo), nil
}

// TotalCount resolves the total number of smart contracts in the list.
func (cl *ContractList) TotalCount() hexutil.Big {
	val := (*hexutil.Big)(big.NewInt(int64(cl.list.Total)))
	return *val
}

// PageInfo resolves the current page information for the smart contract list.
func (cl *ContractList) PageInfo() (*ListPageInfo, error) {
	// do we have any items?
	if cl.list == nil || cl.list.Collection == nil || len(cl.list.Collection) == 0 {
		return NewListPageInfo(nil, nil, false, false)
	}

	// get the first and last elements
	first := Cursor(strconv.FormatUint(cl.list.First, 10))
	last := Cursor(strconv.FormatUint(cl.list.Last, 10))
	return NewListPageInfo(&first, &last, !cl.list.IsEnd, !cl.list.IsStart)
}

// Edges resolves list of edges for the linked smart contract list.
func (cl *ContractList) Edges() []*ContractListEdge {
	// do we have any items? return empty list if not
	if cl.list == nil || cl.list.Collection == nil || len(cl.list.Collection) == 0 {
		return make([]*ContractListEdge, 0)
	}

	// make the list
	edges := make([]*ContractListEdge, len(cl.list.Collection))
	for i, c := range cl.list.Collection {
		// make the contract ref
		ct := NewContract(c, cl.repo)

		// make the element
		edge := ContractListEdge{
			Contract: ct,
			Cursor:   Cursor(strconv.FormatUint(c.OrdinalIndex, 10)),
		}

		// add it to the list
		edges[i] = &edge
	}

	return edges
}
