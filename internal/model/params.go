package model

import (
	"time"
)

var (
	Zero = 0
	One  = 1
	Two  = 2
	Tb   = "tb_"
)

// 分页公共参数
type PageParams struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

// 注册请求参数模型
type SignupParams struct {
	Password     string `json:"password" validate:"required,min=6,max=18"`
	FundPassWord string `json:"fund_password" validate:"required,min=6,max=6"`
	Mobile       string `json:"mobile" validate:"required"`
	MobileCode   string `json:"mobile_code" validate:"required"`
	RegisterCode string `json:"register_code"` // 上级邀请码
}

// 登录请求数据模型
type LoginParams struct {
	Mobile   string `json:"mobile" validate:"required"`
	Password string `json:"password" validate:"required,min=6,max=18"`
}

// 退出请求数据模型
type LogoutParams struct {
	Mobile string `json:"mobile" validate:"required"`
}

// 忘记密码
type ForgetPasswordParams struct {
	Password   string `json:"password" validate:"required,min=6,max=18"`
	MobileCode string `json:"mobile_code" validate:"required"`
	Mobile     string `json:"mobile" validate:"required"`
}

// 修改密码
type UpdatePasswordParams struct {
	Password    string `json:"password" validate:"required,min=6,max=18"`
	OldPassword string `json:"old_password" validate:"required,min=6,max=18"`
	Mobile      string `json:"mobile" validate:"required"`
}

// 修改密码
type UpdateFundPasswordParams struct {
	Password    string `json:"password" validate:"required,min=6,max=6"`
	OldPassword string `json:"old_password" validate:"required,min=6,max=6"`
	Mobile      string `json:"mobile" validate:"required"`
}

// 修改用户等级
type UpdateGradeParams struct {
	UserId uint `json:"user_id" validate:"required"`
	Grade  int  `json:"grade" validate:"required"`
}

// 查看认证等级
type FindCertLevelParams struct {
	PageParams
	UserId uint `json:"user_id"`
}

// 添加实名认证信息
type InsertCertificationParams struct {
	UserId   uint     `json:"user_id" validate:"required"`
	RealName string   `json:"real_name" validate:"required"` // 实名姓名
	Cert     string   `json:"cert" validate:"required"`      //身份证号
	Address  string   `json:"address" validate:"required"`   //身份证地址
	Exp      string   `json:"exp" validate:"required"`       //有效期
	Issue    string   `json:"issue" validate:"required"`     //签发机关
	Identity []string `json:"identity" validate:"required"`  //身份证图片
}

// 查看实名认证
type FindCertificationParams struct {
	PageParams
	UserId uint `json:"user_id"`
}

// 修改实名认证信息
type UpdateCertificationParams struct {
	CertificationId uint     `json:"certification_id" validate:"required"`
	RealName        string   `json:"real_name" validate:"required"` // 实名姓名
	Cert            string   `json:"cert" validate:"required"`      //身份证号
	Address         string   `json:"address" validate:"required"`   //身份证地址
	Exp             string   `json:"exp" validate:"required"`       //有效期
	Issue           string   `json:"issue" validate:"required"`     //签发机关
	Identity        []string `json:"identity" validate:"required"`  //身份证图片
}

// 删除实名认证信息
type DeleteCertificationParams struct {
	CertificationId uint `json:"certification_id" validate:"required"`
}

// 添加轮播图
type InsertBannerParams struct {
	Img string `json:"img" validate:"required"`
}

// 删除轮播图
type DeleteBannerParams struct {
	BannerId uint `json:"banner_id" validate:"required"`
}

// 添加公告
type InsertNoticeParams struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content"`
}

// 删除公告
type DeleteNoticeParams struct {
	NoticeId uint `json:"notice_id" validate:"required"`
}

// 修改公告
type UpdateNoticeParams struct {
	NoticeId uint   `json:"notice_id" validate:"required"`
	Title    string `json:"title" validate:"required"`
	Content  string `json:"content"`
}

// 添加版本
type InsertVersionParams struct {
	Num        string `json:"num"`
	Content    string `json:"content"`
	Link       string `json:"link" validate:"required"`
	Type       int    `json:"type"`
	Constraint int    `json:"constraint"`
}

// 删除版本
type DeleteVersionParams struct {
	VersionId uint `json:"version_id" validate:"required"`
}

// 修改版本
type UpdateVersionParams struct {
	VersionId  uint   `json:"version_id" validate:"required"`
	Num        string `json:"num"`
	Content    string `json:"content"`
	Link       string `json:"link" validate:"required"`
	Type       int    `json:"type"`
	Constraint int    `json:"constraint"`
}

// 查看版本
type FindVersionParams struct {
	PageParams
	Type int `json:"type"`
}

// 添加用户协议
type InsertProtocolParams struct {
	Content string `json:"content" validate:"required"`
	Title   string `json:"title" validate:"required"`
	Type    int    `json:"type" validate:"required"`
}

// 删除协议
type DeleteProtocolParams struct {
	ProtocolId uint `json:"protocol_id" validate:"required"`
}

// 修改协议
type UpdateProtocolParams struct {
	ProtocolId uint   `json:"protocol_id" validate:"required"`
	Content    string `json:"content" validate:"required"`
	Type       int    `json:"type" validate:"required"`
	Title      string `json:"title" validate:"required"`
}

