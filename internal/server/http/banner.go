package http

import (
	"kratos_frame/internal/model"
	err_code "kratos_frame/internal/pkg/ecode"

	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

func InsertBanner(c *bm.Context) {
	var params model.InsertBannerParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	c.JSON(nil, svc.InsertBanner(c.Context, params))
	return
}

func FindBanner(c *bm.Context) {
	var params model.PageParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	// 参数验证， page > 0, page_size > 0
	data, total, err := svc.FindBanner(c.Context, params)
	RespJson(c, data, total, err)
	return
}

func DeleteBanner(c *bm.Context) {
	var params model.DeleteBannerParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	c.JSON(nil, svc.DeleteBanner(c.Context, params))
	return
}
