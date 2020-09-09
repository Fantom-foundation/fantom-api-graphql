// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// CurrentState represents resolvable state detail.
type CurrentState struct {
	repo repository.Repository
}

// State resolves details of the current state of the blockchain and network.
func (rs *rootResolver) State() (CurrentState, error) {
	return CurrentState{
		repo: rs.repo,
	}, nil
}

// SealedEpoch resolves the most recent sealed epoch details.
func (cst CurrentState) SealedEpoch() (types.Epoch, error) {
	// get the sealed epoch
	e, err := cst.repo.CurrentSealedEpoch()
	if err != nil {
		return types.Epoch{}, err
	}

	return *e, nil
}

// Validators resolves the number of validators active in the network.
func (cst CurrentState) Validators() (hexutil.Uint64, error) {
	return cst.repo.StakersNum()
}

// Accounts resolves the number of accounts participating on chain transactions.
func (cst CurrentState) Accounts() (hexutil.Uint64, error) {
	return cst.repo.AccountsActive()
}

// Blocks resolves the total number of blocks in the chain.
func (cst CurrentState) Blocks() (hexutil.Uint64, error) {
	// get the block height of the chain
	h, err := cst.repo.BlockHeight()
	if err != nil {
		return hexutil.Uint64(0), err
	}

	// do we have the block?
	if nil == h || !h.ToInt().IsInt64() {
		return hexutil.Uint64(0), fmt.Errorf("block height not available")
	}

	return hexutil.Uint64(h.ToInt().Uint64()), nil
}

// Transactions resolves the total number of transactions in the chain.
func (cst CurrentState) Transactions() (hexutil.Uint64, error) {
	return cst.repo.TransactionsCount()
}

// SfcLockingEnabled indicates if the stake locking has been enabled in SFC contract.
func (cst CurrentState) SfcLockingEnabled() (bool, error) {
	return cst.repo.LockingAllowed()
}

// SfcVersion resolves the current version of the SFC contract on the connected node.
func (cst CurrentState) SfcVersion() (hexutil.Uint64, error) {
	return cst.repo.SfcVersion()
}
