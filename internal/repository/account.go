/*
Repository package implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Opera/Lachesis full node for blockchain interaction. Mongo database
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
func (p *proxy) Account(addr *common.Address) (*types.Account, error) {
	// try to get the account from cache
	acc := p.cache.PullAccount(addr)

	// we still don't know the account? try to manually construct it if possible
	if acc == nil {
		var err error
		acc, err = p.getAccount(addr)
		if err != nil {
			return nil, err
		}
	}

	// return the account
	return acc, nil
}

// getAccount builds the account representation after validating it against Lachesis node.
func (p *proxy) getAccount(addr *common.Address) (*types.Account, error) {
	// at least we know the account existed
	acc := types.Account{Address: *addr}

	// also keep a copy at the in-memory cache
	err := p.cache.PushAccount(&acc)
	if err != nil {
		p.log.Warningf("can not keep account [%s] information in memory; %s", addr.Hex(), err.Error())
	}

	return &acc, nil
}

// AccountBalance returns the current balance of an account at Opera blockchain.
func (p *proxy) AccountBalance(acc *types.Account) (*hexutil.Big, error) {
	return p.rpc.AccountBalance(&acc.Address)
}

// AccountNonce returns the current number of sent transactions of an account at Opera blockchain.
func (p *proxy) AccountNonce(acc *types.Account) (*hexutil.Uint64, error) {
	val, err := p.rpc.AccountNonce(&acc.Address)
	if err != nil {
		return nil, err
	}

	// make the value and return
	nonce := hexutil.Uint64(val)
	return &nonce, nil
}

// AccountTransactions returns slice of AccountTransaction structure for a given account at Opera blockchain.
func (p *proxy) AccountTransactions(acc *types.Account, cursor *string, count int32) (*types.TransactionHashList, error) {
	// do we have an account?
	if acc == nil {
		return nil, fmt.Errorf("can not get transaction list for empty account")
	}

	// go to the database for the list of hashes of transaction searched
	return p.db.AccountTransactions(acc, cursor, count)
}

// AccountsActive returns total number of accounts known to repository.
func (p *proxy) AccountsActive() (hexutil.Uint64, error) {
	return p.db.AccountCount()
}