// 查看协议
type FindProtocolParams struct {
	PageParams
	Type int `json:"type"`
}

// 添加工单
type InsertWorkParams struct {
	UserId  uint     `json:"user_id" validate:"required"`
	Mobile  string   `json:"mobile" validate:"required"`
	Content string   `json:"content"`
	Title   string   `json:"title" validate:"required"`
	Img     []string `json:"img"`
}

// 删除工单
type DeleteWorkParams struct {
	WorkId uint `json:"work_id" validate:"required"`
}

// 修改工单
type UpdateWorkParams struct {
	WorkId  uint     `json:"work_id" validate:"required"`
	Mobile  string   `json:"mobile" validate:"required"`
	Content string   `json:"content"`
	Title   string   `json:"title" validate:"required"`
	Type    int      `json:"type" validate:"required"`
	Img     []string `json:"img"`
}

// 查看工单
type FindWorkParams struct {
	PageParams
	Type   int    `json:"type"`
	UserId uint   `json:"user_id"`
	Mobile string `json:"mobile"`
}

// 添加菜单
type InsertMenuParams struct {
	SuperiorId uint   `json:"superior_id"`
	Name       string `json:"name" validate:"required"`
	Type       int    `json:"type"`
	Path       string `json:"path"`
	Method     string `json:"method"`
}

// 删除菜单
type DeleteMenuParams struct {
	MenuId uint `json:"menu_id" validate:"required"`
}

// 修改菜单
type UpdateMenuParams struct {
	MenuId     uint   `json:"menu_id" validate:"required"`
	SuperiorId uint   `json:"superior_id"`
	Name       string `json:"name" validate:"required"`
	Type       int    `json:"type"`
	Path       string `json:"path"`
	Method     string `json:"method"`
}

// 查看菜单
type FindMenuParams struct {
	PageParams
	Type int `json:"type"`
}

// 通过 Path 查看菜单
type FindMenuByPathParams struct {
	Path string `json:"path"`
}

// 添加角色
type InsertRoleParams struct {
	Name    string `json:"name" validate:"required"`
	Content string `json:"content"`
}

// 删除角色
type DeleteRoleParams struct {
	RoleId uint `json:"role_id" validate:"required"`
}

// 修改角色
type UpdateRoleParams struct {
	RoleId  uint   `json:"role_id" validate:"required"`
	Name    string `json:"name" validate:"required"`
	Content string `json:"content"`
}

// 查看角色
type FindRoleParams struct {
	PageParams
	RoleId uint `json:"role_id"`
	UserId uint `json:"user_id"`
}

// 根据用户获得角色
type FindRoleByUserParams struct {
	UserId uint `json:"user_id" validate:"required"`
}

// 给用户分配身份
type InsertRoleByUserParams struct {
	Mobile string `json:"mobile" validate:"required"`
	RoleId uint   `json:"role_id" validate:"required"`
}

// 添加角色和菜单的绑定
type InsertRoleMenuParams struct {
	RoleId uint   `json:"role_id" validate:"required"`
	MenuId []uint `json:"menu_id" validate:"required"`
}

// 判断权限是否有效
type EnforcePolicyParams struct {
	Name   string `json:"name"`
	Path   string `json:"path"`
	Method string `json:"method"`
}

// 删除图片
type DelFileParams struct {
	Path string `json:"path" validate:"required"`
}

// 发送短信验证码
type SendSmsParams struct {
	Phone string `json:"phone" validate:"required"`
}

// 发送邮箱验证码
type SendEmailParams struct {
	Email string `json:"email" validate:"required,email"`
}

// 查询谷歌验证
type FindGoogleVerifyParams struct {
	UserId uint `json:"user_id" validate:"required"`
}

// 验证谷歌验证 (需要邮箱验证)
type VerifyGoogleVerifyParams struct {
	UserId     uint   `json:"user_id" validate:"required"`
	GoogleCode string `json:"google_code" validate:"required"`
	EmailCode  string `json:"email_code"`
	PhoneCode  string `json:"phone_code"`
}

// 验证谷歌验证 (需要邮箱验证)
type DeleteUserGoogleVerifyParams struct {
	UserId uint `json:"user_id" validate:"required"`
}

// 查看用户一级下级
type FindUserInviteParams struct {
	PageParams
	UserId uint `json:"user_id" validate:"required"`
}

// 级别查询返回参数
type ReturnUserInviteSubordinateParams struct {
	UserId uint
	Mobile string
	Time   time.Time
}

// 查询邀请相关信息
type ReturnUserInviteParams struct {
	Url            string
	InvitationCode string
	Count          int
}

// 查询用户
type FindUserParams struct {
	PageParams
	UserId  uint   `json:"user_id"`
	Mobile  string `json:"mobile"`
	Between int64  `json:"between"`
	And     int64  `json:"and"`
}

// 用户返回参数
type ReturnUserParams struct {
	UserId     uint
	Uid        string
	Name       string
	Mobile     string
	InviteCode string    // 邀请码
	Status     int       // 是否禁用 0 启用 1 禁用
	IsReal     int       // 是否实名 0待审核 1 已实名 2未实名
	Superior   uint      // 上级用户id
	Time       time.Time // 注册时间
}

// 用户个人中心参数
type ReturnUserInfoParams struct {
	UserId   uint
	Name     string
	Mobile   string
	Cert     string
	Bank     string
	BankCard string
}
