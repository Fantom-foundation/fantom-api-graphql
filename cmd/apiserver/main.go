package main

import (
	"fantom-api-graphql/cmd/apiserver/build"
	"fantom-api-graphql/internal/config"
	"fantom-api-graphql/internal/graphql/resolvers"
	"fantom-api-graphql/internal/handlers"
	"fantom-api-graphql/internal/logger"
	"fantom-api-graphql/internal/repository"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// main initializes the API server and starts it when ready.
func main() {
	// make sure to capture version request
	versionRequest := flag.Bool("v", false, "get the application version")

	// get the configuration to prepare the server
	cfg, err := config.Load()
	if nil != err {
		log.Fatal(err)
	}

	// print the version information
	build.PrintVersion(cfg)
	if *versionRequest {
		return
	}

	// make logger
	lg := logger.New(cfg)

	// create repository for data exchange with the blockchain full node and local persistent storage
	repo, err := repository.New(cfg, lg)
	if err != nil {
		log.Fatal(err)
	}

	// capture termination signals and start listening
	setupSignals(repo, resolver(cfg, lg, repo), lg)
	log.Fatal(http.ListenAndServe(cfg.Server.BindAddress, nil))
}

// initRepository initializes the repository
func resolver(cfg *config.Config, log logger.Logger, repo repository.Repository) resolvers.ApiResolver {
	// create root resolver
	rs := resolvers.New(cfg, log, repo)
	log.Notice("initialized, going live")

	// setup GraphQL API handler
	h := handlers.Api(cfg, log, rs)
	http.Handle("/api", h)
	http.Handle("/graphql", h)

	// handle GraphiQL interface
	http.Handle("/graphi", handlers.GraphiHandler(cfg.Server.DomainAddress, log))

	// log the server opening info
	log.Infof("welcome to Fantom GraphQL API server network interface.")
	log.Infof("listening for requests on [%s]", cfg.Server.BindAddress)

	return rs
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
		log.Notice("server is terminating")
		repo.Close()
		rs.Close()

		log.Info("done")
		os.Exit(0)
	}()
}
