package config

import (
	"app/std/apisix"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Etcd          discov.EtcdConf `json:",optional,inherit"`
	Apisix        apisix.Apisix
	SystemMenuRpc zrpc.RpcClientConf
}
