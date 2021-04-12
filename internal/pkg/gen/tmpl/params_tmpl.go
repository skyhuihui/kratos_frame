package tmpl

import (
	"strings"
)

// 自行定义参数
var paramsTmpl = `
package main

// 使用json tag，做请求参数名， note tag 做参数说明
type Insert{{.ModelName}}Params struct {
}

type Find{{.ModelName}}Params struct {
	PageParams
}

type Delete{{.ModelName}}Params struct {
	{{.ModelName}}Id uint 
}

type Update{{.ModelName}}Params struct {
	{{.ModelName}}Id uint 
}
`

func GenParams(tmplData ModelTmpl) error {
	isExist(params)
	filePath := params + "/" + strings.ToLower(tmplData.ModelName) + ".go"

	return tmpl(filePath, paramsTmpl, tmplData)
}

var genParamsMdTmpl = `
package main

import (
	"kratos_frame/internal/model"
	"kratos_frame/internal/pkg/gen/tmpl"
)

func main() {
	ms := make([]interface{},0)
	ms = append(ms, Insert{{.ModelName}}Params{})
	ms = append(ms, Find{{.ModelName}}Params{})
	ms = append(ms, Delete{{.ModelName}}Params{})
	ms = append(ms, Update{{.ModelName}}Params{})
	for _, m := range ms {
		if err := tmpl.GenPostmanMd(m, model.{{.ModelName}}{}); err != nil {
			panic(err)
		}
	}
}
`

func GenMdTmpl(tmplData ModelTmpl) error {
	isExist(params)
	filePath := params + "/" + strings.ToLower(tmplData.ModelName+"_gen_md") + ".go"

	return tmpl(filePath, genParamsMdTmpl, tmplData)
}
