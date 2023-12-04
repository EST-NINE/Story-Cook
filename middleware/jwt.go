package middleware

import (
	"SparkForge/pkg/ctl"
	"SparkForge/pkg/e"
	"SparkForge/pkg/util"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = e.SUCCESS
		token := c.GetHeader("Authorization")
		if token == "" {
			code = e.ErrorAuthCheckTokenFail
			c.JSON(http.StatusBadRequest, ctl.ErrorResp(errors.New("empty"), e.GetMsg(code), code))
			c.Abort()
			return
		}

		claims, err := util.ParseToken(token)
		if err != nil {
			code = e.ErrorAuthCheckTokenFail
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = e.ErrorAuthCheckTokenTimeout
		}

		if code != e.SUCCESS {
			c.JSON(e.InvalidParams, ctl.ErrorResp(err, e.GetMsg(code), code))
			c.Abort()
			return
		}

		c.Request = c.Request.WithContext(ctl.NewContext(c.Request.Context(), &ctl.UserInfo{Id: claims.Id}))
		c.Next()
	}
}
