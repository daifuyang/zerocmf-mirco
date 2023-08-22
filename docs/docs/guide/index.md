---
nav: 指南
  title: 指南
  order: 0
group:
  title: 基础
  order: -1
---

# 初始化

## 架构公式

biz = core + unit

- biz 产品构成
- core 基础服务，公共服务。由官方维护
- unit 通用业务单元

## 服务编码

- 服务的端口号 = 服务编号 + 服务类型编码

| 服务类型编码 | 说明      | 重要性  |
| ------------ | --------- | ------- |
| 0            | rpc       | ⭐⭐⭐️ |
| 1            | job       | ⭐      |
| 2            | admin api | ⭐      |
| 8            | web api   | ⭐⭐⭐  |

## 服务列表

| 服务            | 编号 | 说明       | 依赖服务 | 服务等级 |
| --------------- | ---- | ---------- | -------- | -------- |
| core/log        | 11X  | 日志服务   | 无       | L0       |
| core/user/admin | 101X | 管理员用户 | 无       | L0       |

## 环境准备

### go

- go 1.18+
- go-zero 1.54+

### cli tool

- [go-task](https://taskfile.dev/#/installation)
- goctl


## 启动项目
```shell
task init
```
