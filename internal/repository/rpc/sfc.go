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

//go:generate abigen --abi ./contracts/abi/sfc-1.1.abi --pkg contracts --type SfcV1Contract --out ./contracts/sfc-v1.go
//go:generate abigen --abi ./contracts/abi/sfc-2.0.2-rc2.abi --pkg contracts --type SfcContract --out ./contracts/sfc-v2.go
//go:generate abigen --abi ./contracts/abi/sfc-tokenizer.abi --pkg contracts --type SfcTokenizer --out ./contracts/sfc_tokenizer.go

import (
	"fantom-api-graphql/internal/repository/rpc/contracts"
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"strconv"
)

// SfcVersion returns current version of the SFC contract as a single number.
func (ftm *FtmBridge) SfcVersion() (hexutil.Uint64, error) {
	// instantiate the contract and display its name
	contract, err := contracts.NewSfcContract(ftm.sfcConfig.SFCContract, ftm.eth)
	if err != nil {
		ftm.log.Criticalf("failed to instantiate SFC contract: %v", err)
		return 0, err
	}

	// get the version information from the contract
	var ver [3]byte
	ver, err = contract.Version(nil)
	if err != nil {
		ftm.log.Criticalf("failed to get the SFC version; %v", err)
		return 0, err
	}

	return hexutil.Uint64((uint64(ver[0]) << 16) | (uint64(ver[1]) << 8) | uint64(ver[2])), nil
}

// CurrentEpoch extract the current epoch id from SFC smart contract.
func (ftm *FtmBridge) CurrentEpoch() (hexutil.Uint64, error) {
	// instantiate the contract and display its name
	contract, err := contracts.NewSfcContract(ftm.sfcConfig.SFCContract, ftm.eth)
	if err != nil {
		ftm.log.Criticalf("failed to instantiate SFC contract: %v", err)
		return 0, err
	}

	// get the value from the contract
	epoch, err := contract.CurrentEpoch(nil)
	if err != nil {
		ftm.log.Errorf("failed to get the current epoch: %v", err)
		return 0, err
	}

	// get the value
	return hexutil.Uint64(epoch.Uint64()), nil
}

// CurrentSealedEpoch extract the current sealed epoch id from SFC smart contract.
func (ftm *FtmBridge) CurrentSealedEpoch() (hexutil.Uint64, error) {
	// instantiate the contract and display its name
	contract, err := contracts.NewSfcContract(ftm.sfcConfig.SFCContract, ftm.eth)
	if err != nil {
		ftm.log.Criticalf("failed to instantiate SFC contract: %v", err)
		return 0, err
	}

	// get the value from the contract
	epoch, err := contract.CurrentSealedEpoch(nil)
	if err != nil {
		ftm.log.Errorf("failed to get the current sealed epoch: %v", err)
		return 0, err
	}

	// get the value
	return hexutil.Uint64(epoch.Uint64()), nil
}

// LastStakerId returns the last staker id in Opera blockchain.
func (ftm *FtmBridge) LastStakerId() (hexutil.Uint64, error) {
	// instantiate the contract and display its name
	contract, err := contracts.NewSfcContract(ftm.sfcConfig.SFCContract, ftm.eth)
	if err != nil {
		ftm.log.Criticalf("failed to instantiate SFC contract: %v", err)
		return 0, err
	}

	// get the value from the contract
	sl, err := contract.StakersLastID(nil)
	if err != nil {
		ftm.log.Errorf("failed to get the last staker ID: %v", err)
		return 0, err
	}

	// get the value
	return hexutil.Uint64(sl.Uint64()), nil
}

// StakersNum returns the number of stakers in Opera blockchain.
func (ftm *FtmBridge) StakersNum() (hexutil.Uint64, error) {
	// instantiate the contract and display its name
	contract, err := contracts.NewSfcContract(ftm.sfcConfig.SFCContract, ftm.eth)
	if err != nil {
		ftm.log.Criticalf("failed to instantiate SFC contract: %v", err)
		return 0, err
	}

	// get the value from the contract
	sn, err := contract.StakersNum(nil)
	if err != nil {
		ftm.log.Errorf("failed to get the current number of stakers: %v", err)
		return 0, err
	}

	// get the value
	return hexutil.Uint64(sn.Uint64()), nil
}

