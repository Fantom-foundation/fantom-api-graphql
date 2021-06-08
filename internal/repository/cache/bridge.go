// Package cache implements bridge to fast in-memory object cache.
package cache

import (
	"fantom-api-graphql/internal/config"
	"fantom-api-graphql/internal/logger"
	"fantom-api-graphql/internal/repository/cache/ring"
	"github.com/allegro/bigcache"
	"time"
)

// TransactionRingCacheSize represents the amount of transactions kept
// in fast in-memory ring cache for fast loading.
const TransactionRingCacheSize = 75

// BlockRingCacheSize represents the amount of blocks kept
// in fast in-memory ring cache for fast loading.
const BlockRingCacheSize = 75

// MemBridge represents BigCache abstraction layer.
type MemBridge struct {
	cache *bigcache.BigCache
	log   logger.Logger

	// ring of the most recent blocks and transactions
	blkRing *ring.Ring
	trxRing *ring.Ring
}

// New creates a new BigCache bridge.
func New(cfg *config.Config, log logger.Logger) (*MemBridge, error) {
	// create the cache
	c, err := bigcache.NewBigCache(cacheConfig(cfg, log))
	if err != nil {
		log.Critical(err)
		return nil, err
	}

	// log the event
	log.Notice("memory cache initialized")

	// make a new Bridge
	return &MemBridge{
		cache: c,
		log:   log,

		// make rings
		blkRing: ring.New(BlockRingCacheSize),
		trxRing: ring.New(TransactionRingCacheSize),
	}, nil
}

// cacheConfig constructs a configuration structure for BigCache initialization.
func cacheConfig(cfg *config.Config, log logger.Logger) bigcache.Config {
	// log the info
	log.Debugf("memory cache eviction set to %s", cfg.Cache.Eviction)

	// return the cache config
	return bigcache.Config{
		// number of shards (must be a power of 2)
		Shards: 2048,

		// time after which entry can be evicted
		LifeWindow: cfg.Cache.Eviction,

		// Interval between removing expired entries (clean up).
		// If set to <= 0 then no action is performed.
		// Setting to < 1 second is counterproductive — big cache has a one second resolution.
		CleanWindow: 5 * time.Minute,

		// rps * lifeWindow, used only in initial memory allocation
		MaxEntriesInWindow: 1000 * 10 * 60,

		// max entry size in bytes, used only in initial memory allocation
		MaxEntrySize: 1024,

		// cache will not allocate more memory than this limit, value in MB
		// if value is reached then the oldest entries can be overridden for the new ones
		// 0 value means no size limit
		HardMaxCacheSize: cfg.Cache.MaxSize,

		// callback fired when the oldest entry is removed because of its expiration time or no space left
		// for the new entry, or because delete was called. A bitmask representing the reason will be returned.
		// Default value is nil which means no callback and it prevents from unwrapping the oldest entry.
		OnRemove: nil,

		// OnRemoveWithReason is a callback fired when the oldest entry is removed because of its expiration time or no space left
		// for the new entry, or because delete was called. A constant representing the reason will be passed through.
		// Default value is nil which means no callback and it prevents from unwrapping the oldest entry.
		// Ignored if OnRemove is specified.
		OnRemoveWithReason: nil,

		// prints information about additional memory allocation
		Verbose: true,

		// Logger is a logging interface and used in combination with `Verbose`
		Logger: log,
	}
}
