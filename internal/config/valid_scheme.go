package config

import (
	"fmt"

	"github.com/ethereum/go-ethereum/core/forkid"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/p2p/enr"
	"github.com/ethereum/go-ethereum/rlp"
)

// List of known secure identity schemes.
var OperaNodeValidSchemes = enr.SchemeMap{
	"v4": operaNodeV4ID{},
}

// operaNodeV4ID is the "v4" identity scheme.
type operaNodeV4ID struct {
	v enode.V4ID
}

// operaNodeEnrEntry is the ENR entry which advertises `eth` protocol on the discovery.
type operaNodeEnrEntry struct {
	ForkID forkid.ID // Fork identifier per EIP-2124

	// Ignore additional fields (for forward compatibility).
	Rest []rlp.RawValue `rlp:"tail"`
}

// ENRKey implements enr.Entry.
func (e operaNodeEnrEntry) ENRKey() string {
	return "opera"
}

func (m operaNodeV4ID) Verify(r *enr.Record, sig []byte) error {
	var opera operaNodeEnrEntry
	err := r.Load(&opera)
	if err != nil {
		return fmt.Errorf("invalid opera node; %s", err.Error())
	}

	return m.v.Verify(r, sig)
}

func (m operaNodeV4ID) NodeAddr(r *enr.Record) []byte {
	return m.v.NodeAddr(r)
}