// stakerStatusFromSfc updates staker information using SFC binding.
func (ftm *FtmBridge) stakerStatusFromSfc(contract *contracts.SfcContract, staker *types.Staker) error {
	// log action
	ftm.log.Debug("updating staker info from SFC")

	// get the value from the contract
	si, err := contract.Stakers(nil, big.NewInt(int64(staker.Id)))
	if err != nil {
		ftm.log.Errorf("failed to get the staker information from SFC: %v", err)
		return err
	}

	// do we have a valid record?
	if si.Status != nil {
		// update some invalid information
		staker.DelegatedMe = (*hexutil.Big)(si.DelegatedMe)
		staker.Stake = (*hexutil.Big)(si.StakeAmount)
		staker.Status = hexutil.Uint64(si.Status.Uint64())

		if staker.Stake != nil && staker.DelegatedMe != nil {
			// recalculate the total stake
			staker.TotalStake = (*hexutil.Big)(big.NewInt(0).Add(si.DelegatedMe, si.StakeAmount))

			// calculate delegation limit
			staker.TotalDelegatedLimit = ftm.maxDelegatedLimit(staker.Stake, contract)

			// calculate available limit for staking
			val := new(big.Int).Sub((*big.Int)(&staker.TotalDelegatedLimit), (*big.Int)(staker.DelegatedMe))
			staker.DelegatedLimit = (hexutil.Big)(*val)
		}
	} else {
		// log issue
		ftm.log.Debug("staker info update from SFC failed, no data received")
	}

	// get the value
	return nil
}

// stakerLockFromSfc updates staker lock details using SFC binding.
func (ftm *FtmBridge) stakerLockFromSfc(contract *contracts.SfcContract, staker *types.Staker) error {
	// log action
	ftm.log.Debug("updating staker locking details from SFC")

	// get staker locking detail
	lock, err := contract.LockedStakes(nil, big.NewInt(int64(staker.Id)))
	if err != nil {
		ftm.log.Errorf("stake lock query failed; %v", err)
		return nil
	}

	// are lock timers available?
	if lock.FromEpoch == nil || lock.EndTime == nil {
		ftm.log.Errorf("stake lock details not available")
		return nil
	}

	// apply the lock values
	staker.LockedFromEpoch = hexutil.Uint64(lock.FromEpoch.Uint64())
	staker.LockedUntil = hexutil.Uint64(lock.EndTime.Uint64())

	// get the value
	return nil
}

// maxDelegatedLimit calculate maximum amount of tokens allowed to be delegated to a staker.
func (ftm *FtmBridge) maxDelegatedLimit(staked *hexutil.Big, contract *contracts.SfcContract) hexutil.Big {
	// if we don't know the staked amount, return zero
	if staked == nil {
		return (hexutil.Big)(*hexutil.MustDecodeBig("0x0"))
	}

	// ratio unit is used to calculate the value (1.000.000)
	// please note this formula is taken from SFC contract and can change
	ratioUnit := hexutil.MustDecodeBig("0xF4240")

	// get delegation ration
	ratio, err := contract.MaxDelegatedRatio(nil)
	if err != nil {
		ftm.log.Errorf("can not get the delegation ratio; %s", err.Error())
		return (hexutil.Big)(*hexutil.MustDecodeBig("0x0"))
	}

	// calculate the delegation limit temp value
	temp := new(big.Int).Mul((*big.Int)(staked), ratio)

	// adjust to percent
	value := new(big.Int).Div(temp, ratioUnit)
	return (hexutil.Big)(*value)
}

// extendStaker extends staker information using SFC contract binding.
func (ftm *FtmBridge) extendStaker(staker *types.Staker) (*types.Staker, error) {
	// instantiate the contract and display its name
	contract, err := contracts.NewSfcContract(ftm.sfcConfig.SFCContract, ftm.eth)
	if err != nil {
		ftm.log.Criticalf("failed to instantiate SFC contract: %v", err)
		return nil, err
	}

	// update status detail
	err = ftm.stakerStatusFromSfc(contract, staker)
	if err != nil {
		ftm.log.Critical("staker status could not be updated from SFC")
	}

	// update locking detail
	err = ftm.stakerLockFromSfc(contract, staker)
	if err != nil {
		ftm.log.Critical("staker locking could not be updated from SFC")
	}

	return staker, nil
}

