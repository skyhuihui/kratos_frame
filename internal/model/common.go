package model

import "github.com/bilibili/kratos/pkg/ecode"

type Resp struct {
	Status  bool        `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func GetResp(code ecode.Code, data interface{}) (r Resp) {
	err := ecode.FromCode(code)
	if err.Code() == 200 {
		r.Status = true
		r.Code = 200
		r.Message = "成功"
		r.Data = data
	}
	r.Status = false
	r.Code = err.Code()
	r.Message = err.Message()
	return
}
