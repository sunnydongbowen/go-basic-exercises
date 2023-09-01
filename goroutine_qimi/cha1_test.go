package goroutine_qimi

import (
	"fmt"
	"testing"
)

// @program:   go-basic-exercises
// @file:      cha1_test.go
// @author:    bowen
// @time:      2023-08-31 18:06
// @description:

func f3(c chan int) {
	for v := range c {
		fmt.Println(v)
	}
}

func TestCh1(t *testing.T) {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)
	f3(ch)
}
