package query

import (
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"mesnier/model"
	"strings"
)

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
