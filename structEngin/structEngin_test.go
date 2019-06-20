package structEngin

import (
	"fmt"
	"github.com/gohouse/gocar/varBindValue"
	"testing"
)

type Nested2 struct {
	F2 int
}
type Nested3 struct {
	F3 int
}
type Nested struct {
	F1 interface{}
	Nested2
	Nested3
	F4 string `gorose:"ignore"`
}

func TestStructEngin_GetStructFields(t *testing.T) {
	e := New()
	var n = new(Nested)
	res := e.GetStructFields(n)
	fmt.Println(res)

	// 绑定值
	for _, item := range res {
		err := varBindValue.BindVal(item, 333)
		fmt.Print(err)
	}
	fmt.Println()
	fmt.Println(n)
}

func TestStructEngin_StructContent2Map(t *testing.T) {
	e := New()
	var n = new(Nested)
	n.F1 = 1
	n.F2 = 2
	n.F3 = 3
	n.F4 = "d"
	var n2 = Nested{F1: 11, F4: "b"}
	var m = []Nested{n2, *n}
	res := e.StructContent2Map(m)

	fmt.Println(res)
}
