package structAutoGenerate

import (
	"fmt"
	"testing"
)

func TestStructAutoGenerate_Generate(t *testing.T) {

	type Users struct {
		Fields        []interface{}
		Result        []map[string]interface{}
		TagName       string
		TagIgnoreName string
	}
	//var a tt
	err := New(&Option{
		Obj:         Users{},
		PackageName: "structAutoGenerate",
		SavePath:    "./users.go",
	}).Generate()

	fmt.Println(err)
}
