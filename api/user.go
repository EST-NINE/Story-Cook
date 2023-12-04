package api

import (
	"SparkForge/pkg/util"
	"SparkForge/service"
	"SparkForge/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserRegisterHandler 用户注册
func UserRegisterHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserServiceReq
		if err := ctx.ShouldBind(&req); err != nil {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
			return
		}

		// 参数校验
		server := service.GetUserSrv()
		resp, err := server.Register(ctx.Request.Context(), &req)
		if err != nil {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusOK, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, resp)

	}
}

// UserLoginHandler 用户登录
func UserLoginHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserServiceReq
		if err := ctx.ShouldBind(&req); err != nil {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
			return
		}

		// 参数校验
		server := service.GetUserSrv()
		resp, err := server.Login(ctx.Request.Context(), &req)
		if err != nil {
			ctx.JSON(http.StatusOK, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}
