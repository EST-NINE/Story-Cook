package service

import (
	"SparkForge/repository/db/dao"
	"SparkForge/repository/db/model"
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"

	"SparkForge/pkg/controller"
	"SparkForge/pkg/util"
	"SparkForge/types"
)

type UserSrv struct {
}

// Register 注册用户
func (s *UserSrv) Register(c context.Context, req *types.UserServiceReq) error {
	userDao := dao.NewUserDao(c)
	user, err := userDao.FindUserByUserName(req.UserName)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			user = &model.User{
				UserName: req.UserName,
				Kitchen:  fmt.Sprint(req.UserName, "的厨房"),
			}
			// 密码加密存储
			if err = user.SetPassword(req.Password); err != nil {
				util.LogrusObj.Info(err)
				return err
			}

			if err = userDao.CreateUser(user); err != nil {
				util.LogrusObj.Info(err)
				return err
			}

			return nil
		}
		return err
	}

	err = errors.New("用户已存在")
	util.LogrusObj.Infoln(err)
	return err
}

// Login 用户登陆函数
func (s *UserSrv) Login(c context.Context, req *types.UserServiceReq) (resp types.TokenDataResp, err error) {
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
		Kitchen:  user.Kitchen,
		Count:    user.GetCount(),
		CreateAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
	}
	return types.TokenDataResp{
		User:  userResp,
		Token: token,
	}, nil
}

// UpdatePwd 用户更改密码
func (s *UserSrv) UpdatePwd(c context.Context, req *types.UserUpdatePwdReq) error {
	// 找到用户
	userInfo, err := controller.GetUserInfo(c)
	if err != nil {
		util.LogrusObj.Info(err)
		return err
	}

	userDao := dao.NewUserDao(c)
	user, err := userDao.FindUserByUserId(userInfo.Id)

	if err != nil {
		util.LogrusObj.Info(err)
		return err
	}

	if !user.CheckPassword(req.OriginPwd) {
		err = errors.New("原密码错误")
		util.LogrusObj.Info(err)
		return err
	}

	if err := user.SetPassword(req.UpdatePwd); err != nil {
		util.LogrusObj.Info(err)
		return err
	}

	err = userDao.UpdateUserById(userInfo.Id, user)
	if err != nil {
		util.LogrusObj.Info(err)
		return err
	}

	return nil
}

// UpdateInfo 用户更改信息
func (s *UserSrv) UpdateInfo(c context.Context, req *types.UserUpdateInfoReq) error {
	// 找到用户
	userInfo, err := controller.GetUserInfo(c)
	if err != nil {
		util.LogrusObj.Info(err)
		return err
	}

	userDao := dao.NewUserDao(c)
	user, err := userDao.FindUserByUserId(userInfo.Id)
	if err != nil {
		util.LogrusObj.Infoln(err)
		return err
	}

	if req.UpdateName != "" {
		_, err := userDao.FindUserByUserName(req.UpdateName)
		if err == nil {
			err = errors.New("用户已存在")
			return err
		}
		user.UserName = req.UpdateName
	}

	if req.Kitchen != "" {
		user.Kitchen = req.Kitchen
	}

	err = userDao.UpdateUserById(userInfo.Id, user)
	if err != nil {
		util.LogrusObj.Infoln(err)
		return err
	}

	return nil
}

// UserInfo 得到用户的信息
func (s *UserSrv) UserInfo(c context.Context) (resp *types.UserResp, err error) {
	// 找到用户
	userInfo, err := controller.GetUserInfo(c)
	if err != nil {
		util.LogrusObj.Infoln(err)
		return
	}

	userDao := dao.NewUserDao(c)
	user, err := userDao.FindUserByUserId(userInfo.Id)

	return &types.UserResp{
		ID:       user.ID,
		UserName: user.UserName,
		Kitchen:  user.Kitchen,
		Count:    user.GetCount(),
		CreateAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}
