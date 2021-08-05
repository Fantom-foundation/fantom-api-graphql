package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// EpochList represents resolvable list of epoch edges structure.
type EpochList struct {
	types.EpochList
}

// EpochListEdge represents a single edge of an epoch list list structure.
type EpochListEdge struct {
	Epoch *Epoch
}

// NewEpochList builds new resolvable list of epochs.
func NewEpochList(el *types.EpochList) *EpochList {
	return &EpochList{EpochList: *el}
}

// Epochs resolves a list of epochs for the given cursor and count.
func (rs *rootResolver) Epochs(args struct {
	Cursor *Cursor
	Count  int32
}) (*EpochList, error) {
	// limit query size; the count can be either positive or negative
	// this controls the loading direction
	args.Count = listLimitCount(args.Count, listMaxEdgesPerRequest)

	// get the transaction hash list from repository
	epl, err := repository.R().Epochs((*string)(args.Cursor), args.Count)
	if err != nil {
		log.Errorf("can not get epoch list; %s", err.Error())
		return nil, err
	}
	return NewEpochList(epl), nil
}

// TotalCount resolves the total number of epochs in the list.
func (el *EpochList) TotalCount() (hexutil.Uint64, error) {
	return repository.R().CurrentEpoch()
}

// PageInfo resolves the current page information for the epoch list.
func (el *EpochList) PageInfo() (*ListPageInfo, error) {
	// do we have any items?
	if el.Collection == nil || len(el.Collection) == 0 {
		return NewListPageInfo(nil, nil, false, false)
	}

	// get the first and last elements
	first := Cursor(el.Collection[0].Id.String())
	last := Cursor(el.Collection[len(el.Collection)-1].Id.String())
	return NewListPageInfo(&first, &last, !el.IsEnd, !el.IsStart)
}

// Edges resolves list of edges for the epochs list.
func (el *EpochList) Edges() []*EpochListEdge {
	// do we have any items? return empty list if not
	if el.Collection == nil || len(el.Collection) == 0 {
		return make([]*EpochListEdge, 0)
	}

	// make the list
	edges := make([]*EpochListEdge, len(el.Collection))
	for i, e := range el.Collection {
		edges[i] = &EpochListEdge{
			Epoch: &Epoch{Epoch: *e},
		}
	}
	return edges
}

// Cursor resolves a cursor of an edge in the edges list.
func (ele *EpochListEdge) Cursor() Cursor {
	return Cursor(ele.Epoch.Id.String())
}
