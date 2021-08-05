// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"crypto/rand"
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"io"
	"regexp"
)

// reExpectedPriceSymbol represents a price symbol expected to be resolved
var reExpectedPriceSymbol = regexp.MustCompile(`^[\w]{2,4}$`)

// Price resolves price details of the Opera blockchain token for the given target symbols.
func (rs *rootResolver) Price(args *struct{ To string }) (types.Price, error) {
	// is the requested denomination even reasonable
	if !reExpectedPriceSymbol.Match([]byte(args.To)) {
		return types.Price{}, fmt.Errorf("invalid denomination received")
	}
	return repository.R().Price(args.To)
}

// GasPrice resolves the current amount of WEI for single Gas.
func (rs *rootResolver) GasPrice() (hexutil.Uint64, error) {
	// get the actual value
	price, err := repository.R().GasPrice()
	if err != nil {
		return hexutil.Uint64(0), err
	}

	// if the price safely within the range
	if !price.ToInt().IsUint64() {
		log.Error("current gas price is too high and can not be extracted")
		return hexutil.Uint64(0), err
	}

	// inform and return
	log.Debugf("current gas price is %d", price.ToInt().Uint64())
	return hexutil.Uint64(price.ToInt().Uint64()), nil
}

// EstimateGas resolves the estimated amount of Gas required to perform
// transaction described by the input params.
func (rs *rootResolver) EstimateGas(args struct {
	From  *common.Address
	To    *common.Address
	Value *hexutil.Big
	Data  *string
}) (*hexutil.Uint64, error) {
	return repository.R().GasEstimate(&args)
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
