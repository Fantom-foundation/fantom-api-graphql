// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// accMaxTransactionsPerRequest maximal number of transaction end-client can request in one query.
const accMaxTransactionsPerRequest = 50

// Account represents resolvable blockchain account structure.
type Account struct {
	repo         repository.Repository
	rfStaker     *types.Staker
	rfDelegation *types.Delegator
	rfBalance    *hexutil.Big
	stashBalance *big.Int

	/* extended delegated amounts pre-loaded */
	dlExtendedAmount           *big.Int
	dlExtendedAmountInWithdraw *big.Int

	types.Account
}

// NewAccount builds new resolvable account structure.
func NewAccount(acc *types.Account, repo repository.Repository) *Account {
	return &Account{
		repo:    repo,
		Account: *acc,
	}
}

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

// Resolves total number of active accounts on the blockchain.
func (rs *rootResolver) AccountsActive() (hexutil.Uint64, error) {
	return rs.repo.AccountsActive()
}

// Balance resolves total balance of the account.
func (acc *Account) Balance() (hexutil.Big, error) {
	if acc.rfBalance == nil {
		// get the sender by address
		bal, err := acc.repo.AccountBalance(&acc.Account)
		if err != nil {
			return hexutil.Big{}, err
		}

		acc.rfBalance = bal
	}

	return *acc.rfBalance, nil
}

// addStashedRewards adds stashed rewards to the provided total value.
func (acc *Account) addStashedRewards(val *big.Int) *big.Int {
	/*
		// get the delegation information
		dl, err := acc.getDelegation()

		// do we have a valid record?
		if err != nil || dl == nil || dl.ClaimedReward == nil {
			return val
		}
	*/

	// get the delegation information
	sb, err := acc.stashedBalance()
	if err != nil {
		return val
	}

	// calculate the new total by adding the stashed amount
	return new(big.Int).Add(val, sb)
}

// stashedBalance returns the current stash balance on account.
func (acc *Account) stashedBalance() (*big.Int, error) {
	// return pre-cached stash balance
	if acc.stashBalance == nil {
		// get the stash amount from SFC
		val, err := acc.repo.RewardsStash(&acc.Address)
		if err != nil {
			return nil, err
		}

		// keep cached
		acc.stashBalance = val
	}

	return acc.stashBalance, nil
}

// Stashed resolves the amount of WEI stashed for this account.
func (acc *Account) Stashed() (hexutil.Big, error) {
	// get the stash amount from SFC
	val, err := acc.stashedBalance()
	if err != nil {
		return hexutil.Big{}, err
	}

	return hexutil.Big(*val), nil
}

// CanUnStash resolves the stash status for this account.
func (acc *Account) CanUnStash() (bool, error) {
	// get the current lock status
	allowed, err := acc.repo.RewardsAllowed()
	if err != nil {
		return false, err
	}

	// rewards are still locked, no luck
	if !allowed {
		return false, nil
	}

	// get the stash amount from SFC
	val, err := acc.stashedBalance()
	if err != nil || val == nil {
		return false, err
	}

	return val.Cmp(big.NewInt(0)) > 0, nil
}

// TotalValue resolves address total value including delegated amount and pending rewards.
func (acc *Account) TotalValue() (hexutil.Big, error) {
	// get the balance
	balance, err := acc.Balance()
	if err != nil {
		return hexutil.Big{}, err
	}

	// try to get delegation
	del, err := acc.getDelegation()
	if err != nil {
		return balance, err
	}

	// do we have a delegation?
	if del != nil {
		var err error

		// try to load the data
		if acc.dlExtendedAmount == nil {
			acc.dlExtendedAmount, acc.dlExtendedAmountInWithdraw, err = acc.repo.DelegatedAmountExtended(del)
			if err != nil {
				return hexutil.Big{}, err
			}
		}

		// add delegated amount to the balance
		val := big.NewInt(0).Add(balance.ToInt(), acc.dlExtendedAmount)

		// get pending rewards
		rw, err := acc.repo.DelegationRewards(acc.Address.String())
		if err != nil {
			return hexutil.Big(*val), err
		}

		// add pending rewards to the final value
		val = big.NewInt(0).Add(val, rw.Amount.ToInt())

		// add claimed/stashed rewards from staking to the total
		val = acc.addStashedRewards(val)
		return hexutil.Big(*val), err
	}

	return balance, nil
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
	// limit query size; the count can be either positive or negative
	// this controls the loading direction
	args.Count = listLimitCount(args.Count, accMaxTransactionsPerRequest)

	// get the transaction hash list from repository
	bl, err := acc.repo.AccountTransactions(&acc.Account, (*string)(args.Cursor), args.Count)
	if err != nil {
		return nil, err
	}

	return NewTransactionList(bl, acc.repo), nil
}

// Staker resolves the account staker detail, if the account is a staker.
func (acc *Account) Staker() (*Staker, error) {
	// get the staker
	st, err := acc.getStaker()
	if err != nil {
		return nil, err
	}

	// not staker
	if st == nil {
		return nil, nil
	}

	return NewStaker(st, acc.repo), nil
}

// Delegation resolves the account delegator detail, if the account is a delegator.
func (acc *Account) Delegation() (*Delegator, error) {
	// try to get the delegator info
	dl, err := acc.repo.Delegation(acc.Address)
	if err != nil {
		return nil, err
	}

	// not delegator
	if dl == nil {
		return nil, nil
	}

	return NewDelegator(dl, acc.repo), nil
}

// Contract resolves the account smart contract detail,
// if the account is a smart contract address.
func (acc *Account) Contract() (*Contract, error) {
	// is this actually a contract account?
	if acc.ContractTx == nil {
		return nil, nil
	}

	// get new contract
	con, err := acc.repo.Contract(&acc.Address)
	if err != nil {
		return nil, err
	}

	return NewContract(con, acc.repo), nil
}

// getStaker returns lazy loaded staker information.
func (acc *Account) getStaker() (*types.Staker, error) {
	// try to get the staker info
	if acc.rfStaker == nil {
		st, err := acc.repo.StakerByAddress(acc.Address)
		if err != nil {
			return nil, err
		}

		// is this a valid staker info?
		if st.Id <= 0 {
			return nil, nil
		}

		// keep the staker info
		acc.rfStaker = st
	}

	return acc.rfStaker, nil
}

// getDelegation return lazy loaded delegation detail for the account.
func (acc *Account) getDelegation() (*types.Delegator, error) {
	// try to get the staker info
	if acc.rfDelegation == nil {
		// try to get the staker info
		dl, err := acc.repo.Delegation(acc.Address)
		if err != nil {
			return nil, err
		}

		// is this a valid staker info?
		if dl.ToStakerId <= 0 {
			return nil, nil
		}

		// keep the delegation reference so we don't have
		// to load it again if needed
		acc.rfDelegation = dl
	}

	return acc.rfDelegation, nil
}
