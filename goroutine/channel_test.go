package goroutine

import (
	"fmt"
	"sync"
	"testing"
)

var wg sync.WaitGroup

// deadlock
func TestChan1(t *testing.T) {
	ch1 := make(chan int)
	ch1 <- 13 // fatal error: all goroutines are asleep - deadlock!    n := <-ch1    println(n)
	n := <-ch1
	println(n)
}

// 一发一收才可以
func TestChan2(t *testing.T) {
	ch1 := make(chan int)
	go func() {
		ch1 <- 13
	}()
	n := <-ch1
	println(n)
}

func TestChan3(t *testing.T) {
	ch1 := make(chan int, 1)
	n := <-ch1 //由于此时ch2的缓冲区中无数据，因此对其进行接收操作将导致goroutine挂起
	fmt.Println(n)

	ch3 := make(chan int, 1)
	ch3 <- 17 //  向ch3发送一个整型数17
	ch3 <- 27 //  由于此时ch3中缓冲区已满，再向ch3发送数据也将导致goroutine挂起

}

func TestChan5(t *testing.T) {
	var ch chan int
	ch <- 1
}

func TestCHAN6(t *testing.T) {
	var ch chan int
	<-ch
}

func TestChan7(t *testing.T) {
	ch := make(chan int)
	<-ch
}

func TestChan8(t *testing.T) {
	var a int
	ch := make(chan int)
	go func() {
		ch <- 16
	}()
	a = <-ch
	fmt.Println(a)

	close(ch)
	fmt.Println(<-ch)
}

func TestChan9(t *testing.T) {
	var a int
	ch := make(chan int)
	go func() {
		ch <- 16
	}()
	a = <-ch
	fmt.Println(a)

	close(ch)
	ch <- 1
}
