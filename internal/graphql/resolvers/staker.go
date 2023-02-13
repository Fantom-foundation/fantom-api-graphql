// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"golang.org/x/sync/singleflight"
	"math/big"
	"time"
)

const (
	// call group keys used to prevent parallel pull of the same value
	// in multithreading processing environment
	stakerCallGroupLock          = "lock"
	stakerCallGroupStake         = "stake"
	stakerCallGroupMaxDelegation = "max_delegation"
	stakerCallGroupDowntime      = "down"

	// SFC status bits
	sfcStatusWithdrawn  = 1
	sfcStatusOffline    = 1 << 3
	sfcStatusDoubleSign = 1 << 7
)

// Staker represents resolvable staker record.
type Staker struct {
	types.Validator
	cg *singleflight.Group
}

// NewStaker creates a new resolvable staker structure
func NewStaker(st *types.Validator) *Staker {
	if nil == st {
		return nil
	}
	return &Staker{Validator: *st, cg: new(singleflight.Group)}
}

// Delegations resolves list of delegations associated with the staker.
func (st Staker) Delegations(args struct {
	Cursor *Cursor
	Count  int32
}) (*DelegationList, error) {
	// limit query size; the count can be either positive or negative
	// this controls the loading direction
	args.Count = listLimitCount(args.Count, accMaxTransactionsPerRequest)

	// get delegations
	dl, err := repository.R().DelegationsOfValidator(&st.Id, (*string)(args.Cursor), args.Count)
	if err != nil {
		return nil, err
	}

	return NewDelegationList(dl), nil
}

// StakerInfo resolves extended staker information if available.
func (st Staker) StakerInfo() *types.StakerInfo {
	return repository.R().RetrieveStakerInfo(&st.Id)
}

// DelegationLock returns information about validator lock.
func (st Staker) DelegationLock() (*types.DelegationLock, error) {
	// load the delegations lock only once
	dl, err, _ := st.cg.Do(stakerCallGroupLock, func() (interface{}, error) {
		return repository.R().DelegationLock(&st.StakerAddress, &st.Id)
	})
	if err != nil {
		return nil, err
	}
	return dl.(*types.DelegationLock), nil
}

// IsStakeLocked signals if the stake is locked right now.
func (st Staker) IsStakeLocked() (bool, error) {
	lock, err := st.DelegationLock()
	if err != nil {
		return false, err
	}
	return lock != nil && 0 > zeroInt.Cmp(lock.LockedAmount.ToInt()) && uint64(lock.LockedUntil) < uint64(time.Now().UTC().Unix()), nil
}

// LockedUntil resolves the end time of delegation.
func (st Staker) LockedUntil() (hexutil.Uint64, error) {
	// get the lock detail
	lock, err := st.DelegationLock()
	if err != nil {
		return hexutil.Uint64(0), err
	}
	// is there any lock in place?
	if lock == nil || 0 <= zeroInt.Cmp(lock.LockedAmount.ToInt()) {
		return hexutil.Uint64(0), nil
	}
	return lock.LockedUntil, nil
}

// LockedFromEpoch resolves the epoch om which the lock has been created.
func (st Staker) LockedFromEpoch() (hexutil.Uint64, error) {
	lock, err := st.DelegationLock()
	if err != nil {
		return hexutil.Uint64(0), err
	}
	// is there any lock in place?
	if lock == nil || 0 <= zeroInt.Cmp(lock.LockedAmount.ToInt()) {
		return hexutil.Uint64(0), nil
	}
	return lock.LockedFromEpoch, nil
}

// WithdrawRequests resolves partial withdraw requests of the staker.
// We load withdraw requests of the stake only, not the stake delegators.
func (st Staker) WithdrawRequests() ([]WithdrawRequest, error) {
	// pull the requests list from remote server
	wwl, err := repository.R().WithdrawRequests(&st.StakerAddress, nil, nil, 50, false)
	if err != nil {
		return nil, err
	}

	// create new result set
	list := make([]WithdrawRequest, len(wwl.Collection))
	for i, req := range wwl.Collection {
		list[i] = NewWithdrawRequest(req)
	}

	// return the final resolvable list
	return list, nil
}

