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
	println("worker is working")
	time.Sleep(1 * time.Second)
}

func spawn(f func()) <-chan signal {
	c := make(chan signal)
	go func() {
		println("worker start to work...")
		f()
		c <- signal{}
	}()
	return c
}

func TestSignal(t *testing.T) {
	println("main start a worker.....")
	c := spawn(worker)
	<-c
	fmt.Println("worker work done!")

}
