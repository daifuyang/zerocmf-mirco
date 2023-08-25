package config

import (
	"app/std/apisix"
	"app/std/database"
	"app/std/oauth"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Etcd         discov.EtcdConf `json:",optional,inherit"`
	UserAdminRpc zrpc.RpcClientConf
	Auth         struct {
		AccessSecret string
	}
	Mysql  database.Mysql
	Oauth  oauth.Config
	Apisix apisix.Apisix
}
