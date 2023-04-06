package goroutine_geek

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

// @program:   go-basic-exercises
// @file:      demo1_test.go
// @author:    bowen
// @time:      2023-04-04 10:57
// @description:
var wg sync.WaitGroup
var notify bool

func TestMerror(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	wg.Add(1)
	go worker(ctx)
	go func() {
		for {
			if notify {
				cancel()
			}
		}

	}()
	time.Sleep(time.Second * 5)
	err := ctx.Err()
	fmt.Println(err)
	wg.Wait()
	fmt.Println("over")

}

func worker(ctx context.Context) {
LOOP:
	for {
		fmt.Println("db connecting")
		notify = true
		time.Sleep(time.Millisecond * 10)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
	fmt.Println("worker done")
	wg.Done()
}
