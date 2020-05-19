package service

import (
	"context"
	"kratos_frame/internal/model"
	err_code "kratos_frame/internal/pkg/ecode"
	"strconv"

	"github.com/bilibili/kratos/pkg/ecode"
	"github.com/jinzhu/gorm"
)

func (s *Service) InsertRole(ctx context.Context, params model.InsertRoleParams) ecode.Code {
	// 参数验证

	userRole, err := s.dao.Insert(ctx, model.Tb+"role", &model.Role{
		Name:    params.Name,
		Content: params.Content,
	})
	if err != nil || userRole.(*model.Role).ID == 0 {
		return err_code.InsertErr
	}

	return err_code.Success
}

func (s *Service) FindRole(ctx context.Context, params model.FindRoleParams) (interface{}, int, ecode.Code) {

	var roles []model.Role
	total := 0
	if params.RoleId != 0 {
		if err := s.dao.Find(ctx, model.Tb+"role", model.Role{
			Model: gorm.Model{ID: params.RoleId},
		}).Find(&roles).Error; err != nil {
			return nil, 0, err_code.FindErr
		}
		if len(roles) == 0 {
			return nil, 0, err_code.NoData
		}

		casBins := s.dao.GetPolicyByName(strconv.Itoa(int(roles[0].ID)))
		var menus []model.Menu
		var superiorId uint
		superiorId = 0
		var mids []string // 所拥有的菜单编号集合
		for _, v := range casBins {
			var menu []model.Menu
			if err := s.dao.Find(ctx, model.Tb+"menu", model.Menu{ // 只查询0级菜单
				Mid:        v[1],
				SuperiorId: &superiorId,
			}).Find(&menu).Error; err != nil {
				return nil, 0, err_code.FindErr
			}
			menus = append(menus, menu...)
			mids = append(mids, v[1])
		}

		m := make([]model.Menu, len(menus))
		for k, v := range menus {
			menu, e := s.FindMenuBySuperiorId(ctx, v.ID, mids)
			if e.Code() != 200 {
				return nil, 0, e
			}
			v.Menus = menu
			m[k] = v
		}
		return struct {
			Role model.Role
			Menu interface{}
		}{
			Role: roles[0],
			Menu: m,
		}, len(menus), err_code.Success

	}

	if params.UserId != 0 {
		var userRoles []model.UserRole
		err := s.dao.Find(ctx, model.Tb+"user_role", model.UserRole{UserId: params.UserId}).Where("deleted_at IS NULL").Count(&total).
			Limit(params.PageSize).Offset((params.Page - 1) * params.PageSize).Find(&userRoles).Error
		if err != nil {
			return nil, 0, err_code.FindErr
		}
		for _, v := range userRoles {
			var role model.Role
			err := s.dao.Find(ctx, model.Tb+"role", model.Role{Model: gorm.Model{ID: v.RoleId}}).First(&role).Error
			if err != nil {
				return nil, 0, err_code.FindErr
			}
			roles = append(roles, role)
		}
		return roles, total, err_code.Success
	}

	err := s.dao.Find(ctx, model.Tb+"role", model.Role{}).Where("deleted_at IS NULL").Count(&total).
		Limit(params.PageSize).Offset((params.Page - 1) * params.PageSize).Find(&roles).Error
	if err != nil {
		return nil, 0, err_code.FindErr
	}
	return roles, total, err_code.Success
}

func (s *Service) DeleteRole(ctx context.Context, params model.DeleteRoleParams) ecode.Code {
	e, i := s.dao.Delete(ctx, model.Tb+"role", model.Role{
		Model: gorm.Model{ID: params.RoleId},
	})
	if e != nil {
		return err_code.DeleteErr
	} else if i == 0 {
		return err_code.NoData
	}
	_, err := s.dao.DelPolicyByName(strconv.Itoa(int(params.RoleId)))
	if err != nil {
		return err_code.DeleteErr
	}

	return err_code.Success
}

func (s *Service) UpdateRole(ctx context.Context, params model.UpdateRoleParams) ecode.Code {
	e, i := s.dao.Update(ctx, model.Tb+"role", model.Role{
		Model: gorm.Model{ID: params.RoleId},
	}, model.Role{
		Name:    params.Name,
		Content: params.Content,
	})
	if e != nil {
		return err_code.UpdateErr
	} else if i == 0 {
		return err_code.NoData
	}

	return err_code.Success
}

// 添加角色和菜单的绑定
func (s *Service) InsertRoleMenu(ctx context.Context, params model.InsertRoleMenuParams) ecode.Code {
	// 参数验证
	var menus []model.Menu
	for _, v := range params.MenuId {
		var menu []model.Menu
		err := s.dao.Find(ctx, model.Tb+"menu", model.Menu{
			Model: gorm.Model{ID: v},
		}).Find(&menu).Error
		if err != nil {
			return err_code.InsertErr
		}
		menus = append(menus, menu...)
	}

	_, err := s.dao.DelPolicyByName(strconv.Itoa(int(params.RoleId)))
	if err != nil {
		return err_code.InsertErr
	}

	for _, v := range menus {
		if v.Method == "" {
			v.Method = "0"
		}
		_, e := s.dao.AddPolicy(strconv.Itoa(int(params.RoleId)), v.Mid, v.Method)
		if e != nil {
			return err_code.InsertErr
		}
	}

	return err_code.Success
}

// 获取用户角色
func (s *Service) FindRoleByUser(ctx context.Context, params model.FindRoleByUserParams) ([]model.UserRole, ecode.Code) {
	var roles []model.UserRole
	err := s.dao.Find(ctx, model.Tb+"user_role", model.UserRole{
		UserId: params.UserId,
	}).Find(&roles).Error
	if err != nil {
		return nil, err_code.FindErr
	}
	return roles, err_code.Success
}

// 判断权限是否有效
func (s *Service) EnforcePolicy(params model.EnforcePolicyParams) (bool, ecode.Code) {
	// 参数验证
	_, err := s.dao.EnforcePolicy(params.Name, params.Path, params.Method)
	if err != nil {
		return false, err_code.VerifyRoleErr
	}
	return true, err_code.Success
}
