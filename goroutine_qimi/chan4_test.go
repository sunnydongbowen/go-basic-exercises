package goroutine_qimi

import (
	"fmt"
	"testing"
)

// @program:   go-basic-exercises
// @file:      chan4_test.go
// @author:    bowen
// @time:      2023-04-11 9:43
// @description:

func Producer2() <-chan int {
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

func Consumer2(ch <-chan int) int {
	sum := 0
	for v := range ch {
		sum += v
	}
	return sum
}

func TestChan41(t *testing.T) {
	ch := Producer()
	sum := Consumer(ch)
	fmt.Println(sum)
}

func TestChan42(t *testing.T) {
	var ch1 = make(chan int, 1)
	ch1 <- 10
	close(ch1)
	// 函数传参时将ch3转为单向通道
	res := Consumer2(ch1)
	fmt.Println(res)

	ch2 := make(chan int, 1)
	ch2 <- 12
	//声明一个只接收通道
	var ch3 <-chan int
	// 赋值时转为单向通道
	ch3 = ch2
	v := <-ch3
	fmt.Println(v)
}
