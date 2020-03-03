// Package gqlschema provides GraphQL schema definition used by GraphQL handler
// to validate requests and build responses on the API interface.
package gqlschema

//go:generate sh ./tools/make_bundle.sh

// Schema provides textual representation of the GraphQL schema content.
func Schema() string {
	return schema
}
