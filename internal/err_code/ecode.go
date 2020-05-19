package err_code

import (
	err_code "kratos_frame/internal/pkg/ecode"

	"github.com/bilibili/kratos/pkg/ecode"
)

func NewECode() {
	cms := map[int]string{
		err_code.Success.Code():            "操作成功",
		err_code.NoLogin.Code():            "用戶未登錄",
		err_code.LoginErr.Code():           "登錄失敗",
		err_code.TokenErr.Code():           "Token認證失敗",
		err_code.NoUser.Code():             "用戶不存在",
		err_code.PhoneErr.Code():           "手機號錯誤",
		err_code.PassWordErr.Code():        "密碼錯誤",
		err_code.ArgsErr.Code():            "參數錯誤",
		err_code.RegisterCodeErr.Code():    "邀請碼驗證失敗",
		err_code.UserAlreadyExists.Code():  "用戶已存在",
		err_code.CreatUserErr.Code():       "創建用戶失敗",
		err_code.CreatUserInviteErr.Code(): "創建用戶級別失敗",
		err_code.InsertErr.Code():          "添加失敗",
		err_code.FindErr.Code():            "查看失敗",
		err_code.UpdateErr.Code():          "修改失敗",
		err_code.DeleteErr.Code():          "刪除失敗",
		err_code.NoData.Code():             "數據不存在",
		err_code.CreatTokenErr.Code():      "創建Token失敗",
		err_code.VerifyTokenErr.Code():     "用戶未登錄",
		err_code.VerifyRoleErr.Code():      "權限驗證失敗",
		err_code.SendSmsErr.Code():         "發送短信驗證碼失敗",
		err_code.SendEmailErr.Code():       "發送郵箱驗證碼失敗",
		err_code.SmsCodeTimeOut.Code():     "短信驗證碼已過期",
		err_code.EmailCodeTimeOut.Code():   "郵箱驗證碼已過期",
		err_code.VerifySmsErr.Code():       "短信驗證失敗",
		err_code.VerifyEmailErr.Code():     "郵箱驗證失敗",
		err_code.GoogleVerifyErr.Code():    "谷歌驗證失敗",
		err_code.UserDisabled.Code():       "賬戶禁用",
		err_code.UserLogged.Code():         "其他設備已登錄",
		err_code.LogoutErr.Code():          "退出登錄失敗",
		err_code.SystemMaintenance.Code():  "系統升級,敬請期待",
	}
	ecode.Register(cms)
}
