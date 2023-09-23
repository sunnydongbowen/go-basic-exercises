package main

// @program:   go-basic-exercises
// @file:      main.go
// @author:    bowen
// @time:      2023-09-06 15:56
// @description:

func main() {
	//name := Closuer("xxx")
	//fmt.Println(name)   // 这个是函数地址
	//fmt.Println(name()) //

	getAge := Closure1()
	println(getAge())
	println(getAge())
	println(getAge())
	println(getAge())
	println(getAge())
}
