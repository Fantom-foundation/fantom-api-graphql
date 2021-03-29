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

// Validator extract a staker information by numeric id.
func (ftm *FtmBridge) Validator(valID *big.Int) (*types.Validator, error) {
	// keep track of the operation
	ftm.log.Debugf("loading validator #%d", valID.Uint64())

	// instantiate the contract and display its name
	contract, err := contracts.NewSfcContract(ftm.sfcConfig.SFCContract, ftm.eth)
	if err != nil {
		ftm.log.Criticalf("failed to instantiate SFC contract; %s", err.Error())
		return nil, err
	}

	return ftm.validatorById(contract, valID)
}

// validatorById loads details of a validator with the specified ID.
func (ftm *FtmBridge) validatorById(contract *contracts.SfcContract, valID *big.Int) (*types.Validator, error) {
	// call for data
	val, err := contract.GetValidator(nil, valID)
	if err != nil {
		ftm.log.Criticalf("failed to load validator #%d from SFC; %s", valID.Uint64(), err.Error())
		return nil, err
	}

	// any creation record?
	if 0 == val.CreatedEpoch.Uint64() {
		ftm.log.Errorf("invalid validator request for #%d", valID.Uint64())
		return nil, fmt.Errorf("unknown validator #%d", valID.Uint64())
	}

	// any deactivation epoch?
	var deaEpoch *hexutil.Uint64
	if nil != val.DeactivatedEpoch {
		dea := hexutil.Uint64(val.DeactivatedEpoch.Uint64())
		deaEpoch = &dea
	}

	// any deactivation time?
	var deaTime *hexutil.Uint64
	if nil != val.DeactivatedTime {
		dea := hexutil.Uint64(val.DeactivatedTime.Uint64())
		deaTime = &dea
	}

	// keep track of the operation
	ftm.log.Debugf("validator #%d is %s", valID.Uint64(), val.Auth.String())
	return &types.Validator{
		Id:               hexutil.Uint64(valID.Uint64()),
		StakerAddress:    val.Auth,
		TotalStake:       (*hexutil.Big)(val.ReceivedStake),
		Status:           hexutil.Uint64(val.Status.Uint64()),
		CreatedEpoch:     hexutil.Uint64(val.CreatedEpoch.Uint64()),
		CreatedTime:      hexutil.Uint64(val.CreatedTime.Uint64()),
		DeactivatedEpoch: deaEpoch,
		DeactivatedTime:  deaTime,
	}, nil
}

// ValidatorAddress extract a staker address for the given staker ID.
func (ftm *FtmBridge) ValidatorAddress(valID *big.Int) (common.Address, error) {
	// instantiate the contract and display its name
	contract, err := contracts.NewSfcContract(ftm.sfcConfig.SFCContract, ftm.eth)
	if err != nil {
		ftm.log.Criticalf("failed to instantiate SFC contract;%s", err.Error())
		return common.Address{}, err
	}

	// do we have an address call?
	val, err := contract.GetValidator(nil, valID)
	if err != nil {
		ftm.log.Error("validator information could not be extracted")
		return common.Address{}, err
	}

	// do we have the val of this number?
	if 0 == val.CreatedEpoch.Uint64() {
		ftm.log.Errorf("invalid validator request for #%d", valID.Uint64())
		return common.Address{}, fmt.Errorf("unknown validator #%d", valID.Uint64())
	}

	ftm.log.Debugf("validator #%d is %s", valID.Uint64(), val.Auth.String())
	return val.Auth, nil
}

// IsValidator returns if the given address is an SFC validator.
func (ftm *FtmBridge) IsValidator(addr *common.Address) (bool, error) {
	// keep track of the operation
	ftm.log.Debugf("verifying validator address %s", addr.String())

	// instantiate the contract and display its name
	contract, err := contracts.NewSfcContract(ftm.sfcConfig.SFCContract, ftm.eth)
	if err != nil {
		ftm.log.Criticalf("failed to instantiate SFC contract; %s", err.Error())
		return false, err
	}

	// try to get the id
	id, err := contract.GetStakerID(nil, *addr)
	if err != nil {
		ftm.log.Criticalf("can not check validator at %s; %s", addr.String(), err.Error())
		return false, err
	}

	return 0 < id.Uint64(), nil
}

// ValidatorByAddress extracts a validator information by address.
func (ftm *FtmBridge) ValidatorByAddress(addr *common.Address) (*types.Validator, error) {
	// keep track of the operation
	ftm.log.Debugf("loading validator with address %s", addr.String())

	// instantiate the contract and display its name
	contract, err := contracts.NewSfcContract(ftm.sfcConfig.SFCContract, ftm.eth)
	if err != nil {
		ftm.log.Criticalf("failed to instantiate SFC contract; %s", err.Error())
		return nil, err
	}

	// try to get the staker id
	id, err := contract.GetStakerID(nil, *addr)
	if err != nil {
		ftm.log.Criticalf("can not check validator at %s; %s", addr.String(), err.Error())
		return nil, err
	}

	// do we have the ID?
	if 0 == id.Uint64() {
		ftm.log.Debugf("validator not found for address %s", addr.String())
		return nil, fmt.Errorf("unknown validator %s", addr.String())
	}
	return ftm.validatorById(contract, id)
}
