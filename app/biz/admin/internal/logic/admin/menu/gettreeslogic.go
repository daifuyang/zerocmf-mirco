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

type GetTreesLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewGetTreesLogic(header *http.Request, svcCtx *svc.ServiceContext) *GetTreesLogic {
	ctx := header.Context()
	return &GetTreesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *GetTreesLogic) GetTrees(req *types.MenusTreeReq) (resp logic.Response) {
	list, err := l.svcCtx.SystemAdminRpc.TreeList(l.ctx, &pb.TreeListReq{})
	if err != nil {
		resp.Error("系统错误！", err.Error())
		return
	}

	if !list.Success {
		resp.Error(list.Message, nil)
		return
	}

	resp.Success("获取成功！", list.Data)
	return
}
