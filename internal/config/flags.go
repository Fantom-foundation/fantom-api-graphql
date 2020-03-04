package config

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Flags usage descriptions
const (
	flagBindAddressUsage   = "Network interface the API server will operate on"
	flagLoggingLevelUsage  = "Log level"
	flagLoggingFormatUsage = "Log format"
	flagLachesisUrlUsage   = "Lachesis RPC address"
	flagMongoUrlUsage      = "MongoDB connection string"
	flagCorsAllowOrigins   = "List of CORS allowed origins"
)

// bindFlags defines command line flags set and binds it with the config.
func bindFlags(cfg *viper.Viper) error {
	// create PFlags
	pflag.String(keyBindAddress, defServerBind, flagBindAddressUsage)
	pflag.String(keyLoggingLevel, defLoggingLevel, flagLoggingLevelUsage)
	pflag.String(keyLoggingFormat, defLoggingFormat, flagLoggingFormatUsage)
	pflag.String(keyLachesisUrl, defLachesisUrl, flagLachesisUrlUsage)
	pflag.String(keyMongoUrl, defMongoUrl, flagMongoUrlUsage)
	pflag.StringSlice(keyCorsAllowOrigins, defCorsAllowOrigins, flagCorsAllowOrigins)

	// parse the options
	pflag.Parse()

	// bind with Viper config
	return cfg.BindPFlags(pflag.CommandLine)
}
