package logic

import (
	"app/core/system/menu/model"
	"context"
	"github.com/jinzhu/copier"
	"strings"
	"time"

	"app/core/system/menu/rpc/internal/svc"
	"app/core/system/menu/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveLogic {
	return &SaveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 保存菜单
func (l *SaveLogic) Save(in *pb.SaveReq) (*pb.SaveResp, error) {
	adminMenuModel := l.svcCtx.AdminMenuModel
	adminApiModel := l.svcCtx.AdminMenuApiModel
	ctx := l.ctx
	adminMenu := model.AdminMenu{}
	err := copier.Copy(&adminMenu, &in)
	if err != nil {
		return nil, err
	}
	id := adminMenu.Id
	if id == 0 {
		adminMenu.CreatedAt = time.Now()
		insert, insertErr := adminMenuModel.Insert(l.ctx, &adminMenu)
		if insertErr != nil {
			return nil, insertErr
		}
		id, err = insert.LastInsertId()
		if err != nil {
			return nil, err
		}
		adminMenu.Id = id
	} else {
		err = adminMenuModel.Update(l.ctx, &adminMenu)
		if err != nil {
			return nil, err
		}
	}

	//	 保存对应的api
	for _, api := range in.ApiList {

		name := api.Name
		if name == "" {
			name = in.Name
		}
		method := strings.ToLower(api.Method)

		apiFirst, apiOneErr := adminApiModel.Where("api = ? AND method = ?", api.Api, method).First(ctx)
		if apiOneErr != nil && apiOneErr != model.ErrNotFound {
			return nil, apiOneErr
		}

		var desc = ""
		if api.Desc != nil {
			desc = *api.Desc
		}

		menuApi := model.AdminMenuApi{
			UserId: 1,
			MenuId: id,
			Name:   name,
			Api:    api.Api,
			Method: method,
			Desc:   desc,
		}
		if apiFirst == nil {
			insert, insertErr := adminApiModel.Insert(ctx, &menuApi)
			if insertErr != nil {
				return nil, insertErr
			}
			menuApi.Id, err = insert.LastInsertId()
			if err != nil {
				return nil, err
			}
		} else {
			menuApi.Id = apiFirst.Id
			err = adminApiModel.Update(ctx, &menuApi)
			if err != nil {
				return nil, err
			}
		}
	}

	message := "新增成功！"
	if in.Id > 0 {
		message = "修改成功！"
	}

	return &pb.SaveResp{
		Success: true,
		Message: message,
		Data:    in,
	}, nil
}
