package slice

import (
	"fmt"
	"testing"
)

// @program:   go-basic-exercises
// @file:      qimi230419_test.go
// @author:    bowen
// @time:      2023-04-19 15:37
// @description:

// 切片初始化的方式
func TestSliceR(t *testing.T) {
	sint := []int{1, 2, 3}
	fmt.Println(sint)

	sm := make([]string, 10, 20)
	sm[0] = "bowen"

}
