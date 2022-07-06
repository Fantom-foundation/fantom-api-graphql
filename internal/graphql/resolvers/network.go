// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"fmt"
	"math"
)

// NetworkNodeGroup represents an aggregated Opera nodes group with a common location.
type NetworkNodeGroup struct {
	list *NetworkNodeGroupList
	types.OperaNodeLocationAggregate
}

// NetworkNodeGroupList represents a list of aggregated Opera nodes
type NetworkNodeGroupList struct {
	Level      string
	TotalCount int32
	Groups     []*NetworkNodeGroup
}

// NetworkNodesAggregated resolves a list of aggregated Opera nodes based on geographic location detail.
func (rs *rootResolver) NetworkNodesAggregated(args struct{ Level string }) (*NetworkNodeGroupList, error) {
	var level int
	switch args.Level {
	case "CONTINENT":
		level = types.OperaNodeGeoAggregationLevelContinent
	case "COUNTRY":
		level = types.OperaNodeGeoAggregationLevelCountry
	case "STATE":
		level = types.OperaNodeGeoAggregationLevelState
	default:
		return nil, fmt.Errorf("unknown precision level %s", args.Level)
	}

	// load the list
	list, err := repository.R().NetworkNodesGeoAggregated(level)
	if err != nil {
		return nil, err
	}

	// make the list
	nl := NetworkNodeGroupList{
		Level:      args.Level,
		TotalCount: 0,
		Groups:     make([]*NetworkNodeGroup, len(list)),
	}

	// populate the members
	for i, ag := range list {
		nl.Groups[i] = &NetworkNodeGroup{
			list:                       &nl,
			OperaNodeLocationAggregate: *ag,
		}
		nl.TotalCount += int32(ag.Count)
	}

	return &nl, nil
}

// Pct provides percentage of the group in the list.
// The number is provided as fixed point integer with 1 decimal precision (i.e. 258 = 25.8%, 1000 = 100%)
func (agm *NetworkNodeGroup) Pct() int32 {
	if agm.list == nil || agm.list.TotalCount == 0 {
		return 0
	}

	return int32(math.Round(float64(agm.Count) * 1000 / float64(agm.list.TotalCount)))
}
