package ecode

import (
	"testing"

	"github.com/siddontang/go/log"

	"github.com/bilibili/kratos/pkg/ecode"
)

type response struct {
	Status  bool
	Code    int
	Message string
	Data    interface{}
}

func resp(code ecode.Code) (r response) {
	err := ecode.FromCode(code)
	if err.Code() != 0 {
		r.Status = false
	}
	r.Code = err.Code()
	r.Message = err.Message()
	return
}

func TestECode(t *testing.T) {
	var (
		UserNotLogin     = ecode.New(123)
		UserPassWordErr  = ecode.New(1234)
		UserPassWordErr1 = ecode.New(1235)
		UserPassWordErr2 = ecode.New(1236)
	)
	cms := map[int]string{
		UserNotLogin.Code(): "很好很强大！",
		-304:                "啥都没变啊~",
		-404:                "啥都没有啊~",
	}
	ecode.Register(cms)
	log.Info(ecode.Cause(UserNotLogin).Code())
	log.Info(ecode.Cause(UserNotLogin).Message())
	log.Info(ecode.FromCode(123).Message())
	log.Info(ecode.FromCode(UserPassWordErr).Code())
	log.Info(ecode.FromCode(UserPassWordErr1).Details())
	log.Info(ecode.FromCode(UserPassWordErr2).Message())

	log.Info(resp(123))
	log.Info(resp(-304))
}
