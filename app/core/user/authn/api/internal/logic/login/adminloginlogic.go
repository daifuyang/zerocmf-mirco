package login

import (
	userAdmin "app/core/user/admin/rpc/service"
	"app/std/apisix"
	"app/std/apisix/plugins/authentication"
	"app/std/net/http/logic"
	"app/std/oauth"
	"context"
	"fmt"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"time"

	"app/core/user/authn/api/internal/svc"
	"app/core/user/authn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminLoginLogic struct {
	logx.Logger
	ctx    context.Context
	header *http.Request
	svcCtx *svc.ServiceContext
}

func NewAdminLoginLogic(header *http.Request, svcCtx *svc.ServiceContext) *AdminLoginLogic {
	ctx := header.Context()
	return &AdminLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		header: header,
		svcCtx: svcCtx,
	}
}

func (l *AdminLoginLogic) AdminLogin(req *types.LoginReq) (resp logic.Response) {

	conf := l.svcCtx.Config

	// 获取当前登录的用户信息
	user, err := l.svcCtx.UserAdminRpc.AdminShow(l.ctx, &userAdmin.AdminShowReq{
		Username: req.Username,
	})
	if err != nil {
		resp.Error("管理员服务系统错误！", err.Error())
	}

	// 密码比对
	password := user.Salt + req.Password
	hashedErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if hashedErr != nil {
		resp.Error("用户名或密码错误！", nil)
		return
	}

	// 删除原来的派发token
	conn := sqlx.NewMysql(conf.Mysql.Dsn())

	likeStr := fmt.Sprintf("'%%\"ClientID\":\"%s\",\"UserID\":\"%d\"%%'", conf.Oauth.ClientID, user.UserId)
	sqlStr := "DELETE from oauth2_token WHERE `data` LIKE " + likeStr
	conn.ExecCtx(l.ctx, sqlStr)

	redirectURL := "http://" + conf.Apisix.Host + ":" + strconv.Itoa(conf.Port)

	oauthServer := oauth.NewServer(oauth.Config{
		ClientID:     conf.Oauth.ClientID,
		ClientSecret: conf.Oauth.ClientSecret,
		RedirectURL:  redirectURL,
		AccessSecret: conf.Auth.AccessSecret,
		Database:     conf.Mysql,
	})
	defer oauthServer.Store.Close()
	srv := oauthServer.Srv
	oauthConfig := oauthServer.Config
	exp := 168

	authReq := &server.AuthorizeRequest{
		RedirectURI:    redirectURL,
		ResponseType:   "code",
		ClientID:       oauthConfig.ClientID,
		State:          "jwt",
		Scope:          "all",
		UserID:         strconv.FormatInt(user.UserId, 10),
		AccessTokenExp: time.Duration(exp) * time.Hour,
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

	// 生成消费者供apisix public-ley消费
	key := strconv.FormatInt(user.UserId, 10)
	var expInt = int64(exp) * 3600

	apisix.NewConsumer(conf.Apisix.ApiKey, conf.Apisix.Host).Add(key, apisix.WithJwtAuth(authentication.JwtAuth{
		Key:    key,
		Secret: conf.Auth.AccessSecret,
		Exp:    expInt,
	}))

	resp.Success("获取成功！", token)
	return
}
