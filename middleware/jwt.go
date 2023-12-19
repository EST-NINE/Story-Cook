package middleware

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"SparkForge/pkg/errCode"
	"SparkForge/pkg/response"
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
			ctx.JSON(http.StatusBadRequest, response.Error(errors.New("empty"), errCode.GetMsg(code), code))
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
			ctx.JSON(errCode.InvalidParams, response.Error(err, errCode.GetMsg(code), code))
			ctx.Abort()
			return
		}

		ctx.Set("claims", claims)
		ctx.Next()
	}
}
