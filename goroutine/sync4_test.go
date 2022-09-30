package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// sync.Mutex 实现对条件轮询等待
var ready bool

func worker3(i int) {
	fmt.Printf("worker %d: is working...\n", i)
	time.Sleep(1 * time.Second)
	fmt.Printf("worker %d: works done\n", i)
}

func spawnGroup3(f func(i int), num int, mu *sync.Mutex) <-chan signal {
	c := make(chan signal)
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			for {
				mu.Lock()
				if !ready {
					mu.Unlock()
					time.Sleep(100 * time.Millisecond)
					continue
				}
				mu.Unlock()
				fmt.Printf("worker %d: start to work...\n", i)
				f(i)
				wg.Done()
				return
			}

		}(i + 1)
	}
	go func() {
		wg.Wait()
		c <- signal(struct{}{})
	}()
	return c
}

func TestSync3(t *testing.T) {
	fmt.Println("start a group of workers...")
	mu := &sync.Mutex{}
	c := spawnGroup3(worker3, 5, mu)

	time.Sleep(5 * time.Second)
	fmt.Println("the group of workers start to work...")

	mu.Lock()
	ready = true
	mu.Unlock()

	<-c
	fmt.Println("the group of workers work done!")
}
