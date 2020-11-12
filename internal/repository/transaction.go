/*
Package repository implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Opera/Lachesis full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import (
	"errors"
	"fantom-api-graphql/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	eth "github.com/ethereum/go-ethereum/rpc"
	"strings"
)

// ErrTransactionNotFound represents an error returned if a transaction can not be found.
var ErrTransactionNotFound = errors.New("requested transaction can not be found in Opera blockchain")

// AddTransaction notifies a new incoming transaction from blockchain to the repository.
func (p *proxy) AddTransaction(block *types.Block, trx *types.Transaction) error {
	// simply pass the transaction to DB handler for adding to off-chain database
	if err := p.db.AddTransaction(block, trx); err != nil {
		return err
	}

	// propagate to accounts
	if err := p.propagateTrxToAccounts(block, trx); err != nil {
		return err
	}

	// add smart contract to the persistent storage, too
	if trx.ContractAddress != nil {
		// check if the smart contract is an official ballot
		if err := p.processBallot(block, trx); err != nil {
			p.log.Critical(err)
			return err
		}
	}

	// everything seems to be ok
	return nil
}

// MarkTransactionProcessed marks given transaction as processed
// and ready to be served full to API users.
// AccountQueue processor call this function as a callback.
func (p *proxy) MarkTransactionProcessed(trx *types.Transaction) {
	// mark the transaction in database
	if err := p.db.MarkTransactionProcessed(trx); err != nil {
		p.log.Errorf("can not finish transaction %s processing", trx.Hash.String())
	}
}

// propagateTrxToAccounts pushes given transaction to accounts on both sides.
// The function is executed in a separate thread so it doesn't block
func (p *proxy) propagateTrxToAccounts(block *types.Block, trx *types.Transaction) error {
	// make sure the sender address is known
	// it's always either a regular wallet, or a contract we already know from it's creation
	p.orc.accountQueue <- &accountQueueRequest{
		blk:         block,
		trx:         trx,
		acc:         &types.Account{Address: *trx.To, ContractTx: nil, Type: types.AccountTypeWallet},
		trxCallback: nil,
	}

	// do we have a receiving account?
	if trx.To != nil {
		// just use the real recipient; if it's a contract we already know it and we don't have to
		// really analyze anything here
		p.orc.accountQueue <- &accountQueueRequest{
			blk:         block,
			trx:         trx,
			acc:         &types.Account{Address: *trx.To, ContractTx: nil, Type: types.AccountTypeWallet},
			trxCallback: nil,
		}
		return nil
	}

	// contract address must exist
	if trx.ContractAddress == nil {
		p.log.Criticalf("contract creation found, but no contract address [%s]", trx.Hash.String())
		return fmt.Errorf("invalid contract creation")
	}

	// add new smart contract
	p.orc.accountQueue <- &accountQueueRequest{
		blk:         block,
		trx:         trx,
		acc:         &types.Account{Address: *trx.ContractAddress, ContractTx: &trx.Hash, Type: types.AccountTypeWallet},
		trxCallback: nil,
	}
	return nil
}

// isBallotContract checks if the given contract transaction is an official
// Fantom ballot smart contract. We expect the ballots to come from configured set
// of internal addresses.
func (p *proxy) isBallotContract(trx *types.Transaction) bool {
	// make sure this is a contract
	if trx.ContractAddress == nil {
		return false
	}

	// make sure we know any ballot sources
	if p.ballotSources == nil {
		return false
	}

	// loop all ballot sources known and compare
	for _, addr := range p.ballotSources {
		if strings.EqualFold(addr, trx.From.String()) {
			return true
		}
	}

	return false
}

// processBallot validates if the given transaction is an official ballot contract
// and adds it into the list of ballots if so.
func (p *proxy) processBallot(block *types.Block, trx *types.Transaction) error {
	// is a ballot contract?
	if p.isBallotContract(trx) {
		// ballotIndex = db.transactionIndex(block, trx)
		ballot := types.Ballot{
			OrdinalIndex: types.TransactionIndex(block, trx),
			Address:      *trx.ContractAddress,
		}

		// collect ballot information from the ballot contract
		if err := p.rpc.LoadBallotDetails(&ballot); err != nil {
			p.log.Critical(err)
			return err
		}

		// add the ballot into the database; capture any possible error
		if err := p.db.AddBallot(&ballot); err != nil {
			p.log.Critical(err)
			return err
		}
	}

	return nil
}

// Transaction returns a transaction at Opera blockchain by a hash, nil if not found.
// If the transaction is not found, ErrTransactionNotFound error is returned.
func (p *proxy) Transaction(hash *types.Hash) (*types.Transaction, error) {
	// log
	p.log.Debugf("requested transaction %s", hash.String())

	// try to use the in-memory cache
	if trx := p.cache.PullTransaction(hash); trx != nil {
		// log and return
		p.log.Infof("transaction %s loaded from cache", hash.String())
		return trx, nil
	}

	// we need to go to RPC
	trx, err := p.rpc.Transaction(hash)
	if err != nil {
		// transaction simply not found?
		if err == eth.ErrNoResult {
			p.log.Warning("transaction not found in the blockchain")
			return nil, ErrTransactionNotFound
		}

		// something went wrong
		return nil, err
	}

	// log and return
	p.log.Debugf("transaction %s loaded from rpc", hash.String())

	// push the transaction to the cache to speed things up next time
	// we don't cache pending transactions since it would cause issues
	// when re-loading data of such transactions on the client side
	if trx.BlockHash != nil {
		err = p.cache.PushTransaction(trx)
		if err != nil {
			p.log.Errorf("can not store transaction in cache; %s", err.Error())
		}
	} else {
		// inform about non-cached trx
		p.log.Debugf("pending transaction [%s] not cached yet", trx.Hash)
	}

	// return the value
	return trx, nil
}

// SendTransaction sends raw signed and RLP encoded transaction to the block chain.
func (p *proxy) SendTransaction(tx hexutil.Bytes) (*types.Transaction, error) {
	// log
	p.log.Debugf("requested transaction submit for %s", tx.String())

	// try to send it and get the tx hash
	hash, err := p.rpc.SendTransaction(tx)
	if err != nil {
		p.log.Errorf("can not send transaction to block chain; %s", err.Error())
		return nil, err
	}

	// we do have the hash so we can use it to get the transaction details
	// we always need to go to RPC and we will not try to store the transaction in cache yet
	trx, err := p.rpc.Transaction(hash)
	if err != nil {
		// transaction simply not found?
		if err == eth.ErrNoResult {
			p.log.Warning("transaction not found in the blockchain")
			return nil, ErrTransactionNotFound
		}

		// something went wrong
		return nil, err
	}

	// log and return
	p.log.Debugf("transaction %s created and found", hash.String())

	// return the value
	return trx, nil
}

// Transactions pulls list of transaction hashes starting on the specified cursor.
// If the initial transaction cursor is not provided, we start on top, or bottom based on count value.
//
// No-number boundaries are handled as follows:
// 	- For positive count we start from the most recent transaction and scan to older transactions.
// 	- For negative count we start from the first transaction and scan to newer transactions.
func (p *proxy) Transactions(cursor *string, count int32) (*types.TransactionHashList, error) {
	// go to the database for the list of hashes of transaction searched
	return p.db.Transactions(cursor, count, nil)
}

// TransactionsCount returns total number of transactions in the block chain.
func (p *proxy) TransactionsCount() (hexutil.Uint64, error) {
	// get the number of transactions registered
	tc, err := p.db.TransactionsCount()
	if err != nil {
		return hexutil.Uint64(0), err
	}

	return hexutil.Uint64(tc), nil
}
