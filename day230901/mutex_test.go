package day230901

import (
	"fmt"
	"sync"
	"testing"
)

// @program:   go-basic-exercises
// @file:      mutex_test.go
// @author:    bowen
// @time:      2023-09-01 19:40
// @description: 笔记复习

type safeResource struct {
	resoure map[string]string
	lock    sync.RWMutex
}

func (s *safeResource) Add(key string, value string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.resoure[key] = value
}

func TestMu(t *testing.T) {
	var s safeResource
	s.resoure = map[string]string{}
	s.Add("name", "enenen")
	fmt.Println(s.resoure)
}

func TestMu1(t *testing.T) {
	ss := safeResource{
		resoure: map[string]string{},
		lock:    sync.RWMutex{},
	}
	ss.Add("key", "value")
	fmt.Println(ss.resoure)
}