// Stake resolves the amount of self staked tokens.
func (st Staker) Stake() (hexutil.Big, error) {
	// load the delegations lock only once
	dl, err, _ := st.cg.Do(stakerCallGroupStake, func() (interface{}, error) {
		return repository.R().DelegationAmountStaked(&st.StakerAddress, &st.Id)
	})
	if err != nil {
		return hexutil.Big{}, err
	}
	return hexutil.Big(*dl.(*big.Int)), err
}

// DelegatedMe resolves the amount of tokens delegated to the validator
// without the self staked amount.
func (st Staker) DelegatedMe() (hexutil.Big, error) {
	// get the amount of self staked tokens
	sf, err := st.Stake()
	if err != nil {
		return hexutil.Big{}, err
	}

	// any total stake?
	if st.TotalStake == nil {
		return hexutil.Big{}, err
	}

	// make a sanity check for the corner amounts
	// the self stake must be lower value than the total stake
	if sf.ToInt().Cmp(st.TotalStake.ToInt()) >= 0 {
		return hexutil.Big{}, err
	}
	return hexutil.Big(*new(big.Int).Sub(st.TotalStake.ToInt(), sf.ToInt())), nil
}

// TotalDelegatedLimit resolves the total max amount of tokens delegated
// to the validator including the self stake.
func (st Staker) TotalDelegatedLimit() (hexutil.Big, error) {
	// calculate the delegation limit
	lim, err, _ := st.cg.Do(stakerCallGroupMaxDelegation, func() (interface{}, error) {
		// pull the amount of self staked tokens
		self, err := st.Stake()
		if err != nil {
			return hexutil.Big{}, err
		}

		// pull the staking ratio
		ratio, err := repository.R().SfcMaxDelegatedRatio()
		if err != nil {
			return hexutil.Big{}, err
		}

		// calculate the value
		val := new(big.Int).Div(new(big.Int).Mul(self.ToInt(), ratio), repository.R().SfcDecimalUnit())
		return hexutil.Big(*val), nil
	})
	if err != nil {
		return hexutil.Big{}, err
	}
	return lim.(hexutil.Big), nil
}

// DelegatedLimit resolves the amount of tokens available to be delegated
// to the validator before their max delegation limit is reached
func (st Staker) DelegatedLimit() (hexutil.Big, error) {
	// get the total limit
	lim, err := st.TotalDelegatedLimit()
	if err != nil {
		return hexutil.Big{}, err
	}

	// make sure the limit is bigger than the current received stake
	if lim.ToInt().Cmp(st.TotalStake.ToInt()) <= 0 {
		return hexutil.Big{}, nil
	}
	return hexutil.Big(*new(big.Int).Sub(lim.ToInt(), st.TotalStake.ToInt())), nil
}

// IsActive signals if the validator is active.
func (st Staker) IsActive() bool {
	return st.Status == 0
}

// IsWithdrawn signals if the validator has withdrawn from the validators.
func (st Staker) IsWithdrawn() bool {
	return st.Status&sfcStatusWithdrawn > 0
}

// IsCheater signals if the validator has a double sign tag active.
func (st Staker) IsCheater() bool {
	return st.Status&sfcStatusDoubleSign > 0
}

// IsOffline signals if the validator is off-line.
func (st Staker) IsOffline() bool {
	return st.Status&sfcStatusOffline > 0
}

// Downtime resolves the amount of time a validator is offline.
func (st Staker) Downtime() (hexutil.Uint64, error) {
	tm, _, err := st.downtime()
	if err != nil {
		return 0, err
	}
	return hexutil.Uint64(tm), err
}

// MissedBlocks resolves the amount of blocks a validator missed recently.
func (st Staker) MissedBlocks() (hexutil.Uint64, error) {
	_, blk, err := st.downtime()
	if err != nil {
		return 0, err
	}
	return hexutil.Uint64(blk), err
}

// downtime pulls information about the validator down time and missed blocks from aBFT API.
func (st Staker) downtime() (uint64, uint64, error) {
	// how the call group responds
	type dt struct {
		Time   uint64
		Blocks uint64
	}

	// pull the values
	val, err, _ := st.cg.Do(stakerCallGroupDowntime, func() (interface{}, error) {
		dtm, blocks, err := repository.R().ValidatorDowntime(&st.Id)
		if err != nil {
			return dt{}, err
		}
		return dt{
			Time:   dtm,
			Blocks: blocks,
		}, nil
	})
	if err != nil {
		return 0, 0, err
	}
	return val.(dt).Time, val.(dt).Blocks, err
}
