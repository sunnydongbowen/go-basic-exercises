package main

import (
	"fmt"
	"os"
)

// @program:   go-basic-exercises
// @file:      main.go
// @author:    bowen
// @time:      2023/1/24 12:12
// @description: go语言圣经

func main() {
	s, seq := "", ""
	for _, arg := range os.Args[1:] {
		s += seq + arg
		seq = " "
		fmt.Println(s)
	}
}
