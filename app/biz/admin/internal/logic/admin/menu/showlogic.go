package menu

import (
	"app/core/system/menu/rpc/pb"
	"app/std/net/http/logic"
	"context"
	"net/http"

	"app/biz/admin/internal/svc"
	"app/biz/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShowLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewShowLogic(header *http.Request, svcCtx *svc.ServiceContext) *ShowLogic {
	ctx := header.Context()
	return &ShowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *ShowLogic) Show(req *types.MenuShowReq) (resp logic.Response) {
	adminMenuRpc := l.svcCtx.SystemMenuRpc
	show, err := adminMenuRpc.Show(l.ctx, &pb.ShowReq{MenuId: req.Id})
	if err != nil {
		resp.Error("获取失败！", err.Error())
		return
	}

	if !show.Success {
		resp.Error(show.Message, nil)
		return
	}

	resp.Success(show.Message, show.Data)
	return
}
