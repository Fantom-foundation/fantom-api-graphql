package config

import (
	"flag"
	"github.com/spf13/viper"
	"log"
)

// Load provides a loaded configuration for Fantom API server.
func Load() (*Config, error) {
	// Get the config reader
	cfg := reader()

	// set default values
	applyDefaults(cfg)

	// Try to read the file
	if err := cfg.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore the error, we may not need the config file
			log.Print("configuration file not found, using default values")
		} else {
			// Config file was found but another error was produced
			log.Printf("can not read the server configuration")
			return nil, err
		}
	}

	// prep the container and try to unmarshal
	// the config file into the config structure
	var config Config
	err := cfg.Unmarshal(&config)
	if err != nil {
		log.Println("can not extract API server configuration")
		log.Println(err.Error())
		return nil, err
	}

	return &config, nil
}

// reader provides instance of the config reader.
// It accepts an explicit path to a config file if it was requested by `cfg` flag.
func reader() *viper.Viper {
	// make new Viper
	cfg := viper.New()

	// what is the expected name of the config file
	cfg.SetConfigName(configFileName)

	// where to look for common files
	cfg.AddConfigPath(defaultConfigDir())
	cfg.AddConfigPath(".")

	// Try to get an explicit configuration file path if present
	var cfgPath string
	flag.StringVar(&cfgPath, keyConfigFilePath, "", "Path to a configuration file")
	flag.Parse()

	// Any path found?
	cfg.SetConfigFile(cfgPath)

	return cfg
}
