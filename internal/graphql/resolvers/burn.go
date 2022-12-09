// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// FtmBurnedTotal resolves total amount of burned FTM tokens in WEI units.
func (rs *rootResolver) FtmBurnedTotal() hexutil.Big {
	val, err := repository.R().FtmBurnTotal()
	if err != nil {
		log.Criticalf("failed to load burned total; %s", err.Error())
		return hexutil.Big{}
	}
	return hexutil.Big(*new(big.Int).Mul(big.NewInt(val), types.BurnDecimalsCorrection))
}

// FtmBurnedTotalAmount resolves total amount of burned FTM tokens in FTM units.
func (rs *rootResolver) FtmBurnedTotalAmount() float64 {
	val, err := repository.R().FtmBurnTotal()
	if err != nil {
		log.Criticalf("failed to load burned total; %s", err.Error())
		return 0
	}
	return float64(val) / types.BurnFTMDecimalsCorrection
}

// FtmLatestBlockBurnList resolves a list of the latest block burns.
func (rs *rootResolver) FtmLatestBlockBurnList(args struct{ Count int32 }) ([]*types.FtmBurn, error) {
	if args.Count < 1 || args.Count > 50 {
		args.Count = 25
	}
	return repository.R().FtmBurnList(int64(args.Count))
}
