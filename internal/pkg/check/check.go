package check

import "regexp"

// 验证资金密码
func CheckFundPwd(password string) bool {
	pattern := `^[0-9]*$`
	reg, _ := regexp.Compile(pattern)
	return reg.MatchString(password)
}
