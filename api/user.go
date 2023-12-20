package api

import (
	"net/http"

	"story-cook-be/pkg/response"
	"story-cook-be/pkg/util"
	"story-cook-be/service"
	"story-cook-be/types"

	"github.com/gin-gonic/gin"
)

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

func GetUserInfoHandler(ctx *gin.Context) {

	// 处理响应
	userSrv := service.UserSrv{}
	resp, err := userSrv.UserInfo(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
	}
	ctx.JSON(http.StatusOK, response.SuccessWithData(resp))
}
