package goroutine_qimi

import (
	"fmt"
	"sync"
	"testing"
)

// @program:   go-basic-exercises
// @file:      mutex_test.go
// @author:    bowen
// @time:      2023-04-11 10:58
// @description:

var (
	x int64
	m sync.Mutex
)

func add() {
	for i := 0; i < 5000; i++ {
		x = x + 1
	}
	wg.Done()
}

func addWithLock() {
	for i := 0; i < 10000; i++ {
		m.Lock()
		x = x + 1
		m.Unlock()
	}
	//fmt.Println(x)
	wg.Done()

}
func TestMutex(t *testing.T) {
	wg.Add(2)
	go addWithLock()
	go addWithLock()
	wg.Wait()
	fmt.Println(x)
}
