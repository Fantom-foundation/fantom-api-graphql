// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Transaction represents resolvable blockchain transaction structure.
type Transaction struct {
	types.Transaction
}

// NewTransaction builds new resolvable transaction structure.
func NewTransaction(trx *types.Transaction) *Transaction {
	return &Transaction{
		Transaction: *trx,
	}
}

// Transaction resolves blockchain transaction by transaction hash.
func (rs *rootResolver) Transaction(args *struct{ Hash common.Hash }) (*Transaction, error) {
	// get the transaction from repository
	trx, err := repository.R().Transaction(&args.Hash)
	if err != nil {
		rs.log.Warningf("can not get transaction %s", args.Hash)
		return nil, err
	}

	return NewTransaction(trx), nil
}

// SendTransaction sends raw signed and RLP encoded transaction to the block chain.
func (rs *rootResolver) SendTransaction(args *struct{ Tx hexutil.Bytes }) (*Transaction, error) {
	// get the transaction from repository
	trx, err := repository.R().SendTransaction(args.Tx)
	if err != nil {
		rs.log.Warningf("can not send transaction %s", err.Error())
		return nil, err
	}

	return NewTransaction(trx), nil
}

// Sender resolves sender's account of the transaction.
func (trx *Transaction) Sender() (*Account, error) {
	// get the sender by address
	acc, err := repository.R().Account(&trx.From)
	if err != nil {
		return nil, err
	}

	return NewAccount(acc), nil
}

// Recipient resolves recipient's account of the transaction.
func (trx *Transaction) Recipient() (*Account, error) {
	// no recipient available
	if trx.To == nil {
		return nil, nil
	}

	// get the recipient by address
	acc, err := repository.R().Account(trx.To)
	if err != nil {
		return nil, err
	}

	return NewAccount(acc), nil
}

// Block resolves block the transaction is bundled in, nil if it's pending and not added to a block yet.
func (trx *Transaction) Block() (*Block, error) {
	// no recipient available
	if trx.BlockNumber == nil {
		return nil, nil
	}

	// get the sender by address
	blk, err := repository.R().BlockByNumber(trx.BlockNumber)
	if err != nil {
		return nil, err
	}

	return NewBlock(blk), nil
}
