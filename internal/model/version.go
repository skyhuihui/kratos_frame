package model

import "github.com/jinzhu/gorm"

// 应用版本更新
type Version struct {
	gorm.Model
	Num        string `gorm:"not null"`               // 版本号
	Content    string `gorm:"type:longtext;not null"` // 更新说明
	Link       string `gorm:"not null;size:500"`      //下载链接
	Type       *int   `gorm:"not null"`               //所属类型 0 安卓 1 ios
	Constraint *int   `gorm:"not null"`               // 是否强制更新
}
