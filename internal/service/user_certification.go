package service

import (
	"context"
	"kratos_frame/internal/model"
	err_code "kratos_frame/internal/pkg/ecode"

	"github.com/bilibili/kratos/pkg/ecode"
	"github.com/jinzhu/gorm"
)

// 新增实名认证
func (s *Service) InsertCertification(ctx context.Context, params model.InsertCertificationParams) ecode.Code {

	if eCode := s.VerifyInsertCertification(ctx, params.UserId, params.Cert); eCode.Code() != 200 {
		return eCode
	}
	var identity string
	for _, v := range params.Identity {
		identity += v + ","
	}
	userCertification, err := s.dao.Insert(ctx, model.Tb+"user_certification", &model.UserCertification{
		UserId:   params.UserId,
		RealName: params.RealName,
		Cert:     params.Cert,
		Address:  params.Address,
		Exp:      params.Exp,
		Issue:    params.Issue,
		Identity: identity[0 : len(identity)-1],
	})
	if userCertification.(*model.UserCertification).ID == 0 || err != nil {
		return err_code.InsertErr
	}

	// 处理认证级别
	if s.UpdateCertLevel(ctx, model.UserCertLevel{
		UserId: params.UserId,
	}, model.UserCertLevel{
		IsReal: &model.One,
	}) != nil {
		return err_code.UpdateErr
	}

	return err_code.Success
}

// 查询实名认证
func (s *Service) FindCertification(ctx context.Context, params model.FindCertificationParams) ([]model.UserCertification, int, ecode.Code) {

	var certification model.UserCertification
	if params.UserId != 0 {
		certification.UserId = params.UserId
	}

	var certifications []model.UserCertification
	total := 0
	err := s.dao.Find(ctx, model.Tb+"user_certification", certification).Where("deleted_at IS NULL").Count(&total).
		Limit(params.PageSize).Offset((params.Page - 1) * params.PageSize).Find(&certifications).Error
	if err != nil {
		return nil, 0, err_code.FindErr
	}
	return certifications, total, err_code.Success
}

// 修改实名认证
func (s *Service) UpdateCertification(ctx context.Context, params model.UpdateCertificationParams) ecode.Code {
	var identity string
	for _, v := range params.Identity {
		identity += v + ","
	}
	e, i := s.dao.Update(ctx, model.Tb+"user_certification", model.UserCertification{
		Model: gorm.Model{ID: params.CertificationId},
	}, model.UserCertification{
		RealName: params.RealName,
		Cert:     params.Cert,
		Address:  params.Address,
		Exp:      params.Exp,
		Issue:    params.Issue,
		Identity: identity[0 : len(identity)-1],
	})
	if e != nil {
		return err_code.UpdateErr
	} else if i == 0 {
		return err_code.NoData
	}

	var cert []model.UserCertification
	s.dao.Find(ctx, model.Tb+"user_certification", model.UserCertification{
		Model: gorm.Model{ID: params.CertificationId},
	}).Find(&cert)

	return err_code.Success
}

// 删除实名认证
func (s *Service) DeleteCertification(ctx context.Context, params model.DeleteCertificationParams) ecode.Code {
	var certifications model.UserCertification
	err := s.dao.Find(ctx, model.Tb+"user_certification", model.UserCertification{Model: gorm.Model{ID: params.CertificationId}}).First(&certifications).Error
	if err != nil {
		return err_code.DeleteErr
	}

	e, i := s.dao.Delete(ctx, model.Tb+"user_certification", model.UserCertification{
		Model: gorm.Model{ID: params.CertificationId},
	})
	if e != nil {
		return err_code.DeleteErr
	} else if i == 0 {
		return err_code.NoData
	}

	// 删除后实名状态设置未实名

	if s.UpdateCertLevel(ctx, model.UserCertLevel{
		UserId: certifications.UserId,
	}, model.UserCertLevel{
		IsReal: &model.Zero,
	}) != nil {
		return err_code.UpdateErr
	}

	return err_code.Success
}
