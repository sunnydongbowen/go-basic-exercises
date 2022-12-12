package fanxing

import (
	"fmt"
	"testing"
)

// @program:     go-basic-exercises
// @file:        maxgeneric_test.go
// @author:      bowen
// @create:      2022-12-11 12:52
// @description:

type ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
}

func maxGenerics[T ordered](sl []T) T {
	if len(sl) == 0 {
		panic("slice is empty")
	}
	max := sl[0]
	for _, v := range sl[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

type mystring string

func TestMaxGeneric(t *testing.T) {
	var m int = maxGenerics([]int{1, 2, -4, -6, 7, 3})
	fmt.Println(m)

	fmt.Println(maxGenerics([]string{"11", "22", "44", "66", "77", "10"}))
	fmt.Println(maxGenerics([]float64{1.01, 2.02, 3.03, 5.05, 7.07, 0.01}))
	fmt.Println(maxGenerics([]int8{1, 2, -4, -6, 7, 0}))
	fmt.Println(maxGenerics([]mystring{"11", "22", "44", "66", "77", "10"}))

}
