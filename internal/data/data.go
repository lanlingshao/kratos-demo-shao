package data

import (
	"github.com/bluele/gcache"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/lanlingshao/kratos-demo-shao/internal/conf"
	"github.com/lanlingshao/kratos-demo-shao/internal/storage/cache"
	"github.com/redis/go-redis/v9"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewRedisClient, NewLocalCacheClient, NewGreeterRepo)

// Data .
type Data struct {
	redisClient redis.UniversalClient
	localCache  gcache.Cache
	logg        *log.Helper
}

// NewData .
func NewData(c *conf.Data, redisClient redis.UniversalClient, localCache gcache.Cache, logger log.Logger) (*Data, func(), error) {
	logg := log.NewHelper(log.With(logger, "module", "internal/data"))
	d := &Data{
		redisClient: redisClient,
		logg:        logg,
		localCache:  localCache,
	}
	cleanup := func() {
		logg.Info("closing the data resources")
		if err := d.redisClient.Close(); err != nil {
			logg.Error(err)
		}
	}
	return d, cleanup, nil
}

func NewRedisClient(conf *conf.Data, logger log.Logger) redis.UniversalClient {
	return cache.NewRedisClient(conf, logger)
}

func NewLocalCacheClient(conf *conf.Data, logger log.Logger) gcache.Cache {
	return cache.NewLocalCacheClient(conf, logger)
}
