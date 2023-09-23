package main

import "fmt"

// @program:   go-basic-exercises
// @file:      en.go
// @author:    bowen
// @time:      2023-09-05 15:05
// @description: 大明初级课笔记

func main() {
	var a int = 123
	var b int = 456
	fmt.Println(a - b)
	fmt.Println(a + b)
	fmt.Println(a * b)
	fmt.Println(a % b)
	if b != 0 {
		fmt.Println(a / b)
	}
}
