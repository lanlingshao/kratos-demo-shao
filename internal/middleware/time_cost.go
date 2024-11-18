package middleware

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	mitlog "gitlab.mitradetech.com/package/mitrade-go-pkg/logger"
	"time"
)

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
