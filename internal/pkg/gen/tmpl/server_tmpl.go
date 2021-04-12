package tmpl

import (
	"strings"
)

var serverTmpl = `
package http

import (
	"kratos_frame/internal/model"
	err_code "kratos_frame/internal/pkg/ecode"

	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

func Insert{{.ModelName}}(c *bm.Context) {
	var params model.Insert{{.ModelName}}Params
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	c.JSON(nil, svc.Insert{{.ModelName}}(c.Context, params))
	return
}

func Find{{.ModelName}}(c *bm.Context) {
	var params model.Find{{.ModelName}}Params
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	data, total, err := svc.Find{{.ModelName}}(c.Context, params)
	RespJson(c, data, total, err)
	return
}

func Delete{{.ModelName}}(c *bm.Context) {
	var params model.Delete{{.ModelName}}Params
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	c.JSON(nil, svc.Delete{{.ModelName}}(c.Context, params))
	return
}
func Update{{.ModelName}}(c *bm.Context) {
	var params model.Update{{.ModelName}}Params
	if err := bindArgs(c, &params); err != nil {
		c.JSON(err.Error(), err_code.ArgsErr)
		return
	}
	c.JSON(nil, svc.Update{{.ModelName}}(c.Context, params))
	return
}
`

func GenServer(tmplData ModelTmpl) error {
	isExist(server)
	filePath := server + "/" + strings.ToLower(tmplData.ModelName) + ".go"

	return tmpl(filePath, serviceTmpl, tmplData)
}
