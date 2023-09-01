package day230831

import (
	"fmt"
	"testing"
)

// @program:   go-basic-exercises
// @file:      cha_test.go
// @author:    bowen
// @time:      2023-08-31 21:27
// @description:

func TestZuse(t *testing.T) {
	ch := make(chan int, 1)
	ch <- 1
	for {
		select {
		case x, ok := <-ch:
			fmt.Println(x, ok)
		default:
		}
	}

}
