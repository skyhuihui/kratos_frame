package service

import (
	"context"
	"kratos_frame/internal/model"
	err_code "kratos_frame/internal/pkg/ecode"
	"kratos_frame/internal/pkg/jwt"
	"kratos_frame/internal/pkg/rand"
	"kratos_frame/internal/pkg/sha256"
	"strconv"

	"github.com/bilibili/kratos/pkg/log"

	"github.com/jinzhu/gorm"

	"github.com/bilibili/kratos/pkg/ecode"
)

//退出登录
func (s *Service) Logout(ctx context.Context, params model.LogoutParams) ecode.Code {
	return err_code.Success
}

// 登录
func (s *Service) Login(ctx context.Context, params model.LoginParams) (interface{}, ecode.Code) {

	// 参数验证
	var users []model.User
	if err := s.dao.Find(ctx, model.Tb+"user", model.User{
		Mobile: params.Mobile,
	}).Find(&users).Error; err != nil {
		return nil, err_code.LoginErr
	}
	if len(users) == 0 { // 用户不存在
		return nil, err_code.NoUser
	}
	if users[0].Password != sha256.Sha(params.Password) { // 密码错误
		return nil, err_code.PassWordErr
	}
	if *users[0].Status == 1 { // 账户禁用
		return nil, err_code.UserDisabled
	}

	token, err := jwt.GenerateToken(users[0].ID)
	if err != nil {
		return nil, err_code.CreatTokenErr
	}

	roles, e := s.FindRoleByUser(ctx, model.FindRoleByUserParams{UserId: users[0].ID})
	if e.Code() != 200 {
		return nil, err_code.LoginErr
	}
	var role []uint
	for _, v := range roles {
		role = append(role, v.RoleId)
	}

	// 登录Token存入缓存
	err = s.dao.RedisSet(context.Background(), strconv.Itoa(int(users[0].ID))+"_token", token)
	if err != nil {
		return nil, err_code.LoginErr
	}

	return struct {
		Token string
		User  model.User
		Role  []uint
	}{
		Token: token,
		User:  users[0],
		Role:  role,
	}, err_code.Success
}

// 注册
func (s *Service) Signup(ctx context.Context, params model.SignupParams) ecode.Code {
	superiorUsers, eCode := s.VerifySignUpSuperior(ctx, params.RegisterCode)
	if eCode.Code() != 200 {
		return eCode
	}

	// 手机号验证码验证
	if eCode := s.VerifySms(ctx, params.Mobile, params.MobileCode); eCode.Code() != 200 {
		return eCode
	}

	// 判断用户存不存在
	var users []model.User
	if err := s.dao.Find(ctx, model.Tb+"user", model.User{Mobile: params.Mobile}).Find(&users).Error; err != nil || len(users) != 0 {
		return err_code.UserAlreadyExists
	}

	// 创建用户
	user := model.User{
		Uid:        rand.GetRandomInt(8),
		Password:   sha256.Sha(params.Password),
		FundPass:   sha256.Sha(params.FundPassWord),
		Mobile:     params.Mobile,
		InviteCode: params.Mobile,
		Status:     &model.Zero,
	}
	userI, err := s.dao.Insert(ctx, model.Tb+"user", &user)
	if err != nil || userI == nil || userI.(*model.User).ID == 0 {
		return err_code.CreatUserErr
	}

	// 处理认证情况
	if err := s.InsertCertLevel(ctx, model.UserCertLevel{
		UserId:       userI.(*model.User).ID,
		IsMobile:     &model.One,
		IsReal:       &model.Zero,
		IsEmail:      &model.Zero,
		IsGooGleCode: &model.Zero,
	}); err != nil {
		return err_code.InsertErr
	}

	// 分配用户身份
	role, err := s.dao.Insert(ctx, model.Tb+"user_role", &model.UserRole{
		UserId: userI.(*model.User).ID,
		RoleId: 2,
	})
	if err != nil || role == nil || role.(*model.UserRole).ID == 0 {
		return err_code.CreatUserErr
	}

	// 处理上下级关系
	if len(superiorUsers) != 0 {
		go func(ctx context.Context, userI interface{}, superiorUsers []model.User) {
			if err := s.InsertInvite(ctx, userI.(*model.User).ID, superiorUsers[0].ID); err != nil {
				log.Error("上下级添加失败", err)
			}
		}(ctx, userI, superiorUsers)
	}
	return err_code.Success
}

// 忘记密码
func (s *Service) ForgetPassword(ctx context.Context, params model.ForgetPasswordParams) ecode.Code {
	// 手机号验证码验证
	if eCode := s.VerifySms(ctx, params.Mobile, params.MobileCode); eCode.Code() != 200 {
		return eCode
	}

	e, i := s.dao.Update(ctx, model.Tb+"user", model.User{
		Mobile: params.Mobile,
	}, model.User{
		Mobile:   params.Mobile,
		Password: sha256.Sha(params.Password),
	})
	if e != nil {
		return err_code.PassWordErr
	} else if i == 0 {
		return err_code.NoData
	}

	return err_code.Success
}

