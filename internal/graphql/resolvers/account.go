// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"golang.org/x/sync/singleflight"
	"math/big"
)

// accMaxTransactionsPerRequest maximal number of transaction end-client can request in one query.
const accMaxTransactionsPerRequest = 250

// Account represents resolvable blockchain account structure.
type Account struct {
	types.Account
	cg singleflight.Group
}

// NewAccount builds new resolvable account structure.
func NewAccount(acc *types.Account) *Account {
	return &Account{
		Account: *acc,
	}
}

// Account resolves blockchain account by address.
func (rs *rootResolver) Account(args struct{ Address common.Address }) (*Account, error) {
	// simply pull the block by hash
	acc, err := repository.R().Account(&args.Address)
	if err != nil {
		log.Errorf("could not get the specified account")
		return nil, err
	}
	return NewAccount(acc), nil
}

// AccountsActive resolves total number of active accounts on the blockchain.
func (rs *rootResolver) AccountsActive() (hexutil.Uint64, error) {
	return repository.R().AccountsActive()
}

// Balance resolves total balance of the account.
func (acc *Account) Balance() (hexutil.Big, error) {
	// get the balance
	val, err, _ := acc.cg.Do("balance", func() (interface{}, error) {
		return repository.R().AccountBalance(&acc.Address)
	})

	// can not get the balance?
	if err != nil {
		return hexutil.Big{}, err
	}
	return *val.(*hexutil.Big), nil
}

// TotalValue resolves account total value including delegated amount and pending rewards.
func (acc *Account) TotalValue() (hexutil.Big, error) {
	// get the balance
	balance, err := acc.Balance()
	if err != nil {
		return hexutil.Big{}, err
	}

	// try to pull the delegations details
	delegated, pendingOut, rewards, err := acc.delegationsTotal()
	if err != nil {
		return hexutil.Big{}, err
	}

	// calc the sum
	val := new(big.Int).Add(new(big.Int).Add(new(big.Int).Add(balance.ToInt(), delegated), rewards), pendingOut)
	return hexutil.Big(*val), nil
}

// TxCount resolves the number of transaction sent by the account, also known as nonce.
func (acc *Account) TxCount() (hexutil.Uint64, error) {
	// get the sender by address
	bal, err := repository.R().AccountNonce(&acc.Address)
	if err != nil {
		return hexutil.Uint64(0), err
	}

	return *bal, nil
}

// TxList resolves list of transaction associated with the account.
func (acc *Account) TxList(args struct {
	Recipient *common.Address
	Cursor    *Cursor
	Count     int32
}) (*TransactionList, error) {
	// limit query size; the count can be either positive or negative
	// this controls the loading direction
	args.Count = listLimitCount(args.Count, accMaxTransactionsPerRequest)

	// get the transaction hash list from repository
	bl, err := repository.R().AccountTransactions(&acc.Address, args.Recipient, (*string)(args.Cursor), args.Count)
	if err != nil {
		return nil, err
	}

	return NewTransactionList(bl), nil
}

// Erc20TxList resolves list of ERC20 transactions associated with the account.
func (acc *Account) Erc20TxList(args struct {
	Cursor *Cursor
	Count  int32
	Token  *common.Address
	TxType *[]string
}) (*ERC20TransactionList, error) {
	// limit query size; the count can be either positive or negative
	// this controls the loading direction
	args.Count = listLimitCount(args.Count, accMaxTransactionsPerRequest)

	// get the transaction hash list from repository
	tl, err := repository.R().TokenTransactions(
		types.AccountTypeERC20Token,
		args.Token,
		nil,
		&acc.Address,
		ercTrxTypesFromNames(args.TxType),
		(*string)(args.Cursor),
		args.Count,
	)
	if err != nil {
		return nil, err
	}

	return NewERC20TransactionList(tl), nil
}

