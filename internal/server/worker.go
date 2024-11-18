package server

import (
	"context"
	"fmt"
	"github.com/lanlingshao/kratos-demo-shao/internal/conf"
	"github.com/xxl-job/xxl-job-executor-go"
	"github.com/xxl-job/xxl-job-executor-go/example/task"
	"log"
)

type CronWorker struct {
	exec xxl.Executor
}

func (c *CronWorker) Start(ctx context.Context) error {
	return c.exec.Run()
}

func (c *CronWorker) Stop(ctx context.Context) error {
	c.exec.Stop()
	return nil
}

func NewCronWorker(c *conf.Server) *CronWorker {
	exec := xxl.NewExecutor(
		xxl.ServerAddr(c.CronWorker.GetAddr()),
		xxl.AccessToken(c.CronWorker.GetAccessToken()),   // 请求令牌(默认为空)
		xxl.ExecutorIp(c.CronWorker.GetExecutorIp()),     // 可自动获取
		xxl.ExecutorPort(c.CronWorker.GetExecutorPort()), // 默认9999（非必填）
		xxl.RegistryKey(c.CronWorker.GetRegistryKey()),   // 执行器名称
		xxl.SetLogger(&logger{}),                         // 自定义日志
	)
	exec.Init()
	exec.Use(customMiddleware)
	// 设置日志查看handler
	exec.LogHandler(customLogHandle)
	// 注册任务handler
	exec.RegTask("task.test", task.Test)
	exec.RegTask("task.test2", task.Test2)
	exec.RegTask("task.panic", task.Panic)
	return &CronWorker{
		exec: exec,
	}
}

// 自定义日志处理器
func customLogHandle(req *xxl.LogReq) *xxl.LogRes {
	return &xxl.LogRes{Code: xxl.SuccessCode, Msg: "", Content: xxl.LogResContent{
		FromLineNum: req.FromLineNum,
		ToLineNum:   2,
		LogContent:  "这个是自定义日志handler",
		IsEnd:       true,
	}}
}

// xxl.Logger接口实现
type logger struct{}

func (l *logger) Info(format string, a ...interface{}) {
	fmt.Println(fmt.Sprintf("自定义日志 - "+format, a...))
}

func (l *logger) Error(format string, a ...interface{}) {
	log.Println(fmt.Sprintf("自定义日志 - "+format, a...))
}

// 自定义中间件
func customMiddleware(tf xxl.TaskFunc) xxl.TaskFunc {
	return func(cxt context.Context, param *xxl.RunReq) string {
		log.Println("I am a middleware start")
		res := tf(cxt, param)
		log.Println("I am a middleware end")
		return res
	}
}