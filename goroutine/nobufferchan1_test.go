package goroutine

import (
	"fmt"
	"testing"
	"time"
)

//geektime 专栏

type signal struct {
}

func worker() {
	println("worker is working...")
	time.Sleep(time.Second)
}

func spawn(f func()) <-chan signal {
	c := make(chan signal)
	go func() {
		println("worker start to work")
		f()
		c <- signal{}
	}()
	return c
}

func TestNobufferChan(t *testing.T) {
	println("start a worker")
	c := spawn(worker)
	<-c
	fmt.Println("worker work done!")
}
