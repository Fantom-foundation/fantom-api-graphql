// Package config handles API server configuration binding and loading.
package config

// default configuration elements and keys
const (
	configFileName = "apiserver"

	// configuration options
	keyAppName           = "app_name"
	keyConfigFilePath    = "cfg"
	keyBindAddress       = "server.bind"
	keyDomainAddress     = "server.domain"
	keyApiPeers          = "server.peers"
	keyApiStateOrigin    = "server.origin"
	keyCorsAllowOrigins  = "server.cors_origins"
	keyLoggingLevel      = "log.level"
	keyLoggingFormat     = "log.format"
	keyLachesisUrl       = "lachesis.url"
	keyMongoUrl          = "db.url"
	keyMongoDatabase     = "db.db"
	keyCacheEvictionTime = "cache.eviction"
	keySolCompilerPath   = "compiler.sol"
	keyVotingSources     = "voting.sources"

	// PoS staking configuration
	keyStakingSfcContract       = "staking.sfc"
	keyStakingTokenizerContract = "staking.tokenizer"
	keyStakingERC20Token        = "staking.token"

	// defi related configs
	keyDefiFMintAddressProvider = "defi.fmint.address_provider"
	keyDefiUniswapCore          = "defi.uniswap.core"
	keyDefiUniswapRouter        = "defi.uniswap.router"
)
