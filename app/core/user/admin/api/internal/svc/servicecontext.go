package svc

import (
	"app/core/user/admin/api/internal/config"
	"app/std/apisix"
	"app/std/apisix/plugins"
	"app/std/util"
)

type ServiceContext struct {
	Config config.Config
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
		Config: c,
	}
}
