package service

import (
	"errors"
	"story-cook-be/dao"
	"story-cook-be/model"
	"story-cook-be/pkg/response"
	"story-cook-be/pkg/util"
	"story-cook-be/types"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MenuSrv struct {
}

// SelectMenu 判断是否是彩蛋
func (s *MenuSrv) SelectMenu(ctx *gin.Context, req *types.SelectMenuReq) (resp *response.MenuResp, err error) {
	menuDao := dao.NewMenuDao(ctx)
	menu, err := menuDao.SelectMenu(req.Keywords)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("没有触发彩蛋哦")
		}
		return
	}

	return response.BuildMenuResp(menu), nil
}

// CreateUserMenu 添加彩蛋用户成就
func (s *MenuSrv) CreateUserMenu(ctx *gin.Context, req *types.CreateUserMenuReq) error {
	claims, _ := ctx.Get("claims")
	userInfo := claims.(*util.Claims)

	user, err := dao.NewUserDao(ctx).FindUserByUserId(userInfo.Id)
	if err != nil {
		util.LogrusObj.Infoln(err)
		return err
	}

	menuDao := dao.NewMenuDao(ctx)
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
func (s *MenuSrv) ListUserMenu(ctx *gin.Context, req *types.ListUserMenuReq) (resp []*response.MenuResp, total int64, err error) {
	claims, _ := ctx.Get("claims")
	userInfo := claims.(*util.Claims)

	userMenus, total, err := dao.NewMenuDao(ctx).ListUserMenu(req.Page, req.Limit, userInfo.Id)
	if err != nil {
		util.LogrusObj.Infoln(err)
		return
	}

	listUserMenuResp := make([]*response.MenuResp, 0)
	for _, userMenu := range userMenus {
		listUserMenuResp = append(listUserMenuResp, response.BuildUserMenuResp(userMenu))
	}

	return listUserMenuResp, total, nil
}
