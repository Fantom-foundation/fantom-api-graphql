// Package gqlschema provides GraphQL schema definition used by GraphQL handler
// to validate requests and build responses on the API interface.
package gqlschema

import (
	"github.com/onsi/gomega"
	"testing"
)

// bundleContentTest represents a single test for GraphQL API bundle content.
type bundleContentTest struct {
	re  string
	msg string
}

// the list of expected content of the GraphQL API definition bundle.
var expectedContent = []bundleContentTest{
	{
		re:  "(?m)^schema\\s+{(\\s+(query|mutation|subscription)\\s*:\\s*(Query|Mutation|Subscription))+\\s*}",
		msg: "root/master schema definition must exists",
	},
	{
		re:  "(?m)^type\\s+Query\\s+{",
		msg: "query schema must exists",
	},
	{
		re:  "(?m)^type\\s+Mutation\\s+{",
		msg: "mutation schema must exists",
	},
	{
		re:  "(?m)^scalar\\s+Long",
		msg: "scalars definition must be included",
	},
	{
		re:  "(?m)^type\\s+CurrentState\\s+{",
		msg: "current state type must exists",
	},
	{
		re:  "(?m)^type\\s+Account\\s+{",
		msg: "account detail type must exists",
	},
	{
		re:  "(?m)^type\\s+Block\\s+{",
		msg: "block detail type must exists",
	},
	{
		re:  "(?m)^type\\s+Epoch\\s+{",
		msg: "SFC epoch type must exists",
	},
	{
		re:  "(?m)^type\\s+Staker\\s+{",
		msg: "SFC staker detail type must exists",
	},
	{
		re:  "(?m)^type\\s+Delegation\\s+{",
		msg: "SFC delegation detail type must exists",
	},
	{
		re:  "(?m)^type\\s+ERC20Token\\s+{",
		msg: "ERC20 token detail type must exists",
	},
	{
		re:  "(?m)^type\\s+FMintAccount\\s+{",
		msg: "fMint account detail type must exists",
	},
	{
		re:  "(?m)^type\\s+DefiSettings\\s+{",
		msg: "DeFi settings detail type must exists",
	},
	{
		re:  "(?m)^type\\s+DefiUniswapVolume\\s+{",
		msg: "DeFi Uniswap Volume detail type must exists",
	},
	{
		re:  "(?m)^type\\s+DefiTimeVolume\\s+{",
		msg: "DeFi Time Volume detail type must exists",
	},
	{
		re:  "(?m)^type\\s+DefiTimePrice\\s+{",
		msg: "DeFi Time Price detail type must exists",
	},
	{
		re:  "(?m)^type\\s+UniswapActionList\\s+{",
		msg: "DeFi Uniswap Action List type must exists",
	},
	{
		re:  "(?m)^type\\s+GovernanceContract\\s+{",
		msg: "governance contract type must exists",
	},
}

// TestBundleContent tests if the bundle contains all the expected elements
// we need to support the expected API interface.
func TestBundleContent(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	s := Schema()

	// test all
	for _, c := range expectedContent {
		g.Expect(s).To(gomega.MatchRegexp(c.re))
	}
}
