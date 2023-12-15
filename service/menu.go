package service

import (
	"SparkForge/repository/db/dao"
	"SparkForge/repository/db/model"
	"context"
	"errors"

	"gorm.io/gorm"

	"SparkForge/pkg/controller"
	"SparkForge/pkg/util"
	"SparkForge/types"
)

type MenuSrv struct {
}

// SelectMenu 判断是否是彩蛋
func (s *MenuSrv) SelectMenu(c context.Context, req *types.SelectMenuReq) (resp *types.MenuResp, err error) {
	menuDao := dao.NewMenuDao(c)
	menu, err := menuDao.SelectMenu(req.Keywords)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("没有触发彩蛋哦")
		}
		return
	}

	return &types.MenuResp{
		ID:       menu.ID,
		Keywords: menu.Keywords,
		Content:  menu.Content,
		CreateAt: menu.CreatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

// CreateUserMenu 添加彩蛋用户成就
func (s *MenuSrv) CreateUserMenu(c context.Context, req *types.CreateUserMenuReq) error {
	userInfo, err := controller.GetUserInfo(c)
	if err != nil {
		util.LogrusObj.Infoln(err)
		return err
	}

	user, err := dao.NewUserDao(c).FindUserByUserId(userInfo.Id)
	if err != nil {
		util.LogrusObj.Infoln(err)
		return err
	}

	menuDao := dao.NewMenuDao(c)
	_, err = menuDao.FindUserMenuByKeywordsAndUserId(userInfo.Id, req.Keywords)
	if err == nil {
		err = errors.New("已经添加过这个成就了哦")
		return err
	}

	userMenu := model.UserMenu{
		User:     *user,
		Keywords: req.Keywords,
		Content:  req.Content,
	}

	err = menuDao.CreateUserMenu(&userMenu)
	if err != nil {
		util.LogrusObj.Infoln(err)
		return err
	}

	return nil
}

// ListUserMenu 得到对应用户的彩蛋成就列表
func (s *MenuSrv) ListUserMenu(c context.Context, req *types.ListUserMenuReq) (resp []*types.MenuResp, total int64, err error) {
	if req.Limit == 0 {
		req.Limit = 15
	}

	userInfo, err := controller.GetUserInfo(c)
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

	return listUserMenuResp, total, nil
}
