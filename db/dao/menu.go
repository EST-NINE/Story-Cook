package dao

import (
	"context"

	"gorm.io/gorm"

	"SparkForge/db/model"
)

type MenuDao struct {
	*gorm.DB
}

func NewMenuDao(c context.Context) *MenuDao {
	if c == nil {
		c = context.Background()
	}
	return &MenuDao{NewDBClient(c)}
}

// SelectMenu 判断是否为彩蛋
func (dao *MenuDao) SelectMenu(keywords string) (menu model.Menu, err error) {
	err = dao.DB.Model(&model.Menu{}).Where("keywords = ?", keywords).First(&menu).Error
	return
}

// FindMenuByKeywords 根据关键词查询彩蛋
func (dao *MenuDao) FindMenuByKeywords(keywords string) (menu *model.Menu, err error) {
	err = dao.DB.Model(&model.Menu{}).Where("keywords = ?", keywords).First(&menu).Error
	return
}

// CreateUserMenu 创建成就
func (dao *MenuDao) CreateUserMenu(userMenu *model.UserMenu) error {
	return dao.DB.Model(&model.UserMenu{}).Create(&userMenu).Error
}

// FindUserMenuByKeywordsAndUserId 根据关键词和用户id查找成就
func (dao *MenuDao) FindUserMenuByKeywordsAndUserId(uid uint, keywords string) (userMenu *model.UserMenu, err error) {
	err = dao.DB.Model(&model.UserMenu{}).Where("uid = ? AND keywords = ? ", uid, keywords).First(&userMenu).Error
	return
}

// ListUserMenu 得到成就列表
func (dao *MenuDao) ListUserMenu(page, limit int, uid uint) (userMenus []model.UserMenu, total int64, err error) {
	err = dao.DB.Model(&model.UserMenu{}).Preload("User").Where("uid = ?", uid).
		Count(&total).
		Order("created_at DESC").
		Limit(limit).Offset((page - 1) * limit).
		Find(&userMenus).Error
	return
}
