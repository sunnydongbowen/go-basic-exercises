package slice

import (
	"fmt"
	"testing"
	"unsafe"
)

// @program:     go-basic-exercises
// @file:        tony_test.go
// @author:      bowen
// @create:      2022-12-01 15:10
// @description: tony白专栏示例

// 数组转切片
func TestSlice1(t *testing.T) {
	a := [3]int{11, 12, 13}
	b := a[:] // 通过切片将数组转换为切片
	b[1] += 10
	fmt.Printf("%v\n", a)
}

// unsafe 将切片转数组
func TestSlice2(t *testing.T) {
	b := []int{11, 12, 13}
	// 这里要理解一下了！
	var p = (*[3]int)(unsafe.Pointer(&b[0]))
	p[1] += 10
	fmt.Printf("%v\n", b)
}

// 黑魔法，直接将切片转为数组
func TestSlice3(t *testing.T) {
	b := []int{11, 12, 13}
	p := (*[3]int)(b)
	p[1] = p[1] + 10
	fmt.Printf("%v\n", b)
	// *[3]int,数组指针
	fmt.Printf("%T\n", p)
}
