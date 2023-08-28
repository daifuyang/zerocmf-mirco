package logic

import (
	"app/core/system/menu/model"
	"app/core/system/menu/rpc/internal/svc"
	"app/core/system/menu/rpc/pb"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type TreeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTreeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TreeListLogic {
	return &TreeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取菜单
func (l *TreeListLogic) TreeList(in *pb.TreeListReq) (*pb.TreeListResp, error) {
	c := l.svcCtx
	adminMenuModel := c.AdminMenuModel
	menus, err := adminMenuModel.Find(l.ctx)
	if err != nil && model.ErrNotFound != err {
		return nil, err
	}

	treeList := recursionMenu(menus, 0)
	return &pb.TreeListResp{
		Success: true,
		Message: "获取成功！",
		Data:    treeList,
	}, nil
}

/**
 * @Author return <1140444693@qq.com>
 * @Description 递归增加子菜单项
 * @Date 2021/11/30 12:50:24
 * @Param
 * @return
 **/

func recursionMenu(menus []*model.AdminMenu, parentId int64) []*pb.TreeList {
	var treeList = make([]*pb.TreeList, 0)
	for _, v := range menus {
		if parentId == v.ParentId {
			menu := &pb.TreeList{
				Id:         v.Id,
				Name:       v.Name,
				Icon:       v.Icon,
				MenuType:   v.MenuType,
				Path:       v.Path,
				Component:  v.Component,
				Access:     v.Access,
				HideInMenu: v.HideInMenu,
				Order:      v.Order,
				CreatedAt:  v.CreatedAt.Unix(),
				CreateTime: v.CreatedAt.Format(time.DateTime),
				Status:     v.Status,
			}

			routes := recursionMenu(menus, v.Id)
			child := make([]*pb.TreeList, len(routes))
			for ri, rv := range routes {
				child[ri] = rv
			}
			menu.Routes = child
			treeList = append(treeList, menu)
		}
	}
	return treeList
}
