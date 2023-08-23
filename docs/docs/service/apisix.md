---
nav:
  title: 服务模块
  order: 1
group:
  title: 通用模块
  order: 1
order: 0
---

# apisix网关支持

## etc配置

### apisix配置
```yaml
Apisix:
  apiKey: edd1c9f034335f136f87ad84b625c8f1
  Host: localhost
  Name: user-admin-api
  upstream:
    nodes:
      - host: 192.168.8.169
        port: 10108
        weight: 1
    timeout:
      connect: 30
      send: 30
      read: 30
    type: roundrobin
    scheme: http
    pass_host: pass
    keepalive_pool:
      idle_timeout: 60
      requests: 1000
      size: 320
```

### Auth授权配置
```yaml
Auth:
  AccessSecret: "youSerret" # 替换成你的加密秘钥
```
## config配置
```go
import (
    "app/std/apisix"
)
type Config struct {
    Apisix apisix.Apisix // 标准库提供
    Auth  struct {
        AccessSecret string
    }
}
```

### svc配置

```go

// 首次初始化，注册网关路由和身份验证
import (
  "app/std/apisix"
  "app/std/apisix/plugins"
  "app/std/util"
)
util.Install(c.Mysql, func() {

    //	注册网关
    routes := []apisix.Route{{
        URI:       "you service route", // 可以使用*。参考apisix路由文档
        Name:      c.Apisix.Name,
        ServiceID: c.Apisix.Name,
        Status:    1,
        Plugins:   plugins.RoutePlugin, // 使用forward-auth进行jwt身份验证
    }}

    err := c.Apisix.Register(routes)
    if err != nil {
        panic(err)
    }
})
```
