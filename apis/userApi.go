package apis

import (
	"github.com/gin-gonic/gin"
	"mesnier/model"
	"mesnier/service"
)

var userService = service.UserService{}

func UserLogin(c *gin.Context) {
	var (
		u model.User
	)
	_ = c.BindJSON(&u)
	token, err := userService.UserLogin(&u)
	SendResponse(c, err, token)
}

func UserList(c *gin.Context) {
	var (
		u model.User
	)
	_ = c.BindJSON(&u)
	users, err := userService.UserList(&u)
	SendResponse(c, err, users)
}
