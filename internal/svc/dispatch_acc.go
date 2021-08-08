// Package svc implements blockchain data processing services.
package svc

import (
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
	inAccount chan *eventAcc
	service
}

// name returns the name of the service used by orchestrator.
func (acd *accDispatcher) name() string {
	return "account dispatcher"
}

// run starts the account queue to life.
func (acd *accDispatcher) run() {
	// make sure we are orchestrated
	if acd.mgr == nil {
		panic(fmt.Errorf("no svc manager set on %s", acd.name()))
	}

	// signal orchestrator we started and go
	acd.mgr.started(acd)
	go acd.execute()
}

// execute runs the main account requests monitor and dispatcher
// loop in a separate thread.
func (acd *accDispatcher) execute() {
	// don't forget to sign off after we are done
	defer func() {
		close(acd.sigStop)
		acd.mgr.finished(acd)
	}()

	// wait for either stop signal, or an account request
	for {
		select {
		case <-acd.sigStop:
			return
		case acc, ok := <-acd.inAccount:
			// is the channel even available for reading
			if !ok {
				log.Noticef("account channel closed, terminating %s", acd.name())
				return
			}

			// do the stuff
			err := acd.process(acc)
			if err != nil {
				log.Errorf("failed account %s processing; %s", acc.addr.String(), err.Error())
			}

			// signal this account has been processed
			acc.watchDog.Done()
		}
	}
}

// processAccount processes account into the database
// based on the account details
func (acd *accDispatcher) process(acc *eventAcc) error {
	// log what we do
	log.Debugf("account %s received for processing", acc.addr.String())

	// check if the account is new; if we already know it, we are done
	if repo.AccountIsKnown(acc.addr) {
		return repo.AccountMarkActivity(acc.addr, uint64(acc.blk.TimeStamp))
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
	log.Debugf("found new account %s", acc.addr.String())

	// check if the target address is not an SFC contract
	acd.checkSfc(acc)

	// add the account into the database
	err := repo.StoreAccount(&types.Account{
		Address:      *acc.addr,
		ContractTx:   acc.deploy,
		Type:         acc.act,
		LastActivity: acc.blk.TimeStamp,
		TrxCounter:   1,
	})
	if err != nil {
		log.Errorf("can not add account %s; %s", acc.addr.String(), err.Error())
	}
	return err
}

// checkSfc verifies if the target account is the SFC contract
// and if so, it adds the SFC target with a different type.
func (acd *accDispatcher) checkSfc(acc *eventAcc) {
	// act on SFC detection
	if uint64(acc.blk.Number) < sfcCheckBelowBlock && repo.IsSfcContract(acc.addr) {
		// change the type to SFC contract
		acc.act = types.AccountTypeSFC

		// get the SFC version
		ver, err := repo.SfcVersion()
		if err == nil {
			// log what we found
			log.Debugf("detected SFC contract %d.%d.%d", byte((ver>>16)&255), byte((ver>>8)&255), byte(ver&255))

			// add the contract
			err = repo.StoreContract(types.NewSfcContract(acc.addr, uint64(ver), acc.blk, acc.trx))
			if err != nil {
				log.Errorf("can not add the SFC contract at %s; %s", acc.addr.String(), err.Error())
			}
		}
	}
}

// processContract processes contract account with detection.
func (acd *accDispatcher) contract(acc *eventAcc) error {
	// log what we do
	log.Debugf("account %s is a smart contract, analyzing", acc.addr.String())

	// detect and identify contract
	con, err := acd.detect(acc.addr, &acc.act, acc.blk, acc.trx)
	if err != nil {
		log.Errorf("can not identify contract at %s; %s", acc.addr.String(), err.Error())
		return err
	}

	// insert the contract record if possible
	if con != nil {
		err = repo.StoreContract(con)
		if err != nil {
			log.Errorf("can not add contract at %s; %s", acc.addr.String(), err.Error())
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
	log.Noticef("unknown contract at %s", addr.String())

	// set as generic contract type if no other has ben detected
	*ct = types.AccountTypeContract
	return types.NewGenericContract(addr, block, trx), nil
}

// detectErc20 identifies ERC20 token contracts by checking common contract end points.
func (acd *accDispatcher) detectErc20(addr *common.Address, block *types.Block, trx *types.Transaction) *types.Contract {
	// try to get the token name
	name, err := repo.Erc20Name(addr)
	if err != nil {
		return nil
	}

	// try to detect symbol
	if _, err := repo.Erc20Symbol(addr); err != nil {
		return nil
	}

	// try to detect balance of
	if _, err := repo.Erc20BalanceOf(addr, &testAddress); err != nil {
		return nil
	}

	// try to detect total supply; if not found this is probably ERC721
	if _, err := repo.Erc20TotalSupply(addr); err != nil {
		log.Noticef("ERC721 token %s detected at %s", name, addr.String())
		return types.NewErcTokenContract(addr, name, block, trx, types.AccountTypeERC20Token, rpc.ERC721TokenABI)
	}

	// log what we do
	log.Noticef("ERC20 token %s detected at %s", name, addr.String())
	return types.NewErcTokenContract(addr, name, block, trx, types.AccountTypeERC721Token, contracts.ERCTwentyABI)
}
