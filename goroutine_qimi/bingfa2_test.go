package goroutine_qimi

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

// @program:   go-basic-exercises
// @file:      bingfa2_test.go
// @author:    bowen
// @time:      2023-04-07 11:10
// @description:
var wg sync.WaitGroup

func hello3() {
	fmt.Println("hello")
	wg.Done() // 告知当前goroutine完成
}

func TestWait(t *testing.T) {
	wg.Add(1)
	go hello3()
	fmt.Println("main:hello")
	wg.Wait() //阻塞等待登记的goroutine完成
}

// -------------------------------------------   //
func helloPra(i int) {
	defer wg.Done()
	fmt.Println("hello", i)
}

func TestWait1(t *testing.T) {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go helloPra(i)
	}
	wg.Wait()
}

// -------------------------------------------   //

func a() {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		time.Sleep(time.Nanosecond)
		fmt.Printf("A:%d\n", i)
	}
}

func b() {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		time.Sleep(time.Nanosecond)
		fmt.Printf("B:%d\n", i)
	}
}

func TestGoMax(t *testing.T) {
	runtime.GOMAXPROCS(6)
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}
