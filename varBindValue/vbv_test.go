package varBindValue

import (
	"fmt"
	"testing"
)

func TestBindVal(t *testing.T) {
	var a int

	err := BindVal(&a, 234)

	fmt.Println(a)
	fmt.Println(err)
}

func TestBindVal2(t *testing.T) {
	var a int

	err := BindVal(a, 234)

	fmt.Println(a)
	fmt.Println(err)
}
