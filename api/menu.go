package api

import (
	"SparkForge/pkg/ctl"
	"SparkForge/pkg/util"
	"SparkForge/service"
	"SparkForge/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SelectMenuHandler 判断是否触发彩蛋成就
func SelectMenuHandler(ctx *gin.Context) {
	var req types.SelectMenuReq
	if err := ctx.ShouldBind(&req); err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusOK, ErrorResponse(err))
		return
	}

	// 处理响应
	menuSrv := service.MenuSrv{}
	resp, err := menuSrv.SelectMenu(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, ctl.SuccessWithDataResp(resp))
}

// CreateUserMenuHandler 给用户添加彩蛋成就
func CreateUserMenuHandler(ctx *gin.Context) {
	var req types.CreateUserMenuReq
	if err := ctx.ShouldBind(&req); err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusOK, ErrorResponse(err))
		return
	}

	// 处理响应
	menuSrv := service.MenuSrv{}
	err := menuSrv.CreateUserMenu(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, ctl.SuccessResp())
}

// ListUserMenuHandler 得到用户的彩蛋成就列表
func ListUserMenuHandler(ctx *gin.Context) {
	var req types.ListUserMenuReq
	if err := ctx.ShouldBind(&req); err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusOK, ErrorResponse(err))
		return
	}

	// 处理响应
	menuSrv := service.MenuSrv{}
	resp, total, err := menuSrv.ListUserMenu(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, ctl.ListResp(resp, total))
}
