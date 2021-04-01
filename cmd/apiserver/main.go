package main

import (
	"fantom-api-graphql/cmd/apiserver/build"
	"fantom-api-graphql/internal/config"
	"fantom-api-graphql/internal/graphql/resolvers"
	"fantom-api-graphql/internal/handlers"
	"flag"
	"net/http"
	"time"

	/* "fantom-api-graphql/internal/handlers" */
	"fantom-api-graphql/internal/logger"
	"fantom-api-graphql/internal/repository"
	"log"
	/* "net/http" */
	"os"
	"os/signal"
	"syscall"
)

// init initializes the package and sets some important app-wide options
func init() {
	/*
	   Safety net for 'too many open files' issue on legacy code.
	   Set a sane timeout duration for the http.DefaultClient, to ensure idle connections are terminated.
	   Reference: https://stackoverflow.com/questions/37454236/net-http-server-too-many-open-files-error
	*/
	http.DefaultClient.Timeout = time.Second * 10
}

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
	repository.SetConfig(cfg)
	repository.SetLogger(lg)

	// start the HTTP server
	startHttpServer(cfg, lg)
}

// startHttpServer starts the HTTP server and begins resolving incoming requests.
func startHttpServer(cfg *config.Config, log logger.Logger) {
	// create request MUXer
	srvMux := new(http.ServeMux)

	// create a HTTP server to handle our requests
	srv := &http.Server{
		Addr:              cfg.Server.BindAddress,
		ReadTimeout:       time.Second * time.Duration(cfg.Server.ReadTimeout),
		WriteTimeout:      time.Second * time.Duration(cfg.Server.WriteTimeout),
		IdleTimeout:       time.Second * time.Duration(cfg.Server.IdleTimeout),
		ReadHeaderTimeout: time.Second * time.Duration(cfg.Server.HeaderTimeout),
		Handler:           srvMux,
	}

	// capture termination signals and start listening
	setupSignals(setupHandlers(srvMux, cfg, log), log)

	// log the server opening info
	log.Infof("welcome to Fantom GraphQL API server network interface.")
	log.Infof("listening for requests on [%s]", cfg.Server.BindAddress)

	// start the server
	log.Fatal(srv.ListenAndServe())
}

// setupHandlers initializes an array of handlers for our HTTP API end-points.
func setupHandlers(mux *http.ServeMux, cfg *config.Config, log logger.Logger) resolvers.ApiResolver {
	// create root resolver
	rs := resolvers.New(cfg, log)
	log.Notice("initialized, going live")

	// setup GraphQL API handler
	h := http.TimeoutHandler(handlers.Api(cfg, log, rs), time.Second*time.Duration(cfg.Server.ResolverTimeout), "Service timeout.")
	mux.Handle("/api", h)
	mux.Handle("/graphql", h)

	// handle GraphiQL interface
	mux.Handle("/graphi", handlers.GraphiHandler(cfg.Server.DomainAddress, log))

	return rs
}

// setupSignals creates a system signal listener and handles graceful termination upon receiving one.
func setupSignals(rs resolvers.ApiResolver, log logger.Logger) {
	// log what we do
	log.Info("os signals captured")

	// make the signal consumer
	ts := make(chan os.Signal, 1)
	signal.Notify(ts, syscall.SIGINT, syscall.SIGTERM)

	// start monitoring
	go func() {
		// wait for the signal
		<-ts

		// log nad close repository
		log.Notice("server is terminating")
		if repo := repository.R(); repo != nil {
			repo.Close()
		}

		// close resolvers
		rs.Close()

		// we are done
		log.Info("done")
		os.Exit(0)
	}()
}
