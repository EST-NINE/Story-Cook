package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"SparkForge/pkg/controller"
	"SparkForge/pkg/util"
	"SparkForge/service"
	"SparkForge/types"
)

// CreateStoryHandler 创建故事
func CreateStoryHandler(ctx *gin.Context) {
	var req types.CreateStoryReq
	if err := ctx.ShouldBind(&req); err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	// 处理响应
	storySrv := service.StorySrv{}
	err := storySrv.CreateStory(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, controller.SuccessResp())
}

// ListStoryHandler 得到对应用户的故事列表
func ListStoryHandler(ctx *gin.Context) {
	var req types.ListStoryReq
	if err := ctx.ShouldBind(&req); err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	// 处理响应
	storySrv := service.StorySrv{}
	resp, total, err := storySrv.ListStory(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, controller.ListResp(resp, total))
}

// DeleteStoryHandler 删除故事
func DeleteStoryHandler(ctx *gin.Context) {
	var req types.DeleteStoryReq
	if err := ctx.ShouldBind(&req); err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	// 处理响应
	storySrv := service.StorySrv{}
	err := storySrv.DeleteStory(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, controller.SuccessResp())
}

// UpdateStoryHandler 更新故事
func UpdateStoryHandler(ctx *gin.Context) {
	var req types.UpdateStoryReq
	if err := ctx.ShouldBind(&req); err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	// 处理响应
	storySrv := service.StorySrv{}
	resp, err := storySrv.UpdateStory(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, controller.SuccessWithDataResp(resp))
}

// SelectStoryHandler 根据mood分类故事
func SelectStoryHandler(ctx *gin.Context) {
	var req types.SelectStoryReq
	if err := ctx.ShouldBind(&req); err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	// 处理响应
	storySrv := service.StorySrv{}
	resp, total, err := storySrv.SelectStory(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, controller.ListResp(resp, total))
}
