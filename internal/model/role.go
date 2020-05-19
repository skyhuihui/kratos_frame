package model

import "github.com/jinzhu/gorm"

// 角色表
type Role struct {
	gorm.Model
	Name    string `gorm:"not null"` // 角色名称
	Content string `gorm:"not null"` // 角色说明
}
