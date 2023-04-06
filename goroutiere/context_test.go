package goroutiere

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// @program:   go-basic-exercises
// @file:      context_test.go
// @author:    bowen
// @time:      2023-04-03 15:06
// @description:
func con2(ctx context.Context) {
	defer wg.Done()
FORLOOP:
	for {
		fmt.Println("002")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break FORLOOP
		default:
		}
	}
}

func fcon(ctx context.Context) {
	defer wg.Done()
	wg.Add(1)
	go con2(ctx)
FORLOOP:
	for {
		fmt.Println("001")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break FORLOOP
		default:
		}
	}
}

func TestContext1(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go fcon(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
}
