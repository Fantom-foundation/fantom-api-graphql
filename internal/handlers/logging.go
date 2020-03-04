package handlers

import (
	flogger "fantom-api-graphql/internal/logger"
	"net/http"
)

// LoggingHandler defines HTTP handler middleware for logging incoming communication through provided Logger.
type LoggingHandler struct {
	logger  flogger.Logger
	handler http.Handler
}

// ServeHTTP handles incoming request by creating a log record with predefined request details
// and passing it to the next handler in the chain.
func (h *LoggingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// We log incoming requests on Debug level since in production the actual incoming traffic is not very important.
	h.logger.Debugf("[%s <- %s] %s %s (%s)", r.Proto, r.RemoteAddr, r.Method, r.URL, r.UserAgent())

	// Pass request down the chain
	h.handler.ServeHTTP(w, r)
}
