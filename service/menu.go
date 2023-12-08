package service

import (
	"SparkForge/db/dao"
	"SparkForge/pkg/ctl"
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
	}

	return ctl.SuccessWithDataResp(menuResp), nil
}
