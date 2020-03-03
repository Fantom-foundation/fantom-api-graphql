package logger

import (
	"fantom-api-graphql/internal/config"
	"github.com/op/go-logging"
	"os"
)

// New provides pre-configured Logger with stderr output and leveled filtering.
// Modules are not supported at the moment, but may be added in the future to make the logging setup more granular.
func New(cfg *config.Config) Logger {
	// Prep the backend for exporting the log records
	// @todo Allow app to define different logging backend by configuration means.
	backend := logging.NewLogBackend(os.Stderr, "", 0)

	// Parse log format from configuration and apply it to the backend
	format := logging.MustStringFormatter(cfg.LoggingFormat)
	fmtBackend := logging.NewBackendFormatter(backend, format)

	// Parse and apply the configured level on which the recording will be emitted
	level, err := logging.LogLevel(cfg.LoggingLevel)
	if err != nil {
		level = logging.INFO
	}
	lvlBackend := logging.AddModuleLevel(fmtBackend)
	lvlBackend.SetLevel(level, "")

	// assign the backend and return the new logger
	logging.SetBackend(lvlBackend)
	return logging.MustGetLogger(cfg.AppName)
}
