package goroutiere

import (
	"context"
	"fmt"
	"testing"
)

// @program:   go-basic-exercises
// @file:      official_test.go
// @author:    bowen
// @time:      2023-04-03 16:28
// @description:

func TestOfficial(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}
