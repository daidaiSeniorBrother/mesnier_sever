package model

type Extension struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Filter struct {
	FilterItems []FilterItems `json:"filterItems"`
}
type FilterItems struct {
}
type Head struct {
	Cid       string      `json:"cid"`
	Ctok      string      `json:"ctok"`
	Cver      string      `json:"cver"`
	Lang      string      `json:"lang"`
	Sid       string      `json:"sid"`
	Syscode   string      `json:"syscode"`
	Auth      string      `json:"auth"`
	Xsid      string      `json:"xsid"`
	Extension []Extension `json:"extension"`
}

type CtripCityAttraction struct {
	Index             int         `json:"index"`
	Count             int         `json:"count"`
	SortType          int         `json:"sortType"`
	IsShowAggregation bool        `json:"isShowAggregation"`
	DistrictId        int         `json:"districtId"`
	Scene             string      `json:"scene"`
	PageId            string      `json:"pageId"`
	TraceId           string      `json:"traceId"`
	Extension         []Extension `json:"extension"`
	Filter            `json:"filter"`
	CrnVersion        string `json:"crnVersion"`
	IsInitialState    bool   `json:"isInitialState"`
	Head              `json:"head"`
}

type CtripCityTemplate struct {
	PoiId int    `json:"poiId"`
	Scene string `json:"scene"`
	Head  `json:"head"`
}
