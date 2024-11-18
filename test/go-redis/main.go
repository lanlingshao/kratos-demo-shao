package main

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/lanlingshao/kratos-demo-shao/internal/resource/cache"
	rds "github.com/redis/go-redis/v9"
	"os"
	"time"
)

func main() {
	ctx := context.Background()
	logger := log.With(log.NewStdLogger(os.Stdout))
	opt := &rds.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
		// ReadTimeout:  time.Duration(2 * time.Millisecond),
		// WriteTimeout: time.Duration(2 * time.Millisecond),
		PoolSize:     10,
		MinIdleConns: 1,
	}
	rdsCli := cache.NewRedisClient(opt, logger)
	res := rdsCli.Set(ctx, "apple", "1", 2*time.Second)
	logger.Log(0, "res:", res.Val())
	val := rdsCli.Get(ctx, "apple")
	logger.Log(0, "val:", val.Val())
	time.Sleep(3 * time.Second)
	val = rdsCli.Get(ctx, "apple")
	logger.Log(0, "val:", val.Val())
}
