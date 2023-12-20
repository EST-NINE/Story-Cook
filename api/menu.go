package api

import (
	"net/http"

	"story-cook-be/pkg/response"
	"story-cook-be/pkg/util"
	"story-cook-be/service"
	"story-cook-be/types"

	"github.com/gin-gonic/gin"
)

func SelectMenuHandler(ctx *gin.Context) {
	var req types.SelectMenuReq
	if err := ctx.ShouldBind(&req); err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusOK, ErrorResponse(err))
		return
	}

	// 处理响应
	menuSrv := service.MenuSrv{}
	resp, err := menuSrv.SelectMenu(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessWithData(resp))
}

func CreateUserMenuHandler(ctx *gin.Context) {
	var req types.CreateUserMenuReq
	if err := ctx.ShouldBind(&req); err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusOK, ErrorResponse(err))
		return
	}

	// 处理响应
	menuSrv := service.MenuSrv{}
	err := menuSrv.CreateUserMenu(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response.Success())
}

func ListUserMenuHandler(ctx *gin.Context) {
	var req types.ListUserMenuReq
	if err := ctx.ShouldBind(&req); err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusOK, ErrorResponse(err))
		return
	}

	// 处理响应
	menuSrv := service.MenuSrv{}
	resp, total, err := menuSrv.ListUserMenu(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response.List(resp, total))
}
