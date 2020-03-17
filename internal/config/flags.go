package config

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Flags usage descriptions
const (
	flagBindAddressUsage   = "Network interface the API server will operate on."
	flagDomainNameUsage    = "Domain name and port the API server will operate on."
	flagLoggingLevelUsage  = "Log level, please use one of these: CRITICAL, ERROR, WARNING, NOTICE, INFO, DEBUG."
	flagLoggingFormatUsage = "Log format."
	flagLachesisUrlUsage   = "Lachesis RPC address."
	flagMongoUrlUsage      = "MongoDB connection string."
	flagCorsAllowOrigins   = "List of CORS allowed origins."
	flagCacheEvictionTime  = "Time after which entry can be evicted from in-memory cache."
)

// bindFlags defines command line flags set and binds it with the config.
func bindFlags(cfg *viper.Viper) error {
	// create PFlags
	pflag.String(keyBindAddress, defServerBind, flagBindAddressUsage)
	pflag.String(keyDomainAddress, defServerDomain, flagDomainNameUsage)
	pflag.String(keyLoggingLevel, defLoggingLevel, flagLoggingLevelUsage)
	pflag.String(keyLoggingFormat, defLoggingFormat, flagLoggingFormatUsage)
	pflag.String(keyLachesisUrl, defLachesisUrl, flagLachesisUrlUsage)
	pflag.String(keyMongoUrl, defMongoUrl, flagMongoUrlUsage)
	pflag.Duration(keyCorsAllowOrigins, defCacheEvictionTime, flagCacheEvictionTime)
	pflag.StringSlice(keyCacheEvictionTime, defCorsAllowOrigins, flagCorsAllowOrigins)

	// parse the options
	pflag.Parse()

	// bind with Viper config
	return cfg.BindPFlags(pflag.CommandLine)
}
