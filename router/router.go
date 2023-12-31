package router

import (
	"net/http"
	"story-cook-be/api"
	"story-cook-be/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.CORS())
	public := r.Group("api/v1")
	{
		public.GET("ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "pong!")
		})

		// 用户操作
		public.POST("user/register", api.UserRegisterHandler)
		public.POST("user/login", api.UserLoginHandler)

		authed := public.Group("/") // 登录保护
		authed.Use(middleware.JWT())
		{
			// 用户操作
			authed.GET("user/isLogin", func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, "登录成功！")
			})
			authed.PUT("user/password", api.UserUpdatePwdHandler)
			authed.PUT("user/information", api.UpdateUserInfoHandler)
			authed.GET("user/information", api.GetUserInfoHandler)
			authed.PUT("user_count", api.UpdateUserCountHandler)

			// 故事操作
			authed.POST("story/generate", api.GenerateStoryHandler)
			authed.POST("story/save", api.CreateStoryHandler)
			authed.POST("story/list", api.ListStoryHandler)
			authed.POST("story/listByMood", api.ListStoryByMoodHandler)
			authed.POST("story/listByTime", api.ListStoryByTimeHandler)
			authed.DELETE("story/:title", api.DeleteStoryHandler)
			authed.PUT("story", api.UpdateStoryHandler)

			// 彩蛋操作
			authed.POST("userMenu/isMenu", api.SelectMenuHandler)
			authed.POST("userMenu/create", api.CreateUserMenuHandler)
			authed.POST("userMenu/list", api.ListUserMenuHandler)
		}
	}

	return r
}
