package model

type CasbinRule struct {
	PType string
	V0    string // 角色
	V1    string // 操作路径
	V2    string // 操作方式
	V3    string
	V4    string
	V5    string
}

// 1、给角色分配菜单 （角色菜单对应表）
// 2. 删除角色时 删除casbin表的所有此角色的记录
