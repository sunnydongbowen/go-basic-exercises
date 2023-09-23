package day230902

import (
	"fmt"
	"sync"
	"testing"
)

// @program:   go-basic-exercises
// @file:      once_test.go
// @author:    bowen
// @time:      2023-09-02 10:48
// @description: Once笔记复习，大明课程

type OnceClose struct {
	close sync.Once
}

func (o *OnceClose) Close() error {
	o.close.Do(func() {
		fmt.Println("close")
	})
	return nil
}

func TestOnceClose(t *testing.T) {
	oc := &OnceClose{}
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		oc.Close()
	}
}

func (no OnceClose) Close1() error {
	no.close.Do(func() {
		fmt.Println("Close")
	})
	return nil
}

func TestClose1(t *testing.T) {
	oc := &OnceClose{}
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		oc.Close1()

	}

}
