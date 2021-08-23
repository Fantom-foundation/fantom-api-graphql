// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/config"
	"fantom-api-graphql/internal/repository"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// CurrentState represents resolvable state detail.
type CurrentState struct {
	config.Staking
}

// State resolves details of the current state of the blockchain and network.
func (rs *rootResolver) State() (CurrentState, error) {
	return CurrentState{cfg.Staking}, nil
}

// SealedEpoch resolves the most recent sealed epoch details.
func (cst CurrentState) SealedEpoch() (Epoch, error) {
	// get the sealed epoch
	e, err := repository.R().CurrentSealedEpoch()
	if err != nil {
		return Epoch{}, err
	}
	return Epoch{*e}, nil
}

// Validators resolves the number of validators active in the network.
func (cst CurrentState) Validators() (hexutil.Uint64, error) {
	val, err := repository.R().ValidatorsCount()
	return hexutil.Uint64(val), err
}

// Accounts resolves the number of accounts participating on chain transactions.
func (cst CurrentState) Accounts() (hexutil.Uint64, error) {
	return repository.R().AccountsActive()
}

// Blocks resolves the total number of blocks in the chain.
func (cst CurrentState) Blocks() (hexutil.Big, error) {
	// get the block height of the chain
	h, err := repository.R().BlockHeight()
	if err != nil {
		return hexutil.Big{}, err
	}
	return *h, nil
}

// Transactions resolves the total number of transactions in the chain.
func (cst CurrentState) Transactions() (hexutil.Uint64, error) {
	return repository.R().EstimateTransactionsCount()
}

// SfcContractAddress resolves address of the SFC contract.
func (cst CurrentState) SfcContractAddress() common.Address {
	return cst.SFCContract
}

// SfcLockingEnabled indicates if the stake locking has been enabled in SFC contract.
func (cst CurrentState) SfcLockingEnabled() (bool, error) {
	return repository.R().LockingAllowed()
}

// SfcVersion resolves the current version of the SFC contract on the connected node.
func (cst CurrentState) SfcVersion() (hexutil.Uint64, error) {
	return repository.R().SfcVersion()
}
