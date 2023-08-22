package svc

import (
	{{.configImport}}
)

type ServiceContext struct {
	OauthConfig {{.config}}
	{{.middleware}}
}

func NewServiceContext(c {{.config}}) *ServiceContext {
	return &ServiceContext{
		OauthConfig: c,
		{{.middlewareAssignment}}
	}
}
