// Package config handles API server configuration binding and loading.
package config

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

// Config defines configuration options structure for Fantom API server.
type Config struct {
	// AppName holds the name of the application
	AppName string `mapstructure:"app_name"`

	// Signature represents a signature of the API server on blockchain and discovery.
	Signature ServerSignature `mapstructure:"me"`

	// Server configuration
	Server Server `mapstructure:"server"`

	// Logger configuration
	Log Log `mapstructure:"log"`

	// Opera represents the node structure
	Opera OperaNode `mapstructure:"node"`

	// OperaNetwork defines peer to peer networking options
	OperaNetwork PeerNetworking `mapstructure:"p2p"`

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

	// Governance configuration
	Governance Governance `mapstructure:"governance"`

	// TokenLogoFilePath contains the path to JSON file with the map
	// of known ERC20 tokens to their logo URLs.
	// The file will be loaded on configuration loading.
	TokenLogoFilePath string `mapstructure:"erc20_tokens_file"`

	// TokenLogo is a list of known ERC20 tokens
	// mapped to URL addresses of their logos.
	TokenLogo map[common.Address]string

	// ReScanBlocks represents the number of blocks to be re-scanned.
	RepoCommand RepoCmd `mapstructure:"cmd"`
}

// RepoCmd represents a repository command configuration.
type RepoCmd struct {
	BlockScanReScan uint64
	RestoreStake    string
}

// Server represents the GraphQL server configuration
type Server struct {
	BindAddress     string   `mapstructure:"bind"`
	DomainAddress   string   `mapstructure:"domain"`
	Origin          string   `mapstructure:"origin"`
	Peers           []string `mapstructure:"peers"`
	CorsOrigin      []string `mapstructure:"cors_origins"`
	ReadTimeout     int64    `mapstructure:"read_timeout"`
	WriteTimeout    int64    `mapstructure:"write_timeout"`
	IdleTimeout     int64    `mapstructure:"idle_timeout"`
	HeaderTimeout   int64    `mapstructure:"header_timeout"`
	ResolverTimeout int64    `mapstructure:"resolver_timeout"`
}

// ServerSignature represents the signature used by this server
// on sending requests to the blockchain, especially signed requests.
type ServerSignature struct {
	Address    common.Address    `mapstructure:"address"`
	PrivateKey *ecdsa.PrivateKey `mapstructure:"pkey"`
}

// Log represents the logger configuration
type Log struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
}

// OperaNode represents the Opera network node access configuration
type OperaNode struct {
	ApiNodeUrl string `mapstructure:"url"`
}

// PeerNetworking defines configuration for Opera p2p protocol.
type PeerNetworking struct {
	DiscoveryUDP   string        `mapstructure:"bind_udp"`
	GeoIPPath      string        `mapstructure:"geo_db"`
	BootstrapNodes []*enode.Node `mapstructure:"bootstrap"`
}

// Database represents the database access configuration.
type Database struct {
	Url    string `mapstructure:"url"`
	DbName string `mapstructure:"db"`
}

// Cache represents the cache sub-system configuration.
type Cache struct {
	Eviction time.Duration `mapstructure:"eviction"`
	MaxSize  int           `mapstructure:"size"`
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
	FLend        DeFiFLend   `mapstructure:"flend"`
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

// DeFiFLend represents the fLend DeFi module configuration.
type DeFiFLend struct {
	LendingPool common.Address `mapstructure:"lending_pool"`
}
