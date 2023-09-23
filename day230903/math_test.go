package day230903

import (
	"fmt"
	"testing"
)

// @program:   go-basic-exercises
// @file:      math_test.go
// @author:    bowen
// @time:      2023-09-03 10:16
// @description:

func TestEle(t *testing.T) {
	var x = 45
	var y = 15
	for i := 0; i < 15; i++ {
		if int(x/y) == 11 && int(x%y) == 0 {
			fmt.Printf("%d 年前，父亲的年龄是儿子的11倍", i)
			fmt.Printf("此时父亲的年龄是%d，儿子的年龄是%d", x, y)
			break
		}
		x--
		y--
	}
}
