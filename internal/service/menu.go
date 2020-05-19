package service

import (
	"context"
	"kratos_frame/internal/model"
	err_code "kratos_frame/internal/pkg/ecode"
	"kratos_frame/internal/pkg/rand"

	"github.com/bilibili/kratos/pkg/ecode"
	"github.com/chenhg5/collection"
	"github.com/jinzhu/gorm"
)

func (s *Service) InsertMenu(ctx context.Context, params model.InsertMenuParams) ecode.Code {
	// 参数验证

	userMenu, err := s.dao.Insert(ctx, model.Tb+"menu", &model.Menu{
		Mid:        rand.GetRandomString(6),
		SuperiorId: &params.SuperiorId,
		Name:       params.Name,
		Type:       &params.Type,
		Path:       params.Path,
		Method:     params.Method,
	})
	if err != nil || userMenu.(*model.Menu).ID == 0 {
		return err_code.InsertErr
	}

	return err_code.Success
}

func (s *Service) FindMenu(ctx context.Context, params model.FindMenuParams) (interface{}, int, ecode.Code) {

	var menu model.Menu
	if params.Type != -1 {
		menu.Type = &params.Type
	}

	var menus []model.Menu
	var superiorId uint
	superiorId = 0
	menu.SuperiorId = &superiorId
	total := 0
	err := s.dao.Find(ctx, model.Tb+"menu", menu).Where("deleted_at IS NULL").Count(&total).
		Limit(params.PageSize).Offset((params.Page - 1) * params.PageSize).Find(&menus).Error
	if err != nil {
		return nil, 0, err_code.FindErr
	}
	m := make([]model.Menu, len(menus))
	for k, v := range menus {
		menu, e := s.FindMenuBySuperiorId(ctx, v.ID, []string{})
		if e.Code() != 200 {
			return nil, 0, e
		}
		v.Menus = menu
		m[k] = v
	}

	return m, total, err_code.Success
}

func (s *Service) FindMenuByPath(ctx context.Context, params model.FindMenuByPathParams) ([]model.Menu, ecode.Code) {
	var menus []model.Menu
	err := s.dao.Find(ctx, model.Tb+"menu", model.Menu{Path: params.Path}).Find(&menus).Error
	if err != nil {
		return nil, err_code.FindErr
	}
	return menus, err_code.Success
}

func (s *Service) DeleteMenu(ctx context.Context, params model.DeleteMenuParams) ecode.Code {
	e, i := s.dao.Delete(ctx, model.Tb+"menu", model.Menu{
		Model: gorm.Model{ID: params.MenuId},
	})
	if e != nil {
		return err_code.DeleteErr
	} else if i == 0 {
		return err_code.NoData
	}

	return err_code.Success
}

func (s *Service) UpdateMenu(ctx context.Context, params model.UpdateMenuParams) ecode.Code {
	e, i := s.dao.Update(ctx, model.Tb+"menu", model.Menu{
		Model: gorm.Model{ID: params.MenuId},
	}, model.Menu{
		SuperiorId: &params.SuperiorId,
		Name:       params.Name,
		Type:       &params.Type,
		Path:       params.Path,
		Method:     params.Method,
	})
	if e != nil {
		return err_code.UpdateErr
	} else if i == 0 {
		return err_code.NoData
	}

	return err_code.Success
}

// 根据上级菜单 查询出所有的下级菜单
// 下级拼成 interface返回
// 第三个参数 对比参数： 如果某条记录不在查询的集合中则忽略 不穿此参数则忽略
func (s *Service) FindMenuBySuperiorId(ctx context.Context, superiorId uint, mids []string) (interface{}, ecode.Code) {
	var menus []model.Menu
	err := s.dao.Find(ctx, model.Tb+"menu", model.Menu{SuperiorId: &superiorId}).Find(&menus).Error
	if err != nil {
		return nil, err_code.FindErr
	}

	m := make([]model.Menu, 0)
	// 查询下级菜单的下级菜单
	for _, v := range menus {
		if !collection.Collect(mids).Contains(v.Mid) && len(mids) != 0 { // 此菜单不在需要查询的菜单列表里
			continue
		}
		menu, e := s.FindMenuBySuperiorId(ctx, v.ID, mids)
		if e.Code() != 200 {
			return nil, e
		}
		v.Menus = menu
		m = append(m, v)
	}
	return m, err_code.Success
}
