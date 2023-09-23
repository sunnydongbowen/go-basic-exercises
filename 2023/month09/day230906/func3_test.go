package main

import (
	"fmt"
	"testing"
)

// @program:   go-basic-exercises
// @file:      func3_test.go
// @author:    bowen
// @time:      2023-09-06 15:38
// @description:

func Func14() string {
	fmt.Println("方法可以作为变量名")
	return "hello"
}

var Abc = func() string {
	return "hello"
}

func Func15(n int) int {
	return n
}

func function16() {
	fn := func() string {
		return ""
	}
	fn()
}
func TestUseFunctional(t *testing.T) {
	myFunc := Func14
	myFunc()

	func1 := Func15(6)
	fmt.Println(func1)

}
