package middleware

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"SparkForge/pkg/controller"
	"SparkForge/pkg/errMsg"
	"SparkForge/pkg/util"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var code int
		code = errMsg.SUCCESS
		token := ctx.GetHeader("Authorization")
		if token == "" {
			code = errMsg.ErrorAuthCheckTokenFail
			ctx.JSON(http.StatusBadRequest, controller.ErrorResp(errors.New("empty"), errMsg.GetMsg(code), code))
			ctx.Abort()
			return
		}

		claims, err := util.ParseToken(token)
		if err != nil {
			code = errMsg.ErrorAuthCheckTokenFail
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = errMsg.ErrorAuthCheckTokenTimeout
		}

		if code != errMsg.SUCCESS {
			ctx.JSON(errMsg.InvalidParams, controller.ErrorResp(err, errMsg.GetMsg(code), code))
			ctx.Abort()
			return
		}

		ctx.Request = ctx.Request.WithContext(controller.NewContext(ctx.Request.Context(), &controller.UserInfo{Id: claims.Id}))
		ctx.Next()
	}
}
