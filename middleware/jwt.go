package middleware

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"SparkForge/pkg/controller"
	"SparkForge/pkg/errCode"
	"SparkForge/pkg/util"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var code int
		code = errCode.SUCCESS
		token := ctx.GetHeader("Authorization")
		if token == "" {
			code = errCode.ErrorAuthCheckTokenFail
			ctx.JSON(http.StatusBadRequest, controller.ErrorResp(errors.New("empty"), errCode.GetMsg(code), code))
			ctx.Abort()
			return
		}

		claims, err := util.ParseToken(token)
		if err != nil {
			code = errCode.ErrorAuthCheckTokenFail
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = errCode.ErrorAuthCheckTokenTimeout
		}

		if code != errCode.SUCCESS {
			ctx.JSON(errCode.InvalidParams, controller.ErrorResp(err, errCode.GetMsg(code), code))
			ctx.Abort()
			return
		}

		ctx.Request = ctx.Request.WithContext(controller.NewContext(ctx.Request.Context(), &controller.UserInfo{Id: claims.Id}))
		ctx.Next()
	}
}
