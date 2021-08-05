// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"github.com/ethereum/go-ethereum/common"
)

// Erc20TokenList resolves an instance of ERC20 token list if available.
func (rs *rootResolver) Erc20TokenList(args struct{ Count int32 }) ([]*ERC20Token, error) {
	// limit query size; the count can be either positive or negative
	// this controls the loading direction
	args.Count = listLimitCount(args.Count, listMaxEdgesPerRequest)

	// get the list of addresses of active tokens
	al, err := repository.R().Erc20TokensList(args.Count)
	if err != nil {
		return nil, err
	}

	// make the container and create resolvables
	list := make([]*ERC20Token, len(al))
	for i, adr := range al {
		list[i] = NewErc20Token(&adr)
	}

	return list, nil
}

// Erc20Assets resolves a list of instances of ERC20 tokens for the given owner.
func (rs *rootResolver) Erc20Assets(args struct {
	Owner common.Address
	Count int32
}) ([]*ERC20Token, error) {
	// limit query size; the count can be either positive or negative
	// this controls the loading direction
	args.Count = listLimitCount(args.Count, listMaxEdgesPerRequest)

	// get the list of addresses of active tokens for the owner
	al, err := repository.R().Erc20Assets(args.Owner, args.Count)
	if err != nil {
		return nil, err
	}

	// make the container and build the list (limit to recognized assets)
	list := make([]*ERC20Token, len(al))
	for i, token := range al {
		list[i] = NewErc20Token(&token)
	}

	return list, nil
}

// ownsErc20Asset checks if the given owner has any tokens of the given ERC20.
func (rs *rootResolver) ownsErc20Asset(token *common.Address, owner *common.Address) bool {
	// get the balance for the owner
	val, err := repository.R().Erc20BalanceOf(token, owner)
	if err != nil {
		log.Errorf("token %s balance can not be loaded for %s; %s", token.String(), owner.String(), err.Error())
		return false
	}

	// any balance?
	return 0 < val.ToInt().Uint64()
}
