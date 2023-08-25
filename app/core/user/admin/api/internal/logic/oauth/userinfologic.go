package oauth

import (
	userAdmin "app/core/user/admin/rpc/service"
	"app/std/net/http/logic"
	"context"
	"net/http"
	"strconv"

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
	userModel := l.svcCtx.UserAdminRpc
	userIdInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		resp.Error("用户id格式化失败：", err.Error())
		return
	}

	user, userErr := userModel.AdminShow(l.ctx, &userAdmin.AdminShowReq{UserId: userIdInt})
	if userErr != nil {
		resp.Error("user rpc adminShow err：", userErr.Error())
		return
	}

	if !user.Success {
		resp.Error(user.Message, nil)
		return
	}

	user.Data.Salt = ""
	user.Data.Password = ""
	resp.Success("登录成功！", user.Data)
	return
}
