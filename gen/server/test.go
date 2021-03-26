package http

import (
	"kratos_frame/internal/model"
	err_code "kratos_frame/internal/pkg/ecode"

	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

func InsertTest(c *bm.Context) {
	var params model.InsertTestParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	c.JSON(nil, svc.InsertTest(c.Context, params))
	return
}

func FindTest(c *bm.Context) {
	var params model.FindTestParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	data, total, err := svc.FindTest(c.Context, params)
	RespJson(c, data, total, err)
	return
}

func DeleteTest(c *bm.Context) {
	var params model.DeleteTestParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	c.JSON(nil, svc.DeleteTest(c.Context, params))
	return
}
func UpdateTest(c *bm.Context) {
	var params model.UpdateTestParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	c.JSON(nil, svc.UpdateTest(c.Context, params))
	return
}
