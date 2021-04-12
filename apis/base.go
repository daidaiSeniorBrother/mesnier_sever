package apis

import (
	"github.com/gin-gonic/gin"
	"mesnier/utils"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	msg := "OK"
	code := 200
	if err != nil {
		if err, ok := err.(*utils.BusinessError); ok {
			msg = err.Err
			code = err.Code
		} else {
			msg = err.Error()
			code = -1
		}
	}
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: msg,
		Data:    data,
	})
	c.Abort()
}
