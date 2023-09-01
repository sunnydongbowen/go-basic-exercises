package day230831

import (
	"fmt"
	"testing"
)

// @program:   go-basic-exercises
// @file:      type_test.go
// @author:    bowen
// @time:      2023-08-31 14:24
// @description: 复习笔记

// 注意着两个的区别
type mytype int
type yourtype = int

func TestType(t *testing.T) {
	var m rune
	fmt.Printf("%T\n", m)

	var y mytype
	fmt.Printf("%T\n", y)

	var r yourtype
	fmt.Printf("%T\n", r)
}
