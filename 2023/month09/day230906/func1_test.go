package main

import "strings"

// @program:   go-basic-exercises
// @file:      func1_test.go
// @author:    bowen
// @time:      2023-09-06 10:22
// @description:

func Func12(abc string) (string, int) {
	segs := strings.Split(abc, " ")
	return segs[0], len(segs)
}

func Func13(abc string) (first string, length int) {
	segs := strings.Split(abc, " ")
	first = segs[0]
	length = len(segs)
	return
}
