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
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"sync"
)

const (
	// accountQueueLength represents how many accounts can be pushed
	// into the queue for processing at once
	accountQueueLength = 20000

	// sfcCheckBelowBlock represents the highest block number we try to detect
	// SFC contract, above this block the contract should already be known and we can
	// skip the check
	sfcCheckBelowBlock = 25000
)

// testAddress represents an address used to test an account reference
var testAddress = common.HexToAddress("0xabc00FA001230012300aBc0012300Fa00FACE000")

// AccountQueueRequest represents an account queue processing request.
// This can be a simple account existence verification, or an extended
// contract analyze.
type accountQueueRequest struct {
	acc         *types.Account
	blk         *types.Block
	trx         *types.Transaction
	trxCallback func(*types.Transaction)
}

// accountQueue implements account analyzer queue
type accountQueue struct {
	service
	buffer chan *accountQueueRequest
}

// newAccountQueue creates new blockchain account analyzer queue service.
func newAccountQueue(buffer chan *accountQueueRequest, repo Repository, log logger.Logger, wg *sync.WaitGroup) *accountQueue {
	// create new instance
	aq := accountQueue{
		service: newService("account queue", repo, log, wg),
		buffer:  buffer,
	}

	// start the scanner job
	aq.run()
	return &aq
}

// run starts the account queue to life.
func (aq *accountQueue) run() {
	// start scanner
	aq.wg.Add(1)
	go aq.monitorAccounts()
}

// monitorAccounts runs the main account requests resolver
// loop in a separate thread.
func (aq *accountQueue) monitorAccounts() {
	// log action
	aq.log.Notice("account queue processing is running")

	// don't forget to sign off after we are done
	defer func() {
		// log finish
		aq.log.Notice("account queue processing is closing")

		// signal to wait group we are done
		aq.wg.Done()
	}()

	// wait for either stop signal, or an account request
	for {
		select {
		case req := <-aq.buffer:
			// log what we do
			aq.log.Debugf("account %s received for processing", req.acc.Address.String())

			// process the account request into the database
			err := aq.processAccount(req.acc, req.blk, req.trx)

			// any callback? notify the transaction to be done
			if err == nil && req.trxCallback != nil {
				req.trxCallback(req.trx)
			}
		case <-aq.sigStop:
			// stop signal received?
			return
		}
	}
}

// processAccount processes account into the database
// based on the account details
func (aq *accountQueue) processAccount(acc *types.Account, block *types.Block, trx *types.Transaction) error {
	// check if the account is new; if we already know it, we are done
	if aq.repo.AccountIsKnown(&acc.Address) {
		return aq.repo.AccountMarkActivity(acc, uint64(block.TimeStamp))
	}

	// simple account; just push it into the database
	if acc.ContractTx == nil {
		// notify new account detected
		aq.log.Noticef("found new account %s", acc.Address.String())

		// check if the target address is not an SFC contract
		aq.checkSfcContract(acc, block, trx)

		// add the account into the database
		err := aq.repo.AccountAdd(acc)
		if err != nil {
			aq.log.Errorf("can not add the account %s; %s", acc.Address.String(), err.Error())
			return err
		}

		// account added
		return nil
	}

	return aq.processAccountContract(acc, block, trx)
}

// checkSfcContract verifies if the target account is the SFC contract
// and if so, it adds the SFC target with a different type.
func (aq *accountQueue) checkSfcContract(acc *types.Account, block *types.Block, trx *types.Transaction) {
	// act on SFC detection
	if uint64(block.Number) < sfcCheckBelowBlock && aq.repo.IsSfcContract(&acc.Address) {
		// change the type to SFC contract
		acc.Type = types.AccountTypeSFC

		// get the SFC version
		ver, err := aq.repo.SfcVersion()
		if err == nil {
			// log what we found
			aq.log.Debugf("detected SFC contract %d.%d.%d", byte((ver>>16)&255), byte((ver>>8)&255), byte(ver&255))

			// add the contract
			if err := aq.repo.ContractAdd(types.NewSfcContract(&acc.Address, uint64(ver), block, trx)); err != nil {
				aq.log.Errorf("can not add the SFC contract at %s; %s", acc.Address.String(), err.Error())
			}
		}
	}
}

// processAccountContract processes contract account with detection.
func (aq *accountQueue) processAccountContract(acc *types.Account, block *types.Block, trx *types.Transaction) error {
	// log what we do
	aq.log.Debugf("account %s is a smart contract, analyzing", acc.Address.String())

	// detect and identify contract
	con, err := aq.detectContract(&acc.Address, &acc.Type, block, trx)
	if err != nil {
		aq.log.Errorf("can not identify the contract at %s; %s", acc.Address.String(), err.Error())
		return err
	}

	// insert the contract record if possible
	if con != nil {
		err = aq.repo.ContractAdd(con)
		if err != nil {
			aq.log.Errorf("can not add the contract at %s; %s", acc.Address.String(), err.Error())
			return err
		}
	}

	// add the account identified into teh database
	err = aq.repo.AccountAdd(acc)
	if err != nil {
		aq.log.Errorf("can not add the account %s; %s", acc.Address.String(), err.Error())
		return err
	}

	return err
}

// detectContract tries to identify the contract type.
func (aq *accountQueue) detectContract(addr *common.Address, cType *string, block *types.Block, trx *types.Transaction) (*types.Contract, error) {
	// identify ERC20 token
	con := aq.detectErc20Token(addr, block, trx)
	if con != nil {
		*cType = types.AccountTypeERC20Token
		return con, nil
	}

	// log that the detection failed
	aq.log.Noticef("unknown contract on %s", addr.String())

	// set as generic contract type if no other has ben detected
	*cType = types.AccountTypeContract
	return types.NewGenericContract(addr, block, trx), nil
}

// detectErc20Token identifies ERC20 token contracts by checking common contract end points.
func (aq *accountQueue) detectErc20Token(addr *common.Address, block *types.Block, trx *types.Transaction) *types.Contract {
	// try to get the token name
	name, err := aq.repo.Erc20Name(addr)
	if err != nil {
		return nil
	}

	// try to detect symbol
	if _, err := aq.repo.Erc20Symbol(addr); err != nil {
		return nil
	}

	// try to detect total supply
	if _, err := aq.repo.Erc20TotalSupply(addr); err != nil {
		return nil
	}

	// try to detect balance of
	if _, err := aq.repo.Erc20BalanceOf(addr, &testAddress); err != nil {
		return nil
	}

	// log what we do
	aq.log.Noticef("ERC20 token %s detected on %s", name, addr.String())
	return types.NewErc20Contract(addr, name, block, trx)
}

// detectStiContract identifies Staker Information contract by checking interface.
func (aq *accountQueue) detectStiContract(addr *common.Address, block *types.Block, trx *types.Transaction) *types.Contract {
	// detect the STI contract
	if aq.repo.IsStiContract(addr) {
		return types.NewStiContract(addr, block, trx)
	}

	return nil
}
