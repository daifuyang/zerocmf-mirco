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
	"time"
)

type ServiceContext struct {
	Config         config.Config
	AdminMenuModel model.AdminMenuModel
}

type adminMenu struct {
	MenuType   int64       `json:"menuType"`
	FormId     string      `json:"form_id"`
	Name       string      `json:"name"`
	Path       string      `json:"path"`
	Component  string      `json:"component"`
	Icon       string      `json:"icon"`
	Order      float64     `json:"order"`
	Access     string      `json:"access"`
	Link       int64       `json:"link"`
	HideInMenu int64       `json:"hideInMenu"`
	Status     *int64      `json:"status"`
	Routes     []adminMenu `json:"routes"`
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
		recursionAddMenu(menus, 0, adminMenuModel)
	})

	return &ServiceContext{
		Config:         c,
		AdminMenuModel: adminMenuModel,
	}
}

func recursionAddMenu(menus []adminMenu, parentId int64, db model.AdminMenuModel) {
	// 增加当前层级
	for _, v := range menus {

		var status int64 = 1
		if v.Status != nil {
			status = *v.Status
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
			Order:      v.Order,
			Status:     status,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}

		hasRoutes := len(v.Routes) > 0

		if hasRoutes {
			menu.MenuType = 0
		}

		// 保存菜单
		var id int64 = 0
		if v.Name != "" {
			first, err := db.Where("name = ?", menu.Name).First(context.Background())
			if err != nil && err != model.ErrNotFound {
				return
			}
			if first == nil {
				var insert sql.Result
				insert, err = db.Insert(context.Background(), &menu)
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
				err = db.Update(context.Background(), &menu)
				if err != nil {
					return
				}
			}
		}

		if hasRoutes {
			recursionAddMenu(v.Routes, id, db)
		}

	}
}
