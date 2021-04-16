package service

import (
	"encoding/json"
	"mesnier/confs"
	"mesnier/model"
	"mesnier/query"
)

type CtripService struct {
}

/**
获取景点详情
*/
func (s *CtripService) GetPoiDetails(ctrip_attraction_id int) (m interface{}, err error) {
	var ctripAttraction = model.CtripAttraction{}
	err = confs.DB.Where("ctrip_attraction_id", ctrip_attraction_id).First(&ctripAttraction).Error
	a := query.GetPoiMoreDetailFunc(ctripAttraction.PoiId)
	p, ok := a["templateList"]
	if ok {
		m = p
	}
	return
}

/**
获取城市列表
*/
func (s *CtripService) GetCityList() (ctripCity []model.CtripCity) {
	confs.DB.Find(&ctripCity)
	return
}

/**
获取某个城市的景点列表
*/

func (s *CtripService) GetCityAttractionList(pageRequest model.PageRequest) (pageResponse model.PageResponse) {
	var (
		PageSize        = pageRequest.PageSize
		PageNum         = pageRequest.PageNum
		total           int64
		ctripAttraction []model.CtripAttraction
	)
	resByre, _ := json.Marshal(pageRequest.QueryModel)
	var newData model.CtripCity
	_ = json.Unmarshal(resByre, &newData)
	confs.DB.Where("ctrip_city_id", newData.CtripCityId).
		Limit(PageNum).
		Offset((PageSize - 1) * PageNum).
		Find(&ctripAttraction)
	confs.DB.Model(&model.CtripAttraction{}).Where("ctrip_city_id", newData.CtripCityId).Count(&total)
	pageResponse.ObjectData = ctripAttraction
	pageResponse.Total = total
	return
}
