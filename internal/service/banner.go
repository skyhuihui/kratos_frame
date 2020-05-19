package service

import (
	"context"
	"kratos_frame/internal/model"
	err_code "kratos_frame/internal/pkg/ecode"

	"github.com/bilibili/kratos/pkg/ecode"
	"github.com/jinzhu/gorm"
)

func (s *Service) InsertBanner(ctx context.Context, params model.InsertBannerParams) ecode.Code {
	// 参数验证
	userCertification, err := s.dao.Insert(ctx, model.Tb+"banner", &model.Banner{
		Img: params.Img,
	})
	if err != nil || userCertification.(*model.Banner).ID == 0 {
		return err_code.InsertErr
	}

	return err_code.Success
}

func (s *Service) FindBanner(ctx context.Context, params model.PageParams) ([]model.Banner, int, ecode.Code) {
	// 参数校验
	var banner []model.Banner
	total := 0
	err := s.dao.Find(ctx, model.Tb+"banner", model.Banner{}).Where("deleted_at IS NULL").Count(&total).
		Limit(params.PageSize).Offset((params.Page - 1) * params.PageSize).Find(&banner).Error
	if err != nil {
		return nil, 0, err_code.FindErr
	}
	return banner, total, err_code.Success
}

func (s *Service) DeleteBanner(ctx context.Context, params model.DeleteBannerParams) ecode.Code {
	e, i := s.dao.Delete(ctx, model.Tb+"banner", model.Banner{
		Model: gorm.Model{ID: params.BannerId},
	})
	if e != nil {
		return err_code.DeleteErr
	} else if i == 0 {
		return err_code.NoData
	}

	return err_code.Success
}
