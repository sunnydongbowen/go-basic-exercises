package day230903

import (
	"fmt"
	"testing"
)

// @program:   go-basic-exercises
// @file:      sum_test.go
// @author:    bowen
// @time:      2023-09-03 11:19
// @description: 笔记复习

func TestSum1(t *testing.T) {
	var a []int
	a = []int{1, 2, 3, 4, 5}
	var sum int
	for _, v := range a {
		sum += v
	}
	fmt.Println(sum)
}

func TestSum2(t *testing.T) {
	var a []int
	a = []int{1, 2, 3, 4, 5}
	var sum int
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	fmt.Println(sum)
}
