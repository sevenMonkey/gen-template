package jwt

import (
	"funds/app"
	"funds/pkg/e"
	"funds/pkg/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			code int
		)
		code = e.SUCCESS
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			code = e.ERROR_INVALID_AUTH
			app.RespJson(c, code, nil)
			c.Abort()
			return
		}
		j := util.NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			default:
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			}
		}
		if code != e.SUCCESS {
			app.RespJson(c, code, nil)
			c.Abort()
			return
		}
		c.Set("claims", claims)
		if len(claims.Id) > 0 {
			c.Set("userId", claims.Id)
			c.Set("role", claims.Role)
			c.Next()
		} else {
			app.RespJson(c, e.ERROR_INVALID_AUTH, nil)
			c.Abort()
			return
		}
	}
}
