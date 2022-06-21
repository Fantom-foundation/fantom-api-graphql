// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"golang.org/x/sync/singleflight"
)

// Transaction represents resolvable blockchain transaction structure.
type Transaction struct {
	types.Transaction
	cg *singleflight.Group
}

// NewTransaction builds new resolvable transaction structure.
func NewTransaction(trx *types.Transaction) *Transaction {
	return &Transaction{
		Transaction: *trx,
		cg:          new(singleflight.Group),
	}
}

// Transaction resolves blockchain transaction by transaction hash.
func (rs *rootResolver) Transaction(args *struct{ Hash common.Hash }) (tx *Transaction, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Criticalf("transaction loader crashed on %s", args.Hash.String())
			err = fmt.Errorf("failed to load transaction %s", args.Hash.String())
			tx = nil
		}
	}()

	// get the transaction from repository
	trx, err := repository.R().Transaction(&args.Hash)
	if err != nil {
		log.Warningf("can not get transaction %s", args.Hash)
		return nil, err
	}

	// transaction not found, yet no error?
	if trx == nil {
		log.Errorf("transaction %s not found", args.Hash.String())
		return nil, fmt.Errorf("transaction %s not found", args.Hash.String())
	}
	return NewTransaction(trx), nil
}

// SendTransaction sends raw signed and RLP encoded transaction to the blockchain.
func (rs *rootResolver) SendTransaction(args *struct{ Tx hexutil.Bytes }) (*Transaction, error) {
	// get the transaction from repository
	trx, err := repository.R().SendTransaction(args.Tx)
	if err != nil {
		log.Warningf("can not send transaction; %s", err.Error())
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

// TokenTransactions resolves list of all ERC20 token transactions involved
// with the base transaction call.
func (trx *Transaction) TokenTransactions() ([]*TokenTransaction, error) {
	// get all the transaction
	tl, err := trx.tokenTransactions()
	if err != nil {
		return nil, err
	}

	// convert to resolvable
	list := make([]*TokenTransaction, len(tl))
	for i, tx := range tl {
		list[i] = NewTokenTransaction(tx)
	}
	return list, nil
}

// tokenTransactions loads list of all token transaction related to this transaction call.
func (trx *Transaction) tokenTransactions() ([]*types.TokenTransaction, error) {
	// call for it only once
	val, err, _ := trx.cg.Do("erc", func() (interface{}, error) {
		log.Noticef("Loading ERC list for %s", trx.Hash.String())
		return repository.R().TokenTransactionsByCall(&trx.Hash)
	})
	if err != nil {
		return nil, err
	}
	return val.([]*types.TokenTransaction), nil
}
