package svc

import (
	"app/core/user/admin/api/internal/config"
	userAdmin "app/core/user/admin/rpc/service"
	"app/std/apisix"
	"app/std/apisix/plugins"
	"app/std/util"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	UserAdminRpc userAdmin.Service
}

func NewServiceContext(c config.Config) *ServiceContext {

	util.Install(c.Mysql, func() {

		//	注册网关
		routes := []apisix.Route{
			{
				URI:       "/api/v1/user/admin/*",
				Name:      c.Apisix.Name,
				ServiceID: c.Apisix.Name,
				Status:    1,
				Plugins:   plugins.RoutePlugin(c.Apisix.Host),
			},
		}

		err := c.Apisix.Register(routes)
		if err != nil {
			panic(err)
		}

	})

	return &ServiceContext{
		Config:       c,
		UserAdminRpc: userAdmin.NewService(zrpc.MustNewClient(c.UserAdminRpc)),
	}
}
