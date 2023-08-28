package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SystemLogModel = (*customSystemLogModel)(nil)

type (
	// SystemLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSystemLogModel.
	SystemLogModel interface {
		systemLogModel
	}

	customSystemLogModel struct {
		*defaultSystemLogModel
	}
)

// NewSystemLogModel returns a model for the database table.
func NewSystemLogModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) SystemLogModel {
	return &customSystemLogModel{
		defaultSystemLogModel: newSystemLogModel(conn, c, opts...),
	}
}

// NewSystemLogSessionModel returns a model for the Session database table.
func NewSystemLogSessionModel(conn sqlx.SqlConn, session sqlx.Session, c cache.CacheConf, opts ...cache.Option) SystemLogModel {
	return &customSystemLogModel{
		defaultSystemLogModel: newSystemLogModel(conn, c, opts...).withSession(session),
	}
}
