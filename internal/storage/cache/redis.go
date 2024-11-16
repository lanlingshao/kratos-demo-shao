package cache

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/lanlingshao/kratos-demo-shao/internal/conf"
	rds "github.com/redis/go-redis/v9"
	"time"
)

func NewRedisClientOld(conf *conf.Data, logger log.Logger) *rds.Client {
	redisConf := &rds.Options{
		Addr:         conf.Redis.Addr,
		Password:     conf.Redis.Password,
		DB:           int(conf.Redis.Database),
		ReadTimeout:  time.Duration(conf.Redis.GetReadTimeout()) * time.Millisecond,
		WriteTimeout: time.Duration(conf.Redis.GetWriteTimeout()) * time.Millisecond,
		PoolSize:     int(conf.Redis.GetPoolSize()),
	}
	redisClient := rds.NewClient(redisConf)
	return redisClient
}

func NewRedisClient(conf *conf.Data, logger log.Logger) *rds.Client {
	logg := log.NewHelper(log.With(logger, "module", "internal/storage/cache/redis"))
	redisConf := &rds.Options{
		Addr:         conf.Redis.Addr,
		Password:     conf.Redis.Password,
		DB:           int(conf.Redis.Database),
		ReadTimeout:  time.Duration(conf.Redis.GetReadTimeout()) * time.Millisecond,
		WriteTimeout: time.Duration(conf.Redis.GetWriteTimeout()) * time.Millisecond,
		PoolSize:     int(conf.Redis.GetPoolSize()),
	}
	redisClient := rds.NewClient(redisConf)
	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()
	err := redisClient.Ping(timeout).Err()
	if err != nil {
		logg.Fatalf("redis connect error: %v", err)
	}
	return redisClient
}
