package service

import (
	"context"
	"fmt"
	"kratos_frame/internal/model"
	err_code "kratos_frame/internal/pkg/ecode"
)

//绑定上下级关系
func (s *Service) InsertInvite(ctx context.Context, userId, superiorId uint) error {

	invite := model.UserInvite{
		UserId:     userId,
		SuperiorId: superiorId,
		Level:      model.One,
	}
	inviteI, err := s.dao.Insert(ctx, model.Tb+"user_invite", &invite)
	if err != nil || inviteI == nil || inviteI.(*model.UserInvite).ID == 0 {
		return fmt.Errorf("添加上级下级直属关系失败")
	}

	invites, err := s.FindInvite(ctx, superiorId, 0, 0)
	if err != nil {
		return err
	}
	for _, v := range invites {
		level := v.Level + 1
		inviteOne := model.UserInvite{
			UserId:     userId,
			SuperiorId: v.SuperiorId,
			Level:      level,
		}
		inviteI, err := s.dao.Insert(ctx, model.Tb+"user_invite", &inviteOne)
		if err != nil || inviteI == nil || inviteI.(*model.UserInvite).ID == 0 {
			return fmt.Errorf("添加上级下级直属关系失败")
		}
	}
	return nil
}

// 1. 只传userId， superiorId是 0 , level 是 0  查询userId 的所有上级
// 2. userId, superiorId都不是零，level是 0 ，判断两个是不是上下级关系
// 3.superiorId 不是 0 , level 是0 查询上级是它的，不论代数， level不是0 ， 指定代数查询
func (s *Service) FindInvite(ctx context.Context, userId, superiorId uint, level int) ([]model.UserInvite, error) {
	var invites []model.UserInvite
	err := s.dao.Find(ctx, model.Tb+"user_invite", model.UserInvite{
		UserId:     userId,
		SuperiorId: superiorId,
		Level:      level,
	}).Find(&invites).Error
	return invites, err
}

// 查询用户邀请相关信息
func (s *Service) FindUserInvite(ctx context.Context, params model.FindUserInviteParams) (model.ReturnUserInviteParams, int, error) {

	invites, err := s.FindInvite(ctx, 0, params.UserId, 1)
	if err != nil {
		return model.ReturnUserInviteParams{}, 0, err_code.FindErr
	}
	user := s.FindUserById(ctx, params.UserId)

	return model.ReturnUserInviteParams{
		Url:            "?Invitationcode=" + user.InviteCode,
		InvitationCode: user.InviteCode,
		Count:          len(invites),
	}, 0, err_code.Success
}

// 查询用户一级下级
func (s *Service) FindUserInviteSubordinate(ctx context.Context, params model.FindUserInviteParams) ([]model.ReturnUserInviteSubordinateParams, int, error) {
	var invites []model.UserInvite
	total := 0

	err := s.dao.Find(ctx, model.Tb+"user_invite", model.UserInvite{SuperiorId: params.UserId, Level: model.One}).
		Where("deleted_at IS NULL").Count(&total).
		Limit(params.PageSize).Offset((params.Page - 1) * params.PageSize).Find(&invites).Error
	if err != nil {
		return nil, 0, err_code.FindErr
	}
	var returnInvites []model.ReturnUserInviteSubordinateParams
	for _, v := range invites {
		user := s.FindUserById(ctx, v.UserId)
		returnInvites = append(returnInvites, model.ReturnUserInviteSubordinateParams{
			UserId: v.UserId,
			Mobile: user.Mobile,
			Time:   user.CreatedAt,
		})
	}

	return returnInvites, total, err_code.Success
}
