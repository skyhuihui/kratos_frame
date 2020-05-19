package service

import (
	"context"
	"fmt"
	"kratos_frame/internal/model"
	err_code "kratos_frame/internal/pkg/ecode"
	"math/rand"
	"time"

	"github.com/bilibili/kratos/pkg/log"

	"github.com/bilibili/kratos/pkg/ecode"
)

func (s *Service) SendSms(ctx context.Context, params model.SendSmsParams) ecode.Code {
	// 参数验证

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	//if err := sms.SendSms(params.Phone, 1, code); err != nil {
	//	return err_code.SendSmsErr
	//}
	e := s.dao.RedisSet(context.Background(), params.Phone+"_sms", code, "EX", 120)
	if e != nil {
		return err_code.SendSmsErr
	}
	log.Info("发送验证码：手机号  : %s , 验证码 :%s", params.Phone, code)
	return err_code.Success
}
