# zerocmf内容管理框架

- [ ] 基于 `go` + `go-zero` + `gRPC`, 以`微服务`方式实现的通用管理平台.
- [ ] 目标:
    - [ ] 综合社区`最佳实践`, 力求`产品级`项目, 而不是 `demo 玩具`.
    - [ ] 通用业务单元: 可以自由组合使用.
- [ ] 文档:

## 在线体验

## 快速开始

### 系统需求

#### go

- go 1.18+
- go-zero 1.5+

#### Cli Tool

- [go-task](https://taskfile.dev/#/installation)
- goctl

#### 中间件

> 核心依赖:

- mysql8.0+: 数据库
- redis: 缓存
- go-queue: 消息队列

> 监控/告警/日志:

- Prometheus: 监控
- Jaeger: 分布式追踪
- Kibana: 日志
- Grafana: 图表

> 网关
- apisix

> 服务发现/注册/配置中心
- etcd

## WIP:

- 🚀 架构公式: [biz](app/biz) = [core](app/core) + [unit](app/unit)
    - ✅ [biz](app/biz): 产品构成
    - ✅ [core](app/core): 基础服务, 公共服务
    - ✅ [unit](app/unit): 业务单元, 通用业务单元

### 产品:

| [产品(app/biz) ](app/biz)             | 说明         |
|:------------------------------------|------------|
| 💖 [app/biz/admin](app/biz/admin)   | 通用后台管理系统   |
| 💖 [app/biz/portal](app/biz/portal) | 通用门户管理系统   |
| 💖 [app/biz/mall](app/biz/mall)     | 通用电商门新零售系统 |


### 鸣谢：

- [gp-zero](https://github.com/zeromicro/go-zero)
- [zeromall](https://github.com/zeromicro/zeromall)