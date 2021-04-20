package confs

import (
	"fmt"
	"mesnier/model"
	"mesnier/query"
	"strconv"
)

func CtripSetUp() {
	if VIPER.GetBool("query.ctrip") {
		initCrtipData()
		InitCrtipAttractionData()
	}
}

func initCrtipData() {
	var (
		cc           []model.CtripCity
		ccCreateList []model.CtripCity
		ccUpdateList []model.CtripCity
		ccMap        = make(map[string]string)
		ccIdMap      = make(map[string]int)
		ccList       = query.QueryCtripCityList()
		err          error
	)
	DB.Find(&cc)
	for _, city := range cc {
		ccMap[city.CtripCityName] = city.CtripCityQuery
		ccIdMap[city.CtripCityName] = city.CtripCityId
	}
	for _, city := range ccList {
		p, ok := ccMap[city.CtripCityName]
		if ok {
			if p != city.CtripCityQuery {
				city.CtripCityId = ccIdMap[city.CtripCityName]
				ccUpdateList = append(ccUpdateList, city)
			}
		} else {
			ccCreateList = append(ccCreateList, city)
		}
	}
	if len(ccCreateList) > 0 {
		if err = DB.Create(&ccCreateList).Error; err != nil {
			fmt.Print(err)
		}
	}
	for _, city := range ccUpdateList {
		DB.Updates(city)
	}
}

func InitCrtipAttractionData() {
	var ctripCity []model.CtripCity
	DB.Find(&ctripCity)
	for _, city := range ctripCity {
		i, _ := strconv.Atoi(city.CtripCityQuery)
		ctripAttractionList := query.GetAttractionListFunc(i, city.CtripCityId)
		if err := DB.Create(&ctripAttractionList).Error; err != nil {
			fmt.Print(err)
		}
	}
}
