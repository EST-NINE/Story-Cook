package service

import (
	"errors"
	"fmt"
	"story-cook-be/dao"
	"story-cook-be/model"
	"story-cook-be/pkg/response"
	"story-cook-be/pkg/util"
	"story-cook-be/types"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserSrv struct {
}

// Register 注册用户
func (s *UserSrv) Register(ctx *gin.Context, req *types.UserServiceReq) (resp response.TokenDataResp, err error) {
	userDao := dao.NewUserDao(ctx)
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
				return
			}

			if err = userDao.CreateUser(user); err != nil {
				util.LogrusObj.Info(err)
				return
			}

			token, _ := util.GenerateToken(user.ID, req.UserName)
			return response.TokenDataResp{
				User:  response.BuildUserResp(user),
				Token: token,
			}, nil
		}
		return
	}

	err = errors.New("用户已存在")
	util.LogrusObj.Infoln(err)
	return
}

// Login 用户登陆函数
func (s *UserSrv) Login(ctx *gin.Context, req *types.UserServiceReq) (resp response.TokenDataResp, err error) {
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

	return response.TokenDataResp{
		User:  response.BuildUserResp(user),
		Token: token,
	}, nil
}

// UpdatePwd 用户更改密码
func (s *UserSrv) UpdatePwd(ctx *gin.Context, req *types.UserUpdatePwdReq) error {
	claims, _ := ctx.Get("claims")
	userInfo := claims.(*util.Claims)

	userDao := dao.NewUserDao(ctx)
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
func (s *UserSrv) UpdateInfo(ctx *gin.Context, req *types.UserUpdateInfoReq) error {
	claims, _ := ctx.Get("claims")
	userInfo := claims.(*util.Claims)

	userDao := dao.NewUserDao(ctx)
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
func (s *UserSrv) UserInfo(ctx *gin.Context) (resp *response.UserResp, err error) {
	claims, _ := ctx.Get("claims")
	userInfo := claims.(*util.Claims)

	userDao := dao.NewUserDao(ctx)
	user, err := userDao.FindUserByUserId(userInfo.Id)

	return response.BuildUserResp(user), nil
}
