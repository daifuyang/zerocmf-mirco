package config

import (
	"app/std/database"
	"app/std/oauth"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Mysql database.Mysql
	Oauth oauth.Config
}
