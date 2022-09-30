package goroutine

import (
	"log"
	"testing"
	"time"
)

var active = make(chan struct{}, 3)
var jobs = make(chan int, 10)

func TestBufferChan(t *testing.T) {
	go func() {
		for i := 0; i < 8; i++ {
			jobs <- (i + 1)
		}
		close(jobs)
	}()

	for j := range jobs {
		wg.Add(1)
		go func(j int) {
			active <- struct{}{}
			log.Printf("handle job: %d\n", j)
			time.Sleep(2 * time.Second)
			<-active
			wg.Done()
		}(j)
	}
	wg.Wait()
}
