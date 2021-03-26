package conversion

import (
	"fmt"
	"testing"
)

type test1 struct {
	Name string
	Age  int
	Sex  string
}

type test2 struct {
	Name string
	Age  int
	Sex  string
}

func TestCopy(t *testing.T) {
	t2 := test2{}
	t1 := test1{
		Name: "xiaoming",
		Age:  0,
		Sex:  "",
	}
	_ = Copy(t1, &t2)
	fmt.Println(t2)
}
