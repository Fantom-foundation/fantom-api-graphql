/*
Package repository implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Opera/Lachesis full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import (
	"fantom-api-graphql/internal/logger"
	"fantom-api-graphql/internal/repository/rpc/contracts"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
	"sync"
)

const (
	// contractCallQueueLength represents how many smart contract calls can be pushed
	// into the queue for processing at once
	contractCallQueueLength = 75000
)

// contractCallDispatcher implements blockchain smart contract calls analyzer queue.
type contractCallDispatcher struct {
	service
	buffer chan *types.Transaction
}

// newContractCallDispatcher creates new blockchain smart contract call analyzer queue service.
func newContractCallDispatcher(
	buffer chan *types.Transaction,
	repo Repository,
	log logger.Logger,
	wg *sync.WaitGroup,
) *contractCallDispatcher {
	// create new instance
	return &contractCallDispatcher{
		service: newService("contract call dispatcher", repo, log, wg),
		buffer:  buffer,
	}
}

// run starts the account queue to life.
func (ccd *contractCallDispatcher) run() {
	// start blockScanner
	ccd.wg.Add(1)
	go ccd.dispatch()
}

// monitorContractCalls runs the main contract calls queue resolver
// loop in a separate thread.
func (ccd *contractCallDispatcher) dispatch() {
	// log action
	ccd.log.Notice("contract call dispatcher is running")

	// don't forget to sign off after we are done
	defer func() {
		// log finish
		ccd.log.Notice("contract call dispatcher is closed")

		// signal to wait group we are done
		ccd.wg.Done()
	}()

	// wait for either stop signal, or an account request
	for {
		select {
		case trx := <-ccd.buffer:
			// log what we do
			ccd.log.Debugf("analyzing contract call at %s", trx.Hash.String())

			// check the call
			ccd.analyzeCall(trx)
		case <-ccd.sigStop:
			// stop signal received?
			return
		}
	}
}

// IsLikelyAContractCall checks if the given transaction can actually be a contract call.
func IsLikelyAContractCall(trx *types.Transaction) bool {
	// the transaction must make sense, has to be addressed to a recipient address,
	// and has to have some input data of at least 4 bytes (function selector)
	// The function selector is created by taking the first 4 bytes from the Keccak hash
	// over the function name and its argument types.
	return trx != nil && trx.To != nil && trx.InputData != nil && 4 <= len(trx.InputData)
}

// analyzeCall analyzes smart contract call and tries to add details
// to the transaction base off-chain data record.
func (ccd *contractCallDispatcher) analyzeCall(trx *types.Transaction) {
	// check if the transaction is likely to be a contract call
	if !IsLikelyAContractCall(trx) {
		ccd.log.Error("analyzer received a non-call transaction")
		return
	}

	// do we know the contract the transaction points to?
	sc, err := ccd.repo.Contract(trx.To)
	if err != nil {
		ccd.log.Errorf("can not analyze call at %s; %s", trx.Hash.String(), err.Error())
	}

	// unknown contract? probably not a call
	if sc == nil {
		ccd.log.Debugf("transaction %s recipient not known, probably not a contract call", trx.Hash.String())
		return
	}

	// assign transaction target contract type
	ccd.updateTargetContractType(trx, sc)

	// decode function of the call
	ccd.updateTargetFunctionSignature(trx, sc)

	// update the transaction in repository
	err = ccd.repo.TransactionUpdate(trx)
	if err != nil {
		ccd.log.Errorf("transaction %s not updated; %s", trx.Hash.String(), err.Error())
	}
}

// updateTargetContractType updates the transaction target contract types
// based on the associated contract.
func (ccd *contractCallDispatcher) updateTargetContractType(trx *types.Transaction, sc *types.Contract) {
	// do the update
	trx.TargetContractType = &sc.Type
	trx.IsErc20Call = strings.EqualFold(sc.Type, types.AccountTypeERC20Token)
}

// updateTargetFunctionSignature detects and decodes the target function signature
// if possible.
func (ccd *contractCallDispatcher) updateTargetFunctionSignature(trx *types.Transaction, sc *types.Contract) {
	// do we have the contract abi? try to use it to find match
	if sc.Abi != "" {
		// mark the ABI parser use
		ccd.log.Debugf("ABI found for contract %s, decoding %s", sc.Address.String(), trx.Hash.String())

		// match the ABI
		ccd.tryMatchWithAbi(trx, &sc.Abi)
	}

	// if we don't have a match and it's the SFC, try previous version of the contract
	if trx.TargetFunctionCall == nil && ccd.repo.IsSfcContract(trx.To) {
		v1Abi := contracts.SfcV1ContractABI
		ccd.tryMatchWithAbi(trx, &v1Abi)
	}

	// do we have the call signature?
	if trx.TargetFunctionCall == nil {
		// log the issue
		ccd.log.Debugf("transaction %s call undefined, generic signature used", trx.Hash.String())

		// use the raw call function signature
		fc := trx.InputData.String()[:8]
		trx.TargetFunctionCall = &fc
	}
}

// tryToMatchAbi tries the given ABI to match and update the contract call.
func (ccd *contractCallDispatcher) tryMatchWithAbi(trx *types.Transaction, inAbi *string) {
	// try to parse the ABI from JSON so we can match function for the call
	parsed, err := abi.JSON(strings.NewReader(*inAbi))
	if err != nil {
		ccd.log.Debugf("failed to parse ABI; %s", err.Error())
		return
	}

	// try to match the trx call with a contract method
	ccd.matchCallMethod(trx, &parsed)
}

// matchCallMethod tries to match call method from the call data and parsed ABI
// of the addressed contract.
func (ccd *contractCallDispatcher) matchCallMethod(trx *types.Transaction, inAbi *abi.ABI) {
	// try to find the method; returns error if not found
	m, err := inAbi.MethodById(trx.InputData)
	if err != nil {
		ccd.log.Errorf("method signature not found at %s; %s", trx.Hash.String(), err.Error())
		return
	}

	// keep the function name for the reference
	trx.TargetFunctionCall = &m.Name
	ccd.log.Debugf("found %s() call at %s", m.Name, trx.Hash.String())
}
