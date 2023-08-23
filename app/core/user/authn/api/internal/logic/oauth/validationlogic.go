package oauth

import (
	"app/std/net/http/logic"
	"app/std/oauth"
	"context"
	"net/http"
	"strconv"
	"strings"

	"app/core/user/authn/api/internal/svc"
	"app/core/user/authn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ValidationLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	writer http.ResponseWriter
	svcCtx *svc.ServiceContext
}

func NewValidationLogic(writer http.ResponseWriter, header *http.Request, svcCtx *svc.ServiceContext) *ValidationLogic {
	ctx := header.Context()
	return &ValidationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		writer: writer,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *ValidationLogic) Validation(req *types.ValidationReq) (resp logic.Response) {
	r := l.header
	w := l.writer
	conf := l.svcCtx.Config
	authHeader := r.Header.Get("Authorization")
	// 判断 Authorization 头部是否以 "Bearer " 开头
	if !strings.HasPrefix(authHeader, "Bearer ") {
		resp.Error("身份验证失败！", nil)
	}
	// 提取 Bearer Token 部分
	bearerToken := authHeader[len("Bearer "):]

	// 在这里您可以使用 bearerToken 进行验证、解析等操作
	// 示例中只是简单地输出 Token

	redirectURL := "http://" + conf.Apisix.Host + ":" + strconv.Itoa(conf.Port)

	oauthServer := oauth.NewServer(oauth.Config{
		ClientID:     conf.Oauth.ClientID,
		ClientSecret: conf.Oauth.ClientSecret,
		RedirectURL:  redirectURL,
		AccessSecret: conf.Auth.AccessSecret,
		Database:     conf.Mysql,
	})

	srv := oauthServer.Srv

	ti, err := srv.Manager.LoadAccessToken(context.Background(), bearerToken)
	if err != nil {
		resp.Error("身份验证失败！", nil)
		return
	}

	w.Header().Add("X-User-ID", ti.GetUserID())

	resp.Success("验证成功！", nil)
	return
}
