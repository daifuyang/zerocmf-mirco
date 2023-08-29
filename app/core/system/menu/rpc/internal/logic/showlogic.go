package logic

import (
	"app/core/system/menu/model"
	"context"

	"app/core/system/menu/rpc/internal/svc"
	"app/core/system/menu/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jinzhu/copier"
)

type ShowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewShowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShowLogic {
	return &ShowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取单个菜单
func (l *ShowLogic) Show(in *pb.ShowReq) (*pb.ShowResp, error) {
	ctx := l.ctx
	adminMenuModel := l.svcCtx.AdminMenuModel
	one, err := adminMenuModel.FindOne(ctx, in.MenuId)
	if err != nil {
		return nil, err
	}
	menuId := one.Id

	// 获取对应的api列表
	adminMenuApi := l.svcCtx.AdminMenuApiModel
	var menus = make([]*model.AdminMenuApi, 0)
	menus, err = adminMenuApi.Where("menu_id = ?", menuId).Find(ctx)
	if err != nil && err != model.ErrNotFound {
		return nil, err
	}

	menu := pb.Menu{}
	err = copier.Copy(&menu, &one)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&menu.ApiList, &menus)
	if err != nil {
		return nil, err
	}

	return &pb.ShowResp{
		Success: true,
		Message: "获取成功！",
		Data:    &menu,
	}, nil
}
