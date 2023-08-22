package logic

import (
	"app/std/net/http/logic"
	"context"
	"net/http"

	"app/core/user/admin/api/internal/svc"
	"app/core/user/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApiLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewApiLogic(header *http.Request, svcCtx *svc.ServiceContext) *ApiLogic {
	ctx := header.Context()
	return &ApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *ApiLogic) Api(req *types.Request) (resp logic.Response) {
	// todo: add your logic here and delete this line

	return
}
