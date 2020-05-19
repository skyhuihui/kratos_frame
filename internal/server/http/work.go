package http

import (
	"kratos_frame/internal/model"
	err_code "kratos_frame/internal/pkg/ecode"

	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

func InsertWork(c *bm.Context) {
	var params model.InsertWorkParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	c.JSON(nil, svc.InsertWork(c.Context, params))
	return
}

func FindWork(c *bm.Context) {
	var params model.FindWorkParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	data, total, err := svc.FindWork(c.Context, params)
	RespJson(c, data, total, err)
	return
}

func DeleteWork(c *bm.Context) {
	var params model.DeleteWorkParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	c.JSON(nil, svc.DeleteWork(c.Context, params))
	return
}
func UpdateWork(c *bm.Context) {
	var params model.UpdateWorkParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	c.JSON(nil, svc.UpdateWork(c.Context, params))
	return
}
