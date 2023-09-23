package main

// @program:   go-basic-exercises
// @file:      defer.go
// @author:    bowen
// @time:      2023-09-07 19:43
// @description:

func Defer() {
	defer func() {
		println("第1个defer")
	}()

	defer func() {
		println("第2个defer")
	}()
}

func DeferClosure() {
	i := 0
	defer func() {
		println(i)
	}()
	i = 1
}

func DeferClosure2() {
	i := 0
	defer func(i int) {
		println(i)
	}(i)
	i = 1
}

func DeferReturn() int {
	a := 0
	defer func() {
		a = 1
	}()
	return a
}

func DeferReturn1() (a int) {
	a = 0
	defer func() {
		a = 1
	}()
	return a
}
