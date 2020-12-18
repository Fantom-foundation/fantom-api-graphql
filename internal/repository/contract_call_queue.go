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

// contractCallQueue implements blockchain smart contract calls analyzer queue.
type contractCallQueue struct {
	service
	buffer chan *types.Transaction
}

// newContractCallQueue creates new blockchain smart contract call analyzer queue service.
func newContractCallQueue(
	buffer chan *types.Transaction,
	repo Repository,
	log logger.Logger,
	wg *sync.WaitGroup,
) *contractCallQueue {
	// create new instance
	return &contractCallQueue{
		service: newService("contract calls queue", repo, log, wg),
		buffer:  buffer,
	}
}

// run starts the account queue to life.
func (cq *contractCallQueue) run() {
	// log the action
	cq.log.Notice("starting contract queue dispatcher")

	// start scanner
	cq.wg.Add(1)
	go cq.monitorContractCalls()
}

// monitorContractCalls runs the main contract calls queue resolver
// loop in a separate thread.
func (cq *contractCallQueue) monitorContractCalls() {
	// log action
	cq.log.Notice("contract calls queue processing is running")

	// don't forget to sign off after we are done
	defer func() {
		// log finish
		cq.log.Notice("contract calls queue processing is closing")

		// signal to wait group we are done
		cq.wg.Done()
	}()

	// wait for either stop signal, or an account request
	for {
		select {
		case trx := <-cq.buffer:
			// log what we do
			cq.log.Debugf("analyzing contract call at %s", trx.Hash.String())

			// check the call
			cq.analyzeCall(trx)
		case <-cq.sigStop:
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
func (cq *contractCallQueue) analyzeCall(trx *types.Transaction) {
	// check if the transaction is likely to be a contract call
	if !IsLikelyAContractCall(trx) {
		cq.log.Error("analyzer received a non-call transaction")
		return
	}

	// do we know the contract the transaction points to?
	sc, err := cq.repo.Contract(trx.To)
	if err != nil {
		cq.log.Errorf("can not analyze call at %s; %s", trx.Hash.String(), err.Error())
	}

	// unknown contract? probably not a call
	if sc == nil {
		cq.log.Debugf("transaction %s recipient not known, probably not a contract call", trx.Hash.String())
		return
	}

	// assign transaction target contract type
	cq.updateTargetContractType(trx, sc)

	// decode function of the call
	cq.updateTargetFunctionSignature(trx, sc)

	// update the transaction in repository
	err = cq.repo.TransactionUpdate(trx)
	if err != nil {
		cq.log.Errorf("transaction %s not updated; %s", trx.Hash.String(), err.Error())
	}
}

// updateTargetContractType updates the transaction target contract types
// based on the associated contract.
func (cq *contractCallQueue) updateTargetContractType(trx *types.Transaction, sc *types.Contract) {
	// do the update
	trx.TargetContractType = &sc.Type
	trx.IsErc20Call = strings.EqualFold(sc.Type, types.AccountTypeERC20Token)
}

// updateTargetFunctionSignature detects and decodes the target function signature
// if possible.
func (cq *contractCallQueue) updateTargetFunctionSignature(trx *types.Transaction, sc *types.Contract) {
	// do we have the contract abi? try to use it to find match
	if sc.Abi != "" {
		// mark the ABI parser use
		cq.log.Debugf("ABI found for contract %s, decoding %s", sc.Address.String(), trx.Hash.String())

		// match the ABI
		cq.tryMatchWithAbi(trx, &sc.Abi)
	}

	// if we don't have a match and it's the SFC, try previous version of the contract
	if trx.TargetFunctionCall == nil && cq.repo.IsSfcContract(trx.To) {
		v1Abi := contracts.SfcV1ContractABI
		cq.tryMatchWithAbi(trx, &v1Abi)
	}

	// do we have the call signature?
	if trx.TargetFunctionCall == nil {
		// log the issue
		cq.log.Debugf("transaction %s call undefined, generic signature used", trx.Hash.String())

		// use the raw call function signature
		fc := trx.InputData.String()[:8]
		trx.TargetFunctionCall = &fc
	}
}

// tryToMatchAbi tries the given ABI to match and update the contract call.
func (cq *contractCallQueue) tryMatchWithAbi(trx *types.Transaction, inAbi *string) {
	// try to parse the ABI from JSON so we can match function for the call
	parsed, err := abi.JSON(strings.NewReader(*inAbi))
	if err != nil {
		cq.log.Debugf("failed to parse ABI; %s", err.Error())
		return
	}

	// try to match the trx call with a contract method
	cq.matchCallMethod(trx, &parsed)
}

// matchCallMethod tries to match call method from the call data and parsed ABI
// of the addressed contract.
func (cq *contractCallQueue) matchCallMethod(trx *types.Transaction, inAbi *abi.ABI) {
	// try to find the method; returns error if not found
	m, err := inAbi.MethodById(trx.InputData)
	if err != nil {
		cq.log.Errorf("method signature not found at %s; %s", trx.Hash.String(), err.Error())
		return
	}

	// keep the function name for the reference
	trx.TargetFunctionCall = &m.Name
	cq.log.Debugf("found %s() call at %s", m.Name, trx.Hash.String())
}
