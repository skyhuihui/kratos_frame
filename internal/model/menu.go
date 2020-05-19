package model

import "github.com/jinzhu/gorm"

//菜单表
type Menu struct {
	gorm.Model
	Mid        string      `gorm:"not null"` // 菜单随机id （前端根据id 来找页面 路径 按钮之类）
	SuperiorId *uint       `gorm:"not null"` // 上级菜单id 没有上级菜单 为 0
	Name       string      `gorm:"not null"` // 菜单名
	Type       *int        `gorm:"not null"` // 菜单类型 0 菜单 1 路径 2 按钮
	Path       string      `gorm:"not null"` // 访问路径
	Method     string      `gorm:"not null"` // 资源请求方式
	Menus      interface{} `gorm:"-"`        // 忽略本字段 查询时组合下级数据
}
