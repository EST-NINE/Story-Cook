package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"SparkForge/pkg/controller"
	"SparkForge/pkg/util"
	"SparkForge/service"
	"SparkForge/types"
)

// CreateStoryHandler 创建历史记录
//
//		@Summary		创建历史记录
//		@Description	创建历史记录
//		@Tags			历史记录操作
//		@Produce		json
//		@Param			story	body		types.CreateStoryReq	true	"创建历史记录请求体"
//	    @Param Authorization header string true "身份验证令牌"
//		@Router			/story/save [post]
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

// ListStoryHandler 得到用户的故事列表
//
//		@Summary		得到用户的故事列表
//		@Description	得到用户的故事列表
//		@Tags			历史记录操作
//		@Produce		json
//		@Param			story	body		types.ListStoryReq	true	"故事列表请求体"
//	    @Param Authorization header string true "身份验证令牌"
//		@Router			/story/list [post]
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

// DeleteStoryHandler 删除历史记录
//
//		@Summary		删除历史记录
//		@Description	删除历史记录
//		@Tags			历史记录操作
//		@Produce		json
//		@Param			story	body		types.DeleteStoryReq	true	"删除历史记录请求体"
//	    @Param Authorization header string true "身份验证令牌"
//		@Router			/story [delete]
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

// UpdateStoryHandler 更新历史记录
//
//		@Summary		更新历史记录
//		@Description	更新历史记录
//		@Tags			历史记录操作
//		@Produce		json
//		@Param			story	body		types.UpdateStoryReq	true	"更新历史记录请求体"
//	    @Param Authorization header string true "身份验证令牌"
//		@Router			/story [put]
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

// SelectStoryHandler 根据mood分类历史记录
//
//		@Summary		根据mood分类历史记录
//		@Description	根据mood分类历史记录
//		@Tags			历史记录操作
//		@Produce		json
//		@Param			story	body		types.SelectStoryReq	true	"分类历史记录请求体"
//	    @Param Authorization header string true "身份验证令牌"
//		@Router			/story/select [post]
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
