# gocar
some useful tool and convertor and etc.

## 目录
[structAutoGenerate](#structAutoGenerate)
[structEngin](#structEngin)
[varBindValue](#varBindValue)

## structAutoGenerate

根据struct, 自动生成对应字段的`Set()`和`Get()`方法, 同时会为这些方法生成对应的接口

```go
package main

import (
	"github.com/gohouse/gocar/structAutoGenerate"
	"fmt"
)

type Users struct {
    Uid  int    `gorose:"uid"`
    Name string `gorose:"name"`
    Age  int    `gorose:"age"`
}
func main() {
    //var a tt
    err := structAutoGenerate.New(&structAutoGenerate.Option{
        // 要生成的机构体对象
        Obj:         Users{},
        // 生成文件包名
        PackageName: "user",
        // 指定生成文件保存目录
        SavePath: "./users.go",
    }).Generate()
    
    fmt.Println(err)
}
```
生成
```go
package user

type IUsers interface {
	SetUid(arg int)
	GetUid() int
	SetName(arg string)
	GetName() string
	SetAge(arg int)
	GetAge() int
}
type Users struct {
	Uid  int    `gorose:"uid"`
	Name string `gorose:"name"`
	Age  int    `gorose:"age"`
}

func NewUsers() *Users {
	return new(Users)
}

func (o *Users) SetUid(arg int) {
	o.Uid = arg
}

func (o *Users) GetUid() int {
	return o.Uid
}

func (o *Users) SetName(arg string) {
	o.Name = arg
}

func (o *Users) GetName() string {
	return o.Name
}

func (o *Users) SetAge(arg int) {
	o.Age = arg
}

func (o *Users) GetAge() int {
	return o.Age
}
```

## structEngin

结构体操作工具包, 目前提供两个工具  
1. 抽取结构体的字段指针, 放入一个 `[]interface{}`中, 同时支持嵌套结构体的循环抽取, 至于抽取出来做什么, 比如可以给他赋值, 在数据库结果集绑定的时候, 很有用  
2. 获取结构体的内容到一个`[]map[string]inteface{}`中, 使用场景: 数据库插入或修改时, 传入结构体数据的解析  
3. 示例  
```go
package structEngin

import (
	"fmt"
	"github.com/gohouse/gocar/varBindValue"
	"github.com/gohouse/gocar/structEngin"
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
	F4 string
}

func main()  {
    testGetStructFields()
    testStructContent2Map()
}

func testGetStructFields() {
	e := structEngin.New()
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

func testStructContent2Map() {
	e := structEngin.New()
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
```
结果
```bash
[0xc00009e2a0 0xc00009e2b0 0xc00009e2b8 0xc00009e2c0]
<nil><nil><nil><nil>
&{333 {333} {333} 333}
--- PASS: TestStructEngin_GetStructFields (0.00s)
=== RUN   TestStructEngin_StructContent2Map
[map[F1:11 F4:b] map[F1:1 F4:d]]
```

## varBindValue

给变量指针绑定值
```go

package main

import (
	"github.com/gohouse/gocar/varBindValue"
	"fmt"
)
func main() {
	var a int

	err := varBindValue.BindVal(&a, 234)

	fmt.Println(a)
	fmt.Println(err)
}
```
结果
```bash
234
<nil>
```
