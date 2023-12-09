package service

import (
	"SparkForge/db/dao"
	"SparkForge/db/model"
	"SparkForge/pkg/ctl"
	"SparkForge/pkg/util"
	"SparkForge/types"
	"context"
	"errors"
	"gorm.io/gorm"
	"sync"
)

type MenuSrv struct {
}

var MenuSrvIns *MenuSrv
var MenuSrvOnce sync.Once

func GetMenuSrv() *MenuSrv {
	MenuSrvOnce.Do(func() {
		MenuSrvIns = &MenuSrv{}
	})
	return MenuSrvIns
}

// SelectMenu 判断是否是彩蛋
func (s *MenuSrv) SelectMenu(c context.Context, req *types.SelectMenuReq) (resp interface{}, err error) {
	menuDao := dao.NewMenuDao(c)
	menu, err := menuDao.SelectMenu(req.Keywords)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("没有触发彩蛋哦")
		}
		return
	}

	menuResp := &types.MenuResp{
		ID:       menu.ID,
		Keywords: menu.Keywords,
		Content:  menu.Content,
		CreateAt: menu.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	return ctl.SuccessWithDataResp(menuResp), nil
}

// CreateUserMenu 添加用户成就
func (s *MenuSrv) CreateUserMenu(c context.Context, req *types.CreateUserMenuReq) (resp interface{}, err error) {
	userInfo, err := ctl.GetUserInfo(c)
	if err != nil {
		util.LogrusObj.Infoln(err)
		return
	}

	user, err := dao.NewUserDao(c).FindUserByUserId(userInfo.Id)
	if err != nil {
		util.LogrusObj.Infoln(err)
		return
	}

	menuDao := dao.NewMenuDao(c)
	_, err = menuDao.FindUserMenuByKeywordsAndUserId(userInfo.Id, req.Keywords)
	if err == nil {
		err = errors.New("已经添加过这个成就了哦")
		return
	}

	userMenu := model.UserMenu{
		User:     *user,
		Keywords: req.Keywords,
		Content:  req.Content,
	}

	err = menuDao.CreateUserMenu(&userMenu)
	if err != nil {
		util.LogrusObj.Infoln(err)
		return
	}

	return ctl.SuccessResp(), nil
}

// ListUserMenu 得到对应用户的故事
func (s *MenuSrv) ListUserMenu(c context.Context, req *types.ListUserMenuReq) (resp interface{}, err error) {
	userInfo, err := ctl.GetUserInfo(c)
	if err != nil {
		util.LogrusObj.Infoln(err)
		return
	}

	userMenus, total, err := dao.NewMenuDao(c).ListUserMenu(req.Page, req.Limit, userInfo.Id)
	if err != nil {
		util.LogrusObj.Infoln(err)
		return
	}

	listUserMenuResp := make([]*types.MenuResp, 0)
	for _, userMenu := range userMenus {
		listUserMenuResp = append(listUserMenuResp, &types.MenuResp{
			ID:       userMenu.ID,
			Keywords: userMenu.Keywords,
			Content:  userMenu.Content,
			CreateAt: userMenu.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return ctl.ListResp(listUserMenuResp, total), nil
}
