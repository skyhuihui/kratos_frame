package service

import (
	"context"
	"kratos_frame/internal/model"
	err_code "kratos_frame/internal/pkg/ecode"

	"github.com/bilibili/kratos/pkg/ecode"
	"github.com/jinzhu/gorm"
)

func (s *Service) InsertWork(ctx context.Context, params model.InsertWorkParams) ecode.Code {
	// 参数验证
	var img string

	for _, v := range params.Img {
		img += v + ","
	}
	userCertification, err := s.dao.Insert(ctx, model.Tb+"work", &model.Work{
		UserId:  params.UserId,
		Title:   params.Title,
		Content: params.Content,
		Type:    &model.Zero,
		Mobile:  params.Mobile,
		Img:     img[0 : len(img)-1],
	})
	if err != nil || userCertification.(*model.Work).ID == 0 {
		return err_code.InsertErr
	}

	return err_code.Success
}

func (s *Service) FindWork(ctx context.Context, params model.FindWorkParams) ([]model.Work, int, ecode.Code) {
	var work model.Work
	if params.Type != -1 {
		work.Type = &params.Type
	}
	if params.UserId != 0 {
		work.UserId = params.UserId
	}
	if params.Mobile != "" {
		work.Mobile = params.Mobile
	}
	var Work []model.Work
	total := 0
	if err := s.dao.Find(ctx, model.Tb+"work", work).Where("deleted_at IS NULL").Count(&total).Order("id desc").
		Limit(params.PageSize).Offset((params.Page - 1) * params.PageSize).Find(&Work).Error; err != nil {
		return nil, 0, err_code.FindErr
	}
	return Work, total, err_code.Success
}

func (s *Service) DeleteWork(ctx context.Context, params model.DeleteWorkParams) ecode.Code {
	e, i := s.dao.Delete(ctx, model.Tb+"work", model.Work{
		Model: gorm.Model{ID: params.WorkId},
	})
	if e != nil {
		return err_code.DeleteErr
	} else if i == 0 {
		return err_code.NoData
	}

	return err_code.Success
}

func (s *Service) UpdateWork(ctx context.Context, params model.UpdateWorkParams) ecode.Code {
	var img string
	for _, v := range params.Img {
		img += v + ","
	}
	e, i := s.dao.Update(ctx, model.Tb+"work", model.Work{
		Model: gorm.Model{ID: params.WorkId},
	}, model.Work{
		Title:   params.Title,
		Content: params.Content,
		Type:    &params.Type,
		Mobile:  params.Mobile,
		Img:     img[0 : len(img)-1],
	})
	if e != nil {
		return err_code.UpdateErr
	} else if i == 0 {
		return err_code.NoData
	}

	return err_code.Success
}
