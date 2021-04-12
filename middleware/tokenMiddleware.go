package middleware

import (
	"github.com/gin-gonic/gin"
	"mesnier/apis"
	"mesnier/utils"
	"net/http"
)

var notVerifyTokenPath = []string{"/api/user/login"}

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.FullPath()
		b := utils.IsContain(notVerifyTokenPath, path)
		if !b {
			token := c.GetHeader("AuthorizationToken")
			token2 := c.Query("AuthorizationToken")
			if token == "" && token2 == "" {
				c.JSON(http.StatusUnauthorized, apis.Response{
					Code:    401,
					Message: "Token失效,请重新登录",
					Data:    "",
				})
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
