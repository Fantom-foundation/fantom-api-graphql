// Package handlers holds HTTP/WS handlers chain along with separate middleware implementations.
package handlers

import (
	"encoding/json"
	"fantom-api-graphql/internal/logger"
	"fantom-api-graphql/internal/repository"
	"net/http"
)

// GasPrice constructs and return the REST API HTTP handler for Gas Price provider.
func GasPrice(log logger.Logger) http.Handler {
	// build the handler function
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get the gas price estimation
		val, err := repository.R().GasPriceExtended()
		if err != nil {
			log.Critical("can not get gas price; %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// respond
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(val)
		if err != nil {
			log.Critical("can not encode gas price structure; %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}
