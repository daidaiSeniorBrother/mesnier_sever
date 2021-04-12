package model

import (
	"gorm.io/gorm"
	"time"
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
		b.CreateTime = time.Now().Format("2006-01-02 15:04:05")
		b.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
	}
	return nil
}

func (b *Base) BeforeUpdate(*gorm.DB) error {
	b.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
	return nil
}
