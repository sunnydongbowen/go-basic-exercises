package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func worker1(i int) {
	fmt.Printf("worker %d: is working...\n", i)
	time.Sleep(1 * time.Second)
	fmt.Printf("worker %d: works done\n", i)
}

func spawnGroup(f func(i int), num int, grouSignal <-chan signal) <-chan signal {

	c := make(chan signal)
	// Num个goroutine
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			<-grouSignal //阻塞点
			fmt.Printf("worker %d: start to work...\n", i)
			f(i)
			wg.Done()
		}(i + 1)
	}

	go func() {
		wg.Wait()
		c <- signal{}
	}()
	return c
}

func TestSignal2(t *testing.T) {
	fmt.Println("start a group of workers...")
	groupSignal := make(chan signal)
	c := spawnGroup(worker1, 5, groupSignal)
	time.Sleep(5 * time.Second)
	fmt.Println("the group of workers start to work...")
	close(groupSignal) // 关闭了
	<-c
	fmt.Println("the group of workers work done!")
}
