package http

import (
	"kratos_frame/internal/model"
	err_code "kratos_frame/internal/pkg/ecode"

	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

// 角色和菜单绑定
func InsertRoleMenu(c *bm.Context) {
	var params model.InsertRoleMenuParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	c.JSON(nil, svc.InsertRoleMenu(c.Context, params))
	return
}

func InsertRole(c *bm.Context) {
	var params model.InsertRoleParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	c.JSON(nil, svc.InsertRole(c.Context, params))
	return
}

func FindRole(c *bm.Context) {
	var params model.FindRoleParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	data, total, err := svc.FindRole(c.Context, params)
	RespJson(c, data, total, err)
	return
}

func DeleteRole(c *bm.Context) {
	var params model.DeleteRoleParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	c.JSON(nil, svc.DeleteRole(c.Context, params))
	return
}
func UpdateRole(c *bm.Context) {
	var params model.UpdateRoleParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	c.JSON(nil, svc.UpdateRole(c.Context, params))
	return
}
