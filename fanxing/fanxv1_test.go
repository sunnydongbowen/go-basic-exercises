package fanxing

import (
	"golang.org/x/exp/constraints"
	"testing"
)

// @program:     go-basic-exercises
// @file:        fanxv1_test.go
// @author:      bowen
// @create:      2022-12-09 20:07
// @description: tony白专栏

func Add[T constraints.Integer](a, b T) T {
	return a + b
}

func TestConsV1(t *testing.T) {
	var m, n int = 5, 6
	println(Add(m, n))

	var i, j int64 = 15, 16
	println(Add(i, j))

	var c, d byte = 0x11, 0x12
	println(Add(c, d))

}
