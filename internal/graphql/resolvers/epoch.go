// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// Epoch represents a resolvable Epoch representation
type Epoch struct {
	types.Epoch
}

// Epoch resolves information about epoch of the given id.
func (rs *rootResolver) Epoch(args *struct{ Id *hexutil.Uint64 }) (Epoch, error) {
	epo, err := repository.R().Epoch(args.Id)
	if err != nil {
		return Epoch{}, err
	}
	return Epoch{*epo}, nil
}

// Duration resolves the time length of the given epoch
func (ep Epoch) Duration() hexutil.Uint64 {
	// no length for the first epochs
	if uint64(ep.Id) < 2 {
		return 0
	}

	// get the previous epoch so we can compare end times
	pid := uint64(ep.Id) - 1
	prev, err := repository.R().Epoch((*hexutil.Uint64)(&pid))
	if err != nil {
		return 0
	}

	// can we even calculate the duration?
	if ep.EndTime < prev.EndTime {
		return 0
	}
	return ep.EndTime - prev.EndTime
}

// FtmTreasuryTotal resolves total amount of FTM tokens in WEI units sent into treasury.
func (rs *rootResolver) FtmTreasuryTotal() hexutil.Big {
	val, err := repository.R().FtmTreasuryTotal()
	if err != nil {
		log.Criticalf("failed to load treasury total; %s", err.Error())
		return hexutil.Big{}
	}
	return hexutil.Big(*new(big.Int).Mul(big.NewInt(val), types.BurnDecimalsCorrection))
}

// FtmTreasuryTotalAmount resolves total amount of FTM tokens in FTM units sent into treasury.
func (rs *rootResolver) FtmTreasuryTotalAmount() float64 {
	val, err := repository.R().FtmTreasuryTotal()
	if err != nil {
		log.Criticalf("failed to load treasury total; %s", err.Error())
		return 0
	}
	return float64(val) / types.BurnFTMDecimalsCorrection
}
