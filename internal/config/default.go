package config

import (
	"github.com/spf13/viper"
	"time"
)

// Default values of configuration options
const (
	// defServerBind holds default API server binding address
	defServerBind = "localhost:16761"

	// defServerDomain holds default API server domain address
	defServerDomain = "localhost:16761"

	// defLoggingLevel holds default Logging level
	// See `godoc.org/github.com/op/go-logging` for the full format specification
	// See `golang.org/pkg/time/` for time format specification
	defLoggingLevel = "INFO"

	// defLoggingFormat holds default format of the Logger output
	defLoggingFormat = "%{color} %{shortfunc} > %{level:.4s} %{id:03x}%{color:reset} %{message}"

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
)

// defCorsAllowOrigins holds CORS default allowed origins.
var defCorsAllowOrigins = []string{"*"}

// default list of API peers
var defVotingSources = make([]string, 0)

// applyDefaults sets default values for configuration options.
func applyDefaults(cfg *viper.Viper) {
	// set simple details
	cfg.SetDefault(keyBindAddress, defServerBind)
	cfg.SetDefault(keyDomainAddress, defServerDomain)
	cfg.SetDefault(keyLoggingLevel, defLoggingLevel)
	cfg.SetDefault(keyLoggingFormat, defLoggingFormat)
	cfg.SetDefault(keyLachesisUrl, defLachesisUrl)
	cfg.SetDefault(keyMongoUrl, defMongoUrl)
	cfg.SetDefault(keyMongoDatabase, defMongoDatabase)
	cfg.SetDefault(keyCacheEvictionTime, defCacheEvictionTime)
	cfg.SetDefault(keySolCompilerPath, defSolCompilerPath)
	cfg.SetDefault(keyApiStateOrigin, defApiStateOrigin)

	// no voting sources by default
	cfg.SetDefault(keyVotingSources, defVotingSources)

	// cors
	cfg.SetDefault(keyCorsAllowOrigins, defCorsAllowOrigins)
}
