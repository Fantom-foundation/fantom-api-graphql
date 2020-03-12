package main

import (
	"fantom-api-graphql/internal/config"
	"fantom-api-graphql/internal/handlers"
	"fantom-api-graphql/internal/logger"
	"fantom-api-graphql/internal/repository"
	"log"
	"net/http"
	"os"
	"os/signal"
)

// main initializes the API server and starts it when ready.
func main() {
	// get the configuration to prepare the server
	cfg, err := config.Load()
	if nil != err {
		log.Fatal(err)
	}

	// make logger
	lg := logger.New(cfg)

	// create repository for data exchange with the blockchain full node and local persistent storage
	repo, err := repository.New(cfg, lg)
	if err != nil {
		log.Fatal(err)
	}

	// capture termination signals
	setupSignals(repo, lg)

	// setup GraphQL API handler
	http.Handle("/api", handlers.Api(cfg, lg, repo))

	// log the server opening info
	lg.Infof("welcome to Fantom GraphQL API server network interface.")
	lg.Infof("listening for requests on [%s]", cfg.BindAddress)

	// start listening
	log.Fatal(http.ListenAndServe(cfg.BindAddress, nil))
}

// setupSignals creates a system signal listener and handles graceful termination upon receiving one.
func setupSignals(repo repository.Repository, log logger.Logger) {
	ts := make(chan os.Signal, 2)
	signal.Notify(ts, os.Interrupt, os.Kill)

	// start monitoring
	go func() {
		// wait for the signal
		<-ts

		// log nad close
		log.Info("server is terminating")
		repo.Close()

		log.Info("have a great day")
		os.Exit(0)
	}()
}
