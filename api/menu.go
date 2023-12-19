package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"SparkForge/pkg/response"
	"SparkForge/pkg/util"
	"SparkForge/service"
	"SparkForge/types"
)

// SelectMenuHandler 判断是否触发彩蛋成就
// @Summary		判断是否触发彩蛋成就
// @Description	判断是否触发彩蛋成就
// @Tags			彩蛋操作
// @Produce		json
// @Param			userMenu	body		types.SelectMenuReq	true	"判断彩蛋成就请求体"
// @Param Authorization header string true "身份验证令牌"
// @Router			/userMenu/isMenu [post]
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

// CreateUserMenuHandler 添加彩蛋成就
// @Summary		添加彩蛋成就
// @Description	添加彩蛋成就
// @Tags			彩蛋操作
// @Produce		json
// @Param			userMenu	body		types.CreateUserMenuReq	true	"添加彩蛋成就请求体"
// @Param Authorization header string true "身份验证令牌"
// @Router			/userMenu/create [post]
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

// ListUserMenuHandler 得到彩蛋成就列表
// @Summary		得到彩蛋成就列表
// @Description	得到彩蛋成就列表
// @Tags			彩蛋操作
// @Produce		json
// @Param			userMenu	body		types.ListUserMenuReq	true	"彩蛋成就列表请求体"
// @Param Authorization header string true "身份验证令牌"
// @Router			/userMenu/list [post]
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
