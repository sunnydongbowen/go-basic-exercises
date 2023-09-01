package day230831

import (
	"fmt"
	"testing"
)

// @program:   go-basic-exercises
// @file:      for_test.go
// @author:    bowen
// @time:      2023-08-31 16:18
// @description:

func TestFor1(t *testing.T) {
	var i = 5
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

func TestFor2(t *testing.T) {
	var sl = []int{1, 2, 3, 4, 5}
	for i, v := range sl {
		fmt.Printf("sl[%d] = %d\n", i, v)
	}
}
