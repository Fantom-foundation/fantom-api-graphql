// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"context"
	"fantom-api-graphql/internal/types"
	"time"
)

// onBlockChannelCapacity is the number of new block events held in memory for being broadcast to subscriber.
const onBlockChannelCapacity = 500

// subscriptOnBlock represents reference to a subscriber to onBlock events broadcast.
type subscriptOnBlock struct {
	stop   <-chan struct{}
	events chan<- *Block
}

// OnBlock resolves subscription to new blocks event broadcast.
func (rs *rootResolver) OnBlock(ctx context.Context) <-chan *Block {
	// make the stream
	c := make(chan *Block, onBlockChannelCapacity)

	// subscribe to event dispatch
	rs.subscribeOnBlock <- &subscriptOnBlock{
		stop:   ctx.Done(),
		events: c,
	}
	return c
}

// addBlockSubscriber adds a new subscription to onBlock events.
func (rs *rootResolver) addBlockSubscriber(sub *subscriptOnBlock) {
	id, err := uuid()
	if err == nil {
		// add the subscriber to the map
		rs.blockSubscribers[id] = sub
	} else {
		// log critical issue
		log.Critical("can not generate UUID for new onBlock subscriber")
		log.Critical(err)
	}
}

// dispatchOnBlock dispatches onBlock event to registered subscribers.
func (rs *rootResolver) dispatchOnBlock(blk *types.Block) {
	// prep the block
	block := NewBlock(blk)

	// broadcast the event in separate go routines so we don't block here
	for id, sub := range rs.blockSubscribers {
		go rs.notifyOnBlock(block, sub, id)
	}
}

// notifyOnBlock broadcasts onBlock event to given subscriber.
func (rs *rootResolver) notifyOnBlock(block *Block, sub *subscriptOnBlock, id string) {
	// check if the context isn't already closed in which case we just unsub and leave
	select {
	case <-sub.stop:
		rs.unsubscribeOnBlock <- id
		return
	default:
	}

	// broadcast
	select {
	case <-sub.stop:
		// just unsub on broken context
		rs.unsubscribeOnBlock <- id

	case sub.events <- block:
		// push the block to subscriber

	case <-time.After(time.Second):
		// timeout reached without response? just remove the subscriber
		rs.unsubscribeOnBlock <- id
	}
}
