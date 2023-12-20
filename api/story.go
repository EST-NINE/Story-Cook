package api

import (
	"net/http"

	"story-cook-be/pkg/response"
	"story-cook-be/pkg/util"
	"story-cook-be/service"
	"story-cook-be/types"

	"github.com/gin-gonic/gin"
)

func CreateStoryHandler(ctx *gin.Context) {
	var req types.CreateStoryReq
	if err := ctx.ShouldBind(&req); err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	// 处理响应
	storySrv := service.StorySrv{}
	err := storySrv.CreateStory(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success())
}

func ListStoryHandler(ctx *gin.Context) {
	var req types.ListStoryReq
	if err := ctx.ShouldBind(&req); err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	// 处理响应
	storySrv := service.StorySrv{}
	resp, total, err := storySrv.ListStory(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response.List(resp, total))
}

func DeleteStoryHandler(ctx *gin.Context) {
	title := ctx.Param("title")
	req := types.DeleteStoryReq{
		Title: title,
	}

	// 处理响应
	storySrv := service.StorySrv{}
	err := storySrv.DeleteStory(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success())
}

func UpdateStoryHandler(ctx *gin.Context) {
	var req types.UpdateStoryReq
	if err := ctx.ShouldBind(&req); err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	// 处理响应
	storySrv := service.StorySrv{}
	resp, err := storySrv.UpdateStory(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessWithData(resp))
}

func ListStoryByMoodHandler(ctx *gin.Context) {
	var req types.ListStoryByMoodReq
	if err := ctx.ShouldBind(&req); err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	// 处理响应
	storySrv := service.StorySrv{}
	resp, total, err := storySrv.ListStoryByMood(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response.List(resp, total))
}

func ListStoryByTimeHandler(ctx *gin.Context) {
	var req types.ListStoryByTimeReq
	if err := ctx.ShouldBind(&req); err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	// 处理响应
	storySrv := service.StorySrv{}
	resp, total, err := storySrv.ListStoryByTime(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response.List(resp, total))
}
