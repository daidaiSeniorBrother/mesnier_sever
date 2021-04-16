package apis

import (
	"github.com/gin-gonic/gin"
	"mesnier/model"
	"mesnier/service"
	"strconv"
)

var ctripService = service.CtripService{}

func GetPoiDetails(c *gin.Context) {
	ctripAttractionId := c.Query("ctrip_attraction_id")
	atom, err := strconv.Atoi(ctripAttractionId)
	m, err := ctripService.GetPoiDetails(atom)
	SendResponse(c, err, m)
}

func GetCityList(c *gin.Context) {
	SendResponse(c, nil, ctripService.GetCityList())
}

func GetCityAttractionList(c *gin.Context) {
	var (
		pageRequest model.PageRequest
	)
	_ = c.BindJSON(&pageRequest)
	SendResponse(c, nil, ctripService.GetCityAttractionList(pageRequest))
}
