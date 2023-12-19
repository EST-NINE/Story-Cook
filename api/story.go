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
//	@Summary		创建历史记录
//	@Description	创建历史记录
//	@Tags			历史记录操作
//	@Produce		json
//	@Param			story	body		types.CreateStoryReq	true	"创建历史记录请求体"
//	@Param Authorization header string true "身份验证令牌"
//	@Router			/story/save [post]
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
	ctx.JSON(http.StatusOK, controller.SuccessResp())
}

// ListStoryHandler 得到用户的故事列表
//
//	@Summary		得到用户的故事列表
//	@Description	得到用户的故事列表
//	@Tags			历史记录操作
//	@Produce		json
//	@Param			story	body		types.ListStoryReq	true	"故事列表请求体"
//	@Param Authorization header string true "身份验证令牌"
//	@Router			/story/list [post]
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
	ctx.JSON(http.StatusOK, controller.ListResp(resp, total))
}

// DeleteStoryHandler 删除历史记录
// @Summary 删除历史记录
// @Description 删除历史记录
// @Tags 历史记录操作
// @Produce json
// @Param title path string true "历史记录标题"
// @Param Authorization header string true "身份验证令牌"
// @Router /story/{title} [delete]
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
	ctx.JSON(http.StatusOK, controller.SuccessResp())
}

// UpdateStoryHandler 更新历史记录
//
//	@Summary		更新历史记录
//	@Description	更新历史记录
//	@Tags			历史记录操作
//	@Produce		json
//	@Param			story	body		types.UpdateStoryReq	true	"更新历史记录请求体"
//	@Param Authorization header string true "身份验证令牌"
//	@Router			/story [put]
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
	ctx.JSON(http.StatusOK, controller.SuccessWithDataResp(resp))
}

// ListStoryByMoodHandler 根据mood分类历史记录
//
//	@Summary		根据mood分类历史记录
//	@Description	根据mood分类历史记录
//	@Tags			历史记录操作
//	@Produce		json
//	@Param			story	body		types.ListStoryByMoodReq	true	"分类历史记录请求体"
//	@Param Authorization header string true "身份验证令牌"
//	@Router			/story/listByMood [post]
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
	ctx.JSON(http.StatusOK, controller.ListResp(resp, total))
}

// ListStoryByTimeHandler 根据time分类历史记录
//
//	@Summary		根据time分类历史记录
//	@Description	根据time分类历史记录
//	@Tags			历史记录操作
//	@Produce		json
//	@Param			story	body		types.ListStoryByTimeReq	true	"分类历史记录请求体"
//	@Param Authorization header string true "身份验证令牌"
//	@Router			/story/listByTime [post]
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
	ctx.JSON(http.StatusOK, controller.ListResp(resp, total))
}
