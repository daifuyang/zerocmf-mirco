package oauth

import (
	"app/std/net/http/logic"
	"app/std/oauth"
	"context"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"strconv"

	"app/core/user/authn/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type TokenLogic struct {
	logx.Logger
	ctx    context.Context
	writer http.ResponseWriter
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewTokenLogic(writer http.ResponseWriter, header *http.Request, svcCtx *svc.ServiceContext) *TokenLogic {
	ctx := header.Context()
	return &TokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		writer: writer,
		svcCtx: svcCtx,
	}
}

func (l *TokenLogic) Token() (resp logic.Response) {
	c := l.svcCtx
	conf := c.Config
	w := l.writer
	r := l.header

	redirectURL := "http://" + conf.Apisix.Host + ":" + strconv.Itoa(conf.Port)

	oauthConf := oauth.Config{
		ClientID:     conf.Oauth.ClientID,
		ClientSecret: conf.Oauth.ClientSecret,
		RedirectURL:  redirectURL,
		AccessSecret: conf.Auth.AccessSecret,
		Database:     conf.Mysql,
	}

	oauthServer := oauth.NewServer(oauthConf)
	defer oauthServer.Store.Close()
	srv := oauthServer.Srv

	err := srv.HandleTokenRequest(w, r)
	if err != nil {
		httpx.OkJsonCtx(r.Context(), w, resp)
	}
	return
}