// 修改密码
func (s *Service) UpdatePassword(ctx context.Context, params model.UpdatePasswordParams) ecode.Code {
	// 参数验证

	e, i := s.dao.Update(ctx, model.Tb+"user", model.User{
		Mobile:   params.Mobile,
		Password: sha256.Sha(params.OldPassword),
	}, model.User{
		Mobile:   params.Mobile,
		Password: sha256.Sha(params.Password),
	})
	if e != nil {
		return err_code.PassWordErr
	} else if i == 0 {
		return err_code.NoData
	}

	return err_code.Success
}

// 修改资金密码
func (s *Service) UpdateFundPassword(ctx context.Context, params model.UpdateFundPasswordParams) ecode.Code {
	// 参数验证

	e, i := s.dao.Update(ctx, model.Tb+"user", model.User{
		Mobile:   params.Mobile,
		FundPass: sha256.Sha(params.OldPassword),
	}, model.User{
		Mobile:   params.Mobile,
		FundPass: sha256.Sha(params.Password),
	})
	if e != nil {
		return err_code.PassWordErr
	} else if i == 0 {
		return err_code.NoData
	}

	return err_code.Success
}

// 查询所有用户
func (s *Service) FindUser(ctx context.Context, params model.FindUserParams) ([]model.ReturnUserParams, int, ecode.Code) {
	between, and := s.ProcessingTime(params.Between, params.And)
	var users []model.User
	total := 0
	if err := s.dao.Find(ctx, model.Tb+"user", model.User{
		Model:  gorm.Model{ID: params.UserId},
		Mobile: params.Mobile,
	}).Where("deleted_at IS NULL").Where("created_at BETWEEN ? AND ?", between, and).Count(&total).
		Limit(params.PageSize).Offset((params.Page - 1) * params.PageSize).Find(&users).Error; err != nil {
		return nil, 0, err_code.NoUser
	}
	if len(users) == 0 {
		return nil, 0, err_code.NoUser
	}

	var returnUsers []model.ReturnUserParams
	for _, v := range users {
		// 查询用户是否实名， 是否谷歌验证
		var certLevels model.UserCertLevel
		err := s.dao.Find(ctx, model.Tb+"user_cert_level", model.UserCertLevel{UserId: v.ID}).First(&certLevels).Error
		if err != nil {
			return nil, 0, err_code.FindErr
		}
		// 查询用户上级
		invites, err := s.FindInvite(ctx, v.ID, 0, 1)
		if err != nil {
			return nil, 0, err_code.FindErr
		}
		var superiorId uint
		if len(invites) != 0 {
			superiorId = invites[0].SuperiorId
		}
		returnUsers = append(returnUsers, model.ReturnUserParams{
			UserId:     v.ID,
			Uid:        v.Uid,
			Name:       v.Name,
			Mobile:     v.Mobile,
			InviteCode: v.InviteCode,
			Status:     *v.Status,
			IsReal:     *certLevels.IsReal,
			Superior:   superiorId,
			Time:       v.CreatedAt,
		})
	}
	return returnUsers, total, err_code.Success
}

func (s *Service) FindUserById(ctx context.Context, id uint) model.User {
	var users []model.User
	if err := s.dao.Find(ctx, model.Tb+"user", model.User{
		Model: gorm.Model{ID: id},
	}).Find(&users).Error; err != nil {
		return model.User{}
	}
	if len(users) == 0 {
		return model.User{}
	}
	return users[0]
}

func (s *Service) FindUserByMobile(ctx context.Context, mobile string) model.User {
	var users []model.User
	if err := s.dao.Find(ctx, model.Tb+"user", model.User{
		Mobile: mobile,
	}).Find(&users).Error; err != nil {
		return model.User{}
	}
	if len(users) == 0 {
		return model.User{}
	}
	return users[0]
}

// 分配身份
func (s *Service) InsertRoleUser(ctx context.Context, params model.InsertRoleByUserParams) ecode.Code {

	var users []model.User
	if err := s.dao.Find(ctx, model.Tb+"user", model.User{Mobile: params.Mobile}).Find(&users).Error; err != nil || len(users) == 0 {
		return err_code.NoData
	}

	role, err := s.dao.Insert(ctx, model.Tb+"user_role", &model.UserRole{
		UserId: users[0].ID,
		RoleId: params.RoleId,
	})
	if err != nil || role == nil || role.(*model.UserRole).ID == 0 {
		return err_code.InsertErr
	}

	return err_code.Success
}

// 查询用户的身份信息
func (s *Service) FindUserInfo(ctx context.Context, params model.FindUserParams) (model.ReturnUserInfoParams, ecode.Code) {
	var info model.ReturnUserInfoParams
	var certification model.UserCertification
	err := s.dao.Find(ctx, model.Tb+"user_certification", model.UserCertification{UserId: params.UserId}).First(&certification).Error
	if err != nil {
		return info, err_code.FindErr
	}

	info.UserId = params.UserId
	info.Mobile = params.Mobile
	info.Name = certification.RealName
	info.Cert = certification.Cert
	return info, err_code.Success
}
