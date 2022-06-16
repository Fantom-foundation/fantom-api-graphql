/*
Package repository implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Opera full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import (
	"bytes"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// PullStakerInfo extracts an extended staker information from smart contact.
func (p *proxy) PullStakerInfo(id *hexutil.Big) (*types.StakerInfo, error) {
	return p.rpc.StakerInfo(id)
}

// StoreStakerInfo stores staker information to in-memory cache for future use.
func (p *proxy) StoreStakerInfo(id *hexutil.Big, sti *types.StakerInfo) error {
	// push to in-memory cache
	err := p.cache.PushStakerInfo(id, sti)
	if err != nil {
		p.log.Error("staker info can net be kept")
		return err
	}
	return nil
}

// RetrieveStakerInfo gets staker information from in-memory if available.
func (p *proxy) RetrieveStakerInfo(id *hexutil.Big) *types.StakerInfo {
	return p.cache.PullStakerInfo(id)
}

// IsStiContract returns true if the given address points to the STI contract.
func (p *proxy) IsStiContract(addr *common.Address) bool {
	return bytes.Equal(addr.Bytes(), p.cfg.Staking.StiContract.Bytes())
}
