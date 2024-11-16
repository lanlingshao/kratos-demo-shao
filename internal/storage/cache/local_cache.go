package cache

import (
	"github.com/bluele/gcache"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/lanlingshao/kratos-demo-shao/internal/conf"
)

func NewLocalCacheClient(conf *conf.Data, logger log.Logger) gcache.Cache {
	// size是缓存可以存的key数量
	localCacheClient := gcache.New(int(conf.LocalCache.GetSize())).LRU().Build()
	return localCacheClient
}
