package logic

import (
	"context"

	"app/core/system/menu/rpc/internal/svc"
	"app/core/system/menu/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveLogic {
	return &SaveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 保存菜单
func (l *SaveLogic) Save(in *pb.SaveReq) (*pb.Response, error) {
	// todo: add your logic here and delete this line

	return &pb.Response{}, nil
}
