package service

import (
	"context"
	"kratos_frame/internal/model"
	err_code "kratos_frame/internal/pkg/ecode"

	"github.com/bilibili/kratos/pkg/ecode"
	"github.com/jinzhu/gorm"
)

func (s *Service) InsertProtocol(ctx context.Context, params model.InsertProtocolParams) ecode.Code {
	// 参数验证
	userCertification, err := s.dao.Insert(ctx, model.Tb+"protocol", &model.Protocol{
		Title:   params.Title,
		Content: params.Content,
		Type:    &params.Type,
	})
	if userCertification.(*model.Protocol).ID == 0 || err != nil {
		return err_code.InsertErr
	}

	return err_code.Success
}

func (s *Service) FindProtocol(ctx context.Context, params model.FindProtocolParams) ([]model.Protocol, int, ecode.Code) {

	var protocol model.Protocol
	if params.Type != -1 {
		protocol.Type = &params.Type
	}

	var protocols []model.Protocol
	total := 0
	if err := s.dao.Find(ctx, model.Tb+"protocol", protocol).Where("deleted_at IS NULL").Count(&total).Order("id desc").
		Limit(params.PageSize).Offset((params.Page - 1) * params.PageSize).Find(&protocols).Error; err != nil {
		return nil, 0, err_code.FindErr
	}
	return protocols, total, err_code.Success
}

func (s *Service) DeleteProtocol(ctx context.Context, params model.DeleteProtocolParams) ecode.Code {
	e, i := s.dao.Delete(ctx, model.Tb+"protocol", model.Protocol{
		Model: gorm.Model{ID: params.ProtocolId},
	})
	if e != nil {
		return err_code.DeleteErr
	} else if i == 0 {
		return err_code.NoData
	}

	return err_code.Success
}

func (s *Service) UpdateProtocol(ctx context.Context, params model.UpdateProtocolParams) ecode.Code {
	e, i := s.dao.Update(ctx, model.Tb+"protocol", model.Protocol{
		Model: gorm.Model{ID: params.ProtocolId},
	}, model.Protocol{
		Title:   params.Title,
		Content: params.Content,
		Type:    &params.Type,
	})
	if e != nil {
		return err_code.UpdateErr
	} else if i == 0 {
		return err_code.NoData
	}

	return err_code.Success
}