// Erc721TxList resolves list of ERC721 transactions associated with the account.
func (acc *Account) Erc721TxList(args struct {
	Cursor  *Cursor
	Count   int32
	Token   *common.Address
	TokenId *hexutil.Big
	TxType  *[]string
}) (*ERC721TransactionList, error) {
	// limit query size; the count can be either positive or negative
	// this controls the loading direction
	args.Count = listLimitCount(args.Count, accMaxTransactionsPerRequest)

	// get the transaction hash list from repository
	tl, err := repository.R().TokenTransactions(
		types.AccountTypeERC721Contract,
		args.Token,
		(*big.Int)(args.TokenId),
		&acc.Address,
		ercTrxTypesFromNames(args.TxType),
		(*string)(args.Cursor),
		args.Count,
	)
	if err != nil {
		return nil, err
	}

	return NewERC721TransactionList(tl), nil
}

// Erc1155TxList resolves list of ERC1155 transactions associated with the account.
func (acc *Account) Erc1155TxList(args struct {
	Cursor  *Cursor
	Count   int32
	Token   *common.Address
	TokenId *hexutil.Big
	TxType  *[]string
}) (*ERC1155TransactionList, error) {
	// limit query size; the count can be either positive or negative
	// this controls the loading direction
	args.Count = listLimitCount(args.Count, accMaxTransactionsPerRequest)

	// get the transaction hash list from repository
	tl, err := repository.R().TokenTransactions(
		types.AccountTypeERC1155Contract,
		args.Token,
		(*big.Int)(args.TokenId),
		&acc.Address,
		ercTrxTypesFromNames(args.TxType),
		(*string)(args.Cursor),
		args.Count,
	)
	if err != nil {
		return nil, err
	}

	return NewERC1155TransactionList(tl), nil
}

// Staker resolves the account staker detail, if the account is a staker.
func (acc *Account) Staker() (*Staker, error) {
	// get the staker
	st, err := repository.R().ValidatorByAddress(&acc.Address)
	if err != nil {
		return nil, err
	}

	// staker not found?
	if st == nil {
		return nil, nil
	}
	return NewStaker(st), nil
}

// Delegations resolves a list of account delegations, if the account is a delegator.
func (acc *Account) Delegations(args *struct {
	Cursor *Cursor
	Count  int32
}) (*DelegationList, error) {
	// limit query size; the count can be either positive or negative
	// this controls the loading direction
	args.Count = listLimitCount(args.Count, listMaxEdgesPerRequest)

	// pull the list
	dl, err := repository.R().DelegationsByAddress(&acc.Address, (*string)(args.Cursor), args.Count)
	if err != nil {
		return nil, err
	}

	// convert to resolvable list
	return NewDelegationList(dl), nil
}

// Contract resolves the account smart contract detail,
// if the account is a smart contract address.
func (acc *Account) Contract() (*Contract, error) {
	// is this actually a contract account?
	if acc.ContractTx == nil {
		return nil, nil
	}

	// get new contract
	con, err := repository.R().Contract(&acc.Address)
	if err != nil {
		return nil, err
	}
	return NewContract(con), nil
}

// delegationsTotal calculates total sum of delegations of the given account including
// pending rewards for those delegations.
func (acc *Account) delegationsTotal() (amount *big.Int, inWithdraw *big.Int, rewards *big.Int, err error) {
	// pull all the delegations of the account
	list, err := repository.R().DelegationsByAddressAll(&acc.Address)
	if err != nil {
		return nil, nil, nil, err
	}

	// prep containers for calculation and loop all delegations found
	amount = new(big.Int)
	rewards = new(big.Int)
	inWithdraw = new(big.Int)
	for _, dlg := range list {
		// any active delegated amount?
		if 0 < dlg.AmountDelegated.ToInt().Uint64() {
			amount = new(big.Int).Add(amount, dlg.AmountDelegated.ToInt())
		}

		// get pending rewards for this delegation (can be stashed)
		rw, err := repository.R().PendingRewards(&acc.Address, dlg.ToStakerId)
		if err != nil {
			return nil, nil, nil, err
		}

		// any rewards?
		if 0 < rw.Amount.ToInt().Uint64() {
			rewards = new(big.Int).Add(rewards, rw.Amount.ToInt())
		}

		// get pending withdrawals
		wd, err := repository.R().WithdrawRequestsPendingTotal(&acc.Address, dlg.ToStakerId)
		if err != nil {
			return nil, nil, nil, err
		}

		// add pending withdrawals value
		if 0 < wd.Uint64() {
			inWithdraw = new(big.Int).Add(inWithdraw, wd)
		}
	}

	return amount, rewards, inWithdraw, nil
}
