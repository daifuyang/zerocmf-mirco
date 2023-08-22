package {{.pkgName}}

import (
    "net/http"
    "app/std/net/http/logic"
	{{.imports}}
)

type {{.logic}} struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func New{{.logic}}(header *http.Request, svcCtx *svc.ServiceContext) *{{.logic}} {
    ctx := header.Context()
	return &{{.logic}}{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *{{.logic}}) {{.function}}({{.request}}) (resp logic.Response) {
	// todo: add your logic here and delete this line

	return
}
