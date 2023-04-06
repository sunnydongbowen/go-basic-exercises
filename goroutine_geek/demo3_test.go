package goroutine_geek

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// @program:   go-basic-exercises
// @file:      demo3_test.go
// @author:    bowen
// @time:      2023-04-04 20:51
// @description:
func TestBussinessTimeout(t *testing.T) {
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, time.Second)
	end := make(chan struct{}, 1)
	go func() {
		Mybusiness()
		end <- struct{}{}
	}()
	select {
	case <-timeoutCtx.Done():
		fmt.Println("timeout")
	case <-end:
		fmt.Println("business end")
	}
	cancel()
	err := timeoutCtx.Err()
	fmt.Println(err)
}

func Mybusiness() {
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("hello Go")
}
