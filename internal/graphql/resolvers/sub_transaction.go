// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"context"
	"fantom-api-graphql/internal/types"
	"time"
)

// onTrxChannelCapacity is the number of new transaction events held in memory for being broadcast to subscriber.
const onTrxChannelCapacity = 500

// subscriptOnTrx represents reference to a subscriber to onTransaction events broadcast.
type subscriptOnTrx struct {
	stop   <-chan struct{}
	events chan<- *Transaction
}

// OnBlock resolves subscription to new blocks event broadcast.
func (rs *rootResolver) OnTransaction(ctx context.Context) <-chan *Transaction {
	// make the stream
	c := make(chan *Transaction, onTrxChannelCapacity)

	// subscribe to event dispatch
	rs.subscribeOnTrx <- &subscriptOnTrx{
		stop:   ctx.Done(),
		events: c,
	}

	return c
}

// addTrxSubscriber adds a new subscription to onTransaction events.
func (rs *rootResolver) addTrxSubscriber(sub *subscriptOnTrx) {
	id, err := uuid()
	if err == nil {
		// add the subscriber to the map
		rs.trxSubscribers[id] = sub
	} else {
		// log critical issue
		log.Critical("can not generate UUID for new onTransaction subscriber")
		log.Critical(err)
	}
}

// dispatchOnTransaction dispatches onTransaction event to registered subscribers.
func (rs *rootResolver) dispatchOnTransaction(trx *types.Transaction) {
	// prep the block
	transaction := NewTransaction(trx)

	// broadcast the event in separate go routines so we don't block here
	for id, sub := range rs.trxSubscribers {
		go rs.notifyOnTransaction(transaction, sub, id)
	}
}

// notifyOnTransaction broadcasts onBlock event to given subscriber.
func (rs *rootResolver) notifyOnTransaction(trx *Transaction, sub *subscriptOnTrx, id string) {
	// check if the context isn't already closed in which case we just unsub and leave
	select {
	case <-sub.stop:
		rs.unsubscribeOnTrx <- id
		return
	default:
	}

	// broadcast
	select {
	case <-sub.stop:
		// just unsub on broken context
		rs.unsubscribeOnTrx <- id

	case sub.events <- trx:
		// push the block to subscriber

	case <-time.After(time.Second):
		// timeout reached without response? just remove the subscriber
		rs.unsubscribeOnTrx <- id
	}
}
