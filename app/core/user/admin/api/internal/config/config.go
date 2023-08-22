package config

import (
	"app/std/database"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Mysql database.Mysql
}
