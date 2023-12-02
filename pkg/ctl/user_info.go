package ctl

import (
	"context"
	"errors"
)

type key int

var userKey key

type UserInfo struct {
	Id uint `json:"id"`
}

func GetUserInfo(c context.Context) (*UserInfo, error) {
	user, ok := FromContext(c) // 从上下文中获取用户信息
	if !ok {
		return nil, errors.New("获取用户信息错误")
	}
	return user, nil
}

func NewContext(c context.Context, u *UserInfo) context.Context {
	return context.WithValue(c, userKey, u) // 使用用户信息创建一个新的上下文
}

func FromContext(c context.Context) (*UserInfo, bool) {
	user, ok := c.Value(userKey).(*UserInfo) // 从上下文中获取用户信息
	return user, ok
}