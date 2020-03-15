// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"context"
	"fantom-api-graphql/internal/logger"
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"sync"
)

// subscriptionQueueCapacity is the amount of subscriptions we buffer for processing.
const subscriptionQueueCapacity = 100
const subscriptionInitialCapacity = 100

// ApiResolver represents the API interface expected to handle API access points
type ApiResolver interface {
	// Account resolves blockchain account by address.
	Account(struct{ Address common.Address }) (*Account, error)

	// Block resolves blockchain block by number or by hash. If neither is provided, the most recent block is given.
	Block(*struct {
		Number *hexutil.Uint64
		Hash   *types.Hash
	}) (*Block, error)

	// Blocks resolves list of blockchain blocks encapsulated in a listable structure.
	Blocks(*struct {
		Cursor *Cursor
		Count  int32
	}) (*BlockList, error)

	// Transaction resolves blockchain transaction by hash.
	Transaction(*struct{ Hash types.Hash }) (*Transaction, error)

	// Transactions resolves list of blockchain transactions encapsulated in a listable structure.
	Transactions(*struct {
		Cursor *Cursor
		Count  int32
	}) (*TransactionList, error)

	// OnBlock resolves subscription to new blocks event broadcast.
	OnBlock(ctx context.Context) <-chan *Block

	// OnTransaction resolves subscription to new transactions event broadcast.
	OnTransaction(ctx context.Context) <-chan *Transaction

	// Close terminates resolver broadcast management.
	Close()
}

// rootResolver represents the ApiResolver implementation.
type rootResolver struct {
	log  logger.Logger
	repo repository.Repository

	// service terminator
	wg      sync.WaitGroup
	sigStop chan bool

	// blocks subscriptions management
	subscribeOnBlock   chan *subscriptOnBlock
	unsubscribeOnBlock chan string
	blockSubscribers   map[string]*subscriptOnBlock
	onBlockEvents      chan *types.Block

	// transaction subscriptions management
	subscribeOnTrx   chan *subscriptOnTrx
	unsubscribeOnTrx chan string
	trxSubscribers   map[string]*subscriptOnTrx
	onTrxEvents      chan *types.Transaction
}

// New creates a new root resolver instance and initializes it's internal structure.
func New(log logger.Logger, repo repository.Repository) ApiResolver {
	// create new resolver
	rs := rootResolver{
		log:  log,
		repo: repo,

		// create terminator
		sigStop: make(chan bool, 1),

		// block events subscription basics
		subscribeOnBlock:   make(chan *subscriptOnBlock, subscriptionQueueCapacity),
		unsubscribeOnBlock: make(chan string, subscriptionQueueCapacity),
		blockSubscribers:   make(map[string]*subscriptOnBlock, subscriptionInitialCapacity),
		onBlockEvents:      make(chan *types.Block, onBlockChannelCapacity),

		// block events subscription basics
		subscribeOnTrx:   make(chan *subscriptOnTrx, subscriptionQueueCapacity),
		unsubscribeOnTrx: make(chan string, subscriptionQueueCapacity),
		trxSubscribers:   make(map[string]*subscriptOnTrx, subscriptionInitialCapacity),
		onTrxEvents:      make(chan *types.Transaction, onBlockChannelCapacity),
	}

	// register event channels with repository
	repo.SetBlockChannel(rs.onBlockEvents)
	repo.SetTrxChannel(rs.onTrxEvents)

	// handle broadcast and subscriptions in a separate routine
	rs.wg.Add(1)
	go rs.run()

	return &rs
}

// Close terminates resolver's broadcast service.
func (rs *rootResolver) Close() {
	// log
	rs.log.Notice("GraphQL resolver is closing")

	// send the signal
	rs.sigStop <- true
	rs.wg.Wait()
}

// run monitors and handles subscriptions and broadcasts incoming events to their subscribers.
func (rs *rootResolver) run() {
	// sign off on leaving
	defer func() {
		// terminate
		rs.log.Notice("GraphQL resolver done")
		rs.wg.Done()
	}()

	// log action
	rs.log.Notice("GraphQL resolver started")

	// main loop waits for data on any channel and act upon it
	for {
		select {
		case <-rs.sigStop:
			return

		case id := <-rs.unsubscribeOnBlock:
			delete(rs.blockSubscribers, id)

		case id := <-rs.unsubscribeOnTrx:
			delete(rs.trxSubscribers, id)

		case sub := <-rs.subscribeOnBlock:
			rs.addBlockSubscriber(sub)

		case sub := <-rs.subscribeOnTrx:
			rs.addTrxSubscriber(sub)

		case evt := <-rs.onBlockEvents:
			rs.dispatchOnBlock(evt)

		case evt := <-rs.onTrxEvents:
			rs.dispatchOnTransaction(evt)
		}
	}
}
