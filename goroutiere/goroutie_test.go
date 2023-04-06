package goroutiere

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// @program:   go-basic-exercises
// @file:      goroutie_test.go
// @author:    bowen
// @time:      2023-04-03 15:52
// @description:

func sun(ctx context.Context) error {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("孙子goroutine: 我饿了")
		time.Sleep(500 * time.Millisecond)
		select {
		case <-ctx.Done():
			fmt.Println("孙子goroutine: 好的，我不叫了")
			break LOOP
		default:
		}
	}
	return ctx.Err()
}

func fa(ctx context.Context) error {
	defer wg.Done()
	wg.Add(1)
	go sun(ctx)
LOOP:
	for {
		fmt.Println("儿子goroutine: 我饿了")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			fmt.Println("儿子goroutine: 好的，我不叫了")
			time.Sleep(10 * time.Nanosecond)
			break LOOP
		default:
		}
	}
	return ctx.Err()
}

func TestGoroutie(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go fa(ctx)
	time.Sleep(5 * time.Second)
	fmt.Println("main:好了，别叫了，开饭了！")
	cancel()
	wg.Wait()
}
