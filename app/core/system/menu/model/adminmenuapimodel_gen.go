// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	adminMenuApiFieldNames          = builder.RawFieldNames(&AdminMenuApi{})
	adminMenuApiRows                = strings.Join(adminMenuApiFieldNames, ",")
	adminMenuApiRowsExpectAutoSet   = strings.Join(stringx.Remove(adminMenuApiFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	adminMenuApiRowsWithPlaceHolder = strings.Join(stringx.Remove(adminMenuApiFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheAdminMenuApiIdPrefix   = "cache:adminMenuApi:id:"
	cacheAdminMenuApiNamePrefix = "cache:adminMenuApi:name:"
)

type (
	adminMenuApiModel interface {
		Where(query string, args ...interface{}) *defaultAdminMenuApiModel
		Limit(limit int) *defaultAdminMenuApiModel
		Offset(offset int) *defaultAdminMenuApiModel
		OrderBy(query string) *defaultAdminMenuApiModel
		First(ctx context.Context) (*AdminMenuApi, error)
		Find(ctx context.Context) ([]*AdminMenuApi, error)
		Count(ctx context.Context) (int64, error)
		Insert(ctx context.Context, data *AdminMenuApi) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*AdminMenuApi, error)
		FindOneByName(ctx context.Context, name string) (*AdminMenuApi, error)
		Update(ctx context.Context, data *AdminMenuApi) error
		Delete(ctx context.Context, id int64) error
	}

	defaultAdminMenuApiModel struct {
		sqlc.CachedConn
		table     string
		query     string
		queryArgs []interface{}
		limit     int
		offset    int
		orderBy   string
	}

	AdminMenuApi struct {
		Id        int64     `db:"id"`
		UserId    int64     `db:"user_id"` // 创建人
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
		DeletedAt time.Time `db:"deleted_at"`
		MenuId    int64     `db:"menu_id"` // 关联菜单
		Name      string    `db:"name"`    // 接口名称
		Api       string    `db:"api"`     // api地址
		Method    string    `db:"method"`  // 请求方法
		Desc      string    `db:"desc"`    // 菜单描述
		Status    int64     `db:"status"`  // 1 =>启用,0 => 停用
	}
)

func newAdminMenuApiModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultAdminMenuApiModel {
	return &defaultAdminMenuApiModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`admin_menu_api`",
	}
}

func (m *defaultAdminMenuApiModel) withSession(session sqlx.Session) *defaultAdminMenuApiModel {
	return &defaultAdminMenuApiModel{
		CachedConn: m.CachedConn.WithSession(session),
		table:      "`admin_menu_api`",
	}
}

func (m *defaultAdminMenuApiModel) Where(query string, args ...interface{}) *defaultAdminMenuApiModel {
	m.query = query
	m.queryArgs = args
	return m
}

func (m *defaultAdminMenuApiModel) Limit(limit int) *defaultAdminMenuApiModel {
	m.limit = limit
	return m
}

func (m *defaultAdminMenuApiModel) Offset(offset int) *defaultAdminMenuApiModel {
	m.offset = offset
	return m
}

func (m *defaultAdminMenuApiModel) OrderBy(orderBy string) *defaultAdminMenuApiModel {
	m.orderBy = orderBy
	return m
}
func (m *defaultAdminMenuApiModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	adminMenuApiIdKey := fmt.Sprintf("%s%v", cacheAdminMenuApiIdPrefix, id)
	adminMenuApiNameKey := fmt.Sprintf("%s%v", cacheAdminMenuApiNamePrefix, data.Name)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (json sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, adminMenuApiIdKey, adminMenuApiNameKey)
	return err
}

func (m *defaultAdminMenuApiModel) FindOne(ctx context.Context, id int64) (*AdminMenuApi, error) {
	adminMenuApiIdKey := fmt.Sprintf("%s%v", cacheAdminMenuApiIdPrefix, id)
	var resp AdminMenuApi
	err := m.QueryRowCtx(ctx, &resp, adminMenuApiIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? AND deleted_at = 0 limit 1", adminMenuApiRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultAdminMenuApiModel) FindOneByName(ctx context.Context, name string) (*AdminMenuApi, error) {
	adminMenuApiNameKey := fmt.Sprintf("%s%v", cacheAdminMenuApiNamePrefix, name)
	var resp AdminMenuApi
	err := m.QueryRowIndexCtx(ctx, &resp, adminMenuApiNameKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `name` = ? AND deleted_at = 0 limit 1", adminMenuApiRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, name); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultAdminMenuApiModel) Insert(ctx context.Context, data *AdminMenuApi) (sql.Result, error) {
	adminMenuApiIdKey := fmt.Sprintf("%s%v", cacheAdminMenuApiIdPrefix, data.Id)
	adminMenuApiNameKey := fmt.Sprintf("%s%v", cacheAdminMenuApiNamePrefix, data.Name)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (json sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, adminMenuApiRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.DeletedAt, data.MenuId, data.Name, data.Api, data.Method, data.Desc, data.Status)
	}, adminMenuApiIdKey, adminMenuApiNameKey)
	return ret, err
}

func (m *defaultAdminMenuApiModel) Update(ctx context.Context, newData *AdminMenuApi) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	adminMenuApiIdKey := fmt.Sprintf("%s%v", cacheAdminMenuApiIdPrefix, data.Id)
	adminMenuApiNameKey := fmt.Sprintf("%s%v", cacheAdminMenuApiNamePrefix, data.Name)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (json sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, adminMenuApiRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.UserId, newData.DeletedAt, newData.MenuId, newData.Name, newData.Api, newData.Method, newData.Desc, newData.Status, newData.Id)
	}, adminMenuApiIdKey, adminMenuApiNameKey)
	return err
}

