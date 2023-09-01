package mutex_geek

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// @program:   go-basic-exercises
// @file:      mutex3_test.go
// @author:    bowen
// @time:      2023-04-13 17:03
// @description:
var wg sync.WaitGroup

func (s *SafeMap[K, V]) LoadOrstoreV2(key K, newValue V) (V, bool) {
	s.lock.RLock()
	oldVal, ok := s.values[key]
	s.lock.RUnlock()
	if ok {
		fmt.Printf("已经有了哦%v\n", s.values)
		return oldVal, true
	}
	time.Sleep(100 * time.Millisecond)

	s.lock.Lock()
	defer s.lock.Unlock()
	s.values[key] = newValue
	fmt.Printf("新的k-v已经放进去啦……%v\n", s.values)
	return newValue, false
}

func TestWithoutDouble(t *testing.T) {
	sm := SafeMap[string, int]{
		values: make(map[string]int, 4),
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Second)
			sm.LoadOrstoreV2("a", i)
		}(i)
	}
	time.Sleep(3 * time.Second)
	wg.Wait()
	fmt.Println(sm)
}
