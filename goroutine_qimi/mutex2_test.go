package goroutine_qimi

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// @progragm:   go-basic-exercises
// @file:      mutex2_test.go
// @author:    bowen
// @time:      2023-04-11 15:26
// @description:

var (
	mutex   sync.Mutex
	rwMutex sync.RWMutex
)

// 使用互斥锁进行的写操作
func writeWithLock() {
	mutex.Lock()
	x = x + 1
	time.Sleep(10 * time.Millisecond)
	mutex.Unlock()
	wg.Done()
}

// 使用互斥锁的读操作
func readWithLock() {
	mutex.Lock()
	time.Sleep(time.Millisecond)
	mutex.Unlock()
	wg.Done()
}

// 使用读写互斥锁的写操作
func writeWithRwLock() {
	rwMutex.Lock()
	x = x + 1
	time.Sleep(10 * time.Millisecond)
	rwMutex.Unlock()
	wg.Done()
}

// 使用读写互斥锁的读操作
func readWithRwLock() {
	rwMutex.RLock()
	time.Sleep(time.Millisecond)
	rwMutex.RUnlock()
	wg.Done()
}

func do(wf, rf func(), wc, rc int) {
	start := time.Now()
	// wc个并发写操作
	for i := 0; i < wc; i++ {
		wg.Add(1)
		go wf()
	}

	// rc个并发读操作
	for i := 0; i < rc; i++ {
		wg.Add(1)
		go rf()
	}
	wg.Wait()
	cost := time.Since(start)
	fmt.Printf("x:%v const:%v\n", x, cost)
}

func TestRW(t *testing.T) {
	do(writeWithLock, readWithLock, 100, 1000)
	do(writeWithRwLock, readWithRwLock, 100, 1000)
}

//func TestWithoutRw(t *testing.T) {
//	do(writeWithLock, readWithLock, 10, 100)
//}
