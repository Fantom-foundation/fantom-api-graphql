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
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	eth "github.com/ethereum/go-ethereum/rpc"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
)

// transactionReScanDepth represents the max number of transactions we rescan
// on contract core update
const transactionReScanDepth = 2147483647

// ErrTransactionNotFound represents an error returned if a transaction can not be found.
var ErrTransactionNotFound = errors.New("requested transaction can not be found in Opera blockchain")

// TransactionAdd notifies a new incoming transaction from blockchain to the repository.
func (p *proxy) TransactionAdd(block *types.Block, trx *types.Transaction) error {
	// simply pass the transaction to DB handler for adding to off-chain database
	if err := p.db.AddTransaction(block, trx); err != nil {
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

	// propagate to accounts
	return p.propagateTrxToAccounts(block, trx)
}

// TransactionUpdate modifies a transaction record in the repository.
func (p *proxy) TransactionUpdate(trx *types.Transaction) error {
	return p.db.UpdateTransaction(trx)
}

// TransactionMarkPropagated marks given transaction as processed
// and ready to be served full to API users.
// AccountQueue processor call this function as a callback.
func (p *proxy) TransactionMarkPropagated(trx *types.Transaction) {
	// mark the transaction in database
	if err := p.db.TransactionMarkPropagated(trx); err != nil {
		p.log.Errorf("failed transaction %s state update; %s", trx.Hash.String(), err.Error())
		return
	}

	// log what we done
	p.log.Infof("transaction %s propagated", trx.Hash.String())

	// may this be a contract call transaction?
	if IsLikelyAContractCall(trx) {
		// log what we done
		p.log.Infof("possible contract call at %s", trx.Hash.String())

		// sign this transaction for additional analysis so we can detect
		// and sort out what type of contract call is this, if any
		p.orc.contractCallsQueue <- trx
	}
}

// Transaction returns a transaction at Opera blockchain by a hash, nil if not found.
// If the transaction is not found, ErrTransactionNotFound error is returned.
func (p *proxy) Transaction(hash *types.Hash) (*types.Transaction, error) {
	// log
	p.log.Debugf("requested transaction %s", hash.String())

	// try to use the in-memory cache
	if trx := p.cache.PullTransaction(hash); trx != nil {
		// log and return
		p.log.Debugf("transaction %s loaded from cache", hash.String())
		return trx, nil
	}

	// return the value
	return p.loadTransaction(hash)
}

// loadTransaction loads the transaction hard way using RPC and DB.
func (p *proxy) loadTransaction(hash *types.Hash) (*types.Transaction, error) {
	// we need to go to RPC
	trx, err := p.rpc.Transaction(hash)
	if err != nil {
		return nil, err
	}

	// push the transaction to the cache to speed things up next time
	// we don't cache pending transactions since it would cause issues
	// when re-loading data of such transactions on the client side
	if trx.BlockHash == nil {
		p.log.Debugf("pending transaction %s found", trx.Hash)
		return trx, nil
	}

	return p.confirmedTransaction(trx)
}

// extendConfirmedTransaction loads additional transaction information if available for transaction
// confirmed inside the blockchain.
func (p *proxy) confirmedTransaction(trx *types.Transaction) (*types.Transaction, error) {
	// do we have the transaction at all?
	if trx == nil {
		p.log.Criticalf("invalid transaction reference received")
		return nil, fmt.Errorf("nil transaction received")
	}

	// pull details of the transaction if available
	err := p.db.TransactionDetails(trx)
	if err != nil {
		p.log.Errorf("can not get transaction %s details from database; %s", trx.Hash.String(), err.Error())
		return trx, err
	}

	// cache the transaction if it has been found in the db already
	if trx.OrdinalIndex > 0 {
		// push the TRX to the cache
		err = p.cache.PushTransaction(trx)
		if err != nil {
			p.log.Errorf("can not store transaction %s in cache; %s", trx.Hash.String(), err.Error())
		}

		// log we have it all
		p.log.Debugf("transaction %s details loaded", trx.Hash.String())
	}

	// log and return
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

// propagateTrxToAccounts pushes given transaction to accounts on both sides.
// The function is executed in a separate thread so it doesn't block
func (p *proxy) propagateTrxToAccounts(block *types.Block, trx *types.Transaction) error {
	// log what we do here
	p.log.Debugf("propagating %s to accounts", trx.Hash.String())

	// the sender is always present
	p.queueAccount(block, trx, trx.From, nil, nil)

	// do we have a recipient? if there is no contract, we are done
	if trx.To != nil && trx.ContractAddress == nil {
		p.queueAccount(block, trx, *trx.To, nil, p.TransactionMarkPropagated)
		return nil
	}

	// recipient address still present?
	if trx.To != nil {
		p.queueAccount(block, trx, *trx.To, nil, nil)
	}

	// log and queue the new contract
	p.log.Debugf("contract %s found at %s", trx.ContractAddress.String(), trx.Hash.String())
	p.queueAccount(block, trx, *trx.ContractAddress, &trx.Hash, p.TransactionMarkPropagated)

	return nil
}

// queueAccount pushes given account to processing queue.
func (p *proxy) queueAccount(
	block *types.Block,
	trx *types.Transaction,
	addr common.Address,
	ccHash *types.Hash,
	callback func(*types.Transaction),
) {
	// what account type is this
	t := types.AccountTypeWallet
	if ccHash != nil {
		t = types.AccountTypeContract
	}

	// push the request
	p.orc.accountQueue <- &accountQueueRequest{
		blk: block,
		trx: trx,
		acc: &types.Account{
			Address:      addr,
			ContractTx:   ccHash,
			Type:         t,
			LastActivity: block.TimeStamp,
		},
		trxCallback: callback,
	}

	// log what we do here
	p.log.Debugf("%s %s queued from %s", t, addr.String(), trx.Hash.String())
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

// transactionRescanContractCalls initiates a rescan of all the contract calls to make sure
// corresponding transactions are up-to-date with their call analysis.
// The function is called in a separate thread to use benefit from parallel approach.
func (p *proxy) transactionRescanContractCalls(con *types.Contract) {
	// log what we do
	p.log.Debugf("initiating calls re-scan for %s", con.Address.String())

	// make the filter for transactions targeted to the contract
	// we use constant for the recipient field here, see db.fiTransactionRecipient
	filter := bson.D{{"to", con.Address.String()}}

	// list all the transaction targeted on the contract so far; we ask for as many trx as we can
	tl, err := p.db.Transactions(nil, transactionReScanDepth, &filter)
	if err != nil {
		p.log.Criticalf("can not pull transactions list for %s; %s", con.Address.String(), err.Error())
		return
	}

	// loop all the transactions and schedule their re-scan
	for _, h := range tl.Collection {
		// pull the transaction
		trx, err := p.Transaction(h)
		if err != nil {
			p.log.Criticalf("can not pull transaction %s; %s", h.String(), err.Error())
			continue
		}

		// add the transaction to the scanner; this will pause the thread if the queue is clogged
		p.orc.contractCallsQueue <- trx
	}
}