// Staker extract a staker information by numeric id.
func (ftm *FtmBridge) Staker(id hexutil.Uint64) (*types.Staker, error) {
	// keep track of the operation
	ftm.log.Debugf("loading staker #%d", id)

	// call for data
	var st types.Staker
	err := ftm.rpc.Call(&st, "sfc_getStaker", id, "0x2")
	if err != nil {
		ftm.log.Error("staker information could not be extracted")
		return nil, err
	}

	// keep track of the operation
	ftm.log.Debugf("staker #%d loaded", id)
	return ftm.extendStaker(&st)
}

// StakerAddress extract a staker address for the given staker ID.
func (ftm *FtmBridge) StakerAddress(id hexutil.Uint64) (common.Address, error) {
	// call for data
	var st types.Staker
	err := ftm.rpc.Call(&st, "sfc_getStaker", id, "0x0")
	if err != nil {
		ftm.log.Error("staker information could not be extracted")
		return common.Address{}, err
	}

	// keep track of the operation
	ftm.log.Debugf("staker #%d identified as %s", id, st.StakerAddress.String())
	return st.StakerAddress, nil
}

// IsStaker returns if the given address is an SFC staker.
func (ftm *FtmBridge) IsStaker(addr *common.Address) (bool, error) {
	// keep track of the operation
	ftm.log.Debugf("loading staker %s", addr.String())

	// check the status
	var st types.Staker
	err := ftm.rpc.Call(&st, "sfc_getStakerByAddress", *addr, "0x0")
	if err != nil {
		ftm.log.Error("staker information could not be extracted")
		return false, err
	}

	// keep track of the operation
	ftm.log.Debugf("staker %s verified to be %s", addr.String(), strconv.FormatBool(st.Id > 0))
	return st.Id > 0, nil
}

// StakerByAddress extracts a staker information by address.
func (ftm *FtmBridge) StakerByAddress(addr common.Address) (*types.Staker, error) {
	// keep track of the operation
	ftm.log.Debugf("loading staker %s", addr.String())

	// call for data
	var st types.Staker
	err := ftm.rpc.Call(&st, "sfc_getStakerByAddress", addr, "0x2")
	if err != nil {
		ftm.log.Error("staker information could not be extracted")
		return nil, err
	}

	// keep track of the operation
	ftm.log.Debugf("staker %s loaded", addr.String())
	return ftm.extendStaker(&st)
}

// Epoch extract information about an epoch from SFC smart contract.
func (ftm *FtmBridge) Epoch(id hexutil.Uint64) (types.Epoch, error) {
	// instantiate the contract and display its name
	contract, err := contracts.NewSfcContract(ftm.sfcConfig.SFCContract, ftm.eth)
	if err != nil {
		ftm.log.Criticalf("failed to instantiate SFC contract: %v", err)
		return types.Epoch{}, err
	}

	// extract epoch snapshot
	epo, err := contract.EpochSnapshots(nil, big.NewInt(int64(id)))
	if err != nil {
		ftm.log.Errorf("failed to extract epoch information: %v", err)
		return types.Epoch{}, err
	}

	return types.Epoch{
		Id:                     id,
		EndTime:                (hexutil.Big)(*epo.EndTime),
		Duration:               (hexutil.Big)(*epo.Duration),
		EpochFee:               (hexutil.Big)(*epo.EpochFee),
		TotalBaseRewardWeight:  (hexutil.Big)(*epo.TotalBaseRewardWeight),
		TotalTxRewardWeight:    (hexutil.Big)(*epo.TotalTxRewardWeight),
		BaseRewardPerSecond:    (hexutil.Big)(*epo.BaseRewardPerSecond),
		StakeTotalAmount:       (hexutil.Big)(*epo.StakeTotalAmount),
		DelegationsTotalAmount: (hexutil.Big)(*epo.DelegationsTotalAmount),
		TotalSupply:            (hexutil.Big)(*epo.TotalSupply),
	}, nil
}

