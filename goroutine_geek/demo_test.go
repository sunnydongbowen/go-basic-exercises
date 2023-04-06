package goroutine_geek

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// @program:   go-basic-exercises
// @file:      demo_test.go
// @author:    bowen
// @time:      2023-04-04 9:58
// @description:

func TestERR(t *testing.T) {
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	time.Sleep(2 * time.Second)
	err := timeoutCtx.Err()
	fmt.Println(err)
}

func TestERR1(t *testing.T) {
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, time.Second)
	time.Sleep(500 * time.Millisecond)
	cancel()
	err := timeoutCtx.Err()
	fmt.Println(err)
}
