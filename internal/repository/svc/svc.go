/*
Package repository implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Opera/Lachesis full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

// Svc defines the interface required for a service
// to be manageable by the orchestrator.
type Svc interface {
	// Name returns the name of the service
	name() string

	// Init initializes the service
	init()

	// Run executes the service
	run()

	// Close signals the service to terminate
	close()
}
