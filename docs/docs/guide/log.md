---
nav: 指南
group:
  title: 基础
  order: 1
order: 3
---

# 系统日志

## 日志微服务

日志微服务所在目录为 app/core/log/rpc。一般由官方维护，如需二次开发，请遵循规范自行修改。

## 日志收集

我们提供了日志服务，包含了系统日志收集，登录日志，和请求日志等。可以直接在 rpc 服务下配置使用。默认 goctl template 会自动接入。手动实现如下：

### 配置 yaml

配置 etc/main.yaml 开启本地日志服务、默认存储在 data/logs 目录下，保留 7 天，最大不超过 2M。

```yaml
Log:
  Mode: file
  Path: data/logs
  KeepDays: 7
  MaxSize: 2

LogRpc:
  Timeout: 30000
  Etcd:
    Key: app.core.log.rpc
```

### 修改config.go

修改 internal/config/config.go

```go
package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type OauthConfig struct {
	LogRpc zrpc.RpcClientConf
}
```

### 初始化服务

修改 internal/svc/servicecontext.go

```go

import (
	log "app/core/log/rpc/service"
  "context"
	"github.com/zeromicro/go-zero/zrpc"
	"runtime"
)

type ServiceContext struct {
	LogRpc         log.Service
	LogError       func(ctx context.Context, err error)
}


func NewServiceContext(c config.OauthConfig) *ServiceContext {
  logRpc := log.NewService(zrpc.MustNewClient(c.LogRpc)) // 实例化日志服务
  return &ServiceContext{
		LogError: func(ctx context.Context, err error) {
			pc, file, line, _ := runtime.Caller(1)
			funcName := runtime.FuncForPC(pc).Name()
			method := fmt.Sprintf("Error occurred in function %s at %s:%d\n", funcName, file, line)
			logRpc.LogSystem(ctx, &log.SystemLogReq{
				AppName:  c.Name,
				LogLevel: "error",
				Method:   method,
				Message:  err.Error(),
			})
		},
	}
}

```

## 使用示例

```go
err := errors.New("系统出错了！")
if err != nil {
		l.svcCtx.LogError(l.ctx, err)
		return nil, err
	}
```
