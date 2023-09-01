package mutex_geek

import (
	"fmt"
	"sync"
	"testing"
)

// @program:   go-basic-exercises
// @file:      mutex1_test.go
// @author:    bowen
// @time:      2023-04-13 13:53
// @description:
type safeResource struct {
	resource map[string]string
	lock     sync.RWMutex
}

func (s *safeResource) Add(key string, value string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.resource[key] = value
}

func TestMu(t *testing.T) {
	var s safeResource
	s.resource = map[string]string{}
	s.Add("name", "Jack")
	fmt.Println(s.resource)
}

func TestMu2(t *testing.T) {
	ss := safeResource{
		resource: map[string]string{}, // 一定要初始化，不然会报错
		lock:     sync.RWMutex{},
	}
	ss.Add("key", "value")
	fmt.Println(ss.resource)
}
