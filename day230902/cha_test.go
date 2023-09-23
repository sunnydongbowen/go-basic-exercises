package day230902

import (
	"fmt"
	"testing"
)

// @program:   go-basic-exercises
// @file:      cha_test.go
// @author:    bowen
// @time:      2023-09-02 20:34
// @description: 笔记复习

func TestCha(t *testing.T) {
	cha1 := make(chan int)
	cha1 <- 13
	n := <-cha1
	println(n)
}

func TestCha2(t *testing.T) {
	ch1 := make(chan int)
	go func() {
		ch1 <- 13
	}()
	n := <-ch1
	fmt.Println(n)
}
