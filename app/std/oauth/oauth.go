package oauth

import (
	"app/std/database"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-oauth2/mysql/v4"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	_ "github.com/go-sql-driver/mysql" // 导入 MySQL 驱动
	"golang.org/x/oauth2"
	"log"
)

type oauth struct {
	Config oauth2.Config
	Srv    *server.Server
	Store  *mysql.Store
}

type Config struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
	Database     database.Mysql `json:",optional"`
}

func NewConf(conf Config) oauth2.Config {
	authServerURL := conf.RedirectURL
	return oauth2.Config{
		ClientID:     conf.ClientID,
		ClientSecret: conf.ClientSecret,
		Scopes:       []string{"all"},
		RedirectURL:  authServerURL,
		Endpoint: oauth2.Endpoint{
			AuthURL:  authServerURL + "/api/authn/authorize",
			TokenURL: authServerURL + "/api/authn/token",
		},
	}
}

func NewServer(conf Config) *oauth {
	authServerURL := conf.RedirectURL
	oauthConfig := NewConf(conf)
	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)
	dsn := conf.Database.Dsn()
	mysqlStore := mysql.NewDefaultStore(
		mysql.NewConfig(dsn),
	)
	// token memory store
	manager.MapTokenStorage(mysqlStore)
	// generate jwt access token
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate("", []byte("00000000"), jwt.SigningMethodHS512))
	clientStore := store.NewClientStore()
	clientStore.Set(oauthConfig.ClientID, &models.Client{
		ID:     oauthConfig.ClientID,
		Secret: oauthConfig.ClientSecret,
		Domain: authServerURL,
	})

	manager.MapClientStorage(clientStore)
	srv := server.NewDefaultServer(manager)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	return &oauth{
		Config: oauthConfig,
		Srv:    srv,
		Store:  mysqlStore,
	}
}
