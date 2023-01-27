package main

import (
	"fmt"
	"os"
)

// @program:   go-basic-exercises
// @file:      main.go
// @author:    bowen
// @time:      2023/1/22 11:59
// @description: go语言圣经ch1

func main() {
	var s, seq string
	//s = os.Args[0]
	for i := 0; i < len(os.Args); i++ {
		s += seq + os.Args[i]
		seq = " "
		fmt.Println(i, s)
	}
	//fmt.Println(s)
}
