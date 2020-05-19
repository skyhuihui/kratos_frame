package service

import (
	"context"
	"kratos_frame/internal/model"
	err_code "kratos_frame/internal/pkg/ecode"

	"github.com/bilibili/kratos/pkg/ecode"
	"github.com/jinzhu/gorm"
)

func (s *Service) InsertNotice(ctx context.Context, params model.InsertNoticeParams) ecode.Code {
	// 参数验证

	userCertification, err := s.dao.Insert(ctx, model.Tb+"notice", &model.Notice{
		Title:   params.Title,
		Content: params.Content,
	})
	if err != nil || userCertification.(*model.Notice).ID == 0 {
		return err_code.InsertErr
	}

	return err_code.Success
}

func (s *Service) FindNotice(ctx context.Context, params model.PageParams) ([]model.Notice, int, ecode.Code) {
	var Notice []model.Notice
	total := 0
	err := s.dao.Find(ctx, model.Tb+"notice", model.Notice{}).Where("deleted_at IS NULL").Count(&total).Order("id desc").
		Limit(params.PageSize).Offset((params.Page - 1) * params.PageSize).Find(&Notice).Error
	if err != nil {
		return nil, 0, err_code.FindErr
	}
	return Notice, total, err_code.Success
}

func (s *Service) DeleteNotice(ctx context.Context, params model.DeleteNoticeParams) ecode.Code {
	e, i := s.dao.Delete(ctx, model.Tb+"notice", model.Notice{
		Model: gorm.Model{ID: params.NoticeId},
	})
	if e != nil {
		return err_code.DeleteErr
	} else if i == 0 {
		return err_code.NoData
	}

	return err_code.Success
}

func (s *Service) UpdateNotice(ctx context.Context, params model.UpdateNoticeParams) ecode.Code {
	e, i := s.dao.Update(ctx, model.Tb+"notice", model.Notice{
		Model: gorm.Model{ID: params.NoticeId},
	}, model.Notice{
		Title:   params.Title,
		Content: params.Content,
	})
	if e != nil {
		return err_code.UpdateErr
	} else if i == 0 {
		return err_code.NoData
	}

	return err_code.Success
}
