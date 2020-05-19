package model

import "github.com/jinzhu/gorm"

//公告
type Notice struct {
	gorm.Model
	Title   string `gorm:"not null"`      //标题
	Content string `gorm:"type:longtext"` //内容
}
