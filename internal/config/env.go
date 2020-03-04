package config

import "github.com/spf13/viper"

// Expected environment variables bound to configuration options
const (
	envBindAddress      = "FANTOM_API_BIND"
	envLoggingLevel     = "FANTOM_API_LOG_LEVEL"
	envLoggingFormat    = "FANTOM_API_LOG_FORMAT"
	envLachesisUrl      = "FANTOM_API_LACHESIS_URL"
	envMongoUrl         = "FANTOM_API_MONGO_URL"
	envCorsAllowOrigins = "FANTOM_API_CORS_ORIGINS"
)

// bindEnv binds configuration options to environment variables.
// We may have actually ignored errors here since they are fired only if the key is missing,
// but it's a good practice to handle them gracefully.
func bindEnv(cfg *viper.Viper) error {
	// listening address
	if err := cfg.BindEnv(keyBindAddress, envBindAddress); nil != err {
		return err
	}

	// log level
	if err := cfg.BindEnv(keyLoggingLevel, envLoggingLevel); nil != err {
		return err
	}

	// log format
	if err := cfg.BindEnv(keyLoggingFormat, envLoggingFormat); nil != err {
		return err
	}

	// lachesis connection URL
	if err := cfg.BindEnv(keyLachesisUrl, envLachesisUrl); nil != err {
		return err
	}

	// Mongo connection URL
	if err := cfg.BindEnv(keyMongoUrl, envMongoUrl); nil != err {
		return err
	}

	// CORS allowed origins
	if err := cfg.BindEnv(keyCorsAllowOrigins, envCorsAllowOrigins); nil != err {
		return err
	}

	return nil
}
