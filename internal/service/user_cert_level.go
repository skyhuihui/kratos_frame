package service

import (
	"context"
	"kratos_frame/internal/model"
	err_code "kratos_frame/internal/pkg/ecode"

	"github.com/bilibili/kratos/pkg/ecode"
)

// 添加认证级别
func (s *Service) InsertCertLevel(ctx context.Context, certLevel model.UserCertLevel) error {
	certLevelI, err := s.dao.Insert(ctx, model.Tb+"user_cert_level", &certLevel)
	if err != nil || certLevelI == nil || certLevelI.(*model.UserCertLevel).ID == 0 {
		return err_code.CreatUserErr
	}
	return nil
}

// 修改认证级别
func (s *Service) UpdateCertLevel(ctx context.Context, whereCertLevel model.UserCertLevel, certLevel model.UserCertLevel) error {
	e, i := s.dao.Update(ctx, model.Tb+"user_cert_level", whereCertLevel, certLevel)
	if e != nil {
		return e
	} else if i == 0 {
		return err_code.UpdateErr
	}
	return nil
}

// 查询认证级别
func (s *Service) FindCertLevel(ctx context.Context, params model.FindCertLevelParams) ([]model.UserCertLevel, int, error) {
	var certLevel model.UserCertLevel
	if params.UserId != 0 {
		certLevel.UserId = params.UserId
	}
	var certLevels []model.UserCertLevel
	total := 0
	err := s.dao.Find(ctx, model.Tb+"user_cert_level", certLevel).Where("deleted_at IS NULL").Count(&total).
		Limit(params.PageSize).Offset((params.Page - 1) * params.PageSize).Find(&certLevels).Error
	if err != nil {
		return nil, 0, err_code.FindErr
	}
	return certLevels, total, err_code.Success
}

func (s *Service) FindCertLevelByUser(ctx context.Context, userId uint) (model.UserCertLevel, ecode.Code) {
	var certLevels model.UserCertLevel
	err := s.dao.Find(ctx, model.Tb+"user_cert_level", model.UserCertLevel{UserId: userId}).First(&certLevels).Error
	if err != nil {
		return certLevels, err_code.FindErr
	}
	return certLevels, err_code.Success
}
