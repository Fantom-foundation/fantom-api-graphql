// Package config handles API server configuration binding and loading.
package config

import "time"

// Config defines configuration options structure for Fantom API server.
type Config struct {
	// AppName holds the name of the application
	AppName string `mapstructure:"app_name"`

	// Server configuration
	Server Server `mapstructure:"server"`

	// Logger configuration
	Log Log `mapstructure:"log"`

	// Lachesis represents the node structure
	Lachesis Lachesis `mapstructure:"node"`

	// Database configuration
	Db Database `mapstructure:"db"`

	// Cache configuration
	Cache Cache `mapstructure:"cache"`

	// Cache configuration
	Compiler Compiler `mapstructure:"compiler"`

	// DeFi configuration
	DeFi DeFi `mapstructure:"defi"`

	// Voting configuration
	Voting Voting `mapstructure:"voting"`

	// Governance configuration
	Governance Governance `mapstructure:"governance"`
}

// Server represents the GraphQL server configuration
type Server struct {
	BindAddress   string   `mapstructure:"bind"`
	DomainAddress string   `mapstructure:"domain"`
	Origin        string   `mapstructure:"origin"`
	Peers         []string `mapstructure:"peers"`
	CorsOrigin    []string `mapstructure:"cors_origins"`
}

// Log represents the logger configuration
type Log struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
}

// Lachesis represents the Lachesis node access configuration
type Lachesis struct {
	Url string `mapstructure:"url"`
}

// Database represents the database access configuration.
type Database struct {
	Url    string `mapstructure:"url"`
	DbName string `mapstructure:"db"`
}

// Cache represents the cache sub-system configuration.
type Cache struct {
	Eviction time.Duration `mapstructure:"eviction"`
}

// Compiler represents the contract compilers configuration.
type Compiler struct {
	DefaultSolCompilerPath string `mapstructure:"sol"`
}

// DeFi represents the DeFi and financial contracts configuration.
type DeFi struct {
	FMint   DeFiFMint   `mapstructure:"fmint"`
	Uniswap DeFiUniswap `mapstructure:"uniswap"`
}

// DeFiFMint represents the fMint DeFi module configuration.
type DeFiFMint struct {
	AddressProvider string `mapstructure:"address_provider"`
}

// DeFiUniswap represents the Uniswap protocol DeFi module configuration.
type DeFiUniswap struct {
	Core   string `mapstructure:"core"`
	Router string `mapstructure:"router"`
}

// Voting represents the simple voting/ballots module configuration.
type Voting struct {
	Sources []string `mapstructure:"sources"`
}

// Governance represents the governance module configuration.
type Governance struct {
	Contracts []GovernanceContract `mapstructure:"contracts"`
}

// GovernanceContract represents a single Governance contract configuration.
type GovernanceContract struct {
	Address string `mapstructure:"address"`
	Name    string `mapstructure:"name"`
	Type    string `mapstructure:"type"`
}
