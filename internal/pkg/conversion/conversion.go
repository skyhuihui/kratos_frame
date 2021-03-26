package conversion

import (
	"fmt"
	"reflect"
)

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
