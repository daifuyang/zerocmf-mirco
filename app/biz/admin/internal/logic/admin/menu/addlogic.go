package menu

import (
	"app/core/system/menu/rpc/pb"
	"app/std/net/http/logic"
	"context"
	"github.com/jinzhu/copier"
	"net/http"

	"app/biz/admin/internal/svc"
	"app/biz/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewAddLogic(header *http.Request, svcCtx *svc.ServiceContext) *AddLogic {
	ctx := header.Context()
	return &AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req *types.MenuSaveReq) (resp logic.Response) {
	adminRpc := l.svcCtx.SystemMenuRpc
	rpcReq := pb.SaveReq{}
	err := copier.Copy(&rpcReq, &req)
	if err != nil {
		resp.Error("系统错误！", err.Error())
		return
	}
	adminResp, saveErr := adminRpc.Save(l.ctx, &rpcReq)
	if saveErr != nil {
		resp.Error("系统错误！", saveErr.Error())
		return
	}
	if !adminResp.Success {
		resp.Error(adminResp.Message, nil)
		return
	}
	resp.Success(adminResp.Message, adminResp.Data)
	return
}
