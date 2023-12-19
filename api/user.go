package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"SparkForge/pkg/response"
	"SparkForge/pkg/util"
	"SparkForge/service"
	"SparkForge/types"
)

// UserRegisterHandler 用户注册
// @Summary		用户注册
// @Description	注册一个新用户
// @Tags			用户操作
// @Produce		json
// @Param			user	body		types.UserServiceReq	true	"用户注册请求体"
// @Router			/user/register [post]
func UserRegisterHandler(ctx *gin.Context) {
	var req types.UserServiceReq
	if err := ctx.ShouldBind(&req); err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	// 处理响应
	userSrv := service.UserSrv{}
	resp, err := userSrv.Register(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

// UserLoginHandler 用户登录
// @Summary		用户登录
// @Description	用户进行登录操作
// @Tags			用户操作
// @Produce		json
// @Param			user	body		types.UserServiceReq	true	"用户信息"
// @Router			/user/login [post]
func UserLoginHandler(ctx *gin.Context) {
	var req types.UserServiceReq
	if err := ctx.ShouldBind(&req); err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	// 处理响应
	userSrv := service.UserSrv{}
	resp, err := userSrv.Login(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

// UserUpdatePwdHandler 用户修改密码
// @Summary		用户修改密码
// @Description	用户修改密码
// @Tags			登录状态下用户操作
// @Produce		json
// @Param			user	body		types.UserUpdatePwdReq	true	"用户修改密码请求体"
// @Param Authorization header string true "身份验证令牌"
// @Router			/user/password [put]
func UserUpdatePwdHandler(ctx *gin.Context) {
	var req types.UserUpdatePwdReq
	if err := ctx.ShouldBind(&req); err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	// 处理响应
	userSrv := service.UserSrv{}
	err := userSrv.UpdatePwd(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success())
}

// UpdateUserInfoHandler 用户修改信息
// @Summary		用户修改信息
// @Description	用户修改信息
// @Tags			登录状态下用户操作
// @Produce		json
// @Param			user	body		types.UserUpdateInfoReq	true	"用户修改信息请求体"
// @Param Authorization header string true "身份验证令牌"
// @Router			/user/information [put]
func UpdateUserInfoHandler(ctx *gin.Context) {
	var req types.UserUpdateInfoReq
	if err := ctx.ShouldBind(&req); err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusOK, ErrorResponse(err))
		return
	}

	// 处理响应
	userSrv := service.UserSrv{}
	err := userSrv.UpdateInfo(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success())
}

// GetUserInfoHandler 得到用户信息
// @Summary		得到用户信息
// @Description	得到用户信息
// @Tags			登录状态下用户操作
// @Produce		json
// @Param Authorization header string true "身份验证令牌"
// @Router			/user/information [get]
func GetUserInfoHandler(ctx *gin.Context) {

	// 处理响应
	userSrv := service.UserSrv{}
	resp, err := userSrv.UserInfo(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
	}
	ctx.JSON(http.StatusOK, response.SuccessWithData(resp))
}
