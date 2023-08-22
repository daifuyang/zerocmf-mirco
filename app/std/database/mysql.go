package database

import "fmt"

type Mysql struct {
	Username  string
	Password  string
	Host      string
	Port      int64
	DbName    string
	Charset   string
	Collate   string
	ParseTime bool
	AuthCode  string `json:",optional"`
}

type Option func(*Mysql)

func (c *Mysql) WithSimple() Option {
	return func(c *Mysql) {
		c.DbName = ""
	}
}

func (c *Mysql) Dsn(options ...Option) string {
	for _, option := range options {
		option(c)
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t",
		c.Username, c.Password, c.Host, c.Port, c.DbName, c.Charset, c.ParseTime)
}
