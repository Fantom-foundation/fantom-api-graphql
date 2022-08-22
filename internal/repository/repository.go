/*
Package repository implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Opera full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import (
	"fantom-api-graphql/internal/config"
	"fantom-api-graphql/internal/logger"
	"fantom-api-graphql/internal/repository/cache"
	"fantom-api-graphql/internal/repository/db"
	"fantom-api-graphql/internal/repository/geoip"
	"fantom-api-graphql/internal/repository/p2p"
	"fantom-api-graphql/internal/repository/rpc"
	"fmt"
	"golang.org/x/sync/singleflight"
	"sync"
)

// repo represents an instance of the Repository manager.
var repo Repository

// onceRepo is the sync object used to make sure the Repository
// is instantiated only once on the first demand.
var onceRepo sync.Once

// config represents the configuration setup used by the repository
// to establish and maintain required connectivity to external services
// as needed.
var cfg *config.Config

// log represents the logger to be used by the repository.
var log logger.Logger

// SetConfig sets the repository configuration to be used to establish
// and maintain external repository connections.
func SetConfig(c *config.Config) {
	cfg = c
}

// SetLogger sets the repository logger to be used to collect logging info.
func SetLogger(l logger.Logger) {
	log = l
}

// R provides access to the singleton instance of the Repository.
func R() Repository {
	// make sure to instantiate the Repository only once
	onceRepo.Do(func() {
		repo = newRepository()
	})
	return repo
}

// Proxy represents Repository interface implementation and controls access to data
// trough several low level bridges.
type proxy struct {
	cache *cache.MemBridge
	db    *db.MongoDbBridge
	rpc   *rpc.FtmBridge
	geoip *geoip.Bridge
	log   logger.Logger
	cfg   *config.Config

	// transaction estimator counter
	txCount uint64

	// we need a Group to use single flight to control price pulls
	apiRequestGroup singleflight.Group

	// governance contracts reference
	govContracts map[string]config.GovernanceContract

	// smart contract compilers
	solCompiler string
}

// newRepository creates new instance of Repository implementation, namely proxy structure.
func newRepository() Repository {
	if cfg == nil {
		panic(fmt.Errorf("missing configuration"))
	}
	if log == nil {
		panic(fmt.Errorf("missing logger"))
	}

	p2p.SetConfig(cfg)
	p2p.SetLogger(log)

	// create connections
	caBridge, dbBridge, rpcBridge, geoBridge, err := connect(cfg, log)
	if err != nil {
		log.Fatal("repository init failed")
		return nil
	}

	// construct the proxy instance
	p := proxy{
		cache: caBridge,
		db:    dbBridge,
		rpc:   rpcBridge,
		geoip: geoBridge,
		log:   log,
		cfg:   cfg,

		// get the map of governance contracts
		govContracts: governanceContractsMap(cfg.Governance),

		// keep reference to the SOL compiler
		solCompiler: cfg.Compiler.DefaultSolCompilerPath,
	}

	// return the proxy
	return &p
}

// governanceContractsMap creates map of governance contracts keyed
// by the contract address.
func governanceContractsMap(cfg config.Governance) map[string]config.GovernanceContract {
	// prep the result set
	res := make(map[string]config.GovernanceContract)

	// collect all the configured governance contracts into the map
	for _, gv := range cfg.Contracts {
		res[gv.Address.String()] = gv
	}
	return res
}

// connect opens connections to the external sources we need.
func connect(cfg *config.Config, log logger.Logger) (*cache.MemBridge, *db.MongoDbBridge, *rpc.FtmBridge, *geoip.Bridge, error) {
	// create new in-memory cache bridge
	caBridge, err := cache.New(cfg, log)
	if err != nil {
		log.Criticalf("can not create in-memory cache bridge, %s", err.Error())
		return nil, nil, nil, nil, err
	}

	// create new database connection bridge
	dbBridge, err := db.New(cfg, log)
	if err != nil {
		log.Criticalf("can not connect backend persistent storage, %s", err.Error())
		return nil, nil, nil, nil, err
	}

	// create new Opera RPC bridge
	rpcBridge, err := rpc.New(cfg, log)
	if err != nil {
		log.Criticalf("can not connect Opera RPC interface, %s", err.Error())
		return nil, nil, nil, nil, err
	}

	// open GeoIP service bridge
	geo, err := geoip.New(&cfg.OperaNetwork, log)
	if err != nil {
		log.Criticalf("GeoIP service failed; %s", err.Error())
		return nil, nil, nil, nil, err
	}

	return caBridge, dbBridge, rpcBridge, geo, nil
}

// Close with close all connections and clean up the pending work for graceful termination.
func (p *proxy) Close() {
	// inform about actions
	p.log.Notice("repository is closing")

	// close connections
	p.geoip.Close()
	p.rpc.Close()
	p.db.Close()

	// inform about actions
	p.log.Notice("repository done")
}
