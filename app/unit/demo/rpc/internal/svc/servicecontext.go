package svc

import (
	log "app/core/log/rpc/service"
	"app/std/util"
	"app/unit/demo/rpc/internal/config"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/zrpc"
	"runtime"
)

type ServiceContext struct {
	Config   config.Config
	LogError func(ctx context.Context, err error)
}

func NewServiceContext(c config.Config) *ServiceContext {
	conf := c.Mysql
	// 数据库初始化
	util.IsInstall(conf)
	return &ServiceContext{
		Config: c,
		LogError: func(ctx context.Context, err error) {
			pc, file, line, _ := runtime.Caller(1)
			funcName := runtime.FuncForPC(pc).Name()
			method := fmt.Sprintf("Error occurred in function %s at %s:%d\n", funcName, file, line)
			logRpc := log.NewService(zrpc.MustNewClient(c.LogRpc))
			logRpc.LogSystem(ctx, &log.SystemLogReq{
				AppName:  c.Name,
				LogLevel: "error",
				Method:   method,
				Message:  err.Error(),
			})
		},
	}
}
