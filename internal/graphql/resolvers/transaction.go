// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
)

// Transaction represents resolvable blockchain transaction structure.
type Transaction struct {
	repo repository.Repository
	types.Transaction
}

// NewTransaction builds new resolvable transaction structure.
func NewTransaction(trx *types.Transaction, repo repository.Repository) *Transaction {
	return &Transaction{
		repo:        repo,
		Transaction: *trx,
	}
}

// Transaction resolves blockchain transaction by transaction hash.
func (rs *rootResolver) Transaction(args *struct{ Hash types.Hash }) (*Transaction, error) {
	// get the transaction from repository
	trx, err := rs.repo.Transaction(&args.Hash)
	if err != nil {
		rs.log.Warningf("can not get transaction %s", args.Hash)
		return nil, err
	}

	return NewTransaction(trx, rs.repo), nil
}

// Sender resolves sender account of the transaction.
func (trx *Transaction) Sender() (*Account, error) {
	// get the sender by address
	acc, err := trx.repo.Account(&trx.From)
	if err != nil {
		return nil, err
	}

	return NewAccount(acc, trx.repo), nil
}

// Sender resolves sender account of the transaction.
func (trx *Transaction) Recipient() (*Account, error) {
	// no recipient available
	if trx.To == nil {
		return nil, nil
	}

	// get the recipient by address
	acc, err := trx.repo.Account(trx.To)
	if err != nil {
		return nil, err
	}

	return NewAccount(acc, trx.repo), nil
}

// Block resolves block the transaction is bundled in, nil if it's pending and not added to a block yet.
func (trx *Transaction) Block() (*Block, error) {
	// no recipient available
	if trx.BlockNumber == nil {
		return nil, nil
	}

	// get the sender by address
	blk, err := trx.repo.BlockByNumber(trx.BlockNumber)
	if err != nil {
		return nil, err
	}

	return NewBlock(blk, trx.repo), nil
}
