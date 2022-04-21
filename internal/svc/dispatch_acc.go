// Package svc implements blockchain data processing services.
package svc

import (
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"time"
)

// testAddress represents an address used to test an account reference
var testAddress = common.HexToAddress("0xabc00FA001230012300aBc0012300Fa00FACE000")

var erc721InterfaceId = [4]byte{0x80, 0xac, 0x58, 0xcd}  // ERC-721: 0x80ac58cd
var erc1155InterfaceId = [4]byte{0xd9, 0xb6, 0x7a, 0x26} // ERC-1155: 0xd9b67a26

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
	defer func() {
		close(acd.sigStop)
		acd.mgr.finished(acd)
	}()

	for {
		select {
		case <-acd.sigStop:
			return

		case acc, ok := <-acd.inAccount:
			if !ok {
				log.Noticef("account channel closed, terminating %s", acd.name())
				return
			}

			if !repo.AccountIsKnown(acc.addr) {
				err := acd.add(acc)
				if err != nil {
					log.Errorf("account %s processing failed; %s", acc.addr.String(), err.Error())
				}
			}

			// the transaction queue waits for this to finish
			acc.watchDog.Done()
		}
	}
}

// add a newly noticed unique account to the API server.
func (acd *accDispatcher) add(acc *eventAcc) error {
	log.Debugf("found a new account %s", acc.addr.String())

	at, name, isContract := acd.accountType(acc.addr)
	act := types.Account{
		Name:            name,
		Address:         *acc.addr,
		AccountType:     at,
		IsContract:      isContract,
		FirstAppearance: time.Unix(int64(acc.blk.TimeStamp), 0),
	}

	if isContract {
		act.DeployedBy = &acc.trx.From
		act.DeploymentTrx = &acc.trx.Hash
	}

	return repo.StoreAccount(&act)
}

// accountType tries to detect the type of the given account.
func (acd *accDispatcher) accountType(adr *common.Address) (int, string, bool) {
	// if there is no code behind the address we assume this is just a wallet
	if !repo.AccountHasCode(adr) {
		return types.AccountTypeWallet, "", false
	}

	log.Debugf("contract found at %s", adr.String())

	// is it a token contract?
	ct, name, err := acd.contractTokenType(adr)
	if err == nil {
		return ct, name, true
	}

	// is it the Opera's special purpose contract?
	if repo.IsSfcContract(adr) {
		return types.AccountTypeSFCContract, "SFC", true
	}

	log.Debugf("unknown contract at %s", adr.String())
	return types.AccountTypeContract, "", true
}

// contractTokenType tries to check if the contract is some type of tokens contract (ERC20 and/or an NFT).
func (acd *accDispatcher) contractTokenType(addr *common.Address) (int, string, error) {
	isErc1155, err := repo.Erc165SupportsInterface(addr, erc1155InterfaceId)
	if err == nil && isErc1155 {
		log.Noticef("ERC1155 multi-token detected at %s", addr.String())
		return types.AccountTypeERC1155Contract, "ERC-1155", nil
	}

	isErc721, name := acd.detectErc721Token(addr)
	if err == nil && isErc721 {
		log.Noticef("ERC721 NFT token detected at %s", addr.String())
		return types.AccountTypeERC721Contract, name, nil
	}

	isErc20, name := acd.detectErc20Token(addr)
	if isErc20 {
		log.Noticef("ERC20 token %s detected at %s", name, addr.String())
		return types.AccountTypeERC20Contract, name, nil
	}

	return types.AccountTypeContract, "", fmt.Errorf("not a detectable token contract")
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

// detectErc721Token checks for ERC-721 interface and extracts collection name, if available.
func (acd *accDispatcher) detectErc721Token(addr *common.Address) (isErc721 bool, name string) {
	isErc721, err := repo.Erc165SupportsInterface(addr, erc721InterfaceId)
	if err != nil || !isErc721 {
		return false, ""
	}

	// try to detect name - but it is optional for ERC721
	name, err = repo.TokenNameAttempt(addr)
	if err != nil {
		return true, "ERC-721"
	}

	return true, name
}
