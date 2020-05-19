package model

import "github.com/jinzhu/gorm"

// 工单
type Work struct {
	gorm.Model
	UserId  uint   `gorm:"not null;index:idx_user_id"`
	Mobile  string `gorm:"not null"`           // 工单联系人
	Title   string `gorm:"not null"`           // 标题
	Content string `gorm:"not null;size:5000"` // 问题描述
	Img     string `gorm:"size:5000"`          // 工单图片， 多个英文逗号隔开
	Type    *int   `gorm:"default 0;not null"` // 状态 0 待回复， 1 已回复 ， 2 已结束
}
