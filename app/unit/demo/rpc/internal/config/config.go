package config

import (
	"app/std/database"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql  database.Mysql
    LogRpc zrpc.RpcClientConf
    Cache cache.CacheConf
}
