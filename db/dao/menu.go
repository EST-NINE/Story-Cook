package dao

import (
	"SparkForge/db/model"
	"context"
	"gorm.io/gorm"
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
