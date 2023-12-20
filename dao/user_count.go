package dao

import (
	"context"
	"gorm.io/gorm"
	"story-cook-be/model"
)

type UserCountDao struct {
	*gorm.DB
}

func NewUseCountDao(c context.Context) *UserCountDao {
	if c == nil {
		c = context.Background()
	}
	return &UserCountDao{NewDBClient(c)}
}

// CreateUserCount 创建次数
func (dao *UserCountDao) CreateUserCount(userCount *model.UserCount) error {
	return dao.DB.Model(&model.UserCount{}).Create(userCount).Error
}

// UpdateUserCount 添加次数
func (dao *UserCountDao) UpdateUserCount(userCount *model.UserCount) error {
	return dao.DB.Model(&model.UserCount{}).Where("uid = ? AND date = ?", userCount.UID, userCount.Date).
		Updates(&userCount).Error
}

// FindUserCountByDate 根据日期查找有无该条用户次数记录
func (dao *UserCountDao) FindUserCountByDate(uid uint, date string) (userCount *model.UserCount, err error) {
	err = dao.DB.Model(&model.UserCount{}).Where("uid = ? AND date = ?", uid, date).
		First(&userCount).Error

	return
}

// FindUserCountByID 根据id查找有无该条用户次数记录
func (dao *UserCountDao) FindUserCountByID(uid uint) (userCount *model.UserCount, err error) {
	err = dao.DB.Model(&model.UserCount{}).Where("uid = ?", uid).
		First(&userCount).Error

	return
}