// 根据条件进行查询一条数据
func (m *defaultAdminMenuApiModel) First(ctx context.Context) (*AdminMenuApi, error) {
	query := m.query

	queryArgs := m.queryArgs
	orderBy := m.orderBy
	var resp AdminMenuApi
	sql := fmt.Sprintf("select %s from %s", adminMenuApiRows, m.table)

	if query != "" {
		sql += " where " + query
	}

	// 排序
	if orderBy != "" {
		sql += fmt.Sprintf(" ORDER BY %s", orderBy)
	}

	sql += " AND deleted_at = 0 limit 1"

	err := m.QueryRowNoCacheCtx(ctx, &resp, sql, queryArgs...)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// 根据条件进行列表查询
func (m *defaultAdminMenuApiModel) Find(ctx context.Context) ([]*AdminMenuApi, error) {

	query := m.query
	queryArgs := m.queryArgs
	orderBy := m.orderBy

	var resp []*AdminMenuApi
	sql := fmt.Sprintf("select %s from %s", adminMenuApiRows, m.table)

	if query != "" {
		sql += " where " + query + " AND deleted_at = 0"
	} else {
		sql += " where deleted_at = 0"
	}

	// 排序
	if orderBy != "" {
		sql += fmt.Sprintf(" ORDER BY %s", orderBy)
	}

	limit := m.limit
	offset := m.offset

	// 查询条件
	if limit > 0 {
		sql += fmt.Sprintf(" LIMIT %d", limit)
	}

	if offset > 0 {
		sql += fmt.Sprintf(" OFFSET %d", offset)
	}

	err := m.QueryRowsNoCacheCtx(ctx, &resp, sql, queryArgs...)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// 统计字段
func (m *defaultAdminMenuApiModel) Count(ctx context.Context) (int64, error) {
	query := m.query
	queryArgs := m.queryArgs
	sql := fmt.Sprintf("select count(`id`) from %s", m.table)
	if query != "" {
		sql += " where " + query
	}
	var resp int64
	err := m.QueryRowNoCacheCtx(ctx, &resp, sql, queryArgs...)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}
func (m *defaultAdminMenuApiModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheAdminMenuApiIdPrefix, primary)
}

func (m *defaultAdminMenuApiModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? AND deleted_at = 0 limit 1", adminMenuApiRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultAdminMenuApiModel) tableName() string {
	return m.table
}
