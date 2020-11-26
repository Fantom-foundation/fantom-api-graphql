package config

import (
	"crypto/ecdsa"
	"fantom-api-graphql/internal/types"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"log"
	"reflect"
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
	err := cfg.Unmarshal(&config, setupConfigUnmarshaler)
	if err != nil {
		log.Println("can not extract API server configuration")
		log.Println(err.Error())
		return nil, err
	}

	return &config, nil
}

// setupConfigUnmarshaler configures the Config loader to properly unmarshal
// special types we use for the API server
func setupConfigUnmarshaler(cfg *mapstructure.DecoderConfig) {
	// add the decoders missing here
	cfg.DecodeHook = mapstructure.ComposeDecodeHookFunc(
		StringToAddressHookFunc(),
		StringToPrivateKeyHookFunc(),
		cfg.DecodeHook)
}

// StringToPrivateKeyHookFunc returns a DecodeHookFunc that converts
// strings to ecdsa.PrivateKey type on config loading.
func StringToPrivateKeyHookFunc() mapstructure.DecodeHookFuncType {
	// return the decoder function
	return func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
		// make sure the input is expected String
		if f.Kind() != reflect.String {
			return data, nil
		}

		// make sure the output is expected to be ecdsa.PrivateKey
		if t != reflect.TypeOf(ecdsa.PrivateKey{}) {
			return data, nil
		}

		// convert input to string
		raw := data.(string)
		if raw == "" {
			return nil, fmt.Errorf("empty private key content")
		}

		// try to decode the key
		key, err := crypto.ToECDSA(common.FromHex(raw))
		if err != nil {
			return nil, err
		}
		return *key, nil
	}
}

// StringToAddressHookFunc returns a DecodeHookFunc that converts
// strings to common.Address type on config loading.
func StringToAddressHookFunc() mapstructure.DecodeHookFuncType {
	// return the decoder function
	return func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
		// make sure the input is expected String
		if f.Kind() != reflect.String {
			return data, nil
		}

		// make sure the output is expected common.Address
		if t == reflect.TypeOf(common.Address{}) {
			raw := data.(string)
			if raw == "" {
				return common.HexToAddress(EmptyAddress), nil
			}
			return stringToCommonAddress(raw)
		}

		// typed address is expected here?
		if t == reflect.TypeOf(types.Address{}) {
			raw := data.(string)
			if raw == "" {
				return stringToAddress(EmptyAddress)
			}
			return stringToAddress(raw)
		}

		// anything found else?
		return data, nil
	}
}

// stringToCommonAddress converts the given String to common Address.
func stringToCommonAddress(str string) (interface{}, error) {
	return common.HexToAddress(str), nil
}

// stringToAddress converts the given String to typed Address.
func stringToAddress(str string) (interface{}, error) {
	return types.Address(common.HexToAddress(str)), nil
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
