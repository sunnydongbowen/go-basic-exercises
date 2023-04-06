package goroutiere

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// @program:   go-basic-exercises
// @file:      deadline_test.go
// @author:    bowen
// @time:      2023-04-03 17:07
// @description:

func TestDeadline(t *testing.T) {
	d := time.Now().Add(5 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("0001")
		//fmt.Println(ctx.Err())
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
	//fmt.Println(ctx.Err())
}
