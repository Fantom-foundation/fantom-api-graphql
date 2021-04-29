package handlers

import (
	"fantom-api-graphql/internal/graphql/resolvers"
	gqlSchema "fantom-api-graphql/internal/graphql/schema"
	"github.com/graph-gophers/graphql-go"
	"github.com/onsi/gomega"
	"testing"
)

// TestGraphQLApiMustCompile tests if the GraphQL API compiles and can be used
// to resolve incoming requests.
func TestGraphQLApiMustCompile(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	// we don't want to write a method for each type field if it could be matched directly
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}

	// try to compile the schema
	_, err := graphql.ParseSchema(gqlSchema.Schema(), resolvers.New(cfg, log, repo), opts...)
	g.Expect(err).Should(gomega.Succeed())
}
