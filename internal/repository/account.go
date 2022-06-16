/*
Package repository implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Opera full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import (
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Account returns account at Opera blockchain for an address, nil if not found.
func (p *proxy) Account(addr *common.Address) (acc *types.Account, err error) {
	// try to get the account from cache
	acc = p.cache.PullAccount(addr)

	// we still don't know the account? try to manually construct it if possible
	if acc == nil {
		acc, err = p.getAccount(addr)
		if err != nil {
			return nil, err
		}
	}

	// return the account
	return acc, nil
}

// getAccount builds the account representation after validating it against Opera node.
func (p *proxy) getAccount(addr *common.Address) (*types.Account, error) {
	// any address given?
	if addr == nil {
		p.log.Error("no address given")
		return nil, fmt.Errorf("no address given")
	}

	// try to get the account from database first
	acc, err := p.db.Account(addr)
	if err != nil {
		p.log.Errorf("can not get the account %s; %s", addr.String(), err.Error())
		return nil, err
	}

	// found the account in database?
	if acc == nil {
		// log an unknown address
		p.log.Debugf("unknown address %s detected", addr.String())

		// at least we know the account existed
		acc = &types.Account{Address: *addr, Type: types.AccountTypeWallet}

		// check if this is a smart contract account; we log the error on the call
		acc.ContractTx, _ = p.db.ContractTransaction(addr)
	}

	// also keep a copy at the in-memory cache
	if err = p.cache.PushAccount(acc); err != nil {
		p.log.Warningf("can not keep account [%s] information in memory; %s", addr.Hex(), err.Error())
	}
	return acc, nil
}

// AccountBalance returns the current balance of an account at Opera blockchain.
func (p *proxy) AccountBalance(addr *common.Address) (*hexutil.Big, error) {
	return p.rpc.AccountBalance(addr)
}

// AccountNonce returns the current number of sent transactions of an account at Opera blockchain.
func (p *proxy) AccountNonce(addr *common.Address) (*hexutil.Uint64, error) {
	return p.rpc.AccountNonce(addr)
}

// AccountTransactions returns slice of AccountTransaction structure for a given account at Opera blockchain.
func (p *proxy) AccountTransactions(addr *common.Address, rec *common.Address, cursor *string, count int32) (*types.TransactionList, error) {
	// do we have an account?
	if addr == nil {
		return nil, fmt.Errorf("can not get transaction list for empty account")
	}

	// go to the database for the list of hashes of transaction searched
	return p.db.AccountTransactions(addr, rec, cursor, count)
}

// AccountsActive returns total number of accounts known to repository.
func (p *proxy) AccountsActive() (hexutil.Uint64, error) {
	val, err := p.db.AccountCount()
	return hexutil.Uint64(val), err
}

// AccountIsKnown checks if the account of the given address is known to the API server.
func (p *proxy) AccountIsKnown(addr *common.Address) bool {
	// try cache first
	stat := p.cache.CheckAccountKnown(addr)
	if nil != stat {
		return *stat
	}

	// check if the database knows the address
	known, err := p.db.IsAccountKnown(addr)
	if err != nil {
		p.log.Errorf("can not check account %s existence; %s", addr.String(), err.Error())
		return false
	}

	// if the account is known already, mark it in cache for faster resolving
	if known {
		p.cache.PushAccountKnown(addr)
	}
	return known
}

// StoreAccount adds specified account detail into the repository.
func (p *proxy) StoreAccount(acc *types.Account) error {
	// add this account to the database and remember it's been added
	err := p.db.AddAccount(acc)
	if err == nil {
		p.cache.PushAccountKnown(&acc.Address)
	}
	return err
}

// AccountMarkActivity marks the latest account activity in the repository.
func (p *proxy) AccountMarkActivity(addr *common.Address, ts uint64) error {
	return p.db.AccountMarkActivity(addr, ts)
}
