// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
)

// Erc721TokenList resolves an instance of ERC721 token list if available.
func (rs *rootResolver) Erc721TokenList(args struct{ Count int32 }) ([]*ERC721Token, error) {
	// limit query size; the count can be either positive or negative
	// this controls the loading direction
	args.Count = listLimitCount(args.Count, listMaxEdgesPerRequest)

	// get the list of addresses of active tokens
	al, err := repository.R().Erc721TokensList(args.Count)
	if err != nil {
		return nil, err
	}

	// make the container and create resolvables
	list := make([]*ERC721Token, len(al))
	for i, adr := range al {
		list[i] = NewErc721Token(&adr)
	}

	return list, nil
}