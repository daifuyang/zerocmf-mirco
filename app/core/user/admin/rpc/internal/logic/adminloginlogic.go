package logic

import (
	"app/core/user/admin/model"
	"app/core/user/admin/rpc/internal/svc"
	"app/core/user/admin/rpc/pb"
	"context"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdminLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminLoginLogic {
	return &AdminLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 管理员登录
func (l *AdminLoginLogic) AdminLogin(in *pb.LoginReq) (*pb.LoginResp, error) {

	config := l.svcCtx.Config

	ctx := l.ctx

	salt := l.svcCtx.Config.Mysql.AuthCode

	adminUserModel := l.svcCtx.AdminUserModel

	username := in.GetUsername()
	if strings.TrimSpace(username) == "" {
		return &pb.LoginResp{
			Success: false,
			Message: "用户名不能为空！",
		}, nil
	}
	password := in.GetPassword()
	if strings.TrimSpace(password) == "" {
		return &pb.LoginResp{
			Success: false,
			Message: "密码不能为空！",
		}, nil
	}

	// 查询当前用户是否存在
	first, firstErr := adminUserModel.Where("username = ?", username).First(ctx)
	if firstErr != nil {
		if errors.Is(firstErr, model.ErrNotFound) {
			return &pb.LoginResp{
				Message: "用户名或密码错误！",
			}, nil
		}
		l.svcCtx.LogError(l.ctx, firstErr)
		return nil, firstErr
	}

	// 密码比对
	hashedErr := bcrypt.CompareHashAndPassword([]byte(first.Password), []byte(salt+password))
	if hashedErr != nil {
		return &pb.LoginResp{
			Message: "用户名或密码错误！",
		}, nil
	}
	// 生成token并下发
}
