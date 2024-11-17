package cache

import (
	"github.com/bluele/gcache"
	"github.com/go-kratos/kratos/v2/log"
)

type LocalCacheOption struct {
	Size int // Size is the amount of keys that the cache can store
}

func NewLocalCacheClient(option *LocalCacheOption, logger log.Logger) gcache.Cache {
	if option == nil || option.Size <= 0 {
		option = &LocalCacheOption{
			Size: 10000,
		}
	}
	localCacheClient := gcache.New(option.Size).LRU().Build()
	return localCacheClient
}
