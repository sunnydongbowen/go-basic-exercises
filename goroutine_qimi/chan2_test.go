package goroutine_qimi

import (
	"fmt"
	"testing"
)

// @program:   go-basic-exercises
// @file:      chan2_test.go
// @author:    bowen
// @time:      2023-04-10 10:41
// @description:

func f(ch chan int) {
	for {
		v, ok := <-ch
		if !ok {
			fmt.Printf("通道已关闭")
			break
		}
		fmt.Printf("value:%#v  ok:%#v\n", v, ok)
	}
}

func TestChan21(t *testing.T) {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)
	f(ch)
}

func TestChan23(t *testing.T) {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
LOOP:
	for {
		select {
		case x, ok := <-ch:
			fmt.Println(x, ok)
		default:
			break LOOP
		}
	}
}
