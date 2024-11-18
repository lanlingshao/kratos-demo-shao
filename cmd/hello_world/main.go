package main

import (
	"context"
	"flag"
	"os"
	"time"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/lanlingshao/kratos-demo-shao/internal/conf"
	"github.com/lanlingshao/kratos-demo-shao/internal/server"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server, ws *server.CronWorker) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
			ws,
		),
		// 服务注册
		// kratos.Registrar(rg),
		kratos.StopTimeout(10*time.Second),
		kratos.RegistrarTimeout(10*time.Second),
		// 增加在服务启动之前执行的操作
		kratos.BeforeStart(func(ctx context.Context) error {
			logger.Log(log.LevelInfo, "start hello world")
			return nil
		}),
		// 增加在服务启动之后执行的操作
		kratos.AfterStart(func(ctx context.Context) error {
			logger.Log(log.LevelInfo, "start after hello world")
			return nil
		}),
		// 增加在服务停止之前执行的操作
		kratos.BeforeStop(func(ctx context.Context) error {
			logger.Log(log.LevelInfo, "start after hello world")
			return nil
		}),
		// 增加在服务停止之后执行的操作
		kratos.AfterStop(func(ctx context.Context) error {
			logger.Log(log.LevelInfo, "start after hello world")
			return nil
		}),
	)
}

func main() {
	flag.Parse()
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	// 读取配置文件，可以改成从 nacos、applo等配置中心获取
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	// 将读取到到配置写入到conf.Bootstrap结构体中
	// 主要包括server配置和data配置，分别在internal/conf/conf.proto中定义，
	// 然后通过protoc-gen-go生成结构体到internal/conf/conf.pb.go中
	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	// 通过 wire 将依赖注入,并调用 newApp 方法，
	// cleanup定义一些应用退出后的清理操作，如：关闭数据库连接等
	app, cleanup, err := wireApp(bc.Server, bc.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	// 调用github.com/go-kratos/kratos/v2/app.go中的 func (a *App) Run() error 启动app，然后等待停止信号，
	// 收到停止信号后调用github.com/go-kratos/kratos/v2/app.go中 func (a *App) Stop() (err error) 停掉app
	if err := app.Run(); err != nil {
		panic(err)
	}
}
