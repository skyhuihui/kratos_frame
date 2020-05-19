package http

import (
	"kratos_frame/internal/model"
	err_code "kratos_frame/internal/pkg/ecode"

	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

func FindGoogleVerify(c *bm.Context) {
	var params model.FindGoogleVerifyParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	data, err := svc.FindGoogleVerify(c.Context, params)
	c.JSON(data, err)
	return
}

func VerifyGoogleVerify(c *bm.Context) {
	var params model.VerifyGoogleVerifyParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	c.JSON(nil, svc.VerifyGoogleVerify(c.Context, params))
	return
}

func DeleteUserGoogleVerify(c *bm.Context) {
	var params model.DeleteUserGoogleVerifyParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	c.JSON(nil, svc.DeleteUserGoogleVerify(c.Context, params))
	return
}
