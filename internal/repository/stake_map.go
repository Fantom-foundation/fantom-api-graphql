/*
Package repository implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Opera/Lachesis full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import (
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"strings"
	"sync"
)

// StakeAmountMap stores a map of delegated stake balances
// to speed up delegation stake amount related processing.
type StakeAmountMap struct {
	mu   *sync.Mutex
	data map[string]uint64
}

// NewStakeAmountMap creates a new instance of the stake amount map.
func NewStakeAmountMap() *StakeAmountMap {
	return &StakeAmountMap{
		mu:   new(sync.Mutex),
		data: make(map[string]uint64),
	}
}

// Update updates the staked amount reference and provides feedback if the update was needed.
func (sam *StakeAmountMap) Update(adr *common.Address, valID *hexutil.Big, amo *big.Int) bool {
	// calculate stored GWei value
	gWei := new(big.Int).Div(amo, types.DelegationDecimalsCorrection).Uint64()

	// generate map key
	var key strings.Builder
	key.Write(adr.Bytes())
	key.Write(valID.ToInt().Bytes())

	// access the map in a thread sec way
	sam.mu.Lock()
	defer sam.mu.Unlock()

	// do we need to update the stored value?
	old, ok := sam.data[key.String()]
	if !ok || old != gWei {
		sam.data[key.String()] = gWei
		return true
	}
	return false
}
