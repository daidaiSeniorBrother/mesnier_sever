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
		tourism := api.Group("/tourism")
		{
			tourism.GET("/city/list", func(context *gin.Context) {
				apis.GetCityList(context)
			})
			tourism.POST("/attraction/list", func(context *gin.Context) {
				apis.GetCityAttractionList(context)
			})
			tourism.GET("/poi/details", func(context *gin.Context) {
				apis.GetPoiDetails(context)
			})

		}
	}
	return r
}
