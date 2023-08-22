package oauth

import (
	"app/std/net/http/logic"
	"context"
	"net/http"

	"app/core/user/authn/api/internal/svc"
	"app/core/user/authn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewRefreshLogic(header *http.Request, svcCtx *svc.ServiceContext) *RefreshLogic {
	ctx := header.Context()
	return &RefreshLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *RefreshLogic) Refresh(req *types.RefreshReq) (resp logic.Response) {
	// todo: add your logic here and delete this line

	return
}
