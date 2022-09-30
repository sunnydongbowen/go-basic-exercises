package main

import "fmt"

// 注释掉了，不然每次运行都会执行这个代码
func init() {
	fmt.Println("我是init")
}

func main() {
	fmt.Println("我是main")
}
