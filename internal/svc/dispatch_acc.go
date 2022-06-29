// Package svc implements blockchain data processing services.
package svc

import (
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

var erc721InterfaceId  = [4]byte { 0x80, 0xac, 0x58, 0xcd } // ERC-721: 0x80ac58cd
var erc1155InterfaceId = [4]byte { 0xd9, 0xb6, 0x7a, 0x26 } // ERC-1155: 0xd9b67a26

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
		err := acd.processContract(acc)
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
func (acd *accDispatcher) processContract(acc *eventAcc) error {
	// log what we do
	log.Debugf("account %s is a smart contract, analyzing", acc.addr.String())

	// detect and identify contract
	contract, accountType, err := acd.detectContract(acc.addr, acc.blk, acc.trx)
	if err != nil {
		log.Errorf("can not identify contract at %s; %s", acc.addr.String(), err.Error())
		return err
	}
	acc.act = accountType

	// insert the contract record if possible
	if contract != nil {
		err = repo.StoreContract(contract)
		if err != nil {
			log.Errorf("can not add contract at %s; %s", acc.addr.String(), err.Error())
			return err
		}
	}
	return nil
}

// detectContract tries to identify the contract type.
func (acd *accDispatcher) detectContract(addr *common.Address, block *types.Block, trx *types.Transaction) (*types.Contract, string, error) {

	isErc1155, err := repo.Erc165SupportsInterface(addr, erc1155InterfaceId)
	if err == nil && isErc1155 {
		log.Noticef("ERC1155 multi-token detected at %s", addr.String())
		contract := types.NewErcTokenContract(addr, "", block, trx, types.AccountTypeERC1155Contract, contracts.ERC1155MetaData.ABI)
		return contract, types.AccountTypeERC1155Contract, nil
	}

	isErc721, name := acd.detectErc721Token(addr)
	if err == nil && isErc721 {
		log.Noticef("ERC721 NFT token detected at %s", addr.String())
		contract := types.NewErcTokenContract(addr, name, block, trx, types.AccountTypeERC721Contract, contracts.ERC721MetaData.ABI)
		return contract, types.AccountTypeERC721Contract, nil
	}

	isErc20, name := acd.detectErc20Token(addr)
	if isErc20 {
		log.Noticef("ERC20 token %s detected at %s", name, addr.String())
		contract := types.NewErcTokenContract(addr, name, block, trx, types.AccountTypeERC20Token, contracts.ERCTwentyMetaData.ABI)
		return contract, types.AccountTypeERC20Token, nil
	}

	// log that the detection failed
	log.Noticef("unknown contract at %s", addr.String())

	// set as generic contract type if no other has been detected
	return types.NewGenericContract(addr, block, trx), types.AccountTypeContract, nil
}

// detectErc20Token identifies ERC20 token contracts by trying to call specific contract methods.
func (acd *accDispatcher) detectErc20Token(addr *common.Address) (isErc20 bool, name string) {
	// try to get the token name
	name, err := repo.Erc20Name(addr)
	if err != nil {
		return false, ""
	}

	// try to detect symbol
	if _, err := repo.Erc20Symbol(addr); err != nil {
		return false, ""
	}

	// try to detect balance of
	if _, err := repo.Erc20BalanceOf(addr, &testAddress); err != nil {
		return false, ""
	}

	// try to detect total supply
	if _, err := repo.Erc20TotalSupply(addr); err != nil {
		return false, ""
	}

	return true, name
}

func (acd *accDispatcher) detectErc721Token(addr *common.Address) (isErc721 bool, name string) {
	isErc721, err := repo.Erc165SupportsInterface(addr, erc721InterfaceId)
	if err != nil || !isErc721 {
		return false, ""
	}

	// try to detect name - but it is optional for ERC721
	name, err = repo.Erc20Name(addr)
	if err != nil {
		return true, ""
	}

	return true, name
}
