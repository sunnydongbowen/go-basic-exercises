package maptest

import (
	"fmt"
	"sync"
	"testing"
)

var wg sync.WaitGroup

func doIteration(m map[int]int) {
	for k, v := range m {
		_ = fmt.Sprintf("[%d,%d] ", k, v)
	}
}
func doWrite(m map[int]int) {
	for k, v := range m {
		m[k] = v + 1
	}
}

func TestMapGoroutine(t *testing.T) {
	m := map[int]int{
		1: 11,
		2: 12,
		3: 13,
	}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			doIteration(m)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			doWrite(m)
		}
	}()
	wg.Wait()
}