// DelegationRewards returns a detail of delegation rewards for the given address.
func (ftm *FtmBridge) DelegationRewards(addr string, staker hexutil.Uint64) (types.PendingRewards, error) {
	// log action
	ftm.log.Debugf("loading delegation rewards for account %s", addr)

	// instantiate the contract and display its name
	contract, err := contracts.NewSfcContract(ftm.sfcConfig.SFCContract, ftm.eth)
	if err != nil {
		ftm.log.Criticalf("failed to instantiate SFC contract: %s", err.Error())
		return types.PendingRewards{}, err
	}

	// get the value from the contract
	epoch, err := contract.CurrentSealedEpoch(nil)
	if err != nil {
		ftm.log.Errorf("failed to get the current sealed epoch: %s", err.Error())
		return types.PendingRewards{}, err
	}

	// do we have the sealed epoch?
	if epoch == nil {
		ftm.log.Errorf("no epoch information available")
		return types.PendingRewards{}, nil
	}

	// get the rewards amount
	amount, fromEpoch, toEpoch, err := contract.CalcDelegationRewards(nil, common.HexToAddress(addr), big.NewInt(int64(staker)), big.NewInt(0), epoch)
	if err != nil {
		ftm.log.Errorf("no delegation rewards for %s -> %d; %s", addr, uint64(staker), err.Error())
		return types.PendingRewards{}, nil
	}

	// return the data
	return types.PendingRewards{
		Staker:    staker,
		Amount:    hexutil.Big(*amount),
		FromEpoch: hexutil.Uint64(fromEpoch.Uint64()),
		ToEpoch:   hexutil.Uint64(toEpoch.Uint64()),
	}, nil
}

// DelegationsOf extract a list of delegations for a given staker.
func (ftm *FtmBridge) DelegationsOf(staker hexutil.Uint64) ([]types.Delegation, error) {
	// keep track of the operation
	ftm.log.Debugf("loading delegations of staker %s", staker)

	// call for data
	dl := make([]types.Delegation, 0)
	err := ftm.rpc.Call(&dl, "sfc_getDelegationsOf", staker, "0x2")
	if err != nil {
		ftm.log.Error("delegations list could not be extracted")
		return nil, err
	}

	// keep track of the operation
	ftm.log.Debugf("delegations of staker %d loaded", staker)
	return dl, nil
}

// IsDelegating returns if the given address is an SFC delegator.
func (ftm *FtmBridge) IsDelegating(addr *common.Address) (bool, error) {
	// keep track of the operation
	ftm.log.Debugf("checking delegations of account %s", addr.String())

	// call for data
	dl := make([]types.Delegation, 0)
	err := ftm.rpc.Call(&dl, "sfc_getDelegationsByAddress", addr.String(), "0x0")
	if err != nil {
		ftm.log.Error("delegations list could not be extracted")
		return false, err
	}

	// keep track of the operation
	ftm.log.Debugf("delegation check of %s is %s", addr.String(), strconv.FormatBool(0 < len(dl)))
	return 0 < len(dl), nil
}

// DelegationsByAddress returns a list of all delegations
// of a given delegator address.
func (ftm *FtmBridge) DelegationsByAddress(addr common.Address) ([]types.Delegation, error) {
	// keep track of the operation
	ftm.log.Debugf("loading delegations of account %s", addr.String())

	// call for data
	dl := make([]types.Delegation, 0)
	err := ftm.rpc.Call(&dl, "sfc_getDelegationsByAddress", addr.String(), "0x2")
	if err != nil {
		ftm.log.Error("delegations list could not be extracted")
		return nil, err
	}

	// keep track of the operation
	ftm.log.Debugf("%d delegations of account %s loaded", len(dl), addr.String())
	return dl, nil
}

// Delegation returns a detail of delegation for the given address.
func (ftm *FtmBridge) Delegation(addr common.Address, staker hexutil.Uint64) (*types.Delegation, error) {
	// keep track of the operation
	ftm.log.Debugf("loading delegation of address %s with staked %d", addr.String(), staker)

	// call for data
	var dl types.Delegation
	err := ftm.rpc.Call(&dl, "sfc_getDelegation", addr, staker.String(), "0x2")
	if err != nil {
		ftm.log.Error("delegation not found")
		return nil, err
	}

	// keep track of the operation
	ftm.log.Debugf("delegation of address %s loaded", addr.String())
	return &dl, nil
}

