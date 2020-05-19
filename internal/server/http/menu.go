package http

import (
	"kratos_frame/internal/model"
	err_code "kratos_frame/internal/pkg/ecode"

	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

func InsertMenu(c *bm.Context) {
	var params model.InsertMenuParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	c.JSON(nil, svc.InsertMenu(c.Context, params))
	return
}

func FindMenu(c *bm.Context) {
	var params model.FindMenuParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	data, total, err := svc.FindMenu(c.Context, params)
	RespJson(c, data, total, err)
	return
}

func DeleteMenu(c *bm.Context) {
	var params model.DeleteMenuParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	c.JSON(nil, svc.DeleteMenu(c.Context, params))
	return
}
func UpdateMenu(c *bm.Context) {
	var params model.UpdateMenuParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	c.JSON(nil, svc.UpdateMenu(c.Context, params))
	return
}
