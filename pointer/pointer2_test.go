package pointer

import (
	"fmt"
	"testing"
)

// @program:   go-basic-exercises
// @file:      pointer2_test.go
// @author:    bowen
// @time:      2023-03-27 17:51
// @description: 复习notion

func TestPointer(t *testing.T) {
	var a *int
	fmt.Println(a) //nil

	var a1 = new(int)
	fmt.Println(a1, *a1) // 这个是有内存地址的，初始值是0

	*a1 = 100
	fmt.Println(a1, *a1)

}

func TestPointer2(t *testing.T) {
	n := 18
	p := &n // 取地址
	fmt.Printf("%v,%T\n", p, p)
}

//	func  convert(a int)    *int{
//		return *a
//	}
func TestP3(t *testing.T) {
	c := 100
	fmt.Print(c, *(&c))
}
