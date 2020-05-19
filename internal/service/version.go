package service

import (
	"context"
	"kratos_frame/internal/model"
	err_code "kratos_frame/internal/pkg/ecode"

	"github.com/bilibili/kratos/pkg/ecode"
	"github.com/jinzhu/gorm"
)

func (s *Service) InsertVersion(ctx context.Context, params model.InsertVersionParams) ecode.Code {
	// 参数验证
	userCertification, err := s.dao.Insert(ctx, model.Tb+"version", &model.Version{
		Num:        params.Num,
		Content:    params.Content,
		Link:       params.Link,
		Type:       &params.Type,
		Constraint: &params.Constraint,
	})
	if userCertification.(*model.Version).ID == 0 || err != nil {
		return err_code.InsertErr
	}

	return err_code.Success
}

func (s *Service) FindVersion(ctx context.Context, params model.FindVersionParams) ([]model.Version, int, ecode.Code) {
	var version model.Version
	if params.Type != -1 {
		version.Type = &params.Type
	}

	var versions []model.Version
	total := 0
	if err := s.dao.Find(ctx, model.Tb+"version", version).Where("deleted_at IS NULL").Count(&total).Order("id desc").
		Limit(params.PageSize).Offset((params.Page - 1) * params.PageSize).Find(&versions).Error; err != nil {
		return nil, 0, err_code.FindErr
	}
	return versions, total, err_code.Success
}

func (s *Service) DeleteVersion(ctx context.Context, params model.DeleteVersionParams) ecode.Code {
	e, i := s.dao.Delete(ctx, model.Tb+"version", model.Version{
		Model: gorm.Model{ID: params.VersionId},
	})
	if e != nil {
		return err_code.DeleteErr
	} else if i == 0 {
		return err_code.NoData
	}

	return err_code.Success
}

func (s *Service) UpdateVersion(ctx context.Context, params model.UpdateVersionParams) ecode.Code {
	e, i := s.dao.Update(ctx, model.Tb+"version", model.Version{
		Model: gorm.Model{ID: params.VersionId},
	}, model.Version{
		Num:        params.Num,
		Content:    params.Content,
		Link:       params.Link,
		Type:       &params.Type,
		Constraint: &params.Constraint,
	})
	if e != nil {
		return err_code.UpdateErr
	} else if i == 0 {
		return err_code.NoData
	}

	return err_code.Success
}
