package main

import (
	"fantom-api-graphql/cmd/apiserver/build"
	"fantom-api-graphql/internal/config"
	"fantom-api-graphql/internal/graphql/resolvers"
	"fantom-api-graphql/internal/handlers"
	"net/http"

	/* "fantom-api-graphql/internal/handlers" */
	"fantom-api-graphql/internal/logger"
	"fantom-api-graphql/internal/repository"
	"log"
	/* "net/http" */
	"os"
	"os/signal"
	"syscall"
)

// main initializes the API server and starts it when ready.
func main() {
	// get the configuration to prepare the server
	cfg, err := config.Load()
	if nil != err {
		log.Fatal(err)
	}

	// print the version information
	build.PrintVersion(cfg)

	// make logger
	lg := logger.New(cfg)

	// create repository for data exchange with the blockchain full node and local persistent storage
	repo, err := repository.New(cfg, lg)
	if err != nil {
		log.Fatal(err)
	}

	// create root resolver
	rs := resolvers.New(cfg, lg, repo)
	lg.Notice("initialized, going live")

	// capture termination signals
	setupSignals(repo, rs, lg)

	// setup GraphQL API handler
	h := handlers.Api(cfg, lg, rs)
	http.Handle("/api", h)
	http.Handle("/graphql", h)

	// handle GraphiQL interface
	http.Handle("/graphi", handlers.GraphiHandler(cfg.Server.DomainAddress, lg))

	// log the server opening info
	lg.Infof("welcome to Fantom GraphQL API server network interface.")
	lg.Infof("listening for requests on [%s]", cfg.Server.BindAddress)

	// start listening
	log.Fatal(http.ListenAndServe(cfg.Server.BindAddress, nil))
}

// setupSignals creates a system signal listener and handles graceful termination upon receiving one.
func setupSignals(repo repository.Repository, rs resolvers.ApiResolver, log logger.Logger) {
	// log what we do
	log.Info("os signals captured")

	// make the signal consumer
	ts := make(chan os.Signal, 1)
	signal.Notify(ts, syscall.SIGINT, syscall.SIGTERM)

	// start monitoring
	go func() {
		// wait for the signal
		<-ts

		// log nad close
		log.Critical("server is terminating")
		repo.Close()
		rs.Close()

		log.Info("done")
		os.Exit(0)
	}()
}
