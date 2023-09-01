package day230831

import (
	"fmt"
	"testing"
)

// @program:   go-basic-exercises
// @file:      struct_test.go
// @author:    bowen
// @time:      2023-08-31 22:08
// @description: 笔记复习

func TestStruct(t *testing.T) {
	obj := Newfile(10, "xxx")
	fmt.Println(obj)

	m := NewMatrix("bowen")
	fmt.Println(m)
}
