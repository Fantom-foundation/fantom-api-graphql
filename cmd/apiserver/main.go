package main

import (
	"fantom-api-graphql/internal/config"
	"fantom-api-graphql/internal/handlers"
	"fantom-api-graphql/internal/logger"
	"log"
	"net/http"
)

// main initializes the API server and starts it when ready.
func main() {
	// get the configuration to prepare the server
	cfg, err := config.Load()
	if nil != err {
		log.Fatalf("can not read application configuration")
	}

	// make logger
	lg := logger.New(cfg)

	// create repository for data exchange with the blockchain full node and local persistent storage
	r, err := repository.New(cfg, lg)
	if err != nil {
		log.Fatalf("can not create application repository")
	}

	// setup GraphQL API handler
	http.Handle("/api", handlers.Api(cfg, lg, r))

	// log the server opening info
	lg.Infof("Welcome to Fantom GraphQL API server network interface.")
	lg.Infof("Listening for requests on [%s]", cfg.BindAddress)

	// start listening
	log.Fatal(http.ListenAndServe(cfg.BindAddress, nil))
}
