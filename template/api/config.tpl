package config

import (
"app/std/database"
{{.authImport}}
)

type Config struct {
	rest.RestConf
	Mysql database.Mysql
	{{.auth}}
	{{.jwtTrans}}
}
