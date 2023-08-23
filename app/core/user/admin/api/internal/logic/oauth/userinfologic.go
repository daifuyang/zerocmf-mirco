package oauth

import (
	"app/std/net/http/logic"
	"context"
	"net/http"

	"app/core/user/admin/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(header *http.Request, svcCtx *svc.ServiceContext) *UserInfoLogic {
	ctx := header.Context()
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp logic.Response) {
	r := l.header
	userId := r.Header.Get("X-User-ID")
	resp.Success("hello admin", userId)
	return
}
