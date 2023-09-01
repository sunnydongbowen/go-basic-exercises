package goroutine_qimi

import (
	"fmt"
	"testing"
	"time"
)

// @program:   go-basic-exercises
// @file:      goroutine1_test.go
// @author:    bowen
// @time:      2023-08-31 14:57
// @description:

func printHello(i int) {
	fmt.Println("hello", i)
}

func TestGoroutine1(t *testing.T) {
	for i := 0; i < 100; i++ {
		go printHello(i)
	}
	fmt.Println("main")
	time.Sleep(time.Second)
}

func TestPrint(t *testing.T) {
	fmt.Println(123)
}
