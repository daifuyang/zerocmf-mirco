package oauth

import (
	"app/std/net/http/logic"
	"app/std/oauth"
	"context"
	"fmt"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"net/http"
	"strconv"
	"time"

	"app/core/user/authn/api/internal/svc"
	"app/core/user/authn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(header *http.Request, svcCtx *svc.ServiceContext) *LoginLogic {
	ctx := header.Context()
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp logic.Response) {

	// 删除原来的派发token
	conn := sqlx.NewMysql(l.svcCtx.Config.Mysql.Dsn())
	conf := l.svcCtx.Config
	likeStr := fmt.Sprintf("'%%\"ClientID\":\"%s\",\"UserID\":\"%d\"%%'", conf.Oauth.ClientID, req.UserId)
	sqlStr := "DELETE from oauth2_token WHERE `data` LIKE " + likeStr
	conn.ExecCtx(l.ctx, sqlStr)

	oauthServer := oauth.NewServer(oauth.Config{
		ClientID:     conf.Oauth.ClientID,
		ClientSecret: conf.Oauth.ClientSecret,
		RedirectURL:  conf.Oauth.RedirectURL,
		Database:     conf.Mysql,
	})
	defer oauthServer.Store.Close()
	srv := oauthServer.Srv
	oauthConfig := oauthServer.Config
	duration := time.Duration(168) * time.Hour

	authReq := &server.AuthorizeRequest{
		RedirectURI:    oauthConfig.RedirectURL,
		ResponseType:   "code",
		ClientID:       oauthConfig.ClientID,
		State:          "jwt",
		Scope:          "all",
		UserID:         strconv.FormatInt(req.UserId, 10),
		AccessTokenExp: duration,
		Request:        l.header,
	}

	ti, err := srv.GetAuthorizeToken(l.ctx, authReq)
	if err != nil {
		resp.Error("系统出错了", err.Error())
		return
	}

	code := ti.GetCode()

	token, tokenErr := oauthConfig.Exchange(context.Background(), code)

	if tokenErr != nil {
		resp.Error("系统出错了", tokenErr.Error())
		return
	}
	resp.Success("获取成功！", token)
	return
}
