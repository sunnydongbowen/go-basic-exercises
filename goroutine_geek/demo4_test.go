package goroutine_geek

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// @program:   go-basic-exercises
// @file:      demo4_test.go
// @author:    bowen
// @time:      2023-04-06 16:32
// @description:

func TestTimeOut(t *testing.T) {
	bg := context.Background()
	timeoutCtx, cancel_parent := context.WithTimeout(bg, time.Second)
	subCtx, cancel_child := context.WithTimeout(timeoutCtx, 10*time.Second)
	go func() {
		fmt.Println(<-subCtx.Done())
		fmt.Println("timeout")
	}()
	time.Sleep(8 * time.Second)
	//手动调用cancel
	cancel_child()
	errsub := subCtx.Err()
	fmt.Println(errsub)
	cancel_parent()
	//
}