// DelegationLock returns delegation lock information using SFC contract binding.
func (ftm *FtmBridge) DelegationLock(delegation *types.Delegation) (*types.DelegationLock, error) {
	// instantiate the contract
	contract, err := contracts.NewSfcContract(ftm.sfcConfig.SFCContract, ftm.eth)
	if err != nil {
		ftm.log.Criticalf("failed to instantiate SFC contract: %v", err)
		return nil, err
	}

	// get staker locking detail
	lock, err := contract.LockedDelegations(nil, delegation.Address, big.NewInt(int64(delegation.ToStakerId)))
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
	dl := new(types.DelegationLock)

	// apply the lock values
	dl.LockedFromEpoch = hexutil.Uint64(lock.FromEpoch.Uint64())
	dl.LockedUntil = hexutil.Uint64(lock.EndTime.Uint64())

	return dl, nil
}

// delegatedAmount calculates total amount currently delegated
// and amount locked in pending un-delegation.
// Partial Un-delegations are subtracted during the preparation
// phase, but total un-delegations are subtracted only when
// the delegation is closed.
func (ftm *FtmBridge) DelegatedAmountExtended(dl *types.Delegation) (*big.Int, *big.Int, error) {
	// base delegated amount is copied here
	amount := big.NewInt(0)
	inWithdraw := big.NewInt(0)

	// do we have any delegation active at all?
	if nil == dl.AmountDelegated {
		ftm.log.Debugf("no delegated amount for %s", dl.Address.String())
		return amount, inWithdraw, nil
	}

	// staker we are dealing with
	stakerId := big.NewInt(int64(dl.ToStakerId))

	// get the amount of pending withdrawals
	pw, err := ftm.PendingWithdrawalsAmount(&dl.Address, stakerId)
	if err != nil {
		ftm.log.Debugf("error loading pending withdrawals; %s", err.Error())
		return nil, nil, err
	}

	// add the pending withdrawal to the base amount value
	// since we want it to include these pending partial un-delegations
	// as well
	amount = new(big.Int).Add(dl.AmountDelegated.ToInt(), pw)

	// try to find pending deactivation for the current staker
	// since if there is one, the whole amount is going
	// to be reclaimed and hence is subject to pending withdrawal
	pd, err := ftm.PendingDeactivation(&dl.Address, stakerId)
	if err != nil {
		ftm.log.Debugf("error loading pending deactivation; %s", err.Error())
		return nil, nil, err
	}

	// calculate amount in withdraw
	if pd != nil {
		inWithdraw = new(big.Int).Add(pw, dl.AmountDelegated.ToInt())
	} else {
		inWithdraw = pw
	}

	// we may need to add partial un-delegations
	return amount, inWithdraw, nil
}

// Stashed returns amount of WEI stashed for the given address.
func (ftm *FtmBridge) Stashed(addr common.Address, stash *big.Int) (*big.Int, error) {
	// keep track of the operation
	ftm.log.Debugf("loading stashed amount of address %s", addr.String())

	// instantiate the contract and display its name
	contract, err := contracts.NewSfcContract(ftm.sfcConfig.SFCContract, ftm.eth)
	if err != nil {
		ftm.log.Criticalf("failed to instantiate SFC contract: %s", err.Error())
		return nil, err
	}

	val, err := contract.RewardsStash(nil, addr, stash)
	if err != nil {
		ftm.log.Errorf("failed to get rewards stash for %s: %v", addr.String(), err)
		return nil, err
	}

	// return the value
	return val, nil
}

// RewardsAllowed returns if the rewards can be manipulated with.
func (ftm *FtmBridge) RewardsAllowed() (bool, error) {
	ftm.log.Debug("rewards lock always open")
	return true, nil
}

// LockingAllowed indicates if the stake locking has been enabled in SFC.
func (ftm *FtmBridge) LockingAllowed() (bool, error) {
	// instantiate the contract and display its name
	contract, err := contracts.NewSfcContract(ftm.sfcConfig.SFCContract, ftm.eth)
	if err != nil {
		ftm.log.Criticalf("failed to instantiate SFC contract: %s", err.Error())
		return false, err
	}

	// get the current value of the first lock
	firstLock, err := contract.FirstLockedUpEpoch(nil)
	if err != nil {
		ftm.log.Errorf("failed to get the first epoch with enabled locking: %s", err.Error())
		return false, err
	}

	// get the current sealed epoch value from the contract
	epoch, err := contract.CurrentSealedEpoch(nil)
	if err != nil {
		ftm.log.Errorf("failed to get the current sealed epoch: %s", err.Error())
		return false, err
	}

	return firstLock.Uint64() > 0 && epoch.Uint64() >= firstLock.Uint64(), nil
}

