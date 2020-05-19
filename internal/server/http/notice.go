package http

import (
	"kratos_frame/internal/model"
	err_code "kratos_frame/internal/pkg/ecode"

	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

func InsertNotice(c *bm.Context) {
	var params model.InsertNoticeParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	c.JSON(nil, svc.InsertNotice(c.Context, params))
	return
}

func FindNotice(c *bm.Context) {
	var params model.PageParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	data, total, err := svc.FindNotice(c.Context, params)
	RespJson(c, data, total, err)
	return
}

func DeleteNotice(c *bm.Context) {
	var params model.DeleteNoticeParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	c.JSON(nil, svc.DeleteNotice(c.Context, params))
	return
}
func UpdateNotice(c *bm.Context) {
	var params model.UpdateNoticeParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	c.JSON(nil, svc.UpdateNotice(c.Context, params))
	return
}
