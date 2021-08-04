// Package svc implements blockchain data processing services.
package svc

import (
	"fantom-api-graphql/internal/config"
	"fantom-api-graphql/internal/logger"
	"fantom-api-graphql/internal/repository"
	"sync"
)

// repo represents the repository used by services to handle data exchange.
var repo repository.Repository

// cfg holds reference to the server configuration.
var cfg *config.Config

// log represents the logger to be used by services to report state.
var log logger.Logger

// manager represents a singleton instance of the service manager.
var manager *ServiceManager

// onceManager is the sync object used to make sure the ServiceManager
// is instantiated only once on the first demand.
var onceManager sync.Once

// SetConfig sets the repository configuration to be used to establish
// and maintain external repository connections.
func SetConfig(c *config.Config) {
	cfg = c
}

// SetLogger sets the repository logger to be used to collect logging info.
func SetLogger(l logger.Logger) {
	log = l
}

// Manager provides access to the singleton instance of the SVC manager.
func Manager() *ServiceManager {
	// make sure to instantiate the Repository only once
	onceManager.Do(func() {
		manager = newServiceManager()
	})
	return manager
}
