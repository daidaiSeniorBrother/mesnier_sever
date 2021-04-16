package dto

type Attraction struct {
	Result         int              `json:"result"`
	TotalCount     int              `json:"totalCount"`
	DistrictId     int              `json:"districtId"`
	AttractionList []AttractionList `json:"attractionList"`
}

type AttractionList struct {
	ShowType int            `json:"showType"`
	Card     AttractionCard `json:"card"`
}
type AttractionCard struct {
	PoiId         int     `json:"poiId"` //景点ID
	BusinessId    int     `json:"businessId"`
	ZoneName      string  `json:"zoneName"`      //地区名称
	PoiName       string  `json:"poiName"`       //景点名称
	CoverImageUrl string  `json:"coverImageUrl"` //缩略图
	DistanceStr   string  `json:"distanceStr"`   //距离
	Price         float64 `json:"price"`         //价格
	IsFree        bool    `json:"isFree"`        //是否免费
	OpenStatus    string  `json:"openStatus"`    //开启状态
}
