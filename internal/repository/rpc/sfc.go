/*
Rpc package implements bridge to Lachesis full node API interface.

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
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// sfcContractAddress represents the address on which the Sfc contract is deployed.
var sfcContractAddress = common.HexToAddress("0xfc00face00000000000000000000000000000000")

// CurrentEpoch extract the current epoch from SFC smart contract.
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

		// recalculate the total stake
		staker.TotalStake = (*hexutil.Big)(big.NewInt(0).Add(si.DelegatedMe, si.StakeAmount))
	} else {
		// log issue
		ftm.log.Debug("staker info update from SFC failed, no data received")
	}

	// get the value
	return nil
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

// Staker extract a staker information by address.
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

// CurrentEpoch extract information about an epoch from SFC smart contract.
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

// Delegation returns a detail of delegation for the given address.
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

	// keep track of the operation
	ftm.log.Debugf("delegation of address %s loaded", addr.String())
	return &dl, nil
}
