package main

import (
	"fmt"
	"testing"
)

// @program:   go-basic-exercises
// @file:      func_test.go
// @author:    bowen
// @time:      2023-09-06 10:10
// @description:

func Func8() (name string, age int) {
	return "xxx", 19
}

func Func9() (name string, age int) {
	name = "ddd"
	age = 19
	return
}

func Func10() (name string, age int) {
	return
}

func Func6(s string) string {
	return s
}

func TestFunc(t *testing.T) {
	name, age := Func9()
	fmt.Println(name, age)
	name1, age := Func10()
	fmt.Println(name1, age)
	Func9()
}
