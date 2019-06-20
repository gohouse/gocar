package structAutoGenerate

import (
	"fmt"
	"testing"
)

func TestStructAutoGenerate_Generate(t *testing.T) {

	type Users struct {
		Uid  int    `gorose:"uid"`
		Name string `gorose:"name"`
		Age  int    `gorose:"age"`
	}
	//var a tt
	err := New(&Option{
		Obj:         Users{},
		PackageName: "structAutoGenerate",
		SavePath: "./users.go",
	}).Generate()

	fmt.Println(err)
}
