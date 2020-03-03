/*
Package config handles API server configuration binding and loading.

The config can be read from an environment (twelve-factor methodology),
from a command flags, and from a config file. We strongly recommend

*/
package config

import (
	"flag"
	"github.com/spf13/viper"
	"log"
)

// default configuration elements and keys
const (
	appName        = "FantomAPI"
	configFileName = "apiserver"

	// configuration options
	keyConfigFilePath = "cfg"
	keyBindAddress    = "server.bind"
	keyLoggingLevel   = "log.level"
	keyLoggingFormat  = "log.format"
	keyLachesisUrl    = "lachesis.url"
	keyMongoUrl       = "mongo.url"
)

// Config defines configuration options structure for Fantom API server.
type Config struct {
	// AppName holds the name of the application
	AppName string

	// BindAddress holds the API server network binding address
	BindAddress string

	// LoggingLevel and LoggingFormat hold configuration for the API server logger
	LoggingLevel  string
	LoggingFormat string

	// LachesisUrl holds address of the Lachesis node we want to communicate with
	LachesisUrl string

	// MongoUrl holds address of the MongoDB we use for persistent storage
	MongoUrl string
}

// Load provides a loaded configuration for Fantom API server.
func Load() (*Config, error) {
	// Get the config reader
	cfg := getReader()

	// Pre-configure the reader by applying different type of bonds
	if err := applyBonds(cfg); nil != err {
		return nil, err
	}

	// Try to read the file
	if err := cfg.ReadInConfig(); nil != err {
		log.Printf("can not read the server configuration")
		return nil, err
	}

	// Build and return the config structure
	return &Config{
		AppName:       appName,
		BindAddress:   cfg.GetString(keyBindAddress),
		LoggingLevel:  cfg.GetString(keyLoggingLevel),
		LoggingFormat: cfg.GetString(keyLoggingFormat),
		LachesisUrl:   cfg.GetString(keyLachesisUrl),
		MongoUrl:      cfg.GetString(keyMongoUrl),
	}, nil
}

// getConfig provides instance of the config reader.
// It accepts an explicit path to a config file if it was requested by `cfg` flag.
func getReader() *viper.Viper {
	// make new Viper
	cfg := viper.New()

	// what is the expected name of the config file
	cfg.SetConfigName(configFileName)

	// where to look for common files
	cfg.AddConfigPath(defaultConfigDir())

	// Try to get an explicit configuration file path if present
	var cfgPath string
	flag.StringVar(&cfgPath, keyConfigFilePath, "", "Path to a configuration file")
	flag.Parse()

	// Any path found?
	cfg.SetConfigFile(cfgPath)

	return cfg
}

// setupConfig prepares config reader by applying default values and bonds
func applyBonds(cfg *viper.Viper) error {
	// set default values
	applyDefaults(cfg)

	// bind environment vars; we want to be able to configure the app deployment in twelve-factor sense
	if err := bindEnv(cfg); nil != err {
		log.Printf("can not bind configuration to environment")
		return err
	}

	// bind command line params
	if err := bindFlags(cfg); nil != err {
		log.Printf("can not bind configuration to command line flags")
		return err
	}

	return nil
}
