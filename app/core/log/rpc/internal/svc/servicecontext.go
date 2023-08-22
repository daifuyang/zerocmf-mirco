package svc

import (
	"app/core/log/model"
	"app/core/log/rpc/internal/config"
	"app/std/util"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config         config.Config
	SystemLogModel model.SystemLogModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conf := c.Mysql
	util.Install(conf)
	return &ServiceContext{
		Config:         c,
		SystemLogModel: model.NewSystemLogModel(sqlx.NewMysql(c.Mysql.Dsn()), c.Cache),
	}
}
