package api

import (
	"SparkForge/pkg/util"
	"SparkForge/service"
	"SparkForge/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserRegisterHandler 用户注册
func UserRegisterHandler(ctx *gin.Context) {
	var req types.UserServiceReq
	if err := ctx.ShouldBind(&req); err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	// 处理响应
	userSrv := service.UserSrv{}
	resp, err := userSrv.Register(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

// UserLoginHandler 用户登录
func UserLoginHandler(ctx *gin.Context) {
	var req types.UserServiceReq
	if err := ctx.ShouldBind(&req); err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	// 处理响应
	userSrv := service.UserSrv{}
	resp, err := userSrv.Login(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

// UserUpdatePwdHandler 用户修改密码
func UserUpdatePwdHandler(ctx *gin.Context) {
	var req types.UserUpdatePwdReq
	if err := ctx.ShouldBind(&req); err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	// 处理响应
	userSrv := service.UserSrv{}
	resp, err := userSrv.UpdatePwd(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

// UserUpdateInfoHandler 用户修改信息
func UserUpdateInfoHandler(ctx *gin.Context) {
	var req types.UseUpdateInfoReq
	if err := ctx.ShouldBind(&req); err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusOK, ErrorResponse(err))
		return
	}

	// 处理响应
	userSrv := service.UserSrv{}
	resp, err := userSrv.UpdateInfo(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

// GetUserInfoHandler 用户信息
func GetUserInfoHandler(ctx *gin.Context) {

	// 处理响应
	userSrv := service.UserSrv{}
	resp, err := userSrv.UserInfo(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
	}
	ctx.JSON(http.StatusOK, resp)
}
