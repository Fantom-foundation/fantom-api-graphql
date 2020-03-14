// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Account represents resolvable blockchain account structure.
type Account struct {
	repo repository.Repository
	types.Account
}

// NewAccount builds new resolvable account structure.
func NewAccount(acc *types.Account, repo repository.Repository) *Account {
	return &Account{
		repo:    repo,
		Account: *acc,
	}
}

// Block resolves blockchain block by number or by hash. If neither is provided, the most recent block is given.
// Account resolves blockchain account by address.
func (rs *rootResolver) Account(args struct{ Address common.Address }) (*Account, error) {
	// simply pull the block by hash
	acc, err := rs.repo.Account(&args.Address)
	if err != nil {
		rs.log.Errorf("could not get the specified account")
		return nil, err
	}

	return NewAccount(acc, rs.repo), nil
}

// Sender resolves sender account of the transaction.
func (acc *Account) Balance() (hexutil.Big, error) {
	// get the sender by address
	bal, err := acc.repo.AccountBalance(&acc.Account)
	if err != nil {
		return hexutil.Big{}, err
	}

	return *bal, nil
}

// TxCount resolves the number of transaction sent by the account, also known as nonce.
func (acc *Account) TxCount() (hexutil.Uint64, error) {
	// get the sender by address
	bal, err := acc.repo.AccountNonce(&acc.Account)
	if err != nil {
		return hexutil.Uint64(0), err
	}

	return *bal, nil
}

// TxList resolves list of transaction associated with the account.
func (acc *Account) TxList(args struct {
	Cursor *Cursor
	Count  int32
}) (*TransactionList, error) {
	// get the transaction hash list from repository
	bl, err := acc.repo.AccountTransactions(&acc.Account, (*string)(args.Cursor), args.Count)
	if err != nil {
		return nil, err
	}

	return NewTransactionList(bl, acc.repo), nil
}
