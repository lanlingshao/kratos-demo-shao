package middleware

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/middleware"
	mitlog "gitlab.mitradetech.com/package/mitrade-go-pkg/logger"
	"time"
)

func TimeCostMiddleware(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
		defer func(startTime time.Time) {
			fmt.Println("time cost", time.Since(startTime))
		}(time.Now())
		fmt.Println("time cost enter in")
		reply, err = handler(ctx, req)
		return
	}
}

func TimeCostMiddleware2() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			defer func(startTime time.Time) {
				mitlog.Info("time cost2<<<", time.Since(startTime))
			}(time.Now())
			return handler(ctx, req)
		}
	}
}
