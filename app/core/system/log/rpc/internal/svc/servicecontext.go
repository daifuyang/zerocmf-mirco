package svc

import (
	"app/core/log/rpc/internal/config"
	"app/std/util"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	conf := c.Mysql
	util.Install(func() {
		util.InitDb(conf)
	})
	return &ServiceContext{
		Config: c,
	}
}
