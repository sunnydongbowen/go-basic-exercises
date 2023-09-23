package main

import "testing"

// @program:   go-basic-exercises
// @file:      byte_test.go
// @author:    bowen
// @time:      2023-09-05 15:32
// @description: 大明初级训练营基础课笔记

func TestByte(t *testing.T) {
	var a byte = 'a'
	println(a)

	var str string = "this is string"
	var bs []byte = []byte(str)
	bs[0] = 'M'
	println(bs)
}
