package config

import (
	"app/std/apisix"
	"app/std/database"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Mysql database.Mysql
	Auth  struct {
		AccessSecret string
	}
	Apisix apisix.Apisix
}
