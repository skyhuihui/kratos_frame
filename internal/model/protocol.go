package model

import "github.com/jinzhu/gorm"

// 用户协议
type Protocol struct {
	gorm.Model
	Title   string `gorm:"not null"`               // 协议标题
	Content string `gorm:"type:longtext;not null"` // 协议内容
	Type    *int   `gorm:"not null"`               // 协议类型 0 注册协议
}
