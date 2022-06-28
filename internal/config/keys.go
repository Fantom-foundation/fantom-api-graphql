// Package config handles API server configuration binding and loading.
package config

// default configuration elements and keys
const (
	configFileName = "apiserver"

	// configuration options
	keyAppName                  = "app_name"
	keyConfigFilePath           = "cfg"
	keyConfigCmdBlockScanStart  = "cmd.blk_from"
	keyConfigCmdBlockScanEnd    = "cmd.blk_to"
	keyConfigCmdBlockScanReScan = "cmd.rescan"
	keyConfigCmdRestoreStake    = "cmd.fix_stake"

	// server related keys
	keyBindAddress      = "server.bind"
	keyDomainAddress    = "server.domain"
	keyApiPeers         = "server.peers"
	keyApiStateOrigin   = "server.origin"
	keyCorsAllowOrigins = "server.cors_origins"

	// server time out related keys
	keyTimeoutRead     = "server.read_timeout"
	keyTimeoutWrite    = "server.write_timeout"
	keyTimeoutIdle     = "server.idle_timeout"
	keyTimeoutHeader   = "server.header_timeout"
	keyTimeoutResolver = "server.resolver_timeout"

	// API server signature related keys
	keySignatureAddress    = "me.address"
	keySignaturePrivateKey = "me.pkey"

	// logging related options
	keyLoggingLevel  = "log.level"
	keyLoggingFormat = "log.format"

	// node connection related options
	keyLachesisUrl = "lachesis.url"

	// off-chain database related options
	keyMongoUrl      = "db.url"
	keyMongoDatabase = "db.db"

	// cache related options
	keyCacheEvictionTime = "cache.eviction"
	keyCacheMaxSize      = "cache.size"

	// contract validation related
	keySolCompilerPath = "compiler.sol"

	// utility options
	keyVotingSources         = "voting.sources"
	keyErc20TokenMapFilePath = "erc20_tokens_file"
	keyErc20Logos            = "erc20_logos"

	// PoS staking configuration
	keyStakingSfcContract       = "staking.sfc"
	keyStakingStiContract       = "staking.sti"
	keyStakingTokenizerContract = "staking.tokenizer"
	keyStakingERC20Token        = "staking.token"

	// defi related configs
	keyDefiFMintAddressProvider = "defi.fmint.address_provider"
	keyDefiUniswapCore          = "defi.uniswap.core"
	keyDefiUniswapRouter        = "defi.uniswap.router"

	keyP2PBindUDP = "p2p.bind_udp"
)
