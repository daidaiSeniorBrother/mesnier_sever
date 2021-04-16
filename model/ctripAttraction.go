package model

type CtripAttraction struct {
	CtripAttractionId int     `gorm:"column:ctrip_attraction_id;primaryKey;autoIncrement" json:"ctrip_attraction_id"`
	PoiId             int     `gorm:"column:poi_id" json:"poi_id"`
	BusinessId        int     `gorm:"column:business_id" json:"business_id"`
	ZoneName          string  `gorm:"column:zone_name" json:"zone_name"`
	PoiName           string  `gorm:"column:poi_name" json:"poi_name"`
	CoverImageUrl     string  `gorm:"column:cover_image_url" json:"cover_image_url"`
	DistanceStr       string  `gorm:"column:distance_str" json:"distance_str"`
	Price             float64 `gorm:"column:price" json:"price"`
	IsFree            bool    `gorm:"column:is_free" json:"is_free"`
	OpenStatus        string  `gorm:"column:open_status" json:"open_status"`
	CtripCityId       int     `gorm:"column:ctrip_city_id" json:"ctrip_city_id"`
	DistrictId        int     `gorm:"column:district_id" json:"district_id"`
}

func (c *CtripAttraction) TableName() string {
	return "ctrip_attraction"
}
