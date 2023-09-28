package algorithm

import (
	"fmt"
	"testing"
)

// @program:   go-basic-exercises
// @file:      计算年龄倍数_test.go
// @author:    bowen
// @time:      2023-09-23 20:24
// @description:  2

func TestAge(t *testing.T) {
	var a = 45
	var b = 15

	for i := 0; i < 15; i++ {
		if int(a/b) == 11 && int(a%b) == 0 {
			fmt.Printf("%d年前，父亲的岁数是孩子的11倍，此时孩子%d岁，父亲%d岁\n", i, b, a)
			break
		}
		a--
		b--
	}

}
