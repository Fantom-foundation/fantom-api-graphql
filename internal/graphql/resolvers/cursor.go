// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"errors"
	"strconv"
)

// Cursor represents a string key of an element position in a sequential list of edges.
type Cursor string

// ImplementsGraphQLType notifies the GraphQL that this type resolves Cursor scalar.
func (Cursor) ImplementsGraphQLType(name string) bool {
	return name == "Cursor"
}

// UnmarshalGraphQL unmarshal incoming Cursor into a local variable.
func (c *Cursor) UnmarshalGraphQL(input interface{}) error {
	var err error
	switch input := input.(type) {
	case string:
		*c = Cursor(input)
	case int32:
		*c = Cursor(strconv.Itoa(int(input)))
	default:
		err = errors.New("wrong cursor type")
	}
	return err
}

// MarshalJSON encodes a cursor to JSON for transport.
func (c Cursor) MarshalJSON() ([]byte, error) {
	return strconv.AppendQuote(nil, string(c)), nil
}
