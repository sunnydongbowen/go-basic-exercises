package mutex_geek

import (
	"fmt"
	"sync"
	"testing"
)

// @program:   go-basic-exercises
// @file:      mutex2_test.go
// @author:    bowen
// @time:      2023-04-13 15:59
// @description:

type SafeMap[K comparable, V any] struct {
	values map[K]V
	lock   sync.RWMutex
}

func (s *SafeMap[K, V]) LoadOrstoreV1(key K, newValue V) (V, bool) {
	s.lock.RLock()
	oldVal, ok := s.values[key]
	s.lock.RUnlock()
	if ok {
		return oldVal, true
	}
	// 加写锁
	fmt.Println("我要加写锁啦")
	s.lock.Lock()
	fmt.Println("我已经加完写锁啦")
	defer s.lock.Unlock()

	//double check
	oldVal, ok = s.values[key]
	if ok {
		return oldVal, true
	}
	s.values[key] = newValue
	return newValue, false
}

func TestDeferRLock(t *testing.T) {
	sm := SafeMap[string, string]{
		values: make(map[string]string, 4),
	}
	sm.LoadOrstoreV1("a", "b")
	fmt.Println(sm.values)
}
