package day230901

import (
	"fmt"
	"sync"
	"testing"
)

// @program:   go-basic-exercises
// @file:      pool_test.go
// @author:    bowen
// @time:      2023-09-01 20:02
// @description: 复习笔记

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() any {
			fmt.Println("hhh new")
			return []byte{}
		},
	}
	for i := 0; i < 5; i++ {
		val := pool.Get()
		fmt.Println(val)
	}
}
