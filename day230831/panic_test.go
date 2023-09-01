package day230831

import (
	"errors"
	"fmt"
	"testing"
)

// @program:   go-basic-exercises
// @file:      panic_test.go
// @author:    bowen
// @time:      2023-08-31 21:43
// @description:  复习笔记

func TestErrors(t *testing.T) {
	var a int
	a = 10
	if a < 18 {
		fmt.Println(errors.New("不能玩手机哦"))
	}
	fmt.Println("over")
}

func TestPanic(t *testing.T) {
	var a int
	a = 10
	if a < 18 {
		//自定义了painc
		panic("不能玩手机哦")
	}
	fmt.Println("over")
}
