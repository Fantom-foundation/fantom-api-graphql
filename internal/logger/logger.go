package logger

import (
	"fantom-api-graphql/internal/config"
	"github.com/op/go-logging"
	"os"
)

// ApiLogger defines extended logger with generic no-level logging option
type ApiLogger struct {
	logging.Logger
}

// Printf implements default non-leveled output.
// We assume the information is low in importance if passed to this function so we relay it to Debug level.
func (a ApiLogger) Printf(format string, args ...interface{}) {
	a.Debugf(format, args...)
}

// New provides pre-configured Logger with stderr output and leveled filtering.
// Modules are not supported at the moment, but may be added in the future to make the logging setup more granular.
func New(cfg *config.Config) Logger {
	// Prep the backend for exporting the log records
	// @todo Allow app to define different logging backend by configuration means.
	backend := logging.NewLogBackend(os.Stderr, "", 0)

	// Parse log format from configuration and apply it to the backend
	format := logging.MustStringFormatter(cfg.Log.Format)
	fmtBackend := logging.NewBackendFormatter(backend, format)

	// Parse and apply the configured level on which the recording will be emitted
	level, err := logging.LogLevel(cfg.Log.Level)
	if err != nil {
		level = logging.INFO
	}
	lvlBackend := logging.AddModuleLevel(fmtBackend)
	lvlBackend.SetLevel(level, "")

	// assign the backend and return the new logger
	logging.SetBackend(lvlBackend)
	l := logging.MustGetLogger(cfg.AppName)

	return &ApiLogger{*l}
}
