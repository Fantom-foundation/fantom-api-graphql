package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"github.com/ethereum/go-ethereum/common"
)

// Erc20Transactions resolves list of ERC20 transactions.
func (rs *rootResolver) Erc20Transactions(args struct {
	Cursor  *Cursor
	Count   int32
	Token   *common.Address
	Account *common.Address
	TxType  *[]string
}) (*ERC20TransactionList, error) {
	// limit query size; the count can be either positive or negative
	// this controls the loading direction
	args.Count = listLimitCount(args.Count, accMaxTransactionsPerRequest)

	// get the transaction hash list from repository
	tl, err := repository.R().TokenTransactions(
		args.Token,
		args.Account,
		ercTrxTypesFromNames(args.TxType),
		(*string)(args.Cursor),
		args.Count,
	)
	if err != nil {
		return nil, err
	}

	return NewERC20TransactionList(tl), nil
}
