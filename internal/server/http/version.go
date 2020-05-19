package http

import (
	"kratos_frame/internal/model"
	err_code "kratos_frame/internal/pkg/ecode"

	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

func InsertVersion(c *bm.Context) {
	var params model.InsertVersionParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	c.JSON(nil, svc.InsertVersion(c.Context, params))
	return
}

func FindVersion(c *bm.Context) {
	var params model.FindVersionParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	data, total, err := svc.FindVersion(c.Context, params)
	RespJson(c, data, total, err)
	return
}

func DeleteVersion(c *bm.Context) {
	var params model.DeleteVersionParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	c.JSON(nil, svc.DeleteVersion(c.Context, params))
	return
}
func UpdateVersion(c *bm.Context) {
	var params model.UpdateVersionParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	c.JSON(nil, svc.UpdateVersion(c.Context, params))
	return
}
