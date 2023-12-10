package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"SparkForge/pkg/ctl"
	"SparkForge/pkg/util"
	"SparkForge/service"
	"SparkForge/types"
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
	err := userSrv.Register(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, ctl.SuccessResp())
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
	err := userSrv.UpdatePwd(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, ctl.SuccessResp())
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
	err := userSrv.UpdateInfo(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, ctl.SuccessResp())
}

// GetUserInfoHandler 用户信息
func GetUserInfoHandler(ctx *gin.Context) {

	// 处理响应
	userSrv := service.UserSrv{}
	resp, err := userSrv.UserInfo(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
	}
	ctx.JSON(http.StatusOK, ctl.SuccessWithDataResp(resp))
}
