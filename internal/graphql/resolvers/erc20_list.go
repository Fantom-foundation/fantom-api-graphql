// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

// Erc20Token resolves an instance of ERC20 token if available.
func (rs *rootResolver) Erc20TokenList(args struct{ Count int32 }) ([]*ERC20Token, error) {
	// limit query size; the count can be either positive or negative
	// this controls the loading direction
	args.Count = listLimitCount(args.Count, listMaxEdgesPerRequest)

	// get the list of addresses of active tokens
	al, err := rs.repo.Erc20TokensList(args.Count)
	if err != nil {
		return nil, err
	}

	// make the container
	list := make([]*ERC20Token, len(al))

	// loop all the
	for i, adr := range al {
		list[i] = NewErc20Token(&adr, rs.repo)
	}

	return list, nil
}
