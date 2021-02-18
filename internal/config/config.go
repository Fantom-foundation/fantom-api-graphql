// Package config handles API server configuration binding and loading.
package config

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"time"
)

// Config defines configuration options structure for Fantom API server.
type Config struct {
	// AppName holds the name of the application
	AppName string `mapstructure:"app_name"`

	// MySignature represents a signature of the server on block chain.
	MySignature ServerSignature `mapstructure:"me"`

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

	// Repository configuration
	Repository Repository `mapstructure:"repository"`

	// Staking configuration
	Staking Staking `mapstructure:"staking"`

	// DeFi configuration
	DeFi DeFi `mapstructure:"defi"`

	// Voting configuration
	Voting Voting `mapstructure:"voting"`

	// Governance configuration
	Governance Governance `mapstructure:"governance"`

	// KnownTokens is a list of known ERC20 tokens
	// mapped to URL addresses of their logos.
	TokenLogo map[common.Address]string `mapstructure:"erc20_logos"`
}

// Server represents the GraphQL server configuration
type Server struct {
	BindAddress   string   `mapstructure:"bind"`
	DomainAddress string   `mapstructure:"domain"`
	Origin        string   `mapstructure:"origin"`
	Peers         []string `mapstructure:"peers"`
	CorsOrigin    []string `mapstructure:"cors_origins"`
}

// ServerSignature represents the signature used by this server
// on sending requests to the block chain, especially signed requests.
type ServerSignature struct {
	Address    common.Address   `mapstructure:"address"`
	PrivateKey ecdsa.PrivateKey `mapstructure:"pkey"`
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
	CompilerTempPath       string `mapstructure:"temp"`
	DefaultSolCompilerPath string `mapstructure:"sol"`
}

// Repository represents the repository configuration.
type Repository struct {
	MonitorStakers bool `mapstructure:"stakers"`
}

// Staking represents the PoS Staking module configuration.
type Staking struct {
	SFCContract         common.Address `mapstructure:"sfc"`
	StiContract         common.Address `mapstructure:"sti"`
	TokenizerContract   common.Address `mapstructure:"tokenizer"`
	TokenizedStakeToken common.Address `mapstructure:"token"`
}

// DeFi represents the DeFi and financial contracts configuration.
type DeFi struct {
	FMint        DeFiFMint   `mapstructure:"fmint"`
	Uniswap      DeFiUniswap `mapstructure:"uniswap"`
	PriceSymbols []string    `mapstructure:"symbols"`
}

// DeFiFMint represents the fMint DeFi module configuration.
type DeFiFMint struct {
	AddressProvider common.Address `mapstructure:"address_provider"`
}

// DeFiUniswap represents the Uniswap protocol DeFi module configuration.
type DeFiUniswap struct {
	Core           common.Address   `mapstructure:"core"`
	Router         common.Address   `mapstructure:"router"`
	PairsWhiteList []common.Address `mapstructure:"whitelist"`
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
	Address    common.Address `mapstructure:"address"`
	Governable common.Address `mapstructure:"governable"`
	Templates  common.Address `mapstructure:"templates"`
	Name       string         `mapstructure:"name"`
	Type       string         `mapstructure:"type"`
}
