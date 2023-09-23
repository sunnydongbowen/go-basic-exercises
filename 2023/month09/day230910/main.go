package main

import "fmt"

// @program:   go-basic-exercises
// @file:      main.go
// @author:    bowen
// @time:      2023-09-10 14:24
// @description:

func main() {
	//var ov1 OuterV1
	//ov1.Dosth() //
	//ov1.Inner.Dosth()
	var array = []int{1, 1, 3, 5, 1, 4, 3, 2}

	m := map[int]bool{}

	for _, v := range array {
		if m[v] == false {
			m[v] = true
		} else {
			fmt.Println(v)
		}
	}
}
