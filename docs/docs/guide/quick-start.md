---
nav: 指南
group:
  title: 基础
  order: 1
order: -1
---

# 快速开始

## 开发 rpc

先确认模块所属的位置，这里以 unit 为例,创建名为 demo 的应用

### 创建文件夹

```shell
mkdir demo
```

### 初始化模块

```shell
go mod init app/unit/demo
go mod tidy
```

### 工具区使用（项目根目录）

```
  go work use app/unit/demo
```

### 创建微服务

```
cd demo
mkdir rpc // 创建应用目录
cd rpc
go mod init app/unit/demo/rpc // 推荐单独生成模块

// 工具区（根目录）下使用
go work use app/unit/demo/rpc

goctl rpc --o pb/demo.proto //生成proto文件
```

### 修改 demo.proto

```proto
syntax = "proto3";

// 控制 grpc 注册服务名
package demo.rpc.v1;

// 统一命名为 pb，无需更改
option go_package="./pb";

message Request {
  string ping = 1;
}

message Response {
  string pong = 1;
}

// 控制生成的 rpc 文件夹名, 统一命名为 Service， 无需更改

service Service {
  rpc Ping(Request) returns(Response);
}
```

### 生成微服务

```shell
cd demo
goctl rpc protoc "./pb/demo.proto" --go_out=./ --go-grpc_out=./ --zrpc_out=./ --style gozero --home ../../../../template
```

最后，你可以在 demo 服务下面创建 Taskfile.yml 文件。用来创建执行任务脚本。

### 最终结构

```shell
tree
.
├── Taskfile.yml
├── go.mod
└── rpc
    ├── data
    │   └── logs
    │       ├── access.log
    │       ├── error.log
    │       ├── severe.log
    │       ├── slow.log
    │       └── stat.log
    ├── demo.go
    ├── etc
    │   └── demo.yaml
    ├── go.mod
    ├── go.sum
    ├── internal
    │   ├── config
    │   │   └── config.go
    │   ├── logic
    │   │   └── pinglogic.go
    │   ├── server
    │   │   └── serviceserver.go
    │   └── svc
    │       └── servicecontext.go
    ├── pb
    │   ├── demo.pb.go
    │   ├── demo.proto
    │   └── demo_grpc.pb.go
    └── service
        └── service.go
```

## 数据库

### 配置 etc/main.yaml

```yaml
# mysql: 参数必带
Mysql:
  Username: root
  Password: '123456'
  Host: localhost
  port: 3306
  DbName: zerocmf_user
  charset: utf8mb4
  Collate: utf8mb4_unicode_ci
  ParseTime: true
  AuthCode: zerocmf
```

### 配置 config.go

修改 internal/config/config.go

```go
type OauthConfig struct {
    Mysql  database.Mysql
    LogRpc zrpc.RpcClientConf
}
```
