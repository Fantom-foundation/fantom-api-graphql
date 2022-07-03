// Package filecache stores files uploaded into IPFS, so they are available
// before the local IPFS node is able to download them.
package filecache

import (
	"errors"
	"fantom-api-graphql/internal/config"
	"fantom-api-graphql/internal/logger"
	"fmt"
	"os"
)

// FileCache represents cache in a filesystem.
type FileCache struct {
	dir string

	// log represents the logger to be used by the filecache.
	log logger.Logger
}

func New(cfg *config.Config, l logger.Logger) *FileCache {
	fc := FileCache{
		dir: cfg.Ipfs.FileCacheDir,
		log: l,
	}
	err := os.MkdirAll(fc.dir, 0700)
	if err != nil {
		panic(fmt.Errorf("unable to create filecache directory \"%s\"; %s", fc.dir, err))
	}
	return &fc
}

func (fc *FileCache) PullIpfsFile(cid string) (content []byte) {
	content, err := os.ReadFile(fc.dir + cid)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		fc.log.Errorf("unable to read filecache; %s", err)
	}
	return content
}

func (fc *FileCache) PushIpfsFile(cid string, content []byte) (err error) {
	return os.WriteFile(fc.dir+cid, content, 0644)
}
