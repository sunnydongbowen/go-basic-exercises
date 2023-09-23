package main

// @program:   go-basic-exercises
// @file:      closure.go
// @author:    bowen
// @time:      2023-09-06 15:56
// @description:

func Closuer(name string) func() string {
	return func() string {
		return "hello " + name
	}
}

func Closure1() func() int {
	var age = 0
	return func() int {
		age++
		return age
	}
}
