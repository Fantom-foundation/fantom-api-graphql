// Package main implements the API server entry point.
package main

import (
	"fantom-api-graphql/cmd/apiserver/build"
	"fantom-api-graphql/internal/config"
	"fantom-api-graphql/internal/logger"
	"fantom-api-graphql/internal/repository"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// apiServer implements the API server application
type apiServer struct {
	cfg          *config.Config
	log          logger.Logger
	srv          *http.Server
	isVersionReq bool
}

// init initializes the API server
func (app *apiServer) init() {
	// make sure to capture version request and rescan depth
	flag.BoolVar(&app.isVersionReq, "v", false, "get the application version")

	// get the configuration including parsing the calling flags
	var err error
	app.cfg, err = config.Load()
	if nil != err {
		log.Fatal(err)
		return
	}

	// configure logger based on the configuration
	app.log = logger.New(app.cfg)

	// make sure to pass logger and config to internals
	repository.SetConfig(app.cfg)
	repository.SetLogger(app.log)

	// make the HTTP server
	app.makeHttpServer()
}

// run executes the API server function.
func (app *apiServer) run() {
	// print the app version and exit if this is the only thing requested
	build.PrintVersion(app.cfg)
	if app.isVersionReq {
		return
	}

	// make sure to capture terminate signals
	app.observeSignals()

	// start responding to requests
	app.log.Infof("welcome to Fantom comparator")

	// terminate the app
	app.terminate()
}

// makeHttpServer creates and configures the HTTP server to be used to serve incoming requests
func (app *apiServer) makeHttpServer() {
	// create request MUXer
	srvMux := new(http.ServeMux)

	// create HTTP server to handle our requests
	app.srv = &http.Server{
		Addr:              app.cfg.Server.BindAddress,
		ReadTimeout:       time.Second * time.Duration(app.cfg.Server.ReadTimeout),
		WriteTimeout:      time.Second * time.Duration(app.cfg.Server.WriteTimeout),
		IdleTimeout:       time.Second * time.Duration(app.cfg.Server.IdleTimeout),
		ReadHeaderTimeout: time.Second * time.Duration(app.cfg.Server.HeaderTimeout),
		Handler:           srvMux,
	}
}

// observeSignals setups terminate signals observation.
func (app *apiServer) observeSignals() {
	// log what we do
	app.log.Info("os signals captured")

	// make the signal consumer
	ts := make(chan os.Signal, 1)
	signal.Notify(ts, syscall.SIGINT, syscall.SIGTERM)

	// start monitoring
	go func() {
		// wait for the signal
		<-ts

		// terminate HTTP responder
		app.log.Notice("closing HTTP server")
		if err := app.srv.Close(); err != nil {
			app.log.Errorf("could not terminate HTTP listener")
			os.Exit(0)
		}
	}()
}

// terminate modules of the API server.
func (app *apiServer) terminate() {
	// terminate connections to DB, blockchain, etc.
	app.log.Notice("closing repository")
	if repo := repository.R(); repo != nil {
		repo.Close()
	}
}
