package svc

import (
	"app/core/user/admin/api/internal/config"
)

type ServiceContext struct {
	OauthConfig config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		OauthConfig: c,
	}
}
