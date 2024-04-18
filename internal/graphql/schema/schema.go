// Package schema provides GraphQL schema definition used by GraphQL handler
// to validate requests and build responses on the API interface.
package schema

import _ "embed"

//go:generate sh ./tools/make_bundle.sh

//go:embed schema.graphql
var schema string

// Schema provides textual representation of the GraphQL schema content.
func Schema() string {
	return schema
}
