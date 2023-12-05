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

// Register 注册用户
func (s *UserSrv) Register(c context.Context, req *types.UserServiceReq) (resp interface{}, err error) {
	userDao := dao.NewUserDao(c)
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

	err = errors.New("用户已存在")
	util.LogrusObj.Infoln(err)
	return nil, err
}

// Login 用户登陆函数
func (s *UserSrv) Login(c context.Context, req *types.UserServiceReq) (resp interface{}, err error) {
	userDao := dao.NewUserDao(c)
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

// UpdatePwd 用户更改密码
func (s *UserSrv) UpdatePwd(c context.Context, req *types.UserUpdatePwdReq) (resp interface{}, err error) {
	// 找到用户
	userInfo, err := ctl.GetUserInfo(c)
	if err != nil {
		util.LogrusObj.Info(err)
		return nil, err
	}

	userDao := dao.NewUserDao(c)
	user, err := userDao.FindUserByUserId(userInfo.Id)

	if err != nil {
		util.LogrusObj.Info(err)
		return nil, err
	}

	if req.UpdatePwd == "" {
		err = errors.New("更改的密码不能为空")
		util.LogrusObj.Info(err)
		return nil, err
	}

	if err := user.SetPassword(req.UpdatePwd); err != nil {
		util.LogrusObj.Info(err)
		return nil, err
	}

	err = userDao.UpdateUserById(userInfo.Id, user)
	if err != nil {
		util.LogrusObj.Info(err)
		return nil, err
	}

	return ctl.SuccessResp(), nil
}

// UpdateInfo 用户更改信息
func (s *UserSrv) UpdateInfo(c context.Context, req *types.UseUpdateInfoReq) (resp interface{}, err error) {
	// 找到用户
	userInfo, err := ctl.GetUserInfo(c)
	if err != nil {
		util.LogrusObj.Info(err)
		return nil, err
	}

	userDao := dao.NewUserDao(c)
	user, err := userDao.FindUserByUserId(userInfo.Id)
	if err != nil {
		util.LogrusObj.Infoln(err)
		return nil, err
	}

	if req.UpdateName == "" {
		err = errors.New("更改的用户信息不能为空")
		util.LogrusObj.Info(err)
		return nil, err
	} else {
		user.UserName = req.UpdateName
	}

	err = userDao.UpdateUserById(userInfo.Id, user)
	if err != nil {
		util.LogrusObj.Infoln(err)
		return nil, err
	}

	return ctl.SuccessResp(), nil
}
