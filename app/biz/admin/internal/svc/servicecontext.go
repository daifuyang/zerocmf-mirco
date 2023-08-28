package svc

import (
	"app/biz/admin/internal/config"
	"app/core/system/menu/rpc/service"
	"app/std/apisix"
	"app/std/apisix/plugins"
	"app/std/util"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	SystemAdminRpc service.Service
}

func NewServiceContext(c config.Config) *ServiceContext {
	util.Install(func() {
		//	注册网关
		routes := []apisix.Route{
			{
				URI:       "/api/v1/admin/*",
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
		Config:         c,
		SystemAdminRpc: service.NewService(zrpc.MustNewClient(c.SystemMenuRpc)),
	}
}
