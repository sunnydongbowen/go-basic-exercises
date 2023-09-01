package goroutine_qimi

import (
	"fmt"
	"testing"
)

// @program:   go-basic-exercises
// @file:      select_test.go
// @author:    bowen
// @time:      2023-04-11 10:44
// @description:

func TestSelect(t *testing.T) {
	ch := make(chan int, 1)
	for i := 1; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}

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
