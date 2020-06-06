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

//go:generate abigen --abi ./contracts/sfc.abi --pkg rpc --type SfcContract --out ./sfc_bind.go

import (
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// sfcContractAddress represents the address on which the Sfc contract is deployed.
var sfcContractAddress = common.HexToAddress("0xfc00face00000000000000000000000000000000")

// CurrentEpoch extract the current epoch id from SFC smart contract.
func (ftm *FtmBridge) CurrentEpoch() (hexutil.Uint64, error) {
	// instantiate the contract and display its name
	contract, err := NewSfcContract(sfcContractAddress, ftm.eth)
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
	contract, err := NewSfcContract(sfcContractAddress, ftm.eth)
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
	contract, err := NewSfcContract(sfcContractAddress, ftm.eth)
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
	contract, err := NewSfcContract(sfcContractAddress, ftm.eth)
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

// stakerUpdateFromSfc updates staker information using SFC binding.
func (ftm *FtmBridge) stakerUpdateFromSfc(staker *types.Staker) error {
	// log action
	ftm.log.Debug("updating staker info from SFC")

	// instantiate the contract and display its name
	contract, err := NewSfcContract(sfcContractAddress, ftm.eth)
	if err != nil {
		ftm.log.Criticalf("failed to instantiate SFC contract: %v", err)
		return err
	}

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

// maxDelegatedLimit calculate maximum amount of tokens allowed to be delegated to a staker.
func (ftm *FtmBridge) maxDelegatedLimit(staked *hexutil.Big, contract *SfcContract) hexutil.Big {
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

	// update detail
	err = ftm.stakerUpdateFromSfc(&st)
	if err != nil {
		ftm.log.Critical("staker information could not be updated from SFC")
	}

	// keep track of the operation
	ftm.log.Debugf("staker #%d loaded", id)
	return &st, nil
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

	// update detail
	err = ftm.stakerUpdateFromSfc(&st)
	if err != nil {
		ftm.log.Critical("staker information could not be updated from SFC")
	}

	// keep track of the operation
	ftm.log.Debugf("staker %s loaded", addr.String())
	return &st, nil
}

// Epoch extract information about an epoch from SFC smart contract.
func (ftm *FtmBridge) Epoch(id hexutil.Uint64) (types.Epoch, error) {
	// instantiate the contract and display its name
	contract, err := NewSfcContract(sfcContractAddress, ftm.eth)
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
func (ftm *FtmBridge) DelegationRewards(addr string) (types.PendingRewards, error) {
	// log action
	ftm.log.Debugf("loading delegation rewards for account %s", addr)

	// instantiate the contract and display its name
	contract, err := NewSfcContract(sfcContractAddress, ftm.eth)
	if err != nil {
		ftm.log.Criticalf("failed to instantiate SFC contract: %v", err)
		return types.PendingRewards{}, err
	}

	// get the value from the contract
	epoch, err := contract.CurrentSealedEpoch(nil)
	if err != nil {
		ftm.log.Errorf("failed to get the current sealed epoch: %v", err)
		return types.PendingRewards{}, err
	}

	// get the rewards amount
	amount, fromEpoch, toEpoch, err := contract.CalcDelegationRewards(nil, common.HexToAddress(addr), big.NewInt(0), epoch)
	if err != nil {
		ftm.log.Errorf("failed to get the delegation rewards: %v", err)
		return types.PendingRewards{}, nil
	}

	// return the data
	return types.PendingRewards{
		Amount:    hexutil.Big(*amount),
		FromEpoch: hexutil.Uint64(fromEpoch.Uint64()),
		ToEpoch:   hexutil.Uint64(toEpoch.Uint64()),
	}, nil
}

// DelegationsOf extract a list of delegations for a given staker.
func (ftm *FtmBridge) DelegationsOf(staker hexutil.Uint64) ([]types.Delegator, error) {
	// keep track of the operation
	ftm.log.Debugf("loading delegations of staker %d", staker)

	// call for data
	dl := make([]types.Delegator, 0)
	err := ftm.rpc.Call(&dl, "sfc_getDelegatorsOf", staker, "0x2")
	if err != nil {
		ftm.log.Error("delegations list could not be extracted")
		return nil, err
	}

	// keep track of the operation
	ftm.log.Debugf("delegations of staker %d loaded", staker)
	return dl, nil
}

// Delegation returns a detail of delegation for the given address.
func (ftm *FtmBridge) Delegation(addr common.Address) (*types.Delegator, error) {
	// keep track of the operation
	ftm.log.Debugf("loading delegation of address %s", addr.String())

	// call for data
	var dl types.Delegator
	err := ftm.rpc.Call(&dl, "sfc_getDelegator", addr, "0x2")
	if err != nil {
		ftm.log.Error("delegation not found")
		return nil, err
	}

	// calculate real amounts delegated and amount in un-delegation
	if err := ftm.delegatedAmount(&dl); err != nil {
		ftm.log.Debugf("can not load delegated amount for %s; %s", addr.String(), err.Error())
		return nil, err
	}

	// keep track of the operation
	ftm.log.Debugf("delegation of address %s loaded", addr.String())
	return &dl, nil
}

// delegatedAmount calculates total amount currently delegated
// and amount locked in pending un-delegation.
// Partial Un-delegations are subtracted during the preparation
// phase, but total un-delegations are subtracted only when
// the delegation is closed.
func (ftm *FtmBridge) delegatedAmount(dl *types.Delegator) error {
	// base delegated amount is copied here
	dl.Amount = (hexutil.Big)(*big.NewInt(0))
	dl.AmountInWithdraw = (hexutil.Big)(*big.NewInt(0))

	// do we have any delegation active at all?
	if nil == dl.AmountDelegated || 0 == dl.AmountDelegated.ToInt().Uint64() {
		return nil
	}

	// staker we are dealing with
	stakerId := big.NewInt(int64(dl.ToStakerId))

	// get the amount of pending withdrawals
	pw, err := ftm.PendingWithdrawalsAmount(&dl.Address, stakerId)
	if err != nil {
		return err
	}

	// add the pending withdrawal to the base amount value
	// since we want it to include these pending partial un-delegations
	// as well
	dl.Amount = (hexutil.Big)(*dl.AmountDelegated.ToInt())
	newAmount := new(big.Int).Add(dl.AmountDelegated.ToInt(), pw)
	dl.Amount = (hexutil.Big)(*newAmount)

	// try to find pending deactivation for the current staker
	// since if there is one, the whole amount is going
	// to be reclaimed and hence is subject to pending withdrawal
	pd, err := ftm.PendingDeactivation(&dl.Address, stakerId)

	// calculate amount in withdraw
	var inWithdraw *big.Int
	if pd != nil {
		inWithdraw = new(big.Int).Add(pw, dl.AmountDelegated.ToInt())
	} else {
		inWithdraw = pw
	}

	// apply to the output
	dl.AmountInWithdraw = (hexutil.Big)(*inWithdraw)

	// we may need to add partial un-delegations
	return nil
}
