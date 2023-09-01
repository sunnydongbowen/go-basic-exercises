package mutex_geek

import (
	"fmt"
	"sync"
	"testing"
)

// @program:   go-basic-exercises
// @file:      mutex5_test.go
// @author:    bowen
// @time:      2023-04-13 20:14
// @description:

type OnceClose struct {
	close sync.Once
}

func (o *OnceClose) Close() error {
	o.close.Do(func() {
		fmt.Println("close")
	})
	return nil
}

func (no OnceClose) Close1() error {
	no.close.Do(func() {
		fmt.Println("Close")
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

func TestOnceClose1(t *testing.T) {
	oc := OnceClose{}
	for i := 0; i < 10; i++ {
		oc.Close1()
	}
}
