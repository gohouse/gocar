package structAutoGenerate

import (
	"fmt"
	"testing"
)

func TestStructAutoGenerate_Generate(t *testing.T) {
	type tt struct {
		A string 
		B int `a:b`
	}
	var a tt
	err := New(&Option{
		Obj:a,
		PackageName:"structAutoGenerate",
	}).Generate()

	fmt.Println(err)
}
