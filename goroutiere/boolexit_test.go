package goroutiere

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// @program:   go-basic-exercises
// @file:      boolexit_test.go
// @author:    bowen
// @time:      2023-04-03 11:29
// @description:

var wg sync.WaitGroup
var notify bool

func f() {
	defer wg.Done() //不加defer的话直接死锁了
	for {
		fmt.Println("bowen")
		time.Sleep(time.Millisecond * 500)
		if notify {
			break
		}
	}
}

func TestBoolExit(t *testing.T) {
	wg.Add(1)
	go f()
	time.Sleep(5 * time.Second)
	// 通知退出
	notify = true
	wg.Wait()
}
