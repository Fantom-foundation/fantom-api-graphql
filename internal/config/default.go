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
	defLoggingFormat = "%{color}%{time:2006-01-02 15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}"

	// defLachesisUrl holds default Lachesis connection string
	defLachesisUrl = "~/.lachesis/data/lachesis.ipc"

	// defMongoUrl holds default MongoDB connection string
	defMongoUrl = "mongodb://localhost:27017"

	// defCacheEvictionTime holds default time for in-memory eviction periods
	defCacheEvictionTime = 30 * time.Minute
)

// defCorsAllowOrigins holds CORS default allowed origins.
var defCorsAllowOrigins = []string{"*"}

// applyDefaults sets default values for configuration options.
func applyDefaults(cfg *viper.Viper) {
	// set simple details
	cfg.SetDefault(keyBindAddress, defServerBind)
	cfg.SetDefault(keyDomainAddress, defServerDomain)
	cfg.SetDefault(keyLoggingLevel, defLoggingLevel)
	cfg.SetDefault(keyLoggingFormat, defLoggingFormat)
	cfg.SetDefault(keyLachesisUrl, defLachesisUrl)
	cfg.SetDefault(keyMongoUrl, defMongoUrl)
	cfg.SetDefault(keyCacheEvictionTime, defCacheEvictionTime)

	// cors
	cfg.SetDefault(keyCorsAllowOrigins, defCorsAllowOrigins)
}
