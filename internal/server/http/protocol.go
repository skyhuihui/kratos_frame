package http

import (
	"kratos_frame/internal/model"
	err_code "kratos_frame/internal/pkg/ecode"

	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

func InsertProtocol(c *bm.Context) {
	var params model.InsertProtocolParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	c.JSON(nil, svc.InsertProtocol(c.Context, params))
	return
}

func FindProtocol(c *bm.Context) {
	var params model.FindProtocolParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	data, total, err := svc.FindProtocol(c.Context, params)
	RespJson(c, data, total, err)
	return
}

func DeleteProtocol(c *bm.Context) {
	var params model.DeleteProtocolParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	c.JSON(nil, svc.DeleteProtocol(c.Context, params))
	return
}
func UpdateProtocol(c *bm.Context) {
	var params model.UpdateProtocolParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	c.JSON(nil, svc.UpdateProtocol(c.Context, params))
	return
}
