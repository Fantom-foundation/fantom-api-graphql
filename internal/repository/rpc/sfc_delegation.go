/*
Package rpc implements bridge to Lachesis full node API interface.

We recommend using local IPC for fast and the most efficient inter-process communication between the API server
and an Opera/Lachesis node. Any remote RPC connection will work, but the performance may be significantly degraded
by extra networking overhead of remote RPC calls.

You should also consider security implications of opening Lachesis RPC interface for a remote access.
If you considering it as your deployment strategy, you should establish encrypted channel between the API server
and Lachesis RPC interface with connection limited to specified endpoints.

We strongly discourage opening Lachesis RPC interface for unrestricted Internet access.
*/
package rpc

import (
	"fantom-api-graphql/internal/repository/rpc/contracts"
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// AmountStaked returns the current amount at stake for the given staker address and target validator
func (ftm *FtmBridge) AmountStaked(addr *common.Address, valID *big.Int) (*big.Int, error) {
	// keep track of the operation
	ftm.log.Debugf("verifying amount staked by %s to %d", addr.String(), valID.Uint64())

	// instantiate the contract and display its name
	contract, err := contracts.NewSfcContract(ftm.sfcConfig.SFCContract, ftm.eth)
	if err != nil {
		ftm.log.Criticalf("failed to instantiate SFC contract; %s", err.Error())
		return nil, err
	}

	// get the amount staked
	return contract.GetStake(nil, *addr, valID)
}

// AmountStakeLocked returns the current locked amount at stake for the given staker address and target validator
func (ftm *FtmBridge) AmountStakeLocked(addr *common.Address, valID *big.Int) (*big.Int, error) {
	// keep track of the operation
	ftm.log.Debugf("verifying amount locked by %s to %d", addr.String(), valID.Uint64())

	// instantiate the contract and display its name
	contract, err := contracts.NewSfcContract(ftm.sfcConfig.SFCContract, ftm.eth)
	if err != nil {
		ftm.log.Criticalf("failed to instantiate SFC contract; %s", err.Error())
		return nil, err
	}

	// get the amount staked
	return contract.GetLockedStake(nil, *addr, valID)
}

// PendingRewards returns a detail of delegation rewards waiting to be claimed for the given delegation.
func (ftm *FtmBridge) PendingRewards(addr *common.Address, valID *big.Int) (*types.PendingRewards, error) {
	// prep the empty value
	pr := types.PendingRewards{
		Staker: hexutil.Uint64(valID.Uint64()),
		Amount: hexutil.Big{},
	}

	// instantiate the contract and display its name
	contract, err := contracts.NewSfcContract(ftm.sfcConfig.SFCContract, ftm.eth)
	if err != nil {
		ftm.log.Criticalf("failed to instantiate SFC contract; %s", err.Error())
		return &pr, nil
	}

	// get the pending rewards amount
	amo, err := contract.PendingRewards(nil, *addr, valID)
	if err != nil {
		ftm.log.Criticalf("can not calculate pending rewards of %s to %d; %s", addr.String(), valID.Uint64(), err.Error())
		return &pr, nil
	}

	// update the amount
	pr.Amount = hexutil.Big(*amo)
	return &pr, nil
}

// DelegationLock returns delegation lock information using SFC contract binding.
func (ftm *FtmBridge) DelegationLock(dlg *types.Delegation) (dll *types.DelegationLock, err error) {
	// recover from panic here
	defer func() {
		if r := recover(); r != nil {
			ftm.log.Criticalf("can not get SFC lock status on delegation %s to %d; SFC call panic", dlg.Address.String(), dlg.ToStakerId.String())
			dll = &types.DelegationLock{}
		}
	}()

	// instantiate the contract
	contract, err := contracts.NewSfcContract(ftm.sfcConfig.SFCContract, ftm.eth)
	if err != nil {
		ftm.log.Criticalf("failed to instantiate SFC contract; %s", err.Error())
		return nil, err
	}

	// get staker locking detail
	lock, err := contract.GetLockupInfo(nil, dlg.Address, dlg.ToStakerId.ToInt())
	if err != nil {
		ftm.log.Errorf("delegation lock query failed; %v", err)
		return nil, err
	}

	// are lock timers available?
	if lock.FromEpoch == nil || lock.EndTime == nil {
		ftm.log.Errorf("delegation lock details not available")
		return nil, fmt.Errorf("delegation lock missing")
	}

	// make a new delegation lock
	return &types.DelegationLock{
		LockedAmount:    hexutil.Big(*lock.LockedStake),
		LockedFromEpoch: hexutil.Uint64(lock.FromEpoch.Uint64()),
		LockedUntil:     hexutil.Uint64(lock.EndTime.Uint64()),
		Duration:        hexutil.Uint64(lock.Duration.Uint64()),
	}, nil
}

// DelegationFluidStakingActive signals if the delegation is upgraded to Fluid Staking model.
func (ftm *FtmBridge) DelegationFluidStakingActive(_ *types.Delegation) (bool, error) {
	ftm.log.Debug("fluid staking is always on with SFCv3")
	return true, nil
}

// DelegationOutstandingSFTM returns the amount of sFTM tokens for the delegation
// identified by the delegator address and the stakerId.
func (ftm *FtmBridge) DelegationOutstandingSFTM(addr *common.Address, valID *big.Int) (*big.Int, error) {
	// log action
	ftm.log.Debugf("checking outstanding sFTM of %s to %d", addr.String(), valID.Uint64())

	// instantiate the contract and display its name
	contract, err := contracts.NewSfcTokenizer(ftm.sfcConfig.TokenizerContract, ftm.eth)
	if err != nil {
		ftm.log.Criticalf("failed to instantiate SFC Tokenizer contract; %s", err.Error())
		return nil, err
	}

	// get the amount of outstanding sFTM
	return contract.OutstandingSFTM(nil, *addr, valID)
}

// DelegationTokenizerUnlocked returns the status of SFC Tokenizer lock
// for a delegation identified by the address and staker id.
func (ftm *FtmBridge) DelegationTokenizerUnlocked(addr *common.Address, valID *big.Int) (bool, error) {
	// log action
	ftm.log.Debugf("checking SFC tokenizer lock of %s to %d", addr.String(), valID.Uint64())

	// instantiate the contract and display its name
	contract, err := contracts.NewSfcTokenizer(ftm.sfcConfig.TokenizerContract, ftm.eth)
	if err != nil {
		ftm.log.Criticalf("failed to instantiate SFC Tokenizer contract: %s", err.Error())
		return false, err
	}

	// get the lock status
	lock, err := contract.AllowedToWithdrawStake(nil, *addr, valID)
	if err != nil {
		ftm.log.Criticalf("failed to get SFC Tokenizer lock status of %s to %d; %s", addr.String(), valID.Uint64(), err.Error())
		return false, err
	}

	return lock, nil
}
