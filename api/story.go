package api

import (
	"SparkForge/pkg/util"
	"SparkForge/service"
	"SparkForge/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateStoryHandler 创建故事
func CreateStoryHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.CreateStoryReq
		if err := c.ShouldBind(&req); err != nil {
			util.LogrusObj.Infoln(err)
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			return
		}

		// 处理响应
		server := service.GetStorySrv()
		resp, err := server.CreateStory(c.Request.Context(), &req)
		if err != nil {
			c.JSON(http.StatusOK, ErrorResponse(err))
			return
		}
		c.JSON(http.StatusOK, resp)
	}
}

// ListStoryHandler 得到对应用户的故事列表
func ListStoryHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.ListStoryReq
		if err := c.ShouldBind(&req); err != nil {
			util.LogrusObj.Infoln(err)
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			return
		}

		// 处理响应
		if req.Limit == 0 {
			req.Limit = 15
		}
		server := service.GetStorySrv()
		resp, err := server.ListStory(c.Request.Context(), &req)
		if err != nil {
			c.JSON(http.StatusOK, ErrorResponse(err))
			return
		}
		c.JSON(http.StatusOK, resp)
	}
}
