package user

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/lanlingshao/kratos-demo-shao/internal/service"
	"github.com/xxl-job/xxl-job-executor-go"
)

type UserCronTask struct {
	greeter *service.GreeterService
}

func (u *UserCronTask) BatchUpdateUser(cxt context.Context, greeter *service.GreeterService) xxl.TaskFunc {
	return func(cxt context.Context, param *xxl.RunReq) string {

		return "BatchUpdateUser done"
	}

}

func BatchUpdateUser(cxt context.Context, param *xxl.RunReq) (msg string) {
	fmt.Println("test one task" + param.ExecutorHandler + " param：" + param.ExecutorParams + " log_id:" + xxl.Int64ToStr(param.LogID))
	return "BatchUpdateUser done"
}

func BatchDeleteUser(cxt context.Context, logger log.Logger) (msg string) {
	return ""
}

func customMiddleware(tf xxl.TaskFunc) xxl.TaskFunc {
	return func(cxt context.Context, param *xxl.RunReq) string {
		fmt.Println("I am a middleware start")
		res := tf(cxt, param)
		fmt.Println("I am a middleware end")
		return res
	}
}
