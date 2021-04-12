package main

import "kratos_frame/internal/pkg/gen/tmpl"

// 自动生成此框架需要的service， controller

// path 写入路径
func Gen(tmplData tmpl.ModelTmpl) {
	if err := tmpl.GenService(tmplData); err != nil {
		panic(err)
	}
	if err := tmpl.GenServer(tmplData); err != nil {
		panic(err)
	}
	if err := tmpl.GenParams(tmplData); err != nil {
		panic(err)
	}
	if err := tmpl.GenMdTmpl(tmplData); err != nil {
		panic(err)
	}
}
