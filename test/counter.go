package test

import (
	"sync"
	"sync/atomic"
)

type bench struct {
	lock    sync.Mutex
	counter int32
	wg      sync.WaitGroup
}

func Counter() int32 {
	var bench bench
	for i := 0; i < 10; i++ {
		bench.wg.Add(1)
		go func() {
			bench.lock.Lock()
			bench.counter++
			bench.lock.Unlock()
			bench.wg.Done()
		}()
	}
	bench.wg.Wait()
	return bench.counter
}

func Counter2() int32 {
	var bench bench
	for i := 0; i < 10; i++ {
		bench.wg.Add(1)
		go func() {
			atomic.AddInt32(&bench.counter, 1)
			bench.wg.Done()
		}()
	}
	bench.wg.Wait()
	return bench.counter
}
