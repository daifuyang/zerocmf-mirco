package logic

import (
	"context"

	"app/core/system/menu/rpc/internal/svc"
	"app/core/system/menu/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除菜单
func (l *DeleteLogic) Delete(in *pb.SaveReq) (*pb.Response, error) {
	// todo: add your logic here and delete this line

	return &pb.Response{}, nil
}
