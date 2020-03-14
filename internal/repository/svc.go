/*
Repository package implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Opera/Lachesis full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import (
	"fantom-api-graphql/internal/logger"
	"sync"
)

// service represents a typical service in repository.
type service struct {
	name    string
	repo    Repository
	log     logger.Logger
	wg      *sync.WaitGroup
	sigStop chan bool
}

// newService creates a new service instance
func newService(name string, repo Repository, log logger.Logger, wg *sync.WaitGroup) service {
	return service{
		name:    name,
		repo:    repo,
		log:     log,
		wg:      wg,
		sigStop: make(chan bool, 1),
	}
}

// Close signals the service to stop.
func (se *service) close() {
	// log action
	se.log.Noticef("%s is closing", se.name)

	// send the signal and close the channel
	se.sigStop <- true
	close(se.sigStop)
}
