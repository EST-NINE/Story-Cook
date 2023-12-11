package controller

import (
	"context"
	"errors"
)

type key int

var userKey key

type UserInfo struct {
	Id uint `json:"id"`
}

// GetUserInfo 从上下文中获取用户信息
func GetUserInfo(c context.Context) (*UserInfo, error) {
	userInfo, ok := c.Value(userKey).(*UserInfo) // 从上下文中获取用户信息
	if !ok {
		return nil, errors.New("获取用户信息错误")
	}
	return userInfo, nil
}

func NewContext(c context.Context, userInfo *UserInfo) context.Context {
	return context.WithValue(c, userKey, userInfo) // 使用用户信息创建一个新的上下文
}
