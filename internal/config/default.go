package config

import (
	"github.com/spf13/viper"
	"time"
)

// Default values of configuration options
const (
	// this defines default application name
	defApplicationName = "GraphQL API Server (unspecified)"

	// this defines empty address
	defNoAddress = "0x0000000000000000000000000000000000000000"

	// defServerBind holds default API server binding address
	defServerBind = "localhost:16761"

	// defServerDomain holds default API server domain address
	defServerDomain = "localhost:16761"

	// defLoggingLevel holds default Logging level
	// See `godoc.org/github.com/op/go-logging` for the full format specification
	// See `golang.org/pkg/time/` for time format specification
	defLoggingLevel = "INFO"

	// defLoggingFormat holds default format of the Logger output
	defLoggingFormat = "%{color}%{level:-8s} %{shortpkg}/%{shortfunc}%{color:reset}: %{message}"

	// defLachesisUrl holds default Lachesis connection string
	defLachesisUrl = "~/.lachesis/data/lachesis.ipc"

	// defMongoUrl holds default MongoDB connection string
	defMongoUrl = "mongodb://localhost:27017"

	// defMongoDatabase holds the default name of the API persistent database
	defMongoDatabase = "fantom"

	// defCacheEvictionTime holds default time for in-memory eviction periods
	defCacheEvictionTime = 30 * time.Minute

	// defSolCompilerPath represents the default SOL compiler path
	defSolCompilerPath = "/usr/bin/solc"

	// defApiStateOrigin represents the default origin used for API state syncing
	defApiStateOrigin = "https://localhost"

	// defDefiFMintAddressProvider represents the address of the fMintAddressProvider
	defDefiFMintAddressProvider = "0x730e27f6c52d07b1a6ab39b639b617dc566c91af"

	// defDefiFMintAddressProvider represents the address of the fMintAddressProvider
	defDefiUniswapCore = defNoAddress

	// defDefiFMintAddressProvider represents the address of the fMintAddressProvider
	defDefiUniswapRouter = defNoAddress
)

// default list of API peers
var defApiPeers = []string{"https://localhost:16761/api"}

// defCorsAllowOrigins holds CORS default allowed origins.
var defCorsAllowOrigins = []string{"*"}

// default list of API peers
var defVotingSources = make([]string, 0)

// applyDefaults sets default values for configuration options.
func applyDefaults(cfg *viper.Viper) {
	// set simple details
	cfg.SetDefault(keyAppName, defApplicationName)
	cfg.SetDefault(keyBindAddress, defServerBind)
	cfg.SetDefault(keyDomainAddress, defServerDomain)
	cfg.SetDefault(keyLoggingLevel, defLoggingLevel)
	cfg.SetDefault(keyLoggingFormat, defLoggingFormat)
	cfg.SetDefault(keyLachesisUrl, defLachesisUrl)
	cfg.SetDefault(keyMongoUrl, defMongoUrl)
	cfg.SetDefault(keyMongoDatabase, defMongoDatabase)
	cfg.SetDefault(keyCacheEvictionTime, defCacheEvictionTime)
	cfg.SetDefault(keySolCompilerPath, defSolCompilerPath)
	cfg.SetDefault(keyApiPeers, defApiPeers)
	cfg.SetDefault(keyApiStateOrigin, defApiStateOrigin)

	// no voting sources by default
	cfg.SetDefault(keyVotingSources, defVotingSources)

	// cors
	cfg.SetDefault(keyCorsAllowOrigins, defCorsAllowOrigins)

	// DeFi configuration
	cfg.SetDefault(keyDefiStakeTokenizer, defNoAddress)
	cfg.SetDefault(keyDefiFMintAddressProvider, defDefiFMintAddressProvider)
	cfg.SetDefault(keyDefiUniswapCore, defDefiUniswapCore)
	cfg.SetDefault(keyDefiUniswapRouter, defDefiUniswapRouter)
}
