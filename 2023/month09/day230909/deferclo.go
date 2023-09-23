package main

import "fmt"

// @program:   go-basic-exercises
// @file:      deferclo.go
// @author:    bowen
// @time:      2023-09-09 9:57
// @description:

func DeferClosuerLoopV3() {
	//var j int
	for i := 0; i < 10; i++ {
		j := i
		defer func() {
			fmt.Printf("j的地址是%p，值是%d\n", &j, j)
		}()
	}
}

func DeferClosuerLoopV2() {
	var j int
	for i := 0; i < 10; i++ {
		j = i
		defer func() {
			fmt.Printf("j的地址是%p，值是%d\n", &j, j)
		}()
	}
}
