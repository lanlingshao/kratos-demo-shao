package cache

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	rds "github.com/redis/go-redis/v9"
	"time"
)

func NewRedisClient(options *rds.Options, logger log.Logger) *rds.Client {
	logg := log.NewHelper(log.With(logger, "module", "internal/storage/cache/redis"))
	redisClient := rds.NewClient(options)
	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()
	err := redisClient.Ping(timeout).Err()
	if err != nil {
		logg.Fatalf("redis connect error: %v", err)
	}
	return redisClient
}
