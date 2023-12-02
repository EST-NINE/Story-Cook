package service

import (
	"SparkForge/pkg/ctl"
	"SparkForge/pkg/util"
	"SparkForge/repository/db/dao"
	"SparkForge/repository/db/model"
	"SparkForge/types"
	"context"
	"errors"
	"gorm.io/gorm"
	"sync"
)

var UserSrvIns *UserSrv
var UserSrvOnce sync.Once

type UserSrv struct {
}

func GetUserSrv() *UserSrv {
	UserSrvOnce.Do(func() {
		UserSrvIns = &UserSrv{}
	})
	return UserSrvIns
}

func (s *UserSrv) Register(ctx context.Context, req *types.UserServiceReq) (resp interface{}, err error) {
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.FindUserByUserName(req.UserName)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			user = &model.User{
				UserName: req.UserName,
			}
			// 密码加密存储
			if err = user.SetPassword(req.Password); err != nil {
				util.LogrusObj.Info(err)
				return
			}

			if err = userDao.CreateUser(user); err != nil {
				util.LogrusObj.Info(err)
				return
			}

			return ctl.SuccessResp(), nil
		}
		return nil, err
	}
	return nil, errors.New("用户已存在")
}

// Login 用户登陆函数
func (s *UserSrv) Login(ctx context.Context, req *types.UserServiceReq) (resp interface{}, err error) {
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.FindUserByUserName(req.UserName)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("用户不存在")
		return
	}

	if !user.CheckPassword(req.Password) {
		err = errors.New("账号/密码错误")
		util.LogrusObj.Info(err)
		return
	}

	token, err := util.GenerateToken(user.ID, req.UserName)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}

	userResp := &types.UserResp{
		ID:       user.ID,
		UserName: user.UserName,
		CreateAt: user.CreatedAt.Unix(),
	}
	uResp := &types.TokenDataResp{
		User:  userResp,
		Token: token,
	}

	return ctl.SuccessWithDataResp(uResp), nil
}
