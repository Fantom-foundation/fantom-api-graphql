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
	"fantom-api-graphql/internal/repository/rpc"
	"fantom-api-graphql/internal/repository/rpc/contracts"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"sync"
)

const (
	// accountQueueLength represents how many accounts can be pushed
	// into the queue for processing at once
	accountQueueLength = 50000

	// sfcCheckBelowBlock represents the highest block number we try to detect
	// SFC contract, above this block the contract should already be known and we can
	// skip the check
	sfcCheckBelowBlock = 100000
)

// testAddress represents an address used to test an account reference
var testAddress = common.HexToAddress("0xabc00FA001230012300aBc0012300Fa00FACE000")

// accountEvent represents an account queue processing request.
// This can be a simple account existence verification, or an extended
// contract analyze.
type accountEvent struct {
	acc *types.Account
	blk *types.Block
	trx *types.Transaction
	wg  *sync.WaitGroup
}

// accountDispatcher implements account analyzer queue
type accountDispatcher struct {
	service
	buffer chan *accountEvent
}

// newAccountDispatcher creates new blockchain account analyzer queue service.
func newAccountDispatcher(buffer chan *accountEvent, repo Repository, log logger.Logger, wg *sync.WaitGroup) *accountDispatcher {
	// create new instance
	return &accountDispatcher{
		service: newService("account dispatcher", repo, log, wg),
		buffer:  buffer,
	}
}

// run starts the account queue to life.
func (acd *accountDispatcher) run() {
	// start blockScanner
	acd.wg.Add(1)
	go acd.dispatch()
}

// dispatch runs the main account requests monitor and dispatcher
// loop in a separate thread.
func (acd *accountDispatcher) dispatch() {
	// log action
	acd.log.Notice("account dispatcher is running")

	// don't forget to sign off after we are done
	defer func() {
		acd.log.Noticef("%s is closed", acd.name)
		acd.wg.Done()
	}()

	// wait for either stop signal, or an account request
	var err error
	for {
		select {
		case req := <-acd.buffer:
			// log what we do
			acd.log.Debugf("account %s received for processing", req.acc.Address.String())

			// process the account request into the database
			if err = acd.processAccount(req.acc, req.blk, req.trx); err != nil {
				acd.log.Errorf("can not process account %s; %s", req.acc.Address.String(), err.Error())
			}

			// we are done with the account
			req.wg.Done()
		case <-acd.sigStop:
			// stop signal received?
			return
		}
	}
}

// processAccount processes account into the database
// based on the account details
func (acd *accountDispatcher) processAccount(acc *types.Account, block *types.Block, trx *types.Transaction) error {
	// check if the account is new; if we already know it, we are done
	if acd.repo.AccountIsKnown(&acc.Address) {
		return acd.repo.AccountMarkActivity(&acc.Address, uint64(block.TimeStamp))
	}

	// is this a simple wallet/account?
	if acc.ContractTx == nil {
		return acd.processSimple(acc, block, trx)
	}
	return acd.processContract(acc, block, trx)
}

// processSimple processes a simple non-contract account into the database
// based on the account details (it still could be the SFC, be cautious about it)
func (acd *accountDispatcher) processSimple(acc *types.Account, block *types.Block, trx *types.Transaction) error {
	// notify new account detected
	acd.log.Debugf("found new account %s", acc.Address.String())

	// check if the target address is not an SFC contract
	acd.checkSfcContract(acc, block, trx)

	// add the account into the database
	err := acd.repo.StoreAccount(acc)
	if err != nil {
		acd.log.Errorf("can not add account %s; %s", acc.Address.String(), err.Error())
		return err
	}
	return nil
}

// checkSfcContract verifies if the target account is the SFC contract
// and if so, it adds the SFC target with a different type.
func (acd *accountDispatcher) checkSfcContract(acc *types.Account, block *types.Block, trx *types.Transaction) {
	// act on SFC detection
	if uint64(block.Number) < sfcCheckBelowBlock && acd.repo.IsSfcContract(&acc.Address) {
		// change the type to SFC contract
		acc.Type = types.AccountTypeSFC

		// get the SFC version
		ver, err := acd.repo.SfcVersion()
		if err == nil {
			// log what we found
			acd.log.Debugf("detected SFC contract %d.%d.%d", byte((ver>>16)&255), byte((ver>>8)&255), byte(ver&255))

			// add the contract
			if err := acd.repo.StoreContract(types.NewSfcContract(&acc.Address, uint64(ver), block, trx)); err != nil {
				acd.log.Errorf("can not add the SFC contract at %s; %s", acc.Address.String(), err.Error())
			}
		}
	}
}

// processContract processes contract account with detection.
func (acd *accountDispatcher) processContract(acc *types.Account, block *types.Block, trx *types.Transaction) error {
	// log what we do
	acd.log.Debugf("account %s is a smart contract, analyzing", acc.Address.String())

	// detect and identify contract
	con, err := acd.detectContract(&acc.Address, &acc.Type, block, trx)
	if err != nil {
		acd.log.Errorf("can not identify contract at %s; %s", acc.Address.String(), err.Error())
		return err
	}

	// insert the contract record if possible
	if con != nil {
		err = acd.repo.StoreContract(con)
		if err != nil {
			acd.log.Errorf("can not add contract at %s; %s", acc.Address.String(), err.Error())
			return err
		}
	}

	// add the account identified into the database
	err = acd.repo.StoreAccount(acc)
	if err != nil {
		acd.log.Errorf("can not add account %s; %s", acc.Address.String(), err.Error())
		return err
	}
	return nil
}

// detectContract tries to identify the contract type.
func (acd *accountDispatcher) detectContract(addr *common.Address, cType *string, block *types.Block, trx *types.Transaction) (*types.Contract, error) {
	// identify ERC20 token
	con := acd.detectErc20(addr, block, trx)
	if con != nil {
		*cType = types.AccountTypeERC20Token
		return con, nil
	}

	// log that the detection failed
	acd.log.Noticef("unknown contract at %s", addr.String())

	// set as generic contract type if no other has ben detected
	*cType = types.AccountTypeContract
	return types.NewGenericContract(addr, block, trx), nil
}

// detectErc20 identifies ERC20 token contracts by checking common contract end points.
func (acd *accountDispatcher) detectErc20(addr *common.Address, block *types.Block, trx *types.Transaction) *types.Contract {
	// try to get the token name
	name, err := acd.repo.Erc20Name(addr)
	if err != nil {
		return nil
	}

	// try to detect symbol
	if _, err := acd.repo.Erc20Symbol(addr); err != nil {
		return nil
	}

	// try to detect balance of
	if _, err := acd.repo.Erc20BalanceOf(addr, &testAddress); err != nil {
		return nil
	}

	// try to detect total supply; if not found this is probably ERC721
	if _, err := acd.repo.Erc20TotalSupply(addr); err != nil {
		acd.log.Noticef("ERC721 token %s detected at %s", name, addr.String())
		return types.NewErcTokenContract(addr, name, block, trx, types.AccountTypeERC20Token, rpc.ERC721TokenABI)
	}

	// log what we do
	acd.log.Noticef("ERC20 token %s detected at %s", name, addr.String())
	return types.NewErcTokenContract(addr, name, block, trx, types.AccountTypeERC721Token, contracts.ERCTwentyABI)
}

// detectStiContract identifies Staker Information contract by checking interface.
func (acd *accountDispatcher) detectStiContract(addr *common.Address, block *types.Block, trx *types.Transaction) *types.Contract {
	// detect the STI contract
	if acd.repo.IsStiContract(addr) {
		return types.NewStiContract(addr, block, trx)
	}
	return nil
}
