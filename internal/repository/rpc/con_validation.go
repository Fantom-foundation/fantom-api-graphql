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
	"fantom-api-graphql/internal/config"
	"fantom-api-graphql/internal/types"
	"fmt"
)

//go:generate abigen --abi ./contracts/con_validation.abi --pkg rpc --type ContractValidator --out ./con_validation_bind.go

// checkValidationConfig verifies that the validation config provided
// by the configuration is valid and can be used to validate
func checkValidationConfig(cfg *config.Config) error {
	// validation repository contract address must be valid
	if 0 == len(cfg.ContractValidatorRepository) {
		return fmt.Errorf("missing validation contract address")
	}

	// validation key file must be provided
	if 0 == len(cfg.ContractValidatorSigKey) {
		return fmt.Errorf("missing validation identity signing key")
	}

	return nil
}

// SubmitValidatedContract submits a validated contract information to the blockchain deployed
// special contract, which keeps track of other contract validations and allows peer API points
// to recognize and process contract validations, even retrospectively.
func (ftm *FtmBridge) SubmitValidatedContract(con *types.Contract) error {
	// instantiate validation binding
	validator, err := NewContractValidator(ftm.valRepository, ftm.eth)
	if err != nil {
		ftm.log.Errorf("can not open contract validator connection; %s", err.Error())
		return err
	}

	// check if this contract validation is actually new
	known, err := ftm.isContractKnown(con, validator)
	if err != nil {
		ftm.log.Errorf("can not open contract validator connection; %s", err.Error())
		return err
	}

	// is the contract known already? exit if it is
	if known {
		ftm.log.Debugf("contract %s is already listed", con.Address.String())
		return nil
	}

	return nil
}

// isContractKnown checks if the provided contract is already listed
// in the on-chain contract validation repository.
func (ftm *FtmBridge) isContractKnown(con *types.Contract, validator *ContractValidator) (bool, error) {
	// try to find the validated contract on chain
	vc, err := validator.Contracts(nil, con.Address)
	if err != nil {
		ftm.log.Errorf("contract validation check failed; %s", err.Error())
		return false, err
	}

	return vc.SrcHash == con.TransactionHash, nil
}
