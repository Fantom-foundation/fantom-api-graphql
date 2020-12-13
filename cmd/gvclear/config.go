package main

import "flag"

// gvConf represents a configuration
// of the Governance Clearing app.
type gvConf struct {
	isVersionQuery bool
	nodeRpcUrl     string
	govAddress     string
	sigKeyFile     string
	sigPasswd      string
	batchLength    uint64
}

// config reads the Governance Clear app command parameters
// and builds the configuration used by the app to process Governance tasks.
func config() *gvConf {
	// prep the config
	var cfg gvConf

	// configure flag readers
	flag.BoolVar(&cfg.isVersionQuery, "version", false, "get the application version")
	flag.StringVar(&cfg.nodeRpcUrl, "rpc", "lachesis.ipc", "URL of the Opera node RPC/IPC interface")
	flag.StringVar(&cfg.govAddress, "gov", "0x0000000000000000000000000000000000000000", "governance contract address")
	flag.StringVar(&cfg.sigKeyFile, "key", "keystore.json", "path to the account key file")
	flag.StringVar(&cfg.sigPasswd, "pwd", "", "account key file password")
	flag.Uint64Var(&cfg.batchLength, "len", 4, "size of a processing batch")

	// parse the flags
	flag.Parse()
	return &cfg
}
