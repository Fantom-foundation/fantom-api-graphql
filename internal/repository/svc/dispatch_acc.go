/*
Package repository implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Opera/Lachesis full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/repository/rpc"
	"fantom-api-graphql/internal/repository/rpc/contracts"
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
)

const (
	// sfcCheckBelowBlock represents the highest block number we try to detect
	// SFC contract, above this block the contract should already be known, and we can
	// skip the check
	sfcCheckBelowBlock = 100000
)

// testAddress represents an address used to test an account reference
var testAddress = common.HexToAddress("0xabc00FA001230012300aBc0012300Fa00FACE000")

// accDispatcher implements account dispatcher queue
type accDispatcher struct {
	or        *Orchestrator
	sigStop   chan bool
	inAccount chan *eventAcc
}

// name returns the name of the service.
func (acd *accDispatcher) name() string {
	return "account dispatcher"
}

// init prepares the account dispatcher to perform its function.
func (acd *accDispatcher) init() {
	acd.sigStop = make(chan bool, 1)
}

// run starts the account queue to life.
func (acd *accDispatcher) run() {
	// make sure we are orchestrated
	if acd.or == nil {
		panic(fmt.Errorf("no orchestrator set"))
	}

	// signal orchestrator we started and go
	acd.or.started(acd)
	go acd.dispatch()
}

// close terminates the block dispatcher.
func (acd *accDispatcher) close() {
	acd.sigStop <- true
}

// dispatch runs the main account requests monitor and dispatcher
// loop in a separate thread.
func (acd *accDispatcher) dispatch() {
	// don't forget to sign off after we are done
	defer func() {
		acd.or.finished(acd)
	}()

	// wait for either stop signal, or an account request
	for {
		select {
		case <-acd.sigStop:
			return
		case acc, ok := <-acd.inAccount:
			// is the channel even available for reading
			if !ok {
				acd.or.log.Notice("account channel closed, terminating %s", acd.name())
				return
			}

			// do the stuff
			err := acd.process(acc)
			if err != nil {
				acd.or.log.Errorf("failed account %s processing; %s", acc.addr.String(), err.Error())
			}
		}
	}
}

// processAccount processes account into the database
// based on the account details
func (acd *accDispatcher) process(acc *eventAcc) error {
	// log what we do
	acd.or.log.Debugf("account %s received for processing", acc.addr.String())

	// check if the account is new; if we already know it, we are done
	if repository.R().AccountIsKnown(acc.addr) {
		return repository.R().AccountMarkActivity(acc.addr, uint64(acc.blk.TimeStamp))
	}

	// is this a simple wallet/account?
	if acc.trx.ContractAddress != nil {
		err := acd.contract(acc)
		if err != nil {
			return err
		}
	}
	return acd.wallet(acc)
}

// wallet processes a simple non-contract wallet account into the database
// based on the account details (it still could be the SFC, be cautious about it)
func (acd *accDispatcher) wallet(acc *eventAcc) error {
	// notify new account detected
	acd.or.log.Debugf("found new account %s", acc.addr.String())

	// check if the target address is not an SFC contract
	acd.checkSfc(acc)

	// add the account into the database
	err := repository.R().StoreAccount(&types.Account{
		Address:      *acc.addr,
		ContractTx:   acc.deploy,
		Type:         acc.act,
		LastActivity: acc.blk.TimeStamp,
		TrxCounter:   1,
	})
	if err != nil {
		acd.or.log.Errorf("can not add account %s; %s", acc.addr.String(), err.Error())
	}
	return err
}

// checkSfc verifies if the target account is the SFC contract
// and if so, it adds the SFC target with a different type.
func (acd *accDispatcher) checkSfc(acc *eventAcc) {
	// act on SFC detection
	if uint64(acc.blk.Number) < sfcCheckBelowBlock && repository.R().IsSfcContract(acc.addr) {
		// change the type to SFC contract
		acc.act = types.AccountTypeSFC

		// get the SFC version
		ver, err := repository.R().SfcVersion()
		if err == nil {
			// log what we found
			acd.or.log.Debugf("detected SFC contract %d.%d.%d", byte((ver>>16)&255), byte((ver>>8)&255), byte(ver&255))

			// add the contract
			err = repository.R().StoreContract(types.NewSfcContract(acc.addr, uint64(ver), acc.blk, acc.trx))
			if err != nil {
				acd.or.log.Errorf("can not add the SFC contract at %s; %s", acc.addr.String(), err.Error())
			}
		}
	}
}

// processContract processes contract account with detection.
func (acd *accDispatcher) contract(acc *eventAcc) error {
	// log what we do
	acd.or.log.Debugf("account %s is a smart contract, analyzing", acc.addr.String())

	// detect and identify contract
	con, err := acd.detect(acc.addr, &acc.act, acc.blk, acc.trx)
	if err != nil {
		acd.or.log.Errorf("can not identify contract at %s; %s", acc.addr.String(), err.Error())
		return err
	}

	// insert the contract record if possible
	if con != nil {
		err = repository.R().StoreContract(con)
		if err != nil {
			acd.or.log.Errorf("can not add contract at %s; %s", acc.addr.String(), err.Error())
			return err
		}
	}
	return nil
}

// detectContract tries to identify the contract type.
func (acd *accDispatcher) detect(addr *common.Address, ct *string, block *types.Block, trx *types.Transaction) (*types.Contract, error) {
	// identify ERC20 token
	con := acd.detectErc20(addr, block, trx)
	if con != nil {
		*ct = types.AccountTypeERC20Token
		return con, nil
	}

	// log that the detection failed
	acd.or.log.Noticef("unknown contract at %s", addr.String())

	// set as generic contract type if no other has ben detected
	*ct = types.AccountTypeContract
	return types.NewGenericContract(addr, block, trx), nil
}

// detectErc20 identifies ERC20 token contracts by checking common contract end points.
func (acd *accDispatcher) detectErc20(addr *common.Address, block *types.Block, trx *types.Transaction) *types.Contract {
	// use repository to check the contract
	r := repository.R()

	// try to get the token name
	name, err := r.Erc20Name(addr)
	if err != nil {
		return nil
	}

	// try to detect symbol
	if _, err := r.Erc20Symbol(addr); err != nil {
		return nil
	}

	// try to detect balance of
	if _, err := r.Erc20BalanceOf(addr, &testAddress); err != nil {
		return nil
	}

	// try to detect total supply; if not found this is probably ERC721
	if _, err := r.Erc20TotalSupply(addr); err != nil {
		acd.or.log.Noticef("ERC721 token %s detected at %s", name, addr.String())
		return types.NewErcTokenContract(addr, name, block, trx, types.AccountTypeERC20Token, rpc.ERC721TokenABI)
	}

	// log what we do
	acd.or.log.Noticef("ERC20 token %s detected at %s", name, addr.String())
	return types.NewErcTokenContract(addr, name, block, trx, types.AccountTypeERC721Token, contracts.ERCTwentyABI)
}
