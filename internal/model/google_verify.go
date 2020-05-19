package model

import "github.com/jinzhu/gorm"

type GoogleVerify struct {
	gorm.Model
	UserId    uint   `gorm:"not null;index:idx_user_id"` //用户id
	Secret    string `gorm:"not null"`                   //用户密钥
	QrCodeUrl string `gorm:"not null"`                   // 打印二维码地址
}
