package model

import (
	"SparkForge/cache"
	"SparkForge/pkg/util"
	"golang.org/x/crypto/bcrypt"
	"strconv"

	"github.com/jinzhu/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	PasswordDigest string
	Kitchen        string
}

const (
	PassWordCost = 12
)

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}

// GetCount 得到当天的合成次数
func (user *User) GetCount() uint64 {
	result, err := cache.RedisClient.Get(cache.UserCountKey(user.ID)).Result()
	if err != nil {
		util.LogrusObj.Infoln(err)
	}

	count, err := strconv.ParseUint(result, 10, 64)
	if err != nil {
		util.LogrusObj.Infoln(err)
	}

	return count
}

// AddCount 添加次数
func (user *User) AddCount() error {
	return cache.RedisClient.Incr(cache.UserCountKey(user.ID)).Err()
}
