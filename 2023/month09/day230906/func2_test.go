package main

// @program:   go-basic-exercises
// @file:      func2_test.go
// @author:    bowen
// @time:      2023-09-06 10:36
// @description:

func re(n int) {
	if n > 10 {
		return
	}
	re(n + 1)
}

func A() {
	B()
}

func B() {
	A()
}

func C() {
	B()
}
