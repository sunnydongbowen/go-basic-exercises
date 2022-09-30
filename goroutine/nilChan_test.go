package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestChan(t *testing.T) {
	ch1, ch2 := make(chan int), make(chan int)
	go func() {
		time.Sleep(time.Second * 5)
		ch1 <- 5
		close(ch1)
	}()
	go func() {
		time.Sleep(time.Second * 7)
		ch2 <- 7
		close(ch2)
	}()
	var ok1, ok2 bool
	for {
		select {
		case x := <-ch1:
			ok1 = true
			time.Sleep(100 * time.Millisecond)
			fmt.Println(x)
		case x := <-ch2:
			ok2 = true
			time.Sleep(100 * time.Millisecond)
			fmt.Println(x)
		}
		if ok1 && ok2 {
			break
		}
	}
	fmt.Println("program end")
}

func TestNilChan(t *testing.T) {
	ch1, ch2 := make(chan int), make(chan int)
	go func() {
		time.Sleep(time.Second * 5)
		ch1 <- 5
		close(ch1)
	}()
	go func() {
		time.Sleep(time.Second * 7)
		ch2 <- 7
		close(ch2)
	}()
	for {
		select {
		case x, ok := <-ch1:
			if !ok {
				ch1 = nil
			} else {
				fmt.Println(x)
			}
		case x, ok := <-ch2:
			if !ok {
				ch2 = nil
			} else {
				fmt.Println(x)
			}
		}
		if ch1 == nil && ch2 == nil {
			break
		}
	}
	fmt.Println("program end")

}
