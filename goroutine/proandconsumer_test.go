package goroutine

import (
	"testing"
	"time"
)

// @program:   go-basic-exercises
// @file:      proandconsumer_test.go
// @author:    bowen
// @time:      2023-04-15 20:41
// @description:

func produce(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i + 1
		time.Sleep(time.Second)
	}
	close(ch)
	defer wg.Done()
}

func consume(ch <-chan int) {
	for v := range ch {
		println(v)
	}
	defer wg.Done()
}
func TestProducer(t *testing.T) {
	ch := make(chan int, 5)
	wg.Add(2)

	go produce(ch)
	go consume(ch)
	wg.Wait()

}
