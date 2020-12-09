package gvclear

import "flag"

// gvConf represents a configuration
// of the Governance Clearing app.
type gvConf struct {
	isVersionQuery bool
	nodeRpcUrl     string
	govAddress     string
}

// config reads the Governance Clear app command parameters
// and builds the configuration used by the app to process Governance tasks.
func config() *gvConf {
	// prep the config
	var cfg gvConf

	// configure flag readers
	flag.BoolVar(&cfg.isVersionQuery, "version", false, "get the application version")
	flag.StringVar(&cfg.nodeRpcUrl, "rpc", "lachesis.ipc", "URL of the Opera node RPC/IPC interface")
	flag.StringVar(&cfg.govAddress, "gov", "", "governance contract address")

	// parse the flags
	flag.Parse()
	return &cfg
}
