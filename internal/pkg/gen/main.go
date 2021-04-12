package main

import "kratos_frame/internal/pkg/gen/tmpl"

func main() {
	ms := make([]interface{}, 0)
	ms = append(ms, tmpl.ModelTmpl{ModelName: "Test", TableName: "test"})
	for _, m := range ms {
		Gen(m.(tmpl.ModelTmpl))
	}
}
