package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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

// Erc721Transactions resolves list of ERC721 transactions.
func (rs *rootResolver) Erc721Transactions(args struct {
	Cursor  *Cursor
	Count   int32
	Token   *common.Address
	TokenId *hexutil.Big
	Account *common.Address
	TxType  *[]string
}) (*ERC721TransactionList, error) {
	// return empty transaction list to keep existing GraphQL schema
	return NewERC721TransactionList(&types.TokenTransactionList{Collection: make([]*types.TokenTransaction, 0)}), nil
}

// Erc1155Transactions resolves list of ERC1155 transactions.
func (rs *rootResolver) Erc1155Transactions(args struct {
	Cursor  *Cursor
	Count   int32
	Token   *common.Address
	TokenId *hexutil.Big
	Account *common.Address
	TxType  *[]string
}) (*ERC1155TransactionList, error) {
	// return empty transaction list to keep existing GraphQL schema
	return NewERC1155TransactionList(&types.TokenTransactionList{Collection: make([]*types.TokenTransaction, 0)}), nil
}
