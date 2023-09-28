package algorithm

import (
	"fmt"
	"testing"
)

// @program:   go-basic-exercises
// @file:      求数组元素之和_test.go
// @author:    bowen
// @time:      2023-09-23 20:34
// @description: 3

func Sum(a []int) int {
	var sum int
	for _, v := range a {
		sum += v
	}
	return sum
}

func TestSum(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(Sum(a))
}
