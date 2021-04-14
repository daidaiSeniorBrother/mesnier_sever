package model

type CtripCity struct {
	CtripCityId    int    `gorm:"column:ctrip_city_id;primaryKey;autoIncrement" json:"ctrip_city_id"`
	CtripCityCode  string `gorm:"column:ctrip_city_code" json:"ctrip_city_code"`
	CtripCityQuery string `gorm:"column:ctrip_city_query" json:"ctrip_city_query"`
	CtripCityName  string `gorm:"column:ctrip_city_name" json:"ctrip_city_name"`
}

func (c *CtripCity) TableName() string {
	return "ctrip_city"
}
