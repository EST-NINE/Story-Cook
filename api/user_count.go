package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"story-cook-be/pkg/response"
	"story-cook-be/service"
)

func UpdateUserCountHandler(ctx *gin.Context) {
	// 处理响应
	userCountSrv := service.UserCountSrv{}
	err := userCountSrv.Update(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success())
}
