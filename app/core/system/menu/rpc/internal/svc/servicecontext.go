package svc

import (
	"app/core/system/menu/model"
	"app/core/system/menu/rpc/internal/config"
	"app/std/util"
	"context"
	"database/sql"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"os"
	"strings"
	"time"
)

type ServiceContext struct {
	Config            config.Config
	AdminMenuModel    model.AdminMenuModel
	AdminMenuApiModel model.AdminMenuApiModel
}

type adminMenuApi struct {
	MenuName string `json:"menu_name"`
	Name     string `json:"name"`
	Api      string `json:"api"`
	Method   string `json:"method"`
	Desc     string `json:"desc"`
	Status   *int64 `json:"status"`
}

type adminMenu struct {
	MenuType   int64          `json:"menuType"`
	FormId     string         `json:"form_id"`
	Name       string         `json:"name"`
	Path       string         `json:"path"`
	Component  string         `json:"component"`
	Icon       string         `json:"icon"`
	Order      *float64       `json:"order"`
	Access     string         `json:"access"`
	Link       int64          `json:"link"`
	HideInMenu int64          `json:"hideInMenu"`
	Status     *int64         `json:"status"`
	Routes     []adminMenu    `json:"routes"`
	Api        []adminMenuApi `json:"api"`
}

func NewServiceContext(c config.Config) *ServiceContext {
	adminMenuModel := model.NewAdminMenuModel(sqlx.NewMysql(c.Mysql.Dsn()), c.Cache)
	util.Install(func() {
		util.InitDb(c.Mysql)
		var menus []adminMenu
		bytes, err := os.ReadFile("data/menu.json")
		if err != nil {
			panic(err.Error())
		}
		err = json.Unmarshal(bytes, &menus)
		recursionAddMenu(menus, 0, c)
	}, true)

	return &ServiceContext{
		Config:            c,
		AdminMenuModel:    adminMenuModel,
		AdminMenuApiModel: model.NewAdminMenuApiModel(sqlx.NewMysql(c.Mysql.Dsn()), c.Cache),
	}
}

func recursionAddMenu(menus []adminMenu, parentId int64, c config.Config) {
	adminMenuModel := model.NewAdminMenuModel(sqlx.NewMysql(c.Mysql.Dsn()), c.Cache)
	adminApiModel := model.NewAdminMenuApiModel(sqlx.NewMysql(c.Mysql.Dsn()), c.Cache)
	// 增加当前层级
	for _, v := range menus {

		var status int64 = 1
		if v.Status != nil {
			status = *v.Status
		}

		var order float64 = 10000
		if v.Order != nil {
			order = *v.Order
		}

		menu := model.AdminMenu{
			ParentId:   parentId,
			FormId:     v.FormId,
			UserId:     1,
			Name:       v.Name,
			Path:       v.Path,
			Component:  v.Component,
			MenuType:   v.MenuType,
			HideInMenu: v.HideInMenu,
			Access:     v.Access,
			Link:       v.Link,
			Order:      order,
			Status:     status,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}

		hasRoutes := len(v.Routes) > 0

		if hasRoutes {
			menu.MenuType = 0
		}

		ctx := context.Background()

		// 保存菜单
		var id int64 = 0
		if v.Name != "" {
			first, err := adminMenuModel.Where("name = ?", menu.Name).First(ctx)
			if err != nil && err != model.ErrNotFound {
				return
			}
			if first == nil {
				var insert sql.Result
				insert, err = adminMenuModel.Insert(ctx, &menu)
				if err != nil {
					return
				}
				id, err = insert.LastInsertId()
				if err != nil {
					return
				}
			} else {
				id = first.Id
				menu.Id = id
				err = adminMenuModel.Update(ctx, &menu)
				if err != nil {
					return
				}
			}

			//	 保存对应的api
			for _, api := range v.Api {

				name := api.Name
				if name == "" {
					name = v.Name
				}
				method := strings.ToLower(api.Method)

				apiFirst, apiOneErr := adminApiModel.Where("api = ? AND method = ?", api.Api, method).First(ctx)
				if apiOneErr != nil && apiOneErr != model.ErrNotFound {
					return
				}
				status = 1
				if api.Status != nil {
					status = *api.Status
				}

				menuApi := model.AdminMenuApi{
					UserId: 1,
					MenuId: id,
					Name:   name,
					Api:    api.Api,
					Method: method,
					Desc:   api.Desc,
					Status: status,
				}
				if apiFirst == nil {
					_, insertErr := adminApiModel.Insert(ctx, &menuApi)
					if insertErr != nil {
						return
					}
				} else {
					menuApi.Id = apiFirst.Id
					err = adminApiModel.Update(ctx, &menuApi)
					if err != nil {
						return
					}
				}
			}
		}

		if hasRoutes {
			recursionAddMenu(v.Routes, id, c)
		}

	}
}
