package server

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	v1 "hello_world/api/helloworld/v1"
	"hello_world/internal/conf"
	"hello_world/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			Auth(),
			Log(),
			authMiddleware,
			loggingMiddleware,
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterGreeterHTTPServer(srv, greeter)
	return srv
}

func Auth() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			fmt.Println("auth middleware start")
			if tr, ok := transport.FromServerContext(ctx); ok {
				tr.RequestHeader().Get("Authorization")
				// do some logic
			}
			return handler(ctx, req)
		}
	}
}

func Log() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			fmt.Println("log middleware start")
			return handler(ctx, req)
		}
	}
}

func authMiddleware(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
		fmt.Println("auth middleware in", req)
		reply, err = handler(ctx, req)
		fmt.Println("auth middleware out", reply)
		return
	}
}

func loggingMiddleware(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
		fmt.Println("logging middleware in", req)
		reply, err = handler(ctx, req)
		fmt.Println("logging middleware out", reply)
		return
	}
}
