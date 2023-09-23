package main

import (
	"testing"
	"unicode/utf8"
)

// @program:   go-basic-exercises
// @file:      string_test.go
// @author:    bowen
// @time:      2023-09-05 15:17
// @description:

func TestString(t *testing.T) {
	// he said: "Hello go"
	println("he said: \"Hello go\"")
	println(len("abc"))
	println(len("你好"))
	println(utf8.RuneCountInString("你好"))

}
