// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

// ListPageInfo represents general resolvable information about the current page of a list of elements.
type ListPageInfo struct {
	First       *Cursor
	Last        *Cursor
	HasNext     bool
	HasPrevious bool
}

// NewListPageInfo creates a new page information structure.
func NewListPageInfo(first *Cursor, last *Cursor, hasNext bool, hasPrevious bool) (*ListPageInfo, error) {
	// make the structure
	return &ListPageInfo{
		First:       first,
		Last:        last,
		HasNext:     hasNext,
		HasPrevious: hasPrevious,
	}, nil
}
