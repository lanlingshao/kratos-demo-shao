package cache

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/lanlingshao/kratos-demo-shao/internal/conf"
	rds "github.com/redis/go-redis/v9"
	"time"
)

func NewRedisClient(ctx context.Context, conf *conf.Data, logger log.Logger) *rds.Client {
	var rdsClient *rds.Client
	redisConf := &rds.Options{
		Addr:     conf.Redis.Addr,
		Password: conf.Redis.Password,
		DB:       int(conf.Redis.Database),
	}
	if conf.Redis.ReadTimeout > 0 {
		redisConf.ReadTimeout = time.Duration(conf.Redis.ReadTimeout) * time.Millisecond
	}
	if conf.Redis.WriteTimeout > 0 {
		redisConf.WriteTimeout = time.Duration(conf.Redis.WriteTimeout) * time.Millisecond
	}
	if conf.Redis.PoolSize > 0 {
		redisConf.PoolSize = int(conf.Redis.PoolSize)
	}
	rdsClient = rds.NewClient(redisConf)
	return rdsClient
}
