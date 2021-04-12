package tmpl

import (
	"reflect"
)

// 自行定义参数
var postmanMdTmpl = `
**参数：**

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
{{range $i, $v := .Args}}
|{{$v.Name}} |是  |{{$v.Type}} | {{$v.Note}}   |
{{end}}

**返回参数说明：**

|参数名|类型|说明|
|:---- |:----- |-----   |
|code   |int | 错误码 |
| message|string | 错误信息 |
|ttl  |int | 暂时忽略 |
{{range $i, $v := .Resps}}
|{{$v.Name}} |{{$v.Type}} | {{$v.Note}}   |
{{end}}

**返回参数：**
`

func GenPostmanMd(reqModel, respModel interface{}) error {
	isExist(md)
	name := reflect.TypeOf(reqModel).Name()
	filePath := md + "/" + name + ".md"

	arg := make([]Arg, 0)
	t := reflect.TypeOf(reqModel)
	for i := 0; i < t.NumField(); i++ {
		note, _ := t.Field(i).Tag.Lookup("note")
		arg = append(arg, Arg{
			Name: t.Field(i).Tag.Get("json"),
			Type: t.Field(i).Type.String(),
			Note: note,
		})
	}

	resp := make([]Resp, 0)
	if name[:4] == "Find" {
		t = reflect.TypeOf(respModel)
		for i := 0; i < t.NumField(); i++ {
			note, _ := t.Field(i).Tag.Lookup("note")
			resp = append(resp, Resp{
				Name: t.Field(i).Name,
				Type: t.Field(i).Type.String(),
				Note: note,
			})
		}
	}

	mdData := MdTmpl{
		Args:  arg,
		Resps: resp,
	}

	return tmpl(filePath, postmanMdTmpl, mdData)
}
