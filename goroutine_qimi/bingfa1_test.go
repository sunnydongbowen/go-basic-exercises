package goroutine_qimi

import (
	"fmt"
	"testing"
	"time"
)

// @program:   go-basic-exercises
// @file:      bingfa1_test.go
// @author:    bowen
// @time:      2023-04-06 17:40
// @description:

func hello() {
	fmt.Println("hello")
}

func TestBin1(t *testing.T) {
	go hello()
	fmt.Println("main goroutine over")
	time.Sleep(time.Second)
}

func TestBin3(t *testing.T) {
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
}

// -------------------------------------------   //
func TestBin4(t *testing.T) {
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
}