// DelegationFluidStakingActive signals if the delegation is upgraded to Fluid Staking model.
func (ftm *FtmBridge) DelegationFluidStakingActive(dl *types.Delegation) (bool, error) {
	// try to get paid info
	paid, err := ftm.DelegationPaidUntilEpoch(dl)
	if err != nil {
		ftm.log.Errorf("delegation information missing; %s", err.Error())
		return false, err
	}

	return paid > 0, nil
}

// DelegationPaidUntilEpoch resolves the id of the last epoch rewards has been paid to."
func (ftm *FtmBridge) DelegationPaidUntilEpoch(dl *types.Delegation) (hexutil.Uint64, error) {
	// instantiate the contract and display its name
	contract, err := contracts.NewSfcContract(ftm.sfcConfig.SFCContract, ftm.eth)
	if err != nil {
		ftm.log.Criticalf("failed to instantiate SFC contract: %s", err.Error())
		return 0, err
	}

	// get delegation info
	dd, err := contract.Delegations(nil, dl.Address, big.NewInt(int64(dl.ToStakerId)))
	if err != nil {
		ftm.log.Errorf("can not pull delegation information for %s / %d; %s", dl.Address.String(), dl.ToStakerId, err.Error())
		return 0, err
	}

	// any data about paid epoch?
	if dd.PaidUntilEpoch == nil {
		ftm.log.Errorf("no paid epoch information for delegation %s / %d", dl.Address.String(), dl.ToStakerId)
		return 0, nil
	}

	return hexutil.Uint64(dd.PaidUntilEpoch.Uint64()), nil
}

// DelegationOutstandingSFTM returns the amount of sFTM tokens for the delegation
// identified by the delegator address and the stakerId.
func (ftm *FtmBridge) DelegationOutstandingSFTM(addr *common.Address, toStaker *hexutil.Uint64) (hexutil.Big, error) {
	// log action
	ftm.log.Debugf("checking outstanding sFTM on %s / %d", addr.String(), uint64(*toStaker))

	// instantiate the contract and display its name
	contract, err := contracts.NewSfcTokenizer(ftm.sfcConfig.TokenizerContract, ftm.eth)
	if err != nil {
		ftm.log.Criticalf("failed to instantiate SFC Tokenizer contract: %s", err.Error())
		return hexutil.Big{}, err
	}

	// get the amount of outstanding sFTM
	val, err := contract.OutstandingSFTM(nil, *addr, new(big.Int).SetUint64(uint64(*toStaker)))
	if err != nil {
		ftm.log.Criticalf("failed to get sFTM amount from %s/%d: %s",
			addr.String(), uint64(*toStaker), err.Error())
		return hexutil.Big{}, err
	}

	return hexutil.Big(*val), nil
}

// DelegationTokenizerUnlocked returns the status of SFC Tokenizer lock
// for a delegation identified by the address and staker id.
func (ftm *FtmBridge) DelegationTokenizerUnlocked(addr *common.Address, toStaker *hexutil.Uint64) (bool, error) {
	// log action
	ftm.log.Debugf("checking SFC tokenizer lock on %s / %d", addr.String(), uint64(*toStaker))

	// instantiate the contract and display its name
	contract, err := contracts.NewSfcTokenizer(ftm.sfcConfig.TokenizerContract, ftm.eth)
	if err != nil {
		ftm.log.Criticalf("failed to instantiate SFC Tokenizer contract: %s", err.Error())
		return false, err
	}

	// get the lock status
	lock, err := contract.AllowedToWithdrawStake(nil, *addr, new(big.Int).SetUint64(uint64(*toStaker)))
	if err != nil {
		ftm.log.Criticalf("failed to get SFC Tokenizer lock status for %s/%d: %s",
			addr.String(), uint64(*toStaker), err.Error())
		return false, err
	}

	return lock, nil
}
