package tmpl

// 生成mvc结构
type ModelTmpl struct {
	ModelName string // 表对应的结构体名称
	TableName string // 表名
}

// 生成 postman markdown的结构
type MdTmpl struct {
	Args  []Arg
	Resps []Resp
}

type Arg struct {
	Name string
	Type string
	Note string
}

type Resp struct {
	Name string
	Type string
	Note string
}
