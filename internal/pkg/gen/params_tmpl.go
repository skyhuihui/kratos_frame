package main

import (
	"os"
	"strings"
	"text/template"
)

// 自行定义参数
var paramsTmpl = `
package model

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

	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		err := os.Remove(filePath)
		if err != nil {
			return err
		}
	}

	f, err := os.Create(filePath)
	if err != nil {
		return err
	}

	tmpl, err := template.New("").Parse(paramsTmpl)
	if err != nil {
		return err
	}

	err = tmpl.Execute(f, tmplData)
	if err != nil {
		return err
	}

	return nil
}
