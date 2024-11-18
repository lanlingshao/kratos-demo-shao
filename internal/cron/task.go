package cron

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/lanlingshao/kratos-demo-shao/internal/service"
	xxl "github.com/xxl-job/xxl-job-executor-go"
	"github.com/xxl-job/xxl-job-executor-go/example/task"
)

type CronTaskOption struct {
	Executor xxl.Executor
	Service  *service.GreeterService
	Log      *log.Helper
}

func RegisterTaskHandler(option CronTaskOption) {
	option.Executor.RegTask("task.test", task.Test)
	option.Executor.RegTask("task.test2", task.Test2)
	option.Executor.RegTask("task.panic", task.Panic)
}
