package svc

import (
	log "app/core/log/rpc/service"
	"app/core/user/admin/model"
	"app/core/user/admin/rpc/internal/config"
	"app/std/util"
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"golang.org/x/crypto/bcrypt"
	"runtime"
	"strings"
)

type ServiceContext struct {
	Config         config.Config
	AdminUserModel model.AdminUserModel
	LogRpc         log.Service
	LogError       func(ctx context.Context, err error)
}

func NewServiceContext(c config.Config) *ServiceContext {
	dsn := c.Mysql.Dsn()
	adminModel := model.NewAdminUserModel(sqlx.NewMysql(dsn), c.Cache)
	util.Install(c.Mysql, func() {

		adminPassword := c.AdminPassword
		if strings.TrimSpace(adminPassword) == "" {
			adminPassword = "123456"
		}

		//	首次运行
		ctx := context.Background()
		one, oneErr := adminModel.FindOne(ctx, 1)
		if oneErr != nil && !errors.Is(oneErr, model.ErrNotFound) {
			panic("创建管理员失败！")
		}

		// 新建管理员
		if one == nil {
			salt := c.Mysql.AuthCode
			hashedPassword, hashedErr := bcrypt.GenerateFromPassword([]byte(salt+adminPassword), bcrypt.DefaultCost)
			if hashedErr != nil {
				panic("哈希密码错误")
			}
			_, insertErr := adminModel.Insert(ctx, &model.AdminUser{
				Username: "admin",
				Password: string(hashedPassword),
				Salt:     salt,
			})
			if insertErr != nil {
				panic("创建管理员失败！")
			}
		}
	})

	return &ServiceContext{
		Config: c,
		LogError: func(ctx context.Context, err error) {
			pc, file, line, _ := runtime.Caller(1)
			funcName := runtime.FuncForPC(pc).Name()
			method := fmt.Sprintf("Error occurred in function %s at %s:%d\n", funcName, file, line)
			logRpc := log.NewService(zrpc.MustNewClient(c.LogRpc))
			logRpc.LogSystem(ctx, &log.SystemLogReq{
				AppName:  c.Name,
				LogLevel: "error",
				Method:   method,
				Message:  err.Error(),
			})
		},
		AdminUserModel: adminModel,
	}
}
