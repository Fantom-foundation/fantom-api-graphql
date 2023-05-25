// Package handlers hold an HTTP/WS handlers chain along with separate middleware implementations.
package handlers

import (
	"embed"
	"encoding/json"
	"fantom-api-graphql/internal/logger"
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"io"
	"net/http"
	"text/template"
)

//go:embed templates
var htmlTemplates embed.FS

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

// ValidatorsDownHandler provides a handler for a textual list of validators with downtime.
func ValidatorsDownHandler(log logger.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(r.Body)

		list, err := repository.R().DownValidators()
		if err != nil {
			log.Criticalf("can not get list of offline server; %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		tmp := template.Must(template.ParseFS(htmlTemplates, "templates/down.html"))

		w.Header().Set("Content-Type", "text/html")
		err = tmp.Execute(w, struct {
			Val   []types.OfflineValidator
			Count int
		}{Val: list, Count: len(list)})
		if err != nil {
			log.Criticalf("can not execute HTML templates; %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}
