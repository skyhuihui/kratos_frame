package service

import (
	"context"
	"kratos_frame/internal/model"
	err_code "kratos_frame/internal/pkg/ecode"
	GoogleVerify "kratos_frame/internal/pkg/google_verify"

	"github.com/bilibili/kratos/pkg/ecode"
	"github.com/gomodule/redigo/redis"
)

func (s *Service) FindGoogleVerify(ctx context.Context, params model.FindGoogleVerifyParams) (model.GoogleVerify, ecode.Code) {
	// 参数校验
	var verify model.GoogleVerify
	err := s.dao.Find(ctx, model.Tb+"google_verify", model.GoogleVerify{UserId: params.UserId}).First(&verify).Error
	if err != nil && err.Error() != "record not found" {
		return model.GoogleVerify{}, err_code.FindErr
	}
	if verify.ID == 0 {
		googleVerify := GoogleVerify.InitAuth("kratos_frame")
		verify.UserId = params.UserId
		verify.Secret = googleVerify.Secret
		verify.QrCodeUrl = googleVerify.QrCodeUrl
		v, err := s.dao.Insert(ctx, model.Tb+"google_verify", &verify)
		if err != nil || v.(*model.GoogleVerify).ID == 0 {
			return model.GoogleVerify{}, err_code.InsertErr
		}
	}
	return verify, err_code.Success
}

func (s *Service) VerifyGoogleVerify(ctx context.Context, params model.VerifyGoogleVerifyParams) ecode.Code {
	user := s.FindUserById(ctx, params.UserId)
	if params.EmailCode != "" {
		emailCode, err := redis.String(s.dao.RedisGet(ctx, user.Email+"_email"))
		if err != nil {
			return err_code.EmailCodeTimeOut
		}
		if emailCode != params.EmailCode {
			return err_code.VerifyEmailErr
		}
		if _, err = s.dao.RedisDel(ctx, user.Email+"_email"); err != nil {
			return err_code.InsertErr
		}
	}

	if params.PhoneCode != "" {
		phoneCode, err := redis.String(s.dao.RedisGet(ctx, user.Mobile+"_sms"))
		if err != nil {
			return err_code.SmsCodeTimeOut
		}
		if phoneCode != params.PhoneCode {
			return err_code.VerifySmsErr
		}
		if _, err = s.dao.RedisDel(ctx, user.Mobile+"_sms"); err != nil {
			return err_code.InsertErr
		}
	}
	// 参数校验
	var verify model.GoogleVerify
	err := s.dao.Find(ctx, model.Tb+"google_verify", model.GoogleVerify{UserId: params.UserId}).First(&verify).Error
	if err != nil {
		return err_code.InsertErr
	}

	// 验证谷歌验证
	code, err := GoogleVerify.NewGoogleAuth().GetCode(verify.Secret)
	if err != nil {
		return err_code.InsertErr
	}
	if code != params.GoogleCode {
		return err_code.GoogleVerifyErr
	}

	if s.UpdateCertLevel(ctx, model.UserCertLevel{
		UserId: params.UserId,
	}, model.UserCertLevel{
		IsGooGleCode: &model.One,
	}) != nil {
		return err_code.InsertErr
	}

	return err_code.Success
}

// 删除用户谷歌认证
func (s *Service) DeleteUserGoogleVerify(ctx context.Context, params model.DeleteUserGoogleVerifyParams) ecode.Code {

	if s.UpdateCertLevel(ctx, model.UserCertLevel{
		UserId: params.UserId,
	}, model.UserCertLevel{
		IsGooGleCode: &model.Zero,
	}) != nil {
		return err_code.InsertErr
	}
	return err_code.Success
}
