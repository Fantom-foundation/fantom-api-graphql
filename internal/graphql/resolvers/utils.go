// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"crypto/rand"
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"io"
)

// Price resolves price details of the Opera blockchain token for the given target symbols.
func (rs *rootResolver) Price(args *struct{ To string }) (types.Price, error) {
	return rs.repo.Price(args.To)
}

// GasPrice resolves the current amount of WEI for single Gas.
func (rs *rootResolver) GasPrice() (hexutil.Uint64, error) {
	return rs.repo.GasPrice()
}

// EstimateGas resolves the estimated amount of Gas required to perform
// transaction described by the input params.
func (rs *rootResolver) EstimateGas(args struct {
	From  *common.Address
	To    *common.Address
	Value *hexutil.Big
	Data  *string
}) *hexutil.Uint64 {
	return rs.repo.GasEstimate(&args)
}

// uuid generates new random subscription UUID
func uuid() (string, error) {
	// prep container
	uuid := make([]byte, 16)

	// try to read random numbers
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}

	// variant bits
	uuid[8] = uuid[8]&^0xc0 | 0x80

	// version 4 (pseudo-random)
	uuid[6] = uuid[6]&^0xf0 | 0x40

	// format
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
