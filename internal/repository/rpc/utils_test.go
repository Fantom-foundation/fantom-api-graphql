package rpc

import (
	"encoding/hex"
	"github.com/onsi/gomega"
	"testing"
)

func TestParseAbiString(t *testing.T) {
	g := gomega.NewWithT(t)

	data, err := hex.DecodeString("0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000f546865526172697479466f726573740000000000000000000000000000000000")
	g.Expect(err).To(gomega.BeNil())

	str, err := parseAbiString(data)
	g.Expect(err).To(gomega.BeNil())
	g.Expect(str).To(gomega.Equal("TheRarityForest"))
}
