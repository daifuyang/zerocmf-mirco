package config

import (
	"app/std/apisix"
	"app/std/database"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Etcd         discov.EtcdConf `json:",optional,inherit"`
	UserAdminRpc zrpc.RpcClientConf
	Mysql        database.Mysql
	Apisix       apisix.Apisix
}
