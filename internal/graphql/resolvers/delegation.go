// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"golang.org/x/sync/singleflight"
	"math/big"
	"strings"
	"time"
)

// Delegation represents resolvable delegator detail.
type Delegation struct {
	types.Delegation
	cg *singleflight.Group
}

// delegationLockSafetyWallSec represents number of seconds used as a safety wall
// for delegation lock evaluation
const delegationLockSafetyWallSec = 300

// NewDelegation creates new instance of resolvable Delegator.
func NewDelegation(d *types.Delegation) *Delegation {
	return &Delegation{Delegation: *d, cg: new(singleflight.Group)}
}

// Delegation resolves details of a delegator by it's address.
func (rs *rootResolver) Delegation(args *struct {
	Address common.Address
	Staker  hexutil.Big
}) (*Delegation, error) {
	// get the delegator detail from backend
	d, err := repository.R().Delegation(&args.Address, &args.Staker)
	if err != nil {
		return nil, err
	}

	return NewDelegation(d), nil
}

// IsSelfStake checks if the delegation is actually a self stake
// of the validator.
func (del Delegation) IsSelfStake() bool {
	return strings.EqualFold(del.Address.String(), del.ToStakerAddress.String())
}

// AmountDelegated resolves the active amount the delegation stands for.
func (del Delegation) AmountDelegated() hexutil.Big {
	if del.Delegation.AmountDelegated == nil {
		return hexutil.Big{}
	}
	return *del.Delegation.AmountDelegated
}

// ToStakerId resolves validator ID the delegation belongs to.
func (del Delegation) ToStakerId() hexutil.Big {
	if del.Delegation.ToStakerId == nil {
		return hexutil.Big{}
	}
	return *del.Delegation.ToStakerId
}

// Amount returns total delegated amount for the delegator.
func (del Delegation) Amount() (hexutil.Big, error) {
	// get the base amount delegated
	base, err := repository.R().DelegationAmountStaked(&del.Address, del.Delegation.ToStakerId)
	if err != nil {
		return hexutil.Big{}, err
	}

	// get the sum of all pending withdrawals
	wd, err := del.pendingWithdrawalsValue()
	if err != nil {
		return hexutil.Big{}, err
	}
	val := new(big.Int).Add(base, wd)
	return (hexutil.Big)(*val), nil
}

// pendingWithdrawalsValue returns total amount of tokens
// locked in pending withdrawals for the delegation.
func (del Delegation) pendingWithdrawalsValue() (*big.Int, error) {
	// call for it only once
	val, err, _ := del.cg.Do("withdraw-total", func() (interface{}, error) {
		return repository.R().WithdrawRequestsPendingTotal(&del.Address, del.Delegation.ToStakerId)
	})
	if err != nil {
		return nil, err
	}
	return val.(*big.Int), nil
}

// AmountInWithdraw returns total delegated amount in pending withdrawals for the delegator.
func (del Delegation) AmountInWithdraw() (hexutil.Big, error) {
	val, err := del.pendingWithdrawalsValue()
	if err != nil {
		return hexutil.Big{}, err
	}
	return (hexutil.Big)(*val), nil
}

// PendingRewards resolves pending rewards for the delegator account.
func (del Delegation) PendingRewards() (types.PendingRewards, error) {
	r, err := repository.R().PendingRewards(&del.Address, del.Delegation.ToStakerId)
	if err != nil {
		return types.PendingRewards{}, err
	}
	return *r, nil
}

// ClaimedReward resolves the total amount of rewards received on the delegation.
func (del Delegation) ClaimedReward() (hexutil.Big, error) {
	val, err := repository.R().RewardsClaimed(&del.Address, (*big.Int)(del.Delegation.ToStakerId), nil, nil)
	if err != nil {
		return hexutil.Big{}, err
	}
	return (hexutil.Big)(*val), nil
}

// WithdrawRequests resolves partial withdraw requests of the delegator.
func (del Delegation) WithdrawRequests(args struct {
	Cursor     *Cursor
	Count      int32
	ActiveOnly bool
}) ([]WithdrawRequest, error) {
	// limit query size; the count can be either positive or negative
	// this controls the loading direction
	args.Count = listLimitCount(args.Count, listMaxEdgesPerRequest)

	// pull list of withdrawals
	wr, err := repository.R().WithdrawRequests(&del.Address, del.Delegation.ToStakerId, (*string)(args.Cursor), args.Count, args.ActiveOnly)
	if err != nil {
		return nil, err
	}

	// create withdrawals list from the collection
	list := make([]WithdrawRequest, len(wr.Collection))
	for i, req := range wr.Collection {
		list[i] = NewWithdrawRequest(req)
	}

	// return the final resolvable list
	return list, nil
}

