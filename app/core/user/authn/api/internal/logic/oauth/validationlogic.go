package oauth

import (
	"app/std/net/http/logic"
	"context"
	"net/http"

	"app/core/user/authn/api/internal/svc"
	"app/core/user/authn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ValidationLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewValidationLogic(header *http.Request, svcCtx *svc.ServiceContext) *ValidationLogic {
	ctx := header.Context()
	return &ValidationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *ValidationLogic) Validation(req *types.ValidationReq) (resp logic.Response) {
	// todo: add your logic here and delete this line

	return
}
