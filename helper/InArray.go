package helper

import "github.com/gohouse/t"

// InArray :给定元素值 是否在 指定的数组中
func InArray(needle, hystack interface{}) bool {
	nt := t.New(needle)
	for _, item := range t.New(hystack).Slice() {
		if nt.String() == item.String() {
			return true
		}
	}
	return false
}
