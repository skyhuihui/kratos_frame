package main

import (
	"fmt"
	"os"
	"reflect"
)

// 自动生成此框架需要的service， controller

type ModelTmpl struct {
	ModelName string
	TableName string
}

const (
	server  = "./gen/server"
	service = "./gen/service"
	params  = "./gen/params"
)

func main() {
	Gen(ModelTmpl{
		ModelName: "Test",
		TableName: "test",
	})
}

// path 写入路径
func Gen(tmplData ModelTmpl) {
	if err := GenService(tmplData); err != nil {
		panic(err)
	}
	if err := GenServer(tmplData); err != nil {
		panic(err)
	}
	if err := GenParams(tmplData); err != nil {
		panic(err)
	}
}

func isExist(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}

//dst type interface 有数据的结构体
//src type interace  要修改的结构体的指针
func Copy(dst interface{}, src interface{}) error {
	if reflect.TypeOf(dst).Kind() != reflect.Struct {
		return fmt.Errorf("must be of structure type")
	}
	if reflect.TypeOf(src).Kind() != reflect.Ptr {
		return fmt.Errorf("must be of ptr type")
	}

	vVal := reflect.ValueOf(dst)        //获取reflect.Type类型
	bVal := reflect.ValueOf(src).Elem() //获取reflect.Type类型
	for i := 0; i < vVal.NumField(); i++ {
		// 在要修改的结构体中查询有数据结构体中相同属性的字段，有则修改其值
		name := vVal.Type().Field(i).Name
		if ok := bVal.FieldByName(name).IsValid(); ok {
			bVal.FieldByName(name).Set(reflect.ValueOf(vVal.Field(i).Interface()))
		}
	}
	return nil
}
