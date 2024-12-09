package data

import (
	"github.com/bluele/gcache"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/lanlingshao/kratos-demo-shao/internal/conf"
	"github.com/lanlingshao/kratos-demo-shao/internal/resource/cache"
	"github.com/lanlingshao/kratos-demo-shao/internal/resource/db"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewMysqlClient, NewRedisClient, NewLocalCacheClient, NewGreeterRepo)

// Data .
type Data struct {
	mysqlClient *gorm.DB
	redisClient redis.UniversalClient
	localCache  gcache.Cache
	logg        *log.Helper
}

// NewData .
func NewData(c *conf.Data, mysqlClient *gorm.DB, redisClient redis.UniversalClient, localCache gcache.Cache, logger log.Logger) (*Data, func(), error) {
	logg := log.NewHelper(log.With(logger, "module", "internal/data"))
	d := &Data{
		mysqlClient: mysqlClient,
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
	redisConf := &redis.Options{
		Addr:         conf.Redis.Addr,
		Password:     conf.Redis.Password,
		DB:           int(conf.Redis.Database),
		ReadTimeout:  time.Duration(conf.Redis.GetReadTimeout()) * time.Millisecond,
		WriteTimeout: time.Duration(conf.Redis.GetWriteTimeout()) * time.Millisecond,
		PoolSize:     int(conf.Redis.GetPoolSize()),
	}

	return cache.NewRedisClient(redisConf, logger)
}

func NewLocalCacheClient(conf *conf.Data, logger log.Logger) gcache.Cache {
	option := &cache.LocalCacheOption{
		Size: int(conf.LocalCache.GetSize()),
	}
	return cache.NewLocalCacheClient(option, logger)
}

func NewMysqlClient(conf *conf.Data, logger log.Logger) *gorm.DB {
	return db.NewMySQLClient(conf, logger)
}
