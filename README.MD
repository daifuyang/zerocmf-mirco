# zerocmf内容管理框架

- [] 基于 `go` + `go-zero` + `gRPC`, 以`微服务`方式实现的通用管理平台.
- [] 目标:
    - [] 综合社区`最佳实践`, 力求`产品级`项目, 而不是 `demo 玩具`.
    - [] 通用业务单元: 可以自由组合使用.
- [] 文档:

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

### 核心服务:

| [核心服务(app/core)](app/core)                 | 服务           |
|:-------------------------------------------|--------------|
| [app/core/user/authn](app/core/user/authn) | 用户服务: 注册/登录  |
| [app/core/user/authz](app/core/user/authz) | 用户服务: 访问权限控制 |
| [app/core/queue](app/core/queue)           | 基础队列服务       |
| [app/core/log](app/core/log)               | 日志记录         |


### 服务矩阵:

| [业务单元(app/unit)](app/unit)                              | 说明                 |
|:--------------------------------------------------------|--------------------|
| ☀️ [app/unit/license/key](app/unit/license/key)         | 注册码服务: 售卖 app `密钥` |
| ☀️ [app/unit/cart](app/unit/cart)                       | 购物车系统              |
| ☀️ [app/unit/order](app/unit/order)                     | 订单系统               |
| ☀️ [app/unit/finance/payment](app/unit/finance/payment) | 支付系统               |


### 鸣谢：
- [gp-zero](https://github.com/zeromicro/go-zero)
- [zeromall](https://github.com/zeromicro/zeromall)