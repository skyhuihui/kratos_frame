package http

import (
	"kratos_frame/internal/model"
	err_code "kratos_frame/internal/pkg/ecode"

	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

func FindUserInfo(c *bm.Context) {
	var params model.FindUserParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	data, err := svc.FindUserInfo(c.Context, params)
	c.JSON(data, err)
	return
}

func FindUser(c *bm.Context) {
	var params model.FindUserParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	data, total, err := svc.FindUser(c.Context, params)
	RespJson(c, data, total, err)
	return
}

func Signup(c *bm.Context) {
	var params model.SignupParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	c.JSON(nil, svc.Signup(c.Context, params))
	return
}

func Login(c *bm.Context) {
	var params model.LoginParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	user, err := svc.Login(c.Context, params)
	c.JSON(user, err)
	return
}

func Logout(c *bm.Context) {
	var params model.LogoutParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	c.JSON(nil, svc.Logout(c.Context, params))
	return
}

func ForgetPassword(c *bm.Context) {
	var params model.ForgetPasswordParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	c.JSON(nil, svc.ForgetPassword(c.Context, params))
	return
}
func UpdatePassword(c *bm.Context) {
	var params model.UpdatePasswordParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	c.JSON(nil, svc.UpdatePassword(c.Context, params))
	return
}

func UpdateFundPassword(c *bm.Context) {
	var params model.UpdateFundPasswordParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	c.JSON(nil, svc.UpdateFundPassword(c.Context, params))
	return
}

// 查看认证级别
func FindCertLevel(c *bm.Context) {
	var params model.FindCertLevelParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	data, total, err := svc.FindCertLevel(c.Context, params)
	RespJson(c, data, total, err)
	return
}

// 查看用户一级下级
func FindUserInviteSubordinate(c *bm.Context) {
	var params model.FindUserInviteParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	data, total, err := svc.FindUserInviteSubordinate(c.Context, params)
	RespJson(c, data, total, err)
	return
}

// 查看用户邀请信息
func FindUserInvite(c *bm.Context) {
	var params model.FindUserInviteParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	data, total, err := svc.FindUserInvite(c.Context, params)
	RespJson(c, data, total, err)
	return
}

// 添加实名认证信息
func InsertCertification(c *bm.Context) {
	var params model.InsertCertificationParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	c.JSON(nil, svc.InsertCertification(c.Context, params))
	return
}

// 查看实名认证信息
func FindCertification(c *bm.Context) {
	var params model.FindCertificationParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	data, total, err := svc.FindCertification(c.Context, params)
	RespJson(c, data, total, err)
	return
}

// 修改实名认证信息
func UpdateCertification(c *bm.Context) {
	var params model.UpdateCertificationParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	c.JSON(nil, svc.UpdateCertification(c.Context, params))
	return
}

// 删除实名认证信息
func DeleteCertification(c *bm.Context) {
	var params model.DeleteCertificationParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	c.JSON(nil, svc.DeleteCertification(c.Context, params))
	return
}

// 给用户分配角色
func InsertRoleUser(c *bm.Context) {
	var params model.InsertRoleByUserParams
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	c.JSON(nil, svc.InsertRoleUser(c.Context, params))
	return
}
