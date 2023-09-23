package main

// @program:   go-basic-exercises
// @file:      ifa.go
// @author:    bowen
// @time:      2023-09-08 15:40
// @description:

func IfElseif(age int) string {
	if age >= 18 {
		println("成年了")
	} else if age >= 20 {
		println("青年 ")
	} else {
		println("小孩子")
	}
	return ""
}

func IfNewVariable(start int, end int) string {
	if distance := end - start; distance > 100 {
		return "太远了"
	} else if distance > 60 {
		return "有点远"
	} else {
		return "还挺好"
	}
}
