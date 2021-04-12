package routers

import (
	"github.com/gin-gonic/gin"
	"mesnier/apis"
	"mesnier/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.VerifyToken())

	api := r.Group("/api")
	{
		user := api.Group("/user")
		{
			user.POST("/login", func(c *gin.Context) {
				apis.UserLogin(c)
			})
			user.POST("/list", func(c *gin.Context) {
				apis.UserList(c)
			})
		}
	}
	return r
}
