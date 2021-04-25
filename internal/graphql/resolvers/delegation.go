// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"sort"
	"time"
)

// Delegation represents resolvable delegator detail.
type Delegation struct {
	repo repository.Repository
	types.Delegation

	/* extended delegated amounts pre-loaded */
	lock                     *types.DelegationLock
	extendedAmount           *big.Int
	extendedAmountInWithdraw *big.Int
}

// DelegationsByAge represents a list of delegations sortable by their age of creation.
type DelegationsByAge []types.Delegation

// NewDelegation creates new instance of resolvable Delegator.
func NewDelegation(d *types.Delegation, repo repository.Repository) *Delegation {
	return &Delegation{
		Delegation: *d,
		repo:       repo,
	}
}

// Delegation resolves details of a delegator by it's address.
func (rs *rootResolver) Delegation(args *struct {
	Address common.Address
	Staker  hexutil.Uint64
}) (*Delegation, error) {
	// get the delegator detail from backend
	d, err := rs.repo.Delegation(args.Address, args.Staker)
	if err != nil {
		return nil, err
	}

	return NewDelegation(d, rs.repo), nil
}

// Amount returns total delegated amount for the delegator.
func (del Delegation) Amount() (hexutil.Big, error) {
	// lazy load data
	if del.extendedAmount == nil {
		var err error

		// try to load the data
		del.extendedAmount, del.extendedAmountInWithdraw, err = del.repo.DelegatedAmountExtended(&del.Delegation)
		if err != nil {
			return hexutil.Big{}, err
		}
	}

	return (hexutil.Big)(*del.extendedAmount), nil
}

// AmountInWithdraw returns total delegated amount in pending withdrawals for the delegator.
func (del Delegation) AmountInWithdraw() (hexutil.Big, error) {
	// lazy load data
	if del.extendedAmountInWithdraw == nil {
		var err error

		// try to load the data
		del.extendedAmount, del.extendedAmountInWithdraw, err = del.repo.DelegatedAmountExtended(&del.Delegation)
		if err != nil {
			return hexutil.Big{}, err
		}
	}

	return (hexutil.Big)(*del.extendedAmountInWithdraw), nil
}

// PendingRewards resolves pending rewards for the delegator account.
func (del Delegation) PendingRewards() (types.PendingRewards, error) {
	// get the rewards
	r, err := del.repo.DelegationRewards(&del.Address, del.ToStakerId)
	if err != nil {
		return types.PendingRewards{}, err
	}

	return r, nil
}

// WithdrawRequests resolves partial withdraw requests of the delegator.
func (del Delegation) WithdrawRequests() ([]WithdrawRequest, error) {
	// pull the requests list from remote server
	wr, err := del.repo.WithdrawRequests(&del.Address, &del.ToStakerId)
	if err != nil {
		return nil, err
	}

	// create new result set
	list := make([]WithdrawRequest, 0)

	// sort the list
	sort.Sort(types.WithdrawRequestsByAge(wr))

	// iterate over the sorted list and populate the output array
	for _, req := range wr {
		list = append(list, NewWithdrawRequest(req, del.repo))
	}

	// return the final resolvable list
	return list, nil
}

// Deactivation resolves deactivated delegation requests of the delegator.
func (del Delegation) Deactivation() ([]DeactivatedDelegation, error) {
	// pull the requests list from remote server
	wr, err := del.repo.DeactivatedDelegation(&del.Address, &del.ToStakerId)
	if err != nil {
		return nil, err
	}

	// sort the list
	sort.Sort(types.DeactivatedDelegationByAge(wr))

	// create new result set
	list := make([]DeactivatedDelegation, 0)

	// iterate over the sorted list and populate the output array
	for _, req := range wr {
		list = append(list, NewDeactivatedDelegation(req, del.repo))
	}

	// return the final resolvable list
	return list, nil
}

// DelegationLock returns information about delegation lock
func (del Delegation) DelegationLock() *types.DelegationLock {
	if nil == del.lock && 0 < del.ToStakerId {
		var err error
		del.lock, err = del.repo.DelegationLock(&del.Delegation)
		if err != nil {
			return nil
		}
	}

	return del.lock
}

// IsDelegationLocked signals if the delegation is locked right now.
func (del Delegation) IsDelegationLocked() bool {
	// get the lock
	lock := del.DelegationLock()
	if lock == nil {
		return false
	}

	// decide based on lock content
	return lock.LockedFromEpoch > 0 && uint64(lock.LockedUntil) > uint64(time.Now().UTC().Unix())
}

// IsFluidStakingActive signals if the delegation is upgraded to Fluid Staking model.
func (del Delegation) IsFluidStakingActive() bool {
	// get the delegation fluid upgrade data
	fluid, err := del.repo.DelegationFluidStakingActive(&del.Delegation)
	if err != nil {
		return false
	}

	return fluid
}

// PaidUntilEpoch resolves the id of the last epoch rewards has been paid to."
func (del Delegation) PaidUntilEpoch() hexutil.Uint64 {
	// get the delegation fluid upgrade data
	paid, err := del.repo.DelegationPaidUntilEpoch(&del.Delegation)
	if err != nil {
		return 0
	}

	return paid
}

// LockedUntil resolves the end time of delegation.
func (del Delegation) LockedUntil() hexutil.Uint64 {
	// get the lock
	lock := del.DelegationLock()
	if lock == nil {
		return hexutil.Uint64(0)
	}

	// return the lock release time stamp
	return lock.LockedUntil
}

// LockedFromEpoch resolves the epoch om which the lock has been created.
func (del Delegation) LockedFromEpoch() hexutil.Uint64 {
	// get the lock
	lock := del.DelegationLock()
	if lock == nil {
		return hexutil.Uint64(0)
	}

	// return the lock creation epoch id
	return lock.LockedFromEpoch
}

// OutstandingSFTM resolves the amount of outstanding sFTM tokens
// minted for this account.
func (del Delegation) OutstandingSFTM() (hexutil.Big, error) {
	return del.repo.DelegationOutstandingSFTM(&del.Address, &del.ToStakerId)
}

// TokenizerAllowedToWithdraw resolves the tokenizer approval
// of the delegation withdrawal.
func (del Delegation) TokenizerAllowedToWithdraw() bool {
	// check the tokenizer lock status
	lock, err := del.repo.DelegationTokenizerUnlocked(&del.Address, &del.ToStakerId)
	if err != nil {
		del.repo.Log().Criticalf("can not check SFC tokenizer status for %s / %d",
			del.Address.String(), uint64(del.ToStakerId))
		return false
	}

	return lock
}

// Len returns size of the delegation list.
func (d DelegationsByAge) Len() int {
	return len(d)
}

// Less compares two delegations and returns true if the first is lower than the last.
// We use it to sort delegations by time created having newer on top.
func (d DelegationsByAge) Less(i, j int) bool {
	return uint64(d[i].CreatedTime) > uint64(d[j].CreatedTime)
}

// Swap changes position of two delegations in the list.
func (d DelegationsByAge) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}
