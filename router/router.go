package router

import (
	"SparkForge/api"
	"SparkForge/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
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
		v1.POST("user/register", api.UserRegisterHandler())
		v1.POST("user/login", api.UserLoginHandler())


		authed := v1.Group("/") // 登录保护
		authed.Use(middleware.JWT())
		{
			// 用户操作
			authed.GET("user/isLogin", func(c *gin.Context) {
				c.JSON(http.StatusOK, "登录成功！")
			})
			authed.POST("user/updatePwd", api.UserUpdatePwdHandler())
			authed.POST("user/updateInfo", api.UserUpdateInfoHandler())
			authed.GET("user/info", api.GetUserInfoHandler())

			// 获取故事
			authed.POST("user/getStory", api.GetStoryHandler())
		}
	}

	return r
}
