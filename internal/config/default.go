package config

import "github.com/spf13/viper"

// Default values of configuration options
const (
	// API server binding address
	defServerBind = "localhost:8008"

	// Logging defaults
	// See `godoc.org/github.com/op/go-logging` for the full format specification
	// See `golang.org/pkg/time/` for time format specification
	defLoggingLevel  = "INFO"
	defLoggingFormat = "%{color}%{time:2006-01-02 15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}"

	// Lachesis connection string
	defLachesisUrl = "~/.lachesis/data/lachesis.ipc"

	// MongoDB connection string
	defMongoUrl = "mongodb://localhost:27017"
)

// applyDefaults sets default values for configuration options.
func applyDefaults(cfg *viper.Viper) {
	cfg.SetDefault(keyBindAddress, defServerBind)
	cfg.SetDefault(keyLoggingLevel, defLoggingLevel)
	cfg.SetDefault(keyLoggingFormat, defLoggingFormat)
	cfg.SetDefault(keyLachesisUrl, defLachesisUrl)
	cfg.SetDefault(keyMongoUrl, defMongoUrl)
}
