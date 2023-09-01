package goroutine_qimi

import (
	"fmt"
	"testing"
)

// @program:   go-basic-exercises
// @file:      chan1_test.go
// @author:    bowen
// @time:      2023-04-07 20:36
// @description:

func TestChan1(t *testing.T) {
	var ch chan int
	fmt.Printf("%T %v\n", ch, ch)
}

func TestChan2(t *testing.T) {
	var ch chan int
	fmt.Println(ch) //nil

	ch = make(chan int, 1)
	fmt.Println(ch) // 0xc00001a230

	fmt.Println(<-ch)
	ch <- 10
	fmt.Println(ch) // 0xc00001a230
}

func TestChan3(t *testing.T) {
	ch := make(chan int)
	ch <- 10
	fmt.Println("send")
}

// -------------------------------------------   //

func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}

func TestChan4(t *testing.T) {
	ch := make(chan int)
	go recv(ch)
	ch <- 10
	fmt.Println("sended")
}

func TestChan5(t *testing.T) {
	ch := make(chan int, 1)
	ch <- 10
	fmt.Println("send")
	ch <- 12 //dead lock
}
