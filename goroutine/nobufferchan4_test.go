package goroutine

import (
	"fmt"
	"testing"
)

type counterchan struct {
	c chan int
	i int
}

func NewCounter() *counterchan {
	cter := &counterchan{
		c: make(chan int),
	}
	go func() {
		for {
			cter.i++
			cter.c <- cter.i
		}
	}()
	return cter
}

func (cter *counterchan) Increase() int {
	return <-cter.c
}

func TestCounterChan(t *testing.T) {
	cter := NewCounter()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			v := cter.Increase()
			fmt.Printf("goroutine-%d: current counter value is %d\n", i, v)
		}(i)
		wg.Done()
	}
	wg.Wait()
}
