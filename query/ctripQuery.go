package query

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"mesnier/model"
	"mesnier/model/dto"
	"mesnier/utils"
	"strings"
)

const GetAttractionList = "https://m.ctrip.com/restapi/soa2/18254/json/getAttractionList"
const GetPoimoreDetail = "https://m.ctrip.com/restapi/soa2/18254/json/getPoiMoreDetail"

func QueryCtripCityList() (ccList []model.CtripCity) {
	data, _ := ioutil.ReadFile("./ctripCityList.html")
	str2 := string(data)
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(str2))
	doc.Find(".c-citylist-parent").Each(func(i int, s *goquery.Selection) {
		p := s.Find("span").Text()
		a := s.Find("li")
		a.Each(func(ii int, ss *goquery.Selection) {
			b, _ := ss.Attr("data-id")
			c, _ := ss.Attr("data-name")
			var cc model.CtripCity
			cc.CtripCityCode = p
			cc.CtripCityQuery = b
			cc.CtripCityName = c
			ccList = append(ccList, cc)
		})
	})
	return
}

/**
  		{
       "index": 1,
       "count": 40,
       "sortType": 1,
       "isShowAggregation": true,
       "districtId": 216,
       "scene": "DISTRICT",
       "pageId": "sight_list",
       "traceId": "ea006adc-ce22-7ec6-92fd-ebb171153340",
       "extension": [
           {
               "name": "osVersion",
               "value": "11.2.2"
           },
           {
               "name": "deviceType",
               "value": "ios"
           }
       ],
       "filter": {
           "filterItems": []
       },
       "crnVersion": "2020-09-01 22:00:45",
       "isInitialState": true,
       "head": {
           "cid": "09031066113126275485",
           "ctok": "",
           "cver": "1.0",
           "lang": "01",
           "sid": "8888",
           "syscode": "09",
           "auth": "",
           "xsid": "",
           "extension": []
       }
    }
*/

func GetAttractionListFunc(districtId int, ctripCityId int) (ctripAttractionList []model.CtripAttraction) {
	a := getAttractionListSubFunc(districtId, 1)
	if a.Result == 0 {
		c := createCtripAttraction(a, districtId, ctripCityId)
		ctripAttractionList = append(ctripAttractionList, c...)
		count := a.TotalCount/20 + 2
		for i := 2; i < count; i++ {
			b := getAttractionListSubFunc(districtId, i)
			d := createCtripAttraction(b, districtId, ctripCityId)
			ctripAttractionList = append(ctripAttractionList, d...)
		}
	}
	return
}

func createCtripAttraction(a dto.Attraction, districtId int, ctripCityId int) (ctripAttractionList []model.CtripAttraction) {
	details := a.AttractionList
	for _, detail := range details {
		var ca = model.CtripAttraction{
			PoiId:         detail.Card.PoiId,
			BusinessId:    detail.Card.BusinessId,
			ZoneName:      detail.Card.ZoneName,
			PoiName:       detail.Card.PoiName,
			CoverImageUrl: detail.Card.CoverImageUrl,
			DistanceStr:   detail.Card.DistanceStr,
			Price:         detail.Card.Price,
			IsFree:        detail.Card.IsFree,
			OpenStatus:    detail.Card.OpenStatus,
			CtripCityId:   ctripCityId,
			DistrictId:    districtId,
		}
		ctripAttractionList = append(ctripAttractionList, ca)
	}
	return
}

func getAttractionListSubFunc(districtId int, index int) (a dto.Attraction) {
	var (
		ctripCityAttraction = dto.CtripCityAttraction{
			Index:             index,
			Count:             20,
			SortType:          1,
			IsShowAggregation: true,
			DistrictId:        districtId,
			Scene:             "DISTRICT",
			PageId:            "sight_list",
			TraceId:           "ea006adc-ce22-7ec6-92fd-ebb171153340",
			Extension:         []dto.Extension{{Name: "osVersion", Value: "11.2.2"}, {Name: "deviceType", Value: "ios"}},
			Filter:            dto.Filter{FilterItems: []dto.FilterItems{}},
			CrnVersion:        "2020-09-01 22:00:45",
			IsInitialState:    true,
			Head: dto.Head{
				Cid:       "09031066113126275485",
				Ctok:      "",
				Cver:      "1.0",
				Lang:      "01",
				Sid:       "8888",
				Syscode:   "09",
				Auth:      "",
				Xsid:      "",
				Extension: []dto.Extension{},
			},
		}
	)
	jsons, _ := json.Marshal(ctripCityAttraction)
	m := utils.SamplePost(GetAttractionList, jsons)
	if err := mapstructure.Decode(m, &a); err != nil {
		fmt.Println(err)
	}
	return
}

/**
{
    "poiId": 23078384,
    "scene": "basic",
    "head": {
        "cid": "09031066113126275485",
        "ctok": "",
        "cver": "1.0",
        "lang": "01",
        "sid": "8888",
        "syscode": "09",
        "auth": "",
        "xsid": "",
        "extension": []
    }
}
*/
func GetPoiMoreDetailFunc(poiId int) (m map[string]interface{}) {
	var (
		ctripCityTemplate = dto.CtripCityTemplate{
			PoiId: poiId,
			Scene: "basic",
			Head: dto.Head{
				Cid:       "09031066113126275485",
				Ctok:      "",
				Cver:      "1.0",
				Lang:      "01",
				Sid:       "8888",
				Syscode:   "09",
				Auth:      "",
				Xsid:      "",
				Extension: []dto.Extension{},
			},
		}
	)
	jsons, _ := json.Marshal(ctripCityTemplate)
	m = utils.SamplePost(GetPoimoreDetail, jsons)
	return
}
