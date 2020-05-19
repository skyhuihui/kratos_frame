package service

import (
	"context"
	"kratos_frame/internal/model"
	err_code "kratos_frame/internal/pkg/ecode"
	"kratos_frame/internal/pkg/sha256"
	"time"

	"github.com/bilibili/kratos/pkg/log"

	"github.com/jinzhu/gorm"

	"github.com/bilibili/kratos/pkg/ecode"
	"github.com/gomodule/redigo/redis"
)

func (s *Service) VerifySms(ctx context.Context, mobile, mobileCode string) ecode.Code {
	// 参数验证
	code, err := redis.String(s.dao.RedisGet(ctx, mobile+"_sms"))
	if err != nil {
		log.Error("短信验证码验证失败：%s", err)
		return err_code.SmsCodeTimeOut
	}
	if code != mobileCode {
		return err_code.VerifySmsErr
	}
	if _, err = s.dao.RedisDel(ctx, mobile+"_sms"); err != nil {
		return err_code.VerifySmsErr
	}
	return err_code.Success
}

func (s *Service) VerifyFundPass(ctx context.Context, userId uint, fundPass string) (model.User, ecode.Code) {
	var user model.User
	if err := s.dao.Find(ctx, model.Tb+"user", model.User{
		Model:    gorm.Model{ID: userId},
		FundPass: sha256.Sha(fundPass),
	}).First(&user).Error; err != nil {
		return user, err_code.FindErr
	}
	if user.ID == 0 {
		return user, err_code.UpdateErr
	}
	return user, err_code.Success
}

// 注册验证上级
func (s *Service) VerifySignUpSuperior(ctx context.Context, registerCode string) ([]model.User, ecode.Code) {
	var superiorUsers []model.User
	if registerCode != "" {
		if err := s.dao.Find(ctx, model.Tb+"user", model.User{InviteCode: registerCode}).Find(&superiorUsers).Error; err != nil || len(superiorUsers) == 0 {
			return superiorUsers, err_code.RegisterCodeErr
		}
	}

	return superiorUsers, err_code.Success
}

func (s *Service) VerifyInsertCertification(ctx context.Context, userId uint, cert string) ecode.Code {
	var certifications model.UserCertification
	err := s.dao.Find(ctx, model.Tb+"user_certification", model.UserCertification{UserId: userId}).First(&certifications).Error
	if (err != nil && err != gorm.ErrRecordNotFound) || certifications.ID != 0 {
		return err_code.InsertErr
	}
	err = s.dao.Find(ctx, model.Tb+"user_certification", model.UserCertification{Cert: cert}).First(&certifications).Error
	if (err != nil && err != gorm.ErrRecordNotFound) || certifications.ID != 0 {
		return err_code.InsertErr
	}
	return err_code.Success
}

// 处理区间时间
func (s *Service) ProcessingTime(between, and int64) (start string, end string) {
	if between == 0 && and == 0 {
		end = time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
	} else {
		end = time.Unix(and, 0).Format("2006-01-02 15:04:05")
	}
	start = time.Unix(between, 0).Format("2006-01-02 15:04:05")
	return
}
