package logic

import (
	"app/core/user/admin/model"
	"context"
	"errors"

	"app/core/user/admin/rpc/internal/svc"
	"app/core/user/admin/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminShowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdminShowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminShowLogic {
	return &AdminShowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取单个管理员
func (l *AdminShowLogic) AdminShow(in *pb.AdminShowReq) (*pb.LoginResp, error) {
	ctx := l.ctx
	adminUserModel := l.svcCtx.AdminUserModel

	var (
		first    *model.AdminUser
		firstErr error
	)

	userId := in.GetUserId()
	username := in.GetUsername()
	if userId != 0 {
		// 查询当前用户是否存在
		first, firstErr = adminUserModel.FindOne(ctx, userId)
	} else if username != "" {
		first, firstErr = adminUserModel.FindOneByUsername(ctx, username)
	}

	if firstErr != nil {
		if errors.Is(firstErr, model.ErrNotFound) {
			return &pb.LoginResp{
				Message: "用户不存在！",
			}, nil
		}
		return nil, firstErr
	}

	return &pb.LoginResp{
		Success: true,
		Message: "获取成功！",
		Data: &pb.LoginData{
			Salt:     first.Salt,
			UserId:   first.Id,
			Username: first.Username,
			Password: first.Password,
		},
	}, nil
}
