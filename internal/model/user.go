package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Uid        string `gorm:"not null"` // 6位随机id
	Name       string
	Password   string `gorm:"not null"`
	Email      string
	Mobile     string `gorm:"not null;unique_index:idx_mobile"`
	InviteCode string // 邀请码
	Status     *int   `gorm:"default 0;not null"` // 是否禁用 0 启用 1 禁用
	FundPass   string // 六位数字资金密码
}

//实名信息
type UserCertification struct {
	gorm.Model
	UserId   uint   `gorm:"not null;index:idx_user_id"`
	RealName string `gorm:"not null"`           // 实名姓名
	Cert     string `gorm:"not null"`           //身份证号
	Address  string `gorm:"not null"`           //身份证地址
	Exp      string `gorm:"not null"`           //有效期
	Issue    string `gorm:"not null"`           //签发机关
	Identity string `gorm:"not null;size:5000"` //身份证图片
}

//上下级 邀请码
type UserInvite struct {
	gorm.Model
	UserId     uint `gorm:"not null;unique_index:idx_user_super"` //用户id
	SuperiorId uint `gorm:"not null;unique_index:idx_user_super"` // 上级id
	Level      int  `gorm:"not null"`                             // 层级
}

// 认证情况
type UserCertLevel struct {
	gorm.Model
	UserId       uint `gorm:"not null;index:idx_user_id"` //用户id
	IsReal       *int `gorm:"default 0;not null"`         //是否实名  0 未实名 1 已实名 2 实名中
	IsEmail      *int `gorm:"default 0;not null"`         //是否认证邮箱
	IsMobile     *int `gorm:"default 0;not null"`         //是否认证手机号
	IsGooGleCode *int `gorm:"default 0;not null"`         //是否完善谷歌验证器
}

// 用户身份对应关系
type UserRole struct {
	gorm.Model
	UserId uint `gorm:"not null;index:idx_user_id"` // 用户id
	RoleId uint `gorm:"not null"`                   // 身份id
}
