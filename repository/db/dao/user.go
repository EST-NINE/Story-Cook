package dao

import (
	"SparkForge/repository/db/model"
	"context"

	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(c context.Context) *UserDao {
	if c == nil {
		c = context.Background()
	}
	return &UserDao{NewDBClient(c)}
}

// CreateUser 创建User
func (dao *UserDao) CreateUser(user *model.User) error {
	return dao.DB.Model(&model.User{}).Create(user).Error
}

// UpdateUserById 根据 id 更新用户信息
func (dao *UserDao) UpdateUserById(uId uint, user *model.User) error {
	return dao.DB.Model(&model.User{}).Where("id=?", uId).
		Updates(&user).Error
}

// FindUserByUserName 根据用户名找到用户
func (dao *UserDao) FindUserByUserName(userName string) (user *model.User, err error) {
	err = dao.DB.Model(&model.User{}).Where("user_name=?", userName).
		First(&user).Error

	return
}

// FindUserByUserId 根据用户id找到用户
func (dao *UserDao) FindUserByUserId(id uint) (user *model.User, err error) {
	err = dao.DB.Model(&model.User{}).Where("id=?", id).
		First(&user).Error

	return
}
