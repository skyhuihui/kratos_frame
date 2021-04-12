package tmpl

import (
	"strings"
)

// 注意新增，查询，删除，修改 均会吧穿参通过反射赋值给相应的处理结构体，请保证参数同名
// 修改，删除请保证id存在

var serviceTmpl = `
package service

import (
	"context"
	"github.com/bilibili/kratos/pkg/log"
	"kratos_frame/internal/model"
	err_code "kratos_frame/internal/pkg/ecode"

	"github.com/bilibili/kratos/pkg/ecode"
	"github.com/jinzhu/gorm"
)

func (s *Service) Insert{{.ModelName}}(ctx context.Context, params model.Insert{{.ModelName}}Params) ecode.Code {

	var args model.{{.ModelName}}
	if err := conversion.Copy(params, &args); err != nil{
		return err_code.ArgsErr
	}

	m, err := s.dao.Insert(ctx, model.Tb+"{{.TableName}}", &args)
	if err != nil || m.(*model.{{.ModelName}}).ID == 0 {
		return err_code.InsertErr
	}

	return err_code.Success
}

func (s *Service) Find{{.ModelName}}(ctx context.Context, params model.Find{{.ModelName}}Params) ([]model.{{.ModelName}}, int, ecode.Code) {

	var args model.{{.ModelName}}
	if err := conversion.Copy(params, &args); err != nil{
		return nil, 0, err_code.ArgsErr
	}

	var {{.ModelName}}s []model.{{.ModelName}}
	total := 0
	err := s.dao.Find(ctx, model.Tb+"{{.TableName}}", args).Where("deleted_at IS NULL").Count(&total).
		Limit(params.PageSize).Offset((params.Page - 1) * params.PageSize).Find(&{{.ModelName}}s).Error
	if err != nil {
		log.Error("数据查询失败 :%s", err.Error())
		return nil, 0, err_code.FindErr
	}
	return {{.ModelName}}s, total, err_code.Success
}

func (s *Service) Delete{{.ModelName}}(ctx context.Context, params model.Delete{{.ModelName}}Params) ecode.Code {
	e, i := s.dao.Delete(ctx, model.Tb+"{{.TableName}}", model.{{.ModelName}}{
		Model: gorm.Model{ID: params.{{.ModelName}}Id},
	})
	if e != nil {
		log.Error("数据删除失败 :%s", e.Error())
		return err_code.DeleteErr
	} else if i == 0 {
		return err_code.NoData
	}

	return err_code.Success
}

func (s *Service) Update{{.ModelName}}(ctx context.Context, params model.Update{{.ModelName}}Params) ecode.Code {

	var args model.{{.ModelName}}
	if err := conversion.Copy(params, &args); err != nil{
		return err_code.ArgsErr
	}	

	e, i := s.dao.Update(ctx, model.Tb+"{{.TableName}}", model.{{.ModelName}}{
		Model: gorm.Model{ID: params.{{.ModelName}}Id},
	}, args)
	if e != nil {
		log.Error("数据修改失败 :%s", e.Error())
		return err_code.UpdateErr
	} else if i == 0 {
		return err_code.NoData
	}

	return err_code.Success
}
`

func GenService(tmplData ModelTmpl) error {
	isExist(service)
	filePath := service + "/" + strings.ToLower(tmplData.ModelName) + ".go"

	return tmpl(filePath, serverTmpl, tmplData)
}
