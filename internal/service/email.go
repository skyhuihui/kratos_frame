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

func (s *Service) SendEmail(ctx context.Context, params model.SendEmailParams) ecode.Code {
	// 参数验证

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	//if err := email.SendEmail(1, []string{params.Email}, code); err != nil {
	//	return err_code.SendSmsErr
	//}
	e := s.dao.RedisSet(context.Background(), params.Email+"_email", code, "EX", 120)
	if e != nil {
		return err_code.SendEmailErr
	}
	log.Info("发送验证码：邮箱 ： %s,  :%s ", params.Email, code)
	return err_code.Success
}
