package service

import (
	"context"
	"github.com/bilibili/kratos/pkg/log"
	"kratos_frame/internal/model"
	err_code "kratos_frame/internal/pkg/ecode"

	"github.com/bilibili/kratos/pkg/ecode"
	"github.com/jinzhu/gorm"
)

func (s *Service) InsertTest(ctx context.Context, params model.InsertTestParams) ecode.Code {

	var args model.Test
	if err := conversion.Copy(params, &args); err != nil {
		return err_code.ArgsErr
	}

	m, err := s.dao.Insert(ctx, model.Tb+"test", &args)
	if err != nil || m.(*model.Test).ID == 0 {
		return err_code.InsertErr
	}

	return err_code.Success
}

func (s *Service) FindTest(ctx context.Context, params model.FindTestParams) ([]model.Test, int, ecode.Code) {

	var args model.Test
	if err := conversion.Copy(params, &args); err != nil {
		return nil, 0, err_code.ArgsErr
	}

	var Tests []model.Test
	total := 0
	err := s.dao.Find(ctx, model.Tb+"test", args).Where("deleted_at IS NULL").Count(&total).
		Limit(params.PageSize).Offset((params.Page - 1) * params.PageSize).Find(&Tests).Error
	if err != nil {
		log.Error("数据查询失败 :%s", err.Error())
		return nil, 0, err_code.FindErr
	}
	return Tests, total, err_code.Success
}

func (s *Service) DeleteTest(ctx context.Context, params model.DeleteTestParams) ecode.Code {
	e, i := s.dao.Delete(ctx, model.Tb+"test", model.Test{
		Model: gorm.Model{ID: params.TestId},
	})
	if e != nil {
		log.Error("数据删除失败 :%s", e.Error())
		return err_code.DeleteErr
	} else if i == 0 {
		return err_code.NoData
	}

	return err_code.Success
}

func (s *Service) UpdateTest(ctx context.Context, params model.UpdateTestParams) ecode.Code {

	var args model.Test
	if err := conversion.Copy(params, &args); err != nil {
		return err_code.ArgsErr
	}

	e, i := s.dao.Update(ctx, model.Tb+"test", model.Test{
		Model: gorm.Model{ID: params.TestId},
	}, args)
	if e != nil {
		log.Error("数据修改失败 :%s", e.Error())
		return err_code.UpdateErr
	} else if i == 0 {
		return err_code.NoData
	}

	return err_code.Success
}
