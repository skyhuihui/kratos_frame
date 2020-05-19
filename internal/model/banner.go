package model

import "github.com/jinzhu/gorm"

type Banner struct {
	gorm.Model
	Img string `gorm:"not null;size:500"` //banner 图地址
}
