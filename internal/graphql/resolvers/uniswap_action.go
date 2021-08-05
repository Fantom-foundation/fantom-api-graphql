// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"strconv"
)

// UniswapAction represents resolvable blockchain uniswap action structure.
type UniswapAction struct {
	types.UniswapAction
	*UniswapPair
}

// NewUniswapAction builds new resolvable uniswap action structure.
func NewUniswapAction(ua *types.UniswapAction, pair *UniswapPair) *UniswapAction {
	return &UniswapAction{
		UniswapAction: *ua,
		UniswapPair:   pair,
	}
}

// UniswapActionList represents resolvable list of blockchain uniswap action edges structure.
type UniswapActionList struct {
	types.UniswapActionList
}

// UniswapActionListEdge represents a single edge of a uniswap action list structure.
type UniswapActionListEdge struct {
	UniswapAction *UniswapAction
	Cursor        Cursor
}

// NewUniswapActionList builds new resolvable list of uniswap actions.
func NewUniswapActionList(ual *types.UniswapActionList) *UniswapActionList {
	return &UniswapActionList{*ual}
}

// DefiUniswapActions resolves list of blockchain uniswap actions encapsulated in a listable structure.
func (rs *rootResolver) DefiUniswapActions(args *struct {
	Cursor      *Cursor
	Count       int32
	PairAddress *common.Address
	ActionType  *int32
}) (*UniswapActionList, error) {
	// limit query size; the count can be either positive or negative
	// this controls the loading direction
	args.Count = listLimitCount(args.Count, listMaxEdgesPerRequest)
	if args.ActionType == nil {
		var t int32 = -1
		args.ActionType = &t
	}

	// get the uniswap action list from repository
	al, err := repository.R().UniswapActions(args.PairAddress, (*string)(args.Cursor), args.Count, *args.ActionType)
	if err != nil {
		log.Errorf("can not get uniswap action list; %s", err.Error())
		return nil, err
	}
	return NewUniswapActionList(al), nil
}

// TotalCount resolves the total number of uniswap actions in the list.
func (cl *UniswapActionList) TotalCount() hexutil.Big {
	val := (*hexutil.Big)(new(big.Int).SetUint64(cl.Total))
	return *val
}

// PageInfo resolves the current page information for the uniswap action list.
func (cl *UniswapActionList) PageInfo() (*ListPageInfo, error) {
	// do we have any items?
	if cl.Collection == nil || len(cl.Collection) == 0 {
		return NewListPageInfo(nil, nil, false, false)
	}

	// get the first and last elements
	first := Cursor(strconv.FormatUint(cl.First, 10))
	last := Cursor(strconv.FormatUint(cl.Last, 10))
	return NewListPageInfo(&first, &last, !cl.IsEnd, !cl.IsStart)
}

// Edges resolves list of edges for the linked uniswap action list.
func (cl *UniswapActionList) Edges() []*UniswapActionListEdge {
	// do we have any items? return empty list if not
	if cl.Collection == nil || len(cl.Collection) == 0 {
		return make([]*UniswapActionListEdge, 0)
	}

	// make the list
	edges := make([]*UniswapActionListEdge, len(cl.Collection))
	for i, c := range cl.Collection {
		edges[i] = &UniswapActionListEdge{
			UniswapAction: NewUniswapAction(c, NewUniswapPair(&c.PairAddress)),
			Cursor:        Cursor(strconv.FormatUint(c.OrdIndex, 10)),
		}
	}

	return edges
}
