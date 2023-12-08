package api

import (
	"SparkForge/pkg/util"
	"SparkForge/service"
	"SparkForge/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SelectMenuHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.SelectMenuReq
		if err := c.ShouldBind(&req); err != nil {
			util.LogrusObj.Infoln(err)
			c.JSON(http.StatusOK, ErrorResponse(err))
			return
		}

		// 处理响应
		server := service.GetMenuSrv()
		resp, err := server.SelectMenu(c.Request.Context(), &req)
		if err != nil {
			c.JSON(http.StatusOK, ErrorResponse(err))
			return
		}

		c.JSON(http.StatusOK, resp)
	}
}
