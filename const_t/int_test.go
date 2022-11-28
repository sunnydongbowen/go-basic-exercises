package const_t

import (
	"fmt"
	"math"
	"testing"
	"unsafe"
)

// @program:     go-basic-exercises
// @file:        int_test.go
// @author:      bowen
// @create:      2022-11-13 20:46
// @description:

func TestInt(t *testing.T) {
	var i1 int = 1
	var i2 int8 = 2
	var i3 int16 = 3
	var i4 int32 = 4
	var i5 int64 = 5
	var i6 uint64 = 6

	fmt.Printf("int    : %v\n", unsafe.Sizeof(i1))
	fmt.Printf("int8    : %v\n", unsafe.Sizeof(i2))
	fmt.Printf("int16    : %v\n", unsafe.Sizeof(i3))
	fmt.Printf("int32    : %v\n", unsafe.Sizeof(i4))
	fmt.Printf("int64    : %v\n", unsafe.Sizeof(i5))
	fmt.Printf("uint64    : %v\n", unsafe.Sizeof(i6))
}

func TestRange(t *testing.T) {
	fmt.Println("int8", math.MinInt8, "~", math.MaxInt8)
	fmt.Println("int16", math.MinInt16, "~", math.MaxInt16)
	fmt.Println("int32", math.MinInt32, "~", math.MaxInt32)
	fmt.Println("int64", math.MinInt64, "~", math.MaxInt64)
	fmt.Println("int", math.MinInt, "~", math.MaxInt)

}
