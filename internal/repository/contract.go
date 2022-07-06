/*
Package repository implements repository for handling fast and efficient access to data required
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
)

// Contract extract a smart contract information by account address, if available.
func (p *proxy) Contract(addr *common.Address) (*types.Contract, error) {
	// try cache first
	sc := p.cache.PullContract(addr)

	// we still don't know the contract? call the db for that
	if sc == nil {
		var err error
		sc, err = p.db.Contract(addr)
		if err != nil {
			return nil, err
		}

		// found the contract? push to cache for future use
		if sc != nil {
			if err = p.cache.PushContract(sc); err != nil {
				p.log.Criticalf("can not cache contract %s; %s", addr.String(), err.Error())
			}
		}
	}

	return sc, nil
}

// Contracts returns list of smart contracts at Opera blockchain.
func (p *proxy) Contracts(validatedOnly bool, cursor *string, count int32) (*types.ContractList, error) {
	// go to the database for the list of contracts searched
	return p.db.Contracts(validatedOnly, cursor, count)
}

// ValidateContract tries to validate contract byte code using
// provided source code. If successful, the contract information
// is updated the repository.
func (p *proxy) ValidateContract(_ *types.Contract) error {
	return fmt.Errorf("contract validation not supported")
}

// StoreContract adds new contract into the repository.
func (p *proxy) StoreContract(con *types.Contract) error {
	// is the known contract which will be updated?
	isUpdate := p.db.IsContractKnown(&con.Address)

	// do the add/update op
	if err := p.db.AddContract(con); err != nil {
		p.log.Errorf("contract %s store failed; %s", con.Address.String(), err.Error())
		return err
	}

	// re-scan transactions of the contract so they are up-to-date with their calls analysis
	if isUpdate {
		// log what we have done here
		p.log.Debugf("updated known contract at %s", con.Address.String())
		p.cache.EvictContract(&con.Address)
	}
	return nil
}
