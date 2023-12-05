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
	return func(c *gin.Context) {
		var req types.UserServiceReq
		if err := c.ShouldBind(&req); err != nil {
			util.LogrusObj.Infoln(err)
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			return
		}

		// 处理响应
		server := service.GetUserSrv()
		resp, err := server.Register(c.Request.Context(), &req)
		if err != nil {
			c.JSON(http.StatusOK, ErrorResponse(err))
			return
		}
		c.JSON(http.StatusOK, resp)

	}
}

// UserLoginHandler 用户登录
func UserLoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.UserServiceReq
		if err := c.ShouldBind(&req); err != nil {
			util.LogrusObj.Infoln(err)
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			return
		}

		// 处理响应
		server := service.GetUserSrv()
		resp, err := server.Login(c.Request.Context(), &req)
		if err != nil {
			c.JSON(http.StatusOK, ErrorResponse(err))
			return
		}
		c.JSON(http.StatusOK, resp)
	}
}

// UserUpdatePwdHandler 用户修改密码
func UserUpdatePwdHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.UserUpdatePwdReq
		if err := c.ShouldBind(&req); err != nil {
			util.LogrusObj.Infoln(err)
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			return
		}

		// 处理响应
		server := service.GetUserSrv()
		resp, err := server.UpdatePwd(c.Request.Context(), &req)
		if err != nil {
			c.JSON(http.StatusOK, ErrorResponse(err))
			return
		}
		c.JSON(http.StatusOK, resp)
	}
}

// UserUpdateInfoHandler 用户修改信息
func UserUpdateInfoHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.UseUpdateInfoReq
		if err := c.ShouldBind(&req); err != nil {
			util.LogrusObj.Infoln(err)
			c.JSON(http.StatusOK, ErrorResponse(err))
			return
		}

		// 处理响应
		server := service.GetUserSrv()
		resp, err := server.UpdateInfo(c.Request.Context(), &req)
		if err != nil {
			c.JSON(http.StatusOK, ErrorResponse(err))
			return
		}
		c.JSON(http.StatusOK, resp)
	}
}

// GetUserInfoHandler 用户信息
func GetUserInfoHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 处理响应
		server := service.GetUserSrv()
		resp, err := server.UserInfo(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusOK, ErrorResponse(err))
		}
		c.JSON(http.StatusOK, resp)
	}
}
