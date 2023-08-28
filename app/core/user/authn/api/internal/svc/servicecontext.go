package svc

import (
	userAdmin "app/core/user/admin/rpc/service"
	"app/core/user/authn/api/internal/config"
	"app/std/apisix"
	"app/std/util"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	UserAdminRpc userAdmin.Service
}

func NewServiceContext(c config.Config) *ServiceContext {

	util.Install(func() {

		util.InitDb(c.Mysql)

		//	注册网关
		routes := []apisix.Route{
			{
				URI:       "/api/authn/*",
				Name:      c.Apisix.Name,
				ServiceID: c.Apisix.Name,
				Status:    1,
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
