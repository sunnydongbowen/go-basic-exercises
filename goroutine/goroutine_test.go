package goroutine

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func spawn1(f func() error) <-chan error {
	c := make(chan error)

	go func() {
		c <- f()
	}()
	return c
}

func TestGorou(t *testing.T) {
	var 小明 int = 5
	fmt.Println(小明)
	c := spawn1(func() error {
		time.Sleep(2 * time.Second)
		return errors.New("timeout")
	})
	fmt.Println(<-c)
}
