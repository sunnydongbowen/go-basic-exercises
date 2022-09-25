package function

import (
	"fmt"
	"testing"
)

type callback func(int, int) int

func doAdd(a, b int, f callback) int {
	fmt.Println("f callback")
	return f(a, b)
}

func add(a, b int) int {
	fmt.Println("add running")
	return a + b
}

func TestCallback(t *testing.T) {
	a, b := 2, 3
	// 回调了这个函数
	fmt.Println(doAdd(a, b, add))
	fmt.Println(doAdd(a, b, func(a int, b int) int {
		return a * b
	}))
}
