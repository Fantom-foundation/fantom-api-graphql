// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
)

// Erc1155ContractList resolves a list of ERC1155 multi-token contracts.
func (rs *rootResolver) Erc1155ContractList(args struct{ Count int32 }) ([]*ERC1155Contract, error) {
	// limit query size; the count can be either positive or negative
	// this controls the loading direction
	args.Count = listLimitCount(args.Count, listMaxEdgesPerRequest)

	// get the list of addresses of active tokens
	al, err := repository.R().Erc1155ContractsList(args.Count)
	if err != nil {
		return nil, err
	}

	list := make([]*ERC1155Contract, len(al))
	for i, adr := range al {
		list[i] = NewErc1155Contract(&adr)
	}

	return list, nil
}