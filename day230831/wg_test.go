package day230831

import (
	"fmt"
	"sync"
	"testing"
)

// @program:   go-basic-exercises
// @file:      wg_test.go
// @author:    bowen
// @time:      2023-08-31 21:32
// @description: 复习笔记

var wg sync.WaitGroup

func hello() {
	fmt.Println("hello")
	wg.Done()
}

func TestWait(t *testing.T) {
	wg.Add(1)
	go hello()
	fmt.Println("main: hello")
	wg.Wait()
}