// RewardClaims resolves list of reward claims of the delegation.
func (del Delegation) RewardClaims(args struct {
	Cursor *Cursor
	Count  int32
}) (*RewardClaimList, error) {
	// limit query size; the count can be either positive or negative
	// this controls the loading direction
	args.Count = listLimitCount(args.Count, listMaxEdgesPerRequest)

	// pull list of withdrawals
	cl, err := repository.R().RewardClaims(&del.Address, (*big.Int)(del.Delegation.ToStakerId), (*string)(args.Cursor), args.Count)
	if err != nil {
		return nil, err
	}

	// return the final resolvable list
	return NewRewardClaimList(cl), nil
}

// DelegationLock returns information about delegation lock
func (del Delegation) DelegationLock() (*types.DelegationLock, error) {
	// load the delegations lock only once
	dl, err, _ := del.cg.Do("lock", func() (interface{}, error) {
		return repository.R().DelegationLock(&del.Address, del.Delegation.ToStakerId)
	})
	if err != nil {
		return nil, err
	}
	return dl.(*types.DelegationLock), nil
}

// IsDelegationLocked signals if the delegation is locked right now.
func (del Delegation) IsDelegationLocked() (bool, error) {
	lock, err := del.DelegationLock()
	if err != nil {
		return false, err
	}
	return lock != nil && 0 > zeroInt.Cmp((*big.Int)(&lock.LockedAmount)) && uint64(lock.LockedUntil) > uint64(time.Now().UTC().Unix()-delegationLockSafetyWallSec), nil
}

// IsFluidStakingActive signals if the delegation is upgraded to Fluid Staking model.
func (del Delegation) IsFluidStakingActive() (bool, error) {
	return repository.R().DelegationFluidStakingActive(&del.Address, del.Delegation.ToStakerId)
}

// LockedUntil resolves the end time of delegation.
func (del Delegation) LockedUntil() (hexutil.Uint64, error) {
	lock, err := del.DelegationLock()
	if err != nil {
		return hexutil.Uint64(0), err
	}
	// is there any lock in place?
	if lock == nil || 0 <= zeroInt.Cmp(lock.LockedAmount.ToInt()) {
		return hexutil.Uint64(0), nil
	}
	return lock.LockedUntil, nil
}

// LockDuration resolves the original duration of the active delegation lock.
func (del Delegation) LockDuration() (hexutil.Uint64, error) {
	lock, err := del.DelegationLock()
	if err != nil {
		return 0, err
	}
	if lock == nil || 0 <= zeroInt.Cmp(lock.LockedAmount.ToInt()) {
		return hexutil.Uint64(0), nil
	}
	return lock.Duration, nil
}

// LockedFromEpoch resolves the epoch om which the lock has been created.
func (del Delegation) LockedFromEpoch() (hexutil.Uint64, error) {
	lock, err := del.DelegationLock()
	if err != nil {
		return hexutil.Uint64(0), err
	}
	// is there any lock in place?
	if lock == nil || 0 <= zeroInt.Cmp(lock.LockedAmount.ToInt()) {
		return hexutil.Uint64(0), nil
	}
	return lock.LockedFromEpoch, nil
}

// LockedAmount resolves the total amount of delegation locked.
func (del Delegation) LockedAmount() (hexutil.Big, error) {
	lock, err := del.DelegationLock()
	if err != nil {
		return hexutil.Big{}, err
	}
	return lock.LockedAmount, nil
}

// UnlockedAmount resolves the total amount of unlocked delegation
// which is available for un-delegate.
func (del Delegation) UnlockedAmount() (hexutil.Big, error) {
	return repository.R().DelegationAmountUnlocked(&del.Address, (*big.Int)(del.Delegation.ToStakerId))
}

// UnlockPenalty resolves the amount of penalty applied to the stake
// on premature unlock request.
func (del Delegation) UnlockPenalty(args struct{ Amount hexutil.Big }) (hexutil.Big, error) {
	return repository.R().DelegationUnlockPenalty(&del.Address, (*big.Int)(del.Delegation.ToStakerId), (*big.Int)(&args.Amount))
}

// OutstandingSFTM resolves the amount of outstanding sFTM tokens
// minted for this account.
func (del Delegation) OutstandingSFTM() (hexutil.Big, error) {
	val, err := repository.R().DelegationOutstandingSFTM(&del.Address, del.Delegation.ToStakerId)
	if err != nil {
		return hexutil.Big{}, err
	}
	return *val, nil
}

// TokenizerAllowedToWithdraw resolves the tokenizer approval
// of the delegation withdrawal.
func (del Delegation) TokenizerAllowedToWithdraw() (bool, error) {
	// check the tokenizer lock status
	lock, err := repository.R().DelegationTokenizerUnlocked(&del.Address, del.Delegation.ToStakerId)
	if err != nil {
		return false, err
	}
	return lock, nil
}
