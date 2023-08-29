package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AdminMenuApiModel = (*customAdminMenuApiModel)(nil)

type (
	// AdminMenuApiModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAdminMenuApiModel.
	AdminMenuApiModel interface {
		adminMenuApiModel
	}

	customAdminMenuApiModel struct {
		*defaultAdminMenuApiModel
	}
)

// NewAdminMenuApiModel returns a model for the database table.
func NewAdminMenuApiModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) AdminMenuApiModel {
	return &customAdminMenuApiModel{
		defaultAdminMenuApiModel: newAdminMenuApiModel(conn, c, opts...),
	}
}

// NewAdminMenuApiSessionModel returns a model for the Session database table.
func NewAdminMenuApiSessionModel(conn sqlx.SqlConn, session sqlx.Session, c cache.CacheConf, opts ...cache.Option) AdminMenuApiModel {
	return &customAdminMenuApiModel{
		defaultAdminMenuApiModel: newAdminMenuApiModel(conn, c, opts...).withSession(session),
	}
}
