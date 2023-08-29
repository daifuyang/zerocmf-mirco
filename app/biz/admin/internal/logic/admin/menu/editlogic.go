package menu

import (
	"app/std/net/http/logic"
	"context"
	"net/http"

	"app/biz/admin/internal/svc"
	"app/biz/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewEditLogic(header *http.Request, svcCtx *svc.ServiceContext) *EditLogic {
	ctx := header.Context()
	return &EditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *EditLogic) Edit(req *types.MenuSaveReq) (resp logic.Response) {
	return NewAddLogic(l.header, l.svcCtx).Add(req)
}
