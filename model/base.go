package model

import (
	"gorm.io/gorm"
	"mesnier/utils"
)

type PageResponse struct {
	ObjectData interface{} `json:"objectData"`
	Total      int64       `json:"total"`
}

type Base struct {
	CreateTime string `gorm:"column:create_time" json:"create_time"`
	UpdateTime string `gorm:"column:update_time" json:"update_time"`
	DeleteFlag int    `gorm:"column:delete_flag" json:"delete_flag"`
}

func (b *Base) BeforeCreate(*gorm.DB) error {
	if len(b.CreateTime) == 0 {
		b.CreateTime = utils.TimeNowStr()
		b.UpdateTime = utils.TimeNowStr()
	}
	return nil
}

func (b *Base) BeforeUpdate(*gorm.DB) error {
	b.UpdateTime = utils.TimeNowStr()
	return nil
}
