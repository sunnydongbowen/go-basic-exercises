package goroutine_qimi

import (
	"fmt"
	"testing"
)

// @program:   go-basic-exercises
// @file:      chan3_test.go
// @author:    bowen
// @time:      2023-04-10 11:00
// @description:

func Producer() chan int {
	ch := make(chan int, 2)
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 1 {
				ch <- i
			}
		}
		close(ch)
	}()
	return ch
}

func Consumer(ch chan int) int {
	sum := 0
	for v := range ch {
		sum += v
	}
	return sum
}

func TestChan31(t *testing.T) {
	ch := Producer()
	res := Consumer(ch)
	fmt.Println(res)
}
