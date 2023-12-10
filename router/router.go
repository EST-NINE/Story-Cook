package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"SparkForge/api"
	"SparkForge/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.CORS())
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, "pong!")
		})

		// 用户操作
		v1.POST("user/register", api.UserRegisterHandler)
		v1.POST("user/login", api.UserLoginHandler)

		authed := v1.Group("/") // 登录保护
		authed.Use(middleware.JWT())
		{
			// 用户操作
			authed.GET("user/isLogin", func(c *gin.Context) {
				c.JSON(http.StatusOK, "登录成功！")
			})
			authed.PUT("user/password", api.UserUpdatePwdHandler)
			authed.PUT("user/information", api.UserUpdateInfoHandler)
			authed.GET("user/information", api.GetUserInfoHandler)

			// 故事操作
			authed.POST("story/generate", api.GenerateStoryHandler)
			authed.POST("story/save", api.CreateStoryHandler)
			authed.POST("story/list", api.ListStoryHandler)
			authed.POST("story/select", api.SelectStoryHandler)
			authed.DELETE("story", api.DeleteStoryHandler)
			authed.PUT("story", api.UpdateStoryHandler)

			// 彩蛋操作
			authed.POST("userMenu/isMenu", api.SelectMenuHandler)
			authed.POST("userMenu/create", api.CreateUserMenuHandler)
			authed.POST("userMenu/list", api.ListUserMenuHandler)
		}
	}

	return r
}
