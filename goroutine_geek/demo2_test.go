package goroutine_geek

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// @program:   go-basic-exercises
// @file:      demo2_test.go
// @author:    bowen
// @time:      2023-04-04 14:59
// @description:
func TestContext1(t *testing.T) {
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, time.Minute*3)
	defer cancel()
	dl, ok := timeoutCtx.Deadline()
	fmt.Println(dl, ok)
}

func TestContext2(t *testing.T) {
	ctx := context.Background()
	dl, ok := ctx.Deadline()
	fmt.Println(dl, ok)
}

func TestContext3(t *testing.T) {
	ctx := context.Background()
	valCtx := context.WithValue(ctx, "abc", 123)
	val := valCtx.Value("abc") //123
	fmt.Println(val)
}

func TestContext4(t *testing.T) {
	ctx := context.Background()
	dlCtx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Minute))
	childCtx := context.WithValue(dlCtx, "key", 123)
	cancel()
	err := childCtx.Err()
	fmt.Println(err)
}

func TestContext5(t *testing.T) {
	ctx := context.Background()
	childCtx := context.WithValue(ctx, "key1", 123)
	ccCtx := context.WithValue(childCtx, "key2", 12345)

	val := childCtx.Value("key2") // 拿它儿子的肯定拿不到
	fmt.Println(val)
	val = ccCtx.Value("key1") //
	fmt.Println(val)

}

func TestContext6(t *testing.T) {
	ctx := context.Background()
	parent := context.WithValue(ctx, "par-key", "par-value")
	child := context.WithValue(parent, "child-key", "child-value")
	fmt.Printf("父亲拿自己的: %v\n", parent.Value("par-key"))
	fmt.Printf("父亲拿儿子的: %v\n", parent.Value("child-key"))
	fmt.Printf("儿子拿自己的: %v\n", child.Value("child-key"))
	fmt.Printf("儿子拿父亲的: %v\n", parent.Value("par-key"))
}

func TestContext7(t *testing.T) {
	ctx := context.Background()
	par := context.WithValue(ctx, "map", map[string]string{})
	child := context.WithValue(par, "key1", "value1")
	m := child.Value("map").(map[string]string)
	fmt.Println(m)
	m["key1"] = "value2"
	val := par.Value("key1")
	fmt.Println(val)
	val = par.Value("map")
	fmt.Println(val)
}
