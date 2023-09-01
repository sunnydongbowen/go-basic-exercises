package goroutine

import (
	"fmt"
	"sync"
	"testing"
)

type counter struct {
	sync.Mutex
	i int
}

var cter counter

// 加锁解锁，i+1
func Increase() int {
	cter.Lock()
	defer cter.Unlock()
	cter.i++
	return cter.i
}

func TestNoBuffer(t *testing.T) {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			v := Increase()
			fmt.Printf("goroutine-%d: current counter value is %d\n", i, v)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
